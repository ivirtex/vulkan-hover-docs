# vkCmdBindDescriptorBufferEmbeddedSamplersEXT(3) Manual Page

## Name

vkCmdBindDescriptorBufferEmbeddedSamplersEXT - Setting embedded immutable samplers offsets in a command buffer



## [](#_c_specification)C Specification

To bind an embedded immutable sampler set to a command buffer, call:

```c++
// Provided by VK_EXT_descriptor_buffer
void vkCmdBindDescriptorBufferEmbeddedSamplersEXT(
    VkCommandBuffer                             commandBuffer,
    VkPipelineBindPoint                         pipelineBindPoint,
    VkPipelineLayout                            layout,
    uint32_t                                    set);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer that the embedded immutable samplers will be bound to.
- `pipelineBindPoint` is a [VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBindPoint.html) indicating the type of the pipeline that will use the embedded immutable samplers.
- `layout` is a [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayout.html) object used to program the bindings.
- `set` is the number of the set to be bound.

## [](#_description)Description

`vkCmdBindDescriptorBufferEmbeddedSamplersEXT` binds the embedded immutable samplers in `set` of `layout` to `set` for the command buffer for subsequent [bound pipeline commands](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-bindpoint-commands) set by `pipelineBindPoint`. Any previous binding to this set by [vkCmdSetDescriptorBufferOffsetsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetDescriptorBufferOffsetsEXT.html) or this command is overwritten. Any sets that were last bound by a call to [vkCmdBindDescriptorSets](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBindDescriptorSets.html) are invalidated upon calling this command. Other sets will also be invalidated upon calling this command if `layout` differs from the pipeline layout used to bind those other sets, as described in [Pipeline Layout Compatibility](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-compatibility).

Valid Usage

- [](#VUID-vkCmdBindDescriptorBufferEmbeddedSamplersEXT-set-08070)VUID-vkCmdBindDescriptorBufferEmbeddedSamplersEXT-set-08070  
  The [VkDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayout.html) at index `set` when `layout` was created **must** have been created with the `VK_DESCRIPTOR_SET_LAYOUT_CREATE_EMBEDDED_IMMUTABLE_SAMPLERS_BIT_EXT` bit set
- [](#VUID-vkCmdBindDescriptorBufferEmbeddedSamplersEXT-set-08071)VUID-vkCmdBindDescriptorBufferEmbeddedSamplersEXT-set-08071  
  `set` **must** be less than or equal to [VkPipelineLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayoutCreateInfo.html)::`setLayoutCount` provided when `layout` was created
- [](#VUID-vkCmdBindDescriptorBufferEmbeddedSamplersEXT-None-08068)VUID-vkCmdBindDescriptorBufferEmbeddedSamplersEXT-None-08068  
  The [`descriptorBuffer`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-descriptorBuffer) feature **must** be enabled
- [](#VUID-vkCmdBindDescriptorBufferEmbeddedSamplersEXT-pipelineBindPoint-08069)VUID-vkCmdBindDescriptorBufferEmbeddedSamplersEXT-pipelineBindPoint-08069  
  `pipelineBindPoint` **must** be supported by the `commandBuffer`’s parent `VkCommandPool`’s queue family

Valid Usage (Implicit)

- [](#VUID-vkCmdBindDescriptorBufferEmbeddedSamplersEXT-commandBuffer-parameter)VUID-vkCmdBindDescriptorBufferEmbeddedSamplersEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdBindDescriptorBufferEmbeddedSamplersEXT-pipelineBindPoint-parameter)VUID-vkCmdBindDescriptorBufferEmbeddedSamplersEXT-pipelineBindPoint-parameter  
  `pipelineBindPoint` **must** be a valid [VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBindPoint.html) value
- [](#VUID-vkCmdBindDescriptorBufferEmbeddedSamplersEXT-layout-parameter)VUID-vkCmdBindDescriptorBufferEmbeddedSamplersEXT-layout-parameter  
  `layout` **must** be a valid [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayout.html) handle
- [](#VUID-vkCmdBindDescriptorBufferEmbeddedSamplersEXT-commandBuffer-recording)VUID-vkCmdBindDescriptorBufferEmbeddedSamplersEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdBindDescriptorBufferEmbeddedSamplersEXT-commandBuffer-cmdpool)VUID-vkCmdBindDescriptorBufferEmbeddedSamplersEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics, or compute operations
- [](#VUID-vkCmdBindDescriptorBufferEmbeddedSamplersEXT-videocoding)VUID-vkCmdBindDescriptorBufferEmbeddedSamplersEXT-videocoding  
  This command **must** only be called outside of a video coding scope
- [](#VUID-vkCmdBindDescriptorBufferEmbeddedSamplersEXT-commonparent)VUID-vkCmdBindDescriptorBufferEmbeddedSamplersEXT-commonparent  
  Both of `commandBuffer`, and `layout` **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

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

State

Conditional Rendering

vkCmdBindDescriptorBufferEmbeddedSamplersEXT is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_EXT\_descriptor\_buffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_descriptor_buffer.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBindPoint.html), [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayout.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdBindDescriptorBufferEmbeddedSamplersEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0