# vkCmdBindTileMemoryQCOM(3) Manual Page

## Name

vkCmdBindTileMemoryQCOM - Bind tile memory to a command buffer



## [](#_c_specification)C Specification

To bind a range of tile memory to the command buffer, call:

```c++
// Provided by VK_QCOM_tile_memory_heap
void vkCmdBindTileMemoryQCOM(
    VkCommandBuffer                             commandBuffer,
    const VkTileMemoryBindInfoQCOM*             pTileMemoryBindInfo);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer that the tile memory will be bound to.
- `pTileMemoryBindInfo` is [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) or a pointer to a [VkTileMemoryBindInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTileMemoryBindInfoQCOM.html) structure defining how tile memory is bound.

## [](#_description)Description

Calling [vkCmdBindTileMemoryQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBindTileMemoryQCOM.html) when `pTileMemoryBindInfo` is [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) is equivalent to binding no tile memory to the command buffer.

Valid Usage (Implicit)

- [](#VUID-vkCmdBindTileMemoryQCOM-commandBuffer-parameter)VUID-vkCmdBindTileMemoryQCOM-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdBindTileMemoryQCOM-pTileMemoryBindInfo-parameter)VUID-vkCmdBindTileMemoryQCOM-pTileMemoryBindInfo-parameter  
  If `pTileMemoryBindInfo` is not `NULL`, `pTileMemoryBindInfo` **must** be a valid pointer to a valid [VkTileMemoryBindInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTileMemoryBindInfoQCOM.html) structure
- [](#VUID-vkCmdBindTileMemoryQCOM-commandBuffer-recording)VUID-vkCmdBindTileMemoryQCOM-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdBindTileMemoryQCOM-commandBuffer-cmdpool)VUID-vkCmdBindTileMemoryQCOM-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support VK\_QUEUE\_COMPUTE\_BIT, or VK\_QUEUE\_GRAPHICS\_BIT operations
- [](#VUID-vkCmdBindTileMemoryQCOM-renderpass)VUID-vkCmdBindTileMemoryQCOM-renderpass  
  This command **must** only be called outside of a render pass instance
- [](#VUID-vkCmdBindTileMemoryQCOM-videocoding)VUID-vkCmdBindTileMemoryQCOM-videocoding  
  This command **must** only be called outside of a video coding scope

Host Synchronization

- Host access to `commandBuffer` **must** be externally synchronized
- Host access to the `VkCommandPool` that `commandBuffer` was allocated from **must** be externally synchronized

Command Properties

     [Command Buffer Levels](#VkCommandBufferLevel) [Render Pass Scope](#vkCmdBeginRenderPass) [Video Coding Scope](#vkCmdBeginVideoCodingKHR) [Supported Queue Types](#VkQueueFlagBits) [Command Type](#fundamentals-queueoperation-command-types)

Primary  
Secondary

Outside

Outside

VK\_QUEUE\_COMPUTE\_BIT  
VK\_QUEUE\_GRAPHICS\_BIT

State

Conditional Rendering

vkCmdBindTileMemoryQCOM is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_QCOM\_tile\_memory\_heap](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_QCOM_tile_memory_heap.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkTileMemoryBindInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTileMemoryBindInfoQCOM.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdBindTileMemoryQCOM).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0