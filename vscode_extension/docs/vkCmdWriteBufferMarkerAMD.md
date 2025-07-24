# vkCmdWriteBufferMarkerAMD(3) Manual Page

## Name

vkCmdWriteBufferMarkerAMD - Execute a pipelined write of a marker value into a buffer



## [](#_c_specification)C Specification

To write a 32-bit marker value into a buffer as a pipelined operation, call:

```c++
// Provided by VK_AMD_buffer_marker
void vkCmdWriteBufferMarkerAMD(
    VkCommandBuffer                             commandBuffer,
    VkPipelineStageFlagBits                     pipelineStage,
    VkBuffer                                    dstBuffer,
    VkDeviceSize                                dstOffset,
    uint32_t                                    marker);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `pipelineStage` is a [VkPipelineStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlagBits.html) value specifying the pipeline stage whose completion triggers the marker write.
- `dstBuffer` is the buffer where the marker will be written to.
- `dstOffset` is the byte offset into the buffer where the marker will be written to.
- `marker` is the 32-bit value of the marker.

## [](#_description)Description

The command will write the 32-bit marker value into the buffer only after all preceding commands have finished executing up to at least the specified pipeline stage. This includes the completion of other preceding `vkCmdWriteBufferMarkerAMD` commands so long as their specified pipeline stages occur either at the same time or earlier than this command’s specified `pipelineStage`.

While consecutive buffer marker writes with the same `pipelineStage` parameter are implicitly complete in submission order, memory and execution dependencies between buffer marker writes and other operations **must** still be explicitly ordered using synchronization commands. The access scope for buffer marker writes falls under the `VK_ACCESS_TRANSFER_WRITE_BIT`, and the pipeline stages for identifying the synchronization scope **must** include both `pipelineStage` and `VK_PIPELINE_STAGE_TRANSFER_BIT`.

Note

Similar to `vkCmdWriteTimestamp`, if an implementation is unable to write a marker at any specific pipeline stage, it **may** instead do so at any logically later stage.

Note

Implementations **may** only support a limited number of pipelined marker write operations in flight at a given time, thus excessive number of marker write operations **may** degrade command execution performance.

Valid Usage

- [](#VUID-vkCmdWriteBufferMarkerAMD-pipelineStage-04074)VUID-vkCmdWriteBufferMarkerAMD-pipelineStage-04074  
  `pipelineStage` **must** be a [valid stage](#synchronization-pipeline-stages-supported) for the queue family that was used to create the command pool that `commandBuffer` was allocated from
- [](#VUID-vkCmdWriteBufferMarkerAMD-pipelineStage-04075)VUID-vkCmdWriteBufferMarkerAMD-pipelineStage-04075  
  If the [`geometryShader`](#features-geometryShader) feature is not enabled, `pipelineStage` **must** not be `VK_PIPELINE_STAGE_GEOMETRY_SHADER_BIT`
- [](#VUID-vkCmdWriteBufferMarkerAMD-pipelineStage-04076)VUID-vkCmdWriteBufferMarkerAMD-pipelineStage-04076  
  If the [`tessellationShader`](#features-tessellationShader) feature is not enabled, `pipelineStage` **must** not be `VK_PIPELINE_STAGE_TESSELLATION_CONTROL_SHADER_BIT` or `VK_PIPELINE_STAGE_TESSELLATION_EVALUATION_SHADER_BIT`
- [](#VUID-vkCmdWriteBufferMarkerAMD-pipelineStage-04077)VUID-vkCmdWriteBufferMarkerAMD-pipelineStage-04077  
  If the [`conditionalRendering`](#features-conditionalRendering) feature is not enabled, `pipelineStage` **must** not be `VK_PIPELINE_STAGE_CONDITIONAL_RENDERING_BIT_EXT`
- [](#VUID-vkCmdWriteBufferMarkerAMD-pipelineStage-04078)VUID-vkCmdWriteBufferMarkerAMD-pipelineStage-04078  
  If the [`fragmentDensityMap`](#features-fragmentDensityMap) feature is not enabled, `pipelineStage` **must** not be `VK_PIPELINE_STAGE_FRAGMENT_DENSITY_PROCESS_BIT_EXT`
- [](#VUID-vkCmdWriteBufferMarkerAMD-pipelineStage-04079)VUID-vkCmdWriteBufferMarkerAMD-pipelineStage-04079  
  If the [`transformFeedback`](#features-transformFeedback) feature is not enabled, `pipelineStage` **must** not be `VK_PIPELINE_STAGE_TRANSFORM_FEEDBACK_BIT_EXT`
- [](#VUID-vkCmdWriteBufferMarkerAMD-pipelineStage-04080)VUID-vkCmdWriteBufferMarkerAMD-pipelineStage-04080  
  If the [`meshShader`](#features-meshShader) feature is not enabled, `pipelineStage` **must** not be `VK_PIPELINE_STAGE_MESH_SHADER_BIT_EXT`
- [](#VUID-vkCmdWriteBufferMarkerAMD-pipelineStage-07077)VUID-vkCmdWriteBufferMarkerAMD-pipelineStage-07077  
  If the [`taskShader`](#features-taskShader) feature is not enabled, `pipelineStage` **must** not be `VK_PIPELINE_STAGE_TASK_SHADER_BIT_EXT`
- [](#VUID-vkCmdWriteBufferMarkerAMD-shadingRateImage-07314)VUID-vkCmdWriteBufferMarkerAMD-shadingRateImage-07314  
  If neither of the [`shadingRateImage`](#features-shadingRateImage) or the [`attachmentFragmentShadingRate`](#features-attachmentFragmentShadingRate) features are enabled, `pipelineStage` **must** not be `VK_PIPELINE_STAGE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR`
- [](#VUID-vkCmdWriteBufferMarkerAMD-synchronization2-06489)VUID-vkCmdWriteBufferMarkerAMD-synchronization2-06489  
  If the [`synchronization2`](#features-synchronization2) feature is not enabled, `pipelineStage` **must** not be `VK_PIPELINE_STAGE_NONE`
- [](#VUID-vkCmdWriteBufferMarkerAMD-rayTracingPipeline-07943)VUID-vkCmdWriteBufferMarkerAMD-rayTracingPipeline-07943  
  If neither of the [VK\_NV\_ray\_tracing](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_ray_tracing.html) extension or the [`rayTracingPipeline`](#features-rayTracingPipeline) feature are enabled, `pipelineStage` **must** not be `VK_PIPELINE_STAGE_RAY_TRACING_SHADER_BIT_KHR`
- [](#VUID-vkCmdWriteBufferMarkerAMD-dstOffset-01798)VUID-vkCmdWriteBufferMarkerAMD-dstOffset-01798  
  `dstOffset` **must** be less than or equal to the size of `dstBuffer` minus `4`
- [](#VUID-vkCmdWriteBufferMarkerAMD-dstBuffer-01799)VUID-vkCmdWriteBufferMarkerAMD-dstBuffer-01799  
  `dstBuffer` **must** have been created with `VK_BUFFER_USAGE_TRANSFER_DST_BIT` usage flag
- [](#VUID-vkCmdWriteBufferMarkerAMD-dstBuffer-01800)VUID-vkCmdWriteBufferMarkerAMD-dstBuffer-01800  
  If `dstBuffer` is non-sparse then it **must** be bound completely and contiguously to a single `VkDeviceMemory` object
- [](#VUID-vkCmdWriteBufferMarkerAMD-dstOffset-01801)VUID-vkCmdWriteBufferMarkerAMD-dstOffset-01801  
  `dstOffset` **must** be a multiple of `4`

Valid Usage (Implicit)

- [](#VUID-vkCmdWriteBufferMarkerAMD-commandBuffer-parameter)VUID-vkCmdWriteBufferMarkerAMD-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdWriteBufferMarkerAMD-pipelineStage-parameter)VUID-vkCmdWriteBufferMarkerAMD-pipelineStage-parameter  
  If `pipelineStage` is not `0`, `pipelineStage` **must** be a valid [VkPipelineStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlagBits.html) value
- [](#VUID-vkCmdWriteBufferMarkerAMD-dstBuffer-parameter)VUID-vkCmdWriteBufferMarkerAMD-dstBuffer-parameter  
  `dstBuffer` **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) handle
- [](#VUID-vkCmdWriteBufferMarkerAMD-commandBuffer-recording)VUID-vkCmdWriteBufferMarkerAMD-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdWriteBufferMarkerAMD-commandBuffer-cmdpool)VUID-vkCmdWriteBufferMarkerAMD-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support transfer, graphics, or compute operations
- [](#VUID-vkCmdWriteBufferMarkerAMD-videocoding)VUID-vkCmdWriteBufferMarkerAMD-videocoding  
  This command **must** only be called outside of a video coding scope
- [](#VUID-vkCmdWriteBufferMarkerAMD-commonparent)VUID-vkCmdWriteBufferMarkerAMD-commonparent  
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

Transfer  
Graphics  
Compute

Action

Conditional Rendering

vkCmdWriteBufferMarkerAMD is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_AMD\_buffer\_marker](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_AMD_buffer_marker.html), [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html), [VkPipelineStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlagBits.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdWriteBufferMarkerAMD)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0