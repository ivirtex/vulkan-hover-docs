# vkCmdEndRendering(3) Manual Page

## Name

vkCmdEndRendering - End a dynamic render pass instance



## [](#_c_specification)C Specification

To end a render pass instance, call:

```c++
// Provided by VK_VERSION_1_3
void vkCmdEndRendering(
    VkCommandBuffer                             commandBuffer);
```

or the equivalent command

```c++
// Provided by VK_KHR_dynamic_rendering
void vkCmdEndRenderingKHR(
    VkCommandBuffer                             commandBuffer);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer in which to record the command.

## [](#_description)Description

If the value of `pRenderingInfo->flags` used to begin this render pass instance included `VK_RENDERING_SUSPENDING_BIT`, then this render pass is suspended and will be resumed later in [submission order](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-submission-order).

Valid Usage

- [](#VUID-vkCmdEndRendering-None-06161)VUID-vkCmdEndRendering-None-06161  
  The current render pass instance **must** have been begun with [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBeginRendering.html)
- [](#VUID-vkCmdEndRendering-commandBuffer-06162)VUID-vkCmdEndRendering-commandBuffer-06162  
  The current render pass instance **must** have been begun in `commandBuffer`
- [](#VUID-vkCmdEndRendering-None-06781)VUID-vkCmdEndRendering-None-06781  
  This command **must** not be recorded when transform feedback is active
- [](#VUID-vkCmdEndRendering-None-06999)VUID-vkCmdEndRendering-None-06999  
  If `vkCmdBeginQuery`* was called within the render pass, the corresponding `vkCmdEndQuery`* **must** have been called subsequently within the same subpass
- [](#VUID-vkCmdEndRendering-None-10645)VUID-vkCmdEndRendering-None-10645  
  This command **must** not be recorded when [per-tile execution model](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-per-tile-execution-model) is enabled

Valid Usage (Implicit)

- [](#VUID-vkCmdEndRendering-commandBuffer-parameter)VUID-vkCmdEndRendering-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdEndRendering-commandBuffer-recording)VUID-vkCmdEndRendering-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdEndRendering-commandBuffer-cmdpool)VUID-vkCmdEndRendering-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics operations
- [](#VUID-vkCmdEndRendering-renderpass)VUID-vkCmdEndRendering-renderpass  
  This command **must** only be called inside of a render pass instance
- [](#VUID-vkCmdEndRendering-videocoding)VUID-vkCmdEndRendering-videocoding  
  This command **must** only be called outside of a video coding scope

Host Synchronization

- Host access to `commandBuffer` **must** be externally synchronized
- Host access to the `VkCommandPool` that `commandBuffer` was allocated from **must** be externally synchronized

Command Properties

     [Command Buffer Levels](#VkCommandBufferLevel) [Render Pass Scope](#vkCmdBeginRenderPass) [Video Coding Scope](#vkCmdBeginVideoCodingKHR) [Supported Queue Types](#VkQueueFlagBits) [Command Type](#fundamentals-queueoperation-command-types)

Primary  
Secondary

Inside

Outside

Graphics

Action  
State

## [](#_see_also)See Also

[VK\_KHR\_dynamic\_rendering](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_dynamic_rendering.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdEndRendering)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0