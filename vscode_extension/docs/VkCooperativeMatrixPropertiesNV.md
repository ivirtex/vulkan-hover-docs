# VkCooperativeMatrixPropertiesNV(3) Manual Page

## Name

VkCooperativeMatrixPropertiesNV - Structure specifying cooperative matrix properties



## [](#_c_specification)C Specification

The `VkCooperativeMatrixPropertiesNV` structure is defined as:

```c++
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

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `MSize` is the number of rows in matrices A, C, and D.
- `KSize` is the number of columns in matrix A and rows in matrix B.
- `NSize` is the number of columns in matrices B, C, D.
- `AType` is the component type of matrix A, of type [VkComponentTypeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComponentTypeNV.html).
- `BType` is the component type of matrix B, of type [VkComponentTypeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComponentTypeNV.html).
- `CType` is the component type of matrix C, of type [VkComponentTypeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComponentTypeNV.html).
- `DType` is the component type of matrix D, of type [VkComponentTypeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComponentTypeNV.html).
- `scope` is the scope of all the matrix types, of type [VkScopeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkScopeNV.html).

## [](#_description)Description

If some types are preferred over other types (e.g. for performance), they **should** appear earlier in the list enumerated by [vkGetPhysicalDeviceCooperativeMatrixPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceCooperativeMatrixPropertiesNV.html).

At least one entry in the list **must** have power of two values for all of `MSize`, `KSize`, and `NSize`.

Valid Usage (Implicit)

- [](#VUID-VkCooperativeMatrixPropertiesNV-sType-sType)VUID-VkCooperativeMatrixPropertiesNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_COOPERATIVE_MATRIX_PROPERTIES_NV`
- [](#VUID-VkCooperativeMatrixPropertiesNV-pNext-pNext)VUID-VkCooperativeMatrixPropertiesNV-pNext-pNext  
  `pNext` **must** be `NULL`

## [](#_see_also)See Also

[VK\_NV\_cooperative\_matrix](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_cooperative_matrix.html), [VkComponentTypeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComponentTypeKHR.html), [VkScopeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkScopeKHR.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetPhysicalDeviceCooperativeMatrixPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceCooperativeMatrixPropertiesNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkCooperativeMatrixPropertiesNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0