# VK_NV_mesh_shader(3) Manual Page

## Name

VK_NV_mesh_shader - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

203

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_spir_v_dependencies" class="anchor"></a>SPIR-V Dependencies

- [SPV_NV_mesh_shader](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/NV/SPV_NV_mesh_shader.html)

## <a href="#_contact" class="anchor"></a>Contact

- Christoph Kubisch <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_mesh_shader%5D%20@pixeljetstream%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_mesh_shader%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>pixeljetstream</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2018-07-19

**Interactions and External Dependencies**  
- This extension provides API support for
  [`GLSL_NV_mesh_shader`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/nv/GLSL_NV_mesh_shader.txt)

**Contributors**  
- Pat Brown, NVIDIA

- Jeff Bolz, NVIDIA

- Daniel Koch, NVIDIA

- Piers Daniell, NVIDIA

- Pierre Boudier, NVIDIA

## <a href="#_description" class="anchor"></a>Description

This extension provides a new mechanism allowing applications to
generate collections of geometric primitives via programmable mesh
shading. It is an alternative to the existing programmable primitive
shading pipeline, which relied on generating input primitives by a fixed
function assembler as well as fixed function vertex fetch.

There are new programmable shader types — the task and mesh shader — to
generate these collections to be processed by fixed-function primitive
assembly and rasterization logic. When task and mesh shaders are
dispatched, they replace the core <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
target="_blank" rel="noopener">pre-rasterization stages</a>, including
vertex array attribute fetching, vertex shader processing, tessellation,
and geometry shader processing.

This extension also adds support for the following SPIR-V extension in
Vulkan:

- [`SPV_NV_mesh_shader`](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/NV/SPV_NV_mesh_shader.html)

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkCmdDrawMeshTasksIndirectCountNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDrawMeshTasksIndirectCountNV.html)

- [vkCmdDrawMeshTasksIndirectNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDrawMeshTasksIndirectNV.html)

- [vkCmdDrawMeshTasksNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDrawMeshTasksNV.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkDrawMeshTasksIndirectCommandNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDrawMeshTasksIndirectCommandNV.html)

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceMeshShaderFeaturesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMeshShaderFeaturesNV.html)

- Extending
  [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html):

  - [VkPhysicalDeviceMeshShaderPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMeshShaderPropertiesNV.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_NV_MESH_SHADER_EXTENSION_NAME`

- `VK_NV_MESH_SHADER_SPEC_VERSION`

- Extending [VkPipelineStageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlagBits.html):

  - `VK_PIPELINE_STAGE_MESH_SHADER_BIT_NV`

  - `VK_PIPELINE_STAGE_TASK_SHADER_BIT_NV`

- Extending [VkShaderStageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderStageFlagBits.html):

  - `VK_SHADER_STAGE_MESH_BIT_NV`

  - `VK_SHADER_STAGE_TASK_BIT_NV`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MESH_SHADER_FEATURES_NV`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MESH_SHADER_PROPERTIES_NV`

## <a href="#_new_or_modified_built_in_variables" class="anchor"></a>New or Modified Built-In Variables

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-taskcount"
  target="_blank" rel="noopener">TaskCountNV</a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-primitivecount"
  target="_blank" rel="noopener">PrimitiveCountNV</a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-primitiveindices"
  target="_blank" rel="noopener">PrimitiveIndicesNV</a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-clipdistancepv"
  target="_blank" rel="noopener">ClipDistancePerViewNV</a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-culldistancepv"
  target="_blank" rel="noopener">CullDistancePerViewNV</a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-layerpv"
  target="_blank" rel="noopener">LayerPerViewNV</a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-meshviewcount"
  target="_blank" rel="noopener">MeshViewCountNV</a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-meshviewindices"
  target="_blank" rel="noopener">MeshViewIndicesNV</a>

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

## <a href="#_new_spir_v_capability" class="anchor"></a>New SPIR-V Capability

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirvenv-capabilities-table-MeshShadingNV"
  target="_blank" rel="noopener"><code>MeshShadingNV</code></a>

## <a href="#_issues" class="anchor"></a>Issues

1.  How to name this extension?

    **RESOLVED**: VK_NV_mesh_shader

    Other options considered:

    - VK_NV_mesh_shading

    - VK_NV_programmable_mesh_shading

    - VK_NV_primitive_group_shading

    - VK_NV_grouped_drawing

2.  Do we need a new VkPrimitiveTopology?

    **RESOLVED**: No. We skip the InputAssembler stage.

3.  Should we allow Instancing?

    **RESOLVED**: No. There is no fixed function input, other than the
    IDs. However, allow offsetting with a “first” value.

4.  Should we use existing vkCmdDraw or introduce new functions?

    **RESOLVED**: Introduce new functions.

    New functions make it easier to separate from “programmable
    primitive shading” chapter, less “dual use” language about existing
    functions having alternative behavior. The text around the existing
    “draws” is heavily based around emitting vertices.

5.  If new functions, how to name?

    **RESOLVED**: CmdDrawMeshTasks\*

    Other options considered:

    - CmdDrawMeshed

    - CmdDrawTasked

    - CmdDrawGrouped

6.  Should VK_SHADER_STAGE_ALL_GRAPHICS be updated to include the new
    stages?

    **RESOLVED**: No. If an application were to be recompiled with
    headers that include additional shader stage bits in
    VK_SHADER_STAGE_ALL_GRAPHICS, then the previously valid application
    would no longer be valid on implementations that do not support mesh
    or task shaders. This means the change would not be backwards
    compatible. It is too bad VkShaderStageFlagBits does not have a
    dedicated “all supported graphics stages” bit like
    VK_PIPELINE_STAGE_ALL_GRAPHICS_BIT, which would have avoided this
    problem.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2018-07-19 (Christoph Kubisch, Daniel Koch)

  - Internal revisions

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_NV_mesh_shader"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
