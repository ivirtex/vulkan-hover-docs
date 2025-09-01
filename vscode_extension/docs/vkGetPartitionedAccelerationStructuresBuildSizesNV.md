# vkGetPartitionedAccelerationStructuresBuildSizesNV(3) Manual Page

## Name

vkGetPartitionedAccelerationStructuresBuildSizesNV - Retrieve the buffer allocation requirements for partitioned acceleration structure command



## [](#_c_specification)C Specification

To determine the memory requirements for a PTAS, call:

```c++
// Provided by VK_NV_partitioned_acceleration_structure
void vkGetPartitionedAccelerationStructuresBuildSizesNV(
    VkDevice                                    device,
    const VkPartitionedAccelerationStructureInstancesInputNV* pInfo,
    VkAccelerationStructureBuildSizesInfoKHR*   pSizeInfo);
```

## [](#_parameters)Parameters

- `device` is the logical device that owns the acceleration structure.
- `pInfo` is a pointer to a [VkPartitionedAccelerationStructureInstancesInputNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPartitionedAccelerationStructureInstancesInputNV.html) structure containing parameters required for the memory requirements query.
- `pSizeInfo` is a pointer to a [VkAccelerationStructureBuildSizesInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureBuildSizesInfoKHR.html) structure which returns the size required for an acceleration structure and the sizes required for the scratch buffers, given the build parameters.

## [](#_description)Description

Valid Usage

- [](#VUID-vkGetPartitionedAccelerationStructuresBuildSizesNV-partitionedAccelerationStructure-10534)VUID-vkGetPartitionedAccelerationStructuresBuildSizesNV-partitionedAccelerationStructure-10534  
  The [`VkPhysicalDevicePartitionedAccelerationStructureFeaturesNV`::`partitionedAccelerationStructure`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-partitionedAccelerationStructure) feature **must** be enabled

Valid Usage (Implicit)

- [](#VUID-vkGetPartitionedAccelerationStructuresBuildSizesNV-device-parameter)VUID-vkGetPartitionedAccelerationStructuresBuildSizesNV-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetPartitionedAccelerationStructuresBuildSizesNV-pInfo-parameter)VUID-vkGetPartitionedAccelerationStructuresBuildSizesNV-pInfo-parameter  
  `pInfo` **must** be a valid pointer to a valid [VkPartitionedAccelerationStructureInstancesInputNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPartitionedAccelerationStructureInstancesInputNV.html) structure
- [](#VUID-vkGetPartitionedAccelerationStructuresBuildSizesNV-pSizeInfo-parameter)VUID-vkGetPartitionedAccelerationStructuresBuildSizesNV-pSizeInfo-parameter  
  `pSizeInfo` **must** be a valid pointer to a [VkAccelerationStructureBuildSizesInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureBuildSizesInfoKHR.html) structure

## [](#_see_also)See Also

[VK\_NV\_partitioned\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_partitioned_acceleration_structure.html), [VkAccelerationStructureBuildSizesInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureBuildSizesInfoKHR.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkPartitionedAccelerationStructureInstancesInputNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPartitionedAccelerationStructureInstancesInputNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetPartitionedAccelerationStructuresBuildSizesNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0