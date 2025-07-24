# vkGetPhysicalDeviceSurfaceCapabilitiesKHR(3) Manual Page

## Name

vkGetPhysicalDeviceSurfaceCapabilitiesKHR - Query surface capabilities



## [](#_c_specification)C Specification

To query the basic capabilities of a surface, needed in order to create a swapchain, call:

```c++
// Provided by VK_KHR_surface
VkResult vkGetPhysicalDeviceSurfaceCapabilitiesKHR(
    VkPhysicalDevice                            physicalDevice,
    VkSurfaceKHR                                surface,
    VkSurfaceCapabilitiesKHR*                   pSurfaceCapabilities);
```

## [](#_parameters)Parameters

- `physicalDevice` is the physical device that will be associated with the swapchain to be created, as described for [vkCreateSwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateSwapchainKHR.html).
- `surface` is the surface that will be associated with the swapchain.
- `pSurfaceCapabilities` is a pointer to a [VkSurfaceCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCapabilitiesKHR.html) structure in which the capabilities are returned.

## [](#_description)Description

Valid Usage

- [](#VUID-vkGetPhysicalDeviceSurfaceCapabilitiesKHR-surface-06211)VUID-vkGetPhysicalDeviceSurfaceCapabilitiesKHR-surface-06211  
  `surface` **must** be supported by `physicalDevice`, as reported by [vkGetPhysicalDeviceSurfaceSupportKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSurfaceSupportKHR.html) or an equivalent platform-specific mechanism

Valid Usage (Implicit)

- [](#VUID-vkGetPhysicalDeviceSurfaceCapabilitiesKHR-physicalDevice-parameter)VUID-vkGetPhysicalDeviceSurfaceCapabilitiesKHR-physicalDevice-parameter  
  `physicalDevice` **must** be a valid [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html) handle
- [](#VUID-vkGetPhysicalDeviceSurfaceCapabilitiesKHR-surface-parameter)VUID-vkGetPhysicalDeviceSurfaceCapabilitiesKHR-surface-parameter  
  `surface` **must** be a valid [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceKHR.html) handle
- [](#VUID-vkGetPhysicalDeviceSurfaceCapabilitiesKHR-pSurfaceCapabilities-parameter)VUID-vkGetPhysicalDeviceSurfaceCapabilitiesKHR-pSurfaceCapabilities-parameter  
  `pSurfaceCapabilities` **must** be a valid pointer to a [VkSurfaceCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCapabilitiesKHR.html) structure
- [](#VUID-vkGetPhysicalDeviceSurfaceCapabilitiesKHR-commonparent)VUID-vkGetPhysicalDeviceSurfaceCapabilitiesKHR-commonparent  
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

[VK\_KHR\_surface](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_surface.html), [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html), [VkSurfaceCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCapabilitiesKHR.html), [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetPhysicalDeviceSurfaceCapabilitiesKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0