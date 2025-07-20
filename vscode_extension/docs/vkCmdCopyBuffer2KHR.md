# vkCmdCopyBuffer2(3) Manual Page

## Name

vkCmdCopyBuffer2 - Copy data between buffer regions



## [](#_c_specification)C Specification

To copy data between buffer objects, call:

```c++
// Provided by VK_VERSION_1_3
void vkCmdCopyBuffer2(
    VkCommandBuffer                             commandBuffer,
    const VkCopyBufferInfo2*                    pCopyBufferInfo);
```

or the equivalent command

```c++
// Provided by VK_KHR_copy_commands2
void vkCmdCopyBuffer2KHR(
    VkCommandBuffer                             commandBuffer,
    const VkCopyBufferInfo2*                    pCopyBufferInfo);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `pCopyBufferInfo` is a pointer to a [VkCopyBufferInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyBufferInfo2.html) structure describing the copy parameters.

## [](#_description)Description

Each source region specified by `pCopyBufferInfo->pRegions` is copied from the source buffer to the destination region of the destination buffer. If any of the specified regions in `pCopyBufferInfo->srcBuffer` overlaps in memory with any of the specified regions in `pCopyBufferInfo->dstBuffer`, values read from those overlapping regions are undefined.

Valid Usage

- [](#VUID-vkCmdCopyBuffer2-commandBuffer-01822)VUID-vkCmdCopyBuffer2-commandBuffer-01822  
  If `commandBuffer` is an unprotected command buffer and [`protectedNoFault`](#limits-protectedNoFault) is not supported, `srcBuffer` **must** not be a protected buffer
- [](#VUID-vkCmdCopyBuffer2-commandBuffer-01823)VUID-vkCmdCopyBuffer2-commandBuffer-01823  
  If `commandBuffer` is an unprotected command buffer and [`protectedNoFault`](#limits-protectedNoFault) is not supported, `dstBuffer` **must** not be a protected buffer
- [](#VUID-vkCmdCopyBuffer2-commandBuffer-01824)VUID-vkCmdCopyBuffer2-commandBuffer-01824  
  If `commandBuffer` is a protected command buffer and [`protectedNoFault`](#limits-protectedNoFault) is not supported, `dstBuffer` **must** not be an unprotected buffer

Valid Usage (Implicit)

- [](#VUID-vkCmdCopyBuffer2-commandBuffer-parameter)VUID-vkCmdCopyBuffer2-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdCopyBuffer2-pCopyBufferInfo-parameter)VUID-vkCmdCopyBuffer2-pCopyBufferInfo-parameter  
  `pCopyBufferInfo` **must** be a valid pointer to a valid [VkCopyBufferInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyBufferInfo2.html) structure
- [](#VUID-vkCmdCopyBuffer2-commandBuffer-recording)VUID-vkCmdCopyBuffer2-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdCopyBuffer2-commandBuffer-cmdpool)VUID-vkCmdCopyBuffer2-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support transfer, graphics, or compute operations
- [](#VUID-vkCmdCopyBuffer2-renderpass)VUID-vkCmdCopyBuffer2-renderpass  
  This command **must** only be called outside of a render pass instance
- [](#VUID-vkCmdCopyBuffer2-videocoding)VUID-vkCmdCopyBuffer2-videocoding  
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

vkCmdCopyBuffer2 is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_KHR\_copy\_commands2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_copy_commands2.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkCopyBufferInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyBufferInfo2.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdCopyBuffer2)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0