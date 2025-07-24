# VkClusterAccelerationStructureBuildTriangleClusterInfoNV(3) Manual Page

## Name

VkClusterAccelerationStructureBuildTriangleClusterInfoNV - Parameters describing build operation for a cluster acceleration structure



## [](#_c_specification)C Specification

The [VkClusterAccelerationStructureBuildTriangleClusterInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureBuildTriangleClusterInfoNV.html) structure is defined as:

```c++
// Provided by VK_NV_cluster_acceleration_structure
typedef struct VkClusterAccelerationStructureBuildTriangleClusterInfoNV {
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
} VkClusterAccelerationStructureBuildTriangleClusterInfoNV;
```

## [](#_members)Members

- `clusterID` is a user specified identifier assigned to this cluster.
- `clusterFlags` is a bitmask of [VkClusterAccelerationStructureClusterFlagBitsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureClusterFlagBitsNV.html) values describing flags how the cluster should be built.
- `triangleCount` is the number of triangles in this cluster.
- `vertexCount` is the number of unique vertices in this cluster.
- `positionTruncateBitCount` is the number of bits starting at the lowest bit (i.e. the LSBs of the mantissa), of each vertex position that will be truncated to zero to improve floating-point compression.
- `indexType` is a single [VkClusterAccelerationStructureIndexFormatFlagBitsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureIndexFormatFlagBitsNV.html) value specifying the index type in `indexBuffer`.
- `opacityMicromapIndexType` is a single [VkClusterAccelerationStructureIndexFormatFlagBitsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureIndexFormatFlagBitsNV.html) value specifying the index type in `opacityMicromapIndexBuffer`.
- `baseGeometryIndexAndGeometryFlags` is a [VkClusterAccelerationStructureGeometryIndexAndGeometryFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureGeometryIndexAndGeometryFlagsNV.html) value specifying the base geometry index and flags for all triangles in the cluster.
- `indexBufferStride` is the stride in bytes in `indexBuffer` with `0` meaning the values are tightly-packed.
- `vertexBufferStride` is the stride in bytes in `vertexBuffer` with `0` meaning the values are tightly-packed.
- `geometryIndexAndFlagsBufferStride` is the stride in bytes in `geometryIndexAndFlagsBuffer` with `0` meaning the values are tightly-packed.
- `opacityMicromapIndexBufferStride` is the stride in bytes in `opacityMicromapIndexBuffer` with `0` meaning the values are tightly-packed.
- `indexBuffer` is a device address containing the indices of the vertices in the cluster and are of type `indexType`.
- `vertexBuffer` is a device address containing the vertex data of the triangles in the cluster with format specified in [VkClusterAccelerationStructureTriangleClusterInputNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureTriangleClusterInputNV.html)::`vertexFormat`.
- []()`geometryIndexAndFlagsBuffer` is either `0` or an address containing strided [VkClusterAccelerationStructureGeometryIndexAndGeometryFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureGeometryIndexAndGeometryFlagsNV.html) values specifying the geometry index and flag for every triangle in the cluster.
- `opacityMicromapArray` is either `0` or specifies the address of a valid opacity micromap array to reference from the cluster acceleration structure. If it is `0`, then opacity micromaps will be disabled for this cluster acceleration structure.
- `opacityMicromapIndexBuffer` is either `0` or specifies the address of a strided array with size equal to the number of triangles or indices into the opacity micromap array. If `opacityMicromapIndexBuffer` is `0` then the index used is the index of the triangle in the geometry.

## [](#_description)Description

The C language specification does not define the ordering of bit-fields, but in practice, this structure produces the correct layout with existing compilers. The intended bit pattern is the following:

- `triangleCount`, `vertexCount`, `positionTruncateBitCount`, `indexType` and `opacityMicromapIndexType` occupy the same memory as if a single `uint32_t` was specified in their place
  
  - `triangleCount` occupies the 9 least significant bits of that memory
  - `vertexCount` occupies the next 9 least significant bits of that memory
  - `positionTruncateBitCount` occupies the next 6 least significant bits of that memory
  - `indexType` occupies the next 4 least significant bits of that memory
  - `opacityMicromapIndexType` occupies the 4 most significant bits of that memory

If a compiler produces code that diverges from that pattern, applications **must** employ another method to set values according to the correct bit pattern.

Valid Usage

- [](#VUID-VkClusterAccelerationStructureBuildTriangleClusterInfoNV-clusterID-10488)VUID-VkClusterAccelerationStructureBuildTriangleClusterInfoNV-clusterID-10488  
  `clusterID` **must** not be 0xFFFFFFFF
- [](#VUID-VkClusterAccelerationStructureBuildTriangleClusterInfoNV-triangleCount-10489)VUID-VkClusterAccelerationStructureBuildTriangleClusterInfoNV-triangleCount-10489  
  `triangleCount` **must** be less than or equal to [VkPhysicalDeviceClusterAccelerationStructurePropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceClusterAccelerationStructurePropertiesNV.html)::`maxTrianglesPerCluster`
- [](#VUID-VkClusterAccelerationStructureBuildTriangleClusterInfoNV-vertexCount-10490)VUID-VkClusterAccelerationStructureBuildTriangleClusterInfoNV-vertexCount-10490  
  `vertexCount` **must** be less than or equal to [VkPhysicalDeviceClusterAccelerationStructurePropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceClusterAccelerationStructurePropertiesNV.html)::`maxVerticesPerCluster`
- [](#VUID-VkClusterAccelerationStructureBuildTriangleClusterInfoNV-indexType-10491)VUID-VkClusterAccelerationStructureBuildTriangleClusterInfoNV-indexType-10491  
  `indexType` **must** only have a single bit set
- [](#VUID-VkClusterAccelerationStructureBuildTriangleClusterInfoNV-opacityMicromapIndexType-10492)VUID-VkClusterAccelerationStructureBuildTriangleClusterInfoNV-opacityMicromapIndexType-10492  
  `opacityMicromapIndexType` **must** only have a single bit set
- [](#VUID-VkClusterAccelerationStructureBuildTriangleClusterInfoNV-positionTruncateBitCount-10493)VUID-VkClusterAccelerationStructureBuildTriangleClusterInfoNV-positionTruncateBitCount-10493  
  `positionTruncateBitCount` **must** be greater than or equal to [VkClusterAccelerationStructureTriangleClusterInputNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureTriangleClusterInputNV.html)::`minPositionTruncateBitCount` and less than or equal to `32`
- [](#VUID-VkClusterAccelerationStructureBuildTriangleClusterInfoNV-indexBufferStride-10494)VUID-VkClusterAccelerationStructureBuildTriangleClusterInfoNV-indexBufferStride-10494  
  `indexBufferStride` **must** be `0` or a multiple of `indexType`
- [](#VUID-VkClusterAccelerationStructureBuildTriangleClusterInfoNV-vertexBufferStride-10495)VUID-VkClusterAccelerationStructureBuildTriangleClusterInfoNV-vertexBufferStride-10495  
  `vertexBufferStride` **must** be `0` or a multiple of value specified in [VkClusterAccelerationStructureTriangleClusterInputNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureTriangleClusterInputNV.html)::`vertexFormat`
- [](#VUID-VkClusterAccelerationStructureBuildTriangleClusterInfoNV-baseGeometryIndex-10496)VUID-VkClusterAccelerationStructureBuildTriangleClusterInfoNV-baseGeometryIndex-10496  
  The maximum geometry index after using the values in `baseGeometryIndex` and `geometryIndexBuffer` **must** be less than [VkPhysicalDeviceClusterAccelerationStructurePropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceClusterAccelerationStructurePropertiesNV.html)::`maxClusterGeometryIndex`
- [](#VUID-VkClusterAccelerationStructureBuildTriangleClusterInfoNV-opacityMicromapArray-10881)VUID-VkClusterAccelerationStructureBuildTriangleClusterInfoNV-opacityMicromapArray-10881  
  If `opacityMicromapArray` is not `0`, then the cluster acceleration structure **must** have been built with `VK_BUILD_ACCELERATION_STRUCTURE_ALLOW_CLUSTER_OPACITY_MICROMAPS_BIT_NV` flag set in [VkClusterAccelerationStructureInputInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureInputInfoNV.html)::`flags`

Valid Usage (Implicit)

- [](#VUID-VkClusterAccelerationStructureBuildTriangleClusterInfoNV-clusterFlags-parameter)VUID-VkClusterAccelerationStructureBuildTriangleClusterInfoNV-clusterFlags-parameter  
  `clusterFlags` **must** be a valid combination of [VkClusterAccelerationStructureClusterFlagBitsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureClusterFlagBitsNV.html) values

## [](#_see_also)See Also

[VK\_NV\_cluster\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_cluster_acceleration_structure.html), [VkClusterAccelerationStructureClusterFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureClusterFlagsNV.html), [VkClusterAccelerationStructureGeometryIndexAndGeometryFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureGeometryIndexAndGeometryFlagsNV.html), [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkClusterAccelerationStructureBuildTriangleClusterInfoNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0