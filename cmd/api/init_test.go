package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"sync"
	"time"

	"github.com/a-novel-kit/context"
	pgctx "github.com/a-novel-kit/context/pgbun"

	"github.com/a-novel/authentication/api/apiclient"
	"github.com/a-novel/authentication/api/codegen"
	"github.com/a-novel/authentication/config"
)

func getServerClient() (*codegen.Client, *apiclient.SecuritySource, error) {
	security := apiclient.NewSecuritySource()

	client, err := codegen.NewClient(fmt.Sprintf("http://127.0.0.1:%v/v1", config.API.Port), security)
	if err != nil {
		return nil, nil, fmt.Errorf("create client: %w", err)
	}

	start := time.Now()
	_, err = client.Ping(context.Background())

	for time.Since(start) < 16*time.Second && err != nil {
		_, err = client.Ping(context.Background())
	}

	if err != nil {
		return nil, nil, fmt.Errorf("ping server: %w", err)
	}

	return client, security, nil
}

// Create a separate database to run integration tests.
func init() {
	ctx, err := pgctx.NewContext(context.Background(), nil)
	if err != nil {
		panic(err)
	}

	db, err := pgctx.Context(ctx)
	if err != nil {
		panic(err)
	}

	_, err = db.NewRaw("CREATE DATABASE integrations_test OWNER test;").Exec(ctx)
	if err != nil {
		panic(err)
	}

	ogDSN := os.Getenv(pgctx.PostgresDSNEnv)

	err = os.Setenv(pgctx.PostgresDSNEnv, "postgres://test:test@localhost:5432/integrations_test?sslmode=disable")
	if err != nil {
		panic(err)
	}

	// Run keys rotation.
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	rotateKeysPath := path.Join(filepath.Dir(pwd), "..", "cmd", "rotatekeys", "main.go")

	out, err := exec.CommandContext(ctx, "go", "run", rotateKeysPath).CombinedOutput()
	if err != nil {
		panic(fmt.Sprintf("rotate keys: %v\n%v", err, string(out)))
	}

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		main()
	}()

	_, _, err = getServerClient()
	if err != nil {
		panic(err)
	}

	err = os.Setenv(pgctx.PostgresDSNEnv, ogDSN)
	if err != nil {
		panic(err)
	}
}
