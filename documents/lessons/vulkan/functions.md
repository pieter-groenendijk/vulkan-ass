# SetGetInstanceProcAddr(getProcAddr unsafe.Pointer)
Sets the `GetInstanceProcAddr` function pointer. This pointer to any
function one might need. E.g. `GetInstanceProcAddr("vkCreateInstance")`. It resolves symbols to the functions.

"GLFW" or "SDL" will help you find and load the correct pointer. For "GLFW", one calls 
`GetVulkanGetInstanceProcAddress`.

