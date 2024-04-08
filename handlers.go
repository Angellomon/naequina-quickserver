package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func MakeSendEmailHandler(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		valid := ValidateHost(r.Host)

		if !valid {
			http.Error(w, "unauthorized host", http.StatusUnauthorized)
		}

		if r.Method != "POST" {
			http.Error(w, "method not supported", http.StatusNotFound)

			return
		}

		var data EmailData

		err := json.NewDecoder(r.Body).Decode(&data)

		defer r.Body.Close()

		if err != nil {
			w.WriteHeader(400)
			fmt.Fprintf(w, "Decode error")
			return
		}

		valid = ValidateRecaptchaToken(ctx, data.RecaptchaToken)

		if !valid {
			http.Error(w, "recaptcha fail", http.StatusBadRequest)

			return
		}

		err = SendEmail(ctx, data)

		if err != nil {
			fmt.Printf("\n%v\n", err)
			fmt.Fprintf(w, "email error")

			return
		}

		fmt.Fprintf(w, "email sent")
	}

}

type AvailableTicketsResponse struct {
	AvailableTickets int `json:"availableTickets"`
}

func MakeEventTicketsAvailableHandler(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "method not supported", http.StatusNotFound)

			return
		}

		availableTickets, err := GetAvailableTickets(ctx)

		if err != nil {
			http.Error(w, "error making the request to eventbrite", http.StatusServiceUnavailable)

			fmt.Println(err)

			return
		}

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(AvailableTicketsResponse{
			AvailableTickets: availableTickets,
		})

	}
}

func MakeGetCertificatePDFHandler(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "method not supported", http.StatusNotFound)

			return
		}

		var data CertificatePDFData

		err := json.NewDecoder(r.Body).Decode(&data)

		if err != nil {
			w.WriteHeader(400)
			fmt.Fprintf(w, "Decode error")
			return
		}

		valid := ValidateRecaptchaToken(ctx, data.RecaptchaToken)

		if !valid {
			http.Error(w, "recaptcha fail", http.StatusBadRequest)

			return
		}

		w.Header().Set("Content-Disposition", "attachment; filename=constancia.pdf")

		err = GetCertificatePDF(ctx, data.Email, w)

		if err != nil {
			http.Error(w, "error generating the pdf certificate", http.StatusServiceUnavailable)
		}
	}
}
