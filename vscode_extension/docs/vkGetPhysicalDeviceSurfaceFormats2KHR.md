# vkGetPhysicalDeviceSurfaceFormats2KHR(3) Manual Page

## Name

vkGetPhysicalDeviceSurfaceFormats2KHR - Query color formats supported by surface



## [](#_c_specification)C Specification

To query the supported swapchain format tuples for a surface, call:

```c++
// Provided by VK_KHR_get_surface_capabilities2
VkResult vkGetPhysicalDeviceSurfaceFormats2KHR(
    VkPhysicalDevice                            physicalDevice,
    const VkPhysicalDeviceSurfaceInfo2KHR*      pSurfaceInfo,
    uint32_t*                                   pSurfaceFormatCount,
    VkSurfaceFormat2KHR*                        pSurfaceFormats);
```

## [](#_parameters)Parameters

- `physicalDevice` is the physical device that will be associated with the swapchain to be created, as described for [vkCreateSwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateSwapchainKHR.html).
- `pSurfaceInfo` is a pointer to a [VkPhysicalDeviceSurfaceInfo2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceSurfaceInfo2KHR.html) structure describing the surface and other fixed parameters that would be consumed by [vkCreateSwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateSwapchainKHR.html).
- `pSurfaceFormatCount` is a pointer to an integer related to the number of format tuples available or queried, as described below.
- `pSurfaceFormats` is either `NULL` or a pointer to an array of [VkSurfaceFormat2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceFormat2KHR.html) structures.

## [](#_description)Description

[vkGetPhysicalDeviceSurfaceFormats2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSurfaceFormats2KHR.html) behaves similarly to [vkGetPhysicalDeviceSurfaceFormatsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSurfaceFormatsKHR.html), with the ability to be extended via `pNext` chains.

If `pSurfaceFormats` is `NULL`, then the number of format tuples supported for the given `surface` is returned in `pSurfaceFormatCount`. Otherwise, `pSurfaceFormatCount` **must** point to a variable set by the application to the number of elements in the `pSurfaceFormats` array, and on return the variable is overwritten with the number of structures actually written to `pSurfaceFormats`. If the value of `pSurfaceFormatCount` is less than the number of format tuples supported, at most `pSurfaceFormatCount` structures will be written, and `VK_INCOMPLETE` will be returned instead of `VK_SUCCESS`, to indicate that not all the available values were returned.

Valid Usage

- [](#VUID-vkGetPhysicalDeviceSurfaceFormats2KHR-pSurfaceInfo-06521)VUID-vkGetPhysicalDeviceSurfaceFormats2KHR-pSurfaceInfo-06521  
  If the `VK_GOOGLE_surfaceless_query` extension is not enabled, `pSurfaceInfo->surface` **must** be a valid [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceKHR.html) handle
- [](#VUID-vkGetPhysicalDeviceSurfaceFormats2KHR-pSurfaceInfo-06522)VUID-vkGetPhysicalDeviceSurfaceFormats2KHR-pSurfaceInfo-06522  
  If `pSurfaceInfo->surface` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `pSurfaceInfo->surface` **must** be supported by `physicalDevice`, as reported by [vkGetPhysicalDeviceSurfaceSupportKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSurfaceSupportKHR.html) or an equivalent platform-specific mechanism

Valid Usage (Implicit)

- [](#VUID-vkGetPhysicalDeviceSurfaceFormats2KHR-physicalDevice-parameter)VUID-vkGetPhysicalDeviceSurfaceFormats2KHR-physicalDevice-parameter  
  `physicalDevice` **must** be a valid [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html) handle
- [](#VUID-vkGetPhysicalDeviceSurfaceFormats2KHR-pSurfaceInfo-parameter)VUID-vkGetPhysicalDeviceSurfaceFormats2KHR-pSurfaceInfo-parameter  
  `pSurfaceInfo` **must** be a valid pointer to a valid [VkPhysicalDeviceSurfaceInfo2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceSurfaceInfo2KHR.html) structure
- [](#VUID-vkGetPhysicalDeviceSurfaceFormats2KHR-pSurfaceFormatCount-parameter)VUID-vkGetPhysicalDeviceSurfaceFormats2KHR-pSurfaceFormatCount-parameter  
  `pSurfaceFormatCount` **must** be a valid pointer to a `uint32_t` value
- [](#VUID-vkGetPhysicalDeviceSurfaceFormats2KHR-pSurfaceFormats-parameter)VUID-vkGetPhysicalDeviceSurfaceFormats2KHR-pSurfaceFormats-parameter  
  If the value referenced by `pSurfaceFormatCount` is not `0`, and `pSurfaceFormats` is not `NULL`, `pSurfaceFormats` **must** be a valid pointer to an array of `pSurfaceFormatCount` [VkSurfaceFormat2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceFormat2KHR.html) structures

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

[VK\_KHR\_get\_surface\_capabilities2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_surface_capabilities2.html), [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html), [VkPhysicalDeviceSurfaceInfo2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceSurfaceInfo2KHR.html), [VkSurfaceFormat2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceFormat2KHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetPhysicalDeviceSurfaceFormats2KHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0