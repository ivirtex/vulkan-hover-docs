# vkCmdCopyImage2(3) Manual Page

## Name

vkCmdCopyImage2 - Copy data between images



## [](#_c_specification)C Specification

To copy data between image objects, call:

```c++
// Provided by VK_VERSION_1_3
void vkCmdCopyImage2(
    VkCommandBuffer                             commandBuffer,
    const VkCopyImageInfo2*                     pCopyImageInfo);
```

or the equivalent command

```c++
// Provided by VK_KHR_copy_commands2
void vkCmdCopyImage2KHR(
    VkCommandBuffer                             commandBuffer,
    const VkCopyImageInfo2*                     pCopyImageInfo);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `pCopyImageInfo` is a pointer to a [VkCopyImageInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyImageInfo2.html) structure describing the copy parameters.

## [](#_description)Description

This command is functionally identical to [vkCmdCopyImage](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyImage.html), but includes extensible sub-structures that include `sType` and `pNext` parameters, allowing them to be more easily extended.

Valid Usage

- [](#VUID-vkCmdCopyImage2-commandBuffer-01825)VUID-vkCmdCopyImage2-commandBuffer-01825  
  If `commandBuffer` is an unprotected command buffer and [`protectedNoFault`](#limits-protectedNoFault) is not supported, `srcImage` **must** not be a protected image
- [](#VUID-vkCmdCopyImage2-commandBuffer-01826)VUID-vkCmdCopyImage2-commandBuffer-01826  
  If `commandBuffer` is an unprotected command buffer and [`protectedNoFault`](#limits-protectedNoFault) is not supported, `dstImage` **must** not be a protected image
- [](#VUID-vkCmdCopyImage2-commandBuffer-01827)VUID-vkCmdCopyImage2-commandBuffer-01827  
  If `commandBuffer` is a protected command buffer and [`protectedNoFault`](#limits-protectedNoFault) is not supported, `dstImage` **must** not be an unprotected image
- [](#VUID-vkCmdCopyImage2-commandBuffer-10217)VUID-vkCmdCopyImage2-commandBuffer-10217  
  If the queue family used to create the [VkCommandPool](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandPool.html) which `commandBuffer` was allocated from does not support `VK_QUEUE_GRAPHICS_BIT`, for each element of `pCopyImageInfo->pRegions`, where the `aspectMask` member of `srcSubresource` is `VK_IMAGE_ASPECT_COLOR_BIT`, the `aspectMask` of `dstSubresource` **must** not be `VK_IMAGE_ASPECT_DEPTH_BIT` or `VK_IMAGE_ASPECT_STENCIL_BIT`
- [](#VUID-vkCmdCopyImage2-commandBuffer-10218)VUID-vkCmdCopyImage2-commandBuffer-10218  
  If the queue family used to create the [VkCommandPool](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandPool.html) which `commandBuffer` was allocated from does not support `VK_QUEUE_GRAPHICS_BIT`, for each element of `pCopyImageInfo->pRegions`, where the `aspectMask` member of `dstSubresource` is `VK_IMAGE_ASPECT_COLOR_BIT` then the `aspectMask` of `srcSubresource` **must** not be `VK_IMAGE_ASPECT_DEPTH_BIT` or `VK_IMAGE_ASPECT_STENCIL_BIT`

Valid Usage (Implicit)

- [](#VUID-vkCmdCopyImage2-commandBuffer-parameter)VUID-vkCmdCopyImage2-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdCopyImage2-pCopyImageInfo-parameter)VUID-vkCmdCopyImage2-pCopyImageInfo-parameter  
  `pCopyImageInfo` **must** be a valid pointer to a valid [VkCopyImageInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyImageInfo2.html) structure
- [](#VUID-vkCmdCopyImage2-commandBuffer-recording)VUID-vkCmdCopyImage2-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdCopyImage2-commandBuffer-cmdpool)VUID-vkCmdCopyImage2-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support transfer, graphics, or compute operations
- [](#VUID-vkCmdCopyImage2-renderpass)VUID-vkCmdCopyImage2-renderpass  
  This command **must** only be called outside of a render pass instance
- [](#VUID-vkCmdCopyImage2-videocoding)VUID-vkCmdCopyImage2-videocoding  
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

## [](#_see_also)See Also

[VK\_KHR\_copy\_commands2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_copy_commands2.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkCopyImageInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyImageInfo2.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdCopyImage2)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0