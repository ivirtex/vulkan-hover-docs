# VkMemoryToImageCopyEXT(3) Manual Page

## Name

VkMemoryToImageCopyEXT - Structure specifying a host memory to image
copy operation



## <a href="#_c_specification" class="anchor"></a>C Specification

Each element of
[VkCopyMemoryToImageInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyMemoryToImageInfoEXT.html)::`pRegions`
is a structure defined as:

``` c
// Provided by VK_EXT_host_image_copy
typedef struct VkMemoryToImageCopyEXT {
    VkStructureType             sType;
    const void*                 pNext;
    const void*                 pHostPointer;
    uint32_t                    memoryRowLength;
    uint32_t                    memoryImageHeight;
    VkImageSubresourceLayers    imageSubresource;
    VkOffset3D                  imageOffset;
    VkExtent3D                  imageExtent;
} VkMemoryToImageCopyEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `pHostPointer` is the host memory address which is the source of the
  copy.

- `memoryRowLength` and `memoryImageHeight` specify in texels a
  subregion of a larger two- or three-dimensional image in host memory,
  and control the addressing calculations. If either of these values is
  zero, that aspect of the host memory is considered to be tightly
  packed according to the `imageExtent`.

- `imageSubresource` is a
  [VkImageSubresourceLayers](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresourceLayers.html) used to
  specify the specific image subresources of the image used for the
  source or destination image data.

- `imageOffset` selects the initial `x`, `y`, `z` offsets in texels of
  the sub-region of the destination image data.

- `imageExtent` is the size in texels of the image to copy in `width`,
  `height` and `depth`.

## <a href="#_description" class="anchor"></a>Description

This structure is functionally similar to
[VkBufferImageCopy2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferImageCopy2.html), except it defines host
memory as the source of copy instead of a buffer. In particular, the
same data packing rules and restrictions as that structure apply here as
well.

Valid Usage

- <a href="#VUID-VkMemoryToImageCopyEXT-pHostPointer-09061"
  id="VUID-VkMemoryToImageCopyEXT-pHostPointer-09061"></a>
  VUID-VkMemoryToImageCopyEXT-pHostPointer-09061  
  `pHostPointer` **must** point to memory that is large enough to
  contain all memory locations that are accessed according to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#copies-buffers-images-addressing"
  target="_blank" rel="noopener">Buffer and Image Addressing</a>, for
  each element of `pRegions`

- <a href="#VUID-VkMemoryToImageCopyEXT-pRegions-09062"
  id="VUID-VkMemoryToImageCopyEXT-pRegions-09062"></a>
  VUID-VkMemoryToImageCopyEXT-pRegions-09062  
  The union of all source regions, and the union of all destination
  regions, specified by the elements of `pRegions`, **must** not overlap
  in memory

<!-- -->

- <a href="#VUID-VkMemoryToImageCopyEXT-memoryRowLength-09101"
  id="VUID-VkMemoryToImageCopyEXT-memoryRowLength-09101"></a>
  VUID-VkMemoryToImageCopyEXT-memoryRowLength-09101  
  `memoryRowLength` **must** be `0`, or greater than or equal to the
  `width` member of `imageExtent`

- <a href="#VUID-VkMemoryToImageCopyEXT-memoryImageHeight-09102"
  id="VUID-VkMemoryToImageCopyEXT-memoryImageHeight-09102"></a>
  VUID-VkMemoryToImageCopyEXT-memoryImageHeight-09102  
  `memoryImageHeight` **must** be `0`, or greater than or equal to the
  `height` member of `imageExtent`

- <a href="#VUID-VkMemoryToImageCopyEXT-aspectMask-09103"
  id="VUID-VkMemoryToImageCopyEXT-aspectMask-09103"></a>
  VUID-VkMemoryToImageCopyEXT-aspectMask-09103  
  The `aspectMask` member of `imageSubresource` **must** only have a
  single bit set

- <a href="#VUID-VkMemoryToImageCopyEXT-imageExtent-06659"
  id="VUID-VkMemoryToImageCopyEXT-imageExtent-06659"></a>
  VUID-VkMemoryToImageCopyEXT-imageExtent-06659  
  `imageExtent.width` **must** not be 0

- <a href="#VUID-VkMemoryToImageCopyEXT-imageExtent-06660"
  id="VUID-VkMemoryToImageCopyEXT-imageExtent-06660"></a>
  VUID-VkMemoryToImageCopyEXT-imageExtent-06660  
  `imageExtent.height` **must** not be 0

- <a href="#VUID-VkMemoryToImageCopyEXT-imageExtent-06661"
  id="VUID-VkMemoryToImageCopyEXT-imageExtent-06661"></a>
  VUID-VkMemoryToImageCopyEXT-imageExtent-06661  
  `imageExtent.depth` **must** not be 0

Valid Usage (Implicit)

- <a href="#VUID-VkMemoryToImageCopyEXT-sType-sType"
  id="VUID-VkMemoryToImageCopyEXT-sType-sType"></a>
  VUID-VkMemoryToImageCopyEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_MEMORY_TO_IMAGE_COPY_EXT`

- <a href="#VUID-VkMemoryToImageCopyEXT-pNext-pNext"
  id="VUID-VkMemoryToImageCopyEXT-pNext-pNext"></a>
  VUID-VkMemoryToImageCopyEXT-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkMemoryToImageCopyEXT-pHostPointer-parameter"
  id="VUID-VkMemoryToImageCopyEXT-pHostPointer-parameter"></a>
  VUID-VkMemoryToImageCopyEXT-pHostPointer-parameter  
  `pHostPointer` **must** be a pointer value

- <a href="#VUID-VkMemoryToImageCopyEXT-imageSubresource-parameter"
  id="VUID-VkMemoryToImageCopyEXT-imageSubresource-parameter"></a>
  VUID-VkMemoryToImageCopyEXT-imageSubresource-parameter  
  `imageSubresource` **must** be a valid
  [VkImageSubresourceLayers](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresourceLayers.html) structure

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_host_image_copy](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_host_image_copy.html),
[VkCopyMemoryToImageInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyMemoryToImageInfoEXT.html),
[VkExtent3D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtent3D.html),
[VkImageSubresourceLayers](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresourceLayers.html),
[VkOffset3D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOffset3D.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkMemoryToImageCopyEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
