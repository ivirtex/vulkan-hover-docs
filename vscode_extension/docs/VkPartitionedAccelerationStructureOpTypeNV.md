# VkPartitionedAccelerationStructureOpTypeNV(3) Manual Page

## Name

VkPartitionedAccelerationStructureOpTypeNV - Enum providing the type of PTLAS operation to perform



## [](#_c_specification)C Specification

Values which **can** be set in [VkPartitionedAccelerationStructureOpTypeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPartitionedAccelerationStructureOpTypeNV.html) are:

```c++
// Provided by VK_NV_partitioned_acceleration_structure
typedef enum VkPartitionedAccelerationStructureOpTypeNV {
    VK_PARTITIONED_ACCELERATION_STRUCTURE_OP_TYPE_WRITE_INSTANCE_NV = 0,
    VK_PARTITIONED_ACCELERATION_STRUCTURE_OP_TYPE_UPDATE_INSTANCE_NV = 1,
    VK_PARTITIONED_ACCELERATION_STRUCTURE_OP_TYPE_WRITE_PARTITION_TRANSLATION_NV = 2,
} VkPartitionedAccelerationStructureOpTypeNV;
```

## [](#_description)Description

- `VK_PARTITIONED_ACCELERATION_STRUCTURE_OP_TYPE_WRITE_INSTANCE_NV` is used to assign a transformed bottom level acceleration structure to an instance and partition. This is similar to [VkAccelerationStructureInstanceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureInstanceKHR.html) that defines the properties and transformations for a single instance in non-partitioned TLAS. Any partition that contains at least one of the affected instances will have their internal acceleration structure rebuilt.
- `VK_PARTITIONED_ACCELERATION_STRUCTURE_OP_TYPE_UPDATE_INSTANCE_NV` specifies that an instance will be updated with a new bottom level acceleration structure.
- `VK_PARTITIONED_ACCELERATION_STRUCTURE_OP_TYPE_WRITE_PARTITION_TRANSLATION_NV` specifies that a partition will be assigned a [translation vector](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#ptlas-partition-translation).

## [](#_see_also)See Also

[VK\_NV\_partitioned\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_partitioned_acceleration_structure.html), [VkBuildPartitionedAccelerationStructureIndirectCommandNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuildPartitionedAccelerationStructureIndirectCommandNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPartitionedAccelerationStructureOpTypeNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0