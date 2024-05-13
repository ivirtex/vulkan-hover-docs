# VkPerformanceQuerySubmitInfoKHR(3) Manual Page

## Name

VkPerformanceQuerySubmitInfoKHR - Structure indicating which counter
pass index is active for performance queries



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPerformanceQuerySubmitInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_performance_query
typedef struct VkPerformanceQuerySubmitInfoKHR {
    VkStructureType    sType;
    const void*        pNext;
    uint32_t           counterPassIndex;
} VkPerformanceQuerySubmitInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `counterPassIndex` specifies which counter pass index is active.

## <a href="#_description" class="anchor"></a>Description

If the `VkSubmitInfo`::`pNext` chain does not include this structure,
the batch defaults to use counter pass index 0.

Valid Usage

- <a href="#VUID-VkPerformanceQuerySubmitInfoKHR-counterPassIndex-03221"
  id="VUID-VkPerformanceQuerySubmitInfoKHR-counterPassIndex-03221"></a>
  VUID-VkPerformanceQuerySubmitInfoKHR-counterPassIndex-03221  
  `counterPassIndex` **must** be less than the number of counter passes
  required by any queries within the batch. The required number of
  counter passes for a performance query is obtained by calling
  [vkGetPhysicalDeviceQueueFamilyPerformanceQueryPassesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceQueueFamilyPerformanceQueryPassesKHR.html)

Valid Usage (Implicit)

- <a href="#VUID-VkPerformanceQuerySubmitInfoKHR-sType-sType"
  id="VUID-VkPerformanceQuerySubmitInfoKHR-sType-sType"></a>
  VUID-VkPerformanceQuerySubmitInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PERFORMANCE_QUERY_SUBMIT_INFO_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_performance_query](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_performance_query.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPerformanceQuerySubmitInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
