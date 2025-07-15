package main

import (
	"example/vulkan-breathes/app"
	"log"
)


func main() {
	_app := app.New()

	defer _app.Clean()

	err := _app.Setup()
	if err != nil {
		log.Fatal(err)
	}

	_app.Loop()
}
