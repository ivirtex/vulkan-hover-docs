# vkGetPhysicalDeviceSurfaceFormats2KHR(3) Manual Page

## Name

vkGetPhysicalDeviceSurfaceFormats2KHR - Query color formats supported by
surface



## <a href="#_c_specification" class="anchor"></a>C Specification

To query the supported swapchain format tuples for a surface, call:

``` c
// Provided by VK_KHR_get_surface_capabilities2
VkResult vkGetPhysicalDeviceSurfaceFormats2KHR(
    VkPhysicalDevice                            physicalDevice,
    const VkPhysicalDeviceSurfaceInfo2KHR*      pSurfaceInfo,
    uint32_t*                                   pSurfaceFormatCount,
    VkSurfaceFormat2KHR*                        pSurfaceFormats);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `physicalDevice` is the physical device that will be associated with
  the swapchain to be created, as described for
  [vkCreateSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateSwapchainKHR.html).

- `pSurfaceInfo` is a pointer to a
  [VkPhysicalDeviceSurfaceInfo2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSurfaceInfo2KHR.html)
  structure describing the surface and other fixed parameters that would
  be consumed by [vkCreateSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateSwapchainKHR.html).

- `pSurfaceFormatCount` is a pointer to an integer related to the number
  of format tuples available or queried, as described below.

- `pSurfaceFormats` is either `NULL` or a pointer to an array of
  [VkSurfaceFormat2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceFormat2KHR.html) structures.

## <a href="#_description" class="anchor"></a>Description

[vkGetPhysicalDeviceSurfaceFormats2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSurfaceFormats2KHR.html)
behaves similarly to
[vkGetPhysicalDeviceSurfaceFormatsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSurfaceFormatsKHR.html),
with the ability to be extended via `pNext` chains.

If `pSurfaceFormats` is `NULL`, then the number of format tuples
supported for the given `surface` is returned in `pSurfaceFormatCount`.
Otherwise, `pSurfaceFormatCount` **must** point to a variable set by the
application to the number of elements in the `pSurfaceFormats` array,
and on return the variable is overwritten with the number of structures
actually written to `pSurfaceFormats`. If the value of
`pSurfaceFormatCount` is less than the number of format tuples
supported, at most `pSurfaceFormatCount` structures will be written, and
`VK_INCOMPLETE` will be returned instead of `VK_SUCCESS`, to indicate
that not all the available values were returned.

Valid Usage

- <a href="#VUID-vkGetPhysicalDeviceSurfaceFormats2KHR-pSurfaceInfo-06521"
  id="VUID-vkGetPhysicalDeviceSurfaceFormats2KHR-pSurfaceInfo-06521"></a>
  VUID-vkGetPhysicalDeviceSurfaceFormats2KHR-pSurfaceInfo-06521  
  If the
  [`VK_GOOGLE_surfaceless_query`](VK_GOOGLE_surfaceless_query.html)
  extension is not enabled, `pSurfaceInfo->surface` **must** be a valid
  [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceKHR.html) handle

- <a href="#VUID-vkGetPhysicalDeviceSurfaceFormats2KHR-pSurfaceInfo-06522"
  id="VUID-vkGetPhysicalDeviceSurfaceFormats2KHR-pSurfaceInfo-06522"></a>
  VUID-vkGetPhysicalDeviceSurfaceFormats2KHR-pSurfaceInfo-06522  
  If `pSurfaceInfo->surface` is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `pSurfaceInfo->surface`
  **must** be supported by `physicalDevice`, as reported by
  [vkGetPhysicalDeviceSurfaceSupportKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSurfaceSupportKHR.html)
  or an equivalent platform-specific mechanism

Valid Usage (Implicit)

- <a
  href="#VUID-vkGetPhysicalDeviceSurfaceFormats2KHR-physicalDevice-parameter"
  id="VUID-vkGetPhysicalDeviceSurfaceFormats2KHR-physicalDevice-parameter"></a>
  VUID-vkGetPhysicalDeviceSurfaceFormats2KHR-physicalDevice-parameter  
  `physicalDevice` **must** be a valid
  [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html) handle

- <a
  href="#VUID-vkGetPhysicalDeviceSurfaceFormats2KHR-pSurfaceInfo-parameter"
  id="VUID-vkGetPhysicalDeviceSurfaceFormats2KHR-pSurfaceInfo-parameter"></a>
  VUID-vkGetPhysicalDeviceSurfaceFormats2KHR-pSurfaceInfo-parameter  
  `pSurfaceInfo` **must** be a valid pointer to a valid
  [VkPhysicalDeviceSurfaceInfo2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSurfaceInfo2KHR.html)
  structure

- <a
  href="#VUID-vkGetPhysicalDeviceSurfaceFormats2KHR-pSurfaceFormatCount-parameter"
  id="VUID-vkGetPhysicalDeviceSurfaceFormats2KHR-pSurfaceFormatCount-parameter"></a>
  VUID-vkGetPhysicalDeviceSurfaceFormats2KHR-pSurfaceFormatCount-parameter  
  `pSurfaceFormatCount` **must** be a valid pointer to a `uint32_t`
  value

- <a
  href="#VUID-vkGetPhysicalDeviceSurfaceFormats2KHR-pSurfaceFormats-parameter"
  id="VUID-vkGetPhysicalDeviceSurfaceFormats2KHR-pSurfaceFormats-parameter"></a>
  VUID-vkGetPhysicalDeviceSurfaceFormats2KHR-pSurfaceFormats-parameter  
  If the value referenced by `pSurfaceFormatCount` is not `0`, and
  `pSurfaceFormats` is not `NULL`, `pSurfaceFormats` **must** be a valid
  pointer to an array of `pSurfaceFormatCount`
  [VkSurfaceFormat2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceFormat2KHR.html) structures

Return Codes

On success, this command returns  
- `VK_SUCCESS`

- `VK_INCOMPLETE`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

- `VK_ERROR_SURFACE_LOST_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_get_surface_capabilities2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_surface_capabilities2.html),
[VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html),
[VkPhysicalDeviceSurfaceInfo2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSurfaceInfo2KHR.html),
[VkSurfaceFormat2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceFormat2KHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetPhysicalDeviceSurfaceFormats2KHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
