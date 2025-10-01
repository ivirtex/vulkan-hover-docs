# vkCmdBeginQueryIndexedEXT(3) Manual Page

## Name

vkCmdBeginQueryIndexedEXT - Begin an indexed query



## [](#_c_specification)C Specification

To begin an indexed query, call:

```c++
// Provided by VK_EXT_transform_feedback
void vkCmdBeginQueryIndexedEXT(
    VkCommandBuffer                             commandBuffer,
    VkQueryPool                                 queryPool,
    uint32_t                                    query,
    VkQueryControlFlags                         flags,
    uint32_t                                    index);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which this command will be recorded.
- `queryPool` is the query pool that will manage the results of the query.
- `query` is the query index within the query pool that will contain the results.
- `flags` is a bitmask of [VkQueryControlFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryControlFlagBits.html) specifying constraints on the types of queries that **can** be performed.
- `index` is the query type specific index. When the query type is `VK_QUERY_TYPE_TRANSFORM_FEEDBACK_STREAM_EXT` or `VK_QUERY_TYPE_PRIMITIVES_GENERATED_EXT`, the index represents the vertex stream.

## [](#_description)Description

The `vkCmdBeginQueryIndexedEXT` command operates the same as the [vkCmdBeginQuery](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBeginQuery.html) command, except that it also accepts a query type specific `index` parameter.

This command defines an execution dependency between other query commands that reference the same query index.

The first [synchronization scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-scopes) includes all commands which reference the queries in `queryPool` indicated by `query` and `index` that occur earlier in [submission order](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-submission-order).

The second [synchronization scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-scopes) includes all commands which reference the queries in `queryPool` indicated by `query` and `index` that occur later in [submission order](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-submission-order).

The operation of this command happens after the first scope and happens before the second scope.

Valid Usage

- [](#VUID-vkCmdBeginQueryIndexedEXT-None-00807)VUID-vkCmdBeginQueryIndexedEXT-None-00807  
  All queries used by the command **must** be *unavailable*
- [](#VUID-vkCmdBeginQueryIndexedEXT-queryType-02804)VUID-vkCmdBeginQueryIndexedEXT-queryType-02804  
  The `queryType` used to create `queryPool` **must** not be `VK_QUERY_TYPE_TIMESTAMP`
- [](#VUID-vkCmdBeginQueryIndexedEXT-queryType-04728)VUID-vkCmdBeginQueryIndexedEXT-queryType-04728  
  The `queryType` used to create `queryPool` **must** not be `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_COMPACTED_SIZE_KHR` or `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_SERIALIZATION_SIZE_KHR`
- [](#VUID-vkCmdBeginQueryIndexedEXT-queryType-06741)VUID-vkCmdBeginQueryIndexedEXT-queryType-06741  
  The `queryType` used to create `queryPool` **must** not be `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_SIZE_KHR` or `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_SERIALIZATION_BOTTOM_LEVEL_POINTERS_KHR`
- [](#VUID-vkCmdBeginQueryIndexedEXT-queryType-04729)VUID-vkCmdBeginQueryIndexedEXT-queryType-04729  
  The `queryType` used to create `queryPool` **must** not be `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_COMPACTED_SIZE_NV`
- [](#VUID-vkCmdBeginQueryIndexedEXT-queryType-00800)VUID-vkCmdBeginQueryIndexedEXT-queryType-00800  
  If the [`occlusionQueryPrecise`](#features-occlusionQueryPrecise) feature is not enabled, or the `queryType` used to create `queryPool` was not `VK_QUERY_TYPE_OCCLUSION`, `flags` **must** not contain `VK_QUERY_CONTROL_PRECISE_BIT`
- [](#VUID-vkCmdBeginQueryIndexedEXT-query-00802)VUID-vkCmdBeginQueryIndexedEXT-query-00802  
  `query` **must** be less than the number of queries in `queryPool`
- [](#VUID-vkCmdBeginQueryIndexedEXT-queryType-00803)VUID-vkCmdBeginQueryIndexedEXT-queryType-00803  
  If the `queryType` used to create `queryPool` was `VK_QUERY_TYPE_OCCLUSION`, the `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics operations
- [](#VUID-vkCmdBeginQueryIndexedEXT-queryType-00804)VUID-vkCmdBeginQueryIndexedEXT-queryType-00804  
  If the `queryType` used to create `queryPool` was `VK_QUERY_TYPE_PIPELINE_STATISTICS` and any of the `pipelineStatistics` indicate graphics operations, the `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics operations
- [](#VUID-vkCmdBeginQueryIndexedEXT-queryType-00805)VUID-vkCmdBeginQueryIndexedEXT-queryType-00805  
  If the `queryType` used to create `queryPool` was `VK_QUERY_TYPE_PIPELINE_STATISTICS` and any of the `pipelineStatistics` indicate compute operations, the `VkCommandPool` that `commandBuffer` was allocated from **must** support compute operations
- [](#VUID-vkCmdBeginQueryIndexedEXT-commandBuffer-01885)VUID-vkCmdBeginQueryIndexedEXT-commandBuffer-01885  
  `commandBuffer` **must** not be a protected command buffer
- [](#VUID-vkCmdBeginQueryIndexedEXT-query-00808)VUID-vkCmdBeginQueryIndexedEXT-query-00808  
  If called within a render pass instance, the sum of `query` and the number of bits set in the current subpass’s view mask **must** be less than or equal to the number of queries in `queryPool`
- [](#VUID-vkCmdBeginQueryIndexedEXT-queryType-07126)VUID-vkCmdBeginQueryIndexedEXT-queryType-07126  
  If the `queryType` used to create `queryPool` was `VK_QUERY_TYPE_RESULT_STATUS_ONLY_KHR`, then the `VkCommandPool` that `commandBuffer` was allocated from **must** have been created with a queue family index that supports [result status queries](#queries-result-status-only), as indicated by [VkQueueFamilyQueryResultStatusPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueFamilyQueryResultStatusPropertiesKHR.html)::`queryResultStatusSupport`
- [](#VUID-vkCmdBeginQueryIndexedEXT-None-07127)VUID-vkCmdBeginQueryIndexedEXT-None-07127  
  If there is a bound video session, then there **must** be no [active](#queries-operation-active) queries
- [](#VUID-vkCmdBeginQueryIndexedEXT-None-08370)VUID-vkCmdBeginQueryIndexedEXT-None-08370  
  If there is a bound video session, then it **must** not have been created with `VK_VIDEO_SESSION_CREATE_INLINE_QUERIES_BIT_KHR`
- [](#VUID-vkCmdBeginQueryIndexedEXT-queryType-07128)VUID-vkCmdBeginQueryIndexedEXT-queryType-07128  
  If the `queryType` used to create `queryPool` was `VK_QUERY_TYPE_RESULT_STATUS_ONLY_KHR` and there is a bound video session, then `queryPool` **must** have been created with a [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoProfileInfoKHR.html) structure included in the `pNext` chain of [VkQueryPoolCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPoolCreateInfo.html) identical to the one specified in [VkVideoSessionCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionCreateInfoKHR.html)::`pVideoProfile` the bound video session was created with
- [](#VUID-vkCmdBeginQueryIndexedEXT-queryType-07129)VUID-vkCmdBeginQueryIndexedEXT-queryType-07129  
  If the `queryType` used to create `queryPool` was `VK_QUERY_TYPE_VIDEO_ENCODE_FEEDBACK_KHR`, then there **must** be a bound video session
- [](#VUID-vkCmdBeginQueryIndexedEXT-queryType-07130)VUID-vkCmdBeginQueryIndexedEXT-queryType-07130  
  If the `queryType` used to create `queryPool` was `VK_QUERY_TYPE_VIDEO_ENCODE_FEEDBACK_KHR` and there is a bound video session, then `queryPool` **must** have been created with a [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoProfileInfoKHR.html) structure included in the `pNext` chain of [VkQueryPoolCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPoolCreateInfo.html) identical to the one specified in [VkVideoSessionCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionCreateInfoKHR.html)::`pVideoProfile` the bound video session was created with
- [](#VUID-vkCmdBeginQueryIndexedEXT-queryType-07131)VUID-vkCmdBeginQueryIndexedEXT-queryType-07131  
  If the `queryType` used to create `queryPool` was not `VK_QUERY_TYPE_RESULT_STATUS_ONLY_KHR` or `VK_QUERY_TYPE_VIDEO_ENCODE_FEEDBACK_KHR`, then there **must** be no bound video session
- [](#VUID-vkCmdBeginQueryIndexedEXT-None-10681)VUID-vkCmdBeginQueryIndexedEXT-None-10681  
  This command **must** not be recorded when [per-tile execution model](#renderpass-per-tile-execution-model) is enabled
- [](#VUID-vkCmdBeginQueryIndexedEXT-queryPool-04753)VUID-vkCmdBeginQueryIndexedEXT-queryPool-04753  
  If the `queryPool` was created with the same `queryType` as that of another [active](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#queries-operation-active) query within `commandBuffer`, then `index` **must** not match the index used for the active query
- [](#VUID-vkCmdBeginQueryIndexedEXT-queryType-02338)VUID-vkCmdBeginQueryIndexedEXT-queryType-02338  
  If the `queryType` used to create `queryPool` was `VK_QUERY_TYPE_TRANSFORM_FEEDBACK_STREAM_EXT` the `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics operations
- [](#VUID-vkCmdBeginQueryIndexedEXT-queryType-02339)VUID-vkCmdBeginQueryIndexedEXT-queryType-02339  
  If the `queryType` used to create `queryPool` was `VK_QUERY_TYPE_TRANSFORM_FEEDBACK_STREAM_EXT` the `index` parameter **must** be less than `VkPhysicalDeviceTransformFeedbackPropertiesEXT`::`maxTransformFeedbackStreams`
- [](#VUID-vkCmdBeginQueryIndexedEXT-queryType-06692)VUID-vkCmdBeginQueryIndexedEXT-queryType-06692  
  If the `queryType` used to create `queryPool` was not `VK_QUERY_TYPE_TRANSFORM_FEEDBACK_STREAM_EXT` and not `VK_QUERY_TYPE_PRIMITIVES_GENERATED_EXT`, the `index` **must** be zero
- [](#VUID-vkCmdBeginQueryIndexedEXT-queryType-06689)VUID-vkCmdBeginQueryIndexedEXT-queryType-06689  
  If the `queryType` used to create `queryPool` was `VK_QUERY_TYPE_PRIMITIVES_GENERATED_EXT` the `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics operations
- [](#VUID-vkCmdBeginQueryIndexedEXT-queryType-06690)VUID-vkCmdBeginQueryIndexedEXT-queryType-06690  
  If the `queryType` used to create `queryPool` was `VK_QUERY_TYPE_PRIMITIVES_GENERATED_EXT` the `index` parameter **must** be less than `VkPhysicalDeviceTransformFeedbackPropertiesEXT`::`maxTransformFeedbackStreams`
- [](#VUID-vkCmdBeginQueryIndexedEXT-queryType-06691)VUID-vkCmdBeginQueryIndexedEXT-queryType-06691  
  If the `queryType` used to create `queryPool` was `VK_QUERY_TYPE_PRIMITIVES_GENERATED_EXT` and the [`primitivesGeneratedQueryWithNonZeroStreams`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-primitivesGeneratedQueryWithNonZeroStreams) feature is not enabled, the `index` parameter **must** be zero
- [](#VUID-vkCmdBeginQueryIndexedEXT-queryType-06693)VUID-vkCmdBeginQueryIndexedEXT-queryType-06693  
  If the `queryType` used to create `queryPool` was `VK_QUERY_TYPE_PRIMITIVES_GENERATED_EXT` then [`primitivesGeneratedQuery`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-primitivesGeneratedQuery) **must** be enabled
- [](#VUID-vkCmdBeginQueryIndexedEXT-queryType-02341)VUID-vkCmdBeginQueryIndexedEXT-queryType-02341  
  If the `queryType` used to create `queryPool` was `VK_QUERY_TYPE_TRANSFORM_FEEDBACK_STREAM_EXT` then `VkPhysicalDeviceTransformFeedbackPropertiesEXT`::`transformFeedbackQueries` **must** be supported
- [](#VUID-vkCmdBeginQueryIndexedEXT-queryType-07071)VUID-vkCmdBeginQueryIndexedEXT-queryType-07071  
  The `queryType` used to create `queryPool` **must** not be `VK_QUERY_TYPE_MESH_PRIMITIVES_GENERATED_EXT`

<!--THE END-->

- [](#VUID-vkCmdBeginQueryIndexedEXT-queryPool-07289)VUID-vkCmdBeginQueryIndexedEXT-queryPool-07289  
  If `queryPool` was created with a `queryType` of `VK_QUERY_TYPE_PERFORMANCE_QUERY_KHR`, then the [VkQueryPoolPerformanceCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPoolPerformanceCreateInfoKHR.html)::`queueFamilyIndex` `queryPool` was created with **must** equal the queue family index of the `VkCommandPool` that `commandBuffer` was allocated from
- [](#VUID-vkCmdBeginQueryIndexedEXT-queryPool-03223)VUID-vkCmdBeginQueryIndexedEXT-queryPool-03223  
  If `queryPool` was created with a `queryType` of `VK_QUERY_TYPE_PERFORMANCE_QUERY_KHR`, the [profiling lock](#profiling-lock) **must** have been held before [vkBeginCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/vkBeginCommandBuffer.html) was called on `commandBuffer`
- [](#VUID-vkCmdBeginQueryIndexedEXT-queryPool-03224)VUID-vkCmdBeginQueryIndexedEXT-queryPool-03224  
  If `queryPool` was created with a `queryType` of `VK_QUERY_TYPE_PERFORMANCE_QUERY_KHR` and one of the counters used to create `queryPool` was `VK_PERFORMANCE_COUNTER_SCOPE_COMMAND_BUFFER_KHR`, the query begin **must** be the first recorded command in `commandBuffer`
- [](#VUID-vkCmdBeginQueryIndexedEXT-queryPool-03225)VUID-vkCmdBeginQueryIndexedEXT-queryPool-03225  
  If `queryPool` was created with a `queryType` of `VK_QUERY_TYPE_PERFORMANCE_QUERY_KHR` and one of the counters used to create `queryPool` was `VK_PERFORMANCE_COUNTER_SCOPE_RENDER_PASS_KHR`, the begin command **must** not be recorded within a render pass instance
- [](#VUID-vkCmdBeginQueryIndexedEXT-queryPool-03226)VUID-vkCmdBeginQueryIndexedEXT-queryPool-03226  
  If `queryPool` was created with a `queryType` of `VK_QUERY_TYPE_PERFORMANCE_QUERY_KHR` and another query pool with a `queryType` `VK_QUERY_TYPE_PERFORMANCE_QUERY_KHR` has been used within `commandBuffer`, its parent primary command buffer or secondary command buffer recorded within the same parent primary command buffer as `commandBuffer`, the [`performanceCounterMultipleQueryPools`](#features-performanceCounterMultipleQueryPools) feature **must** be enabled
- [](#VUID-vkCmdBeginQueryIndexedEXT-None-02863)VUID-vkCmdBeginQueryIndexedEXT-None-02863  
  If `queryPool` was created with a `queryType` of `VK_QUERY_TYPE_PERFORMANCE_QUERY_KHR`, this command **must** not be recorded in a command buffer that, either directly or through secondary command buffers, also contains a `vkCmdResetQueryPool` command affecting the same query

Valid Usage (Implicit)

- [](#VUID-vkCmdBeginQueryIndexedEXT-commandBuffer-parameter)VUID-vkCmdBeginQueryIndexedEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdBeginQueryIndexedEXT-queryPool-parameter)VUID-vkCmdBeginQueryIndexedEXT-queryPool-parameter  
  `queryPool` **must** be a valid [VkQueryPool](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPool.html) handle
- [](#VUID-vkCmdBeginQueryIndexedEXT-flags-parameter)VUID-vkCmdBeginQueryIndexedEXT-flags-parameter  
  `flags` **must** be a valid combination of [VkQueryControlFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryControlFlagBits.html) values
- [](#VUID-vkCmdBeginQueryIndexedEXT-commandBuffer-recording)VUID-vkCmdBeginQueryIndexedEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdBeginQueryIndexedEXT-commandBuffer-cmdpool)VUID-vkCmdBeginQueryIndexedEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support VK\_QUEUE\_COMPUTE\_BIT, VK\_QUEUE\_GRAPHICS\_BIT, VK\_QUEUE\_VIDEO\_DECODE\_BIT\_KHR, or VK\_QUEUE\_VIDEO\_ENCODE\_BIT\_KHR operations
- [](#VUID-vkCmdBeginQueryIndexedEXT-videocoding)VUID-vkCmdBeginQueryIndexedEXT-videocoding  
  This command **must** only be called outside of a video coding scope
- [](#VUID-vkCmdBeginQueryIndexedEXT-commonparent)VUID-vkCmdBeginQueryIndexedEXT-commonparent  
  Both of `commandBuffer`, and `queryPool` **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

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
VK\_QUEUE\_VIDEO\_DECODE\_BIT\_KHR  
VK\_QUEUE\_VIDEO\_ENCODE\_BIT\_KHR

Action  
State

Conditional Rendering

vkCmdBeginQueryIndexedEXT is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_EXT\_transform\_feedback](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_transform_feedback.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkQueryControlFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryControlFlags.html), [VkQueryPool](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPool.html), [vkCmdBeginQuery](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBeginQuery.html), [vkCmdEndQuery](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdEndQuery.html), [vkCmdEndQueryIndexedEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdEndQueryIndexedEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdBeginQueryIndexedEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0