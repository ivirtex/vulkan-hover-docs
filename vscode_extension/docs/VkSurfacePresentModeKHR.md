# VkSurfacePresentModeKHR(3) Manual Page

## Name

VkSurfacePresentModeKHR - Structure describing present mode of a surface



## [](#_c_specification)C Specification

The `VkSurfacePresentModeKHR` structure is defined as:

```c++
// Provided by VK_KHR_surface_maintenance1
typedef struct VkSurfacePresentModeKHR {
    VkStructureType     sType;
    void*               pNext;
    VkPresentModeKHR    presentMode;
} VkSurfacePresentModeKHR;
```

or the equivalent

```c++
// Provided by VK_EXT_surface_maintenance1
typedef VkSurfacePresentModeKHR VkSurfacePresentModeEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `presentMode` is the presentation mode the swapchain will use.

## [](#_description)Description

If the `VkSurfacePresentModeKHR` structure is included in the `pNext` chain of [VkPhysicalDeviceSurfaceInfo2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceSurfaceInfo2KHR.html), the values returned in [VkSurfaceCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCapabilitiesKHR.html)::`minImageCount`, [VkSurfaceCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCapabilitiesKHR.html)::`maxImageCount`, [VkSurfacePresentScalingCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfacePresentScalingCapabilitiesKHR.html)::`minScaledImageExtent`, and [VkSurfacePresentScalingCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfacePresentScalingCapabilitiesKHR.html)::`maxScaledImageExtent` are valid only for the specified `presentMode`. If `presentMode` is `VK_PRESENT_MODE_SHARED_DEMAND_REFRESH_KHR` or `VK_PRESENT_MODE_SHARED_CONTINUOUS_REFRESH_KHR`, the per-present mode image counts **must** both be one. The per-present mode image counts **may** be less-than or greater-than the image counts returned when `VkSurfacePresentModeKHR` is not provided.

Note

If [VkSwapchainPresentModesCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainPresentModesCreateInfoKHR.html) is provided to swapchain creation, the requirements for forward progress may be less strict. For example, a FIFO swapchain might only require 2 images to guarantee forward progress, but a MAILBOX one might require 4. Without the per-present image counts, such an implementation would have to return 4 in [VkSurfaceCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCapabilitiesKHR.html)::`minImageCount`, which pessimizes FIFO. Conversely, an implementation may return a low number for minImageCount, but internally bump the image count when application queries [vkGetSwapchainImagesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetSwapchainImagesKHR.html), which can surprise applications, and is not discoverable until swapchain creation. Using `VkSurfacePresentModeKHR` and [VkSwapchainPresentModesCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainPresentModesCreateInfoKHR.html) together effectively removes this problem.

[VkSwapchainPresentModesCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainPresentModesCreateInfoKHR.html) is required for the specification to be backwards compatible with applications that do not know about, or make use of this feature.

Valid Usage

- [](#VUID-VkSurfacePresentModeKHR-presentMode-07780)VUID-VkSurfacePresentModeKHR-presentMode-07780  
  `presentMode` **must** be a value reported by [vkGetPhysicalDeviceSurfacePresentModesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSurfacePresentModesKHR.html) for the specified surface

Valid Usage (Implicit)

- [](#VUID-VkSurfacePresentModeKHR-sType-sType)VUID-VkSurfacePresentModeKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SURFACE_PRESENT_MODE_KHR`
- [](#VUID-VkSurfacePresentModeKHR-presentMode-parameter)VUID-VkSurfacePresentModeKHR-presentMode-parameter  
  `presentMode` **must** be a valid [VkPresentModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentModeKHR.html) value

## [](#_see_also)See Also

[VK\_EXT\_surface\_maintenance1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_surface_maintenance1.html), [VK\_KHR\_surface\_maintenance1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_surface_maintenance1.html), [VkPresentModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentModeKHR.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSurfacePresentModeKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0