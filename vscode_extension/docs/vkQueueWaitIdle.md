# vkQueueWaitIdle(3) Manual Page

## Name

vkQueueWaitIdle - Wait for a queue to become idle



## [](#_c_specification)C Specification

To wait on the host for the completion of outstanding queue operations for a given queue, call:

```c++
// Provided by VK_VERSION_1_0
VkResult vkQueueWaitIdle(
    VkQueue                                     queue);
```

## [](#_parameters)Parameters

- `queue` is the queue on which to wait.

## [](#_description)Description

`vkQueueWaitIdle` is equivalent to having submitted a valid fence to every previously executed [queue submission command](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#devsandqueues-submission) that accepts a fence, then waiting for all of those fences to signal using [vkWaitForFences](https://registry.khronos.org/vulkan/specs/latest/man/html/vkWaitForFences.html) with an infinite timeout and `waitAll` set to `VK_TRUE`.

Valid Usage (Implicit)

- [](#VUID-vkQueueWaitIdle-queue-parameter)VUID-vkQueueWaitIdle-queue-parameter  
  `queue` **must** be a valid [VkQueue](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueue.html) handle

Host Synchronization

- Host access to `queue` **must** be externally synchronized

Command Properties

     [Command Buffer Levels](#VkCommandBufferLevel) [Render Pass Scope](#vkCmdBeginRenderPass) [Video Coding Scope](#vkCmdBeginVideoCodingKHR) [Supported Queue Types](#VkQueueFlagBits) [Command Type](#fundamentals-queueoperation-command-types)

\-

\-

\-

Any

\-

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_DEVICE_LOST`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkQueue](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueue.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkQueueWaitIdle)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0