# VkCooperativeMatrixPropertiesKHR(3) Manual Page

## Name

VkCooperativeMatrixPropertiesKHR - Structure specifying cooperative matrix properties



## [](#_c_specification)C Specification

The `VkCooperativeMatrixPropertiesKHR` structure is defined as:

```c++
// Provided by VK_KHR_cooperative_matrix
typedef struct VkCooperativeMatrixPropertiesKHR {
    VkStructureType       sType;
    void*                 pNext;
    uint32_t              MSize;
    uint32_t              NSize;
    uint32_t              KSize;
    VkComponentTypeKHR    AType;
    VkComponentTypeKHR    BType;
    VkComponentTypeKHR    CType;
    VkComponentTypeKHR    ResultType;
    VkBool32              saturatingAccumulation;
    VkScopeKHR            scope;
} VkCooperativeMatrixPropertiesKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `MSize` is the number of rows in matrices `A`, `C`, and `Result`.
- `KSize` is the number of columns in matrix `A` and rows in matrix `B`.
- `NSize` is the number of columns in matrices `B`, `C`, `Result`.
- `AType` is the component type of matrix `A`, of type [VkComponentTypeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComponentTypeKHR.html).
- `BType` is the component type of matrix `B`, of type [VkComponentTypeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComponentTypeKHR.html).
- `CType` is the component type of matrix `C`, of type [VkComponentTypeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComponentTypeKHR.html).
- `ResultType` is the component type of matrix `Result`, of type [VkComponentTypeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComponentTypeKHR.html).
- `saturatingAccumulation` indicates whether the `SaturatingAccumulation` operand to `OpCooperativeMatrixMulAddKHR` **must** be present or not. If it is `VK_TRUE`, the `SaturatingAccumulation` operand **must** be present. If it is `VK_FALSE`, the `SaturatingAccumulation` operand **must** not be present.
- `scope` is the scope of all the matrix types, of type [VkScopeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkScopeKHR.html).

## [](#_description)Description

If some types are preferred over other types (e.g. for performance), they **should** appear earlier in the list enumerated by [vkGetPhysicalDeviceCooperativeMatrixPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceCooperativeMatrixPropertiesKHR.html).

At least one entry in the list **must** have power of two values for all of `MSize`, `KSize`, and `NSize`.

If the [`cooperativeMatrixWorkgroupScope`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-cooperativeMatrixWorkgroupScope) feature is not supported, `scope` **must** be `VK_SCOPE_SUBGROUP_KHR`.

Valid Usage (Implicit)

- [](#VUID-VkCooperativeMatrixPropertiesKHR-sType-sType)VUID-VkCooperativeMatrixPropertiesKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_COOPERATIVE_MATRIX_PROPERTIES_KHR`
- [](#VUID-VkCooperativeMatrixPropertiesKHR-pNext-pNext)VUID-VkCooperativeMatrixPropertiesKHR-pNext-pNext  
  `pNext` **must** be `NULL`

## [](#_see_also)See Also

[VK\_KHR\_cooperative\_matrix](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_cooperative_matrix.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkComponentTypeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComponentTypeKHR.html), [VkScopeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkScopeKHR.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetPhysicalDeviceCooperativeMatrixPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceCooperativeMatrixPropertiesKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkCooperativeMatrixPropertiesKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0