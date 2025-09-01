# VK\_NV\_partitioned\_acceleration\_structure(3) Manual Page

## Name

VK\_NV\_partitioned\_acceleration\_structure - device extension



## [](#_registered_extension_number)Registered Extension Number

571

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_acceleration_structure.html)

## [](#_contact)Contact

- Vikram Kushwaha [\[GitHub\]vkushwaha](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_partitioned_acceleration_structure%5D%20%40vkushwaha%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_partitioned_acceleration_structure%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_NV\_partitioned\_acceleration\_structure](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_NV_partitioned_acceleration_structure.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2025-01-09

**Contributors**

- Vikram Kushwaha, NVIDIA
- Eric Werness, NVIDIA
- Christoph Kubisch, NVIDIA
- Jan Schmid, NVIDIA
- Pyarelal Knowles, NVIDIA

## [](#_description)Description

With an increase in scene complexity and expansive game worlds, the number of instances has surged in ray tracing over the last few years. The current Top Level Acceleration Structure (TLAS) API necessitates a full rebuild of the entire data structure even when only a few instances are modified.

This extension introduces Partitioned Top Level Acceleration Structures (PTLAS) as an alternative to the existing TLAS. PTLAS enables the efficient reuse of previously constructed parts of the acceleration structure, resulting in much faster build times and supporting a higher number of instances.

## [](#_new_commands)New Commands

- [vkCmdBuildPartitionedAccelerationStructuresNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBuildPartitionedAccelerationStructuresNV.html)
- [vkGetPartitionedAccelerationStructuresBuildSizesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPartitionedAccelerationStructuresBuildSizesNV.html)

## [](#_new_structures)New Structures

- [VkBuildPartitionedAccelerationStructureIndirectCommandNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuildPartitionedAccelerationStructureIndirectCommandNV.html)
- [VkBuildPartitionedAccelerationStructureInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuildPartitionedAccelerationStructureInfoNV.html)
- [VkPartitionedAccelerationStructureInstancesInputNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPartitionedAccelerationStructureInstancesInputNV.html)
- [VkPartitionedAccelerationStructureUpdateInstanceDataNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPartitionedAccelerationStructureUpdateInstanceDataNV.html)
- [VkPartitionedAccelerationStructureWriteInstanceDataNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPartitionedAccelerationStructureWriteInstanceDataNV.html)
- [VkPartitionedAccelerationStructureWritePartitionTranslationDataNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPartitionedAccelerationStructureWritePartitionTranslationDataNV.html)
- Extending [VkPartitionedAccelerationStructureInstancesInputNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPartitionedAccelerationStructureInstancesInputNV.html):
  
  - [VkPartitionedAccelerationStructureFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPartitionedAccelerationStructureFlagsNV.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDevicePartitionedAccelerationStructureFeaturesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevicePartitionedAccelerationStructureFeaturesNV.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDevicePartitionedAccelerationStructurePropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevicePartitionedAccelerationStructurePropertiesNV.html)
- Extending [VkWriteDescriptorSet](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWriteDescriptorSet.html):
  
  - [VkWriteDescriptorSetPartitionedAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWriteDescriptorSetPartitionedAccelerationStructureNV.html)

## [](#_new_enums)New Enums

- [VkPartitionedAccelerationStructureInstanceFlagBitsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPartitionedAccelerationStructureInstanceFlagBitsNV.html)
- [VkPartitionedAccelerationStructureOpTypeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPartitionedAccelerationStructureOpTypeNV.html)

## [](#_new_bitmasks)New Bitmasks

- [VkPartitionedAccelerationStructureInstanceFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPartitionedAccelerationStructureInstanceFlagsNV.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_NV_PARTITIONED_ACCELERATION_STRUCTURE_EXTENSION_NAME`
- `VK_NV_PARTITIONED_ACCELERATION_STRUCTURE_SPEC_VERSION`
- `VK_PARTITIONED_ACCELERATION_STRUCTURE_PARTITION_INDEX_GLOBAL_NV`
- Extending [VkDescriptorType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorType.html):
  
  - `VK_DESCRIPTOR_TYPE_PARTITIONED_ACCELERATION_STRUCTURE_NV`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_BUILD_PARTITIONED_ACCELERATION_STRUCTURE_INFO_NV`
  - `VK_STRUCTURE_TYPE_PARTITIONED_ACCELERATION_STRUCTURE_FLAGS_NV`
  - `VK_STRUCTURE_TYPE_PARTITIONED_ACCELERATION_STRUCTURE_INSTANCES_INPUT_NV`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PARTITIONED_ACCELERATION_STRUCTURE_FEATURES_NV`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PARTITIONED_ACCELERATION_STRUCTURE_PROPERTIES_NV`
  - `VK_STRUCTURE_TYPE_WRITE_DESCRIPTOR_SET_PARTITIONED_ACCELERATION_STRUCTURE_NV`

## [](#_version_history)Version History

- Revision 1, 2025-01-09 (Vikram Kushwaha)
  
  - Initial draft

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_NV_partitioned_acceleration_structure).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0