# VkBufferImageCopy2(3) Manual Page

## Name

VkBufferImageCopy2 - Structure specifying a buffer image copy operation



## <a href="#_c_specification" class="anchor"></a>C Specification

For both [vkCmdCopyBufferToImage2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyBufferToImage2.html) and
[vkCmdCopyImageToBuffer2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyImageToBuffer2.html), each element of
`pRegions` is a structure defined as:

``` c
// Provided by VK_VERSION_1_3
typedef struct VkBufferImageCopy2 {
    VkStructureType             sType;
    const void*                 pNext;
    VkDeviceSize                bufferOffset;
    uint32_t                    bufferRowLength;
    uint32_t                    bufferImageHeight;
    VkImageSubresourceLayers    imageSubresource;
    VkOffset3D                  imageOffset;
    VkExtent3D                  imageExtent;
} VkBufferImageCopy2;
```

or the equivalent

``` c
// Provided by VK_KHR_copy_commands2
typedef VkBufferImageCopy2 VkBufferImageCopy2KHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `bufferOffset` is the offset in bytes from the start of the buffer
  object where the image data is copied from or to.

- `bufferRowLength` and `bufferImageHeight` specify in texels a
  subregion of a larger two- or three-dimensional image in buffer
  memory, and control the addressing calculations. If either of these
  values is zero, that aspect of the buffer memory is considered to be
  tightly packed according to the `imageExtent`.

- `imageSubresource` is a
  [VkImageSubresourceLayers](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresourceLayers.html) used to
  specify the specific image subresources of the image used for the
  source or destination image data.

- `imageOffset` selects the initial `x`, `y`, `z` offsets in texels of
  the sub-region of the source or destination image data.

- `imageExtent` is the size in texels of the image to copy in `width`,
  `height` and `depth`.

## <a href="#_description" class="anchor"></a>Description

This structure is functionally identical to
[VkBufferImageCopy](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferImageCopy.html), but adds `sType` and
`pNext` parameters, allowing it to be more easily extended.

Valid Usage

- <a href="#VUID-VkBufferImageCopy2-bufferRowLength-09101"
  id="VUID-VkBufferImageCopy2-bufferRowLength-09101"></a>
  VUID-VkBufferImageCopy2-bufferRowLength-09101  
  `bufferRowLength` **must** be `0`, or greater than or equal to the
  `width` member of `imageExtent`

- <a href="#VUID-VkBufferImageCopy2-bufferImageHeight-09102"
  id="VUID-VkBufferImageCopy2-bufferImageHeight-09102"></a>
  VUID-VkBufferImageCopy2-bufferImageHeight-09102  
  `bufferImageHeight` **must** be `0`, or greater than or equal to the
  `height` member of `imageExtent`

- <a href="#VUID-VkBufferImageCopy2-aspectMask-09103"
  id="VUID-VkBufferImageCopy2-aspectMask-09103"></a>
  VUID-VkBufferImageCopy2-aspectMask-09103  
  The `aspectMask` member of `imageSubresource` **must** only have a
  single bit set

- <a href="#VUID-VkBufferImageCopy2-imageExtent-06659"
  id="VUID-VkBufferImageCopy2-imageExtent-06659"></a>
  VUID-VkBufferImageCopy2-imageExtent-06659  
  `imageExtent.width` **must** not be 0

- <a href="#VUID-VkBufferImageCopy2-imageExtent-06660"
  id="VUID-VkBufferImageCopy2-imageExtent-06660"></a>
  VUID-VkBufferImageCopy2-imageExtent-06660  
  `imageExtent.height` **must** not be 0

- <a href="#VUID-VkBufferImageCopy2-imageExtent-06661"
  id="VUID-VkBufferImageCopy2-imageExtent-06661"></a>
  VUID-VkBufferImageCopy2-imageExtent-06661  
  `imageExtent.depth` **must** not be 0

Valid Usage (Implicit)

- <a href="#VUID-VkBufferImageCopy2-sType-sType"
  id="VUID-VkBufferImageCopy2-sType-sType"></a>
  VUID-VkBufferImageCopy2-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_BUFFER_IMAGE_COPY_2`

- <a href="#VUID-VkBufferImageCopy2-pNext-pNext"
  id="VUID-VkBufferImageCopy2-pNext-pNext"></a>
  VUID-VkBufferImageCopy2-pNext-pNext  
  `pNext` **must** be `NULL` or a pointer to a valid instance of
  [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyCommandTransformInfoQCOM.html)

- <a href="#VUID-VkBufferImageCopy2-sType-unique"
  id="VUID-VkBufferImageCopy2-sType-unique"></a>
  VUID-VkBufferImageCopy2-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

- <a href="#VUID-VkBufferImageCopy2-imageSubresource-parameter"
  id="VUID-VkBufferImageCopy2-imageSubresource-parameter"></a>
  VUID-VkBufferImageCopy2-imageSubresource-parameter  
  `imageSubresource` **must** be a valid
  [VkImageSubresourceLayers](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresourceLayers.html) structure

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_copy_commands2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_copy_commands2.html),
[VK_VERSION_1_3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_3.html),
[VkCopyBufferToImageInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyBufferToImageInfo2.html),
[VkCopyImageToBufferInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyImageToBufferInfo2.html),
[VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html), [VkExtent3D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtent3D.html),
[VkImageSubresourceLayers](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresourceLayers.html),
[VkOffset3D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOffset3D.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkBufferImageCopy2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
