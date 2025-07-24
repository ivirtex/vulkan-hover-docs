# vkCmdSetCoverageReductionModeNV(3) Manual Page

## Name

vkCmdSetCoverageReductionModeNV - Specify the coverage reduction mode dynamically for a command buffer



## [](#_c_specification)C Specification

To [dynamically set](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-dynamic-state) the `coverageReductionMode` state, call:

```c++
// Provided by VK_EXT_extended_dynamic_state3 with VK_NV_coverage_reduction_mode, VK_EXT_shader_object with VK_NV_coverage_reduction_mode
void vkCmdSetCoverageReductionModeNV(
    VkCommandBuffer                             commandBuffer,
    VkCoverageReductionModeNV                   coverageReductionMode);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `coverageReductionMode` specifies the `coverageReductionMode` state.

## [](#_description)Description

This command sets the `coverageReductionMode` state for subsequent drawing commands when drawing using [shader objects](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-objects), or when the graphics pipeline is created with `VK_DYNAMIC_STATE_COVERAGE_REDUCTION_MODE_NV` set in [VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`. Otherwise, this state is specified by the [VkPipelineCoverageReductionStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCoverageReductionStateCreateInfoNV.html)::`coverageReductionMode` value used to create the currently active pipeline.

Valid Usage

- [](#VUID-vkCmdSetCoverageReductionModeNV-None-09423)VUID-vkCmdSetCoverageReductionModeNV-None-09423  
  At least one of the following **must** be true:
  
  - The [`extendedDynamicState3CoverageReductionMode`](#features-extendedDynamicState3CoverageReductionMode) feature is enabled
  - The [`shaderObject`](#features-shaderObject) feature is enabled

Valid Usage (Implicit)

- [](#VUID-vkCmdSetCoverageReductionModeNV-commandBuffer-parameter)VUID-vkCmdSetCoverageReductionModeNV-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdSetCoverageReductionModeNV-coverageReductionMode-parameter)VUID-vkCmdSetCoverageReductionModeNV-coverageReductionMode-parameter  
  `coverageReductionMode` **must** be a valid [VkCoverageReductionModeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCoverageReductionModeNV.html) value
- [](#VUID-vkCmdSetCoverageReductionModeNV-commandBuffer-recording)VUID-vkCmdSetCoverageReductionModeNV-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdSetCoverageReductionModeNV-commandBuffer-cmdpool)VUID-vkCmdSetCoverageReductionModeNV-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics operations
- [](#VUID-vkCmdSetCoverageReductionModeNV-videocoding)VUID-vkCmdSetCoverageReductionModeNV-videocoding  
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

vkCmdSetCoverageReductionModeNV is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_EXT\_extended\_dynamic\_state3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_extended_dynamic_state3.html), [VK\_EXT\_shader\_object](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_shader_object.html), [VK\_NV\_coverage\_reduction\_mode](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_coverage_reduction_mode.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkCoverageReductionModeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCoverageReductionModeNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdSetCoverageReductionModeNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0