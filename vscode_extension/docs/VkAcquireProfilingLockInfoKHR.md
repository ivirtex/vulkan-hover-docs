# VkAcquireProfilingLockInfoKHR(3) Manual Page

## Name

VkAcquireProfilingLockInfoKHR - Structure specifying parameters to acquire the profiling lock



## [](#_c_specification)C Specification

The `VkAcquireProfilingLockInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_performance_query
typedef struct VkAcquireProfilingLockInfoKHR {
    VkStructureType                   sType;
    const void*                       pNext;
    VkAcquireProfilingLockFlagsKHR    flags;
    uint64_t                          timeout;
} VkAcquireProfilingLockInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is reserved for future use.
- `timeout` indicates how long the function waits, in nanoseconds, if the profiling lock is not available.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkAcquireProfilingLockInfoKHR-sType-sType)VUID-VkAcquireProfilingLockInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_ACQUIRE_PROFILING_LOCK_INFO_KHR`
- [](#VUID-VkAcquireProfilingLockInfoKHR-pNext-pNext)VUID-VkAcquireProfilingLockInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkAcquireProfilingLockInfoKHR-flags-zerobitmask)VUID-VkAcquireProfilingLockInfoKHR-flags-zerobitmask  
  `flags` **must** be `0`

If `timeout` is 0, `vkAcquireProfilingLockKHR` will not block while attempting to acquire the profiling lock. If `timeout` is `UINT64_MAX`, the function will not return until the profiling lock was acquired.

## [](#_see_also)See Also

[VK\_KHR\_performance\_query](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_performance_query.html), [VkAcquireProfilingLockFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAcquireProfilingLockFlagsKHR.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkAcquireProfilingLockKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkAcquireProfilingLockKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkAcquireProfilingLockInfoKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0