# vkCmdSetBlendConstants(3) Manual Page

## Name

vkCmdSetBlendConstants - Set the values of blend constants



## [](#_c_specification)C Specification

To [dynamically set and change](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-dynamic-state) the blend constants, call:

```c++
// Provided by VK_VERSION_1_0
void vkCmdSetBlendConstants(
    VkCommandBuffer                             commandBuffer,
    const float                                 blendConstants[4]);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `blendConstants` is a pointer to an array of four values specifying the Rc, Gc, Bc, and Ac components of the blend constant color used in blending, depending on the [blend factor](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#framebuffer-blendfactors).

## [](#_description)Description

This command sets blend constants for subsequent drawing commands when when drawing using [shader objects](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-objects), or the graphics pipeline is created with `VK_DYNAMIC_STATE_BLEND_CONSTANTS` set in [VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`. Otherwise, this state is specified by the [VkPipelineColorBlendStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineColorBlendStateCreateInfo.html)::`blendConstants` values used to create the currently active pipeline.

Valid Usage (Implicit)

- [](#VUID-vkCmdSetBlendConstants-commandBuffer-parameter)VUID-vkCmdSetBlendConstants-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdSetBlendConstants-commandBuffer-recording)VUID-vkCmdSetBlendConstants-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdSetBlendConstants-commandBuffer-cmdpool)VUID-vkCmdSetBlendConstants-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics operations
- [](#VUID-vkCmdSetBlendConstants-videocoding)VUID-vkCmdSetBlendConstants-videocoding  
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

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdSetBlendConstants)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0