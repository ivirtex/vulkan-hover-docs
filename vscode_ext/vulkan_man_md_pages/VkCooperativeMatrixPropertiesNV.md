# VkCooperativeMatrixPropertiesNV(3) Manual Page

## Name

VkCooperativeMatrixPropertiesNV - Structure specifying cooperative
matrix properties



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkCooperativeMatrixPropertiesNV` structure is defined as:

``` c
// Provided by VK_NV_cooperative_matrix
typedef struct VkCooperativeMatrixPropertiesNV {
    VkStructureType      sType;
    void*                pNext;
    uint32_t             MSize;
    uint32_t             NSize;
    uint32_t             KSize;
    VkComponentTypeNV    AType;
    VkComponentTypeNV    BType;
    VkComponentTypeNV    CType;
    VkComponentTypeNV    DType;
    VkScopeNV            scope;
} VkCooperativeMatrixPropertiesNV;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `MSize` is the number of rows in matrices A, C, and D.

- `KSize` is the number of columns in matrix A and rows in matrix B.

- `NSize` is the number of columns in matrices B, C, D.

- `AType` is the component type of matrix A, of type
  [VkComponentTypeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkComponentTypeNV.html).

- `BType` is the component type of matrix B, of type
  [VkComponentTypeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkComponentTypeNV.html).

- `CType` is the component type of matrix C, of type
  [VkComponentTypeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkComponentTypeNV.html).

- `DType` is the component type of matrix D, of type
  [VkComponentTypeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkComponentTypeNV.html).

- `scope` is the scope of all the matrix types, of type
  [VkScopeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkScopeNV.html).

## <a href="#_description" class="anchor"></a>Description

If some types are preferred over other types (e.g. for performance),
they **should** appear earlier in the list enumerated by
[vkGetPhysicalDeviceCooperativeMatrixPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceCooperativeMatrixPropertiesNV.html).

At least one entry in the list **must** have power of two values for all
of `MSize`, `KSize`, and `NSize`.

Valid Usage (Implicit)

- <a href="#VUID-VkCooperativeMatrixPropertiesNV-sType-sType"
  id="VUID-VkCooperativeMatrixPropertiesNV-sType-sType"></a>
  VUID-VkCooperativeMatrixPropertiesNV-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_COOPERATIVE_MATRIX_PROPERTIES_NV`

- <a href="#VUID-VkCooperativeMatrixPropertiesNV-pNext-pNext"
  id="VUID-VkCooperativeMatrixPropertiesNV-pNext-pNext"></a>
  VUID-VkCooperativeMatrixPropertiesNV-pNext-pNext  
  `pNext` **must** be `NULL`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_cooperative_matrix](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_cooperative_matrix.html),
[VkComponentTypeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkComponentTypeNV.html),
[VkScopeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkScopeNV.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetPhysicalDeviceCooperativeMatrixPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceCooperativeMatrixPropertiesNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkCooperativeMatrixPropertiesNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
