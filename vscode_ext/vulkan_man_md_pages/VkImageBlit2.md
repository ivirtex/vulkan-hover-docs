# VkImageBlit2(3) Manual Page

## Name

VkImageBlit2 - Structure specifying an image blit operation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkImageBlit2` structure is defined as:

``` c
// Provided by VK_VERSION_1_3
typedef struct VkImageBlit2 {
    VkStructureType             sType;
    const void*                 pNext;
    VkImageSubresourceLayers    srcSubresource;
    VkOffset3D                  srcOffsets[2];
    VkImageSubresourceLayers    dstSubresource;
    VkOffset3D                  dstOffsets[2];
} VkImageBlit2;
```

or the equivalent

``` c
// Provided by VK_KHR_copy_commands2
typedef VkImageBlit2 VkImageBlit2KHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `srcSubresource` is the subresource to blit from.

- `srcOffsets` is a pointer to an array of two
  [VkOffset3D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOffset3D.html) structures specifying the bounds of the
  source region within `srcSubresource`.

- `dstSubresource` is the subresource to blit into.

- `dstOffsets` is a pointer to an array of two
  [VkOffset3D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOffset3D.html) structures specifying the bounds of the
  destination region within `dstSubresource`.

## <a href="#_description" class="anchor"></a>Description

For each element of the `pRegions` array, a blit operation is performed
for the specified source and destination regions.

Valid Usage

- <a href="#VUID-VkImageBlit2-aspectMask-00238"
  id="VUID-VkImageBlit2-aspectMask-00238"></a>
  VUID-VkImageBlit2-aspectMask-00238  
  The `aspectMask` member of `srcSubresource` and `dstSubresource`
  **must** match

- <a href="#VUID-VkImageBlit2-layerCount-08800"
  id="VUID-VkImageBlit2-layerCount-08800"></a>
  VUID-VkImageBlit2-layerCount-08800  
  If neither of the `layerCount` members of `srcSubresource` or
  `dstSubresource` are `VK_REMAINING_ARRAY_LAYERS`, the `layerCount`
  members of `srcSubresource` or `dstSubresource` **must** match

- <a href="#VUID-VkImageBlit2-layerCount-08801"
  id="VUID-VkImageBlit2-layerCount-08801"></a>
  VUID-VkImageBlit2-layerCount-08801  
  If one of the `layerCount` members of `srcSubresource` or
  `dstSubresource` is `VK_REMAINING_ARRAY_LAYERS`, the other member
  **must** be either `VK_REMAINING_ARRAY_LAYERS` or equal to the
  `arrayLayers` member of the
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) used to create the image
  minus `baseArrayLayer`

Valid Usage (Implicit)

- <a href="#VUID-VkImageBlit2-sType-sType"
  id="VUID-VkImageBlit2-sType-sType"></a>
  VUID-VkImageBlit2-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_IMAGE_BLIT_2`

- <a href="#VUID-VkImageBlit2-pNext-pNext"
  id="VUID-VkImageBlit2-pNext-pNext"></a>
  VUID-VkImageBlit2-pNext-pNext  
  `pNext` **must** be `NULL` or a pointer to a valid instance of
  [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyCommandTransformInfoQCOM.html)

- <a href="#VUID-VkImageBlit2-sType-unique"
  id="VUID-VkImageBlit2-sType-unique"></a>
  VUID-VkImageBlit2-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

- <a href="#VUID-VkImageBlit2-srcSubresource-parameter"
  id="VUID-VkImageBlit2-srcSubresource-parameter"></a>
  VUID-VkImageBlit2-srcSubresource-parameter  
  `srcSubresource` **must** be a valid
  [VkImageSubresourceLayers](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresourceLayers.html) structure

- <a href="#VUID-VkImageBlit2-dstSubresource-parameter"
  id="VUID-VkImageBlit2-dstSubresource-parameter"></a>
  VUID-VkImageBlit2-dstSubresource-parameter  
  `dstSubresource` **must** be a valid
  [VkImageSubresourceLayers](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresourceLayers.html) structure

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_copy_commands2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_copy_commands2.html),
[VK_VERSION_1_3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_3.html),
[VkBlitImageInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBlitImageInfo2.html),
[VkImageSubresourceLayers](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresourceLayers.html),
[VkOffset3D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOffset3D.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkImageBlit2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
