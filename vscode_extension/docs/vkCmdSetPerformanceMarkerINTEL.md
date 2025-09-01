# vkCmdSetPerformanceMarkerINTEL(3) Manual Page

## Name

vkCmdSetPerformanceMarkerINTEL - Markers



## [](#_c_specification)C Specification

To help associate query results with a particular point at which an application emitted commands, markers can be set into the command buffers with the call:

```c++
// Provided by VK_INTEL_performance_query
VkResult vkCmdSetPerformanceMarkerINTEL(
    VkCommandBuffer                             commandBuffer,
    const VkPerformanceMarkerInfoINTEL*         pMarkerInfo);
```

## [](#_description)Description

The last marker set onto a command buffer before the end of a query will be part of the query result.

Valid Usage (Implicit)

- [](#VUID-vkCmdSetPerformanceMarkerINTEL-commandBuffer-parameter)VUID-vkCmdSetPerformanceMarkerINTEL-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdSetPerformanceMarkerINTEL-pMarkerInfo-parameter)VUID-vkCmdSetPerformanceMarkerINTEL-pMarkerInfo-parameter  
  `pMarkerInfo` **must** be a valid pointer to a valid [VkPerformanceMarkerInfoINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceMarkerInfoINTEL.html) structure
- [](#VUID-vkCmdSetPerformanceMarkerINTEL-commandBuffer-recording)VUID-vkCmdSetPerformanceMarkerINTEL-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdSetPerformanceMarkerINTEL-commandBuffer-cmdpool)VUID-vkCmdSetPerformanceMarkerINTEL-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics, compute, or transfer operations
- [](#VUID-vkCmdSetPerformanceMarkerINTEL-videocoding)VUID-vkCmdSetPerformanceMarkerINTEL-videocoding  
  This command **must** only be called outside of a video coding scope

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
Compute  
Transfer

Action  
State

Conditional Rendering

vkCmdSetPerformanceMarkerINTEL is not affected by [conditional rendering](#drawing-conditional-rendering)

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_TOO_MANY_OBJECTS`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_INTEL\_performance\_query](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_INTEL_performance_query.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkPerformanceMarkerInfoINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceMarkerInfoINTEL.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdSetPerformanceMarkerINTEL).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0