# VkSwapchainPresentModesCreateInfoKHR(3) Manual Page

## Name

VkSwapchainPresentModesCreateInfoKHR - All presentation modes usable by the swapchain



## [](#_c_specification)C Specification

Applications **can** modify the presentation mode used by the swapchain on a per-presentation basis. However, all presentation modes the application intends to use with the swapchain **must** be specified at swapchain creation time. To specify more than one presentation mode when creating a swapchain, include the `VkSwapchainPresentModesCreateInfoKHR` structure in the `pNext` chain of the [VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainCreateInfoKHR.html) structure.

The `VkSwapchainPresentModesCreateInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_swapchain_maintenance1
typedef struct VkSwapchainPresentModesCreateInfoKHR {
    VkStructureType            sType;
    const void*                pNext;
    uint32_t                   presentModeCount;
    const VkPresentModeKHR*    pPresentModes;
} VkSwapchainPresentModesCreateInfoKHR;
```

or the equivalent

```c++
// Provided by VK_EXT_swapchain_maintenance1
typedef VkSwapchainPresentModesCreateInfoKHR VkSwapchainPresentModesCreateInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `presentModeCount` is the number of presentation modes provided.
- `pPresentModes` is a list of presentation modes with `presentModeCount` entries

## [](#_description)Description

Valid Usage

- [](#VUID-VkSwapchainPresentModesCreateInfoKHR-None-07762)VUID-VkSwapchainPresentModesCreateInfoKHR-None-07762  
  Each entry in pPresentModes **must** be one of the [VkPresentModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentModeKHR.html) values returned by `vkGetPhysicalDeviceSurfacePresentModesKHR` for the surface
- [](#VUID-VkSwapchainPresentModesCreateInfoKHR-presentModeFifoLatestReady-10160)VUID-VkSwapchainPresentModesCreateInfoKHR-presentModeFifoLatestReady-10160  
  If the [`presentModeFifoLatestReady`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-presentModeFifoLatestReady) feature is not enabled, pPresentModes **must** not contain `VK_PRESENT_MODE_FIFO_LATEST_READY_KHR`
- [](#VUID-VkSwapchainPresentModesCreateInfoKHR-pPresentModes-07763)VUID-VkSwapchainPresentModesCreateInfoKHR-pPresentModes-07763  
  The entries in pPresentModes **must** be a subset of the present modes returned in [VkSurfacePresentModeCompatibilityKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfacePresentModeCompatibilityKHR.html)::`pPresentModes`, given [VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainCreateInfoKHR.html)::`presentMode` in [VkSurfacePresentModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfacePresentModeKHR.html)
- [](#VUID-VkSwapchainPresentModesCreateInfoKHR-presentMode-07764)VUID-VkSwapchainPresentModesCreateInfoKHR-presentMode-07764  
  [VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainCreateInfoKHR.html)::`presentMode` **must** be included in `pPresentModes`

Valid Usage (Implicit)

- [](#VUID-VkSwapchainPresentModesCreateInfoKHR-sType-sType)VUID-VkSwapchainPresentModesCreateInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SWAPCHAIN_PRESENT_MODES_CREATE_INFO_KHR`
- [](#VUID-VkSwapchainPresentModesCreateInfoKHR-pPresentModes-parameter)VUID-VkSwapchainPresentModesCreateInfoKHR-pPresentModes-parameter  
  `pPresentModes` **must** be a valid pointer to an array of `presentModeCount` valid [VkPresentModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentModeKHR.html) values
- [](#VUID-VkSwapchainPresentModesCreateInfoKHR-presentModeCount-arraylength)VUID-VkSwapchainPresentModesCreateInfoKHR-presentModeCount-arraylength  
  `presentModeCount` **must** be greater than `0`

## [](#_see_also)See Also

[VK\_EXT\_swapchain\_maintenance1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_swapchain_maintenance1.html), [VK\_KHR\_swapchain\_maintenance1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_swapchain_maintenance1.html), [VkPresentModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentModeKHR.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSwapchainPresentModesCreateInfoKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0