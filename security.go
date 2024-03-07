package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// medida provisional de seguridad, sólo asegura que la conexión provenga del proxy
func ValidateHost(reqHost string) bool {
	mode := os.Getenv("MODE")

	var allowedHosts []string

	if mode == "DEV" {
		allowedHosts = []string{
			"localhost:5005",
			"naequina.com",
		}
	} else if mode == "PROD" {
		allowedHosts = []string{
			"naequina.com",
		}
	}

	valid := false

	for _, host := range allowedHosts {
		if reqHost == host {
			valid = true
			break
		}
	}

	return valid
}

func ValidateRecaptchaToken(ctx context.Context, token string) bool {
	var captchaVerifyResponse interface{}

	key := ctx.Value(ContextKey("Recaptcha")).(string)

	captchaPayloadRequest := url.Values{}
	captchaPayloadRequest.Set("secret", key)
	captchaPayloadRequest.Set("response", token)

	verifyCaptchaRequest, _ := http.NewRequest("POST", "https://www.google.com/recaptcha/api/siteverify", strings.NewReader(captchaPayloadRequest.Encode()))
	verifyCaptchaRequest.Header.Add("content-type", "application/x-www-form-urlencoded")
	verifyCaptchaRequest.Header.Add("cache-control", "no-cache")

	verifyCaptchaResponse, _ := http.DefaultClient.Do(verifyCaptchaRequest)

	decoder := json.NewDecoder(verifyCaptchaResponse.Body)
	decoderErr := decoder.Decode(&captchaVerifyResponse)

	defer verifyCaptchaResponse.Body.Close()

	fmt.Printf("err == nil: %v\n", decoderErr == nil)

	fmt.Println(decoderErr)

	return decoderErr != nil
}
