# VkCooperativeMatrixFlexibleDimensionsPropertiesNV(3) Manual Page

## Name

VkCooperativeMatrixFlexibleDimensionsPropertiesNV - Structure specifying cooperative matrix properties



## [](#_c_specification)C Specification

The `VkCooperativeMatrixFlexibleDimensionsPropertiesNV` structure is defined as:

```c++
// Provided by VK_NV_cooperative_matrix2
typedef struct VkCooperativeMatrixFlexibleDimensionsPropertiesNV {
    VkStructureType       sType;
    void*                 pNext;
    uint32_t              MGranularity;
    uint32_t              NGranularity;
    uint32_t              KGranularity;
    VkComponentTypeKHR    AType;
    VkComponentTypeKHR    BType;
    VkComponentTypeKHR    CType;
    VkComponentTypeKHR    ResultType;
    VkBool32              saturatingAccumulation;
    VkScopeKHR            scope;
    uint32_t              workgroupInvocations;
} VkCooperativeMatrixFlexibleDimensionsPropertiesNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `MGranularity` is the granularity of the number of rows in matrices `A`, `C`, and `Result`. The rows **must** be an integer multiple of this value.
- `KGranularity` is the granularity of columns in matrix `A` and rows in matrix `B`. The columns/rows **must** be an integer multiple of this value.
- `NGranularity` is the granularity of columns in matrices `B`, `C`, `Result`. The columns **must** be an integer multiple of this value.
- `AType` is the component type of matrix `A`, of type [VkComponentTypeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComponentTypeKHR.html).
- `BType` is the component type of matrix `B`, of type [VkComponentTypeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComponentTypeKHR.html).
- `CType` is the component type of matrix `C`, of type [VkComponentTypeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComponentTypeKHR.html).
- `ResultType` is the component type of matrix `Result`, of type [VkComponentTypeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComponentTypeKHR.html).
- `saturatingAccumulation` indicates whether the `SaturatingAccumulation` operand to `OpCooperativeMatrixMulAddKHR` **must** be present or not. If it is `VK_TRUE`, the `SaturatingAccumulation` operand **must** be present. If it is `VK_FALSE`, the `SaturatingAccumulation` operand **must** not be present.
- `scope` is the scope of all the matrix types, of type [VkScopeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkScopeKHR.html).
- `workgroupInvocations` is the number of invocations in the local workgroup when this combination of values is supported.

## [](#_description)Description

Rather than explicitly enumerating a list of supported sizes, `VkCooperativeMatrixFlexibleDimensionsPropertiesNV` advertises size granularities, where the matrix **must** be a multiple of the advertised size. The M and K granularities apply to rows and columns of matrices with `Use` of `MatrixA`, K, and N apply to rows and columns of matrices with `Use` of `MatrixB`, M, and N apply to rows and columns of matrices with `Use` of `MatrixAccumulator`.

For a given type combination, if multiple workgroup sizes are supported there **may** be multiple `VkCooperativeMatrixFlexibleDimensionsPropertiesNV` structures with different granularities.

All granularity values **must** be powers of two.

Note

Different A/B types may require different granularities but share the same accumulator type. In such a case, the supported granularity for a matrix with the accumulator type would be the smallest advertised granularity.

Valid Usage (Implicit)

- [](#VUID-VkCooperativeMatrixFlexibleDimensionsPropertiesNV-sType-sType)VUID-VkCooperativeMatrixFlexibleDimensionsPropertiesNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_COOPERATIVE_MATRIX_FLEXIBLE_DIMENSIONS_PROPERTIES_NV`
- [](#VUID-VkCooperativeMatrixFlexibleDimensionsPropertiesNV-pNext-pNext)VUID-VkCooperativeMatrixFlexibleDimensionsPropertiesNV-pNext-pNext  
  `pNext` **must** be `NULL`

## [](#_see_also)See Also

[VK\_NV\_cooperative\_matrix2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_cooperative_matrix2.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkComponentTypeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComponentTypeKHR.html), [VkScopeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkScopeKHR.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetPhysicalDeviceCooperativeMatrixFlexibleDimensionsPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceCooperativeMatrixFlexibleDimensionsPropertiesNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkCooperativeMatrixFlexibleDimensionsPropertiesNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0