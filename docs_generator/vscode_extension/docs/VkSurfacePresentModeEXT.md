# VkSurfacePresentModeEXT(3) Manual Page

## Name

VkSurfacePresentModeEXT - Structure describing present mode of a surface



## [](#_c_specification)C Specification

The `VkSurfacePresentModeEXT` structure is defined as:

```c++
// Provided by VK_EXT_surface_maintenance1
typedef struct VkSurfacePresentModeEXT {
    VkStructureType     sType;
    void*               pNext;
    VkPresentModeKHR    presentMode;
} VkSurfacePresentModeEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `presentMode` is the presentation mode the swapchain will use.

## [](#_description)Description

If the `VkSurfacePresentModeEXT` structure is included in the `pNext` chain of [VkPhysicalDeviceSurfaceInfo2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceSurfaceInfo2KHR.html), the values returned in [VkSurfaceCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCapabilitiesKHR.html)::`minImageCount`, [VkSurfaceCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCapabilitiesKHR.html)::`maxImageCount`, [VkSurfacePresentScalingCapabilitiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfacePresentScalingCapabilitiesEXT.html)::`minScaledImageExtent`, and [VkSurfacePresentScalingCapabilitiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfacePresentScalingCapabilitiesEXT.html)::`maxScaledImageExtent` are valid only for the specified `presentMode`. If `presentMode` is `VK_PRESENT_MODE_SHARED_DEMAND_REFRESH_KHR` or `VK_PRESENT_MODE_SHARED_CONTINUOUS_REFRESH_KHR`, the per-present mode image counts **must** both be one. The per-present mode image counts **may** be less-than or greater-than the image counts returned when `VkSurfacePresentModeEXT` is not provided.

Note

If [VkSwapchainPresentModesCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainPresentModesCreateInfoEXT.html) is provided to swapchain creation, the requirements for forward progress may be less strict. For example, a FIFO swapchain might only require 2 images to guarantee forward progress, but a MAILBOX one might require 4. Without the per-present image counts, such an implementation would have to return 4 in [VkSurfaceCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCapabilitiesKHR.html)::`minImageCount`, which pessimizes FIFO. Conversely, an implementation may return a low number for minImageCount, but internally bump the image count when application queries [vkGetSwapchainImagesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetSwapchainImagesKHR.html), which can surprise applications, and is not discoverable until swapchain creation. Using `VkSurfacePresentModeEXT` and [VkSwapchainPresentModesCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainPresentModesCreateInfoEXT.html) together effectively removes this problem.

[VkSwapchainPresentModesCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainPresentModesCreateInfoEXT.html) is required for the specification to be backwards compatible with applications that do not know about, or make use of this feature.

Valid Usage

- [](#VUID-VkSurfacePresentModeEXT-presentMode-07780)VUID-VkSurfacePresentModeEXT-presentMode-07780  
  `presentMode` **must** be a value reported by [vkGetPhysicalDeviceSurfacePresentModesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSurfacePresentModesKHR.html) for the specified surface

Valid Usage (Implicit)

- [](#VUID-VkSurfacePresentModeEXT-sType-sType)VUID-VkSurfacePresentModeEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SURFACE_PRESENT_MODE_EXT`
- [](#VUID-VkSurfacePresentModeEXT-presentMode-parameter)VUID-VkSurfacePresentModeEXT-presentMode-parameter  
  `presentMode` **must** be a valid [VkPresentModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentModeKHR.html) value

## [](#_see_also)See Also

[VK\_EXT\_surface\_maintenance1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_surface_maintenance1.html), [VkPresentModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentModeKHR.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSurfacePresentModeEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0