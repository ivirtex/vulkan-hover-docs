# vkCmdBindIndexBuffer(3) Manual Page

## Name

vkCmdBindIndexBuffer - Bind an index buffer to a command buffer



## [](#_c_specification)C Specification

To bind an index buffer to a command buffer, call:

```c++
// Provided by VK_VERSION_1_0
void vkCmdBindIndexBuffer(
    VkCommandBuffer                             commandBuffer,
    VkBuffer                                    buffer,
    VkDeviceSize                                offset,
    VkIndexType                                 indexType);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command is recorded.
- `buffer` is the buffer being bound.
- `offset` is the starting offset in bytes within `buffer` used in index buffer address calculations.
- `indexType` is a [VkIndexType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndexType.html) value specifying the size of the indices.

## [](#_description)Description

If the [`maintenance6`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-maintenance6) feature is enabled, `buffer` **can** be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html). If `buffer` is [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) and the [`nullDescriptor`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-nullDescriptor) feature is enabled, every index fetched results in a value of zero.

Valid Usage

- [](#VUID-vkCmdBindIndexBuffer-offset-08782)VUID-vkCmdBindIndexBuffer-offset-08782  
  `offset` **must** be less than the size of `buffer`
- [](#VUID-vkCmdBindIndexBuffer-offset-08783)VUID-vkCmdBindIndexBuffer-offset-08783  
  The sum of `offset` and the base address of the range of `VkDeviceMemory` object that is backing `buffer`, **must** be a multiple of the size of the type indicated by `indexType`
- [](#VUID-vkCmdBindIndexBuffer-buffer-08784)VUID-vkCmdBindIndexBuffer-buffer-08784  
  `buffer` **must** have been created with the `VK_BUFFER_USAGE_INDEX_BUFFER_BIT` flag
- [](#VUID-vkCmdBindIndexBuffer-buffer-08785)VUID-vkCmdBindIndexBuffer-buffer-08785  
  If `buffer` is non-sparse then it **must** be bound completely and contiguously to a single `VkDeviceMemory` object
- [](#VUID-vkCmdBindIndexBuffer-indexType-08786)VUID-vkCmdBindIndexBuffer-indexType-08786  
  `indexType` **must** not be `VK_INDEX_TYPE_NONE_KHR`
- [](#VUID-vkCmdBindIndexBuffer-indexType-08787)VUID-vkCmdBindIndexBuffer-indexType-08787  
  If `indexType` is `VK_INDEX_TYPE_UINT8`, the [`indexTypeUint8`](#features-indexTypeUint8) feature **must** be enabled
- [](#VUID-vkCmdBindIndexBuffer-None-09493)VUID-vkCmdBindIndexBuffer-None-09493  
  If the [`maintenance6`](#features-maintenance6) feature is not enabled, `buffer` **must** not be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html)
- [](#VUID-vkCmdBindIndexBuffer-buffer-09494)VUID-vkCmdBindIndexBuffer-buffer-09494  
  If `buffer` is [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), offset **must** be zero

Valid Usage (Implicit)

- [](#VUID-vkCmdBindIndexBuffer-commandBuffer-parameter)VUID-vkCmdBindIndexBuffer-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdBindIndexBuffer-buffer-parameter)VUID-vkCmdBindIndexBuffer-buffer-parameter  
  If `buffer` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `buffer` **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) handle
- [](#VUID-vkCmdBindIndexBuffer-indexType-parameter)VUID-vkCmdBindIndexBuffer-indexType-parameter  
  `indexType` **must** be a valid [VkIndexType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndexType.html) value
- [](#VUID-vkCmdBindIndexBuffer-commandBuffer-recording)VUID-vkCmdBindIndexBuffer-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdBindIndexBuffer-commandBuffer-cmdpool)VUID-vkCmdBindIndexBuffer-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics operations
- [](#VUID-vkCmdBindIndexBuffer-videocoding)VUID-vkCmdBindIndexBuffer-videocoding  
  This command **must** only be called outside of a video coding scope
- [](#VUID-vkCmdBindIndexBuffer-commonparent)VUID-vkCmdBindIndexBuffer-commonparent  
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

Conditional Rendering

vkCmdBindIndexBuffer is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html), [VkIndexType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndexType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdBindIndexBuffer)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0