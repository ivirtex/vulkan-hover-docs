# VkImageViewMinLodCreateInfoEXT(3) Manual Page

## Name

VkImageViewMinLodCreateInfoEXT - Structure describing the minimum LOD of an image view



## [](#_c_specification)C Specification

The `VkImageViewMinLodCreateInfoEXT` structure is defined as:

```c++
// Provided by VK_EXT_image_view_min_lod
typedef struct VkImageViewMinLodCreateInfoEXT {
    VkStructureType    sType;
    const void*        pNext;
    float              minLod;
} VkImageViewMinLodCreateInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `minLod` is the value to clamp the minimum LOD accessible by this [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html).

## [](#_description)Description

If the `pNext` chain includes a `VkImageViewMinLodCreateInfoEXT` structure, then that structure includes a parameter specifying a value to clamp the minimum LOD value during [Image Level(s) Selection](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#textures-image-level-selection), [Texel Gathering](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#textures-gather) and [Integer Texel Coordinate Operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#textures-integer-coordinate-operations).

If the image view contains `VkImageViewMinLodCreateInfoEXT` and it is used as part of a sampling operation:

minLodFloatimageView = `minLod`

otherwise:

minLodFloatimageView = 0.0

An integer variant of this parameter is also defined for sampling operations which access integer mipmap levels:

minLodIntegerimageView = ⌊minLodFloatimageView⌋

Valid Usage

- [](#VUID-VkImageViewMinLodCreateInfoEXT-minLod-06455)VUID-VkImageViewMinLodCreateInfoEXT-minLod-06455  
  If the [`minLod`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-minLod) feature is not enabled, `minLod` **must** be `0.0`
- [](#VUID-VkImageViewMinLodCreateInfoEXT-minLod-06456)VUID-VkImageViewMinLodCreateInfoEXT-minLod-06456  
  `minLod` **must** be less or equal to the index of the last mipmap level accessible to the view

Valid Usage (Implicit)

- [](#VUID-VkImageViewMinLodCreateInfoEXT-sType-sType)VUID-VkImageViewMinLodCreateInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_IMAGE_VIEW_MIN_LOD_CREATE_INFO_EXT`

## [](#_see_also)See Also

[VK\_EXT\_image\_view\_min\_lod](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_image_view_min_lod.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkImageViewMinLodCreateInfoEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0