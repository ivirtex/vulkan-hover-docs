# VkClusterAccelerationStructureGeometryFlagBitsNV(3) Manual Page

## Name

VkClusterAccelerationStructureGeometryFlagBitsNV - Bitmask specifying geometry flags for cluster acceleration structure



## [](#_c_specification)C Specification

Bits which **can** be set in [VkClusterAccelerationStructureGeometryIndexAndGeometryFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureGeometryIndexAndGeometryFlagsNV.html)::`geometryFlags`, specifying geometry flags for cluster acceleration structure, are:

```c++
// Provided by VK_NV_cluster_acceleration_structure
typedef enum VkClusterAccelerationStructureGeometryFlagBitsNV {
    VK_CLUSTER_ACCELERATION_STRUCTURE_GEOMETRY_CULL_DISABLE_BIT_NV = 0x00000001,
    VK_CLUSTER_ACCELERATION_STRUCTURE_GEOMETRY_NO_DUPLICATE_ANYHIT_INVOCATION_BIT_NV = 0x00000002,
    VK_CLUSTER_ACCELERATION_STRUCTURE_GEOMETRY_OPAQUE_BIT_NV = 0x00000004,
} VkClusterAccelerationStructureGeometryFlagBitsNV;
```

## [](#_description)Description

- `VK_CLUSTER_ACCELERATION_STRUCTURE_GEOMETRY_CULL_DISABLE_BIT_NV` disables face culling for this geometry.
- `VK_CLUSTER_ACCELERATION_STRUCTURE_GEOMETRY_NO_DUPLICATE_ANYHIT_INVOCATION_BIT_NV` specifies that the implementation **must** only call the any-hit shader a single time for each primitive in this geometry. If this bit is absent an implementation **may** invoke the any-hit shader more than once for this geometry.
- `VK_CLUSTER_ACCELERATION_STRUCTURE_GEOMETRY_OPAQUE_BIT_NV` specifies that this geometry does not invoke the any-hit shaders even if present in a hit group.

## [](#_see_also)See Also

[VK\_NV\_cluster\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_cluster_acceleration_structure.html), [VkClusterAccelerationStructureGeometryFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureGeometryFlagsNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkClusterAccelerationStructureGeometryFlagBitsNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0