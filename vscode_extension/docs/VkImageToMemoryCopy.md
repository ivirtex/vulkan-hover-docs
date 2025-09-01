# VkImageToMemoryCopy(3) Manual Page

## Name

VkImageToMemoryCopy - Structure specifying an image to host memory copy operation



## [](#_c_specification)C Specification

Each element of [VkCopyImageToMemoryInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyImageToMemoryInfo.html)::`pRegions` is a structure defined as:

```c++
// Provided by VK_VERSION_1_4
typedef struct VkImageToMemoryCopy {
    VkStructureType             sType;
    const void*                 pNext;
    void*                       pHostPointer;
    uint32_t                    memoryRowLength;
    uint32_t                    memoryImageHeight;
    VkImageSubresourceLayers    imageSubresource;
    VkOffset3D                  imageOffset;
    VkExtent3D                  imageExtent;
} VkImageToMemoryCopy;
```

or the equivalent

```c++
// Provided by VK_EXT_host_image_copy
typedef VkImageToMemoryCopy VkImageToMemoryCopyEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `pHostPointer` is the host memory address which is the destination of the copy.
- `memoryRowLength` and `memoryImageHeight` specify in texels a subregion of a larger two- or three-dimensional image in host memory, and control the addressing calculations. If either of these values is zero, that aspect of the host memory is considered to be tightly packed according to the `imageExtent`.
- `imageSubresource` is a [VkImageSubresourceLayers](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresourceLayers.html) used to specify the specific image subresources of the image used for the source or destination image data.
- `imageOffset` selects the initial `x`, `y`, `z` offsets in texels of the sub-region of the source image data.
- `imageExtent` is the size in texels of the image to copy in `width`, `height` and `depth`.

## [](#_description)Description

This structure is functionally similar to [VkBufferImageCopy2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferImageCopy2.html), except it defines host memory as the target of copy instead of a buffer. In particular, the same data packing rules and restrictions as that structure apply here as well.

Valid Usage

- [](#VUID-VkImageToMemoryCopy-pHostPointer-09066)VUID-VkImageToMemoryCopy-pHostPointer-09066  
  `pHostPointer` **must** point to memory that is large enough to contain all memory locations that are accessed according to [Buffer and Image Addressing](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#copies-buffers-images-addressing), for each element of `pRegions`
- [](#VUID-VkImageToMemoryCopy-pRegions-09067)VUID-VkImageToMemoryCopy-pRegions-09067  
  The union of all source regions, and the union of all destination regions, specified by the elements of `pRegions`, **must** not overlap in memory

<!--THE END-->

- [](#VUID-VkImageToMemoryCopy-memoryRowLength-09101)VUID-VkImageToMemoryCopy-memoryRowLength-09101  
  `memoryRowLength` **must** be `0`, or greater than or equal to the `width` member of `imageExtent`
- [](#VUID-VkImageToMemoryCopy-memoryImageHeight-09102)VUID-VkImageToMemoryCopy-memoryImageHeight-09102  
  `memoryImageHeight` **must** be `0`, or greater than or equal to the `height` member of `imageExtent`
- [](#VUID-VkImageToMemoryCopy-aspectMask-09103)VUID-VkImageToMemoryCopy-aspectMask-09103  
  The `aspectMask` member of `imageSubresource` **must** only have a single bit set
- [](#VUID-VkImageToMemoryCopy-imageExtent-06659)VUID-VkImageToMemoryCopy-imageExtent-06659  
  `imageExtent.width` **must** not be 0
- [](#VUID-VkImageToMemoryCopy-imageExtent-06660)VUID-VkImageToMemoryCopy-imageExtent-06660  
  `imageExtent.height` **must** not be 0
- [](#VUID-VkImageToMemoryCopy-imageExtent-06661)VUID-VkImageToMemoryCopy-imageExtent-06661  
  `imageExtent.depth` **must** not be 0

Valid Usage (Implicit)

- [](#VUID-VkImageToMemoryCopy-sType-sType)VUID-VkImageToMemoryCopy-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_IMAGE_TO_MEMORY_COPY`
- [](#VUID-VkImageToMemoryCopy-pNext-pNext)VUID-VkImageToMemoryCopy-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkImageToMemoryCopy-pHostPointer-parameter)VUID-VkImageToMemoryCopy-pHostPointer-parameter  
  `pHostPointer` **must** be a pointer value
- [](#VUID-VkImageToMemoryCopy-imageSubresource-parameter)VUID-VkImageToMemoryCopy-imageSubresource-parameter  
  `imageSubresource` **must** be a valid [VkImageSubresourceLayers](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresourceLayers.html) structure

## [](#_see_also)See Also

[VK\_EXT\_host\_image\_copy](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_host_image_copy.html), [VK\_VERSION\_1\_4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_4.html), [VkCopyImageToMemoryInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyImageToMemoryInfo.html), [VkExtent3D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtent3D.html), [VkImageSubresourceLayers](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresourceLayers.html), [VkOffset3D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOffset3D.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkImageToMemoryCopy).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0