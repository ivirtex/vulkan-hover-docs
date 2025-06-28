# VkQueryPoolSamplingModeINTEL(3) Manual Page

## Name

VkQueryPoolSamplingModeINTEL - Enum specifying how performance queries should be captured



## [](#_c_specification)C Specification

Possible values of [VkQueryPoolPerformanceQueryCreateInfoINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPoolPerformanceQueryCreateInfoINTEL.html)::`performanceCountersSampling` are:

```c++
// Provided by VK_INTEL_performance_query
typedef enum VkQueryPoolSamplingModeINTEL {
    VK_QUERY_POOL_SAMPLING_MODE_MANUAL_INTEL = 0,
} VkQueryPoolSamplingModeINTEL;
```

## [](#_description)Description

- `VK_QUERY_POOL_SAMPLING_MODE_MANUAL_INTEL` is the default mode in which the application calls [vkCmdBeginQuery](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBeginQuery.html) and [vkCmdEndQuery](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdEndQuery.html) to record performance data.

## [](#_see_also)See Also

[VK\_INTEL\_performance\_query](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_INTEL_performance_query.html), [VkQueryPoolPerformanceQueryCreateInfoINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPoolPerformanceQueryCreateInfoINTEL.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkQueryPoolSamplingModeINTEL)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0