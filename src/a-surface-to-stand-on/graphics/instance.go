package graphics

import (
	"fmt"

	vk "github.com/vulkan-go/vulkan"
)

func initInstance(extensions []string) (vk.Instance, error) {
	appInfo := &vk.ApplicationInfo{
		SType: vk.StructureTypeApplicationInfo,
		PApplicationName: "Vulkan Baby",
		ApplicationVersion: vk.MakeVersion(1, 0, 0),
		PEngineName: "No engine",
		EngineVersion: vk.MakeVersion(1, 0, 0),
		ApiVersion: vk.ApiVersion10,
	}

	numOfExtensions := uint32(len(extensions))
	createInfo := &vk.InstanceCreateInfo{
		SType: vk.StructureTypeInstanceCreateInfo,
		PApplicationInfo: appInfo,

		EnabledExtensionCount: numOfExtensions,
		PpEnabledExtensionNames: extensions,

		EnabledLayerCount: 0,
	}

	var instance vk.Instance
	result := vk.CreateInstance(createInfo, nil, &instance)
	if result != vk.Success {
		return nil, fmt.Errorf("failed to create instance, VkResult: %v", result)
	}

	return instance, nil
}
