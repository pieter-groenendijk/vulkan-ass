package main

import (
	"fmt"
	"log"

	"github.com/go-gl/glfw/v3.3/glfw"
	vk "github.com/vulkan-go/vulkan"
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
	err := initWindow();
	if err != nil {
		log.Fatal("failed to initialize window:", err)
	}

	vk.SetGetInstanceProcAddr(glfw.GetVulkanGetInstanceProcAddress())

	err = vk.Init()
	if err != nil {
		log.Fatal("failed to initialize vulkan:", err)
	}

	fmt.Println("Vulkan initialized")
}

func initWindow() error {
	err := glfw.Init()
	if err != nil {
		return fmt.Errorf("failed to initialize GLFW: %w", err)
	}

	glfw.WindowHint(glfw.ClientAPI, glfw.NoAPI)
	glfw.WindowHint(glfw.Resizable, glfw.False)
	window, err = glfw.CreateWindow(windowWidth, windowHeight, "Vulkan", nil, nil)
	if err != nil {
		return fmt.Errorf("failed to create window: %w", err)
	}

	return nil
}

func loop() {
	for ; window.ShouldClose(); {
		glfw.PollEvents()
	}
}

func clean() {

}
