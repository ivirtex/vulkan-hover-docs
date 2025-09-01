# vkCmdSetCoverageToColorLocationNV(3) Manual Page

## Name

vkCmdSetCoverageToColorLocationNV - Specify the coverage to color location dynamically for a command buffer



## [](#_c_specification)C Specification

To [dynamically set](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-dynamic-state) the `coverageToColorLocation` state, call:

```c++
// Provided by VK_EXT_extended_dynamic_state3 with VK_NV_fragment_coverage_to_color, VK_EXT_shader_object with VK_NV_fragment_coverage_to_color
void vkCmdSetCoverageToColorLocationNV(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    coverageToColorLocation);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `coverageToColorLocation` specifies the `coverageToColorLocation` state.

## [](#_description)Description

This command sets the `coverageToColorLocation` state for subsequent drawing commands when drawing using [shader objects](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-objects), or when the graphics pipeline is created with `VK_DYNAMIC_STATE_COVERAGE_TO_COLOR_LOCATION_NV` set in [VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`. Otherwise, this state is specified by the [VkPipelineCoverageToColorStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCoverageToColorStateCreateInfoNV.html)::`coverageToColorLocation` value used to create the currently active pipeline.

Valid Usage

- [](#VUID-vkCmdSetCoverageToColorLocationNV-None-09423)VUID-vkCmdSetCoverageToColorLocationNV-None-09423  
  At least one of the following **must** be true:
  
  - The [`extendedDynamicState3CoverageToColorLocation`](#features-extendedDynamicState3CoverageToColorLocation) feature is enabled
  - The [`shaderObject`](#features-shaderObject) feature is enabled

Valid Usage (Implicit)

- [](#VUID-vkCmdSetCoverageToColorLocationNV-commandBuffer-parameter)VUID-vkCmdSetCoverageToColorLocationNV-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdSetCoverageToColorLocationNV-commandBuffer-recording)VUID-vkCmdSetCoverageToColorLocationNV-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdSetCoverageToColorLocationNV-commandBuffer-cmdpool)VUID-vkCmdSetCoverageToColorLocationNV-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics operations
- [](#VUID-vkCmdSetCoverageToColorLocationNV-videocoding)VUID-vkCmdSetCoverageToColorLocationNV-videocoding  
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

vkCmdSetCoverageToColorLocationNV is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_EXT\_extended\_dynamic\_state3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_extended_dynamic_state3.html), [VK\_EXT\_shader\_object](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_shader_object.html), [VK\_NV\_fragment\_coverage\_to\_color](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_fragment_coverage_to_color.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdSetCoverageToColorLocationNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0