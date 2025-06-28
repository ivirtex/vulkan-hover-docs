# VkAccessFlagBits2(3) Manual Page

## Name

VkAccessFlagBits2 - Access flags for VkAccessFlags2



## [](#_c_specification)C Specification

Bits which **can** be set in the `srcAccessMask` and `dstAccessMask` members of [VkMemoryBarrier2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryBarrier2KHR.html), [VkImageMemoryBarrier2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageMemoryBarrier2KHR.html), and [VkBufferMemoryBarrier2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferMemoryBarrier2KHR.html), specifying access behavior, are:

```c++
// Provided by VK_VERSION_1_3
// Flag bits for VkAccessFlagBits2
typedef VkFlags64 VkAccessFlagBits2;
static const VkAccessFlagBits2 VK_ACCESS_2_NONE = 0ULL;
static const VkAccessFlagBits2 VK_ACCESS_2_INDIRECT_COMMAND_READ_BIT = 0x00000001ULL;
static const VkAccessFlagBits2 VK_ACCESS_2_INDEX_READ_BIT = 0x00000002ULL;
static const VkAccessFlagBits2 VK_ACCESS_2_VERTEX_ATTRIBUTE_READ_BIT = 0x00000004ULL;
static const VkAccessFlagBits2 VK_ACCESS_2_UNIFORM_READ_BIT = 0x00000008ULL;
static const VkAccessFlagBits2 VK_ACCESS_2_INPUT_ATTACHMENT_READ_BIT = 0x00000010ULL;
static const VkAccessFlagBits2 VK_ACCESS_2_SHADER_READ_BIT = 0x00000020ULL;
static const VkAccessFlagBits2 VK_ACCESS_2_SHADER_WRITE_BIT = 0x00000040ULL;
static const VkAccessFlagBits2 VK_ACCESS_2_COLOR_ATTACHMENT_READ_BIT = 0x00000080ULL;
static const VkAccessFlagBits2 VK_ACCESS_2_COLOR_ATTACHMENT_WRITE_BIT = 0x00000100ULL;
static const VkAccessFlagBits2 VK_ACCESS_2_DEPTH_STENCIL_ATTACHMENT_READ_BIT = 0x00000200ULL;
static const VkAccessFlagBits2 VK_ACCESS_2_DEPTH_STENCIL_ATTACHMENT_WRITE_BIT = 0x00000400ULL;
static const VkAccessFlagBits2 VK_ACCESS_2_TRANSFER_READ_BIT = 0x00000800ULL;
static const VkAccessFlagBits2 VK_ACCESS_2_TRANSFER_WRITE_BIT = 0x00001000ULL;
static const VkAccessFlagBits2 VK_ACCESS_2_HOST_READ_BIT = 0x00002000ULL;
static const VkAccessFlagBits2 VK_ACCESS_2_HOST_WRITE_BIT = 0x00004000ULL;
static const VkAccessFlagBits2 VK_ACCESS_2_MEMORY_READ_BIT = 0x00008000ULL;
static const VkAccessFlagBits2 VK_ACCESS_2_MEMORY_WRITE_BIT = 0x00010000ULL;
static const VkAccessFlagBits2 VK_ACCESS_2_SHADER_SAMPLED_READ_BIT = 0x100000000ULL;
static const VkAccessFlagBits2 VK_ACCESS_2_SHADER_STORAGE_READ_BIT = 0x200000000ULL;
static const VkAccessFlagBits2 VK_ACCESS_2_SHADER_STORAGE_WRITE_BIT = 0x400000000ULL;
// Provided by VK_KHR_video_decode_queue
static const VkAccessFlagBits2 VK_ACCESS_2_VIDEO_DECODE_READ_BIT_KHR = 0x800000000ULL;
// Provided by VK_KHR_video_decode_queue
static const VkAccessFlagBits2 VK_ACCESS_2_VIDEO_DECODE_WRITE_BIT_KHR = 0x1000000000ULL;
// Provided by VK_KHR_video_encode_queue
static const VkAccessFlagBits2 VK_ACCESS_2_VIDEO_ENCODE_READ_BIT_KHR = 0x2000000000ULL;
// Provided by VK_KHR_video_encode_queue
static const VkAccessFlagBits2 VK_ACCESS_2_VIDEO_ENCODE_WRITE_BIT_KHR = 0x4000000000ULL;
// Provided by VK_QCOM_tile_shading
static const VkAccessFlagBits2 VK_ACCESS_2_SHADER_TILE_ATTACHMENT_READ_BIT_QCOM = 0x8000000000000ULL;
// Provided by VK_QCOM_tile_shading
static const VkAccessFlagBits2 VK_ACCESS_2_SHADER_TILE_ATTACHMENT_WRITE_BIT_QCOM = 0x10000000000000ULL;
// Provided by VK_KHR_synchronization2
static const VkAccessFlagBits2 VK_ACCESS_2_NONE_KHR = 0ULL;
// Provided by VK_KHR_synchronization2
static const VkAccessFlagBits2 VK_ACCESS_2_INDIRECT_COMMAND_READ_BIT_KHR = 0x00000001ULL;
// Provided by VK_KHR_synchronization2
static const VkAccessFlagBits2 VK_ACCESS_2_INDEX_READ_BIT_KHR = 0x00000002ULL;
// Provided by VK_KHR_synchronization2
static const VkAccessFlagBits2 VK_ACCESS_2_VERTEX_ATTRIBUTE_READ_BIT_KHR = 0x00000004ULL;
// Provided by VK_KHR_synchronization2
static const VkAccessFlagBits2 VK_ACCESS_2_UNIFORM_READ_BIT_KHR = 0x00000008ULL;
// Provided by VK_KHR_synchronization2
static const VkAccessFlagBits2 VK_ACCESS_2_INPUT_ATTACHMENT_READ_BIT_KHR = 0x00000010ULL;
// Provided by VK_KHR_synchronization2
static const VkAccessFlagBits2 VK_ACCESS_2_SHADER_READ_BIT_KHR = 0x00000020ULL;
// Provided by VK_KHR_synchronization2
static const VkAccessFlagBits2 VK_ACCESS_2_SHADER_WRITE_BIT_KHR = 0x00000040ULL;
// Provided by VK_KHR_synchronization2
static const VkAccessFlagBits2 VK_ACCESS_2_COLOR_ATTACHMENT_READ_BIT_KHR = 0x00000080ULL;
// Provided by VK_KHR_synchronization2
static const VkAccessFlagBits2 VK_ACCESS_2_COLOR_ATTACHMENT_WRITE_BIT_KHR = 0x00000100ULL;
// Provided by VK_KHR_synchronization2
static const VkAccessFlagBits2 VK_ACCESS_2_DEPTH_STENCIL_ATTACHMENT_READ_BIT_KHR = 0x00000200ULL;
// Provided by VK_KHR_synchronization2
static const VkAccessFlagBits2 VK_ACCESS_2_DEPTH_STENCIL_ATTACHMENT_WRITE_BIT_KHR = 0x00000400ULL;
// Provided by VK_KHR_synchronization2
static const VkAccessFlagBits2 VK_ACCESS_2_TRANSFER_READ_BIT_KHR = 0x00000800ULL;
// Provided by VK_KHR_synchronization2
static const VkAccessFlagBits2 VK_ACCESS_2_TRANSFER_WRITE_BIT_KHR = 0x00001000ULL;
// Provided by VK_KHR_synchronization2
static const VkAccessFlagBits2 VK_ACCESS_2_HOST_READ_BIT_KHR = 0x00002000ULL;
// Provided by VK_KHR_synchronization2
static const VkAccessFlagBits2 VK_ACCESS_2_HOST_WRITE_BIT_KHR = 0x00004000ULL;
// Provided by VK_KHR_synchronization2
static const VkAccessFlagBits2 VK_ACCESS_2_MEMORY_READ_BIT_KHR = 0x00008000ULL;
// Provided by VK_KHR_synchronization2
static const VkAccessFlagBits2 VK_ACCESS_2_MEMORY_WRITE_BIT_KHR = 0x00010000ULL;
// Provided by VK_KHR_synchronization2
static const VkAccessFlagBits2 VK_ACCESS_2_SHADER_SAMPLED_READ_BIT_KHR = 0x100000000ULL;
// Provided by VK_KHR_synchronization2
static const VkAccessFlagBits2 VK_ACCESS_2_SHADER_STORAGE_READ_BIT_KHR = 0x200000000ULL;
// Provided by VK_KHR_synchronization2
static const VkAccessFlagBits2 VK_ACCESS_2_SHADER_STORAGE_WRITE_BIT_KHR = 0x400000000ULL;
// Provided by VK_KHR_synchronization2 with VK_EXT_transform_feedback
static const VkAccessFlagBits2 VK_ACCESS_2_TRANSFORM_FEEDBACK_WRITE_BIT_EXT = 0x02000000ULL;
// Provided by VK_KHR_synchronization2 with VK_EXT_transform_feedback
static const VkAccessFlagBits2 VK_ACCESS_2_TRANSFORM_FEEDBACK_COUNTER_READ_BIT_EXT = 0x04000000ULL;
// Provided by VK_KHR_synchronization2 with VK_EXT_transform_feedback
static const VkAccessFlagBits2 VK_ACCESS_2_TRANSFORM_FEEDBACK_COUNTER_WRITE_BIT_EXT = 0x08000000ULL;
// Provided by VK_KHR_synchronization2 with VK_EXT_conditional_rendering
static const VkAccessFlagBits2 VK_ACCESS_2_CONDITIONAL_RENDERING_READ_BIT_EXT = 0x00100000ULL;
// Provided by VK_KHR_synchronization2 with VK_NV_device_generated_commands
static const VkAccessFlagBits2 VK_ACCESS_2_COMMAND_PREPROCESS_READ_BIT_NV = 0x00020000ULL;
// Provided by VK_KHR_synchronization2 with VK_NV_device_generated_commands
static const VkAccessFlagBits2 VK_ACCESS_2_COMMAND_PREPROCESS_WRITE_BIT_NV = 0x00040000ULL;
// Provided by VK_KHR_synchronization2 with VK_EXT_device_generated_commands
static const VkAccessFlagBits2 VK_ACCESS_2_COMMAND_PREPROCESS_READ_BIT_EXT = 0x00020000ULL;
// Provided by VK_KHR_synchronization2 with VK_EXT_device_generated_commands
static const VkAccessFlagBits2 VK_ACCESS_2_COMMAND_PREPROCESS_WRITE_BIT_EXT = 0x00040000ULL;
// Provided by VK_KHR_fragment_shading_rate with VK_KHR_synchronization2
static const VkAccessFlagBits2 VK_ACCESS_2_FRAGMENT_SHADING_RATE_ATTACHMENT_READ_BIT_KHR = 0x00800000ULL;
// Provided by VK_KHR_synchronization2 with VK_NV_shading_rate_image
static const VkAccessFlagBits2 VK_ACCESS_2_SHADING_RATE_IMAGE_READ_BIT_NV = 0x00800000ULL;
// Provided by VK_KHR_acceleration_structure with VK_KHR_synchronization2
static const VkAccessFlagBits2 VK_ACCESS_2_ACCELERATION_STRUCTURE_READ_BIT_KHR = 0x00200000ULL;
// Provided by VK_KHR_acceleration_structure with VK_KHR_synchronization2
static const VkAccessFlagBits2 VK_ACCESS_2_ACCELERATION_STRUCTURE_WRITE_BIT_KHR = 0x00400000ULL;
// Provided by VK_KHR_synchronization2 with VK_NV_ray_tracing
static const VkAccessFlagBits2 VK_ACCESS_2_ACCELERATION_STRUCTURE_READ_BIT_NV = 0x00200000ULL;
// Provided by VK_KHR_synchronization2 with VK_NV_ray_tracing
static const VkAccessFlagBits2 VK_ACCESS_2_ACCELERATION_STRUCTURE_WRITE_BIT_NV = 0x00400000ULL;
// Provided by VK_KHR_synchronization2 with VK_EXT_fragment_density_map
static const VkAccessFlagBits2 VK_ACCESS_2_FRAGMENT_DENSITY_MAP_READ_BIT_EXT = 0x01000000ULL;
// Provided by VK_KHR_synchronization2 with VK_EXT_blend_operation_advanced
static const VkAccessFlagBits2 VK_ACCESS_2_COLOR_ATTACHMENT_READ_NONCOHERENT_BIT_EXT = 0x00080000ULL;
// Provided by VK_EXT_descriptor_buffer
static const VkAccessFlagBits2 VK_ACCESS_2_DESCRIPTOR_BUFFER_READ_BIT_EXT = 0x20000000000ULL;
// Provided by VK_HUAWEI_invocation_mask
static const VkAccessFlagBits2 VK_ACCESS_2_INVOCATION_MASK_READ_BIT_HUAWEI = 0x8000000000ULL;
// Provided by VK_KHR_ray_tracing_maintenance1 with (VK_KHR_synchronization2 or VK_VERSION_1_3) and VK_KHR_ray_tracing_pipeline
static const VkAccessFlagBits2 VK_ACCESS_2_SHADER_BINDING_TABLE_READ_BIT_KHR = 0x10000000000ULL;
// Provided by VK_EXT_opacity_micromap
static const VkAccessFlagBits2 VK_ACCESS_2_MICROMAP_READ_BIT_EXT = 0x100000000000ULL;
// Provided by VK_EXT_opacity_micromap
static const VkAccessFlagBits2 VK_ACCESS_2_MICROMAP_WRITE_BIT_EXT = 0x200000000000ULL;
// Provided by VK_NV_optical_flow
static const VkAccessFlagBits2 VK_ACCESS_2_OPTICAL_FLOW_READ_BIT_NV = 0x40000000000ULL;
// Provided by VK_NV_optical_flow
static const VkAccessFlagBits2 VK_ACCESS_2_OPTICAL_FLOW_WRITE_BIT_NV = 0x80000000000ULL;
// Provided by VK_ARM_data_graph
static const VkAccessFlagBits2 VK_ACCESS_2_DATA_GRAPH_READ_BIT_ARM = 0x800000000000ULL;
// Provided by VK_ARM_data_graph
static const VkAccessFlagBits2 VK_ACCESS_2_DATA_GRAPH_WRITE_BIT_ARM = 0x1000000000000ULL;
```

or the equivalent

```c++
// Provided by VK_KHR_synchronization2
typedef VkAccessFlagBits2 VkAccessFlagBits2KHR;
```

## [](#_description)Description

- `VK_ACCESS_2_NONE` specifies no accesses.
- `VK_ACCESS_2_MEMORY_READ_BIT` specifies all read accesses. It is always valid in any access mask, and is treated as equivalent to setting all `READ` access flags that are valid where it is used.
- `VK_ACCESS_2_MEMORY_WRITE_BIT` specifies all write accesses. It is always valid in any access mask, and is treated as equivalent to setting all `WRITE` access flags that are valid where it is used.
- `VK_ACCESS_2_INDIRECT_COMMAND_READ_BIT` specifies read access to command data read from indirect buffers as part of an indirect build, trace, drawing or dispatch command. Such access occurs in the `VK_PIPELINE_STAGE_2_DRAW_INDIRECT_BIT` pipeline stage.
- `VK_ACCESS_2_INDEX_READ_BIT` specifies read access to an index buffer as part of an indexed drawing command, bound by [vkCmdBindIndexBuffer2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBindIndexBuffer2.html) and [vkCmdBindIndexBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBindIndexBuffer.html). Such access occurs in the `VK_PIPELINE_STAGE_2_INDEX_INPUT_BIT` pipeline stage.
- `VK_ACCESS_2_VERTEX_ATTRIBUTE_READ_BIT` specifies read access to a vertex buffer as part of a drawing command, bound by [vkCmdBindVertexBuffers](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBindVertexBuffers.html). Such access occurs in the `VK_PIPELINE_STAGE_2_VERTEX_ATTRIBUTE_INPUT_BIT` pipeline stage.
- `VK_ACCESS_2_UNIFORM_READ_BIT` specifies read access to a [uniform buffer](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-uniformbuffer) in any shader pipeline stage.
- `VK_ACCESS_2_INPUT_ATTACHMENT_READ_BIT` specifies read access to an [input attachment](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass) within a render pass during subpass shading or fragment shading. Such access occurs in the `VK_PIPELINE_STAGE_2_SUBPASS_SHADER_BIT_HUAWEI` or `VK_PIPELINE_STAGE_2_FRAGMENT_SHADER_BIT` pipeline stage.
- `VK_ACCESS_2_SHADER_SAMPLED_READ_BIT` specifies read access to a [uniform texel buffer](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-uniformtexelbuffer) or [sampled image](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-sampledimage) in any shader pipeline stage.
- `VK_ACCESS_2_SHADER_STORAGE_READ_BIT` specifies read access to a [storage buffer](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-storagebuffer), [physical storage buffer](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-physical-storage-buffer), [storage texel buffer](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-storagetexelbuffer), or [storage image](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-storageimage) in any shader pipeline stage.
- `VK_ACCESS_2_SHADER_BINDING_TABLE_READ_BIT_KHR` specifies read access to a [shader binding table](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shader-binding-table) in any shader pipeline stage.
- `VK_ACCESS_2_SHADER_READ_BIT` is equivalent to the logical OR of:
  
  - `VK_ACCESS_2_SHADER_SAMPLED_READ_BIT`
  - `VK_ACCESS_2_SHADER_STORAGE_READ_BIT`
  - `VK_ACCESS_2_SHADER_TILE_ATTACHMENT_READ_BIT_QCOM`
- `VK_ACCESS_2_SHADER_STORAGE_WRITE_BIT` specifies write access to a [storage buffer](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-storagebuffer), [physical storage buffer](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-physical-storage-buffer), [storage texel buffer](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-storagetexelbuffer), or [storage image](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-storageimage) in any shader pipeline stage.
- `VK_ACCESS_2_SHADER_WRITE_BIT` is equivalent to `VK_ACCESS_2_SHADER_STORAGE_WRITE_BIT`.
- `VK_ACCESS_2_COLOR_ATTACHMENT_READ_BIT` specifies read access to a [color attachment](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass), such as via [blending](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#framebuffer-blending) (other than [advanced blend operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#framebuffer-blend-advanced)), [logic operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#framebuffer-logicop) or certain [render pass load operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-load-operations) in the `VK_PIPELINE_STAGE_2_COLOR_ATTACHMENT_OUTPUT_BIT` pipeline stage or via [fragment shader tile image reads](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#fragops-shader-tileimage-reads) in the `VK_PIPELINE_STAGE_2_FRAGMENT_SHADER_BIT` pipeline stage.
- `VK_ACCESS_2_COLOR_ATTACHMENT_WRITE_BIT` specifies write access to a [color attachment](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass) during a [render pass](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass) or via certain render pass [load](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-load-operations), [store](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-store-operations), and [multisample resolve](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-resolve-operations) operations. Such access occurs in the `VK_PIPELINE_STAGE_2_COLOR_ATTACHMENT_OUTPUT_BIT` pipeline stage.
- `VK_ACCESS_2_DEPTH_STENCIL_ATTACHMENT_READ_BIT` specifies read access to a [depth/stencil attachment](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass), via [depth or stencil operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#fragops-ds-state) or certain [render pass load operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-load-operations) in the `VK_PIPELINE_STAGE_2_EARLY_FRAGMENT_TESTS_BIT` or `VK_PIPELINE_STAGE_2_LATE_FRAGMENT_TESTS_BIT` pipeline stages or via [fragment shader tile image reads](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#fragops-shader-tileimage-reads) in the `VK_PIPELINE_STAGE_2_FRAGMENT_SHADER_BIT` pipeline stage.
- `VK_ACCESS_2_DEPTH_STENCIL_ATTACHMENT_WRITE_BIT` specifies write access to a [depth/stencil attachment](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass), via [depth or stencil operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#fragops-ds-state) or certain render pass [load](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-load-operations) and [store](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-store-operations) operations. Such access occurs in the `VK_PIPELINE_STAGE_2_EARLY_FRAGMENT_TESTS_BIT` or `VK_PIPELINE_STAGE_2_LATE_FRAGMENT_TESTS_BIT` pipeline stages.
- `VK_ACCESS_2_TRANSFER_READ_BIT` specifies read access to an image or buffer in a [copy](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#copies) operation. Such access occurs in the `VK_PIPELINE_STAGE_2_COPY_BIT`, `VK_PIPELINE_STAGE_2_BLIT_BIT`, or `VK_PIPELINE_STAGE_2_RESOLVE_BIT` pipeline stages.
- `VK_ACCESS_2_TRANSFER_WRITE_BIT` specifies write access to an image or buffer in a [clear](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#clears) or [copy](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#copies) operation. Such access occurs in the `VK_PIPELINE_STAGE_2_COPY_BIT`, `VK_PIPELINE_STAGE_2_BLIT_BIT`, `VK_PIPELINE_STAGE_2_CLEAR_BIT`, or `VK_PIPELINE_STAGE_2_RESOLVE_BIT` pipeline stages.
- `VK_ACCESS_2_HOST_READ_BIT` specifies read access by a host operation. Accesses of this type are not performed through a resource, but directly on memory. Such access occurs in the `VK_PIPELINE_STAGE_2_HOST_BIT` pipeline stage.
- `VK_ACCESS_2_HOST_WRITE_BIT` specifies write access by a host operation. Accesses of this type are not performed through a resource, but directly on memory. Such access occurs in the `VK_PIPELINE_STAGE_2_HOST_BIT` pipeline stage.
- `VK_ACCESS_2_CONDITIONAL_RENDERING_READ_BIT_EXT` specifies read access to a predicate as part of conditional rendering. Such access occurs in the `VK_PIPELINE_STAGE_2_CONDITIONAL_RENDERING_BIT_EXT` pipeline stage.
- `VK_ACCESS_2_TRANSFORM_FEEDBACK_WRITE_BIT_EXT` specifies write access to a transform feedback buffer made when transform feedback is active. Such access occurs in the `VK_PIPELINE_STAGE_2_TRANSFORM_FEEDBACK_BIT_EXT` pipeline stage.
- `VK_ACCESS_2_TRANSFORM_FEEDBACK_COUNTER_READ_BIT_EXT` specifies read access to a transform feedback counter buffer which is read when [vkCmdBeginTransformFeedbackEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBeginTransformFeedbackEXT.html) executes. Such access occurs in the `VK_PIPELINE_STAGE_2_TRANSFORM_FEEDBACK_BIT_EXT` pipeline stage.
- `VK_ACCESS_2_TRANSFORM_FEEDBACK_COUNTER_WRITE_BIT_EXT` specifies write access to a transform feedback counter buffer which is written when [vkCmdEndTransformFeedbackEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdEndTransformFeedbackEXT.html) executes. Such access occurs in the `VK_PIPELINE_STAGE_2_TRANSFORM_FEEDBACK_BIT_EXT` pipeline stage.
- `VK_ACCESS_2_COMMAND_PREPROCESS_READ_BIT_NV` specifies reads from buffer inputs to [vkCmdPreprocessGeneratedCommandsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdPreprocessGeneratedCommandsNV.html). Such access occurs in the `VK_PIPELINE_STAGE_2_COMMAND_PREPROCESS_BIT_NV` pipeline stage.
- `VK_ACCESS_2_COMMAND_PREPROCESS_WRITE_BIT_NV` specifies writes to the target command buffer preprocess outputs. Such access occurs in the `VK_PIPELINE_STAGE_2_COMMAND_PREPROCESS_BIT_NV` pipeline stage.
- `VK_ACCESS_2_COMMAND_PREPROCESS_READ_BIT_EXT` specifies reads from buffer inputs to [vkCmdPreprocessGeneratedCommandsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdPreprocessGeneratedCommandsEXT.html). Such access occurs in the `VK_PIPELINE_STAGE_2_COMMAND_PREPROCESS_BIT_EXT` pipeline stage.
- `VK_ACCESS_2_COMMAND_PREPROCESS_WRITE_BIT_EXT` specifies writes to the target command buffer preprocess outputs. Such access occurs in the `VK_PIPELINE_STAGE_2_COMMAND_PREPROCESS_BIT_EXT` pipeline stage.
- `VK_ACCESS_2_COLOR_ATTACHMENT_READ_NONCOHERENT_BIT_EXT` specifies read access to [color attachments](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass), including [advanced blend operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#framebuffer-blend-advanced). Such access occurs in the `VK_PIPELINE_STAGE_2_COLOR_ATTACHMENT_OUTPUT_BIT` pipeline stage.
- `VK_ACCESS_2_INVOCATION_MASK_READ_BIT_HUAWEI` specifies read access to an invocation mask image in the `VK_PIPELINE_STAGE_2_INVOCATION_MASK_BIT_HUAWEI` pipeline stage.
- `VK_ACCESS_2_ACCELERATION_STRUCTURE_READ_BIT_KHR` specifies read access to an acceleration structure as part of a trace, build, or copy command, or to an [acceleration structure scratch buffer](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#acceleration-structure-scratch) as part of a build command. Such access occurs in the `VK_PIPELINE_STAGE_2_RAY_TRACING_SHADER_BIT_KHR` pipeline stage or `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR` pipeline stage.
- `VK_ACCESS_2_ACCELERATION_STRUCTURE_WRITE_BIT_KHR` specifies write access to an acceleration structure or [acceleration structure scratch buffer](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#acceleration-structure-scratch) as part of a build or copy command. Such access occurs in the `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR` pipeline stage.
- `VK_ACCESS_2_FRAGMENT_DENSITY_MAP_READ_BIT_EXT` specifies read access to a [fragment density map attachment](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-fragmentdensitymapattachment) during dynamic [fragment density map operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#fragmentdensitymapops). Such access occurs in the `VK_PIPELINE_STAGE_2_FRAGMENT_DENSITY_PROCESS_BIT_EXT` pipeline stage.
- `VK_ACCESS_2_FRAGMENT_SHADING_RATE_ATTACHMENT_READ_BIT_KHR` specifies read access to a fragment shading rate attachment during rasterization. Such access occurs in the `VK_PIPELINE_STAGE_2_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR` pipeline stage.
- `VK_ACCESS_2_SHADING_RATE_IMAGE_READ_BIT_NV` specifies read access to a shading rate image during rasterization. Such access occurs in the `VK_PIPELINE_STAGE_2_SHADING_RATE_IMAGE_BIT_NV` pipeline stage. It is equivalent to `VK_ACCESS_2_FRAGMENT_SHADING_RATE_ATTACHMENT_READ_BIT_KHR`.
- `VK_ACCESS_2_VIDEO_DECODE_READ_BIT_KHR` specifies read access to an image or buffer resource in a [video decode operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#video-decode-operations). Such access occurs in the `VK_PIPELINE_STAGE_2_VIDEO_DECODE_BIT_KHR` pipeline stage.
- `VK_ACCESS_2_VIDEO_DECODE_WRITE_BIT_KHR` specifies write access to an image or buffer resource in a [video decode operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#video-decode-operations). Such access occurs in the `VK_PIPELINE_STAGE_2_VIDEO_DECODE_BIT_KHR` pipeline stage.
- `VK_ACCESS_2_VIDEO_ENCODE_READ_BIT_KHR` specifies read access to an image or buffer resource in a [video encode operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#video-encode-operations). Such access occurs in the `VK_PIPELINE_STAGE_2_VIDEO_ENCODE_BIT_KHR` pipeline stage.
- `VK_ACCESS_2_VIDEO_ENCODE_WRITE_BIT_KHR` specifies write access to an image or buffer resource in a [video encode operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#video-encode-operations). Such access occurs in the `VK_PIPELINE_STAGE_2_VIDEO_ENCODE_BIT_KHR` pipeline stage.
- `VK_ACCESS_2_DESCRIPTOR_BUFFER_READ_BIT_EXT` specifies read access to a [descriptor buffer](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorbuffers) in any shader pipeline stage.
- `VK_ACCESS_2_OPTICAL_FLOW_READ_BIT_NV` specifies read access to an image or buffer resource as part of a [optical flow operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#opticalflow-operations). Such access occurs in the `VK_PIPELINE_STAGE_2_OPTICAL_FLOW_BIT_NV` pipeline stage.
- `VK_ACCESS_2_OPTICAL_FLOW_WRITE_BIT_NV` specifies write access to an image or buffer resource as part of a [optical flow operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#opticalflow-operations). Such access occurs in the `VK_PIPELINE_STAGE_2_OPTICAL_FLOW_BIT_NV` pipeline stage.
- `VK_ACCESS_2_MICROMAP_WRITE_BIT_EXT` specifies write access to a micromap object. Such access occurs in the `VK_PIPELINE_STAGE_2_MICROMAP_BUILD_BIT_EXT` pipeline stage.
- `VK_ACCESS_2_MICROMAP_READ_BIT_EXT` specifies read access to a micromap object. Such access occurs in the `VK_PIPELINE_STAGE_2_MICROMAP_BUILD_BIT_EXT` and `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR` pipeline stages.
- `VK_ACCESS_2_SHADER_TILE_ATTACHMENT_READ_BIT_QCOM` specifies read access to a [tile attachment](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-tile-shading-attachment-access). Such access occurs in the `VK_PIPELINE_STAGE_2_FRAGMENT_SHADER_BIT` or `VK_PIPELINE_STAGE_2_COMPUTE_SHADER_BIT` pipeline stages.
- `VK_ACCESS_2_SHADER_TILE_ATTACHMENT_WRITE_BIT_QCOM` specifies write access to a [tile attachment](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-tile-shading-attachment-access). Such access occurs in the `VK_PIPELINE_STAGE_2_FRAGMENT_SHADER_BIT` or `VK_PIPELINE_STAGE_2_COMPUTE_SHADER_BIT` pipeline stages.
- `VK_ACCESS_2_DATA_GRAPH_READ_BIT_ARM` specifies read access to resources in the `VK_PIPELINE_STAGE_2_DATA_GRAPH_BIT_ARM` pipeline stage.
- `VK_ACCESS_2_DATA_GRAPH_WRITE_BIT_ARM` specifies write access to resources in the `VK_PIPELINE_STAGE_2_DATA_GRAPH_BIT_ARM` pipeline stage.

Note

In situations where an application wishes to select all access types for a given set of pipeline stages, `VK_ACCESS_2_MEMORY_READ_BIT` or `VK_ACCESS_2_MEMORY_WRITE_BIT` can be used. This is particularly useful when specifying stages that only have a single access type.

Note

The `VkAccessFlags2` bitmask goes beyond the 31 individual bit flags allowable within a C99 enum, which is how [VkAccessFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlagBits.html) is defined. The first 31 values are common to both, and are interchangeable.

## [](#_see_also)See Also

[VK\_KHR\_synchronization2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_synchronization2.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkAccessFlags2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlags2.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkAccessFlagBits2)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0