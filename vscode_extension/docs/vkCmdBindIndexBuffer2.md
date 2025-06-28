# vkCmdBindIndexBuffer2(3) Manual Page

## Name

vkCmdBindIndexBuffer2 - Bind an index buffer to a command buffer



## [](#_c_specification)C Specification

To bind an index buffer, along with its size, to a command buffer, call:

```c++
// Provided by VK_VERSION_1_4
void vkCmdBindIndexBuffer2(
    VkCommandBuffer                             commandBuffer,
    VkBuffer                                    buffer,
    VkDeviceSize                                offset,
    VkDeviceSize                                size,
    VkIndexType                                 indexType);
```

or the equivalent command

```c++
// Provided by VK_KHR_maintenance5
void vkCmdBindIndexBuffer2KHR(
    VkCommandBuffer                             commandBuffer,
    VkBuffer                                    buffer,
    VkDeviceSize                                offset,
    VkDeviceSize                                size,
    VkIndexType                                 indexType);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command is recorded.
- `buffer` is the buffer being bound.
- `offset` is the starting offset in bytes within `buffer` used in index buffer address calculations.
- `size` is the size in bytes of index data bound from `buffer`.
- `indexType` is a [VkIndexType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndexType.html) value specifying the size of the indices.

## [](#_description)Description

`size` specifies the bound size of the index buffer starting from `offset`. If `size` is `VK_WHOLE_SIZE` then the bound size is from `offset` to the end of the `buffer`.

If the [`maintenance6`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-maintenance6) feature is enabled, `buffer` **can** be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html). If `buffer` is [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) and the [`nullDescriptor`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-nullDescriptor) feature is enabled, every index fetched results in a value of zero.

Valid Usage

- [](#VUID-vkCmdBindIndexBuffer2-offset-08782)VUID-vkCmdBindIndexBuffer2-offset-08782  
  `offset` **must** be less than the size of `buffer`
- [](#VUID-vkCmdBindIndexBuffer2-offset-08783)VUID-vkCmdBindIndexBuffer2-offset-08783  
  The sum of `offset` and the base address of the range of `VkDeviceMemory` object that is backing `buffer`, **must** be a multiple of the size of the type indicated by `indexType`
- [](#VUID-vkCmdBindIndexBuffer2-buffer-08784)VUID-vkCmdBindIndexBuffer2-buffer-08784  
  `buffer` **must** have been created with the `VK_BUFFER_USAGE_INDEX_BUFFER_BIT` flag
- [](#VUID-vkCmdBindIndexBuffer2-buffer-08785)VUID-vkCmdBindIndexBuffer2-buffer-08785  
  If `buffer` is non-sparse then it **must** be bound completely and contiguously to a single `VkDeviceMemory` object
- [](#VUID-vkCmdBindIndexBuffer2-indexType-08786)VUID-vkCmdBindIndexBuffer2-indexType-08786  
  `indexType` **must** not be `VK_INDEX_TYPE_NONE_KHR`
- [](#VUID-vkCmdBindIndexBuffer2-indexType-08787)VUID-vkCmdBindIndexBuffer2-indexType-08787  
  If `indexType` is `VK_INDEX_TYPE_UINT8`, the [`indexTypeUint8`](#features-indexTypeUint8) feature **must** be enabled
- [](#VUID-vkCmdBindIndexBuffer2-None-09493)VUID-vkCmdBindIndexBuffer2-None-09493  
  If the [`maintenance6`](#features-maintenance6) feature is not enabled, `buffer` **must** not be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html)
- [](#VUID-vkCmdBindIndexBuffer2-buffer-09494)VUID-vkCmdBindIndexBuffer2-buffer-09494  
  If `buffer` is [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), offset **must** be zero
- [](#VUID-vkCmdBindIndexBuffer2-size-08767)VUID-vkCmdBindIndexBuffer2-size-08767  
  If `size` is not `VK_WHOLE_SIZE`, `size` **must** be a multiple of the size of the type indicated by `indexType`
- [](#VUID-vkCmdBindIndexBuffer2-size-08768)VUID-vkCmdBindIndexBuffer2-size-08768  
  If `size` is not `VK_WHOLE_SIZE`, the sum of `offset` and `size` **must** be less than or equal to the size of `buffer`

Valid Usage (Implicit)

- [](#VUID-vkCmdBindIndexBuffer2-commandBuffer-parameter)VUID-vkCmdBindIndexBuffer2-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdBindIndexBuffer2-buffer-parameter)VUID-vkCmdBindIndexBuffer2-buffer-parameter  
  If `buffer` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `buffer` **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) handle
- [](#VUID-vkCmdBindIndexBuffer2-indexType-parameter)VUID-vkCmdBindIndexBuffer2-indexType-parameter  
  `indexType` **must** be a valid [VkIndexType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndexType.html) value
- [](#VUID-vkCmdBindIndexBuffer2-commandBuffer-recording)VUID-vkCmdBindIndexBuffer2-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdBindIndexBuffer2-commandBuffer-cmdpool)VUID-vkCmdBindIndexBuffer2-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics operations
- [](#VUID-vkCmdBindIndexBuffer2-videocoding)VUID-vkCmdBindIndexBuffer2-videocoding  
  This command **must** only be called outside of a video coding scope
- [](#VUID-vkCmdBindIndexBuffer2-commonparent)VUID-vkCmdBindIndexBuffer2-commonparent  
  Both of `buffer`, and `commandBuffer` that are valid handles of non-ignored parameters **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

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

State

## [](#_see_also)See Also

[VK\_KHR\_maintenance5](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance5.html), [VK\_VERSION\_1\_4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_4.html), [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html), [VkIndexType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndexType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdBindIndexBuffer2)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0