# vkCmdSetColorBlendEquationEXT(3) Manual Page

## Name

vkCmdSetColorBlendEquationEXT - Specify the blend factors and operations dynamically for a command buffer



## [](#_c_specification)C Specification

To [dynamically set](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-dynamic-state) color blend factors and operations, call:

```c++
// Provided by VK_EXT_extended_dynamic_state3, VK_EXT_shader_object
void vkCmdSetColorBlendEquationEXT(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    firstAttachment,
    uint32_t                                    attachmentCount,
    const VkColorBlendEquationEXT*              pColorBlendEquations);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `firstAttachment` the first color attachment the color blend factors and operations apply to.
- `attachmentCount` the number of [VkColorBlendEquationEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkColorBlendEquationEXT.html) elements in the `pColorBlendEquations` array.
- `pColorBlendEquations` an array of [VkColorBlendEquationEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkColorBlendEquationEXT.html) structs that specify the color blend factors and operations for the corresponding attachments.

## [](#_description)Description

This command sets the color blending factors and operations of the specified attachments for subsequent drawing commands when drawing using [shader objects](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-objects), or when the graphics pipeline is created with `VK_DYNAMIC_STATE_COLOR_BLEND_EQUATION_EXT` set in [VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`. Otherwise, this state is specified by the [VkPipelineColorBlendAttachmentState](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineColorBlendAttachmentState.html)::`srcColorBlendFactor`, [VkPipelineColorBlendAttachmentState](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineColorBlendAttachmentState.html)::`dstColorBlendFactor`, [VkPipelineColorBlendAttachmentState](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineColorBlendAttachmentState.html)::`colorBlendOp`, [VkPipelineColorBlendAttachmentState](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineColorBlendAttachmentState.html)::`srcAlphaBlendFactor`, [VkPipelineColorBlendAttachmentState](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineColorBlendAttachmentState.html)::`dstAlphaBlendFactor`, and [VkPipelineColorBlendAttachmentState](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineColorBlendAttachmentState.html)::`alphaBlendOp` values used to create the currently active pipeline.

Valid Usage

- [](#VUID-vkCmdSetColorBlendEquationEXT-None-09423)VUID-vkCmdSetColorBlendEquationEXT-None-09423  
  At least one of the following **must** be true:
  
  - The [`extendedDynamicState3ColorBlendEquation`](#features-extendedDynamicState3ColorBlendEquation) feature is enabled
  - The [`shaderObject`](#features-shaderObject) feature is enabled

Valid Usage (Implicit)

- [](#VUID-vkCmdSetColorBlendEquationEXT-commandBuffer-parameter)VUID-vkCmdSetColorBlendEquationEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdSetColorBlendEquationEXT-pColorBlendEquations-parameter)VUID-vkCmdSetColorBlendEquationEXT-pColorBlendEquations-parameter  
  `pColorBlendEquations` **must** be a valid pointer to an array of `attachmentCount` valid [VkColorBlendEquationEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkColorBlendEquationEXT.html) structures
- [](#VUID-vkCmdSetColorBlendEquationEXT-commandBuffer-recording)VUID-vkCmdSetColorBlendEquationEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdSetColorBlendEquationEXT-commandBuffer-cmdpool)VUID-vkCmdSetColorBlendEquationEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics operations
- [](#VUID-vkCmdSetColorBlendEquationEXT-videocoding)VUID-vkCmdSetColorBlendEquationEXT-videocoding  
  This command **must** only be called outside of a video coding scope
- [](#VUID-vkCmdSetColorBlendEquationEXT-attachmentCount-arraylength)VUID-vkCmdSetColorBlendEquationEXT-attachmentCount-arraylength  
  `attachmentCount` **must** be greater than `0`

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

vkCmdSetColorBlendEquationEXT is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_EXT\_extended\_dynamic\_state3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_extended_dynamic_state3.html), [VK\_EXT\_shader\_object](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_shader_object.html), [VkColorBlendEquationEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkColorBlendEquationEXT.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdSetColorBlendEquationEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0