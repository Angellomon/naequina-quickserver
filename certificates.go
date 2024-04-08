package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

type CertificatePDFError struct {
	Message string
}

func (e *CertificatePDFError) Error() string {
	return e.Message
}

type CertificatePDFData struct {
	Email          string `json:"email"`
	RecaptchaToken string `json:"recaptchaToken"`
}

func GetCertificatePDF(ctx context.Context, email string, w io.Writer) error {
	umbrellaEventId := ctx.Value(ContextKey("UmbrellaEventId")).(string)

	umbrellaEventUrl := fmt.Sprintf("https://constancias.umbrellaservices.angellos.net/events/pdf/%v?correo=%v", umbrellaEventId, email)

	res, err := http.Get(umbrellaEventUrl)

	if err != nil {
		fmt.Println("Error making the request (certificate pdf)")
		return &CertificatePDFError{Message: "Error making the request"}
	}

	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Println("Error reading the body (certificate pdf)")
		return &CertificatePDFError{Message: "Error reading the body"}
	}

	w.Write(b)

	return nil
}
