package main

import (
	"log"

	"github.com/homebrew-arcade/rpi5-ebitengine-kiosk/pkg/kiosk"
)

func main() {
	err := kiosk.Main()
	if err != nil {
		log.Fatal(err)
	}
}
