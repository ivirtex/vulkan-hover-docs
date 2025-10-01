# VK\_NV\_mesh\_shader(3) Manual Page

## Name

VK\_NV\_mesh\_shader - device extension



## [](#_registered_extension_number)Registered Extension Number

203

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_api_interactions)API Interactions

- Interacts with VK\_VERSION\_1\_2
- Interacts with VK\_AMD\_draw\_indirect\_count
- Interacts with VK\_EXT\_device\_generated\_commands
- Interacts with VK\_KHR\_draw\_indirect\_count

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_NV\_mesh\_shader](https://github.khronos.org/SPIRV-Registry/extensions/NV/SPV_NV_mesh_shader.html)

## [](#_contact)Contact

- Christoph Kubisch [\[GitHub\]pixeljetstream](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_mesh_shader%5D%20%40pixeljetstream%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_mesh_shader%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2018-07-19

**Interactions and External Dependencies**

- This extension provides API support for [`GLSL_NV_mesh_shader`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/nv/GLSL_NV_mesh_shader.txt)

**Contributors**

- Pat Brown, NVIDIA
- Jeff Bolz, NVIDIA
- Daniel Koch, NVIDIA
- Piers Daniell, NVIDIA
- Pierre Boudier, NVIDIA

## [](#_description)Description

This extension provides a new mechanism allowing applications to generate collections of geometric primitives via programmable mesh shading. It is an alternative to the existing programmable primitive shading pipeline, which relied on generating input primitives by a fixed function assembler as well as fixed function vertex fetch.

There are new programmable shader types — the task and mesh shader — to generate these collections to be processed by fixed-function primitive assembly and rasterization logic. When task and mesh shaders are dispatched, they replace the core [pre-rasterization stages](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization), including vertex array attribute fetching, vertex shader processing, tessellation, and geometry shader processing.

This extension also adds support for the following SPIR-V extension in Vulkan:

- [`SPV_NV_mesh_shader`](https://github.khronos.org/SPIRV-Registry/extensions/NV/SPV_NV_mesh_shader.html)

## [](#_new_commands)New Commands

- [vkCmdDrawMeshTasksIndirectNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDrawMeshTasksIndirectNV.html)
- [vkCmdDrawMeshTasksNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDrawMeshTasksNV.html)

If [Vulkan Version 1.2](#versions-1.2) or [VK\_KHR\_draw\_indirect\_count](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_draw_indirect_count.html) or [VK\_AMD\_draw\_indirect\_count](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_AMD_draw_indirect_count.html) is supported:

- [vkCmdDrawMeshTasksIndirectCountNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDrawMeshTasksIndirectCountNV.html)

## [](#_new_structures)New Structures

- [VkDrawMeshTasksIndirectCommandNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDrawMeshTasksIndirectCommandNV.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceMeshShaderFeaturesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderFeaturesNV.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceMeshShaderPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMeshShaderPropertiesNV.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_NV_MESH_SHADER_EXTENSION_NAME`
- `VK_NV_MESH_SHADER_SPEC_VERSION`
- Extending [VkPipelineStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlagBits.html):
  
  - `VK_PIPELINE_STAGE_MESH_SHADER_BIT_NV`
  - `VK_PIPELINE_STAGE_TASK_SHADER_BIT_NV`
- Extending [VkShaderStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderStageFlagBits.html):
  
  - `VK_SHADER_STAGE_MESH_BIT_NV`
  - `VK_SHADER_STAGE_TASK_BIT_NV`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MESH_SHADER_FEATURES_NV`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MESH_SHADER_PROPERTIES_NV`

If [VK\_EXT\_device\_generated\_commands](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_device_generated_commands.html) is supported:

- Extending [VkIndirectCommandsTokenTypeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsTokenTypeEXT.html):
  
  - `VK_INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_MESH_TASKS_COUNT_NV_EXT`
  - `VK_INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_MESH_TASKS_NV_EXT`

## [](#_new_or_modified_built_in_variables)New or Modified Built-In Variables

- [TaskCountNV](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-taskcount)
- [PrimitiveCountNV](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-primitivecount)
- [PrimitiveIndicesNV](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-primitiveindices)
- [ClipDistancePerViewNV](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-clipdistancepv)
- [CullDistancePerViewNV](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-culldistancepv)
- [LayerPerViewNV](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-layerpv)
- [MeshViewCountNV](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-meshviewcount)
- [MeshViewIndicesNV](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-meshviewindices)
- (modified)`Position`
- (modified)`PointSize`
- (modified)`ClipDistance`
- (modified)`CullDistance`
- (modified)`PrimitiveId`
- (modified)`Layer`
- (modified)`ViewportIndex`
- (modified)`WorkgroupSize`
- (modified)`WorkgroupId`
- (modified)`LocalInvocationId`
- (modified)`GlobalInvocationId`
- (modified)`LocalInvocationIndex`
- (modified)`DrawIndex`
- (modified)`ViewportMaskNV`
- (modified)`PositionPerViewNV`
- (modified)`ViewportMaskPerViewNV`

## [](#_new_spir_v_capability)New SPIR-V Capability

- [`MeshShadingNV`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-MeshShadingNV)

## [](#_issues)Issues

1. How to name this extension?
   
   **RESOLVED**: VK\_NV\_mesh\_shader
   
   Other options considered:
   
   - VK\_NV\_mesh\_shading
   - VK\_NV\_programmable\_mesh\_shading
   - VK\_NV\_primitive\_group\_shading
   - VK\_NV\_grouped\_drawing
2. Do we need a new VkPrimitiveTopology?
   
   **RESOLVED**: No. We skip the InputAssembler stage.
3. Should we allow Instancing?
   
   **RESOLVED**: No. There is no fixed function input, other than the IDs. However, allow offsetting with a “first” value.
4. Should we use existing vkCmdDraw or introduce new functions?
   
   **RESOLVED**: Introduce new functions.
   
   New functions make it easier to separate from “programmable primitive shading” chapter, less “dual use” language about existing functions having alternative behavior. The text around the existing “draws” is heavily based around emitting vertices.
5. If new functions, how to name?
   
   **RESOLVED**: CmdDrawMeshTasks*
   
   Other options considered:
   
   - CmdDrawMeshed
   - CmdDrawTasked
   - CmdDrawGrouped
6. Should VK\_SHADER\_STAGE\_ALL\_GRAPHICS be updated to include the new stages?
   
   **RESOLVED**: No. If an application were to be recompiled with headers that include additional shader stage bits in VK\_SHADER\_STAGE\_ALL\_GRAPHICS, then the previously valid application would no longer be valid on implementations that do not support mesh or task shaders. This means the change would not be backwards compatible. It is too bad VkShaderStageFlagBits does not have a dedicated “all supported graphics stages” bit like VK\_PIPELINE\_STAGE\_ALL\_GRAPHICS\_BIT, which would have avoided this problem.

## [](#_version_history)Version History

- Revision 1, 2018-07-19 (Christoph Kubisch, Daniel Koch)
  
  - Internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_NV_mesh_shader).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0