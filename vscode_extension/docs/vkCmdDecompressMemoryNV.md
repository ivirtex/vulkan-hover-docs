# vkCmdDecompressMemoryNV(3) Manual Page

## Name

vkCmdDecompressMemoryNV - Decompress data between memory regions



## [](#_c_specification)C Specification

To decompress data between one or more memory regions call:

```c++
// Provided by VK_NV_memory_decompression
void vkCmdDecompressMemoryNV(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    decompressRegionCount,
    const VkDecompressMemoryRegionNV*           pDecompressMemoryRegions);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `decompressRegionCount` is the number of memory regions to decompress.
- `pDecompressMemoryRegions` is a pointer to an array of `decompressRegionCount` [VkDecompressMemoryRegionNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDecompressMemoryRegionNV.html) structures specifying decompression parameters.

## [](#_description)Description

Each region specified in `pDecompressMemoryRegions` is decompressed from the source to destination region based on the specified decompression method.

Valid Usage

- [](#VUID-vkCmdDecompressMemoryNV-None-07684)VUID-vkCmdDecompressMemoryNV-None-07684  
  The [`memoryDecompression`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-memoryDecompression) feature **must** be enabled

Valid Usage (Implicit)

- [](#VUID-vkCmdDecompressMemoryNV-commandBuffer-parameter)VUID-vkCmdDecompressMemoryNV-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdDecompressMemoryNV-pDecompressMemoryRegions-parameter)VUID-vkCmdDecompressMemoryNV-pDecompressMemoryRegions-parameter  
  `pDecompressMemoryRegions` **must** be a valid pointer to an array of `decompressRegionCount` valid [VkDecompressMemoryRegionNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDecompressMemoryRegionNV.html) structures
- [](#VUID-vkCmdDecompressMemoryNV-commandBuffer-recording)VUID-vkCmdDecompressMemoryNV-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdDecompressMemoryNV-commandBuffer-cmdpool)VUID-vkCmdDecompressMemoryNV-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support VK\_QUEUE\_COMPUTE\_BIT, or VK\_QUEUE\_GRAPHICS\_BIT operations
- [](#VUID-vkCmdDecompressMemoryNV-renderpass)VUID-vkCmdDecompressMemoryNV-renderpass  
  This command **must** only be called outside of a render pass instance
- [](#VUID-vkCmdDecompressMemoryNV-videocoding)VUID-vkCmdDecompressMemoryNV-videocoding  
  This command **must** only be called outside of a video coding scope
- [](#VUID-vkCmdDecompressMemoryNV-decompressRegionCount-arraylength)VUID-vkCmdDecompressMemoryNV-decompressRegionCount-arraylength  
  `decompressRegionCount` **must** be greater than `0`

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

Action

Conditional Rendering

vkCmdDecompressMemoryNV is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_NV\_memory\_decompression](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_memory_decompression.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkDecompressMemoryRegionNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDecompressMemoryRegionNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdDecompressMemoryNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0