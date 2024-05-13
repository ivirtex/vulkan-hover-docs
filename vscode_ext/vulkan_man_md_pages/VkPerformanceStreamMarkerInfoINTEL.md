# VkPerformanceStreamMarkerInfoINTEL(3) Manual Page

## Name

VkPerformanceStreamMarkerInfoINTEL - Structure specifying stream
performance markers



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPerformanceStreamMarkerInfoINTEL` structure is defined as:

``` c
// Provided by VK_INTEL_performance_query
typedef struct VkPerformanceStreamMarkerInfoINTEL {
    VkStructureType    sType;
    const void*        pNext;
    uint32_t           marker;
} VkPerformanceStreamMarkerInfoINTEL;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `marker` is the marker value that will be recorded into the reports
  consumed by an external application.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkPerformanceStreamMarkerInfoINTEL-marker-02735"
  id="VUID-VkPerformanceStreamMarkerInfoINTEL-marker-02735"></a>
  VUID-VkPerformanceStreamMarkerInfoINTEL-marker-02735  
  The value written by the application into `marker` **must** only used
  the valid bits as reported by
  [vkGetPerformanceParameterINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPerformanceParameterINTEL.html)
  with the
  `VK_PERFORMANCE_PARAMETER_TYPE_STREAM_MARKER_VALID_BITS_INTEL`

Valid Usage (Implicit)

- <a href="#VUID-VkPerformanceStreamMarkerInfoINTEL-sType-sType"
  id="VUID-VkPerformanceStreamMarkerInfoINTEL-sType-sType"></a>
  VUID-VkPerformanceStreamMarkerInfoINTEL-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PERFORMANCE_STREAM_MARKER_INFO_INTEL`

- <a href="#VUID-VkPerformanceStreamMarkerInfoINTEL-pNext-pNext"
  id="VUID-VkPerformanceStreamMarkerInfoINTEL-pNext-pNext"></a>
  VUID-VkPerformanceStreamMarkerInfoINTEL-pNext-pNext  
  `pNext` **must** be `NULL`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_INTEL_performance_query](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_INTEL_performance_query.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCmdSetPerformanceStreamMarkerINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetPerformanceStreamMarkerINTEL.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPerformanceStreamMarkerInfoINTEL"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
