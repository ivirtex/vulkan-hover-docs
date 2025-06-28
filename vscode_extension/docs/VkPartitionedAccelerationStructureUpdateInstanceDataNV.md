# VkPartitionedAccelerationStructureUpdateInstanceDataNV(3) Manual Page

## Name

VkPartitionedAccelerationStructureUpdateInstanceDataNV - Structure describing instance data to update in PTLAS



## [](#_c_specification)C Specification

The [VkPartitionedAccelerationStructureUpdateInstanceDataNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPartitionedAccelerationStructureUpdateInstanceDataNV.html) structure is defined as:

```c++
// Provided by VK_NV_partitioned_acceleration_structure
typedef struct VkPartitionedAccelerationStructureUpdateInstanceDataNV {
    uint32_t           instanceIndex;
    uint32_t           instanceContributionToHitGroupIndex;
    VkDeviceAddress    accelerationStructure;
} VkPartitionedAccelerationStructureUpdateInstanceDataNV;
```

## [](#_members)Members

- `instanceIndex` is the index of the instance being updated.
- `instanceContributionToHitGroupIndex` is a 24-bit per instance value added in the indexing into the shader binding table to fetch the hit group to use.
- `accelerationStructure` is the device address of the bottom level acceleration structure or a clustered bottom level acceleration structure whose instance is being updated. The instance is disabled if the device address is `0`.

## [](#_description)Description

If the instance was originally disabled by specifying a `0` in [VkPartitionedAccelerationStructureWriteInstanceDataNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPartitionedAccelerationStructureWriteInstanceDataNV.html)::`accelerationStructure`, it can not be updated to a new acceleration structure as the instance **may** have been permanently disabled by the implementation.

To avoid a refit, the new acceleration structure **must** be within the bounding box specified by [VkPartitionedAccelerationStructureWriteInstanceDataNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPartitionedAccelerationStructureWriteInstanceDataNV.html)::`explicitAABB` when the instance was first created.

Valid Usage

- [](#VUID-VkPartitionedAccelerationStructureUpdateInstanceDataNV-instanceContributionToHitGroupIndex-10571)VUID-VkPartitionedAccelerationStructureUpdateInstanceDataNV-instanceContributionToHitGroupIndex-10571  
  The most significant 8 bits of `instanceContributionToHitGroupIndex` **must** be `0`
- [](#VUID-VkPartitionedAccelerationStructureUpdateInstanceDataNV-None-10572)VUID-VkPartitionedAccelerationStructureUpdateInstanceDataNV-None-10572  
  The instance **must** have either been created with flag `VK_PARTITIONED_ACCELERATION_STRUCTURE_INSTANCE_FLAG_ENABLE_EXPLICIT_BOUNDING_BOX_NV` or did not have an acceleration structure assigned with [VkPartitionedAccelerationStructureWriteInstanceDataNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPartitionedAccelerationStructureWriteInstanceDataNV.html)
- [](#VUID-VkPartitionedAccelerationStructureUpdateInstanceDataNV-instanceIndex-10573)VUID-VkPartitionedAccelerationStructureUpdateInstanceDataNV-instanceIndex-10573  
  `instanceIndex` **must** be less than [VkBuildPartitionedAccelerationStructureInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuildPartitionedAccelerationStructureInfoNV.html)::`input`::`instanceCount`

## [](#_see_also)See Also

[VK\_NV\_partitioned\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_partitioned_acceleration_structure.html), [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPartitionedAccelerationStructureUpdateInstanceDataNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0