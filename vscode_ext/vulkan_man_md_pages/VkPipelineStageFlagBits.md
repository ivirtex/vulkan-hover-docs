# VkPipelineStageFlagBits(3) Manual Page

## Name

VkPipelineStageFlagBits - Bitmask specifying pipeline stages



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **can** be set in a
[VkPipelineStageFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlags.html) mask, specifying
stages of execution, are:

``` c
// Provided by VK_VERSION_1_0
typedef enum VkPipelineStageFlagBits {
    VK_PIPELINE_STAGE_TOP_OF_PIPE_BIT = 0x00000001,
    VK_PIPELINE_STAGE_DRAW_INDIRECT_BIT = 0x00000002,
    VK_PIPELINE_STAGE_VERTEX_INPUT_BIT = 0x00000004,
    VK_PIPELINE_STAGE_VERTEX_SHADER_BIT = 0x00000008,
    VK_PIPELINE_STAGE_TESSELLATION_CONTROL_SHADER_BIT = 0x00000010,
    VK_PIPELINE_STAGE_TESSELLATION_EVALUATION_SHADER_BIT = 0x00000020,
    VK_PIPELINE_STAGE_GEOMETRY_SHADER_BIT = 0x00000040,
    VK_PIPELINE_STAGE_FRAGMENT_SHADER_BIT = 0x00000080,
    VK_PIPELINE_STAGE_EARLY_FRAGMENT_TESTS_BIT = 0x00000100,
    VK_PIPELINE_STAGE_LATE_FRAGMENT_TESTS_BIT = 0x00000200,
    VK_PIPELINE_STAGE_COLOR_ATTACHMENT_OUTPUT_BIT = 0x00000400,
    VK_PIPELINE_STAGE_COMPUTE_SHADER_BIT = 0x00000800,
    VK_PIPELINE_STAGE_TRANSFER_BIT = 0x00001000,
    VK_PIPELINE_STAGE_BOTTOM_OF_PIPE_BIT = 0x00002000,
    VK_PIPELINE_STAGE_HOST_BIT = 0x00004000,
    VK_PIPELINE_STAGE_ALL_GRAPHICS_BIT = 0x00008000,
    VK_PIPELINE_STAGE_ALL_COMMANDS_BIT = 0x00010000,
  // Provided by VK_VERSION_1_3
    VK_PIPELINE_STAGE_NONE = 0,
  // Provided by VK_EXT_transform_feedback
    VK_PIPELINE_STAGE_TRANSFORM_FEEDBACK_BIT_EXT = 0x01000000,
  // Provided by VK_EXT_conditional_rendering
    VK_PIPELINE_STAGE_CONDITIONAL_RENDERING_BIT_EXT = 0x00040000,
  // Provided by VK_KHR_acceleration_structure
    VK_PIPELINE_STAGE_ACCELERATION_STRUCTURE_BUILD_BIT_KHR = 0x02000000,
  // Provided by VK_KHR_ray_tracing_pipeline
    VK_PIPELINE_STAGE_RAY_TRACING_SHADER_BIT_KHR = 0x00200000,
  // Provided by VK_EXT_fragment_density_map
    VK_PIPELINE_STAGE_FRAGMENT_DENSITY_PROCESS_BIT_EXT = 0x00800000,
  // Provided by VK_KHR_fragment_shading_rate
    VK_PIPELINE_STAGE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR = 0x00400000,
  // Provided by VK_NV_device_generated_commands
    VK_PIPELINE_STAGE_COMMAND_PREPROCESS_BIT_NV = 0x00020000,
  // Provided by VK_EXT_mesh_shader
    VK_PIPELINE_STAGE_TASK_SHADER_BIT_EXT = 0x00080000,
  // Provided by VK_EXT_mesh_shader
    VK_PIPELINE_STAGE_MESH_SHADER_BIT_EXT = 0x00100000,
  // Provided by VK_NV_shading_rate_image
    VK_PIPELINE_STAGE_SHADING_RATE_IMAGE_BIT_NV = VK_PIPELINE_STAGE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR,
  // Provided by VK_NV_ray_tracing
    VK_PIPELINE_STAGE_RAY_TRACING_SHADER_BIT_NV = VK_PIPELINE_STAGE_RAY_TRACING_SHADER_BIT_KHR,
  // Provided by VK_NV_ray_tracing
    VK_PIPELINE_STAGE_ACCELERATION_STRUCTURE_BUILD_BIT_NV = VK_PIPELINE_STAGE_ACCELERATION_STRUCTURE_BUILD_BIT_KHR,
  // Provided by VK_NV_mesh_shader
    VK_PIPELINE_STAGE_TASK_SHADER_BIT_NV = VK_PIPELINE_STAGE_TASK_SHADER_BIT_EXT,
  // Provided by VK_NV_mesh_shader
    VK_PIPELINE_STAGE_MESH_SHADER_BIT_NV = VK_PIPELINE_STAGE_MESH_SHADER_BIT_EXT,
  // Provided by VK_KHR_synchronization2
    VK_PIPELINE_STAGE_NONE_KHR = VK_PIPELINE_STAGE_NONE,
} VkPipelineStageFlagBits;
```

## <a href="#_description" class="anchor"></a>Description

These values all have the same meaning as the equivalently named values
for [VkPipelineStageFlags2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlags2.html).

- `VK_PIPELINE_STAGE_NONE` specifies no stages of execution.

- `VK_PIPELINE_STAGE_DRAW_INDIRECT_BIT` specifies the stage of the
  pipeline where `VkDrawIndirect*` / `VkDispatchIndirect*` /
  `VkTraceRaysIndirect*` data structures are consumed. This stage also
  includes reading commands written by
  [vkCmdExecuteGeneratedCommandsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdExecuteGeneratedCommandsNV.html).

- `VK_PIPELINE_STAGE_TASK_SHADER_BIT_EXT` specifies the task shader
  stage.

- `VK_PIPELINE_STAGE_MESH_SHADER_BIT_EXT` specifies the mesh shader
  stage.

- `VK_PIPELINE_STAGE_VERTEX_INPUT_BIT` specifies the stage of the
  pipeline where vertex and index buffers are consumed.

- `VK_PIPELINE_STAGE_VERTEX_SHADER_BIT` specifies the vertex shader
  stage.

- `VK_PIPELINE_STAGE_TESSELLATION_CONTROL_SHADER_BIT` specifies the
  tessellation control shader stage.

- `VK_PIPELINE_STAGE_TESSELLATION_EVALUATION_SHADER_BIT` specifies the
  tessellation evaluation shader stage.

- `VK_PIPELINE_STAGE_GEOMETRY_SHADER_BIT` specifies the geometry shader
  stage.

- `VK_PIPELINE_STAGE_FRAGMENT_SHADER_BIT` specifies the fragment shader
  stage.

- `VK_PIPELINE_STAGE_EARLY_FRAGMENT_TESTS_BIT` specifies the stage of
  the pipeline where early fragment tests (depth and stencil tests
  before fragment shading) are performed. This stage also includes <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass-load-operations"
  target="_blank" rel="noopener">render pass load operations</a> for
  framebuffer attachments with a depth/stencil format.

- `VK_PIPELINE_STAGE_LATE_FRAGMENT_TESTS_BIT` specifies the stage of the
  pipeline where late fragment tests (depth and stencil tests after
  fragment shading) are performed. This stage also includes <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass-store-operations"
  target="_blank" rel="noopener">render pass store operations</a> for
  framebuffer attachments with a depth/stencil format.

- `VK_PIPELINE_STAGE_COLOR_ATTACHMENT_OUTPUT_BIT` specifies the stage of
  the pipeline after blending where the final color values are output
  from the pipeline. This stage includes <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#framebuffer-blending"
  target="_blank" rel="noopener">blending</a>, <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#framebuffer-logicop"
  target="_blank" rel="noopener">logic operations</a>, render pass <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass-load-operations"
  target="_blank" rel="noopener">load</a> and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass-store-operations"
  target="_blank" rel="noopener">store</a> operations for color
  attachments, <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass-resolve-operations"
  target="_blank" rel="noopener">render pass multisample resolve
  operations</a>, and
  [vkCmdClearAttachments](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdClearAttachments.html).

- `VK_PIPELINE_STAGE_COMPUTE_SHADER_BIT` specifies the execution of a
  compute shader.

- <span id="synchronization-pipeline-stages-transfer"></span>
  `VK_PIPELINE_STAGE_TRANSFER_BIT` specifies the following commands:

  - All <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#copies"
    target="_blank" rel="noopener">copy commands</a>, including
    [vkCmdCopyQueryPoolResults](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyQueryPoolResults.html)

  - [vkCmdBlitImage2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBlitImage2.html) and
    [vkCmdBlitImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBlitImage.html)

  - [vkCmdResolveImage2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdResolveImage2.html) and
    [vkCmdResolveImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdResolveImage.html)

  - All <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#clears"
    target="_blank" rel="noopener">clear commands</a>, with the
    exception of [vkCmdClearAttachments](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdClearAttachments.html)

- `VK_PIPELINE_STAGE_HOST_BIT` specifies a pseudo-stage indicating
  execution on the host of reads/writes of device memory. This stage is
  not invoked by any commands recorded in a command buffer.

- `VK_PIPELINE_STAGE_ACCELERATION_STRUCTURE_BUILD_BIT_KHR` specifies the
  execution of
  [vkCmdBuildAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBuildAccelerationStructureNV.html),
  [vkCmdCopyAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyAccelerationStructureNV.html),
  [vkCmdWriteAccelerationStructuresPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdWriteAccelerationStructuresPropertiesNV.html)
  ,
  [vkCmdBuildAccelerationStructuresKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBuildAccelerationStructuresKHR.html),
  [vkCmdBuildAccelerationStructuresIndirectKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBuildAccelerationStructuresIndirectKHR.html),
  [vkCmdCopyAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyAccelerationStructureKHR.html),
  [vkCmdCopyAccelerationStructureToMemoryKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyAccelerationStructureToMemoryKHR.html),
  [vkCmdCopyMemoryToAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyMemoryToAccelerationStructureKHR.html),
  and
  [vkCmdWriteAccelerationStructuresPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdWriteAccelerationStructuresPropertiesKHR.html).

- `VK_PIPELINE_STAGE_RAY_TRACING_SHADER_BIT_KHR` specifies the execution
  of the ray tracing shader stages, via
  [vkCmdTraceRaysNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdTraceRaysNV.html) ,
  [vkCmdTraceRaysKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdTraceRaysKHR.html), or
  [vkCmdTraceRaysIndirectKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdTraceRaysIndirectKHR.html)

- `VK_PIPELINE_STAGE_ALL_GRAPHICS_BIT` specifies the execution of all
  graphics pipeline stages, and is equivalent to the logical OR of:

  - `VK_PIPELINE_STAGE_DRAW_INDIRECT_BIT`

  - `VK_PIPELINE_STAGE_TASK_SHADER_BIT_EXT`

  - `VK_PIPELINE_STAGE_MESH_SHADER_BIT_EXT`

  - `VK_PIPELINE_STAGE_VERTEX_INPUT_BIT`

  - `VK_PIPELINE_STAGE_VERTEX_SHADER_BIT`

  - `VK_PIPELINE_STAGE_TESSELLATION_CONTROL_SHADER_BIT`

  - `VK_PIPELINE_STAGE_TESSELLATION_EVALUATION_SHADER_BIT`

  - `VK_PIPELINE_STAGE_GEOMETRY_SHADER_BIT`

  - `VK_PIPELINE_STAGE_FRAGMENT_SHADER_BIT`

  - `VK_PIPELINE_STAGE_EARLY_FRAGMENT_TESTS_BIT`

  - `VK_PIPELINE_STAGE_LATE_FRAGMENT_TESTS_BIT`

  - `VK_PIPELINE_STAGE_COLOR_ATTACHMENT_OUTPUT_BIT`

  - `VK_PIPELINE_STAGE_CONDITIONAL_RENDERING_BIT_EXT`

  - `VK_PIPELINE_STAGE_TRANSFORM_FEEDBACK_BIT_EXT`

  - `VK_PIPELINE_STAGE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR`

  - `VK_PIPELINE_STAGE_FRAGMENT_DENSITY_PROCESS_BIT_EXT`

- `VK_PIPELINE_STAGE_ALL_COMMANDS_BIT` specifies all operations
  performed by all commands supported on the queue it is used with.

- `VK_PIPELINE_STAGE_CONDITIONAL_RENDERING_BIT_EXT` specifies the stage
  of the pipeline where the predicate of conditional rendering is
  consumed.

- `VK_PIPELINE_STAGE_TRANSFORM_FEEDBACK_BIT_EXT` specifies the stage of
  the pipeline where vertex attribute output values are written to the
  transform feedback buffers.

- `VK_PIPELINE_STAGE_COMMAND_PREPROCESS_BIT_NV` specifies the stage of
  the pipeline where device-side preprocessing for generated commands
  via
  [vkCmdPreprocessGeneratedCommandsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdPreprocessGeneratedCommandsNV.html)
  is handled.

- `VK_PIPELINE_STAGE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR` specifies
  the stage of the pipeline where the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-fragment-shading-rate-attachment"
  target="_blank" rel="noopener">fragment shading rate attachment</a> or
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-shading-rate-image"
  target="_blank" rel="noopener">shading rate image</a> is read to
  determine the fragment shading rate for portions of a rasterized
  primitive.

- `VK_PIPELINE_STAGE_FRAGMENT_DENSITY_PROCESS_BIT_EXT` specifies the
  stage of the pipeline where the fragment density map is read to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fragmentdensitymapops"
  target="_blank" rel="noopener">generate the fragment areas</a>.

- `VK_PIPELINE_STAGE_TOP_OF_PIPE_BIT` is equivalent to
  `VK_PIPELINE_STAGE_ALL_COMMANDS_BIT` with
  [VkAccessFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccessFlags.html) set to `0` when specified in the
  second synchronization scope, but specifies no stage of execution when
  specified in the first scope.

- `VK_PIPELINE_STAGE_BOTTOM_OF_PIPE_BIT` is equivalent to
  `VK_PIPELINE_STAGE_ALL_COMMANDS_BIT` with
  [VkAccessFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccessFlags.html) set to `0` when specified in the
  first synchronization scope, but specifies no stage of execution when
  specified in the second scope.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkCheckpointDataNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCheckpointDataNV.html),
[VkPipelineStageFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlags.html),
[vkCmdWriteBufferMarkerAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdWriteBufferMarkerAMD.html),
[vkCmdWriteTimestamp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdWriteTimestamp.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPipelineStageFlagBits"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
