# VkGetLatencyMarkerInfoNV(3) Manual Page

## Name

VkGetLatencyMarkerInfoNV - Structure specifying the parameters of
vkGetLatencyTimingsNV



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkGetLatencyMarkerInfoNV` structure is defined as:

``` c
// Provided by VK_NV_low_latency2
typedef struct VkGetLatencyMarkerInfoNV {
    VkStructureType                   sType;
    const void*                       pNext;
    uint32_t                          timingCount;
    VkLatencyTimingsFrameReportNV*    pTimings;
} VkGetLatencyMarkerInfoNV;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is either `NULL` or a pointer to a structure extending this
  structure.

- `timingCount` is an integer related to the number of previous frames
  of latency data available or queried, as described below.

- `pTimings` is either `NULL` or a pointer to an array of
  [VkLatencyTimingsFrameReportNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLatencyTimingsFrameReportNV.html)
  structures.

## <a href="#_description" class="anchor"></a>Description

If `pTimings` is `NULL` then the maximum number of queryable frame data
is returned in `timingCount`. Otherwise, `timingCount` **must** be set
by the user to the number of elements in the `pTimings` array, and on
return the variable is overwritten with the number of values actually
written to `pTimings`. The elements of `pTimings` are arranged in the
order they were requested in, with the oldest data in the first entry.

Valid Usage (Implicit)

- <a href="#VUID-VkGetLatencyMarkerInfoNV-sType-sType"
  id="VUID-VkGetLatencyMarkerInfoNV-sType-sType"></a>
  VUID-VkGetLatencyMarkerInfoNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_GET_LATENCY_MARKER_INFO_NV`

- <a href="#VUID-VkGetLatencyMarkerInfoNV-pTimings-parameter"
  id="VUID-VkGetLatencyMarkerInfoNV-pTimings-parameter"></a>
  VUID-VkGetLatencyMarkerInfoNV-pTimings-parameter  
  If `timingCount` is not `0`, and `pTimings` is not `NULL`, `pTimings`
  **must** be a valid pointer to an array of `timingCount`
  [VkLatencyTimingsFrameReportNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLatencyTimingsFrameReportNV.html)
  structures

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_low_latency2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_low_latency2.html),
[VkLatencyTimingsFrameReportNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLatencyTimingsFrameReportNV.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetLatencyTimingsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetLatencyTimingsNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkGetLatencyMarkerInfoNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
