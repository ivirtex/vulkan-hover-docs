# VkCopyMemoryToImageIndirectCommandNV(3) Manual Page

## Name

VkCopyMemoryToImageIndirectCommandNV - Structure specifying indirect buffer image copy operation



## [](#_c_specification)C Specification

The `VkCopyMemoryToImageIndirectCommandNV` is defined as:

```c++
// Provided by VK_NV_copy_memory_indirect
typedef struct VkCopyMemoryToImageIndirectCommandNV {
    VkDeviceAddress             srcAddress;
    uint32_t                    bufferRowLength;
    uint32_t                    bufferImageHeight;
    VkImageSubresourceLayers    imageSubresource;
    VkOffset3D                  imageOffset;
    VkExtent3D                  imageExtent;
} VkCopyMemoryToImageIndirectCommandNV;
```

## [](#_members)Members

- `srcAddress` is the starting address of the source device memory to copy from.
- `bufferRowLength` and `bufferImageHeight` specify in texels a subregion of a larger two- or three-dimensional image in buffer memory, and control the addressing calculations. If either of these values is zero, that aspect of the buffer memory is considered to be tightly packed according to the `imageExtent`.
- `imageSubresource` is a [VkImageSubresourceLayers](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresourceLayers.html) used to specify the specific image subresources of the image used for the destination image data, which **must** match the values specified in `pImageSubresources` parameter of [vkCmdCopyMemoryToImageIndirectNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyMemoryToImageIndirectNV.html) during command recording.
- `imageOffset` selects the initial `x`, `y`, `z` offsets in texels of the sub-region of the destination image data.
- `imageExtent` is the size in texels of the destination image in `width`, `height` and `depth`.

## [](#_description)Description

Valid Usage

- [](#VUID-VkCopyMemoryToImageIndirectCommandNV-srcAddress-07678)VUID-VkCopyMemoryToImageIndirectCommandNV-srcAddress-07678  
  The `srcAddress` **must** be 4 byte aligned
- [](#VUID-VkCopyMemoryToImageIndirectCommandNV-bufferRowLength-07679)VUID-VkCopyMemoryToImageIndirectCommandNV-bufferRowLength-07679  
  `bufferRowLength` **must** be `0`, or greater than or equal to the `width` member of `imageExtent`
- [](#VUID-VkCopyMemoryToImageIndirectCommandNV-bufferImageHeight-07680)VUID-VkCopyMemoryToImageIndirectCommandNV-bufferImageHeight-07680  
  `bufferImageHeight` **must** be `0`, or greater than or equal to the `height` member of `imageExtent`
- [](#VUID-VkCopyMemoryToImageIndirectCommandNV-imageOffset-07681)VUID-VkCopyMemoryToImageIndirectCommandNV-imageOffset-07681  
  `imageOffset` **must** specify a valid offset in the destination image
- [](#VUID-VkCopyMemoryToImageIndirectCommandNV-imageExtent-07682)VUID-VkCopyMemoryToImageIndirectCommandNV-imageExtent-07682  
  `imageExtent` **must** specify a valid region in the destination image and can be `0`

Valid Usage (Implicit)

- [](#VUID-VkCopyMemoryToImageIndirectCommandNV-imageSubresource-parameter)VUID-VkCopyMemoryToImageIndirectCommandNV-imageSubresource-parameter  
  `imageSubresource` **must** be a valid [VkImageSubresourceLayers](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresourceLayers.html) structure

## [](#_see_also)See Also

[VK\_NV\_copy\_memory\_indirect](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_copy_memory_indirect.html), [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html), [VkExtent3D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtent3D.html), [VkImageSubresourceLayers](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresourceLayers.html), [VkOffset3D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOffset3D.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkCopyMemoryToImageIndirectCommandNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0