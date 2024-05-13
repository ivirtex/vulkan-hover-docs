# VkImageResolve2(3) Manual Page

## Name

VkImageResolve2 - Structure specifying an image resolve operation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkImageResolve2` structure is defined as:

``` c
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

``` c
// Provided by VK_KHR_copy_commands2
typedef VkImageResolve2 VkImageResolve2KHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `srcSubresource` and `dstSubresource` are
  [VkImageSubresourceLayers](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresourceLayers.html) structures
  specifying the image subresources of the images used for the source
  and destination image data, respectively. Resolve of depth/stencil
  images is not supported.

- `srcOffset` and `dstOffset` select the initial `x`, `y`, and `z`
  offsets in texels of the sub-regions of the source and destination
  image data.

- `extent` is the size in texels of the source image to resolve in
  `width`, `height` and `depth`.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkImageResolve2-aspectMask-00266"
  id="VUID-VkImageResolve2-aspectMask-00266"></a>
  VUID-VkImageResolve2-aspectMask-00266  
  The `aspectMask` member of `srcSubresource` and `dstSubresource`
  **must** only contain `VK_IMAGE_ASPECT_COLOR_BIT`

- <a href="#VUID-VkImageResolve2-layerCount-08803"
  id="VUID-VkImageResolve2-layerCount-08803"></a>
  VUID-VkImageResolve2-layerCount-08803  
  If neither of the `layerCount` members of `srcSubresource` or
  `dstSubresource` are `VK_REMAINING_ARRAY_LAYERS`, the `layerCount`
  member of `srcSubresource` and `dstSubresource` **must** match

- <a href="#VUID-VkImageResolve2-layerCount-08804"
  id="VUID-VkImageResolve2-layerCount-08804"></a>
  VUID-VkImageResolve2-layerCount-08804  
  If one of the `layerCount` members of `srcSubresource` or
  `dstSubresource` is `VK_REMAINING_ARRAY_LAYERS`, the other member
  **must** be either `VK_REMAINING_ARRAY_LAYERS` or equal to the
  `arrayLayers` member of the
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) used to create the image
  minus `baseArrayLayer`

Valid Usage (Implicit)

- <a href="#VUID-VkImageResolve2-sType-sType"
  id="VUID-VkImageResolve2-sType-sType"></a>
  VUID-VkImageResolve2-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_IMAGE_RESOLVE_2`

- <a href="#VUID-VkImageResolve2-pNext-pNext"
  id="VUID-VkImageResolve2-pNext-pNext"></a>
  VUID-VkImageResolve2-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkImageResolve2-srcSubresource-parameter"
  id="VUID-VkImageResolve2-srcSubresource-parameter"></a>
  VUID-VkImageResolve2-srcSubresource-parameter  
  `srcSubresource` **must** be a valid
  [VkImageSubresourceLayers](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresourceLayers.html) structure

- <a href="#VUID-VkImageResolve2-dstSubresource-parameter"
  id="VUID-VkImageResolve2-dstSubresource-parameter"></a>
  VUID-VkImageResolve2-dstSubresource-parameter  
  `dstSubresource` **must** be a valid
  [VkImageSubresourceLayers](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresourceLayers.html) structure

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_copy_commands2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_copy_commands2.html),
[VK_VERSION_1_3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_3.html), [VkExtent3D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtent3D.html),
[VkImageSubresourceLayers](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresourceLayers.html),
[VkOffset3D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOffset3D.html),
[VkResolveImageInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResolveImageInfo2.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkImageResolve2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
