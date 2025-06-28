# VkBufferImageCopy2(3) Manual Page

## Name

VkBufferImageCopy2 - Structure specifying a buffer image copy operation



## [](#_c_specification)C Specification

For both [vkCmdCopyBufferToImage2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyBufferToImage2.html) and [vkCmdCopyImageToBuffer2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyImageToBuffer2.html), each element of `pRegions` is a structure defined as:

```c++
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

```c++
// Provided by VK_KHR_copy_commands2
typedef VkBufferImageCopy2 VkBufferImageCopy2KHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `bufferOffset` is the offset in bytes from the start of the buffer object where the image data is copied from or to.
- `bufferRowLength` and `bufferImageHeight` specify in texels a subregion of a larger two- or three-dimensional image in buffer memory, and control the addressing calculations. If either of these values is zero, that aspect of the buffer memory is considered to be tightly packed according to the `imageExtent`.
- `imageSubresource` is a [VkImageSubresourceLayers](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresourceLayers.html) used to specify the specific image subresources of the image used for the source or destination image data.
- `imageOffset` selects the initial `x`, `y`, `z` offsets in texels of the sub-region of the source or destination image data.
- `imageExtent` is the size in texels of the image to copy in `width`, `height` and `depth`.

## [](#_description)Description

This structure is functionally identical to [VkBufferImageCopy](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferImageCopy.html), but adds `sType` and `pNext` parameters, allowing it to be more easily extended.

Valid Usage

- [](#VUID-VkBufferImageCopy2-bufferRowLength-09101)VUID-VkBufferImageCopy2-bufferRowLength-09101  
  `bufferRowLength` **must** be `0`, or greater than or equal to the `width` member of `imageExtent`
- [](#VUID-VkBufferImageCopy2-bufferImageHeight-09102)VUID-VkBufferImageCopy2-bufferImageHeight-09102  
  `bufferImageHeight` **must** be `0`, or greater than or equal to the `height` member of `imageExtent`
- [](#VUID-VkBufferImageCopy2-aspectMask-09103)VUID-VkBufferImageCopy2-aspectMask-09103  
  The `aspectMask` member of `imageSubresource` **must** only have a single bit set
- [](#VUID-VkBufferImageCopy2-imageExtent-06659)VUID-VkBufferImageCopy2-imageExtent-06659  
  `imageExtent.width` **must** not be 0
- [](#VUID-VkBufferImageCopy2-imageExtent-06660)VUID-VkBufferImageCopy2-imageExtent-06660  
  `imageExtent.height` **must** not be 0
- [](#VUID-VkBufferImageCopy2-imageExtent-06661)VUID-VkBufferImageCopy2-imageExtent-06661  
  `imageExtent.depth` **must** not be 0

Valid Usage (Implicit)

- [](#VUID-VkBufferImageCopy2-sType-sType)VUID-VkBufferImageCopy2-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_BUFFER_IMAGE_COPY_2`
- [](#VUID-VkBufferImageCopy2-pNext-pNext)VUID-VkBufferImageCopy2-pNext-pNext  
  `pNext` **must** be `NULL` or a pointer to a valid instance of [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html)
- [](#VUID-VkBufferImageCopy2-sType-unique)VUID-VkBufferImageCopy2-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique
- [](#VUID-VkBufferImageCopy2-imageSubresource-parameter)VUID-VkBufferImageCopy2-imageSubresource-parameter  
  `imageSubresource` **must** be a valid [VkImageSubresourceLayers](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresourceLayers.html) structure

## [](#_see_also)See Also

[VK\_KHR\_copy\_commands2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_copy_commands2.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkCopyBufferToImageInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyBufferToImageInfo2.html), [VkCopyImageToBufferInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyImageToBufferInfo2.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html), [VkExtent3D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtent3D.html), [VkImageSubresourceLayers](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresourceLayers.html), [VkOffset3D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOffset3D.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkBufferImageCopy2)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0