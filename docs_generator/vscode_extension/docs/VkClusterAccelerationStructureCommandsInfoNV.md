# VkClusterAccelerationStructureCommandsInfoNV(3) Manual Page

## Name

VkClusterAccelerationStructureCommandsInfoNV - Structure describing parameters for building for moving an acceleration structure



## [](#_c_specification)C Specification

The [VkClusterAccelerationStructureCommandsInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureCommandsInfoNV.html) structure is defined as:

```c++
// Provided by VK_NV_cluster_acceleration_structure
typedef struct VkClusterAccelerationStructureCommandsInfoNV {
    VkStructureType                                           sType;
    void*                                                     pNext;
    VkClusterAccelerationStructureInputInfoNV                 input;
    VkDeviceAddress                                           dstImplicitData;
    VkDeviceAddress                                           scratchData;
    VkStridedDeviceAddressRegionKHR                           dstAddressesArray;
    VkStridedDeviceAddressRegionKHR                           dstSizesArray;
    VkStridedDeviceAddressRegionKHR                           srcInfosArray;
    VkDeviceAddress                                           srcInfosCount;
    VkClusterAccelerationStructureAddressResolutionFlagsNV    addressResolutionFlags;
} VkClusterAccelerationStructureCommandsInfoNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `input` is [VkClusterAccelerationStructureInputInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureInputInfoNV.html) structure describing the build or move parameters for the cluster acceleration structure.
- `dstImplicitData` is the device address for memory where the implicit build of cluster acceleration structure will be saved. If [VkClusterAccelerationStructureInputInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureInputInfoNV.html)::`opMode` is `VK_CLUSTER_ACCELERATION_STRUCTURE_OP_MODE_EXPLICIT_DESTINATIONS_NV` or `VK_CLUSTER_ACCELERATION_STRUCTURE_OP_MODE_COMPUTE_SIZES_NV`, this value is ignored.
- `scratchData` is the device address of scratch memory that will be used during cluster acceleration structure move or build.
- `dstAddressesArray` is a [VkStridedDeviceAddressRegionKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStridedDeviceAddressRegionKHR.html) where the individual addresses and stride of moved or built cluster acceleration structures will be saved or read from depending on [VkClusterAccelerationStructureInputInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureInputInfoNV.html)::`opMode`. If [VkClusterAccelerationStructureInputInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureInputInfoNV.html)::`opMode` is `VK_CLUSTER_ACCELERATION_STRUCTURE_OP_MODE_IMPLICIT_DESTINATIONS_NV` and the address in `dstAddressesArray` is not `0`, then the addresses are saved. If [VkClusterAccelerationStructureInputInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureInputInfoNV.html)::`opMode` is `VK_CLUSTER_ACCELERATION_STRUCTURE_OP_MODE_EXPLICIT_DESTINATIONS_NV`, then the addresses are read from. If [VkClusterAccelerationStructureInputInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureInputInfoNV.html)::`opMode` is `VK_CLUSTER_ACCELERATION_STRUCTURE_OP_MODE_COMPUTE_SIZES_NV`, then this value is ignored and **may** be `0`.
- `dstSizesArray` is `NULL` or a [VkStridedDeviceAddressRegionKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStridedDeviceAddressRegionKHR.html) containing sizes of moved or built cluster acceleration structures. Similar to `dstAddressesArray`, if [VkClusterAccelerationStructureInputInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureInputInfoNV.html)::`opMode` is `VK_CLUSTER_ACCELERATION_STRUCTURE_OP_MODE_IMPLICIT_DESTINATIONS_NV`, then the sizes are saved. If [VkClusterAccelerationStructureInputInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureInputInfoNV.html)::`opMode` is `VK_CLUSTER_ACCELERATION_STRUCTURE_OP_MODE_EXPLICIT_DESTINATIONS_NV`, then the sizes are read from.
- `srcInfosArray` is a [VkStridedDeviceAddressRegionKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStridedDeviceAddressRegionKHR.html) where input data for the build or move operation is read from. If the stride is `0`, the structures are assumed to be packed tightly. Its format is dependent on [VkClusterAccelerationStructureInputInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureInputInfoNV.html)::`opType` as per the table below.

`input`::`opType`

Format of `srcInfosArray`

`VK_CLUSTER_ACCELERATION_STRUCTURE_OP_TYPE_MOVE_OBJECTS_NV`

[VkClusterAccelerationStructureMoveObjectsInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureMoveObjectsInfoNV.html)

`VK_CLUSTER_ACCELERATION_STRUCTURE_OP_TYPE_BUILD_CLUSTERS_BOTTOM_LEVEL_NV`

[VkClusterAccelerationStructureBuildClustersBottomLevelInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureBuildClustersBottomLevelInfoNV.html)

`VK_CLUSTER_ACCELERATION_STRUCTURE_OP_TYPE_BUILD_TRIANGLE_CLUSTER_NV`

[VkClusterAccelerationStructureBuildTriangleClusterInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureBuildTriangleClusterInfoNV.html)

`VK_CLUSTER_ACCELERATION_STRUCTURE_OP_TYPE_BUILD_TRIANGLE_CLUSTER_TEMPLATE_NV`

[VkClusterAccelerationStructureBuildTriangleClusterTemplateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureBuildTriangleClusterTemplateInfoNV.html)

`VK_CLUSTER_ACCELERATION_STRUCTURE_OP_TYPE_INSTANTIATE_TRIANGLE_CLUSTER_NV`

[VkClusterAccelerationStructureInstantiateClusterInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureInstantiateClusterInfoNV.html)

`VK_CLUSTER_ACCELERATION_STRUCTURE_OP_TYPE_GET_CLUSTER_TEMPLATE_INDICES_NV`

[VkClusterAccelerationStructureGetTemplateIndicesInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureGetTemplateIndicesInfoNV.html)

- `srcInfosCount` is the device address of memory containing the count of number of build or move operations to perform. The actual value is the minimum of this value and the value specified in `input`::`maxAccelerationStructureCount`. If this value is `0`, the count is determined by `input`::`maxAccelerationStructureCount` alone.
- `addressResolutionFlags` is a bitmask of [VkClusterAccelerationStructureAddressResolutionFlagBitsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureAddressResolutionFlagBitsNV.html) values specifying how an implementation will interpret the device addresses in this structure.

## [](#_description)Description

Valid Usage

- [](#VUID-VkClusterAccelerationStructureCommandsInfoNV-opMode-10466)VUID-VkClusterAccelerationStructureCommandsInfoNV-opMode-10466  
  If [VkClusterAccelerationStructureInputInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureInputInfoNV.html)::`opMode` is `VK_CLUSTER_ACCELERATION_STRUCTURE_OP_MODE_IMPLICIT_DESTINATIONS_NV`, `dstImplicitData` **must** be a valid address
- [](#VUID-VkClusterAccelerationStructureCommandsInfoNV-opMode-10467)VUID-VkClusterAccelerationStructureCommandsInfoNV-opMode-10467  
  If [VkClusterAccelerationStructureInputInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureInputInfoNV.html)::`opMode` is `VK_CLUSTER_ACCELERATION_STRUCTURE_OP_MODE_IMPLICIT_DESTINATIONS_NV` and `input`::`opType` is not `VK_CLUSTER_ACCELERATION_STRUCTURE_OP_TYPE_MOVE_OBJECTS_NV`, the memory in `dstImplicitData` **must** be equal to or larger than the [VkAccelerationStructureBuildSizesInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureBuildSizesInfoKHR.html)::`accelerationStructureSize` value returned from [vkGetClusterAccelerationStructureBuildSizesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetClusterAccelerationStructureBuildSizesNV.html) with same input parameters
- [](#VUID-VkClusterAccelerationStructureCommandsInfoNV-opMode-10468)VUID-VkClusterAccelerationStructureCommandsInfoNV-opMode-10468  
  If [VkClusterAccelerationStructureInputInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureInputInfoNV.html)::`opMode` is `VK_CLUSTER_ACCELERATION_STRUCTURE_OP_MODE_IMPLICIT_DESTINATIONS_NV` and `input`::`opType` is `VK_CLUSTER_ACCELERATION_STRUCTURE_OP_TYPE_MOVE_OBJECTS_NV`, the memory in `dstImplicitData` **must** be equal to or larger than the sum of all the built acceleration structures that are being moved
- [](#VUID-VkClusterAccelerationStructureCommandsInfoNV-input-10469)VUID-VkClusterAccelerationStructureCommandsInfoNV-input-10469  
  If `input`::`opType` is `VK_CLUSTER_ACCELERATION_STRUCTURE_OP_TYPE_MOVE_OBJECTS_NV`, the total memory moved **must** not be larger than the size provided in [VkClusterAccelerationStructureMoveObjectsInputNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureMoveObjectsInputNV.html)::`maxMovedBytes`
- [](#VUID-VkClusterAccelerationStructureCommandsInfoNV-opMode-10470)VUID-VkClusterAccelerationStructureCommandsInfoNV-opMode-10470  
  If [VkClusterAccelerationStructureInputInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureInputInfoNV.html)::`opMode` is `VK_CLUSTER_ACCELERATION_STRUCTURE_OP_MODE_COMPUTE_SIZES_NV`, `dstSizesArray` **must** be a valid address
- [](#VUID-VkClusterAccelerationStructureCommandsInfoNV-opMode-10471)VUID-VkClusterAccelerationStructureCommandsInfoNV-opMode-10471  
  If [VkClusterAccelerationStructureInputInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureInputInfoNV.html)::`opMode` is `VK_CLUSTER_ACCELERATION_STRUCTURE_OP_MODE_EXPLICIT_DESTINATIONS_NV`, the address in `dstAddressesArray` **must** be a valid address with sizes of individual buffers large enough to accommodate built or moved clusters
- [](#VUID-VkClusterAccelerationStructureCommandsInfoNV-opMode-10472)VUID-VkClusterAccelerationStructureCommandsInfoNV-opMode-10472  
  If [VkClusterAccelerationStructureInputInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureInputInfoNV.html)::`opMode` is `VK_CLUSTER_ACCELERATION_STRUCTURE_OP_MODE_EXPLICIT_DESTINATIONS_NV`, the buffers in `dstAddressesArray` **must** not overlap
- [](#VUID-VkClusterAccelerationStructureCommandsInfoNV-opMode-10473)VUID-VkClusterAccelerationStructureCommandsInfoNV-opMode-10473  
  If [VkClusterAccelerationStructureInputInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureInputInfoNV.html)::`opMode` is `VK_CLUSTER_ACCELERATION_STRUCTURE_OP_MODE_EXPLICIT_DESTINATIONS_NV`, the addresses in `dstAddressesArray` **must** be aligned based on the cluster acceleration structure type and its alignment properties as described in [VkPhysicalDeviceClusterAccelerationStructurePropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceClusterAccelerationStructurePropertiesNV.html)
- [](#VUID-VkClusterAccelerationStructureCommandsInfoNV-dstAddressesArray-10474)VUID-VkClusterAccelerationStructureCommandsInfoNV-dstAddressesArray-10474  
  The stride in `dstAddressesArray` **must** be greater than or equal to 8
- [](#VUID-VkClusterAccelerationStructureCommandsInfoNV-dstSizesArray-10475)VUID-VkClusterAccelerationStructureCommandsInfoNV-dstSizesArray-10475  
  The stride in `dstSizesArray` **must** be greater than or equal to 4
- [](#VUID-VkClusterAccelerationStructureCommandsInfoNV-srcInfosArray-10476)VUID-VkClusterAccelerationStructureCommandsInfoNV-srcInfosArray-10476  
  The stride in `srcInfosArray` **must** be greater than the type of structure the address is describing
- [](#VUID-VkClusterAccelerationStructureCommandsInfoNV-input-10477)VUID-VkClusterAccelerationStructureCommandsInfoNV-input-10477  
  If `input`::`opType` is `VK_CLUSTER_ACCELERATION_STRUCTURE_OP_TYPE_BUILD_TRIANGLE_CLUSTER_NV`, then depending on the [VkClusterAccelerationStructureInputInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureInputInfoNV.html)::`opMode`, `dstImplicitData` or addresses specified in `dstAddressesArray` **must** be aligned to [VkPhysicalDeviceClusterAccelerationStructurePropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceClusterAccelerationStructurePropertiesNV.html)::`clusterByteAlignment`
- [](#VUID-VkClusterAccelerationStructureCommandsInfoNV-input-10478)VUID-VkClusterAccelerationStructureCommandsInfoNV-input-10478  
  If `input`::`opType` is `VK_CLUSTER_ACCELERATION_STRUCTURE_OP_TYPE_BUILD_TRIANGLE_CLUSTER_TEMPLATE_NV`, then depending on the [VkClusterAccelerationStructureInputInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureInputInfoNV.html)::`opMode`, `dstImplicitData` or addresses specified in `dstAddressesArray` **must** be aligned to [VkPhysicalDeviceClusterAccelerationStructurePropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceClusterAccelerationStructurePropertiesNV.html)::`clusterTemplateByteAlignment`
- [](#VUID-VkClusterAccelerationStructureCommandsInfoNV-input-10479)VUID-VkClusterAccelerationStructureCommandsInfoNV-input-10479  
  If `input`::`opType` is `VK_CLUSTER_ACCELERATION_STRUCTURE_OP_TYPE_INSTANTIATE_TRIANGLE_CLUSTER_NV`, then depending on the [VkClusterAccelerationStructureInputInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureInputInfoNV.html)::`opMode`, `dstImplicitData` or addresses specified in `dstAddressesArray` **must** be aligned to [VkPhysicalDeviceClusterAccelerationStructurePropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceClusterAccelerationStructurePropertiesNV.html)::`clusterByteAlignment`
- [](#VUID-VkClusterAccelerationStructureCommandsInfoNV-scratchData-10480)VUID-VkClusterAccelerationStructureCommandsInfoNV-scratchData-10480  
  `scratchData` **must** be aligned to [VkPhysicalDeviceClusterAccelerationStructurePropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceClusterAccelerationStructurePropertiesNV.html)::`clusterScratchByteAlignment`
- [](#VUID-VkClusterAccelerationStructureCommandsInfoNV-srcInfosCount-10481)VUID-VkClusterAccelerationStructureCommandsInfoNV-srcInfosCount-10481  
  `srcInfosCount` **must** be 4-byte aligned
- [](#VUID-VkClusterAccelerationStructureCommandsInfoNV-input-10482)VUID-VkClusterAccelerationStructureCommandsInfoNV-input-10482  
  If `input`::`opType` is `VK_CLUSTER_ACCELERATION_STRUCTURE_OP_TYPE_BUILD_CLUSTERS_BOTTOM_LEVEL_NV`, the total and per argument number of cluster acceleration structures referenced in `srcInfosArray` **must** be equal or less than the maximum values with which memory requirements were queried in [vkGetClusterAccelerationStructureBuildSizesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetClusterAccelerationStructureBuildSizesNV.html) with [VkClusterAccelerationStructureOpInputNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureOpInputNV.html)::`pClustersBottomLevel`

Valid Usage (Implicit)

- [](#VUID-VkClusterAccelerationStructureCommandsInfoNV-sType-sType)VUID-VkClusterAccelerationStructureCommandsInfoNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_CLUSTER_ACCELERATION_STRUCTURE_COMMANDS_INFO_NV`
- [](#VUID-VkClusterAccelerationStructureCommandsInfoNV-pNext-pNext)VUID-VkClusterAccelerationStructureCommandsInfoNV-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkClusterAccelerationStructureCommandsInfoNV-input-parameter)VUID-VkClusterAccelerationStructureCommandsInfoNV-input-parameter  
  `input` **must** be a valid [VkClusterAccelerationStructureInputInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureInputInfoNV.html) structure
- [](#VUID-VkClusterAccelerationStructureCommandsInfoNV-addressResolutionFlags-parameter)VUID-VkClusterAccelerationStructureCommandsInfoNV-addressResolutionFlags-parameter  
  `addressResolutionFlags` **must** be a valid combination of [VkClusterAccelerationStructureAddressResolutionFlagBitsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureAddressResolutionFlagBitsNV.html) values

## [](#_see_also)See Also

[VK\_NV\_cluster\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_cluster_acceleration_structure.html), [VkClusterAccelerationStructureAddressResolutionFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureAddressResolutionFlagsNV.html), [VkClusterAccelerationStructureInputInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureInputInfoNV.html), [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html), [VkStridedDeviceAddressRegionKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStridedDeviceAddressRegionKHR.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCmdBuildClusterAccelerationStructureIndirectNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBuildClusterAccelerationStructureIndirectNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkClusterAccelerationStructureCommandsInfoNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0