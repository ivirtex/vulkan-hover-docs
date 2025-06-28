# vkCmdSetViewportWScalingEnableNV(3) Manual Page

## Name

vkCmdSetViewportWScalingEnableNV - Specify the viewport W scaling enable state dynamically for a command buffer



## [](#_c_specification)C Specification

To [dynamically set](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-dynamic-state) the `viewportWScalingEnable` state, call:

```c++
// Provided by VK_EXT_extended_dynamic_state3 with VK_NV_clip_space_w_scaling, VK_EXT_shader_object with VK_NV_clip_space_w_scaling
void vkCmdSetViewportWScalingEnableNV(
    VkCommandBuffer                             commandBuffer,
    VkBool32                                    viewportWScalingEnable);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `viewportWScalingEnable` specifies the `viewportWScalingEnable` state.

## [](#_description)Description

This command sets the `viewportWScalingEnable` state for subsequent drawing commands when drawing using [shader objects](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-objects), or when the graphics pipeline is created with `VK_DYNAMIC_STATE_VIEWPORT_W_SCALING_ENABLE_NV` set in [VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`. Otherwise, this state is specified by the [VkPipelineViewportWScalingStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineViewportWScalingStateCreateInfoNV.html)::`viewportWScalingEnable` value used to create the currently active pipeline.

Valid Usage

- [](#VUID-vkCmdSetViewportWScalingEnableNV-None-09423)VUID-vkCmdSetViewportWScalingEnableNV-None-09423  
  At least one of the following **must** be true:
  
  - The [`extendedDynamicState3ViewportWScalingEnable`](#features-extendedDynamicState3ViewportWScalingEnable) feature is enabled
  - The [`shaderObject`](#features-shaderObject) feature is enabled

Valid Usage (Implicit)

- [](#VUID-vkCmdSetViewportWScalingEnableNV-commandBuffer-parameter)VUID-vkCmdSetViewportWScalingEnableNV-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdSetViewportWScalingEnableNV-commandBuffer-recording)VUID-vkCmdSetViewportWScalingEnableNV-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdSetViewportWScalingEnableNV-commandBuffer-cmdpool)VUID-vkCmdSetViewportWScalingEnableNV-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics operations
- [](#VUID-vkCmdSetViewportWScalingEnableNV-videocoding)VUID-vkCmdSetViewportWScalingEnableNV-videocoding  
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

## [](#_see_also)See Also

[VK\_EXT\_extended\_dynamic\_state3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_extended_dynamic_state3.html), [VK\_EXT\_shader\_object](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_shader_object.html), [VK\_NV\_clip\_space\_w\_scaling](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_clip_space_w_scaling.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdSetViewportWScalingEnableNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0