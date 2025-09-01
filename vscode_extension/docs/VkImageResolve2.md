# VkImageResolve2(3) Manual Page

## Name

VkImageResolve2 - Structure specifying an image resolve operation



## [](#_c_specification)C Specification

The `VkImageResolve2` structure is defined as:

```c++
// Provided by VK_VERSION_1_3
typedef struct VkImageResolve2 {
    VkStructureType             sType;
    const void*                 pNext;
    VkImageSubresourceLayers    srcSubresource;
    VkOffset3D                  srcOffset;
    VkImageSubresourceLayers    dstSubresource;
    VkOffset3D                  dstOffset;
    VkExtent3D                  extent;
} VkImageResolve2;
```

or the equivalent

```c++
// Provided by VK_KHR_copy_commands2
typedef VkImageResolve2 VkImageResolve2KHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `srcSubresource` and `dstSubresource` are [VkImageSubresourceLayers](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresourceLayers.html) structures specifying the image subresources of the images used for the source and destination image data, respectively. Resolve of depth/stencil images is not supported.
- `srcOffset` and `dstOffset` select the initial `x`, `y`, and `z` offsets in texels of the sub-regions of the source and destination image data.
- `extent` is the size in texels of the source image to resolve in `width`, `height` and `depth`.

## [](#_description)Description

Valid Usage

- [](#VUID-VkImageResolve2-aspectMask-00266)VUID-VkImageResolve2-aspectMask-00266  
  The `aspectMask` member of `srcSubresource` and `dstSubresource` **must** only contain `VK_IMAGE_ASPECT_COLOR_BIT`
- [](#VUID-VkImageResolve2-layerCount-08803)VUID-VkImageResolve2-layerCount-08803  
  If neither of the `layerCount` members of `srcSubresource` or `dstSubresource` are `VK_REMAINING_ARRAY_LAYERS`, the `layerCount` member of `srcSubresource` and `dstSubresource` **must** match
- [](#VUID-VkImageResolve2-layerCount-08804)VUID-VkImageResolve2-layerCount-08804  
  If one of the `layerCount` members of `srcSubresource` or `dstSubresource` is `VK_REMAINING_ARRAY_LAYERS`, the other member **must** be either `VK_REMAINING_ARRAY_LAYERS` or equal to the `arrayLayers` member of the [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) used to create the image minus `baseArrayLayer`

Valid Usage (Implicit)

- [](#VUID-VkImageResolve2-sType-sType)VUID-VkImageResolve2-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_IMAGE_RESOLVE_2`
- [](#VUID-VkImageResolve2-pNext-pNext)VUID-VkImageResolve2-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkImageResolve2-srcSubresource-parameter)VUID-VkImageResolve2-srcSubresource-parameter  
  `srcSubresource` **must** be a valid [VkImageSubresourceLayers](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresourceLayers.html) structure
- [](#VUID-VkImageResolve2-dstSubresource-parameter)VUID-VkImageResolve2-dstSubresource-parameter  
  `dstSubresource` **must** be a valid [VkImageSubresourceLayers](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresourceLayers.html) structure

## [](#_see_also)See Also

[VK\_KHR\_copy\_commands2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_copy_commands2.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkExtent3D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtent3D.html), [VkImageSubresourceLayers](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresourceLayers.html), [VkOffset3D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOffset3D.html), [VkResolveImageInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkResolveImageInfo2.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkImageResolve2).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0