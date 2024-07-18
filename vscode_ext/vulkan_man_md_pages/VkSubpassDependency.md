# VkSubpassDependency(3) Manual Page

## Name

VkSubpassDependency - Structure specifying a subpass dependency



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkSubpassDependency` structure is defined as:

``` c
// Provided by VK_VERSION_1_0
typedef struct VkSubpassDependency {
    uint32_t                srcSubpass;
    uint32_t                dstSubpass;
    VkPipelineStageFlags    srcStageMask;
    VkPipelineStageFlags    dstStageMask;
    VkAccessFlags           srcAccessMask;
    VkAccessFlags           dstAccessMask;
    VkDependencyFlags       dependencyFlags;
} VkSubpassDependency;
```

## <a href="#_members" class="anchor"></a>Members

- `srcSubpass` is the subpass index of the first subpass in the
  dependency, or `VK_SUBPASS_EXTERNAL`.

- `dstSubpass` is the subpass index of the second subpass in the
  dependency, or `VK_SUBPASS_EXTERNAL`.

- `srcStageMask` is a bitmask of
  [VkPipelineStageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlagBits.html) specifying the
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-pipeline-stages-masks"
  target="_blank" rel="noopener">source stage mask</a>.

- `dstStageMask` is a bitmask of
  [VkPipelineStageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlagBits.html) specifying the
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-pipeline-stages-masks"
  target="_blank" rel="noopener">destination stage mask</a>

- `srcAccessMask` is a bitmask of
  [VkAccessFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccessFlagBits.html) specifying a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-access-masks"
  target="_blank" rel="noopener">source access mask</a>.

- `dstAccessMask` is a bitmask of
  [VkAccessFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccessFlagBits.html) specifying a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-access-masks"
  target="_blank" rel="noopener">destination access mask</a>.

- `dependencyFlags` is a bitmask of
  [VkDependencyFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDependencyFlagBits.html).

## <a href="#_description" class="anchor"></a>Description

If `srcSubpass` is equal to `dstSubpass` then the
[VkSubpassDependency](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassDependency.html) does not directly define
a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies"
target="_blank" rel="noopener">dependency</a>. Instead, it enables
pipeline barriers to be used in a render pass instance within the
identified subpass, where the scopes of one pipeline barrier **must** be
a subset of those described by one subpass dependency. Subpass
dependencies specified in this way that include <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-framebuffer-regions"
target="_blank" rel="noopener">framebuffer-space stages</a> in the
`srcStageMask` **must** only include <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-framebuffer-regions"
target="_blank" rel="noopener">framebuffer-space stages</a> in
`dstStageMask`, and **must** include `VK_DEPENDENCY_BY_REGION_BIT`. When
a subpass dependency is specified in this way for a subpass that has
more than one view in its view mask, its `dependencyFlags` **must**
include `VK_DEPENDENCY_VIEW_LOCAL_BIT`.

If `srcSubpass` and `dstSubpass` are not equal, when a render pass
instance which includes a subpass dependency is submitted to a queue, it
defines a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies"
target="_blank" rel="noopener">dependency</a> between the subpasses
identified by `srcSubpass` and `dstSubpass`.

If `srcSubpass` is equal to `VK_SUBPASS_EXTERNAL`, the first <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies-scopes"
target="_blank" rel="noopener">synchronization scope</a> includes
commands that occur earlier in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-submission-order"
target="_blank" rel="noopener">submission order</a> than the
[vkCmdBeginRenderPass](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRenderPass.html) used to begin the
render pass instance. Otherwise, the first set of commands includes all
commands submitted as part of the subpass instance identified by
`srcSubpass` and any <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass-load-operations"
target="_blank" rel="noopener">load</a>, <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass-store-operations"
target="_blank" rel="noopener">store</a>, or <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass-resolve-operations"
target="_blank" rel="noopener">multisample resolve</a> operations on
attachments used in `srcSubpass`. In either case, the first
synchronization scope is limited to operations on the pipeline stages
determined by the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-pipeline-stages-masks"
target="_blank" rel="noopener">source stage mask</a> specified by
`srcStageMask`.

If `dstSubpass` is equal to `VK_SUBPASS_EXTERNAL`, the second <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies-scopes"
target="_blank" rel="noopener">synchronization scope</a> includes
commands that occur later in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-submission-order"
target="_blank" rel="noopener">submission order</a> than the
[vkCmdEndRenderPass](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdEndRenderPass.html) used to end the render
pass instance. Otherwise, the second set of commands includes all
commands submitted as part of the subpass instance identified by
`dstSubpass` and any <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass-load-operations"
target="_blank" rel="noopener">load</a>, <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass-store-operations"
target="_blank" rel="noopener">store</a>, and <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass-resolve-operations"
target="_blank" rel="noopener">multisample resolve</a> operations on
attachments used in `dstSubpass`. In either case, the second
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
`srcStageMask`. It is also limited to access types in the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-access-masks"
target="_blank" rel="noopener">source access mask</a> specified by
`srcAccessMask`.

The second <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies-access-scopes"
target="_blank" rel="noopener">access scope</a> is limited to accesses
in the pipeline stages determined by the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-pipeline-stages-masks"
target="_blank" rel="noopener">destination stage mask</a> specified by
`dstStageMask`. It is also limited to access types in the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-access-masks"
target="_blank" rel="noopener">destination access mask</a> specified by
`dstAccessMask`.

The <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies-available-and-visible"
target="_blank" rel="noopener">availability and visibility
operations</a> defined by a subpass dependency affect the execution of
<a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass-layout-transitions"
target="_blank" rel="noopener">image layout transitions</a> within the
render pass.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>For non-attachment resources, the memory dependency expressed by
subpass dependency is nearly identical to that of a <a
href="VkMemoryBarrier.html">VkMemoryBarrier</a> (with matching
<code>srcAccessMask</code> and <code>dstAccessMask</code> parameters)
submitted as a part of a <a
href="vkCmdPipelineBarrier.html">vkCmdPipelineBarrier</a> (with matching
<code>srcStageMask</code> and <code>dstStageMask</code> parameters). The
only difference being that its scopes are limited to the identified
subpasses rather than potentially affecting everything before and
after.</p>
<p>For attachments however, subpass dependencies work more like a <a
href="VkImageMemoryBarrier.html">VkImageMemoryBarrier</a> defined
similarly to the <a href="VkMemoryBarrier.html">VkMemoryBarrier</a>
above, the queue family indices set to
<code>VK_QUEUE_FAMILY_IGNORED</code>, and layouts as follows:</p>
<ul>
<li><p>The equivalent to <code>oldLayout</code> is the attachment’s
layout according to the subpass description for
<code>srcSubpass</code>.</p></li>
<li><p>The equivalent to <code>newLayout</code> is the attachment’s
layout according to the subpass description for
<code>dstSubpass</code>.</p></li>
</ul></td>
</tr>
</tbody>
</table>

Valid Usage

- <a href="#VUID-VkSubpassDependency-srcStageMask-04090"
  id="VUID-VkSubpassDependency-srcStageMask-04090"></a>
  VUID-VkSubpassDependency-srcStageMask-04090  
  If the [`geometryShader`](#features-geometryShader) feature is not
  enabled, `srcStageMask` **must** not contain
  `VK_PIPELINE_STAGE_GEOMETRY_SHADER_BIT`

- <a href="#VUID-VkSubpassDependency-srcStageMask-04091"
  id="VUID-VkSubpassDependency-srcStageMask-04091"></a>
  VUID-VkSubpassDependency-srcStageMask-04091  
  If the [`tessellationShader`](#features-tessellationShader) feature is
  not enabled, `srcStageMask` **must** not contain
  `VK_PIPELINE_STAGE_TESSELLATION_CONTROL_SHADER_BIT` or
  `VK_PIPELINE_STAGE_TESSELLATION_EVALUATION_SHADER_BIT`

- <a href="#VUID-VkSubpassDependency-srcStageMask-04092"
  id="VUID-VkSubpassDependency-srcStageMask-04092"></a>
  VUID-VkSubpassDependency-srcStageMask-04092  
  If the [`conditionalRendering`](#features-conditionalRendering)
  feature is not enabled, `srcStageMask` **must** not contain
  `VK_PIPELINE_STAGE_CONDITIONAL_RENDERING_BIT_EXT`

- <a href="#VUID-VkSubpassDependency-srcStageMask-04093"
  id="VUID-VkSubpassDependency-srcStageMask-04093"></a>
  VUID-VkSubpassDependency-srcStageMask-04093  
  If the [`fragmentDensityMap`](#features-fragmentDensityMap) feature is
  not enabled, `srcStageMask` **must** not contain
  `VK_PIPELINE_STAGE_FRAGMENT_DENSITY_PROCESS_BIT_EXT`

- <a href="#VUID-VkSubpassDependency-srcStageMask-04094"
  id="VUID-VkSubpassDependency-srcStageMask-04094"></a>
  VUID-VkSubpassDependency-srcStageMask-04094  
  If the [`transformFeedback`](#features-transformFeedback) feature is
  not enabled, `srcStageMask` **must** not contain
  `VK_PIPELINE_STAGE_TRANSFORM_FEEDBACK_BIT_EXT`

- <a href="#VUID-VkSubpassDependency-srcStageMask-04095"
  id="VUID-VkSubpassDependency-srcStageMask-04095"></a>
  VUID-VkSubpassDependency-srcStageMask-04095  
  If the [`meshShader`](#features-meshShader) feature is not enabled,
  `srcStageMask` **must** not contain
  `VK_PIPELINE_STAGE_MESH_SHADER_BIT_EXT`

- <a href="#VUID-VkSubpassDependency-srcStageMask-04096"
  id="VUID-VkSubpassDependency-srcStageMask-04096"></a>
  VUID-VkSubpassDependency-srcStageMask-04096  
  If the [`taskShader`](#features-taskShader) feature is not enabled,
  `srcStageMask` **must** not contain
  `VK_PIPELINE_STAGE_TASK_SHADER_BIT_EXT`

- <a href="#VUID-VkSubpassDependency-srcStageMask-07318"
  id="VUID-VkSubpassDependency-srcStageMask-07318"></a>
  VUID-VkSubpassDependency-srcStageMask-07318  
  If neither the [`shadingRateImage`](#features-shadingRateImage) or
  [`attachmentFragmentShadingRate`](#features-attachmentFragmentShadingRate)
  are enabled, `srcStageMask` **must** not contain
  `VK_PIPELINE_STAGE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR`

- <a href="#VUID-VkSubpassDependency-srcStageMask-03937"
  id="VUID-VkSubpassDependency-srcStageMask-03937"></a>
  VUID-VkSubpassDependency-srcStageMask-03937  
  If the [`synchronization2`](#features-synchronization2) feature is not
  enabled, `srcStageMask` **must** not be `0`

- <a href="#VUID-VkSubpassDependency-srcStageMask-07949"
  id="VUID-VkSubpassDependency-srcStageMask-07949"></a>
  VUID-VkSubpassDependency-srcStageMask-07949  
  If neither the [VK_NV_ray_tracing](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_ray_tracing.html) extension
  or [`rayTracingPipeline` feature](#features-rayTracingPipeline) are
  enabled, `srcStageMask` **must** not contain
  `VK_PIPELINE_STAGE_RAY_TRACING_SHADER_BIT_KHR`

<!-- -->

- <a href="#VUID-VkSubpassDependency-dstStageMask-04090"
  id="VUID-VkSubpassDependency-dstStageMask-04090"></a>
  VUID-VkSubpassDependency-dstStageMask-04090  
  If the [`geometryShader`](#features-geometryShader) feature is not
  enabled, `dstStageMask` **must** not contain
  `VK_PIPELINE_STAGE_GEOMETRY_SHADER_BIT`

- <a href="#VUID-VkSubpassDependency-dstStageMask-04091"
  id="VUID-VkSubpassDependency-dstStageMask-04091"></a>
  VUID-VkSubpassDependency-dstStageMask-04091  
  If the [`tessellationShader`](#features-tessellationShader) feature is
  not enabled, `dstStageMask` **must** not contain
  `VK_PIPELINE_STAGE_TESSELLATION_CONTROL_SHADER_BIT` or
  `VK_PIPELINE_STAGE_TESSELLATION_EVALUATION_SHADER_BIT`

- <a href="#VUID-VkSubpassDependency-dstStageMask-04092"
  id="VUID-VkSubpassDependency-dstStageMask-04092"></a>
  VUID-VkSubpassDependency-dstStageMask-04092  
  If the [`conditionalRendering`](#features-conditionalRendering)
  feature is not enabled, `dstStageMask` **must** not contain
  `VK_PIPELINE_STAGE_CONDITIONAL_RENDERING_BIT_EXT`

- <a href="#VUID-VkSubpassDependency-dstStageMask-04093"
  id="VUID-VkSubpassDependency-dstStageMask-04093"></a>
  VUID-VkSubpassDependency-dstStageMask-04093  
  If the [`fragmentDensityMap`](#features-fragmentDensityMap) feature is
  not enabled, `dstStageMask` **must** not contain
  `VK_PIPELINE_STAGE_FRAGMENT_DENSITY_PROCESS_BIT_EXT`

- <a href="#VUID-VkSubpassDependency-dstStageMask-04094"
  id="VUID-VkSubpassDependency-dstStageMask-04094"></a>
  VUID-VkSubpassDependency-dstStageMask-04094  
  If the [`transformFeedback`](#features-transformFeedback) feature is
  not enabled, `dstStageMask` **must** not contain
  `VK_PIPELINE_STAGE_TRANSFORM_FEEDBACK_BIT_EXT`

- <a href="#VUID-VkSubpassDependency-dstStageMask-04095"
  id="VUID-VkSubpassDependency-dstStageMask-04095"></a>
  VUID-VkSubpassDependency-dstStageMask-04095  
  If the [`meshShader`](#features-meshShader) feature is not enabled,
  `dstStageMask` **must** not contain
  `VK_PIPELINE_STAGE_MESH_SHADER_BIT_EXT`

- <a href="#VUID-VkSubpassDependency-dstStageMask-04096"
  id="VUID-VkSubpassDependency-dstStageMask-04096"></a>
  VUID-VkSubpassDependency-dstStageMask-04096  
  If the [`taskShader`](#features-taskShader) feature is not enabled,
  `dstStageMask` **must** not contain
  `VK_PIPELINE_STAGE_TASK_SHADER_BIT_EXT`

- <a href="#VUID-VkSubpassDependency-dstStageMask-07318"
  id="VUID-VkSubpassDependency-dstStageMask-07318"></a>
  VUID-VkSubpassDependency-dstStageMask-07318  
  If neither the [`shadingRateImage`](#features-shadingRateImage) or
  [`attachmentFragmentShadingRate`](#features-attachmentFragmentShadingRate)
  are enabled, `dstStageMask` **must** not contain
  `VK_PIPELINE_STAGE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR`

- <a href="#VUID-VkSubpassDependency-dstStageMask-03937"
  id="VUID-VkSubpassDependency-dstStageMask-03937"></a>
  VUID-VkSubpassDependency-dstStageMask-03937  
  If the [`synchronization2`](#features-synchronization2) feature is not
  enabled, `dstStageMask` **must** not be `0`

- <a href="#VUID-VkSubpassDependency-dstStageMask-07949"
  id="VUID-VkSubpassDependency-dstStageMask-07949"></a>
  VUID-VkSubpassDependency-dstStageMask-07949  
  If neither the [VK_NV_ray_tracing](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_ray_tracing.html) extension
  or [`rayTracingPipeline` feature](#features-rayTracingPipeline) are
  enabled, `dstStageMask` **must** not contain
  `VK_PIPELINE_STAGE_RAY_TRACING_SHADER_BIT_KHR`

- <a href="#VUID-VkSubpassDependency-srcSubpass-00864"
  id="VUID-VkSubpassDependency-srcSubpass-00864"></a>
  VUID-VkSubpassDependency-srcSubpass-00864  
  `srcSubpass` **must** be less than or equal to `dstSubpass`, unless
  one of them is `VK_SUBPASS_EXTERNAL`, to avoid cyclic dependencies and
  ensure a valid execution order

- <a href="#VUID-VkSubpassDependency-srcSubpass-00865"
  id="VUID-VkSubpassDependency-srcSubpass-00865"></a>
  VUID-VkSubpassDependency-srcSubpass-00865  
  `srcSubpass` and `dstSubpass` **must** not both be equal to
  `VK_SUBPASS_EXTERNAL`

- <a href="#VUID-VkSubpassDependency-srcSubpass-06809"
  id="VUID-VkSubpassDependency-srcSubpass-06809"></a>
  VUID-VkSubpassDependency-srcSubpass-06809  
  If `srcSubpass` is equal to `dstSubpass` and `srcStageMask` includes a
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-framebuffer-regions"
  target="_blank" rel="noopener">framebuffer-space stage</a>,
  `dstStageMask` **must** only contain <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-framebuffer-regions"
  target="_blank" rel="noopener">framebuffer-space stages</a>

- <a href="#VUID-VkSubpassDependency-srcAccessMask-00868"
  id="VUID-VkSubpassDependency-srcAccessMask-00868"></a>
  VUID-VkSubpassDependency-srcAccessMask-00868  
  Any access flag included in `srcAccessMask` **must** be supported by
  one of the pipeline stages in `srcStageMask`, as specified in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-access-types-supported"
  target="_blank" rel="noopener">table of supported access types</a>

- <a href="#VUID-VkSubpassDependency-dstAccessMask-00869"
  id="VUID-VkSubpassDependency-dstAccessMask-00869"></a>
  VUID-VkSubpassDependency-dstAccessMask-00869  
  Any access flag included in `dstAccessMask` **must** be supported by
  one of the pipeline stages in `dstStageMask`, as specified in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-access-types-supported"
  target="_blank" rel="noopener">table of supported access types</a>

- <a href="#VUID-VkSubpassDependency-srcSubpass-02243"
  id="VUID-VkSubpassDependency-srcSubpass-02243"></a>
  VUID-VkSubpassDependency-srcSubpass-02243  
  If `srcSubpass` equals `dstSubpass`, and `srcStageMask` and
  `dstStageMask` both include a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-framebuffer-regions"
  target="_blank" rel="noopener">framebuffer-space stage</a>, then
  `dependencyFlags` **must** include `VK_DEPENDENCY_BY_REGION_BIT`

- <a href="#VUID-VkSubpassDependency-dependencyFlags-02520"
  id="VUID-VkSubpassDependency-dependencyFlags-02520"></a>
  VUID-VkSubpassDependency-dependencyFlags-02520  
  If `dependencyFlags` includes `VK_DEPENDENCY_VIEW_LOCAL_BIT`,
  `srcSubpass` **must** not be equal to `VK_SUBPASS_EXTERNAL`

- <a href="#VUID-VkSubpassDependency-dependencyFlags-02521"
  id="VUID-VkSubpassDependency-dependencyFlags-02521"></a>
  VUID-VkSubpassDependency-dependencyFlags-02521  
  If `dependencyFlags` includes `VK_DEPENDENCY_VIEW_LOCAL_BIT`,
  `dstSubpass` **must** not be equal to `VK_SUBPASS_EXTERNAL`

- <a href="#VUID-VkSubpassDependency-srcSubpass-00872"
  id="VUID-VkSubpassDependency-srcSubpass-00872"></a>
  VUID-VkSubpassDependency-srcSubpass-00872  
  If `srcSubpass` equals `dstSubpass` and that subpass has more than one
  bit set in the view mask, then `dependencyFlags` **must** include
  `VK_DEPENDENCY_VIEW_LOCAL_BIT`

Valid Usage (Implicit)

- <a href="#VUID-VkSubpassDependency-srcStageMask-parameter"
  id="VUID-VkSubpassDependency-srcStageMask-parameter"></a>
  VUID-VkSubpassDependency-srcStageMask-parameter  
  `srcStageMask` **must** be a valid combination of
  [VkPipelineStageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlagBits.html) values

- <a href="#VUID-VkSubpassDependency-dstStageMask-parameter"
  id="VUID-VkSubpassDependency-dstStageMask-parameter"></a>
  VUID-VkSubpassDependency-dstStageMask-parameter  
  `dstStageMask` **must** be a valid combination of
  [VkPipelineStageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlagBits.html) values

- <a href="#VUID-VkSubpassDependency-srcAccessMask-parameter"
  id="VUID-VkSubpassDependency-srcAccessMask-parameter"></a>
  VUID-VkSubpassDependency-srcAccessMask-parameter  
  `srcAccessMask` **must** be a valid combination of
  [VkAccessFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccessFlagBits.html) values

- <a href="#VUID-VkSubpassDependency-dstAccessMask-parameter"
  id="VUID-VkSubpassDependency-dstAccessMask-parameter"></a>
  VUID-VkSubpassDependency-dstAccessMask-parameter  
  `dstAccessMask` **must** be a valid combination of
  [VkAccessFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccessFlagBits.html) values

- <a href="#VUID-VkSubpassDependency-dependencyFlags-parameter"
  id="VUID-VkSubpassDependency-dependencyFlags-parameter"></a>
  VUID-VkSubpassDependency-dependencyFlags-parameter  
  `dependencyFlags` **must** be a valid combination of
  [VkDependencyFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDependencyFlagBits.html) values

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkAccessFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccessFlags.html),
[VkDependencyFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDependencyFlags.html),
[VkPipelineStageFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlags.html),
[VkRenderPassCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassCreateInfo.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSubpassDependency"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
