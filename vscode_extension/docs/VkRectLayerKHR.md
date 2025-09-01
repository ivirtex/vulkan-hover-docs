# VkRectLayerKHR(3) Manual Page

## Name

VkRectLayerKHR - Structure containing a rectangle, including layer, changed by vkQueuePresentKHR for a given VkImage



## [](#_c_specification)C Specification

The `VkRectLayerKHR` structure is defined as:

```c++
// Provided by VK_KHR_incremental_present
typedef struct VkRectLayerKHR {
    VkOffset2D    offset;
    VkExtent2D    extent;
    uint32_t      layer;
} VkRectLayerKHR;
```

## [](#_members)Members

- `offset` is the origin of the rectangle, in pixels.
- `extent` is the size of the rectangle, in pixels.
- `layer` is the layer of the image. For images with only one layer, the value of `layer` **must** be 0.

## [](#_description)Description

Some platforms allow the size of a surface to change, and then scale the pixels of the image to fit the surface. `VkRectLayerKHR` specifies pixels of the swapchain’s image(s), which will be constant for the life of the swapchain.

Valid Usage

- [](#VUID-VkRectLayerKHR-offset-04864)VUID-VkRectLayerKHR-offset-04864  
  The sum of `offset` and `extent`, after being transformed according to the `preTransform` member of the [VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainCreateInfoKHR.html) structure, **must** be no greater than the `imageExtent` member of the [VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainCreateInfoKHR.html) structure passed to [vkCreateSwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateSwapchainKHR.html)
- [](#VUID-VkRectLayerKHR-layer-01262)VUID-VkRectLayerKHR-layer-01262  
  `layer` **must** be less than the `imageArrayLayers` member of the [VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainCreateInfoKHR.html) structure passed to [vkCreateSwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateSwapchainKHR.html)

## [](#_see_also)See Also

[VK\_KHR\_incremental\_present](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_incremental_present.html), [VkExtent2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtent2D.html), [VkOffset2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOffset2D.html), [VkPresentRegionKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentRegionKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkRectLayerKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0