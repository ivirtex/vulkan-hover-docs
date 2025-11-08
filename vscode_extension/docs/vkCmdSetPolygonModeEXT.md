# vkCmdSetPolygonModeEXT(3) Manual Page

## Name

vkCmdSetPolygonModeEXT - Specify polygon mode dynamically for a command buffer



## [](#_c_specification)C Specification

To [dynamically set](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-dynamic-state) the polygon mode, call:

```c++
// Provided by VK_EXT_extended_dynamic_state3, VK_EXT_shader_object
void vkCmdSetPolygonModeEXT(
    VkCommandBuffer                             commandBuffer,
    VkPolygonMode                               polygonMode);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `polygonMode` specifies polygon mode.

## [](#_description)Description

This command sets the polygon mode for subsequent drawing commands when drawing using [shader objects](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-objects), or when the graphics pipeline is created with `VK_DYNAMIC_STATE_POLYGON_MODE_EXT` set in [VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`. Otherwise, this state is specified by the [VkPipelineRasterizationStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRasterizationStateCreateInfo.html)::`polygonMode` value used to create the currently active pipeline.

Valid Usage

- [](#VUID-vkCmdSetPolygonModeEXT-None-09423)VUID-vkCmdSetPolygonModeEXT-None-09423  
  At least one of the following **must** be true:
  
  - The [`extendedDynamicState3PolygonMode`](#features-extendedDynamicState3PolygonMode) feature is enabled
  - The [`shaderObject`](#features-shaderObject) feature is enabled
- [](#VUID-vkCmdSetPolygonModeEXT-fillModeNonSolid-07424)VUID-vkCmdSetPolygonModeEXT-fillModeNonSolid-07424  
  If the [`fillModeNonSolid`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-fillModeNonSolid) feature is not enabled, `polygonMode` **must** be `VK_POLYGON_MODE_FILL` or `VK_POLYGON_MODE_FILL_RECTANGLE_NV`
- [](#VUID-vkCmdSetPolygonModeEXT-polygonMode-07425)VUID-vkCmdSetPolygonModeEXT-polygonMode-07425  
  If the `VK_NV_fill_rectangle` extension is not enabled, `polygonMode` **must** not be `VK_POLYGON_MODE_FILL_RECTANGLE_NV`

Valid Usage (Implicit)

- [](#VUID-vkCmdSetPolygonModeEXT-commandBuffer-parameter)VUID-vkCmdSetPolygonModeEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdSetPolygonModeEXT-polygonMode-parameter)VUID-vkCmdSetPolygonModeEXT-polygonMode-parameter  
  `polygonMode` **must** be a valid [VkPolygonMode](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPolygonMode.html) value
- [](#VUID-vkCmdSetPolygonModeEXT-commandBuffer-recording)VUID-vkCmdSetPolygonModeEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdSetPolygonModeEXT-commandBuffer-cmdpool)VUID-vkCmdSetPolygonModeEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support VK\_QUEUE\_GRAPHICS\_BIT operations
- [](#VUID-vkCmdSetPolygonModeEXT-videocoding)VUID-vkCmdSetPolygonModeEXT-videocoding  
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

vkCmdSetPolygonModeEXT is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_EXT\_extended\_dynamic\_state3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_extended_dynamic_state3.html), [VK\_EXT\_shader\_object](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_shader_object.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkPolygonMode](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPolygonMode.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdSetPolygonModeEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0