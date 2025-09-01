# VkClusterAccelerationStructureGeometryIndexAndGeometryFlagsNV(3) Manual Page

## Name

VkClusterAccelerationStructureGeometryIndexAndGeometryFlagsNV - Parameters describing geometry index and flags values for cluster acceleration structure



## [](#_c_specification)C Specification

The [VkClusterAccelerationStructureGeometryIndexAndGeometryFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureGeometryIndexAndGeometryFlagsNV.html) structure is defined as:

```c++
// Provided by VK_NV_cluster_acceleration_structure
typedef struct VkClusterAccelerationStructureGeometryIndexAndGeometryFlagsNV {
    uint32_t    geometryIndex:24;
    uint32_t    reserved:5;
    uint32_t    geometryFlags:3;
} VkClusterAccelerationStructureGeometryIndexAndGeometryFlagsNV;
```

## [](#_members)Members

- `geometryIndex` specifies the geometry index for all triangles in the cluster acceleration structure.
- `reserved` is reserved for future use.
- `geometryFlags` is a bitmask of [VkClusterAccelerationStructureGeometryFlagBitsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureGeometryFlagBitsNV.html) values describing geometry flags for the cluster acceleration structure.

## [](#_description)Description

The C language specification does not define the ordering of bit-fields, but in practice, this structure produces the correct layout with existing compilers. The intended bit pattern is the following:

- `geometryIndex`, `reserved` and `mask` occupy the same memory as if a single `uint32_t` was specified in their place
  
  - `geometryIndex` occupies the 24 least significant bits of that memory
  - `geometryFlags` occupies the 3 most significant bits of that memory

If a compiler produces code that diverges from that pattern, applications **must** employ another method to set values according to the correct bit pattern.

Valid Usage

- [](#VUID-VkClusterAccelerationStructureGeometryIndexAndGeometryFlagsNV-reserved-10487)VUID-VkClusterAccelerationStructureGeometryIndexAndGeometryFlagsNV-reserved-10487  
  `reserved` **must** be `0`

## [](#_see_also)See Also

[VK\_NV\_cluster\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_cluster_acceleration_structure.html), [VkClusterAccelerationStructureBuildTriangleClusterInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureBuildTriangleClusterInfoNV.html), [VkClusterAccelerationStructureBuildTriangleClusterTemplateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureBuildTriangleClusterTemplateInfoNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkClusterAccelerationStructureGeometryIndexAndGeometryFlagsNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0