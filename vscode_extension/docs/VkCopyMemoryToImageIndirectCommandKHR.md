# VkCopyMemoryToImageIndirectCommandKHR(3) Manual Page

## Name

VkCopyMemoryToImageIndirectCommandKHR - Structure specifying indirect memory region to image copy operation



## [](#_c_specification)C Specification

The structure describing source and destination memory regions, `VkCopyMemoryToImageIndirectCommandKHR` is defined as:

```c++
// Provided by VK_KHR_copy_memory_indirect
typedef struct VkCopyMemoryToImageIndirectCommandKHR {
    VkDeviceAddress             srcAddress;
    uint32_t                    bufferRowLength;
    uint32_t                    bufferImageHeight;
    VkImageSubresourceLayers    imageSubresource;
    VkOffset3D                  imageOffset;
    VkExtent3D                  imageExtent;
} VkCopyMemoryToImageIndirectCommandKHR;
```

or the equivalent

```c++
// Provided by VK_NV_copy_memory_indirect
typedef VkCopyMemoryToImageIndirectCommandKHR VkCopyMemoryToImageIndirectCommandNV;
```

## [](#_members)Members

- `srcAddress` is the starting address of the source device memory to copy from.
- `bufferRowLength` and `bufferImageHeight` specify in texels a subregion of a larger two- or three-dimensional image in buffer memory, and control the addressing calculations. If either of these values is zero, that aspect of the buffer memory is considered to be tightly packed according to the `imageExtent`.
- `imageSubresource` is a [VkImageSubresourceLayers](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresourceLayers.html) structure used to specify the specific image subresources of the image used for the destination image data, which **must** match the value specified in corresponding index of the `pCopyMemoryToImageIndirectInfo->pImageSubresources` array of [vkCmdCopyMemoryToImageIndirectKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyMemoryToImageIndirectKHR.html) during command recording.
- `imageOffset` selects the initial `x`, `y`, `z` offsets in texels of the sub-region of the destination image data.
- `imageExtent` is the size in texels of the destination image in `width`, `height` and `depth`.

## [](#_description)Description

Valid Usage

- [](#VUID-VkCopyMemoryToImageIndirectCommandKHR-srcAddress-10963)VUID-VkCopyMemoryToImageIndirectCommandKHR-srcAddress-10963  
  The `srcAddress` **must** be 4 byte aligned
- [](#VUID-VkCopyMemoryToImageIndirectCommandKHR-bufferRowLength-10964)VUID-VkCopyMemoryToImageIndirectCommandKHR-bufferRowLength-10964  
  `bufferRowLength` **must** be `0`, or greater than or equal to the `width` member of `imageExtent`
- [](#VUID-VkCopyMemoryToImageIndirectCommandKHR-bufferImageHeight-10965)VUID-VkCopyMemoryToImageIndirectCommandKHR-bufferImageHeight-10965  
  `bufferImageHeight` **must** be `0`, or greater than or equal to the `height` member of `imageExtent`
- [](#VUID-VkCopyMemoryToImageIndirectCommandKHR-imageOffset-10966)VUID-VkCopyMemoryToImageIndirectCommandKHR-imageOffset-10966  
  `imageOffset` **must** specify a valid offset in the destination image
- [](#VUID-VkCopyMemoryToImageIndirectCommandKHR-imageExtent-10967)VUID-VkCopyMemoryToImageIndirectCommandKHR-imageExtent-10967  
  `imageExtent` **must** specify a valid region in the destination image and **can** be `0`
- [](#VUID-VkCopyMemoryToImageIndirectCommandKHR-srcAddress-10968)VUID-VkCopyMemoryToImageIndirectCommandKHR-srcAddress-10968  
  The memory region starting at `srcAddress` and described by `bufferRowLength` and `bufferImageHeight` **must** not exceed the bounds of the memory allocation backing memory at `srcAddress`
- [](#VUID-VkCopyMemoryToImageIndirectCommandKHR-imageOffset-10969)VUID-VkCopyMemoryToImageIndirectCommandKHR-imageOffset-10969  
  The `imageOffset` and `imageExtent` members of each region **must** respect the image transfer granularity requirements of `commandBuffer`’s command pool’s queue family, as described in [VkQueueFamilyProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueFamilyProperties.html)
- [](#VUID-VkCopyMemoryToImageIndirectCommandKHR-imageOffset-10970)VUID-VkCopyMemoryToImageIndirectCommandKHR-imageOffset-10970  
  For each destination region, `imageOffset.x` and (`imageExtent.width` + `imageOffset.x`) **must** both be greater than or equal to `0` and less than or equal to the width of the specified subresource
- [](#VUID-VkCopyMemoryToImageIndirectCommandKHR-imageOffset-10971)VUID-VkCopyMemoryToImageIndirectCommandKHR-imageOffset-10971  
  For each destination region, `imageOffset.y` and (`imageExtent.height` + `imageOffset.y`) **must** both be greater than or equal to `0` and less than or equal to the height of the specified subresource
- [](#VUID-VkCopyMemoryToImageIndirectCommandKHR-dstImage-10972)VUID-VkCopyMemoryToImageIndirectCommandKHR-dstImage-10972  
  If `dstImage` is of type `VK_IMAGE_TYPE_3D`, for each destination region, `imageSubresource.baseArrayLayer` and (`imageSubresource.baseArrayLayer` + `imageSubresource.layerCount`) **must** both be greater than or equal to `0` and less than or equal to the depth of the specified subresource
- [](#VUID-VkCopyMemoryToImageIndirectCommandKHR-dstImage-10973)VUID-VkCopyMemoryToImageIndirectCommandKHR-dstImage-10973  
  If `dstImage` is of type `VK_IMAGE_TYPE_3D`, `imageOffset.z` and `imageExtent.depth` are ignored as base slice and number of slices to copy are read from `baseArrayLayer` and `layerCount` from [VkImageSubresourceLayers](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresourceLayers.html) in [VkCopyMemoryToImageIndirectInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyMemoryToImageIndirectInfoKHR.html) at the time of command recording

Valid Usage (Implicit)

- [](#VUID-VkCopyMemoryToImageIndirectCommandKHR-srcAddress-parameter)VUID-VkCopyMemoryToImageIndirectCommandKHR-srcAddress-parameter  
  `srcAddress` **must** be a valid [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html) value
- [](#VUID-VkCopyMemoryToImageIndirectCommandKHR-imageSubresource-parameter)VUID-VkCopyMemoryToImageIndirectCommandKHR-imageSubresource-parameter  
  `imageSubresource` **must** be a valid [VkImageSubresourceLayers](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresourceLayers.html) structure

## [](#_see_also)See Also

[VK\_KHR\_copy\_memory\_indirect](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_copy_memory_indirect.html), [VK\_NV\_copy\_memory\_indirect](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_copy_memory_indirect.html), [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html), [VkExtent3D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtent3D.html), [VkImageSubresourceLayers](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresourceLayers.html), [VkOffset3D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOffset3D.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkCopyMemoryToImageIndirectCommandKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0