# VkCopyMemoryToImageIndirectInfoKHR(3) Manual Page

## Name

VkCopyMemoryToImageIndirectInfoKHR - Parameters describing indirect image copy parameters



## [](#_c_specification)C Specification

The [VkCopyMemoryToImageIndirectInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyMemoryToImageIndirectInfoKHR.html) structure is defined as:

```c++
// Provided by VK_KHR_copy_memory_indirect
typedef struct VkCopyMemoryToImageIndirectInfoKHR {
    VkStructureType                    sType;
    const void*                        pNext;
    VkAddressCopyFlagsKHR              srcCopyFlags;
    uint32_t                           copyCount;
    VkStridedDeviceAddressRangeKHR     copyAddressRange;
    VkImage                            dstImage;
    VkImageLayout                      dstImageLayout;
    const VkImageSubresourceLayers*    pImageSubresources;
} VkCopyMemoryToImageIndirectInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `srcCopyFlags` is a [VkAddressCopyFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAddressCopyFlagsKHR.html) value defining the copy flags for the source address range.
- `copyCount` is the number of copies to execute, and **can** be zero.
- `copyAddressRange` is a memory region specifying the copy parameters. It is laid out as an array of [VkCopyMemoryToImageIndirectCommandKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyMemoryToImageIndirectCommandKHR.html) structures.
- `dstImage` is the destination image.
- `dstImageLayout` is the layout of the destination image subresources for the copy.
- `pImageSubresources` is a pointer to an array of `copyCount` [VkImageSubresourceLayers](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresourceLayers.html) structures, specifying the image subresources of the destination image data for the copy operation.

## [](#_description)Description

Valid Usage

- [](#VUID-VkCopyMemoryToImageIndirectInfoKHR-srcCopyFlags-10950)VUID-VkCopyMemoryToImageIndirectInfoKHR-srcCopyFlags-10950  
  If `srcCopyFlags` contains `VK_ADDRESS_COPY_SPARSE_BIT_KHR`, the source memory regions accessed **must** be fully-resident
- [](#VUID-VkCopyMemoryToImageIndirectInfoKHR-copyCount-10951)VUID-VkCopyMemoryToImageIndirectInfoKHR-copyCount-10951  
  `copyCount` **must** be less than or equal to `copyAddressRange.pname`:size / `copyAddressRange.pname`:stride
- [](#VUID-VkCopyMemoryToImageIndirectInfoKHR-copyAddressRange-10952)VUID-VkCopyMemoryToImageIndirectInfoKHR-copyAddressRange-10952  
  `copyAddressRange.pname`:address **must** be 4 byte aligned
- [](#VUID-VkCopyMemoryToImageIndirectInfoKHR-copyAddressRange-10953)VUID-VkCopyMemoryToImageIndirectInfoKHR-copyAddressRange-10953  
  `copyAddressRange.pname`:stride **must** be a multiple of `4` and **must** be greater than or equal to sizeof([VkCopyMemoryToImageIndirectCommandKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyMemoryToImageIndirectCommandKHR.html))
- [](#VUID-VkCopyMemoryToImageIndirectInfoKHR-copyAddressRange-10954)VUID-VkCopyMemoryToImageIndirectInfoKHR-copyAddressRange-10954  
  Any of the source or destination memory regions specified in `copyAddressRange` **must** not overlap with any of the specified destination memory regions
- [](#VUID-VkCopyMemoryToImageIndirectInfoKHR-dstImage-10955)VUID-VkCopyMemoryToImageIndirectInfoKHR-dstImage-10955  
  The format features of `dstImage` **must** contain `VK_FORMAT_FEATURE_2_COPY_IMAGE_INDIRECT_DST_BIT_KHR`

<!--THE END-->

- [](#VUID-VkCopyMemoryToImageIndirectInfoKHR-dstImage-07661)VUID-VkCopyMemoryToImageIndirectInfoKHR-dstImage-07661  
  `dstImage` **must** not be a protected image
- [](#VUID-VkCopyMemoryToImageIndirectInfoKHR-aspectMask-07662)VUID-VkCopyMemoryToImageIndirectInfoKHR-aspectMask-07662  
  The `aspectMask` member for every subresource in `pImageSubresources` **must** only have a single bit set
- [](#VUID-VkCopyMemoryToImageIndirectInfoKHR-dstImage-07663)VUID-VkCopyMemoryToImageIndirectInfoKHR-dstImage-07663  
  The image region specified by each element in `copyBufferAddress` **must** be a region that is contained within `dstImage`
- [](#VUID-VkCopyMemoryToImageIndirectInfoKHR-dstImage-07664)VUID-VkCopyMemoryToImageIndirectInfoKHR-dstImage-07664  
  `dstImage` **must** have been created with `VK_IMAGE_USAGE_TRANSFER_DST_BIT` usage flag
- [](#VUID-VkCopyMemoryToImageIndirectInfoKHR-dstImage-07665)VUID-VkCopyMemoryToImageIndirectInfoKHR-dstImage-07665  
  If `dstImage` is non-sparse then it **must** be bound completely and contiguously to a single `VkDeviceMemory` object

<!--THE END-->

- [](#VUID-VkCopyMemoryToImageIndirectInfoKHR-dstImage-07973)VUID-VkCopyMemoryToImageIndirectInfoKHR-dstImage-07973  
  `dstImage` **must** have a sample count equal to `VK_SAMPLE_COUNT_1_BIT`
- [](#VUID-VkCopyMemoryToImageIndirectInfoKHR-dstImageLayout-07667)VUID-VkCopyMemoryToImageIndirectInfoKHR-dstImageLayout-07667  
  `dstImageLayout` **must** specify the layout of the image subresources of `dstImage` at the time this command is executed on a `VkDevice`
- [](#VUID-VkCopyMemoryToImageIndirectInfoKHR-dstImageLayout-07669)VUID-VkCopyMemoryToImageIndirectInfoKHR-dstImageLayout-07669  
  `dstImageLayout` **must** be `VK_IMAGE_LAYOUT_TRANSFER_DST_OPTIMAL`, `VK_IMAGE_LAYOUT_SHARED_PRESENT_KHR`, or `VK_IMAGE_LAYOUT_GENERAL`
- [](#VUID-VkCopyMemoryToImageIndirectInfoKHR-mipLevel-07670)VUID-VkCopyMemoryToImageIndirectInfoKHR-mipLevel-07670  
  The specified `mipLevel` of each region in `pImageSubresources` **must** be less than the `mipLevels` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) when `dstImage` was created
- [](#VUID-VkCopyMemoryToImageIndirectInfoKHR-layerCount-08764)VUID-VkCopyMemoryToImageIndirectInfoKHR-layerCount-08764  
  If the specified `layerCount` of each region in `pImageSubresources` is not `VK_REMAINING_ARRAY_LAYERS`, the specified `baseArrayLayer` + `layerCount` of each region in `pImageSubresources` **must** be less than or equal to the `arrayLayers` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) when `dstImage` was created
- [](#VUID-VkCopyMemoryToImageIndirectInfoKHR-dstImage-07673)VUID-VkCopyMemoryToImageIndirectInfoKHR-dstImage-07673  
  `dstImage` **must** not have been created with `flags` containing `VK_IMAGE_CREATE_SUBSAMPLED_BIT_EXT`
- [](#VUID-VkCopyMemoryToImageIndirectInfoKHR-commandBuffer-07674)VUID-VkCopyMemoryToImageIndirectInfoKHR-commandBuffer-07674  
  If the queue family used to create the [VkCommandPool](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandPool.html) which `commandBuffer` was allocated from does not support `VK_QUEUE_GRAPHICS_BIT`, for each region, the `aspectMask` member of `pImageSubresources` **must** not be `VK_IMAGE_ASPECT_DEPTH_BIT` or `VK_IMAGE_ASPECT_STENCIL_BIT`
- [](#VUID-VkCopyMemoryToImageIndirectInfoKHR-dstImage-10974)VUID-VkCopyMemoryToImageIndirectInfoKHR-dstImage-10974  
  The format features of `dstImage` **must** contain `VK_FORMAT_FEATURE_TRANSFER_DST_BIT`
- [](#VUID-VkCopyMemoryToImageIndirectInfoKHR-copyBufferAddress-10975)VUID-VkCopyMemoryToImageIndirectInfoKHR-copyBufferAddress-10975  
  Any of the source or destination memory regions specified in `copyBufferAddress` **must** not overlap with any of the specified destination memory regions

Valid Usage (Implicit)

- [](#VUID-VkCopyMemoryToImageIndirectInfoKHR-sType-sType)VUID-VkCopyMemoryToImageIndirectInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_COPY_MEMORY_TO_IMAGE_INDIRECT_INFO_KHR`
- [](#VUID-VkCopyMemoryToImageIndirectInfoKHR-pNext-pNext)VUID-VkCopyMemoryToImageIndirectInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkCopyMemoryToImageIndirectInfoKHR-srcCopyFlags-parameter)VUID-VkCopyMemoryToImageIndirectInfoKHR-srcCopyFlags-parameter  
  `srcCopyFlags` **must** be a valid combination of [VkAddressCopyFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAddressCopyFlagBitsKHR.html) values
- [](#VUID-VkCopyMemoryToImageIndirectInfoKHR-srcCopyFlags-requiredbitmask)VUID-VkCopyMemoryToImageIndirectInfoKHR-srcCopyFlags-requiredbitmask  
  `srcCopyFlags` **must** not be `0`
- [](#VUID-VkCopyMemoryToImageIndirectInfoKHR-dstImage-parameter)VUID-VkCopyMemoryToImageIndirectInfoKHR-dstImage-parameter  
  `dstImage` **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html) handle
- [](#VUID-VkCopyMemoryToImageIndirectInfoKHR-dstImageLayout-parameter)VUID-VkCopyMemoryToImageIndirectInfoKHR-dstImageLayout-parameter  
  `dstImageLayout` **must** be a valid [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html) value
- [](#VUID-VkCopyMemoryToImageIndirectInfoKHR-pImageSubresources-parameter)VUID-VkCopyMemoryToImageIndirectInfoKHR-pImageSubresources-parameter  
  `pImageSubresources` **must** be a valid pointer to an array of `copyCount` valid [VkImageSubresourceLayers](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresourceLayers.html) structures
- [](#VUID-VkCopyMemoryToImageIndirectInfoKHR-copyCount-arraylength)VUID-VkCopyMemoryToImageIndirectInfoKHR-copyCount-arraylength  
  `copyCount` **must** be greater than `0`

## [](#_see_also)See Also

[VK\_KHR\_copy\_memory\_indirect](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_copy_memory_indirect.html), [VkAddressCopyFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAddressCopyFlagsKHR.html), [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html), [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html), [VkImageSubresourceLayers](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresourceLayers.html), [VkStridedDeviceAddressRangeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStridedDeviceAddressRangeKHR.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCmdCopyMemoryToImageIndirectKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyMemoryToImageIndirectKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkCopyMemoryToImageIndirectInfoKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0