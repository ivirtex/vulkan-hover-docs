# VkSetLatencyMarkerInfoNV(3) Manual Page

## Name

VkSetLatencyMarkerInfoNV - Structure specifying the parameters of
vkSetLatencyMarkerNV



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkSetLatencyMarkerInfoNV` structure is defined as:

``` c
// Provided by VK_NV_low_latency2
typedef struct VkSetLatencyMarkerInfoNV {
    VkStructureType      sType;
    const void*          pNext;
    uint64_t             presentID;
    VkLatencyMarkerNV    marker;
} VkSetLatencyMarkerInfoNV;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `presentId` is an application provided value that is used to associate
  the timestamp with a `vkQueuePresentKHR` command using
  [VkPresentIdKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentIdKHR.html)::`pPresentIds` for a given
  present.

- `marker` is the type of timestamp to be recorded.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkSetLatencyMarkerInfoNV-sType-sType"
  id="VUID-VkSetLatencyMarkerInfoNV-sType-sType"></a>
  VUID-VkSetLatencyMarkerInfoNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SET_LATENCY_MARKER_INFO_NV`

- <a href="#VUID-VkSetLatencyMarkerInfoNV-marker-parameter"
  id="VUID-VkSetLatencyMarkerInfoNV-marker-parameter"></a>
  VUID-VkSetLatencyMarkerInfoNV-marker-parameter  
  `marker` **must** be a valid
  [VkLatencyMarkerNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLatencyMarkerNV.html) value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_low_latency2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_low_latency2.html),
[VkLatencyMarkerNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLatencyMarkerNV.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkSetLatencyMarkerNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkSetLatencyMarkerNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSetLatencyMarkerInfoNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
