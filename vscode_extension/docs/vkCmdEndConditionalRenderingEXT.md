# vkCmdEndConditionalRenderingEXT(3) Manual Page

## Name

vkCmdEndConditionalRenderingEXT - Define the end of a conditional rendering block



## [](#_c_specification)C Specification

To end conditional rendering, call:

```c++
// Provided by VK_EXT_conditional_rendering
void vkCmdEndConditionalRenderingEXT(
    VkCommandBuffer                             commandBuffer);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which this command will be recorded.

## [](#_description)Description

Once ended, conditional rendering becomes inactive.

Valid Usage

- [](#VUID-vkCmdEndConditionalRenderingEXT-None-01985)VUID-vkCmdEndConditionalRenderingEXT-None-01985  
  Conditional rendering **must** be [active](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#active-conditional-rendering)
- [](#VUID-vkCmdEndConditionalRenderingEXT-None-01986)VUID-vkCmdEndConditionalRenderingEXT-None-01986  
  If conditional rendering was made [active](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#active-conditional-rendering) outside of a render pass instance, it **must** not be ended inside a render pass instance
- [](#VUID-vkCmdEndConditionalRenderingEXT-None-01987)VUID-vkCmdEndConditionalRenderingEXT-None-01987  
  If conditional rendering was made [active](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#active-conditional-rendering) within a subpass it **must** be ended in the same subpass

Valid Usage (Implicit)

- [](#VUID-vkCmdEndConditionalRenderingEXT-commandBuffer-parameter)VUID-vkCmdEndConditionalRenderingEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdEndConditionalRenderingEXT-commandBuffer-recording)VUID-vkCmdEndConditionalRenderingEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdEndConditionalRenderingEXT-commandBuffer-cmdpool)VUID-vkCmdEndConditionalRenderingEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics, or compute operations
- [](#VUID-vkCmdEndConditionalRenderingEXT-videocoding)VUID-vkCmdEndConditionalRenderingEXT-videocoding  
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

Action  
State

Conditional Rendering

vkCmdEndConditionalRenderingEXT is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_EXT\_conditional\_rendering](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_conditional_rendering.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdEndConditionalRenderingEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0