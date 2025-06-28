# vkQueueSubmit2(3) Manual Page

## Name

vkQueueSubmit2 - Submits command buffers to a queue



## [](#_c_specification)C Specification

To submit command buffers to a queue, call:

```c++
// Provided by VK_VERSION_1_3
VkResult vkQueueSubmit2(
    VkQueue                                     queue,
    uint32_t                                    submitCount,
    const VkSubmitInfo2*                        pSubmits,
    VkFence                                     fence);
```

or the equivalent command

```c++
// Provided by VK_KHR_synchronization2
VkResult vkQueueSubmit2KHR(
    VkQueue                                     queue,
    uint32_t                                    submitCount,
    const VkSubmitInfo2*                        pSubmits,
    VkFence                                     fence);
```

## [](#_parameters)Parameters

- `queue` is the queue that the command buffers will be submitted to.
- `submitCount` is the number of elements in the `pSubmits` array.
- `pSubmits` is a pointer to an array of [VkSubmitInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubmitInfo2.html) structures, each specifying a command buffer submission batch. Command buffers and semaphores specified in this array **may** be accessed at any point until the [queue operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#devsandqueues-submission) they define complete execution on the device.
- `fence` is an **optional** handle to a fence to be signaled once all submitted command buffers have completed execution. If `fence` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), it defines a [fence signal operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-fences-signaling). If it is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `fence` **may** be accessed at any point until this command completes on the device.

## [](#_description)Description

`vkQueueSubmit2` is a [queue submission command](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#devsandqueues-submission), with each batch defined by an element of `pSubmits`.

Semaphore operations submitted with [vkQueueSubmit2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkQueueSubmit2.html) have additional ordering constraints compared to other submission commands, with dependencies involving previous and subsequent queue operations. Information about these additional constraints can be found in the [semaphore](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-semaphores) section of [the synchronization chapter](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization).

If any command buffer submitted to this queue is in the [executable state](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#commandbuffers-lifecycle), it is moved to the [pending state](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#commandbuffers-lifecycle). Once execution of all submissions of a command buffer complete, it moves from the [pending state](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#commandbuffers-lifecycle), back to the [executable state](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#commandbuffers-lifecycle). If a command buffer was recorded with the `VK_COMMAND_BUFFER_USAGE_ONE_TIME_SUBMIT_BIT` flag, it instead moves back to the [invalid state](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#commandbuffers-lifecycle).

If `vkQueueSubmit2` fails, it **may** return `VK_ERROR_OUT_OF_HOST_MEMORY` or `VK_ERROR_OUT_OF_DEVICE_MEMORY`. If it does, the implementation **must** ensure that the state and contents of any resources or synchronization primitives referenced by the submitted command buffers and any semaphores referenced by `pSubmits` is unaffected by the call or its failure. If `vkQueueSubmit2` fails in such a way that the implementation is unable to make that guarantee, the implementation **must** return `VK_ERROR_DEVICE_LOST`. See [Lost Device](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#devsandqueues-lost-device).

Valid Usage

- [](#VUID-vkQueueSubmit2-fence-04894)VUID-vkQueueSubmit2-fence-04894  
  If `fence` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `fence` **must** be unsignaled
- [](#VUID-vkQueueSubmit2-fence-04895)VUID-vkQueueSubmit2-fence-04895  
  If `fence` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `fence` **must** not be associated with any other queue command that has not yet completed execution on that queue
- [](#VUID-vkQueueSubmit2-synchronization2-03866)VUID-vkQueueSubmit2-synchronization2-03866  
  The [`synchronization2`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-synchronization2) feature **must** be enabled
- [](#VUID-vkQueueSubmit2-commandBuffer-03867)VUID-vkQueueSubmit2-commandBuffer-03867  
  If a command recorded into the `commandBuffer` member of any element of the `pCommandBufferInfos` member of any element of `pSubmits` referenced a [VkEvent](https://registry.khronos.org/vulkan/specs/latest/man/html/VkEvent.html), that event **must** not be referenced by a command that has been submitted to another queue and is still in the *pending state*
- [](#VUID-vkQueueSubmit2-semaphore-03868)VUID-vkQueueSubmit2-semaphore-03868  
  The `semaphore` member of any binary semaphore element of the `pSignalSemaphoreInfos` member of any element of `pSubmits` **must** be unsignaled when the semaphore signal operation it defines is executed on the device
- [](#VUID-vkQueueSubmit2-stageMask-03869)VUID-vkQueueSubmit2-stageMask-03869  
  The `stageMask` member of any element of the `pSignalSemaphoreInfos` member of any element of `pSubmits` **must** only include pipeline stages that are supported by the queue family which `queue` belongs to
- [](#VUID-vkQueueSubmit2-stageMask-03870)VUID-vkQueueSubmit2-stageMask-03870  
  The `stageMask` member of any element of the `pWaitSemaphoreInfos` member of any element of `pSubmits` **must** only include pipeline stages that are supported by the queue family which `queue` belongs to
- [](#VUID-vkQueueSubmit2-semaphore-03871)VUID-vkQueueSubmit2-semaphore-03871  
  When a semaphore wait operation for a binary semaphore is executed, as defined by the `semaphore` member of any element of the `pWaitSemaphoreInfos` member of any element of `pSubmits`, there **must** be no other queues waiting on the same semaphore
- [](#VUID-vkQueueSubmit2-semaphore-03873)VUID-vkQueueSubmit2-semaphore-03873  
  The `semaphore` member of any element of the `pWaitSemaphoreInfos` member of any element of `pSubmits` that was created with a [VkSemaphoreType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreType.html) of `VK_SEMAPHORE_TYPE_BINARY` **must** reference a semaphore signal operation that has been submitted for execution and any [semaphore signal operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-semaphores-signaling) on which it depends **must** have also been submitted for execution
- [](#VUID-vkQueueSubmit2-commandBuffer-03874)VUID-vkQueueSubmit2-commandBuffer-03874  
  The `commandBuffer` member of any element of the `pCommandBufferInfos` member of any element of `pSubmits` **must** be in the [pending or executable state](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#commandbuffers-lifecycle)
- [](#VUID-vkQueueSubmit2-commandBuffer-03875)VUID-vkQueueSubmit2-commandBuffer-03875  
  If a command recorded into the `commandBuffer` member of any element of the `pCommandBufferInfos` member of any element of `pSubmits` was not recorded with the `VK_COMMAND_BUFFER_USAGE_SIMULTANEOUS_USE_BIT`, it **must** not be in the [pending state](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#commandbuffers-lifecycle)
- [](#VUID-vkQueueSubmit2-commandBuffer-03876)VUID-vkQueueSubmit2-commandBuffer-03876  
  Any [secondary command buffers recorded](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#commandbuffers-secondary) into the `commandBuffer` member of any element of the `pCommandBufferInfos` member of any element of `pSubmits` **must** be in the [pending or executable state](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#commandbuffers-lifecycle)
- [](#VUID-vkQueueSubmit2-commandBuffer-03877)VUID-vkQueueSubmit2-commandBuffer-03877  
  If any [secondary command buffers recorded](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#commandbuffers-secondary) into the `commandBuffer` member of any element of the `pCommandBufferInfos` member of any element of `pSubmits` was not recorded with the `VK_COMMAND_BUFFER_USAGE_SIMULTANEOUS_USE_BIT`, it **must** not be in the [pending state](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#commandbuffers-lifecycle)
- [](#VUID-vkQueueSubmit2-commandBuffer-03878)VUID-vkQueueSubmit2-commandBuffer-03878  
  The `commandBuffer` member of any element of the `pCommandBufferInfos` member of any element of `pSubmits` **must** have been allocated from a `VkCommandPool` that was created for the same queue family `queue` belongs to
- [](#VUID-vkQueueSubmit2-commandBuffer-03879)VUID-vkQueueSubmit2-commandBuffer-03879  
  If a command recorded into the `commandBuffer` member of any element of the `pCommandBufferInfos` member of any element of `pSubmits` includes a [Queue Family Ownership Transfer Acquire Operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-queue-transfers-acquire), there **must** exist a previously submitted [Queue Family Ownership Transfer Release Operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-queue-transfers-release) on a queue in the queue family identified by the acquire operation, with parameters matching the acquire operation as defined in the definition of such [acquire operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-queue-transfers-acquire), and which happens before the acquire operation
- [](#VUID-vkQueueSubmit2-commandBuffer-03880)VUID-vkQueueSubmit2-commandBuffer-03880  
  If a command recorded into the `commandBuffer` member of any element of the `pCommandBufferInfos` member of any element of `pSubmits` was a [vkCmdBeginQuery](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBeginQuery.html) whose `queryPool` was created with a `queryType` of `VK_QUERY_TYPE_PERFORMANCE_QUERY_KHR`, the [profiling lock](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#profiling-lock) **must** have been held continuously on the `VkDevice` that `queue` was retrieved from, throughout recording of those command buffers
- [](#VUID-vkQueueSubmit2-queue-06447)VUID-vkQueueSubmit2-queue-06447  
  If `queue` was not created with `VK_DEVICE_QUEUE_CREATE_PROTECTED_BIT`, the `flags` member of any element of `pSubmits` **must** not include `VK_SUBMIT_PROTECTED_BIT_KHR`

Valid Usage (Implicit)

- [](#VUID-vkQueueSubmit2-queue-parameter)VUID-vkQueueSubmit2-queue-parameter  
  `queue` **must** be a valid [VkQueue](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueue.html) handle
- [](#VUID-vkQueueSubmit2-pSubmits-parameter)VUID-vkQueueSubmit2-pSubmits-parameter  
  If `submitCount` is not `0`, `pSubmits` **must** be a valid pointer to an array of `submitCount` valid [VkSubmitInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubmitInfo2.html) structures
- [](#VUID-vkQueueSubmit2-fence-parameter)VUID-vkQueueSubmit2-fence-parameter  
  If `fence` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `fence` **must** be a valid [VkFence](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFence.html) handle
- [](#VUID-vkQueueSubmit2-commonparent)VUID-vkQueueSubmit2-commonparent  
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

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_DEVICE_LOST`

## [](#_see_also)See Also

[VK\_KHR\_synchronization2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_synchronization2.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkFence](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFence.html), [VkQueue](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueue.html), [VkSubmitInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubmitInfo2.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkQueueSubmit2)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0