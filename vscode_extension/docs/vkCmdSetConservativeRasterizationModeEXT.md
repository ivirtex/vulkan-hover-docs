# vkCmdSetConservativeRasterizationModeEXT(3) Manual Page

## Name

vkCmdSetConservativeRasterizationModeEXT - Specify the conservative rasterization mode dynamically for a command buffer



## [](#_c_specification)C Specification

To [dynamically set](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-dynamic-state) the `conservativeRasterizationMode`, call:

```c++
// Provided by VK_EXT_conservative_rasterization with VK_EXT_extended_dynamic_state3, VK_EXT_conservative_rasterization with VK_EXT_shader_object
void vkCmdSetConservativeRasterizationModeEXT(
    VkCommandBuffer                             commandBuffer,
    VkConservativeRasterizationModeEXT          conservativeRasterizationMode);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `conservativeRasterizationMode` specifies the `conservativeRasterizationMode` state.

## [](#_description)Description

This command sets the `conservativeRasterizationMode` state for subsequent drawing commands when drawing using [shader objects](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-objects), or when the graphics pipeline is created with `VK_DYNAMIC_STATE_CONSERVATIVE_RASTERIZATION_MODE_EXT` set in [VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`. Otherwise, this state is specified by the [VkPipelineRasterizationConservativeStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRasterizationConservativeStateCreateInfoEXT.html)::`conservativeRasterizationMode` value used to create the currently active pipeline.

Valid Usage

- [](#VUID-vkCmdSetConservativeRasterizationModeEXT-None-09423)VUID-vkCmdSetConservativeRasterizationModeEXT-None-09423  
  At least one of the following **must** be true:
  
  - The [`extendedDynamicState3ConservativeRasterizationMode`](#features-extendedDynamicState3ConservativeRasterizationMode) feature is enabled
  - The [`shaderObject`](#features-shaderObject) feature is enabled

Valid Usage (Implicit)

- [](#VUID-vkCmdSetConservativeRasterizationModeEXT-commandBuffer-parameter)VUID-vkCmdSetConservativeRasterizationModeEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdSetConservativeRasterizationModeEXT-conservativeRasterizationMode-parameter)VUID-vkCmdSetConservativeRasterizationModeEXT-conservativeRasterizationMode-parameter  
  `conservativeRasterizationMode` **must** be a valid [VkConservativeRasterizationModeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkConservativeRasterizationModeEXT.html) value
- [](#VUID-vkCmdSetConservativeRasterizationModeEXT-commandBuffer-recording)VUID-vkCmdSetConservativeRasterizationModeEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdSetConservativeRasterizationModeEXT-commandBuffer-cmdpool)VUID-vkCmdSetConservativeRasterizationModeEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics operations
- [](#VUID-vkCmdSetConservativeRasterizationModeEXT-videocoding)VUID-vkCmdSetConservativeRasterizationModeEXT-videocoding  
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

vkCmdSetConservativeRasterizationModeEXT is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_EXT\_conservative\_rasterization](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_conservative_rasterization.html), [VK\_EXT\_extended\_dynamic\_state3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_extended_dynamic_state3.html), [VK\_EXT\_shader\_object](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_shader_object.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkConservativeRasterizationModeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkConservativeRasterizationModeEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdSetConservativeRasterizationModeEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0