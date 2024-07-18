# vkCmdBeginQuery(3) Manual Page

## Name

vkCmdBeginQuery - Begin a query



## <a href="#_c_specification" class="anchor"></a>C Specification

To begin a query, call:

``` c
// Provided by VK_VERSION_1_0
void vkCmdBeginQuery(
    VkCommandBuffer                             commandBuffer,
    VkQueryPool                                 queryPool,
    uint32_t                                    query,
    VkQueryControlFlags                         flags);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which this command will be
  recorded.

- `queryPool` is the query pool that will manage the results of the
  query.

- `query` is the query index within the query pool that will contain the
  results.

- `flags` is a bitmask of
  [VkQueryControlFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryControlFlagBits.html) specifying
  constraints on the types of queries that **can** be performed.

## <a href="#_description" class="anchor"></a>Description

If the `queryType` of the pool is `VK_QUERY_TYPE_OCCLUSION` and `flags`
contains `VK_QUERY_CONTROL_PRECISE_BIT`, an implementation **must**
return a result that matches the actual number of samples passed. This
is described in more detail in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#queries-occlusion"
target="_blank" rel="noopener">Occlusion Queries</a>.

Calling `vkCmdBeginQuery` is equivalent to calling
[vkCmdBeginQueryIndexedEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginQueryIndexedEXT.html) with the
`index` parameter set to zero.

After beginning a query, that query is considered *active* within the
command buffer it was called in until that same query is ended. Queries
active in a primary command buffer when secondary command buffers are
executed are considered active for those secondary command buffers.

Furthermore, if the query is started within a video coding scope, the
following command buffer states are initialized for the query type:

- <span id="queries-operation-active-query-index"></span> The
  *active_query_index* is set to the value specified by `query`.

- <span id="queries-operation-last-activatable-query-index"></span> The
  *last activatable query index* is also set to the value specified by
  `query`.

Each <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-coding"
target="_blank" rel="noopener">video coding operation</a> stores a
result to the query corresponding to the current active query index,
followed by incrementing the active query index. If the active query
index gets incremented past the last activatable query index, issuing
any further video coding operations results in undefined behavior.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>In practice, this means that currently no more than a single video
coding operation <strong>must</strong> be issued between a begin and end
query pair.</p></td>
</tr>
</tbody>
</table>

This command defines an execution dependency between other query
commands that reference the same query.

The first <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies-scopes"
target="_blank" rel="noopener">synchronization scope</a> includes all
commands which reference the queries in `queryPool` indicated by `query`
that occur earlier in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-submission-order"
target="_blank" rel="noopener">submission order</a>.

The second <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies-scopes"
target="_blank" rel="noopener">synchronization scope</a> includes all
commands which reference the queries in `queryPool` indicated by `query`
that occur later in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-submission-order"
target="_blank" rel="noopener">submission order</a>.

The operation of this command happens after the first scope and happens
before the second scope.

Valid Usage

- <a href="#VUID-vkCmdBeginQuery-None-00807"
  id="VUID-vkCmdBeginQuery-None-00807"></a>
  VUID-vkCmdBeginQuery-None-00807  
  All queries used by the command **must** be *unavailable*

- <a href="#VUID-vkCmdBeginQuery-queryType-02804"
  id="VUID-vkCmdBeginQuery-queryType-02804"></a>
  VUID-vkCmdBeginQuery-queryType-02804  
  The `queryType` used to create `queryPool` **must** not be
  `VK_QUERY_TYPE_TIMESTAMP`

- <a href="#VUID-vkCmdBeginQuery-queryType-04728"
  id="VUID-vkCmdBeginQuery-queryType-04728"></a>
  VUID-vkCmdBeginQuery-queryType-04728  
  The `queryType` used to create `queryPool` **must** not be
  `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_COMPACTED_SIZE_KHR` or
  `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_SERIALIZATION_SIZE_KHR`

- <a href="#VUID-vkCmdBeginQuery-queryType-06741"
  id="VUID-vkCmdBeginQuery-queryType-06741"></a>
  VUID-vkCmdBeginQuery-queryType-06741  
  The `queryType` used to create `queryPool` **must** not be
  `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_SIZE_KHR` or
  `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_SERIALIZATION_BOTTOM_LEVEL_POINTERS_KHR`

- <a href="#VUID-vkCmdBeginQuery-queryType-04729"
  id="VUID-vkCmdBeginQuery-queryType-04729"></a>
  VUID-vkCmdBeginQuery-queryType-04729  
  The `queryType` used to create `queryPool` **must** not be
  `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_COMPACTED_SIZE_NV`

- <a href="#VUID-vkCmdBeginQuery-queryType-00800"
  id="VUID-vkCmdBeginQuery-queryType-00800"></a>
  VUID-vkCmdBeginQuery-queryType-00800  
  If the [`occlusionQueryPrecise`](#features-occlusionQueryPrecise)
  feature is not enabled, or the `queryType` used to create `queryPool`
  was not `VK_QUERY_TYPE_OCCLUSION`, `flags` **must** not contain
  `VK_QUERY_CONTROL_PRECISE_BIT`

- <a href="#VUID-vkCmdBeginQuery-query-00802"
  id="VUID-vkCmdBeginQuery-query-00802"></a>
  VUID-vkCmdBeginQuery-query-00802  
  `query` **must** be less than the number of queries in `queryPool`

- <a href="#VUID-vkCmdBeginQuery-queryType-00803"
  id="VUID-vkCmdBeginQuery-queryType-00803"></a>
  VUID-vkCmdBeginQuery-queryType-00803  
  If the `queryType` used to create `queryPool` was
  `VK_QUERY_TYPE_OCCLUSION`, the `VkCommandPool` that `commandBuffer`
  was allocated from **must** support graphics operations

- <a href="#VUID-vkCmdBeginQuery-queryType-00804"
  id="VUID-vkCmdBeginQuery-queryType-00804"></a>
  VUID-vkCmdBeginQuery-queryType-00804  
  If the `queryType` used to create `queryPool` was
  `VK_QUERY_TYPE_PIPELINE_STATISTICS` and any of the
  `pipelineStatistics` indicate graphics operations, the `VkCommandPool`
  that `commandBuffer` was allocated from **must** support graphics
  operations

- <a href="#VUID-vkCmdBeginQuery-queryType-00805"
  id="VUID-vkCmdBeginQuery-queryType-00805"></a>
  VUID-vkCmdBeginQuery-queryType-00805  
  If the `queryType` used to create `queryPool` was
  `VK_QUERY_TYPE_PIPELINE_STATISTICS` and any of the
  `pipelineStatistics` indicate compute operations, the `VkCommandPool`
  that `commandBuffer` was allocated from **must** support compute
  operations

- <a href="#VUID-vkCmdBeginQuery-commandBuffer-01885"
  id="VUID-vkCmdBeginQuery-commandBuffer-01885"></a>
  VUID-vkCmdBeginQuery-commandBuffer-01885  
  `commandBuffer` **must** not be a protected command buffer

- <a href="#VUID-vkCmdBeginQuery-query-00808"
  id="VUID-vkCmdBeginQuery-query-00808"></a>
  VUID-vkCmdBeginQuery-query-00808  
  If called within a render pass instance, the sum of `query` and the
  number of bits set in the current subpass’s view mask **must** be less
  than or equal to the number of queries in `queryPool`

- <a href="#VUID-vkCmdBeginQuery-queryType-07126"
  id="VUID-vkCmdBeginQuery-queryType-07126"></a>
  VUID-vkCmdBeginQuery-queryType-07126  
  If the `queryType` used to create `queryPool` was
  `VK_QUERY_TYPE_RESULT_STATUS_ONLY_KHR`, then the `VkCommandPool` that
  `commandBuffer` was allocated from **must** have been created with a
  queue family index that supports [result status
  queries](#queries-result-status-only), as indicated by
  [VkQueueFamilyQueryResultStatusPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueueFamilyQueryResultStatusPropertiesKHR.html)::`queryResultStatusSupport`

- <a href="#VUID-vkCmdBeginQuery-None-07127"
  id="VUID-vkCmdBeginQuery-None-07127"></a>
  VUID-vkCmdBeginQuery-None-07127  
  If there is a bound video session, then there **must** be no
  [active](#queries-operation-active) queries

- <a href="#VUID-vkCmdBeginQuery-None-08370"
  id="VUID-vkCmdBeginQuery-None-08370"></a>
  VUID-vkCmdBeginQuery-None-08370  
  If there is a bound video session, then it **must** not have been
  created with `VK_VIDEO_SESSION_CREATE_INLINE_QUERIES_BIT_KHR`

- <a href="#VUID-vkCmdBeginQuery-queryType-07128"
  id="VUID-vkCmdBeginQuery-queryType-07128"></a>
  VUID-vkCmdBeginQuery-queryType-07128  
  If the `queryType` used to create `queryPool` was
  `VK_QUERY_TYPE_RESULT_STATUS_ONLY_KHR` and there is a bound video
  session, then `queryPool` **must** have been created with a
  [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileInfoKHR.html) structure included
  in the `pNext` chain of
  [VkQueryPoolCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPoolCreateInfo.html) identical to the
  one specified in
  [VkVideoSessionCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionCreateInfoKHR.html)::`pVideoProfile`
  the bound video session was created with

- <a href="#VUID-vkCmdBeginQuery-queryType-07129"
  id="VUID-vkCmdBeginQuery-queryType-07129"></a>
  VUID-vkCmdBeginQuery-queryType-07129  
  If the `queryType` used to create `queryPool` was
  `VK_QUERY_TYPE_VIDEO_ENCODE_FEEDBACK_KHR`, then there **must** be a
  bound video session

- <a href="#VUID-vkCmdBeginQuery-queryType-07130"
  id="VUID-vkCmdBeginQuery-queryType-07130"></a>
  VUID-vkCmdBeginQuery-queryType-07130  
  If the `queryType` used to create `queryPool` was
  `VK_QUERY_TYPE_VIDEO_ENCODE_FEEDBACK_KHR` and there is a bound video
  session, then `queryPool` **must** have been created with a
  [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileInfoKHR.html) structure included
  in the `pNext` chain of
  [VkQueryPoolCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPoolCreateInfo.html) identical to the
  one specified in
  [VkVideoSessionCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionCreateInfoKHR.html)::`pVideoProfile`
  the bound video session was created with

- <a href="#VUID-vkCmdBeginQuery-queryType-07131"
  id="VUID-vkCmdBeginQuery-queryType-07131"></a>
  VUID-vkCmdBeginQuery-queryType-07131  
  If the `queryType` used to create `queryPool` was not
  `VK_QUERY_TYPE_RESULT_STATUS_ONLY_KHR` or
  `VK_QUERY_TYPE_VIDEO_ENCODE_FEEDBACK_KHR`, then there **must** be no
  bound video session

- <a href="#VUID-vkCmdBeginQuery-queryPool-01922"
  id="VUID-vkCmdBeginQuery-queryPool-01922"></a>
  VUID-vkCmdBeginQuery-queryPool-01922  
  `queryPool` **must** have been created with a `queryType` that differs
  from that of any queries that are <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#queries-operation-active"
  target="_blank" rel="noopener">active</a> within `commandBuffer`

- <a href="#VUID-vkCmdBeginQuery-queryType-07070"
  id="VUID-vkCmdBeginQuery-queryType-07070"></a>
  VUID-vkCmdBeginQuery-queryType-07070  
  If the `queryType` used to create `queryPool` was
  `VK_QUERY_TYPE_MESH_PRIMITIVES_GENERATED_EXT` the `VkCommandPool` that
  `commandBuffer` was allocated from **must** support graphics
  operations

- <a href="#VUID-vkCmdBeginQuery-queryType-02327"
  id="VUID-vkCmdBeginQuery-queryType-02327"></a>
  VUID-vkCmdBeginQuery-queryType-02327  
  If the `queryType` used to create `queryPool` was
  `VK_QUERY_TYPE_TRANSFORM_FEEDBACK_STREAM_EXT` the `VkCommandPool` that
  `commandBuffer` was allocated from **must** support graphics
  operations

- <a href="#VUID-vkCmdBeginQuery-queryType-02328"
  id="VUID-vkCmdBeginQuery-queryType-02328"></a>
  VUID-vkCmdBeginQuery-queryType-02328  
  If the `queryType` used to create `queryPool` was
  `VK_QUERY_TYPE_TRANSFORM_FEEDBACK_STREAM_EXT` then
  `VkPhysicalDeviceTransformFeedbackPropertiesEXT`::`transformFeedbackQueries`
  **must** be supported

- <a href="#VUID-vkCmdBeginQuery-queryType-06687"
  id="VUID-vkCmdBeginQuery-queryType-06687"></a>
  VUID-vkCmdBeginQuery-queryType-06687  
  If the `queryType` used to create `queryPool` was
  `VK_QUERY_TYPE_PRIMITIVES_GENERATED_EXT` the `VkCommandPool` that
  `commandBuffer` was allocated from **must** support graphics
  operations

- <a href="#VUID-vkCmdBeginQuery-queryType-06688"
  id="VUID-vkCmdBeginQuery-queryType-06688"></a>
  VUID-vkCmdBeginQuery-queryType-06688  
  If the `queryType` used to create `queryPool` was
  `VK_QUERY_TYPE_PRIMITIVES_GENERATED_EXT` then <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-primitivesGeneratedQuery"
  target="_blank" rel="noopener"><code>primitivesGeneratedQuery</code></a>
  **must** be enabled

<!-- -->

- <a href="#VUID-vkCmdBeginQuery-queryPool-07289"
  id="VUID-vkCmdBeginQuery-queryPool-07289"></a>
  VUID-vkCmdBeginQuery-queryPool-07289  
  If `queryPool` was created with a `queryType` of
  `VK_QUERY_TYPE_PERFORMANCE_QUERY_KHR`, then the
  [VkQueryPoolPerformanceCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPoolPerformanceCreateInfoKHR.html)::`queueFamilyIndex`
  `queryPool` was created with **must** equal the queue family index of
  the `VkCommandPool` that `commandBuffer` was allocated from

- <a href="#VUID-vkCmdBeginQuery-queryPool-03223"
  id="VUID-vkCmdBeginQuery-queryPool-03223"></a>
  VUID-vkCmdBeginQuery-queryPool-03223  
  If `queryPool` was created with a `queryType` of
  `VK_QUERY_TYPE_PERFORMANCE_QUERY_KHR`, the [profiling
  lock](#profiling-lock) **must** have been held before
  [vkBeginCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkBeginCommandBuffer.html) was called on
  `commandBuffer`

- <a href="#VUID-vkCmdBeginQuery-queryPool-03224"
  id="VUID-vkCmdBeginQuery-queryPool-03224"></a>
  VUID-vkCmdBeginQuery-queryPool-03224  
  If `queryPool` was created with a `queryType` of
  `VK_QUERY_TYPE_PERFORMANCE_QUERY_KHR` and one of the counters used to
  create `queryPool` was
  `VK_PERFORMANCE_COUNTER_SCOPE_COMMAND_BUFFER_KHR`, the query begin
  **must** be the first recorded command in `commandBuffer`

- <a href="#VUID-vkCmdBeginQuery-queryPool-03225"
  id="VUID-vkCmdBeginQuery-queryPool-03225"></a>
  VUID-vkCmdBeginQuery-queryPool-03225  
  If `queryPool` was created with a `queryType` of
  `VK_QUERY_TYPE_PERFORMANCE_QUERY_KHR` and one of the counters used to
  create `queryPool` was `VK_PERFORMANCE_COUNTER_SCOPE_RENDER_PASS_KHR`,
  the begin command **must** not be recorded within a render pass
  instance

- <a href="#VUID-vkCmdBeginQuery-queryPool-03226"
  id="VUID-vkCmdBeginQuery-queryPool-03226"></a>
  VUID-vkCmdBeginQuery-queryPool-03226  
  If `queryPool` was created with a `queryType` of
  `VK_QUERY_TYPE_PERFORMANCE_QUERY_KHR` and another query pool with a
  `queryType` `VK_QUERY_TYPE_PERFORMANCE_QUERY_KHR` has been used within
  `commandBuffer`, its parent primary command buffer or secondary
  command buffer recorded within the same parent primary command buffer
  as `commandBuffer`, the
  [`performanceCounterMultipleQueryPools`](#features-performanceCounterMultipleQueryPools)
  feature **must** be enabled

- <a href="#VUID-vkCmdBeginQuery-None-02863"
  id="VUID-vkCmdBeginQuery-None-02863"></a>
  VUID-vkCmdBeginQuery-None-02863  
  If `queryPool` was created with a `queryType` of
  `VK_QUERY_TYPE_PERFORMANCE_QUERY_KHR`, this command **must** not be
  recorded in a command buffer that, either directly or through
  secondary command buffers, also contains a `vkCmdResetQueryPool`
  command affecting the same query

Valid Usage (Implicit)

- <a href="#VUID-vkCmdBeginQuery-commandBuffer-parameter"
  id="VUID-vkCmdBeginQuery-commandBuffer-parameter"></a>
  VUID-vkCmdBeginQuery-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdBeginQuery-queryPool-parameter"
  id="VUID-vkCmdBeginQuery-queryPool-parameter"></a>
  VUID-vkCmdBeginQuery-queryPool-parameter  
  `queryPool` **must** be a valid [VkQueryPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPool.html) handle

- <a href="#VUID-vkCmdBeginQuery-flags-parameter"
  id="VUID-vkCmdBeginQuery-flags-parameter"></a>
  VUID-vkCmdBeginQuery-flags-parameter  
  `flags` **must** be a valid combination of
  [VkQueryControlFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryControlFlagBits.html) values

- <a href="#VUID-vkCmdBeginQuery-commandBuffer-recording"
  id="VUID-vkCmdBeginQuery-commandBuffer-recording"></a>
  VUID-vkCmdBeginQuery-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdBeginQuery-commandBuffer-cmdpool"
  id="VUID-vkCmdBeginQuery-commandBuffer-cmdpool"></a>
  VUID-vkCmdBeginQuery-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics, compute, decode, or encode operations

- <a href="#VUID-vkCmdBeginQuery-commonparent"
  id="VUID-vkCmdBeginQuery-commonparent"></a>
  VUID-vkCmdBeginQuery-commonparent  
  Both of `commandBuffer`, and `queryPool` **must** have been created,
  allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

Host Synchronization

- Host access to `commandBuffer` **must** be externally synchronized

- Host access to the `VkCommandPool` that `commandBuffer` was allocated
  from **must** be externally synchronized

Command Properties

<table class="tableblock frame-all grid-all stretch">
<colgroup>
<col style="width: 20%" />
<col style="width: 20%" />
<col style="width: 20%" />
<col style="width: 20%" />
<col style="width: 20%" />
</colgroup>
<thead>
<tr>
<th class="tableblock halign-left valign-top"><a
href="#VkCommandBufferLevel">Command Buffer Levels</a></th>
<th class="tableblock halign-left valign-top"><a
href="#vkCmdBeginRenderPass">Render Pass Scope</a></th>
<th class="tableblock halign-left valign-top"><a
href="#vkCmdBeginVideoCodingKHR">Video Coding Scope</a></th>
<th class="tableblock halign-left valign-top"><a
href="#VkQueueFlagBits">Supported Queue Types</a></th>
<th class="tableblock halign-left valign-top"><a
href="#fundamentals-queueoperation-command-types">Command Type</a></th>
</tr>
</thead>
<tbody>
<tr>
<td class="tableblock halign-left valign-top"><p>Primary<br />
Secondary</p></td>
<td class="tableblock halign-left valign-top"><p>Both</p></td>
<td class="tableblock halign-left valign-top"><p>Both</p></td>
<td class="tableblock halign-left valign-top"><p>Graphics<br />
Compute<br />
Decode<br />
Encode</p></td>
<td class="tableblock halign-left valign-top"><p>Action<br />
State</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkQueryControlFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryControlFlags.html),
[VkQueryPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPool.html),
[vkCmdBeginQueryIndexedEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginQueryIndexedEXT.html),
[vkCmdEndQuery](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdEndQuery.html),
[vkCmdEndQueryIndexedEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdEndQueryIndexedEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdBeginQuery"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
