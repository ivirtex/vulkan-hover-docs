# VK_EXT_mesh_shader(3) Manual Page

## Name

VK_EXT_mesh_shader - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

329

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_spirv_1_4](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_spirv_1_4.html)  

## <a href="#_api_interactions" class="anchor"></a>API Interactions

- Interacts with VK_NV_device_generated_commands

## <a href="#_spir_v_dependencies" class="anchor"></a>SPIR-V Dependencies

- [SPV_EXT_mesh_shader](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/EXT/SPV_EXT_mesh_shader.html)

## <a href="#_contact" class="anchor"></a>Contact

- Christoph Kubisch <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_mesh_shader%5D%20@pixeljetstream%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_mesh_shader%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>pixeljetstream</a>

## <a href="#_extension_proposal" class="anchor"></a>Extension Proposal

[VK_EXT_mesh_shader](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_EXT_mesh_shader.adoc)

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2022-01-20

**Interactions and External Dependencies**  
- This extension provides API support for
  [`GLSL_EXT_mesh_shader`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/ext/GLSL_EXT_mesh_shader.txt)

- Interacts with Vulkan 1.1

- Interacts with [`VK_KHR_multiview`](VK_KHR_multiview.html)

- Interacts with
  [`VK_KHR_fragment_shading_rate`](VK_KHR_fragment_shading_rate.html)

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

## <a href="#_description" class="anchor"></a>Description

This extension provides a new mechanism allowing applications to
generate collections of geometric primitives via programmable mesh
shading. It is an alternative to the existing programmable primitive
shading pipeline, which relied on generating input primitives by a fixed
function assembler as well as fixed function vertex fetch.

This extension also adds support for the following SPIR-V extension in
Vulkan:

- [`SPV_EXT_mesh_shader`](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/EXT/SPV_EXT_mesh_shader.html)

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkCmdDrawMeshTasksEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDrawMeshTasksEXT.html)

- [vkCmdDrawMeshTasksIndirectCountEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDrawMeshTasksIndirectCountEXT.html)

- [vkCmdDrawMeshTasksIndirectEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDrawMeshTasksIndirectEXT.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkDrawMeshTasksIndirectCommandEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDrawMeshTasksIndirectCommandEXT.html)

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceMeshShaderFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMeshShaderFeaturesEXT.html)

- Extending
  [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html):

  - [VkPhysicalDeviceMeshShaderPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMeshShaderPropertiesEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_MESH_SHADER_EXTENSION_NAME`

- `VK_EXT_MESH_SHADER_SPEC_VERSION`

- Extending [VkPipelineStageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlagBits.html):

  - `VK_PIPELINE_STAGE_MESH_SHADER_BIT_EXT`

  - `VK_PIPELINE_STAGE_TASK_SHADER_BIT_EXT`

- Extending
  [VkQueryPipelineStatisticFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPipelineStatisticFlagBits.html):

  - `VK_QUERY_PIPELINE_STATISTIC_MESH_SHADER_INVOCATIONS_BIT_EXT`

  - `VK_QUERY_PIPELINE_STATISTIC_TASK_SHADER_INVOCATIONS_BIT_EXT`

- Extending [VkQueryType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryType.html):

  - `VK_QUERY_TYPE_MESH_PRIMITIVES_GENERATED_EXT`

- Extending [VkShaderStageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderStageFlagBits.html):

  - `VK_SHADER_STAGE_MESH_BIT_EXT`

  - `VK_SHADER_STAGE_TASK_BIT_EXT`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MESH_SHADER_FEATURES_EXT`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MESH_SHADER_PROPERTIES_EXT`

If
[VK_NV_device_generated_commands](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_device_generated_commands.html)
is supported:

- Extending
  [VkIndirectCommandsTokenTypeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIndirectCommandsTokenTypeNV.html):

  - `VK_INDIRECT_COMMANDS_TOKEN_TYPE_DRAW_MESH_TASKS_NV`

## <a href="#_new_or_modified_built_in_variables" class="anchor"></a>New or Modified Built-In Variables

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-cullprimitive"
  target="_blank" rel="noopener">CullPrimitiveEXT</a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-primitivepointindices"
  target="_blank" rel="noopener">PrimitivePointIndicesEXT</a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-primitivelineindices"
  target="_blank" rel="noopener">PrimitiveLineIndicesEXT</a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-primitivetriangleindices"
  target="_blank" rel="noopener">PrimitiveTriangleIndicesEXT</a>

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

## <a href="#_new_spir_v_capability" class="anchor"></a>New SPIR-V Capability

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirvenv-capabilities-table-MeshShadingEXT"
  target="_blank" rel="noopener"><code>MeshShadingEXT</code></a>

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2022-03-08 (Christoph Kubisch, Daniel Koch, Patrick Mours)

  - Initial revision

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_mesh_shader"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
