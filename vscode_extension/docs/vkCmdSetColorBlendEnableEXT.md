# vkCmdSetColorBlendEnableEXT(3) Manual Page

## Name

vkCmdSetColorBlendEnableEXT - Specify the pname:blendEnable for each attachment dynamically for a command buffer



## [](#_c_specification)C Specification

To [dynamically set](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-dynamic-state) `blendEnable`, call:

```c++
// Provided by VK_EXT_extended_dynamic_state3, VK_EXT_shader_object
void vkCmdSetColorBlendEnableEXT(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    firstAttachment,
    uint32_t                                    attachmentCount,
    const VkBool32*                             pColorBlendEnables);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `firstAttachment` the first color attachment the color blending enable applies.
- `attachmentCount` the number of color blending enables in the `pColorBlendEnables` array.
- `pColorBlendEnables` an array of booleans to indicate whether color blending is enabled for the corresponding attachment.

## [](#_description)Description

This command sets the color blending enable of the specified color attachments for subsequent drawing commands when drawing using [shader objects](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-objects), or when the graphics pipeline is created with `VK_DYNAMIC_STATE_COLOR_BLEND_ENABLE_EXT` set in [VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`. Otherwise, this state is specified by the [VkPipelineColorBlendAttachmentState](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineColorBlendAttachmentState.html)::`blendEnable` values used to create the currently active pipeline.

Valid Usage

- [](#VUID-vkCmdSetColorBlendEnableEXT-None-09423)VUID-vkCmdSetColorBlendEnableEXT-None-09423  
  At least one of the following **must** be true:
  
  - The [`extendedDynamicState3ColorBlendEnable`](#features-extendedDynamicState3ColorBlendEnable) feature is enabled
  - The [`shaderObject`](#features-shaderObject) feature is enabled

Valid Usage (Implicit)

- [](#VUID-vkCmdSetColorBlendEnableEXT-commandBuffer-parameter)VUID-vkCmdSetColorBlendEnableEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdSetColorBlendEnableEXT-pColorBlendEnables-parameter)VUID-vkCmdSetColorBlendEnableEXT-pColorBlendEnables-parameter  
  `pColorBlendEnables` **must** be a valid pointer to an array of `attachmentCount` [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html) values
- [](#VUID-vkCmdSetColorBlendEnableEXT-commandBuffer-recording)VUID-vkCmdSetColorBlendEnableEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdSetColorBlendEnableEXT-commandBuffer-cmdpool)VUID-vkCmdSetColorBlendEnableEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics operations
- [](#VUID-vkCmdSetColorBlendEnableEXT-videocoding)VUID-vkCmdSetColorBlendEnableEXT-videocoding  
  This command **must** only be called outside of a video coding scope
- [](#VUID-vkCmdSetColorBlendEnableEXT-attachmentCount-arraylength)VUID-vkCmdSetColorBlendEnableEXT-attachmentCount-arraylength  
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

vkCmdSetColorBlendEnableEXT is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_EXT\_extended\_dynamic\_state3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_extended_dynamic_state3.html), [VK\_EXT\_shader\_object](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_shader_object.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdSetColorBlendEnableEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0