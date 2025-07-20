# vkCmdCopyBufferToImage2(3) Manual Page

## Name

vkCmdCopyBufferToImage2 - Copy data from a buffer into an image



## [](#_c_specification)C Specification

To copy data from a buffer object to an image object, call:

```c++
// Provided by VK_VERSION_1_3
void vkCmdCopyBufferToImage2(
    VkCommandBuffer                             commandBuffer,
    const VkCopyBufferToImageInfo2*             pCopyBufferToImageInfo);
```

or the equivalent command

```c++
// Provided by VK_KHR_copy_commands2
void vkCmdCopyBufferToImage2KHR(
    VkCommandBuffer                             commandBuffer,
    const VkCopyBufferToImageInfo2*             pCopyBufferToImageInfo);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `pCopyBufferToImageInfo` is a pointer to a [VkCopyBufferToImageInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyBufferToImageInfo2.html) structure describing the copy parameters.

## [](#_description)Description

This command is functionally identical to [vkCmdCopyBufferToImage](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyBufferToImage.html), but includes extensible sub-structures that include `sType` and `pNext` parameters, allowing them to be more easily extended.

Valid Usage

- [](#VUID-vkCmdCopyBufferToImage2-commandBuffer-01828)VUID-vkCmdCopyBufferToImage2-commandBuffer-01828  
  If `commandBuffer` is an unprotected command buffer and [`protectedNoFault`](#limits-protectedNoFault) is not supported, `srcBuffer` **must** not be a protected buffer
- [](#VUID-vkCmdCopyBufferToImage2-commandBuffer-01829)VUID-vkCmdCopyBufferToImage2-commandBuffer-01829  
  If `commandBuffer` is an unprotected command buffer and [`protectedNoFault`](#limits-protectedNoFault) is not supported, `dstImage` **must** not be a protected image
- [](#VUID-vkCmdCopyBufferToImage2-commandBuffer-01830)VUID-vkCmdCopyBufferToImage2-commandBuffer-01830  
  If `commandBuffer` is a protected command buffer and [`protectedNoFault`](#limits-protectedNoFault) is not supported, `dstImage` **must** not be an unprotected image
- [](#VUID-vkCmdCopyBufferToImage2-commandBuffer-07737)VUID-vkCmdCopyBufferToImage2-commandBuffer-07737  
  If the queue family used to create the [VkCommandPool](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandPool.html) which `commandBuffer` was allocated from does not support `VK_QUEUE_GRAPHICS_BIT` or `VK_QUEUE_COMPUTE_BIT`, the `bufferOffset` member of any element of `pCopyBufferToImageInfo->pRegions` **must** be a multiple of `4`
- [](#VUID-vkCmdCopyBufferToImage2-imageOffset-07738)VUID-vkCmdCopyBufferToImage2-imageOffset-07738  
  The `imageOffset` and `imageExtent` members of each element of `pCopyBufferToImageInfo->pRegions` **must** respect the image transfer granularity requirements of `commandBuffer`’s command pool’s queue family, as described in [VkQueueFamilyProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueFamilyProperties.html)
- [](#VUID-vkCmdCopyBufferToImage2-commandBuffer-07739)VUID-vkCmdCopyBufferToImage2-commandBuffer-07739  
  If the queue family used to create the [VkCommandPool](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandPool.html) which `commandBuffer` was allocated from does not support `VK_QUEUE_GRAPHICS_BIT`, for each element of `pCopyBufferToImageInfo->pRegions`, the `aspectMask` member of `imageSubresource` **must** not be `VK_IMAGE_ASPECT_DEPTH_BIT` or `VK_IMAGE_ASPECT_STENCIL_BIT`

Valid Usage (Implicit)

- [](#VUID-vkCmdCopyBufferToImage2-commandBuffer-parameter)VUID-vkCmdCopyBufferToImage2-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdCopyBufferToImage2-pCopyBufferToImageInfo-parameter)VUID-vkCmdCopyBufferToImage2-pCopyBufferToImageInfo-parameter  
  `pCopyBufferToImageInfo` **must** be a valid pointer to a valid [VkCopyBufferToImageInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyBufferToImageInfo2.html) structure
- [](#VUID-vkCmdCopyBufferToImage2-commandBuffer-recording)VUID-vkCmdCopyBufferToImage2-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdCopyBufferToImage2-commandBuffer-cmdpool)VUID-vkCmdCopyBufferToImage2-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support transfer, graphics, or compute operations
- [](#VUID-vkCmdCopyBufferToImage2-renderpass)VUID-vkCmdCopyBufferToImage2-renderpass  
  This command **must** only be called outside of a render pass instance
- [](#VUID-vkCmdCopyBufferToImage2-videocoding)VUID-vkCmdCopyBufferToImage2-videocoding  
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

Transfer  
Graphics  
Compute

Action

Conditional Rendering

vkCmdCopyBufferToImage2 is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_KHR\_copy\_commands2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_copy_commands2.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkCopyBufferToImageInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyBufferToImageInfo2.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdCopyBufferToImage2)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0