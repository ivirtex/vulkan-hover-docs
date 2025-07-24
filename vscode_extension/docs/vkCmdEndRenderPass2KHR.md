# vkCmdEndRenderPass2(3) Manual Page

## Name

vkCmdEndRenderPass2 - End the current render pass



## [](#_c_specification)C Specification

To record a command to end a render pass instance after recording the commands for the last subpass, call:

Warning

This functionality is deprecated by [Vulkan Version 1.4](#versions-1.4). See [Deprecated Functionality](#deprecation-dynamicrendering) for more information.

```c++
// Provided by VK_VERSION_1_2
void vkCmdEndRenderPass2(
    VkCommandBuffer                             commandBuffer,
    const VkSubpassEndInfo*                     pSubpassEndInfo);
```

or the equivalent command

```c++
// Provided by VK_KHR_create_renderpass2
void vkCmdEndRenderPass2KHR(
    VkCommandBuffer                             commandBuffer,
    const VkSubpassEndInfo*                     pSubpassEndInfo);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer in which to end the current render pass instance.
- `pSubpassEndInfo` is a pointer to a [VkSubpassEndInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassEndInfo.html) structure containing information about how the last subpass will be ended.

## [](#_description)Description

`vkCmdEndRenderPass2` is semantically identical to [vkCmdEndRenderPass](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdEndRenderPass.html), except that it is extensible.

Valid Usage

- [](#VUID-vkCmdEndRenderPass2-None-03103)VUID-vkCmdEndRenderPass2-None-03103  
  The current subpass index **must** be equal to the number of subpasses in the render pass minus one
- [](#VUID-vkCmdEndRenderPass2-None-02352)VUID-vkCmdEndRenderPass2-None-02352  
  This command **must** not be recorded when transform feedback is active
- [](#VUID-vkCmdEndRenderPass2-None-06171)VUID-vkCmdEndRenderPass2-None-06171  
  The current render pass instance **must** not have been begun with [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBeginRendering.html)
- [](#VUID-vkCmdEndRenderPass2-None-07005)VUID-vkCmdEndRenderPass2-None-07005  
  If `vkCmdBeginQuery`* was called within a subpass of the render pass, the corresponding `vkCmdEndQuery`* **must** have been called subsequently within the same subpass

Valid Usage (Implicit)

- [](#VUID-vkCmdEndRenderPass2-commandBuffer-parameter)VUID-vkCmdEndRenderPass2-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdEndRenderPass2-pSubpassEndInfo-parameter)VUID-vkCmdEndRenderPass2-pSubpassEndInfo-parameter  
  `pSubpassEndInfo` **must** be a valid pointer to a valid [VkSubpassEndInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassEndInfo.html) structure
- [](#VUID-vkCmdEndRenderPass2-commandBuffer-recording)VUID-vkCmdEndRenderPass2-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdEndRenderPass2-commandBuffer-cmdpool)VUID-vkCmdEndRenderPass2-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics operations
- [](#VUID-vkCmdEndRenderPass2-renderpass)VUID-vkCmdEndRenderPass2-renderpass  
  This command **must** only be called inside of a render pass instance
- [](#VUID-vkCmdEndRenderPass2-videocoding)VUID-vkCmdEndRenderPass2-videocoding  
  This command **must** only be called outside of a video coding scope
- [](#VUID-vkCmdEndRenderPass2-bufferlevel)VUID-vkCmdEndRenderPass2-bufferlevel  
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

Conditional Rendering

vkCmdEndRenderPass2 is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_KHR\_create\_renderpass2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_create_renderpass2.html), [VK\_VERSION\_1\_2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_2.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkSubpassEndInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassEndInfo.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdEndRenderPass2)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0