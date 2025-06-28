# VkImageViewSlicedCreateInfoEXT(3) Manual Page

## Name

VkImageViewSlicedCreateInfoEXT - Specify the subset of 3D slices of an image view



## [](#_c_specification)C Specification

The range of 3D slices for the created image view **can** be restricted to a subset of the parent image’s Z range by adding a `VkImageViewSlicedCreateInfoEXT` structure to the `pNext` chain of [VkImageViewCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewCreateInfo.html).

The `VkImageViewSlicedCreateInfoEXT` structure is defined as:

```c++
// Provided by VK_EXT_image_sliced_view_of_3d
typedef struct VkImageViewSlicedCreateInfoEXT {
    VkStructureType    sType;
    const void*        pNext;
    uint32_t           sliceOffset;
    uint32_t           sliceCount;
} VkImageViewSlicedCreateInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `sliceOffset` is the Z-offset for the first 3D slice accessible to the image view.
- `sliceCount` is the number of 3D slices accessible to the image view.

## [](#_description)Description

When this structure is chained to [VkImageViewCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewCreateInfo.html) the `sliceOffset` field is treated as a Z-offset for the sliced view and `sliceCount` specifies the range. Shader accesses using a Z coordinate of 0 will access the depth slice corresponding to `sliceOffset` in the image, and in a shader, the maximum in-bounds Z coordinate for the view is `sliceCount` - 1.

A sliced 3D view **must** only be used with a single mip level. The slice coordinates are integer coordinates within the `subresourceRange.baseMipLevel` used to create the image view.

The effective view depth is equal to `extent.depth` used to create the `image` for this view adjusted by `subresourceRange.baseMipLevel` as specified in [Image Mip Level Sizing](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#resources-image-mip-level-sizing).

Shader access to this image view is only affected by `VkImageViewSlicedCreateInfoEXT` if it uses a descriptor of type `VK_DESCRIPTOR_TYPE_STORAGE_IMAGE`. For access using any other descriptor type, the contents of `VkImageViewSlicedCreateInfoEXT` are ignored; instead, `sliceOffset` is treated as being equal to 0, and `sliceCount` is treated as being equal to `VK_REMAINING_3D_SLICES_EXT`.

Valid Usage

- [](#VUID-VkImageViewSlicedCreateInfoEXT-sliceOffset-07867)VUID-VkImageViewSlicedCreateInfoEXT-sliceOffset-07867  
  `sliceOffset` **must** be less than the effective view depth as specified in [Image Mip Level Sizing](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#resources-image-mip-level-sizing)
- [](#VUID-VkImageViewSlicedCreateInfoEXT-sliceCount-07868)VUID-VkImageViewSlicedCreateInfoEXT-sliceCount-07868  
  If `sliceCount` is not `VK_REMAINING_3D_SLICES_EXT`, it **must** be non-zero and `sliceOffset` + `sliceCount` **must** be less than or equal to the effective view depth as specified in [Image Mip Level Sizing](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#resources-image-mip-level-sizing)
- [](#VUID-VkImageViewSlicedCreateInfoEXT-image-07869)VUID-VkImageViewSlicedCreateInfoEXT-image-07869  
  `image` **must** have been created with `imageType` equal to `VK_IMAGE_TYPE_3D`
- [](#VUID-VkImageViewSlicedCreateInfoEXT-viewType-07909)VUID-VkImageViewSlicedCreateInfoEXT-viewType-07909  
  `viewType` **must** be `VK_IMAGE_VIEW_TYPE_3D`
- [](#VUID-VkImageViewSlicedCreateInfoEXT-None-07870)VUID-VkImageViewSlicedCreateInfoEXT-None-07870  
  The image view **must** reference exactly 1 mip level
- [](#VUID-VkImageViewSlicedCreateInfoEXT-None-07871)VUID-VkImageViewSlicedCreateInfoEXT-None-07871  
  The [imageSlicedViewOf3D](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-imageSlicedViewOf3D) feature **must** be enabled

Valid Usage (Implicit)

- [](#VUID-VkImageViewSlicedCreateInfoEXT-sType-sType)VUID-VkImageViewSlicedCreateInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_IMAGE_VIEW_SLICED_CREATE_INFO_EXT`

## [](#_see_also)See Also

[VK\_EXT\_image\_sliced\_view\_of\_3d](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_image_sliced_view_of_3d.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkImageViewSlicedCreateInfoEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0