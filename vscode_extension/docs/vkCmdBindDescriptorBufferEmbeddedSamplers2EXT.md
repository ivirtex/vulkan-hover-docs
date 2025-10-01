# vkCmdBindDescriptorBufferEmbeddedSamplers2EXT(3) Manual Page

## Name

vkCmdBindDescriptorBufferEmbeddedSamplers2EXT - Setting embedded immutable samplers offsets in a command buffer



## [](#_c_specification)C Specification

Alternatively, to bind an embedded immutable sampler set to a command buffer, call:

```c++
// Provided by VK_KHR_maintenance6 with VK_EXT_descriptor_buffer
void vkCmdBindDescriptorBufferEmbeddedSamplers2EXT(
    VkCommandBuffer                             commandBuffer,
    const VkBindDescriptorBufferEmbeddedSamplersInfoEXT* pBindDescriptorBufferEmbeddedSamplersInfo);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer that the embedded immutable samplers will be bound to.
- `pBindDescriptorBufferEmbeddedSamplersInfo` is a pointer to a `VkBindDescriptorBufferEmbeddedSamplersInfoEXT` structure.

## [](#_description)Description

Valid Usage

- [](#VUID-vkCmdBindDescriptorBufferEmbeddedSamplers2EXT-descriptorBuffer-09472)VUID-vkCmdBindDescriptorBufferEmbeddedSamplers2EXT-descriptorBuffer-09472  
  The [`descriptorBuffer`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-descriptorBuffer) feature **must** be enabled
- [](#VUID-vkCmdBindDescriptorBufferEmbeddedSamplers2EXT-pBindDescriptorBufferEmbeddedSamplersInfo-09473)VUID-vkCmdBindDescriptorBufferEmbeddedSamplers2EXT-pBindDescriptorBufferEmbeddedSamplersInfo-09473  
  Each bit in `pBindDescriptorBufferEmbeddedSamplersInfo->stageFlags` **must** be a stage supported by the `commandBuffer`’s parent `VkCommandPool`’s queue family

Valid Usage (Implicit)

- [](#VUID-vkCmdBindDescriptorBufferEmbeddedSamplers2EXT-commandBuffer-parameter)VUID-vkCmdBindDescriptorBufferEmbeddedSamplers2EXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdBindDescriptorBufferEmbeddedSamplers2EXT-pBindDescriptorBufferEmbeddedSamplersInfo-parameter)VUID-vkCmdBindDescriptorBufferEmbeddedSamplers2EXT-pBindDescriptorBufferEmbeddedSamplersInfo-parameter  
  `pBindDescriptorBufferEmbeddedSamplersInfo` **must** be a valid pointer to a valid [VkBindDescriptorBufferEmbeddedSamplersInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindDescriptorBufferEmbeddedSamplersInfoEXT.html) structure
- [](#VUID-vkCmdBindDescriptorBufferEmbeddedSamplers2EXT-commandBuffer-recording)VUID-vkCmdBindDescriptorBufferEmbeddedSamplers2EXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdBindDescriptorBufferEmbeddedSamplers2EXT-commandBuffer-cmdpool)VUID-vkCmdBindDescriptorBufferEmbeddedSamplers2EXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support VK\_QUEUE\_COMPUTE\_BIT, or VK\_QUEUE\_GRAPHICS\_BIT operations
- [](#VUID-vkCmdBindDescriptorBufferEmbeddedSamplers2EXT-videocoding)VUID-vkCmdBindDescriptorBufferEmbeddedSamplers2EXT-videocoding  
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

VK\_QUEUE\_COMPUTE\_BIT  
VK\_QUEUE\_GRAPHICS\_BIT

State

Conditional Rendering

vkCmdBindDescriptorBufferEmbeddedSamplers2EXT is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_EXT\_descriptor\_buffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_descriptor_buffer.html), [VK\_KHR\_maintenance6](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance6.html), [VkBindDescriptorBufferEmbeddedSamplersInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindDescriptorBufferEmbeddedSamplersInfoEXT.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdBindDescriptorBufferEmbeddedSamplers2EXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0