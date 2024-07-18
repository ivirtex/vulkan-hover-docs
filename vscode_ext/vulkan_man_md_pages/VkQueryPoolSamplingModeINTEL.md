# VkQueryPoolSamplingModeINTEL(3) Manual Page

## Name

VkQueryPoolSamplingModeINTEL - Enum specifying how performance queries
should be captured



## <a href="#_c_specification" class="anchor"></a>C Specification

Possible values of
[VkQueryPoolPerformanceQueryCreateInfoINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPoolPerformanceQueryCreateInfoINTEL.html)::`performanceCountersSampling`
are:

``` c
// Provided by VK_INTEL_performance_query
typedef enum VkQueryPoolSamplingModeINTEL {
    VK_QUERY_POOL_SAMPLING_MODE_MANUAL_INTEL = 0,
} VkQueryPoolSamplingModeINTEL;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_QUERY_POOL_SAMPLING_MODE_MANUAL_INTEL` is the default mode in
  which the application calls [vkCmdBeginQuery](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginQuery.html)
  and [vkCmdEndQuery](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdEndQuery.html) to record performance data.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_INTEL_performance_query](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_INTEL_performance_query.html),
[VkQueryPoolPerformanceQueryCreateInfoINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPoolPerformanceQueryCreateInfoINTEL.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkQueryPoolSamplingModeINTEL"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
