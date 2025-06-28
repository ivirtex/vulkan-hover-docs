# VkPerformanceStreamMarkerInfoINTEL(3) Manual Page

## Name

VkPerformanceStreamMarkerInfoINTEL - Structure specifying stream performance markers



## [](#_c_specification)C Specification

The `VkPerformanceStreamMarkerInfoINTEL` structure is defined as:

```c++
// Provided by VK_INTEL_performance_query
typedef struct VkPerformanceStreamMarkerInfoINTEL {
    VkStructureType    sType;
    const void*        pNext;
    uint32_t           marker;
} VkPerformanceStreamMarkerInfoINTEL;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `marker` is the marker value that will be recorded into the reports consumed by an external application.

## [](#_description)Description

Valid Usage

- [](#VUID-VkPerformanceStreamMarkerInfoINTEL-marker-02735)VUID-VkPerformanceStreamMarkerInfoINTEL-marker-02735  
  The value written by the application into `marker` **must** only used the valid bits as reported by [vkGetPerformanceParameterINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPerformanceParameterINTEL.html) with the `VK_PERFORMANCE_PARAMETER_TYPE_STREAM_MARKER_VALID_BITS_INTEL`

Valid Usage (Implicit)

- [](#VUID-VkPerformanceStreamMarkerInfoINTEL-sType-sType)VUID-VkPerformanceStreamMarkerInfoINTEL-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PERFORMANCE_STREAM_MARKER_INFO_INTEL`
- [](#VUID-VkPerformanceStreamMarkerInfoINTEL-pNext-pNext)VUID-VkPerformanceStreamMarkerInfoINTEL-pNext-pNext  
  `pNext` **must** be `NULL`

## [](#_see_also)See Also

[VK\_INTEL\_performance\_query](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_INTEL_performance_query.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCmdSetPerformanceStreamMarkerINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetPerformanceStreamMarkerINTEL.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPerformanceStreamMarkerInfoINTEL)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0