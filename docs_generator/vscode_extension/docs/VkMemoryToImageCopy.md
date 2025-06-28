# VkMemoryToImageCopy(3) Manual Page

## Name

VkMemoryToImageCopy - Structure specifying a host memory to image copy operation



## [](#_c_specification)C Specification

Each element of [VkCopyMemoryToImageInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyMemoryToImageInfo.html)::`pRegions` is a structure defined as:

```c++
// Provided by VK_VERSION_1_4
typedef struct VkMemoryToImageCopy {
    VkStructureType             sType;
    const void*                 pNext;
    const void*                 pHostPointer;
    uint32_t                    memoryRowLength;
    uint32_t                    memoryImageHeight;
    VkImageSubresourceLayers    imageSubresource;
    VkOffset3D                  imageOffset;
    VkExtent3D                  imageExtent;
} VkMemoryToImageCopy;
```

or the equivalent

```c++
// Provided by VK_EXT_host_image_copy
typedef VkMemoryToImageCopy VkMemoryToImageCopyEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `pHostPointer` is the host memory address which is the source of the copy.
- `memoryRowLength` and `memoryImageHeight` specify in texels a subregion of a larger two- or three-dimensional image in host memory, and control the addressing calculations. If either of these values is zero, that aspect of the host memory is considered to be tightly packed according to the `imageExtent`.
- `imageSubresource` is a [VkImageSubresourceLayers](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresourceLayers.html) used to specify the specific image subresources of the image used for the source or destination image data.
- `imageOffset` selects the initial `x`, `y`, `z` offsets in texels of the sub-region of the destination image data.
- `imageExtent` is the size in texels of the image to copy in `width`, `height` and `depth`.

## [](#_description)Description

This structure is functionally similar to [VkBufferImageCopy2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferImageCopy2.html), except it defines host memory as the source of copy instead of a buffer. In particular, the same data packing rules and restrictions as that structure apply here as well.

Valid Usage

- [](#VUID-VkMemoryToImageCopy-pHostPointer-09061)VUID-VkMemoryToImageCopy-pHostPointer-09061  
  `pHostPointer` **must** point to memory that is large enough to contain all memory locations that are accessed according to [Buffer and Image Addressing](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#copies-buffers-images-addressing), for each element of `pRegions`
- [](#VUID-VkMemoryToImageCopy-pRegions-09062)VUID-VkMemoryToImageCopy-pRegions-09062  
  The union of all source regions, and the union of all destination regions, specified by the elements of `pRegions`, **must** not overlap in memory

<!--THE END-->

- [](#VUID-VkMemoryToImageCopy-memoryRowLength-09101)VUID-VkMemoryToImageCopy-memoryRowLength-09101  
  `memoryRowLength` **must** be `0`, or greater than or equal to the `width` member of `imageExtent`
- [](#VUID-VkMemoryToImageCopy-memoryImageHeight-09102)VUID-VkMemoryToImageCopy-memoryImageHeight-09102  
  `memoryImageHeight` **must** be `0`, or greater than or equal to the `height` member of `imageExtent`
- [](#VUID-VkMemoryToImageCopy-aspectMask-09103)VUID-VkMemoryToImageCopy-aspectMask-09103  
  The `aspectMask` member of `imageSubresource` **must** only have a single bit set
- [](#VUID-VkMemoryToImageCopy-imageExtent-06659)VUID-VkMemoryToImageCopy-imageExtent-06659  
  `imageExtent.width` **must** not be 0
- [](#VUID-VkMemoryToImageCopy-imageExtent-06660)VUID-VkMemoryToImageCopy-imageExtent-06660  
  `imageExtent.height` **must** not be 0
- [](#VUID-VkMemoryToImageCopy-imageExtent-06661)VUID-VkMemoryToImageCopy-imageExtent-06661  
  `imageExtent.depth` **must** not be 0

Valid Usage (Implicit)

- [](#VUID-VkMemoryToImageCopy-sType-sType)VUID-VkMemoryToImageCopy-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_MEMORY_TO_IMAGE_COPY`
- [](#VUID-VkMemoryToImageCopy-pNext-pNext)VUID-VkMemoryToImageCopy-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkMemoryToImageCopy-pHostPointer-parameter)VUID-VkMemoryToImageCopy-pHostPointer-parameter  
  `pHostPointer` **must** be a pointer value
- [](#VUID-VkMemoryToImageCopy-imageSubresource-parameter)VUID-VkMemoryToImageCopy-imageSubresource-parameter  
  `imageSubresource` **must** be a valid [VkImageSubresourceLayers](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresourceLayers.html) structure

## [](#_see_also)See Also

[VK\_EXT\_host\_image\_copy](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_host_image_copy.html), [VK\_VERSION\_1\_4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_4.html), [VkCopyMemoryToImageInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyMemoryToImageInfo.html), [VkExtent3D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtent3D.html), [VkImageSubresourceLayers](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresourceLayers.html), [VkOffset3D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOffset3D.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkMemoryToImageCopy)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0