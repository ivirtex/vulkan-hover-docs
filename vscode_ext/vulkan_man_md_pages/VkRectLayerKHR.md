# VkRectLayerKHR(3) Manual Page

## Name

VkRectLayerKHR - Structure containing a rectangle, including layer,
changed by vkQueuePresentKHR for a given VkImage



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkRectLayerKHR` structure is defined as:

``` c
// Provided by VK_KHR_incremental_present
typedef struct VkRectLayerKHR {
    VkOffset2D    offset;
    VkExtent2D    extent;
    uint32_t      layer;
} VkRectLayerKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `offset` is the origin of the rectangle, in pixels.

- `extent` is the size of the rectangle, in pixels.

- `layer` is the layer of the image. For images with only one layer, the
  value of `layer` **must** be 0.

## <a href="#_description" class="anchor"></a>Description

Some platforms allow the size of a surface to change, and then scale the
pixels of the image to fit the surface. `VkRectLayerKHR` specifies
pixels of the swapchain’s image(s), which will be constant for the life
of the swapchain.

Valid Usage

- <a href="#VUID-VkRectLayerKHR-offset-04864"
  id="VUID-VkRectLayerKHR-offset-04864"></a>
  VUID-VkRectLayerKHR-offset-04864  
  The sum of `offset` and `extent`, after being transformed according to
  the `preTransform` member of the
  [VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainCreateInfoKHR.html) structure,
  **must** be no greater than the `imageExtent` member of the
  [VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainCreateInfoKHR.html) structure
  passed to [vkCreateSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateSwapchainKHR.html)

- <a href="#VUID-VkRectLayerKHR-layer-01262"
  id="VUID-VkRectLayerKHR-layer-01262"></a>
  VUID-VkRectLayerKHR-layer-01262  
  `layer` **must** be less than the `imageArrayLayers` member of the
  [VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainCreateInfoKHR.html) structure
  passed to [vkCreateSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateSwapchainKHR.html)

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_incremental_present](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_incremental_present.html),
[VkExtent2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtent2D.html), [VkOffset2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOffset2D.html),
[VkPresentRegionKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentRegionKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkRectLayerKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
