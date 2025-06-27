# VkQueryPoolPerformanceCreateInfoKHR(3) Manual Page

## Name

VkQueryPoolPerformanceCreateInfoKHR - Structure specifying parameters of
a newly created performance query pool



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkQueryPoolPerformanceCreateInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_performance_query
typedef struct VkQueryPoolPerformanceCreateInfoKHR {
    VkStructureType    sType;
    const void*        pNext;
    uint32_t           queueFamilyIndex;
    uint32_t           counterIndexCount;
    const uint32_t*    pCounterIndices;
} VkQueryPoolPerformanceCreateInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `queueFamilyIndex` is the queue family index to create this
  performance query pool for.

- `counterIndexCount` is the length of the `pCounterIndices` array.

- `pCounterIndices` is a pointer to an array of indices into the
  [vkEnumeratePhysicalDeviceQueueFamilyPerformanceQueryCountersKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkEnumeratePhysicalDeviceQueueFamilyPerformanceQueryCountersKHR.html)::`pCounters`
  to enable in this performance query pool.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a
  href="#VUID-VkQueryPoolPerformanceCreateInfoKHR-queueFamilyIndex-03236"
  id="VUID-VkQueryPoolPerformanceCreateInfoKHR-queueFamilyIndex-03236"></a>
  VUID-VkQueryPoolPerformanceCreateInfoKHR-queueFamilyIndex-03236  
  `queueFamilyIndex` **must** be a valid queue family index of the
  device

- <a
  href="#VUID-VkQueryPoolPerformanceCreateInfoKHR-performanceCounterQueryPools-03237"
  id="VUID-VkQueryPoolPerformanceCreateInfoKHR-performanceCounterQueryPools-03237"></a>
  VUID-VkQueryPoolPerformanceCreateInfoKHR-performanceCounterQueryPools-03237  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-performanceCounterQueryPools"
  target="_blank"
  rel="noopener"><code>performanceCounterQueryPools</code></a> feature
  **must** be enabled

- <a
  href="#VUID-VkQueryPoolPerformanceCreateInfoKHR-pCounterIndices-03321"
  id="VUID-VkQueryPoolPerformanceCreateInfoKHR-pCounterIndices-03321"></a>
  VUID-VkQueryPoolPerformanceCreateInfoKHR-pCounterIndices-03321  
  Each element of `pCounterIndices` **must** be in the range of counters
  reported by
  `vkEnumeratePhysicalDeviceQueueFamilyPerformanceQueryCountersKHR` for
  the queue family specified in `queueFamilyIndex`

Valid Usage (Implicit)

- <a href="#VUID-VkQueryPoolPerformanceCreateInfoKHR-sType-sType"
  id="VUID-VkQueryPoolPerformanceCreateInfoKHR-sType-sType"></a>
  VUID-VkQueryPoolPerformanceCreateInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_QUERY_POOL_PERFORMANCE_CREATE_INFO_KHR`

- <a
  href="#VUID-VkQueryPoolPerformanceCreateInfoKHR-pCounterIndices-parameter"
  id="VUID-VkQueryPoolPerformanceCreateInfoKHR-pCounterIndices-parameter"></a>
  VUID-VkQueryPoolPerformanceCreateInfoKHR-pCounterIndices-parameter  
  `pCounterIndices` **must** be a valid pointer to an array of
  `counterIndexCount` `uint32_t` values

- <a
  href="#VUID-VkQueryPoolPerformanceCreateInfoKHR-counterIndexCount-arraylength"
  id="VUID-VkQueryPoolPerformanceCreateInfoKHR-counterIndexCount-arraylength"></a>
  VUID-VkQueryPoolPerformanceCreateInfoKHR-counterIndexCount-arraylength  
  `counterIndexCount` **must** be greater than `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_performance_query](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_performance_query.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetPhysicalDeviceQueueFamilyPerformanceQueryPassesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceQueueFamilyPerformanceQueryPassesKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkQueryPoolPerformanceCreateInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
