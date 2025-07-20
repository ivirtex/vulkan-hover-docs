# vkQueueSubmit(3) Manual Page

## Name

vkQueueSubmit - Submits a sequence of semaphores or command buffers to a queue



## [](#_c_specification)C Specification

To submit command buffers to a queue, call:

```c++
// Provided by VK_VERSION_1_0
VkResult vkQueueSubmit(
    VkQueue                                     queue,
    uint32_t                                    submitCount,
    const VkSubmitInfo*                         pSubmits,
    VkFence                                     fence);
```

## [](#_parameters)Parameters

- `queue` is the queue that the command buffers will be submitted to.
- `submitCount` is the number of elements in the `pSubmits` array.
- `pSubmits` is a pointer to an array of [VkSubmitInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubmitInfo.html) structures, each specifying a command buffer submission batch. Command buffers and semaphores specified in this array **may** be accessed at any point until the [queue operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#devsandqueues-submission) they define complete execution on the device.
- `fence` is an **optional** handle to a fence to be signaled once all submitted command buffers have completed execution. If `fence` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), it defines a [fence signal operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-fences-signaling). If it is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `fence` **may** be accessed at any point until this command completes on the device.

## [](#_description)Description

`vkQueueSubmit` is a [queue submission command](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#devsandqueues-submission), with each batch defined by an element of `pSubmits`. Batches begin execution in the order they appear in `pSubmits`, but **may** complete out of order.

Fence and semaphore operations submitted with [vkQueueSubmit](https://registry.khronos.org/vulkan/specs/latest/man/html/vkQueueSubmit.html) have additional ordering constraints compared to other submission commands, with dependencies involving previous and subsequent queue operations. Information about these additional constraints can be found in the [semaphore](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-semaphores) and [fence](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-fences) sections of [the synchronization chapter](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization).

Details on the interaction of `pWaitDstStageMask` with synchronization are described in the [semaphore wait operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-semaphores-waiting) section of [the synchronization chapter](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization).

The order that batches appear in `pSubmits` is used to determine [submission order](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-submission-order), and thus all the [implicit ordering guarantees](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-implicit) that respect it. Other than these implicit ordering guarantees and any [explicit synchronization primitives](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization), these batches **may** overlap or otherwise execute out of order.

If any command buffer submitted to this queue is in the [executable state](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#commandbuffers-lifecycle), it is moved to the [pending state](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#commandbuffers-lifecycle). Once execution of all submissions of a command buffer complete, it moves from the [pending state](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#commandbuffers-lifecycle), back to the [executable state](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#commandbuffers-lifecycle). If a command buffer was recorded with the `VK_COMMAND_BUFFER_USAGE_ONE_TIME_SUBMIT_BIT` flag, it instead moves to the [invalid state](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#commandbuffers-lifecycle).

If `vkQueueSubmit` fails, it **may** return `VK_ERROR_OUT_OF_HOST_MEMORY` or `VK_ERROR_OUT_OF_DEVICE_MEMORY`. If it does, the implementation **must** ensure that the state and contents of any resources or synchronization primitives referenced by the submitted command buffers and any semaphores referenced by `pSubmits` is unaffected by the call or its failure. If `vkQueueSubmit` fails in such a way that the implementation is unable to make that guarantee, the implementation **must** return `VK_ERROR_DEVICE_LOST`. See [Lost Device](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#devsandqueues-lost-device).

Valid Usage

- [](#VUID-vkQueueSubmit-fence-00063)VUID-vkQueueSubmit-fence-00063  
  If `fence` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `fence` **must** be unsignaled
- [](#VUID-vkQueueSubmit-fence-00064)VUID-vkQueueSubmit-fence-00064  
  If `fence` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `fence` **must** not be associated with any other queue command that has not yet completed execution on that queue
- [](#VUID-vkQueueSubmit-pCommandBuffers-00065)VUID-vkQueueSubmit-pCommandBuffers-00065  
  Any calls to [vkCmdSetEvent](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetEvent.html), [vkCmdResetEvent](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdResetEvent.html) or [vkCmdWaitEvents](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdWaitEvents.html) that have been recorded into any of the command buffer elements of the `pCommandBuffers` member of any element of `pSubmits`, **must** not reference any [VkEvent](https://registry.khronos.org/vulkan/specs/latest/man/html/VkEvent.html) that is referenced by any of those commands in a command buffer that has been submitted to another queue and is still in the *pending state*
- [](#VUID-vkQueueSubmit-pWaitDstStageMask-00066)VUID-vkQueueSubmit-pWaitDstStageMask-00066  
  Any stage flag included in any element of the `pWaitDstStageMask` member of any element of `pSubmits` **must** be a pipeline stage supported by one of the capabilities of `queue`, as specified in the [table of supported pipeline stages](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-pipeline-stages-supported)
- [](#VUID-vkQueueSubmit-pSignalSemaphores-00067)VUID-vkQueueSubmit-pSignalSemaphores-00067  
  Each binary semaphore element of the `pSignalSemaphores` member of any element of `pSubmits` **must** be unsignaled when the semaphore signal operation it defines is executed on the device
- [](#VUID-vkQueueSubmit-pWaitSemaphores-00068)VUID-vkQueueSubmit-pWaitSemaphores-00068  
  When a semaphore wait operation referring to a binary semaphore defined by any element of the `pWaitSemaphores` member of any element of `pSubmits` executes on `queue`, there **must** be no other queues waiting on the same semaphore
- [](#VUID-vkQueueSubmit-pWaitSemaphores-03238)VUID-vkQueueSubmit-pWaitSemaphores-03238  
  All elements of the `pWaitSemaphores` member of all elements of `pSubmits` created with a [VkSemaphoreType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreType.html) of `VK_SEMAPHORE_TYPE_BINARY` **must** reference a semaphore signal operation that has been submitted for execution and any [semaphore signal operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-semaphores-signaling) on which it depends **must** have also been submitted for execution
- [](#VUID-vkQueueSubmit-pCommandBuffers-00070)VUID-vkQueueSubmit-pCommandBuffers-00070  
  Each element of the `pCommandBuffers` member of each element of `pSubmits` **must** be in the [pending or executable state](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#commandbuffers-lifecycle)
- [](#VUID-vkQueueSubmit-pCommandBuffers-00071)VUID-vkQueueSubmit-pCommandBuffers-00071  
  If any element of the `pCommandBuffers` member of any element of `pSubmits` was not recorded with the `VK_COMMAND_BUFFER_USAGE_SIMULTANEOUS_USE_BIT`, it **must** not be in the [pending state](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#commandbuffers-lifecycle)
- [](#VUID-vkQueueSubmit-pCommandBuffers-00072)VUID-vkQueueSubmit-pCommandBuffers-00072  
  Any [secondary command buffers recorded](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#commandbuffers-secondary) into any element of the `pCommandBuffers` member of any element of `pSubmits` **must** be in the [pending or executable state](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#commandbuffers-lifecycle)
- [](#VUID-vkQueueSubmit-pCommandBuffers-00073)VUID-vkQueueSubmit-pCommandBuffers-00073  
  If any [secondary command buffers recorded](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#commandbuffers-secondary) into any element of the `pCommandBuffers` member of any element of `pSubmits` was not recorded with the `VK_COMMAND_BUFFER_USAGE_SIMULTANEOUS_USE_BIT`, it **must** not be in the [pending state](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#commandbuffers-lifecycle)
- [](#VUID-vkQueueSubmit-pCommandBuffers-00074)VUID-vkQueueSubmit-pCommandBuffers-00074  
  Each element of the `pCommandBuffers` member of each element of `pSubmits` **must** have been allocated from a `VkCommandPool` that was created for the same queue family `queue` belongs to
- [](#VUID-vkQueueSubmit-pSubmits-02207)VUID-vkQueueSubmit-pSubmits-02207  
  If any element of `pSubmits->pCommandBuffers` includes a [Queue Family Ownership Transfer Acquire Operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-queue-transfers-acquire), there **must** exist a previously submitted [Queue Family Ownership Transfer Release Operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-queue-transfers-release) on a queue in the queue family identified by the acquire operation, with parameters matching the acquire operation as defined in the definition of such [acquire operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-queue-transfers-acquire), and which happens-before the acquire operation
- [](#VUID-vkQueueSubmit-pCommandBuffers-03220)VUID-vkQueueSubmit-pCommandBuffers-03220  
  If a command recorded into any element of `pCommandBuffers` was a [vkCmdBeginQuery](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBeginQuery.html) whose `queryPool` was created with a `queryType` of `VK_QUERY_TYPE_PERFORMANCE_QUERY_KHR`, the [profiling lock](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#profiling-lock) **must** have been held continuously on the `VkDevice` that `queue` was retrieved from, throughout recording of those command buffers
- [](#VUID-vkQueueSubmit-pSubmits-02808)VUID-vkQueueSubmit-pSubmits-02808  
  Any resource created with `VK_SHARING_MODE_EXCLUSIVE` that is read by an operation specified by `pSubmits` **must** not be owned by any queue family other than the one which `queue` belongs to, at the time it is executed
- [](#VUID-vkQueueSubmit-pSubmits-04626)VUID-vkQueueSubmit-pSubmits-04626  
  Any resource created with `VK_SHARING_MODE_CONCURRENT` that is accessed by an operation specified by `pSubmits` **must** have included the queue family of `queue` at resource creation time
- [](#VUID-vkQueueSubmit-queue-06448)VUID-vkQueueSubmit-queue-06448  
  If `queue` was not created with `VK_DEVICE_QUEUE_CREATE_PROTECTED_BIT`, there **must** be no element of `pSubmits` that includes a [VkProtectedSubmitInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkProtectedSubmitInfo.html) structure in its `pNext` chain with `protectedSubmit` equal to `VK_TRUE`

Valid Usage (Implicit)

- [](#VUID-vkQueueSubmit-queue-parameter)VUID-vkQueueSubmit-queue-parameter  
  `queue` **must** be a valid [VkQueue](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueue.html) handle
- [](#VUID-vkQueueSubmit-pSubmits-parameter)VUID-vkQueueSubmit-pSubmits-parameter  
  If `submitCount` is not `0`, `pSubmits` **must** be a valid pointer to an array of `submitCount` valid [VkSubmitInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubmitInfo.html) structures
- [](#VUID-vkQueueSubmit-fence-parameter)VUID-vkQueueSubmit-fence-parameter  
  If `fence` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `fence` **must** be a valid [VkFence](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFence.html) handle
- [](#VUID-vkQueueSubmit-commonparent)VUID-vkQueueSubmit-commonparent  
  Both of `fence`, and `queue` that are valid handles of non-ignored parameters **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

Host Synchronization

- Host access to `queue` **must** be externally synchronized
- Host access to `fence` **must** be externally synchronized

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

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkFence](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFence.html), [VkQueue](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueue.html), [VkSubmitInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubmitInfo.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkQueueSubmit)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0