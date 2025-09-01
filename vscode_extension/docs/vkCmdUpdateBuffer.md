# vkCmdUpdateBuffer(3) Manual Page

## Name

vkCmdUpdateBuffer - Update a buffer's contents from host memory



## [](#_c_specification)C Specification

To update buffer data inline in a command buffer, call:

```c++
// Provided by VK_VERSION_1_0
void vkCmdUpdateBuffer(
    VkCommandBuffer                             commandBuffer,
    VkBuffer                                    dstBuffer,
    VkDeviceSize                                dstOffset,
    VkDeviceSize                                dataSize,
    const void*                                 pData);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `dstBuffer` is a handle to the buffer to be updated.
- `dstOffset` is the byte offset into the buffer to start updating, and **must** be a multiple of 4.
- `dataSize` is the number of bytes to update, and **must** be a multiple of 4.
- `pData` is a pointer to the source data for the buffer update, and **must** be at least `dataSize` bytes in size.

## [](#_description)Description

`dataSize` **must** be less than or equal to 65536 bytes. For larger updates, applications **can** use buffer to buffer [copies](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#copies-buffers).

Note

Buffer updates performed with `vkCmdUpdateBuffer` first copy the data into command buffer memory when the command is recorded (which requires additional storage and may incur an additional allocation), and then copy the data from the command buffer into `dstBuffer` when the command is executed on a device.

The additional cost of this functionality compared to [buffer to buffer copies](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#copies-buffers) means it should only be used for very small amounts of data, and is why it is limited to at most 65536 bytes. Applications **can** work around this restriction by issuing multiple `vkCmdUpdateBuffer` commands to different ranges of the same buffer, but doing so is not recommended.

The source data is copied from `pData` to the command buffer when the command is called.

`vkCmdUpdateBuffer` is only allowed outside of a render pass. This command is treated as a “transfer” operation for the purposes of synchronization barriers. The `VK_BUFFER_USAGE_TRANSFER_DST_BIT` **must** be specified in `usage` of [VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCreateInfo.html) in order for the buffer to be compatible with `vkCmdUpdateBuffer`.

Valid Usage

- [](#VUID-vkCmdUpdateBuffer-dstOffset-00032)VUID-vkCmdUpdateBuffer-dstOffset-00032  
  `dstOffset` **must** be less than the size of `dstBuffer`
- [](#VUID-vkCmdUpdateBuffer-dataSize-00033)VUID-vkCmdUpdateBuffer-dataSize-00033  
  `dataSize` **must** be less than or equal to the size of `dstBuffer` minus `dstOffset`
- [](#VUID-vkCmdUpdateBuffer-dstBuffer-00034)VUID-vkCmdUpdateBuffer-dstBuffer-00034  
  `dstBuffer` **must** have been created with `VK_BUFFER_USAGE_TRANSFER_DST_BIT` usage flag
- [](#VUID-vkCmdUpdateBuffer-dstBuffer-00035)VUID-vkCmdUpdateBuffer-dstBuffer-00035  
  If `dstBuffer` is non-sparse then it **must** be bound completely and contiguously to a single `VkDeviceMemory` object
- [](#VUID-vkCmdUpdateBuffer-dstOffset-00036)VUID-vkCmdUpdateBuffer-dstOffset-00036  
  `dstOffset` **must** be a multiple of `4`
- [](#VUID-vkCmdUpdateBuffer-dataSize-00037)VUID-vkCmdUpdateBuffer-dataSize-00037  
  `dataSize` **must** be less than or equal to `65536`
- [](#VUID-vkCmdUpdateBuffer-dataSize-00038)VUID-vkCmdUpdateBuffer-dataSize-00038  
  `dataSize` **must** be a multiple of `4`
- [](#VUID-vkCmdUpdateBuffer-commandBuffer-01813)VUID-vkCmdUpdateBuffer-commandBuffer-01813  
  If `commandBuffer` is an unprotected command buffer and [`protectedNoFault`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-protectedNoFault) is not supported, `dstBuffer` **must** not be a protected buffer
- [](#VUID-vkCmdUpdateBuffer-commandBuffer-01814)VUID-vkCmdUpdateBuffer-commandBuffer-01814  
  If `commandBuffer` is a protected command buffer and [`protectedNoFault`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-protectedNoFault) is not supported, `dstBuffer` **must** not be an unprotected buffer

Valid Usage (Implicit)

- [](#VUID-vkCmdUpdateBuffer-commandBuffer-parameter)VUID-vkCmdUpdateBuffer-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdUpdateBuffer-dstBuffer-parameter)VUID-vkCmdUpdateBuffer-dstBuffer-parameter  
  `dstBuffer` **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) handle
- [](#VUID-vkCmdUpdateBuffer-pData-parameter)VUID-vkCmdUpdateBuffer-pData-parameter  
  `pData` **must** be a valid pointer to an array of `dataSize` bytes
- [](#VUID-vkCmdUpdateBuffer-commandBuffer-recording)VUID-vkCmdUpdateBuffer-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdUpdateBuffer-commandBuffer-cmdpool)VUID-vkCmdUpdateBuffer-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support transfer, graphics, or compute operations
- [](#VUID-vkCmdUpdateBuffer-renderpass)VUID-vkCmdUpdateBuffer-renderpass  
  This command **must** only be called outside of a render pass instance
- [](#VUID-vkCmdUpdateBuffer-videocoding)VUID-vkCmdUpdateBuffer-videocoding  
  This command **must** only be called outside of a video coding scope
- [](#VUID-vkCmdUpdateBuffer-dataSize-arraylength)VUID-vkCmdUpdateBuffer-dataSize-arraylength  
  `dataSize` **must** be greater than `0`
- [](#VUID-vkCmdUpdateBuffer-commonparent)VUID-vkCmdUpdateBuffer-commonparent  
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

Conditional Rendering

vkCmdUpdateBuffer is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdUpdateBuffer).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0