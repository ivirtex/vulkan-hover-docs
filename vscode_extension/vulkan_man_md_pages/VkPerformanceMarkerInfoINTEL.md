# VkPerformanceMarkerInfoINTEL(3) Manual Page

## Name

VkPerformanceMarkerInfoINTEL - Structure specifying performance markers



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPerformanceMarkerInfoINTEL` structure is defined as:

``` c
// Provided by VK_INTEL_performance_query
typedef struct VkPerformanceMarkerInfoINTEL {
    VkStructureType    sType;
    const void*        pNext;
    uint64_t           marker;
} VkPerformanceMarkerInfoINTEL;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `marker` is the marker value that will be recorded into the opaque
  query results.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkPerformanceMarkerInfoINTEL-sType-sType"
  id="VUID-VkPerformanceMarkerInfoINTEL-sType-sType"></a>
  VUID-VkPerformanceMarkerInfoINTEL-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PERFORMANCE_MARKER_INFO_INTEL`

- <a href="#VUID-VkPerformanceMarkerInfoINTEL-pNext-pNext"
  id="VUID-VkPerformanceMarkerInfoINTEL-pNext-pNext"></a>
  VUID-VkPerformanceMarkerInfoINTEL-pNext-pNext  
  `pNext` **must** be `NULL`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_INTEL_performance_query](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_INTEL_performance_query.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCmdSetPerformanceMarkerINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetPerformanceMarkerINTEL.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPerformanceMarkerInfoINTEL"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
