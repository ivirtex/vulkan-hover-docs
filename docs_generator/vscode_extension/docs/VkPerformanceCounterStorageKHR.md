# VkPerformanceCounterStorageKHR(3) Manual Page

## Name

VkPerformanceCounterStorageKHR - Supported counter storage types



## [](#_c_specification)C Specification

Performance counters have an associated storage. This storage describes the payload of a counter result.

The performance counter storage types which **may** be returned in [VkPerformanceCounterKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceCounterKHR.html)::`storage` are:

```c++
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

## [](#_description)Description

- `VK_PERFORMANCE_COUNTER_STORAGE_INT32_KHR` - the performance counter storage is a 32-bit signed integer.
- `VK_PERFORMANCE_COUNTER_STORAGE_INT64_KHR` - the performance counter storage is a 64-bit signed integer.
- `VK_PERFORMANCE_COUNTER_STORAGE_UINT32_KHR` - the performance counter storage is a 32-bit unsigned integer.
- `VK_PERFORMANCE_COUNTER_STORAGE_UINT64_KHR` - the performance counter storage is a 64-bit unsigned integer.
- `VK_PERFORMANCE_COUNTER_STORAGE_FLOAT32_KHR` - the performance counter storage is a 32-bit floating-point.
- `VK_PERFORMANCE_COUNTER_STORAGE_FLOAT64_KHR` - the performance counter storage is a 64-bit floating-point.

## [](#_see_also)See Also

[VK\_KHR\_performance\_query](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_performance_query.html), [VkPerformanceCounterKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceCounterKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPerformanceCounterStorageKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0