# vkCmdBeginQuery(3) Manual Page

## Name

vkCmdBeginQuery - Begin a query



## [](#_c_specification)C Specification

To begin a query, call:

```c++
// Provided by VK_VERSION_1_0
void vkCmdBeginQuery(
    VkCommandBuffer                             commandBuffer,
    VkQueryPool                                 queryPool,
    uint32_t                                    query,
    VkQueryControlFlags                         flags);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which this command will be recorded.
- `queryPool` is the query pool that will manage the results of the query.
- `query` is the query index within the query pool that will contain the results.
- `flags` is a bitmask of [VkQueryControlFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryControlFlagBits.html) specifying constraints on the types of queries that **can** be performed.

## [](#_description)Description

If the `queryType` of the pool is `VK_QUERY_TYPE_OCCLUSION` and `flags` contains `VK_QUERY_CONTROL_PRECISE_BIT`, an implementation **must** return a result that matches the actual number of samples passed. This is described in more detail in [Occlusion Queries](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#queries-occlusion).

Calling `vkCmdBeginQuery` is equivalent to calling [vkCmdBeginQueryIndexedEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBeginQueryIndexedEXT.html) with the `index` parameter set to zero.

After beginning a query, that query is considered *active* within the command buffer it was called in until that same query is ended. Queries active in a primary command buffer when secondary command buffers are executed are considered active for those secondary command buffers.

Furthermore, if the query is started within a video coding scope, the following command buffer states are initialized for the query type:

- []()The *active\_query\_index* is set to the value specified by `query`.
- []()The *last activatable query index* is also set to the value specified by `query`.

Each [video coding operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#video-coding) stores a result to the query corresponding to the current active query index, followed by incrementing the active query index. If the active query index gets incremented past the last activatable query index, issuing any further video coding operations results in undefined behavior.

Note

In practice, this means that currently no more than a single video coding operation **must** be issued between a begin and end query pair.

This command defines an execution dependency between other query commands that reference the same query.

The first [synchronization scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-scopes) includes all commands which reference the queries in `queryPool` indicated by `query` that occur earlier in [submission order](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-submission-order).

The second [synchronization scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-scopes) includes all commands which reference the queries in `queryPool` indicated by `query` that occur later in [submission order](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-submission-order).

The operation of this command happens after the first scope and happens before the second scope.

Valid Usage

- [](#VUID-vkCmdBeginQuery-None-00807)VUID-vkCmdBeginQuery-None-00807  
  All queries used by the command **must** be *unavailable*
- [](#VUID-vkCmdBeginQuery-queryType-02804)VUID-vkCmdBeginQuery-queryType-02804  
  The `queryType` used to create `queryPool` **must** not be `VK_QUERY_TYPE_TIMESTAMP`
- [](#VUID-vkCmdBeginQuery-queryType-04728)VUID-vkCmdBeginQuery-queryType-04728  
  The `queryType` used to create `queryPool` **must** not be `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_COMPACTED_SIZE_KHR` or `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_SERIALIZATION_SIZE_KHR`
- [](#VUID-vkCmdBeginQuery-queryType-06741)VUID-vkCmdBeginQuery-queryType-06741  
  The `queryType` used to create `queryPool` **must** not be `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_SIZE_KHR` or `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_SERIALIZATION_BOTTOM_LEVEL_POINTERS_KHR`
- [](#VUID-vkCmdBeginQuery-queryType-04729)VUID-vkCmdBeginQuery-queryType-04729  
  The `queryType` used to create `queryPool` **must** not be `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_COMPACTED_SIZE_NV`
- [](#VUID-vkCmdBeginQuery-queryType-00800)VUID-vkCmdBeginQuery-queryType-00800  
  If the [`occlusionQueryPrecise`](#features-occlusionQueryPrecise) feature is not enabled, or the `queryType` used to create `queryPool` was not `VK_QUERY_TYPE_OCCLUSION`, `flags` **must** not contain `VK_QUERY_CONTROL_PRECISE_BIT`
- [](#VUID-vkCmdBeginQuery-query-00802)VUID-vkCmdBeginQuery-query-00802  
  `query` **must** be less than the number of queries in `queryPool`
- [](#VUID-vkCmdBeginQuery-queryType-00803)VUID-vkCmdBeginQuery-queryType-00803  
  If the `queryType` used to create `queryPool` was `VK_QUERY_TYPE_OCCLUSION`, the `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics operations
- [](#VUID-vkCmdBeginQuery-queryType-00804)VUID-vkCmdBeginQuery-queryType-00804  
  If the `queryType` used to create `queryPool` was `VK_QUERY_TYPE_PIPELINE_STATISTICS` and any of the `pipelineStatistics` indicate graphics operations, the `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics operations
- [](#VUID-vkCmdBeginQuery-queryType-00805)VUID-vkCmdBeginQuery-queryType-00805  
  If the `queryType` used to create `queryPool` was `VK_QUERY_TYPE_PIPELINE_STATISTICS` and any of the `pipelineStatistics` indicate compute operations, the `VkCommandPool` that `commandBuffer` was allocated from **must** support compute operations
- [](#VUID-vkCmdBeginQuery-commandBuffer-01885)VUID-vkCmdBeginQuery-commandBuffer-01885  
  `commandBuffer` **must** not be a protected command buffer
- [](#VUID-vkCmdBeginQuery-query-00808)VUID-vkCmdBeginQuery-query-00808  
  If called within a render pass instance, the sum of `query` and the number of bits set in the current subpass’s view mask **must** be less than or equal to the number of queries in `queryPool`
- [](#VUID-vkCmdBeginQuery-queryType-07126)VUID-vkCmdBeginQuery-queryType-07126  
  If the `queryType` used to create `queryPool` was `VK_QUERY_TYPE_RESULT_STATUS_ONLY_KHR`, then the `VkCommandPool` that `commandBuffer` was allocated from **must** have been created with a queue family index that supports [result status queries](#queries-result-status-only), as indicated by [VkQueueFamilyQueryResultStatusPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueFamilyQueryResultStatusPropertiesKHR.html)::`queryResultStatusSupport`
- [](#VUID-vkCmdBeginQuery-None-07127)VUID-vkCmdBeginQuery-None-07127  
  If there is a bound video session, then there **must** be no [active](#queries-operation-active) queries
- [](#VUID-vkCmdBeginQuery-None-08370)VUID-vkCmdBeginQuery-None-08370  
  If there is a bound video session, then it **must** not have been created with `VK_VIDEO_SESSION_CREATE_INLINE_QUERIES_BIT_KHR`
- [](#VUID-vkCmdBeginQuery-queryType-07128)VUID-vkCmdBeginQuery-queryType-07128  
  If the `queryType` used to create `queryPool` was `VK_QUERY_TYPE_RESULT_STATUS_ONLY_KHR` and there is a bound video session, then `queryPool` **must** have been created with a [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoProfileInfoKHR.html) structure included in the `pNext` chain of [VkQueryPoolCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPoolCreateInfo.html) identical to the one specified in [VkVideoSessionCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionCreateInfoKHR.html)::`pVideoProfile` the bound video session was created with
- [](#VUID-vkCmdBeginQuery-queryType-07129)VUID-vkCmdBeginQuery-queryType-07129  
  If the `queryType` used to create `queryPool` was `VK_QUERY_TYPE_VIDEO_ENCODE_FEEDBACK_KHR`, then there **must** be a bound video session
- [](#VUID-vkCmdBeginQuery-queryType-07130)VUID-vkCmdBeginQuery-queryType-07130  
  If the `queryType` used to create `queryPool` was `VK_QUERY_TYPE_VIDEO_ENCODE_FEEDBACK_KHR` and there is a bound video session, then `queryPool` **must** have been created with a [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoProfileInfoKHR.html) structure included in the `pNext` chain of [VkQueryPoolCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPoolCreateInfo.html) identical to the one specified in [VkVideoSessionCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionCreateInfoKHR.html)::`pVideoProfile` the bound video session was created with
- [](#VUID-vkCmdBeginQuery-queryType-07131)VUID-vkCmdBeginQuery-queryType-07131  
  If the `queryType` used to create `queryPool` was not `VK_QUERY_TYPE_RESULT_STATUS_ONLY_KHR` or `VK_QUERY_TYPE_VIDEO_ENCODE_FEEDBACK_KHR`, then there **must** be no bound video session
- [](#VUID-vkCmdBeginQuery-None-10681)VUID-vkCmdBeginQuery-None-10681  
  This command **must** not be recorded when [per-tile execution model](#renderpass-per-tile-execution-model) is enabled
- [](#VUID-vkCmdBeginQuery-queryPool-01922)VUID-vkCmdBeginQuery-queryPool-01922  
  `queryPool` **must** have been created with a `queryType` that differs from that of any queries that are [active](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#queries-operation-active) within `commandBuffer`
- [](#VUID-vkCmdBeginQuery-queryType-07070)VUID-vkCmdBeginQuery-queryType-07070  
  If the `queryType` used to create `queryPool` was `VK_QUERY_TYPE_MESH_PRIMITIVES_GENERATED_EXT` the `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics operations
- [](#VUID-vkCmdBeginQuery-queryType-02327)VUID-vkCmdBeginQuery-queryType-02327  
  If the `queryType` used to create `queryPool` was `VK_QUERY_TYPE_TRANSFORM_FEEDBACK_STREAM_EXT` the `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics operations
- [](#VUID-vkCmdBeginQuery-queryType-02328)VUID-vkCmdBeginQuery-queryType-02328  
  If the `queryType` used to create `queryPool` was `VK_QUERY_TYPE_TRANSFORM_FEEDBACK_STREAM_EXT` then `VkPhysicalDeviceTransformFeedbackPropertiesEXT`::`transformFeedbackQueries` **must** be supported
- [](#VUID-vkCmdBeginQuery-queryType-06687)VUID-vkCmdBeginQuery-queryType-06687  
  If the `queryType` used to create `queryPool` was `VK_QUERY_TYPE_PRIMITIVES_GENERATED_EXT` the `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics operations
- [](#VUID-vkCmdBeginQuery-queryType-06688)VUID-vkCmdBeginQuery-queryType-06688  
  If the `queryType` used to create `queryPool` was `VK_QUERY_TYPE_PRIMITIVES_GENERATED_EXT` then [`primitivesGeneratedQuery`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-primitivesGeneratedQuery) **must** be enabled

<!--THE END-->

- [](#VUID-vkCmdBeginQuery-queryPool-07289)VUID-vkCmdBeginQuery-queryPool-07289  
  If `queryPool` was created with a `queryType` of `VK_QUERY_TYPE_PERFORMANCE_QUERY_KHR`, then the [VkQueryPoolPerformanceCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPoolPerformanceCreateInfoKHR.html)::`queueFamilyIndex` `queryPool` was created with **must** equal the queue family index of the `VkCommandPool` that `commandBuffer` was allocated from
- [](#VUID-vkCmdBeginQuery-queryPool-03223)VUID-vkCmdBeginQuery-queryPool-03223  
  If `queryPool` was created with a `queryType` of `VK_QUERY_TYPE_PERFORMANCE_QUERY_KHR`, the [profiling lock](#profiling-lock) **must** have been held before [vkBeginCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/vkBeginCommandBuffer.html) was called on `commandBuffer`
- [](#VUID-vkCmdBeginQuery-queryPool-03224)VUID-vkCmdBeginQuery-queryPool-03224  
  If `queryPool` was created with a `queryType` of `VK_QUERY_TYPE_PERFORMANCE_QUERY_KHR` and one of the counters used to create `queryPool` was `VK_PERFORMANCE_COUNTER_SCOPE_COMMAND_BUFFER_KHR`, the query begin **must** be the first recorded command in `commandBuffer`
- [](#VUID-vkCmdBeginQuery-queryPool-03225)VUID-vkCmdBeginQuery-queryPool-03225  
  If `queryPool` was created with a `queryType` of `VK_QUERY_TYPE_PERFORMANCE_QUERY_KHR` and one of the counters used to create `queryPool` was `VK_PERFORMANCE_COUNTER_SCOPE_RENDER_PASS_KHR`, the begin command **must** not be recorded within a render pass instance
- [](#VUID-vkCmdBeginQuery-queryPool-03226)VUID-vkCmdBeginQuery-queryPool-03226  
  If `queryPool` was created with a `queryType` of `VK_QUERY_TYPE_PERFORMANCE_QUERY_KHR` and another query pool with a `queryType` `VK_QUERY_TYPE_PERFORMANCE_QUERY_KHR` has been used within `commandBuffer`, its parent primary command buffer or secondary command buffer recorded within the same parent primary command buffer as `commandBuffer`, the [`performanceCounterMultipleQueryPools`](#features-performanceCounterMultipleQueryPools) feature **must** be enabled
- [](#VUID-vkCmdBeginQuery-None-02863)VUID-vkCmdBeginQuery-None-02863  
  If `queryPool` was created with a `queryType` of `VK_QUERY_TYPE_PERFORMANCE_QUERY_KHR`, this command **must** not be recorded in a command buffer that, either directly or through secondary command buffers, also contains a `vkCmdResetQueryPool` command affecting the same query

Valid Usage (Implicit)

- [](#VUID-vkCmdBeginQuery-commandBuffer-parameter)VUID-vkCmdBeginQuery-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdBeginQuery-queryPool-parameter)VUID-vkCmdBeginQuery-queryPool-parameter  
  `queryPool` **must** be a valid [VkQueryPool](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPool.html) handle
- [](#VUID-vkCmdBeginQuery-flags-parameter)VUID-vkCmdBeginQuery-flags-parameter  
  `flags` **must** be a valid combination of [VkQueryControlFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryControlFlagBits.html) values
- [](#VUID-vkCmdBeginQuery-commandBuffer-recording)VUID-vkCmdBeginQuery-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdBeginQuery-commandBuffer-cmdpool)VUID-vkCmdBeginQuery-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics, compute, decode, or encode operations
- [](#VUID-vkCmdBeginQuery-commonparent)VUID-vkCmdBeginQuery-commonparent  
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

Graphics  
Compute  
Decode  
Encode

Action  
State

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkQueryControlFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryControlFlags.html), [VkQueryPool](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPool.html), [vkCmdBeginQueryIndexedEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBeginQueryIndexedEXT.html), [vkCmdEndQuery](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdEndQuery.html), [vkCmdEndQueryIndexedEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdEndQueryIndexedEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdBeginQuery)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0