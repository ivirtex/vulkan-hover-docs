# vkCmdSetExclusiveScissorNV(3) Manual Page

## Name

vkCmdSetExclusiveScissorNV - Set exclusive scissor rectangles dynamically for a command buffer



## [](#_c_specification)C Specification

To [dynamically set](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-dynamic-state) the exclusive scissor rectangles, call:

```c++
// Provided by VK_NV_scissor_exclusive
void vkCmdSetExclusiveScissorNV(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    firstExclusiveScissor,
    uint32_t                                    exclusiveScissorCount,
    const VkRect2D*                             pExclusiveScissors);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `firstExclusiveScissor` is the index of the first exclusive scissor rectangle whose state is updated by the command.
- `exclusiveScissorCount` is the number of exclusive scissor rectangles updated by the command.
- `pExclusiveScissors` is a pointer to an array of [VkRect2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRect2D.html) structures defining exclusive scissor rectangles.

## [](#_description)Description

The scissor rectangles taken from element i of `pExclusiveScissors` replace the current state for the scissor index `firstExclusiveScissor` + i, for i in [0, `exclusiveScissorCount`).

This command sets the exclusive scissor rectangles for subsequent drawing commands when drawing using [shader objects](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-objects), or when the graphics pipeline is created with `VK_DYNAMIC_STATE_EXCLUSIVE_SCISSOR_NV` set in [VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`. Otherwise, this state is specified by the [VkPipelineViewportExclusiveScissorStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineViewportExclusiveScissorStateCreateInfoNV.html)::`pExclusiveScissors` values used to create the currently active pipeline.

Valid Usage

- [](#VUID-vkCmdSetExclusiveScissorNV-None-02031)VUID-vkCmdSetExclusiveScissorNV-None-02031  
  The [`exclusiveScissor`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-exclusiveScissor) feature **must** be enabled
- [](#VUID-vkCmdSetExclusiveScissorNV-firstExclusiveScissor-02034)VUID-vkCmdSetExclusiveScissorNV-firstExclusiveScissor-02034  
  The sum of `firstExclusiveScissor` and `exclusiveScissorCount` **must** be between `1` and `VkPhysicalDeviceLimits`::`maxViewports`, inclusive
- [](#VUID-vkCmdSetExclusiveScissorNV-firstExclusiveScissor-02035)VUID-vkCmdSetExclusiveScissorNV-firstExclusiveScissor-02035  
  If the [`multiViewport`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-multiViewport) feature is not enabled, `firstExclusiveScissor` **must** be `0`
- [](#VUID-vkCmdSetExclusiveScissorNV-exclusiveScissorCount-02036)VUID-vkCmdSetExclusiveScissorNV-exclusiveScissorCount-02036  
  If the [`multiViewport`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-multiViewport) feature is not enabled, `exclusiveScissorCount` **must** be `1`
- [](#VUID-vkCmdSetExclusiveScissorNV-x-02037)VUID-vkCmdSetExclusiveScissorNV-x-02037  
  The `x` and `y` members of `offset` in each member of `pExclusiveScissors` **must** be greater than or equal to `0`
- [](#VUID-vkCmdSetExclusiveScissorNV-offset-02038)VUID-vkCmdSetExclusiveScissorNV-offset-02038  
  Evaluation of (`offset.x` + `extent.width`) for each member of `pExclusiveScissors` **must** not cause a signed integer addition overflow
- [](#VUID-vkCmdSetExclusiveScissorNV-offset-02039)VUID-vkCmdSetExclusiveScissorNV-offset-02039  
  Evaluation of (`offset.y` + `extent.height`) for each member of `pExclusiveScissors` **must** not cause a signed integer addition overflow

Valid Usage (Implicit)

- [](#VUID-vkCmdSetExclusiveScissorNV-commandBuffer-parameter)VUID-vkCmdSetExclusiveScissorNV-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdSetExclusiveScissorNV-pExclusiveScissors-parameter)VUID-vkCmdSetExclusiveScissorNV-pExclusiveScissors-parameter  
  `pExclusiveScissors` **must** be a valid pointer to an array of `exclusiveScissorCount` [VkRect2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRect2D.html) structures
- [](#VUID-vkCmdSetExclusiveScissorNV-commandBuffer-recording)VUID-vkCmdSetExclusiveScissorNV-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdSetExclusiveScissorNV-commandBuffer-cmdpool)VUID-vkCmdSetExclusiveScissorNV-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics operations
- [](#VUID-vkCmdSetExclusiveScissorNV-videocoding)VUID-vkCmdSetExclusiveScissorNV-videocoding  
  This command **must** only be called outside of a video coding scope
- [](#VUID-vkCmdSetExclusiveScissorNV-exclusiveScissorCount-arraylength)VUID-vkCmdSetExclusiveScissorNV-exclusiveScissorCount-arraylength  
  `exclusiveScissorCount` **must** be greater than `0`

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

## [](#_see_also)See Also

[VK\_NV\_scissor\_exclusive](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_scissor_exclusive.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkRect2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRect2D.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdSetExclusiveScissorNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0