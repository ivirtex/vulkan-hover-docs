# VkAccelerationStructureTrianglesDisplacementMicromapNV(3) Manual Page

## Name

VkAccelerationStructureTrianglesDisplacementMicromapNV - Structure specifying a displacement micromap in a bottom-level acceleration structure



## [](#_c_specification)C Specification

The `VkAccelerationStructureTrianglesDisplacementMicromapNV` structure is defined as:

```c++
// Provided by VK_NV_displacement_micromap
typedef struct VkAccelerationStructureTrianglesDisplacementMicromapNV {
    VkStructureType                     sType;
    void*                               pNext;
    VkFormat                            displacementBiasAndScaleFormat;
    VkFormat                            displacementVectorFormat;
    VkDeviceOrHostAddressConstKHR       displacementBiasAndScaleBuffer;
    VkDeviceSize                        displacementBiasAndScaleStride;
    VkDeviceOrHostAddressConstKHR       displacementVectorBuffer;
    VkDeviceSize                        displacementVectorStride;
    VkDeviceOrHostAddressConstKHR       displacedMicromapPrimitiveFlags;
    VkDeviceSize                        displacedMicromapPrimitiveFlagsStride;
    VkIndexType                         indexType;
    VkDeviceOrHostAddressConstKHR       indexBuffer;
    VkDeviceSize                        indexStride;
    uint32_t                            baseTriangle;
    uint32_t                            usageCountsCount;
    const VkMicromapUsageEXT*           pUsageCounts;
    const VkMicromapUsageEXT* const*    ppUsageCounts;
    VkMicromapEXT                       micromap;
} VkAccelerationStructureTrianglesDisplacementMicromapNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `displacementBiasAndScaleFormat` is the format of displacement bias and scale used in this displacement micromap reference.
- `displacementVectorFormat` is the format of displacement vector used in this displacement micromap reference.
- `displacementBiasAndScaleBuffer` is the address containing the bias and scale.
- `displacementBiasAndScaleStride` is the byte stride between bias and scale values.
- `displacementVectorBuffer` is the address containing the displacement vector values.
- `displacementVectorStride` is the byte stride between displacement vector values.
- `displacedMicromapPrimitiveFlags` is the address containing the primitive flags.
- `displacedMicromapPrimitiveFlagsStride` is the byte stride between primitive flag values.
- `indexType` is the type of triangle indices used when indexing this micromap.
- `indexBuffer` is the address containing the triangle indices.
- `indexStride` is the byte stride between triangle indices.
- `baseTriangle` is the base value added to the non-negative triangle indices.
- `usageCountsCount` specifies the number of usage counts structures that will be used to determine the size of this micromap.
- `pUsageCounts` is a pointer to an array of [VkMicromapUsageEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapUsageEXT.html) structures.
- `ppUsageCounts` is a pointer to an array of pointers to [VkMicromapUsageEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapUsageEXT.html) structures.
- `micromap` is the handle to the micromap object to include in this geometry.

## [](#_description)Description

If `VkAccelerationStructureTrianglesDisplacementMicromapNV` is included in the `pNext` chain of a [VkAccelerationStructureGeometryTrianglesDataKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureGeometryTrianglesDataKHR.html) structure, that geometry will reference that micromap.

For each triangle in the geometry, the acceleration structure build fetches an index from `indexBuffer` using `indexType` and `indexStride`. That triangle uses the displacement micromap information from `micromap` at that index plus `baseTriangle`.

Only one of `pUsageCounts` or `ppUsageCounts` **can** be a valid pointer, the other **must** be `NULL`. The elements of the non-`NULL` array describe the total count used to build this geometry. For a given `format` and `subdivisionLevel` the number of triangles in this geometry matching those values after indirection **must** be equal to the sum of matching `count` provided.

Valid Usage

- [](#VUID-VkAccelerationStructureTrianglesDisplacementMicromapNV-displacementBiasAndScaleFormat-09501)VUID-VkAccelerationStructureTrianglesDisplacementMicromapNV-displacementBiasAndScaleFormat-09501  
  `displacementBiasAndScaleFormat` **must** not be `VK_FORMAT_UNDEFINED`
- [](#VUID-VkAccelerationStructureTrianglesDisplacementMicromapNV-displacementVectorFormat-09502)VUID-VkAccelerationStructureTrianglesDisplacementMicromapNV-displacementVectorFormat-09502  
  `displacementVectorFormat` **must** not be `VK_FORMAT_UNDEFINED`
- [](#VUID-VkAccelerationStructureTrianglesDisplacementMicromapNV-pUsageCounts-07992)VUID-VkAccelerationStructureTrianglesDisplacementMicromapNV-pUsageCounts-07992  
  Only one of `pUsageCounts` or `ppUsageCounts` **can** be a valid pointer, the other **must** be `NULL`

Valid Usage (Implicit)

- [](#VUID-VkAccelerationStructureTrianglesDisplacementMicromapNV-sType-sType)VUID-VkAccelerationStructureTrianglesDisplacementMicromapNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_TRIANGLES_DISPLACEMENT_MICROMAP_NV`
- [](#VUID-VkAccelerationStructureTrianglesDisplacementMicromapNV-displacementBiasAndScaleFormat-parameter)VUID-VkAccelerationStructureTrianglesDisplacementMicromapNV-displacementBiasAndScaleFormat-parameter  
  `displacementBiasAndScaleFormat` **must** be a valid [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) value
- [](#VUID-VkAccelerationStructureTrianglesDisplacementMicromapNV-displacementVectorFormat-parameter)VUID-VkAccelerationStructureTrianglesDisplacementMicromapNV-displacementVectorFormat-parameter  
  `displacementVectorFormat` **must** be a valid [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) value
- [](#VUID-VkAccelerationStructureTrianglesDisplacementMicromapNV-indexType-parameter)VUID-VkAccelerationStructureTrianglesDisplacementMicromapNV-indexType-parameter  
  `indexType` **must** be a valid [VkIndexType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndexType.html) value
- [](#VUID-VkAccelerationStructureTrianglesDisplacementMicromapNV-pUsageCounts-parameter)VUID-VkAccelerationStructureTrianglesDisplacementMicromapNV-pUsageCounts-parameter  
  If `usageCountsCount` is not `0`, and `pUsageCounts` is not `NULL`, `pUsageCounts` **must** be a valid pointer to an array of `usageCountsCount` [VkMicromapUsageEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapUsageEXT.html) structures
- [](#VUID-VkAccelerationStructureTrianglesDisplacementMicromapNV-ppUsageCounts-parameter)VUID-VkAccelerationStructureTrianglesDisplacementMicromapNV-ppUsageCounts-parameter  
  If `usageCountsCount` is not `0`, and `ppUsageCounts` is not `NULL`, `ppUsageCounts` **must** be a valid pointer to an array of `usageCountsCount` valid pointers to [VkMicromapUsageEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapUsageEXT.html) structures
- [](#VUID-VkAccelerationStructureTrianglesDisplacementMicromapNV-micromap-parameter)VUID-VkAccelerationStructureTrianglesDisplacementMicromapNV-micromap-parameter  
  If `micromap` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `micromap` **must** be a valid [VkMicromapEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapEXT.html) handle

## [](#_see_also)See Also

[VK\_NV\_displacement\_micromap](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_displacement_micromap.html), [VkDeviceOrHostAddressConstKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceOrHostAddressConstKHR.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html), [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html), [VkIndexType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndexType.html), [VkMicromapEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapEXT.html), [VkMicromapUsageEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapUsageEXT.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkAccelerationStructureTrianglesDisplacementMicromapNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0