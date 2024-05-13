# vkGetPhysicalDeviceSurfaceFormatsKHR(3) Manual Page

## Name

vkGetPhysicalDeviceSurfaceFormatsKHR - Query color formats supported by
surface



## <a href="#_c_specification" class="anchor"></a>C Specification

To query the supported swapchain format-color space pairs for a surface,
call:

``` c
// Provided by VK_KHR_surface
VkResult vkGetPhysicalDeviceSurfaceFormatsKHR(
    VkPhysicalDevice                            physicalDevice,
    VkSurfaceKHR                                surface,
    uint32_t*                                   pSurfaceFormatCount,
    VkSurfaceFormatKHR*                         pSurfaceFormats);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `physicalDevice` is the physical device that will be associated with
  the swapchain to be created, as described for
  [vkCreateSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateSwapchainKHR.html).

- `surface` is the surface that will be associated with the swapchain.

- `pSurfaceFormatCount` is a pointer to an integer related to the number
  of format pairs available or queried, as described below.

- `pSurfaceFormats` is either `NULL` or a pointer to an array of
  `VkSurfaceFormatKHR` structures.

## <a href="#_description" class="anchor"></a>Description

If `pSurfaceFormats` is `NULL`, then the number of format pairs
supported for the given `surface` is returned in `pSurfaceFormatCount`.
Otherwise, `pSurfaceFormatCount` **must** point to a variable set by the
user to the number of elements in the `pSurfaceFormats` array, and on
return the variable is overwritten with the number of structures
actually written to `pSurfaceFormats`. If the value of
`pSurfaceFormatCount` is less than the number of format pairs supported,
at most `pSurfaceFormatCount` structures will be written, and
`VK_INCOMPLETE` will be returned instead of `VK_SUCCESS`, to indicate
that not all the available format pairs were returned.

The number of format pairs supported **must** be greater than or equal
to 1. `pSurfaceFormats` **must** not contain an entry whose value for
`format` is `VK_FORMAT_UNDEFINED`.

If `pSurfaceFormats` includes an entry whose value for `colorSpace` is
`VK_COLOR_SPACE_SRGB_NONLINEAR_KHR` and whose value for `format` is a
UNORM (or SRGB) format and the corresponding SRGB (or UNORM) format is a
color renderable format for `VK_IMAGE_TILING_OPTIMAL`, then
`pSurfaceFormats` **must** also contain an entry with the same value for
`colorSpace` and `format` equal to the corresponding SRGB (or UNORM)
format.

If the [`VK_GOOGLE_surfaceless_query`](VK_GOOGLE_surfaceless_query.html)
extension is enabled, the values returned in `pSurfaceFormats` will be
identical for every valid surface created on this physical device, and
so `surface` **can** be [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html).

Valid Usage

- <a href="#VUID-vkGetPhysicalDeviceSurfaceFormatsKHR-surface-06524"
  id="VUID-vkGetPhysicalDeviceSurfaceFormatsKHR-surface-06524"></a>
  VUID-vkGetPhysicalDeviceSurfaceFormatsKHR-surface-06524  
  If the
  [`VK_GOOGLE_surfaceless_query`](VK_GOOGLE_surfaceless_query.html)
  extension is not enabled, `surface` **must** be a valid
  [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceKHR.html) handle

- <a href="#VUID-vkGetPhysicalDeviceSurfaceFormatsKHR-surface-06525"
  id="VUID-vkGetPhysicalDeviceSurfaceFormatsKHR-surface-06525"></a>
  VUID-vkGetPhysicalDeviceSurfaceFormatsKHR-surface-06525  
  If `surface` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `surface`
  **must** be supported by `physicalDevice`, as reported by
  [vkGetPhysicalDeviceSurfaceSupportKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSurfaceSupportKHR.html)
  or an equivalent platform-specific mechanism

Valid Usage (Implicit)

- <a
  href="#VUID-vkGetPhysicalDeviceSurfaceFormatsKHR-physicalDevice-parameter"
  id="VUID-vkGetPhysicalDeviceSurfaceFormatsKHR-physicalDevice-parameter"></a>
  VUID-vkGetPhysicalDeviceSurfaceFormatsKHR-physicalDevice-parameter  
  `physicalDevice` **must** be a valid
  [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html) handle

- <a href="#VUID-vkGetPhysicalDeviceSurfaceFormatsKHR-surface-parameter"
  id="VUID-vkGetPhysicalDeviceSurfaceFormatsKHR-surface-parameter"></a>
  VUID-vkGetPhysicalDeviceSurfaceFormatsKHR-surface-parameter  
  If `surface` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `surface`
  **must** be a valid [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceKHR.html) handle

- <a
  href="#VUID-vkGetPhysicalDeviceSurfaceFormatsKHR-pSurfaceFormatCount-parameter"
  id="VUID-vkGetPhysicalDeviceSurfaceFormatsKHR-pSurfaceFormatCount-parameter"></a>
  VUID-vkGetPhysicalDeviceSurfaceFormatsKHR-pSurfaceFormatCount-parameter  
  `pSurfaceFormatCount` **must** be a valid pointer to a `uint32_t`
  value

- <a
  href="#VUID-vkGetPhysicalDeviceSurfaceFormatsKHR-pSurfaceFormats-parameter"
  id="VUID-vkGetPhysicalDeviceSurfaceFormatsKHR-pSurfaceFormats-parameter"></a>
  VUID-vkGetPhysicalDeviceSurfaceFormatsKHR-pSurfaceFormats-parameter  
  If the value referenced by `pSurfaceFormatCount` is not `0`, and
  `pSurfaceFormats` is not `NULL`, `pSurfaceFormats` **must** be a valid
  pointer to an array of `pSurfaceFormatCount`
  [VkSurfaceFormatKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceFormatKHR.html) structures

- <a href="#VUID-vkGetPhysicalDeviceSurfaceFormatsKHR-commonparent"
  id="VUID-vkGetPhysicalDeviceSurfaceFormatsKHR-commonparent"></a>
  VUID-vkGetPhysicalDeviceSurfaceFormatsKHR-commonparent  
  Both of `physicalDevice`, and `surface` that are valid handles of
  non-ignored parameters **must** have been created, allocated, or
  retrieved from the same [VkInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstance.html)

Return Codes

On success, this command returns  
- `VK_SUCCESS`

- `VK_INCOMPLETE`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

- `VK_ERROR_SURFACE_LOST_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_surface](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_surface.html),
[VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html),
[VkSurfaceFormatKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceFormatKHR.html),
[VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetPhysicalDeviceSurfaceFormatsKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
