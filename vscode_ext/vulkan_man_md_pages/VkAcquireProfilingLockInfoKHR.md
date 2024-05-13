# VkAcquireProfilingLockInfoKHR(3) Manual Page

## Name

VkAcquireProfilingLockInfoKHR - Structure specifying parameters to
acquire the profiling lock



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkAcquireProfilingLockInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_performance_query
typedef struct VkAcquireProfilingLockInfoKHR {
    VkStructureType                   sType;
    const void*                       pNext;
    VkAcquireProfilingLockFlagsKHR    flags;
    uint64_t                          timeout;
} VkAcquireProfilingLockInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is reserved for future use.

- `timeout` indicates how long the function waits, in nanoseconds, if
  the profiling lock is not available.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkAcquireProfilingLockInfoKHR-sType-sType"
  id="VUID-VkAcquireProfilingLockInfoKHR-sType-sType"></a>
  VUID-VkAcquireProfilingLockInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_ACQUIRE_PROFILING_LOCK_INFO_KHR`

- <a href="#VUID-VkAcquireProfilingLockInfoKHR-pNext-pNext"
  id="VUID-VkAcquireProfilingLockInfoKHR-pNext-pNext"></a>
  VUID-VkAcquireProfilingLockInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkAcquireProfilingLockInfoKHR-flags-zerobitmask"
  id="VUID-VkAcquireProfilingLockInfoKHR-flags-zerobitmask"></a>
  VUID-VkAcquireProfilingLockInfoKHR-flags-zerobitmask  
  `flags` **must** be `0`

If `timeout` is 0, `vkAcquireProfilingLockKHR` will not block while
attempting to acquire the profiling lock. If `timeout` is `UINT64_MAX`,
the function will not return until the profiling lock was acquired.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_performance_query](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_performance_query.html),
[VkAcquireProfilingLockFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAcquireProfilingLockFlagsKHR.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkAcquireProfilingLockKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkAcquireProfilingLockKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkAcquireProfilingLockInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
