# vkCmdSetLineStippleEnableEXT(3) Manual Page

## Name

vkCmdSetLineStippleEnableEXT - Specify the line stipple enable dynamically for a command buffer



## [](#_c_specification)C Specification

To [dynamically set](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-dynamic-state) the `stippledLineEnable` state, call:

```c++
// Provided by VK_EXT_extended_dynamic_state3 with VK_EXT_line_rasterization, VK_EXT_line_rasterization with VK_EXT_shader_object
void vkCmdSetLineStippleEnableEXT(
    VkCommandBuffer                             commandBuffer,
    VkBool32                                    stippledLineEnable);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `stippledLineEnable` specifies the `stippledLineEnable` state.

## [](#_description)Description

This command sets the `stippledLineEnable` state for subsequent drawing commands when drawing using [shader objects](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-objects), or when the graphics pipeline is created with `VK_DYNAMIC_STATE_LINE_STIPPLE_ENABLE_EXT` set in [VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`. Otherwise, this state is specified by the [VkPipelineRasterizationLineStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRasterizationLineStateCreateInfo.html)::`stippledLineEnable` value used to create the currently active pipeline.

Valid Usage

- [](#VUID-vkCmdSetLineStippleEnableEXT-None-09423)VUID-vkCmdSetLineStippleEnableEXT-None-09423  
  At least one of the following **must** be true:
  
  - The [`extendedDynamicState3LineStippleEnable`](#features-extendedDynamicState3LineStippleEnable) feature is enabled
  - The [`shaderObject`](#features-shaderObject) feature is enabled

Valid Usage (Implicit)

- [](#VUID-vkCmdSetLineStippleEnableEXT-commandBuffer-parameter)VUID-vkCmdSetLineStippleEnableEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdSetLineStippleEnableEXT-commandBuffer-recording)VUID-vkCmdSetLineStippleEnableEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdSetLineStippleEnableEXT-commandBuffer-cmdpool)VUID-vkCmdSetLineStippleEnableEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support VK\_QUEUE\_GRAPHICS\_BIT operations
- [](#VUID-vkCmdSetLineStippleEnableEXT-videocoding)VUID-vkCmdSetLineStippleEnableEXT-videocoding  
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

vkCmdSetLineStippleEnableEXT is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_EXT\_extended\_dynamic\_state3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_extended_dynamic_state3.html), [VK\_EXT\_line\_rasterization](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_line_rasterization.html), [VK\_EXT\_shader\_object](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_shader_object.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdSetLineStippleEnableEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0