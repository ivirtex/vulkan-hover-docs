# vkGetPhysicalDeviceSurfaceCapabilities2EXT(3) Manual Page

## Name

vkGetPhysicalDeviceSurfaceCapabilities2EXT - Query surface capabilities



## [](#_c_specification)C Specification

To query the basic capabilities of a surface, needed in order to create a swapchain, call:

```c++
// Provided by VK_EXT_display_surface_counter
VkResult vkGetPhysicalDeviceSurfaceCapabilities2EXT(
    VkPhysicalDevice                            physicalDevice,
    VkSurfaceKHR                                surface,
    VkSurfaceCapabilities2EXT*                  pSurfaceCapabilities);
```

## [](#_parameters)Parameters

- `physicalDevice` is the physical device that will be associated with the swapchain to be created, as described for [vkCreateSwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateSwapchainKHR.html).
- `surface` is the surface that will be associated with the swapchain.
- `pSurfaceCapabilities` is a pointer to a [VkSurfaceCapabilities2EXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCapabilities2EXT.html) structure in which the capabilities are returned.

## [](#_description)Description

`vkGetPhysicalDeviceSurfaceCapabilities2EXT` behaves similarly to [vkGetPhysicalDeviceSurfaceCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSurfaceCapabilitiesKHR.html), with the ability to return extended information by adding extending structures to the `pNext` chain of its `pSurfaceCapabilities` parameter.

Valid Usage

- [](#VUID-vkGetPhysicalDeviceSurfaceCapabilities2EXT-surface-06211)VUID-vkGetPhysicalDeviceSurfaceCapabilities2EXT-surface-06211  
  `surface` **must** be supported by `physicalDevice`, as reported by [vkGetPhysicalDeviceSurfaceSupportKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSurfaceSupportKHR.html) or an equivalent platform-specific mechanism

Valid Usage (Implicit)

- [](#VUID-vkGetPhysicalDeviceSurfaceCapabilities2EXT-physicalDevice-parameter)VUID-vkGetPhysicalDeviceSurfaceCapabilities2EXT-physicalDevice-parameter  
  `physicalDevice` **must** be a valid [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html) handle
- [](#VUID-vkGetPhysicalDeviceSurfaceCapabilities2EXT-surface-parameter)VUID-vkGetPhysicalDeviceSurfaceCapabilities2EXT-surface-parameter  
  `surface` **must** be a valid [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceKHR.html) handle
- [](#VUID-vkGetPhysicalDeviceSurfaceCapabilities2EXT-pSurfaceCapabilities-parameter)VUID-vkGetPhysicalDeviceSurfaceCapabilities2EXT-pSurfaceCapabilities-parameter  
  `pSurfaceCapabilities` **must** be a valid pointer to a [VkSurfaceCapabilities2EXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCapabilities2EXT.html) structure
- [](#VUID-vkGetPhysicalDeviceSurfaceCapabilities2EXT-commonparent)VUID-vkGetPhysicalDeviceSurfaceCapabilities2EXT-commonparent  
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

[VK\_EXT\_display\_surface\_counter](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_display_surface_counter.html), [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html), [VkSurfaceCapabilities2EXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCapabilities2EXT.html), [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetPhysicalDeviceSurfaceCapabilities2EXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0