# vkCmdCopyTensorARM(3) Manual Page

## Name

vkCmdCopyTensorARM - Copy data between tensors



## [](#_c_specification)C Specification

To copy data between tensor objects, call:

```c++
// Provided by VK_ARM_tensors
void vkCmdCopyTensorARM(
    VkCommandBuffer                             commandBuffer,
     const VkCopyTensorInfoARM*                 pCopyTensorInfo);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `pCopyTensorInfo` is a pointer to [VkCopyTensorInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyTensorInfoARM.html) structure describing the copy parameters.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkCmdCopyTensorARM-commandBuffer-parameter)VUID-vkCmdCopyTensorARM-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdCopyTensorARM-pCopyTensorInfo-parameter)VUID-vkCmdCopyTensorARM-pCopyTensorInfo-parameter  
  `pCopyTensorInfo` **must** be a valid pointer to a valid [VkCopyTensorInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyTensorInfoARM.html) structure
- [](#VUID-vkCmdCopyTensorARM-commandBuffer-recording)VUID-vkCmdCopyTensorARM-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdCopyTensorARM-commandBuffer-cmdpool)VUID-vkCmdCopyTensorARM-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support transfer, graphics, or compute operations
- [](#VUID-vkCmdCopyTensorARM-renderpass)VUID-vkCmdCopyTensorARM-renderpass  
  This command **must** only be called outside of a render pass instance
- [](#VUID-vkCmdCopyTensorARM-videocoding)VUID-vkCmdCopyTensorARM-videocoding  
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

vkCmdCopyTensorARM is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_ARM\_tensors](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_tensors.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkCopyTensorInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyTensorInfoARM.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdCopyTensorARM)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0