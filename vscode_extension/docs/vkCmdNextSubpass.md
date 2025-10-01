# vkCmdNextSubpass(3) Manual Page

## Name

vkCmdNextSubpass - Transition to the next subpass of a render pass



## [](#_c_specification)C Specification

To transition to the next subpass in the render pass instance after recording the commands for a subpass, call:

Warning

This functionality is deprecated by [Vulkan Version 1.2](#versions-1.2). See [Deprecated Functionality](#deprecation-renderpass2) for more information.

```c++
// Provided by VK_VERSION_1_0
void vkCmdNextSubpass(
    VkCommandBuffer                             commandBuffer,
    VkSubpassContents                           contents);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer in which to record the command.
- `contents` specifies how the commands in the next subpass will be provided, in the same fashion as the corresponding parameter of [vkCmdBeginRenderPass](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBeginRenderPass.html).

## [](#_description)Description

The subpass index for a render pass begins at zero when `vkCmdBeginRenderPass` is recorded, and increments each time `vkCmdNextSubpass` is recorded.

After transitioning to the next subpass, the application **can** record the commands for that subpass.

Valid Usage

- [](#VUID-vkCmdNextSubpass-None-00909)VUID-vkCmdNextSubpass-None-00909  
  The current subpass index **must** be less than the number of subpasses in the render pass minus one
- [](#VUID-vkCmdNextSubpass-None-02349)VUID-vkCmdNextSubpass-None-02349  
  This command **must** not be recorded when transform feedback is active

Valid Usage (Implicit)

- [](#VUID-vkCmdNextSubpass-commandBuffer-parameter)VUID-vkCmdNextSubpass-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdNextSubpass-contents-parameter)VUID-vkCmdNextSubpass-contents-parameter  
  `contents` **must** be a valid [VkSubpassContents](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassContents.html) value
- [](#VUID-vkCmdNextSubpass-commandBuffer-recording)VUID-vkCmdNextSubpass-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdNextSubpass-commandBuffer-cmdpool)VUID-vkCmdNextSubpass-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support VK\_QUEUE\_GRAPHICS\_BIT operations
- [](#VUID-vkCmdNextSubpass-renderpass)VUID-vkCmdNextSubpass-renderpass  
  This command **must** only be called inside of a render pass instance
- [](#VUID-vkCmdNextSubpass-videocoding)VUID-vkCmdNextSubpass-videocoding  
  This command **must** only be called outside of a video coding scope
- [](#VUID-vkCmdNextSubpass-bufferlevel)VUID-vkCmdNextSubpass-bufferlevel  
  `commandBuffer` **must** be a primary `VkCommandBuffer`

Host Synchronization

- Host access to `commandBuffer` **must** be externally synchronized
- Host access to the `VkCommandPool` that `commandBuffer` was allocated from **must** be externally synchronized

Command Properties

     [Command Buffer Levels](#VkCommandBufferLevel) [Render Pass Scope](#vkCmdBeginRenderPass) [Video Coding Scope](#vkCmdBeginVideoCodingKHR) [Supported Queue Types](#VkQueueFlagBits) [Command Type](#fundamentals-queueoperation-command-types)

Primary

Inside

Outside

VK\_QUEUE\_GRAPHICS\_BIT

Action  
State  
Synchronization

Conditional Rendering

vkCmdNextSubpass is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkSubpassContents](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassContents.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdNextSubpass).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0