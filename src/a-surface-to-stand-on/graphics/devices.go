package graphics

import (
	"fmt"

	"github.com/vulkan-go/vulkan"
	vk "github.com/vulkan-go/vulkan"
)

func GetPhysicalDevices(instance vulkan.Instance) ([]vk.PhysicalDevice, error) {
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

func GetPhysicalDeviceProperties(device vk.PhysicalDevice) (*vk.PhysicalDeviceProperties) {
	props := new(vk.PhysicalDeviceProperties)
	vk.GetPhysicalDeviceProperties(device, props)

	return props 
}
