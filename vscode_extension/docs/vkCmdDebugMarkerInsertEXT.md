# vkCmdDebugMarkerInsertEXT(3) Manual Page

## Name

vkCmdDebugMarkerInsertEXT - Insert a marker label into a command buffer



## [](#_c_specification)C Specification

A single marker label can be inserted into a command buffer by calling:

```c++
// Provided by VK_EXT_debug_marker
void vkCmdDebugMarkerInsertEXT(
    VkCommandBuffer                             commandBuffer,
    const VkDebugMarkerMarkerInfoEXT*           pMarkerInfo);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command is recorded.
- `pMarkerInfo` is a pointer to a [VkDebugMarkerMarkerInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugMarkerMarkerInfoEXT.html) structure specifying the parameters of the marker to insert.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkCmdDebugMarkerInsertEXT-commandBuffer-parameter)VUID-vkCmdDebugMarkerInsertEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdDebugMarkerInsertEXT-pMarkerInfo-parameter)VUID-vkCmdDebugMarkerInsertEXT-pMarkerInfo-parameter  
  `pMarkerInfo` **must** be a valid pointer to a valid [VkDebugMarkerMarkerInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugMarkerMarkerInfoEXT.html) structure
- [](#VUID-vkCmdDebugMarkerInsertEXT-commandBuffer-recording)VUID-vkCmdDebugMarkerInsertEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdDebugMarkerInsertEXT-commandBuffer-cmdpool)VUID-vkCmdDebugMarkerInsertEXT-commandBuffer-cmdpool  
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

vkCmdDebugMarkerInsertEXT is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_EXT\_debug\_marker](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_debug_marker.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkDebugMarkerMarkerInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugMarkerMarkerInfoEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdDebugMarkerInsertEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0