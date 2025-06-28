# VkSubmitInfo2(3) Manual Page

## Name

VkSubmitInfo2 - Structure specifying a queue submit operation



## [](#_c_specification)C Specification

The `VkSubmitInfo2` structure is defined as:

```c++
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

```c++
// Provided by VK_KHR_synchronization2
typedef VkSubmitInfo2 VkSubmitInfo2KHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is a bitmask of [VkSubmitFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubmitFlagBits.html).
- `waitSemaphoreInfoCount` is the number of elements in `pWaitSemaphoreInfos`.
- `pWaitSemaphoreInfos` is a pointer to an array of [VkSemaphoreSubmitInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreSubmitInfo.html) structures defining [semaphore wait operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-semaphores-waiting).
- `commandBufferInfoCount` is the number of elements in `pCommandBufferInfos` and the number of command buffers to execute in the batch.
- `pCommandBufferInfos` is a pointer to an array of [VkCommandBufferSubmitInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferSubmitInfo.html) structures describing command buffers to execute in the batch.
- `signalSemaphoreInfoCount` is the number of elements in `pSignalSemaphoreInfos`.
- `pSignalSemaphoreInfos` is a pointer to an array of [VkSemaphoreSubmitInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreSubmitInfo.html) describing [semaphore signal operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-semaphores-signaling).

## [](#_description)Description

Valid Usage

- [](#VUID-VkSubmitInfo2-semaphore-03881)VUID-VkSubmitInfo2-semaphore-03881  
  If the same semaphore is used as the `semaphore` member of both an element of `pSignalSemaphoreInfos` and `pWaitSemaphoreInfos`, and that semaphore is a timeline semaphore, the `value` member of the `pSignalSemaphoreInfos` element **must** be greater than the `value` member of the `pWaitSemaphoreInfos` element
- [](#VUID-VkSubmitInfo2-semaphore-03882)VUID-VkSubmitInfo2-semaphore-03882  
  If the `semaphore` member of any element of `pSignalSemaphoreInfos` is a timeline semaphore, the `value` member of that element **must** have a value greater than the current value of the semaphore when the [semaphore signal operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-semaphores-signaling) is executed
- [](#VUID-VkSubmitInfo2-semaphore-03883)VUID-VkSubmitInfo2-semaphore-03883  
  If the `semaphore` member of any element of `pSignalSemaphoreInfos` is a timeline semaphore, the `value` member of that element **must** have a value which does not differ from the current value of the semaphore or the value of any outstanding semaphore wait or signal operation on that semaphore by more than [`maxTimelineSemaphoreValueDifference`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxTimelineSemaphoreValueDifference)
- [](#VUID-VkSubmitInfo2-semaphore-03884)VUID-VkSubmitInfo2-semaphore-03884  
  If the `semaphore` member of any element of `pWaitSemaphoreInfos` is a timeline semaphore, the `value` member of that element **must** have a value which does not differ from the current value of the semaphore or the value of any outstanding semaphore wait or signal operation on that semaphore by more than [`maxTimelineSemaphoreValueDifference`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxTimelineSemaphoreValueDifference)
- [](#VUID-VkSubmitInfo2-flags-03886)VUID-VkSubmitInfo2-flags-03886  
  If `flags` includes `VK_SUBMIT_PROTECTED_BIT`, all elements of `pCommandBuffers` **must** be protected command buffers
- [](#VUID-VkSubmitInfo2-flags-03887)VUID-VkSubmitInfo2-flags-03887  
  If `flags` does not include `VK_SUBMIT_PROTECTED_BIT`, each element of `pCommandBuffers` **must** not be a protected command buffer
- [](#VUID-VkSubmitInfo2-commandBuffer-06192)VUID-VkSubmitInfo2-commandBuffer-06192  
  If any `commandBuffer` member of an element of `pCommandBufferInfos` contains any [resumed render pass instances](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-suspension), they **must** be suspended by a render pass instance earlier in submission order within `pCommandBufferInfos`
- [](#VUID-VkSubmitInfo2-commandBuffer-06010)VUID-VkSubmitInfo2-commandBuffer-06010  
  If any `commandBuffer` member of an element of `pCommandBufferInfos` contains any [suspended render pass instances](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-suspension), they **must** be resumed by a render pass instance later in submission order within `pCommandBufferInfos`
- [](#VUID-VkSubmitInfo2-commandBuffer-06011)VUID-VkSubmitInfo2-commandBuffer-06011  
  If any `commandBuffer` member of an element of `pCommandBufferInfos` contains any [suspended render pass instances](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-suspension), there **must** be no action or synchronization commands between that render pass instance and the render pass instance that resumes it
- [](#VUID-VkSubmitInfo2-commandBuffer-06012)VUID-VkSubmitInfo2-commandBuffer-06012  
  If any `commandBuffer` member of an element of `pCommandBufferInfos` contains any [suspended render pass instances](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-suspension), there **must** be no render pass instances between that render pass instance and the render pass instance that resumes it
- [](#VUID-VkSubmitInfo2-variableSampleLocations-06013)VUID-VkSubmitInfo2-variableSampleLocations-06013  
  If the [`variableSampleLocations`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-variableSampleLocations) limit is not supported, and any `commandBuffer` member of an element of `pCommandBufferInfos` contains any [suspended render pass instances](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-suspension), where a graphics pipeline has been bound, any pipelines bound in the render pass instance that resumes it, or any subsequent render pass instances that resume from that one and so on, **must** use the same sample locations
- [](#VUID-VkSubmitInfo2-pNext-09682)VUID-VkSubmitInfo2-pNext-09682  
  If the `pNext` chain of this structure includes a [VkFrameBoundaryTensorsARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFrameBoundaryTensorsARM.html) structure then it **must** also include a [VkFrameBoundaryEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFrameBoundaryEXT.html) structure.
- [](#VUID-VkSubmitInfo2-pCommandBufferInfos-09933)VUID-VkSubmitInfo2-pCommandBufferInfos-09933  
  If at least one [VkCommandBufferSubmitInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferSubmitInfo.html) structure in `pCommandBufferInfos` references a `commandBuffer` allocated from a pool that was created with a [VkDataGraphProcessingEngineCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphProcessingEngineCreateInfoARM.html) structure in the `pNext` chain of [VkCommandPoolCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandPoolCreateInfo.html) that included a foreign data graph processing engine in its `pProcessingEngines` member, then `pWaitSemaphoreInfos` and `pSignalSemaphoreInfos` **must** only reference `semaphore` objects that were created from external handle types reported as supported in a [VkQueueFamilyDataGraphProcessingEnginePropertiesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueFamilyDataGraphProcessingEnginePropertiesARM.html)::`foreignSemaphoreHandleTypes` structure via [vkGetPhysicalDeviceQueueFamilyDataGraphProcessingEnginePropertiesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceQueueFamilyDataGraphProcessingEnginePropertiesARM.html) with a `queueFamilyIndex` matching the one the command pool was created for, for all the foreign data graph processing engines that were part of the [VkDataGraphProcessingEngineCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphProcessingEngineCreateInfoARM.html) used to create the command pool

Valid Usage (Implicit)

- [](#VUID-VkSubmitInfo2-sType-sType)VUID-VkSubmitInfo2-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SUBMIT_INFO_2`
- [](#VUID-VkSubmitInfo2-pNext-pNext)VUID-VkSubmitInfo2-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the `pNext` chain **must** be either `NULL` or a pointer to a valid instance of [VkFrameBoundaryEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFrameBoundaryEXT.html), [VkFrameBoundaryTensorsARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFrameBoundaryTensorsARM.html), [VkLatencySubmissionPresentIdNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkLatencySubmissionPresentIdNV.html), [VkPerformanceQuerySubmitInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceQuerySubmitInfoKHR.html), [VkWin32KeyedMutexAcquireReleaseInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWin32KeyedMutexAcquireReleaseInfoKHR.html), or [VkWin32KeyedMutexAcquireReleaseInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWin32KeyedMutexAcquireReleaseInfoNV.html)
- [](#VUID-VkSubmitInfo2-sType-unique)VUID-VkSubmitInfo2-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique
- [](#VUID-VkSubmitInfo2-flags-parameter)VUID-VkSubmitInfo2-flags-parameter  
  `flags` **must** be a valid combination of [VkSubmitFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubmitFlagBits.html) values
- [](#VUID-VkSubmitInfo2-pWaitSemaphoreInfos-parameter)VUID-VkSubmitInfo2-pWaitSemaphoreInfos-parameter  
  If `waitSemaphoreInfoCount` is not `0`, `pWaitSemaphoreInfos` **must** be a valid pointer to an array of `waitSemaphoreInfoCount` valid [VkSemaphoreSubmitInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreSubmitInfo.html) structures
- [](#VUID-VkSubmitInfo2-pCommandBufferInfos-parameter)VUID-VkSubmitInfo2-pCommandBufferInfos-parameter  
  If `commandBufferInfoCount` is not `0`, `pCommandBufferInfos` **must** be a valid pointer to an array of `commandBufferInfoCount` valid [VkCommandBufferSubmitInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferSubmitInfo.html) structures
- [](#VUID-VkSubmitInfo2-pSignalSemaphoreInfos-parameter)VUID-VkSubmitInfo2-pSignalSemaphoreInfos-parameter  
  If `signalSemaphoreInfoCount` is not `0`, `pSignalSemaphoreInfos` **must** be a valid pointer to an array of `signalSemaphoreInfoCount` valid [VkSemaphoreSubmitInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreSubmitInfo.html) structures

## [](#_see_also)See Also

[VK\_KHR\_synchronization2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_synchronization2.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkCommandBufferSubmitInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferSubmitInfo.html), [VkSemaphoreSubmitInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreSubmitInfo.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkSubmitFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubmitFlags.html), [vkQueueSubmit2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkQueueSubmit2.html), [vkQueueSubmit2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkQueueSubmit2KHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSubmitInfo2)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0