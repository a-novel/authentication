package dao

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/uptrace/bun/driver/pgdriver"

	"github.com/a-novel-kit/context"
	pgctx "github.com/a-novel-kit/context/pgbun"

	"github.com/a-novel/authentication/models"
)

// InsertShortCodeData is the input used to perform the InsertShortCodeRepository.InsertShortCode action.
type InsertShortCodeData struct {
	// ID of the short code. It must be unique (random).
	ID uuid.UUID

	// The encrypted code. A clear version of this code is sent to the target.
	Code string
	// Information about the resource the short code grants access to.
	Usage models.ShortCodeUsage
	// The target that is allowed to access the resource. Only this target can retrieve the short code.
	Target string
	// Data used for the targeted resource. It can contain any information required to perform a specific action.
	Data []byte

	// Time at which the short code was created.
	Now time.Time
	// Expiration of the short code. Each short code is REQUIRED to expire past a certain time. Once the expiration date
	// is reached, the short code can no longer be used or retrieved.
	ExpiresAt time.Time

	// Whe true, automatically expires any existing short code with the same target and usage.
	// Otherwise, the presence of a duplicate will trigger an error.
	//
	// Deleted duplicates will have their deletion comment set to DeleteCommentOverrideWithNewerKey.
	Override bool
}

// InsertShortCodeRepository is the repository used to perform the InsertShortCodeRepository.InsertShortCode action.
//
// You may create one using the NewInsertShortCodeRepository function.
type InsertShortCodeRepository struct{}

// InsertShortCode inserts a new short code in the database.
//
// The clear value of the short code MUST not be saved with this repository, but instead sent to the target using a
// secure channel. The encrypted version of the short code is stored in the database.
//
// A short code is uniquely identified by its target and usage. If another code with the same target and usage is
// found, the insertion will fail. You can override this behavior, and use the InsertShortCodeData.Override field to
// automatically expire any conflicting short code.
func (repository *InsertShortCodeRepository) InsertShortCode(
	ctx context.Context, data InsertShortCodeData,
) (*ShortCodeEntity, error) {
	// Retrieve a connection to postgres from the context.
	db, err := pgctx.Context(ctx)
	if err != nil {
		return nil, fmt.Errorf("(InsertShortCodeRepository.InsertShortCode) get postgres client: %w", err)
	}

	// Since we may be performing 2 operations depending on the parameters, create a new transaction to prevent any
	// data corruption.
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("(InsertShortCodeRepository.InsertShortCode) start transaction: %w", err)
	}
	// Make sure to roll back the transaction if an error occurs.
	defer func() { _ = tx.Rollback() }()

	if data.Override {
		oldEntity := &ShortCodeEntity{
			// So as we switch between Go/Pg timestamps, and transactions / views / triggers,
			// there is a mismatch when comparing dates which can sometimes lead Postgres to
			// believe this (expired) row is still part of the active short codes view.
			//
			// To make sure any conflicting short code is considered expired by the time
			// we perform the insertion, we cheat and make the deletion date one second
			// older. This is usually enough to prevent any issue.
			DeletedAt:      lo.ToPtr(data.Now.Add(-time.Second)),
			DeletedComment: lo.ToPtr(DeleteCommentOverrideWithNewerKey),
		}

		// Discard any conflicting short codes before starting.
		_, err = tx.NewUpdate().
			Model(oldEntity).
			Column("deleted_at", "deleted_comment").
			Where("target = ?", data.Target).
			Where("usage = ?", data.Usage).
			Where("COALESCE(deleted_at, expires_at) >= clock_timestamp()").
			Returning("*").
			Exec(ctx)
		if err != nil {
			return nil, fmt.Errorf("(InsertShortCodeRepository.InsertShortCode) discard old short codes: %w", err)
		}
	}

	// Insert the new short code.
	newEntity := &ShortCodeEntity{
		ID:        data.ID,
		Code:      data.Code,
		Usage:     data.Usage,
		Target:    data.Target,
		Data:      data.Data,
		CreatedAt: data.Now,
		ExpiresAt: data.ExpiresAt,
	}

	// Execute query.
	_, err = tx.NewInsert().Model(newEntity).Returning("*").Exec(ctx)
	if err != nil {
		var pgErr pgdriver.Error
		if errors.As(err, &pgErr) && pgErr.Field('C') == "23505" {
			return nil, errors.Join(
				fmt.Errorf("(InsertShortCodeRepository.InsertShortCode): %w", err),
				ErrShortCodeAlreadyExists,
			)
		}

		return nil, fmt.Errorf("(InsertShortCodeRepository.InsertShortCode) insert short code: %w", err)
	}

	// Commit the transaction.
	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("(InsertShortCodeRepository.InsertShortCode) commit transaction: %w", err)
	}

	return newEntity, nil
}

func NewInsertShortCodeRepository() *InsertShortCodeRepository {
	return &InsertShortCodeRepository{}
}
