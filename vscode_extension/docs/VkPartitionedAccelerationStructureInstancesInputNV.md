# VkPartitionedAccelerationStructureInstancesInputNV(3) Manual Page

## Name

VkPartitionedAccelerationStructureInstancesInputNV - Parameters describing a PTLAS structure



## [](#_c_specification)C Specification

The [VkPartitionedAccelerationStructureInstancesInputNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPartitionedAccelerationStructureInstancesInputNV.html) structure is defined as:

```c++
// Provided by VK_NV_partitioned_acceleration_structure
typedef struct VkPartitionedAccelerationStructureInstancesInputNV {
    VkStructureType                         sType;
    void*                                   pNext;
    VkBuildAccelerationStructureFlagsKHR    flags;
    uint32_t                                instanceCount;
    uint32_t                                maxInstancePerPartitionCount;
    uint32_t                                partitionCount;
    uint32_t                                maxInstanceInGlobalPartitionCount;
} VkPartitionedAccelerationStructureInstancesInputNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is a bitmask of [VkBuildAccelerationStructureFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuildAccelerationStructureFlagsKHR.html) specifying flags for the PTLAS build operation.
- `instanceCount` is the number of instances in this PTLAS.
- `maxInstancePerPartitionCount` is the maximum number of instances per partition in the PTLAS.
- `partitionCount` is the number of partitions in the PTLAS.
- `maxInstanceInGlobalPartitionCount` is maximum number of instances in the [global partition](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#ptlas-global-partition).

## [](#_description)Description

If the `pNext` chain includes a [VkPartitionedAccelerationStructureFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPartitionedAccelerationStructureFlagsNV.html) structure, then that structure specifies additional flags for the PTLAS.

Valid Usage

- [](#VUID-VkPartitionedAccelerationStructureInstancesInputNV-partitionCount-10535)VUID-VkPartitionedAccelerationStructureInstancesInputNV-partitionCount-10535  
  The sum of `partitionCount` and `maxInstanceInGlobalPartitionCount` **must** be less than or equal to [VkPhysicalDevicePartitionedAccelerationStructurePropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevicePartitionedAccelerationStructurePropertiesNV.html)::`maxPartitionCount`

Valid Usage (Implicit)

- [](#VUID-VkPartitionedAccelerationStructureInstancesInputNV-sType-sType)VUID-VkPartitionedAccelerationStructureInstancesInputNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PARTITIONED_ACCELERATION_STRUCTURE_INSTANCES_INPUT_NV`
- [](#VUID-VkPartitionedAccelerationStructureInstancesInputNV-pNext-pNext)VUID-VkPartitionedAccelerationStructureInstancesInputNV-pNext-pNext  
  `pNext` **must** be `NULL` or a pointer to a valid instance of [VkPartitionedAccelerationStructureFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPartitionedAccelerationStructureFlagsNV.html)
- [](#VUID-VkPartitionedAccelerationStructureInstancesInputNV-sType-unique)VUID-VkPartitionedAccelerationStructureInstancesInputNV-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique
- [](#VUID-VkPartitionedAccelerationStructureInstancesInputNV-flags-parameter)VUID-VkPartitionedAccelerationStructureInstancesInputNV-flags-parameter  
  `flags` **must** be a valid combination of [VkBuildAccelerationStructureFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuildAccelerationStructureFlagBitsKHR.html) values

## [](#_see_also)See Also

[VK\_NV\_partitioned\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_partitioned_acceleration_structure.html), [VkBuildAccelerationStructureFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuildAccelerationStructureFlagsKHR.html), [VkBuildPartitionedAccelerationStructureInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuildPartitionedAccelerationStructureInfoNV.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetPartitionedAccelerationStructuresBuildSizesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPartitionedAccelerationStructuresBuildSizesNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPartitionedAccelerationStructureInstancesInputNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0