# vkCmdNextSubpass2(3) Manual Page

## Name

vkCmdNextSubpass2 - Transition to the next subpass of a render pass



## [](#_c_specification)C Specification

To transition to the next subpass in the render pass instance after recording the commands for a subpass, call:

Warning

This functionality is deprecated by [Vulkan Version 1.4](#versions-1.4). See [Deprecated Functionality](#deprecation-dynamicrendering) for more information.

```c++
// Provided by VK_VERSION_1_2
void vkCmdNextSubpass2(
    VkCommandBuffer                             commandBuffer,
    const VkSubpassBeginInfo*                   pSubpassBeginInfo,
    const VkSubpassEndInfo*                     pSubpassEndInfo);
```

or the equivalent command

```c++
// Provided by VK_KHR_create_renderpass2
void vkCmdNextSubpass2KHR(
    VkCommandBuffer                             commandBuffer,
    const VkSubpassBeginInfo*                   pSubpassBeginInfo,
    const VkSubpassEndInfo*                     pSubpassEndInfo);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer in which to record the command.
- `pSubpassBeginInfo` is a pointer to a [VkSubpassBeginInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassBeginInfo.html) structure containing information about the subpass which is about to begin rendering.
- `pSubpassEndInfo` is a pointer to a [VkSubpassEndInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassEndInfo.html) structure containing information about how the previous subpass will be ended.

## [](#_description)Description

`vkCmdNextSubpass2` is semantically identical to [vkCmdNextSubpass](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdNextSubpass.html), except that it is extensible, and that `contents` is provided as part of an extensible structure instead of as a flat parameter.

Valid Usage

- [](#VUID-vkCmdNextSubpass2-None-03102)VUID-vkCmdNextSubpass2-None-03102  
  The current subpass index **must** be less than the number of subpasses in the render pass minus one
- [](#VUID-vkCmdNextSubpass2-None-02350)VUID-vkCmdNextSubpass2-None-02350  
  This command **must** not be recorded when transform feedback is active

Valid Usage (Implicit)

- [](#VUID-vkCmdNextSubpass2-commandBuffer-parameter)VUID-vkCmdNextSubpass2-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdNextSubpass2-pSubpassBeginInfo-parameter)VUID-vkCmdNextSubpass2-pSubpassBeginInfo-parameter  
  `pSubpassBeginInfo` **must** be a valid pointer to a valid [VkSubpassBeginInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassBeginInfo.html) structure
- [](#VUID-vkCmdNextSubpass2-pSubpassEndInfo-parameter)VUID-vkCmdNextSubpass2-pSubpassEndInfo-parameter  
  `pSubpassEndInfo` **must** be a valid pointer to a valid [VkSubpassEndInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassEndInfo.html) structure
- [](#VUID-vkCmdNextSubpass2-commandBuffer-recording)VUID-vkCmdNextSubpass2-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdNextSubpass2-commandBuffer-cmdpool)VUID-vkCmdNextSubpass2-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support VK\_QUEUE\_GRAPHICS\_BIT operations
- [](#VUID-vkCmdNextSubpass2-renderpass)VUID-vkCmdNextSubpass2-renderpass  
  This command **must** only be called inside of a render pass instance
- [](#VUID-vkCmdNextSubpass2-videocoding)VUID-vkCmdNextSubpass2-videocoding  
  This command **must** only be called outside of a video coding scope
- [](#VUID-vkCmdNextSubpass2-bufferlevel)VUID-vkCmdNextSubpass2-bufferlevel  
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

vkCmdNextSubpass2 is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_KHR\_create\_renderpass2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_create_renderpass2.html), [VK\_VERSION\_1\_2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_2.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkSubpassBeginInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassBeginInfo.html), [VkSubpassEndInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassEndInfo.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdNextSubpass2).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0