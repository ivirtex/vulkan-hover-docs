# vkCmdSetRasterizerDiscardEnable(3) Manual Page

## Name

vkCmdSetRasterizerDiscardEnable - Control whether primitives are discarded before the rasterization stage dynamically for a command buffer



## [](#_c_specification)C Specification

To [dynamically enable](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-dynamic-state) whether primitives are discarded before the rasterization stage, call:

```c++
// Provided by VK_VERSION_1_3
void vkCmdSetRasterizerDiscardEnable(
    VkCommandBuffer                             commandBuffer,
    VkBool32                                    rasterizerDiscardEnable);
```

or the equivalent command

```c++
// Provided by VK_EXT_extended_dynamic_state2, VK_EXT_shader_object
void vkCmdSetRasterizerDiscardEnableEXT(
    VkCommandBuffer                             commandBuffer,
    VkBool32                                    rasterizerDiscardEnable);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `rasterizerDiscardEnable` controls whether primitives are discarded immediately before the rasterization stage.

## [](#_description)Description

This command sets the discard enable for subsequent drawing commands when drawing using [shader objects](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-objects), or when the graphics pipeline is created with `VK_DYNAMIC_STATE_RASTERIZER_DISCARD_ENABLE` set in [VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`. Otherwise, this state is specified by the [VkPipelineRasterizationStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRasterizationStateCreateInfo.html)::`rasterizerDiscardEnable` value used to create the currently active pipeline.

Valid Usage

- [](#VUID-vkCmdSetRasterizerDiscardEnable-None-08970)VUID-vkCmdSetRasterizerDiscardEnable-None-08970  
  At least one of the following **must** be true:
  
  - the [`extendedDynamicState2`](#features-extendedDynamicState2) feature is enabled
  - the [`shaderObject`](#features-shaderObject) feature is enabled
  - the value of [VkApplicationInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkApplicationInfo.html)::`apiVersion` used to create the [VkInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstance.html) parent of `commandBuffer` is greater than or equal to Version 1.3

Valid Usage (Implicit)

- [](#VUID-vkCmdSetRasterizerDiscardEnable-commandBuffer-parameter)VUID-vkCmdSetRasterizerDiscardEnable-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdSetRasterizerDiscardEnable-commandBuffer-recording)VUID-vkCmdSetRasterizerDiscardEnable-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdSetRasterizerDiscardEnable-commandBuffer-cmdpool)VUID-vkCmdSetRasterizerDiscardEnable-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics operations
- [](#VUID-vkCmdSetRasterizerDiscardEnable-videocoding)VUID-vkCmdSetRasterizerDiscardEnable-videocoding  
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

[VK\_EXT\_extended\_dynamic\_state2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_extended_dynamic_state2.html), [VK\_EXT\_shader\_object](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_shader_object.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdSetRasterizerDiscardEnable)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0