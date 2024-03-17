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
	ctx = context.WithValue(ctx, ContextKey("SendGrid"), os.Getenv("SENDGRID_API_KEY"))
	ctx = context.WithValue(ctx, ContextKey("Recaptcha"), os.Getenv("CAPTCHA_SECRET_KEY"))
	ctx = context.WithValue(ctx, ContextKey("Mode"), os.Getenv("MODE"))
	ctx = context.WithValue(ctx, ContextKey("EventId"), os.Getenv("EVENTBRITE_EVENT_ID"))
	ctx = context.WithValue(ctx, ContextKey("EventbritePrivateToken"), os.Getenv("EVENTBRITE_PRIVATE_TOKEN"))

	fmt.Println("mode" + os.Getenv("MODE"))

	t, err := ReadTemplateStr("contact-template.html")

	if err != nil {
		log.Println(err)
		return
	}

	ctx = context.WithValue(ctx, ContextKey("TemplateContacto"), t)

	// mux := http.NewServeMux()

	http.HandleFunc("/send-email", MakeSendEmailHandler(ctx))
	// corsHandler, err := SetupCors(mux)

	http.HandleFunc("/available-tickets", MakeEventTicketsAvailableHandler(ctx))

	fmt.Printf("starting server at port 5005\n")

	if err := http.ListenAndServe(":5005", nil); err != nil {
		log.Fatal(err)
	}
}
