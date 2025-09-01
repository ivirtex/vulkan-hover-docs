# VkPresentRegionsKHR(3) Manual Page

## Name

VkPresentRegionsKHR - Structure hint of rectangular regions changed by vkQueuePresentKHR



## [](#_c_specification)C Specification

When the `VK_KHR_incremental_present` extension is enabled, additional fields **can** be specified that allow an application to specify that only certain rectangular regions of the presentable images of a swapchain are changed. This is an optimization hint that a presentation engine **may** use to only update the region of a surface that is actually changing. The application still **must** ensure that all pixels of a presented image contain the desired values, in case the presentation engine ignores this hint. An application **can** provide this hint by adding a `VkPresentRegionsKHR` structure to the `pNext` chain of the `VkPresentInfoKHR` structure.

The `VkPresentRegionsKHR` structure is defined as:

```c++
// Provided by VK_KHR_incremental_present
typedef struct VkPresentRegionsKHR {
    VkStructureType              sType;
    const void*                  pNext;
    uint32_t                     swapchainCount;
    const VkPresentRegionKHR*    pRegions;
} VkPresentRegionsKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `swapchainCount` is the number of swapchains being presented to by this command.
- `pRegions` is `NULL` or a pointer to an array of `VkPresentRegionKHR` elements with `swapchainCount` entries. If not `NULL`, each element of `pRegions` contains the region that has changed since the last present to the swapchain in the corresponding entry in the `VkPresentInfoKHR`::`pSwapchains` array.

## [](#_description)Description

Valid Usage

- [](#VUID-VkPresentRegionsKHR-swapchainCount-01260)VUID-VkPresentRegionsKHR-swapchainCount-01260  
  `swapchainCount` **must** be the same value as `VkPresentInfoKHR`::`swapchainCount`, where `VkPresentInfoKHR` is included in the `pNext` chain of this `VkPresentRegionsKHR` structure

Valid Usage (Implicit)

- [](#VUID-VkPresentRegionsKHR-sType-sType)VUID-VkPresentRegionsKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PRESENT_REGIONS_KHR`
- [](#VUID-VkPresentRegionsKHR-pRegions-parameter)VUID-VkPresentRegionsKHR-pRegions-parameter  
  If `pRegions` is not `NULL`, `pRegions` **must** be a valid pointer to an array of `swapchainCount` valid [VkPresentRegionKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentRegionKHR.html) structures
- [](#VUID-VkPresentRegionsKHR-swapchainCount-arraylength)VUID-VkPresentRegionsKHR-swapchainCount-arraylength  
  `swapchainCount` **must** be greater than `0`

## [](#_see_also)See Also

[VK\_KHR\_incremental\_present](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_incremental_present.html), [VkPresentRegionKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentRegionKHR.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPresentRegionsKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0