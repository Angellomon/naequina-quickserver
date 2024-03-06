package main

import (
	"context"
	"fmt"
	"log"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type EmailData struct {
	contactEmail string
	message      string
	name         string
}

func SendEmail(ctx context.Context, data EmailData) error {
	sendgridAPIKey := ctx.Value(ContextKey("SendGrid")).(string)
	template := ctx.Value(ContextKey("TemplateContacto")).(string)

	m := mail.NewV3Mail()
	content := mail.NewContent("text/html", template)

	fromMail := mail.NewEmail("NAEQUINA", "naequina@gponutec.com")

	m.SetFrom(fromMail)
	m.AddContent(content)

	personalization := mail.NewPersonalization()

	tos := []*mail.Email{
		mail.NewEmail("Angel", "angelmtzdiaz@gmail.com"),
		// mail.NewEmail("rpereza", "rpereza@unam.mx"),
		// mail.NewEmail("rpereza", "rpereza@unam.mx"),
	}

	personalization.AddTos(tos...)

	personalization.SetSubstitution("%nombre%", data.name)
	personalization.SetSubstitution("%correo%", data.contactEmail)
	personalization.SetSubstitution("%mensaje%", data.message)

	personalization.Subject = "Solicitud de Contacto"

	m.AddPersonalizations(personalization)

	request := sendgrid.GetRequest(sendgridAPIKey, "/v3/mail/send", "https://api.sendgrid.com")

	request.Method = "POST"
	request.Body = mail.GetRequestBody(m)
	response, err := sendgrid.API(request)

	fmt.Println(response)

	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}

	return err
}
