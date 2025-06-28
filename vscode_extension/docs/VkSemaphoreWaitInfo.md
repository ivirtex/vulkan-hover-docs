# VkSemaphoreWaitInfo(3) Manual Page

## Name

VkSemaphoreWaitInfo - Structure containing information about the semaphore wait condition



## [](#_c_specification)C Specification

The `VkSemaphoreWaitInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_2
typedef struct VkSemaphoreWaitInfo {
    VkStructureType         sType;
    const void*             pNext;
    VkSemaphoreWaitFlags    flags;
    uint32_t                semaphoreCount;
    const VkSemaphore*      pSemaphores;
    const uint64_t*         pValues;
} VkSemaphoreWaitInfo;
```

or the equivalent

```c++
// Provided by VK_KHR_timeline_semaphore
typedef VkSemaphoreWaitInfo VkSemaphoreWaitInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is a bitmask of [VkSemaphoreWaitFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreWaitFlagBits.html) specifying additional parameters for the semaphore wait operation.
- `semaphoreCount` is the number of semaphores to wait on.
- `pSemaphores` is a pointer to an array of `semaphoreCount` semaphore handles to wait on.
- `pValues` is a pointer to an array of `semaphoreCount` timeline semaphore values.

## [](#_description)Description

Valid Usage

- [](#VUID-VkSemaphoreWaitInfo-pSemaphores-03256)VUID-VkSemaphoreWaitInfo-pSemaphores-03256  
  All of the elements of `pSemaphores` **must** reference a semaphore that was created with a [VkSemaphoreType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreType.html) of `VK_SEMAPHORE_TYPE_TIMELINE`

Valid Usage (Implicit)

- [](#VUID-VkSemaphoreWaitInfo-sType-sType)VUID-VkSemaphoreWaitInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SEMAPHORE_WAIT_INFO`
- [](#VUID-VkSemaphoreWaitInfo-pNext-pNext)VUID-VkSemaphoreWaitInfo-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkSemaphoreWaitInfo-flags-parameter)VUID-VkSemaphoreWaitInfo-flags-parameter  
  `flags` **must** be a valid combination of [VkSemaphoreWaitFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreWaitFlagBits.html) values
- [](#VUID-VkSemaphoreWaitInfo-pSemaphores-parameter)VUID-VkSemaphoreWaitInfo-pSemaphores-parameter  
  `pSemaphores` **must** be a valid pointer to an array of `semaphoreCount` valid [VkSemaphore](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphore.html) handles
- [](#VUID-VkSemaphoreWaitInfo-pValues-parameter)VUID-VkSemaphoreWaitInfo-pValues-parameter  
  `pValues` **must** be a valid pointer to an array of `semaphoreCount` `uint64_t` values
- [](#VUID-VkSemaphoreWaitInfo-semaphoreCount-arraylength)VUID-VkSemaphoreWaitInfo-semaphoreCount-arraylength  
  `semaphoreCount` **must** be greater than `0`

## [](#_see_also)See Also

[VK\_KHR\_timeline\_semaphore](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_timeline_semaphore.html), [VK\_VERSION\_1\_2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_2.html), [VkSemaphore](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphore.html), [VkSemaphoreWaitFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreWaitFlags.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkWaitSemaphores](https://registry.khronos.org/vulkan/specs/latest/man/html/vkWaitSemaphores.html), [vkWaitSemaphoresKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkWaitSemaphoresKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSemaphoreWaitInfo)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0