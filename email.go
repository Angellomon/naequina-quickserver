package main

import (
	"context"
	"fmt"
	"log"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type EmailData struct {
	ContactEmail   string `json:"correo"`
	Message        string `json:"mensaje"`
	Name           string `json:"nombre"`
	RecaptchaToken string `json:"recaptchaToken"`
}

func SendEmail(ctx context.Context, data EmailData) error {
	sendgridAPIKey := ctx.Value(ContextKey("SendGrid")).(string)
	template := ctx.Value(ContextKey("TemplateContacto")).(string)
	mode := ctx.Value(ContextKey("Mode")).(string)

	m := mail.NewV3Mail()
	content := mail.NewContent("text/html", template)

	fromMail := mail.NewEmail("NAEQUINA", "naequina@gponutec.com")

	m.SetFrom(fromMail)
	m.AddContent(content)

	personalization := mail.NewPersonalization()

	var tos []*mail.Email

	if mode == "DEV" {
		tos = []*mail.Email{
			mail.NewEmail("Angel", "angelmtzdiaz@gmail.com"),
		}
	} else {
		tos = []*mail.Email{
			mail.NewEmail("rpereza", "rpereza@unam.mx"),
			mail.NewEmail("evelazquez", "evelazquez@gponutec.com"),
		}
	}

	personalization.AddTos(tos...)

	personalization.SetSubstitution("%nombre%", data.Name)
	personalization.SetSubstitution("%correo%", data.ContactEmail)
	personalization.SetSubstitution("%mensaje%", data.Message)

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
