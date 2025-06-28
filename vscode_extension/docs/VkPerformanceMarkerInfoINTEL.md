# VkPerformanceMarkerInfoINTEL(3) Manual Page

## Name

VkPerformanceMarkerInfoINTEL - Structure specifying performance markers



## [](#_c_specification)C Specification

The `VkPerformanceMarkerInfoINTEL` structure is defined as:

```c++
// Provided by VK_INTEL_performance_query
typedef struct VkPerformanceMarkerInfoINTEL {
    VkStructureType    sType;
    const void*        pNext;
    uint64_t           marker;
} VkPerformanceMarkerInfoINTEL;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `marker` is the marker value that will be recorded into the opaque query results.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkPerformanceMarkerInfoINTEL-sType-sType)VUID-VkPerformanceMarkerInfoINTEL-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PERFORMANCE_MARKER_INFO_INTEL`
- [](#VUID-VkPerformanceMarkerInfoINTEL-pNext-pNext)VUID-VkPerformanceMarkerInfoINTEL-pNext-pNext  
  `pNext` **must** be `NULL`

## [](#_see_also)See Also

[VK\_INTEL\_performance\_query](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_INTEL_performance_query.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCmdSetPerformanceMarkerINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetPerformanceMarkerINTEL.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPerformanceMarkerInfoINTEL)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0