# vkCmdSetLineRasterizationModeEXT(3) Manual Page

## Name

vkCmdSetLineRasterizationModeEXT - Specify the line rasterization mode dynamically for a command buffer



## [](#_c_specification)C Specification

To [dynamically set](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-dynamic-state) the `lineRasterizationMode` state, call:

```c++
// Provided by VK_EXT_extended_dynamic_state3 with VK_EXT_line_rasterization, VK_EXT_line_rasterization with VK_EXT_shader_object
void vkCmdSetLineRasterizationModeEXT(
    VkCommandBuffer                             commandBuffer,
    VkLineRasterizationModeEXT                  lineRasterizationMode);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `lineRasterizationMode` specifies the `lineRasterizationMode` state.

## [](#_description)Description

This command sets the `lineRasterizationMode` state for subsequent drawing commands when drawing using [shader objects](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-objects), or when the graphics pipeline is created with `VK_DYNAMIC_STATE_LINE_RASTERIZATION_MODE_EXT` set in [VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`. Otherwise, this state is specified by the [VkPipelineRasterizationLineStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRasterizationLineStateCreateInfo.html)::`lineRasterizationMode` value used to create the currently active pipeline.

Valid Usage

- [](#VUID-vkCmdSetLineRasterizationModeEXT-None-09423)VUID-vkCmdSetLineRasterizationModeEXT-None-09423  
  At least one of the following **must** be true:
  
  - The [`extendedDynamicState3LineRasterizationMode`](#features-extendedDynamicState3LineRasterizationMode) feature is enabled
  - The [`shaderObject`](#features-shaderObject) feature is enabled
- [](#VUID-vkCmdSetLineRasterizationModeEXT-lineRasterizationMode-07418)VUID-vkCmdSetLineRasterizationModeEXT-lineRasterizationMode-07418  
  If `lineRasterizationMode` is `VK_LINE_RASTERIZATION_MODE_RECTANGULAR`, then the [`rectangularLines`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-rectangularLines) feature **must** be enabled
- [](#VUID-vkCmdSetLineRasterizationModeEXT-lineRasterizationMode-07419)VUID-vkCmdSetLineRasterizationModeEXT-lineRasterizationMode-07419  
  If `lineRasterizationMode` is `VK_LINE_RASTERIZATION_MODE_BRESENHAM`, then the [`bresenhamLines`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-bresenhamLines) feature **must** be enabled
- [](#VUID-vkCmdSetLineRasterizationModeEXT-lineRasterizationMode-07420)VUID-vkCmdSetLineRasterizationModeEXT-lineRasterizationMode-07420  
  If `lineRasterizationMode` is `VK_LINE_RASTERIZATION_MODE_RECTANGULAR_SMOOTH`, then the [`smoothLines`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-smoothLines) feature **must** be enabled

Valid Usage (Implicit)

- [](#VUID-vkCmdSetLineRasterizationModeEXT-commandBuffer-parameter)VUID-vkCmdSetLineRasterizationModeEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdSetLineRasterizationModeEXT-lineRasterizationMode-parameter)VUID-vkCmdSetLineRasterizationModeEXT-lineRasterizationMode-parameter  
  `lineRasterizationMode` **must** be a valid [VkLineRasterizationModeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkLineRasterizationModeEXT.html) value
- [](#VUID-vkCmdSetLineRasterizationModeEXT-commandBuffer-recording)VUID-vkCmdSetLineRasterizationModeEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdSetLineRasterizationModeEXT-commandBuffer-cmdpool)VUID-vkCmdSetLineRasterizationModeEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics operations
- [](#VUID-vkCmdSetLineRasterizationModeEXT-videocoding)VUID-vkCmdSetLineRasterizationModeEXT-videocoding  
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

Graphics

State

Conditional Rendering

vkCmdSetLineRasterizationModeEXT is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_EXT\_extended\_dynamic\_state3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_extended_dynamic_state3.html), [VK\_EXT\_line\_rasterization](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_line_rasterization.html), [VK\_EXT\_shader\_object](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_shader_object.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkLineRasterizationMode](https://registry.khronos.org/vulkan/specs/latest/man/html/VkLineRasterizationMode.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdSetLineRasterizationModeEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0