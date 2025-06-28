# VK\_AMDX\_shader\_enqueue(3) Manual Page

## Name

VK\_AMDX\_shader\_enqueue - device extension



## [](#_registered_extension_number)Registered Extension Number

135

## [](#_revision)Revision

2

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

         [VK\_KHR\_synchronization2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_synchronization2.html)  
         and  
         [VK\_KHR\_spirv\_1\_4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_spirv_1_4.html)  
         and  
         [VK\_EXT\_extended\_dynamic\_state](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_extended_dynamic_state.html)  
     or  
     [Vulkan Version 1.3](#versions-1.3)  
and  
[VK\_KHR\_maintenance5](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance5.html)  
and  
[VK\_KHR\_pipeline\_library](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_pipeline_library.html)

- **This is a *provisional* extension and must be used with caution. See the [description](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#boilerplate-provisional-header) of provisional header files for enablement and stability details.**

## [](#_api_interactions)API Interactions

- Interacts with VK\_VERSION\_1\_4
- Interacts with VK\_EXT\_mesh\_shader
- Interacts with VK\_KHR\_maintenance5

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_AMDX\_shader\_enqueue](https://github.khronos.org/SPIRV-Registry/extensions/AMDX/SPV_AMDX_shader_enqueue.html)

## [](#_contact)Contact

- Tobias Hector [\[GitHub\]tobski](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_AMDX_shader_enqueue%5D%20%40tobski%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_AMDX_shader_enqueue%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_AMDX\_shader\_enqueue](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_AMDX_shader_enqueue.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2024-07-17

**Provisional**

**This extension is *provisional* and should not be used in production applications. The functionality may change in ways that break backwards compatibility between revisions, and before final release.**

**Contributors**

- Tobias Hector, AMD
- Matthaeus Chajdas, AMD
- Maciej Jesionowski, AMD
- Robert Martin, AMD
- Qun Lin, AMD
- Rex Xu, AMD
- Dominik Witczak, AMD
- Karthik Srinivasan, AMD
- Nicolai Haehnle, AMD
- Stuart Smith, AMD

## [](#_description)Description

This extension adds the ability for developers to enqueue mesh and compute shader workgroups from other compute shaders.

## [](#_new_commands)New Commands

- [vkCmdDispatchGraphAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDispatchGraphAMDX.html)
- [vkCmdDispatchGraphIndirectAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDispatchGraphIndirectAMDX.html)
- [vkCmdDispatchGraphIndirectCountAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDispatchGraphIndirectCountAMDX.html)
- [vkCmdInitializeGraphScratchMemoryAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdInitializeGraphScratchMemoryAMDX.html)
- [vkCreateExecutionGraphPipelinesAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateExecutionGraphPipelinesAMDX.html)
- [vkGetExecutionGraphPipelineNodeIndexAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetExecutionGraphPipelineNodeIndexAMDX.html)
- [vkGetExecutionGraphPipelineScratchSizeAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetExecutionGraphPipelineScratchSizeAMDX.html)

## [](#_new_structures)New Structures

- [VkDispatchGraphCountInfoAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDispatchGraphCountInfoAMDX.html)
- [VkDispatchGraphInfoAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDispatchGraphInfoAMDX.html)
- [VkExecutionGraphPipelineCreateInfoAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExecutionGraphPipelineCreateInfoAMDX.html)
- [VkExecutionGraphPipelineScratchSizeAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExecutionGraphPipelineScratchSizeAMDX.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceShaderEnqueueFeaturesAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceShaderEnqueueFeaturesAMDX.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceShaderEnqueuePropertiesAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceShaderEnqueuePropertiesAMDX.html)
- Extending [VkPipelineShaderStageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineShaderStageCreateInfo.html):
  
  - [VkPipelineShaderStageNodeCreateInfoAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineShaderStageNodeCreateInfoAMDX.html)

## [](#_new_unions)New Unions

- [VkDeviceOrHostAddressConstAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceOrHostAddressConstAMDX.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_AMDX_SHADER_ENQUEUE_EXTENSION_NAME`
- `VK_AMDX_SHADER_ENQUEUE_SPEC_VERSION`
- `VK_SHADER_INDEX_UNUSED_AMDX`
- Extending [VkBufferUsageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferUsageFlagBits.html):
  
  - `VK_BUFFER_USAGE_EXECUTION_GRAPH_SCRATCH_BIT_AMDX`
- Extending [VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBindPoint.html):
  
  - `VK_PIPELINE_BIND_POINT_EXECUTION_GRAPH_AMDX`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_EXECUTION_GRAPH_PIPELINE_CREATE_INFO_AMDX`
  - `VK_STRUCTURE_TYPE_EXECUTION_GRAPH_PIPELINE_SCRATCH_SIZE_AMDX`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_ENQUEUE_FEATURES_AMDX`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_ENQUEUE_PROPERTIES_AMDX`
  - `VK_STRUCTURE_TYPE_PIPELINE_SHADER_STAGE_NODE_CREATE_INFO_AMDX`

If [VK\_KHR\_maintenance5](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance5.html) or [Vulkan Version 1.4](#versions-1.4) is supported:

- Extending [VkBufferUsageFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferUsageFlagBits2.html):
  
  - `VK_BUFFER_USAGE_2_EXECUTION_GRAPH_SCRATCH_BIT_AMDX`
- Extending [VkPipelineCreateFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlagBits2.html):
  
  - `VK_PIPELINE_CREATE_2_EXECUTION_GRAPH_BIT_AMDX`

## [](#_version_history)Version History

- Revision 2, 2024-07-17 (Tobias Hector)
  
  - Add mesh nodes
- Revision 1, 2021-07-22 (Tobias Hector)
  
  - Initial revision

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_AMDX_shader_enqueue)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0