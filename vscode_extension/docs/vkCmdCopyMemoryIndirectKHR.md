# vkCmdCopyMemoryIndirectKHR(3) Manual Page

## Name

vkCmdCopyMemoryIndirectKHR - Copy data between memory regions



## [](#_c_specification)C Specification

To copy data between two memory regions by specifying copy parameters indirectly in memory, call:

```c++
// Provided by VK_KHR_copy_memory_indirect
void vkCmdCopyMemoryIndirectKHR(
    VkCommandBuffer                             commandBuffer,
    const VkCopyMemoryIndirectInfoKHR*          pCopyMemoryIndirectInfo);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `pCopyMemoryIndirectInfo` is a pointer to a [VkCopyMemoryIndirectInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyMemoryIndirectInfoKHR.html) structure containing the copy parameters, including the number of copies to execute and a strided array of [VkCopyMemoryIndirectCommandKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyMemoryIndirectCommandKHR.html) structures.

## [](#_description)Description

Each region specified in the memory referenced by `pCopyMemoryIndirectInfo->copyAddressRange` is copied from the source region to the specified destination region. The results are undefined if any of the source and destination regions overlap in memory.

Valid Usage

- [](#VUID-vkCmdCopyMemoryIndirectKHR-indirectMemoryCopy-10935)VUID-vkCmdCopyMemoryIndirectKHR-indirectMemoryCopy-10935  
  The [`indirectMemoryCopy`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-indirectMemoryCopy) feature **must** be enabled
- [](#VUID-vkCmdCopyMemoryIndirectKHR-commandBuffer-10936)VUID-vkCmdCopyMemoryIndirectKHR-commandBuffer-10936  
  The [VkCommandPool](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandPool.html) that `commandBuffer` was allocated from **must** support at least one of the queue types specified in [VkPhysicalDeviceCopyMemoryIndirectPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceCopyMemoryIndirectPropertiesKHR.html)::`supportedQueues`
- [](#VUID-vkCmdCopyMemoryIndirectKHR-commandBuffer-10937)VUID-vkCmdCopyMemoryIndirectKHR-commandBuffer-10937  
  `commandBuffer` must not be a protected command buffer

Valid Usage (Implicit)

- [](#VUID-vkCmdCopyMemoryIndirectKHR-commandBuffer-parameter)VUID-vkCmdCopyMemoryIndirectKHR-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdCopyMemoryIndirectKHR-pCopyMemoryIndirectInfo-parameter)VUID-vkCmdCopyMemoryIndirectKHR-pCopyMemoryIndirectInfo-parameter  
  `pCopyMemoryIndirectInfo` **must** be a valid pointer to a valid [VkCopyMemoryIndirectInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyMemoryIndirectInfoKHR.html) structure
- [](#VUID-vkCmdCopyMemoryIndirectKHR-commandBuffer-recording)VUID-vkCmdCopyMemoryIndirectKHR-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdCopyMemoryIndirectKHR-commandBuffer-cmdpool)VUID-vkCmdCopyMemoryIndirectKHR-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support VK\_QUEUE\_COMPUTE\_BIT, VK\_QUEUE\_GRAPHICS\_BIT, or VK\_QUEUE\_TRANSFER\_BIT operations
- [](#VUID-vkCmdCopyMemoryIndirectKHR-renderpass)VUID-vkCmdCopyMemoryIndirectKHR-renderpass  
  This command **must** only be called outside of a render pass instance
- [](#VUID-vkCmdCopyMemoryIndirectKHR-videocoding)VUID-vkCmdCopyMemoryIndirectKHR-videocoding  
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

vkCmdCopyMemoryIndirectKHR is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_KHR\_copy\_memory\_indirect](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_copy_memory_indirect.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkCopyMemoryIndirectInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyMemoryIndirectInfoKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdCopyMemoryIndirectKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0