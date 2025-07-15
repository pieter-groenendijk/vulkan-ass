package main

import (
	"fmt"

	"github.com/go-gl/glfw/v3.3/glfw"
)

func initWindow() error {
	err := glfw.Init()
	if err != nil {
		return fmt.Errorf("failed to initialize GLFW: %w", err)
	}

	glfw.WindowHint(glfw.ClientAPI, glfw.NoAPI)
	window, err = glfw.CreateWindow(windowWidth, windowHeight, "Vulkan", nil, nil)
	if err != nil {
		return fmt.Errorf("failed to create window: %w", err)
	}

	return nil
}
