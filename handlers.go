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
