# Project Setup
## Golang
Using a binding library to be able to call vulkan, which is C. This way I'm
not forced to write C.

## GLFW
My graphics need a place to be. I will be using [GLFW](https://www.glfw.org/) for this exact purpose. It provides
a library to create windows, contexts (vulkan) and surfaces, receiving input and events.

Since this is written is C, and I'm using Go, I will also need to use a binding for this one. I have chosen 
[this one](https://github.com/go-gl/glfw).

_The mainstream alternative is SDL2, which is a more complete package._

