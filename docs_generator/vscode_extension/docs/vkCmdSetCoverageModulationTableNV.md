# vkCmdSetCoverageModulationTableNV(3) Manual Page

## Name

vkCmdSetCoverageModulationTableNV - Specify the coverage modulation table dynamically for a command buffer



## [](#_c_specification)C Specification

To [dynamically set](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-dynamic-state) the `pCoverageModulationTable` state, call:

```c++
// Provided by VK_EXT_extended_dynamic_state3 with VK_NV_framebuffer_mixed_samples, VK_EXT_shader_object with VK_NV_framebuffer_mixed_samples
void vkCmdSetCoverageModulationTableNV(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    coverageModulationTableCount,
    const float*                                pCoverageModulationTable);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `coverageModulationTableCount` specifies the number of elements in `pCoverageModulationTable`.
- `pCoverageModulationTable` specifies the table of modulation factors containing a value for each number of covered samples.

## [](#_description)Description

This command sets the table of modulation factors for subsequent drawing commands when drawing using [shader objects](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-objects), or when the graphics pipeline is created with `VK_DYNAMIC_STATE_COVERAGE_MODULATION_TABLE_NV` set in [VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`. Otherwise, this state is specified by the [VkPipelineCoverageModulationStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCoverageModulationStateCreateInfoNV.html)::`coverageModulationTableCount`, and [VkPipelineCoverageModulationStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCoverageModulationStateCreateInfoNV.html)::`pCoverageModulationTable` values used to create the currently active pipeline.

Valid Usage

- [](#VUID-vkCmdSetCoverageModulationTableNV-None-09423)VUID-vkCmdSetCoverageModulationTableNV-None-09423  
  At least one of the following **must** be true:
  
  - The [`extendedDynamicState3CoverageModulationTable`](#features-extendedDynamicState3CoverageModulationTable) feature is enabled
  - The [`shaderObject`](#features-shaderObject) feature is enabled

Valid Usage (Implicit)

- [](#VUID-vkCmdSetCoverageModulationTableNV-commandBuffer-parameter)VUID-vkCmdSetCoverageModulationTableNV-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdSetCoverageModulationTableNV-pCoverageModulationTable-parameter)VUID-vkCmdSetCoverageModulationTableNV-pCoverageModulationTable-parameter  
  `pCoverageModulationTable` **must** be a valid pointer to an array of `coverageModulationTableCount` `float` values
- [](#VUID-vkCmdSetCoverageModulationTableNV-commandBuffer-recording)VUID-vkCmdSetCoverageModulationTableNV-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdSetCoverageModulationTableNV-commandBuffer-cmdpool)VUID-vkCmdSetCoverageModulationTableNV-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics operations
- [](#VUID-vkCmdSetCoverageModulationTableNV-videocoding)VUID-vkCmdSetCoverageModulationTableNV-videocoding  
  This command **must** only be called outside of a video coding scope
- [](#VUID-vkCmdSetCoverageModulationTableNV-coverageModulationTableCount-arraylength)VUID-vkCmdSetCoverageModulationTableNV-coverageModulationTableCount-arraylength  
  `coverageModulationTableCount` **must** be greater than `0`

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

[VK\_EXT\_extended\_dynamic\_state3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_extended_dynamic_state3.html), [VK\_EXT\_shader\_object](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_shader_object.html), [VK\_NV\_framebuffer\_mixed\_samples](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_framebuffer_mixed_samples.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdSetCoverageModulationTableNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0