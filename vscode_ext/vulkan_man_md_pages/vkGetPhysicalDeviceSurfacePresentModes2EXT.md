# vkGetPhysicalDeviceSurfacePresentModes2EXT(3) Manual Page

## Name

vkGetPhysicalDeviceSurfacePresentModes2EXT - Query supported
presentation modes



## <a href="#_c_specification" class="anchor"></a>C Specification

Alternatively, to query the supported presentation modes for a surface
combined with select other fixed swapchain creation parameters, call:

``` c
// Provided by VK_EXT_full_screen_exclusive
VkResult vkGetPhysicalDeviceSurfacePresentModes2EXT(
    VkPhysicalDevice                            physicalDevice,
    const VkPhysicalDeviceSurfaceInfo2KHR*      pSurfaceInfo,
    uint32_t*                                   pPresentModeCount,
    VkPresentModeKHR*                           pPresentModes);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `physicalDevice` is the physical device that will be associated with
  the swapchain to be created, as described for
  [vkCreateSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateSwapchainKHR.html).

- `pSurfaceInfo` is a pointer to a
  [VkPhysicalDeviceSurfaceInfo2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSurfaceInfo2KHR.html)
  structure describing the surface and other fixed parameters that would
  be consumed by [vkCreateSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateSwapchainKHR.html).

- `pPresentModeCount` is a pointer to an integer related to the number
  of presentation modes available or queried, as described below.

- `pPresentModes` is either `NULL` or a pointer to an array of
  [VkPresentModeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentModeKHR.html) values, indicating the
  supported presentation modes.

## <a href="#_description" class="anchor"></a>Description

`vkGetPhysicalDeviceSurfacePresentModes2EXT` behaves similarly to
[vkGetPhysicalDeviceSurfacePresentModesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSurfacePresentModesKHR.html),
with the ability to specify extended inputs via chained input
structures.

Valid Usage

- <a
  href="#VUID-vkGetPhysicalDeviceSurfacePresentModes2EXT-pSurfaceInfo-06521"
  id="VUID-vkGetPhysicalDeviceSurfacePresentModes2EXT-pSurfaceInfo-06521"></a>
  VUID-vkGetPhysicalDeviceSurfacePresentModes2EXT-pSurfaceInfo-06521  
  If the
  [`VK_GOOGLE_surfaceless_query`](VK_GOOGLE_surfaceless_query.html)
  extension is not enabled, `pSurfaceInfo->surface` **must** be a valid
  [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceKHR.html) handle

- <a
  href="#VUID-vkGetPhysicalDeviceSurfacePresentModes2EXT-pSurfaceInfo-06522"
  id="VUID-vkGetPhysicalDeviceSurfacePresentModes2EXT-pSurfaceInfo-06522"></a>
  VUID-vkGetPhysicalDeviceSurfacePresentModes2EXT-pSurfaceInfo-06522  
  If `pSurfaceInfo->surface` is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `pSurfaceInfo->surface`
  **must** be supported by `physicalDevice`, as reported by
  [vkGetPhysicalDeviceSurfaceSupportKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSurfaceSupportKHR.html)
  or an equivalent platform-specific mechanism

Valid Usage (Implicit)

- <a
  href="#VUID-vkGetPhysicalDeviceSurfacePresentModes2EXT-physicalDevice-parameter"
  id="VUID-vkGetPhysicalDeviceSurfacePresentModes2EXT-physicalDevice-parameter"></a>
  VUID-vkGetPhysicalDeviceSurfacePresentModes2EXT-physicalDevice-parameter  
  `physicalDevice` **must** be a valid
  [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html) handle

- <a
  href="#VUID-vkGetPhysicalDeviceSurfacePresentModes2EXT-pSurfaceInfo-parameter"
  id="VUID-vkGetPhysicalDeviceSurfacePresentModes2EXT-pSurfaceInfo-parameter"></a>
  VUID-vkGetPhysicalDeviceSurfacePresentModes2EXT-pSurfaceInfo-parameter  
  `pSurfaceInfo` **must** be a valid pointer to a valid
  [VkPhysicalDeviceSurfaceInfo2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSurfaceInfo2KHR.html)
  structure

- <a
  href="#VUID-vkGetPhysicalDeviceSurfacePresentModes2EXT-pPresentModeCount-parameter"
  id="VUID-vkGetPhysicalDeviceSurfacePresentModes2EXT-pPresentModeCount-parameter"></a>
  VUID-vkGetPhysicalDeviceSurfacePresentModes2EXT-pPresentModeCount-parameter  
  `pPresentModeCount` **must** be a valid pointer to a `uint32_t` value

- <a
  href="#VUID-vkGetPhysicalDeviceSurfacePresentModes2EXT-pPresentModes-parameter"
  id="VUID-vkGetPhysicalDeviceSurfacePresentModes2EXT-pPresentModes-parameter"></a>
  VUID-vkGetPhysicalDeviceSurfacePresentModes2EXT-pPresentModes-parameter  
  If the value referenced by `pPresentModeCount` is not `0`, and
  `pPresentModes` is not `NULL`, `pPresentModes` **must** be a valid
  pointer to an array of `pPresentModeCount`
  [VkPresentModeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentModeKHR.html) values

Return Codes

On success, this command returns  
- `VK_SUCCESS`

- `VK_INCOMPLETE`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

- `VK_ERROR_SURFACE_LOST_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_full_screen_exclusive](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_full_screen_exclusive.html),
[VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html),
[VkPhysicalDeviceSurfaceInfo2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSurfaceInfo2KHR.html),
[VkPresentModeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentModeKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetPhysicalDeviceSurfacePresentModes2EXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
