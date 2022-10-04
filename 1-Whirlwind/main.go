package main

import (
	"github.com/jesseobrien/codemash-go-2023/logs"
)

func main() {
	log := logs.New()

	// Emoji support and utf in the language!
	log.Info("👋 Hello CodeMash attendees!")
	log.Info("こんにちは!")
	log.Info("你好!")
	log.Info("Bonjour!")
	log.Info("Hola!")
	log.Info("привіт!")
	log.Info("привет!")

	log.Info("https://go.dev/play/")
}
