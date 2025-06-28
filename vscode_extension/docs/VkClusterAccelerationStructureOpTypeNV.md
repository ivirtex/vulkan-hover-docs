# VkClusterAccelerationStructureOpTypeNV(3) Manual Page

## Name

VkClusterAccelerationStructureOpTypeNV - Enum providing the type of operation



## [](#_c_specification)C Specification

Values which **can** be set in `VkClusterAccelerationStructureOpTypeNV` are:

```c++
// Provided by VK_NV_cluster_acceleration_structure
typedef enum VkClusterAccelerationStructureOpTypeNV {
    VK_CLUSTER_ACCELERATION_STRUCTURE_OP_TYPE_MOVE_OBJECTS_NV = 0,
    VK_CLUSTER_ACCELERATION_STRUCTURE_OP_TYPE_BUILD_CLUSTERS_BOTTOM_LEVEL_NV = 1,
    VK_CLUSTER_ACCELERATION_STRUCTURE_OP_TYPE_BUILD_TRIANGLE_CLUSTER_NV = 2,
    VK_CLUSTER_ACCELERATION_STRUCTURE_OP_TYPE_BUILD_TRIANGLE_CLUSTER_TEMPLATE_NV = 3,
    VK_CLUSTER_ACCELERATION_STRUCTURE_OP_TYPE_INSTANTIATE_TRIANGLE_CLUSTER_NV = 4,
    VK_CLUSTER_ACCELERATION_STRUCTURE_OP_TYPE_GET_CLUSTER_TEMPLATE_INDICES_NV = 5,
} VkClusterAccelerationStructureOpTypeNV;
```

## [](#_description)Description

- `VK_CLUSTER_ACCELERATION_STRUCTURE_OP_TYPE_MOVE_OBJECTS_NV` specifies that a cluster acceleration structure, cluster acceleration structure template or a bottom level acceleration structure built from cluster acceleration structures will be moved. If a cluster acceleration structure is moved, the bottom level cluster acceleration structures containing it will have to be re-built. If used with `VK_CLUSTER_ACCELERATION_STRUCTURE_OP_MODE_COMPUTE_SIZES_NV`, it returns the size of existing cluster acceleration structures.
- `VK_CLUSTER_ACCELERATION_STRUCTURE_OP_TYPE_BUILD_CLUSTERS_BOTTOM_LEVEL_NV` specifies that bottom level cluster acceleration structures will be built.
- `VK_CLUSTER_ACCELERATION_STRUCTURE_OP_TYPE_BUILD_TRIANGLE_CLUSTER_NV` specifies that cluster acceleration structures will be built.
- `VK_CLUSTER_ACCELERATION_STRUCTURE_OP_TYPE_BUILD_TRIANGLE_CLUSTER_TEMPLATE_NV` specifies that a template for cluster acceleration structure will be built.
- `VK_CLUSTER_ACCELERATION_STRUCTURE_OP_TYPE_INSTANTIATE_TRIANGLE_CLUSTER_NV` specifies that a template for a cluster acceleration structure will be instantiated, resulting in a built cluster acceleration structure.
- `VK_CLUSTER_ACCELERATION_STRUCTURE_OP_TYPE_GET_CLUSTER_TEMPLATE_INDICES_NV` specifies that the vertex indices of the cluster template acceleration structure will be fetched.

## [](#_see_also)See Also

[VK\_NV\_cluster\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_cluster_acceleration_structure.html), [VkClusterAccelerationStructureInputInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureInputInfoNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkClusterAccelerationStructureOpTypeNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0