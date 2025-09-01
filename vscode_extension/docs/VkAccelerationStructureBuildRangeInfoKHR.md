# VkAccelerationStructureBuildRangeInfoKHR(3) Manual Page

## Name

VkAccelerationStructureBuildRangeInfoKHR - Structure specifying build offsets and counts for acceleration structure builds



## [](#_c_specification)C Specification

`VkAccelerationStructureBuildRangeInfoKHR` is defined as:

```c++
// Provided by VK_KHR_acceleration_structure
typedef struct VkAccelerationStructureBuildRangeInfoKHR {
    uint32_t    primitiveCount;
    uint32_t    primitiveOffset;
    uint32_t    firstVertex;
    uint32_t    transformOffset;
} VkAccelerationStructureBuildRangeInfoKHR;
```

## [](#_members)Members

- `primitiveCount` defines the number of primitives for a corresponding acceleration structure geometry.
- `primitiveOffset` defines an offset in bytes into the memory where primitive data is defined.
- `firstVertex` is the index of the first vertex to build from for triangle geometry.
- `transformOffset` defines an offset in bytes into the memory where a transform matrix is defined.

## [](#_description)Description

The primitive count and primitive offset are interpreted differently depending on the [VkGeometryTypeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeometryTypeKHR.html) used:

- For geometries of type `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, `primitiveCount` is the number of triangles to be built, where each triangle is treated as 3 vertices.
  
  - If the geometry uses indices, `primitiveCount` × 3 indices are consumed from [VkAccelerationStructureGeometryTrianglesDataKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureGeometryTrianglesDataKHR.html)::`indexData`, starting at an offset of `primitiveOffset`. The value of `firstVertex` is added to the index values before fetching vertices.
  - If the geometry does not use indices, `primitiveCount` × 3 vertices are consumed from [VkAccelerationStructureGeometryTrianglesDataKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureGeometryTrianglesDataKHR.html)::`vertexData`, starting at an offset of `primitiveOffset` + [VkAccelerationStructureGeometryTrianglesDataKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureGeometryTrianglesDataKHR.html)::`vertexStride` × `firstVertex`.
  - If [VkAccelerationStructureGeometryTrianglesDataKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureGeometryTrianglesDataKHR.html)::`transformData` is not `NULL`, a single [VkTransformMatrixKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTransformMatrixKHR.html) structure is consumed from [VkAccelerationStructureGeometryTrianglesDataKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureGeometryTrianglesDataKHR.html)::`transformData`, at an offset of `transformOffset`. This matrix describes a transformation from the space in which the vertices for all triangles in this geometry are described to the space in which the acceleration structure is defined.
- For geometries of type `VK_GEOMETRY_TYPE_AABBS_KHR`, `primitiveCount` is the number of axis-aligned bounding boxes. `primitiveCount` [VkAabbPositionsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAabbPositionsKHR.html) structures are consumed from [VkAccelerationStructureGeometryAabbsDataKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureGeometryAabbsDataKHR.html)::`data`, starting at an offset of `primitiveOffset`.
- For geometries of type `VK_GEOMETRY_TYPE_SPHERES_NV`, `primitiveCount` is the number of spheres to be built, where each sphere is treated as 1 vertex.
  
  - If the geometry uses indices, `primitiveCount` indices are consumed from [VkAccelerationStructureGeometrySpheresDataNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureGeometrySpheresDataNV.html)::`indexData`, starting at an offset of `primitiveOffset`. The value of `firstVertex` is added to the index values before fetching vertices and radii.
  - If the geometry does not use indices, `primitiveCount` vertices and radii are consumed from [VkAccelerationStructureGeometrySpheresDataNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureGeometrySpheresDataNV.html)::`vertexData`, starting at an offset of `primitiveOffset` + [VkAccelerationStructureGeometrySpheresDataNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureGeometrySpheresDataNV.html)::`vertexStride` × `firstVertex` and [VkAccelerationStructureGeometrySpheresDataNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureGeometrySpheresDataNV.html)::`radiusData`, starting at an offset of `primitiveOffset` + [VkAccelerationStructureGeometrySpheresDataNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureGeometrySpheresDataNV.html)::`radiusStride` × `firstVertex` respectively.
- For geometries of type `VK_GEOMETRY_TYPE_LINEAR_SWEPT_SPHERES_NV`, `primitiveCount` is the number of LSS primitives to be built, where each LSS primitive is treated as 2 vertices.
  
  - If the geometry uses indices, `primitiveCount` × 2 indices are consumed from [VkAccelerationStructureGeometryLinearSweptSpheresDataNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureGeometryLinearSweptSpheresDataNV.html)::`indexData`, starting at an offset of `primitiveOffset`. The value of `firstVertex` is added to the index values before fetching vertices and radii.
  - If the geometry does not use indices, `primitiveCount` × 2 vertices and radii are consumed from [VkAccelerationStructureGeometryLinearSweptSpheresDataNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureGeometryLinearSweptSpheresDataNV.html)::`vertexData`, starting at an offset of `primitiveOffset` + [VkAccelerationStructureGeometryLinearSweptSpheresDataNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureGeometryLinearSweptSpheresDataNV.html)::`vertexStride` × `firstVertex` and [VkAccelerationStructureGeometryLinearSweptSpheresDataNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureGeometryLinearSweptSpheresDataNV.html)::`radiusData`, starting at an offset of `primitiveOffset` + [VkAccelerationStructureGeometryLinearSweptSpheresDataNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureGeometryLinearSweptSpheresDataNV.html)::`radiusStride` × `firstVertex` respectively.
- For geometries of type `VK_GEOMETRY_TYPE_INSTANCES_KHR`, `primitiveCount` is the number of acceleration structures. `primitiveCount` [VkAccelerationStructureInstanceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureInstanceKHR.html) or [VkAccelerationStructureMotionInstanceNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureMotionInstanceNV.html) structures are consumed from [VkAccelerationStructureGeometryInstancesDataKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureGeometryInstancesDataKHR.html)::`data`, starting at an offset of `primitiveOffset`.

Valid Usage

- [](#VUID-VkAccelerationStructureBuildRangeInfoKHR-vertexData-10418)VUID-VkAccelerationStructureBuildRangeInfoKHR-vertexData-10418  
  The number of vertices consumed from [VkAccelerationStructureGeometryTrianglesDataKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureGeometryTrianglesDataKHR.html)::`vertexData` **must** be less than or equal to [VkAccelerationStructureGeometryTrianglesDataKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureGeometryTrianglesDataKHR.html)::`maxVertex` + 1
- [](#VUID-VkAccelerationStructureBuildRangeInfoKHR-primitiveOffset-03656)VUID-VkAccelerationStructureBuildRangeInfoKHR-primitiveOffset-03656  
  For geometries of type `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, if the geometry uses indices, the offset `primitiveOffset` from [VkAccelerationStructureGeometryTrianglesDataKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureGeometryTrianglesDataKHR.html)::`indexData` **must** be a multiple of the element size of [VkAccelerationStructureGeometryTrianglesDataKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureGeometryTrianglesDataKHR.html)::`indexType`
- [](#VUID-VkAccelerationStructureBuildRangeInfoKHR-primitiveOffset-03657)VUID-VkAccelerationStructureBuildRangeInfoKHR-primitiveOffset-03657  
  For geometries of type `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, if the geometry does not use indices, the offset `primitiveOffset` from [VkAccelerationStructureGeometryTrianglesDataKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureGeometryTrianglesDataKHR.html)::`vertexData` **must** be a multiple of:
  
  - the [size of the format](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats) specified in [VkAccelerationStructureGeometryTrianglesDataKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureGeometryTrianglesDataKHR.html)::`vertexFormat`, if that format is a [packed format](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats-packed)
  - the [component size](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats) of the [VkAccelerationStructureGeometryTrianglesDataKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureGeometryTrianglesDataKHR.html)::`vertexFormat`, if that format is not a [packed format](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats-packed)
- [](#VUID-VkAccelerationStructureBuildRangeInfoKHR-maxVertex-10774)VUID-VkAccelerationStructureBuildRangeInfoKHR-maxVertex-10774  
  For geometries of type `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, if the geometry uses indices, then [VkAccelerationStructureGeometryTrianglesDataKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureGeometryTrianglesDataKHR.html)::`maxVertex` **must** be greater than or equal to `firstVertex` plus the maximum index value found in the [VkAccelerationStructureGeometryTrianglesDataKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureGeometryTrianglesDataKHR.html)::`indexData` in the range [`primitiveOffset`, `primitiveOffset`  
  `primitiveCount` x 3]
- [](#VUID-VkAccelerationStructureBuildRangeInfoKHR-None-10775)VUID-VkAccelerationStructureBuildRangeInfoKHR-None-10775  
  For geometries of type `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, if the geometry does not use indices, then [VkAccelerationStructureGeometryTrianglesDataKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureGeometryTrianglesDataKHR.html)::maxVertex **must** be greater than or equal to firstVertex + primitiveCount x 3 - 1
- [](#VUID-VkAccelerationStructureBuildRangeInfoKHR-transformOffset-03658)VUID-VkAccelerationStructureBuildRangeInfoKHR-transformOffset-03658  
  For geometries of type `VK_GEOMETRY_TYPE_TRIANGLES_KHR`, the offset `transformOffset` from [VkAccelerationStructureGeometryTrianglesDataKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureGeometryTrianglesDataKHR.html)::`transformData` **must** be a multiple of 16
- [](#VUID-VkAccelerationStructureBuildRangeInfoKHR-primitiveOffset-03659)VUID-VkAccelerationStructureBuildRangeInfoKHR-primitiveOffset-03659  
  For geometries of type `VK_GEOMETRY_TYPE_AABBS_KHR`, the offset `primitiveOffset` from [VkAccelerationStructureGeometryAabbsDataKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureGeometryAabbsDataKHR.html)::`data` **must** be a multiple of 8
- [](#VUID-VkAccelerationStructureBuildRangeInfoKHR-primitiveOffset-03660)VUID-VkAccelerationStructureBuildRangeInfoKHR-primitiveOffset-03660  
  For geometries of type `VK_GEOMETRY_TYPE_INSTANCES_KHR`, the offset `primitiveOffset` from [VkAccelerationStructureGeometryInstancesDataKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureGeometryInstancesDataKHR.html)::`data` **must** be a multiple of 16

## [](#_see_also)See Also

[VK\_KHR\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_acceleration_structure.html), [vkBuildAccelerationStructuresKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkBuildAccelerationStructuresKHR.html), [vkCmdBuildAccelerationStructuresKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBuildAccelerationStructuresKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkAccelerationStructureBuildRangeInfoKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0