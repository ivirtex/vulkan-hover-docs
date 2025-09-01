# VkPartitionedAccelerationStructureWritePartitionTranslationDataNV(3) Manual Page

## Name

VkPartitionedAccelerationStructureWritePartitionTranslationDataNV - Structure describing partition translation data to write in PTLAS



## [](#_c_specification)C Specification

The [VkPartitionedAccelerationStructureWritePartitionTranslationDataNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPartitionedAccelerationStructureWritePartitionTranslationDataNV.html) structure is defined as:

```c++
// Provided by VK_NV_partitioned_acceleration_structure
typedef struct VkPartitionedAccelerationStructureWritePartitionTranslationDataNV {
    uint32_t    partitionIndex;
    float       partitionTranslation[3];
} VkPartitionedAccelerationStructureWritePartitionTranslationDataNV;
```

## [](#_members)Members

- `partitionIndex` is the index of partition to write. [Global partition](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#ptlas-global-partition) is referred to by `VK_PARTITIONED_ACCELERATION_STRUCTURE_PARTITION_INDEX_GLOBAL_NV`.
- `partitionTranslation` sets the [translation vector](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#ptlas-partition-translation) for this partition. When tracing this partition, the contained instances will behave as if the partition translation was added to the translation component of the instance transform. This translation vector is also added to the instances in the partition that had their bounding box specified.

## [](#_description)Description

Valid Usage

- [](#VUID-VkPartitionedAccelerationStructureWritePartitionTranslationDataNV-partitionIndex-10574)VUID-VkPartitionedAccelerationStructureWritePartitionTranslationDataNV-partitionIndex-10574  
  `partitionIndex` **must** be less than [VkBuildPartitionedAccelerationStructureInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuildPartitionedAccelerationStructureInfoNV.html)::`input.partitionCount`
- [](#VUID-VkPartitionedAccelerationStructureWritePartitionTranslationDataNV-enablePartitionTranslation-10575)VUID-VkPartitionedAccelerationStructureWritePartitionTranslationDataNV-enablePartitionTranslation-10575  
  The partitioned acceleration structure **must** have the [VkPartitionedAccelerationStructureFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPartitionedAccelerationStructureFlagsNV.html)::`enablePartitionTranslation` flag set

## [](#_see_also)See Also

[VK\_NV\_partitioned\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_partitioned_acceleration_structure.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPartitionedAccelerationStructureWritePartitionTranslationDataNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0