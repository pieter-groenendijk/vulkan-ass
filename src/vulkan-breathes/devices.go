package main

import (
	"fmt"
	vk "github.com/vulkan-go/vulkan"
)

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

func getPhysicalDeviceProperties(device vk.PhysicalDevice) (*vk.PhysicalDeviceProperties) {
	props := new(vk.PhysicalDeviceProperties)
	vk.GetPhysicalDeviceProperties(device, props)

	return props 
}
