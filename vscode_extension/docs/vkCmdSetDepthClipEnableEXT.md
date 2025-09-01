# vkCmdSetDepthClipEnableEXT(3) Manual Page

## Name

vkCmdSetDepthClipEnableEXT - Specify dynamically whether depth clipping is enabled in the command buffer



## [](#_c_specification)C Specification

To [dynamically set](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-dynamic-state) enable or disable depth clipping, call:

```c++
// Provided by VK_EXT_depth_clip_enable with VK_EXT_extended_dynamic_state3, VK_EXT_depth_clip_enable with VK_EXT_shader_object
void vkCmdSetDepthClipEnableEXT(
    VkCommandBuffer                             commandBuffer,
    VkBool32                                    depthClipEnable);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `depthClipEnable` specifies whether depth clipping is enabled.

## [](#_description)Description

This command sets whether depth clipping is enabled or disabled for subsequent drawing commands when drawing using [shader objects](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-objects), or when the graphics pipeline is created with `VK_DYNAMIC_STATE_DEPTH_CLIP_ENABLE_EXT` set in [VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`. Otherwise, this state is specified by the [VkPipelineRasterizationDepthClipStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRasterizationDepthClipStateCreateInfoEXT.html)::`depthClipEnable` value used to create the currently active pipeline, or by the inverse of [VkPipelineRasterizationStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRasterizationStateCreateInfo.html)::`depthClampEnable` if `VkPipelineRasterizationDepthClipStateCreateInfoEXT` is not specified.

Valid Usage

- [](#VUID-vkCmdSetDepthClipEnableEXT-None-09423)VUID-vkCmdSetDepthClipEnableEXT-None-09423  
  At least one of the following **must** be true:
  
  - The [`extendedDynamicState3DepthClipEnable`](#features-extendedDynamicState3DepthClipEnable) feature is enabled
  - The [`shaderObject`](#features-shaderObject) feature is enabled
- [](#VUID-vkCmdSetDepthClipEnableEXT-depthClipEnable-07451)VUID-vkCmdSetDepthClipEnableEXT-depthClipEnable-07451  
  The [`depthClipEnable`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-depthClipEnable) feature **must** be enabled

Valid Usage (Implicit)

- [](#VUID-vkCmdSetDepthClipEnableEXT-commandBuffer-parameter)VUID-vkCmdSetDepthClipEnableEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdSetDepthClipEnableEXT-commandBuffer-recording)VUID-vkCmdSetDepthClipEnableEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdSetDepthClipEnableEXT-commandBuffer-cmdpool)VUID-vkCmdSetDepthClipEnableEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics operations
- [](#VUID-vkCmdSetDepthClipEnableEXT-videocoding)VUID-vkCmdSetDepthClipEnableEXT-videocoding  
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

vkCmdSetDepthClipEnableEXT is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_EXT\_depth\_clip\_enable](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_depth_clip_enable.html), [VK\_EXT\_extended\_dynamic\_state3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_extended_dynamic_state3.html), [VK\_EXT\_shader\_object](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_shader_object.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdSetDepthClipEnableEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0