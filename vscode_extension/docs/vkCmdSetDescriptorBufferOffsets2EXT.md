# vkCmdSetDescriptorBufferOffsets2EXT(3) Manual Page

## Name

vkCmdSetDescriptorBufferOffsets2EXT - Setting descriptor buffer offsets in a command buffer



## [](#_c_specification)C Specification

Alternatively, to set descriptor buffer offsets in a command buffer, call:

```c++
// Provided by VK_KHR_maintenance6 with VK_EXT_descriptor_buffer
void vkCmdSetDescriptorBufferOffsets2EXT(
    VkCommandBuffer                             commandBuffer,
    const VkSetDescriptorBufferOffsetsInfoEXT*  pSetDescriptorBufferOffsetsInfo);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer in which the descriptor buffer offsets will be set.
- `pSetDescriptorBufferOffsetsInfo` is a pointer to a `VkSetDescriptorBufferOffsetsInfoEXT` structure.

## [](#_description)Description

Valid Usage

- [](#VUID-vkCmdSetDescriptorBufferOffsets2EXT-descriptorBuffer-09470)VUID-vkCmdSetDescriptorBufferOffsets2EXT-descriptorBuffer-09470  
  The [`descriptorBuffer`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-descriptorBuffer) feature **must** be enabled
- [](#VUID-vkCmdSetDescriptorBufferOffsets2EXT-pSetDescriptorBufferOffsetsInfo-09471)VUID-vkCmdSetDescriptorBufferOffsets2EXT-pSetDescriptorBufferOffsetsInfo-09471  
  Each bit in `pSetDescriptorBufferOffsetsInfo->stageFlags` **must** be a stage supported by the `commandBuffer`’s parent `VkCommandPool`’s queue family

Valid Usage (Implicit)

- [](#VUID-vkCmdSetDescriptorBufferOffsets2EXT-commandBuffer-parameter)VUID-vkCmdSetDescriptorBufferOffsets2EXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdSetDescriptorBufferOffsets2EXT-pSetDescriptorBufferOffsetsInfo-parameter)VUID-vkCmdSetDescriptorBufferOffsets2EXT-pSetDescriptorBufferOffsetsInfo-parameter  
  `pSetDescriptorBufferOffsetsInfo` **must** be a valid pointer to a valid [VkSetDescriptorBufferOffsetsInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSetDescriptorBufferOffsetsInfoEXT.html) structure
- [](#VUID-vkCmdSetDescriptorBufferOffsets2EXT-commandBuffer-recording)VUID-vkCmdSetDescriptorBufferOffsets2EXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdSetDescriptorBufferOffsets2EXT-commandBuffer-cmdpool)VUID-vkCmdSetDescriptorBufferOffsets2EXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics, compute, or data\_graph operations
- [](#VUID-vkCmdSetDescriptorBufferOffsets2EXT-videocoding)VUID-vkCmdSetDescriptorBufferOffsets2EXT-videocoding  
  This command **must** only be called outside of a video coding scope

Host Synchronization

- Host access to `commandBuffer` **must** be externally synchronized
- Host access to the `VkCommandPool` that `commandBuffer` was allocated from **must** be externally synchronized

Command Properties

     [Command Buffer Levels](#VkCommandBufferLevel) [Render Pass Scope](#vkCmdBeginRenderPass) [Video Coding Scope](#vkCmdBeginVideoCodingKHR) [Supported Queue Types](#VkQueueFlagBits) [Command Type](#fundamentals-queueoperation-command-types)

Primary  
Secondary

Both

Outside

Graphics  
Compute  
Data\_Graph

State

Conditional Rendering

vkCmdSetDescriptorBufferOffsets2EXT is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_EXT\_descriptor\_buffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_descriptor_buffer.html), [VK\_KHR\_maintenance6](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance6.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkSetDescriptorBufferOffsetsInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSetDescriptorBufferOffsetsInfoEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdSetDescriptorBufferOffsets2EXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0