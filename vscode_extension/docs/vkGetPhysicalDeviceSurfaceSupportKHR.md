# vkGetPhysicalDeviceSurfaceSupportKHR(3) Manual Page

## Name

vkGetPhysicalDeviceSurfaceSupportKHR - Query if presentation is supported



## [](#_c_specification)C Specification

To determine whether a queue family of a physical device supports presentation to a given surface, call:

```c++
// Provided by VK_KHR_surface
VkResult vkGetPhysicalDeviceSurfaceSupportKHR(
    VkPhysicalDevice                            physicalDevice,
    uint32_t                                    queueFamilyIndex,
    VkSurfaceKHR                                surface,
    VkBool32*                                   pSupported);
```

## [](#_parameters)Parameters

- `physicalDevice` is the physical device.
- `queueFamilyIndex` is the queue family.
- `surface` is the surface.
- `pSupported` is a pointer to a [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html). `VK_TRUE` indicates support, and `VK_FALSE` indicates no support.

## [](#_description)Description

Valid Usage

- [](#VUID-vkGetPhysicalDeviceSurfaceSupportKHR-queueFamilyIndex-01269)VUID-vkGetPhysicalDeviceSurfaceSupportKHR-queueFamilyIndex-01269  
  `queueFamilyIndex` **must** be less than `pQueueFamilyPropertyCount` returned by `vkGetPhysicalDeviceQueueFamilyProperties` for the given `physicalDevice`

Valid Usage (Implicit)

- [](#VUID-vkGetPhysicalDeviceSurfaceSupportKHR-physicalDevice-parameter)VUID-vkGetPhysicalDeviceSurfaceSupportKHR-physicalDevice-parameter  
  `physicalDevice` **must** be a valid [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html) handle
- [](#VUID-vkGetPhysicalDeviceSurfaceSupportKHR-surface-parameter)VUID-vkGetPhysicalDeviceSurfaceSupportKHR-surface-parameter  
  `surface` **must** be a valid [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceKHR.html) handle
- [](#VUID-vkGetPhysicalDeviceSurfaceSupportKHR-pSupported-parameter)VUID-vkGetPhysicalDeviceSurfaceSupportKHR-pSupported-parameter  
  `pSupported` **must** be a valid pointer to a [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html) value
- [](#VUID-vkGetPhysicalDeviceSurfaceSupportKHR-commonparent)VUID-vkGetPhysicalDeviceSurfaceSupportKHR-commonparent  
  Both of `physicalDevice`, and `surface` **must** have been created, allocated, or retrieved from the same [VkInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstance.html)

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_SURFACE_LOST_KHR`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_KHR\_surface](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_surface.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html), [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetPhysicalDeviceSurfaceSupportKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0