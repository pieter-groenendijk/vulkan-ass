package main

import (
	"fmt"

	"github.com/go-gl/glfw/v3.3/glfw"
	vk "github.com/vulkan-go/vulkan"
)

var instance vk.Instance

func initVulkan() error {
	vk.SetGetInstanceProcAddr(glfw.GetVulkanGetInstanceProcAddress())

	err := vk.Init()
	if err != nil {
		return fmt.Errorf("failed to initialize vulkan: %w", err)
	}

	err = initInstance()
	if err != nil {
		return fmt.Errorf("failed to initialize instance: %w", err)
	}

	return nil
}

func destroyVulkan() {
	vk.DestroyInstance(instance, nil)
}

func initInstance() error {
	appInfo := &vk.ApplicationInfo{
		SType: vk.StructureTypeApplicationInfo,
		PApplicationName: "Vulkan Baby",
		ApplicationVersion: vk.MakeVersion(1, 0, 0),
		PEngineName: "No engine",
		EngineVersion: vk.MakeVersion(1, 0, 0),
		ApiVersion: vk.ApiVersion10,
	}

	extensions := window.GetRequiredInstanceExtensions()
	if extensions == nil {
		return fmt.Errorf("Vulkan is available but no set of extensions allowing window surface creation was found")
	}
	var numOfExtensions uint32 = uint32(len(extensions))

	createInfo := &vk.InstanceCreateInfo{
		SType: vk.StructureTypeInstanceCreateInfo,
		PApplicationInfo: appInfo,

		EnabledExtensionCount: numOfExtensions,
		PpEnabledExtensionNames: extensions,

		EnabledLayerCount: 0,
	}

	result := vk.CreateInstance(createInfo, nil, &instance)
	if result != vk.Success {
		return fmt.Errorf("failed to create instance, VkResult: %v", result)
	}

	return nil
}

func getPhysicalDevices() ([]vk.PhysicalDevice, error) {
	var amount uint32

	// if `pPhysicalDevices == nil` then function queries the available amount, and replaces `pPhysicalDeviceCount`
	result := vk.EnumeratePhysicalDevices(instance, &amount, nil)
	if result != vk.Success {
		return nil, fmt.Errorf("failed to query physical devices amount, VkResult: %v", result)
	}

	devices := make([]vk.PhysicalDevice, amount)

	// if `pPhysicalDevices != nil` then function queries the physical devices, and inserts them into `pPhysicalDevices`
	result = vk.EnumeratePhysicalDevices(instance, &amount, devices)
	if result != vk.Success {
		return nil, fmt.Errorf("failed to query physical devices, VkResult : %v", result)
	}

	return devices, nil
}
