# vkCmdSetProvokingVertexModeEXT(3) Manual Page

## Name

vkCmdSetProvokingVertexModeEXT - Specify the provoking vertex mode dynamically for a command buffer



## [](#_c_specification)C Specification

To [dynamically set](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-dynamic-state) the `provokingVertexMode` state, call:

```c++
// Provided by VK_EXT_extended_dynamic_state3 with VK_EXT_provoking_vertex, VK_EXT_provoking_vertex with VK_EXT_shader_object
void vkCmdSetProvokingVertexModeEXT(
    VkCommandBuffer                             commandBuffer,
    VkProvokingVertexModeEXT                    provokingVertexMode);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `provokingVertexMode` specifies the `provokingVertexMode` state.

## [](#_description)Description

This command sets the `provokingVertexMode` state for subsequent drawing commands when drawing using [shader objects](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-objects), or when the graphics pipeline is created with `VK_DYNAMIC_STATE_PROVOKING_VERTEX_MODE_EXT` set in [VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`. Otherwise, this state is specified by the [VkPipelineRasterizationProvokingVertexStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRasterizationProvokingVertexStateCreateInfoEXT.html)::`provokingVertexMode` value used to create the currently active pipeline.

Valid Usage

- [](#VUID-vkCmdSetProvokingVertexModeEXT-None-09423)VUID-vkCmdSetProvokingVertexModeEXT-None-09423  
  At least one of the following **must** be true:
  
  - The [`extendedDynamicState3ProvokingVertexMode`](#features-extendedDynamicState3ProvokingVertexMode) feature is enabled
  - The [`shaderObject`](#features-shaderObject) feature is enabled
- [](#VUID-vkCmdSetProvokingVertexModeEXT-provokingVertexMode-07447)VUID-vkCmdSetProvokingVertexModeEXT-provokingVertexMode-07447  
  If `provokingVertexMode` is `VK_PROVOKING_VERTEX_MODE_LAST_VERTEX_EXT`, then the [`provokingVertexLast`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-provokingVertexLast) feature **must** be enabled

Valid Usage (Implicit)

- [](#VUID-vkCmdSetProvokingVertexModeEXT-commandBuffer-parameter)VUID-vkCmdSetProvokingVertexModeEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdSetProvokingVertexModeEXT-provokingVertexMode-parameter)VUID-vkCmdSetProvokingVertexModeEXT-provokingVertexMode-parameter  
  `provokingVertexMode` **must** be a valid [VkProvokingVertexModeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkProvokingVertexModeEXT.html) value
- [](#VUID-vkCmdSetProvokingVertexModeEXT-commandBuffer-recording)VUID-vkCmdSetProvokingVertexModeEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdSetProvokingVertexModeEXT-commandBuffer-cmdpool)VUID-vkCmdSetProvokingVertexModeEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support VK\_QUEUE\_GRAPHICS\_BIT operations
- [](#VUID-vkCmdSetProvokingVertexModeEXT-videocoding)VUID-vkCmdSetProvokingVertexModeEXT-videocoding  
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

vkCmdSetProvokingVertexModeEXT is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_EXT\_extended\_dynamic\_state3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_extended_dynamic_state3.html), [VK\_EXT\_provoking\_vertex](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_provoking_vertex.html), [VK\_EXT\_shader\_object](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_shader_object.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkProvokingVertexModeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkProvokingVertexModeEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdSetProvokingVertexModeEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0