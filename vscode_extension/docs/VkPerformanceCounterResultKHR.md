# VkPerformanceCounterResultKHR(3) Manual Page

## Name

VkPerformanceCounterResultKHR - Union containing a performance counter result



## [](#_c_specification)C Specification

The `VkPerformanceCounterResultKHR` union is defined as:

```c++
// Provided by VK_KHR_performance_query
typedef union VkPerformanceCounterResultKHR {
    int32_t     int32;
    int64_t     int64;
    uint32_t    uint32;
    uint64_t    uint64;
    float       float32;
    double      float64;
} VkPerformanceCounterResultKHR;
```

## [](#_members)Members

- `int32` is a 32-bit signed integer value.
- `int64` is a 64-bit signed integer value.
- `uint32` is a 32-bit unsigned integer value.
- `uint64` is a 64-bit unsigned integer value.
- `float32` is a 32-bit floating-point value.
- `float64` is a 64-bit floating-point value.

## [](#_description)Description

Performance query results are returned in an array of `VkPerformanceCounterResultKHR` unions containing the data associated with each counter in the query, stored in the same order as the counters supplied in `pCounterIndices` when creating the performance query. [VkPerformanceCounterKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceCounterKHR.html)::`storage` specifies how to parse the counter data.

## [](#_see_also)See Also

[VK\_KHR\_performance\_query](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_performance_query.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPerformanceCounterResultKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0