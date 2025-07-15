package graphics

import (
	"fmt"

	"github.com/go-gl/glfw/v3.3/glfw"
	vk "github.com/vulkan-go/vulkan"
)

func New(window *glfw.Window) (vk.Instance, error) {
	vk.SetGetInstanceProcAddr(glfw.GetVulkanGetInstanceProcAddress())

	err := vk.Init()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize vulkan: %w", err)
	}

	instance, err := initInstance(window.GetRequiredInstanceExtensions())
	if err != nil {
		return nil, fmt.Errorf("failed to initialize instance: %w", err)
	}

	return instance, nil
}

func Clean(instance vk.Instance) {
	vk.DestroyInstance(instance, nil)
}


