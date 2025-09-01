# VkQueryPoolPerformanceCreateInfoKHR(3) Manual Page

## Name

VkQueryPoolPerformanceCreateInfoKHR - Structure specifying parameters of a newly created performance query pool



## [](#_c_specification)C Specification

The `VkQueryPoolPerformanceCreateInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_performance_query
typedef struct VkQueryPoolPerformanceCreateInfoKHR {
    VkStructureType    sType;
    const void*        pNext;
    uint32_t           queueFamilyIndex;
    uint32_t           counterIndexCount;
    const uint32_t*    pCounterIndices;
} VkQueryPoolPerformanceCreateInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `queueFamilyIndex` is the queue family index to create this performance query pool for.
- `counterIndexCount` is the length of the `pCounterIndices` array.
- `pCounterIndices` is a pointer to an array of indices into the [vkEnumeratePhysicalDeviceQueueFamilyPerformanceQueryCountersKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkEnumeratePhysicalDeviceQueueFamilyPerformanceQueryCountersKHR.html)::`pCounters` to enable in this performance query pool.

## [](#_description)Description

Valid Usage

- [](#VUID-VkQueryPoolPerformanceCreateInfoKHR-queueFamilyIndex-03236)VUID-VkQueryPoolPerformanceCreateInfoKHR-queueFamilyIndex-03236  
  `queueFamilyIndex` **must** be a valid queue family index of the device
- [](#VUID-VkQueryPoolPerformanceCreateInfoKHR-performanceCounterQueryPools-03237)VUID-VkQueryPoolPerformanceCreateInfoKHR-performanceCounterQueryPools-03237  
  The [`performanceCounterQueryPools`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-performanceCounterQueryPools) feature **must** be enabled
- [](#VUID-VkQueryPoolPerformanceCreateInfoKHR-pCounterIndices-03321)VUID-VkQueryPoolPerformanceCreateInfoKHR-pCounterIndices-03321  
  Each element of `pCounterIndices` **must** be in the range of counters reported by `vkEnumeratePhysicalDeviceQueueFamilyPerformanceQueryCountersKHR` for the queue family specified in `queueFamilyIndex`

Valid Usage (Implicit)

- [](#VUID-VkQueryPoolPerformanceCreateInfoKHR-sType-sType)VUID-VkQueryPoolPerformanceCreateInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_QUERY_POOL_PERFORMANCE_CREATE_INFO_KHR`
- [](#VUID-VkQueryPoolPerformanceCreateInfoKHR-pCounterIndices-parameter)VUID-VkQueryPoolPerformanceCreateInfoKHR-pCounterIndices-parameter  
  `pCounterIndices` **must** be a valid pointer to an array of `counterIndexCount` `uint32_t` values
- [](#VUID-VkQueryPoolPerformanceCreateInfoKHR-counterIndexCount-arraylength)VUID-VkQueryPoolPerformanceCreateInfoKHR-counterIndexCount-arraylength  
  `counterIndexCount` **must** be greater than `0`

## [](#_see_also)See Also

[VK\_KHR\_performance\_query](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_performance_query.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetPhysicalDeviceQueueFamilyPerformanceQueryPassesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceQueueFamilyPerformanceQueryPassesKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkQueryPoolPerformanceCreateInfoKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0