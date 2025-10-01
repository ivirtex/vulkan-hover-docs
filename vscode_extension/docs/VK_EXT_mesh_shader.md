# VK\_EXT\_mesh\_shader(3) Manual Page

## Name

VK\_EXT\_mesh\_shader - device extension



## [](#_registered_extension_number)Registered Extension Number

329

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_spirv\_1\_4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_spirv_1_4.html)  
or  
[Vulkan Version 1.2](#versions-1.2)

## [](#_api_interactions)API Interactions

- Interacts with VK\_VERSION\_1\_2
- Interacts with VK\_AMD\_draw\_indirect\_count
- Interacts with VK\_EXT\_device\_generated\_commands
- Interacts with VK\_KHR\_draw\_indirect\_count
- Interacts with VK\_KHR\_fragment\_shading\_rate
- Interacts with VK\_NV\_device\_generated\_commands
- Interacts with VkPhysicalDeviceMeshShaderFeaturesEXT::primitiveFragmentShadingRateMeshShader

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_EXT\_mesh\_shader](https://github.khronos.org/SPIRV-Registry/extensions/EXT/SPV_EXT_mesh_shader.html)

## [](#_contact)Contact

- Christoph Kubisch [\[GitHub\]pixeljetstream](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_mesh_shader%5D%20%40pixeljetstream%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_mesh_shader%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_EXT\_mesh\_shader](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_EXT_mesh_shader.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2022-01-20

**Interactions and External Dependencies**

- This extension provides API support for [`GLSL_EXT_mesh_shader`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/ext/GLSL_EXT_mesh_shader.txt)
- Interacts with Vulkan 1.1
- Interacts with `VK_KHR_multiview`
- Interacts with `VK_KHR_fragment_shading_rate`

**Contributors**

- Christoph Kubisch, NVIDIA
- Pat Brown, NVIDIA
- Jeff Bolz, NVIDIA
- Daniel Koch, NVIDIA
- Piers Daniell, NVIDIA
- Pierre Boudier, NVIDIA
- Patrick Mours, NVIDIA
- David Zhao Akeley, NVIDIA
- Kedarnath Thangudu, NVIDIA
- Timur Kristóf, Valve
- Hans-Kristian Arntzen, Valve
- Philip Rebohle, Valve
- Mike Blumenkrantz, Valve
- Slawomir Grajewski, Intel
- Michal Pietrasiuk, Intel
- Mariusz Merecki, Intel
- Tom Olson, ARM
- Jan-Harald Fredriksen, ARM
- Sandeep Kakarlapudi, ARM
- Ruihao Zhang, QUALCOMM
- Ricardo Garcia, Igalia, S.L.
- Tobias Hector, AMD
- Stu Smith, AMD

## [](#_description)Description

This extension provides a new mechanism allowing applications to generate collections of geometric primitives via programmable mesh shading. It is an alternative to the existing programmable primitive shading pipeline, which relied on generating input primitives by a fixed function assembler as well as fixed function vertex fetch.

This extension also adds support for the following SPIR-V extension in Vulkan:

- [`SPV_EXT_mesh_shader`](https://github.khronos.org/SPIRV-Registry/extensions/EXT/SPV_EXT_mesh_shader.html)

## [](#_new_commands)New Commands

- [vkCmdDrawMeshTasksEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDrawMeshTasksEXT.html)
- [vkCmdDrawMeshTasksIndirectEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDrawMeshTasksIndirectEXT.html)

If [Vulkan Version 1.2](#versions-1.2) or [VK\_KHR\_draw\_indirect\_count](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_draw_indirect_count.html) or [VK\_AMD\_draw\_indirect\_count](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_AMD_draw_indirect_count.html) is supported:

- [vkCmdDrawMeshTasksIndirectCountEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDrawMeshTasksIndirectCountEXT.html)

## [](#_new_structures)New Structures

- [VkDrawMeshTasksIndirectCommandEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDrawMeshTasksIndirectCommandEXT.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceMeshShaderFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderFeaturesEXT.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceMeshShaderPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_MESH_SHADER_EXTENSION_NAME`
- `VK_EXT_MESH_SHADER_SPEC_VERSION`
- Extending [VkPipelineStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlagBits.html):
  
  - `VK_PIPELINE_STAGE_MESH_SHADER_BIT_EXT`
  - `VK_PIPELINE_STAGE_TASK_SHADER_BIT_EXT`
- Extending [VkQueryPipelineStatisticFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPipelineStatisticFlagBits.html):
  
  - `VK_QUERY_PIPELINE_STATISTIC_MESH_SHADER_INVOCATIONS_BIT_EXT`
  - `VK_QUERY_PIPELINE_STATISTIC_TASK_SHADER_INVOCATIONS_BIT_EXT`
- Extending [VkQueryType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryType.html):
  
  - `VK_QUERY_TYPE_MESH_PRIMITIVES_GENERATED_EXT`
- Extending [VkShaderStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderStageFlagBits.html):
  
  - `VK_SHADER_STAGE_MESH_BIT_EXT`
  - `VK_SHADER_STAGE_TASK_BIT_EXT`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MESH_SHADER_FEATURES_EXT`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MESH_SHADER_PROPERTIES_EXT`

If [VK\_EXT\_device\_generated\_commands](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_device_generated_commands.html) is supported:

- Extending [VkIndirectCommandsTokenTypeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsTokenTypeEXT.html):
  
  - `VK_INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_MESH_TASKS_COUNT_EXT`
  - `VK_INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_MESH_TASKS_EXT`

If [VK\_NV\_device\_generated\_commands](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_device_generated_commands.html) is supported:

- Extending [VkIndirectCommandsTokenTypeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsTokenTypeNV.html):
  
  - `VK_INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_MESH_TASKS_NV`

## [](#_new_or_modified_built_in_variables)New or Modified Built-In Variables

- [CullPrimitiveEXT](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-cullprimitive)
- [PrimitivePointIndicesEXT](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-primitivepointindices)
- [PrimitiveLineIndicesEXT](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-primitivelineindices)
- [PrimitiveTriangleIndicesEXT](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-primitivetriangleindices)
- (modified)`Position`
- (modified)`PointSize`
- (modified)`ClipDistance`
- (modified)`CullDistance`
- (modified)`PrimitiveId`
- (modified)`Layer`
- (modified)`ViewportIndex`
- (modified)`NumWorkgroups`
- (modified)`WorkgroupSize`
- (modified)`WorkgroupId`
- (modified)`LocalInvocationId`
- (modified)`GlobalInvocationId`
- (modified)`LocalInvocationIndex`
- (modified)`NumSubgroups`
- (modified)`SubgroupId`
- (modified)`DrawIndex`
- (modified)`PrimitiveShadingRateKHR`
- (modified)`ViewIndex`

## [](#_new_spir_v_capability)New SPIR-V Capability

- [`MeshShadingEXT`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-MeshShadingEXT)

## [](#_version_history)Version History

- Revision 1, 2022-03-08 (Christoph Kubisch, Daniel Koch, Patrick Mours)
  
  - Initial revision

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_mesh_shader).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0