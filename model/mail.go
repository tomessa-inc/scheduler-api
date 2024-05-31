package model

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	awslocal "scheduler-api/aws"
	"scheduler-api/conversion"
	e "scheduler-api/entity"
)

const (
	// Replace sender@example.com with your "From" address.
	// This address must be verified with Amazon SES.
	Sender = "tomc@tomvisions.com"

	// Replace recipient@example.com with a "To" address. If your account
	// is still in the sandbox, this address must be verified.
	Recipient = "tcruicksh@gmail.com"

	// Specify a configuration set. To use a configuration
	// set, comment the next line and line 92.
	//ConfigurationSet = "ConfigSet"

	// The subject line for the email.
	Subject = "Amazon SES Test (AWS SDK for Go)"

	// The HTML body for the email.
	HtmlBodyUpcoming = "<h1>Week of %s</h1><p>The following ushers are using for the mass</p>%s" +
		"<p>If anyone is unavailable this week please click on unavailable Link <a href='https://aws.amazon.com/ses/'>Amazon SES</a> using the " +
		"<a href='https://aws.amazon.com/sdk-for-go/'>AWS SDK for Go</a>.</p>"

	HtmlBody = "<h1>Amazon SES Test Email (AWS SDK for Go)</h1><p>This email was sent with " +
		"<a href='https://aws.amazon.com/ses/'>Amazon SES</a> using the " +
		"<a href='https://aws.amazon.com/sdk-for-go/'>AWS SDK for Go</a>.</p>"

	//The email body for recipients with non-HTML email clients.
	TextBody = "This email was sent with Amazon SES using the AWS SDK for Go."

	// The character encoding for the email.
	CharSet = "UTF-8"
)

func PrepareEmail(users e.ScheduleUser, emailType string, token ...string) {
	ctx := context.Background()
	noreply := "tomc@tomvisions.com"
	to := "tcruicksh@gmail.com"
	var htmlData []byte

	// Initialise AWS config.
	cfg, err := awslocal.NewConfig(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	// Initialise SES client.
	ses := awslocal.NewSES(cfg, noreply)
	switch emailType {
	case "week-off":
		htmlData, err = ioutil.ReadFile("../data/week-off.html")
		fmt.Println(htmlData)
		if err != nil {

		}
		break

	case "week-on":
		htmlData, err = ioutil.ReadFile("../data/week-on.html")

		if err != nil {

		}
		break

	case "unavailable":
		htmlData, err = ioutil.ReadFile("../data/unavailable.html")

		if err != nil {

		}

		break
	case "available":
		htmlData, err = ioutil.ReadFile("../data/available.html")

		if err != nil {

		}
		break
	}

	templateData, err := conversion.ConvertArrayBytesToString(htmlData)

	// Send templated email from AWS.
	err = ses.SendTemplatedEmail(ctx, awslocal.SESSendTemplatedEmailArgs{
		TOs:          []string{to},
		TemplateName: "DefaultEmailTemplate",
		TemplateData: templateData,
	})

}
