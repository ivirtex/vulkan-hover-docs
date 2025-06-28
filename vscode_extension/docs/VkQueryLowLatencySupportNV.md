# VkQueryLowLatencySupportNV(3) Manual Page

## Name

VkQueryLowLatencySupportNV - Structure used for NVIDIA Reflex Support



## [](#_c_specification)C Specification

The `VkQueryLowLatencySupportNV` structure is defined as:

```c++
// Provided by VK_NV_low_latency
typedef struct VkQueryLowLatencySupportNV {
    VkStructureType    sType;
    const void*        pNext;
    void*              pQueriedLowLatencyData;
} VkQueryLowLatencySupportNV;
```

## [](#_members)Members

This structure describes the following feature:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `pQueriedLowLatencyData` is used for NVIDIA Reflex Support.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkQueryLowLatencySupportNV-sType-sType)VUID-VkQueryLowLatencySupportNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_QUERY_LOW_LATENCY_SUPPORT_NV`
- [](#VUID-VkQueryLowLatencySupportNV-pQueriedLowLatencyData-parameter)VUID-VkQueryLowLatencySupportNV-pQueriedLowLatencyData-parameter  
  `pQueriedLowLatencyData` **must** be a pointer value

## [](#_see_also)See Also

[VK\_NV\_low\_latency](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_low_latency.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkQueryLowLatencySupportNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0