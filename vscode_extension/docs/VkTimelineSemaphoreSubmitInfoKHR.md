# VkTimelineSemaphoreSubmitInfo(3) Manual Page

## Name

VkTimelineSemaphoreSubmitInfo - Structure specifying signal and wait values for timeline semaphores



## [](#_c_specification)C Specification

To specify the values to use when waiting for and signaling semaphores created with a [VkSemaphoreType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreType.html) of `VK_SEMAPHORE_TYPE_TIMELINE`, add a [VkTimelineSemaphoreSubmitInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTimelineSemaphoreSubmitInfo.html) structure to the `pNext` chain of the [VkSubmitInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubmitInfo.html) structure when using [vkQueueSubmit](https://registry.khronos.org/vulkan/specs/latest/man/html/vkQueueSubmit.html) or the [VkBindSparseInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindSparseInfo.html) structure when using [vkQueueBindSparse](https://registry.khronos.org/vulkan/specs/latest/man/html/vkQueueBindSparse.html) . The `VkTimelineSemaphoreSubmitInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_2
typedef struct VkTimelineSemaphoreSubmitInfo {
    VkStructureType    sType;
    const void*        pNext;
    uint32_t           waitSemaphoreValueCount;
    const uint64_t*    pWaitSemaphoreValues;
    uint32_t           signalSemaphoreValueCount;
    const uint64_t*    pSignalSemaphoreValues;
} VkTimelineSemaphoreSubmitInfo;
```

or the equivalent

```c++
// Provided by VK_KHR_timeline_semaphore
typedef VkTimelineSemaphoreSubmitInfo VkTimelineSemaphoreSubmitInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `waitSemaphoreValueCount` is the number of semaphore wait values specified in `pWaitSemaphoreValues`.
- `pWaitSemaphoreValues` is a pointer to an array of `waitSemaphoreValueCount` values for the corresponding semaphores in [VkSubmitInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubmitInfo.html)::`pWaitSemaphores` to wait for.
- `signalSemaphoreValueCount` is the number of semaphore signal values specified in `pSignalSemaphoreValues`.
- `pSignalSemaphoreValues` is a pointer to an array `signalSemaphoreValueCount` values for the corresponding semaphores in [VkSubmitInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubmitInfo.html)::`pSignalSemaphores` to set when signaled.

## [](#_description)Description

If the semaphore in [VkSubmitInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubmitInfo.html)::`pWaitSemaphores` or [VkSubmitInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubmitInfo.html)::`pSignalSemaphores` corresponding to an entry in `pWaitSemaphoreValues` or `pSignalSemaphoreValues` respectively was not created with a [VkSemaphoreType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreType.html) of `VK_SEMAPHORE_TYPE_TIMELINE`, the implementation **must** ignore the value in the `pWaitSemaphoreValues` or `pSignalSemaphoreValues` entry.

Valid Usage (Implicit)

- [](#VUID-VkTimelineSemaphoreSubmitInfo-sType-sType)VUID-VkTimelineSemaphoreSubmitInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_TIMELINE_SEMAPHORE_SUBMIT_INFO`
- [](#VUID-VkTimelineSemaphoreSubmitInfo-pWaitSemaphoreValues-parameter)VUID-VkTimelineSemaphoreSubmitInfo-pWaitSemaphoreValues-parameter  
  If `waitSemaphoreValueCount` is not `0`, and `pWaitSemaphoreValues` is not `NULL`, `pWaitSemaphoreValues` **must** be a valid pointer to an array of `waitSemaphoreValueCount` `uint64_t` values
- [](#VUID-VkTimelineSemaphoreSubmitInfo-pSignalSemaphoreValues-parameter)VUID-VkTimelineSemaphoreSubmitInfo-pSignalSemaphoreValues-parameter  
  If `signalSemaphoreValueCount` is not `0`, and `pSignalSemaphoreValues` is not `NULL`, `pSignalSemaphoreValues` **must** be a valid pointer to an array of `signalSemaphoreValueCount` `uint64_t` values

## [](#_see_also)See Also

[VK\_KHR\_timeline\_semaphore](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_timeline_semaphore.html), [VK\_VERSION\_1\_2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_2.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkTimelineSemaphoreSubmitInfo)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0