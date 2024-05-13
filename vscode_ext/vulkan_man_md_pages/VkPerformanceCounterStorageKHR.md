# VkPerformanceCounterStorageKHR(3) Manual Page

## Name

VkPerformanceCounterStorageKHR - Supported counter storage types



## <a href="#_c_specification" class="anchor"></a>C Specification

Performance counters have an associated storage. This storage describes
the payload of a counter result.

The performance counter storage types which **may** be returned in
[VkPerformanceCounterKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceCounterKHR.html)::`storage` are:

``` c
// Provided by VK_KHR_performance_query
typedef enum VkPerformanceCounterStorageKHR {
    VK_PERFORMANCE_COUNTER_STORAGE_INT32_KHR = 0,
    VK_PERFORMANCE_COUNTER_STORAGE_INT64_KHR = 1,
    VK_PERFORMANCE_COUNTER_STORAGE_UINT32_KHR = 2,
    VK_PERFORMANCE_COUNTER_STORAGE_UINT64_KHR = 3,
    VK_PERFORMANCE_COUNTER_STORAGE_FLOAT32_KHR = 4,
    VK_PERFORMANCE_COUNTER_STORAGE_FLOAT64_KHR = 5,
} VkPerformanceCounterStorageKHR;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_PERFORMANCE_COUNTER_STORAGE_INT32_KHR` - the performance counter
  storage is a 32-bit signed integer.

- `VK_PERFORMANCE_COUNTER_STORAGE_INT64_KHR` - the performance counter
  storage is a 64-bit signed integer.

- `VK_PERFORMANCE_COUNTER_STORAGE_UINT32_KHR` - the performance counter
  storage is a 32-bit unsigned integer.

- `VK_PERFORMANCE_COUNTER_STORAGE_UINT64_KHR` - the performance counter
  storage is a 64-bit unsigned integer.

- `VK_PERFORMANCE_COUNTER_STORAGE_FLOAT32_KHR` - the performance counter
  storage is a 32-bit floating-point.

- `VK_PERFORMANCE_COUNTER_STORAGE_FLOAT64_KHR` - the performance counter
  storage is a 64-bit floating-point.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_performance_query](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_performance_query.html),
[VkPerformanceCounterKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceCounterKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPerformanceCounterStorageKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
