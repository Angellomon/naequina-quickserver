package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx := context.Background()
	ctxVal := context.WithValue(ctx, ContextKey("SendGrid"), os.Getenv("SENDGRID_API_KEY"))
	ctxVal = context.WithValue(ctxVal, ContextKey("Recaptcha"), os.Getenv("CAPTCHA_SECRET_KEY"))
	ctxVal = context.WithValue(ctxVal, ContextKey("Mode"), os.Getenv("MODE"))

	t, err := ReadTemplateStr("contact-template.html")

	if err != nil {
		log.Println(err)
		return
	}

	ctxVal = context.WithValue(ctxVal, ContextKey("TemplateContacto"), t)

	// mux := http.NewServeMux()

	http.HandleFunc("/send-email", MakeSendEmailHandler(ctxVal))
	// corsHandler, err := SetupCors(mux)

	if err != nil {
		log.Fatal(err)

	}

	fmt.Printf("starting server at port 5005\n")

	if err := http.ListenAndServe(":5005", nil); err != nil {
		log.Fatal(err)
	}
}
