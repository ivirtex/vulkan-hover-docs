# vkCmdSetAttachmentFeedbackLoopEnableEXT(3) Manual Page

## Name

vkCmdSetAttachmentFeedbackLoopEnableEXT - Specify whether attachment feedback loops are enabled dynamically on a command buffer



## [](#_c_specification)C Specification

To [dynamically set](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-dynamic-state) whether a pipeline **can** access a resource as a non-attachment while it is also used as an attachment that is written to, call:

```c++
// Provided by VK_EXT_attachment_feedback_loop_dynamic_state
void vkCmdSetAttachmentFeedbackLoopEnableEXT(
    VkCommandBuffer                             commandBuffer,
    VkImageAspectFlags                          aspectMask);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `aspectMask` specifies the types of attachments for which feedback loops will be enabled. Attachment types whose aspects are not included in `aspectMask` will have feedback loops disabled.

## [](#_description)Description

For attachments that are written to in a render pass, only attachments with the aspects specified in `aspectMask` **can** be accessed as non-attachments by subsequent [drawing commands](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#drawing).

Valid Usage

- [](#VUID-vkCmdSetAttachmentFeedbackLoopEnableEXT-attachmentFeedbackLoopDynamicState-08862)VUID-vkCmdSetAttachmentFeedbackLoopEnableEXT-attachmentFeedbackLoopDynamicState-08862  
  The [`attachmentFeedbackLoopDynamicState`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-attachmentFeedbackLoopDynamicState) feature **must** be enabled
- [](#VUID-vkCmdSetAttachmentFeedbackLoopEnableEXT-aspectMask-08863)VUID-vkCmdSetAttachmentFeedbackLoopEnableEXT-aspectMask-08863  
  `aspectMask` **must** only include `VK_IMAGE_ASPECT_NONE`, `VK_IMAGE_ASPECT_COLOR_BIT`, `VK_IMAGE_ASPECT_DEPTH_BIT`, and `VK_IMAGE_ASPECT_STENCIL_BIT`
- [](#VUID-vkCmdSetAttachmentFeedbackLoopEnableEXT-attachmentFeedbackLoopLayout-08864)VUID-vkCmdSetAttachmentFeedbackLoopEnableEXT-attachmentFeedbackLoopLayout-08864  
  If the [`attachmentFeedbackLoopLayout`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-attachmentFeedbackLoopLayout) feature is not enabled, `aspectMask` **must** be `VK_IMAGE_ASPECT_NONE`

Valid Usage (Implicit)

- [](#VUID-vkCmdSetAttachmentFeedbackLoopEnableEXT-commandBuffer-parameter)VUID-vkCmdSetAttachmentFeedbackLoopEnableEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdSetAttachmentFeedbackLoopEnableEXT-aspectMask-parameter)VUID-vkCmdSetAttachmentFeedbackLoopEnableEXT-aspectMask-parameter  
  `aspectMask` **must** be a valid combination of [VkImageAspectFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageAspectFlagBits.html) values
- [](#VUID-vkCmdSetAttachmentFeedbackLoopEnableEXT-commandBuffer-recording)VUID-vkCmdSetAttachmentFeedbackLoopEnableEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdSetAttachmentFeedbackLoopEnableEXT-commandBuffer-cmdpool)VUID-vkCmdSetAttachmentFeedbackLoopEnableEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics operations
- [](#VUID-vkCmdSetAttachmentFeedbackLoopEnableEXT-videocoding)VUID-vkCmdSetAttachmentFeedbackLoopEnableEXT-videocoding  
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

State

Conditional Rendering

vkCmdSetAttachmentFeedbackLoopEnableEXT is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_EXT\_attachment\_feedback\_loop\_dynamic\_state](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_attachment_feedback_loop_dynamic_state.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkImageAspectFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageAspectFlags.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdSetAttachmentFeedbackLoopEnableEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0