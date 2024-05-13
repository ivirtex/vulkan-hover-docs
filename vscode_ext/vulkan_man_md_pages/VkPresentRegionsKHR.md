# VkPresentRegionsKHR(3) Manual Page

## Name

VkPresentRegionsKHR - Structure hint of rectangular regions changed by
vkQueuePresentKHR



## <a href="#_c_specification" class="anchor"></a>C Specification

When the [`VK_KHR_incremental_present`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_incremental_present.html)
extension is enabled, additional fields **can** be specified that allow
an application to specify that only certain rectangular regions of the
presentable images of a swapchain are changed. This is an optimization
hint that a presentation engine **may** use to only update the region of
a surface that is actually changing. The application still **must**
ensure that all pixels of a presented image contain the desired values,
in case the presentation engine ignores this hint. An application
**can** provide this hint by adding a `VkPresentRegionsKHR` structure to
the `pNext` chain of the `VkPresentInfoKHR` structure.

The `VkPresentRegionsKHR` structure is defined as:

``` c
// Provided by VK_KHR_incremental_present
typedef struct VkPresentRegionsKHR {
    VkStructureType              sType;
    const void*                  pNext;
    uint32_t                     swapchainCount;
    const VkPresentRegionKHR*    pRegions;
} VkPresentRegionsKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `swapchainCount` is the number of swapchains being presented to by
  this command.

- `pRegions` is `NULL` or a pointer to an array of `VkPresentRegionKHR`
  elements with `swapchainCount` entries. If not `NULL`, each element of
  `pRegions` contains the region that has changed since the last present
  to the swapchain in the corresponding entry in the
  `VkPresentInfoKHR`::`pSwapchains` array.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkPresentRegionsKHR-swapchainCount-01260"
  id="VUID-VkPresentRegionsKHR-swapchainCount-01260"></a>
  VUID-VkPresentRegionsKHR-swapchainCount-01260  
  `swapchainCount` **must** be the same value as
  `VkPresentInfoKHR`::`swapchainCount`, where `VkPresentInfoKHR` is
  included in the `pNext` chain of this `VkPresentRegionsKHR` structure

Valid Usage (Implicit)

- <a href="#VUID-VkPresentRegionsKHR-sType-sType"
  id="VUID-VkPresentRegionsKHR-sType-sType"></a>
  VUID-VkPresentRegionsKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PRESENT_REGIONS_KHR`

- <a href="#VUID-VkPresentRegionsKHR-pRegions-parameter"
  id="VUID-VkPresentRegionsKHR-pRegions-parameter"></a>
  VUID-VkPresentRegionsKHR-pRegions-parameter  
  If `pRegions` is not `NULL`, `pRegions` **must** be a valid pointer to
  an array of `swapchainCount` valid
  [VkPresentRegionKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentRegionKHR.html) structures

- <a href="#VUID-VkPresentRegionsKHR-swapchainCount-arraylength"
  id="VUID-VkPresentRegionsKHR-swapchainCount-arraylength"></a>
  VUID-VkPresentRegionsKHR-swapchainCount-arraylength  
  `swapchainCount` **must** be greater than `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_incremental_present](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_incremental_present.html),
[VkPresentRegionKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentRegionKHR.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPresentRegionsKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
