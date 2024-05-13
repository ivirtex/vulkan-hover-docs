# vkGetPhysicalDeviceSurfaceSupportKHR(3) Manual Page

## Name

vkGetPhysicalDeviceSurfaceSupportKHR - Query if presentation is
supported



## <a href="#_c_specification" class="anchor"></a>C Specification

To determine whether a queue family of a physical device supports
presentation to a given surface, call:

``` c
// Provided by VK_KHR_surface
VkResult vkGetPhysicalDeviceSurfaceSupportKHR(
    VkPhysicalDevice                            physicalDevice,
    uint32_t                                    queueFamilyIndex,
    VkSurfaceKHR                                surface,
    VkBool32*                                   pSupported);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `physicalDevice` is the physical device.

- `queueFamilyIndex` is the queue family.

- `surface` is the surface.

- `pSupported` is a pointer to a [VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), which is set
  to `VK_TRUE` to indicate support, and `VK_FALSE` otherwise.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a
  href="#VUID-vkGetPhysicalDeviceSurfaceSupportKHR-queueFamilyIndex-01269"
  id="VUID-vkGetPhysicalDeviceSurfaceSupportKHR-queueFamilyIndex-01269"></a>
  VUID-vkGetPhysicalDeviceSurfaceSupportKHR-queueFamilyIndex-01269  
  `queueFamilyIndex` **must** be less than `pQueueFamilyPropertyCount`
  returned by `vkGetPhysicalDeviceQueueFamilyProperties` for the given
  `physicalDevice`

Valid Usage (Implicit)

- <a
  href="#VUID-vkGetPhysicalDeviceSurfaceSupportKHR-physicalDevice-parameter"
  id="VUID-vkGetPhysicalDeviceSurfaceSupportKHR-physicalDevice-parameter"></a>
  VUID-vkGetPhysicalDeviceSurfaceSupportKHR-physicalDevice-parameter  
  `physicalDevice` **must** be a valid
  [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html) handle

- <a href="#VUID-vkGetPhysicalDeviceSurfaceSupportKHR-surface-parameter"
  id="VUID-vkGetPhysicalDeviceSurfaceSupportKHR-surface-parameter"></a>
  VUID-vkGetPhysicalDeviceSurfaceSupportKHR-surface-parameter  
  `surface` **must** be a valid [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceKHR.html) handle

- <a
  href="#VUID-vkGetPhysicalDeviceSurfaceSupportKHR-pSupported-parameter"
  id="VUID-vkGetPhysicalDeviceSurfaceSupportKHR-pSupported-parameter"></a>
  VUID-vkGetPhysicalDeviceSurfaceSupportKHR-pSupported-parameter  
  `pSupported` **must** be a valid pointer to a
  [VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html) value

- <a href="#VUID-vkGetPhysicalDeviceSurfaceSupportKHR-commonparent"
  id="VUID-vkGetPhysicalDeviceSurfaceSupportKHR-commonparent"></a>
  VUID-vkGetPhysicalDeviceSurfaceSupportKHR-commonparent  
  Both of `physicalDevice`, and `surface` **must** have been created,
  allocated, or retrieved from the same [VkInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstance.html)

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

- `VK_ERROR_SURFACE_LOST_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_surface](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_surface.html), [VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html),
[VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html),
[VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetPhysicalDeviceSurfaceSupportKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
