# vkQueueBindSparse(3) Manual Page

## Name

vkQueueBindSparse - Bind device memory to a sparse resource object



## [](#_c_specification)C Specification

To submit sparse binding operations to a queue, call:

```c++
// Provided by VK_VERSION_1_0
VkResult vkQueueBindSparse(
    VkQueue                                     queue,
    uint32_t                                    bindInfoCount,
    const VkBindSparseInfo*                     pBindInfo,
    VkFence                                     fence);
```

## [](#_parameters)Parameters

- `queue` is the queue that the sparse binding operations will be submitted to.
- `bindInfoCount` is the number of elements in the `pBindInfo` array.
- `pBindInfo` is a pointer to an array of [VkBindSparseInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindSparseInfo.html) structures, each specifying a sparse binding submission batch.
- `fence` is an **optional** handle to a fence to be signaled. If `fence` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), it defines a [fence signal operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-fences-signaling).

## [](#_description)Description

`vkQueueBindSparse` is a [queue submission command](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#devsandqueues-submission), with each batch defined by an element of `pBindInfo` as a [VkBindSparseInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindSparseInfo.html) structure. Batches begin execution in the order they appear in `pBindInfo`, but **may** complete out of order.

Within a batch, a given range of a resource **must** not be bound more than once. Across batches, if a range is to be bound to one allocation and offset and then to another allocation and offset, then the application **must** guarantee (usually using semaphores) that the binding operations are executed in the correct order, as well as to order binding operations against the execution of command buffer submissions.

As no operation to [vkQueueBindSparse](https://registry.khronos.org/vulkan/specs/latest/man/html/vkQueueBindSparse.html) causes any pipeline stage to access memory, synchronization primitives used in this command effectively only define execution dependencies.

Additional information about fence and semaphore operation is described in [the synchronization chapter](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization).

Valid Usage

- [](#VUID-vkQueueBindSparse-fence-01113)VUID-vkQueueBindSparse-fence-01113  
  If `fence` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `fence` **must** be unsignaled
- [](#VUID-vkQueueBindSparse-fence-01114)VUID-vkQueueBindSparse-fence-01114  
  If `fence` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `fence` **must** not be associated with any other queue command that has not yet completed execution on that queue
- [](#VUID-vkQueueBindSparse-pSignalSemaphores-01115)VUID-vkQueueBindSparse-pSignalSemaphores-01115  
  Each element of the `pSignalSemaphores` member of each element of `pBindInfo` **must** be unsignaled when the semaphore signal operation it defines is executed on the device
- [](#VUID-vkQueueBindSparse-pWaitSemaphores-01116)VUID-vkQueueBindSparse-pWaitSemaphores-01116  
  When a semaphore wait operation referring to a binary semaphore defined by any element of the `pWaitSemaphores` member of any element of `pBindInfo` executes on `queue`, there **must** be no other queues waiting on the same semaphore
- [](#VUID-vkQueueBindSparse-pWaitSemaphores-03245)VUID-vkQueueBindSparse-pWaitSemaphores-03245  
  All elements of the `pWaitSemaphores` member of all elements of `pBindInfo` referring to a semaphore created with a [VkSemaphoreType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreType.html) of `VK_SEMAPHORE_TYPE_BINARY` **must** reference a semaphore signal operation that has been submitted for execution and any [semaphore signal operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-semaphores-signaling) on which it depends **must** have also been submitted for execution

Valid Usage (Implicit)

- [](#VUID-vkQueueBindSparse-queue-parameter)VUID-vkQueueBindSparse-queue-parameter  
  `queue` **must** be a valid [VkQueue](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueue.html) handle
- [](#VUID-vkQueueBindSparse-pBindInfo-parameter)VUID-vkQueueBindSparse-pBindInfo-parameter  
  If `bindInfoCount` is not `0`, `pBindInfo` **must** be a valid pointer to an array of `bindInfoCount` valid [VkBindSparseInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindSparseInfo.html) structures
- [](#VUID-vkQueueBindSparse-fence-parameter)VUID-vkQueueBindSparse-fence-parameter  
  If `fence` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `fence` **must** be a valid [VkFence](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFence.html) handle
- [](#VUID-vkQueueBindSparse-queuetype)VUID-vkQueueBindSparse-queuetype  
  The `queue` **must** support VK\_QUEUE\_SPARSE\_BINDING\_BIT operations
- [](#VUID-vkQueueBindSparse-commonparent)VUID-vkQueueBindSparse-commonparent  
  Both of `fence`, and `queue` that are valid handles of non-ignored parameters **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

Host Synchronization

- Host access to `queue` **must** be externally synchronized
- Host access to `fence` **must** be externally synchronized

Command Properties

     [Command Buffer Levels](#VkCommandBufferLevel) [Render Pass Scope](#vkCmdBeginRenderPass) [Video Coding Scope](#vkCmdBeginVideoCodingKHR) [Supported Queue Types](#VkQueueFlagBits) [Command Type](#fundamentals-queueoperation-command-types)

\-

\-

\-

VK\_QUEUE\_SPARSE\_BINDING\_BIT

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

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkBindSparseInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindSparseInfo.html), [VkFence](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFence.html), [VkQueue](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueue.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkQueueBindSparse).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0