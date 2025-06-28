# vkCmdSetCoverageModulationModeNV(3) Manual Page

## Name

vkCmdSetCoverageModulationModeNV - Specify the coverage modulation mode dynamically for a command buffer



## [](#_c_specification)C Specification

To [dynamically set](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-dynamic-state) the `coverageModulationMode` state, call:

```c++
// Provided by VK_EXT_extended_dynamic_state3 with VK_NV_framebuffer_mixed_samples, VK_EXT_shader_object with VK_NV_framebuffer_mixed_samples
void vkCmdSetCoverageModulationModeNV(
    VkCommandBuffer                             commandBuffer,
    VkCoverageModulationModeNV                  coverageModulationMode);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `coverageModulationMode` specifies the `coverageModulationMode` state.

## [](#_description)Description

This command sets the `coverageModulationMode` state for subsequent drawing commands when drawing using [shader objects](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-objects), or the graphics pipeline is created with `VK_DYNAMIC_STATE_COVERAGE_MODULATION_MODE_NV` set in [VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`. Otherwise, this state is specified by the [VkPipelineCoverageModulationStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCoverageModulationStateCreateInfoNV.html)::`coverageModulationMode` value used to create the currently active pipeline.

Valid Usage

- [](#VUID-vkCmdSetCoverageModulationModeNV-None-09423)VUID-vkCmdSetCoverageModulationModeNV-None-09423  
  At least one of the following **must** be true:
  
  - The [`extendedDynamicState3CoverageModulationMode`](#features-extendedDynamicState3CoverageModulationMode) feature is enabled
  - The [`shaderObject`](#features-shaderObject) feature is enabled

Valid Usage (Implicit)

- [](#VUID-vkCmdSetCoverageModulationModeNV-commandBuffer-parameter)VUID-vkCmdSetCoverageModulationModeNV-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdSetCoverageModulationModeNV-coverageModulationMode-parameter)VUID-vkCmdSetCoverageModulationModeNV-coverageModulationMode-parameter  
  `coverageModulationMode` **must** be a valid [VkCoverageModulationModeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCoverageModulationModeNV.html) value
- [](#VUID-vkCmdSetCoverageModulationModeNV-commandBuffer-recording)VUID-vkCmdSetCoverageModulationModeNV-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdSetCoverageModulationModeNV-commandBuffer-cmdpool)VUID-vkCmdSetCoverageModulationModeNV-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics operations
- [](#VUID-vkCmdSetCoverageModulationModeNV-videocoding)VUID-vkCmdSetCoverageModulationModeNV-videocoding  
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

[VK\_EXT\_extended\_dynamic\_state3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_extended_dynamic_state3.html), [VK\_EXT\_shader\_object](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_shader_object.html), [VK\_NV\_framebuffer\_mixed\_samples](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_framebuffer_mixed_samples.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkCoverageModulationModeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCoverageModulationModeNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdSetCoverageModulationModeNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0