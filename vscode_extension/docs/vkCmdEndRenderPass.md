# vkCmdEndRenderPass(3) Manual Page

## Name

vkCmdEndRenderPass - End the current render pass



## [](#_c_specification)C Specification

To record a command to end a render pass instance after recording the commands for the last subpass, call:

Warning

This functionality is deprecated by [Vulkan Version 1.2](#versions-1.2). See [Deprecated Functionality](#deprecation-renderpass2) for more information.

```c++
// Provided by VK_VERSION_1_0
void vkCmdEndRenderPass(
    VkCommandBuffer                             commandBuffer);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer in which to end the current render pass instance.

## [](#_description)Description

Ending a render pass instance performs any multisample resolve operations on the final subpass.

Valid Usage

- [](#VUID-vkCmdEndRenderPass-None-00910)VUID-vkCmdEndRenderPass-None-00910  
  The current subpass index **must** be equal to the number of subpasses in the render pass minus one
- [](#VUID-vkCmdEndRenderPass-None-02351)VUID-vkCmdEndRenderPass-None-02351  
  This command **must** not be recorded when transform feedback is active
- [](#VUID-vkCmdEndRenderPass-None-06170)VUID-vkCmdEndRenderPass-None-06170  
  The current render pass instance **must** not have been begun with [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBeginRendering.html)
- [](#VUID-vkCmdEndRenderPass-None-07004)VUID-vkCmdEndRenderPass-None-07004  
  If `vkCmdBeginQuery`* was called within a subpass of the render pass, the corresponding `vkCmdEndQuery`* **must** have been called subsequently within the same subpass
- [](#VUID-vkCmdEndRenderPass-None-10653)VUID-vkCmdEndRenderPass-None-10653  
  This command **must** not be recorded when [per-tile execution model](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-per-tile-execution-model) is enabled

Valid Usage (Implicit)

- [](#VUID-vkCmdEndRenderPass-commandBuffer-parameter)VUID-vkCmdEndRenderPass-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdEndRenderPass-commandBuffer-recording)VUID-vkCmdEndRenderPass-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdEndRenderPass-commandBuffer-cmdpool)VUID-vkCmdEndRenderPass-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics operations
- [](#VUID-vkCmdEndRenderPass-renderpass)VUID-vkCmdEndRenderPass-renderpass  
  This command **must** only be called inside of a render pass instance
- [](#VUID-vkCmdEndRenderPass-videocoding)VUID-vkCmdEndRenderPass-videocoding  
  This command **must** only be called outside of a video coding scope
- [](#VUID-vkCmdEndRenderPass-bufferlevel)VUID-vkCmdEndRenderPass-bufferlevel  
  `commandBuffer` **must** be a primary `VkCommandBuffer`

Host Synchronization

- Host access to `commandBuffer` **must** be externally synchronized
- Host access to the `VkCommandPool` that `commandBuffer` was allocated from **must** be externally synchronized

Command Properties

     [Command Buffer Levels](#VkCommandBufferLevel) [Render Pass Scope](#vkCmdBeginRenderPass) [Video Coding Scope](#vkCmdBeginVideoCodingKHR) [Supported Queue Types](#VkQueueFlagBits) [Command Type](#fundamentals-queueoperation-command-types)

Primary

Inside

Outside

Graphics

Action  
State  
Synchronization

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdEndRenderPass)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0