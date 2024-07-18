# VkQueryPoolPerformanceQueryCreateInfoINTEL(3) Manual Page

## Name

VkQueryPoolPerformanceQueryCreateInfoINTEL - Structure specifying
parameters to create a pool of performance queries



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkQueryPoolPerformanceQueryCreateInfoINTEL` structure is defined
as:

``` c
// Provided by VK_INTEL_performance_query
typedef struct VkQueryPoolPerformanceQueryCreateInfoINTEL {
    VkStructureType                 sType;
    const void*                     pNext;
    VkQueryPoolSamplingModeINTEL    performanceCountersSampling;
} VkQueryPoolPerformanceQueryCreateInfoINTEL;
```

``` c
// Provided by VK_INTEL_performance_query
typedef VkQueryPoolPerformanceQueryCreateInfoINTEL VkQueryPoolCreateInfoINTEL;
```

## <a href="#_members" class="anchor"></a>Members

To create a pool for Intel performance queries, set
[VkQueryPoolCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPoolCreateInfo.html)::`queryType` to
`VK_QUERY_TYPE_PERFORMANCE_QUERY_INTEL` and add a
`VkQueryPoolPerformanceQueryCreateInfoINTEL` structure to the `pNext`
chain of the [VkQueryPoolCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPoolCreateInfo.html)
structure.

## <a href="#_description" class="anchor"></a>Description

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `performanceCountersSampling` describe how performance queries should
  be captured.

Valid Usage (Implicit)

- <a href="#VUID-VkQueryPoolPerformanceQueryCreateInfoINTEL-sType-sType"
  id="VUID-VkQueryPoolPerformanceQueryCreateInfoINTEL-sType-sType"></a>
  VUID-VkQueryPoolPerformanceQueryCreateInfoINTEL-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_QUERY_POOL_PERFORMANCE_QUERY_CREATE_INFO_INTEL`

- <a
  href="#VUID-VkQueryPoolPerformanceQueryCreateInfoINTEL-performanceCountersSampling-parameter"
  id="VUID-VkQueryPoolPerformanceQueryCreateInfoINTEL-performanceCountersSampling-parameter"></a>
  VUID-VkQueryPoolPerformanceQueryCreateInfoINTEL-performanceCountersSampling-parameter  
  `performanceCountersSampling` **must** be a valid
  [VkQueryPoolSamplingModeINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPoolSamplingModeINTEL.html)
  value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_INTEL_performance_query](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_INTEL_performance_query.html),
[VkQueryPoolSamplingModeINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPoolSamplingModeINTEL.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkQueryPoolPerformanceQueryCreateInfoINTEL"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
