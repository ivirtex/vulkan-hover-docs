# vkCmdSetColorBlendAdvancedEXT(3) Manual Page

## Name

vkCmdSetColorBlendAdvancedEXT - Specify the advanced color blend state dynamically for a command buffer



## [](#_c_specification)C Specification

To [dynamically set](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-dynamic-state) the advanced blend state, call:

```c++
// Provided by VK_EXT_blend_operation_advanced with VK_EXT_extended_dynamic_state3, VK_EXT_blend_operation_advanced with VK_EXT_shader_object
void vkCmdSetColorBlendAdvancedEXT(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    firstAttachment,
    uint32_t                                    attachmentCount,
    const VkColorBlendAdvancedEXT*              pColorBlendAdvanced);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `firstAttachment` the first color attachment the advanced blend parameters apply to.
- `attachmentCount` the number of [VkColorBlendAdvancedEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkColorBlendAdvancedEXT.html) elements in the `pColorBlendAdvanced` array.
- `pColorBlendAdvanced` an array of [VkColorBlendAdvancedEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkColorBlendAdvancedEXT.html) structs that specify the advanced color blend parameters for the corresponding attachments.

## [](#_description)Description

This command sets the advanced blend operation parameters of the specified attachments for subsequent drawing commands when drawing using [shader objects](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-objects), or when the graphics pipeline is created with `VK_DYNAMIC_STATE_COLOR_BLEND_ADVANCED_EXT` set in [VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`. Otherwise, this state is specified by the [VkPipelineColorBlendAdvancedStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineColorBlendAdvancedStateCreateInfoEXT.html)::`srcPremultiplied`, [VkPipelineColorBlendAdvancedStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineColorBlendAdvancedStateCreateInfoEXT.html)::`dstPremultiplied`, and [VkPipelineColorBlendAdvancedStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineColorBlendAdvancedStateCreateInfoEXT.html)::`blendOverlap` values used to create the currently active pipeline.

Valid Usage

- [](#VUID-vkCmdSetColorBlendAdvancedEXT-None-09423)VUID-vkCmdSetColorBlendAdvancedEXT-None-09423  
  At least one of the following **must** be true:
  
  - The [`extendedDynamicState3ColorBlendAdvanced`](#features-extendedDynamicState3ColorBlendAdvanced) feature is enabled
  - The [`shaderObject`](#features-shaderObject) feature is enabled

Valid Usage (Implicit)

- [](#VUID-vkCmdSetColorBlendAdvancedEXT-commandBuffer-parameter)VUID-vkCmdSetColorBlendAdvancedEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdSetColorBlendAdvancedEXT-pColorBlendAdvanced-parameter)VUID-vkCmdSetColorBlendAdvancedEXT-pColorBlendAdvanced-parameter  
  `pColorBlendAdvanced` **must** be a valid pointer to an array of `attachmentCount` valid [VkColorBlendAdvancedEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkColorBlendAdvancedEXT.html) structures
- [](#VUID-vkCmdSetColorBlendAdvancedEXT-commandBuffer-recording)VUID-vkCmdSetColorBlendAdvancedEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdSetColorBlendAdvancedEXT-commandBuffer-cmdpool)VUID-vkCmdSetColorBlendAdvancedEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics operations
- [](#VUID-vkCmdSetColorBlendAdvancedEXT-videocoding)VUID-vkCmdSetColorBlendAdvancedEXT-videocoding  
  This command **must** only be called outside of a video coding scope
- [](#VUID-vkCmdSetColorBlendAdvancedEXT-attachmentCount-arraylength)VUID-vkCmdSetColorBlendAdvancedEXT-attachmentCount-arraylength  
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

vkCmdSetColorBlendAdvancedEXT is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_EXT\_blend\_operation\_advanced](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_blend_operation_advanced.html), [VK\_EXT\_extended\_dynamic\_state3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_extended_dynamic_state3.html), [VK\_EXT\_shader\_object](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_shader_object.html), [VkColorBlendAdvancedEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkColorBlendAdvancedEXT.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdSetColorBlendAdvancedEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0