# VkImageViewMinLodCreateInfoEXT(3) Manual Page

## Name

VkImageViewMinLodCreateInfoEXT - Structure describing the minimum LOD of
an image view



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkImageViewMinLodCreateInfoEXT` structure is defined as:

``` c
// Provided by VK_EXT_image_view_min_lod
typedef struct VkImageViewMinLodCreateInfoEXT {
    VkStructureType    sType;
    const void*        pNext;
    float              minLod;
} VkImageViewMinLodCreateInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `minLod` is the value to clamp the minimum LOD accessible by this
  [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html).

## <a href="#_description" class="anchor"></a>Description

If the `pNext` chain includes a `VkImageViewMinLodCreateInfoEXT`
structure, then that structure includes a parameter specifying a value
to clamp the minimum LOD value during <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#textures-image-level-selection"
target="_blank" rel="noopener">Image Level(s) Selection</a>, <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#textures-gather"
target="_blank" rel="noopener">Texel Gathering</a> and <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#textures-integer-coordinate-operations"
target="_blank" rel="noopener">Integer Texel Coordinate Operations</a>.

If the image view contains `VkImageViewMinLodCreateInfoEXT` and it is
used as part of a sampling operation:

minLodFloat<sub>imageView</sub> = `minLod`

otherwise:

minLodFloat<sub>imageView</sub> = 0.0

An integer variant of this parameter is also defined for sampling
operations which access integer mipmap levels:

minLodInteger<sub>imageView</sub> = ⌊minLodFloat<sub>imageView</sub>⌋

Valid Usage

- <a href="#VUID-VkImageViewMinLodCreateInfoEXT-minLod-06455"
  id="VUID-VkImageViewMinLodCreateInfoEXT-minLod-06455"></a>
  VUID-VkImageViewMinLodCreateInfoEXT-minLod-06455  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-minLod"
  target="_blank" rel="noopener"><code>minLod</code></a> feature is not
  enabled, `minLod` **must** be `0.0`

- <a href="#VUID-VkImageViewMinLodCreateInfoEXT-minLod-06456"
  id="VUID-VkImageViewMinLodCreateInfoEXT-minLod-06456"></a>
  VUID-VkImageViewMinLodCreateInfoEXT-minLod-06456  
  `minLod` **must** be less or equal to the index of the last mipmap
  level accessible to the view

Valid Usage (Implicit)

- <a href="#VUID-VkImageViewMinLodCreateInfoEXT-sType-sType"
  id="VUID-VkImageViewMinLodCreateInfoEXT-sType-sType"></a>
  VUID-VkImageViewMinLodCreateInfoEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_IMAGE_VIEW_MIN_LOD_CREATE_INFO_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_image_view_min_lod](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_image_view_min_lod.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkImageViewMinLodCreateInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
