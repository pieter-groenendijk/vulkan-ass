package main

import (
	"fmt"
	"log"

	"github.com/go-gl/glfw/v3.3/glfw"
	vk "github.com/vulkan-go/vulkan"
)

func main() {
	err := glfw.Init()
	if err != nil {
		log.Fatal("failed to initialize GLFW:", err)
	}
	defer glfw.Terminate()


	// Further research required! 
	// This is our glfw passing pointers for the "proc", process?
	// So that vulkan knows where to bla bla to (define bla bla).
	vk.SetGetInstanceProcAddr(glfw.GetVulkanGetInstanceProcAddress())

	err = vk.Init()
	if err != nil {
		log.Fatal("failed to initialize vulkan", err)
	}

	fmt.Println("Vulkan initialized")
}

