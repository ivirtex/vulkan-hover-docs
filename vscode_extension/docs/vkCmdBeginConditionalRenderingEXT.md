# vkCmdBeginConditionalRenderingEXT(3) Manual Page

## Name

vkCmdBeginConditionalRenderingEXT - Define the beginning of a conditional rendering block



## [](#_c_specification)C Specification

To begin conditional rendering, call:

```c++
// Provided by VK_EXT_conditional_rendering
void vkCmdBeginConditionalRenderingEXT(
    VkCommandBuffer                             commandBuffer,
    const VkConditionalRenderingBeginInfoEXT*   pConditionalRenderingBegin);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which this command will be recorded.
- `pConditionalRenderingBegin` is a pointer to a [VkConditionalRenderingBeginInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkConditionalRenderingBeginInfoEXT.html) structure specifying parameters of conditional rendering.

## [](#_description)Description

Valid Usage

- [](#VUID-vkCmdBeginConditionalRenderingEXT-None-01980)VUID-vkCmdBeginConditionalRenderingEXT-None-01980  
  Conditional rendering **must** not already be [active](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#active-conditional-rendering)

Valid Usage (Implicit)

- [](#VUID-vkCmdBeginConditionalRenderingEXT-commandBuffer-parameter)VUID-vkCmdBeginConditionalRenderingEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdBeginConditionalRenderingEXT-pConditionalRenderingBegin-parameter)VUID-vkCmdBeginConditionalRenderingEXT-pConditionalRenderingBegin-parameter  
  `pConditionalRenderingBegin` **must** be a valid pointer to a valid [VkConditionalRenderingBeginInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkConditionalRenderingBeginInfoEXT.html) structure
- [](#VUID-vkCmdBeginConditionalRenderingEXT-commandBuffer-recording)VUID-vkCmdBeginConditionalRenderingEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdBeginConditionalRenderingEXT-commandBuffer-cmdpool)VUID-vkCmdBeginConditionalRenderingEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support VK\_QUEUE\_COMPUTE\_BIT, or VK\_QUEUE\_GRAPHICS\_BIT operations
- [](#VUID-vkCmdBeginConditionalRenderingEXT-videocoding)VUID-vkCmdBeginConditionalRenderingEXT-videocoding  
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

Action  
State

Conditional Rendering

vkCmdBeginConditionalRenderingEXT is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_EXT\_conditional\_rendering](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_conditional_rendering.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkConditionalRenderingBeginInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkConditionalRenderingBeginInfoEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdBeginConditionalRenderingEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0