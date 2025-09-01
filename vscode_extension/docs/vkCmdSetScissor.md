# vkCmdSetScissor(3) Manual Page

## Name

vkCmdSetScissor - Set scissor rectangles dynamically for a command buffer



## [](#_c_specification)C Specification

To [dynamically set](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-dynamic-state) the scissor rectangles, call:

```c++
// Provided by VK_VERSION_1_0
void vkCmdSetScissor(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    firstScissor,
    uint32_t                                    scissorCount,
    const VkRect2D*                             pScissors);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `firstScissor` is the index of the first scissor whose state is updated by the command.
- `scissorCount` is the number of scissors whose rectangles are updated by the command.
- `pScissors` is a pointer to an array of [VkRect2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRect2D.html) structures defining scissor rectangles.

## [](#_description)Description

The scissor rectangles taken from element i of `pScissors` replace the current state for the scissor index `firstScissor` + i, for i in [0, `scissorCount`).

This command sets the scissor rectangles for subsequent drawing commands when drawing using [shader objects](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-objects), or when the graphics pipeline is created with `VK_DYNAMIC_STATE_SCISSOR` set in [VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`. Otherwise, this state is specified by the [VkPipelineViewportStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineViewportStateCreateInfo.html)::`pScissors` values used to create the currently active pipeline.

Valid Usage

- [](#VUID-vkCmdSetScissor-firstScissor-00592)VUID-vkCmdSetScissor-firstScissor-00592  
  The sum of `firstScissor` and `scissorCount` **must** be between `1` and `VkPhysicalDeviceLimits`::`maxViewports`, inclusive
- [](#VUID-vkCmdSetScissor-firstScissor-00593)VUID-vkCmdSetScissor-firstScissor-00593  
  If the [`multiViewport`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-multiViewport) feature is not enabled, `firstScissor` **must** be `0`
- [](#VUID-vkCmdSetScissor-scissorCount-00594)VUID-vkCmdSetScissor-scissorCount-00594  
  If the [`multiViewport`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-multiViewport) feature is not enabled, `scissorCount` **must** be `1`
- [](#VUID-vkCmdSetScissor-x-00595)VUID-vkCmdSetScissor-x-00595  
  The `x` and `y` members of `offset` member of any element of `pScissors` **must** be greater than or equal to `0`
- [](#VUID-vkCmdSetScissor-offset-00596)VUID-vkCmdSetScissor-offset-00596  
  Evaluation of (`offset.x` + `extent.width`) **must** not cause a signed integer addition overflow for any element of `pScissors`
- [](#VUID-vkCmdSetScissor-offset-00597)VUID-vkCmdSetScissor-offset-00597  
  Evaluation of (`offset.y` + `extent.height`) **must** not cause a signed integer addition overflow for any element of `pScissors`
- [](#VUID-vkCmdSetScissor-viewportScissor2D-04789)VUID-vkCmdSetScissor-viewportScissor2D-04789  
  If this command is recorded in a secondary command buffer with [VkCommandBufferInheritanceViewportScissorInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferInheritanceViewportScissorInfoNV.html)::`viewportScissor2D` enabled, then this function **must** not be called

Valid Usage (Implicit)

- [](#VUID-vkCmdSetScissor-commandBuffer-parameter)VUID-vkCmdSetScissor-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdSetScissor-pScissors-parameter)VUID-vkCmdSetScissor-pScissors-parameter  
  `pScissors` **must** be a valid pointer to an array of `scissorCount` [VkRect2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRect2D.html) structures
- [](#VUID-vkCmdSetScissor-commandBuffer-recording)VUID-vkCmdSetScissor-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdSetScissor-commandBuffer-cmdpool)VUID-vkCmdSetScissor-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics operations
- [](#VUID-vkCmdSetScissor-videocoding)VUID-vkCmdSetScissor-videocoding  
  This command **must** only be called outside of a video coding scope
- [](#VUID-vkCmdSetScissor-scissorCount-arraylength)VUID-vkCmdSetScissor-scissorCount-arraylength  
  `scissorCount` **must** be greater than `0`

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

vkCmdSetScissor is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkRect2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRect2D.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdSetScissor).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0