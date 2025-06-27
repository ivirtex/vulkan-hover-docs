# VkPerformanceCounterResultKHR(3) Manual Page

## Name

VkPerformanceCounterResultKHR - Union containing a performance counter
result



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPerformanceCounterResultKHR` union is defined as:

``` c
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

## <a href="#_members" class="anchor"></a>Members

- `int32` is a 32-bit signed integer value.

- `int64` is a 64-bit signed integer value.

- `uint32` is a 32-bit unsigned integer value.

- `uint64` is a 64-bit unsigned integer value.

- `float32` is a 32-bit floating-point value.

- `float64` is a 64-bit floating-point value.

## <a href="#_description" class="anchor"></a>Description

Performance query results are returned in an array of
`VkPerformanceCounterResultKHR` unions containing the data associated
with each counter in the query, stored in the same order as the counters
supplied in `pCounterIndices` when creating the performance query.
[VkPerformanceCounterKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceCounterKHR.html)::`storage`
specifies how to parse the counter data.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_performance_query](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_performance_query.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPerformanceCounterResultKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
