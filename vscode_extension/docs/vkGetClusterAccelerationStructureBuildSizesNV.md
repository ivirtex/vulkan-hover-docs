# vkGetClusterAccelerationStructureBuildSizesNV(3) Manual Page

## Name

vkGetClusterAccelerationStructureBuildSizesNV - Retrieve the buffer allocation requirements for cluster geometry command



## [](#_c_specification)C Specification

These cluster acceleration structures **can** be built or moved by a single versatile multi-indirect function [vkCmdBuildClusterAccelerationStructureIndirectNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBuildClusterAccelerationStructureIndirectNV.html). To determine the memory requirements for executing this function, call:

```c++
// Provided by VK_NV_cluster_acceleration_structure
void vkGetClusterAccelerationStructureBuildSizesNV(
    VkDevice                                    device,
    const VkClusterAccelerationStructureInputInfoNV* pInfo,
    VkAccelerationStructureBuildSizesInfoKHR*   pSizeInfo);
```

## [](#_parameters)Parameters

- `device` is the logical device that owns the acceleration structure.
- `pInfo` is a pointer to a [VkClusterAccelerationStructureInputInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureInputInfoNV.html) structure containing parameters required for the memory requirements query.
- `pSizeInfo` is a pointer to a [VkAccelerationStructureBuildSizesInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureBuildSizesInfoKHR.html) structure which returns the size required for an acceleration structure and scratch buffer, given the build parameters.

## [](#_description)Description

If [VkClusterAccelerationStructureInputInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureInputInfoNV.html)::`opMode` is `VK_CLUSTER_ACCELERATION_STRUCTURE_OP_MODE_IMPLICIT_DESTINATIONS_NV`, acceleration structure and scratch memory sizes are returned for all [VkClusterAccelerationStructureInputInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureInputInfoNV.html)::`maxAccelerationStructureCount` acceleration structures. If [VkClusterAccelerationStructureInputInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureInputInfoNV.html)::`opMode` is `VK_CLUSTER_ACCELERATION_STRUCTURE_OP_MODE_EXPLICIT_DESTINATIONS_NV`, scratch memory size for all [VkClusterAccelerationStructureInputInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureInputInfoNV.html)::`maxAccelerationStructureCount` acceleration structures and the acceleration structure memory size for a single acceleration structure is returned. If [VkClusterAccelerationStructureInputInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureInputInfoNV.html)::`opMode` is `VK_CLUSTER_ACCELERATION_STRUCTURE_OP_MODE_COMPUTE_SIZES_NV`, only scratch memory size is returned for the requested acceleration structures.

Valid Usage

- [](#VUID-vkGetClusterAccelerationStructureBuildSizesNV-clusterAccelerationStructure-10438)VUID-vkGetClusterAccelerationStructureBuildSizesNV-clusterAccelerationStructure-10438  
  The [`VkPhysicalDeviceClusterAccelerationStructureFeaturesNV`::`clusterAccelerationStructure`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-clusterAccelerationStructure) feature **must** be enabled

Valid Usage (Implicit)

- [](#VUID-vkGetClusterAccelerationStructureBuildSizesNV-device-parameter)VUID-vkGetClusterAccelerationStructureBuildSizesNV-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetClusterAccelerationStructureBuildSizesNV-pInfo-parameter)VUID-vkGetClusterAccelerationStructureBuildSizesNV-pInfo-parameter  
  `pInfo` **must** be a valid pointer to a valid [VkClusterAccelerationStructureInputInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureInputInfoNV.html) structure
- [](#VUID-vkGetClusterAccelerationStructureBuildSizesNV-pSizeInfo-parameter)VUID-vkGetClusterAccelerationStructureBuildSizesNV-pSizeInfo-parameter  
  `pSizeInfo` **must** be a valid pointer to a [VkAccelerationStructureBuildSizesInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureBuildSizesInfoKHR.html) structure

## [](#_see_also)See Also

[VK\_NV\_cluster\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_cluster_acceleration_structure.html), [VkAccelerationStructureBuildSizesInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureBuildSizesInfoKHR.html), [VkClusterAccelerationStructureInputInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureInputInfoNV.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetClusterAccelerationStructureBuildSizesNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0