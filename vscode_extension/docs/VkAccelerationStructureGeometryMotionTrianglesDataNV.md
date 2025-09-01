# VkAccelerationStructureGeometryMotionTrianglesDataNV(3) Manual Page

## Name

VkAccelerationStructureGeometryMotionTrianglesDataNV - Structure specifying vertex motion in a bottom-level acceleration structure



## [](#_c_specification)C Specification

The `VkAccelerationStructureGeometryMotionTrianglesDataNV` structure is defined as:

```c++
// Provided by VK_NV_ray_tracing_motion_blur
typedef struct VkAccelerationStructureGeometryMotionTrianglesDataNV {
    VkStructureType                  sType;
    const void*                      pNext;
    VkDeviceOrHostAddressConstKHR    vertexData;
} VkAccelerationStructureGeometryMotionTrianglesDataNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `vertexData` is a pointer to vertex data for this geometry at time 1.0

## [](#_description)Description

If `VkAccelerationStructureGeometryMotionTrianglesDataNV` is included in the `pNext` chain of a [VkAccelerationStructureGeometryTrianglesDataKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureGeometryTrianglesDataKHR.html) structure, the basic vertex positions are used for the position of the triangles in the geometry at time 0.0 and the `vertexData` in `VkAccelerationStructureGeometryMotionTrianglesDataNV` is used for the vertex positions at time 1.0, with positions linearly interpolated at intermediate times.

Indexing for `VkAccelerationStructureGeometryMotionTrianglesDataNV` `vertexData` is equivalent to the basic vertex position data.

Valid Usage (Implicit)

- [](#VUID-VkAccelerationStructureGeometryMotionTrianglesDataNV-sType-sType)VUID-VkAccelerationStructureGeometryMotionTrianglesDataNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_GEOMETRY_MOTION_TRIANGLES_DATA_NV`

## [](#_see_also)See Also

[VK\_NV\_ray\_tracing\_motion\_blur](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_ray_tracing_motion_blur.html), [VkDeviceOrHostAddressConstKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceOrHostAddressConstKHR.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkAccelerationStructureGeometryMotionTrianglesDataNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0