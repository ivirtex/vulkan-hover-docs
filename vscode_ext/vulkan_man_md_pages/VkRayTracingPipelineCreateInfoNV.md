# VkRayTracingPipelineCreateInfoNV(3) Manual Page

## Name

VkRayTracingPipelineCreateInfoNV - Structure specifying parameters of a
newly created ray tracing pipeline



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkRayTracingPipelineCreateInfoNV` structure is defined as:

``` c
// Provided by VK_NV_ray_tracing
typedef struct VkRayTracingPipelineCreateInfoNV {
    VkStructureType                               sType;
    const void*                                   pNext;
    VkPipelineCreateFlags                         flags;
    uint32_t                                      stageCount;
    const VkPipelineShaderStageCreateInfo*        pStages;
    uint32_t                                      groupCount;
    const VkRayTracingShaderGroupCreateInfoNV*    pGroups;
    uint32_t                                      maxRecursionDepth;
    VkPipelineLayout                              layout;
    VkPipeline                                    basePipelineHandle;
    int32_t                                       basePipelineIndex;
} VkRayTracingPipelineCreateInfoNV;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is a bitmask of
  [VkPipelineCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlagBits.html) specifying
  how the pipeline will be generated.

- `stageCount` is the number of entries in the `pStages` array.

- `pStages` is a pointer to an array of
  [VkPipelineShaderStageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageCreateInfo.html)
  structures specifying the set of the shader stages to be included in
  the ray tracing pipeline.

- `groupCount` is the number of entries in the `pGroups` array.

- `pGroups` is a pointer to an array of
  [VkRayTracingShaderGroupCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingShaderGroupCreateInfoNV.html)
  structures describing the set of the shader stages to be included in
  each shader group in the ray tracing pipeline.

- `maxRecursionDepth` is the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#ray-tracing-recursion-depth"
  target="_blank" rel="noopener">maximum recursion depth</a> of shaders
  executed by this pipeline.

- `layout` is the description of binding locations used by both the
  pipeline and descriptor sets used with the pipeline.

- `basePipelineHandle` is a pipeline to derive from.

- `basePipelineIndex` is an index into the `pCreateInfos` parameter to
  use as a pipeline to derive from.

## <a href="#_description" class="anchor"></a>Description

The parameters `basePipelineHandle` and `basePipelineIndex` are
described in more detail in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-pipeline-derivatives"
target="_blank" rel="noopener">Pipeline Derivatives</a>.

If a
[VkPipelineCreateFlags2CreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlags2CreateInfoKHR.html)
structure is present in the `pNext` chain,
[VkPipelineCreateFlags2CreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlags2CreateInfoKHR.html)::`flags`
from that structure is used instead of `flags` from this structure.

Valid Usage

- <a href="#VUID-VkRayTracingPipelineCreateInfoNV-None-09497"
  id="VUID-VkRayTracingPipelineCreateInfoNV-None-09497"></a>
  VUID-VkRayTracingPipelineCreateInfoNV-None-09497  
  If the `pNext` chain does not include a
  [VkPipelineCreateFlags2CreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlags2CreateInfoKHR.html)
  structure, `flags` must be a valid combination of
  [VkPipelineCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlagBits.html) values

- <a href="#VUID-VkRayTracingPipelineCreateInfoNV-flags-07984"
  id="VUID-VkRayTracingPipelineCreateInfoNV-flags-07984"></a>
  VUID-VkRayTracingPipelineCreateInfoNV-flags-07984  
  If `flags` contains the `VK_PIPELINE_CREATE_DERIVATIVE_BIT` flag, and
  `basePipelineIndex` is -1, `basePipelineHandle` **must** be a valid
  ray tracing `VkPipeline` handle

- <a href="#VUID-VkRayTracingPipelineCreateInfoNV-flags-07985"
  id="VUID-VkRayTracingPipelineCreateInfoNV-flags-07985"></a>
  VUID-VkRayTracingPipelineCreateInfoNV-flags-07985  
  If `flags` contains the `VK_PIPELINE_CREATE_DERIVATIVE_BIT` flag, and
  `basePipelineHandle` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `basePipelineIndex` **must** be a valid index into the calling
  command’s `pCreateInfos` parameter

- <a href="#VUID-VkRayTracingPipelineCreateInfoNV-flags-07986"
  id="VUID-VkRayTracingPipelineCreateInfoNV-flags-07986"></a>
  VUID-VkRayTracingPipelineCreateInfoNV-flags-07986  
  If `flags` contains the `VK_PIPELINE_CREATE_DERIVATIVE_BIT` flag,
  `basePipelineIndex` **must** be -1 or `basePipelineHandle` **must** be
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

- <a href="#VUID-VkRayTracingPipelineCreateInfoNV-layout-07987"
  id="VUID-VkRayTracingPipelineCreateInfoNV-layout-07987"></a>
  VUID-VkRayTracingPipelineCreateInfoNV-layout-07987  
  If a push constant block is declared in a shader, a push constant
  range in `layout` **must** match both the shader stage and range

- <a href="#VUID-VkRayTracingPipelineCreateInfoNV-layout-07988"
  id="VUID-VkRayTracingPipelineCreateInfoNV-layout-07988"></a>
  VUID-VkRayTracingPipelineCreateInfoNV-layout-07988  
  If a [resource variables](#interfaces-resources) is declared in a
  shader, a descriptor slot in `layout` **must** match the shader stage

- <a href="#VUID-VkRayTracingPipelineCreateInfoNV-layout-07990"
  id="VUID-VkRayTracingPipelineCreateInfoNV-layout-07990"></a>
  VUID-VkRayTracingPipelineCreateInfoNV-layout-07990  
  If a [resource variables](#interfaces-resources) is declared in a
  shader, and the descriptor type is not
  `VK_DESCRIPTOR_TYPE_MUTABLE_EXT`, a descriptor slot in `layout`
  **must** match the descriptor type

- <a href="#VUID-VkRayTracingPipelineCreateInfoNV-layout-07991"
  id="VUID-VkRayTracingPipelineCreateInfoNV-layout-07991"></a>
  VUID-VkRayTracingPipelineCreateInfoNV-layout-07991  
  If a [resource variables](#interfaces-resources) is declared in a
  shader as an array, a descriptor slot in `layout` **must** match the
  descriptor count

<!-- -->

- <a href="#VUID-VkRayTracingPipelineCreateInfoNV-pStages-03426"
  id="VUID-VkRayTracingPipelineCreateInfoNV-pStages-03426"></a>
  VUID-VkRayTracingPipelineCreateInfoNV-pStages-03426  
  The shader code for the entry points identified by `pStages`, and the
  rest of the state identified by this structure **must** adhere to the
  pipeline linking rules described in the [Shader
  Interfaces](#interfaces) chapter

- <a href="#VUID-VkRayTracingPipelineCreateInfoNV-layout-03428"
  id="VUID-VkRayTracingPipelineCreateInfoNV-layout-03428"></a>
  VUID-VkRayTracingPipelineCreateInfoNV-layout-03428  
  The number of resources in `layout` accessible to each shader stage
  that is used by the pipeline **must** be less than or equal to
  [VkPhysicalDeviceLimits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceLimits.html)::`maxPerStageResources`

- <a href="#VUID-VkRayTracingPipelineCreateInfoNV-flags-02904"
  id="VUID-VkRayTracingPipelineCreateInfoNV-flags-02904"></a>
  VUID-VkRayTracingPipelineCreateInfoNV-flags-02904  
  `flags` **must** not include
  `VK_PIPELINE_CREATE_INDIRECT_BINDABLE_BIT_NV`

- <a
  href="#VUID-VkRayTracingPipelineCreateInfoNV-pipelineCreationCacheControl-02905"
  id="VUID-VkRayTracingPipelineCreateInfoNV-pipelineCreationCacheControl-02905"></a>
  VUID-VkRayTracingPipelineCreateInfoNV-pipelineCreationCacheControl-02905  
  If the
  [`pipelineCreationCacheControl`](#features-pipelineCreationCacheControl)
  feature is not enabled, `flags` **must** not include
  `VK_PIPELINE_CREATE_FAIL_ON_PIPELINE_COMPILE_REQUIRED_BIT` or
  `VK_PIPELINE_CREATE_EARLY_RETURN_ON_FAILURE_BIT`

- <a href="#VUID-VkRayTracingPipelineCreateInfoNV-stage-06232"
  id="VUID-VkRayTracingPipelineCreateInfoNV-stage-06232"></a>
  VUID-VkRayTracingPipelineCreateInfoNV-stage-06232  
  The `stage` member of at least one element of `pStages` **must** be
  `VK_SHADER_STAGE_RAYGEN_BIT_KHR`

- <a href="#VUID-VkRayTracingPipelineCreateInfoNV-flags-03456"
  id="VUID-VkRayTracingPipelineCreateInfoNV-flags-03456"></a>
  VUID-VkRayTracingPipelineCreateInfoNV-flags-03456  
  `flags` **must** not include `VK_PIPELINE_CREATE_LIBRARY_BIT_KHR`

- <a href="#VUID-VkRayTracingPipelineCreateInfoNV-maxRecursionDepth-03457"
  id="VUID-VkRayTracingPipelineCreateInfoNV-maxRecursionDepth-03457"></a>
  VUID-VkRayTracingPipelineCreateInfoNV-maxRecursionDepth-03457  
  `maxRecursionDepth` **must** be less than or equal to
  [VkPhysicalDeviceRayTracingPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceRayTracingPropertiesNV.html)::`maxRecursionDepth`

- <a href="#VUID-VkRayTracingPipelineCreateInfoNV-flags-03458"
  id="VUID-VkRayTracingPipelineCreateInfoNV-flags-03458"></a>
  VUID-VkRayTracingPipelineCreateInfoNV-flags-03458  
  `flags` **must** not include
  `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_ANY_HIT_SHADERS_BIT_KHR`

- <a href="#VUID-VkRayTracingPipelineCreateInfoNV-flags-03459"
  id="VUID-VkRayTracingPipelineCreateInfoNV-flags-03459"></a>
  VUID-VkRayTracingPipelineCreateInfoNV-flags-03459  
  `flags` **must** not include
  `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_CLOSEST_HIT_SHADERS_BIT_KHR`

- <a href="#VUID-VkRayTracingPipelineCreateInfoNV-flags-03460"
  id="VUID-VkRayTracingPipelineCreateInfoNV-flags-03460"></a>
  VUID-VkRayTracingPipelineCreateInfoNV-flags-03460  
  `flags` **must** not include
  `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_MISS_SHADERS_BIT_KHR`

- <a href="#VUID-VkRayTracingPipelineCreateInfoNV-flags-03461"
  id="VUID-VkRayTracingPipelineCreateInfoNV-flags-03461"></a>
  VUID-VkRayTracingPipelineCreateInfoNV-flags-03461  
  `flags` **must** not include
  `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_INTERSECTION_SHADERS_BIT_KHR`

- <a href="#VUID-VkRayTracingPipelineCreateInfoNV-flags-03462"
  id="VUID-VkRayTracingPipelineCreateInfoNV-flags-03462"></a>
  VUID-VkRayTracingPipelineCreateInfoNV-flags-03462  
  `flags` **must** not include
  `VK_PIPELINE_CREATE_RAY_TRACING_SKIP_AABBS_BIT_KHR`

- <a href="#VUID-VkRayTracingPipelineCreateInfoNV-flags-03463"
  id="VUID-VkRayTracingPipelineCreateInfoNV-flags-03463"></a>
  VUID-VkRayTracingPipelineCreateInfoNV-flags-03463  
  `flags` **must** not include
  `VK_PIPELINE_CREATE_RAY_TRACING_SKIP_TRIANGLES_BIT_KHR`

- <a href="#VUID-VkRayTracingPipelineCreateInfoNV-flags-03588"
  id="VUID-VkRayTracingPipelineCreateInfoNV-flags-03588"></a>
  VUID-VkRayTracingPipelineCreateInfoNV-flags-03588  
  `flags` **must** not include
  `VK_PIPELINE_CREATE_RAY_TRACING_SHADER_GROUP_HANDLE_CAPTURE_REPLAY_BIT_KHR`

- <a href="#VUID-VkRayTracingPipelineCreateInfoNV-flags-04948"
  id="VUID-VkRayTracingPipelineCreateInfoNV-flags-04948"></a>
  VUID-VkRayTracingPipelineCreateInfoNV-flags-04948  
  `flags` **must** not include
  `VK_PIPELINE_CREATE_RAY_TRACING_ALLOW_MOTION_BIT_NV`

- <a href="#VUID-VkRayTracingPipelineCreateInfoNV-flags-02957"
  id="VUID-VkRayTracingPipelineCreateInfoNV-flags-02957"></a>
  VUID-VkRayTracingPipelineCreateInfoNV-flags-02957  
  `flags` **must** not include both
  `VK_PIPELINE_CREATE_DEFER_COMPILE_BIT_NV` and
  `VK_PIPELINE_CREATE_FAIL_ON_PIPELINE_COMPILE_REQUIRED_BIT` at the same
  time

- <a
  href="#VUID-VkRayTracingPipelineCreateInfoNV-pipelineStageCreationFeedbackCount-06651"
  id="VUID-VkRayTracingPipelineCreateInfoNV-pipelineStageCreationFeedbackCount-06651"></a>
  VUID-VkRayTracingPipelineCreateInfoNV-pipelineStageCreationFeedbackCount-06651  
  If
  [VkPipelineCreationFeedbackCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreationFeedbackCreateInfo.html)::`pipelineStageCreationFeedbackCount`
  is not `0`, it **must** be equal to `stageCount`

- <a href="#VUID-VkRayTracingPipelineCreateInfoNV-stage-06898"
  id="VUID-VkRayTracingPipelineCreateInfoNV-stage-06898"></a>
  VUID-VkRayTracingPipelineCreateInfoNV-stage-06898  
  The `stage` value in all `pStages` elements **must** be one of
  `VK_SHADER_STAGE_RAYGEN_BIT_KHR`, `VK_SHADER_STAGE_ANY_HIT_BIT_KHR`,
  `VK_SHADER_STAGE_CLOSEST_HIT_BIT_KHR`, `VK_SHADER_STAGE_MISS_BIT_KHR`,
  `VK_SHADER_STAGE_INTERSECTION_BIT_KHR`, or
  `VK_SHADER_STAGE_CALLABLE_BIT_KHR`

- <a href="#VUID-VkRayTracingPipelineCreateInfoNV-flags-07402"
  id="VUID-VkRayTracingPipelineCreateInfoNV-flags-07402"></a>
  VUID-VkRayTracingPipelineCreateInfoNV-flags-07402  
  `flags` **must** not include
  `VK_PIPELINE_CREATE_RAY_TRACING_OPACITY_MICROMAP_BIT_EXT`

- <a href="#VUID-VkRayTracingPipelineCreateInfoNV-flags-07998"
  id="VUID-VkRayTracingPipelineCreateInfoNV-flags-07998"></a>
  VUID-VkRayTracingPipelineCreateInfoNV-flags-07998  
  `flags` **must** not include
  `VK_PIPELINE_CREATE_RAY_TRACING_DISPLACEMENT_MICROMAP_BIT_NV`

Valid Usage (Implicit)

- <a href="#VUID-VkRayTracingPipelineCreateInfoNV-sType-sType"
  id="VUID-VkRayTracingPipelineCreateInfoNV-sType-sType"></a>
  VUID-VkRayTracingPipelineCreateInfoNV-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_RAY_TRACING_PIPELINE_CREATE_INFO_NV`

- <a href="#VUID-VkRayTracingPipelineCreateInfoNV-pNext-pNext"
  id="VUID-VkRayTracingPipelineCreateInfoNV-pNext-pNext"></a>
  VUID-VkRayTracingPipelineCreateInfoNV-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the
  `pNext` chain **must** be either `NULL` or a pointer to a valid
  instance of
  [VkPipelineCreateFlags2CreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlags2CreateInfoKHR.html)
  or
  [VkPipelineCreationFeedbackCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreationFeedbackCreateInfo.html)

- <a href="#VUID-VkRayTracingPipelineCreateInfoNV-sType-unique"
  id="VUID-VkRayTracingPipelineCreateInfoNV-sType-unique"></a>
  VUID-VkRayTracingPipelineCreateInfoNV-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

- <a href="#VUID-VkRayTracingPipelineCreateInfoNV-pStages-parameter"
  id="VUID-VkRayTracingPipelineCreateInfoNV-pStages-parameter"></a>
  VUID-VkRayTracingPipelineCreateInfoNV-pStages-parameter  
  `pStages` **must** be a valid pointer to an array of `stageCount`
  valid
  [VkPipelineShaderStageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageCreateInfo.html)
  structures

- <a href="#VUID-VkRayTracingPipelineCreateInfoNV-pGroups-parameter"
  id="VUID-VkRayTracingPipelineCreateInfoNV-pGroups-parameter"></a>
  VUID-VkRayTracingPipelineCreateInfoNV-pGroups-parameter  
  `pGroups` **must** be a valid pointer to an array of `groupCount`
  valid
  [VkRayTracingShaderGroupCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingShaderGroupCreateInfoNV.html)
  structures

- <a href="#VUID-VkRayTracingPipelineCreateInfoNV-layout-parameter"
  id="VUID-VkRayTracingPipelineCreateInfoNV-layout-parameter"></a>
  VUID-VkRayTracingPipelineCreateInfoNV-layout-parameter  
  `layout` **must** be a valid [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html)
  handle

- <a href="#VUID-VkRayTracingPipelineCreateInfoNV-stageCount-arraylength"
  id="VUID-VkRayTracingPipelineCreateInfoNV-stageCount-arraylength"></a>
  VUID-VkRayTracingPipelineCreateInfoNV-stageCount-arraylength  
  `stageCount` **must** be greater than `0`

- <a href="#VUID-VkRayTracingPipelineCreateInfoNV-groupCount-arraylength"
  id="VUID-VkRayTracingPipelineCreateInfoNV-groupCount-arraylength"></a>
  VUID-VkRayTracingPipelineCreateInfoNV-groupCount-arraylength  
  `groupCount` **must** be greater than `0`

- <a href="#VUID-VkRayTracingPipelineCreateInfoNV-commonparent"
  id="VUID-VkRayTracingPipelineCreateInfoNV-commonparent"></a>
  VUID-VkRayTracingPipelineCreateInfoNV-commonparent  
  Both of `basePipelineHandle`, and `layout` that are valid handles of
  non-ignored parameters **must** have been created, allocated, or
  retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_ray_tracing](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_ray_tracing.html),
[VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html),
[VkPipelineCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlags.html),
[VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html),
[VkPipelineShaderStageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageCreateInfo.html),
[VkRayTracingShaderGroupCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingShaderGroupCreateInfoNV.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCreateRayTracingPipelinesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateRayTracingPipelinesNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkRayTracingPipelineCreateInfoNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
