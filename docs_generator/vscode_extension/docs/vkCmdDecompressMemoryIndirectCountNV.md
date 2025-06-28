# vkCmdDecompressMemoryIndirectCountNV(3) Manual Page

## Name

vkCmdDecompressMemoryIndirectCountNV - Indirect decompress data between memory regions



## [](#_c_specification)C Specification

To decompress data between one or more memory regions by specifying decompression parameters indirectly in a buffer, call:

```c++
// Provided by VK_NV_memory_decompression
void vkCmdDecompressMemoryIndirectCountNV(
    VkCommandBuffer                             commandBuffer,
    VkDeviceAddress                             indirectCommandsAddress,
    VkDeviceAddress                             indirectCommandsCountAddress,
    uint32_t                                    stride);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `indirectCommandsAddress` is the device address containing decompression parameters laid out as an array of [VkDecompressMemoryRegionNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDecompressMemoryRegionNV.html) structures.
- `indirectCommandsCountAddress` is the device address containing the decompression count.
- `stride` is the byte stride between successive sets of decompression parameters located starting from `indirectCommandsAddress`.

## [](#_description)Description

Each region specified in `indirectCommandsAddress` is decompressed from the source to destination region based on the specified decompression method.

Valid Usage

- [](#VUID-vkCmdDecompressMemoryIndirectCountNV-None-07692)VUID-vkCmdDecompressMemoryIndirectCountNV-None-07692  
  The [`memoryDecompression`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-memoryDecompression) feature **must** be enabled
- [](#VUID-vkCmdDecompressMemoryIndirectCountNV-indirectCommandsAddress-07693)VUID-vkCmdDecompressMemoryIndirectCountNV-indirectCommandsAddress-07693  
  If `indirectCommandsAddress` comes from a non-sparse buffer then it **must** be bound completely and contiguously to a single `VkDeviceMemory` object
- [](#VUID-vkCmdDecompressMemoryIndirectCountNV-indirectCommandsAddress-07694)VUID-vkCmdDecompressMemoryIndirectCountNV-indirectCommandsAddress-07694  
  The [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) that `indirectCommandsAddress` comes from **must** have been created with the `VK_BUFFER_USAGE_INDIRECT_BUFFER_BIT` bit set
- [](#VUID-vkCmdDecompressMemoryIndirectCountNV-offset-07695)VUID-vkCmdDecompressMemoryIndirectCountNV-offset-07695  
  `offset` **must** be a multiple of `4`
- [](#VUID-vkCmdDecompressMemoryIndirectCountNV-indirectCommandsCountAddress-07696)VUID-vkCmdDecompressMemoryIndirectCountNV-indirectCommandsCountAddress-07696  
  If `indirectCommandsCountAddress` comes from a non-sparse buffer then it **must** be bound completely and contiguously to a single `VkDeviceMemory` object
- [](#VUID-vkCmdDecompressMemoryIndirectCountNV-indirectCommandsCountAddress-07697)VUID-vkCmdDecompressMemoryIndirectCountNV-indirectCommandsCountAddress-07697  
  The [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) that `indirectCommandsCountAddress` comes from **must** have been created with the `VK_BUFFER_USAGE_INDIRECT_BUFFER_BIT` bit set
- [](#VUID-vkCmdDecompressMemoryIndirectCountNV-indirectCommandsCountAddress-07698)VUID-vkCmdDecompressMemoryIndirectCountNV-indirectCommandsCountAddress-07698  
  `indirectCommandsCountAddress` **must** be a multiple of `4`
- [](#VUID-vkCmdDecompressMemoryIndirectCountNV-indirectCommandsCountAddress-07699)VUID-vkCmdDecompressMemoryIndirectCountNV-indirectCommandsCountAddress-07699  
  The count stored in `indirectCommandsCountAddress` **must** be less than or equal to `VkPhysicalDeviceMemoryDecompressionPropertiesNV`::`maxDecompressionIndirectCount`
- [](#VUID-vkCmdDecompressMemoryIndirectCountNV-stride-07700)VUID-vkCmdDecompressMemoryIndirectCountNV-stride-07700  
  `stride` **must** be a multiple of `4` and **must** be greater than or equal to sizeof(`VkDecompressMemoryRegionNV`)
- [](#VUID-vkCmdDecompressMemoryIndirectCountNV-indirectCommandsCountAddress-07701)VUID-vkCmdDecompressMemoryIndirectCountNV-indirectCommandsCountAddress-07701  
  If the count stored in `indirectCommandsCountAddress` is equal to `1`, (`offset` + sizeof(`VkDecompressMemoryRegionNV`)) **must** be less than or equal to the size of the [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) that `indirectCommandsAddress` comes from
- [](#VUID-vkCmdDecompressMemoryIndirectCountNV-indirectCommandsCountAddress-07702)VUID-vkCmdDecompressMemoryIndirectCountNV-indirectCommandsCountAddress-07702  
  If the count stored in `indirectCommandsCountAddress` is greater than `1`, `indirectCommandsAddress` + sizeof(`VkDecompressMemoryRegionNV`) + (`stride` × (count stored in `countBuffer` - 1)) **must** be less than or equal to the last valid address in the [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) that `indirectCommandsAddress` was created from

Valid Usage (Implicit)

- [](#VUID-vkCmdDecompressMemoryIndirectCountNV-commandBuffer-parameter)VUID-vkCmdDecompressMemoryIndirectCountNV-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdDecompressMemoryIndirectCountNV-commandBuffer-recording)VUID-vkCmdDecompressMemoryIndirectCountNV-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdDecompressMemoryIndirectCountNV-commandBuffer-cmdpool)VUID-vkCmdDecompressMemoryIndirectCountNV-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics, or compute operations
- [](#VUID-vkCmdDecompressMemoryIndirectCountNV-renderpass)VUID-vkCmdDecompressMemoryIndirectCountNV-renderpass  
  This command **must** only be called outside of a render pass instance
- [](#VUID-vkCmdDecompressMemoryIndirectCountNV-videocoding)VUID-vkCmdDecompressMemoryIndirectCountNV-videocoding  
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

Graphics  
Compute

Action

## [](#_see_also)See Also

[VK\_NV\_memory\_decompression](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_memory_decompression.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdDecompressMemoryIndirectCountNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0