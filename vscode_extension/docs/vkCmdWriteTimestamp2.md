# vkCmdWriteTimestamp2(3) Manual Page

## Name

vkCmdWriteTimestamp2 - Write a device timestamp into a query object



## [](#_c_specification)C Specification

To request a timestamp and write the value to memory, call:

```c++
// Provided by VK_VERSION_1_3
void vkCmdWriteTimestamp2(
    VkCommandBuffer                             commandBuffer,
    VkPipelineStageFlags2                       stage,
    VkQueryPool                                 queryPool,
    uint32_t                                    query);
```

or the equivalent command

```c++
// Provided by VK_KHR_synchronization2
void vkCmdWriteTimestamp2KHR(
    VkCommandBuffer                             commandBuffer,
    VkPipelineStageFlags2                       stage,
    VkQueryPool                                 queryPool,
    uint32_t                                    query);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `stage` specifies a stage of the pipeline.
- `queryPool` is the query pool that will manage the timestamp.
- `query` is the query within the query pool that will contain the timestamp.

## [](#_description)Description

When `vkCmdWriteTimestamp2` is submitted to a queue, it defines an execution dependency on commands that were submitted before it, and writes a timestamp to a query pool.

The first [synchronization scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-scopes) includes all commands that occur earlier in [submission order](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-submission-order). The synchronization scope is limited to operations on the pipeline stage specified by `stage`.

The second [synchronization scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-scopes) includes only the timestamp write operation.

Note

Implementations may write the timestamp at any stage that is [logically later](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-pipeline-stages-order) than `stage`.

Any timestamp write that [happens-after](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-execution) another timestamp write in the same submission **must** not have a lower value unless its value overflows the maximum supported integer bit width of the query. If `VK_KHR_calibrated_timestamps` or `VK_EXT_calibrated_timestamps` is enabled, this extends to timestamp writes across all submissions on the same logical device: any timestamp write that [happens-after](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-execution) another **must** not have a lower value unless its value overflows the maximum supported integer bit width of the query. Timestamps written by this command **must** be in the `VK_TIME_DOMAIN_DEVICE_KHR` [time domain](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTimeDomainKHR.html). If an overflow occurs, the timestamp value **must** wrap back to zero.

If `vkCmdWriteTimestamp2` is called while executing a render pass instance that has multiview enabled, the timestamp uses N consecutive query indices in the query pool (starting at `query`) where N is the number of bits set in the view mask of the subpass or dynamic render pass the command is executed in. The resulting query values are determined by an implementation-dependent choice of one of the following behaviors:

- The first query is a timestamp value and (if more than one bit is set in the view mask) zero is written to the remaining queries.
- All N queries are timestamp values.

Either way, if two timestamps are written in the same subpass or dynamic render pass with multiview enabled, each of the N consecutive queries written for a timestamp **must** not have a lower value than the queries with corresponding indices written by the timestamp that happens-before unless the value overflows the maximum supported integer bit width of the query.

Valid Usage

- [](#VUID-vkCmdWriteTimestamp2-stage-03929)VUID-vkCmdWriteTimestamp2-stage-03929  
  If the [`geometryShader`](#features-geometryShader) feature is not enabled, `stage` **must** not contain `VK_PIPELINE_STAGE_2_GEOMETRY_SHADER_BIT`
- [](#VUID-vkCmdWriteTimestamp2-stage-03930)VUID-vkCmdWriteTimestamp2-stage-03930  
  If the [`tessellationShader`](#features-tessellationShader) feature is not enabled, `stage` **must** not contain `VK_PIPELINE_STAGE_2_TESSELLATION_CONTROL_SHADER_BIT` or `VK_PIPELINE_STAGE_2_TESSELLATION_EVALUATION_SHADER_BIT`
- [](#VUID-vkCmdWriteTimestamp2-stage-03931)VUID-vkCmdWriteTimestamp2-stage-03931  
  If the [`conditionalRendering`](#features-conditionalRendering) feature is not enabled, `stage` **must** not contain `VK_PIPELINE_STAGE_2_CONDITIONAL_RENDERING_BIT_EXT`
- [](#VUID-vkCmdWriteTimestamp2-stage-03932)VUID-vkCmdWriteTimestamp2-stage-03932  
  If the [`fragmentDensityMap`](#features-fragmentDensityMap) feature is not enabled, `stage` **must** not contain `VK_PIPELINE_STAGE_2_FRAGMENT_DENSITY_PROCESS_BIT_EXT`
- [](#VUID-vkCmdWriteTimestamp2-stage-03933)VUID-vkCmdWriteTimestamp2-stage-03933  
  If the [`transformFeedback`](#features-transformFeedback) feature is not enabled, `stage` **must** not contain `VK_PIPELINE_STAGE_2_TRANSFORM_FEEDBACK_BIT_EXT`
- [](#VUID-vkCmdWriteTimestamp2-stage-03934)VUID-vkCmdWriteTimestamp2-stage-03934  
  If the [`meshShader`](#features-meshShader) feature is not enabled, `stage` **must** not contain `VK_PIPELINE_STAGE_2_MESH_SHADER_BIT_EXT`
- [](#VUID-vkCmdWriteTimestamp2-stage-03935)VUID-vkCmdWriteTimestamp2-stage-03935  
  If the [`taskShader`](#features-taskShader) feature is not enabled, `stage` **must** not contain `VK_PIPELINE_STAGE_2_TASK_SHADER_BIT_EXT`
- [](#VUID-vkCmdWriteTimestamp2-stage-07316)VUID-vkCmdWriteTimestamp2-stage-07316  
  If neither of the [`shadingRateImage`](#features-shadingRateImage) or the [`attachmentFragmentShadingRate`](#features-attachmentFragmentShadingRate) features are enabled, `stage` **must** not contain `VK_PIPELINE_STAGE_2_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR`
- [](#VUID-vkCmdWriteTimestamp2-stage-04957)VUID-vkCmdWriteTimestamp2-stage-04957  
  If the [`subpassShading`](#features-subpassShading) feature is not enabled, `stage` **must** not contain `VK_PIPELINE_STAGE_2_SUBPASS_SHADER_BIT_HUAWEI`
- [](#VUID-vkCmdWriteTimestamp2-stage-04995)VUID-vkCmdWriteTimestamp2-stage-04995  
  If the [`invocationMask`](#features-invocationMask) feature is not enabled, `stage` **must** not contain `VK_PIPELINE_STAGE_2_INVOCATION_MASK_BIT_HUAWEI`
- [](#VUID-vkCmdWriteTimestamp2-stage-07946)VUID-vkCmdWriteTimestamp2-stage-07946  
  If neither the [VK\_NV\_ray\_tracing](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_ray_tracing.html) extension or the [`rayTracingPipeline`](#features-rayTracingPipeline) feature are enabled, `stage` **must** not contain `VK_PIPELINE_STAGE_2_RAY_TRACING_SHADER_BIT_KHR`
- [](#VUID-vkCmdWriteTimestamp2-stage-10751)VUID-vkCmdWriteTimestamp2-stage-10751  
  If the [`accelerationStructure`](#features-accelerationStructure) feature is not enabled, `stage` **must** not contain `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR`
- [](#VUID-vkCmdWriteTimestamp2-stage-10752)VUID-vkCmdWriteTimestamp2-stage-10752  
  If the [`rayTracingMaintenance1`](#features-rayTracingMaintenance1) feature is not enabled, `stage` **must** not contain `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_COPY_BIT_KHR`
- [](#VUID-vkCmdWriteTimestamp2-stage-10753)VUID-vkCmdWriteTimestamp2-stage-10753  
  If the [`micromap`](#features-micromap) feature is not enabled, `stage` **must** not contain `VK_PIPELINE_STAGE_2_MICROMAP_BUILD_BIT_EXT`
- [](#VUID-vkCmdWriteTimestamp2-synchronization2-03858)VUID-vkCmdWriteTimestamp2-synchronization2-03858  
  The [`synchronization2`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-synchronization2) feature **must** be enabled
- [](#VUID-vkCmdWriteTimestamp2-stage-03859)VUID-vkCmdWriteTimestamp2-stage-03859  
  `stage` **must** only include a single pipeline stage
- [](#VUID-vkCmdWriteTimestamp2-stage-03860)VUID-vkCmdWriteTimestamp2-stage-03860  
  `stage` **must** only include stages valid for the queue family that was used to create the command pool that `commandBuffer` was allocated from
- [](#VUID-vkCmdWriteTimestamp2-queryPool-03861)VUID-vkCmdWriteTimestamp2-queryPool-03861  
  `queryPool` **must** have been created with a `queryType` of `VK_QUERY_TYPE_TIMESTAMP`
- [](#VUID-vkCmdWriteTimestamp2-timestampValidBits-03863)VUID-vkCmdWriteTimestamp2-timestampValidBits-03863  
  The command pool’s queue family **must** support a non-zero `timestampValidBits`
- [](#VUID-vkCmdWriteTimestamp2-query-04903)VUID-vkCmdWriteTimestamp2-query-04903  
  `query` **must** be less than the number of queries in `queryPool`
- [](#VUID-vkCmdWriteTimestamp2-None-03864)VUID-vkCmdWriteTimestamp2-None-03864  
  All queries used by the command **must** be *unavailable*
- [](#VUID-vkCmdWriteTimestamp2-query-03865)VUID-vkCmdWriteTimestamp2-query-03865  
  If `vkCmdWriteTimestamp2` is called within a render pass instance, the sum of `query` and the number of bits set in the current subpass’s view mask **must** be less than or equal to the number of queries in `queryPool`
- [](#VUID-vkCmdWriteTimestamp2-None-10639)VUID-vkCmdWriteTimestamp2-None-10639  
  This command **must** not be recorded when [per-tile execution model](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-per-tile-execution-model) is enabled

Valid Usage (Implicit)

- [](#VUID-vkCmdWriteTimestamp2-commandBuffer-parameter)VUID-vkCmdWriteTimestamp2-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdWriteTimestamp2-stage-parameter)VUID-vkCmdWriteTimestamp2-stage-parameter  
  `stage` **must** be a valid combination of [VkPipelineStageFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlagBits2.html) values
- [](#VUID-vkCmdWriteTimestamp2-queryPool-parameter)VUID-vkCmdWriteTimestamp2-queryPool-parameter  
  `queryPool` **must** be a valid [VkQueryPool](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPool.html) handle
- [](#VUID-vkCmdWriteTimestamp2-commandBuffer-recording)VUID-vkCmdWriteTimestamp2-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdWriteTimestamp2-commandBuffer-cmdpool)VUID-vkCmdWriteTimestamp2-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support VK\_QUEUE\_COMPUTE\_BIT, VK\_QUEUE\_GRAPHICS\_BIT, VK\_QUEUE\_TRANSFER\_BIT, VK\_QUEUE\_VIDEO\_DECODE\_BIT\_KHR, or VK\_QUEUE\_VIDEO\_ENCODE\_BIT\_KHR operations
- [](#VUID-vkCmdWriteTimestamp2-commonparent)VUID-vkCmdWriteTimestamp2-commonparent  
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

VK\_QUEUE\_COMPUTE\_BIT  
VK\_QUEUE\_GRAPHICS\_BIT  
VK\_QUEUE\_TRANSFER\_BIT  
VK\_QUEUE\_VIDEO\_DECODE\_BIT\_KHR  
VK\_QUEUE\_VIDEO\_ENCODE\_BIT\_KHR

Action

Conditional Rendering

vkCmdWriteTimestamp2 is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_KHR\_synchronization2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_synchronization2.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkPipelineStageFlags2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlags2.html), [VkQueryPool](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPool.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdWriteTimestamp2).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0