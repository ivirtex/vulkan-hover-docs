# vkCmdWriteTimestamp2(3) Manual Page

## Name

vkCmdWriteTimestamp2 - Write a device timestamp into a query object



## <a href="#_c_specification" class="anchor"></a>C Specification

To request a timestamp and write the value to memory, call:

``` c
// Provided by VK_VERSION_1_3
void vkCmdWriteTimestamp2(
    VkCommandBuffer                             commandBuffer,
    VkPipelineStageFlags2                       stage,
    VkQueryPool                                 queryPool,
    uint32_t                                    query);
```

or the equivalent command

``` c
// Provided by VK_KHR_synchronization2
void vkCmdWriteTimestamp2KHR(
    VkCommandBuffer                             commandBuffer,
    VkPipelineStageFlags2                       stage,
    VkQueryPool                                 queryPool,
    uint32_t                                    query);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `stage` specifies a stage of the pipeline.

- `queryPool` is the query pool that will manage the timestamp.

- `query` is the query within the query pool that will contain the
  timestamp.

## <a href="#_description" class="anchor"></a>Description

When `vkCmdWriteTimestamp2` is submitted to a queue, it defines an
execution dependency on commands that were submitted before it, and
writes a timestamp to a query pool.

The first <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies-scopes"
target="_blank" rel="noopener">synchronization scope</a> includes all
commands that occur earlier in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-submission-order"
target="_blank" rel="noopener">submission order</a>. The synchronization
scope is limited to operations on the pipeline stage specified by
`stage`.

The second <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies-scopes"
target="_blank" rel="noopener">synchronization scope</a> includes only
the timestamp write operation.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>Implementations may write the timestamp at any stage that is <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-pipeline-stages-order"
target="_blank" rel="noopener">logically later</a> than
<code>stage</code>.</p></td>
</tr>
</tbody>
</table>

Any timestamp write that <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies-execution"
target="_blank" rel="noopener">happens-after</a> another timestamp write
in the same submission **must** not have a lower value unless its value
overflows the maximum supported integer bit width of the query. If
[`VK_KHR_calibrated_timestamps`](VK_KHR_calibrated_timestamps.html) or
[`VK_EXT_calibrated_timestamps`](VK_EXT_calibrated_timestamps.html) is
enabled, this extends to timestamp writes across all submissions on the
same logical device: any timestamp write that <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies-execution"
target="_blank" rel="noopener">happens-after</a> another **must** not
have a lower value unless its value overflows the maximum supported
integer bit width of the query. Timestamps written by this command
**must** be in the `VK_TIME_DOMAIN_DEVICE_KHR`
<a href="VkTimeDomainKHR.html" target="_blank" rel="noopener">time
domain</a>. If an overflow occurs, the timestamp value **must** wrap
back to zero.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>Comparisons between timestamps should be done between timestamps
where they are guaranteed to not decrease. For example, subtracting an
older timestamp from a newer one to determine the execution time of a
sequence of commands is only a reliable measurement if the two timestamp
writes were performed in the same submission, or if the writes were
performed on the same logical device and <a
href="VK_KHR_calibrated_timestamps.html"><code>VK_KHR_calibrated_timestamps</code></a>
or <a
href="VK_EXT_calibrated_timestamps.html"><code>VK_EXT_calibrated_timestamps</code></a>
is enabled.</p></td>
</tr>
</tbody>
</table>

If `vkCmdWriteTimestamp2` is called while executing a render pass
instance that has multiview enabled, the timestamp uses N consecutive
query indices in the query pool (starting at `query`) where N is the
number of bits set in the view mask of the subpass the command is
executed in. The resulting query values are determined by an
implementation-dependent choice of one of the following behaviors:

- The first query is a timestamp value and (if more than one bit is set
  in the view mask) zero is written to the remaining queries. If two
  timestamps are written in the same subpass, the sum of the execution
  time of all views between those commands is the difference between the
  first query written by each command.

- All N queries are timestamp values. If two timestamps are written in
  the same subpass, the sum of the execution time of all views between
  those commands is the sum of the difference between corresponding
  queries written by each command. The difference between corresponding
  queries **may** be the execution time of a single view.

In either case, the application **can** sum the differences between all
N queries to determine the total execution time.

Valid Usage

- <a href="#VUID-vkCmdWriteTimestamp2-stage-03929"
  id="VUID-vkCmdWriteTimestamp2-stage-03929"></a>
  VUID-vkCmdWriteTimestamp2-stage-03929  
  If the [`geometryShader`](#features-geometryShader) feature is not
  enabled, `stage` **must** not contain
  `VK_PIPELINE_STAGE_2_GEOMETRY_SHADER_BIT`

- <a href="#VUID-vkCmdWriteTimestamp2-stage-03930"
  id="VUID-vkCmdWriteTimestamp2-stage-03930"></a>
  VUID-vkCmdWriteTimestamp2-stage-03930  
  If the [`tessellationShader`](#features-tessellationShader) feature is
  not enabled, `stage` **must** not contain
  `VK_PIPELINE_STAGE_2_TESSELLATION_CONTROL_SHADER_BIT` or
  `VK_PIPELINE_STAGE_2_TESSELLATION_EVALUATION_SHADER_BIT`

- <a href="#VUID-vkCmdWriteTimestamp2-stage-03931"
  id="VUID-vkCmdWriteTimestamp2-stage-03931"></a>
  VUID-vkCmdWriteTimestamp2-stage-03931  
  If the [`conditionalRendering`](#features-conditionalRendering)
  feature is not enabled, `stage` **must** not contain
  `VK_PIPELINE_STAGE_2_CONDITIONAL_RENDERING_BIT_EXT`

- <a href="#VUID-vkCmdWriteTimestamp2-stage-03932"
  id="VUID-vkCmdWriteTimestamp2-stage-03932"></a>
  VUID-vkCmdWriteTimestamp2-stage-03932  
  If the [`fragmentDensityMap`](#features-fragmentDensityMap) feature is
  not enabled, `stage` **must** not contain
  `VK_PIPELINE_STAGE_2_FRAGMENT_DENSITY_PROCESS_BIT_EXT`

- <a href="#VUID-vkCmdWriteTimestamp2-stage-03933"
  id="VUID-vkCmdWriteTimestamp2-stage-03933"></a>
  VUID-vkCmdWriteTimestamp2-stage-03933  
  If the [`transformFeedback`](#features-transformFeedback) feature is
  not enabled, `stage` **must** not contain
  `VK_PIPELINE_STAGE_2_TRANSFORM_FEEDBACK_BIT_EXT`

- <a href="#VUID-vkCmdWriteTimestamp2-stage-03934"
  id="VUID-vkCmdWriteTimestamp2-stage-03934"></a>
  VUID-vkCmdWriteTimestamp2-stage-03934  
  If the [`meshShader`](#features-meshShader) feature is not enabled,
  `stage` **must** not contain `VK_PIPELINE_STAGE_2_MESH_SHADER_BIT_EXT`

- <a href="#VUID-vkCmdWriteTimestamp2-stage-03935"
  id="VUID-vkCmdWriteTimestamp2-stage-03935"></a>
  VUID-vkCmdWriteTimestamp2-stage-03935  
  If the [`taskShader`](#features-taskShader) feature is not enabled,
  `stage` **must** not contain `VK_PIPELINE_STAGE_2_TASK_SHADER_BIT_EXT`

- <a href="#VUID-vkCmdWriteTimestamp2-stage-07316"
  id="VUID-vkCmdWriteTimestamp2-stage-07316"></a>
  VUID-vkCmdWriteTimestamp2-stage-07316  
  If neither the [`shadingRateImage`](#features-shadingRateImage) or
  [`attachmentFragmentShadingRate`](#features-attachmentFragmentShadingRate)
  are enabled, `stage` **must** not contain
  `VK_PIPELINE_STAGE_2_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR`

- <a href="#VUID-vkCmdWriteTimestamp2-stage-04957"
  id="VUID-vkCmdWriteTimestamp2-stage-04957"></a>
  VUID-vkCmdWriteTimestamp2-stage-04957  
  If the [`subpassShading`](#features-subpassShading) feature is not
  enabled, `stage` **must** not contain
  `VK_PIPELINE_STAGE_2_SUBPASS_SHADER_BIT_HUAWEI`

- <a href="#VUID-vkCmdWriteTimestamp2-stage-04995"
  id="VUID-vkCmdWriteTimestamp2-stage-04995"></a>
  VUID-vkCmdWriteTimestamp2-stage-04995  
  If the [`invocationMask`](#features-invocationMask) feature is not
  enabled, `stage` **must** not contain
  `VK_PIPELINE_STAGE_2_INVOCATION_MASK_BIT_HUAWEI`

- <a href="#VUID-vkCmdWriteTimestamp2-stage-07946"
  id="VUID-vkCmdWriteTimestamp2-stage-07946"></a>
  VUID-vkCmdWriteTimestamp2-stage-07946  
  If neither the [VK_NV_ray_tracing](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_ray_tracing.html) extension
  or [`rayTracingPipeline` feature](#features-rayTracingPipeline) are
  enabled, `stage` **must** not contain
  `VK_PIPELINE_STAGE_2_RAY_TRACING_SHADER_BIT_KHR`

- <a href="#VUID-vkCmdWriteTimestamp2-synchronization2-03858"
  id="VUID-vkCmdWriteTimestamp2-synchronization2-03858"></a>
  VUID-vkCmdWriteTimestamp2-synchronization2-03858  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-synchronization2"
  target="_blank" rel="noopener"><code>synchronization2</code></a>
  feature **must** be enabled

- <a href="#VUID-vkCmdWriteTimestamp2-stage-03859"
  id="VUID-vkCmdWriteTimestamp2-stage-03859"></a>
  VUID-vkCmdWriteTimestamp2-stage-03859  
  `stage` **must** only include a single pipeline stage

- <a href="#VUID-vkCmdWriteTimestamp2-stage-03860"
  id="VUID-vkCmdWriteTimestamp2-stage-03860"></a>
  VUID-vkCmdWriteTimestamp2-stage-03860  
  `stage` **must** only include stages valid for the queue family that
  was used to create the command pool that `commandBuffer` was allocated
  from

- <a href="#VUID-vkCmdWriteTimestamp2-queryPool-03861"
  id="VUID-vkCmdWriteTimestamp2-queryPool-03861"></a>
  VUID-vkCmdWriteTimestamp2-queryPool-03861  
  `queryPool` **must** have been created with a `queryType` of
  `VK_QUERY_TYPE_TIMESTAMP`

- <a href="#VUID-vkCmdWriteTimestamp2-timestampValidBits-03863"
  id="VUID-vkCmdWriteTimestamp2-timestampValidBits-03863"></a>
  VUID-vkCmdWriteTimestamp2-timestampValidBits-03863  
  The command pool’s queue family **must** support a non-zero
  `timestampValidBits`

- <a href="#VUID-vkCmdWriteTimestamp2-query-04903"
  id="VUID-vkCmdWriteTimestamp2-query-04903"></a>
  VUID-vkCmdWriteTimestamp2-query-04903  
  `query` **must** be less than the number of queries in `queryPool`

- <a href="#VUID-vkCmdWriteTimestamp2-None-03864"
  id="VUID-vkCmdWriteTimestamp2-None-03864"></a>
  VUID-vkCmdWriteTimestamp2-None-03864  
  All queries used by the command **must** be *unavailable*

- <a href="#VUID-vkCmdWriteTimestamp2-query-03865"
  id="VUID-vkCmdWriteTimestamp2-query-03865"></a>
  VUID-vkCmdWriteTimestamp2-query-03865  
  If `vkCmdWriteTimestamp2` is called within a render pass instance, the
  sum of `query` and the number of bits set in the current subpass’s
  view mask **must** be less than or equal to the number of queries in
  `queryPool`

Valid Usage (Implicit)

- <a href="#VUID-vkCmdWriteTimestamp2-commandBuffer-parameter"
  id="VUID-vkCmdWriteTimestamp2-commandBuffer-parameter"></a>
  VUID-vkCmdWriteTimestamp2-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdWriteTimestamp2-stage-parameter"
  id="VUID-vkCmdWriteTimestamp2-stage-parameter"></a>
  VUID-vkCmdWriteTimestamp2-stage-parameter  
  `stage` **must** be a valid combination of
  [VkPipelineStageFlagBits2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlagBits2.html) values

- <a href="#VUID-vkCmdWriteTimestamp2-queryPool-parameter"
  id="VUID-vkCmdWriteTimestamp2-queryPool-parameter"></a>
  VUID-vkCmdWriteTimestamp2-queryPool-parameter  
  `queryPool` **must** be a valid [VkQueryPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPool.html) handle

- <a href="#VUID-vkCmdWriteTimestamp2-commandBuffer-recording"
  id="VUID-vkCmdWriteTimestamp2-commandBuffer-recording"></a>
  VUID-vkCmdWriteTimestamp2-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdWriteTimestamp2-commandBuffer-cmdpool"
  id="VUID-vkCmdWriteTimestamp2-commandBuffer-cmdpool"></a>
  VUID-vkCmdWriteTimestamp2-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support transfer, graphics, compute, decode, or encode operations

- <a href="#VUID-vkCmdWriteTimestamp2-commonparent"
  id="VUID-vkCmdWriteTimestamp2-commonparent"></a>
  VUID-vkCmdWriteTimestamp2-commonparent  
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
<tr class="header">
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
<tr class="odd">
<td class="tableblock halign-left valign-top"><p>Primary<br />
Secondary</p></td>
<td class="tableblock halign-left valign-top"><p>Both</p></td>
<td class="tableblock halign-left valign-top"><p>Both</p></td>
<td class="tableblock halign-left valign-top"><p>Transfer<br />
Graphics<br />
Compute<br />
Decode<br />
Encode</p></td>
<td class="tableblock halign-left valign-top"><p>Action</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_synchronization2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_synchronization2.html),
[VK_VERSION_1_3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_3.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkPipelineStageFlags2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlags2.html),
[VkQueryPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPool.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdWriteTimestamp2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
