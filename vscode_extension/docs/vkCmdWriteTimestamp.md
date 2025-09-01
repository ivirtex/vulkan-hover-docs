# vkCmdWriteTimestamp(3) Manual Page

## Name

vkCmdWriteTimestamp - Write a device timestamp into a query object



## [](#_c_specification)C Specification

To request a timestamp and write the value to memory, call:

```c++
// Provided by VK_VERSION_1_0
void vkCmdWriteTimestamp(
    VkCommandBuffer                             commandBuffer,
    VkPipelineStageFlagBits                     pipelineStage,
    VkQueryPool                                 queryPool,
    uint32_t                                    query);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `pipelineStage` is a [VkPipelineStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlagBits.html) value, specifying a stage of the pipeline.
- `queryPool` is the query pool that will manage the timestamp.
- `query` is the query within the query pool that will contain the timestamp.

## [](#_description)Description

When `vkCmdWriteTimestamp` is submitted to a queue, it defines an execution dependency on commands that were submitted before it, and writes a timestamp to a query pool.

The first [synchronization scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-scopes) includes all commands that occur earlier in [submission order](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-submission-order). The synchronization scope is limited to operations on the pipeline stage specified by `pipelineStage`.

The second [synchronization scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-scopes) includes only the timestamp write operation.

Note

Implementations may write the timestamp at any stage that is [logically later](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-pipeline-stages-order) than `stage`.

Any timestamp write that [happens-after](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-execution) another timestamp write in the same submission **must** not have a lower value unless its value overflows the maximum supported integer bit width of the query. If `VK_KHR_calibrated_timestamps` or `VK_EXT_calibrated_timestamps` is enabled, this extends to timestamp writes across all submissions on the same logical device: any timestamp write that [happens-after](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-execution) another **must** not have a lower value unless its value overflows the maximum supported integer bit width of the query. Timestamps written by this command **must** be in the `VK_TIME_DOMAIN_DEVICE_KHR` [time domain](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTimeDomainKHR.html). If an overflow occurs, the timestamp value **must** wrap back to zero.

If `vkCmdWriteTimestamp` is called while executing a render pass instance that has multiview enabled, the timestamp uses N consecutive query indices in the query pool (starting at `query`) where N is the number of bits set in the view mask of the subpass or dynamic render pass the command is executed in. The resulting query values are determined by an implementation-dependent choice of one of the following behaviors:

- The first query is a timestamp value and (if more than one bit is set in the view mask) zero is written to the remaining queries.
- All N queries are timestamp values.

Either way, if two timestamps are written in the same subpass or dynamic render pass with multiview enabled, each of the N consecutive queries written for a timestamp **must** not have a lower value than the queries with corresponding indices written by the timestamp that happens-before unless the value overflows the maximum supported integer bit width of the query.

Valid Usage

- [](#VUID-vkCmdWriteTimestamp-pipelineStage-04074)VUID-vkCmdWriteTimestamp-pipelineStage-04074  
  `pipelineStage` **must** be a [valid stage](#synchronization-pipeline-stages-supported) for the queue family that was used to create the command pool that `commandBuffer` was allocated from
- [](#VUID-vkCmdWriteTimestamp-pipelineStage-04075)VUID-vkCmdWriteTimestamp-pipelineStage-04075  
  If the [`geometryShader`](#features-geometryShader) feature is not enabled, `pipelineStage` **must** not be `VK_PIPELINE_STAGE_GEOMETRY_SHADER_BIT`
- [](#VUID-vkCmdWriteTimestamp-pipelineStage-04076)VUID-vkCmdWriteTimestamp-pipelineStage-04076  
  If the [`tessellationShader`](#features-tessellationShader) feature is not enabled, `pipelineStage` **must** not be `VK_PIPELINE_STAGE_TESSELLATION_CONTROL_SHADER_BIT` or `VK_PIPELINE_STAGE_TESSELLATION_EVALUATION_SHADER_BIT`
- [](#VUID-vkCmdWriteTimestamp-pipelineStage-04077)VUID-vkCmdWriteTimestamp-pipelineStage-04077  
  If the [`conditionalRendering`](#features-conditionalRendering) feature is not enabled, `pipelineStage` **must** not be `VK_PIPELINE_STAGE_CONDITIONAL_RENDERING_BIT_EXT`
- [](#VUID-vkCmdWriteTimestamp-pipelineStage-04078)VUID-vkCmdWriteTimestamp-pipelineStage-04078  
  If the [`fragmentDensityMap`](#features-fragmentDensityMap) feature is not enabled, `pipelineStage` **must** not be `VK_PIPELINE_STAGE_FRAGMENT_DENSITY_PROCESS_BIT_EXT`
- [](#VUID-vkCmdWriteTimestamp-pipelineStage-04079)VUID-vkCmdWriteTimestamp-pipelineStage-04079  
  If the [`transformFeedback`](#features-transformFeedback) feature is not enabled, `pipelineStage` **must** not be `VK_PIPELINE_STAGE_TRANSFORM_FEEDBACK_BIT_EXT`
- [](#VUID-vkCmdWriteTimestamp-pipelineStage-04080)VUID-vkCmdWriteTimestamp-pipelineStage-04080  
  If the [`meshShader`](#features-meshShader) feature is not enabled, `pipelineStage` **must** not be `VK_PIPELINE_STAGE_MESH_SHADER_BIT_EXT`
- [](#VUID-vkCmdWriteTimestamp-pipelineStage-07077)VUID-vkCmdWriteTimestamp-pipelineStage-07077  
  If the [`taskShader`](#features-taskShader) feature is not enabled, `pipelineStage` **must** not be `VK_PIPELINE_STAGE_TASK_SHADER_BIT_EXT`
- [](#VUID-vkCmdWriteTimestamp-shadingRateImage-07314)VUID-vkCmdWriteTimestamp-shadingRateImage-07314  
  If neither of the [`shadingRateImage`](#features-shadingRateImage) or the [`attachmentFragmentShadingRate`](#features-attachmentFragmentShadingRate) features are enabled, `pipelineStage` **must** not be `VK_PIPELINE_STAGE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR`
- [](#VUID-vkCmdWriteTimestamp-synchronization2-06489)VUID-vkCmdWriteTimestamp-synchronization2-06489  
  If the [`synchronization2`](#features-synchronization2) feature is not enabled, `pipelineStage` **must** not be `VK_PIPELINE_STAGE_NONE`
- [](#VUID-vkCmdWriteTimestamp-rayTracingPipeline-07943)VUID-vkCmdWriteTimestamp-rayTracingPipeline-07943  
  If neither of the [VK\_NV\_ray\_tracing](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_ray_tracing.html) extension or the [`rayTracingPipeline`](#features-rayTracingPipeline) feature are enabled, `pipelineStage` **must** not be `VK_PIPELINE_STAGE_RAY_TRACING_SHADER_BIT_KHR`
- [](#VUID-vkCmdWriteTimestamp-queryPool-01416)VUID-vkCmdWriteTimestamp-queryPool-01416  
  `queryPool` **must** have been created with a `queryType` of `VK_QUERY_TYPE_TIMESTAMP`
- [](#VUID-vkCmdWriteTimestamp-timestampValidBits-00829)VUID-vkCmdWriteTimestamp-timestampValidBits-00829  
  The command pool’s queue family **must** support a non-zero `timestampValidBits`
- [](#VUID-vkCmdWriteTimestamp-query-04904)VUID-vkCmdWriteTimestamp-query-04904  
  `query` **must** be less than the number of queries in `queryPool`
- [](#VUID-vkCmdWriteTimestamp-None-00830)VUID-vkCmdWriteTimestamp-None-00830  
  All queries used by the command **must** be *unavailable*
- [](#VUID-vkCmdWriteTimestamp-query-00831)VUID-vkCmdWriteTimestamp-query-00831  
  If `vkCmdWriteTimestamp` is called within a render pass instance, the sum of `query` and the number of bits set in the current subpass’s view mask **must** be less than or equal to the number of queries in `queryPool`
- [](#VUID-vkCmdWriteTimestamp-None-10640)VUID-vkCmdWriteTimestamp-None-10640  
  This command **must** not be recorded when [per-tile execution model](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-per-tile-execution-model) is enabled

Valid Usage (Implicit)

- [](#VUID-vkCmdWriteTimestamp-commandBuffer-parameter)VUID-vkCmdWriteTimestamp-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdWriteTimestamp-pipelineStage-parameter)VUID-vkCmdWriteTimestamp-pipelineStage-parameter  
  `pipelineStage` **must** be a valid [VkPipelineStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlagBits.html) value
- [](#VUID-vkCmdWriteTimestamp-queryPool-parameter)VUID-vkCmdWriteTimestamp-queryPool-parameter  
  `queryPool` **must** be a valid [VkQueryPool](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPool.html) handle
- [](#VUID-vkCmdWriteTimestamp-commandBuffer-recording)VUID-vkCmdWriteTimestamp-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdWriteTimestamp-commandBuffer-cmdpool)VUID-vkCmdWriteTimestamp-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support transfer, graphics, compute, decode, encode, or optical flow operations
- [](#VUID-vkCmdWriteTimestamp-commonparent)VUID-vkCmdWriteTimestamp-commonparent  
  Both of `commandBuffer`, and `queryPool` **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

Host Synchronization

- Host access to `commandBuffer` **must** be externally synchronized
- Host access to the `VkCommandPool` that `commandBuffer` was allocated from **must** be externally synchronized

Command Properties

     [Command Buffer Levels](#VkCommandBufferLevel) [Render Pass Scope](#vkCmdBeginRenderPass) [Video Coding Scope](#vkCmdBeginVideoCodingKHR) [Supported Queue Types](#VkQueueFlagBits) [Command Type](#fundamentals-queueoperation-command-types)

Primary  
Secondary

Both

Both

Transfer  
Graphics  
Compute  
Decode  
Encode  
Opticalflow

Action

Conditional Rendering

vkCmdWriteTimestamp is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkPipelineStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlagBits.html), [VkQueryPool](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPool.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdWriteTimestamp).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0