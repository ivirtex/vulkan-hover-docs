# VkSwapchainPresentModesCreateInfoEXT(3) Manual Page

## Name

VkSwapchainPresentModesCreateInfoEXT - All presentation modes usable by the swapchain



## [](#_c_specification)C Specification

Applications **can** modify the presentation mode used by the swapchain on a per-presentation basis. However, all presentation modes the application intends to use with the swapchain **must** be specified at swapchain creation time. To specify more than one presentation mode when creating a swapchain, include the `VkSwapchainPresentModesCreateInfoEXT` structure in the `pNext` chain of the [VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainCreateInfoKHR.html) structure.

The `VkSwapchainPresentModesCreateInfoEXT` structure is defined as:

```c++
// Provided by VK_EXT_swapchain_maintenance1
typedef struct VkSwapchainPresentModesCreateInfoEXT {
    VkStructureType            sType;
    const void*                pNext;
    uint32_t                   presentModeCount;
    const VkPresentModeKHR*    pPresentModes;
} VkSwapchainPresentModesCreateInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `presentModeCount` is the number of presentation modes provided.
- `pPresentModes` is a list of presentation modes with `presentModeCount` entries

## [](#_description)Description

Valid Usage

- [](#VUID-VkSwapchainPresentModesCreateInfoEXT-None-07762)VUID-VkSwapchainPresentModesCreateInfoEXT-None-07762  
  Each entry in pPresentModes **must** be one of the [VkPresentModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentModeKHR.html) values returned by `vkGetPhysicalDeviceSurfacePresentModesKHR` for the surface
- [](#VUID-VkSwapchainPresentModesCreateInfoEXT-presentModeFifoLatestReady-10160)VUID-VkSwapchainPresentModesCreateInfoEXT-presentModeFifoLatestReady-10160  
  If the [`presentModeFifoLatestReady`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-presentModeFifoLatestReady) feature is not enabled, pPresentModes **must** not contain `VK_PRESENT_MODE_FIFO_LATEST_READY_EXT`
- [](#VUID-VkSwapchainPresentModesCreateInfoEXT-pPresentModes-07763)VUID-VkSwapchainPresentModesCreateInfoEXT-pPresentModes-07763  
  The entries in pPresentModes **must** be a subset of the present modes returned in [VkSurfacePresentModeCompatibilityEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfacePresentModeCompatibilityEXT.html)::`pPresentModes`, given [VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainCreateInfoKHR.html)::`presentMode` in [VkSurfacePresentModeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfacePresentModeEXT.html)
- [](#VUID-VkSwapchainPresentModesCreateInfoEXT-presentMode-07764)VUID-VkSwapchainPresentModesCreateInfoEXT-presentMode-07764  
  [VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainCreateInfoKHR.html)::`presentMode` **must** be included in `pPresentModes`

Valid Usage (Implicit)

- [](#VUID-VkSwapchainPresentModesCreateInfoEXT-sType-sType)VUID-VkSwapchainPresentModesCreateInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SWAPCHAIN_PRESENT_MODES_CREATE_INFO_EXT`
- [](#VUID-VkSwapchainPresentModesCreateInfoEXT-pPresentModes-parameter)VUID-VkSwapchainPresentModesCreateInfoEXT-pPresentModes-parameter  
  `pPresentModes` **must** be a valid pointer to an array of `presentModeCount` valid [VkPresentModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentModeKHR.html) values
- [](#VUID-VkSwapchainPresentModesCreateInfoEXT-presentModeCount-arraylength)VUID-VkSwapchainPresentModesCreateInfoEXT-presentModeCount-arraylength  
  `presentModeCount` **must** be greater than `0`

## [](#_see_also)See Also

[VK\_EXT\_swapchain\_maintenance1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_swapchain_maintenance1.html), [VkPresentModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentModeKHR.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSwapchainPresentModesCreateInfoEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0