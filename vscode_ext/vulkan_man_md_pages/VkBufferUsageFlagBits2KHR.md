# VkBufferUsageFlagBits2KHR(3) Manual Page

## Name

VkBufferUsageFlagBits2KHR - Bitmask controlling how a pipeline is
created



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **can** be set in
[VkBufferUsageFlags2CreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferUsageFlags2CreateInfoKHR.html)::`usage`,
specifying usage behavior of a buffer, are:

``` c
// Provided by VK_KHR_maintenance5
// Flag bits for VkBufferUsageFlagBits2KHR
typedef VkFlags64 VkBufferUsageFlagBits2KHR;
static const VkBufferUsageFlagBits2KHR VK_BUFFER_USAGE_2_TRANSFER_SRC_BIT_KHR = 0x00000001ULL;
static const VkBufferUsageFlagBits2KHR VK_BUFFER_USAGE_2_TRANSFER_DST_BIT_KHR = 0x00000002ULL;
static const VkBufferUsageFlagBits2KHR VK_BUFFER_USAGE_2_UNIFORM_TEXEL_BUFFER_BIT_KHR = 0x00000004ULL;
static const VkBufferUsageFlagBits2KHR VK_BUFFER_USAGE_2_STORAGE_TEXEL_BUFFER_BIT_KHR = 0x00000008ULL;
static const VkBufferUsageFlagBits2KHR VK_BUFFER_USAGE_2_UNIFORM_BUFFER_BIT_KHR = 0x00000010ULL;
static const VkBufferUsageFlagBits2KHR VK_BUFFER_USAGE_2_STORAGE_BUFFER_BIT_KHR = 0x00000020ULL;
static const VkBufferUsageFlagBits2KHR VK_BUFFER_USAGE_2_INDEX_BUFFER_BIT_KHR = 0x00000040ULL;
static const VkBufferUsageFlagBits2KHR VK_BUFFER_USAGE_2_VERTEX_BUFFER_BIT_KHR = 0x00000080ULL;
static const VkBufferUsageFlagBits2KHR VK_BUFFER_USAGE_2_INDIRECT_BUFFER_BIT_KHR = 0x00000100ULL;
// Provided by VK_KHR_maintenance5 with VK_AMDX_shader_enqueue
static const VkBufferUsageFlagBits2KHR VK_BUFFER_USAGE_2_EXECUTION_GRAPH_SCRATCH_BIT_AMDX = 0x02000000ULL;
// Provided by VK_KHR_maintenance5 with VK_EXT_conditional_rendering
static const VkBufferUsageFlagBits2KHR VK_BUFFER_USAGE_2_CONDITIONAL_RENDERING_BIT_EXT = 0x00000200ULL;
// Provided by VK_KHR_maintenance5 with VK_KHR_ray_tracing_pipeline
static const VkBufferUsageFlagBits2KHR VK_BUFFER_USAGE_2_SHADER_BINDING_TABLE_BIT_KHR = 0x00000400ULL;
// Provided by VK_KHR_maintenance5 with VK_NV_ray_tracing
static const VkBufferUsageFlagBits2KHR VK_BUFFER_USAGE_2_RAY_TRACING_BIT_NV = 0x00000400ULL;
// Provided by VK_KHR_maintenance5 with VK_EXT_transform_feedback
static const VkBufferUsageFlagBits2KHR VK_BUFFER_USAGE_2_TRANSFORM_FEEDBACK_BUFFER_BIT_EXT = 0x00000800ULL;
// Provided by VK_KHR_maintenance5 with VK_EXT_transform_feedback
static const VkBufferUsageFlagBits2KHR VK_BUFFER_USAGE_2_TRANSFORM_FEEDBACK_COUNTER_BUFFER_BIT_EXT = 0x00001000ULL;
// Provided by VK_KHR_maintenance5 with VK_KHR_video_decode_queue
static const VkBufferUsageFlagBits2KHR VK_BUFFER_USAGE_2_VIDEO_DECODE_SRC_BIT_KHR = 0x00002000ULL;
// Provided by VK_KHR_maintenance5 with VK_KHR_video_decode_queue
static const VkBufferUsageFlagBits2KHR VK_BUFFER_USAGE_2_VIDEO_DECODE_DST_BIT_KHR = 0x00004000ULL;
// Provided by VK_KHR_maintenance5 with VK_KHR_video_encode_queue
static const VkBufferUsageFlagBits2KHR VK_BUFFER_USAGE_2_VIDEO_ENCODE_DST_BIT_KHR = 0x00008000ULL;
// Provided by VK_KHR_maintenance5 with VK_KHR_video_encode_queue
static const VkBufferUsageFlagBits2KHR VK_BUFFER_USAGE_2_VIDEO_ENCODE_SRC_BIT_KHR = 0x00010000ULL;
// Provided by VK_KHR_maintenance5 with (VK_VERSION_1_2 or VK_KHR_buffer_device_address) or VK_EXT_buffer_device_address
static const VkBufferUsageFlagBits2KHR VK_BUFFER_USAGE_2_SHADER_DEVICE_ADDRESS_BIT_KHR = 0x00020000ULL;
// Provided by VK_KHR_acceleration_structure with VK_KHR_maintenance5
static const VkBufferUsageFlagBits2KHR VK_BUFFER_USAGE_2_ACCELERATION_STRUCTURE_BUILD_INPUT_READ_ONLY_BIT_KHR = 0x00080000ULL;
// Provided by VK_KHR_acceleration_structure with VK_KHR_maintenance5
static const VkBufferUsageFlagBits2KHR VK_BUFFER_USAGE_2_ACCELERATION_STRUCTURE_STORAGE_BIT_KHR = 0x00100000ULL;
// Provided by VK_KHR_maintenance5 with VK_EXT_descriptor_buffer
static const VkBufferUsageFlagBits2KHR VK_BUFFER_USAGE_2_SAMPLER_DESCRIPTOR_BUFFER_BIT_EXT = 0x00200000ULL;
// Provided by VK_KHR_maintenance5 with VK_EXT_descriptor_buffer
static const VkBufferUsageFlagBits2KHR VK_BUFFER_USAGE_2_RESOURCE_DESCRIPTOR_BUFFER_BIT_EXT = 0x00400000ULL;
// Provided by VK_KHR_maintenance5 with VK_EXT_descriptor_buffer
static const VkBufferUsageFlagBits2KHR VK_BUFFER_USAGE_2_PUSH_DESCRIPTORS_DESCRIPTOR_BUFFER_BIT_EXT = 0x04000000ULL;
// Provided by VK_KHR_maintenance5 with VK_EXT_opacity_micromap
static const VkBufferUsageFlagBits2KHR VK_BUFFER_USAGE_2_MICROMAP_BUILD_INPUT_READ_ONLY_BIT_EXT = 0x00800000ULL;
// Provided by VK_KHR_maintenance5 with VK_EXT_opacity_micromap
static const VkBufferUsageFlagBits2KHR VK_BUFFER_USAGE_2_MICROMAP_STORAGE_BIT_EXT = 0x01000000ULL;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_BUFFER_USAGE_2_TRANSFER_SRC_BIT_KHR` specifies that the buffer
  **can** be used as the source of a *transfer command* (see the
  definition of <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-pipeline-stages-transfer"
  target="_blank"
  rel="noopener"><code>VK_PIPELINE_STAGE_TRANSFER_BIT</code></a>).

- `VK_BUFFER_USAGE_2_TRANSFER_DST_BIT_KHR` specifies that the buffer
  **can** be used as the destination of a transfer command.

- `VK_BUFFER_USAGE_2_UNIFORM_TEXEL_BUFFER_BIT_KHR` specifies that the
  buffer **can** be used to create a `VkBufferView` suitable for
  occupying a `VkDescriptorSet` slot of type
  `VK_DESCRIPTOR_TYPE_UNIFORM_TEXEL_BUFFER`.

- `VK_BUFFER_USAGE_2_STORAGE_TEXEL_BUFFER_BIT_KHR` specifies that the
  buffer **can** be used to create a `VkBufferView` suitable for
  occupying a `VkDescriptorSet` slot of type
  `VK_DESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFER`.

- `VK_BUFFER_USAGE_2_UNIFORM_BUFFER_BIT_KHR` specifies that the buffer
  **can** be used in a `VkDescriptorBufferInfo` suitable for occupying a
  `VkDescriptorSet` slot either of type
  `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER` or
  `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC`.

- `VK_BUFFER_USAGE_2_STORAGE_BUFFER_BIT_KHR` specifies that the buffer
  **can** be used in a `VkDescriptorBufferInfo` suitable for occupying a
  `VkDescriptorSet` slot either of type
  `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER` or
  `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC`.

- `VK_BUFFER_USAGE_2_INDEX_BUFFER_BIT_KHR` specifies that the buffer is
  suitable for passing as the `buffer` parameter to
  [vkCmdBindIndexBuffer2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBindIndexBuffer2KHR.html) and
  [vkCmdBindIndexBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBindIndexBuffer.html).

- `VK_BUFFER_USAGE_2_VERTEX_BUFFER_BIT_KHR` specifies that the buffer is
  suitable for passing as an element of the `pBuffers` array to
  [vkCmdBindVertexBuffers](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBindVertexBuffers.html).

- `VK_BUFFER_USAGE_2_INDIRECT_BUFFER_BIT_KHR` specifies that the buffer
  is suitable for passing as the `buffer` parameter to
  [vkCmdDrawIndirect](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDrawIndirect.html),
  [vkCmdDrawIndexedIndirect](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDrawIndexedIndirect.html),
  [vkCmdDrawMeshTasksIndirectNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDrawMeshTasksIndirectNV.html),
  [vkCmdDrawMeshTasksIndirectCountNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDrawMeshTasksIndirectCountNV.html),
  `vkCmdDrawMeshTasksIndirectEXT`, `vkCmdDrawMeshTasksIndirectCountEXT`,
  [vkCmdDrawClusterIndirectHUAWEI](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDrawClusterIndirectHUAWEI.html),
  or [vkCmdDispatchIndirect](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDispatchIndirect.html). It is also
  suitable for passing as the `buffer` member of
  `VkIndirectCommandsStreamNV`, or `sequencesCountBuffer` or
  `sequencesIndexBuffer` or `preprocessedBuffer` member of
  `VkGeneratedCommandsInfoNV`

- `VK_BUFFER_USAGE_2_CONDITIONAL_RENDERING_BIT_EXT` specifies that the
  buffer is suitable for passing as the `buffer` parameter to
  [vkCmdBeginConditionalRenderingEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginConditionalRenderingEXT.html).

- `VK_BUFFER_USAGE_2_TRANSFORM_FEEDBACK_BUFFER_BIT_EXT` specifies that
  the buffer is suitable for using for binding as a transform feedback
  buffer with
  [vkCmdBindTransformFeedbackBuffersEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBindTransformFeedbackBuffersEXT.html).

- `VK_BUFFER_USAGE_2_TRANSFORM_FEEDBACK_COUNTER_BUFFER_BIT_EXT`
  specifies that the buffer is suitable for using as a counter buffer
  with
  [vkCmdBeginTransformFeedbackEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginTransformFeedbackEXT.html)
  and [vkCmdEndTransformFeedbackEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdEndTransformFeedbackEXT.html).

- `VK_BUFFER_USAGE_2_SAMPLER_DESCRIPTOR_BUFFER_BIT_EXT` specifies that
  the buffer is suitable to contain sampler and combined image sampler
  descriptors when bound as a descriptor buffer. Buffers containing
  combined image sampler descriptors **must** also specify
  `VK_BUFFER_USAGE_2_RESOURCE_DESCRIPTOR_BUFFER_BIT_EXT`.

- `VK_BUFFER_USAGE_2_RESOURCE_DESCRIPTOR_BUFFER_BIT_EXT` specifies that
  the buffer is suitable to contain resource descriptors when bound as a
  descriptor buffer.

- `VK_BUFFER_USAGE_2_PUSH_DESCRIPTORS_DESCRIPTOR_BUFFER_BIT_EXT`
  specifies that the buffer, when bound, **can** be used by the
  implementation to support push descriptors when using descriptor
  buffers.

- `VK_BUFFER_USAGE_2_RAY_TRACING_BIT_NV` specifies that the buffer is
  suitable for use in [vkCmdTraceRaysNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdTraceRaysNV.html).

- `VK_BUFFER_USAGE_2_SHADER_BINDING_TABLE_BIT_KHR` specifies that the
  buffer is suitable for use as a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#shader-binding-table"
  target="_blank" rel="noopener">Shader Binding Table</a>.

- `VK_BUFFER_USAGE_2_ACCELERATION_STRUCTURE_BUILD_INPUT_READ_ONLY_BIT_KHR`
  specifies that the buffer is suitable for use as a read-only input to
  an <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#acceleration-structure-building"
  target="_blank" rel="noopener">acceleration structure build</a>.

- `VK_BUFFER_USAGE_2_ACCELERATION_STRUCTURE_STORAGE_BIT_KHR` specifies
  that the buffer is suitable for storage space for a
  [VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureKHR.html).

- `VK_BUFFER_USAGE_2_SHADER_DEVICE_ADDRESS_BIT_KHR` specifies that the
  buffer **can** be used to retrieve a buffer device address via
  [vkGetBufferDeviceAddress](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetBufferDeviceAddress.html) and use that
  address to access the buffer’s memory from a shader.

- `VK_BUFFER_USAGE_2_VIDEO_DECODE_SRC_BIT_KHR` specifies that the buffer
  **can** be used as the source video bitstream buffer in a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-decode-operations"
  target="_blank" rel="noopener">video decode operation</a>.

- `VK_BUFFER_USAGE_2_VIDEO_DECODE_DST_BIT_KHR` is reserved for future
  use.

- `VK_BUFFER_USAGE_2_VIDEO_ENCODE_DST_BIT_KHR` specifies that the buffer
  **can** be used as the destination video bitstream buffer in a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-encode-operations"
  target="_blank" rel="noopener">video encode operation</a>.

- `VK_BUFFER_USAGE_2_VIDEO_ENCODE_SRC_BIT_KHR` is reserved for future
  use.

- `VK_BUFFER_USAGE_2_EXECUTION_GRAPH_SCRATCH_BIT_AMDX` specifies that
  the buffer **can** be used for as scratch memory for <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#executiongraphs"
  target="_blank" rel="noopener">execution graph dispatch</a>.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_maintenance5](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_maintenance5.html),
[VkBufferUsageFlags2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferUsageFlags2KHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkBufferUsageFlagBits2KHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
