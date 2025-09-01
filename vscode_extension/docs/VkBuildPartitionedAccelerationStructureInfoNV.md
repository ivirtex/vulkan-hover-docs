# VkBuildPartitionedAccelerationStructureInfoNV(3) Manual Page

## Name

VkBuildPartitionedAccelerationStructureInfoNV - Structure describing build parameters for a PTLAS



## [](#_c_specification)C Specification

The [VkBuildPartitionedAccelerationStructureInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuildPartitionedAccelerationStructureInfoNV.html) structure is defined as:

```c++
// Provided by VK_NV_partitioned_acceleration_structure
typedef struct VkBuildPartitionedAccelerationStructureInfoNV {
    VkStructureType                                       sType;
    void*                                                 pNext;
    VkPartitionedAccelerationStructureInstancesInputNV    input;
    VkDeviceAddress                                       srcAccelerationStructureData;
    VkDeviceAddress                                       dstAccelerationStructureData;
    VkDeviceAddress                                       scratchData;
    VkDeviceAddress                                       srcInfos;
    VkDeviceAddress                                       srcInfosCount;
} VkBuildPartitionedAccelerationStructureInfoNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `input` is a [VkPartitionedAccelerationStructureInstancesInputNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPartitionedAccelerationStructureInstancesInputNV.html) structure describing the instance and partition count information in the PTLAS.
- `srcAccelerationStructureData` is `NULL` or an address of a previously built PTLAS. If non-`NULL`, the PTLAS stored at this address is used as a basis to create new PTLAS.
- `dstAccelerationStructureData` is the address to store the built PTLAS.
- `scratchData` is the device address of scratch memory that will be used during PTLAS build.
- `srcInfos` is the device address of an array of [VkBuildPartitionedAccelerationStructureIndirectCommandNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuildPartitionedAccelerationStructureIndirectCommandNV.html) structures describing the type of operation to perform.
- `srcInfosCount` is a device address containing the size of `srcInfos` array.

## [](#_description)Description

Members `srcAccelerationStructureData` and `dstAccelerationStructureData` **may** be the same or different. If they are the same, the update happens in-place. Otherwise, the destination acceleration structure is updated and the source is not modified.

Valid Usage

- [](#VUID-VkBuildPartitionedAccelerationStructureInfoNV-scratchData-10558)VUID-VkBuildPartitionedAccelerationStructureInfoNV-scratchData-10558  
  `scratchData` **must** not be `NULL`
- [](#VUID-VkBuildPartitionedAccelerationStructureInfoNV-scratchData-10559)VUID-VkBuildPartitionedAccelerationStructureInfoNV-scratchData-10559  
  Memory at `scratchData` **must** be equal or larger than the [VkAccelerationStructureBuildSizesInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureBuildSizesInfoKHR.html)::`buildScratchSize` value returned from [vkGetPartitionedAccelerationStructuresBuildSizesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPartitionedAccelerationStructuresBuildSizesNV.html) with the same build parameters
- [](#VUID-VkBuildPartitionedAccelerationStructureInfoNV-srcAccelerationStructureData-10560)VUID-VkBuildPartitionedAccelerationStructureInfoNV-srcAccelerationStructureData-10560  
  If `srcAccelerationStructureData` is not `NULL`, it **must** have previously been built as a PTLAS
- [](#VUID-VkBuildPartitionedAccelerationStructureInfoNV-dstAccelerationStructureData-10561)VUID-VkBuildPartitionedAccelerationStructureInfoNV-dstAccelerationStructureData-10561  
  `dstAccelerationStructureData` **must** not be `NULL`
- [](#VUID-VkBuildPartitionedAccelerationStructureInfoNV-dstAccelerationStructureData-10562)VUID-VkBuildPartitionedAccelerationStructureInfoNV-dstAccelerationStructureData-10562  
  Memory at `dstAccelerationStructureData` **must** be equal or larger than the [VkAccelerationStructureBuildSizesInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureBuildSizesInfoKHR.html)::`accelerationStructureSize` value returned from [vkGetPartitionedAccelerationStructuresBuildSizesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPartitionedAccelerationStructuresBuildSizesNV.html) with the same build parameters
- [](#VUID-VkBuildPartitionedAccelerationStructureInfoNV-srcInfosCount-10563)VUID-VkBuildPartitionedAccelerationStructureInfoNV-srcInfosCount-10563  
  `srcInfosCount` **must** be 4-byte aligned
- [](#VUID-VkBuildPartitionedAccelerationStructureInfoNV-srcInfos-10564)VUID-VkBuildPartitionedAccelerationStructureInfoNV-srcInfos-10564  
  Each element of `srcInfos` array **must** have a unique [VkBuildPartitionedAccelerationStructureIndirectCommandNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuildPartitionedAccelerationStructureIndirectCommandNV.html)::`opType`

Valid Usage (Implicit)

- [](#VUID-VkBuildPartitionedAccelerationStructureInfoNV-sType-sType)VUID-VkBuildPartitionedAccelerationStructureInfoNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_BUILD_PARTITIONED_ACCELERATION_STRUCTURE_INFO_NV`
- [](#VUID-VkBuildPartitionedAccelerationStructureInfoNV-pNext-pNext)VUID-VkBuildPartitionedAccelerationStructureInfoNV-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkBuildPartitionedAccelerationStructureInfoNV-input-parameter)VUID-VkBuildPartitionedAccelerationStructureInfoNV-input-parameter  
  `input` **must** be a valid [VkPartitionedAccelerationStructureInstancesInputNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPartitionedAccelerationStructureInstancesInputNV.html) structure
- [](#VUID-VkBuildPartitionedAccelerationStructureInfoNV-srcAccelerationStructureData-parameter)VUID-VkBuildPartitionedAccelerationStructureInfoNV-srcAccelerationStructureData-parameter  
  `srcAccelerationStructureData` **must** be a valid [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html) value
- [](#VUID-VkBuildPartitionedAccelerationStructureInfoNV-dstAccelerationStructureData-parameter)VUID-VkBuildPartitionedAccelerationStructureInfoNV-dstAccelerationStructureData-parameter  
  `dstAccelerationStructureData` **must** be a valid [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html) value
- [](#VUID-VkBuildPartitionedAccelerationStructureInfoNV-scratchData-parameter)VUID-VkBuildPartitionedAccelerationStructureInfoNV-scratchData-parameter  
  `scratchData` **must** be a valid [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html) value
- [](#VUID-VkBuildPartitionedAccelerationStructureInfoNV-srcInfos-parameter)VUID-VkBuildPartitionedAccelerationStructureInfoNV-srcInfos-parameter  
  `srcInfos` **must** be a valid [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html) value
- [](#VUID-VkBuildPartitionedAccelerationStructureInfoNV-srcInfosCount-parameter)VUID-VkBuildPartitionedAccelerationStructureInfoNV-srcInfosCount-parameter  
  `srcInfosCount` **must** be a valid [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html) value

## [](#_see_also)See Also

[VK\_NV\_partitioned\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_partitioned_acceleration_structure.html), [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html), [VkPartitionedAccelerationStructureInstancesInputNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPartitionedAccelerationStructureInstancesInputNV.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCmdBuildPartitionedAccelerationStructuresNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBuildPartitionedAccelerationStructuresNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkBuildPartitionedAccelerationStructureInfoNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0