# VkAccessFlagBits(3) Manual Page

## Name

VkAccessFlagBits - Bitmask specifying memory access types that will participate in a memory dependency



## [](#_c_specification)C Specification

Bits which **can** be set in the `srcAccessMask` and `dstAccessMask` members of [VkSubpassDependency](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassDependency.html), [VkSubpassDependency2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassDependency2.html), [VkMemoryBarrier](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryBarrier.html), [VkBufferMemoryBarrier](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferMemoryBarrier.html), and [VkImageMemoryBarrier](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageMemoryBarrier.html), specifying access behavior, are:

```c++
// Provided by VK_VERSION_1_0
typedef enum VkAccessFlagBits {
    VK_ACCESS_INDIRECT_COMMAND_READ_BIT = 0x00000001,
    VK_ACCESS_INDEX_READ_BIT = 0x00000002,
    VK_ACCESS_VERTEX_ATTRIBUTE_READ_BIT = 0x00000004,
    VK_ACCESS_UNIFORM_READ_BIT = 0x00000008,
    VK_ACCESS_INPUT_ATTACHMENT_READ_BIT = 0x00000010,
    VK_ACCESS_SHADER_READ_BIT = 0x00000020,
    VK_ACCESS_SHADER_WRITE_BIT = 0x00000040,
    VK_ACCESS_COLOR_ATTACHMENT_READ_BIT = 0x00000080,
    VK_ACCESS_COLOR_ATTACHMENT_WRITE_BIT = 0x00000100,
    VK_ACCESS_DEPTH_STENCIL_ATTACHMENT_READ_BIT = 0x00000200,
    VK_ACCESS_DEPTH_STENCIL_ATTACHMENT_WRITE_BIT = 0x00000400,
    VK_ACCESS_TRANSFER_READ_BIT = 0x00000800,
    VK_ACCESS_TRANSFER_WRITE_BIT = 0x00001000,
    VK_ACCESS_HOST_READ_BIT = 0x00002000,
    VK_ACCESS_HOST_WRITE_BIT = 0x00004000,
    VK_ACCESS_MEMORY_READ_BIT = 0x00008000,
    VK_ACCESS_MEMORY_WRITE_BIT = 0x00010000,
  // Provided by VK_VERSION_1_3
    VK_ACCESS_NONE = 0,
  // Provided by VK_EXT_transform_feedback
    VK_ACCESS_TRANSFORM_FEEDBACK_WRITE_BIT_EXT = 0x02000000,
  // Provided by VK_EXT_transform_feedback
    VK_ACCESS_TRANSFORM_FEEDBACK_COUNTER_READ_BIT_EXT = 0x04000000,
  // Provided by VK_EXT_transform_feedback
    VK_ACCESS_TRANSFORM_FEEDBACK_COUNTER_WRITE_BIT_EXT = 0x08000000,
  // Provided by VK_EXT_conditional_rendering
    VK_ACCESS_CONDITIONAL_RENDERING_READ_BIT_EXT = 0x00100000,
  // Provided by VK_EXT_blend_operation_advanced
    VK_ACCESS_COLOR_ATTACHMENT_READ_NONCOHERENT_BIT_EXT = 0x00080000,
  // Provided by VK_KHR_acceleration_structure
    VK_ACCESS_ACCELERATION_STRUCTURE_READ_BIT_KHR = 0x00200000,
  // Provided by VK_KHR_acceleration_structure
    VK_ACCESS_ACCELERATION_STRUCTURE_WRITE_BIT_KHR = 0x00400000,
  // Provided by VK_EXT_fragment_density_map
    VK_ACCESS_FRAGMENT_DENSITY_MAP_READ_BIT_EXT = 0x01000000,
  // Provided by VK_KHR_fragment_shading_rate
    VK_ACCESS_FRAGMENT_SHADING_RATE_ATTACHMENT_READ_BIT_KHR = 0x00800000,
  // Provided by VK_EXT_device_generated_commands
    VK_ACCESS_COMMAND_PREPROCESS_READ_BIT_EXT = 0x00020000,
  // Provided by VK_EXT_device_generated_commands
    VK_ACCESS_COMMAND_PREPROCESS_WRITE_BIT_EXT = 0x00040000,
  // Provided by VK_NV_shading_rate_image
    VK_ACCESS_SHADING_RATE_IMAGE_READ_BIT_NV = VK_ACCESS_FRAGMENT_SHADING_RATE_ATTACHMENT_READ_BIT_KHR,
  // Provided by VK_NV_ray_tracing
    VK_ACCESS_ACCELERATION_STRUCTURE_READ_BIT_NV = VK_ACCESS_ACCELERATION_STRUCTURE_READ_BIT_KHR,
  // Provided by VK_NV_ray_tracing
    VK_ACCESS_ACCELERATION_STRUCTURE_WRITE_BIT_NV = VK_ACCESS_ACCELERATION_STRUCTURE_WRITE_BIT_KHR,
  // Provided by VK_NV_device_generated_commands
    VK_ACCESS_COMMAND_PREPROCESS_READ_BIT_NV = VK_ACCESS_COMMAND_PREPROCESS_READ_BIT_EXT,
  // Provided by VK_NV_device_generated_commands
    VK_ACCESS_COMMAND_PREPROCESS_WRITE_BIT_NV = VK_ACCESS_COMMAND_PREPROCESS_WRITE_BIT_EXT,
  // Provided by VK_KHR_synchronization2
    VK_ACCESS_NONE_KHR = VK_ACCESS_NONE,
} VkAccessFlagBits;
```

## [](#_description)Description

These values all have the same meaning as the equivalently named values for [VkAccessFlags2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlags2.html).

- `VK_ACCESS_NONE` specifies no accesses.
- `VK_ACCESS_MEMORY_READ_BIT` specifies all read accesses. It is always valid in any access mask, and is treated as equivalent to setting all `READ` access flags that are valid where it is used.
- `VK_ACCESS_MEMORY_WRITE_BIT` specifies all write accesses. It is always valid in any access mask, and is treated as equivalent to setting all `WRITE` access flags that are valid where it is used.
- `VK_ACCESS_INDIRECT_COMMAND_READ_BIT` specifies read access to indirect command data read as part of an indirect build, trace, drawing or dispatching command. Such access occurs in the `VK_PIPELINE_STAGE_DRAW_INDIRECT_BIT` pipeline stage.
- `VK_ACCESS_INDEX_READ_BIT` specifies read access to an index buffer as part of an indexed drawing command, bound by [vkCmdBindIndexBuffer2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBindIndexBuffer2.html) and [vkCmdBindIndexBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBindIndexBuffer.html). Such access occurs in the `VK_PIPELINE_STAGE_VERTEX_INPUT_BIT` pipeline stage.
- `VK_ACCESS_VERTEX_ATTRIBUTE_READ_BIT` specifies read access to a vertex buffer as part of a drawing command, bound by [vkCmdBindVertexBuffers](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBindVertexBuffers.html). Such access occurs in the `VK_PIPELINE_STAGE_VERTEX_INPUT_BIT` pipeline stage.
- `VK_ACCESS_UNIFORM_READ_BIT` specifies read access to a [uniform buffer](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-uniformbuffer) in any shader pipeline stage.
- `VK_ACCESS_INPUT_ATTACHMENT_READ_BIT` specifies read access to an [input attachment](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass) within a render pass during subpass shading or fragment shading. Such access occurs in the `VK_PIPELINE_STAGE_2_SUBPASS_SHADER_BIT_HUAWEI` or `VK_PIPELINE_STAGE_FRAGMENT_SHADER_BIT` pipeline stage.
- `VK_ACCESS_SHADER_READ_BIT` specifies read access to a [uniform texel buffer](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-uniformtexelbuffer), [sampled image](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-sampledimage), [storage buffer](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-storagebuffer), [physical storage buffer](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-physical-storage-buffer), [shader binding table](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shader-binding-table), [storage tensor](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-storagetensor), [storage texel buffer](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-storagetexelbuffer), or [storage image](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-storageimage) in any shader pipeline stage.
- `VK_ACCESS_SHADER_WRITE_BIT` specifies write access to a [storage buffer](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-storagebuffer), [physical storage buffer](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-physical-storage-buffer), [storage tensor](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-storagetensor), [storage texel buffer](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-storagetexelbuffer), or [storage image](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-storageimage) in any shader pipeline stage.
- `VK_ACCESS_COLOR_ATTACHMENT_READ_BIT` specifies read access to a [color attachment](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass), such as via [blending](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#framebuffer-blending) (other than [advanced blend operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#framebuffer-blend-advanced)), [logic operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#framebuffer-logicop) or certain [render pass load operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-load-operations) in the `VK_PIPELINE_STAGE_COLOR_ATTACHMENT_OUTPUT_BIT` pipeline stage or via [fragment shader tile image reads](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#fragops-shader-tileimage-reads) in the `VK_PIPELINE_STAGE_FRAGMENT_SHADER_BIT` pipeline stage.
- `VK_ACCESS_COLOR_ATTACHMENT_WRITE_BIT` specifies write access to a [color, resolve, or depth/stencil resolve attachment](#renderpass) during a [render pass](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass) or via certain render pass [load](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-load-operations) and [store](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-store-operations) operations. Such access occurs in the `VK_PIPELINE_STAGE_COLOR_ATTACHMENT_OUTPUT_BIT` pipeline stage.
- `VK_ACCESS_DEPTH_STENCIL_ATTACHMENT_READ_BIT` specifies read access to a [depth/stencil attachment](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass), via [depth or stencil operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#fragops-ds-state) or certain [render pass load operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-load-operations) in the `VK_PIPELINE_STAGE_EARLY_FRAGMENT_TESTS_BIT` or `VK_PIPELINE_STAGE_LATE_FRAGMENT_TESTS_BIT` pipeline stages or via [fragment shader tile image reads](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#fragops-shader-tileimage-reads) in the `VK_PIPELINE_STAGE_FRAGMENT_SHADER_BIT` pipeline stage.
- `VK_ACCESS_DEPTH_STENCIL_ATTACHMENT_WRITE_BIT` specifies write access to a [depth/stencil attachment](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass), via [depth or stencil operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#fragops-ds-state) or certain render pass [load](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-load-operations) and [store](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-store-operations) operations. Such access occurs in the `VK_PIPELINE_STAGE_EARLY_FRAGMENT_TESTS_BIT` or `VK_PIPELINE_STAGE_LATE_FRAGMENT_TESTS_BIT` pipeline stages.
- `VK_ACCESS_TRANSFER_READ_BIT` specifies read access to an image, tensor, or buffer in a [copy](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#copies) operation. Such access occurs in the `VK_PIPELINE_STAGE_2_ALL_TRANSFER_BIT` pipeline stage.
- `VK_ACCESS_TRANSFER_WRITE_BIT` specifies write access to an image, tensor, or buffer in a [clear](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#clears) or [copy](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#copies) operation. Such access occurs in the `VK_PIPELINE_STAGE_2_ALL_TRANSFER_BIT` pipeline stage.
- `VK_ACCESS_HOST_READ_BIT` specifies read access by a host operation. Accesses of this type are not performed through a resource, but directly on memory. Such access occurs in the `VK_PIPELINE_STAGE_HOST_BIT` pipeline stage.
- `VK_ACCESS_HOST_WRITE_BIT` specifies write access by a host operation. Accesses of this type are not performed through a resource, but directly on memory. Such access occurs in the `VK_PIPELINE_STAGE_HOST_BIT` pipeline stage.
- `VK_ACCESS_CONDITIONAL_RENDERING_READ_BIT_EXT` specifies read access to a predicate as part of conditional rendering. Such access occurs in the `VK_PIPELINE_STAGE_CONDITIONAL_RENDERING_BIT_EXT` pipeline stage.
- `VK_ACCESS_TRANSFORM_FEEDBACK_WRITE_BIT_EXT` specifies write access to a transform feedback buffer made when transform feedback is active. Such access occurs in the `VK_PIPELINE_STAGE_TRANSFORM_FEEDBACK_BIT_EXT` pipeline stage.
- `VK_ACCESS_TRANSFORM_FEEDBACK_COUNTER_READ_BIT_EXT` specifies read access to a transform feedback counter buffer which is read when `vkCmdBeginTransformFeedbackEXT` executes. Such access occurs in the `VK_PIPELINE_STAGE_TRANSFORM_FEEDBACK_BIT_EXT` pipeline stage.
- `VK_ACCESS_TRANSFORM_FEEDBACK_COUNTER_WRITE_BIT_EXT` specifies write access to a transform feedback counter buffer which is written when `vkCmdEndTransformFeedbackEXT` executes. Such access occurs in the `VK_PIPELINE_STAGE_TRANSFORM_FEEDBACK_BIT_EXT` pipeline stage.
- `VK_ACCESS_COMMAND_PREPROCESS_READ_BIT_NV` specifies reads from buffer inputs to [vkCmdPreprocessGeneratedCommandsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdPreprocessGeneratedCommandsNV.html). Such access occurs in the `VK_PIPELINE_STAGE_COMMAND_PREPROCESS_BIT_NV` pipeline stage.
- `VK_ACCESS_COMMAND_PREPROCESS_WRITE_BIT_NV` specifies writes to the target command buffer preprocess outputs in [vkCmdPreprocessGeneratedCommandsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdPreprocessGeneratedCommandsNV.html). Such access occurs in the `VK_PIPELINE_STAGE_COMMAND_PREPROCESS_BIT_NV` pipeline stage.
- `VK_ACCESS_COMMAND_PREPROCESS_READ_BIT_EXT` specifies reads from buffer inputs to [vkCmdPreprocessGeneratedCommandsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdPreprocessGeneratedCommandsEXT.html). Such access occurs in the `VK_PIPELINE_STAGE_COMMAND_PREPROCESS_BIT_EXT` pipeline stage.
- `VK_ACCESS_COMMAND_PREPROCESS_WRITE_BIT_EXT` specifies writes to the target command buffer preprocess outputs in [vkCmdPreprocessGeneratedCommandsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdPreprocessGeneratedCommandsEXT.html). Such access occurs in the `VK_PIPELINE_STAGE_COMMAND_PREPROCESS_BIT_EXT` pipeline stage.
- `VK_ACCESS_COLOR_ATTACHMENT_READ_NONCOHERENT_BIT_EXT` specifies read access to [color attachments](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass), including [advanced blend operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#framebuffer-blend-advanced). Such access occurs in the `VK_PIPELINE_STAGE_COLOR_ATTACHMENT_OUTPUT_BIT` pipeline stage.
- `VK_ACCESS_2_INVOCATION_MASK_READ_BIT_HUAWEI` specifies read access to an invocation mask image in the `VK_PIPELINE_STAGE_2_INVOCATION_MASK_BIT_HUAWEI` pipeline stage.
- `VK_ACCESS_ACCELERATION_STRUCTURE_READ_BIT_KHR` specifies read access to an acceleration structure as part of a trace, build, or copy command, or to an [acceleration structure scratch buffer](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#acceleration-structure-scratch) as part of a build command. Such access occurs in the `VK_PIPELINE_STAGE_RAY_TRACING_SHADER_BIT_KHR` pipeline stage or `VK_PIPELINE_STAGE_ACCELERATION_STRUCTURE_BUILD_BIT_KHR` pipeline stage.
- `VK_ACCESS_ACCELERATION_STRUCTURE_WRITE_BIT_KHR` specifies write access to an acceleration structure or [acceleration structure scratch buffer](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#acceleration-structure-scratch) as part of a build or copy command. Such access occurs in the `VK_PIPELINE_STAGE_ACCELERATION_STRUCTURE_BUILD_BIT_KHR` pipeline stage.
- `VK_ACCESS_FRAGMENT_DENSITY_MAP_READ_BIT_EXT` specifies read access to a [fragment density map attachment](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-fragmentdensitymapattachment) during dynamic [fragment density map operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#fragmentdensitymapops) Such access occurs in the `VK_PIPELINE_STAGE_FRAGMENT_DENSITY_PROCESS_BIT_EXT` pipeline stage.
- `VK_ACCESS_FRAGMENT_SHADING_RATE_ATTACHMENT_READ_BIT_KHR` specifies read access to a fragment shading rate attachment during rasterization. Such access occurs in the `VK_PIPELINE_STAGE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR` pipeline stage.
- `VK_ACCESS_SHADING_RATE_IMAGE_READ_BIT_NV` specifies read access to a shading rate image during rasterization. Such access occurs in the `VK_PIPELINE_STAGE_SHADING_RATE_IMAGE_BIT_NV` pipeline stage. It is equivalent to `VK_PIPELINE_STAGE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR`.

Certain access types are only performed by a subset of pipeline stages. Any synchronization command that takes both stage masks and access masks uses both to define the [access scopes](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-access-scopes) - only the specified access types performed by the specified stages are included in the access scope. An application **must** not specify an access flag in a synchronization command if it does not include a pipeline stage in the corresponding stage mask that is able to perform accesses of that type. The following table lists, for each access flag, which pipeline stages **can** perform that type of access.

Table 1. Supported Access Types   Access flag Supported pipeline stages

`VK_ACCESS_2_NONE`

Any

`VK_ACCESS_2_INDIRECT_COMMAND_READ_BIT`

`VK_PIPELINE_STAGE_2_DRAW_INDIRECT_BIT`, `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR`

`VK_ACCESS_2_INDEX_READ_BIT`

`VK_PIPELINE_STAGE_2_VERTEX_INPUT_BIT`, `VK_PIPELINE_STAGE_2_INDEX_INPUT_BIT`

`VK_ACCESS_2_VERTEX_ATTRIBUTE_READ_BIT`

`VK_PIPELINE_STAGE_2_VERTEX_INPUT_BIT`, `VK_PIPELINE_STAGE_2_VERTEX_ATTRIBUTE_INPUT_BIT`

`VK_ACCESS_2_UNIFORM_READ_BIT`

`VK_PIPELINE_STAGE_2_VERTEX_SHADER_BIT`, `VK_PIPELINE_STAGE_2_TESSELLATION_CONTROL_SHADER_BIT`, `VK_PIPELINE_STAGE_2_TESSELLATION_EVALUATION_SHADER_BIT`, `VK_PIPELINE_STAGE_2_GEOMETRY_SHADER_BIT`, `VK_PIPELINE_STAGE_2_FRAGMENT_SHADER_BIT`, `VK_PIPELINE_STAGE_2_COMPUTE_SHADER_BIT`, `VK_PIPELINE_STAGE_2_RAY_TRACING_SHADER_BIT_KHR`, `VK_PIPELINE_STAGE_2_TASK_SHADER_BIT_EXT`, `VK_PIPELINE_STAGE_2_MESH_SHADER_BIT_EXT`, `VK_PIPELINE_STAGE_2_SUBPASS_SHADER_BIT_HUAWEI`, `VK_PIPELINE_STAGE_2_CLUSTER_CULLING_SHADER_BIT_HUAWEI`

`VK_ACCESS_2_INPUT_ATTACHMENT_READ_BIT`

`VK_PIPELINE_STAGE_2_FRAGMENT_SHADER_BIT`, `VK_PIPELINE_STAGE_2_SUBPASS_SHADER_BIT_HUAWEI`

`VK_ACCESS_2_SHADER_READ_BIT`

`VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR`, `VK_PIPELINE_STAGE_2_MICROMAP_BUILD_BIT_EXT`, `VK_PIPELINE_STAGE_2_VERTEX_SHADER_BIT`, `VK_PIPELINE_STAGE_2_TESSELLATION_CONTROL_SHADER_BIT`, `VK_PIPELINE_STAGE_2_TESSELLATION_EVALUATION_SHADER_BIT`, `VK_PIPELINE_STAGE_2_GEOMETRY_SHADER_BIT`, `VK_PIPELINE_STAGE_2_FRAGMENT_SHADER_BIT`, `VK_PIPELINE_STAGE_2_COMPUTE_SHADER_BIT`, `VK_PIPELINE_STAGE_2_RAY_TRACING_SHADER_BIT_KHR`, `VK_PIPELINE_STAGE_2_TASK_SHADER_BIT_EXT`, `VK_PIPELINE_STAGE_2_MESH_SHADER_BIT_EXT`, `VK_PIPELINE_STAGE_2_SUBPASS_SHADER_BIT_HUAWEI`, `VK_PIPELINE_STAGE_2_CLUSTER_CULLING_SHADER_BIT_HUAWEI`

`VK_ACCESS_2_SHADER_WRITE_BIT`

`VK_PIPELINE_STAGE_2_VERTEX_SHADER_BIT`, `VK_PIPELINE_STAGE_2_TESSELLATION_CONTROL_SHADER_BIT`, `VK_PIPELINE_STAGE_2_TESSELLATION_EVALUATION_SHADER_BIT`, `VK_PIPELINE_STAGE_2_GEOMETRY_SHADER_BIT`, `VK_PIPELINE_STAGE_2_FRAGMENT_SHADER_BIT`, `VK_PIPELINE_STAGE_2_COMPUTE_SHADER_BIT`, `VK_PIPELINE_STAGE_2_RAY_TRACING_SHADER_BIT_KHR`, `VK_PIPELINE_STAGE_2_TASK_SHADER_BIT_EXT`, `VK_PIPELINE_STAGE_2_MESH_SHADER_BIT_EXT`, `VK_PIPELINE_STAGE_2_SUBPASS_SHADER_BIT_HUAWEI`, `VK_PIPELINE_STAGE_2_CLUSTER_CULLING_SHADER_BIT_HUAWEI`

`VK_ACCESS_2_COLOR_ATTACHMENT_READ_BIT`

`VK_PIPELINE_STAGE_2_FRAGMENT_SHADER_BIT`, `VK_PIPELINE_STAGE_2_COLOR_ATTACHMENT_OUTPUT_BIT`

`VK_ACCESS_2_COLOR_ATTACHMENT_WRITE_BIT`

`VK_PIPELINE_STAGE_2_COLOR_ATTACHMENT_OUTPUT_BIT`

`VK_ACCESS_2_DEPTH_STENCIL_ATTACHMENT_READ_BIT`

`VK_PIPELINE_STAGE_2_FRAGMENT_SHADER_BIT`, `VK_PIPELINE_STAGE_2_EARLY_FRAGMENT_TESTS_BIT`, `VK_PIPELINE_STAGE_2_LATE_FRAGMENT_TESTS_BIT`

`VK_ACCESS_2_DEPTH_STENCIL_ATTACHMENT_WRITE_BIT`

`VK_PIPELINE_STAGE_2_EARLY_FRAGMENT_TESTS_BIT`, `VK_PIPELINE_STAGE_2_LATE_FRAGMENT_TESTS_BIT`

`VK_ACCESS_2_TRANSFER_READ_BIT`

`VK_PIPELINE_STAGE_2_ALL_TRANSFER_BIT`, `VK_PIPELINE_STAGE_2_COPY_BIT`, `VK_PIPELINE_STAGE_2_RESOLVE_BIT`, `VK_PIPELINE_STAGE_2_BLIT_BIT`, `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR`, `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_COPY_BIT_KHR`, `VK_PIPELINE_STAGE_2_MICROMAP_BUILD_BIT_EXT`, `VK_PIPELINE_STAGE_2_CONVERT_COOPERATIVE_VECTOR_MATRIX_BIT_NV`

`VK_ACCESS_2_TRANSFER_WRITE_BIT`

`VK_PIPELINE_STAGE_2_ALL_TRANSFER_BIT`, `VK_PIPELINE_STAGE_2_COPY_BIT`, `VK_PIPELINE_STAGE_2_RESOLVE_BIT`, `VK_PIPELINE_STAGE_2_BLIT_BIT`, `VK_PIPELINE_STAGE_2_CLEAR_BIT`, `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR`, `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_COPY_BIT_KHR`, `VK_PIPELINE_STAGE_2_MICROMAP_BUILD_BIT_EXT`, `VK_PIPELINE_STAGE_2_CONVERT_COOPERATIVE_VECTOR_MATRIX_BIT_NV`

`VK_ACCESS_2_HOST_READ_BIT`

`VK_PIPELINE_STAGE_2_HOST_BIT`

`VK_ACCESS_2_HOST_WRITE_BIT`

`VK_PIPELINE_STAGE_2_HOST_BIT`

`VK_ACCESS_2_MEMORY_READ_BIT`

Any

`VK_ACCESS_2_MEMORY_WRITE_BIT`

Any

`VK_ACCESS_2_SHADER_SAMPLED_READ_BIT`

`VK_PIPELINE_STAGE_2_VERTEX_SHADER_BIT`, `VK_PIPELINE_STAGE_2_TESSELLATION_CONTROL_SHADER_BIT`, `VK_PIPELINE_STAGE_2_TESSELLATION_EVALUATION_SHADER_BIT`, `VK_PIPELINE_STAGE_2_GEOMETRY_SHADER_BIT`, `VK_PIPELINE_STAGE_2_FRAGMENT_SHADER_BIT`, `VK_PIPELINE_STAGE_2_COMPUTE_SHADER_BIT`, `VK_PIPELINE_STAGE_2_RAY_TRACING_SHADER_BIT_KHR`, `VK_PIPELINE_STAGE_2_TASK_SHADER_BIT_EXT`, `VK_PIPELINE_STAGE_2_MESH_SHADER_BIT_EXT`, `VK_PIPELINE_STAGE_2_SUBPASS_SHADER_BIT_HUAWEI`, `VK_PIPELINE_STAGE_2_CLUSTER_CULLING_SHADER_BIT_HUAWEI`

`VK_ACCESS_2_SHADER_STORAGE_READ_BIT`

`VK_PIPELINE_STAGE_2_VERTEX_SHADER_BIT`, `VK_PIPELINE_STAGE_2_TESSELLATION_CONTROL_SHADER_BIT`, `VK_PIPELINE_STAGE_2_TESSELLATION_EVALUATION_SHADER_BIT`, `VK_PIPELINE_STAGE_2_GEOMETRY_SHADER_BIT`, `VK_PIPELINE_STAGE_2_FRAGMENT_SHADER_BIT`, `VK_PIPELINE_STAGE_2_COMPUTE_SHADER_BIT`, `VK_PIPELINE_STAGE_2_RAY_TRACING_SHADER_BIT_KHR`, `VK_PIPELINE_STAGE_2_TASK_SHADER_BIT_EXT`, `VK_PIPELINE_STAGE_2_MESH_SHADER_BIT_EXT`, `VK_PIPELINE_STAGE_2_SUBPASS_SHADER_BIT_HUAWEI`, `VK_PIPELINE_STAGE_2_CLUSTER_CULLING_SHADER_BIT_HUAWEI`

`VK_ACCESS_2_SHADER_STORAGE_WRITE_BIT`

`VK_PIPELINE_STAGE_2_VERTEX_SHADER_BIT`, `VK_PIPELINE_STAGE_2_TESSELLATION_CONTROL_SHADER_BIT`, `VK_PIPELINE_STAGE_2_TESSELLATION_EVALUATION_SHADER_BIT`, `VK_PIPELINE_STAGE_2_GEOMETRY_SHADER_BIT`, `VK_PIPELINE_STAGE_2_FRAGMENT_SHADER_BIT`, `VK_PIPELINE_STAGE_2_COMPUTE_SHADER_BIT`, `VK_PIPELINE_STAGE_2_RAY_TRACING_SHADER_BIT_KHR`, `VK_PIPELINE_STAGE_2_TASK_SHADER_BIT_EXT`, `VK_PIPELINE_STAGE_2_MESH_SHADER_BIT_EXT`, `VK_PIPELINE_STAGE_2_SUBPASS_SHADER_BIT_HUAWEI`, `VK_PIPELINE_STAGE_2_CLUSTER_CULLING_SHADER_BIT_HUAWEI`

`VK_ACCESS_2_VIDEO_DECODE_READ_BIT_KHR`

`VK_PIPELINE_STAGE_2_VIDEO_DECODE_BIT_KHR`

`VK_ACCESS_2_VIDEO_DECODE_WRITE_BIT_KHR`

`VK_PIPELINE_STAGE_2_VIDEO_DECODE_BIT_KHR`

`VK_ACCESS_2_VIDEO_ENCODE_READ_BIT_KHR`

`VK_PIPELINE_STAGE_2_VIDEO_ENCODE_BIT_KHR`

`VK_ACCESS_2_VIDEO_ENCODE_WRITE_BIT_KHR`

`VK_PIPELINE_STAGE_2_VIDEO_ENCODE_BIT_KHR`

`VK_ACCESS_2_TRANSFORM_FEEDBACK_WRITE_BIT_EXT`

`VK_PIPELINE_STAGE_2_TRANSFORM_FEEDBACK_BIT_EXT`

`VK_ACCESS_2_TRANSFORM_FEEDBACK_COUNTER_READ_BIT_EXT`

`VK_PIPELINE_STAGE_2_DRAW_INDIRECT_BIT`, `VK_PIPELINE_STAGE_2_TRANSFORM_FEEDBACK_BIT_EXT`

`VK_ACCESS_2_TRANSFORM_FEEDBACK_COUNTER_WRITE_BIT_EXT`

`VK_PIPELINE_STAGE_2_TRANSFORM_FEEDBACK_BIT_EXT`

`VK_ACCESS_2_CONDITIONAL_RENDERING_READ_BIT_EXT`

`VK_PIPELINE_STAGE_2_CONDITIONAL_RENDERING_BIT_EXT`

`VK_ACCESS_2_COMMAND_PREPROCESS_READ_BIT_EXT`

`VK_PIPELINE_STAGE_2_COMMAND_PREPROCESS_BIT_EXT`

`VK_ACCESS_2_COMMAND_PREPROCESS_WRITE_BIT_EXT`

`VK_PIPELINE_STAGE_2_COMMAND_PREPROCESS_BIT_EXT`

`VK_ACCESS_2_FRAGMENT_SHADING_RATE_ATTACHMENT_READ_BIT_KHR`

`VK_PIPELINE_STAGE_2_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR`

`VK_ACCESS_2_ACCELERATION_STRUCTURE_READ_BIT_KHR`

`VK_PIPELINE_STAGE_2_VERTEX_SHADER_BIT`, `VK_PIPELINE_STAGE_2_TESSELLATION_CONTROL_SHADER_BIT`, `VK_PIPELINE_STAGE_2_TESSELLATION_EVALUATION_SHADER_BIT`, `VK_PIPELINE_STAGE_2_GEOMETRY_SHADER_BIT`, `VK_PIPELINE_STAGE_2_FRAGMENT_SHADER_BIT`, `VK_PIPELINE_STAGE_2_COMPUTE_SHADER_BIT`, `VK_PIPELINE_STAGE_2_RAY_TRACING_SHADER_BIT_KHR`, `VK_PIPELINE_STAGE_2_TASK_SHADER_BIT_EXT`, `VK_PIPELINE_STAGE_2_MESH_SHADER_BIT_EXT`, `VK_PIPELINE_STAGE_2_CLUSTER_CULLING_SHADER_BIT_HUAWEI`, `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR`, `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_COPY_BIT_KHR`, `VK_PIPELINE_STAGE_2_SUBPASS_SHADER_BIT_HUAWEI`

`VK_ACCESS_2_ACCELERATION_STRUCTURE_WRITE_BIT_KHR`

`VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR`, `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_COPY_BIT_KHR`

`VK_ACCESS_2_FRAGMENT_DENSITY_MAP_READ_BIT_EXT`

`VK_PIPELINE_STAGE_2_FRAGMENT_DENSITY_PROCESS_BIT_EXT`

`VK_ACCESS_2_COLOR_ATTACHMENT_READ_NONCOHERENT_BIT_EXT`

`VK_PIPELINE_STAGE_2_COLOR_ATTACHMENT_OUTPUT_BIT`

`VK_ACCESS_2_DESCRIPTOR_BUFFER_READ_BIT_EXT`

`VK_PIPELINE_STAGE_2_VERTEX_SHADER_BIT`, `VK_PIPELINE_STAGE_2_TESSELLATION_CONTROL_SHADER_BIT`, `VK_PIPELINE_STAGE_2_TESSELLATION_EVALUATION_SHADER_BIT`, `VK_PIPELINE_STAGE_2_GEOMETRY_SHADER_BIT`, `VK_PIPELINE_STAGE_2_FRAGMENT_SHADER_BIT`, `VK_PIPELINE_STAGE_2_COMPUTE_SHADER_BIT`, `VK_PIPELINE_STAGE_2_RAY_TRACING_SHADER_BIT_KHR`, `VK_PIPELINE_STAGE_2_TASK_SHADER_BIT_EXT`, `VK_PIPELINE_STAGE_2_MESH_SHADER_BIT_EXT`, `VK_PIPELINE_STAGE_2_SUBPASS_SHADER_BIT_HUAWEI`, `VK_PIPELINE_STAGE_2_CLUSTER_CULLING_SHADER_BIT_HUAWEI`

`VK_ACCESS_2_INVOCATION_MASK_READ_BIT_HUAWEI`

`VK_PIPELINE_STAGE_2_INVOCATION_MASK_BIT_HUAWEI`

`VK_ACCESS_2_MICROMAP_READ_BIT_EXT`

`VK_PIPELINE_STAGE_2_MICROMAP_BUILD_BIT_EXT`, `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR`

`VK_ACCESS_2_MICROMAP_WRITE_BIT_EXT`

`VK_PIPELINE_STAGE_2_MICROMAP_BUILD_BIT_EXT`

`VK_ACCESS_2_OPTICAL_FLOW_READ_BIT_NV`

`VK_PIPELINE_STAGE_2_OPTICAL_FLOW_BIT_NV`

`VK_ACCESS_2_OPTICAL_FLOW_WRITE_BIT_NV`

`VK_PIPELINE_STAGE_2_OPTICAL_FLOW_BIT_NV`

`VK_ACCESS_2_SHADER_TILE_ATTACHMENT_READ_BIT_QCOM`

`VK_PIPELINE_STAGE_2_FRAGMENT_SHADER_BIT`, `VK_PIPELINE_STAGE_2_COMPUTE_SHADER_BIT`

`VK_ACCESS_2_SHADER_TILE_ATTACHMENT_WRITE_BIT_QCOM`

`VK_PIPELINE_STAGE_2_FRAGMENT_SHADER_BIT`, `VK_PIPELINE_STAGE_2_COMPUTE_SHADER_BIT`

`VK_ACCESS_2_DATA_GRAPH_READ_BIT_ARM`

`VK_PIPELINE_STAGE_2_DATA_GRAPH_BIT_ARM`

`VK_ACCESS_2_DATA_GRAPH_WRITE_BIT_ARM`

`VK_PIPELINE_STAGE_2_DATA_GRAPH_BIT_ARM`

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkAccessFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlags.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkAccessFlagBits)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0