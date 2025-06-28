# VkSetLatencyMarkerInfoNV(3) Manual Page

## Name

VkSetLatencyMarkerInfoNV - Structure specifying the parameters of vkSetLatencyMarkerNV



## [](#_c_specification)C Specification

The `VkSetLatencyMarkerInfoNV` structure is defined as:

```c++
// Provided by VK_NV_low_latency2
typedef struct VkSetLatencyMarkerInfoNV {
    VkStructureType      sType;
    const void*          pNext;
    uint64_t             presentID;
    VkLatencyMarkerNV    marker;
} VkSetLatencyMarkerInfoNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `presentID` is an application provided value that is used to associate the timestamp with a `vkQueuePresentKHR` command using [VkPresentIdKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentIdKHR.html)::`pPresentIds` or [VkPresentId2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentId2KHR.html)::`pPresentIds` for a given present.
- `marker` is the type of timestamp to be recorded.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkSetLatencyMarkerInfoNV-sType-sType)VUID-VkSetLatencyMarkerInfoNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SET_LATENCY_MARKER_INFO_NV`
- [](#VUID-VkSetLatencyMarkerInfoNV-marker-parameter)VUID-VkSetLatencyMarkerInfoNV-marker-parameter  
  `marker` **must** be a valid [VkLatencyMarkerNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkLatencyMarkerNV.html) value

## [](#_see_also)See Also

[VK\_NV\_low\_latency2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_low_latency2.html), [VkLatencyMarkerNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkLatencyMarkerNV.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkSetLatencyMarkerNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkSetLatencyMarkerNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSetLatencyMarkerInfoNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0