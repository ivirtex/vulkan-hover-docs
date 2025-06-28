# vkGetPhysicalDeviceSurfacePresentModesKHR(3) Manual Page

## Name

vkGetPhysicalDeviceSurfacePresentModesKHR - Query supported presentation modes



## [](#_c_specification)C Specification

To query the supported presentation modes for a surface, call:

```c++
// Provided by VK_KHR_surface
VkResult vkGetPhysicalDeviceSurfacePresentModesKHR(
    VkPhysicalDevice                            physicalDevice,
    VkSurfaceKHR                                surface,
    uint32_t*                                   pPresentModeCount,
    VkPresentModeKHR*                           pPresentModes);
```

## [](#_parameters)Parameters

- `physicalDevice` is the physical device that will be associated with the swapchain to be created, as described for [vkCreateSwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateSwapchainKHR.html).
- `surface` is the surface that will be associated with the swapchain.
- `pPresentModeCount` is a pointer to an integer related to the number of presentation modes available or queried, as described below.
- `pPresentModes` is either `NULL` or a pointer to an array of [VkPresentModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentModeKHR.html) values, indicating the supported presentation modes.

## [](#_description)Description

If `pPresentModes` is `NULL`, then the number of presentation modes supported for the given `surface` is returned in `pPresentModeCount`. Otherwise, `pPresentModeCount` **must** point to a variable set by the application to the number of elements in the `pPresentModes` array, and on return the variable is overwritten with the number of values actually written to `pPresentModes`. If the value of `pPresentModeCount` is less than the number of presentation modes supported, at most `pPresentModeCount` values will be written, and `VK_INCOMPLETE` will be returned instead of `VK_SUCCESS`, to indicate that not all the available modes were returned.

If the `VK_GOOGLE_surfaceless_query` extension is enabled and `surface` is [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), the values returned in `pPresentModes` will only indicate support for `VK_PRESENT_MODE_FIFO_KHR`, `VK_PRESENT_MODE_SHARED_DEMAND_REFRESH_KHR`, and `VK_PRESENT_MODE_SHARED_CONTINUOUS_REFRESH_KHR`. To query support for any other present mode, a valid handle **must** be provided in `surface`.

Valid Usage

- [](#VUID-vkGetPhysicalDeviceSurfacePresentModesKHR-surface-06524)VUID-vkGetPhysicalDeviceSurfacePresentModesKHR-surface-06524  
  If the `VK_GOOGLE_surfaceless_query` extension is not enabled, `surface` **must** be a valid [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceKHR.html) handle
- [](#VUID-vkGetPhysicalDeviceSurfacePresentModesKHR-surface-06525)VUID-vkGetPhysicalDeviceSurfacePresentModesKHR-surface-06525  
  If `surface` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `surface` **must** be supported by `physicalDevice`, as reported by [vkGetPhysicalDeviceSurfaceSupportKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSurfaceSupportKHR.html) or an equivalent platform-specific mechanism

Valid Usage (Implicit)

- [](#VUID-vkGetPhysicalDeviceSurfacePresentModesKHR-physicalDevice-parameter)VUID-vkGetPhysicalDeviceSurfacePresentModesKHR-physicalDevice-parameter  
  `physicalDevice` **must** be a valid [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html) handle
- [](#VUID-vkGetPhysicalDeviceSurfacePresentModesKHR-surface-parameter)VUID-vkGetPhysicalDeviceSurfacePresentModesKHR-surface-parameter  
  If `surface` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `surface` **must** be a valid [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceKHR.html) handle
- [](#VUID-vkGetPhysicalDeviceSurfacePresentModesKHR-pPresentModeCount-parameter)VUID-vkGetPhysicalDeviceSurfacePresentModesKHR-pPresentModeCount-parameter  
  `pPresentModeCount` **must** be a valid pointer to a `uint32_t` value
- [](#VUID-vkGetPhysicalDeviceSurfacePresentModesKHR-pPresentModes-parameter)VUID-vkGetPhysicalDeviceSurfacePresentModesKHR-pPresentModes-parameter  
  If the value referenced by `pPresentModeCount` is not `0`, and `pPresentModes` is not `NULL`, `pPresentModes` **must** be a valid pointer to an array of `pPresentModeCount` [VkPresentModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentModeKHR.html) values
- [](#VUID-vkGetPhysicalDeviceSurfacePresentModesKHR-commonparent)VUID-vkGetPhysicalDeviceSurfacePresentModesKHR-commonparent  
  Both of `physicalDevice`, and `surface` that are valid handles of non-ignored parameters **must** have been created, allocated, or retrieved from the same [VkInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstance.html)

Return Codes

On success, this command returns

- `VK_SUCCESS`
- `VK_INCOMPLETE`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_SURFACE_LOST_KHR`

## [](#_see_also)See Also

[VK\_KHR\_surface](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_surface.html), [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html), [VkPresentModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentModeKHR.html), [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetPhysicalDeviceSurfacePresentModesKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0