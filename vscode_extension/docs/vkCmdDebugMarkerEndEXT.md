# vkCmdDebugMarkerEndEXT(3) Manual Page

## Name

vkCmdDebugMarkerEndEXT - Close a command buffer marker region



## [](#_c_specification)C Specification

A marker region can be closed by calling:

```c++
// Provided by VK_EXT_debug_marker
void vkCmdDebugMarkerEndEXT(
    VkCommandBuffer                             commandBuffer);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command is recorded.

## [](#_description)Description

An application **may** open a marker region in one command buffer and close it in another, or otherwise split marker regions across multiple command buffers or multiple queue submissions. When viewed from the linear series of submissions to a single queue, the calls to `vkCmdDebugMarkerBeginEXT` and `vkCmdDebugMarkerEndEXT` **must** be matched and balanced.

Valid Usage

- [](#VUID-vkCmdDebugMarkerEndEXT-commandBuffer-01239)VUID-vkCmdDebugMarkerEndEXT-commandBuffer-01239  
  There **must** be an outstanding [vkCmdDebugMarkerBeginEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDebugMarkerBeginEXT.html) command prior to the `vkCmdDebugMarkerEndEXT` on the queue that `commandBuffer` is submitted to
- [](#VUID-vkCmdDebugMarkerEndEXT-commandBuffer-01240)VUID-vkCmdDebugMarkerEndEXT-commandBuffer-01240  
  If `commandBuffer` is a secondary command buffer, there **must** be an outstanding [vkCmdDebugMarkerBeginEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDebugMarkerBeginEXT.html) command recorded to `commandBuffer` that has not previously been ended by a call to [vkCmdDebugMarkerEndEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDebugMarkerEndEXT.html)
- [](#VUID-vkCmdDebugMarkerEndEXT-None-10615)VUID-vkCmdDebugMarkerEndEXT-None-10615  
  This command **must** not be recorded when [per-tile execution model](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-per-tile-execution-model) is enabled

Valid Usage (Implicit)

- [](#VUID-vkCmdDebugMarkerEndEXT-commandBuffer-parameter)VUID-vkCmdDebugMarkerEndEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdDebugMarkerEndEXT-commandBuffer-recording)VUID-vkCmdDebugMarkerEndEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdDebugMarkerEndEXT-commandBuffer-cmdpool)VUID-vkCmdDebugMarkerEndEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support transfer, graphics, compute, decode, encode, or optical flow operations

Host Synchronization

- Host access to `commandBuffer` **must** be externally synchronized
- Host access to the `VkCommandPool` that `commandBuffer` was allocated from **must** be externally synchronized

Command Properties

     [Command Buffer Levels](#VkCommandBufferLevel) [Render Pass Scope](#vkCmdBeginRenderPass) [Video Coding Scope](#vkCmdBeginVideoCodingKHR) [Supported Queue Types](#VkQueueFlagBits) [Command Type](#fundamentals-queueoperation-command-types)

Primary  
Secondary

Both

Both

Transfer  
Graphics  
Compute  
Decode  
Encode  
Opticalflow

Action

Conditional Rendering

vkCmdDebugMarkerEndEXT is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_EXT\_debug\_marker](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_debug_marker.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdDebugMarkerEndEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0