# VkSubpassDependency2(3) Manual Page

## Name

VkSubpassDependency2 - Structure specifying a subpass dependency



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkSubpassDependency2` structure is defined as:

``` c
// Provided by VK_VERSION_1_2
typedef struct VkSubpassDependency2 {
    VkStructureType         sType;
    const void*             pNext;
    uint32_t                srcSubpass;
    uint32_t                dstSubpass;
    VkPipelineStageFlags    srcStageMask;
    VkPipelineStageFlags    dstStageMask;
    VkAccessFlags           srcAccessMask;
    VkAccessFlags           dstAccessMask;
    VkDependencyFlags       dependencyFlags;
    int32_t                 viewOffset;
} VkSubpassDependency2;
```

or the equivalent

``` c
// Provided by VK_KHR_create_renderpass2
typedef VkSubpassDependency2 VkSubpassDependency2KHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

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

- `viewOffset` controls which views in the source subpass the views in
  the destination subpass depend on.

## <a href="#_description" class="anchor"></a>Description

Parameters defined by this structure with the same name as those in
[VkSubpassDependency](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassDependency.html) have the identical
effect to those parameters.

`viewOffset` has the same effect for the described subpass dependency as
[VkRenderPassMultiviewCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassMultiviewCreateInfo.html)::`pViewOffsets`
has on each corresponding subpass dependency.

If a [VkMemoryBarrier2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryBarrier2.html) is included in the
`pNext` chain, `srcStageMask`, `dstStageMask`, `srcAccessMask`, and
`dstAccessMask` parameters are ignored. The synchronization and access
scopes instead are defined by the parameters of
[VkMemoryBarrier2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryBarrier2.html).

Valid Usage

- <a href="#VUID-VkSubpassDependency2-srcStageMask-04090"
  id="VUID-VkSubpassDependency2-srcStageMask-04090"></a>
  VUID-VkSubpassDependency2-srcStageMask-04090  
  If the [`geometryShader`](#features-geometryShader) feature is not
  enabled, `srcStageMask` **must** not contain
  `VK_PIPELINE_STAGE_GEOMETRY_SHADER_BIT`

- <a href="#VUID-VkSubpassDependency2-srcStageMask-04091"
  id="VUID-VkSubpassDependency2-srcStageMask-04091"></a>
  VUID-VkSubpassDependency2-srcStageMask-04091  
  If the [`tessellationShader`](#features-tessellationShader) feature is
  not enabled, `srcStageMask` **must** not contain
  `VK_PIPELINE_STAGE_TESSELLATION_CONTROL_SHADER_BIT` or
  `VK_PIPELINE_STAGE_TESSELLATION_EVALUATION_SHADER_BIT`

- <a href="#VUID-VkSubpassDependency2-srcStageMask-04092"
  id="VUID-VkSubpassDependency2-srcStageMask-04092"></a>
  VUID-VkSubpassDependency2-srcStageMask-04092  
  If the [`conditionalRendering`](#features-conditionalRendering)
  feature is not enabled, `srcStageMask` **must** not contain
  `VK_PIPELINE_STAGE_CONDITIONAL_RENDERING_BIT_EXT`

- <a href="#VUID-VkSubpassDependency2-srcStageMask-04093"
  id="VUID-VkSubpassDependency2-srcStageMask-04093"></a>
  VUID-VkSubpassDependency2-srcStageMask-04093  
  If the [`fragmentDensityMap`](#features-fragmentDensityMap) feature is
  not enabled, `srcStageMask` **must** not contain
  `VK_PIPELINE_STAGE_FRAGMENT_DENSITY_PROCESS_BIT_EXT`

- <a href="#VUID-VkSubpassDependency2-srcStageMask-04094"
  id="VUID-VkSubpassDependency2-srcStageMask-04094"></a>
  VUID-VkSubpassDependency2-srcStageMask-04094  
  If the [`transformFeedback`](#features-transformFeedback) feature is
  not enabled, `srcStageMask` **must** not contain
  `VK_PIPELINE_STAGE_TRANSFORM_FEEDBACK_BIT_EXT`

- <a href="#VUID-VkSubpassDependency2-srcStageMask-04095"
  id="VUID-VkSubpassDependency2-srcStageMask-04095"></a>
  VUID-VkSubpassDependency2-srcStageMask-04095  
  If the [`meshShader`](#features-meshShader) feature is not enabled,
  `srcStageMask` **must** not contain
  `VK_PIPELINE_STAGE_MESH_SHADER_BIT_EXT`

- <a href="#VUID-VkSubpassDependency2-srcStageMask-04096"
  id="VUID-VkSubpassDependency2-srcStageMask-04096"></a>
  VUID-VkSubpassDependency2-srcStageMask-04096  
  If the [`taskShader`](#features-taskShader) feature is not enabled,
  `srcStageMask` **must** not contain
  `VK_PIPELINE_STAGE_TASK_SHADER_BIT_EXT`

- <a href="#VUID-VkSubpassDependency2-srcStageMask-07318"
  id="VUID-VkSubpassDependency2-srcStageMask-07318"></a>
  VUID-VkSubpassDependency2-srcStageMask-07318  
  If neither the [`shadingRateImage`](#features-shadingRateImage) or
  [`attachmentFragmentShadingRate`](#features-attachmentFragmentShadingRate)
  are enabled, `srcStageMask` **must** not contain
  `VK_PIPELINE_STAGE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR`

- <a href="#VUID-VkSubpassDependency2-srcStageMask-03937"
  id="VUID-VkSubpassDependency2-srcStageMask-03937"></a>
  VUID-VkSubpassDependency2-srcStageMask-03937  
  If the [`synchronization2`](#features-synchronization2) feature is not
  enabled, `srcStageMask` **must** not be `0`

- <a href="#VUID-VkSubpassDependency2-srcStageMask-07949"
  id="VUID-VkSubpassDependency2-srcStageMask-07949"></a>
  VUID-VkSubpassDependency2-srcStageMask-07949  
  If neither the [VK_NV_ray_tracing](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_ray_tracing.html) extension
  or [`rayTracingPipeline` feature](#features-rayTracingPipeline) are
  enabled, `srcStageMask` **must** not contain
  `VK_PIPELINE_STAGE_RAY_TRACING_SHADER_BIT_KHR`

<!-- -->

- <a href="#VUID-VkSubpassDependency2-dstStageMask-04090"
  id="VUID-VkSubpassDependency2-dstStageMask-04090"></a>
  VUID-VkSubpassDependency2-dstStageMask-04090  
  If the [`geometryShader`](#features-geometryShader) feature is not
  enabled, `dstStageMask` **must** not contain
  `VK_PIPELINE_STAGE_GEOMETRY_SHADER_BIT`

- <a href="#VUID-VkSubpassDependency2-dstStageMask-04091"
  id="VUID-VkSubpassDependency2-dstStageMask-04091"></a>
  VUID-VkSubpassDependency2-dstStageMask-04091  
  If the [`tessellationShader`](#features-tessellationShader) feature is
  not enabled, `dstStageMask` **must** not contain
  `VK_PIPELINE_STAGE_TESSELLATION_CONTROL_SHADER_BIT` or
  `VK_PIPELINE_STAGE_TESSELLATION_EVALUATION_SHADER_BIT`

- <a href="#VUID-VkSubpassDependency2-dstStageMask-04092"
  id="VUID-VkSubpassDependency2-dstStageMask-04092"></a>
  VUID-VkSubpassDependency2-dstStageMask-04092  
  If the [`conditionalRendering`](#features-conditionalRendering)
  feature is not enabled, `dstStageMask` **must** not contain
  `VK_PIPELINE_STAGE_CONDITIONAL_RENDERING_BIT_EXT`

- <a href="#VUID-VkSubpassDependency2-dstStageMask-04093"
  id="VUID-VkSubpassDependency2-dstStageMask-04093"></a>
  VUID-VkSubpassDependency2-dstStageMask-04093  
  If the [`fragmentDensityMap`](#features-fragmentDensityMap) feature is
  not enabled, `dstStageMask` **must** not contain
  `VK_PIPELINE_STAGE_FRAGMENT_DENSITY_PROCESS_BIT_EXT`

- <a href="#VUID-VkSubpassDependency2-dstStageMask-04094"
  id="VUID-VkSubpassDependency2-dstStageMask-04094"></a>
  VUID-VkSubpassDependency2-dstStageMask-04094  
  If the [`transformFeedback`](#features-transformFeedback) feature is
  not enabled, `dstStageMask` **must** not contain
  `VK_PIPELINE_STAGE_TRANSFORM_FEEDBACK_BIT_EXT`

- <a href="#VUID-VkSubpassDependency2-dstStageMask-04095"
  id="VUID-VkSubpassDependency2-dstStageMask-04095"></a>
  VUID-VkSubpassDependency2-dstStageMask-04095  
  If the [`meshShader`](#features-meshShader) feature is not enabled,
  `dstStageMask` **must** not contain
  `VK_PIPELINE_STAGE_MESH_SHADER_BIT_EXT`

- <a href="#VUID-VkSubpassDependency2-dstStageMask-04096"
  id="VUID-VkSubpassDependency2-dstStageMask-04096"></a>
  VUID-VkSubpassDependency2-dstStageMask-04096  
  If the [`taskShader`](#features-taskShader) feature is not enabled,
  `dstStageMask` **must** not contain
  `VK_PIPELINE_STAGE_TASK_SHADER_BIT_EXT`

- <a href="#VUID-VkSubpassDependency2-dstStageMask-07318"
  id="VUID-VkSubpassDependency2-dstStageMask-07318"></a>
  VUID-VkSubpassDependency2-dstStageMask-07318  
  If neither the [`shadingRateImage`](#features-shadingRateImage) or
  [`attachmentFragmentShadingRate`](#features-attachmentFragmentShadingRate)
  are enabled, `dstStageMask` **must** not contain
  `VK_PIPELINE_STAGE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR`

- <a href="#VUID-VkSubpassDependency2-dstStageMask-03937"
  id="VUID-VkSubpassDependency2-dstStageMask-03937"></a>
  VUID-VkSubpassDependency2-dstStageMask-03937  
  If the [`synchronization2`](#features-synchronization2) feature is not
  enabled, `dstStageMask` **must** not be `0`

- <a href="#VUID-VkSubpassDependency2-dstStageMask-07949"
  id="VUID-VkSubpassDependency2-dstStageMask-07949"></a>
  VUID-VkSubpassDependency2-dstStageMask-07949  
  If neither the [VK_NV_ray_tracing](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_ray_tracing.html) extension
  or [`rayTracingPipeline` feature](#features-rayTracingPipeline) are
  enabled, `dstStageMask` **must** not contain
  `VK_PIPELINE_STAGE_RAY_TRACING_SHADER_BIT_KHR`

- <a href="#VUID-VkSubpassDependency2-srcSubpass-03084"
  id="VUID-VkSubpassDependency2-srcSubpass-03084"></a>
  VUID-VkSubpassDependency2-srcSubpass-03084  
  `srcSubpass` **must** be less than or equal to `dstSubpass`, unless
  one of them is `VK_SUBPASS_EXTERNAL`, to avoid cyclic dependencies and
  ensure a valid execution order

- <a href="#VUID-VkSubpassDependency2-srcSubpass-03085"
  id="VUID-VkSubpassDependency2-srcSubpass-03085"></a>
  VUID-VkSubpassDependency2-srcSubpass-03085  
  `srcSubpass` and `dstSubpass` **must** not both be equal to
  `VK_SUBPASS_EXTERNAL`

- <a href="#VUID-VkSubpassDependency2-srcSubpass-06810"
  id="VUID-VkSubpassDependency2-srcSubpass-06810"></a>
  VUID-VkSubpassDependency2-srcSubpass-06810  
  If `srcSubpass` is equal to `dstSubpass` and `srcStageMask` includes a
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-framebuffer-regions"
  target="_blank" rel="noopener">framebuffer-space stage</a>,
  `dstStageMask` **must** only contain <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-framebuffer-regions"
  target="_blank" rel="noopener">framebuffer-space stages</a>

- <a href="#VUID-VkSubpassDependency2-srcAccessMask-03088"
  id="VUID-VkSubpassDependency2-srcAccessMask-03088"></a>
  VUID-VkSubpassDependency2-srcAccessMask-03088  
  Any access flag included in `srcAccessMask` **must** be supported by
  one of the pipeline stages in `srcStageMask`, as specified in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-access-types-supported"
  target="_blank" rel="noopener">table of supported access types</a>

- <a href="#VUID-VkSubpassDependency2-dstAccessMask-03089"
  id="VUID-VkSubpassDependency2-dstAccessMask-03089"></a>
  VUID-VkSubpassDependency2-dstAccessMask-03089  
  Any access flag included in `dstAccessMask` **must** be supported by
  one of the pipeline stages in `dstStageMask`, as specified in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-access-types-supported"
  target="_blank" rel="noopener">table of supported access types</a>

- <a href="#VUID-VkSubpassDependency2-dependencyFlags-03090"
  id="VUID-VkSubpassDependency2-dependencyFlags-03090"></a>
  VUID-VkSubpassDependency2-dependencyFlags-03090  
  If `dependencyFlags` includes `VK_DEPENDENCY_VIEW_LOCAL_BIT`,
  `srcSubpass` **must** not be equal to `VK_SUBPASS_EXTERNAL`

- <a href="#VUID-VkSubpassDependency2-dependencyFlags-03091"
  id="VUID-VkSubpassDependency2-dependencyFlags-03091"></a>
  VUID-VkSubpassDependency2-dependencyFlags-03091  
  If `dependencyFlags` includes `VK_DEPENDENCY_VIEW_LOCAL_BIT`,
  `dstSubpass` **must** not be equal to `VK_SUBPASS_EXTERNAL`

- <a href="#VUID-VkSubpassDependency2-srcSubpass-02245"
  id="VUID-VkSubpassDependency2-srcSubpass-02245"></a>
  VUID-VkSubpassDependency2-srcSubpass-02245  
  If `srcSubpass` equals `dstSubpass`, and `srcStageMask` and
  `dstStageMask` both include a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-framebuffer-regions"
  target="_blank" rel="noopener">framebuffer-space stage</a>, then
  `dependencyFlags` **must** include `VK_DEPENDENCY_BY_REGION_BIT`

- <a href="#VUID-VkSubpassDependency2-viewOffset-02530"
  id="VUID-VkSubpassDependency2-viewOffset-02530"></a>
  VUID-VkSubpassDependency2-viewOffset-02530  
  If `viewOffset` is not equal to `0`, `srcSubpass` **must** not be
  equal to `dstSubpass`

- <a href="#VUID-VkSubpassDependency2-dependencyFlags-03092"
  id="VUID-VkSubpassDependency2-dependencyFlags-03092"></a>
  VUID-VkSubpassDependency2-dependencyFlags-03092  
  If `dependencyFlags` does not include `VK_DEPENDENCY_VIEW_LOCAL_BIT`,
  `viewOffset` **must** be `0`

Valid Usage (Implicit)

- <a href="#VUID-VkSubpassDependency2-sType-sType"
  id="VUID-VkSubpassDependency2-sType-sType"></a>
  VUID-VkSubpassDependency2-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SUBPASS_DEPENDENCY_2`

- <a href="#VUID-VkSubpassDependency2-pNext-pNext"
  id="VUID-VkSubpassDependency2-pNext-pNext"></a>
  VUID-VkSubpassDependency2-pNext-pNext  
  `pNext` **must** be `NULL` or a pointer to a valid instance of
  [VkMemoryBarrier2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryBarrier2.html)

- <a href="#VUID-VkSubpassDependency2-sType-unique"
  id="VUID-VkSubpassDependency2-sType-unique"></a>
  VUID-VkSubpassDependency2-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

- <a href="#VUID-VkSubpassDependency2-srcStageMask-parameter"
  id="VUID-VkSubpassDependency2-srcStageMask-parameter"></a>
  VUID-VkSubpassDependency2-srcStageMask-parameter  
  `srcStageMask` **must** be a valid combination of
  [VkPipelineStageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlagBits.html) values

- <a href="#VUID-VkSubpassDependency2-dstStageMask-parameter"
  id="VUID-VkSubpassDependency2-dstStageMask-parameter"></a>
  VUID-VkSubpassDependency2-dstStageMask-parameter  
  `dstStageMask` **must** be a valid combination of
  [VkPipelineStageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlagBits.html) values

- <a href="#VUID-VkSubpassDependency2-srcAccessMask-parameter"
  id="VUID-VkSubpassDependency2-srcAccessMask-parameter"></a>
  VUID-VkSubpassDependency2-srcAccessMask-parameter  
  `srcAccessMask` **must** be a valid combination of
  [VkAccessFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccessFlagBits.html) values

- <a href="#VUID-VkSubpassDependency2-dstAccessMask-parameter"
  id="VUID-VkSubpassDependency2-dstAccessMask-parameter"></a>
  VUID-VkSubpassDependency2-dstAccessMask-parameter  
  `dstAccessMask` **must** be a valid combination of
  [VkAccessFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccessFlagBits.html) values

- <a href="#VUID-VkSubpassDependency2-dependencyFlags-parameter"
  id="VUID-VkSubpassDependency2-dependencyFlags-parameter"></a>
  VUID-VkSubpassDependency2-dependencyFlags-parameter  
  `dependencyFlags` **must** be a valid combination of
  [VkDependencyFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDependencyFlagBits.html) values

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_create_renderpass2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_create_renderpass2.html),
[VK_VERSION_1_2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_2.html),
[VkAccessFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccessFlags.html),
[VkDependencyFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDependencyFlags.html),
[VkPipelineStageFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlags.html),
[VkRenderPassCreateInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassCreateInfo2.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSubpassDependency2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
