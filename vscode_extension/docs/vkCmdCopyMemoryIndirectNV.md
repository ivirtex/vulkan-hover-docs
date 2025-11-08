# vkCmdCopyMemoryIndirectNV(3) Manual Page

## Name

vkCmdCopyMemoryIndirectNV - Copy data between memory regions



## [](#_c_specification)C Specification

To copy data between two memory regions by specifying copy parameters indirectly in memory, call:

```c++
// Provided by VK_NV_copy_memory_indirect
void vkCmdCopyMemoryIndirectNV(
    VkCommandBuffer                             commandBuffer,
    VkDeviceAddress                             copyBufferAddress,
    uint32_t                                    copyCount,
    uint32_t                                    stride);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `copyBufferAddress` is the memory address specifying the copy parameters. It is laid out as an array of [VkCopyMemoryIndirectCommandNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyMemoryIndirectCommandNV.html) structures.
- `copyCount` is the number of copies to execute, and **can** be zero.
- `stride` is the stride in bytes between successive sets of copy parameters.

## [](#_description)Description

Each region read from `copyBufferAddress` is copied from the source region to the specified destination region. The results are undefined if any of the source and destination regions overlap in memory.

Valid Usage

- [](#VUID-vkCmdCopyMemoryIndirectNV-None-07653)VUID-vkCmdCopyMemoryIndirectNV-None-07653  
  The [`indirectCopy`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-indirectCopy) feature **must** be enabled
- [](#VUID-vkCmdCopyMemoryIndirectNV-copyBufferAddress-07654)VUID-vkCmdCopyMemoryIndirectNV-copyBufferAddress-07654  
  `copyBufferAddress` **must** be 4 byte aligned
- [](#VUID-vkCmdCopyMemoryIndirectNV-stride-07655)VUID-vkCmdCopyMemoryIndirectNV-stride-07655  
  `stride` **must** be a multiple of `4` and **must** be greater than or equal to sizeof([VkCopyMemoryIndirectCommandNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyMemoryIndirectCommandNV.html))
- [](#VUID-vkCmdCopyMemoryIndirectNV-commandBuffer-07656)VUID-vkCmdCopyMemoryIndirectNV-commandBuffer-07656  
  The [VkCommandPool](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandPool.html) that `commandBuffer` was allocated from **must** support at least one of the queue types specified in [VkPhysicalDeviceCopyMemoryIndirectPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceCopyMemoryIndirectPropertiesKHR.html)::`supportedQueues`
- [](#VUID-vkCmdCopyMemoryIndirectNV-copyBufferAddress-10946)VUID-vkCmdCopyMemoryIndirectNV-copyBufferAddress-10946  
  Any of the source or destination memory regions specified in `copyBufferAddress` **must** not overlap with any of the specified destination memory regions

Valid Usage (Implicit)

- [](#VUID-vkCmdCopyMemoryIndirectNV-commandBuffer-parameter)VUID-vkCmdCopyMemoryIndirectNV-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdCopyMemoryIndirectNV-copyBufferAddress-parameter)VUID-vkCmdCopyMemoryIndirectNV-copyBufferAddress-parameter  
  `copyBufferAddress` **must** be a valid [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html) value
- [](#VUID-vkCmdCopyMemoryIndirectNV-commandBuffer-recording)VUID-vkCmdCopyMemoryIndirectNV-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdCopyMemoryIndirectNV-commandBuffer-cmdpool)VUID-vkCmdCopyMemoryIndirectNV-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support VK\_QUEUE\_COMPUTE\_BIT, VK\_QUEUE\_GRAPHICS\_BIT, or VK\_QUEUE\_TRANSFER\_BIT operations
- [](#VUID-vkCmdCopyMemoryIndirectNV-renderpass)VUID-vkCmdCopyMemoryIndirectNV-renderpass  
  This command **must** only be called outside of a render pass instance
- [](#VUID-vkCmdCopyMemoryIndirectNV-videocoding)VUID-vkCmdCopyMemoryIndirectNV-videocoding  
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
VK\_QUEUE\_TRANSFER\_BIT

Action

Conditional Rendering

vkCmdCopyMemoryIndirectNV is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_NV\_copy\_memory\_indirect](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_copy_memory_indirect.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdCopyMemoryIndirectNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0