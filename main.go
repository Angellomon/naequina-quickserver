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

	t, err := ReadTemplateStr("contact-template.html")

	if err != nil {
		log.Println(err)
		return
	}

	ctxVal2 := context.WithValue(ctxVal, ContextKey("TemplateContacto"), t)

	http.HandleFunc("/send-email", MakeSendEmailHandler(ctxVal2))

	fmt.Printf("starting server at port 5005")
	if err := http.ListenAndServe(":5005", nil); err != nil {
		log.Fatal(err)
	}
}
