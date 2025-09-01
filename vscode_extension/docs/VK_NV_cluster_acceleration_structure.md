# VK\_NV\_cluster\_acceleration\_structure(3) Manual Page

## Name

VK\_NV\_cluster\_acceleration\_structure - device extension



## [](#_registered_extension_number)Registered Extension Number

570

## [](#_revision)Revision

4

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_acceleration_structure.html)

## [](#_api_interactions)API Interactions

- Interacts with VK\_EXT\_opacity\_micromap
- Interacts with VK\_KHR\_ray\_tracing\_pipeline

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_NV\_cluster\_acceleration\_structure](https://github.khronos.org/SPIRV-Registry/extensions/NV/SPV_NV_cluster_acceleration_structure.html)

## [](#_contact)Contact

- Vikram Kushwaha [\[GitHub\]vkushwaha](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_cluster_acceleration_structure%5D%20%40vkushwaha%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_cluster_acceleration_structure%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_NV\_cluster\_acceleration\_structure](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_NV_cluster_acceleration_structure.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2024-09-09

**Contributors**

- Vikram Kushwaha, NVIDIA
- Eric Werness, NVIDIA
- Christoph Kubisch, NVIDIA
- Jan Schmid, NVIDIA
- Pyarelal Knowles, NVIDIA

## [](#_description)Description

Acceleration structure build times can become a bottleneck in ray tracing applications dealing with extensive dynamic geometry. This extension addresses the problem by enabling applications to construct bottom-level acceleration structures (BLAS) from pre-generated acceleration structures based on clusters of triangles (CLAS), leading to significant improvements in build times.

It provides a host-side query function to fetch the requirements and a versatile multi-indirect call for managing cluster geometry. This call enables applications to generate cluster geometry, construct Cluster BLAS from CLAS lists, and move or copy CLAS and BLAS. By sourcing inputs from device memory and processing multiple elements simultaneously, the call reduces the host-side costs associated with traditional acceleration structure functions.

This extension adds support for the following SPIR-V extension in Vulkan:

- `SPV_NV_cluster_acceleration_structure`

## [](#_new_commands)New Commands

- [vkCmdBuildClusterAccelerationStructureIndirectNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBuildClusterAccelerationStructureIndirectNV.html)
- [vkGetClusterAccelerationStructureBuildSizesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetClusterAccelerationStructureBuildSizesNV.html)

## [](#_new_structures)New Structures

- [VkClusterAccelerationStructureBuildClustersBottomLevelInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureBuildClustersBottomLevelInfoNV.html)
- [VkClusterAccelerationStructureBuildTriangleClusterInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureBuildTriangleClusterInfoNV.html)
- [VkClusterAccelerationStructureBuildTriangleClusterTemplateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureBuildTriangleClusterTemplateInfoNV.html)
- [VkClusterAccelerationStructureClustersBottomLevelInputNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureClustersBottomLevelInputNV.html)
- [VkClusterAccelerationStructureCommandsInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureCommandsInfoNV.html)
- [VkClusterAccelerationStructureGeometryIndexAndGeometryFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureGeometryIndexAndGeometryFlagsNV.html)
- [VkClusterAccelerationStructureGetTemplateIndicesInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureGetTemplateIndicesInfoNV.html)
- [VkClusterAccelerationStructureInputInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureInputInfoNV.html)
- [VkClusterAccelerationStructureInstantiateClusterInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureInstantiateClusterInfoNV.html)
- [VkClusterAccelerationStructureMoveObjectsInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureMoveObjectsInfoNV.html)
- [VkClusterAccelerationStructureMoveObjectsInputNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureMoveObjectsInputNV.html)
- [VkClusterAccelerationStructureTriangleClusterInputNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureTriangleClusterInputNV.html)
- [VkStridedDeviceAddressNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStridedDeviceAddressNV.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceClusterAccelerationStructureFeaturesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceClusterAccelerationStructureFeaturesNV.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceClusterAccelerationStructurePropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceClusterAccelerationStructurePropertiesNV.html)

If [VK\_KHR\_ray\_tracing\_pipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_ray_tracing_pipeline.html) is supported:

- Extending [VkRayTracingPipelineCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineCreateInfoKHR.html):
  
  - [VkRayTracingPipelineClusterAccelerationStructureCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineClusterAccelerationStructureCreateInfoNV.html)

## [](#_new_unions)New Unions

- [VkClusterAccelerationStructureOpInputNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureOpInputNV.html)

## [](#_new_enums)New Enums

- [VkClusterAccelerationStructureAddressResolutionFlagBitsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureAddressResolutionFlagBitsNV.html)
- [VkClusterAccelerationStructureClusterFlagBitsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureClusterFlagBitsNV.html)
- [VkClusterAccelerationStructureGeometryFlagBitsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureGeometryFlagBitsNV.html)
- [VkClusterAccelerationStructureIndexFormatFlagBitsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureIndexFormatFlagBitsNV.html)
- [VkClusterAccelerationStructureOpModeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureOpModeNV.html)
- [VkClusterAccelerationStructureOpTypeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureOpTypeNV.html)
- [VkClusterAccelerationStructureTypeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureTypeNV.html)

## [](#_new_bitmasks)New Bitmasks

- [VkClusterAccelerationStructureAddressResolutionFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureAddressResolutionFlagsNV.html)
- [VkClusterAccelerationStructureClusterFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureClusterFlagsNV.html)
- [VkClusterAccelerationStructureGeometryFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureGeometryFlagsNV.html)
- [VkClusterAccelerationStructureIndexFormatFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureIndexFormatFlagsNV.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_NV_CLUSTER_ACCELERATION_STRUCTURE_EXTENSION_NAME`
- `VK_NV_CLUSTER_ACCELERATION_STRUCTURE_SPEC_VERSION`
- Extending [VkBuildAccelerationStructureFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuildAccelerationStructureFlagBitsKHR.html):
  
  - `VK_BUILD_ACCELERATION_STRUCTURE_ALLOW_CLUSTER_OPACITY_MICROMAPS_BIT_NV`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_CLUSTER_ACCELERATION_STRUCTURE_CLUSTERS_BOTTOM_LEVEL_INPUT_NV`
  - `VK_STRUCTURE_TYPE_CLUSTER_ACCELERATION_STRUCTURE_COMMANDS_INFO_NV`
  - `VK_STRUCTURE_TYPE_CLUSTER_ACCELERATION_STRUCTURE_INPUT_INFO_NV`
  - `VK_STRUCTURE_TYPE_CLUSTER_ACCELERATION_STRUCTURE_MOVE_OBJECTS_INPUT_NV`
  - `VK_STRUCTURE_TYPE_CLUSTER_ACCELERATION_STRUCTURE_TRIANGLE_CLUSTER_INPUT_NV`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_CLUSTER_ACCELERATION_STRUCTURE_FEATURES_NV`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_CLUSTER_ACCELERATION_STRUCTURE_PROPERTIES_NV`
  - `VK_STRUCTURE_TYPE_RAY_TRACING_PIPELINE_CLUSTER_ACCELERATION_STRUCTURE_CREATE_INFO_NV`

If [VK\_EXT\_opacity\_micromap](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_opacity_micromap.html) is supported:

- Extending [VkOpacityMicromapSpecialIndexEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpacityMicromapSpecialIndexEXT.html):
  
  - `VK_OPACITY_MICROMAP_SPECIAL_INDEX_CLUSTER_GEOMETRY_DISABLE_OPACITY_MICROMAP_NV`

## [](#_new_or_modified_built_in_variables)New or Modified Built-In Variables

- [`ClusterIDNV`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-clusteridnv)

## [](#_new_spir_v_capability)New SPIR-V Capability

- [`RayTracingClusterAccelerationStructureNV`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-RayTracingClusterAccelerationStructureNV)

## [](#_version_history)Version History

- Revision 4, 2025-07-16 (Vikram Kushwaha)
  
  - Adding build flag to enable OMM in cluster acceleration structure
- Revision 3, 2025-06-18 (Vikram Kushwaha)
  
  - Adding a OpType to get template’s index data
- Revision 2, 2024-09-09 (Vikram Kushwaha)
  
  - Changes to some structures causing incompatibility with Revision 1
- Revision 1, 2024-08-29 (Vikram Kushwaha)
  
  - Initial draft

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_NV_cluster_acceleration_structure).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0