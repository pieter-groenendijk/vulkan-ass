package app

import (
	"example/vulkan-breathes/graphics"
	"example/vulkan-breathes/windows"
	"fmt"
	"log"

	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/vulkan-go/vulkan"
)

type App struct {
	vkInstance vulkan.Instance
	vkSurface vulkan.Surface
	window *glfw.Window
}

func New() *App {
	return &App{}
}

func (app *App) Setup() error {
	window, err := windows.New()
	if err != nil {
		log.Fatal("failed to initialize window: ", err)
	}

	vkInstance, err := graphics.New(window)
	if err != nil {
		log.Fatal("failed to initialize vulkan: ", err)
	}

	devices, err := graphics.GetPhysicalDevices(vkInstance)
	if err != nil {
		log.Fatal("failed to get physical devices: ", err)
	}
	fmt.Println(devices)
	if len(devices) == 0 {
		log.Fatal("no physical devices available")
	}
	fmt.Printf("%+v\n", graphics.GetPhysicalDeviceProperties(devices[0]))

	app.window = window
	app.vkInstance = vkInstance

	return nil
}

func (app *App) Loop() {
	for ; !app.window.ShouldClose(); {
		glfw.PollEvents()
	}
}

func (app *App) Clean() {
	graphics.Clean(app.vkInstance)
	app.window.Destroy()
	glfw.Terminate()
}
