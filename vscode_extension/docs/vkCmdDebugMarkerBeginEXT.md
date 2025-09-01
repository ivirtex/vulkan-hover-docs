# vkCmdDebugMarkerBeginEXT(3) Manual Page

## Name

vkCmdDebugMarkerBeginEXT - Open a command buffer marker region



## [](#_c_specification)C Specification

A marker region can be opened by calling:

```c++
// Provided by VK_EXT_debug_marker
void vkCmdDebugMarkerBeginEXT(
    VkCommandBuffer                             commandBuffer,
    const VkDebugMarkerMarkerInfoEXT*           pMarkerInfo);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command is recorded.
- `pMarkerInfo` is a pointer to a [VkDebugMarkerMarkerInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugMarkerMarkerInfoEXT.html) structure specifying the parameters of the marker region to open.

## [](#_description)Description

Valid Usage

- [](#VUID-vkCmdDebugMarkerBeginEXT-None-10614)VUID-vkCmdDebugMarkerBeginEXT-None-10614  
  This command **must** not be recorded when [per-tile execution model](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-per-tile-execution-model) is enabled

Valid Usage (Implicit)

- [](#VUID-vkCmdDebugMarkerBeginEXT-commandBuffer-parameter)VUID-vkCmdDebugMarkerBeginEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdDebugMarkerBeginEXT-pMarkerInfo-parameter)VUID-vkCmdDebugMarkerBeginEXT-pMarkerInfo-parameter  
  `pMarkerInfo` **must** be a valid pointer to a valid [VkDebugMarkerMarkerInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugMarkerMarkerInfoEXT.html) structure
- [](#VUID-vkCmdDebugMarkerBeginEXT-commandBuffer-recording)VUID-vkCmdDebugMarkerBeginEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdDebugMarkerBeginEXT-commandBuffer-cmdpool)VUID-vkCmdDebugMarkerBeginEXT-commandBuffer-cmdpool  
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

vkCmdDebugMarkerBeginEXT is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_EXT\_debug\_marker](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_debug_marker.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkDebugMarkerMarkerInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugMarkerMarkerInfoEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdDebugMarkerBeginEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0