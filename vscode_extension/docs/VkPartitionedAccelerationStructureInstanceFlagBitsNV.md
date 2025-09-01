# VkPartitionedAccelerationStructureInstanceFlagBitsNV(3) Manual Page

## Name

VkPartitionedAccelerationStructureInstanceFlagBitsNV - Bitmask specifying flags for PTLAS instances



## [](#_c_specification)C Specification

Bits which **can** be set in [VkPartitionedAccelerationStructureWriteInstanceDataNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPartitionedAccelerationStructureWriteInstanceDataNV.html)::`instanceFlags`, specifying flags for instances, are:

```c++
// Provided by VK_NV_partitioned_acceleration_structure
typedef enum VkPartitionedAccelerationStructureInstanceFlagBitsNV {
    VK_PARTITIONED_ACCELERATION_STRUCTURE_INSTANCE_FLAG_TRIANGLE_FACING_CULL_DISABLE_BIT_NV = 0x00000001,
    VK_PARTITIONED_ACCELERATION_STRUCTURE_INSTANCE_FLAG_TRIANGLE_FLIP_FACING_BIT_NV = 0x00000002,
    VK_PARTITIONED_ACCELERATION_STRUCTURE_INSTANCE_FLAG_FORCE_OPAQUE_BIT_NV = 0x00000004,
    VK_PARTITIONED_ACCELERATION_STRUCTURE_INSTANCE_FLAG_FORCE_NO_OPAQUE_BIT_NV = 0x00000008,
    VK_PARTITIONED_ACCELERATION_STRUCTURE_INSTANCE_FLAG_ENABLE_EXPLICIT_BOUNDING_BOX_NV = 0x00000010,
} VkPartitionedAccelerationStructureInstanceFlagBitsNV;
```

## [](#_description)Description

- `VK_PARTITIONED_ACCELERATION_STRUCTURE_INSTANCE_FLAG_TRIANGLE_FACING_CULL_DISABLE_BIT_NV` disables face culling for this instance.
- `VK_PARTITIONED_ACCELERATION_STRUCTURE_INSTANCE_FLAG_TRIANGLE_FLIP_FACING_BIT_NV` specifies that the [facing determination](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#ray-traversal-culling-face) for geometry in this instance is inverted.
- `VK_PARTITIONED_ACCELERATION_STRUCTURE_INSTANCE_FLAG_FORCE_OPAQUE_BIT_NV` causes this instance to act as though `VK_GEOMETRY_OPAQUE_BIT_KHR` were specified on all geometries referenced by this instance.
- `VK_PARTITIONED_ACCELERATION_STRUCTURE_INSTANCE_FLAG_FORCE_NO_OPAQUE_BIT_NV` causes this instance to act as though `VK_GEOMETRY_OPAQUE_BIT_KHR` were not specified on all geometries referenced by this instance.
- `VK_PARTITIONED_ACCELERATION_STRUCTURE_INSTANCE_FLAG_ENABLE_EXPLICIT_BOUNDING_BOX_NV` enables use of an explicit bounding box for this instance.

## [](#_see_also)See Also

[VK\_NV\_partitioned\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_partitioned_acceleration_structure.html), [VkPartitionedAccelerationStructureInstanceFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPartitionedAccelerationStructureInstanceFlagsNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPartitionedAccelerationStructureInstanceFlagBitsNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0