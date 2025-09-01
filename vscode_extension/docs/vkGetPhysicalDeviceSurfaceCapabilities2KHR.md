# vkGetPhysicalDeviceSurfaceCapabilities2KHR(3) Manual Page

## Name

vkGetPhysicalDeviceSurfaceCapabilities2KHR - Reports capabilities of a surface on a physical device



## [](#_c_specification)C Specification

To query the basic capabilities of a surface defined by the core or extensions, call:

```c++
// Provided by VK_KHR_get_surface_capabilities2
VkResult vkGetPhysicalDeviceSurfaceCapabilities2KHR(
    VkPhysicalDevice                            physicalDevice,
    const VkPhysicalDeviceSurfaceInfo2KHR*      pSurfaceInfo,
    VkSurfaceCapabilities2KHR*                  pSurfaceCapabilities);
```

## [](#_parameters)Parameters

- `physicalDevice` is the physical device that will be associated with the swapchain to be created, as described for [vkCreateSwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateSwapchainKHR.html).
- `pSurfaceInfo` is a pointer to a [VkPhysicalDeviceSurfaceInfo2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceSurfaceInfo2KHR.html) structure describing the surface and other fixed parameters that would be consumed by [vkCreateSwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateSwapchainKHR.html).
- `pSurfaceCapabilities` is a pointer to a [VkSurfaceCapabilities2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCapabilities2KHR.html) structure in which the capabilities are returned.

## [](#_description)Description

`vkGetPhysicalDeviceSurfaceCapabilities2KHR` behaves similarly to [vkGetPhysicalDeviceSurfaceCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSurfaceCapabilitiesKHR.html), with the ability to specify extended inputs via chained input structures, and to return extended information via chained output structures.

Valid Usage

- [](#VUID-vkGetPhysicalDeviceSurfaceCapabilities2KHR-pSurfaceInfo-06521)VUID-vkGetPhysicalDeviceSurfaceCapabilities2KHR-pSurfaceInfo-06521  
  If the `VK_GOOGLE_surfaceless_query` extension is not enabled, `pSurfaceInfo->surface` **must** be a valid [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceKHR.html) handle
- [](#VUID-vkGetPhysicalDeviceSurfaceCapabilities2KHR-pSurfaceInfo-06522)VUID-vkGetPhysicalDeviceSurfaceCapabilities2KHR-pSurfaceInfo-06522  
  If `pSurfaceInfo->surface` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `pSurfaceInfo->surface` **must** be supported by `physicalDevice`, as reported by [vkGetPhysicalDeviceSurfaceSupportKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSurfaceSupportKHR.html) or an equivalent platform-specific mechanism

<!--THE END-->

- [](#VUID-vkGetPhysicalDeviceSurfaceCapabilities2KHR-pNext-02671)VUID-vkGetPhysicalDeviceSurfaceCapabilities2KHR-pNext-02671  
  If a [VkSurfaceCapabilitiesFullScreenExclusiveEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCapabilitiesFullScreenExclusiveEXT.html) structure is included in the `pNext` chain of `pSurfaceCapabilities`, a [VkSurfaceFullScreenExclusiveWin32InfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceFullScreenExclusiveWin32InfoEXT.html) structure **must** be included in the `pNext` chain of `pSurfaceInfo`
- [](#VUID-vkGetPhysicalDeviceSurfaceCapabilities2KHR-pNext-07776)VUID-vkGetPhysicalDeviceSurfaceCapabilities2KHR-pNext-07776  
  If a [VkSurfacePresentModeCompatibilityKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfacePresentModeCompatibilityKHR.html) structure is included in the `pNext` chain of `pSurfaceCapabilities`, a [VkSurfacePresentModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfacePresentModeKHR.html) structure **must** be included in the `pNext` chain of `pSurfaceInfo`
- [](#VUID-vkGetPhysicalDeviceSurfaceCapabilities2KHR-pNext-07777)VUID-vkGetPhysicalDeviceSurfaceCapabilities2KHR-pNext-07777  
  If a [VkSurfacePresentScalingCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfacePresentScalingCapabilitiesKHR.html) structure is included in the `pNext` chain of `pSurfaceCapabilities`, a [VkSurfacePresentModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfacePresentModeKHR.html) structure **must** be included in the `pNext` chain of `pSurfaceInfo`
- [](#VUID-vkGetPhysicalDeviceSurfaceCapabilities2KHR-pNext-07778)VUID-vkGetPhysicalDeviceSurfaceCapabilities2KHR-pNext-07778  
  If a [VkSurfacePresentModeCompatibilityKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfacePresentModeCompatibilityKHR.html) structure is included in the `pNext` chain of `pSurfaceCapabilities`, `pSurfaceInfo->surface` **must** be a valid [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceKHR.html) handle
- [](#VUID-vkGetPhysicalDeviceSurfaceCapabilities2KHR-pNext-07779)VUID-vkGetPhysicalDeviceSurfaceCapabilities2KHR-pNext-07779  
  If a [VkSurfacePresentScalingCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfacePresentScalingCapabilitiesKHR.html) structure is included in the `pNext` chain of `pSurfaceCapabilities`, `pSurfaceInfo->surface` **must** be a valid [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceKHR.html) handle

Valid Usage (Implicit)

- [](#VUID-vkGetPhysicalDeviceSurfaceCapabilities2KHR-physicalDevice-parameter)VUID-vkGetPhysicalDeviceSurfaceCapabilities2KHR-physicalDevice-parameter  
  `physicalDevice` **must** be a valid [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html) handle
- [](#VUID-vkGetPhysicalDeviceSurfaceCapabilities2KHR-pSurfaceInfo-parameter)VUID-vkGetPhysicalDeviceSurfaceCapabilities2KHR-pSurfaceInfo-parameter  
  `pSurfaceInfo` **must** be a valid pointer to a valid [VkPhysicalDeviceSurfaceInfo2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceSurfaceInfo2KHR.html) structure
- [](#VUID-vkGetPhysicalDeviceSurfaceCapabilities2KHR-pSurfaceCapabilities-parameter)VUID-vkGetPhysicalDeviceSurfaceCapabilities2KHR-pSurfaceCapabilities-parameter  
  `pSurfaceCapabilities` **must** be a valid pointer to a [VkSurfaceCapabilities2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCapabilities2KHR.html) structure

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

[VK\_KHR\_get\_surface\_capabilities2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_surface_capabilities2.html), [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html), [VkPhysicalDeviceSurfaceInfo2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceSurfaceInfo2KHR.html), [VkSurfaceCapabilities2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCapabilities2KHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetPhysicalDeviceSurfaceCapabilities2KHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0