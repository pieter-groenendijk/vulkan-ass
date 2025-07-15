package windows

import (
	"fmt"

	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/vulkan-go/vulkan"
)

func CreateSurface(instance vulkan.Instance, window *glfw.Window) (vulkan.Surface, error) {
	surface, err := window.CreateWindowSurface(instance, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create window surface: %w", err)
	}

	return vulkan.SurfaceFromPointer(surface), nil
}
