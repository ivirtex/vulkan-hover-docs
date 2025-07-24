# vkCmdSetColorWriteMaskEXT(3) Manual Page

## Name

vkCmdSetColorWriteMaskEXT - Specify the color write masks for each attachment dynamically for a command buffer



## [](#_c_specification)C Specification

To [dynamically set](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-dynamic-state) the color write masks, call:

```c++
// Provided by VK_EXT_extended_dynamic_state3, VK_EXT_shader_object
void vkCmdSetColorWriteMaskEXT(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    firstAttachment,
    uint32_t                                    attachmentCount,
    const VkColorComponentFlags*                pColorWriteMasks);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `firstAttachment` the first color attachment the color write masks apply to.
- `attachmentCount` the number of [VkColorComponentFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkColorComponentFlags.html) values in the `pColorWriteMasks` array.
- `pColorWriteMasks` an array of [VkColorComponentFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkColorComponentFlags.html) values that specify the color write masks of the corresponding attachments.

## [](#_description)Description

This command sets the color write masks of the specified attachments for subsequent drawing commands when drawing using [shader objects](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-objects), or when the graphics pipeline is created with `VK_DYNAMIC_STATE_COLOR_WRITE_MASK_EXT` set in [VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`. Otherwise, this state is specified by the [VkPipelineColorBlendAttachmentState](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineColorBlendAttachmentState.html)::`colorWriteMask` values used to create the currently active pipeline.

Note

Formats with bits that are shared between components specified by [VkColorComponentFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkColorComponentFlagBits.html), such as `VK_FORMAT_E5B9G9R9_UFLOAT_PACK32`, cannot have their channels individually masked by this functionality; either all components that share bits have to be enabled, or none of them.

Valid Usage

- [](#VUID-vkCmdSetColorWriteMaskEXT-None-09423)VUID-vkCmdSetColorWriteMaskEXT-None-09423  
  At least one of the following **must** be true:
  
  - The [`extendedDynamicState3ColorWriteMask`](#features-extendedDynamicState3ColorWriteMask) feature is enabled
  - The [`shaderObject`](#features-shaderObject) feature is enabled

Valid Usage (Implicit)

- [](#VUID-vkCmdSetColorWriteMaskEXT-commandBuffer-parameter)VUID-vkCmdSetColorWriteMaskEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdSetColorWriteMaskEXT-pColorWriteMasks-parameter)VUID-vkCmdSetColorWriteMaskEXT-pColorWriteMasks-parameter  
  `pColorWriteMasks` **must** be a valid pointer to an array of `attachmentCount` valid combinations of [VkColorComponentFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkColorComponentFlagBits.html) values
- [](#VUID-vkCmdSetColorWriteMaskEXT-commandBuffer-recording)VUID-vkCmdSetColorWriteMaskEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdSetColorWriteMaskEXT-commandBuffer-cmdpool)VUID-vkCmdSetColorWriteMaskEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics operations
- [](#VUID-vkCmdSetColorWriteMaskEXT-videocoding)VUID-vkCmdSetColorWriteMaskEXT-videocoding  
  This command **must** only be called outside of a video coding scope
- [](#VUID-vkCmdSetColorWriteMaskEXT-attachmentCount-arraylength)VUID-vkCmdSetColorWriteMaskEXT-attachmentCount-arraylength  
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

vkCmdSetColorWriteMaskEXT is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_EXT\_extended\_dynamic\_state3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_extended_dynamic_state3.html), [VK\_EXT\_shader\_object](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_shader_object.html), [VkColorComponentFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkColorComponentFlags.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdSetColorWriteMaskEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0