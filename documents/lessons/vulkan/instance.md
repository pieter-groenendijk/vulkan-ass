# Instance
Vulkan itself is stateless; there is no global state. Yet, an application has some state.. Vulkan provides `vkInstance` 
likely for this purpose. It allows for per-application state. The instance is what you'll be 
interacting with through the object model.

From a more practical perspective. It allows developers to define how they want to vulkan to behave; the configuration. Vulkan 
knows how it should behave, specifically according to the configuration of the instance used.
