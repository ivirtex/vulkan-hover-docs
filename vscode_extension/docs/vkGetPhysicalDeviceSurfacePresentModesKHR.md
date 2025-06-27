# vkGetPhysicalDeviceSurfacePresentModesKHR(3) Manual Page

## Name

vkGetPhysicalDeviceSurfacePresentModesKHR - Query supported presentation
modes



## <a href="#_c_specification" class="anchor"></a>C Specification

To query the supported presentation modes for a surface, call:

``` c
// Provided by VK_KHR_surface
VkResult vkGetPhysicalDeviceSurfacePresentModesKHR(
    VkPhysicalDevice                            physicalDevice,
    VkSurfaceKHR                                surface,
    uint32_t*                                   pPresentModeCount,
    VkPresentModeKHR*                           pPresentModes);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `physicalDevice` is the physical device that will be associated with
  the swapchain to be created, as described for
  [vkCreateSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateSwapchainKHR.html).

- `surface` is the surface that will be associated with the swapchain.

- `pPresentModeCount` is a pointer to an integer related to the number
  of presentation modes available or queried, as described below.

- `pPresentModes` is either `NULL` or a pointer to an array of
  [VkPresentModeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentModeKHR.html) values, indicating the
  supported presentation modes.

## <a href="#_description" class="anchor"></a>Description

If `pPresentModes` is `NULL`, then the number of presentation modes
supported for the given `surface` is returned in `pPresentModeCount`.
Otherwise, `pPresentModeCount` **must** point to a variable set by the
application to the number of elements in the `pPresentModes` array, and
on return the variable is overwritten with the number of values actually
written to `pPresentModes`. If the value of `pPresentModeCount` is less
than the number of presentation modes supported, at most
`pPresentModeCount` values will be written, and `VK_INCOMPLETE` will be
returned instead of `VK_SUCCESS`, to indicate that not all the available
modes were returned.

If the [`VK_GOOGLE_surfaceless_query`](VK_GOOGLE_surfaceless_query.html)
extension is enabled and `surface` is
[VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the values returned in
`pPresentModes` will only indicate support for
`VK_PRESENT_MODE_FIFO_KHR`, `VK_PRESENT_MODE_SHARED_DEMAND_REFRESH_KHR`,
and `VK_PRESENT_MODE_SHARED_CONTINUOUS_REFRESH_KHR`. To query support
for any other present mode, a valid handle **must** be provided in
`surface`.

Valid Usage

- <a href="#VUID-vkGetPhysicalDeviceSurfacePresentModesKHR-surface-06524"
  id="VUID-vkGetPhysicalDeviceSurfacePresentModesKHR-surface-06524"></a>
  VUID-vkGetPhysicalDeviceSurfacePresentModesKHR-surface-06524  
  If the
  [`VK_GOOGLE_surfaceless_query`](VK_GOOGLE_surfaceless_query.html)
  extension is not enabled, `surface` **must** be a valid
  [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceKHR.html) handle

- <a href="#VUID-vkGetPhysicalDeviceSurfacePresentModesKHR-surface-06525"
  id="VUID-vkGetPhysicalDeviceSurfacePresentModesKHR-surface-06525"></a>
  VUID-vkGetPhysicalDeviceSurfacePresentModesKHR-surface-06525  
  If `surface` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `surface`
  **must** be supported by `physicalDevice`, as reported by
  [vkGetPhysicalDeviceSurfaceSupportKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSurfaceSupportKHR.html)
  or an equivalent platform-specific mechanism

Valid Usage (Implicit)

- <a
  href="#VUID-vkGetPhysicalDeviceSurfacePresentModesKHR-physicalDevice-parameter"
  id="VUID-vkGetPhysicalDeviceSurfacePresentModesKHR-physicalDevice-parameter"></a>
  VUID-vkGetPhysicalDeviceSurfacePresentModesKHR-physicalDevice-parameter  
  `physicalDevice` **must** be a valid
  [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html) handle

- <a
  href="#VUID-vkGetPhysicalDeviceSurfacePresentModesKHR-surface-parameter"
  id="VUID-vkGetPhysicalDeviceSurfacePresentModesKHR-surface-parameter"></a>
  VUID-vkGetPhysicalDeviceSurfacePresentModesKHR-surface-parameter  
  If `surface` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `surface`
  **must** be a valid [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceKHR.html) handle

- <a
  href="#VUID-vkGetPhysicalDeviceSurfacePresentModesKHR-pPresentModeCount-parameter"
  id="VUID-vkGetPhysicalDeviceSurfacePresentModesKHR-pPresentModeCount-parameter"></a>
  VUID-vkGetPhysicalDeviceSurfacePresentModesKHR-pPresentModeCount-parameter  
  `pPresentModeCount` **must** be a valid pointer to a `uint32_t` value

- <a
  href="#VUID-vkGetPhysicalDeviceSurfacePresentModesKHR-pPresentModes-parameter"
  id="VUID-vkGetPhysicalDeviceSurfacePresentModesKHR-pPresentModes-parameter"></a>
  VUID-vkGetPhysicalDeviceSurfacePresentModesKHR-pPresentModes-parameter  
  If the value referenced by `pPresentModeCount` is not `0`, and
  `pPresentModes` is not `NULL`, `pPresentModes` **must** be a valid
  pointer to an array of `pPresentModeCount`
  [VkPresentModeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentModeKHR.html) values

- <a href="#VUID-vkGetPhysicalDeviceSurfacePresentModesKHR-commonparent"
  id="VUID-vkGetPhysicalDeviceSurfacePresentModesKHR-commonparent"></a>
  VUID-vkGetPhysicalDeviceSurfacePresentModesKHR-commonparent  
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
[VkPresentModeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentModeKHR.html),
[VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetPhysicalDeviceSurfacePresentModesKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
