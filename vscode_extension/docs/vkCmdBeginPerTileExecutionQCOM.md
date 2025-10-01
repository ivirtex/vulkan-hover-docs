# vkCmdBeginPerTileExecutionQCOM(3) Manual Page

## Name

vkCmdBeginPerTileExecutionQCOM - Begin per-tile execution mode



## [](#_c_specification)C Specification

To enable the per-tile execution model, call:

```c++
// Provided by VK_QCOM_tile_shading
void vkCmdBeginPerTileExecutionQCOM(
    VkCommandBuffer                             commandBuffer,
    const VkPerTileBeginInfoQCOM*               pPerTileBeginInfo);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer in which to record the command.
- `pPerTileBeginInfo` is a pointer to a [VkPerTileBeginInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerTileBeginInfoQCOM.html) structure containing information about how the *per-tile execution model* is started.

## [](#_description)Description

When *per-tile execution model* is enabled, recorded `vkCmdDraw*` or `vkCmdDispatch*` commands are invoked per tile. That is, the recorded draw or dispatch is invoked exactly once for each *covered tile*. The set of *covered tiles* for a given render pass instance consists of the set of render pass tiles, which **can** be queried with `VK_QCOM_tile_properties`, that are completely or partially covered by the `renderArea` for the render pass instance. The draw or dispatch commands **may** be invoked for uncovered tiles.

Each per-tile command invocation is associated with a single tile, the *active tile*. These per-tile invocations are not specified to execute in any particular order, but the size and offset of the *active tile* is available via shader built-ins.

When *per-tile execution model* is enabled, the following restrictions apply:

- Transform feedback commands such as [vkCmdBeginTransformFeedbackEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBeginTransformFeedbackEXT.html), [vkCmdEndTransformFeedbackEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdEndTransformFeedbackEXT.html), [vkCmdBeginQueryIndexedEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBeginQueryIndexedEXT.html), and [vkCmdEndQueryIndexedEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdEndQueryIndexedEXT.html), **must** not be recorded.
- Query commands such as [vkCmdWriteTimestamp](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdWriteTimestamp.html), [vkCmdDebugMarkerBeginEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDebugMarkerBeginEXT.html), [vkCmdDebugMarkerEndEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDebugMarkerEndEXT.html), [vkCmdDebugMarkerInsertEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDebugMarkerInsertEXT.html), [vkCmdBeginQuery](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBeginQuery.html), and [vkCmdEndQuery](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdEndQuery.html), **must** not be recorded.
- Event commands such as [vkCmdWaitEvents2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdWaitEvents2.html) and [vkCmdWaitEvents](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdWaitEvents.html) **must** not be recorded.
- Render pass clears like [vkCmdClearAttachments](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdClearAttachments.html) **must** not be recorded
- Access of an attachment with layout `VK_IMAGE_LAYOUT_ATTACHMENT_FEEDBACK_LOOP_OPTIMAL_EXT` as provided by `VK_EXT_attachment_feedback_loop_layout` is disallowed
- Any commands that would cause a invocations of one of the following shader stages are not allowed
  
  - tessellation
  - geometry
  - ray tracing
  - mesh shading

Valid Usage

- [](#VUID-vkCmdBeginPerTileExecutionQCOM-None-10664)VUID-vkCmdBeginPerTileExecutionQCOM-None-10664  
  The current render pass **must** be a [tile shading render pass](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-tile-shading)
- [](#VUID-vkCmdBeginPerTileExecutionQCOM-None-10665)VUID-vkCmdBeginPerTileExecutionQCOM-None-10665  
  The [tileShadingPerTileDispatch](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-tileShadingPerTileDispatch) or [tileShadingPerTileDraw](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-tileShadingPerTileDraw) feature must be enabled

Valid Usage (Implicit)

- [](#VUID-vkCmdBeginPerTileExecutionQCOM-commandBuffer-parameter)VUID-vkCmdBeginPerTileExecutionQCOM-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdBeginPerTileExecutionQCOM-pPerTileBeginInfo-parameter)VUID-vkCmdBeginPerTileExecutionQCOM-pPerTileBeginInfo-parameter  
  `pPerTileBeginInfo` **must** be a valid pointer to a valid [VkPerTileBeginInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerTileBeginInfoQCOM.html) structure
- [](#VUID-vkCmdBeginPerTileExecutionQCOM-commandBuffer-recording)VUID-vkCmdBeginPerTileExecutionQCOM-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdBeginPerTileExecutionQCOM-commandBuffer-cmdpool)VUID-vkCmdBeginPerTileExecutionQCOM-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support VK\_QUEUE\_COMPUTE\_BIT, or VK\_QUEUE\_GRAPHICS\_BIT operations
- [](#VUID-vkCmdBeginPerTileExecutionQCOM-renderpass)VUID-vkCmdBeginPerTileExecutionQCOM-renderpass  
  This command **must** only be called inside of a render pass instance
- [](#VUID-vkCmdBeginPerTileExecutionQCOM-videocoding)VUID-vkCmdBeginPerTileExecutionQCOM-videocoding  
  This command **must** only be called outside of a video coding scope

Host Synchronization

- Host access to the `VkCommandPool` that `commandBuffer` was allocated from **must** be externally synchronized

Command Properties

     [Command Buffer Levels](#VkCommandBufferLevel) [Render Pass Scope](#vkCmdBeginRenderPass) [Video Coding Scope](#vkCmdBeginVideoCodingKHR) [Supported Queue Types](#VkQueueFlagBits) [Command Type](#fundamentals-queueoperation-command-types)

Primary  
Secondary

Inside

Outside

VK\_QUEUE\_COMPUTE\_BIT  
VK\_QUEUE\_GRAPHICS\_BIT

State

Conditional Rendering

vkCmdBeginPerTileExecutionQCOM is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_QCOM\_tile\_shading](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_QCOM_tile_shading.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkPerTileBeginInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerTileBeginInfoQCOM.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdBeginPerTileExecutionQCOM).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0