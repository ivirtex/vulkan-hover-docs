# vkCmdCopyImageToBuffer2(3) Manual Page

## Name

vkCmdCopyImageToBuffer2 - Copy image data into a buffer



## [](#_c_specification)C Specification

To copy data from an image object to a buffer object, call:

```c++
// Provided by VK_VERSION_1_3
void vkCmdCopyImageToBuffer2(
    VkCommandBuffer                             commandBuffer,
    const VkCopyImageToBufferInfo2*             pCopyImageToBufferInfo);
```

or the equivalent command

```c++
// Provided by VK_KHR_copy_commands2
void vkCmdCopyImageToBuffer2KHR(
    VkCommandBuffer                             commandBuffer,
    const VkCopyImageToBufferInfo2*             pCopyImageToBufferInfo);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `pCopyImageToBufferInfo` is a pointer to a [VkCopyImageToBufferInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyImageToBufferInfo2.html) structure describing the copy parameters.

## [](#_description)Description

This command is functionally identical to [vkCmdCopyImageToBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyImageToBuffer.html), but includes extensible sub-structures that include `sType` and `pNext` parameters, allowing them to be more easily extended.

Valid Usage

- [](#VUID-vkCmdCopyImageToBuffer2-commandBuffer-01831)VUID-vkCmdCopyImageToBuffer2-commandBuffer-01831  
  If `commandBuffer` is an unprotected command buffer and [`protectedNoFault`](#limits-protectedNoFault) is not supported, `srcImage` **must** not be a protected image
- [](#VUID-vkCmdCopyImageToBuffer2-commandBuffer-01832)VUID-vkCmdCopyImageToBuffer2-commandBuffer-01832  
  If `commandBuffer` is an unprotected command buffer and [`protectedNoFault`](#limits-protectedNoFault) is not supported, `dstBuffer` **must** not be a protected buffer
- [](#VUID-vkCmdCopyImageToBuffer2-commandBuffer-01833)VUID-vkCmdCopyImageToBuffer2-commandBuffer-01833  
  If `commandBuffer` is a protected command buffer and [`protectedNoFault`](#limits-protectedNoFault) is not supported, `dstBuffer` **must** not be an unprotected buffer
- [](#VUID-vkCmdCopyImageToBuffer2-commandBuffer-07746)VUID-vkCmdCopyImageToBuffer2-commandBuffer-07746  
  If the queue family used to create the [VkCommandPool](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandPool.html) which `commandBuffer` was allocated from does not support `VK_QUEUE_GRAPHICS_BIT` or `VK_QUEUE_COMPUTE_BIT`, the `bufferOffset` member of any element of `pCopyImageToBufferInfo->pRegions` **must** be a multiple of `4`
- [](#VUID-vkCmdCopyImageToBuffer2-imageOffset-07747)VUID-vkCmdCopyImageToBuffer2-imageOffset-07747  
  The `imageOffset` and `imageExtent` members of each element of `pCopyImageToBufferInfo->pRegions` **must** respect the image transfer granularity requirements of `commandBuffer`’s command pool’s queue family, as described in [VkQueueFamilyProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueFamilyProperties.html)
- [](#VUID-vkCmdCopyImageToBuffer2-commandBuffer-10216)VUID-vkCmdCopyImageToBuffer2-commandBuffer-10216  
  If the queue family used to create the [VkCommandPool](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandPool.html) which `commandBuffer` was allocated from does not support `VK_QUEUE_GRAPHICS_BIT`, for each element of `pCopyImageToBufferInfo->pRegions`, the `aspectMask` member of `imageSubresource` **must** not be `VK_IMAGE_ASPECT_DEPTH_BIT` or `VK_IMAGE_ASPECT_STENCIL_BIT`

Valid Usage (Implicit)

- [](#VUID-vkCmdCopyImageToBuffer2-commandBuffer-parameter)VUID-vkCmdCopyImageToBuffer2-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdCopyImageToBuffer2-pCopyImageToBufferInfo-parameter)VUID-vkCmdCopyImageToBuffer2-pCopyImageToBufferInfo-parameter  
  `pCopyImageToBufferInfo` **must** be a valid pointer to a valid [VkCopyImageToBufferInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyImageToBufferInfo2.html) structure
- [](#VUID-vkCmdCopyImageToBuffer2-commandBuffer-recording)VUID-vkCmdCopyImageToBuffer2-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdCopyImageToBuffer2-commandBuffer-cmdpool)VUID-vkCmdCopyImageToBuffer2-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support VK\_QUEUE\_COMPUTE\_BIT, VK\_QUEUE\_GRAPHICS\_BIT, or VK\_QUEUE\_TRANSFER\_BIT operations
- [](#VUID-vkCmdCopyImageToBuffer2-renderpass)VUID-vkCmdCopyImageToBuffer2-renderpass  
  This command **must** only be called outside of a render pass instance
- [](#VUID-vkCmdCopyImageToBuffer2-videocoding)VUID-vkCmdCopyImageToBuffer2-videocoding  
  This command **must** only be called outside of a video coding scope

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

vkCmdCopyImageToBuffer2 is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_KHR\_copy\_commands2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_copy_commands2.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkCopyImageToBufferInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyImageToBufferInfo2.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdCopyImageToBuffer2).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0