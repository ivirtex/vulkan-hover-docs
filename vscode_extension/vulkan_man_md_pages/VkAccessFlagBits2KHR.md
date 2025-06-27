# VkAccessFlagBits2(3) Manual Page

## Name

VkAccessFlagBits2 - Access flags for VkAccessFlags2



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **can** be set in the `srcAccessMask` and `dstAccessMask`
members of [VkMemoryBarrier2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryBarrier2KHR.html),
[VkImageMemoryBarrier2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageMemoryBarrier2KHR.html), and
[VkBufferMemoryBarrier2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferMemoryBarrier2KHR.html), specifying
access behavior, are:

``` c
// Provided by VK_VERSION_1_3
// Flag bits for VkAccessFlagBits2
typedef VkFlags64 VkAccessFlagBits2;
static const VkAccessFlagBits2 VK_ACCESS_2_NONE = 0ULL;
static const VkAccessFlagBits2 VK_ACCESS_2_NONE_KHR = 0ULL;
static const VkAccessFlagBits2 VK_ACCESS_2_INDIRECT_COMMAND_READ_BIT = 0x00000001ULL;
static const VkAccessFlagBits2 VK_ACCESS_2_INDIRECT_COMMAND_READ_BIT_KHR = 0x00000001ULL;
static const VkAccessFlagBits2 VK_ACCESS_2_INDEX_READ_BIT = 0x00000002ULL;
static const VkAccessFlagBits2 VK_ACCESS_2_INDEX_READ_BIT_KHR = 0x00000002ULL;
static const VkAccessFlagBits2 VK_ACCESS_2_VERTEX_ATTRIBUTE_READ_BIT = 0x00000004ULL;
static const VkAccessFlagBits2 VK_ACCESS_2_VERTEX_ATTRIBUTE_READ_BIT_KHR = 0x00000004ULL;
static const VkAccessFlagBits2 VK_ACCESS_2_UNIFORM_READ_BIT = 0x00000008ULL;
static const VkAccessFlagBits2 VK_ACCESS_2_UNIFORM_READ_BIT_KHR = 0x00000008ULL;
static const VkAccessFlagBits2 VK_ACCESS_2_INPUT_ATTACHMENT_READ_BIT = 0x00000010ULL;
static const VkAccessFlagBits2 VK_ACCESS_2_INPUT_ATTACHMENT_READ_BIT_KHR = 0x00000010ULL;
static const VkAccessFlagBits2 VK_ACCESS_2_SHADER_READ_BIT = 0x00000020ULL;
static const VkAccessFlagBits2 VK_ACCESS_2_SHADER_READ_BIT_KHR = 0x00000020ULL;
static const VkAccessFlagBits2 VK_ACCESS_2_SHADER_WRITE_BIT = 0x00000040ULL;
static const VkAccessFlagBits2 VK_ACCESS_2_SHADER_WRITE_BIT_KHR = 0x00000040ULL;
static const VkAccessFlagBits2 VK_ACCESS_2_COLOR_ATTACHMENT_READ_BIT = 0x00000080ULL;
static const VkAccessFlagBits2 VK_ACCESS_2_COLOR_ATTACHMENT_READ_BIT_KHR = 0x00000080ULL;
static const VkAccessFlagBits2 VK_ACCESS_2_COLOR_ATTACHMENT_WRITE_BIT = 0x00000100ULL;
static const VkAccessFlagBits2 VK_ACCESS_2_COLOR_ATTACHMENT_WRITE_BIT_KHR = 0x00000100ULL;
static const VkAccessFlagBits2 VK_ACCESS_2_DEPTH_STENCIL_ATTACHMENT_READ_BIT = 0x00000200ULL;
static const VkAccessFlagBits2 VK_ACCESS_2_DEPTH_STENCIL_ATTACHMENT_READ_BIT_KHR = 0x00000200ULL;
static const VkAccessFlagBits2 VK_ACCESS_2_DEPTH_STENCIL_ATTACHMENT_WRITE_BIT = 0x00000400ULL;
static const VkAccessFlagBits2 VK_ACCESS_2_DEPTH_STENCIL_ATTACHMENT_WRITE_BIT_KHR = 0x00000400ULL;
static const VkAccessFlagBits2 VK_ACCESS_2_TRANSFER_READ_BIT = 0x00000800ULL;
static const VkAccessFlagBits2 VK_ACCESS_2_TRANSFER_READ_BIT_KHR = 0x00000800ULL;
static const VkAccessFlagBits2 VK_ACCESS_2_TRANSFER_WRITE_BIT = 0x00001000ULL;
static const VkAccessFlagBits2 VK_ACCESS_2_TRANSFER_WRITE_BIT_KHR = 0x00001000ULL;
static const VkAccessFlagBits2 VK_ACCESS_2_HOST_READ_BIT = 0x00002000ULL;
static const VkAccessFlagBits2 VK_ACCESS_2_HOST_READ_BIT_KHR = 0x00002000ULL;
static const VkAccessFlagBits2 VK_ACCESS_2_HOST_WRITE_BIT = 0x00004000ULL;
static const VkAccessFlagBits2 VK_ACCESS_2_HOST_WRITE_BIT_KHR = 0x00004000ULL;
static const VkAccessFlagBits2 VK_ACCESS_2_MEMORY_READ_BIT = 0x00008000ULL;
static const VkAccessFlagBits2 VK_ACCESS_2_MEMORY_READ_BIT_KHR = 0x00008000ULL;
static const VkAccessFlagBits2 VK_ACCESS_2_MEMORY_WRITE_BIT = 0x00010000ULL;
static const VkAccessFlagBits2 VK_ACCESS_2_MEMORY_WRITE_BIT_KHR = 0x00010000ULL;
static const VkAccessFlagBits2 VK_ACCESS_2_SHADER_SAMPLED_READ_BIT = 0x100000000ULL;
static const VkAccessFlagBits2 VK_ACCESS_2_SHADER_SAMPLED_READ_BIT_KHR = 0x100000000ULL;
static const VkAccessFlagBits2 VK_ACCESS_2_SHADER_STORAGE_READ_BIT = 0x200000000ULL;
static const VkAccessFlagBits2 VK_ACCESS_2_SHADER_STORAGE_READ_BIT_KHR = 0x200000000ULL;
static const VkAccessFlagBits2 VK_ACCESS_2_SHADER_STORAGE_WRITE_BIT = 0x400000000ULL;
static const VkAccessFlagBits2 VK_ACCESS_2_SHADER_STORAGE_WRITE_BIT_KHR = 0x400000000ULL;
// Provided by VK_KHR_video_decode_queue
static const VkAccessFlagBits2 VK_ACCESS_2_VIDEO_DECODE_READ_BIT_KHR = 0x800000000ULL;
// Provided by VK_KHR_video_decode_queue
static const VkAccessFlagBits2 VK_ACCESS_2_VIDEO_DECODE_WRITE_BIT_KHR = 0x1000000000ULL;
// Provided by VK_KHR_video_encode_queue
static const VkAccessFlagBits2 VK_ACCESS_2_VIDEO_ENCODE_READ_BIT_KHR = 0x2000000000ULL;
// Provided by VK_KHR_video_encode_queue
static const VkAccessFlagBits2 VK_ACCESS_2_VIDEO_ENCODE_WRITE_BIT_KHR = 0x4000000000ULL;
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
```

or the equivalent

``` c
// Provided by VK_KHR_synchronization2
typedef VkAccessFlagBits2 VkAccessFlagBits2KHR;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_ACCESS_2_NONE` specifies no accesses.

- `VK_ACCESS_2_MEMORY_READ_BIT` specifies all read accesses. It is
  always valid in any access mask, and is treated as equivalent to
  setting all `READ` access flags that are valid where it is used.

- `VK_ACCESS_2_MEMORY_WRITE_BIT` specifies all write accesses. It is
  always valid in any access mask, and is treated as equivalent to
  setting all `WRITE` access flags that are valid where it is used.

- `VK_ACCESS_2_INDIRECT_COMMAND_READ_BIT` specifies read access to
  command data read from indirect buffers as part of an indirect build,
  trace, drawing or dispatch command. Such access occurs in the
  `VK_PIPELINE_STAGE_2_DRAW_INDIRECT_BIT` pipeline stage.

- `VK_ACCESS_2_INDEX_READ_BIT` specifies read access to an index buffer
  as part of an indexed drawing command, bound by
  [vkCmdBindIndexBuffer2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBindIndexBuffer2KHR.html) and
  [vkCmdBindIndexBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBindIndexBuffer.html). Such access occurs
  in the `VK_PIPELINE_STAGE_2_INDEX_INPUT_BIT` pipeline stage.

- `VK_ACCESS_2_VERTEX_ATTRIBUTE_READ_BIT` specifies read access to a
  vertex buffer as part of a drawing command, bound by
  [vkCmdBindVertexBuffers](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBindVertexBuffers.html). Such access
  occurs in the `VK_PIPELINE_STAGE_2_VERTEX_ATTRIBUTE_INPUT_BIT`
  pipeline stage.

- `VK_ACCESS_2_UNIFORM_READ_BIT` specifies read access to a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-uniformbuffer"
  target="_blank" rel="noopener">uniform buffer</a> in any shader
  pipeline stage.

- `VK_ACCESS_2_INPUT_ATTACHMENT_READ_BIT` specifies read access to an <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass"
  target="_blank" rel="noopener">input attachment</a> within a render
  pass during subpass shading or fragment shading. Such access occurs in
  the `VK_PIPELINE_STAGE_2_SUBPASS_SHADER_BIT_HUAWEI` or
  `VK_PIPELINE_STAGE_2_FRAGMENT_SHADER_BIT` pipeline stage.

- `VK_ACCESS_2_SHADER_SAMPLED_READ_BIT` specifies read access to a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-uniformtexelbuffer"
  target="_blank" rel="noopener">uniform texel buffer</a> or <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-sampledimage"
  target="_blank" rel="noopener">sampled image</a> in any shader
  pipeline stage.

- `VK_ACCESS_2_SHADER_STORAGE_READ_BIT` specifies read access to a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-storagebuffer"
  target="_blank" rel="noopener">storage buffer</a>, <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-physical-storage-buffer"
  target="_blank" rel="noopener">physical storage buffer</a>, <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-storagetexelbuffer"
  target="_blank" rel="noopener">storage texel buffer</a>, or <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-storageimage"
  target="_blank" rel="noopener">storage image</a> in any shader
  pipeline stage.

- `VK_ACCESS_2_SHADER_BINDING_TABLE_READ_BIT_KHR` specifies read access
  to a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#shader-binding-table"
  target="_blank" rel="noopener">shader binding table</a> in any shader
  pipeline stage.

- `VK_ACCESS_2_SHADER_READ_BIT` is equivalent to the logical OR of:

  - `VK_ACCESS_2_SHADER_SAMPLED_READ_BIT`

  - `VK_ACCESS_2_SHADER_STORAGE_READ_BIT`

- `VK_ACCESS_2_SHADER_STORAGE_WRITE_BIT` specifies write access to a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-storagebuffer"
  target="_blank" rel="noopener">storage buffer</a>, <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-physical-storage-buffer"
  target="_blank" rel="noopener">physical storage buffer</a>, <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-storagetexelbuffer"
  target="_blank" rel="noopener">storage texel buffer</a>, or <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-storageimage"
  target="_blank" rel="noopener">storage image</a> in any shader
  pipeline stage.

- `VK_ACCESS_2_SHADER_WRITE_BIT` is equivalent to
  `VK_ACCESS_2_SHADER_STORAGE_WRITE_BIT`.

- `VK_ACCESS_2_COLOR_ATTACHMENT_READ_BIT` specifies read access to a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass"
  target="_blank" rel="noopener">color attachment</a>, such as via <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#framebuffer-blending"
  target="_blank" rel="noopener">blending</a> (other than <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#framebuffer-blend-advanced"
  target="_blank" rel="noopener">advanced blend operations</a>), <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#framebuffer-logicop"
  target="_blank" rel="noopener">logic operations</a> or certain <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass-load-operations"
  target="_blank" rel="noopener">render pass load operations</a> in the
  `VK_PIPELINE_STAGE_2_COLOR_ATTACHMENT_OUTPUT_BIT` pipeline stage or
  via <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fragops-shader-tileimage-reads"
  target="_blank" rel="noopener">fragment shader tile image reads</a> in
  the `VK_PIPELINE_STAGE_2_FRAGMENT_SHADER_BIT` pipeline stage.

- `VK_ACCESS_2_COLOR_ATTACHMENT_WRITE_BIT` specifies write access to a
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass"
  target="_blank" rel="noopener">color attachment</a> during a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass"
  target="_blank" rel="noopener">render pass</a> or via certain render
  pass <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass-load-operations"
  target="_blank" rel="noopener">load</a>, <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass-store-operations"
  target="_blank" rel="noopener">store</a>, and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass-resolve-operations"
  target="_blank" rel="noopener">multisample resolve</a> operations.
  Such access occurs in the
  `VK_PIPELINE_STAGE_2_COLOR_ATTACHMENT_OUTPUT_BIT` pipeline stage.

- `VK_ACCESS_2_DEPTH_STENCIL_ATTACHMENT_READ_BIT` specifies read access
  to a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass"
  target="_blank" rel="noopener">depth/stencil attachment</a>, via <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fragops-ds-state"
  target="_blank" rel="noopener">depth or stencil operations</a> or
  certain <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass-load-operations"
  target="_blank" rel="noopener">render pass load operations</a> in the
  `VK_PIPELINE_STAGE_2_EARLY_FRAGMENT_TESTS_BIT` or
  `VK_PIPELINE_STAGE_2_LATE_FRAGMENT_TESTS_BIT` pipeline stages or via
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fragops-shader-tileimage-reads"
  target="_blank" rel="noopener">fragment shader tile image reads</a> in
  the `VK_PIPELINE_STAGE_2_FRAGMENT_SHADER_BIT` pipeline stage.

- `VK_ACCESS_2_DEPTH_STENCIL_ATTACHMENT_WRITE_BIT` specifies write
  access to a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass"
  target="_blank" rel="noopener">depth/stencil attachment</a>, via <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fragops-ds-state"
  target="_blank" rel="noopener">depth or stencil operations</a> or
  certain render pass <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass-load-operations"
  target="_blank" rel="noopener">load</a> and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass-store-operations"
  target="_blank" rel="noopener">store</a> operations. Such access
  occurs in the `VK_PIPELINE_STAGE_2_EARLY_FRAGMENT_TESTS_BIT` or
  `VK_PIPELINE_STAGE_2_LATE_FRAGMENT_TESTS_BIT` pipeline stages.

- `VK_ACCESS_2_TRANSFER_READ_BIT` specifies read access to an image or
  buffer in a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#copies"
  target="_blank" rel="noopener">copy</a> operation. Such access occurs
  in the `VK_PIPELINE_STAGE_2_COPY_BIT`, `VK_PIPELINE_STAGE_2_BLIT_BIT`,
  or `VK_PIPELINE_STAGE_2_RESOLVE_BIT` pipeline stages.

- `VK_ACCESS_2_TRANSFER_WRITE_BIT` specifies write access to an image or
  buffer in a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#clears"
  target="_blank" rel="noopener">clear</a> or <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#copies"
  target="_blank" rel="noopener">copy</a> operation. Such access occurs
  in the `VK_PIPELINE_STAGE_2_COPY_BIT`, `VK_PIPELINE_STAGE_2_BLIT_BIT`,
  `VK_PIPELINE_STAGE_2_CLEAR_BIT`, or `VK_PIPELINE_STAGE_2_RESOLVE_BIT`
  pipeline stages.

- `VK_ACCESS_2_HOST_READ_BIT` specifies read access by a host operation.
  Accesses of this type are not performed through a resource, but
  directly on memory. Such access occurs in the
  `VK_PIPELINE_STAGE_2_HOST_BIT` pipeline stage.

- `VK_ACCESS_2_HOST_WRITE_BIT` specifies write access by a host
  operation. Accesses of this type are not performed through a resource,
  but directly on memory. Such access occurs in the
  `VK_PIPELINE_STAGE_2_HOST_BIT` pipeline stage.

- `VK_ACCESS_2_CONDITIONAL_RENDERING_READ_BIT_EXT` specifies read access
  to a predicate as part of conditional rendering. Such access occurs in
  the `VK_PIPELINE_STAGE_2_CONDITIONAL_RENDERING_BIT_EXT` pipeline
  stage.

- `VK_ACCESS_2_TRANSFORM_FEEDBACK_WRITE_BIT_EXT` specifies write access
  to a transform feedback buffer made when transform feedback is active.
  Such access occurs in the
  `VK_PIPELINE_STAGE_2_TRANSFORM_FEEDBACK_BIT_EXT` pipeline stage.

- `VK_ACCESS_2_TRANSFORM_FEEDBACK_COUNTER_READ_BIT_EXT` specifies read
  access to a transform feedback counter buffer which is read when
  [vkCmdBeginTransformFeedbackEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginTransformFeedbackEXT.html)
  executes. Such access occurs in the
  `VK_PIPELINE_STAGE_2_TRANSFORM_FEEDBACK_BIT_EXT` pipeline stage.

- `VK_ACCESS_2_TRANSFORM_FEEDBACK_COUNTER_WRITE_BIT_EXT` specifies write
  access to a transform feedback counter buffer which is written when
  [vkCmdEndTransformFeedbackEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdEndTransformFeedbackEXT.html)
  executes. Such access occurs in the
  `VK_PIPELINE_STAGE_2_TRANSFORM_FEEDBACK_BIT_EXT` pipeline stage.

- `VK_ACCESS_2_COMMAND_PREPROCESS_READ_BIT_NV` specifies reads from
  buffer inputs to
  [vkCmdPreprocessGeneratedCommandsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdPreprocessGeneratedCommandsNV.html).
  Such access occurs in the
  `VK_PIPELINE_STAGE_2_COMMAND_PREPROCESS_BIT_NV` pipeline stage.

- `VK_ACCESS_2_COMMAND_PREPROCESS_WRITE_BIT_NV` specifies writes to the
  target command buffer preprocess outputs. Such access occurs in the
  `VK_PIPELINE_STAGE_2_COMMAND_PREPROCESS_BIT_NV` pipeline stage.

- `VK_ACCESS_2_COLOR_ATTACHMENT_READ_NONCOHERENT_BIT_EXT` specifies read
  access to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass"
  target="_blank" rel="noopener">color attachments</a>, including <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#framebuffer-blend-advanced"
  target="_blank" rel="noopener">advanced blend operations</a>. Such
  access occurs in the `VK_PIPELINE_STAGE_2_COLOR_ATTACHMENT_OUTPUT_BIT`
  pipeline stage.

- `VK_ACCESS_2_INVOCATION_MASK_READ_BIT_HUAWEI` specifies read access to
  an invocation mask image in the
  `VK_PIPELINE_STAGE_2_INVOCATION_MASK_BIT_HUAWEI` pipeline stage.

- `VK_ACCESS_2_ACCELERATION_STRUCTURE_READ_BIT_KHR` specifies read
  access to an acceleration structure as part of a trace, build, or copy
  command, or to an <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#acceleration-structure-scratch"
  target="_blank" rel="noopener">acceleration structure scratch buffer</a>
  as part of a build command. Such access occurs in the
  `VK_PIPELINE_STAGE_2_RAY_TRACING_SHADER_BIT_KHR` pipeline stage or
  `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR` pipeline
  stage.

- `VK_ACCESS_2_ACCELERATION_STRUCTURE_WRITE_BIT_KHR` specifies write
  access to an acceleration structure or <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#acceleration-structure-scratch"
  target="_blank" rel="noopener">acceleration structure scratch buffer</a>
  as part of a build or copy command. Such access occurs in the
  `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR` pipeline
  stage.

- `VK_ACCESS_2_FRAGMENT_DENSITY_MAP_READ_BIT_EXT` specifies read access
  to a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass-fragmentdensitymapattachment"
  target="_blank" rel="noopener">fragment density map attachment</a>
  during dynamic <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fragmentdensitymapops"
  target="_blank" rel="noopener">fragment density map operations</a>.
  Such access occurs in the
  `VK_PIPELINE_STAGE_2_FRAGMENT_DENSITY_PROCESS_BIT_EXT` pipeline stage.

- `VK_ACCESS_2_FRAGMENT_SHADING_RATE_ATTACHMENT_READ_BIT_KHR` specifies
  read access to a fragment shading rate attachment during
  rasterization. Such access occurs in the
  `VK_PIPELINE_STAGE_2_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR`
  pipeline stage.

- `VK_ACCESS_2_SHADING_RATE_IMAGE_READ_BIT_NV` specifies read access to
  a shading rate image during rasterization. Such access occurs in the
  `VK_PIPELINE_STAGE_2_SHADING_RATE_IMAGE_BIT_NV` pipeline stage. It is
  equivalent to
  `VK_ACCESS_2_FRAGMENT_SHADING_RATE_ATTACHMENT_READ_BIT_KHR`.

- `VK_ACCESS_2_VIDEO_DECODE_READ_BIT_KHR` specifies read access to an
  image or buffer resource in a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-decode-operations"
  target="_blank" rel="noopener">video decode operation</a>. Such access
  occurs in the `VK_PIPELINE_STAGE_2_VIDEO_DECODE_BIT_KHR` pipeline
  stage.

- `VK_ACCESS_2_VIDEO_DECODE_WRITE_BIT_KHR` specifies write access to an
  image or buffer resource in a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-decode-operations"
  target="_blank" rel="noopener">video decode operation</a>. Such access
  occurs in the `VK_PIPELINE_STAGE_2_VIDEO_DECODE_BIT_KHR` pipeline
  stage.

- `VK_ACCESS_2_VIDEO_ENCODE_READ_BIT_KHR` specifies read access to an
  image or buffer resource in a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-encode-operations"
  target="_blank" rel="noopener">video encode operation</a>. Such access
  occurs in the `VK_PIPELINE_STAGE_2_VIDEO_ENCODE_BIT_KHR` pipeline
  stage.

- `VK_ACCESS_2_VIDEO_ENCODE_WRITE_BIT_KHR` specifies write access to an
  image or buffer resource in a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-encode-operations"
  target="_blank" rel="noopener">video encode operation</a>. Such access
  occurs in the `VK_PIPELINE_STAGE_2_VIDEO_ENCODE_BIT_KHR` pipeline
  stage.

- `VK_ACCESS_2_DESCRIPTOR_BUFFER_READ_BIT_EXT` specifies read access to
  a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorbuffers"
  target="_blank" rel="noopener">descriptor buffer</a> in any shader
  pipeline stage.

- `VK_ACCESS_2_OPTICAL_FLOW_READ_BIT_NV` specifies read access to an
  image or buffer resource as part of a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#opticalflow-operations"
  target="_blank" rel="noopener">optical flow operation</a>. Such access
  occurs in the `VK_PIPELINE_STAGE_2_OPTICAL_FLOW_BIT_NV` pipeline
  stage.

- `VK_ACCESS_2_OPTICAL_FLOW_WRITE_BIT_NV` specifies write access to an
  image or buffer resource as part of a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#opticalflow-operations"
  target="_blank" rel="noopener">optical flow operation</a>. Such access
  occurs in the `VK_PIPELINE_STAGE_2_OPTICAL_FLOW_BIT_NV` pipeline
  stage.

- `VK_ACCESS_2_MICROMAP_WRITE_BIT_EXT` specifies write access to a
  micromap object. Such access occurs in the
  `VK_PIPELINE_STAGE_2_MICROMAP_BUILD_BIT_EXT` pipeline stage.

- `VK_ACCESS_2_MICROMAP_READ_BIT_EXT` specifies read access to a
  micromap object. Such access occurs in the
  `VK_PIPELINE_STAGE_2_MICROMAP_BUILD_BIT_EXT` and
  `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR` pipeline
  stages.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>In situations where an application wishes to select all access types
for a given set of pipeline stages,
<code>VK_ACCESS_2_MEMORY_READ_BIT</code> or
<code>VK_ACCESS_2_MEMORY_WRITE_BIT</code> can be used. This is
particularly useful when specifying stages that only have a single
access type.</p></td>
</tr>
</tbody>
</table>

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>The <code>VkAccessFlags2</code> bitmask goes beyond the 31 individual
bit flags allowable within a C99 enum, which is how <a
href="VkAccessFlagBits.html">VkAccessFlagBits</a> is defined. The first
31 values are common to both, and are interchangeable.</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_synchronization2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_synchronization2.html),
[VK_VERSION_1_3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_3.html),
[VkAccessFlags2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccessFlags2.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkAccessFlagBits2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
