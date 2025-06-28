# vkCmdFillBuffer(3) Manual Page

## Name

vkCmdFillBuffer - Fill a region of a buffer with a fixed value



## [](#_c_specification)C Specification

To clear buffer data, call:

```c++
// Provided by VK_VERSION_1_0
void vkCmdFillBuffer(
    VkCommandBuffer                             commandBuffer,
    VkBuffer                                    dstBuffer,
    VkDeviceSize                                dstOffset,
    VkDeviceSize                                size,
    uint32_t                                    data);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `dstBuffer` is the buffer to be filled.
- `dstOffset` is the byte offset into the buffer at which to start filling, and **must** be a multiple of 4.
- `size` is the number of bytes to fill, and **must** be either a multiple of 4, or `VK_WHOLE_SIZE` to fill the range from `offset` to the end of the buffer. If `VK_WHOLE_SIZE` is used and the remaining size of the buffer is not a multiple of 4, then the nearest smaller multiple is used.
- `data` is the 4-byte word written repeatedly to the buffer to fill `size` bytes of data. The data word is written to memory according to the host endianness.

## [](#_description)Description

`vkCmdFillBuffer` is treated as a “transfer” operation for the purposes of synchronization barriers. The `VK_BUFFER_USAGE_TRANSFER_DST_BIT` **must** be specified in `usage` of [VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCreateInfo.html) in order for the buffer to be compatible with `vkCmdFillBuffer`.

Valid Usage

- [](#VUID-vkCmdFillBuffer-dstOffset-00024)VUID-vkCmdFillBuffer-dstOffset-00024  
  `dstOffset` **must** be less than the size of `dstBuffer`
- [](#VUID-vkCmdFillBuffer-dstOffset-00025)VUID-vkCmdFillBuffer-dstOffset-00025  
  `dstOffset` **must** be a multiple of `4`
- [](#VUID-vkCmdFillBuffer-size-00026)VUID-vkCmdFillBuffer-size-00026  
  If `size` is not equal to `VK_WHOLE_SIZE`, `size` **must** be greater than `0`
- [](#VUID-vkCmdFillBuffer-size-00027)VUID-vkCmdFillBuffer-size-00027  
  If `size` is not equal to `VK_WHOLE_SIZE`, `size` **must** be less than or equal to the size of `dstBuffer` minus `dstOffset`
- [](#VUID-vkCmdFillBuffer-size-00028)VUID-vkCmdFillBuffer-size-00028  
  If `size` is not equal to `VK_WHOLE_SIZE`, `size` **must** be a multiple of `4`
- [](#VUID-vkCmdFillBuffer-dstBuffer-00029)VUID-vkCmdFillBuffer-dstBuffer-00029  
  `dstBuffer` **must** have been created with `VK_BUFFER_USAGE_TRANSFER_DST_BIT` usage flag
- [](#VUID-vkCmdFillBuffer-apiVersion-07894)VUID-vkCmdFillBuffer-apiVersion-07894  
  If the [VK\_KHR\_maintenance1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance1.html) extension is not enabled and [VkPhysicalDeviceProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties.html)::`apiVersion` is less than Vulkan 1.1, the [VkCommandPool](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandPool.html) that `commandBuffer` was allocated from **must** support graphics or compute operations
- [](#VUID-vkCmdFillBuffer-dstBuffer-00031)VUID-vkCmdFillBuffer-dstBuffer-00031  
  If `dstBuffer` is non-sparse then it **must** be bound completely and contiguously to a single [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) object
- [](#VUID-vkCmdFillBuffer-commandBuffer-01811)VUID-vkCmdFillBuffer-commandBuffer-01811  
  If `commandBuffer` is an unprotected command buffer and [`protectedNoFault`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-protectedNoFault) is not supported, `dstBuffer` **must** not be a protected buffer
- [](#VUID-vkCmdFillBuffer-commandBuffer-01812)VUID-vkCmdFillBuffer-commandBuffer-01812  
  If `commandBuffer` is a protected command buffer and [`protectedNoFault`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-protectedNoFault) is not supported, `dstBuffer` **must** not be an unprotected buffer

Valid Usage (Implicit)

- [](#VUID-vkCmdFillBuffer-commandBuffer-parameter)VUID-vkCmdFillBuffer-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdFillBuffer-dstBuffer-parameter)VUID-vkCmdFillBuffer-dstBuffer-parameter  
  `dstBuffer` **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) handle
- [](#VUID-vkCmdFillBuffer-commandBuffer-recording)VUID-vkCmdFillBuffer-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdFillBuffer-commandBuffer-cmdpool)VUID-vkCmdFillBuffer-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support transfer, graphics or compute operations
- [](#VUID-vkCmdFillBuffer-renderpass)VUID-vkCmdFillBuffer-renderpass  
  This command **must** only be called outside of a render pass instance
- [](#VUID-vkCmdFillBuffer-videocoding)VUID-vkCmdFillBuffer-videocoding  
  This command **must** only be called outside of a video coding scope
- [](#VUID-vkCmdFillBuffer-commonparent)VUID-vkCmdFillBuffer-commonparent  
  Both of `commandBuffer`, and `dstBuffer` **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

Host Synchronization

- Host access to `commandBuffer` **must** be externally synchronized
- Host access to the `VkCommandPool` that `commandBuffer` was allocated from **must** be externally synchronized

Command Properties

     [Command Buffer Levels](#VkCommandBufferLevel) [Render Pass Scope](#vkCmdBeginRenderPass) [Video Coding Scope](#vkCmdBeginVideoCodingKHR) [Supported Queue Types](#VkQueueFlagBits) [Command Type](#fundamentals-queueoperation-command-types)

Primary  
Secondary

Outside

Outside

Transfer  
Graphics  
Compute

Action

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdFillBuffer)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0