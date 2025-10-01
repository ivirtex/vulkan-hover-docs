# vkCmdWriteBufferMarker2AMD(3) Manual Page

## Name

vkCmdWriteBufferMarker2AMD - Execute a pipelined write of a marker value into a buffer



## [](#_c_specification)C Specification

To write a 32-bit marker value into a buffer as a pipelined operation, call:

```c++
// Provided by VK_AMD_buffer_marker with VK_VERSION_1_3 or VK_KHR_synchronization2
void vkCmdWriteBufferMarker2AMD(
    VkCommandBuffer                             commandBuffer,
    VkPipelineStageFlags2                       stage,
    VkBuffer                                    dstBuffer,
    VkDeviceSize                                dstOffset,
    uint32_t                                    marker);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `stage` specifies the pipeline stage whose completion triggers the marker write.
- `dstBuffer` is the buffer where the marker will be written.
- `dstOffset` is the byte offset into the buffer where the marker will be written.
- `marker` is the 32-bit value of the marker.

## [](#_description)Description

When `vkCmdWriteBufferMarker2AMD` is submitted to a queue, it defines an execution dependency between prior operations and writing the marker value, as well as a memory dependency from earlier [buffer marker write commands](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#copies-buffer-markers).

The first [synchronization scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-scopes) includes operations performed by operations that occur earlier in [submission order](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-submission-order) in the pipeline stage identified by `pipelineStage`. It additionally includes other [buffer marker write commands](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#copies-buffer-markers) that occur earlier in [submission order](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-submission-order) that specified either the same `pipelineStage` or a stage that is [logically earlier](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-pipeline-stages-order).

The second [synchronization scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-scopes) includes only the buffer marker write.

The first [access scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-access-scopes) includes only accesses performed by other [buffer marker write commands](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#copies-buffer-markers).

The second [access scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-access-scopes) is empty.

The access scope for buffer marker writes falls under the `VK_ACCESS_TRANSFER_WRITE_BIT` flag, and is performed by either `pipelineStage` or `VK_PIPELINE_STAGE_TRANSFER_BIT`. [Synchronization commands](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization) should specify this access flag and both pipeline stages when defining dependencies with this command.

Note

Similar to `vkCmdWriteTimestamp2`, if an implementation is unable to write a marker at any specific pipeline stage, it **may** instead do so at any logically later stage.

Note

Implementations **may** only support a limited number of pipelined marker write operations in flight at a given time. Thus an excessive number of marker write operations **may** degrade command execution performance.

Valid Usage

- [](#VUID-vkCmdWriteBufferMarker2AMD-stage-03929)VUID-vkCmdWriteBufferMarker2AMD-stage-03929  
  If the [`geometryShader`](#features-geometryShader) feature is not enabled, `stage` **must** not contain `VK_PIPELINE_STAGE_2_GEOMETRY_SHADER_BIT`
- [](#VUID-vkCmdWriteBufferMarker2AMD-stage-03930)VUID-vkCmdWriteBufferMarker2AMD-stage-03930  
  If the [`tessellationShader`](#features-tessellationShader) feature is not enabled, `stage` **must** not contain `VK_PIPELINE_STAGE_2_TESSELLATION_CONTROL_SHADER_BIT` or `VK_PIPELINE_STAGE_2_TESSELLATION_EVALUATION_SHADER_BIT`
- [](#VUID-vkCmdWriteBufferMarker2AMD-stage-03931)VUID-vkCmdWriteBufferMarker2AMD-stage-03931  
  If the [`conditionalRendering`](#features-conditionalRendering) feature is not enabled, `stage` **must** not contain `VK_PIPELINE_STAGE_2_CONDITIONAL_RENDERING_BIT_EXT`
- [](#VUID-vkCmdWriteBufferMarker2AMD-stage-03932)VUID-vkCmdWriteBufferMarker2AMD-stage-03932  
  If the [`fragmentDensityMap`](#features-fragmentDensityMap) feature is not enabled, `stage` **must** not contain `VK_PIPELINE_STAGE_2_FRAGMENT_DENSITY_PROCESS_BIT_EXT`
- [](#VUID-vkCmdWriteBufferMarker2AMD-stage-03933)VUID-vkCmdWriteBufferMarker2AMD-stage-03933  
  If the [`transformFeedback`](#features-transformFeedback) feature is not enabled, `stage` **must** not contain `VK_PIPELINE_STAGE_2_TRANSFORM_FEEDBACK_BIT_EXT`
- [](#VUID-vkCmdWriteBufferMarker2AMD-stage-03934)VUID-vkCmdWriteBufferMarker2AMD-stage-03934  
  If the [`meshShader`](#features-meshShader) feature is not enabled, `stage` **must** not contain `VK_PIPELINE_STAGE_2_MESH_SHADER_BIT_EXT`
- [](#VUID-vkCmdWriteBufferMarker2AMD-stage-03935)VUID-vkCmdWriteBufferMarker2AMD-stage-03935  
  If the [`taskShader`](#features-taskShader) feature is not enabled, `stage` **must** not contain `VK_PIPELINE_STAGE_2_TASK_SHADER_BIT_EXT`
- [](#VUID-vkCmdWriteBufferMarker2AMD-stage-07316)VUID-vkCmdWriteBufferMarker2AMD-stage-07316  
  If neither of the [`shadingRateImage`](#features-shadingRateImage) or the [`attachmentFragmentShadingRate`](#features-attachmentFragmentShadingRate) features are enabled, `stage` **must** not contain `VK_PIPELINE_STAGE_2_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR`
- [](#VUID-vkCmdWriteBufferMarker2AMD-stage-04957)VUID-vkCmdWriteBufferMarker2AMD-stage-04957  
  If the [`subpassShading`](#features-subpassShading) feature is not enabled, `stage` **must** not contain `VK_PIPELINE_STAGE_2_SUBPASS_SHADER_BIT_HUAWEI`
- [](#VUID-vkCmdWriteBufferMarker2AMD-stage-04995)VUID-vkCmdWriteBufferMarker2AMD-stage-04995  
  If the [`invocationMask`](#features-invocationMask) feature is not enabled, `stage` **must** not contain `VK_PIPELINE_STAGE_2_INVOCATION_MASK_BIT_HUAWEI`
- [](#VUID-vkCmdWriteBufferMarker2AMD-stage-07946)VUID-vkCmdWriteBufferMarker2AMD-stage-07946  
  If neither the [VK\_NV\_ray\_tracing](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_ray_tracing.html) extension or the [`rayTracingPipeline`](#features-rayTracingPipeline) feature are enabled, `stage` **must** not contain `VK_PIPELINE_STAGE_2_RAY_TRACING_SHADER_BIT_KHR`
- [](#VUID-vkCmdWriteBufferMarker2AMD-stage-10751)VUID-vkCmdWriteBufferMarker2AMD-stage-10751  
  If the [`accelerationStructure`](#features-accelerationStructure) feature is not enabled, `stage` **must** not contain `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR`
- [](#VUID-vkCmdWriteBufferMarker2AMD-stage-10752)VUID-vkCmdWriteBufferMarker2AMD-stage-10752  
  If the [`rayTracingMaintenance1`](#features-rayTracingMaintenance1) feature is not enabled, `stage` **must** not contain `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_COPY_BIT_KHR`
- [](#VUID-vkCmdWriteBufferMarker2AMD-stage-10753)VUID-vkCmdWriteBufferMarker2AMD-stage-10753  
  If the [`micromap`](#features-micromap) feature is not enabled, `stage` **must** not contain `VK_PIPELINE_STAGE_2_MICROMAP_BUILD_BIT_EXT`
- [](#VUID-vkCmdWriteBufferMarker2AMD-synchronization2-03893)VUID-vkCmdWriteBufferMarker2AMD-synchronization2-03893  
  The [`synchronization2`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-synchronization2) feature **must** be enabled
- [](#VUID-vkCmdWriteBufferMarker2AMD-stage-03894)VUID-vkCmdWriteBufferMarker2AMD-stage-03894  
  `stage` **must** include only a single pipeline stage
- [](#VUID-vkCmdWriteBufferMarker2AMD-stage-03895)VUID-vkCmdWriteBufferMarker2AMD-stage-03895  
  `stage` **must** include only stages that are valid for the queue family that was used to create the command pool that `commandBuffer` was allocated from
- [](#VUID-vkCmdWriteBufferMarker2AMD-dstOffset-03896)VUID-vkCmdWriteBufferMarker2AMD-dstOffset-03896  
  `dstOffset` **must** be less than or equal to the size of `dstBuffer` minus `4`
- [](#VUID-vkCmdWriteBufferMarker2AMD-dstBuffer-03897)VUID-vkCmdWriteBufferMarker2AMD-dstBuffer-03897  
  `dstBuffer` **must** have been created with the `VK_BUFFER_USAGE_TRANSFER_DST_BIT` usage flag
- [](#VUID-vkCmdWriteBufferMarker2AMD-dstBuffer-03898)VUID-vkCmdWriteBufferMarker2AMD-dstBuffer-03898  
  If `dstBuffer` is non-sparse then it **must** be bound completely and contiguously to a single `VkDeviceMemory` object
- [](#VUID-vkCmdWriteBufferMarker2AMD-dstOffset-03899)VUID-vkCmdWriteBufferMarker2AMD-dstOffset-03899  
  `dstOffset` **must** be a multiple of `4`

Valid Usage (Implicit)

- [](#VUID-vkCmdWriteBufferMarker2AMD-commandBuffer-parameter)VUID-vkCmdWriteBufferMarker2AMD-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdWriteBufferMarker2AMD-stage-parameter)VUID-vkCmdWriteBufferMarker2AMD-stage-parameter  
  `stage` **must** be a valid combination of [VkPipelineStageFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlagBits2.html) values
- [](#VUID-vkCmdWriteBufferMarker2AMD-dstBuffer-parameter)VUID-vkCmdWriteBufferMarker2AMD-dstBuffer-parameter  
  `dstBuffer` **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) handle
- [](#VUID-vkCmdWriteBufferMarker2AMD-commandBuffer-recording)VUID-vkCmdWriteBufferMarker2AMD-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdWriteBufferMarker2AMD-commandBuffer-cmdpool)VUID-vkCmdWriteBufferMarker2AMD-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support VK\_QUEUE\_COMPUTE\_BIT, VK\_QUEUE\_GRAPHICS\_BIT, or VK\_QUEUE\_TRANSFER\_BIT operations
- [](#VUID-vkCmdWriteBufferMarker2AMD-videocoding)VUID-vkCmdWriteBufferMarker2AMD-videocoding  
  This command **must** only be called outside of a video coding scope
- [](#VUID-vkCmdWriteBufferMarker2AMD-commonparent)VUID-vkCmdWriteBufferMarker2AMD-commonparent  
  Both of `commandBuffer`, and `dstBuffer` **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

Host Synchronization

- Host access to `commandBuffer` **must** be externally synchronized
- Host access to the `VkCommandPool` that `commandBuffer` was allocated from **must** be externally synchronized

Command Properties

     [Command Buffer Levels](#VkCommandBufferLevel) [Render Pass Scope](#vkCmdBeginRenderPass) [Video Coding Scope](#vkCmdBeginVideoCodingKHR) [Supported Queue Types](#VkQueueFlagBits) [Command Type](#fundamentals-queueoperation-command-types)

Primary  
Secondary

Both

Outside

VK\_QUEUE\_COMPUTE\_BIT  
VK\_QUEUE\_GRAPHICS\_BIT  
VK\_QUEUE\_TRANSFER\_BIT

Action

Conditional Rendering

vkCmdWriteBufferMarker2AMD is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_AMD\_buffer\_marker](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_AMD_buffer_marker.html), [VK\_KHR\_synchronization2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_synchronization2.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html), [VkPipelineStageFlags2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlags2.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdWriteBufferMarker2AMD).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0