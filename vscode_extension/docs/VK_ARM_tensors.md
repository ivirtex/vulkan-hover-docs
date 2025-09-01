# VK\_ARM\_tensors(3) Manual Page

## Name

VK\_ARM\_tensors - device extension



## [](#_registered_extension_number)Registered Extension Number

461

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[Vulkan Version 1.3](#versions-1.3)

## [](#_api_interactions)API Interactions

- Interacts with VK\_EXT\_descriptor\_buffer
- Interacts with VK\_EXT\_frame\_boundary

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_ARM\_tensors](https://github.khronos.org/SPIRV-Registry/extensions/ARM/SPV_ARM_tensors.html)

## [](#_contact)Contact

- Kevin Petit [\[GitHub\]kpet](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_ARM_tensors%5D%20%40kpet%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_ARM_tensors%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_ARM\_tensors](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_ARM_tensors.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2025-06-03

**Interactions and External Dependencies**

- This extension requires [`SPV_ARM_tensors`](https://github.khronos.org/SPIRV-Registry/extensions/ARM/SPV_ARM_tensors.html)
- This extension provides API support for [`GL_ARM_tensors`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/arm/GL_ARM_tensors.txt)
- This extension interacts with `VK_EXT_mutable_descriptor_type`
- This extension interacts with `VK_EXT_descriptor_buffer`
- This extension interacts with `VK_EXT_frame_boundary`
- This extension interacts with `VK_EXT_robustness2`

**IP Status**

No known IP claims.

**Contributors**

- Kévin Petit, Arm Ltd.
- Einar Hov, Arm Ltd.
- Dominic Symes, Arm Ltd.
- Jan-Harald Fredriksen, Arm Ltd.
- Marco Cattani, Arm Ltd.
- Lisa Wu, Arm Ltd.
- Robert Hughes, Arm Ltd.
- David Garbett, Arm Ltd.
- Oualid Khelifi, Arm Ltd.

## [](#_description)Description

This extension adds support for tensors.

## [](#_new_object_types)New Object Types

- [VkTensorARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorARM.html)
- [VkTensorViewARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorViewARM.html)

## [](#_new_commands)New Commands

- [vkBindTensorMemoryARM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkBindTensorMemoryARM.html)
- [vkCmdCopyTensorARM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyTensorARM.html)
- [vkCreateTensorARM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateTensorARM.html)
- [vkCreateTensorViewARM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateTensorViewARM.html)
- [vkDestroyTensorARM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkDestroyTensorARM.html)
- [vkDestroyTensorViewARM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkDestroyTensorViewARM.html)
- [vkGetDeviceTensorMemoryRequirementsARM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceTensorMemoryRequirementsARM.html)
- [vkGetPhysicalDeviceExternalTensorPropertiesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceExternalTensorPropertiesARM.html)
- [vkGetTensorMemoryRequirementsARM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetTensorMemoryRequirementsARM.html)

If [VK\_EXT\_descriptor\_buffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_descriptor_buffer.html) is supported:

- [vkGetTensorOpaqueCaptureDescriptorDataARM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetTensorOpaqueCaptureDescriptorDataARM.html)
- [vkGetTensorViewOpaqueCaptureDescriptorDataARM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetTensorViewOpaqueCaptureDescriptorDataARM.html)

## [](#_new_structures)New Structures

- [VkBindTensorMemoryInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindTensorMemoryInfoARM.html)
- [VkCopyTensorInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyTensorInfoARM.html)
- [VkDeviceTensorMemoryRequirementsARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceTensorMemoryRequirementsARM.html)
- [VkExternalTensorPropertiesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalTensorPropertiesARM.html)
- [VkPhysicalDeviceExternalTensorInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceExternalTensorInfoARM.html)
- [VkTensorCopyARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorCopyARM.html)
- [VkTensorCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorCreateInfoARM.html)
- [VkTensorMemoryRequirementsInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorMemoryRequirementsInfoARM.html)
- [VkTensorViewCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorViewCreateInfoARM.html)
- Extending [VkDataGraphPipelineResourceInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineResourceInfoARM.html), [VkDataGraphPipelineConstantARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineConstantARM.html):
  
  - [VkTensorDescriptionARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorDescriptionARM.html)
- Extending [VkDependencyInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDependencyInfo.html):
  
  - [VkTensorDependencyInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorDependencyInfoARM.html)
  - [VkTensorMemoryBarrierARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorMemoryBarrierARM.html)
- Extending [VkFormatProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatProperties2.html):
  
  - [VkTensorFormatPropertiesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorFormatPropertiesARM.html)
- Extending [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryAllocateInfo.html):
  
  - [VkMemoryDedicatedAllocateInfoTensorARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryDedicatedAllocateInfoTensorARM.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceTensorFeaturesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceTensorFeaturesARM.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceTensorPropertiesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceTensorPropertiesARM.html)
- Extending [VkTensorCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorCreateInfoARM.html):
  
  - [VkExternalMemoryTensorCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryTensorCreateInfoARM.html)
- Extending [VkWriteDescriptorSet](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWriteDescriptorSet.html):
  
  - [VkWriteDescriptorSetTensorARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWriteDescriptorSetTensorARM.html)

If [VK\_EXT\_descriptor\_buffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_descriptor_buffer.html) is supported:

- [VkTensorCaptureDescriptorDataInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorCaptureDescriptorDataInfoARM.html)
- [VkTensorViewCaptureDescriptorDataInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorViewCaptureDescriptorDataInfoARM.html)
- Extending [VkDescriptorGetInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorGetInfoEXT.html):
  
  - [VkDescriptorGetTensorInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorGetTensorInfoARM.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceDescriptorBufferTensorFeaturesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDescriptorBufferTensorFeaturesARM.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceDescriptorBufferTensorPropertiesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDescriptorBufferTensorPropertiesARM.html)

If [VK\_EXT\_frame\_boundary](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_frame_boundary.html) is supported:

- Extending [VkSubmitInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubmitInfo.html), [VkSubmitInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubmitInfo2.html), [VkPresentInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentInfoKHR.html), [VkBindSparseInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindSparseInfo.html):
  
  - [VkFrameBoundaryTensorsARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFrameBoundaryTensorsARM.html)

## [](#_new_enums)New Enums

- [VkTensorCreateFlagBitsARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorCreateFlagBitsARM.html)
- [VkTensorTilingARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorTilingARM.html)
- [VkTensorUsageFlagBitsARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorUsageFlagBitsARM.html)
- [VkTensorViewCreateFlagBitsARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorViewCreateFlagBitsARM.html)

## [](#_new_bitmasks)New Bitmasks

- [VkTensorCreateFlagsARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorCreateFlagsARM.html)
- [VkTensorUsageFlagsARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorUsageFlagsARM.html)
- [VkTensorViewCreateFlagsARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorViewCreateFlagsARM.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_ARM_TENSORS_EXTENSION_NAME`
- `VK_ARM_TENSORS_SPEC_VERSION`
- Extending [VkDescriptorType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorType.html):
  
  - `VK_DESCRIPTOR_TYPE_TENSOR_ARM`
- Extending [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html):
  
  - `VK_FORMAT_R8_BOOL_ARM`
- Extending [VkFormatFeatureFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatFeatureFlagBits2.html):
  
  - `VK_FORMAT_FEATURE_2_TENSOR_IMAGE_ALIASING_BIT_ARM`
  - `VK_FORMAT_FEATURE_2_TENSOR_SHADER_BIT_ARM`
- Extending [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html):
  
  - `VK_IMAGE_LAYOUT_TENSOR_ALIASING_ARM`
- Extending [VkImageUsageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageUsageFlagBits.html):
  
  - `VK_IMAGE_USAGE_TENSOR_ALIASING_BIT_ARM`
- Extending [VkObjectType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkObjectType.html):
  
  - `VK_OBJECT_TYPE_TENSOR_ARM`
  - `VK_OBJECT_TYPE_TENSOR_VIEW_ARM`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_BIND_TENSOR_MEMORY_INFO_ARM`
  - `VK_STRUCTURE_TYPE_COPY_TENSOR_INFO_ARM`
  - `VK_STRUCTURE_TYPE_DEVICE_TENSOR_MEMORY_REQUIREMENTS_ARM`
  - `VK_STRUCTURE_TYPE_EXTERNAL_MEMORY_TENSOR_CREATE_INFO_ARM`
  - `VK_STRUCTURE_TYPE_EXTERNAL_TENSOR_PROPERTIES_ARM`
  - `VK_STRUCTURE_TYPE_MEMORY_DEDICATED_ALLOCATE_INFO_TENSOR_ARM`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_TENSOR_INFO_ARM`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_TENSOR_FEATURES_ARM`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_TENSOR_PROPERTIES_ARM`
  - `VK_STRUCTURE_TYPE_TENSOR_COPY_ARM`
  - `VK_STRUCTURE_TYPE_TENSOR_CREATE_INFO_ARM`
  - `VK_STRUCTURE_TYPE_TENSOR_DEPENDENCY_INFO_ARM`
  - `VK_STRUCTURE_TYPE_TENSOR_DESCRIPTION_ARM`
  - `VK_STRUCTURE_TYPE_TENSOR_FORMAT_PROPERTIES_ARM`
  - `VK_STRUCTURE_TYPE_TENSOR_MEMORY_BARRIER_ARM`
  - `VK_STRUCTURE_TYPE_TENSOR_MEMORY_REQUIREMENTS_INFO_ARM`
  - `VK_STRUCTURE_TYPE_TENSOR_VIEW_CREATE_INFO_ARM`
  - `VK_STRUCTURE_TYPE_WRITE_DESCRIPTOR_SET_TENSOR_ARM`

If [VK\_EXT\_descriptor\_buffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_descriptor_buffer.html) is supported:

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_DESCRIPTOR_GET_TENSOR_INFO_ARM`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DESCRIPTOR_BUFFER_TENSOR_FEATURES_ARM`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DESCRIPTOR_BUFFER_TENSOR_PROPERTIES_ARM`
  - `VK_STRUCTURE_TYPE_TENSOR_CAPTURE_DESCRIPTOR_DATA_INFO_ARM`
  - `VK_STRUCTURE_TYPE_TENSOR_VIEW_CAPTURE_DESCRIPTOR_DATA_INFO_ARM`
- Extending [VkTensorCreateFlagBitsARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorCreateFlagBitsARM.html):
  
  - `VK_TENSOR_CREATE_DESCRIPTOR_BUFFER_CAPTURE_REPLAY_BIT_ARM`
- Extending [VkTensorViewCreateFlagBitsARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorViewCreateFlagBitsARM.html):
  
  - `VK_TENSOR_VIEW_CREATE_DESCRIPTOR_BUFFER_CAPTURE_REPLAY_BIT_ARM`

If [VK\_EXT\_frame\_boundary](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_frame_boundary.html) is supported:

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_FRAME_BOUNDARY_TENSORS_ARM`

## [](#_new_spir_v_capabilities)New SPIR-V Capabilities

- [`TensorsARM`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-TensorsARM)
- [`StorageTensorArrayDynamicIndexingARM`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-StorageTensorArrayDynamicIndexingARM)
- [`StorageTensorArrayNonUniformIndexingARM`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-StorageTensorArrayNonUniformIndexingARM)

## [](#_issues)Issues

1\) Should tensor strides be passed in elements or in bytes?

**RESOLVED**: Strides are passed in bytes but are required to be a multiple of the tensor element size. Passing strides in bytes makes it possible to relax this requirement in the future without an interface change. It also makes it easier to describe memory alignment requirements.

2\) Should there be commands to copy data between tensors and buffers/images?

**RESOLVED**: Adding these commands would result in a rather large API surface and not insignificant implementation and validation cost. The same outcome can be achieved with memory aliasing and tensor to tensor copy operations.

3\) Should this extension define transpose and/or other data reorganization operations?

**RESOLVED**: These operations are useful to expose but this extension is only meant to add base support for tensors. Additional operations should be layered on top and defined in other extensions.

4\) Why are tensor strides described using signed integers?

**RESOLVED**: Negative strides make it possible to describe different linear data layouts. While this extension does not allow negative strides, it uses signed integers for strides to make it possible to relax this limitation in future extensions.

## [](#_version_history)Version History

- Revision 1, 2025-06-03 (Kévin Petit)
  
  - Initial revision

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_ARM_tensors).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0