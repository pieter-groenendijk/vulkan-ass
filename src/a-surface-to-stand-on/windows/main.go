package windows

import (
	"fmt"
	"github.com/go-gl/glfw/v3.3/glfw"
)

const (
	windowWidth = 800
	windowHeight = 600
)

func New() (*glfw.Window, error) {
	err := glfw.Init()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize GLFW: %w", err)
	}

	glfw.WindowHint(glfw.ClientAPI, glfw.NoAPI)
	window, err := glfw.CreateWindow(windowWidth, windowHeight, "Vulkan", nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create window: %w", err)
	}

	return window, nil
}

