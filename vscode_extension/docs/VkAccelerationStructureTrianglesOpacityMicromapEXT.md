# VkAccelerationStructureTrianglesOpacityMicromapEXT(3) Manual Page

## Name

VkAccelerationStructureTrianglesOpacityMicromapEXT - Structure specifying an opacity micromap in a bottom-level acceleration structure



## [](#_c_specification)C Specification

The `VkAccelerationStructureTrianglesOpacityMicromapEXT` structure is defined as:

```c++
// Provided by VK_EXT_opacity_micromap
typedef struct VkAccelerationStructureTrianglesOpacityMicromapEXT {
    VkStructureType                     sType;
    void*                               pNext;
    VkIndexType                         indexType;
    VkDeviceOrHostAddressConstKHR       indexBuffer;
    VkDeviceSize                        indexStride;
    uint32_t                            baseTriangle;
    uint32_t                            usageCountsCount;
    const VkMicromapUsageEXT*           pUsageCounts;
    const VkMicromapUsageEXT* const*    ppUsageCounts;
    VkMicromapEXT                       micromap;
} VkAccelerationStructureTrianglesOpacityMicromapEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `indexType` is the type of triangle indices used when indexing this micromap
- `indexBuffer` is the address containing the triangle indices
- `indexStride` is the byte stride between triangle indices
- `baseTriangle` is the base value added to the non-negative triangle indices
- `usageCountsCount` specifies the number of usage counts structures that will be used to determine the size of this micromap.
- `pUsageCounts` is a pointer to an array of [VkMicromapUsageEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapUsageEXT.html) structures.
- `ppUsageCounts` is a pointer to an array of pointers to [VkMicromapUsageEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapUsageEXT.html) structures.
- `micromap` is the handle to the micromap object to include in this geometry

## [](#_description)Description

If `VkAccelerationStructureTrianglesOpacityMicromapEXT` is included in the `pNext` chain of a [VkAccelerationStructureDenseGeometryFormatTrianglesDataAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureDenseGeometryFormatTrianglesDataAMDX.html) or [VkAccelerationStructureGeometryTrianglesDataKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureGeometryTrianglesDataKHR.html) structure, that geometry will reference that micromap.

For each triangle in the geometry, the acceleration structure build fetches an index from `indexBuffer` using `indexType` and `indexStride` if present. If `indexBuffer` is `NULL` then the index used is the index of the triangle in the geometry.

If that value is the unsigned cast of one of the values from [VkOpacityMicromapSpecialIndexEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpacityMicromapSpecialIndexEXT.html) then that triangle behaves as described for that special value in [Ray Opacity Micromap](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#ray-opacity-micromap).

Otherwise that triangle uses the opacity micromap information from `micromap` at that index plus `baseTriangle`.

Only one of `pUsageCounts` or `ppUsageCounts` **can** be a valid pointer, the other **must** be `NULL`. The elements of the non-`NULL` array describe the total count used to build this geometry. For a given `format` and `subdivisionLevel` the number of triangles in this geometry matching those values after indirection and special index handling **must** be equal to the sum of matching `count` provided.

If `micromap` is [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), then every value read from `indexBuffer` **must** be one of the values in `VkOpacityMicromapSpecialIndexEXT`.

Valid Usage

- [](#VUID-VkAccelerationStructureTrianglesOpacityMicromapEXT-pUsageCounts-07335)VUID-VkAccelerationStructureTrianglesOpacityMicromapEXT-pUsageCounts-07335  
  Only one of `pUsageCounts` or `ppUsageCounts` **can** be a valid pointer, the other **must** be `NULL`
- [](#VUID-VkAccelerationStructureTrianglesOpacityMicromapEXT-indexType-10719)VUID-VkAccelerationStructureTrianglesOpacityMicromapEXT-indexType-10719  
  `indexType` **must** be `VK_INDEX_TYPE_UINT16`, `VK_INDEX_TYPE_UINT32`, or `VK_INDEX_TYPE_NONE_KHR`
- [](#VUID-VkAccelerationStructureTrianglesOpacityMicromapEXT-indexType-10722)VUID-VkAccelerationStructureTrianglesOpacityMicromapEXT-indexType-10722  
  If `indexType` is not `VK_INDEX_TYPE_NONE_KHR`, then `indexStride` **must** be a multiple of the size in bytes of `indexType`
- [](#VUID-VkAccelerationStructureTrianglesOpacityMicromapEXT-indexType-10723)VUID-VkAccelerationStructureTrianglesOpacityMicromapEXT-indexType-10723  
  If `indexType` is not `VK_INDEX_TYPE_NONE_KHR`, then `indexStride` **must** be less than or equal to 232-1

Valid Usage (Implicit)

- [](#VUID-VkAccelerationStructureTrianglesOpacityMicromapEXT-sType-sType)VUID-VkAccelerationStructureTrianglesOpacityMicromapEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_TRIANGLES_OPACITY_MICROMAP_EXT`
- [](#VUID-VkAccelerationStructureTrianglesOpacityMicromapEXT-indexType-parameter)VUID-VkAccelerationStructureTrianglesOpacityMicromapEXT-indexType-parameter  
  `indexType` **must** be a valid [VkIndexType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndexType.html) value
- [](#VUID-VkAccelerationStructureTrianglesOpacityMicromapEXT-pUsageCounts-parameter)VUID-VkAccelerationStructureTrianglesOpacityMicromapEXT-pUsageCounts-parameter  
  If `usageCountsCount` is not `0`, and `pUsageCounts` is not `NULL`, `pUsageCounts` **must** be a valid pointer to an array of `usageCountsCount` [VkMicromapUsageEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapUsageEXT.html) structures
- [](#VUID-VkAccelerationStructureTrianglesOpacityMicromapEXT-ppUsageCounts-parameter)VUID-VkAccelerationStructureTrianglesOpacityMicromapEXT-ppUsageCounts-parameter  
  If `usageCountsCount` is not `0`, and `ppUsageCounts` is not `NULL`, `ppUsageCounts` **must** be a valid pointer to an array of `usageCountsCount` valid pointers to [VkMicromapUsageEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapUsageEXT.html) structures
- [](#VUID-VkAccelerationStructureTrianglesOpacityMicromapEXT-micromap-parameter)VUID-VkAccelerationStructureTrianglesOpacityMicromapEXT-micromap-parameter  
  If `micromap` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `micromap` **must** be a valid [VkMicromapEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapEXT.html) handle

## [](#_see_also)See Also

[VK\_EXT\_opacity\_micromap](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_opacity_micromap.html), [VkDeviceOrHostAddressConstKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceOrHostAddressConstKHR.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html), [VkIndexType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndexType.html), [VkMicromapEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapEXT.html), [VkMicromapUsageEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapUsageEXT.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkAccelerationStructureTrianglesOpacityMicromapEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0