# vkCmdSetDiscardRectangleModeEXT(3) Manual Page

## Name

vkCmdSetDiscardRectangleModeEXT - Sets the discard rectangle mode dynamically for a command buffer



## [](#_c_specification)C Specification

To [dynamically set](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-dynamic-state) the discard rectangle mode, call:

```c++
// Provided by VK_EXT_discard_rectangles
void vkCmdSetDiscardRectangleModeEXT(
    VkCommandBuffer                             commandBuffer,
    VkDiscardRectangleModeEXT                   discardRectangleMode);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `discardRectangleMode` specifies the discard rectangle mode for all discard rectangles, either inclusive or exclusive.

## [](#_description)Description

This command sets the discard rectangle mode for subsequent drawing commands when drawing using [shader objects](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-objects), or when the graphics pipeline is created with `VK_DYNAMIC_STATE_DISCARD_RECTANGLE_MODE_EXT` set in [VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`. Otherwise, this state is specified by the [VkPipelineDiscardRectangleStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineDiscardRectangleStateCreateInfoEXT.html)::`discardRectangleMode` value used to create the currently active pipeline.

Valid Usage

- [](#VUID-vkCmdSetDiscardRectangleModeEXT-specVersion-07852)VUID-vkCmdSetDiscardRectangleModeEXT-specVersion-07852  
  The `VK_EXT_discard_rectangles` extension **must** be enabled, and the implementation **must** support at least `specVersion` `2` of this extension

Valid Usage (Implicit)

- [](#VUID-vkCmdSetDiscardRectangleModeEXT-commandBuffer-parameter)VUID-vkCmdSetDiscardRectangleModeEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdSetDiscardRectangleModeEXT-discardRectangleMode-parameter)VUID-vkCmdSetDiscardRectangleModeEXT-discardRectangleMode-parameter  
  `discardRectangleMode` **must** be a valid [VkDiscardRectangleModeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDiscardRectangleModeEXT.html) value
- [](#VUID-vkCmdSetDiscardRectangleModeEXT-commandBuffer-recording)VUID-vkCmdSetDiscardRectangleModeEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdSetDiscardRectangleModeEXT-commandBuffer-cmdpool)VUID-vkCmdSetDiscardRectangleModeEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support VK\_QUEUE\_GRAPHICS\_BIT operations
- [](#VUID-vkCmdSetDiscardRectangleModeEXT-videocoding)VUID-vkCmdSetDiscardRectangleModeEXT-videocoding  
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

VK\_QUEUE\_GRAPHICS\_BIT

State

Conditional Rendering

vkCmdSetDiscardRectangleModeEXT is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_EXT\_discard\_rectangles](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_discard_rectangles.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkDiscardRectangleModeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDiscardRectangleModeEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdSetDiscardRectangleModeEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0