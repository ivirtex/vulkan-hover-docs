# vkCmdWaitEvents(3) Manual Page

## Name

vkCmdWaitEvents - Wait for one or more events and insert a set of memory



## <a href="#_c_specification" class="anchor"></a>C Specification

To wait for one or more events to enter the signaled state on a device,
call:

``` c
// Provided by VK_VERSION_1_0
void vkCmdWaitEvents(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    eventCount,
    const VkEvent*                              pEvents,
    VkPipelineStageFlags                        srcStageMask,
    VkPipelineStageFlags                        dstStageMask,
    uint32_t                                    memoryBarrierCount,
    const VkMemoryBarrier*                      pMemoryBarriers,
    uint32_t                                    bufferMemoryBarrierCount,
    const VkBufferMemoryBarrier*                pBufferMemoryBarriers,
    uint32_t                                    imageMemoryBarrierCount,
    const VkImageMemoryBarrier*                 pImageMemoryBarriers);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command is
  recorded.

- `eventCount` is the length of the `pEvents` array.

- `pEvents` is a pointer to an array of event object handles to wait on.

- `srcStageMask` is a bitmask of
  [VkPipelineStageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlagBits.html) specifying the
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-pipeline-stages"
  target="_blank" rel="noopener">source stage mask</a>.

- `dstStageMask` is a bitmask of
  [VkPipelineStageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlagBits.html) specifying the
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-pipeline-stages"
  target="_blank" rel="noopener">destination stage mask</a>.

- `memoryBarrierCount` is the length of the `pMemoryBarriers` array.

- `pMemoryBarriers` is a pointer to an array of
  [VkMemoryBarrier](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryBarrier.html) structures.

- `bufferMemoryBarrierCount` is the length of the
  `pBufferMemoryBarriers` array.

- `pBufferMemoryBarriers` is a pointer to an array of
  [VkBufferMemoryBarrier](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferMemoryBarrier.html) structures.

- `imageMemoryBarrierCount` is the length of the `pImageMemoryBarriers`
  array.

- `pImageMemoryBarriers` is a pointer to an array of
  [VkImageMemoryBarrier](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageMemoryBarrier.html) structures.

## <a href="#_description" class="anchor"></a>Description

`vkCmdWaitEvents` is largely similar to
[vkCmdWaitEvents2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdWaitEvents2.html), but **can** only wait on
signal operations defined by [vkCmdSetEvent](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetEvent.html). As
[vkCmdSetEvent](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetEvent.html) does not define any access scopes,
`vkCmdWaitEvents` defines the first access scope for each event signal
operation in addition to its own access scopes.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>Since <a href="https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetEvent.html">vkCmdSetEvent</a> does not have
any dependency information beyond a stage mask, implementations do not
have the same opportunity to perform <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies-available-and-visible"
target="_blank" rel="noopener">availability and visibility
operations</a> or <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-image-layout-transitions"
target="_blank" rel="noopener">image layout transitions</a> in advance
as they do with <a href="https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetEvent2.html">vkCmdSetEvent2</a> and <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdWaitEvents2.html">vkCmdWaitEvents2</a>.</p></td>
</tr>
</tbody>
</table>

When `vkCmdWaitEvents` is submitted to a queue, it defines a memory
dependency between prior event signal operations on the same queue or
the host, and subsequent commands. `vkCmdWaitEvents` **must** not be
used to wait on event signal operations occurring on other queues.

The first synchronization scope only includes event signal operations
that operate on members of `pEvents`, and the operations that
happened-before the event signal operations. Event signal operations
performed by [vkCmdSetEvent](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetEvent.html) that occur earlier in
<a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-submission-order"
target="_blank" rel="noopener">submission order</a> are included in the
first synchronization scope, if the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-pipeline-stages-order"
target="_blank" rel="noopener">logically latest</a> pipeline stage in
their `stageMask` parameter is <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-pipeline-stages-order"
target="_blank" rel="noopener">logically earlier</a> than or equal to
the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-pipeline-stages-order"
target="_blank" rel="noopener">logically latest</a> pipeline stage in
`srcStageMask`. Event signal operations performed by
[vkSetEvent](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkSetEvent.html) are only included in the first
synchronization scope if `VK_PIPELINE_STAGE_HOST_BIT` is included in
`srcStageMask`.

The second <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies-scopes"
target="_blank" rel="noopener">synchronization scope</a> includes all
commands that occur later in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-submission-order"
target="_blank" rel="noopener">submission order</a>. The second
synchronization scope is limited to operations on the pipeline stages
determined by the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-pipeline-stages-masks"
target="_blank" rel="noopener">destination stage mask</a> specified by
`dstStageMask`.

The first <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies-access-scopes"
target="_blank" rel="noopener">access scope</a> is limited to accesses
in the pipeline stages determined by the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-pipeline-stages-masks"
target="_blank" rel="noopener">source stage mask</a> specified by
`srcStageMask`. Within that, the first access scope only includes the
first access scopes defined by elements of the `pMemoryBarriers`,
`pBufferMemoryBarriers` and `pImageMemoryBarriers` arrays, which each
define a set of <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-memory-barriers"
target="_blank" rel="noopener">memory barriers</a>. If no memory
barriers are specified, then the first access scope includes no
accesses.

The second <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies-access-scopes"
target="_blank" rel="noopener">access scope</a> is limited to accesses
in the pipeline stages determined by the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-pipeline-stages-masks"
target="_blank" rel="noopener">destination stage mask</a> specified by
`dstStageMask`. Within that, the second access scope only includes the
second access scopes defined by elements of the `pMemoryBarriers`,
`pBufferMemoryBarriers` and `pImageMemoryBarriers` arrays, which each
define a set of <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-memory-barriers"
target="_blank" rel="noopener">memory barriers</a>. If no memory
barriers are specified, then the second access scope includes no
accesses.

Valid Usage

- <a href="#VUID-vkCmdWaitEvents-srcStageMask-04090"
  id="VUID-vkCmdWaitEvents-srcStageMask-04090"></a>
  VUID-vkCmdWaitEvents-srcStageMask-04090  
  If the [`geometryShader`](#features-geometryShader) feature is not
  enabled, `srcStageMask` **must** not contain
  `VK_PIPELINE_STAGE_GEOMETRY_SHADER_BIT`

- <a href="#VUID-vkCmdWaitEvents-srcStageMask-04091"
  id="VUID-vkCmdWaitEvents-srcStageMask-04091"></a>
  VUID-vkCmdWaitEvents-srcStageMask-04091  
  If the [`tessellationShader`](#features-tessellationShader) feature is
  not enabled, `srcStageMask` **must** not contain
  `VK_PIPELINE_STAGE_TESSELLATION_CONTROL_SHADER_BIT` or
  `VK_PIPELINE_STAGE_TESSELLATION_EVALUATION_SHADER_BIT`

- <a href="#VUID-vkCmdWaitEvents-srcStageMask-04092"
  id="VUID-vkCmdWaitEvents-srcStageMask-04092"></a>
  VUID-vkCmdWaitEvents-srcStageMask-04092  
  If the [`conditionalRendering`](#features-conditionalRendering)
  feature is not enabled, `srcStageMask` **must** not contain
  `VK_PIPELINE_STAGE_CONDITIONAL_RENDERING_BIT_EXT`

- <a href="#VUID-vkCmdWaitEvents-srcStageMask-04093"
  id="VUID-vkCmdWaitEvents-srcStageMask-04093"></a>
  VUID-vkCmdWaitEvents-srcStageMask-04093  
  If the [`fragmentDensityMap`](#features-fragmentDensityMap) feature is
  not enabled, `srcStageMask` **must** not contain
  `VK_PIPELINE_STAGE_FRAGMENT_DENSITY_PROCESS_BIT_EXT`

- <a href="#VUID-vkCmdWaitEvents-srcStageMask-04094"
  id="VUID-vkCmdWaitEvents-srcStageMask-04094"></a>
  VUID-vkCmdWaitEvents-srcStageMask-04094  
  If the [`transformFeedback`](#features-transformFeedback) feature is
  not enabled, `srcStageMask` **must** not contain
  `VK_PIPELINE_STAGE_TRANSFORM_FEEDBACK_BIT_EXT`

- <a href="#VUID-vkCmdWaitEvents-srcStageMask-04095"
  id="VUID-vkCmdWaitEvents-srcStageMask-04095"></a>
  VUID-vkCmdWaitEvents-srcStageMask-04095  
  If the [`meshShader`](#features-meshShader) feature is not enabled,
  `srcStageMask` **must** not contain
  `VK_PIPELINE_STAGE_MESH_SHADER_BIT_EXT`

- <a href="#VUID-vkCmdWaitEvents-srcStageMask-04096"
  id="VUID-vkCmdWaitEvents-srcStageMask-04096"></a>
  VUID-vkCmdWaitEvents-srcStageMask-04096  
  If the [`taskShader`](#features-taskShader) feature is not enabled,
  `srcStageMask` **must** not contain
  `VK_PIPELINE_STAGE_TASK_SHADER_BIT_EXT`

- <a href="#VUID-vkCmdWaitEvents-srcStageMask-07318"
  id="VUID-vkCmdWaitEvents-srcStageMask-07318"></a>
  VUID-vkCmdWaitEvents-srcStageMask-07318  
  If neither the [`shadingRateImage`](#features-shadingRateImage) or
  [`attachmentFragmentShadingRate`](#features-attachmentFragmentShadingRate)
  are enabled, `srcStageMask` **must** not contain
  `VK_PIPELINE_STAGE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR`

- <a href="#VUID-vkCmdWaitEvents-srcStageMask-03937"
  id="VUID-vkCmdWaitEvents-srcStageMask-03937"></a>
  VUID-vkCmdWaitEvents-srcStageMask-03937  
  If the [`synchronization2`](#features-synchronization2) feature is not
  enabled, `srcStageMask` **must** not be `0`

- <a href="#VUID-vkCmdWaitEvents-srcStageMask-07949"
  id="VUID-vkCmdWaitEvents-srcStageMask-07949"></a>
  VUID-vkCmdWaitEvents-srcStageMask-07949  
  If neither the [VK_NV_ray_tracing](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_ray_tracing.html) extension
  or [`rayTracingPipeline` feature](#features-rayTracingPipeline) are
  enabled, `srcStageMask` **must** not contain
  `VK_PIPELINE_STAGE_RAY_TRACING_SHADER_BIT_KHR`

<!-- -->

- <a href="#VUID-vkCmdWaitEvents-srcAccessMask-06257"
  id="VUID-vkCmdWaitEvents-srcAccessMask-06257"></a>
  VUID-vkCmdWaitEvents-srcAccessMask-06257  
  If the [`rayQuery`](#features-rayQuery) feature is not enabled and a
  memory barrier `srcAccessMask` includes
  `VK_ACCESS_ACCELERATION_STRUCTURE_READ_BIT_KHR`, `srcStageMask`
  **must** not include any of the `VK_PIPELINE_STAGE_*_SHADER_BIT`
  stages except `VK_PIPELINE_STAGE_RAY_TRACING_SHADER_BIT_KHR`

<!-- -->

- <a href="#VUID-vkCmdWaitEvents-dstStageMask-04090"
  id="VUID-vkCmdWaitEvents-dstStageMask-04090"></a>
  VUID-vkCmdWaitEvents-dstStageMask-04090  
  If the [`geometryShader`](#features-geometryShader) feature is not
  enabled, `dstStageMask` **must** not contain
  `VK_PIPELINE_STAGE_GEOMETRY_SHADER_BIT`

- <a href="#VUID-vkCmdWaitEvents-dstStageMask-04091"
  id="VUID-vkCmdWaitEvents-dstStageMask-04091"></a>
  VUID-vkCmdWaitEvents-dstStageMask-04091  
  If the [`tessellationShader`](#features-tessellationShader) feature is
  not enabled, `dstStageMask` **must** not contain
  `VK_PIPELINE_STAGE_TESSELLATION_CONTROL_SHADER_BIT` or
  `VK_PIPELINE_STAGE_TESSELLATION_EVALUATION_SHADER_BIT`

- <a href="#VUID-vkCmdWaitEvents-dstStageMask-04092"
  id="VUID-vkCmdWaitEvents-dstStageMask-04092"></a>
  VUID-vkCmdWaitEvents-dstStageMask-04092  
  If the [`conditionalRendering`](#features-conditionalRendering)
  feature is not enabled, `dstStageMask` **must** not contain
  `VK_PIPELINE_STAGE_CONDITIONAL_RENDERING_BIT_EXT`

- <a href="#VUID-vkCmdWaitEvents-dstStageMask-04093"
  id="VUID-vkCmdWaitEvents-dstStageMask-04093"></a>
  VUID-vkCmdWaitEvents-dstStageMask-04093  
  If the [`fragmentDensityMap`](#features-fragmentDensityMap) feature is
  not enabled, `dstStageMask` **must** not contain
  `VK_PIPELINE_STAGE_FRAGMENT_DENSITY_PROCESS_BIT_EXT`

- <a href="#VUID-vkCmdWaitEvents-dstStageMask-04094"
  id="VUID-vkCmdWaitEvents-dstStageMask-04094"></a>
  VUID-vkCmdWaitEvents-dstStageMask-04094  
  If the [`transformFeedback`](#features-transformFeedback) feature is
  not enabled, `dstStageMask` **must** not contain
  `VK_PIPELINE_STAGE_TRANSFORM_FEEDBACK_BIT_EXT`

- <a href="#VUID-vkCmdWaitEvents-dstStageMask-04095"
  id="VUID-vkCmdWaitEvents-dstStageMask-04095"></a>
  VUID-vkCmdWaitEvents-dstStageMask-04095  
  If the [`meshShader`](#features-meshShader) feature is not enabled,
  `dstStageMask` **must** not contain
  `VK_PIPELINE_STAGE_MESH_SHADER_BIT_EXT`

- <a href="#VUID-vkCmdWaitEvents-dstStageMask-04096"
  id="VUID-vkCmdWaitEvents-dstStageMask-04096"></a>
  VUID-vkCmdWaitEvents-dstStageMask-04096  
  If the [`taskShader`](#features-taskShader) feature is not enabled,
  `dstStageMask` **must** not contain
  `VK_PIPELINE_STAGE_TASK_SHADER_BIT_EXT`

- <a href="#VUID-vkCmdWaitEvents-dstStageMask-07318"
  id="VUID-vkCmdWaitEvents-dstStageMask-07318"></a>
  VUID-vkCmdWaitEvents-dstStageMask-07318  
  If neither the [`shadingRateImage`](#features-shadingRateImage) or
  [`attachmentFragmentShadingRate`](#features-attachmentFragmentShadingRate)
  are enabled, `dstStageMask` **must** not contain
  `VK_PIPELINE_STAGE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR`

- <a href="#VUID-vkCmdWaitEvents-dstStageMask-03937"
  id="VUID-vkCmdWaitEvents-dstStageMask-03937"></a>
  VUID-vkCmdWaitEvents-dstStageMask-03937  
  If the [`synchronization2`](#features-synchronization2) feature is not
  enabled, `dstStageMask` **must** not be `0`

- <a href="#VUID-vkCmdWaitEvents-dstStageMask-07949"
  id="VUID-vkCmdWaitEvents-dstStageMask-07949"></a>
  VUID-vkCmdWaitEvents-dstStageMask-07949  
  If neither the [VK_NV_ray_tracing](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_ray_tracing.html) extension
  or [`rayTracingPipeline` feature](#features-rayTracingPipeline) are
  enabled, `dstStageMask` **must** not contain
  `VK_PIPELINE_STAGE_RAY_TRACING_SHADER_BIT_KHR`

<!-- -->

- <a href="#VUID-vkCmdWaitEvents-dstAccessMask-06257"
  id="VUID-vkCmdWaitEvents-dstAccessMask-06257"></a>
  VUID-vkCmdWaitEvents-dstAccessMask-06257  
  If the [`rayQuery`](#features-rayQuery) feature is not enabled and a
  memory barrier `dstAccessMask` includes
  `VK_ACCESS_ACCELERATION_STRUCTURE_READ_BIT_KHR`, `dstStageMask`
  **must** not include any of the `VK_PIPELINE_STAGE_*_SHADER_BIT`
  stages except `VK_PIPELINE_STAGE_RAY_TRACING_SHADER_BIT_KHR`

<!-- -->

- <a href="#VUID-vkCmdWaitEvents-srcAccessMask-02815"
  id="VUID-vkCmdWaitEvents-srcAccessMask-02815"></a>
  VUID-vkCmdWaitEvents-srcAccessMask-02815  
  The `srcAccessMask` member of each element of `pMemoryBarriers`
  **must** only include access flags that are supported by one or more
  of the pipeline stages in `srcStageMask`, as specified in the [table
  of supported access types](#synchronization-access-types-supported)

- <a href="#VUID-vkCmdWaitEvents-dstAccessMask-02816"
  id="VUID-vkCmdWaitEvents-dstAccessMask-02816"></a>
  VUID-vkCmdWaitEvents-dstAccessMask-02816  
  The `dstAccessMask` member of each element of `pMemoryBarriers`
  **must** only include access flags that are supported by one or more
  of the pipeline stages in `dstStageMask`, as specified in the [table
  of supported access types](#synchronization-access-types-supported)

- <a href="#VUID-vkCmdWaitEvents-pBufferMemoryBarriers-02817"
  id="VUID-vkCmdWaitEvents-pBufferMemoryBarriers-02817"></a>
  VUID-vkCmdWaitEvents-pBufferMemoryBarriers-02817  
  For any element of `pBufferMemoryBarriers`, if its
  `srcQueueFamilyIndex` and `dstQueueFamilyIndex` members are equal, or
  if its `srcQueueFamilyIndex` is the queue family index that was used
  to create the command pool that `commandBuffer` was allocated from,
  then its `srcAccessMask` member **must** only contain access flags
  that are supported by one or more of the pipeline stages in
  `srcStageMask`, as specified in the [table of supported access
  types](#synchronization-access-types-supported)

- <a href="#VUID-vkCmdWaitEvents-pBufferMemoryBarriers-02818"
  id="VUID-vkCmdWaitEvents-pBufferMemoryBarriers-02818"></a>
  VUID-vkCmdWaitEvents-pBufferMemoryBarriers-02818  
  For any element of `pBufferMemoryBarriers`, if its
  `srcQueueFamilyIndex` and `dstQueueFamilyIndex` members are equal, or
  if its `dstQueueFamilyIndex` is the queue family index that was used
  to create the command pool that `commandBuffer` was allocated from,
  then its `dstAccessMask` member **must** only contain access flags
  that are supported by one or more of the pipeline stages in
  `dstStageMask`, as specified in the [table of supported access
  types](#synchronization-access-types-supported)

- <a href="#VUID-vkCmdWaitEvents-pImageMemoryBarriers-02819"
  id="VUID-vkCmdWaitEvents-pImageMemoryBarriers-02819"></a>
  VUID-vkCmdWaitEvents-pImageMemoryBarriers-02819  
  For any element of `pImageMemoryBarriers`, if its
  `srcQueueFamilyIndex` and `dstQueueFamilyIndex` members are equal, or
  if its `srcQueueFamilyIndex` is the queue family index that was used
  to create the command pool that `commandBuffer` was allocated from,
  then its `srcAccessMask` member **must** only contain access flags
  that are supported by one or more of the pipeline stages in
  `srcStageMask`, as specified in the [table of supported access
  types](#synchronization-access-types-supported)

- <a href="#VUID-vkCmdWaitEvents-pImageMemoryBarriers-02820"
  id="VUID-vkCmdWaitEvents-pImageMemoryBarriers-02820"></a>
  VUID-vkCmdWaitEvents-pImageMemoryBarriers-02820  
  For any element of `pImageMemoryBarriers`, if its
  `srcQueueFamilyIndex` and `dstQueueFamilyIndex` members are equal, or
  if its `dstQueueFamilyIndex` is the queue family index that was used
  to create the command pool that `commandBuffer` was allocated from,
  then its `dstAccessMask` member **must** only contain access flags
  that are supported by one or more of the pipeline stages in
  `dstStageMask`, as specified in the [table of supported access
  types](#synchronization-access-types-supported)

- <a href="#VUID-vkCmdWaitEvents-srcStageMask-06459"
  id="VUID-vkCmdWaitEvents-srcStageMask-06459"></a>
  VUID-vkCmdWaitEvents-srcStageMask-06459  
  Any pipeline stage included in `srcStageMask` **must** be supported by
  the capabilities of the queue family specified by the
  `queueFamilyIndex` member of the
  [VkCommandPoolCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandPoolCreateInfo.html) structure that
  was used to create the `VkCommandPool` that `commandBuffer` was
  allocated from, as specified in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-pipeline-stages-supported"
  target="_blank" rel="noopener">table of supported pipeline stages</a>

- <a href="#VUID-vkCmdWaitEvents-dstStageMask-06460"
  id="VUID-vkCmdWaitEvents-dstStageMask-06460"></a>
  VUID-vkCmdWaitEvents-dstStageMask-06460  
  Any pipeline stage included in `dstStageMask` **must** be supported by
  the capabilities of the queue family specified by the
  `queueFamilyIndex` member of the
  [VkCommandPoolCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandPoolCreateInfo.html) structure that
  was used to create the `VkCommandPool` that `commandBuffer` was
  allocated from, as specified in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-pipeline-stages-supported"
  target="_blank" rel="noopener">table of supported pipeline stages</a>

- <a href="#VUID-vkCmdWaitEvents-srcStageMask-01158"
  id="VUID-vkCmdWaitEvents-srcStageMask-01158"></a>
  VUID-vkCmdWaitEvents-srcStageMask-01158  
  `srcStageMask` **must** be the bitwise OR of the `stageMask` parameter
  used in previous calls to `vkCmdSetEvent` with any of the elements of
  `pEvents` and `VK_PIPELINE_STAGE_HOST_BIT` if any of the elements of
  `pEvents` was set using `vkSetEvent`

- <a href="#VUID-vkCmdWaitEvents-srcStageMask-07308"
  id="VUID-vkCmdWaitEvents-srcStageMask-07308"></a>
  VUID-vkCmdWaitEvents-srcStageMask-07308  
  If `vkCmdWaitEvents` is being called inside a render pass instance,
  `srcStageMask` **must** not include `VK_PIPELINE_STAGE_HOST_BIT`

- <a href="#VUID-vkCmdWaitEvents-srcQueueFamilyIndex-02803"
  id="VUID-vkCmdWaitEvents-srcQueueFamilyIndex-02803"></a>
  VUID-vkCmdWaitEvents-srcQueueFamilyIndex-02803  
  The `srcQueueFamilyIndex` and `dstQueueFamilyIndex` members of any
  element of `pBufferMemoryBarriers` or `pImageMemoryBarriers` **must**
  be equal

- <a href="#VUID-vkCmdWaitEvents-commandBuffer-01167"
  id="VUID-vkCmdWaitEvents-commandBuffer-01167"></a>
  VUID-vkCmdWaitEvents-commandBuffer-01167  
  `commandBuffer`’s current device mask **must** include exactly one
  physical device

- <a href="#VUID-vkCmdWaitEvents-pEvents-03847"
  id="VUID-vkCmdWaitEvents-pEvents-03847"></a>
  VUID-vkCmdWaitEvents-pEvents-03847  
  Elements of `pEvents` **must** not have been signaled by
  [vkCmdSetEvent2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetEvent2.html)

Valid Usage (Implicit)

- <a href="#VUID-vkCmdWaitEvents-commandBuffer-parameter"
  id="VUID-vkCmdWaitEvents-commandBuffer-parameter"></a>
  VUID-vkCmdWaitEvents-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdWaitEvents-pEvents-parameter"
  id="VUID-vkCmdWaitEvents-pEvents-parameter"></a>
  VUID-vkCmdWaitEvents-pEvents-parameter  
  `pEvents` **must** be a valid pointer to an array of `eventCount`
  valid [VkEvent](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkEvent.html) handles

- <a href="#VUID-vkCmdWaitEvents-srcStageMask-parameter"
  id="VUID-vkCmdWaitEvents-srcStageMask-parameter"></a>
  VUID-vkCmdWaitEvents-srcStageMask-parameter  
  `srcStageMask` **must** be a valid combination of
  [VkPipelineStageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlagBits.html) values

- <a href="#VUID-vkCmdWaitEvents-dstStageMask-parameter"
  id="VUID-vkCmdWaitEvents-dstStageMask-parameter"></a>
  VUID-vkCmdWaitEvents-dstStageMask-parameter  
  `dstStageMask` **must** be a valid combination of
  [VkPipelineStageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlagBits.html) values

- <a href="#VUID-vkCmdWaitEvents-pMemoryBarriers-parameter"
  id="VUID-vkCmdWaitEvents-pMemoryBarriers-parameter"></a>
  VUID-vkCmdWaitEvents-pMemoryBarriers-parameter  
  If `memoryBarrierCount` is not `0`, `pMemoryBarriers` **must** be a
  valid pointer to an array of `memoryBarrierCount` valid
  [VkMemoryBarrier](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryBarrier.html) structures

- <a href="#VUID-vkCmdWaitEvents-pBufferMemoryBarriers-parameter"
  id="VUID-vkCmdWaitEvents-pBufferMemoryBarriers-parameter"></a>
  VUID-vkCmdWaitEvents-pBufferMemoryBarriers-parameter  
  If `bufferMemoryBarrierCount` is not `0`, `pBufferMemoryBarriers`
  **must** be a valid pointer to an array of `bufferMemoryBarrierCount`
  valid [VkBufferMemoryBarrier](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferMemoryBarrier.html) structures

- <a href="#VUID-vkCmdWaitEvents-pImageMemoryBarriers-parameter"
  id="VUID-vkCmdWaitEvents-pImageMemoryBarriers-parameter"></a>
  VUID-vkCmdWaitEvents-pImageMemoryBarriers-parameter  
  If `imageMemoryBarrierCount` is not `0`, `pImageMemoryBarriers`
  **must** be a valid pointer to an array of `imageMemoryBarrierCount`
  valid [VkImageMemoryBarrier](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageMemoryBarrier.html) structures

- <a href="#VUID-vkCmdWaitEvents-commandBuffer-recording"
  id="VUID-vkCmdWaitEvents-commandBuffer-recording"></a>
  VUID-vkCmdWaitEvents-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdWaitEvents-commandBuffer-cmdpool"
  id="VUID-vkCmdWaitEvents-commandBuffer-cmdpool"></a>
  VUID-vkCmdWaitEvents-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics, compute, decode, or encode operations

- <a href="#VUID-vkCmdWaitEvents-eventCount-arraylength"
  id="VUID-vkCmdWaitEvents-eventCount-arraylength"></a>
  VUID-vkCmdWaitEvents-eventCount-arraylength  
  `eventCount` **must** be greater than `0`

- <a href="#VUID-vkCmdWaitEvents-commonparent"
  id="VUID-vkCmdWaitEvents-commonparent"></a>
  VUID-vkCmdWaitEvents-commonparent  
  Both of `commandBuffer`, and the elements of `pEvents` **must** have
  been created, allocated, or retrieved from the same
  [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

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
<td
class="tableblock halign-left valign-top"><p>Synchronization</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkBufferMemoryBarrier](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferMemoryBarrier.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html), [VkEvent](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkEvent.html),
[VkImageMemoryBarrier](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageMemoryBarrier.html),
[VkMemoryBarrier](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryBarrier.html),
[VkPipelineStageFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlags.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdWaitEvents"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
