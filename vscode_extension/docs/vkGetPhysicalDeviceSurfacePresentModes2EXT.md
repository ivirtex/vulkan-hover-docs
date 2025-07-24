# vkGetPhysicalDeviceSurfacePresentModes2EXT(3) Manual Page

## Name

vkGetPhysicalDeviceSurfacePresentModes2EXT - Query supported presentation modes



## [](#_c_specification)C Specification

Alternatively, to query the supported presentation modes for a surface combined with select other fixed swapchain creation parameters, call:

```c++
// Provided by VK_EXT_full_screen_exclusive
VkResult vkGetPhysicalDeviceSurfacePresentModes2EXT(
    VkPhysicalDevice                            physicalDevice,
    const VkPhysicalDeviceSurfaceInfo2KHR*      pSurfaceInfo,
    uint32_t*                                   pPresentModeCount,
    VkPresentModeKHR*                           pPresentModes);
```

## [](#_parameters)Parameters

- `physicalDevice` is the physical device that will be associated with the swapchain to be created, as described for [vkCreateSwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateSwapchainKHR.html).
- `pSurfaceInfo` is a pointer to a [VkPhysicalDeviceSurfaceInfo2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceSurfaceInfo2KHR.html) structure describing the surface and other fixed parameters that would be consumed by [vkCreateSwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateSwapchainKHR.html).
- `pPresentModeCount` is a pointer to an integer related to the number of presentation modes available or queried, as described below.
- `pPresentModes` is either `NULL` or a pointer to an array of [VkPresentModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentModeKHR.html) values, indicating the supported presentation modes.

## [](#_description)Description

`vkGetPhysicalDeviceSurfacePresentModes2EXT` behaves similarly to [vkGetPhysicalDeviceSurfacePresentModesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSurfacePresentModesKHR.html), with the ability to specify extended inputs via chained input structures.

Valid Usage

- [](#VUID-vkGetPhysicalDeviceSurfacePresentModes2EXT-pSurfaceInfo-06521)VUID-vkGetPhysicalDeviceSurfacePresentModes2EXT-pSurfaceInfo-06521  
  If the `VK_GOOGLE_surfaceless_query` extension is not enabled, `pSurfaceInfo->surface` **must** be a valid [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceKHR.html) handle
- [](#VUID-vkGetPhysicalDeviceSurfacePresentModes2EXT-pSurfaceInfo-06522)VUID-vkGetPhysicalDeviceSurfacePresentModes2EXT-pSurfaceInfo-06522  
  If `pSurfaceInfo->surface` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `pSurfaceInfo->surface` **must** be supported by `physicalDevice`, as reported by [vkGetPhysicalDeviceSurfaceSupportKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSurfaceSupportKHR.html) or an equivalent platform-specific mechanism

Valid Usage (Implicit)

- [](#VUID-vkGetPhysicalDeviceSurfacePresentModes2EXT-physicalDevice-parameter)VUID-vkGetPhysicalDeviceSurfacePresentModes2EXT-physicalDevice-parameter  
  `physicalDevice` **must** be a valid [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html) handle
- [](#VUID-vkGetPhysicalDeviceSurfacePresentModes2EXT-pSurfaceInfo-parameter)VUID-vkGetPhysicalDeviceSurfacePresentModes2EXT-pSurfaceInfo-parameter  
  `pSurfaceInfo` **must** be a valid pointer to a valid [VkPhysicalDeviceSurfaceInfo2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceSurfaceInfo2KHR.html) structure
- [](#VUID-vkGetPhysicalDeviceSurfacePresentModes2EXT-pPresentModeCount-parameter)VUID-vkGetPhysicalDeviceSurfacePresentModes2EXT-pPresentModeCount-parameter  
  `pPresentModeCount` **must** be a valid pointer to a `uint32_t` value
- [](#VUID-vkGetPhysicalDeviceSurfacePresentModes2EXT-pPresentModes-parameter)VUID-vkGetPhysicalDeviceSurfacePresentModes2EXT-pPresentModes-parameter  
  If the value referenced by `pPresentModeCount` is not `0`, and `pPresentModes` is not `NULL`, `pPresentModes` **must** be a valid pointer to an array of `pPresentModeCount` [VkPresentModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentModeKHR.html) values

Return Codes

On success, this command returns

- `VK_INCOMPLETE`
- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_SURFACE_LOST_KHR`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_EXT\_full\_screen\_exclusive](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_full_screen_exclusive.html), [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html), [VkPhysicalDeviceSurfaceInfo2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceSurfaceInfo2KHR.html), [VkPresentModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentModeKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetPhysicalDeviceSurfacePresentModes2EXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0