# vkCmdEndPerTileExecutionQCOM(3) Manual Page

## Name

vkCmdEndPerTileExecutionQCOM - End per-tile execution mode



## [](#_c_specification)C Specification

To disable per-tile execution model, call:

```c++
// Provided by VK_QCOM_tile_shading
void vkCmdEndPerTileExecutionQCOM(
    VkCommandBuffer                             commandBuffer,
    const VkPerTileEndInfoQCOM*                 pPerTileEndInfo);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer in which to record the command.
- `pPerTileEndInfo` is a pointer to a [VkPerTileEndInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerTileEndInfoQCOM.html) structure containing information about how the *per-tile execution model* is ended.

## [](#_description)Description

This command disables *per-tile execution model*.

Valid Usage

- [](#VUID-vkCmdEndPerTileExecutionQCOM-None-10666)VUID-vkCmdEndPerTileExecutionQCOM-None-10666  
  The *per-tile execution model* **must** have been enabled in the current render pass
- [](#VUID-vkCmdEndPerTileExecutionQCOM-None-10667)VUID-vkCmdEndPerTileExecutionQCOM-None-10667  
  The current render pass **must** be a [tile shading render pass](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-tile-shading)

Valid Usage (Implicit)

- [](#VUID-vkCmdEndPerTileExecutionQCOM-commandBuffer-parameter)VUID-vkCmdEndPerTileExecutionQCOM-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdEndPerTileExecutionQCOM-pPerTileEndInfo-parameter)VUID-vkCmdEndPerTileExecutionQCOM-pPerTileEndInfo-parameter  
  `pPerTileEndInfo` **must** be a valid pointer to a valid [VkPerTileEndInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerTileEndInfoQCOM.html) structure
- [](#VUID-vkCmdEndPerTileExecutionQCOM-commandBuffer-recording)VUID-vkCmdEndPerTileExecutionQCOM-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdEndPerTileExecutionQCOM-commandBuffer-cmdpool)VUID-vkCmdEndPerTileExecutionQCOM-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics, or compute operations
- [](#VUID-vkCmdEndPerTileExecutionQCOM-renderpass)VUID-vkCmdEndPerTileExecutionQCOM-renderpass  
  This command **must** only be called inside of a render pass instance
- [](#VUID-vkCmdEndPerTileExecutionQCOM-videocoding)VUID-vkCmdEndPerTileExecutionQCOM-videocoding  
  This command **must** only be called outside of a video coding scope

Host Synchronization

- Host access to the `VkCommandPool` that `commandBuffer` was allocated from **must** be externally synchronized

Command Properties

     [Command Buffer Levels](#VkCommandBufferLevel) [Render Pass Scope](#vkCmdBeginRenderPass) [Video Coding Scope](#vkCmdBeginVideoCodingKHR) [Supported Queue Types](#VkQueueFlagBits) [Command Type](#fundamentals-queueoperation-command-types)

Primary  
Secondary

Inside

Outside

Graphics  
Compute

State

Conditional Rendering

vkCmdEndPerTileExecutionQCOM is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_QCOM\_tile\_shading](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_QCOM_tile_shading.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkPerTileEndInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerTileEndInfoQCOM.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdEndPerTileExecutionQCOM).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0