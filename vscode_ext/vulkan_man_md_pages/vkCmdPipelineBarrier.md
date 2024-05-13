# vkCmdPipelineBarrier(3) Manual Page

## Name

vkCmdPipelineBarrier - Insert a memory dependency



## <a href="#_c_specification" class="anchor"></a>C Specification

To record a pipeline barrier, call:

``` c
// Provided by VK_VERSION_1_0
void vkCmdPipelineBarrier(
    VkCommandBuffer                             commandBuffer,
    VkPipelineStageFlags                        srcStageMask,
    VkPipelineStageFlags                        dstStageMask,
    VkDependencyFlags                           dependencyFlags,
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

- `srcStageMask` is a bitmask of
  [VkPipelineStageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlagBits.html) specifying the
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-pipeline-stages-masks"
  target="_blank" rel="noopener">source stages</a>.

- `dstStageMask` is a bitmask of
  [VkPipelineStageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlagBits.html) specifying the
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-pipeline-stages-masks"
  target="_blank" rel="noopener">destination stages</a>.

- `dependencyFlags` is a bitmask of
  [VkDependencyFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDependencyFlagBits.html) specifying how
  execution and memory dependencies are formed.

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

`vkCmdPipelineBarrier` operates almost identically to
[vkCmdPipelineBarrier2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdPipelineBarrier2.html), except that the
scopes and barriers are defined as direct parameters rather than being
defined by a [VkDependencyInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDependencyInfo.html).

When [vkCmdPipelineBarrier](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdPipelineBarrier.html) is submitted to a
queue, it defines a memory dependency between commands that were
submitted to the same queue before it, and those submitted to the same
queue after it.

If [vkCmdPipelineBarrier](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdPipelineBarrier.html) was recorded
outside a render pass instance, the first <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies-scopes"
target="_blank" rel="noopener">synchronization scope</a> includes all
commands that occur earlier in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-submission-order"
target="_blank" rel="noopener">submission order</a>. If
[vkCmdPipelineBarrier](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdPipelineBarrier.html) was recorded inside a
render pass instance, the first synchronization scope includes only
commands that occur earlier in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-submission-order"
target="_blank" rel="noopener">submission order</a> within the same
subpass. In either case, the first synchronization scope is limited to
operations on the pipeline stages determined by the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-pipeline-stages-masks"
target="_blank" rel="noopener">source stage mask</a> specified by
`srcStageMask`.

If [vkCmdPipelineBarrier](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdPipelineBarrier.html) was recorded
outside a render pass instance, the second <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies-scopes"
target="_blank" rel="noopener">synchronization scope</a> includes all
commands that occur later in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-submission-order"
target="_blank" rel="noopener">submission order</a>. If
[vkCmdPipelineBarrier](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdPipelineBarrier.html) was recorded inside a
render pass instance, the second synchronization scope includes only
commands that occur later in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-submission-order"
target="_blank" rel="noopener">submission order</a> within the same
subpass. In either case, the second synchronization scope is limited to
operations on the pipeline stages determined by the <a
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

If `dependencyFlags` includes `VK_DEPENDENCY_BY_REGION_BIT`, then any
dependency between <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-framebuffer-regions"
target="_blank" rel="noopener">framebuffer-space</a> pipeline stages is
<a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-framebuffer-regions"
target="_blank" rel="noopener">framebuffer-local</a> - otherwise it is
<a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-framebuffer-regions"
target="_blank" rel="noopener">framebuffer-global</a>.

Valid Usage

- <a href="#VUID-vkCmdPipelineBarrier-srcStageMask-04090"
  id="VUID-vkCmdPipelineBarrier-srcStageMask-04090"></a>
  VUID-vkCmdPipelineBarrier-srcStageMask-04090  
  If the [`geometryShader`](#features-geometryShader) feature is not
  enabled, `srcStageMask` **must** not contain
  `VK_PIPELINE_STAGE_GEOMETRY_SHADER_BIT`

- <a href="#VUID-vkCmdPipelineBarrier-srcStageMask-04091"
  id="VUID-vkCmdPipelineBarrier-srcStageMask-04091"></a>
  VUID-vkCmdPipelineBarrier-srcStageMask-04091  
  If the [`tessellationShader`](#features-tessellationShader) feature is
  not enabled, `srcStageMask` **must** not contain
  `VK_PIPELINE_STAGE_TESSELLATION_CONTROL_SHADER_BIT` or
  `VK_PIPELINE_STAGE_TESSELLATION_EVALUATION_SHADER_BIT`

- <a href="#VUID-vkCmdPipelineBarrier-srcStageMask-04092"
  id="VUID-vkCmdPipelineBarrier-srcStageMask-04092"></a>
  VUID-vkCmdPipelineBarrier-srcStageMask-04092  
  If the [`conditionalRendering`](#features-conditionalRendering)
  feature is not enabled, `srcStageMask` **must** not contain
  `VK_PIPELINE_STAGE_CONDITIONAL_RENDERING_BIT_EXT`

- <a href="#VUID-vkCmdPipelineBarrier-srcStageMask-04093"
  id="VUID-vkCmdPipelineBarrier-srcStageMask-04093"></a>
  VUID-vkCmdPipelineBarrier-srcStageMask-04093  
  If the [`fragmentDensityMap`](#features-fragmentDensityMap) feature is
  not enabled, `srcStageMask` **must** not contain
  `VK_PIPELINE_STAGE_FRAGMENT_DENSITY_PROCESS_BIT_EXT`

- <a href="#VUID-vkCmdPipelineBarrier-srcStageMask-04094"
  id="VUID-vkCmdPipelineBarrier-srcStageMask-04094"></a>
  VUID-vkCmdPipelineBarrier-srcStageMask-04094  
  If the [`transformFeedback`](#features-transformFeedback) feature is
  not enabled, `srcStageMask` **must** not contain
  `VK_PIPELINE_STAGE_TRANSFORM_FEEDBACK_BIT_EXT`

- <a href="#VUID-vkCmdPipelineBarrier-srcStageMask-04095"
  id="VUID-vkCmdPipelineBarrier-srcStageMask-04095"></a>
  VUID-vkCmdPipelineBarrier-srcStageMask-04095  
  If the [`meshShader`](#features-meshShader) feature is not enabled,
  `srcStageMask` **must** not contain
  `VK_PIPELINE_STAGE_MESH_SHADER_BIT_EXT`

- <a href="#VUID-vkCmdPipelineBarrier-srcStageMask-04096"
  id="VUID-vkCmdPipelineBarrier-srcStageMask-04096"></a>
  VUID-vkCmdPipelineBarrier-srcStageMask-04096  
  If the [`taskShader`](#features-taskShader) feature is not enabled,
  `srcStageMask` **must** not contain
  `VK_PIPELINE_STAGE_TASK_SHADER_BIT_EXT`

- <a href="#VUID-vkCmdPipelineBarrier-srcStageMask-07318"
  id="VUID-vkCmdPipelineBarrier-srcStageMask-07318"></a>
  VUID-vkCmdPipelineBarrier-srcStageMask-07318  
  If neither the [`shadingRateImage`](#features-shadingRateImage) or
  [`attachmentFragmentShadingRate`](#features-attachmentFragmentShadingRate)
  are enabled, `srcStageMask` **must** not contain
  `VK_PIPELINE_STAGE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR`

- <a href="#VUID-vkCmdPipelineBarrier-srcStageMask-03937"
  id="VUID-vkCmdPipelineBarrier-srcStageMask-03937"></a>
  VUID-vkCmdPipelineBarrier-srcStageMask-03937  
  If the [`synchronization2`](#features-synchronization2) feature is not
  enabled, `srcStageMask` **must** not be `0`

- <a href="#VUID-vkCmdPipelineBarrier-srcStageMask-07949"
  id="VUID-vkCmdPipelineBarrier-srcStageMask-07949"></a>
  VUID-vkCmdPipelineBarrier-srcStageMask-07949  
  If neither the [VK_NV_ray_tracing](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_ray_tracing.html) extension
  or [`rayTracingPipeline` feature](#features-rayTracingPipeline) are
  enabled, `srcStageMask` **must** not contain
  `VK_PIPELINE_STAGE_RAY_TRACING_SHADER_BIT_KHR`

<!-- -->

- <a href="#VUID-vkCmdPipelineBarrier-srcAccessMask-06257"
  id="VUID-vkCmdPipelineBarrier-srcAccessMask-06257"></a>
  VUID-vkCmdPipelineBarrier-srcAccessMask-06257  
  If the [`rayQuery`](#features-rayQuery) feature is not enabled and a
  memory barrier `srcAccessMask` includes
  `VK_ACCESS_ACCELERATION_STRUCTURE_READ_BIT_KHR`, `srcStageMask`
  **must** not include any of the `VK_PIPELINE_STAGE_*_SHADER_BIT`
  stages except `VK_PIPELINE_STAGE_RAY_TRACING_SHADER_BIT_KHR`

<!-- -->

- <a href="#VUID-vkCmdPipelineBarrier-dstStageMask-04090"
  id="VUID-vkCmdPipelineBarrier-dstStageMask-04090"></a>
  VUID-vkCmdPipelineBarrier-dstStageMask-04090  
  If the [`geometryShader`](#features-geometryShader) feature is not
  enabled, `dstStageMask` **must** not contain
  `VK_PIPELINE_STAGE_GEOMETRY_SHADER_BIT`

- <a href="#VUID-vkCmdPipelineBarrier-dstStageMask-04091"
  id="VUID-vkCmdPipelineBarrier-dstStageMask-04091"></a>
  VUID-vkCmdPipelineBarrier-dstStageMask-04091  
  If the [`tessellationShader`](#features-tessellationShader) feature is
  not enabled, `dstStageMask` **must** not contain
  `VK_PIPELINE_STAGE_TESSELLATION_CONTROL_SHADER_BIT` or
  `VK_PIPELINE_STAGE_TESSELLATION_EVALUATION_SHADER_BIT`

- <a href="#VUID-vkCmdPipelineBarrier-dstStageMask-04092"
  id="VUID-vkCmdPipelineBarrier-dstStageMask-04092"></a>
  VUID-vkCmdPipelineBarrier-dstStageMask-04092  
  If the [`conditionalRendering`](#features-conditionalRendering)
  feature is not enabled, `dstStageMask` **must** not contain
  `VK_PIPELINE_STAGE_CONDITIONAL_RENDERING_BIT_EXT`

- <a href="#VUID-vkCmdPipelineBarrier-dstStageMask-04093"
  id="VUID-vkCmdPipelineBarrier-dstStageMask-04093"></a>
  VUID-vkCmdPipelineBarrier-dstStageMask-04093  
  If the [`fragmentDensityMap`](#features-fragmentDensityMap) feature is
  not enabled, `dstStageMask` **must** not contain
  `VK_PIPELINE_STAGE_FRAGMENT_DENSITY_PROCESS_BIT_EXT`

- <a href="#VUID-vkCmdPipelineBarrier-dstStageMask-04094"
  id="VUID-vkCmdPipelineBarrier-dstStageMask-04094"></a>
  VUID-vkCmdPipelineBarrier-dstStageMask-04094  
  If the [`transformFeedback`](#features-transformFeedback) feature is
  not enabled, `dstStageMask` **must** not contain
  `VK_PIPELINE_STAGE_TRANSFORM_FEEDBACK_BIT_EXT`

- <a href="#VUID-vkCmdPipelineBarrier-dstStageMask-04095"
  id="VUID-vkCmdPipelineBarrier-dstStageMask-04095"></a>
  VUID-vkCmdPipelineBarrier-dstStageMask-04095  
  If the [`meshShader`](#features-meshShader) feature is not enabled,
  `dstStageMask` **must** not contain
  `VK_PIPELINE_STAGE_MESH_SHADER_BIT_EXT`

- <a href="#VUID-vkCmdPipelineBarrier-dstStageMask-04096"
  id="VUID-vkCmdPipelineBarrier-dstStageMask-04096"></a>
  VUID-vkCmdPipelineBarrier-dstStageMask-04096  
  If the [`taskShader`](#features-taskShader) feature is not enabled,
  `dstStageMask` **must** not contain
  `VK_PIPELINE_STAGE_TASK_SHADER_BIT_EXT`

- <a href="#VUID-vkCmdPipelineBarrier-dstStageMask-07318"
  id="VUID-vkCmdPipelineBarrier-dstStageMask-07318"></a>
  VUID-vkCmdPipelineBarrier-dstStageMask-07318  
  If neither the [`shadingRateImage`](#features-shadingRateImage) or
  [`attachmentFragmentShadingRate`](#features-attachmentFragmentShadingRate)
  are enabled, `dstStageMask` **must** not contain
  `VK_PIPELINE_STAGE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR`

- <a href="#VUID-vkCmdPipelineBarrier-dstStageMask-03937"
  id="VUID-vkCmdPipelineBarrier-dstStageMask-03937"></a>
  VUID-vkCmdPipelineBarrier-dstStageMask-03937  
  If the [`synchronization2`](#features-synchronization2) feature is not
  enabled, `dstStageMask` **must** not be `0`

- <a href="#VUID-vkCmdPipelineBarrier-dstStageMask-07949"
  id="VUID-vkCmdPipelineBarrier-dstStageMask-07949"></a>
  VUID-vkCmdPipelineBarrier-dstStageMask-07949  
  If neither the [VK_NV_ray_tracing](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_ray_tracing.html) extension
  or [`rayTracingPipeline` feature](#features-rayTracingPipeline) are
  enabled, `dstStageMask` **must** not contain
  `VK_PIPELINE_STAGE_RAY_TRACING_SHADER_BIT_KHR`

<!-- -->

- <a href="#VUID-vkCmdPipelineBarrier-dstAccessMask-06257"
  id="VUID-vkCmdPipelineBarrier-dstAccessMask-06257"></a>
  VUID-vkCmdPipelineBarrier-dstAccessMask-06257  
  If the [`rayQuery`](#features-rayQuery) feature is not enabled and a
  memory barrier `dstAccessMask` includes
  `VK_ACCESS_ACCELERATION_STRUCTURE_READ_BIT_KHR`, `dstStageMask`
  **must** not include any of the `VK_PIPELINE_STAGE_*_SHADER_BIT`
  stages except `VK_PIPELINE_STAGE_RAY_TRACING_SHADER_BIT_KHR`

<!-- -->

- <a href="#VUID-vkCmdPipelineBarrier-srcAccessMask-02815"
  id="VUID-vkCmdPipelineBarrier-srcAccessMask-02815"></a>
  VUID-vkCmdPipelineBarrier-srcAccessMask-02815  
  The `srcAccessMask` member of each element of `pMemoryBarriers`
  **must** only include access flags that are supported by one or more
  of the pipeline stages in `srcStageMask`, as specified in the [table
  of supported access types](#synchronization-access-types-supported)

- <a href="#VUID-vkCmdPipelineBarrier-dstAccessMask-02816"
  id="VUID-vkCmdPipelineBarrier-dstAccessMask-02816"></a>
  VUID-vkCmdPipelineBarrier-dstAccessMask-02816  
  The `dstAccessMask` member of each element of `pMemoryBarriers`
  **must** only include access flags that are supported by one or more
  of the pipeline stages in `dstStageMask`, as specified in the [table
  of supported access types](#synchronization-access-types-supported)

- <a href="#VUID-vkCmdPipelineBarrier-pBufferMemoryBarriers-02817"
  id="VUID-vkCmdPipelineBarrier-pBufferMemoryBarriers-02817"></a>
  VUID-vkCmdPipelineBarrier-pBufferMemoryBarriers-02817  
  For any element of `pBufferMemoryBarriers`, if its
  `srcQueueFamilyIndex` and `dstQueueFamilyIndex` members are equal, or
  if its `srcQueueFamilyIndex` is the queue family index that was used
  to create the command pool that `commandBuffer` was allocated from,
  then its `srcAccessMask` member **must** only contain access flags
  that are supported by one or more of the pipeline stages in
  `srcStageMask`, as specified in the [table of supported access
  types](#synchronization-access-types-supported)

- <a href="#VUID-vkCmdPipelineBarrier-pBufferMemoryBarriers-02818"
  id="VUID-vkCmdPipelineBarrier-pBufferMemoryBarriers-02818"></a>
  VUID-vkCmdPipelineBarrier-pBufferMemoryBarriers-02818  
  For any element of `pBufferMemoryBarriers`, if its
  `srcQueueFamilyIndex` and `dstQueueFamilyIndex` members are equal, or
  if its `dstQueueFamilyIndex` is the queue family index that was used
  to create the command pool that `commandBuffer` was allocated from,
  then its `dstAccessMask` member **must** only contain access flags
  that are supported by one or more of the pipeline stages in
  `dstStageMask`, as specified in the [table of supported access
  types](#synchronization-access-types-supported)

- <a href="#VUID-vkCmdPipelineBarrier-pImageMemoryBarriers-02819"
  id="VUID-vkCmdPipelineBarrier-pImageMemoryBarriers-02819"></a>
  VUID-vkCmdPipelineBarrier-pImageMemoryBarriers-02819  
  For any element of `pImageMemoryBarriers`, if its
  `srcQueueFamilyIndex` and `dstQueueFamilyIndex` members are equal, or
  if its `srcQueueFamilyIndex` is the queue family index that was used
  to create the command pool that `commandBuffer` was allocated from,
  then its `srcAccessMask` member **must** only contain access flags
  that are supported by one or more of the pipeline stages in
  `srcStageMask`, as specified in the [table of supported access
  types](#synchronization-access-types-supported)

- <a href="#VUID-vkCmdPipelineBarrier-pImageMemoryBarriers-02820"
  id="VUID-vkCmdPipelineBarrier-pImageMemoryBarriers-02820"></a>
  VUID-vkCmdPipelineBarrier-pImageMemoryBarriers-02820  
  For any element of `pImageMemoryBarriers`, if its
  `srcQueueFamilyIndex` and `dstQueueFamilyIndex` members are equal, or
  if its `dstQueueFamilyIndex` is the queue family index that was used
  to create the command pool that `commandBuffer` was allocated from,
  then its `dstAccessMask` member **must** only contain access flags
  that are supported by one or more of the pipeline stages in
  `dstStageMask`, as specified in the [table of supported access
  types](#synchronization-access-types-supported)

<!-- -->

- <a href="#VUID-vkCmdPipelineBarrier-None-07889"
  id="VUID-vkCmdPipelineBarrier-None-07889"></a>
  VUID-vkCmdPipelineBarrier-None-07889  
  If `vkCmdPipelineBarrier` is called within a render pass instance
  using a [VkRenderPass](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPass.html) object, the render pass
  **must** have been created with at least one subpass dependency that
  expresses a dependency from the current subpass to itself, does not
  include `VK_DEPENDENCY_BY_REGION_BIT` if this command does not, does
  not include `VK_DEPENDENCY_VIEW_LOCAL_BIT` if this command does not,
  and has [synchronization scopes](#synchronization-dependencies-scopes)
  and [access scopes](#synchronization-dependencies-access-scopes) that
  are all supersets of the scopes defined in this command

- <a href="#VUID-vkCmdPipelineBarrier-bufferMemoryBarrierCount-01178"
  id="VUID-vkCmdPipelineBarrier-bufferMemoryBarrierCount-01178"></a>
  VUID-vkCmdPipelineBarrier-bufferMemoryBarrierCount-01178  
  If `vkCmdPipelineBarrier` is called within a render pass instance
  using a [VkRenderPass](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPass.html) object, it **must** not
  include any buffer memory barriers

- <a href="#VUID-vkCmdPipelineBarrier-image-04073"
  id="VUID-vkCmdPipelineBarrier-image-04073"></a>
  VUID-vkCmdPipelineBarrier-image-04073  
  If `vkCmdPipelineBarrier` is called within a render pass instance
  using a [VkRenderPass](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPass.html) object, the `image` member
  of any image memory barrier included in this command **must** be an
  attachment used in the current subpass both as an input attachment,
  and as either a color, color resolve, or depth/stencil attachment

- <a href="#VUID-vkCmdPipelineBarrier-image-09373"
  id="VUID-vkCmdPipelineBarrier-image-09373"></a>
  VUID-vkCmdPipelineBarrier-image-09373  
  If `vkCmdPipelineBarrier` is called within a render pass instance
  using a [VkRenderPass](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPass.html) object, and the `image`
  member of any image memory barrier is a color resolve attachment, the
  corresponding color attachment **must** be `VK_ATTACHMENT_UNUSED`

- <a href="#VUID-vkCmdPipelineBarrier-image-09374"
  id="VUID-vkCmdPipelineBarrier-image-09374"></a>
  VUID-vkCmdPipelineBarrier-image-09374  
  If `vkCmdPipelineBarrier` is called within a render pass instance
  using a [VkRenderPass](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPass.html) object, and the `image`
  member of any image memory barrier is a color resolve attachment, it
  **must** have been created with a non-zero
  [VkExternalFormatANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFormatANDROID.html)::`externalFormat`
  value

- <a href="#VUID-vkCmdPipelineBarrier-oldLayout-01181"
  id="VUID-vkCmdPipelineBarrier-oldLayout-01181"></a>
  VUID-vkCmdPipelineBarrier-oldLayout-01181  
  If `vkCmdPipelineBarrier` is called within a render pass instance, the
  `oldLayout` and `newLayout` members of any image memory barrier
  included in this command **must** be equal

- <a href="#VUID-vkCmdPipelineBarrier-srcQueueFamilyIndex-01182"
  id="VUID-vkCmdPipelineBarrier-srcQueueFamilyIndex-01182"></a>
  VUID-vkCmdPipelineBarrier-srcQueueFamilyIndex-01182  
  If `vkCmdPipelineBarrier` is called within a render pass instance, the
  `srcQueueFamilyIndex` and `dstQueueFamilyIndex` members of any memory
  barrier included in this command **must** be equal

- <a href="#VUID-vkCmdPipelineBarrier-None-07890"
  id="VUID-vkCmdPipelineBarrier-None-07890"></a>
  VUID-vkCmdPipelineBarrier-None-07890  
  If `vkCmdPipelineBarrier` is called within a render pass instance, and
  the source stage masks of any memory barriers include
  [framebuffer-space stages](#synchronization-framebuffer-regions),
  destination stage masks of all memory barriers **must** only include
  [framebuffer-space stages](#synchronization-framebuffer-regions)

- <a href="#VUID-vkCmdPipelineBarrier-dependencyFlags-07891"
  id="VUID-vkCmdPipelineBarrier-dependencyFlags-07891"></a>
  VUID-vkCmdPipelineBarrier-dependencyFlags-07891  
  If `vkCmdPipelineBarrier` is called within a render pass instance, and
  the source stage masks of any memory barriers include
  [framebuffer-space stages](#synchronization-framebuffer-regions), then
  `dependencyFlags` **must** include `VK_DEPENDENCY_BY_REGION_BIT`

- <a href="#VUID-vkCmdPipelineBarrier-None-07892"
  id="VUID-vkCmdPipelineBarrier-None-07892"></a>
  VUID-vkCmdPipelineBarrier-None-07892  
  If `vkCmdPipelineBarrier` is called within a render pass instance, the
  source and destination stage masks of any memory barriers **must**
  only include graphics pipeline stages

- <a href="#VUID-vkCmdPipelineBarrier-dependencyFlags-01186"
  id="VUID-vkCmdPipelineBarrier-dependencyFlags-01186"></a>
  VUID-vkCmdPipelineBarrier-dependencyFlags-01186  
  If `vkCmdPipelineBarrier` is called outside of a render pass instance,
  the dependency flags **must** not include
  `VK_DEPENDENCY_VIEW_LOCAL_BIT`

- <a href="#VUID-vkCmdPipelineBarrier-None-07893"
  id="VUID-vkCmdPipelineBarrier-None-07893"></a>
  VUID-vkCmdPipelineBarrier-None-07893  
  If `vkCmdPipelineBarrier` is called inside a render pass instance, and
  there is more than one view in the current subpass, dependency flags
  **must** include `VK_DEPENDENCY_VIEW_LOCAL_BIT`

- <a href="#VUID-vkCmdPipelineBarrier-None-09553"
  id="VUID-vkCmdPipelineBarrier-None-09553"></a>
  VUID-vkCmdPipelineBarrier-None-09553  
  If none of the
  [`shaderTileImageColorReadAccess`](#features-shaderTileImageColorReadAccess),
  [`shaderTileImageStencilReadAccess`](#features-shaderTileImageStencilReadAccess),
  or
  [`shaderTileImageDepthReadAccess`](#features-shaderTileImageDepthReadAccess)
  features are enabled, and the
  [`dynamicRenderingLocalRead`](#features-dynamicRenderingLocalRead)
  feature is not enabled, `vkCmdPipelineBarrier` **must** not be called
  within a render pass instance started with
  [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html)

- <a href="#VUID-vkCmdPipelineBarrier-None-09554"
  id="VUID-vkCmdPipelineBarrier-None-09554"></a>
  VUID-vkCmdPipelineBarrier-None-09554  
  If the
  [`dynamicRenderingLocalRead`](#features-dynamicRenderingLocalRead)
  feature is not enabled, and `vkCmdPipelineBarrier` is called within a
  render pass instance started with
  [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html), there **must** be no
  buffer or image memory barriers specified by this command

- <a href="#VUID-vkCmdPipelineBarrier-None-09586"
  id="VUID-vkCmdPipelineBarrier-None-09586"></a>
  VUID-vkCmdPipelineBarrier-None-09586  
  If the
  [`dynamicRenderingLocalRead`](#features-dynamicRenderingLocalRead)
  feature is not enabled, and `vkCmdPipelineBarrier` is called within a
  render pass instance started with
  [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html), memory barriers
  specified by this command **must** only include
  `VK_ACCESS_2_COLOR_ATTACHMENT_READ_BIT`,
  `VK_ACCESS_2_COLOR_ATTACHMENT_WRITE_BIT`,
  `VK_ACCESS_2_DEPTH_STENCIL_ATTACHMENT_READ_BIT`, or
  `VK_ACCESS_2_DEPTH_STENCIL_ATTACHMENT_WRITE_BIT` in their access masks

- <a href="#VUID-vkCmdPipelineBarrier-image-09555"
  id="VUID-vkCmdPipelineBarrier-image-09555"></a>
  VUID-vkCmdPipelineBarrier-image-09555  
  If `vkCmdPipelineBarrier` is called within a render pass instance
  started with [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html), and the
  `image` member of any image memory barrier is used as an attachment in
  the current render pass instance, it **must** be in the
  `VK_IMAGE_LAYOUT_RENDERING_LOCAL_READ_KHR` or
  `VK_IMAGE_LAYOUT_GENERAL` layout

- <a href="#VUID-vkCmdPipelineBarrier-srcStageMask-09556"
  id="VUID-vkCmdPipelineBarrier-srcStageMask-09556"></a>
  VUID-vkCmdPipelineBarrier-srcStageMask-09556  
  If `vkCmdPipelineBarrier` is called within a render pass instance
  started with [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html), this
  command **must** only specify [framebuffer-space
  stages](#synchronization-framebuffer-regions) in `srcStageMask` and
  `dstStageMask`

- <a href="#VUID-vkCmdPipelineBarrier-srcStageMask-06461"
  id="VUID-vkCmdPipelineBarrier-srcStageMask-06461"></a>
  VUID-vkCmdPipelineBarrier-srcStageMask-06461  
  Any pipeline stage included in `srcStageMask` **must** be supported by
  the capabilities of the queue family specified by the
  `queueFamilyIndex` member of the
  [VkCommandPoolCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandPoolCreateInfo.html) structure that
  was used to create the `VkCommandPool` that `commandBuffer` was
  allocated from, as specified in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-pipeline-stages-supported"
  target="_blank" rel="noopener">table of supported pipeline stages</a>

- <a href="#VUID-vkCmdPipelineBarrier-dstStageMask-06462"
  id="VUID-vkCmdPipelineBarrier-dstStageMask-06462"></a>
  VUID-vkCmdPipelineBarrier-dstStageMask-06462  
  Any pipeline stage included in `dstStageMask` **must** be supported by
  the capabilities of the queue family specified by the
  `queueFamilyIndex` member of the
  [VkCommandPoolCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandPoolCreateInfo.html) structure that
  was used to create the `VkCommandPool` that `commandBuffer` was
  allocated from, as specified in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-pipeline-stages-supported"
  target="_blank" rel="noopener">table of supported pipeline stages</a>

- <a href="#VUID-vkCmdPipelineBarrier-srcStageMask-09633"
  id="VUID-vkCmdPipelineBarrier-srcStageMask-09633"></a>
  VUID-vkCmdPipelineBarrier-srcStageMask-09633  
  If either `srcStageMask` or `dstStageMask` includes
  `VK_PIPELINE_STAGE_HOST_BIT`, for any element of
  `pImageMemoryBarriers`, `srcQueueFamilyIndex` and
  `dstQueueFamilyIndex` **must** be equal

- <a href="#VUID-vkCmdPipelineBarrier-srcStageMask-09634"
  id="VUID-vkCmdPipelineBarrier-srcStageMask-09634"></a>
  VUID-vkCmdPipelineBarrier-srcStageMask-09634  
  If either `srcStageMask` or `dstStageMask` includes
  `VK_PIPELINE_STAGE_HOST_BIT`, for any element of
  `pBufferMemoryBarriers`, `srcQueueFamilyIndex` and
  `dstQueueFamilyIndex` **must** be equal

Valid Usage (Implicit)

- <a href="#VUID-vkCmdPipelineBarrier-commandBuffer-parameter"
  id="VUID-vkCmdPipelineBarrier-commandBuffer-parameter"></a>
  VUID-vkCmdPipelineBarrier-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdPipelineBarrier-srcStageMask-parameter"
  id="VUID-vkCmdPipelineBarrier-srcStageMask-parameter"></a>
  VUID-vkCmdPipelineBarrier-srcStageMask-parameter  
  `srcStageMask` **must** be a valid combination of
  [VkPipelineStageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlagBits.html) values

- <a href="#VUID-vkCmdPipelineBarrier-dstStageMask-parameter"
  id="VUID-vkCmdPipelineBarrier-dstStageMask-parameter"></a>
  VUID-vkCmdPipelineBarrier-dstStageMask-parameter  
  `dstStageMask` **must** be a valid combination of
  [VkPipelineStageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlagBits.html) values

- <a href="#VUID-vkCmdPipelineBarrier-dependencyFlags-parameter"
  id="VUID-vkCmdPipelineBarrier-dependencyFlags-parameter"></a>
  VUID-vkCmdPipelineBarrier-dependencyFlags-parameter  
  `dependencyFlags` **must** be a valid combination of
  [VkDependencyFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDependencyFlagBits.html) values

- <a href="#VUID-vkCmdPipelineBarrier-pMemoryBarriers-parameter"
  id="VUID-vkCmdPipelineBarrier-pMemoryBarriers-parameter"></a>
  VUID-vkCmdPipelineBarrier-pMemoryBarriers-parameter  
  If `memoryBarrierCount` is not `0`, `pMemoryBarriers` **must** be a
  valid pointer to an array of `memoryBarrierCount` valid
  [VkMemoryBarrier](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryBarrier.html) structures

- <a href="#VUID-vkCmdPipelineBarrier-pBufferMemoryBarriers-parameter"
  id="VUID-vkCmdPipelineBarrier-pBufferMemoryBarriers-parameter"></a>
  VUID-vkCmdPipelineBarrier-pBufferMemoryBarriers-parameter  
  If `bufferMemoryBarrierCount` is not `0`, `pBufferMemoryBarriers`
  **must** be a valid pointer to an array of `bufferMemoryBarrierCount`
  valid [VkBufferMemoryBarrier](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferMemoryBarrier.html) structures

- <a href="#VUID-vkCmdPipelineBarrier-pImageMemoryBarriers-parameter"
  id="VUID-vkCmdPipelineBarrier-pImageMemoryBarriers-parameter"></a>
  VUID-vkCmdPipelineBarrier-pImageMemoryBarriers-parameter  
  If `imageMemoryBarrierCount` is not `0`, `pImageMemoryBarriers`
  **must** be a valid pointer to an array of `imageMemoryBarrierCount`
  valid [VkImageMemoryBarrier](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageMemoryBarrier.html) structures

- <a href="#VUID-vkCmdPipelineBarrier-commandBuffer-recording"
  id="VUID-vkCmdPipelineBarrier-commandBuffer-recording"></a>
  VUID-vkCmdPipelineBarrier-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdPipelineBarrier-commandBuffer-cmdpool"
  id="VUID-vkCmdPipelineBarrier-commandBuffer-cmdpool"></a>
  VUID-vkCmdPipelineBarrier-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support transfer, graphics, compute, decode, or encode operations

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
<td
class="tableblock halign-left valign-top"><p>Synchronization</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkBufferMemoryBarrier](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferMemoryBarrier.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkDependencyFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDependencyFlags.html),
[VkImageMemoryBarrier](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageMemoryBarrier.html),
[VkMemoryBarrier](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryBarrier.html),
[VkPipelineStageFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlags.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdPipelineBarrier"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
