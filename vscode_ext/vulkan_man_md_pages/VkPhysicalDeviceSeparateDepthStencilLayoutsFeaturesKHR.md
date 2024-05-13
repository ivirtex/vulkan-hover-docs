# VkPhysicalDeviceSeparateDepthStencilLayoutsFeatures(3) Manual Page

## Name

VkPhysicalDeviceSeparateDepthStencilLayoutsFeatures - Structure
describing whether the implementation can do depth and stencil image
barriers separately



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceSeparateDepthStencilLayoutsFeatures` structure is
defined as:

``` c
// Provided by VK_VERSION_1_2
typedef struct VkPhysicalDeviceSeparateDepthStencilLayoutsFeatures {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           separateDepthStencilLayouts;
} VkPhysicalDeviceSeparateDepthStencilLayoutsFeatures;
```

or the equivalent

``` c
// Provided by VK_KHR_separate_depth_stencil_layouts
typedef VkPhysicalDeviceSeparateDepthStencilLayoutsFeatures VkPhysicalDeviceSeparateDepthStencilLayoutsFeaturesKHR;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following feature:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

## <a href="#_description" class="anchor"></a>Description

- <span id="extension-features-separateDepthStencilLayouts"></span>
  `separateDepthStencilLayouts` indicates whether the implementation
  supports a `VkImageMemoryBarrier` for a depth/stencil image with only
  one of `VK_IMAGE_ASPECT_DEPTH_BIT` or `VK_IMAGE_ASPECT_STENCIL_BIT`
  set, and whether `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL`,
  `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL`,
  `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL`, or
  `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL` can be used.

If the `VkPhysicalDeviceSeparateDepthStencilLayoutsFeatures` structure
is included in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceSeparateDepthStencilLayoutsFeatures` **can** also be
used in the `pNext` chain of
[VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to selectively enable
these features.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceSeparateDepthStencilLayoutsFeatures-sType-sType"
  id="VUID-VkPhysicalDeviceSeparateDepthStencilLayoutsFeatures-sType-sType"></a>
  VUID-VkPhysicalDeviceSeparateDepthStencilLayoutsFeatures-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SEPARATE_DEPTH_STENCIL_LAYOUTS_FEATURES`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_separate_depth_stencil_layouts](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_separate_depth_stencil_layouts.html),
[VK_VERSION_1_2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_2.html), [VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceSeparateDepthStencilLayoutsFeatures"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
