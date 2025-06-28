# vkCmdEndRendering2EXT(3) Manual Page

## Name

vkCmdEndRendering2EXT - End a dynamic render pass instance



## [](#_c_specification)C Specification

Alternatively, to end a render pass instance, call:

```c++
// Provided by VK_EXT_fragment_density_map_offset
void vkCmdEndRendering2EXT(
    VkCommandBuffer                             commandBuffer,
    const VkRenderingEndInfoEXT*                pRenderingEndInfo);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer in which to record the command.
- `pRenderingEndInfo` is `NULL` or a pointer to a [VkRenderingEndInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingEndInfoEXT.html) structure containing information about how the render pass will be ended.

## [](#_description)Description

If the value of `pRenderingInfo->flags` used to begin this render pass instance included `VK_RENDERING_SUSPENDING_BIT`, then this render pass is suspended and will be resumed later in [submission order](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-submission-order).

Valid Usage

- [](#VUID-vkCmdEndRendering2EXT-None-10610)VUID-vkCmdEndRendering2EXT-None-10610  
  The current render pass instance **must** have been begun with [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBeginRendering.html)
- [](#VUID-vkCmdEndRendering2EXT-commandBuffer-10611)VUID-vkCmdEndRendering2EXT-commandBuffer-10611  
  The current render pass instance **must** have been begun in `commandBuffer`
- [](#VUID-vkCmdEndRendering2EXT-None-10612)VUID-vkCmdEndRendering2EXT-None-10612  
  This command **must** not be recorded when transform feedback is active
- [](#VUID-vkCmdEndRendering2EXT-None-10613)VUID-vkCmdEndRendering2EXT-None-10613  
  If `vkCmdBeginQuery`* was called within the render pass, the corresponding `vkCmdEndQuery`* **must** have been called subsequently within the same subpass

Valid Usage (Implicit)

- [](#VUID-vkCmdEndRendering2EXT-commandBuffer-parameter)VUID-vkCmdEndRendering2EXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdEndRendering2EXT-pRenderingEndInfo-parameter)VUID-vkCmdEndRendering2EXT-pRenderingEndInfo-parameter  
  If `pRenderingEndInfo` is not `NULL`, `pRenderingEndInfo` **must** be a valid pointer to a valid [VkRenderingEndInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingEndInfoEXT.html) structure
- [](#VUID-vkCmdEndRendering2EXT-commandBuffer-recording)VUID-vkCmdEndRendering2EXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdEndRendering2EXT-commandBuffer-cmdpool)VUID-vkCmdEndRendering2EXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics operations
- [](#VUID-vkCmdEndRendering2EXT-renderpass)VUID-vkCmdEndRendering2EXT-renderpass  
  This command **must** only be called inside of a render pass instance
- [](#VUID-vkCmdEndRendering2EXT-videocoding)VUID-vkCmdEndRendering2EXT-videocoding  
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

Action  
State

## [](#_see_also)See Also

[VK\_EXT\_fragment\_density\_map\_offset](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_fragment_density_map_offset.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkRenderingEndInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingEndInfoEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdEndRendering2EXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0