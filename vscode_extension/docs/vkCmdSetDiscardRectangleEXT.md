# vkCmdSetDiscardRectangleEXT(3) Manual Page

## Name

vkCmdSetDiscardRectangleEXT - Set discard rectangles dynamically for a command buffer



## [](#_c_specification)C Specification

To [dynamically set](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-dynamic-state) the discard rectangles, call:

```c++
// Provided by VK_EXT_discard_rectangles
void vkCmdSetDiscardRectangleEXT(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    firstDiscardRectangle,
    uint32_t                                    discardRectangleCount,
    const VkRect2D*                             pDiscardRectangles);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `firstDiscardRectangle` is the index of the first discard rectangle whose state is updated by the command.
- `discardRectangleCount` is the number of discard rectangles whose state are updated by the command.
- `pDiscardRectangles` is a pointer to an array of [VkRect2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRect2D.html) structures specifying discard rectangles.

## [](#_description)Description

The discard rectangle taken from element i of `pDiscardRectangles` replace the current state for the discard rectangle at index `firstDiscardRectangle` + i, for i in [0, `discardRectangleCount`).

This command sets the discard rectangles for subsequent drawing commands when drawing using [shader objects](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-objects), or when the graphics pipeline is created with `VK_DYNAMIC_STATE_DISCARD_RECTANGLE_EXT` set in [VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`. Otherwise, this state is specified by the [VkPipelineDiscardRectangleStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineDiscardRectangleStateCreateInfoEXT.html)::`pDiscardRectangles` values used to create the currently active pipeline.

Valid Usage

- [](#VUID-vkCmdSetDiscardRectangleEXT-firstDiscardRectangle-00585)VUID-vkCmdSetDiscardRectangleEXT-firstDiscardRectangle-00585  
  The sum of `firstDiscardRectangle` and `discardRectangleCount` **must** be less than or equal to [VkPhysicalDeviceDiscardRectanglePropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDiscardRectanglePropertiesEXT.html)::`maxDiscardRectangles`
- [](#VUID-vkCmdSetDiscardRectangleEXT-x-00587)VUID-vkCmdSetDiscardRectangleEXT-x-00587  
  The `x` and `y` member of `offset` in each [VkRect2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRect2D.html) element of `pDiscardRectangles` **must** be greater than or equal to `0`
- [](#VUID-vkCmdSetDiscardRectangleEXT-offset-00588)VUID-vkCmdSetDiscardRectangleEXT-offset-00588  
  Evaluation of (`offset.x` + `extent.width`) in each [VkRect2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRect2D.html) element of `pDiscardRectangles` **must** not cause a signed integer addition overflow
- [](#VUID-vkCmdSetDiscardRectangleEXT-offset-00589)VUID-vkCmdSetDiscardRectangleEXT-offset-00589  
  Evaluation of (`offset.y` + `extent.height`) in each [VkRect2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRect2D.html) element of `pDiscardRectangles` **must** not cause a signed integer addition overflow
- [](#VUID-vkCmdSetDiscardRectangleEXT-viewportScissor2D-04788)VUID-vkCmdSetDiscardRectangleEXT-viewportScissor2D-04788  
  If this command is recorded in a secondary command buffer with [VkCommandBufferInheritanceViewportScissorInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferInheritanceViewportScissorInfoNV.html)::`viewportScissor2D` enabled, then this function **must** not be called

Valid Usage (Implicit)

- [](#VUID-vkCmdSetDiscardRectangleEXT-commandBuffer-parameter)VUID-vkCmdSetDiscardRectangleEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdSetDiscardRectangleEXT-pDiscardRectangles-parameter)VUID-vkCmdSetDiscardRectangleEXT-pDiscardRectangles-parameter  
  `pDiscardRectangles` **must** be a valid pointer to an array of `discardRectangleCount` [VkRect2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRect2D.html) structures
- [](#VUID-vkCmdSetDiscardRectangleEXT-commandBuffer-recording)VUID-vkCmdSetDiscardRectangleEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdSetDiscardRectangleEXT-commandBuffer-cmdpool)VUID-vkCmdSetDiscardRectangleEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support VK\_QUEUE\_GRAPHICS\_BIT operations
- [](#VUID-vkCmdSetDiscardRectangleEXT-videocoding)VUID-vkCmdSetDiscardRectangleEXT-videocoding  
  This command **must** only be called outside of a video coding scope
- [](#VUID-vkCmdSetDiscardRectangleEXT-discardRectangleCount-arraylength)VUID-vkCmdSetDiscardRectangleEXT-discardRectangleCount-arraylength  
  `discardRectangleCount` **must** be greater than `0`

Host Synchronization

- Host access to `commandBuffer` **must** be externally synchronized
- Host access to the `VkCommandPool` that `commandBuffer` was allocated from **must** be externally synchronized

Command Properties

     [Command Buffer Levels](#VkCommandBufferLevel) [Render Pass Scope](#vkCmdBeginRenderPass) [Video Coding Scope](#vkCmdBeginVideoCodingKHR) [Supported Queue Types](#VkQueueFlagBits) [Command Type](#fundamentals-queueoperation-command-types)

Primary  
Secondary

Both

Outside

VK\_QUEUE\_GRAPHICS\_BIT

State

Conditional Rendering

vkCmdSetDiscardRectangleEXT is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_EXT\_discard\_rectangles](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_discard_rectangles.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkRect2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRect2D.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdSetDiscardRectangleEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0