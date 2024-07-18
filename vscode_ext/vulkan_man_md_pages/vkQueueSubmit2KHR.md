# vkQueueSubmit2(3) Manual Page

## Name

vkQueueSubmit2 - Submits command buffers to a queue



## <a href="#_c_specification" class="anchor"></a>C Specification

To submit command buffers to a queue, call:

``` c
// Provided by VK_VERSION_1_3
VkResult vkQueueSubmit2(
    VkQueue                                     queue,
    uint32_t                                    submitCount,
    const VkSubmitInfo2*                        pSubmits,
    VkFence                                     fence);
```

or the equivalent command

``` c
// Provided by VK_KHR_synchronization2
VkResult vkQueueSubmit2KHR(
    VkQueue                                     queue,
    uint32_t                                    submitCount,
    const VkSubmitInfo2*                        pSubmits,
    VkFence                                     fence);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `queue` is the queue that the command buffers will be submitted to.

- `submitCount` is the number of elements in the `pSubmits` array.

- `pSubmits` is a pointer to an array of
  [VkSubmitInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubmitInfo2.html) structures, each specifying a
  command buffer submission batch.

- `fence` is an **optional** handle to a fence to be signaled once all
  submitted command buffers have completed execution. If `fence` is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), it defines a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-fences-signaling"
  target="_blank" rel="noopener">fence signal operation</a>.

## <a href="#_description" class="anchor"></a>Description

`vkQueueSubmit2` is a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#devsandqueues-submission"
target="_blank" rel="noopener">queue submission command</a>, with each
batch defined by an element of `pSubmits`.

Semaphore operations submitted with
[vkQueueSubmit2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkQueueSubmit2.html) have additional ordering
constraints compared to other submission commands, with dependencies
involving previous and subsequent queue operations. Information about
these additional constraints can be found in the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-semaphores"
target="_blank" rel="noopener">semaphore</a> section of <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization"
target="_blank" rel="noopener">the synchronization chapter</a>.

If any command buffer submitted to this queue is in the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#commandbuffers-lifecycle"
target="_blank" rel="noopener">executable state</a>, it is moved to the
<a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#commandbuffers-lifecycle"
target="_blank" rel="noopener">pending state</a>. Once execution of all
submissions of a command buffer complete, it moves from the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#commandbuffers-lifecycle"
target="_blank" rel="noopener">pending state</a>, back to the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#commandbuffers-lifecycle"
target="_blank" rel="noopener">executable state</a>. If a command buffer
was recorded with the `VK_COMMAND_BUFFER_USAGE_ONE_TIME_SUBMIT_BIT`
flag, it instead moves back to the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#commandbuffers-lifecycle"
target="_blank" rel="noopener">invalid state</a>.

If `vkQueueSubmit2` fails, it **may** return
`VK_ERROR_OUT_OF_HOST_MEMORY` or `VK_ERROR_OUT_OF_DEVICE_MEMORY`. If it
does, the implementation **must** ensure that the state and contents of
any resources or synchronization primitives referenced by the submitted
command buffers and any semaphores referenced by `pSubmits` is
unaffected by the call or its failure. If `vkQueueSubmit2` fails in such
a way that the implementation is unable to make that guarantee, the
implementation **must** return `VK_ERROR_DEVICE_LOST`. See <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#devsandqueues-lost-device"
target="_blank" rel="noopener">Lost Device</a>.

Valid Usage

- <a href="#VUID-vkQueueSubmit2-fence-04894"
  id="VUID-vkQueueSubmit2-fence-04894"></a>
  VUID-vkQueueSubmit2-fence-04894  
  If `fence` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `fence`
  **must** be unsignaled

- <a href="#VUID-vkQueueSubmit2-fence-04895"
  id="VUID-vkQueueSubmit2-fence-04895"></a>
  VUID-vkQueueSubmit2-fence-04895  
  If `fence` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `fence`
  **must** not be associated with any other queue command that has not
  yet completed execution on that queue

- <a href="#VUID-vkQueueSubmit2-synchronization2-03866"
  id="VUID-vkQueueSubmit2-synchronization2-03866"></a>
  VUID-vkQueueSubmit2-synchronization2-03866  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-synchronization2"
  target="_blank" rel="noopener"><code>synchronization2</code></a>
  feature **must** be enabled

- <a href="#VUID-vkQueueSubmit2-commandBuffer-03867"
  id="VUID-vkQueueSubmit2-commandBuffer-03867"></a>
  VUID-vkQueueSubmit2-commandBuffer-03867  
  If a command recorded into the `commandBuffer` member of any element
  of the `pCommandBufferInfos` member of any element of `pSubmits`
  referenced a [VkEvent](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkEvent.html), that event **must** not be
  referenced by a command that has been submitted to another queue and
  is still in the *pending state*

- <a href="#VUID-vkQueueSubmit2-semaphore-03868"
  id="VUID-vkQueueSubmit2-semaphore-03868"></a>
  VUID-vkQueueSubmit2-semaphore-03868  
  The `semaphore` member of any binary semaphore element of the
  `pSignalSemaphoreInfos` member of any element of `pSubmits` **must**
  be unsignaled when the semaphore signal operation it defines is
  executed on the device

- <a href="#VUID-vkQueueSubmit2-stageMask-03869"
  id="VUID-vkQueueSubmit2-stageMask-03869"></a>
  VUID-vkQueueSubmit2-stageMask-03869  
  The `stageMask` member of any element of the `pSignalSemaphoreInfos`
  member of any element of `pSubmits` **must** only include pipeline
  stages that are supported by the queue family which `queue` belongs to

- <a href="#VUID-vkQueueSubmit2-stageMask-03870"
  id="VUID-vkQueueSubmit2-stageMask-03870"></a>
  VUID-vkQueueSubmit2-stageMask-03870  
  The `stageMask` member of any element of the `pWaitSemaphoreInfos`
  member of any element of `pSubmits` **must** only include pipeline
  stages that are supported by the queue family which `queue` belongs to

- <a href="#VUID-vkQueueSubmit2-semaphore-03871"
  id="VUID-vkQueueSubmit2-semaphore-03871"></a>
  VUID-vkQueueSubmit2-semaphore-03871  
  When a semaphore wait operation for a binary semaphore is executed, as
  defined by the `semaphore` member of any element of the
  `pWaitSemaphoreInfos` member of any element of `pSubmits`, there
  **must** be no other queues waiting on the same semaphore

- <a href="#VUID-vkQueueSubmit2-semaphore-03873"
  id="VUID-vkQueueSubmit2-semaphore-03873"></a>
  VUID-vkQueueSubmit2-semaphore-03873  
  The `semaphore` member of any element of the `pWaitSemaphoreInfos`
  member of any element of `pSubmits` that was created with a
  [VkSemaphoreTypeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreTypeKHR.html) of
  `VK_SEMAPHORE_TYPE_BINARY_KHR` **must** reference a semaphore signal
  operation that has been submitted for execution and any <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-semaphores-signaling"
  target="_blank" rel="noopener">semaphore signal operations</a> on
  which it depends **must** have also been submitted for execution

- <a href="#VUID-vkQueueSubmit2-commandBuffer-03874"
  id="VUID-vkQueueSubmit2-commandBuffer-03874"></a>
  VUID-vkQueueSubmit2-commandBuffer-03874  
  The `commandBuffer` member of any element of the `pCommandBufferInfos`
  member of any element of `pSubmits` **must** be in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#commandbuffers-lifecycle"
  target="_blank" rel="noopener">pending or executable state</a>

- <a href="#VUID-vkQueueSubmit2-commandBuffer-03875"
  id="VUID-vkQueueSubmit2-commandBuffer-03875"></a>
  VUID-vkQueueSubmit2-commandBuffer-03875  
  If a command recorded into the `commandBuffer` member of any element
  of the `pCommandBufferInfos` member of any element of `pSubmits` was
  not recorded with the `VK_COMMAND_BUFFER_USAGE_SIMULTANEOUS_USE_BIT`,
  it **must** not be in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#commandbuffers-lifecycle"
  target="_blank" rel="noopener">pending state</a>

- <a href="#VUID-vkQueueSubmit2-commandBuffer-03876"
  id="VUID-vkQueueSubmit2-commandBuffer-03876"></a>
  VUID-vkQueueSubmit2-commandBuffer-03876  
  Any <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#commandbuffers-secondary"
  target="_blank" rel="noopener">secondary command buffers recorded</a>
  into the `commandBuffer` member of any element of the
  `pCommandBufferInfos` member of any element of `pSubmits` **must** be
  in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#commandbuffers-lifecycle"
  target="_blank" rel="noopener">pending or executable state</a>

- <a href="#VUID-vkQueueSubmit2-commandBuffer-03877"
  id="VUID-vkQueueSubmit2-commandBuffer-03877"></a>
  VUID-vkQueueSubmit2-commandBuffer-03877  
  If any <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#commandbuffers-secondary"
  target="_blank" rel="noopener">secondary command buffers recorded</a>
  into the `commandBuffer` member of any element of the
  `pCommandBufferInfos` member of any element of `pSubmits` was not
  recorded with the `VK_COMMAND_BUFFER_USAGE_SIMULTANEOUS_USE_BIT`, it
  **must** not be in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#commandbuffers-lifecycle"
  target="_blank" rel="noopener">pending state</a>

- <a href="#VUID-vkQueueSubmit2-commandBuffer-03878"
  id="VUID-vkQueueSubmit2-commandBuffer-03878"></a>
  VUID-vkQueueSubmit2-commandBuffer-03878  
  The `commandBuffer` member of any element of the `pCommandBufferInfos`
  member of any element of `pSubmits` **must** have been allocated from
  a `VkCommandPool` that was created for the same queue family `queue`
  belongs to

- <a href="#VUID-vkQueueSubmit2-commandBuffer-03879"
  id="VUID-vkQueueSubmit2-commandBuffer-03879"></a>
  VUID-vkQueueSubmit2-commandBuffer-03879  
  If a command recorded into the `commandBuffer` member of any element
  of the `pCommandBufferInfos` member of any element of `pSubmits`
  includes a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-queue-transfers-acquire"
  target="_blank" rel="noopener">Queue Family Ownership Transfer Acquire
  Operation</a>, there **must** exist a previously submitted <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-queue-transfers-release"
  target="_blank" rel="noopener">Queue Family Ownership Transfer Release
  Operation</a> on a queue in the queue family identified by the acquire
  operation, with parameters matching the acquire operation as defined
  in the definition of such <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-queue-transfers-acquire"
  target="_blank" rel="noopener">acquire operations</a>, and which
  happens before the acquire operation

- <a href="#VUID-vkQueueSubmit2-commandBuffer-03880"
  id="VUID-vkQueueSubmit2-commandBuffer-03880"></a>
  VUID-vkQueueSubmit2-commandBuffer-03880  
  If a command recorded into the `commandBuffer` member of any element
  of the `pCommandBufferInfos` member of any element of `pSubmits` was a
  [vkCmdBeginQuery](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginQuery.html) whose `queryPool` was created
  with a `queryType` of `VK_QUERY_TYPE_PERFORMANCE_QUERY_KHR`, the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#profiling-lock"
  target="_blank" rel="noopener">profiling lock</a> **must** have been
  held continuously on the `VkDevice` that `queue` was retrieved from,
  throughout recording of those command buffers

- <a href="#VUID-vkQueueSubmit2-queue-06447"
  id="VUID-vkQueueSubmit2-queue-06447"></a>
  VUID-vkQueueSubmit2-queue-06447  
  If `queue` was not created with
  `VK_DEVICE_QUEUE_CREATE_PROTECTED_BIT`, the `flags` member of any
  element of `pSubmits` **must** not include
  `VK_SUBMIT_PROTECTED_BIT_KHR`

Valid Usage (Implicit)

- <a href="#VUID-vkQueueSubmit2-queue-parameter"
  id="VUID-vkQueueSubmit2-queue-parameter"></a>
  VUID-vkQueueSubmit2-queue-parameter  
  `queue` **must** be a valid [VkQueue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueue.html) handle

- <a href="#VUID-vkQueueSubmit2-pSubmits-parameter"
  id="VUID-vkQueueSubmit2-pSubmits-parameter"></a>
  VUID-vkQueueSubmit2-pSubmits-parameter  
  If `submitCount` is not `0`, `pSubmits` **must** be a valid pointer to
  an array of `submitCount` valid [VkSubmitInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubmitInfo2.html)
  structures

- <a href="#VUID-vkQueueSubmit2-fence-parameter"
  id="VUID-vkQueueSubmit2-fence-parameter"></a>
  VUID-vkQueueSubmit2-fence-parameter  
  If `fence` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `fence`
  **must** be a valid [VkFence](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFence.html) handle

- <a href="#VUID-vkQueueSubmit2-commonparent"
  id="VUID-vkQueueSubmit2-commonparent"></a>
  VUID-vkQueueSubmit2-commonparent  
  Both of `fence`, and `queue` that are valid handles of non-ignored
  parameters **must** have been created, allocated, or retrieved from
  the same [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

Host Synchronization

- Host access to `queue` **must** be externally synchronized

- Host access to `fence` **must** be externally synchronized

Command Properties

| [Command Buffer Levels](#VkCommandBufferLevel) | [Render Pass Scope](#vkCmdBeginRenderPass) | [Video Coding Scope](#vkCmdBeginVideoCodingKHR) | [Supported Queue Types](#VkQueueFlagBits) | [Command Type](#fundamentals-queueoperation-command-types) |
|----|----|----|----|----|
| \- | \- | \- | Any | \- |

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

- `VK_ERROR_DEVICE_LOST`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_synchronization2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_synchronization2.html),
[VK_VERSION_1_3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_3.html), [VkFence](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFence.html),
[VkQueue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueue.html), [VkSubmitInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubmitInfo2.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkQueueSubmit2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
