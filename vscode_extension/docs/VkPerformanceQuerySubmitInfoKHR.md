# VkPerformanceQuerySubmitInfoKHR(3) Manual Page

## Name

VkPerformanceQuerySubmitInfoKHR - Structure indicating which counter pass index is active for performance queries



## [](#_c_specification)C Specification

The `VkPerformanceQuerySubmitInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_performance_query
typedef struct VkPerformanceQuerySubmitInfoKHR {
    VkStructureType    sType;
    const void*        pNext;
    uint32_t           counterPassIndex;
} VkPerformanceQuerySubmitInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `counterPassIndex` specifies which counter pass index is active.

## [](#_description)Description

If the `VkSubmitInfo`::`pNext` chain does not include this structure, the batch defaults to use counter pass index 0.

Valid Usage

- [](#VUID-VkPerformanceQuerySubmitInfoKHR-counterPassIndex-03221)VUID-VkPerformanceQuerySubmitInfoKHR-counterPassIndex-03221  
  `counterPassIndex` **must** be less than the number of counter passes required by any queries within the batch. The required number of counter passes for a performance query is obtained by calling [vkGetPhysicalDeviceQueueFamilyPerformanceQueryPassesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceQueueFamilyPerformanceQueryPassesKHR.html)

Valid Usage (Implicit)

- [](#VUID-VkPerformanceQuerySubmitInfoKHR-sType-sType)VUID-VkPerformanceQuerySubmitInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PERFORMANCE_QUERY_SUBMIT_INFO_KHR`

## [](#_see_also)See Also

[VK\_KHR\_performance\_query](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_performance_query.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPerformanceQuerySubmitInfoKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0