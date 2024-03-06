package main

import (
	"os"
)

func ReadTemplateStr(tempalteName string) (string, error) {
	b, err := os.ReadFile("templates/" + tempalteName)

	if err != nil {
		return "", err
	}

	template := string(b)

	return template, nil
}
