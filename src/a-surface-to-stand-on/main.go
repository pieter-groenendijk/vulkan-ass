package main

import (
	"fmt"
	"log"

	"github.com/go-gl/glfw/v3.3/glfw"
)

const (
	windowWidth = 800
	windowHeight = 600
)

var window *glfw.Window

func main() {
	setup();
	loop();
	clean();
}

func setup() {
	err := initWindow()
	if err != nil {
		log.Fatal("failed to initialize window: ", err)
	}

	err = initVulkan()
	if err != nil {
		log.Fatal("failed to initialize vulkan: ", err)
	}

	devices, err := getPhysicalDevices()
	if err != nil {
		log.Fatal("failed to get physical devices: ", err)
	}
	fmt.Println(devices)
	if len(devices) == 0 {
		log.Fatal("no physical devices available")
	}
	fmt.Printf("%+v\n", getPhysicalDeviceProperties(devices[0]))
}

func loop() {
	for ; !window.ShouldClose(); {
		glfw.PollEvents()
	}
}

func clean() {
	destroyVulkan()
	window.Destroy()
	glfw.Terminate()
}
