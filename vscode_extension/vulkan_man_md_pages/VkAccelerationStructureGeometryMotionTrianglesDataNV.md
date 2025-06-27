# VkAccelerationStructureGeometryMotionTrianglesDataNV(3) Manual Page

## Name

VkAccelerationStructureGeometryMotionTrianglesDataNV - Structure
specifying vertex motion in a bottom-level acceleration structure



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkAccelerationStructureGeometryMotionTrianglesDataNV` structure is
defined as:

``` c
// Provided by VK_NV_ray_tracing_motion_blur
typedef struct VkAccelerationStructureGeometryMotionTrianglesDataNV {
    VkStructureType                  sType;
    const void*                      pNext;
    VkDeviceOrHostAddressConstKHR    vertexData;
} VkAccelerationStructureGeometryMotionTrianglesDataNV;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `vertexData` is a pointer to vertex data for this geometry at time 1.0

## <a href="#_description" class="anchor"></a>Description

If `VkAccelerationStructureGeometryMotionTrianglesDataNV` is included in
the `pNext` chain of a
[VkAccelerationStructureGeometryTrianglesDataKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureGeometryTrianglesDataKHR.html)
structure, the basic vertex positions are used for the position of the
triangles in the geometry at time 0.0 and the `vertexData` in
`VkAccelerationStructureGeometryMotionTrianglesDataNV` is used for the
vertex positions at time 1.0, with positions linearly interpolated at
intermediate times.

Indexing for `VkAccelerationStructureGeometryMotionTrianglesDataNV`
`vertexData` is equivalent to the basic vertex position data.

Valid Usage (Implicit)

- <a
  href="#VUID-VkAccelerationStructureGeometryMotionTrianglesDataNV-sType-sType"
  id="VUID-VkAccelerationStructureGeometryMotionTrianglesDataNV-sType-sType"></a>
  VUID-VkAccelerationStructureGeometryMotionTrianglesDataNV-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_GEOMETRY_MOTION_TRIANGLES_DATA_NV`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_ray_tracing_motion_blur](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_ray_tracing_motion_blur.html),
[VkDeviceOrHostAddressConstKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceOrHostAddressConstKHR.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkAccelerationStructureGeometryMotionTrianglesDataNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
