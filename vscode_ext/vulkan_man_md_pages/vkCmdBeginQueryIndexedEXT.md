# vkCmdBeginQueryIndexedEXT(3) Manual Page

## Name

vkCmdBeginQueryIndexedEXT - Begin an indexed query



## <a href="#_c_specification" class="anchor"></a>C Specification

To begin an indexed query, call:

``` c
// Provided by VK_EXT_transform_feedback
void vkCmdBeginQueryIndexedEXT(
    VkCommandBuffer                             commandBuffer,
    VkQueryPool                                 queryPool,
    uint32_t                                    query,
    VkQueryControlFlags                         flags,
    uint32_t                                    index);
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

- `index` is the query type specific index. When the query type is
  `VK_QUERY_TYPE_TRANSFORM_FEEDBACK_STREAM_EXT` or
  `VK_QUERY_TYPE_PRIMITIVES_GENERATED_EXT`, the index represents the
  vertex stream.

## <a href="#_description" class="anchor"></a>Description

The `vkCmdBeginQueryIndexedEXT` command operates the same as the
[vkCmdBeginQuery](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginQuery.html) command, except that it also
accepts a query type specific `index` parameter.

This command defines an execution dependency between other query
commands that reference the same query index.

The first <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies-scopes"
target="_blank" rel="noopener">synchronization scope</a> includes all
commands which reference the queries in `queryPool` indicated by `query`
and `index` that occur earlier in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-submission-order"
target="_blank" rel="noopener">submission order</a>.

The second <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies-scopes"
target="_blank" rel="noopener">synchronization scope</a> includes all
commands which reference the queries in `queryPool` indicated by `query`
and `index` that occur later in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-submission-order"
target="_blank" rel="noopener">submission order</a>.

The operation of this command happens after the first scope and happens
before the second scope.

Valid Usage

- <a href="#VUID-vkCmdBeginQueryIndexedEXT-None-00807"
  id="VUID-vkCmdBeginQueryIndexedEXT-None-00807"></a>
  VUID-vkCmdBeginQueryIndexedEXT-None-00807  
  All queries used by the command **must** be *unavailable*

- <a href="#VUID-vkCmdBeginQueryIndexedEXT-queryType-02804"
  id="VUID-vkCmdBeginQueryIndexedEXT-queryType-02804"></a>
  VUID-vkCmdBeginQueryIndexedEXT-queryType-02804  
  The `queryType` used to create `queryPool` **must** not be
  `VK_QUERY_TYPE_TIMESTAMP`

- <a href="#VUID-vkCmdBeginQueryIndexedEXT-queryType-04728"
  id="VUID-vkCmdBeginQueryIndexedEXT-queryType-04728"></a>
  VUID-vkCmdBeginQueryIndexedEXT-queryType-04728  
  The `queryType` used to create `queryPool` **must** not be
  `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_COMPACTED_SIZE_KHR` or
  `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_SERIALIZATION_SIZE_KHR`

- <a href="#VUID-vkCmdBeginQueryIndexedEXT-queryType-06741"
  id="VUID-vkCmdBeginQueryIndexedEXT-queryType-06741"></a>
  VUID-vkCmdBeginQueryIndexedEXT-queryType-06741  
  The `queryType` used to create `queryPool` **must** not be
  `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_SIZE_KHR` or
  `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_SERIALIZATION_BOTTOM_LEVEL_POINTERS_KHR`

- <a href="#VUID-vkCmdBeginQueryIndexedEXT-queryType-04729"
  id="VUID-vkCmdBeginQueryIndexedEXT-queryType-04729"></a>
  VUID-vkCmdBeginQueryIndexedEXT-queryType-04729  
  The `queryType` used to create `queryPool` **must** not be
  `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_COMPACTED_SIZE_NV`

- <a href="#VUID-vkCmdBeginQueryIndexedEXT-queryType-00800"
  id="VUID-vkCmdBeginQueryIndexedEXT-queryType-00800"></a>
  VUID-vkCmdBeginQueryIndexedEXT-queryType-00800  
  If the [`occlusionQueryPrecise`](#features-occlusionQueryPrecise)
  feature is not enabled, or the `queryType` used to create `queryPool`
  was not `VK_QUERY_TYPE_OCCLUSION`, `flags` **must** not contain
  `VK_QUERY_CONTROL_PRECISE_BIT`

- <a href="#VUID-vkCmdBeginQueryIndexedEXT-query-00802"
  id="VUID-vkCmdBeginQueryIndexedEXT-query-00802"></a>
  VUID-vkCmdBeginQueryIndexedEXT-query-00802  
  `query` **must** be less than the number of queries in `queryPool`

- <a href="#VUID-vkCmdBeginQueryIndexedEXT-queryType-00803"
  id="VUID-vkCmdBeginQueryIndexedEXT-queryType-00803"></a>
  VUID-vkCmdBeginQueryIndexedEXT-queryType-00803  
  If the `queryType` used to create `queryPool` was
  `VK_QUERY_TYPE_OCCLUSION`, the `VkCommandPool` that `commandBuffer`
  was allocated from **must** support graphics operations

- <a href="#VUID-vkCmdBeginQueryIndexedEXT-queryType-00804"
  id="VUID-vkCmdBeginQueryIndexedEXT-queryType-00804"></a>
  VUID-vkCmdBeginQueryIndexedEXT-queryType-00804  
  If the `queryType` used to create `queryPool` was
  `VK_QUERY_TYPE_PIPELINE_STATISTICS` and any of the
  `pipelineStatistics` indicate graphics operations, the `VkCommandPool`
  that `commandBuffer` was allocated from **must** support graphics
  operations

- <a href="#VUID-vkCmdBeginQueryIndexedEXT-queryType-00805"
  id="VUID-vkCmdBeginQueryIndexedEXT-queryType-00805"></a>
  VUID-vkCmdBeginQueryIndexedEXT-queryType-00805  
  If the `queryType` used to create `queryPool` was
  `VK_QUERY_TYPE_PIPELINE_STATISTICS` and any of the
  `pipelineStatistics` indicate compute operations, the `VkCommandPool`
  that `commandBuffer` was allocated from **must** support compute
  operations

- <a href="#VUID-vkCmdBeginQueryIndexedEXT-commandBuffer-01885"
  id="VUID-vkCmdBeginQueryIndexedEXT-commandBuffer-01885"></a>
  VUID-vkCmdBeginQueryIndexedEXT-commandBuffer-01885  
  `commandBuffer` **must** not be a protected command buffer

- <a href="#VUID-vkCmdBeginQueryIndexedEXT-query-00808"
  id="VUID-vkCmdBeginQueryIndexedEXT-query-00808"></a>
  VUID-vkCmdBeginQueryIndexedEXT-query-00808  
  If called within a render pass instance, the sum of `query` and the
  number of bits set in the current subpass’s view mask **must** be less
  than or equal to the number of queries in `queryPool`

- <a href="#VUID-vkCmdBeginQueryIndexedEXT-queryType-07126"
  id="VUID-vkCmdBeginQueryIndexedEXT-queryType-07126"></a>
  VUID-vkCmdBeginQueryIndexedEXT-queryType-07126  
  If the `queryType` used to create `queryPool` was
  `VK_QUERY_TYPE_RESULT_STATUS_ONLY_KHR`, then the `VkCommandPool` that
  `commandBuffer` was allocated from **must** have been created with a
  queue family index that supports [result status
  queries](#queries-result-status-only), as indicated by
  [VkQueueFamilyQueryResultStatusPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueueFamilyQueryResultStatusPropertiesKHR.html)::`queryResultStatusSupport`

- <a href="#VUID-vkCmdBeginQueryIndexedEXT-None-07127"
  id="VUID-vkCmdBeginQueryIndexedEXT-None-07127"></a>
  VUID-vkCmdBeginQueryIndexedEXT-None-07127  
  If there is a bound video session, then there **must** be no
  [active](#queries-operation-active) queries

- <a href="#VUID-vkCmdBeginQueryIndexedEXT-None-08370"
  id="VUID-vkCmdBeginQueryIndexedEXT-None-08370"></a>
  VUID-vkCmdBeginQueryIndexedEXT-None-08370  
  If there is a bound video session, then it **must** not have been
  created with `VK_VIDEO_SESSION_CREATE_INLINE_QUERIES_BIT_KHR`

- <a href="#VUID-vkCmdBeginQueryIndexedEXT-queryType-07128"
  id="VUID-vkCmdBeginQueryIndexedEXT-queryType-07128"></a>
  VUID-vkCmdBeginQueryIndexedEXT-queryType-07128  
  If the `queryType` used to create `queryPool` was
  `VK_QUERY_TYPE_RESULT_STATUS_ONLY_KHR` and there is a bound video
  session, then `queryPool` **must** have been created with a
  [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileInfoKHR.html) structure included
  in the `pNext` chain of
  [VkQueryPoolCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPoolCreateInfo.html) identical to the
  one specified in
  [VkVideoSessionCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionCreateInfoKHR.html)::`pVideoProfile`
  the bound video session was created with

- <a href="#VUID-vkCmdBeginQueryIndexedEXT-queryType-07129"
  id="VUID-vkCmdBeginQueryIndexedEXT-queryType-07129"></a>
  VUID-vkCmdBeginQueryIndexedEXT-queryType-07129  
  If the `queryType` used to create `queryPool` was
  `VK_QUERY_TYPE_VIDEO_ENCODE_FEEDBACK_KHR`, then there **must** be a
  bound video session

- <a href="#VUID-vkCmdBeginQueryIndexedEXT-queryType-07130"
  id="VUID-vkCmdBeginQueryIndexedEXT-queryType-07130"></a>
  VUID-vkCmdBeginQueryIndexedEXT-queryType-07130  
  If the `queryType` used to create `queryPool` was
  `VK_QUERY_TYPE_VIDEO_ENCODE_FEEDBACK_KHR` and there is a bound video
  session, then `queryPool` **must** have been created with a
  [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileInfoKHR.html) structure included
  in the `pNext` chain of
  [VkQueryPoolCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPoolCreateInfo.html) identical to the
  one specified in
  [VkVideoSessionCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionCreateInfoKHR.html)::`pVideoProfile`
  the bound video session was created with

- <a href="#VUID-vkCmdBeginQueryIndexedEXT-queryType-07131"
  id="VUID-vkCmdBeginQueryIndexedEXT-queryType-07131"></a>
  VUID-vkCmdBeginQueryIndexedEXT-queryType-07131  
  If the `queryType` used to create `queryPool` was not
  `VK_QUERY_TYPE_RESULT_STATUS_ONLY_KHR` or
  `VK_QUERY_TYPE_VIDEO_ENCODE_FEEDBACK_KHR`, then there **must** be no
  bound video session

- <a href="#VUID-vkCmdBeginQueryIndexedEXT-queryPool-04753"
  id="VUID-vkCmdBeginQueryIndexedEXT-queryPool-04753"></a>
  VUID-vkCmdBeginQueryIndexedEXT-queryPool-04753  
  If the `queryPool` was created with the same `queryType` as that of
  another <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#queries-operation-active"
  target="_blank" rel="noopener">active</a> query within
  `commandBuffer`, then `index` **must** not match the index used for
  the active query

- <a href="#VUID-vkCmdBeginQueryIndexedEXT-queryType-02338"
  id="VUID-vkCmdBeginQueryIndexedEXT-queryType-02338"></a>
  VUID-vkCmdBeginQueryIndexedEXT-queryType-02338  
  If the `queryType` used to create `queryPool` was
  `VK_QUERY_TYPE_TRANSFORM_FEEDBACK_STREAM_EXT` the `VkCommandPool` that
  `commandBuffer` was allocated from **must** support graphics
  operations

- <a href="#VUID-vkCmdBeginQueryIndexedEXT-queryType-02339"
  id="VUID-vkCmdBeginQueryIndexedEXT-queryType-02339"></a>
  VUID-vkCmdBeginQueryIndexedEXT-queryType-02339  
  If the `queryType` used to create `queryPool` was
  `VK_QUERY_TYPE_TRANSFORM_FEEDBACK_STREAM_EXT` the `index` parameter
  **must** be less than
  `VkPhysicalDeviceTransformFeedbackPropertiesEXT`::`maxTransformFeedbackStreams`

- <a href="#VUID-vkCmdBeginQueryIndexedEXT-queryType-06692"
  id="VUID-vkCmdBeginQueryIndexedEXT-queryType-06692"></a>
  VUID-vkCmdBeginQueryIndexedEXT-queryType-06692  
  If the `queryType` used to create `queryPool` was not
  `VK_QUERY_TYPE_TRANSFORM_FEEDBACK_STREAM_EXT` and not
  `VK_QUERY_TYPE_PRIMITIVES_GENERATED_EXT`, the `index` **must** be zero

- <a href="#VUID-vkCmdBeginQueryIndexedEXT-queryType-06689"
  id="VUID-vkCmdBeginQueryIndexedEXT-queryType-06689"></a>
  VUID-vkCmdBeginQueryIndexedEXT-queryType-06689  
  If the `queryType` used to create `queryPool` was
  `VK_QUERY_TYPE_PRIMITIVES_GENERATED_EXT` the `VkCommandPool` that
  `commandBuffer` was allocated from **must** support graphics
  operations

- <a href="#VUID-vkCmdBeginQueryIndexedEXT-queryType-06690"
  id="VUID-vkCmdBeginQueryIndexedEXT-queryType-06690"></a>
  VUID-vkCmdBeginQueryIndexedEXT-queryType-06690  
  If the `queryType` used to create `queryPool` was
  `VK_QUERY_TYPE_PRIMITIVES_GENERATED_EXT` the `index` parameter
  **must** be less than
  `VkPhysicalDeviceTransformFeedbackPropertiesEXT`::`maxTransformFeedbackStreams`

- <a href="#VUID-vkCmdBeginQueryIndexedEXT-queryType-06691"
  id="VUID-vkCmdBeginQueryIndexedEXT-queryType-06691"></a>
  VUID-vkCmdBeginQueryIndexedEXT-queryType-06691  
  If the `queryType` used to create `queryPool` was
  `VK_QUERY_TYPE_PRIMITIVES_GENERATED_EXT` and the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-primitivesGeneratedQueryWithNonZeroStreams"
  target="_blank"
  rel="noopener"><code>primitivesGeneratedQueryWithNonZeroStreams</code></a>
  feature is not enabled, the `index` parameter **must** be zero

- <a href="#VUID-vkCmdBeginQueryIndexedEXT-queryType-06693"
  id="VUID-vkCmdBeginQueryIndexedEXT-queryType-06693"></a>
  VUID-vkCmdBeginQueryIndexedEXT-queryType-06693  
  If the `queryType` used to create `queryPool` was
  `VK_QUERY_TYPE_PRIMITIVES_GENERATED_EXT` then <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-primitivesGeneratedQuery"
  target="_blank" rel="noopener"><code>primitivesGeneratedQuery</code></a>
  **must** be enabled

- <a href="#VUID-vkCmdBeginQueryIndexedEXT-queryType-02341"
  id="VUID-vkCmdBeginQueryIndexedEXT-queryType-02341"></a>
  VUID-vkCmdBeginQueryIndexedEXT-queryType-02341  
  If the `queryType` used to create `queryPool` was
  `VK_QUERY_TYPE_TRANSFORM_FEEDBACK_STREAM_EXT` then
  `VkPhysicalDeviceTransformFeedbackPropertiesEXT`::`transformFeedbackQueries`
  **must** be supported

- <a href="#VUID-vkCmdBeginQueryIndexedEXT-queryType-07071"
  id="VUID-vkCmdBeginQueryIndexedEXT-queryType-07071"></a>
  VUID-vkCmdBeginQueryIndexedEXT-queryType-07071  
  The `queryType` used to create `queryPool` **must** not be
  `VK_QUERY_TYPE_MESH_PRIMITIVES_GENERATED_EXT`

<!-- -->

- <a href="#VUID-vkCmdBeginQueryIndexedEXT-queryPool-07289"
  id="VUID-vkCmdBeginQueryIndexedEXT-queryPool-07289"></a>
  VUID-vkCmdBeginQueryIndexedEXT-queryPool-07289  
  If `queryPool` was created with a `queryType` of
  `VK_QUERY_TYPE_PERFORMANCE_QUERY_KHR`, then the
  [VkQueryPoolPerformanceCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPoolPerformanceCreateInfoKHR.html)::`queueFamilyIndex`
  `queryPool` was created with **must** equal the queue family index of
  the `VkCommandPool` that `commandBuffer` was allocated from

- <a href="#VUID-vkCmdBeginQueryIndexedEXT-queryPool-03223"
  id="VUID-vkCmdBeginQueryIndexedEXT-queryPool-03223"></a>
  VUID-vkCmdBeginQueryIndexedEXT-queryPool-03223  
  If `queryPool` was created with a `queryType` of
  `VK_QUERY_TYPE_PERFORMANCE_QUERY_KHR`, the [profiling
  lock](#profiling-lock) **must** have been held before
  [vkBeginCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkBeginCommandBuffer.html) was called on
  `commandBuffer`

- <a href="#VUID-vkCmdBeginQueryIndexedEXT-queryPool-03224"
  id="VUID-vkCmdBeginQueryIndexedEXT-queryPool-03224"></a>
  VUID-vkCmdBeginQueryIndexedEXT-queryPool-03224  
  If `queryPool` was created with a `queryType` of
  `VK_QUERY_TYPE_PERFORMANCE_QUERY_KHR` and one of the counters used to
  create `queryPool` was
  `VK_PERFORMANCE_COUNTER_SCOPE_COMMAND_BUFFER_KHR`, the query begin
  **must** be the first recorded command in `commandBuffer`

- <a href="#VUID-vkCmdBeginQueryIndexedEXT-queryPool-03225"
  id="VUID-vkCmdBeginQueryIndexedEXT-queryPool-03225"></a>
  VUID-vkCmdBeginQueryIndexedEXT-queryPool-03225  
  If `queryPool` was created with a `queryType` of
  `VK_QUERY_TYPE_PERFORMANCE_QUERY_KHR` and one of the counters used to
  create `queryPool` was `VK_PERFORMANCE_COUNTER_SCOPE_RENDER_PASS_KHR`,
  the begin command **must** not be recorded within a render pass
  instance

- <a href="#VUID-vkCmdBeginQueryIndexedEXT-queryPool-03226"
  id="VUID-vkCmdBeginQueryIndexedEXT-queryPool-03226"></a>
  VUID-vkCmdBeginQueryIndexedEXT-queryPool-03226  
  If `queryPool` was created with a `queryType` of
  `VK_QUERY_TYPE_PERFORMANCE_QUERY_KHR` and another query pool with a
  `queryType` `VK_QUERY_TYPE_PERFORMANCE_QUERY_KHR` has been used within
  `commandBuffer`, its parent primary command buffer or secondary
  command buffer recorded within the same parent primary command buffer
  as `commandBuffer`, the
  [`performanceCounterMultipleQueryPools`](#features-performanceCounterMultipleQueryPools)
  feature **must** be enabled

- <a href="#VUID-vkCmdBeginQueryIndexedEXT-None-02863"
  id="VUID-vkCmdBeginQueryIndexedEXT-None-02863"></a>
  VUID-vkCmdBeginQueryIndexedEXT-None-02863  
  If `queryPool` was created with a `queryType` of
  `VK_QUERY_TYPE_PERFORMANCE_QUERY_KHR`, this command **must** not be
  recorded in a command buffer that, either directly or through
  secondary command buffers, also contains a `vkCmdResetQueryPool`
  command affecting the same query

Valid Usage (Implicit)

- <a href="#VUID-vkCmdBeginQueryIndexedEXT-commandBuffer-parameter"
  id="VUID-vkCmdBeginQueryIndexedEXT-commandBuffer-parameter"></a>
  VUID-vkCmdBeginQueryIndexedEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdBeginQueryIndexedEXT-queryPool-parameter"
  id="VUID-vkCmdBeginQueryIndexedEXT-queryPool-parameter"></a>
  VUID-vkCmdBeginQueryIndexedEXT-queryPool-parameter  
  `queryPool` **must** be a valid [VkQueryPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPool.html) handle

- <a href="#VUID-vkCmdBeginQueryIndexedEXT-flags-parameter"
  id="VUID-vkCmdBeginQueryIndexedEXT-flags-parameter"></a>
  VUID-vkCmdBeginQueryIndexedEXT-flags-parameter  
  `flags` **must** be a valid combination of
  [VkQueryControlFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryControlFlagBits.html) values

- <a href="#VUID-vkCmdBeginQueryIndexedEXT-commandBuffer-recording"
  id="VUID-vkCmdBeginQueryIndexedEXT-commandBuffer-recording"></a>
  VUID-vkCmdBeginQueryIndexedEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdBeginQueryIndexedEXT-commandBuffer-cmdpool"
  id="VUID-vkCmdBeginQueryIndexedEXT-commandBuffer-cmdpool"></a>
  VUID-vkCmdBeginQueryIndexedEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics, compute, decode, or encode operations

- <a href="#VUID-vkCmdBeginQueryIndexedEXT-videocoding"
  id="VUID-vkCmdBeginQueryIndexedEXT-videocoding"></a>
  VUID-vkCmdBeginQueryIndexedEXT-videocoding  
  This command **must** only be called outside of a video coding scope

- <a href="#VUID-vkCmdBeginQueryIndexedEXT-commonparent"
  id="VUID-vkCmdBeginQueryIndexedEXT-commonparent"></a>
  VUID-vkCmdBeginQueryIndexedEXT-commonparent  
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
<td class="tableblock halign-left valign-top"><p>Outside</p></td>
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

[VK_EXT_transform_feedback](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_transform_feedback.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkQueryControlFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryControlFlags.html),
[VkQueryPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPool.html),
[vkCmdBeginQuery](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginQuery.html),
[vkCmdEndQuery](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdEndQuery.html),
[vkCmdEndQueryIndexedEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdEndQueryIndexedEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdBeginQueryIndexedEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
