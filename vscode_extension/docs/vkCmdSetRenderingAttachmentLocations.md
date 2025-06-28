# vkCmdSetRenderingAttachmentLocations(3) Manual Page

## Name

vkCmdSetRenderingAttachmentLocations - Set color attachment location mappings for a command buffer



## [](#_c_specification)C Specification

To set the fragment output location mappings during rendering, call:

```c++
// Provided by VK_VERSION_1_4
void vkCmdSetRenderingAttachmentLocations(
    VkCommandBuffer                             commandBuffer,
    const VkRenderingAttachmentLocationInfo*    pLocationInfo);
```

or the equivalent command

```c++
// Provided by VK_KHR_dynamic_rendering_local_read
void vkCmdSetRenderingAttachmentLocationsKHR(
    VkCommandBuffer                             commandBuffer,
    const VkRenderingAttachmentLocationInfo*    pLocationInfo);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `pLocationInfo` is a [VkRenderingAttachmentLocationInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingAttachmentLocationInfo.html) structure indicating the new mappings.

## [](#_description)Description

This command sets the attachment location mappings for subsequent drawing commands, and **must** match the mappings provided to the bound pipeline, if one is bound, which **can** be set by chaining [VkRenderingAttachmentLocationInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingAttachmentLocationInfo.html) to [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineCreateInfo.html).

Until this command is called, mappings in the command buffer state are treated as each color attachment specified in [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBeginRendering.html) having a location equal to its index in [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingInfo.html)::`pColorAttachments`. This state is reset whenever [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBeginRendering.html) is called.

Valid Usage

- [](#VUID-vkCmdSetRenderingAttachmentLocations-dynamicRenderingLocalRead-09509)VUID-vkCmdSetRenderingAttachmentLocations-dynamicRenderingLocalRead-09509  
  [`dynamicRenderingLocalRead`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-dynamicRenderingLocalRead) **must** be enabled
- [](#VUID-vkCmdSetRenderingAttachmentLocations-pLocationInfo-09510)VUID-vkCmdSetRenderingAttachmentLocations-pLocationInfo-09510  
  `pLocationInfo->colorAttachmentCount` **must** be equal to the value of [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingInfo.html)::`colorAttachmentCount` used to begin the current render pass instance
- [](#VUID-vkCmdSetRenderingAttachmentLocations-commandBuffer-09511)VUID-vkCmdSetRenderingAttachmentLocations-commandBuffer-09511  
  The current render pass instance **must** have been started or resumed by [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBeginRendering.html) in this `commandBuffer`

Valid Usage (Implicit)

- [](#VUID-vkCmdSetRenderingAttachmentLocations-commandBuffer-parameter)VUID-vkCmdSetRenderingAttachmentLocations-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdSetRenderingAttachmentLocations-pLocationInfo-parameter)VUID-vkCmdSetRenderingAttachmentLocations-pLocationInfo-parameter  
  `pLocationInfo` **must** be a valid pointer to a valid [VkRenderingAttachmentLocationInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingAttachmentLocationInfo.html) structure
- [](#VUID-vkCmdSetRenderingAttachmentLocations-commandBuffer-recording)VUID-vkCmdSetRenderingAttachmentLocations-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdSetRenderingAttachmentLocations-commandBuffer-cmdpool)VUID-vkCmdSetRenderingAttachmentLocations-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics operations
- [](#VUID-vkCmdSetRenderingAttachmentLocations-renderpass)VUID-vkCmdSetRenderingAttachmentLocations-renderpass  
  This command **must** only be called inside of a render pass instance
- [](#VUID-vkCmdSetRenderingAttachmentLocations-videocoding)VUID-vkCmdSetRenderingAttachmentLocations-videocoding  
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

[VK\_KHR\_dynamic\_rendering\_local\_read](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_dynamic_rendering_local_read.html), [VK\_VERSION\_1\_4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_4.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkRenderingAttachmentLocationInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingAttachmentLocationInfo.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdSetRenderingAttachmentLocations)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0