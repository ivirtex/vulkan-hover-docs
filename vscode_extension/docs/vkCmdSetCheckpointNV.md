# vkCmdSetCheckpointNV(3) Manual Page

## Name

vkCmdSetCheckpointNV - Insert diagnostic checkpoint in command stream



## [](#_c_specification)C Specification

Device diagnostic checkpoints are inserted into the command stream by calling [vkCmdSetCheckpointNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetCheckpointNV.html).

```c++
// Provided by VK_NV_device_diagnostic_checkpoints
void vkCmdSetCheckpointNV(
    VkCommandBuffer                             commandBuffer,
    const void*                                 pCheckpointMarker);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer that will receive the marker
- `pCheckpointMarker` is an opaque application-provided value that will be associated with the checkpoint.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkCmdSetCheckpointNV-commandBuffer-parameter)VUID-vkCmdSetCheckpointNV-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdSetCheckpointNV-commandBuffer-recording)VUID-vkCmdSetCheckpointNV-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdSetCheckpointNV-commandBuffer-cmdpool)VUID-vkCmdSetCheckpointNV-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics, compute, or transfer operations
- [](#VUID-vkCmdSetCheckpointNV-videocoding)VUID-vkCmdSetCheckpointNV-videocoding  
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

Conditional Rendering

vkCmdSetCheckpointNV is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_NV\_device\_diagnostic\_checkpoints](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_device_diagnostic_checkpoints.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdSetCheckpointNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0