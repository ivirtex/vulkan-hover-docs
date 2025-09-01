# VkInitializePerformanceApiInfoINTEL(3) Manual Page

## Name

VkInitializePerformanceApiInfoINTEL - Structure specifying parameters of initialize of the device



## [](#_c_specification)C Specification

The `VkInitializePerformanceApiInfoINTEL` structure is defined as :

```c++
// Provided by VK_INTEL_performance_query
typedef struct VkInitializePerformanceApiInfoINTEL {
    VkStructureType    sType;
    const void*        pNext;
    void*              pUserData;
} VkInitializePerformanceApiInfoINTEL;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `pUserData` is a pointer for application data.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkInitializePerformanceApiInfoINTEL-sType-sType)VUID-VkInitializePerformanceApiInfoINTEL-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_INITIALIZE_PERFORMANCE_API_INFO_INTEL`
- [](#VUID-VkInitializePerformanceApiInfoINTEL-pNext-pNext)VUID-VkInitializePerformanceApiInfoINTEL-pNext-pNext  
  `pNext` **must** be `NULL`

## [](#_see_also)See Also

[VK\_INTEL\_performance\_query](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_INTEL_performance_query.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkInitializePerformanceApiINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/vkInitializePerformanceApiINTEL.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkInitializePerformanceApiInfoINTEL).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0