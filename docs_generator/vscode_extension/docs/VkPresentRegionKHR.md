# VkPresentRegionKHR(3) Manual Page

## Name

VkPresentRegionKHR - Structure containing rectangular region changed by vkQueuePresentKHR for a given VkImage



## [](#_c_specification)C Specification

For a given image and swapchain, the region to present is specified by the `VkPresentRegionKHR` structure, which is defined as:

```c++
// Provided by VK_KHR_incremental_present
typedef struct VkPresentRegionKHR {
    uint32_t                 rectangleCount;
    const VkRectLayerKHR*    pRectangles;
} VkPresentRegionKHR;
```

## [](#_members)Members

- `rectangleCount` is the number of rectangles in `pRectangles`, or zero if the entire image has changed and should be presented.
- `pRectangles` is either `NULL` or a pointer to an array of `VkRectLayerKHR` structures. The `VkRectLayerKHR` structure is the framebuffer coordinates, plus layer, of a portion of a presentable image that has changed and **must** be presented. If non-`NULL`, each entry in `pRectangles` is a rectangle of the given image that has changed since the last image was presented to the given swapchain. The rectangles **must** be specified relative to [VkSurfaceCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCapabilitiesKHR.html)::`currentTransform`, regardless of the swapchain’s `preTransform`. The presentation engine will apply the `preTransform` transformation to the rectangles, along with any further transformation it applies to the image content.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkPresentRegionKHR-pRectangles-parameter)VUID-VkPresentRegionKHR-pRectangles-parameter  
  If `rectangleCount` is not `0`, and `pRectangles` is not `NULL`, `pRectangles` **must** be a valid pointer to an array of `rectangleCount` valid [VkRectLayerKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRectLayerKHR.html) structures

## [](#_see_also)See Also

[VK\_KHR\_incremental\_present](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_incremental_present.html), [VkPresentRegionsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentRegionsKHR.html), [VkRectLayerKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRectLayerKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPresentRegionKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0