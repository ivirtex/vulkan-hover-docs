# vkCmdSetRenderingInputAttachmentIndices(3) Manual Page

## Name

vkCmdSetRenderingInputAttachmentIndices - Set input attachment index mappings for a command buffer



## [](#_c_specification)C Specification

To set the input attachment index mappings during dynamic rendering, call:

```c++
// Provided by VK_VERSION_1_4
void vkCmdSetRenderingInputAttachmentIndices(
    VkCommandBuffer                             commandBuffer,
    const VkRenderingInputAttachmentIndexInfo*  pInputAttachmentIndexInfo);
```

or the equivalent command

```c++
// Provided by VK_KHR_dynamic_rendering_local_read
void vkCmdSetRenderingInputAttachmentIndicesKHR(
    VkCommandBuffer                             commandBuffer,
    const VkRenderingInputAttachmentIndexInfo*  pInputAttachmentIndexInfo);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `pInputAttachmentIndexInfo` is a [VkRenderingInputAttachmentIndexInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingInputAttachmentIndexInfo.html) structure indicating the new mappings.

## [](#_description)Description

This command sets the input attachment index mappings for subsequent drawing commands, and **must** match the mappings provided to the bound pipeline, if one is bound, which **can** be set by chaining [VkRenderingInputAttachmentIndexInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingInputAttachmentIndexInfo.html) to [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineCreateInfo.html).

Until this command is called, mappings in the command buffer state are treated as each color attachment specified in [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBeginRendering.html) mapping to subpass inputs with a `InputAttachmentIndex` equal to its index in [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingInfo.html)::`pColorAttachments`, and depth/stencil attachments mapping to input attachments without these decorations. This state is reset whenever [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBeginRendering.html) is called.

Valid Usage

- [](#VUID-vkCmdSetRenderingInputAttachmentIndices-dynamicRenderingLocalRead-09516)VUID-vkCmdSetRenderingInputAttachmentIndices-dynamicRenderingLocalRead-09516  
  [`dynamicRenderingLocalRead`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-dynamicRenderingLocalRead) **must** be enabled
- [](#VUID-vkCmdSetRenderingInputAttachmentIndices-pInputAttachmentIndexInfo-09517)VUID-vkCmdSetRenderingInputAttachmentIndices-pInputAttachmentIndexInfo-09517  
  `pInputAttachmentIndexInfo->colorAttachmentCount` **must** be equal to the value of [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingInfo.html)::`colorAttachmentCount` used to begin the current render pass instance
- [](#VUID-vkCmdSetRenderingInputAttachmentIndices-commandBuffer-09518)VUID-vkCmdSetRenderingInputAttachmentIndices-commandBuffer-09518  
  The current render pass instance **must** have been started or resumed by [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBeginRendering.html) in this `commandBuffer`

Valid Usage (Implicit)

- [](#VUID-vkCmdSetRenderingInputAttachmentIndices-commandBuffer-parameter)VUID-vkCmdSetRenderingInputAttachmentIndices-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdSetRenderingInputAttachmentIndices-pInputAttachmentIndexInfo-parameter)VUID-vkCmdSetRenderingInputAttachmentIndices-pInputAttachmentIndexInfo-parameter  
  `pInputAttachmentIndexInfo` **must** be a valid pointer to a valid [VkRenderingInputAttachmentIndexInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingInputAttachmentIndexInfo.html) structure
- [](#VUID-vkCmdSetRenderingInputAttachmentIndices-commandBuffer-recording)VUID-vkCmdSetRenderingInputAttachmentIndices-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdSetRenderingInputAttachmentIndices-commandBuffer-cmdpool)VUID-vkCmdSetRenderingInputAttachmentIndices-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics operations
- [](#VUID-vkCmdSetRenderingInputAttachmentIndices-renderpass)VUID-vkCmdSetRenderingInputAttachmentIndices-renderpass  
  This command **must** only be called inside of a render pass instance
- [](#VUID-vkCmdSetRenderingInputAttachmentIndices-videocoding)VUID-vkCmdSetRenderingInputAttachmentIndices-videocoding  
  This command **must** only be called outside of a video coding scope

Host Synchronization

- Host access to `commandBuffer` **must** be externally synchronized
- Host access to the `VkCommandPool` that `commandBuffer` was allocated from **must** be externally synchronized

Command Properties

     [Command Buffer Levels](#VkCommandBufferLevel) [Render Pass Scope](#vkCmdBeginRenderPass) [Video Coding Scope](#vkCmdBeginVideoCodingKHR) [Supported Queue Types](#VkQueueFlagBits) [Command Type](#fundamentals-queueoperation-command-types)

Primary  
Secondary

Inside

Outside

Graphics

State

## [](#_see_also)See Also

[VK\_KHR\_dynamic\_rendering\_local\_read](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_dynamic_rendering_local_read.html), [VK\_VERSION\_1\_4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_4.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkRenderingInputAttachmentIndexInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingInputAttachmentIndexInfo.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdSetRenderingInputAttachmentIndices)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0