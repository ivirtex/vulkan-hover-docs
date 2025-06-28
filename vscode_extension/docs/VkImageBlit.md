# VkImageBlit(3) Manual Page

## Name

VkImageBlit - Structure specifying an image blit operation



## [](#_c_specification)C Specification

The `VkImageBlit` structure is defined as:

```c++
// Provided by VK_VERSION_1_0
typedef struct VkImageBlit {
    VkImageSubresourceLayers    srcSubresource;
    VkOffset3D                  srcOffsets[2];
    VkImageSubresourceLayers    dstSubresource;
    VkOffset3D                  dstOffsets[2];
} VkImageBlit;
```

## [](#_members)Members

- `srcSubresource` is the subresource to blit from.
- `srcOffsets` is a pointer to an array of two [VkOffset3D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOffset3D.html) structures specifying the bounds of the source region within `srcSubresource`.
- `dstSubresource` is the subresource to blit into.
- `dstOffsets` is a pointer to an array of two [VkOffset3D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOffset3D.html) structures specifying the bounds of the destination region within `dstSubresource`.

## [](#_description)Description

For each element of the `pRegions` array, a blit operation is performed for the specified source and destination regions.

Valid Usage

- [](#VUID-VkImageBlit-aspectMask-00238)VUID-VkImageBlit-aspectMask-00238  
  The `aspectMask` member of `srcSubresource` and `dstSubresource` **must** match
- [](#VUID-VkImageBlit-layerCount-08800)VUID-VkImageBlit-layerCount-08800  
  If neither of the `layerCount` members of `srcSubresource` or `dstSubresource` are `VK_REMAINING_ARRAY_LAYERS`, the `layerCount` members of `srcSubresource` or `dstSubresource` **must** match
- [](#VUID-VkImageBlit-layerCount-08801)VUID-VkImageBlit-layerCount-08801  
  If one of the `layerCount` members of `srcSubresource` or `dstSubresource` is `VK_REMAINING_ARRAY_LAYERS`, the other member **must** be either `VK_REMAINING_ARRAY_LAYERS` or equal to the `arrayLayers` member of the [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) used to create the image minus `baseArrayLayer`

Valid Usage (Implicit)

- [](#VUID-VkImageBlit-srcSubresource-parameter)VUID-VkImageBlit-srcSubresource-parameter  
  `srcSubresource` **must** be a valid [VkImageSubresourceLayers](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresourceLayers.html) structure
- [](#VUID-VkImageBlit-dstSubresource-parameter)VUID-VkImageBlit-dstSubresource-parameter  
  `dstSubresource` **must** be a valid [VkImageSubresourceLayers](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresourceLayers.html) structure

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkImageSubresourceLayers](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresourceLayers.html), [VkOffset3D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOffset3D.html), [vkCmdBlitImage](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBlitImage.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkImageBlit)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0