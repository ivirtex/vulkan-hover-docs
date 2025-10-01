# vkCmdCopyMemoryToImageIndirectNV(3) Manual Page

## Name

vkCmdCopyMemoryToImageIndirectNV - Copy data from a memory region to an image object



## [](#_c_specification)C Specification

To copy data from a memory region to an image object by specifying copy parameters in memory, call:

```c++
// Provided by VK_NV_copy_memory_indirect
void vkCmdCopyMemoryToImageIndirectNV(
    VkCommandBuffer                             commandBuffer,
    VkDeviceAddress                             copyBufferAddress,
    uint32_t                                    copyCount,
    uint32_t                                    stride,
    VkImage                                     dstImage,
    VkImageLayout                               dstImageLayout,
    const VkImageSubresourceLayers*             pImageSubresources);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `copyBufferAddress` is the address specifying the copy parameters which are laid out in memory as an array of [VkCopyMemoryToImageIndirectCommandNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyMemoryToImageIndirectCommandNV.html) structures.
- `copyCount` is the number of copies to execute, and **can** be zero.
- `stride` is the byte stride between successive sets of copy parameters.
- `dstImage` is the destination image.
- `dstImageLayout` is the layout of the destination image subresources for the copy.
- `pImageSubresources` is a pointer to an array of `copyCount` [VkImageSubresourceLayers](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresourceLayers.html) structures, specifying the image subresources of the destination image data for the copy operation.

## [](#_description)Description

Each region in `copyBufferAddress` is copied from the source memory region to an image region in the destination image. If the destination image is of type `VK_IMAGE_TYPE_3D`, the starting slice and number of slices to copy are specified in `pImageSubresources->baseArrayLayer` and `pImageSubresources->layerCount` respectively.

Valid Usage

- [](#VUID-vkCmdCopyMemoryToImageIndirectNV-None-07660)VUID-vkCmdCopyMemoryToImageIndirectNV-None-07660  
  The [`indirectCopy`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-indirectCopy) feature **must** be enabled
- [](#VUID-vkCmdCopyMemoryToImageIndirectNV-offset-07676)VUID-vkCmdCopyMemoryToImageIndirectNV-offset-07676  
  `copyBufferAddress` **must** be 4 byte aligned
- [](#VUID-vkCmdCopyMemoryToImageIndirectNV-stride-07677)VUID-vkCmdCopyMemoryToImageIndirectNV-stride-07677  
  `stride` **must** be a multiple of `4` and **must** be greater than or equal to sizeof([VkCopyMemoryToImageIndirectCommandNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyMemoryToImageIndirectCommandNV.html))
- [](#VUID-vkCmdCopyMemoryToImageIndirectNV-commandBuffer-10956)VUID-vkCmdCopyMemoryToImageIndirectNV-commandBuffer-10956  
  The [VkCommandPool](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandPool.html) that `commandBuffer` was allocated from **must** support at least one of the queue types specified in [VkPhysicalDeviceCopyMemoryIndirectPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceCopyMemoryIndirectPropertiesKHR.html)::`supportedQueues`

<!--THE END-->

- [](#VUID-vkCmdCopyMemoryToImageIndirectNV-dstImage-07661)VUID-vkCmdCopyMemoryToImageIndirectNV-dstImage-07661  
  `dstImage` **must** not be a protected image
- [](#VUID-vkCmdCopyMemoryToImageIndirectNV-aspectMask-07662)VUID-vkCmdCopyMemoryToImageIndirectNV-aspectMask-07662  
  The `aspectMask` member for every subresource in `pImageSubresources` **must** only have a single bit set
- [](#VUID-vkCmdCopyMemoryToImageIndirectNV-dstImage-07663)VUID-vkCmdCopyMemoryToImageIndirectNV-dstImage-07663  
  The image region specified by each element in `copyBufferAddress` **must** be a region that is contained within `dstImage`
- [](#VUID-vkCmdCopyMemoryToImageIndirectNV-dstImage-07664)VUID-vkCmdCopyMemoryToImageIndirectNV-dstImage-07664  
  `dstImage` **must** have been created with `VK_IMAGE_USAGE_TRANSFER_DST_BIT` usage flag
- [](#VUID-vkCmdCopyMemoryToImageIndirectNV-dstImage-07665)VUID-vkCmdCopyMemoryToImageIndirectNV-dstImage-07665  
  If `dstImage` is non-sparse then it **must** be bound completely and contiguously to a single `VkDeviceMemory` object

<!--THE END-->

- [](#VUID-vkCmdCopyMemoryToImageIndirectNV-dstImage-07973)VUID-vkCmdCopyMemoryToImageIndirectNV-dstImage-07973  
  `dstImage` **must** have a sample count equal to `VK_SAMPLE_COUNT_1_BIT`
- [](#VUID-vkCmdCopyMemoryToImageIndirectNV-dstImageLayout-07667)VUID-vkCmdCopyMemoryToImageIndirectNV-dstImageLayout-07667  
  `dstImageLayout` **must** specify the layout of the image subresources of `dstImage` at the time this command is executed on a `VkDevice`
- [](#VUID-vkCmdCopyMemoryToImageIndirectNV-dstImageLayout-07669)VUID-vkCmdCopyMemoryToImageIndirectNV-dstImageLayout-07669  
  `dstImageLayout` **must** be `VK_IMAGE_LAYOUT_TRANSFER_DST_OPTIMAL`, `VK_IMAGE_LAYOUT_SHARED_PRESENT_KHR`, or `VK_IMAGE_LAYOUT_GENERAL`
- [](#VUID-vkCmdCopyMemoryToImageIndirectNV-mipLevel-07670)VUID-vkCmdCopyMemoryToImageIndirectNV-mipLevel-07670  
  The specified `mipLevel` of each region in `pImageSubresources` **must** be less than the `mipLevels` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) when `dstImage` was created
- [](#VUID-vkCmdCopyMemoryToImageIndirectNV-layerCount-08764)VUID-vkCmdCopyMemoryToImageIndirectNV-layerCount-08764  
  If the specified `layerCount` of each region in `pImageSubresources` is not `VK_REMAINING_ARRAY_LAYERS`, the specified `baseArrayLayer` + `layerCount` of each region in `pImageSubresources` **must** be less than or equal to the `arrayLayers` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) when `dstImage` was created
- [](#VUID-vkCmdCopyMemoryToImageIndirectNV-dstImage-07673)VUID-vkCmdCopyMemoryToImageIndirectNV-dstImage-07673  
  `dstImage` **must** not have been created with `flags` containing `VK_IMAGE_CREATE_SUBSAMPLED_BIT_EXT`
- [](#VUID-vkCmdCopyMemoryToImageIndirectNV-commandBuffer-07674)VUID-vkCmdCopyMemoryToImageIndirectNV-commandBuffer-07674  
  If the queue family used to create the [VkCommandPool](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandPool.html) which `commandBuffer` was allocated from does not support `VK_QUEUE_GRAPHICS_BIT`, for each region, the `aspectMask` member of `pImageSubresources` **must** not be `VK_IMAGE_ASPECT_DEPTH_BIT` or `VK_IMAGE_ASPECT_STENCIL_BIT`
- [](#VUID-vkCmdCopyMemoryToImageIndirectNV-dstImage-10974)VUID-vkCmdCopyMemoryToImageIndirectNV-dstImage-10974  
  The format features of `dstImage` **must** contain `VK_FORMAT_FEATURE_TRANSFER_DST_BIT`
- [](#VUID-vkCmdCopyMemoryToImageIndirectNV-copyBufferAddress-10975)VUID-vkCmdCopyMemoryToImageIndirectNV-copyBufferAddress-10975  
  Any of the source or destination memory regions specified in `copyBufferAddress` **must** not overlap with any of the specified destination memory regions

Valid Usage (Implicit)

- [](#VUID-vkCmdCopyMemoryToImageIndirectNV-commandBuffer-parameter)VUID-vkCmdCopyMemoryToImageIndirectNV-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdCopyMemoryToImageIndirectNV-copyBufferAddress-parameter)VUID-vkCmdCopyMemoryToImageIndirectNV-copyBufferAddress-parameter  
  `copyBufferAddress` **must** be a valid [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html) value
- [](#VUID-vkCmdCopyMemoryToImageIndirectNV-dstImage-parameter)VUID-vkCmdCopyMemoryToImageIndirectNV-dstImage-parameter  
  `dstImage` **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html) handle
- [](#VUID-vkCmdCopyMemoryToImageIndirectNV-dstImageLayout-parameter)VUID-vkCmdCopyMemoryToImageIndirectNV-dstImageLayout-parameter  
  `dstImageLayout` **must** be a valid [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html) value
- [](#VUID-vkCmdCopyMemoryToImageIndirectNV-pImageSubresources-parameter)VUID-vkCmdCopyMemoryToImageIndirectNV-pImageSubresources-parameter  
  `pImageSubresources` **must** be a valid pointer to an array of `copyCount` valid [VkImageSubresourceLayers](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresourceLayers.html) structures
- [](#VUID-vkCmdCopyMemoryToImageIndirectNV-commandBuffer-recording)VUID-vkCmdCopyMemoryToImageIndirectNV-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdCopyMemoryToImageIndirectNV-commandBuffer-cmdpool)VUID-vkCmdCopyMemoryToImageIndirectNV-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support VK\_QUEUE\_COMPUTE\_BIT, VK\_QUEUE\_GRAPHICS\_BIT, or VK\_QUEUE\_TRANSFER\_BIT operations
- [](#VUID-vkCmdCopyMemoryToImageIndirectNV-renderpass)VUID-vkCmdCopyMemoryToImageIndirectNV-renderpass  
  This command **must** only be called outside of a render pass instance
- [](#VUID-vkCmdCopyMemoryToImageIndirectNV-videocoding)VUID-vkCmdCopyMemoryToImageIndirectNV-videocoding  
  This command **must** only be called outside of a video coding scope
- [](#VUID-vkCmdCopyMemoryToImageIndirectNV-copyCount-arraylength)VUID-vkCmdCopyMemoryToImageIndirectNV-copyCount-arraylength  
  `copyCount` **must** be greater than `0`
- [](#VUID-vkCmdCopyMemoryToImageIndirectNV-commonparent)VUID-vkCmdCopyMemoryToImageIndirectNV-commonparent  
  Both of `commandBuffer`, and `dstImage` **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

Host Synchronization

- Host access to `commandBuffer` **must** be externally synchronized
- Host access to the `VkCommandPool` that `commandBuffer` was allocated from **must** be externally synchronized

Command Properties

     [Command Buffer Levels](#VkCommandBufferLevel) [Render Pass Scope](#vkCmdBeginRenderPass) [Video Coding Scope](#vkCmdBeginVideoCodingKHR) [Supported Queue Types](#VkQueueFlagBits) [Command Type](#fundamentals-queueoperation-command-types)

Primary  
Secondary

Outside

Outside

VK\_QUEUE\_COMPUTE\_BIT  
VK\_QUEUE\_GRAPHICS\_BIT  
VK\_QUEUE\_TRANSFER\_BIT

Action

Conditional Rendering

vkCmdCopyMemoryToImageIndirectNV is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_NV\_copy\_memory\_indirect](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_copy_memory_indirect.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html), [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html), [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html), [VkImageSubresourceLayers](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresourceLayers.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdCopyMemoryToImageIndirectNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0