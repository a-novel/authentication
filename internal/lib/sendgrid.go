package lib

import (
	"net/http"

	"github.com/rs/zerolog"
	"github.com/samber/lo"
	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"

	"github.com/a-novel-kit/context"

	"github.com/a-novel/authentication/config"
)

func SendMail(ctx context.Context, message *mail.SGMailV3) {
	logger := zerolog.Ctx(ctx)
	client := sendgrid.NewSendClient(config.Sendgrid.APIKey)

	message.MailSettings = &mail.MailSettings{
		BypassListManagement: mail.NewSetting(true),
		SandboxMode:          mail.NewSetting(config.Sendgrid.Sandbox),
	}

	var (
		recipents           []string
		bccs                []string
		ccs                 []string
		dynamicTemplateData = make(map[string]interface{})

		response *rest.Response
		err      error
	)

	for _, personnalization := range message.Personalizations {
		for _, to := range personnalization.To {
			recipents = append(recipents, to.Address)
		}

		for _, bcc := range personnalization.BCC {
			bccs = append(bccs, bcc.Address)
		}

		for _, cc := range personnalization.CC {
			ccs = append(ccs, cc.Address)
		}

		for key, value := range personnalization.DynamicTemplateData {
			dynamicTemplateData[key] = value
		}
	}

	// If sandbox mode is enabled, sendgrid will just parse the template, without actually sending the email.
	response, err = client.SendWithContext(ctx, message)

	if err != nil {
		logger.Error().
			Err(err).
			Str("templateID", message.TemplateID).
			Str("from", message.From.Address).
			Strs("to", recipents).
			Strs("bccs", bccs).
			Strs("ccs", ccs).
			Msg("send transactional email")
	} else {
		event := lo.Ternary(
			response.StatusCode >= http.StatusBadRequest, logger.Error(), logger.Info(), //nolint:zerologlint
		)

		// Only log the mail body if the sandbox mode is activated.
		if config.Sendgrid.Sandbox {
			event = event.
				Str("mail", string(mail.GetRequestBody(message))).
				Interface("dynamicTemplateData", dynamicTemplateData)
		}

		event.
			Str("templateID", message.TemplateID).
			Str("from", message.From.Address).
			Strs("to", recipents).
			Strs("bccs", bccs).
			Strs("ccs", ccs).
			Str("response", response.Body).
			Int("status", response.StatusCode).
			Msg("send transactional email")
	}
}
