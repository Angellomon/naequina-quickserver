package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	// "net/url"
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

	umbrellaEventStr := fmt.Sprintf("https://constancias.umbrellaservices.angellos.net/pdf/%v?correo=%v", umbrellaEventId, email)

	// umbrellaEventUrl := url.QueryEscape(umbrellaEventStr)

	res, err := http.Get(umbrellaEventStr)

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

	size, _ := w.Write(b)

	fmt.Println("wrote", size, "bytes")
	fmt.Println("content length", res.ContentLength)

	return nil
}
