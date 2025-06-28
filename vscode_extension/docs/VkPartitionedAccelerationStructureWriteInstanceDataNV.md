# VkPartitionedAccelerationStructureWriteInstanceDataNV(3) Manual Page

## Name

VkPartitionedAccelerationStructureWriteInstanceDataNV - Structure describing instance data to write in PTLAS



## [](#_c_specification)C Specification

The [VkPartitionedAccelerationStructureWriteInstanceDataNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPartitionedAccelerationStructureWriteInstanceDataNV.html) structure is defined as:

```c++
// Provided by VK_NV_partitioned_acceleration_structure
typedef struct VkPartitionedAccelerationStructureWriteInstanceDataNV {
    VkTransformMatrixKHR                                 transform;
    float                                                explicitAABB[6];
    uint32_t                                             instanceID;
    uint32_t                                             instanceMask;
    uint32_t                                             instanceContributionToHitGroupIndex;
    VkPartitionedAccelerationStructureInstanceFlagsNV    instanceFlags;
    uint32_t                                             instanceIndex;
    uint32_t                                             partitionIndex;
    VkDeviceAddress                                      accelerationStructure;
} VkPartitionedAccelerationStructureWriteInstanceDataNV;
```

## [](#_members)Members

- `transform` is a [VkTransformMatrixKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTransformMatrixKHR.html) structure describing the transformation to be applied to the instance in PTLAS.
- `explicitAABB` specifies an axis aligned bounding box representing the maximum extent of any vertex within the used acceleration structure after applying the instance-to-world transformation. The [partition translation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#ptlas-partition-translation) is not applied to the bounding box.
- `instanceID` is a user specified constant assigned to an instance in the PTLAS.
- `instanceMask` is a 8-bit mask assigned to the instance that **may** be used to include or reject group of instances.
- `instanceContributionToHitGroupIndex` is a 24-bit per application specified instance value added in the indexing into the shader binding table to fetch the hit group to use.
- `instanceFlags` is a bitmask of [VkPartitionedAccelerationStructureInstanceFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPartitionedAccelerationStructureInstanceFlagsNV.html) specifying flags an instance in the PTLAS.
- `instanceIndex` is the index of the instance within the PTLAS.
- `partitionIndex` is the index of the partition to which this instance belongs. [Global partitions](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#ptlas-global-partition) are referred to by `VK_PARTITIONED_ACCELERATION_STRUCTURE_PARTITION_INDEX_GLOBAL_NV`.
- `accelerationStructure` is the device address of the bottom level acceleration structure or a clustered bottom level acceleration structure that is being instanced. This instance is disabled if the device address is `0`.

## [](#_description)Description

Valid Usage

- [](#VUID-VkPartitionedAccelerationStructureWriteInstanceDataNV-instanceMask-10566)VUID-VkPartitionedAccelerationStructureWriteInstanceDataNV-instanceMask-10566  
  The most significant 24 bits of `instanceMask` **must** be `0`
- [](#VUID-VkPartitionedAccelerationStructureWriteInstanceDataNV-instanceContributionToHitGroupIndex-10567)VUID-VkPartitionedAccelerationStructureWriteInstanceDataNV-instanceContributionToHitGroupIndex-10567  
  The most significant 8 bits of `instanceContributionToHitGroupIndex` **must** be `0`
- [](#VUID-VkPartitionedAccelerationStructureWriteInstanceDataNV-instanceIndex-10568)VUID-VkPartitionedAccelerationStructureWriteInstanceDataNV-instanceIndex-10568  
  `instanceIndex` **must** be less than [VkBuildPartitionedAccelerationStructureInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuildPartitionedAccelerationStructureInfoNV.html)::`input`::`instanceCount`
- [](#VUID-VkPartitionedAccelerationStructureWriteInstanceDataNV-partitionIndex-10569)VUID-VkPartitionedAccelerationStructureWriteInstanceDataNV-partitionIndex-10569  
  `partitionIndex` **must** be less than [VkBuildPartitionedAccelerationStructureInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuildPartitionedAccelerationStructureInfoNV.html)::`input`::`partitionCount`
- [](#VUID-VkPartitionedAccelerationStructureWriteInstanceDataNV-explicitAABB-10570)VUID-VkPartitionedAccelerationStructureWriteInstanceDataNV-explicitAABB-10570  
  `explicitAABB` **must** be a valid bounding box if instance was created with flag `VK_PARTITIONED_ACCELERATION_STRUCTURE_INSTANCE_FLAG_ENABLE_EXPLICIT_BOUNDING_BOX_NV` set

Valid Usage (Implicit)

- [](#VUID-VkPartitionedAccelerationStructureWriteInstanceDataNV-instanceFlags-parameter)VUID-VkPartitionedAccelerationStructureWriteInstanceDataNV-instanceFlags-parameter  
  `instanceFlags` **must** be a valid combination of [VkPartitionedAccelerationStructureInstanceFlagBitsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPartitionedAccelerationStructureInstanceFlagBitsNV.html) values

## [](#_see_also)See Also

[VK\_NV\_partitioned\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_partitioned_acceleration_structure.html), [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html), [VkPartitionedAccelerationStructureInstanceFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPartitionedAccelerationStructureInstanceFlagsNV.html), [VkTransformMatrixKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTransformMatrixKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPartitionedAccelerationStructureWriteInstanceDataNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0