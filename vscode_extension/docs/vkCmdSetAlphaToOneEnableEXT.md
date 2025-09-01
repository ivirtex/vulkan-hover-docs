# vkCmdSetAlphaToOneEnableEXT(3) Manual Page

## Name

vkCmdSetAlphaToOneEnableEXT - Specify the alpha to one enable state dynamically for a command buffer



## [](#_c_specification)C Specification

To [dynamically set](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-dynamic-state) the `alphaToOneEnable` state, call:

```c++
// Provided by VK_EXT_extended_dynamic_state3, VK_EXT_shader_object
void vkCmdSetAlphaToOneEnableEXT(
    VkCommandBuffer                             commandBuffer,
    VkBool32                                    alphaToOneEnable);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `alphaToOneEnable` specifies the `alphaToOneEnable` state.

## [](#_description)Description

This command sets the `alphaToOneEnable` state for subsequent drawing commands when drawing using [shader objects](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-objects), or when the graphics pipeline is created with `VK_DYNAMIC_STATE_ALPHA_TO_ONE_ENABLE_EXT` set in [VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`. Otherwise, this state is specified by the [VkPipelineMultisampleStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineMultisampleStateCreateInfo.html)::`alphaToOneEnable` value used to create the currently active pipeline.

Valid Usage

- [](#VUID-vkCmdSetAlphaToOneEnableEXT-None-09423)VUID-vkCmdSetAlphaToOneEnableEXT-None-09423  
  At least one of the following **must** be true:
  
  - The [`extendedDynamicState3AlphaToOneEnable`](#features-extendedDynamicState3AlphaToOneEnable) feature is enabled
  - The [`shaderObject`](#features-shaderObject) feature is enabled
- [](#VUID-vkCmdSetAlphaToOneEnableEXT-alphaToOne-07607)VUID-vkCmdSetAlphaToOneEnableEXT-alphaToOne-07607  
  If the [`alphaToOne`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-alphaToOne) feature is not enabled, `alphaToOneEnable` **must** be `VK_FALSE`

Valid Usage (Implicit)

- [](#VUID-vkCmdSetAlphaToOneEnableEXT-commandBuffer-parameter)VUID-vkCmdSetAlphaToOneEnableEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdSetAlphaToOneEnableEXT-commandBuffer-recording)VUID-vkCmdSetAlphaToOneEnableEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdSetAlphaToOneEnableEXT-commandBuffer-cmdpool)VUID-vkCmdSetAlphaToOneEnableEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics operations
- [](#VUID-vkCmdSetAlphaToOneEnableEXT-videocoding)VUID-vkCmdSetAlphaToOneEnableEXT-videocoding  
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

vkCmdSetAlphaToOneEnableEXT is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_EXT\_extended\_dynamic\_state3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_extended_dynamic_state3.html), [VK\_EXT\_shader\_object](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_shader_object.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdSetAlphaToOneEnableEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0