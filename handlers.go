package main

import (
	"context"
	"fmt"
	"net/http"
)

func MakeSendEmailHandler(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/send-email" {
			http.Error(w, "404 not found", http.StatusNotFound)

			return
		}

		if r.Method != "POST" {
			http.Error(w, "method not supported", http.StatusNotFound)

			return
		}

		data := EmailData{
			contactEmail: "test@test.com",
			message:      "Test",
			name:         "Angel",
		}

		err := SendEmail(ctx, data)

		if err != nil {
			fmt.Printf("\n%v\n", err)
			fmt.Fprintf(w, "email error")

			return
		}

		fmt.Fprintf(w, "email sent")
	}

}
