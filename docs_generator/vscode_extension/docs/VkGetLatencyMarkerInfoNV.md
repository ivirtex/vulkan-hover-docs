# VkGetLatencyMarkerInfoNV(3) Manual Page

## Name

VkGetLatencyMarkerInfoNV - Structure specifying the parameters of vkGetLatencyTimingsNV



## [](#_c_specification)C Specification

The `VkGetLatencyMarkerInfoNV` structure is defined as:

```c++
// Provided by VK_NV_low_latency2
typedef struct VkGetLatencyMarkerInfoNV {
    VkStructureType                   sType;
    const void*                       pNext;
    uint32_t                          timingCount;
    VkLatencyTimingsFrameReportNV*    pTimings;
} VkGetLatencyMarkerInfoNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is either `NULL` or a pointer to a structure extending this structure.
- `timingCount` is an integer related to the number of previous frames of latency data available or queried, as described below.
- `pTimings` is either `NULL` or a pointer to an array of [VkLatencyTimingsFrameReportNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkLatencyTimingsFrameReportNV.html) structures.

## [](#_description)Description

If `pTimings` is `NULL` then the maximum number of queryable frame data is returned in `timingCount`. Otherwise, `timingCount` **must** be set by the application to the number of elements in the `pTimings` array, and on return the variable is overwritten with the number of values actually written to `pTimings`. The elements of `pTimings` are arranged in the order they were requested in, with the oldest data in the first entry.

Valid Usage (Implicit)

- [](#VUID-VkGetLatencyMarkerInfoNV-sType-sType)VUID-VkGetLatencyMarkerInfoNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_GET_LATENCY_MARKER_INFO_NV`
- [](#VUID-VkGetLatencyMarkerInfoNV-pTimings-parameter)VUID-VkGetLatencyMarkerInfoNV-pTimings-parameter  
  If `timingCount` is not `0`, and `pTimings` is not `NULL`, `pTimings` **must** be a valid pointer to an array of `timingCount` [VkLatencyTimingsFrameReportNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkLatencyTimingsFrameReportNV.html) structures

## [](#_see_also)See Also

[VK\_NV\_low\_latency2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_low_latency2.html), [VkLatencyTimingsFrameReportNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkLatencyTimingsFrameReportNV.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetLatencyTimingsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetLatencyTimingsNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkGetLatencyMarkerInfoNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0