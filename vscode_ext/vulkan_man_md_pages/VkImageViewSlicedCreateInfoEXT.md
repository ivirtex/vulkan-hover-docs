# VkImageViewSlicedCreateInfoEXT(3) Manual Page

## Name

VkImageViewSlicedCreateInfoEXT - Specify the subset of 3D slices of an
image view



## <a href="#_c_specification" class="anchor"></a>C Specification

The range of 3D slices for the created image view **can** be restricted
to a subset of the parent image’s Z range by adding a
`VkImageViewSlicedCreateInfoEXT` structure to the `pNext` chain of
[VkImageViewCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewCreateInfo.html).

The `VkImageViewSlicedCreateInfoEXT` structure is defined as:

``` c
// Provided by VK_EXT_image_sliced_view_of_3d
typedef struct VkImageViewSlicedCreateInfoEXT {
    VkStructureType    sType;
    const void*        pNext;
    uint32_t           sliceOffset;
    uint32_t           sliceCount;
} VkImageViewSlicedCreateInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `sliceOffset` is the Z-offset for the first 3D slice accessible to the
  image view.

- `sliceCount` is the number of 3D slices accessible to the image view.

## <a href="#_description" class="anchor"></a>Description

When this structure is chained to
[VkImageViewCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewCreateInfo.html) the `sliceOffset`
field is treated as a Z-offset for the sliced view and `sliceCount`
specifies the range. Shader accesses using a Z coordinate of 0 will
access the depth slice corresponding to `sliceOffset` in the image, and
in a shader, the maximum in-bounds Z coordinate for the view is
`sliceCount` - 1.

A sliced 3D view **must** only be used with a single mip level. The
slice coordinates are integer coordinates within the
`subresourceRange.baseMipLevel` used to create the image view.

The effective view depth is equal to `extent.depth` used to create the
`image` for this view adjusted by `subresourceRange.baseMipLevel` as
specified in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-image-mip-level-sizing"
target="_blank" rel="noopener">Image Mip Level Sizing</a>.

Shader access to this image view is only affected by
`VkImageViewSlicedCreateInfoEXT` if it uses a descriptor of type
`VK_DESCRIPTOR_TYPE_STORAGE_IMAGE`. For access using any other
descriptor type, the contents of `VkImageViewSlicedCreateInfoEXT` are
ignored; instead, `sliceOffset` is treated as being equal to 0, and
`sliceCount` is treated as being equal to `VK_REMAINING_3D_SLICES_EXT`.

Valid Usage

- <a href="#VUID-VkImageViewSlicedCreateInfoEXT-sliceOffset-07867"
  id="VUID-VkImageViewSlicedCreateInfoEXT-sliceOffset-07867"></a>
  VUID-VkImageViewSlicedCreateInfoEXT-sliceOffset-07867  
  `sliceOffset` **must** be less than the effective view depth as
  specified in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-image-mip-level-sizing"
  target="_blank" rel="noopener">Image Mip Level Sizing</a>

- <a href="#VUID-VkImageViewSlicedCreateInfoEXT-sliceCount-07868"
  id="VUID-VkImageViewSlicedCreateInfoEXT-sliceCount-07868"></a>
  VUID-VkImageViewSlicedCreateInfoEXT-sliceCount-07868  
  If `sliceCount` is not `VK_REMAINING_3D_SLICES_EXT`, it **must** be
  non-zero and `sliceOffset` + `sliceCount` **must** be less than or
  equal to the effective view depth as specified in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-image-mip-level-sizing"
  target="_blank" rel="noopener">Image Mip Level Sizing</a>

- <a href="#VUID-VkImageViewSlicedCreateInfoEXT-image-07869"
  id="VUID-VkImageViewSlicedCreateInfoEXT-image-07869"></a>
  VUID-VkImageViewSlicedCreateInfoEXT-image-07869  
  `image` **must** have been created with `imageType` equal to
  `VK_IMAGE_TYPE_3D`

- <a href="#VUID-VkImageViewSlicedCreateInfoEXT-viewType-07909"
  id="VUID-VkImageViewSlicedCreateInfoEXT-viewType-07909"></a>
  VUID-VkImageViewSlicedCreateInfoEXT-viewType-07909  
  `viewType` **must** be `VK_IMAGE_VIEW_TYPE_3D`

- <a href="#VUID-VkImageViewSlicedCreateInfoEXT-None-07870"
  id="VUID-VkImageViewSlicedCreateInfoEXT-None-07870"></a>
  VUID-VkImageViewSlicedCreateInfoEXT-None-07870  
  The image view **must** reference exactly 1 mip level

- <a href="#VUID-VkImageViewSlicedCreateInfoEXT-None-07871"
  id="VUID-VkImageViewSlicedCreateInfoEXT-None-07871"></a>
  VUID-VkImageViewSlicedCreateInfoEXT-None-07871  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-imageSlicedViewOf3D"
  target="_blank" rel="noopener">imageSlicedViewOf3D</a> feature
  **must** be enabled on the device

Valid Usage (Implicit)

- <a href="#VUID-VkImageViewSlicedCreateInfoEXT-sType-sType"
  id="VUID-VkImageViewSlicedCreateInfoEXT-sType-sType"></a>
  VUID-VkImageViewSlicedCreateInfoEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_IMAGE_VIEW_SLICED_CREATE_INFO_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_image_sliced_view_of_3d](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_image_sliced_view_of_3d.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkImageViewSlicedCreateInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
