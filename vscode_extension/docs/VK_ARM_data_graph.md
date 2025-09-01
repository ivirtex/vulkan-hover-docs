# VK\_ARM\_data\_graph(3) Manual Page

## Name

VK\_ARM\_data\_graph - device extension



## [](#_registered_extension_number)Registered Extension Number

508

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[Vulkan Version 1.3](#versions-1.3)  
and  
[VK\_KHR\_maintenance5](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance5.html)  
and  
[VK\_KHR\_deferred\_host\_operations](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_deferred_host_operations.html)

## [](#_api_interactions)API Interactions

- Interacts with VK\_ARM\_tensors

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_ARM\_graph](https://github.khronos.org/SPIRV-Registry/extensions/ARM/SPV_ARM_graph.html)

## [](#_contact)Contact

- Kevin Petit [\[GitHub\]kpet](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_ARM_data_graph%5D%20%40kpet%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_ARM_data_graph%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2025-06-18

**Interactions and External Dependencies**

- This extension requires [`SPV_ARM_graph`](https://github.khronos.org/SPIRV-Registry/extensions/EXT/SPV_ARM_graph.html)
- This extension interacts with `VK_EXT_mutable_descriptor_type`
- This extension interacts with `VK_EXT_pipeline_protected_access`
- This extension interacts with `VK_ARM_tensors`
- This extension interacts with `VK_EXT_descriptor_buffer`
- This extension interacts with `VK_KHR_maintenance6`

**IP Status**

No known IP claims.

**Contributors**

- Kévin Petit, Arm Ltd.
- Emma Ben Yossef, Arm Ltd.
- Stefano Bucciarelli, Arm Ltd.
- Marco Cattani, Arm Ltd.
- Aaron DeBattista, Arm Ltd.
- Jan-Harald Fredriksen, Arm Ltd.
- Einar Hov, Arm Ltd.
- Robert Hughes, Arm Ltd.
- Oualid Khelifi, Arm Ltd.
- Derek Lamberti, Arm Ltd.
- Chetan Mistry, Arm Ltd.
- Georgios Teneketzis, Arm Ltd.
- Matthew Netsch, Qualcomm Technologies, Inc

## [](#_description)Description

This extension adds support for a new type of pipeline, data graph pipelines, that provide an encapsulation construct for computational graphs operating on full resources (e.g. ML/AI graphs, image processing pipelines, etc). This extension only supports tensor resources and does not define any operations that can be used within those graphs. These operations will be defined by separate extensions.

## [](#_new_object_types)New Object Types

- [VkDataGraphPipelineSessionARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineSessionARM.html)

## [](#_new_commands)New Commands

- [vkBindDataGraphPipelineSessionMemoryARM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkBindDataGraphPipelineSessionMemoryARM.html)
- [vkCmdDispatchDataGraphARM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDispatchDataGraphARM.html)
- [vkCreateDataGraphPipelineSessionARM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateDataGraphPipelineSessionARM.html)
- [vkCreateDataGraphPipelinesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateDataGraphPipelinesARM.html)
- [vkDestroyDataGraphPipelineSessionARM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkDestroyDataGraphPipelineSessionARM.html)
- [vkGetDataGraphPipelineAvailablePropertiesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDataGraphPipelineAvailablePropertiesARM.html)
- [vkGetDataGraphPipelinePropertiesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDataGraphPipelinePropertiesARM.html)
- [vkGetDataGraphPipelineSessionBindPointRequirementsARM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDataGraphPipelineSessionBindPointRequirementsARM.html)
- [vkGetDataGraphPipelineSessionMemoryRequirementsARM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDataGraphPipelineSessionMemoryRequirementsARM.html)
- [vkGetPhysicalDeviceQueueFamilyDataGraphProcessingEnginePropertiesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceQueueFamilyDataGraphProcessingEnginePropertiesARM.html)
- [vkGetPhysicalDeviceQueueFamilyDataGraphPropertiesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceQueueFamilyDataGraphPropertiesARM.html)

## [](#_new_structures)New Structures

- [VkBindDataGraphPipelineSessionMemoryInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindDataGraphPipelineSessionMemoryInfoARM.html)
- [VkDataGraphPipelineConstantARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineConstantARM.html)
- [VkDataGraphPipelineCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineCreateInfoARM.html)
- [VkDataGraphPipelineDispatchInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineDispatchInfoARM.html)
- [VkDataGraphPipelineInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineInfoARM.html)
- [VkDataGraphPipelinePropertyQueryResultARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelinePropertyQueryResultARM.html)
- [VkDataGraphPipelineResourceInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineResourceInfoARM.html)
- [VkDataGraphPipelineSessionBindPointRequirementARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineSessionBindPointRequirementARM.html)
- [VkDataGraphPipelineSessionBindPointRequirementsInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineSessionBindPointRequirementsInfoARM.html)
- [VkDataGraphPipelineSessionCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineSessionCreateInfoARM.html)
- [VkDataGraphPipelineSessionMemoryRequirementsInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineSessionMemoryRequirementsInfoARM.html)
- [VkPhysicalDeviceDataGraphOperationSupportARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDataGraphOperationSupportARM.html)
- [VkPhysicalDeviceDataGraphProcessingEngineARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDataGraphProcessingEngineARM.html)
- [VkPhysicalDeviceQueueFamilyDataGraphProcessingEngineInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceQueueFamilyDataGraphProcessingEngineInfoARM.html)
- [VkQueueFamilyDataGraphProcessingEnginePropertiesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueFamilyDataGraphProcessingEnginePropertiesARM.html)
- [VkQueueFamilyDataGraphPropertiesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueFamilyDataGraphPropertiesARM.html)
- Extending [VkDataGraphPipelineCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineCreateInfoARM.html):
  
  - [VkDataGraphPipelineCompilerControlCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineCompilerControlCreateInfoARM.html)
  - [VkDataGraphPipelineIdentifierCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineIdentifierCreateInfoARM.html)
  - [VkDataGraphPipelineShaderModuleCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineShaderModuleCreateInfoARM.html)
- Extending [VkDataGraphPipelineCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineCreateInfoARM.html), [VkDescriptorPoolCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorPoolCreateInfo.html), [VkCommandPoolCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandPoolCreateInfo.html):
  
  - [VkDataGraphProcessingEngineCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphProcessingEngineCreateInfoARM.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceDataGraphFeaturesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDataGraphFeaturesARM.html)

If [VK\_ARM\_tensors](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_tensors.html) is supported:

- Extending [VkDataGraphPipelineConstantARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineConstantARM.html):
  
  - [VkDataGraphPipelineConstantTensorSemiStructuredSparsityInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineConstantTensorSemiStructuredSparsityInfoARM.html)

## [](#_new_enums)New Enums

- [VkDataGraphPipelineDispatchFlagBitsARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineDispatchFlagBitsARM.html)
- [VkDataGraphPipelinePropertyARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelinePropertyARM.html)
- [VkDataGraphPipelineSessionBindPointARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineSessionBindPointARM.html)
- [VkDataGraphPipelineSessionBindPointTypeARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineSessionBindPointTypeARM.html)
- [VkDataGraphPipelineSessionCreateFlagBitsARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineSessionCreateFlagBitsARM.html)
- [VkPhysicalDeviceDataGraphOperationTypeARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDataGraphOperationTypeARM.html)
- [VkPhysicalDeviceDataGraphProcessingEngineTypeARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDataGraphProcessingEngineTypeARM.html)

## [](#_new_bitmasks)New Bitmasks

- [VkDataGraphPipelineDispatchFlagsARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineDispatchFlagsARM.html)
- [VkDataGraphPipelineSessionCreateFlagsARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineSessionCreateFlagsARM.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_ARM_DATA_GRAPH_EXTENSION_NAME`
- `VK_ARM_DATA_GRAPH_SPEC_VERSION`
- `VK_MAX_PHYSICAL_DEVICE_DATA_GRAPH_OPERATION_SET_NAME_SIZE_ARM`
- Extending [VkAccessFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlagBits2.html):
  
  - `VK_ACCESS_2_DATA_GRAPH_READ_BIT_ARM`
  - `VK_ACCESS_2_DATA_GRAPH_WRITE_BIT_ARM`
- Extending [VkBufferUsageFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferUsageFlagBits2.html):
  
  - `VK_BUFFER_USAGE_2_DATA_GRAPH_FOREIGN_DESCRIPTOR_BIT_ARM`
- Extending [VkFormatFeatureFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatFeatureFlagBits2.html):
  
  - `VK_FORMAT_FEATURE_2_TENSOR_DATA_GRAPH_BIT_ARM`
- Extending [VkObjectType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkObjectType.html):
  
  - `VK_OBJECT_TYPE_DATA_GRAPH_PIPELINE_SESSION_ARM`
- Extending [VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBindPoint.html):
  
  - `VK_PIPELINE_BIND_POINT_DATA_GRAPH_ARM`
- Extending [VkPipelineStageFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlagBits2.html):
  
  - `VK_PIPELINE_STAGE_2_DATA_GRAPH_BIT_ARM`
- Extending [VkQueueFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueFlagBits.html):
  
  - `VK_QUEUE_DATA_GRAPH_BIT_ARM`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_BIND_DATA_GRAPH_PIPELINE_SESSION_MEMORY_INFO_ARM`
  - `VK_STRUCTURE_TYPE_DATA_GRAPH_PIPELINE_COMPILER_CONTROL_CREATE_INFO_ARM`
  - `VK_STRUCTURE_TYPE_DATA_GRAPH_PIPELINE_CONSTANT_ARM`
  - `VK_STRUCTURE_TYPE_DATA_GRAPH_PIPELINE_CREATE_INFO_ARM`
  - `VK_STRUCTURE_TYPE_DATA_GRAPH_PIPELINE_DISPATCH_INFO_ARM`
  - `VK_STRUCTURE_TYPE_DATA_GRAPH_PIPELINE_IDENTIFIER_CREATE_INFO_ARM`
  - `VK_STRUCTURE_TYPE_DATA_GRAPH_PIPELINE_INFO_ARM`
  - `VK_STRUCTURE_TYPE_DATA_GRAPH_PIPELINE_PROPERTY_QUERY_RESULT_ARM`
  - `VK_STRUCTURE_TYPE_DATA_GRAPH_PIPELINE_RESOURCE_INFO_ARM`
  - `VK_STRUCTURE_TYPE_DATA_GRAPH_PIPELINE_SESSION_BIND_POINT_REQUIREMENTS_INFO_ARM`
  - `VK_STRUCTURE_TYPE_DATA_GRAPH_PIPELINE_SESSION_BIND_POINT_REQUIREMENT_ARM`
  - `VK_STRUCTURE_TYPE_DATA_GRAPH_PIPELINE_SESSION_CREATE_INFO_ARM`
  - `VK_STRUCTURE_TYPE_DATA_GRAPH_PIPELINE_SESSION_MEMORY_REQUIREMENTS_INFO_ARM`
  - `VK_STRUCTURE_TYPE_DATA_GRAPH_PIPELINE_SHADER_MODULE_CREATE_INFO_ARM`
  - `VK_STRUCTURE_TYPE_DATA_GRAPH_PROCESSING_ENGINE_CREATE_INFO_ARM`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DATA_GRAPH_FEATURES_ARM`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_QUEUE_FAMILY_DATA_GRAPH_PROCESSING_ENGINE_INFO_ARM`
  - `VK_STRUCTURE_TYPE_QUEUE_FAMILY_DATA_GRAPH_PROCESSING_ENGINE_PROPERTIES_ARM`
  - `VK_STRUCTURE_TYPE_QUEUE_FAMILY_DATA_GRAPH_PROPERTIES_ARM`
- Extending [VkTensorUsageFlagBitsARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorUsageFlagBitsARM.html):
  
  - `VK_TENSOR_USAGE_DATA_GRAPH_BIT_ARM`

If [VK\_ARM\_tensors](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_tensors.html) is supported:

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_DATA_GRAPH_PIPELINE_CONSTANT_TENSOR_SEMI_STRUCTURED_SPARSITY_INFO_ARM`

## [](#_new_spir_v_capabilities)New SPIR-V Capabilities

- [GraphARM](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-GraphARM)

## [](#_issues)Issues

1\) Should graph pipeline resource info structures be integrated into pipeline layouts? Would a new graph pipeline layout be a better fit?

**RESOLVED**: Graph pipeline resource info are passed separately at pipeline creation time.

2\) Do we need a new shader stage for graph pipelines for use in creating descriptor set layouts?

**RESOLVED**: Currently using `VK_SHADER_STAGE_ALL`.

3\) Should this extension provide applications with a way of knowing which combinations of sparsity information implementations can take advantage of when processing graph constants?

**RESOLVED**: No. Describing the exact combinations is in some cases complex and it is always valid for implementations to ignore the sparsity information and treat the data as dense. Specific implementations can provide guidance to application writers if they so desire and applications are encouraged to always provide sparsity information that they have.

## [](#_version_history)Version History

- Revision 1, 2025-06-18 (Kévin Petit)
  
  - Initial revision

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_ARM_data_graph).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0