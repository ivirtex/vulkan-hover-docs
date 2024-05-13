# VkBufferUsageFlagBits(3) Manual Page

## Name

VkBufferUsageFlagBits - Bitmask specifying allowed usage of a buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **can** be set in
[VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCreateInfo.html)::`usage`, specifying usage
behavior of a buffer, are:

``` c
// Provided by VK_VERSION_1_0
typedef enum VkBufferUsageFlagBits {
    VK_BUFFER_USAGE_TRANSFER_SRC_BIT = 0x00000001,
    VK_BUFFER_USAGE_TRANSFER_DST_BIT = 0x00000002,
    VK_BUFFER_USAGE_UNIFORM_TEXEL_BUFFER_BIT = 0x00000004,
    VK_BUFFER_USAGE_STORAGE_TEXEL_BUFFER_BIT = 0x00000008,
    VK_BUFFER_USAGE_UNIFORM_BUFFER_BIT = 0x00000010,
    VK_BUFFER_USAGE_STORAGE_BUFFER_BIT = 0x00000020,
    VK_BUFFER_USAGE_INDEX_BUFFER_BIT = 0x00000040,
    VK_BUFFER_USAGE_VERTEX_BUFFER_BIT = 0x00000080,
    VK_BUFFER_USAGE_INDIRECT_BUFFER_BIT = 0x00000100,
  // Provided by VK_VERSION_1_2
    VK_BUFFER_USAGE_SHADER_DEVICE_ADDRESS_BIT = 0x00020000,
  // Provided by VK_KHR_video_decode_queue
    VK_BUFFER_USAGE_VIDEO_DECODE_SRC_BIT_KHR = 0x00002000,
  // Provided by VK_KHR_video_decode_queue
    VK_BUFFER_USAGE_VIDEO_DECODE_DST_BIT_KHR = 0x00004000,
  // Provided by VK_EXT_transform_feedback
    VK_BUFFER_USAGE_TRANSFORM_FEEDBACK_BUFFER_BIT_EXT = 0x00000800,
  // Provided by VK_EXT_transform_feedback
    VK_BUFFER_USAGE_TRANSFORM_FEEDBACK_COUNTER_BUFFER_BIT_EXT = 0x00001000,
  // Provided by VK_EXT_conditional_rendering
    VK_BUFFER_USAGE_CONDITIONAL_RENDERING_BIT_EXT = 0x00000200,
#ifdef VK_ENABLE_BETA_EXTENSIONS
  // Provided by VK_AMDX_shader_enqueue
    VK_BUFFER_USAGE_EXECUTION_GRAPH_SCRATCH_BIT_AMDX = 0x02000000,
#endif
  // Provided by VK_KHR_acceleration_structure
    VK_BUFFER_USAGE_ACCELERATION_STRUCTURE_BUILD_INPUT_READ_ONLY_BIT_KHR = 0x00080000,
  // Provided by VK_KHR_acceleration_structure
    VK_BUFFER_USAGE_ACCELERATION_STRUCTURE_STORAGE_BIT_KHR = 0x00100000,
  // Provided by VK_KHR_ray_tracing_pipeline
    VK_BUFFER_USAGE_SHADER_BINDING_TABLE_BIT_KHR = 0x00000400,
  // Provided by VK_KHR_video_encode_queue
    VK_BUFFER_USAGE_VIDEO_ENCODE_DST_BIT_KHR = 0x00008000,
  // Provided by VK_KHR_video_encode_queue
    VK_BUFFER_USAGE_VIDEO_ENCODE_SRC_BIT_KHR = 0x00010000,
  // Provided by VK_EXT_descriptor_buffer
    VK_BUFFER_USAGE_SAMPLER_DESCRIPTOR_BUFFER_BIT_EXT = 0x00200000,
  // Provided by VK_EXT_descriptor_buffer
    VK_BUFFER_USAGE_RESOURCE_DESCRIPTOR_BUFFER_BIT_EXT = 0x00400000,
  // Provided by VK_EXT_descriptor_buffer
    VK_BUFFER_USAGE_PUSH_DESCRIPTORS_DESCRIPTOR_BUFFER_BIT_EXT = 0x04000000,
  // Provided by VK_EXT_opacity_micromap
    VK_BUFFER_USAGE_MICROMAP_BUILD_INPUT_READ_ONLY_BIT_EXT = 0x00800000,
  // Provided by VK_EXT_opacity_micromap
    VK_BUFFER_USAGE_MICROMAP_STORAGE_BIT_EXT = 0x01000000,
  // Provided by VK_NV_ray_tracing
    VK_BUFFER_USAGE_RAY_TRACING_BIT_NV = VK_BUFFER_USAGE_SHADER_BINDING_TABLE_BIT_KHR,
  // Provided by VK_EXT_buffer_device_address
    VK_BUFFER_USAGE_SHADER_DEVICE_ADDRESS_BIT_EXT = VK_BUFFER_USAGE_SHADER_DEVICE_ADDRESS_BIT,
  // Provided by VK_KHR_buffer_device_address
    VK_BUFFER_USAGE_SHADER_DEVICE_ADDRESS_BIT_KHR = VK_BUFFER_USAGE_SHADER_DEVICE_ADDRESS_BIT,
} VkBufferUsageFlagBits;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_BUFFER_USAGE_TRANSFER_SRC_BIT` specifies that the buffer **can**
  be used as the source of a *transfer command* (see the definition of
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-pipeline-stages-transfer"
  target="_blank"
  rel="noopener"><code>VK_PIPELINE_STAGE_TRANSFER_BIT</code></a>).

- `VK_BUFFER_USAGE_TRANSFER_DST_BIT` specifies that the buffer **can**
  be used as the destination of a transfer command.

- `VK_BUFFER_USAGE_UNIFORM_TEXEL_BUFFER_BIT` specifies that the buffer
  **can** be used to create a `VkBufferView` suitable for occupying a
  `VkDescriptorSet` slot of type
  `VK_DESCRIPTOR_TYPE_UNIFORM_TEXEL_BUFFER`.

- `VK_BUFFER_USAGE_STORAGE_TEXEL_BUFFER_BIT` specifies that the buffer
  **can** be used to create a `VkBufferView` suitable for occupying a
  `VkDescriptorSet` slot of type
  `VK_DESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFER`.

- `VK_BUFFER_USAGE_UNIFORM_BUFFER_BIT` specifies that the buffer **can**
  be used in a `VkDescriptorBufferInfo` suitable for occupying a
  `VkDescriptorSet` slot either of type
  `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER` or
  `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC`.

- `VK_BUFFER_USAGE_STORAGE_BUFFER_BIT` specifies that the buffer **can**
  be used in a `VkDescriptorBufferInfo` suitable for occupying a
  `VkDescriptorSet` slot either of type
  `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER` or
  `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC`.

- `VK_BUFFER_USAGE_INDEX_BUFFER_BIT` specifies that the buffer is
  suitable for passing as the `buffer` parameter to
  [vkCmdBindIndexBuffer2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBindIndexBuffer2KHR.html) and
  [vkCmdBindIndexBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBindIndexBuffer.html).

- `VK_BUFFER_USAGE_VERTEX_BUFFER_BIT` specifies that the buffer is
  suitable for passing as an element of the `pBuffers` array to
  [vkCmdBindVertexBuffers](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBindVertexBuffers.html).

- `VK_BUFFER_USAGE_INDIRECT_BUFFER_BIT` specifies that the buffer is
  suitable for passing as the `buffer` parameter to
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

- `VK_BUFFER_USAGE_CONDITIONAL_RENDERING_BIT_EXT` specifies that the
  buffer is suitable for passing as the `buffer` parameter to
  [vkCmdBeginConditionalRenderingEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginConditionalRenderingEXT.html).

- `VK_BUFFER_USAGE_TRANSFORM_FEEDBACK_BUFFER_BIT_EXT` specifies that the
  buffer is suitable for using for binding as a transform feedback
  buffer with
  [vkCmdBindTransformFeedbackBuffersEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBindTransformFeedbackBuffersEXT.html).

- `VK_BUFFER_USAGE_TRANSFORM_FEEDBACK_COUNTER_BUFFER_BIT_EXT` specifies
  that the buffer is suitable for using as a counter buffer with
  [vkCmdBeginTransformFeedbackEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginTransformFeedbackEXT.html)
  and [vkCmdEndTransformFeedbackEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdEndTransformFeedbackEXT.html).

- `VK_BUFFER_USAGE_SAMPLER_DESCRIPTOR_BUFFER_BIT_EXT` specifies that the
  buffer is suitable to contain sampler and combined image sampler
  descriptors when bound as a descriptor buffer. Buffers containing
  combined image sampler descriptors **must** also specify
  `VK_BUFFER_USAGE_RESOURCE_DESCRIPTOR_BUFFER_BIT_EXT`.

- `VK_BUFFER_USAGE_RESOURCE_DESCRIPTOR_BUFFER_BIT_EXT` specifies that
  the buffer is suitable to contain resource descriptors when bound as a
  descriptor buffer.

- `VK_BUFFER_USAGE_PUSH_DESCRIPTORS_DESCRIPTOR_BUFFER_BIT_EXT` specifies
  that the buffer, when bound, **can** be used by the implementation to
  support push descriptors when using descriptor buffers.

- `VK_BUFFER_USAGE_RAY_TRACING_BIT_NV` specifies that the buffer is
  suitable for use in [vkCmdTraceRaysNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdTraceRaysNV.html).

- `VK_BUFFER_USAGE_SHADER_BINDING_TABLE_BIT_KHR` specifies that the
  buffer is suitable for use as a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#shader-binding-table"
  target="_blank" rel="noopener">Shader Binding Table</a>.

- `VK_BUFFER_USAGE_ACCELERATION_STRUCTURE_BUILD_INPUT_READ_ONLY_BIT_KHR`
  specifies that the buffer is suitable for use as a read-only input to
  an <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#acceleration-structure-building"
  target="_blank" rel="noopener">acceleration structure build</a>.

- `VK_BUFFER_USAGE_ACCELERATION_STRUCTURE_STORAGE_BIT_KHR` specifies
  that the buffer is suitable for storage space for a
  [VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureKHR.html).

- `VK_BUFFER_USAGE_SHADER_DEVICE_ADDRESS_BIT` specifies that the buffer
  **can** be used to retrieve a buffer device address via
  [vkGetBufferDeviceAddress](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetBufferDeviceAddress.html) and use that
  address to access the buffer’s memory from a shader.

- `VK_BUFFER_USAGE_VIDEO_DECODE_SRC_BIT_KHR` specifies that the buffer
  **can** be used as the source video bitstream buffer in a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-decode-operations"
  target="_blank" rel="noopener">video decode operation</a>.

- `VK_BUFFER_USAGE_VIDEO_DECODE_DST_BIT_KHR` is reserved for future use.

- `VK_BUFFER_USAGE_VIDEO_ENCODE_DST_BIT_KHR` specifies that the buffer
  **can** be used as the destination video bitstream buffer in a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-encode-operations"
  target="_blank" rel="noopener">video encode operation</a>.

- `VK_BUFFER_USAGE_VIDEO_ENCODE_SRC_BIT_KHR` is reserved for future use.

- `VK_BUFFER_USAGE_EXECUTION_GRAPH_SCRATCH_BIT_AMDX` specifies that the
  buffer **can** be used for as scratch memory for <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#executiongraphs"
  target="_blank" rel="noopener">execution graph dispatch</a>.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkBufferUsageFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferUsageFlags.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkBufferUsageFlagBits"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
