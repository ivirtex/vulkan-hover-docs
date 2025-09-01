# VkSemaphoreSignalInfo(3) Manual Page

## Name

VkSemaphoreSignalInfo - Structure containing information about a semaphore signal operation



## [](#_c_specification)C Specification

The `VkSemaphoreSignalInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_2
typedef struct VkSemaphoreSignalInfo {
    VkStructureType    sType;
    const void*        pNext;
    VkSemaphore        semaphore;
    uint64_t           value;
} VkSemaphoreSignalInfo;
```

or the equivalent

```c++
// Provided by VK_KHR_timeline_semaphore
typedef VkSemaphoreSignalInfo VkSemaphoreSignalInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `semaphore` is the handle of the semaphore to signal.
- `value` is the value to signal.

## [](#_description)Description

Valid Usage

- [](#VUID-VkSemaphoreSignalInfo-semaphore-03257)VUID-VkSemaphoreSignalInfo-semaphore-03257  
  `semaphore` **must** have been created with a [VkSemaphoreType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreType.html) of `VK_SEMAPHORE_TYPE_TIMELINE`
- [](#VUID-VkSemaphoreSignalInfo-value-03258)VUID-VkSemaphoreSignalInfo-value-03258  
  `value` **must** have a value greater than the current value of the semaphore
- [](#VUID-VkSemaphoreSignalInfo-value-03259)VUID-VkSemaphoreSignalInfo-value-03259  
  `value` **must** be less than the value of any pending semaphore signal operations
- [](#VUID-VkSemaphoreSignalInfo-value-03260)VUID-VkSemaphoreSignalInfo-value-03260  
  `value` **must** have a value which does not differ from the current value of the semaphore or the value of any outstanding semaphore wait or signal operation on `semaphore` by more than [`maxTimelineSemaphoreValueDifference`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxTimelineSemaphoreValueDifference)

Valid Usage (Implicit)

- [](#VUID-VkSemaphoreSignalInfo-sType-sType)VUID-VkSemaphoreSignalInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SEMAPHORE_SIGNAL_INFO`
- [](#VUID-VkSemaphoreSignalInfo-pNext-pNext)VUID-VkSemaphoreSignalInfo-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkSemaphoreSignalInfo-semaphore-parameter)VUID-VkSemaphoreSignalInfo-semaphore-parameter  
  `semaphore` **must** be a valid [VkSemaphore](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphore.html) handle

## [](#_see_also)See Also

[VK\_KHR\_timeline\_semaphore](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_timeline_semaphore.html), [VK\_VERSION\_1\_2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_2.html), [VkSemaphore](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphore.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkSignalSemaphore](https://registry.khronos.org/vulkan/specs/latest/man/html/vkSignalSemaphore.html), [vkSignalSemaphore](https://registry.khronos.org/vulkan/specs/latest/man/html/vkSignalSemaphore.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSemaphoreSignalInfo).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0