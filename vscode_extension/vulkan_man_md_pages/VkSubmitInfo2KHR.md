# VkSubmitInfo2(3) Manual Page

## Name

VkSubmitInfo2 - Structure specifying a queue submit operation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkSubmitInfo2` structure is defined as:

``` c
// Provided by VK_VERSION_1_3
typedef struct VkSubmitInfo2 {
    VkStructureType                     sType;
    const void*                         pNext;
    VkSubmitFlags                       flags;
    uint32_t                            waitSemaphoreInfoCount;
    const VkSemaphoreSubmitInfo*        pWaitSemaphoreInfos;
    uint32_t                            commandBufferInfoCount;
    const VkCommandBufferSubmitInfo*    pCommandBufferInfos;
    uint32_t                            signalSemaphoreInfoCount;
    const VkSemaphoreSubmitInfo*        pSignalSemaphoreInfos;
} VkSubmitInfo2;
```

or the equivalent

``` c
// Provided by VK_KHR_synchronization2
typedef VkSubmitInfo2 VkSubmitInfo2KHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is a bitmask of [VkSubmitFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubmitFlagBits.html).

- `waitSemaphoreInfoCount` is the number of elements in
  `pWaitSemaphoreInfos`.

- `pWaitSemaphoreInfos` is a pointer to an array of
  [VkSemaphoreSubmitInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreSubmitInfo.html) structures
  defining <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-semaphores-waiting"
  target="_blank" rel="noopener">semaphore wait operations</a>.

- `commandBufferInfoCount` is the number of elements in
  `pCommandBufferInfos` and the number of command buffers to execute in
  the batch.

- `pCommandBufferInfos` is a pointer to an array of
  [VkCommandBufferSubmitInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferSubmitInfo.html) structures
  describing command buffers to execute in the batch.

- `signalSemaphoreInfoCount` is the number of elements in
  `pSignalSemaphoreInfos`.

- `pSignalSemaphoreInfos` is a pointer to an array of
  [VkSemaphoreSubmitInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreSubmitInfo.html) describing <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-semaphores-signaling"
  target="_blank" rel="noopener">semaphore signal operations</a>.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkSubmitInfo2-semaphore-03881"
  id="VUID-VkSubmitInfo2-semaphore-03881"></a>
  VUID-VkSubmitInfo2-semaphore-03881  
  If the same semaphore is used as the `semaphore` member of both an
  element of `pSignalSemaphoreInfos` and `pWaitSemaphoreInfos`, and that
  semaphore is a timeline semaphore, the `value` member of the
  `pSignalSemaphoreInfos` element **must** be greater than the `value`
  member of the `pWaitSemaphoreInfos` element

- <a href="#VUID-VkSubmitInfo2-semaphore-03882"
  id="VUID-VkSubmitInfo2-semaphore-03882"></a>
  VUID-VkSubmitInfo2-semaphore-03882  
  If the `semaphore` member of any element of `pSignalSemaphoreInfos` is
  a timeline semaphore, the `value` member of that element **must** have
  a value greater than the current value of the semaphore when the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-semaphores-signaling"
  target="_blank" rel="noopener">semaphore signal operation</a> is
  executed

- <a href="#VUID-VkSubmitInfo2-semaphore-03883"
  id="VUID-VkSubmitInfo2-semaphore-03883"></a>
  VUID-VkSubmitInfo2-semaphore-03883  
  If the `semaphore` member of any element of `pSignalSemaphoreInfos` is
  a timeline semaphore, the `value` member of that element **must** have
  a value which does not differ from the current value of the semaphore
  or the value of any outstanding semaphore wait or signal operation on
  that semaphore by more than <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxTimelineSemaphoreValueDifference"
  target="_blank"
  rel="noopener"><code>maxTimelineSemaphoreValueDifference</code></a>

- <a href="#VUID-VkSubmitInfo2-semaphore-03884"
  id="VUID-VkSubmitInfo2-semaphore-03884"></a>
  VUID-VkSubmitInfo2-semaphore-03884  
  If the `semaphore` member of any element of `pWaitSemaphoreInfos` is a
  timeline semaphore, the `value` member of that element **must** have a
  value which does not differ from the current value of the semaphore or
  the value of any outstanding semaphore wait or signal operation on
  that semaphore by more than <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxTimelineSemaphoreValueDifference"
  target="_blank"
  rel="noopener"><code>maxTimelineSemaphoreValueDifference</code></a>

- <a href="#VUID-VkSubmitInfo2-flags-03886"
  id="VUID-VkSubmitInfo2-flags-03886"></a>
  VUID-VkSubmitInfo2-flags-03886  
  If `flags` includes `VK_SUBMIT_PROTECTED_BIT`, all elements of
  `pCommandBuffers` **must** be protected command buffers

- <a href="#VUID-VkSubmitInfo2-flags-03887"
  id="VUID-VkSubmitInfo2-flags-03887"></a>
  VUID-VkSubmitInfo2-flags-03887  
  If `flags` does not include `VK_SUBMIT_PROTECTED_BIT`, each element of
  `pCommandBuffers` **must** not be a protected command buffer

- <a href="#VUID-VkSubmitInfo2KHR-commandBuffer-06192"
  id="VUID-VkSubmitInfo2KHR-commandBuffer-06192"></a>
  VUID-VkSubmitInfo2KHR-commandBuffer-06192  
  If any `commandBuffer` member of an element of `pCommandBufferInfos`
  contains any <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass-suspension"
  target="_blank" rel="noopener">resumed render pass instances</a>, they
  **must** be suspended by a render pass instance earlier in submission
  order within `pCommandBufferInfos`

- <a href="#VUID-VkSubmitInfo2KHR-commandBuffer-06010"
  id="VUID-VkSubmitInfo2KHR-commandBuffer-06010"></a>
  VUID-VkSubmitInfo2KHR-commandBuffer-06010  
  If any `commandBuffer` member of an element of `pCommandBufferInfos`
  contains any <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass-suspension"
  target="_blank" rel="noopener">suspended render pass instances</a>,
  they **must** be resumed by a render pass instance later in submission
  order within `pCommandBufferInfos`

- <a href="#VUID-VkSubmitInfo2KHR-commandBuffer-06011"
  id="VUID-VkSubmitInfo2KHR-commandBuffer-06011"></a>
  VUID-VkSubmitInfo2KHR-commandBuffer-06011  
  If any `commandBuffer` member of an element of `pCommandBufferInfos`
  contains any <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass-suspension"
  target="_blank" rel="noopener">suspended render pass instances</a>,
  there **must** be no action or synchronization commands between that
  render pass instance and the render pass instance that resumes it

- <a href="#VUID-VkSubmitInfo2KHR-commandBuffer-06012"
  id="VUID-VkSubmitInfo2KHR-commandBuffer-06012"></a>
  VUID-VkSubmitInfo2KHR-commandBuffer-06012  
  If any `commandBuffer` member of an element of `pCommandBufferInfos`
  contains any <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass-suspension"
  target="_blank" rel="noopener">suspended render pass instances</a>,
  there **must** be no render pass instances between that render pass
  instance and the render pass instance that resumes it

- <a href="#VUID-VkSubmitInfo2KHR-variableSampleLocations-06013"
  id="VUID-VkSubmitInfo2KHR-variableSampleLocations-06013"></a>
  VUID-VkSubmitInfo2KHR-variableSampleLocations-06013  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-variableSampleLocations"
  target="_blank" rel="noopener"><code>variableSampleLocations</code></a>
  limit is not supported, and any `commandBuffer` member of an element
  of `pCommandBufferInfos` contains any <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass-suspension"
  target="_blank" rel="noopener">suspended render pass instances</a>,
  where a graphics pipeline has been bound, any pipelines bound in the
  render pass instance that resumes it, or any subsequent render pass
  instances that resume from that one and so on, **must** use the same
  sample locations

Valid Usage (Implicit)

- <a href="#VUID-VkSubmitInfo2-sType-sType"
  id="VUID-VkSubmitInfo2-sType-sType"></a>
  VUID-VkSubmitInfo2-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SUBMIT_INFO_2`

- <a href="#VUID-VkSubmitInfo2-pNext-pNext"
  id="VUID-VkSubmitInfo2-pNext-pNext"></a>
  VUID-VkSubmitInfo2-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the
  `pNext` chain **must** be either `NULL` or a pointer to a valid
  instance of [VkFrameBoundaryEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFrameBoundaryEXT.html),
  [VkLatencySubmissionPresentIdNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLatencySubmissionPresentIdNV.html),
  [VkPerformanceQuerySubmitInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceQuerySubmitInfoKHR.html),
  [VkWin32KeyedMutexAcquireReleaseInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkWin32KeyedMutexAcquireReleaseInfoKHR.html),
  or
  [VkWin32KeyedMutexAcquireReleaseInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkWin32KeyedMutexAcquireReleaseInfoNV.html)

- <a href="#VUID-VkSubmitInfo2-sType-unique"
  id="VUID-VkSubmitInfo2-sType-unique"></a>
  VUID-VkSubmitInfo2-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

- <a href="#VUID-VkSubmitInfo2-flags-parameter"
  id="VUID-VkSubmitInfo2-flags-parameter"></a>
  VUID-VkSubmitInfo2-flags-parameter  
  `flags` **must** be a valid combination of
  [VkSubmitFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubmitFlagBits.html) values

- <a href="#VUID-VkSubmitInfo2-pWaitSemaphoreInfos-parameter"
  id="VUID-VkSubmitInfo2-pWaitSemaphoreInfos-parameter"></a>
  VUID-VkSubmitInfo2-pWaitSemaphoreInfos-parameter  
  If `waitSemaphoreInfoCount` is not `0`, `pWaitSemaphoreInfos` **must**
  be a valid pointer to an array of `waitSemaphoreInfoCount` valid
  [VkSemaphoreSubmitInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreSubmitInfo.html) structures

- <a href="#VUID-VkSubmitInfo2-pCommandBufferInfos-parameter"
  id="VUID-VkSubmitInfo2-pCommandBufferInfos-parameter"></a>
  VUID-VkSubmitInfo2-pCommandBufferInfos-parameter  
  If `commandBufferInfoCount` is not `0`, `pCommandBufferInfos` **must**
  be a valid pointer to an array of `commandBufferInfoCount` valid
  [VkCommandBufferSubmitInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferSubmitInfo.html) structures

- <a href="#VUID-VkSubmitInfo2-pSignalSemaphoreInfos-parameter"
  id="VUID-VkSubmitInfo2-pSignalSemaphoreInfos-parameter"></a>
  VUID-VkSubmitInfo2-pSignalSemaphoreInfos-parameter  
  If `signalSemaphoreInfoCount` is not `0`, `pSignalSemaphoreInfos`
  **must** be a valid pointer to an array of `signalSemaphoreInfoCount`
  valid [VkSemaphoreSubmitInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreSubmitInfo.html) structures

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_synchronization2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_synchronization2.html),
[VK_VERSION_1_3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_3.html),
[VkCommandBufferSubmitInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferSubmitInfo.html),
[VkSemaphoreSubmitInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreSubmitInfo.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkSubmitFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubmitFlags.html),
[vkQueueSubmit2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkQueueSubmit2.html),
[vkQueueSubmit2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkQueueSubmit2KHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSubmitInfo2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
