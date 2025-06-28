# VkBuildPartitionedAccelerationStructureIndirectCommandNV(3) Manual Page

## Name

VkBuildPartitionedAccelerationStructureIndirectCommandNV - Structure describing PTLAS operation to perform



## [](#_c_specification)C Specification

The [VkBuildPartitionedAccelerationStructureIndirectCommandNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuildPartitionedAccelerationStructureIndirectCommandNV.html) structure is defined as:

```c++
// Provided by VK_NV_partitioned_acceleration_structure
typedef struct VkBuildPartitionedAccelerationStructureIndirectCommandNV {
    VkPartitionedAccelerationStructureOpTypeNV    opType;
    uint32_t                                      argCount;
    VkStridedDeviceAddressNV                      argData;
} VkBuildPartitionedAccelerationStructureIndirectCommandNV;
```

## [](#_members)Members

- `opType` is a [VkPartitionedAccelerationStructureOpTypeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPartitionedAccelerationStructureOpTypeNV.html) describing the type of operation.
- `argCount` the number of structures in `argData` array.
- `argData` is an array of [VkStridedDeviceAddressNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStridedDeviceAddressNV.html) structures containing the write or update data for instances and partitions in the PTLAS. The structure is dependent on `opType` as shown in the table below.

`opType`

Format of `argData`

`VK_PARTITIONED_ACCELERATION_STRUCTURE_OP_TYPE_WRITE_INSTANCE_NV`

[VkPartitionedAccelerationStructureWriteInstanceDataNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPartitionedAccelerationStructureWriteInstanceDataNV.html)

`VK_PARTITIONED_ACCELERATION_STRUCTURE_OP_TYPE_UPDATE_INSTANCE_NV`

[VkPartitionedAccelerationStructureUpdateInstanceDataNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPartitionedAccelerationStructureUpdateInstanceDataNV.html)

`VK_PARTITIONED_ACCELERATION_STRUCTURE_OP_TYPE_WRITE_PARTITION_TRANSLATION_NV`

[VkPartitionedAccelerationStructureWritePartitionTranslationDataNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPartitionedAccelerationStructureWritePartitionTranslationDataNV.html)

## [](#_description)Description

Valid Usage

- [](#VUID-VkBuildPartitionedAccelerationStructureIndirectCommandNV-argData-10565)VUID-VkBuildPartitionedAccelerationStructureIndirectCommandNV-argData-10565  
  An instance index **must** not be referenced by more than one structure in `argData`

Valid Usage (Implicit)

- [](#VUID-VkBuildPartitionedAccelerationStructureIndirectCommandNV-opType-parameter)VUID-VkBuildPartitionedAccelerationStructureIndirectCommandNV-opType-parameter  
  `opType` **must** be a valid [VkPartitionedAccelerationStructureOpTypeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPartitionedAccelerationStructureOpTypeNV.html) value

## [](#_see_also)See Also

[VK\_NV\_partitioned\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_partitioned_acceleration_structure.html), [VkPartitionedAccelerationStructureOpTypeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPartitionedAccelerationStructureOpTypeNV.html), [VkStridedDeviceAddressNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStridedDeviceAddressNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkBuildPartitionedAccelerationStructureIndirectCommandNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0