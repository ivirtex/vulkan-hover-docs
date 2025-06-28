# VkQueryPoolPerformanceQueryCreateInfoINTEL(3) Manual Page

## Name

VkQueryPoolPerformanceQueryCreateInfoINTEL - Structure specifying parameters to create a pool of performance queries



## [](#_c_specification)C Specification

The `VkQueryPoolPerformanceQueryCreateInfoINTEL` structure is defined as:

```c++
// Provided by VK_INTEL_performance_query
typedef struct VkQueryPoolPerformanceQueryCreateInfoINTEL {
    VkStructureType                 sType;
    const void*                     pNext;
    VkQueryPoolSamplingModeINTEL    performanceCountersSampling;
} VkQueryPoolPerformanceQueryCreateInfoINTEL;
```

```c++
// Provided by VK_INTEL_performance_query
typedef VkQueryPoolPerformanceQueryCreateInfoINTEL VkQueryPoolCreateInfoINTEL;
```

## [](#_members)Members

To create a pool for Intel performance queries, set [VkQueryPoolCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPoolCreateInfo.html)::`queryType` to `VK_QUERY_TYPE_PERFORMANCE_QUERY_INTEL` and add a `VkQueryPoolPerformanceQueryCreateInfoINTEL` structure to the `pNext` chain of the [VkQueryPoolCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPoolCreateInfo.html) structure.

## [](#_description)Description

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `performanceCountersSampling` describe how performance queries should be captured.

Valid Usage (Implicit)

- [](#VUID-VkQueryPoolPerformanceQueryCreateInfoINTEL-sType-sType)VUID-VkQueryPoolPerformanceQueryCreateInfoINTEL-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_QUERY_POOL_PERFORMANCE_QUERY_CREATE_INFO_INTEL`
- [](#VUID-VkQueryPoolPerformanceQueryCreateInfoINTEL-performanceCountersSampling-parameter)VUID-VkQueryPoolPerformanceQueryCreateInfoINTEL-performanceCountersSampling-parameter  
  `performanceCountersSampling` **must** be a valid [VkQueryPoolSamplingModeINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPoolSamplingModeINTEL.html) value

## [](#_see_also)See Also

[VK\_INTEL\_performance\_query](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_INTEL_performance_query.html), [VkQueryPoolSamplingModeINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPoolSamplingModeINTEL.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkQueryPoolPerformanceQueryCreateInfoINTEL)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0