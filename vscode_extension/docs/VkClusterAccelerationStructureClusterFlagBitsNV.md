# VkClusterAccelerationStructureClusterFlagBitsNV(3) Manual Page

## Name

VkClusterAccelerationStructureClusterFlagBitsNV - Bitmask specifying cluster acceleration structure flags



## [](#_c_specification)C Specification

Bits which **can** be set in [VkClusterAccelerationStructureBuildTriangleClusterInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureBuildTriangleClusterInfoNV.html)::`clusterFlags`, specifying flags for clusters in an acceleration structure, are:

```c++
// Provided by VK_NV_cluster_acceleration_structure
typedef enum VkClusterAccelerationStructureClusterFlagBitsNV {
    VK_CLUSTER_ACCELERATION_STRUCTURE_CLUSTER_ALLOW_DISABLE_OPACITY_MICROMAPS_NV = 0x00000001,
} VkClusterAccelerationStructureClusterFlagBitsNV;
```

## [](#_description)Description

- `VK_CLUSTER_ACCELERATION_STRUCTURE_CLUSTER_ALLOW_DISABLE_OPACITY_MICROMAPS_NV` specifies that the specified cluster acceleration structure **may** be referenced in an instance with `VK_GEOMETRY_INSTANCE_DISABLE_OPACITY_MICROMAPS_BIT_EXT` set.

## [](#_see_also)See Also

[VK\_NV\_cluster\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_cluster_acceleration_structure.html), [VkClusterAccelerationStructureClusterFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureClusterFlagsNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkClusterAccelerationStructureClusterFlagBitsNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0