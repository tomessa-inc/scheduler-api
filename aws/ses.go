package awslocal

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
)

type SES struct {
	noReply string
	client  *ses.Client
}

func NewSES(cfg aws.Config, noReply string) SES {
	return SES{
		noReply: noReply,
		client:  ses.NewFromConfig(cfg),
	}
}

func (s SES) SendTemplatedEmail(ctx context.Context, args SESSendTemplatedEmailArgs) error {
	if len(args.TOs) == 0 {
		return errors.New("recipient email is missing")
	}

	if args.TemplateName == "" {
		return errors.New("template name is missing")
	}

	var data string

	if len(args.TemplateData) != 0 {
		val, err := json.Marshal(args.TemplateData)
		if err != nil {
			return fmt.Errorf("unable to marshal template data: %w", err)
		}
		data = string(val)
	}

	_, err := s.client.SendTemplatedEmail(ctx, &ses.SendTemplatedEmailInput{
		Source: aws.String(s.noReply),
		Destination: &types.Destination{
			ToAddresses: args.TOs,
			CcAddresses: args.CCs,
		},
		Template:     &args.TemplateName,
		TemplateData: &data,
	})
	if err != nil {
		return fmt.Errorf("unable to send email: %w", err)
	}

	return nil
}

func (s SES) SendEmail(ctx context.Context, args SESSendEmailArgs) error {
	if len(args.TOs) == 0 {
		return errors.New("recipient email is missing")
	}

	if args.HtmlBody == "" && args.TextBody == "" {
		return errors.New("body is missing")
	}

	var body types.Body
	if args.HtmlBody != "" {
		body.Html = &types.Content{
			Charset: aws.String("UTF-8"),
			Data:    aws.String(args.HtmlBody),
		}
	}
	if args.TextBody != "" {
		body.Text = &types.Content{
			Charset: aws.String("UTF-8"),
			Data:    aws.String(args.TextBody),
		}
	}

	_, err := s.client.SendEmail(ctx, &ses.SendEmailInput{
		Source: aws.String(s.noReply),
		Destination: &types.Destination{
			ToAddresses: args.TOs,
			CcAddresses: args.CCs,
		},
		Message: &types.Message{
			Subject: &types.Content{
				Charset: aws.String("UTF-8"),
				Data:    aws.String(args.Subject),
			},
			Body: &body,
		},
	})
	if err != nil {
		return fmt.Errorf("unable to send email: %w", err)
	}

	return nil
}
