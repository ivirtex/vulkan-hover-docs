# VkQueryLowLatencySupportNV(3) Manual Page

## Name

VkQueryLowLatencySupportNV - Structure used for NVIDIA Reflex Support



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkQueryLowLatencySupportNV` structure is defined as:

``` c
// Provided by VK_NV_low_latency
typedef struct VkQueryLowLatencySupportNV {
    VkStructureType    sType;
    const void*        pNext;
    void*              pQueriedLowLatencyData;
} VkQueryLowLatencySupportNV;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following feature:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `pQueriedLowLatencyData` is used for NVIDIA Reflex Support.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkQueryLowLatencySupportNV-sType-sType"
  id="VUID-VkQueryLowLatencySupportNV-sType-sType"></a>
  VUID-VkQueryLowLatencySupportNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_QUERY_LOW_LATENCY_SUPPORT_NV`

- <a
  href="#VUID-VkQueryLowLatencySupportNV-pQueriedLowLatencyData-parameter"
  id="VUID-VkQueryLowLatencySupportNV-pQueriedLowLatencyData-parameter"></a>
  VUID-VkQueryLowLatencySupportNV-pQueriedLowLatencyData-parameter  
  `pQueriedLowLatencyData` **must** be a pointer value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_low_latency](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_low_latency.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkQueryLowLatencySupportNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
