# VkClusterAccelerationStructureBuildTriangleClusterTemplateInfoNV(3) Manual Page

## Name

VkClusterAccelerationStructureBuildTriangleClusterTemplateInfoNV - Parameters describing build operation for a template cluster acceleration structure



## [](#_c_specification)C Specification

The [VkClusterAccelerationStructureBuildTriangleClusterTemplateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureBuildTriangleClusterTemplateInfoNV.html) structure is defined as:

```c++
// Provided by VK_NV_cluster_acceleration_structure
typedef struct VkClusterAccelerationStructureBuildTriangleClusterTemplateInfoNV {
    uint32_t                                                         clusterID;
    VkClusterAccelerationStructureClusterFlagsNV                     clusterFlags;
    uint32_t                                                         triangleCount:9;
    uint32_t                                                         vertexCount:9;
    uint32_t                                                         positionTruncateBitCount:6;
    uint32_t                                                         indexType:4;
    uint32_t                                                         opacityMicromapIndexType:4;
    VkClusterAccelerationStructureGeometryIndexAndGeometryFlagsNV    baseGeometryIndexAndGeometryFlags;
    uint16_t                                                         indexBufferStride;
    uint16_t                                                         vertexBufferStride;
    uint16_t                                                         geometryIndexAndFlagsBufferStride;
    uint16_t                                                         opacityMicromapIndexBufferStride;
    VkDeviceAddress                                                  indexBuffer;
    VkDeviceAddress                                                  vertexBuffer;
    VkDeviceAddress                                                  geometryIndexAndFlagsBuffer;
    VkDeviceAddress                                                  opacityMicromapArray;
    VkDeviceAddress                                                  opacityMicromapIndexBuffer;
    VkDeviceAddress                                                  instantiationBoundingBoxLimit;
} VkClusterAccelerationStructureBuildTriangleClusterTemplateInfoNV;
```

## [](#_members)Members

- `clusterID` is a user specified identifier assigned to this cluster template.
- `clusterFlags` is a bitmask of [VkClusterAccelerationStructureClusterFlagBitsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureClusterFlagBitsNV.html) values describing flags how the cluster template should be built.
- `triangleCount` is the number of triangles in this cluster.
- `vertexCount` is the number of unique vertices in this cluster.
- `positionTruncateBitCount` is the number of bits starting at the lowest bit (i.e. the LSBs of the mantissa), of each vertex position that will be truncated to zero to improve floating-point compression.
- `indexType` is a single [VkClusterAccelerationStructureIndexFormatFlagBitsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureIndexFormatFlagBitsNV.html) value specifying the index type in `indexBuffer`.
- `opacityMicromapIndexType` is a single [VkClusterAccelerationStructureIndexFormatFlagBitsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureIndexFormatFlagBitsNV.html) value specifying the index type in `opacityMicromapIndexBuffer`.
- `baseGeometryIndexAndGeometryFlags` is a [VkClusterAccelerationStructureGeometryIndexAndGeometryFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureGeometryIndexAndGeometryFlagsNV.html) value specifying the base geometry index and flags for all triangles in the cluster template.
- `indexBufferStride` is the stride in bytes in `indexBuffer`.
- `vertexBufferStride` is the stride in bytes in `vertexBuffer`.
- `geometryIndexAndFlagsBufferStride` is the stride in bytes in `geometryIndexAndFlagsBuffer`.
- `opacityMicromapIndexBufferStride` is the stride in bytes in `opacityMicromapIndexBuffer`.
- `indexBuffer` contains the indices of vertices in the cluster and is of type `indexType`.
- `vertexBuffer` is either `NULL` or specifies the vertex data of the triangles in the cluster template with format specified in [VkClusterAccelerationStructureTriangleClusterInputNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureTriangleClusterInputNV.html)::`vertexFormat`.
- `geometryIndexAndFlagsBuffer` is either `NULL` or an address containing strided [VkClusterAccelerationStructureGeometryIndexAndGeometryFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureGeometryIndexAndGeometryFlagsNV.html) values specifying the geometry index and flag for every triangle in the cluster.
- `opacityMicromapArray` is either `NULL` or specifies the address of a valid opacity micromap array to reference from the cluster acceleration structure. If it is `NULL`, then opacity micromaps will be disabled for this cluster acceleration structure.
- `opacityMicromapIndexBuffer` is either `NULL` or specifies the address of a strided array with size equal to the number of triangles or indices into the opacity micromap array.
- `instantiationBoundingBoxLimit` is either `NULL` or specifies the address of a bounding box within which all instantiated clusters **must** lie. The bounding box is specified by six 32-bit floating-point values in the order MinX, MinY, MinZ, MaxX, MaxY, MaxZ.

## [](#_description)Description

The C language specification does not define the ordering of bit-fields, but in practice, this structure produces the correct layout with existing compilers. The intended bit pattern is the following:

- `triangleCount`, `vertexCount`, `positionTruncateBitCount`, `indexType` and `opacityMicromapIndexType` occupy the same memory as if a single `uint32_t` was specified in their place
  
  - `triangleCount` occupies the 9 least significant bits of that memory
  - `vertexCount` occupies the next 9 least significant bits of that memory
  - `positionTruncateBitCount` occupies the next 6 least significant bits of that memory
  - `indexType` occupies the next 4 least significant bits of that memory
  - `opacityMicromapIndexType` occupies the 4 most significant bits of that memory

If a compiler produces code that diverges from that pattern, applications **must** employ another method to set values according to the correct bit pattern.

Cluster templates cannot be directly used to build bottom level acceleration structures, instead, they **must** be instantiated into [CLAS objects](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#acceleration-structure-clas-geometry).

Valid Usage

- [](#VUID-VkClusterAccelerationStructureBuildTriangleClusterTemplateInfoNV-clusterID-10497)VUID-VkClusterAccelerationStructureBuildTriangleClusterTemplateInfoNV-clusterID-10497  
  `clusterID` **must** not be 0xFFFFFFFF
- [](#VUID-VkClusterAccelerationStructureBuildTriangleClusterTemplateInfoNV-triangleCount-10498)VUID-VkClusterAccelerationStructureBuildTriangleClusterTemplateInfoNV-triangleCount-10498  
  `triangleCount` **must** be less than or equal to [VkPhysicalDeviceClusterAccelerationStructurePropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceClusterAccelerationStructurePropertiesNV.html)::`maxTrianglesPerCluster`
- [](#VUID-VkClusterAccelerationStructureBuildTriangleClusterTemplateInfoNV-vertexCount-10499)VUID-VkClusterAccelerationStructureBuildTriangleClusterTemplateInfoNV-vertexCount-10499  
  `vertexCount` **must** be less than or equal to [VkPhysicalDeviceClusterAccelerationStructurePropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceClusterAccelerationStructurePropertiesNV.html)::`maxVerticesPerCluster`
- [](#VUID-VkClusterAccelerationStructureBuildTriangleClusterTemplateInfoNV-indexType-10500)VUID-VkClusterAccelerationStructureBuildTriangleClusterTemplateInfoNV-indexType-10500  
  `indexType` **must** only have a single bit set
- [](#VUID-VkClusterAccelerationStructureBuildTriangleClusterTemplateInfoNV-opacityMicromapIndexType-10501)VUID-VkClusterAccelerationStructureBuildTriangleClusterTemplateInfoNV-opacityMicromapIndexType-10501  
  `opacityMicromapIndexType` **must** only have a single bit set
- [](#VUID-VkClusterAccelerationStructureBuildTriangleClusterTemplateInfoNV-positionTruncateBitCount-10502)VUID-VkClusterAccelerationStructureBuildTriangleClusterTemplateInfoNV-positionTruncateBitCount-10502  
  `positionTruncateBitCount` **must** be greater than or equal to [VkClusterAccelerationStructureTriangleClusterInputNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureTriangleClusterInputNV.html)::`minPositionTruncateBitCount` and less than or equal to `32`
- [](#VUID-VkClusterAccelerationStructureBuildTriangleClusterTemplateInfoNV-indexBufferStride-10503)VUID-VkClusterAccelerationStructureBuildTriangleClusterTemplateInfoNV-indexBufferStride-10503  
  `indexBufferStride` **must** be `0` or a multiple of `indexType`
- [](#VUID-VkClusterAccelerationStructureBuildTriangleClusterTemplateInfoNV-vertexBufferStride-10504)VUID-VkClusterAccelerationStructureBuildTriangleClusterTemplateInfoNV-vertexBufferStride-10504  
  `vertexBufferStride` **must** be `0` or a multiple of value specified in [VkClusterAccelerationStructureTriangleClusterInputNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureTriangleClusterInputNV.html)::`vertexFormat`
- [](#VUID-VkClusterAccelerationStructureBuildTriangleClusterTemplateInfoNV-instantiationBoundingBoxLimit-10505)VUID-VkClusterAccelerationStructureBuildTriangleClusterTemplateInfoNV-instantiationBoundingBoxLimit-10505  
  `instantiationBoundingBoxLimit` **must** be aligned to [VkPhysicalDeviceClusterAccelerationStructurePropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceClusterAccelerationStructurePropertiesNV.html)::`clusterTemplateBoundsByteAlignment`
- [](#VUID-VkClusterAccelerationStructureBuildTriangleClusterTemplateInfoNV-baseGeometryIndex-10506)VUID-VkClusterAccelerationStructureBuildTriangleClusterTemplateInfoNV-baseGeometryIndex-10506  
  The maximum geometry index after using the values in `baseGeometryIndex` and `geometryIndexBuffer` **must** be less than [VkPhysicalDeviceClusterAccelerationStructurePropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceClusterAccelerationStructurePropertiesNV.html)::`maxClusterGeometryIndex`

Valid Usage (Implicit)

- [](#VUID-VkClusterAccelerationStructureBuildTriangleClusterTemplateInfoNV-clusterFlags-parameter)VUID-VkClusterAccelerationStructureBuildTriangleClusterTemplateInfoNV-clusterFlags-parameter  
  `clusterFlags` **must** be a valid combination of [VkClusterAccelerationStructureClusterFlagBitsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureClusterFlagBitsNV.html) values

## [](#_see_also)See Also

[VK\_NV\_cluster\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_cluster_acceleration_structure.html), [VkClusterAccelerationStructureClusterFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureClusterFlagsNV.html), [VkClusterAccelerationStructureGeometryIndexAndGeometryFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureGeometryIndexAndGeometryFlagsNV.html), [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkClusterAccelerationStructureBuildTriangleClusterTemplateInfoNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0