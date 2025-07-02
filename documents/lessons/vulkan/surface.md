# Surface
Vulkan itself is completely platform agnostic, which is why we need to
use the Window System Interface (WSI) extension to interact with the 
window manager. A surface is an abstraction over windows to render to. 
Windowing libraries like "GLFW" and "SDL" have helper functions to handle
with these platform-specific details.
