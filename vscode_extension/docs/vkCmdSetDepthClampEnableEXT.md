# vkCmdSetDepthClampEnableEXT(3) Manual Page

## Name

vkCmdSetDepthClampEnableEXT - Specify dynamically whether depth clamping is enabled in the command buffer



## [](#_c_specification)C Specification

To [dynamically set](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-dynamic-state) enable or disable depth clamping, call:

```c++
// Provided by VK_EXT_extended_dynamic_state3, VK_EXT_shader_object
void vkCmdSetDepthClampEnableEXT(
    VkCommandBuffer                             commandBuffer,
    VkBool32                                    depthClampEnable);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `depthClampEnable` specifies whether depth clamping is enabled.

## [](#_description)Description

This command sets whether depth clamping is enabled or disabled for subsequent drawing commands when drawing using [shader objects](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-objects), or when the graphics pipeline is created with `VK_DYNAMIC_STATE_DEPTH_CLAMP_ENABLE_EXT` set in [VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`. Otherwise, this state is specified by the [VkPipelineRasterizationStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRasterizationStateCreateInfo.html)::`depthClampEnable` value used to create the currently active pipeline.

If the depth clamping state is changed dynamically, and the pipeline was not created with `VK_DYNAMIC_STATE_DEPTH_CLIP_ENABLE_EXT` enabled, then depth clipping is enabled when depth clamping is disabled and vice versa.

Valid Usage

- [](#VUID-vkCmdSetDepthClampEnableEXT-None-09423)VUID-vkCmdSetDepthClampEnableEXT-None-09423  
  At least one of the following **must** be true:
  
  - The [`extendedDynamicState3DepthClampEnable`](#features-extendedDynamicState3DepthClampEnable) feature is enabled
  - The [`shaderObject`](#features-shaderObject) feature is enabled
- [](#VUID-vkCmdSetDepthClampEnableEXT-depthClamp-07449)VUID-vkCmdSetDepthClampEnableEXT-depthClamp-07449  
  If the [`depthClamp`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-depthClamp) feature is not enabled, `depthClampEnable` **must** be `VK_FALSE`

Valid Usage (Implicit)

- [](#VUID-vkCmdSetDepthClampEnableEXT-commandBuffer-parameter)VUID-vkCmdSetDepthClampEnableEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdSetDepthClampEnableEXT-commandBuffer-recording)VUID-vkCmdSetDepthClampEnableEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdSetDepthClampEnableEXT-commandBuffer-cmdpool)VUID-vkCmdSetDepthClampEnableEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support VK\_QUEUE\_GRAPHICS\_BIT operations
- [](#VUID-vkCmdSetDepthClampEnableEXT-videocoding)VUID-vkCmdSetDepthClampEnableEXT-videocoding  
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

vkCmdSetDepthClampEnableEXT is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_EXT\_extended\_dynamic\_state3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_extended_dynamic_state3.html), [VK\_EXT\_shader\_object](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_shader_object.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdSetDepthClampEnableEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0