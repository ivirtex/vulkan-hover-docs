# VkInitializePerformanceApiInfoINTEL(3) Manual Page

## Name

VkInitializePerformanceApiInfoINTEL - Structure specifying parameters of
initialize of the device



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkInitializePerformanceApiInfoINTEL` structure is defined as :

``` c
// Provided by VK_INTEL_performance_query
typedef struct VkInitializePerformanceApiInfoINTEL {
    VkStructureType    sType;
    const void*        pNext;
    void*              pUserData;
} VkInitializePerformanceApiInfoINTEL;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `pUserData` is a pointer for application data.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkInitializePerformanceApiInfoINTEL-sType-sType"
  id="VUID-VkInitializePerformanceApiInfoINTEL-sType-sType"></a>
  VUID-VkInitializePerformanceApiInfoINTEL-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_INITIALIZE_PERFORMANCE_API_INFO_INTEL`

- <a href="#VUID-VkInitializePerformanceApiInfoINTEL-pNext-pNext"
  id="VUID-VkInitializePerformanceApiInfoINTEL-pNext-pNext"></a>
  VUID-VkInitializePerformanceApiInfoINTEL-pNext-pNext  
  `pNext` **must** be `NULL`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_INTEL_performance_query](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_INTEL_performance_query.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkInitializePerformanceApiINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkInitializePerformanceApiINTEL.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkInitializePerformanceApiInfoINTEL"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
