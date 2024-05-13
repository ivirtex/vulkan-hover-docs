# vkQueueSubmit(3) Manual Page

## Name

vkQueueSubmit - Submits a sequence of semaphores or command buffers to a
queue



## <a href="#_c_specification" class="anchor"></a>C Specification

To submit command buffers to a queue, call:

``` c
// Provided by VK_VERSION_1_0
VkResult vkQueueSubmit(
    VkQueue                                     queue,
    uint32_t                                    submitCount,
    const VkSubmitInfo*                         pSubmits,
    VkFence                                     fence);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `queue` is the queue that the command buffers will be submitted to.

- `submitCount` is the number of elements in the `pSubmits` array.

- `pSubmits` is a pointer to an array of
  [VkSubmitInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubmitInfo.html) structures, each specifying a
  command buffer submission batch.

- `fence` is an **optional** handle to a fence to be signaled once all
  submitted command buffers have completed execution. If `fence` is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), it defines a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-fences-signaling"
  target="_blank" rel="noopener">fence signal operation</a>.

## <a href="#_description" class="anchor"></a>Description

`vkQueueSubmit` is a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#devsandqueues-submission"
target="_blank" rel="noopener">queue submission command</a>, with each
batch defined by an element of `pSubmits`. Batches begin execution in
the order they appear in `pSubmits`, but **may** complete out of order.

Fence and semaphore operations submitted with
[vkQueueSubmit](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkQueueSubmit.html) have additional ordering constraints
compared to other submission commands, with dependencies involving
previous and subsequent queue operations. Information about these
additional constraints can be found in the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-semaphores"
target="_blank" rel="noopener">semaphore</a> and <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-fences"
target="_blank" rel="noopener">fence</a> sections of <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization"
target="_blank" rel="noopener">the synchronization chapter</a>.

Details on the interaction of `pWaitDstStageMask` with synchronization
are described in the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-semaphores-waiting"
target="_blank" rel="noopener">semaphore wait operation</a> section of
<a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization"
target="_blank" rel="noopener">the synchronization chapter</a>.

The order that batches appear in `pSubmits` is used to determine <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-submission-order"
target="_blank" rel="noopener">submission order</a>, and thus all the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-implicit"
target="_blank" rel="noopener">implicit ordering guarantees</a> that
respect it. Other than these implicit ordering guarantees and any <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization"
target="_blank" rel="noopener">explicit synchronization primitives</a>,
these batches **may** overlap or otherwise execute out of order.

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
flag, it instead moves to the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#commandbuffers-lifecycle"
target="_blank" rel="noopener">invalid state</a>.

If `vkQueueSubmit` fails, it **may** return
`VK_ERROR_OUT_OF_HOST_MEMORY` or `VK_ERROR_OUT_OF_DEVICE_MEMORY`. If it
does, the implementation **must** ensure that the state and contents of
any resources or synchronization primitives referenced by the submitted
command buffers and any semaphores referenced by `pSubmits` is
unaffected by the call or its failure. If `vkQueueSubmit` fails in such
a way that the implementation is unable to make that guarantee, the
implementation **must** return `VK_ERROR_DEVICE_LOST`. See <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#devsandqueues-lost-device"
target="_blank" rel="noopener">Lost Device</a>.

Valid Usage

- <a href="#VUID-vkQueueSubmit-fence-00063"
  id="VUID-vkQueueSubmit-fence-00063"></a>
  VUID-vkQueueSubmit-fence-00063  
  If `fence` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `fence`
  **must** be unsignaled

- <a href="#VUID-vkQueueSubmit-fence-00064"
  id="VUID-vkQueueSubmit-fence-00064"></a>
  VUID-vkQueueSubmit-fence-00064  
  If `fence` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `fence`
  **must** not be associated with any other queue command that has not
  yet completed execution on that queue

- <a href="#VUID-vkQueueSubmit-pCommandBuffers-00065"
  id="VUID-vkQueueSubmit-pCommandBuffers-00065"></a>
  VUID-vkQueueSubmit-pCommandBuffers-00065  
  Any calls to [vkCmdSetEvent](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetEvent.html),
  [vkCmdResetEvent](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdResetEvent.html) or
  [vkCmdWaitEvents](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdWaitEvents.html) that have been recorded into
  any of the command buffer elements of the `pCommandBuffers` member of
  any element of `pSubmits`, **must** not reference any
  [VkEvent](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkEvent.html) that is referenced by any of those commands in
  a command buffer that has been submitted to another queue and is still
  in the *pending state*

- <a href="#VUID-vkQueueSubmit-pWaitDstStageMask-00066"
  id="VUID-vkQueueSubmit-pWaitDstStageMask-00066"></a>
  VUID-vkQueueSubmit-pWaitDstStageMask-00066  
  Any stage flag included in any element of the `pWaitDstStageMask`
  member of any element of `pSubmits` **must** be a pipeline stage
  supported by one of the capabilities of `queue`, as specified in the
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-pipeline-stages-supported"
  target="_blank" rel="noopener">table of supported pipeline stages</a>

- <a href="#VUID-vkQueueSubmit-pSignalSemaphores-00067"
  id="VUID-vkQueueSubmit-pSignalSemaphores-00067"></a>
  VUID-vkQueueSubmit-pSignalSemaphores-00067  
  Each binary semaphore element of the `pSignalSemaphores` member of any
  element of `pSubmits` **must** be unsignaled when the semaphore signal
  operation it defines is executed on the device

- <a href="#VUID-vkQueueSubmit-pWaitSemaphores-00068"
  id="VUID-vkQueueSubmit-pWaitSemaphores-00068"></a>
  VUID-vkQueueSubmit-pWaitSemaphores-00068  
  When a semaphore wait operation referring to a binary semaphore
  defined by any element of the `pWaitSemaphores` member of any element
  of `pSubmits` executes on `queue`, there **must** be no other queues
  waiting on the same semaphore

- <a href="#VUID-vkQueueSubmit-pWaitSemaphores-03238"
  id="VUID-vkQueueSubmit-pWaitSemaphores-03238"></a>
  VUID-vkQueueSubmit-pWaitSemaphores-03238  
  All elements of the `pWaitSemaphores` member of all elements of
  `pSubmits` created with a [VkSemaphoreType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreType.html) of
  `VK_SEMAPHORE_TYPE_BINARY` **must** reference a semaphore signal
  operation that has been submitted for execution and any <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-semaphores-signaling"
  target="_blank" rel="noopener">semaphore signal operations</a> on
  which it depends **must** have also been submitted for execution

- <a href="#VUID-vkQueueSubmit-pCommandBuffers-00070"
  id="VUID-vkQueueSubmit-pCommandBuffers-00070"></a>
  VUID-vkQueueSubmit-pCommandBuffers-00070  
  Each element of the `pCommandBuffers` member of each element of
  `pSubmits` **must** be in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#commandbuffers-lifecycle"
  target="_blank" rel="noopener">pending or executable state</a>

- <a href="#VUID-vkQueueSubmit-pCommandBuffers-00071"
  id="VUID-vkQueueSubmit-pCommandBuffers-00071"></a>
  VUID-vkQueueSubmit-pCommandBuffers-00071  
  If any element of the `pCommandBuffers` member of any element of
  `pSubmits` was not recorded with the
  `VK_COMMAND_BUFFER_USAGE_SIMULTANEOUS_USE_BIT`, it **must** not be in
  the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#commandbuffers-lifecycle"
  target="_blank" rel="noopener">pending state</a>

- <a href="#VUID-vkQueueSubmit-pCommandBuffers-00072"
  id="VUID-vkQueueSubmit-pCommandBuffers-00072"></a>
  VUID-vkQueueSubmit-pCommandBuffers-00072  
  Any <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#commandbuffers-secondary"
  target="_blank" rel="noopener">secondary command buffers recorded</a>
  into any element of the `pCommandBuffers` member of any element of
  `pSubmits` **must** be in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#commandbuffers-lifecycle"
  target="_blank" rel="noopener">pending or executable state</a>

- <a href="#VUID-vkQueueSubmit-pCommandBuffers-00073"
  id="VUID-vkQueueSubmit-pCommandBuffers-00073"></a>
  VUID-vkQueueSubmit-pCommandBuffers-00073  
  If any <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#commandbuffers-secondary"
  target="_blank" rel="noopener">secondary command buffers recorded</a>
  into any element of the `pCommandBuffers` member of any element of
  `pSubmits` was not recorded with the
  `VK_COMMAND_BUFFER_USAGE_SIMULTANEOUS_USE_BIT`, it **must** not be in
  the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#commandbuffers-lifecycle"
  target="_blank" rel="noopener">pending state</a>

- <a href="#VUID-vkQueueSubmit-pCommandBuffers-00074"
  id="VUID-vkQueueSubmit-pCommandBuffers-00074"></a>
  VUID-vkQueueSubmit-pCommandBuffers-00074  
  Each element of the `pCommandBuffers` member of each element of
  `pSubmits` **must** have been allocated from a `VkCommandPool` that
  was created for the same queue family `queue` belongs to

- <a href="#VUID-vkQueueSubmit-pSubmits-02207"
  id="VUID-vkQueueSubmit-pSubmits-02207"></a>
  VUID-vkQueueSubmit-pSubmits-02207  
  If any element of `pSubmits->pCommandBuffers` includes a <a
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
  happens-before the acquire operation

- <a href="#VUID-vkQueueSubmit-pCommandBuffers-03220"
  id="VUID-vkQueueSubmit-pCommandBuffers-03220"></a>
  VUID-vkQueueSubmit-pCommandBuffers-03220  
  If a command recorded into any element of `pCommandBuffers` was a
  [vkCmdBeginQuery](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginQuery.html) whose `queryPool` was created
  with a `queryType` of `VK_QUERY_TYPE_PERFORMANCE_QUERY_KHR`, the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#profiling-lock"
  target="_blank" rel="noopener">profiling lock</a> **must** have been
  held continuously on the `VkDevice` that `queue` was retrieved from,
  throughout recording of those command buffers

- <a href="#VUID-vkQueueSubmit-pSubmits-02808"
  id="VUID-vkQueueSubmit-pSubmits-02808"></a>
  VUID-vkQueueSubmit-pSubmits-02808  
  Any resource created with `VK_SHARING_MODE_EXCLUSIVE` that is read by
  an operation specified by `pSubmits` **must** not be owned by any
  queue family other than the one which `queue` belongs to, at the time
  it is executed

- <a href="#VUID-vkQueueSubmit-pSubmits-04626"
  id="VUID-vkQueueSubmit-pSubmits-04626"></a>
  VUID-vkQueueSubmit-pSubmits-04626  
  Any resource created with `VK_SHARING_MODE_CONCURRENT` that is
  accessed by an operation specified by `pSubmits` **must** have
  included the queue family of `queue` at resource creation time

- <a href="#VUID-vkQueueSubmit-queue-06448"
  id="VUID-vkQueueSubmit-queue-06448"></a>
  VUID-vkQueueSubmit-queue-06448  
  If `queue` was not created with
  `VK_DEVICE_QUEUE_CREATE_PROTECTED_BIT`, there **must** be no element
  of `pSubmits` that includes a
  [VkProtectedSubmitInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkProtectedSubmitInfo.html) structure in its
  `pNext` chain with `protectedSubmit` equal to `VK_TRUE`

Valid Usage (Implicit)

- <a href="#VUID-vkQueueSubmit-queue-parameter"
  id="VUID-vkQueueSubmit-queue-parameter"></a>
  VUID-vkQueueSubmit-queue-parameter  
  `queue` **must** be a valid [VkQueue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueue.html) handle

- <a href="#VUID-vkQueueSubmit-pSubmits-parameter"
  id="VUID-vkQueueSubmit-pSubmits-parameter"></a>
  VUID-vkQueueSubmit-pSubmits-parameter  
  If `submitCount` is not `0`, `pSubmits` **must** be a valid pointer to
  an array of `submitCount` valid [VkSubmitInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubmitInfo.html)
  structures

- <a href="#VUID-vkQueueSubmit-fence-parameter"
  id="VUID-vkQueueSubmit-fence-parameter"></a>
  VUID-vkQueueSubmit-fence-parameter  
  If `fence` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `fence`
  **must** be a valid [VkFence](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFence.html) handle

- <a href="#VUID-vkQueueSubmit-commonparent"
  id="VUID-vkQueueSubmit-commonparent"></a>
  VUID-vkQueueSubmit-commonparent  
  Both of `fence`, and `queue` that are valid handles of non-ignored
  parameters **must** have been created, allocated, or retrieved from
  the same [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

Host Synchronization

- Host access to `queue` **must** be externally synchronized

- Host access to `fence` **must** be externally synchronized

Command Properties

| [Command Buffer Levels](#VkCommandBufferLevel) | [Render Pass Scope](#vkCmdBeginRenderPass) | [Video Coding Scope](#vkCmdBeginVideoCodingKHR) | [Supported Queue Types](#VkQueueFlagBits) | [Command Type](#fundamentals-queueoperation-command-types) |
|------------------------------------------------|--------------------------------------------|-------------------------------------------------|-------------------------------------------|------------------------------------------------------------|
| \-                                             | \-                                         | \-                                              | Any                                       | \-                                                         |

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

- `VK_ERROR_DEVICE_LOST`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html), [VkFence](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFence.html),
[VkQueue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueue.html), [VkSubmitInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubmitInfo.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkQueueSubmit"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
