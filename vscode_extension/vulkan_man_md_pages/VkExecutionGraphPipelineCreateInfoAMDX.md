# VkExecutionGraphPipelineCreateInfoAMDX(3) Manual Page

## Name

VkExecutionGraphPipelineCreateInfoAMDX - Structure specifying parameters
of a newly created execution graph pipeline



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkExecutionGraphPipelineCreateInfoAMDX` structure is defined as:

``` c
// Provided by VK_AMDX_shader_enqueue
typedef struct VkExecutionGraphPipelineCreateInfoAMDX {
    VkStructureType                           sType;
    const void*                               pNext;
    VkPipelineCreateFlags                     flags;
    uint32_t                                  stageCount;
    const VkPipelineShaderStageCreateInfo*    pStages;
    const VkPipelineLibraryCreateInfoKHR*     pLibraryInfo;
    VkPipelineLayout                          layout;
    VkPipeline                                basePipelineHandle;
    int32_t                                   basePipelineIndex;
} VkExecutionGraphPipelineCreateInfoAMDX;
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

- `pStages` is a pointer to an array of `stageCount`
  [VkPipelineShaderStageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageCreateInfo.html)
  structures describing the set of the shader stages to be included in
  the execution graph pipeline.

- `pLibraryInfo` is a pointer to a
  [VkPipelineLibraryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLibraryCreateInfoKHR.html)
  structure defining pipeline libraries to include.

- `layout` is the description of binding locations used by both the
  pipeline and descriptor sets used with the pipeline.

- `basePipelineHandle` is a pipeline to derive from

- `basePipelineIndex` is an index into the `pCreateInfos` parameter to
  use as a pipeline to derive from

## <a href="#_description" class="anchor"></a>Description

The parameters `basePipelineHandle` and `basePipelineIndex` are
described in more detail in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-pipeline-derivatives"
target="_blank" rel="noopener">Pipeline Derivatives</a>.

Each shader stage provided when creating an execution graph pipeline
(including those in libraries) is associated with a name and an index,
determined by the inclusion or omission of a
[VkPipelineShaderStageNodeCreateInfoAMDX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageNodeCreateInfoAMDX.html)
structure in its `pNext` chain.

In addition to the shader name and index, an internal "node index" is
also generated for each node, which can be queried with
[vkGetExecutionGraphPipelineNodeIndexAMDX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetExecutionGraphPipelineNodeIndexAMDX.html),
and is used exclusively for initial dispatch of an execution graph.

Valid Usage

- <a href="#VUID-VkExecutionGraphPipelineCreateInfoAMDX-None-09497"
  id="VUID-VkExecutionGraphPipelineCreateInfoAMDX-None-09497"></a>
  VUID-VkExecutionGraphPipelineCreateInfoAMDX-None-09497  
  If the `pNext` chain does not include a
  [VkPipelineCreateFlags2CreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlags2CreateInfoKHR.html)
  structure, `flags` must be a valid combination of
  [VkPipelineCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlagBits.html) values

- <a href="#VUID-VkExecutionGraphPipelineCreateInfoAMDX-flags-07984"
  id="VUID-VkExecutionGraphPipelineCreateInfoAMDX-flags-07984"></a>
  VUID-VkExecutionGraphPipelineCreateInfoAMDX-flags-07984  
  If `flags` contains the `VK_PIPELINE_CREATE_DERIVATIVE_BIT` flag, and
  `basePipelineIndex` is -1, `basePipelineHandle` **must** be a valid
  execution graph `VkPipeline` handle

- <a href="#VUID-VkExecutionGraphPipelineCreateInfoAMDX-flags-07985"
  id="VUID-VkExecutionGraphPipelineCreateInfoAMDX-flags-07985"></a>
  VUID-VkExecutionGraphPipelineCreateInfoAMDX-flags-07985  
  If `flags` contains the `VK_PIPELINE_CREATE_DERIVATIVE_BIT` flag, and
  `basePipelineHandle` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `basePipelineIndex` **must** be a valid index into the calling
  command’s `pCreateInfos` parameter

- <a href="#VUID-VkExecutionGraphPipelineCreateInfoAMDX-flags-07986"
  id="VUID-VkExecutionGraphPipelineCreateInfoAMDX-flags-07986"></a>
  VUID-VkExecutionGraphPipelineCreateInfoAMDX-flags-07986  
  If `flags` contains the `VK_PIPELINE_CREATE_DERIVATIVE_BIT` flag,
  `basePipelineIndex` **must** be -1 or `basePipelineHandle` **must** be
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

- <a href="#VUID-VkExecutionGraphPipelineCreateInfoAMDX-layout-07987"
  id="VUID-VkExecutionGraphPipelineCreateInfoAMDX-layout-07987"></a>
  VUID-VkExecutionGraphPipelineCreateInfoAMDX-layout-07987  
  If a push constant block is declared in a shader, a push constant
  range in `layout` **must** match both the shader stage and range

- <a href="#VUID-VkExecutionGraphPipelineCreateInfoAMDX-layout-07988"
  id="VUID-VkExecutionGraphPipelineCreateInfoAMDX-layout-07988"></a>
  VUID-VkExecutionGraphPipelineCreateInfoAMDX-layout-07988  
  If a [resource variables](#interfaces-resources) is declared in a
  shader, a descriptor slot in `layout` **must** match the shader stage

- <a href="#VUID-VkExecutionGraphPipelineCreateInfoAMDX-layout-07990"
  id="VUID-VkExecutionGraphPipelineCreateInfoAMDX-layout-07990"></a>
  VUID-VkExecutionGraphPipelineCreateInfoAMDX-layout-07990  
  If a [resource variables](#interfaces-resources) is declared in a
  shader, and the descriptor type is not
  `VK_DESCRIPTOR_TYPE_MUTABLE_EXT`, a descriptor slot in `layout`
  **must** match the descriptor type

- <a href="#VUID-VkExecutionGraphPipelineCreateInfoAMDX-layout-07991"
  id="VUID-VkExecutionGraphPipelineCreateInfoAMDX-layout-07991"></a>
  VUID-VkExecutionGraphPipelineCreateInfoAMDX-layout-07991  
  If a [resource variables](#interfaces-resources) is declared in a
  shader as an array, a descriptor slot in `layout` **must** match the
  descriptor count

<!-- -->

- <a href="#VUID-VkExecutionGraphPipelineCreateInfoAMDX-flags-03365"
  id="VUID-VkExecutionGraphPipelineCreateInfoAMDX-flags-03365"></a>
  VUID-VkExecutionGraphPipelineCreateInfoAMDX-flags-03365  
  `flags` **must** not include
  `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_ANY_HIT_SHADERS_BIT_KHR`

- <a href="#VUID-VkExecutionGraphPipelineCreateInfoAMDX-flags-03366"
  id="VUID-VkExecutionGraphPipelineCreateInfoAMDX-flags-03366"></a>
  VUID-VkExecutionGraphPipelineCreateInfoAMDX-flags-03366  
  `flags` **must** not include
  `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_CLOSEST_HIT_SHADERS_BIT_KHR`

- <a href="#VUID-VkExecutionGraphPipelineCreateInfoAMDX-flags-03367"
  id="VUID-VkExecutionGraphPipelineCreateInfoAMDX-flags-03367"></a>
  VUID-VkExecutionGraphPipelineCreateInfoAMDX-flags-03367  
  `flags` **must** not include
  `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_MISS_SHADERS_BIT_KHR`

- <a href="#VUID-VkExecutionGraphPipelineCreateInfoAMDX-flags-03368"
  id="VUID-VkExecutionGraphPipelineCreateInfoAMDX-flags-03368"></a>
  VUID-VkExecutionGraphPipelineCreateInfoAMDX-flags-03368  
  `flags` **must** not include
  `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_INTERSECTION_SHADERS_BIT_KHR`

- <a href="#VUID-VkExecutionGraphPipelineCreateInfoAMDX-flags-03369"
  id="VUID-VkExecutionGraphPipelineCreateInfoAMDX-flags-03369"></a>
  VUID-VkExecutionGraphPipelineCreateInfoAMDX-flags-03369  
  `flags` **must** not include
  `VK_PIPELINE_CREATE_RAY_TRACING_SKIP_TRIANGLES_BIT_KHR`

- <a href="#VUID-VkExecutionGraphPipelineCreateInfoAMDX-flags-03370"
  id="VUID-VkExecutionGraphPipelineCreateInfoAMDX-flags-03370"></a>
  VUID-VkExecutionGraphPipelineCreateInfoAMDX-flags-03370  
  `flags` **must** not include
  `VK_PIPELINE_CREATE_RAY_TRACING_SKIP_AABBS_BIT_KHR`

- <a href="#VUID-VkExecutionGraphPipelineCreateInfoAMDX-flags-03576"
  id="VUID-VkExecutionGraphPipelineCreateInfoAMDX-flags-03576"></a>
  VUID-VkExecutionGraphPipelineCreateInfoAMDX-flags-03576  
  `flags` **must** not include
  `VK_PIPELINE_CREATE_RAY_TRACING_SHADER_GROUP_HANDLE_CAPTURE_REPLAY_BIT_KHR`

- <a href="#VUID-VkExecutionGraphPipelineCreateInfoAMDX-flags-04945"
  id="VUID-VkExecutionGraphPipelineCreateInfoAMDX-flags-04945"></a>
  VUID-VkExecutionGraphPipelineCreateInfoAMDX-flags-04945  
  `flags` **must** not include
  `VK_PIPELINE_CREATE_RAY_TRACING_ALLOW_MOTION_BIT_NV`

- <a href="#VUID-VkExecutionGraphPipelineCreateInfoAMDX-flags-09007"
  id="VUID-VkExecutionGraphPipelineCreateInfoAMDX-flags-09007"></a>
  VUID-VkExecutionGraphPipelineCreateInfoAMDX-flags-09007  
  If the
  [`VkPhysicalDeviceDeviceGeneratedCommandsComputeFeaturesNV`::`deviceGeneratedComputePipelines`](#features-deviceGeneratedComputePipelines)
  is not enabled, `flags` **must** not include
  `VK_PIPELINE_CREATE_INDIRECT_BINDABLE_BIT_NV`

- <a href="#VUID-VkExecutionGraphPipelineCreateInfoAMDX-flags-09008"
  id="VUID-VkExecutionGraphPipelineCreateInfoAMDX-flags-09008"></a>
  VUID-VkExecutionGraphPipelineCreateInfoAMDX-flags-09008  
  If `flags` includes `VK_PIPELINE_CREATE_INDIRECT_BINDABLE_BIT_NV`,
  then the `pNext` chain **must** include a pointer to a valid instance
  of
  [VkComputePipelineIndirectBufferInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkComputePipelineIndirectBufferInfoNV.html)
  specifying the address where the pipeline’s metadata will be saved

- <a
  href="#VUID-VkExecutionGraphPipelineCreateInfoAMDX-pipelineCreationCacheControl-02875"
  id="VUID-VkExecutionGraphPipelineCreateInfoAMDX-pipelineCreationCacheControl-02875"></a>
  VUID-VkExecutionGraphPipelineCreateInfoAMDX-pipelineCreationCacheControl-02875  
  If the
  [`pipelineCreationCacheControl`](#features-pipelineCreationCacheControl)
  feature is not enabled, `flags` **must** not include
  `VK_PIPELINE_CREATE_FAIL_ON_PIPELINE_COMPILE_REQUIRED_BIT` or
  `VK_PIPELINE_CREATE_EARLY_RETURN_ON_FAILURE_BIT`

- <a href="#VUID-VkExecutionGraphPipelineCreateInfoAMDX-stage-09128"
  id="VUID-VkExecutionGraphPipelineCreateInfoAMDX-stage-09128"></a>
  VUID-VkExecutionGraphPipelineCreateInfoAMDX-stage-09128  
  The `stage` member of any element of `pStages` **must** be
  `VK_SHADER_STAGE_COMPUTE_BIT`

- <a href="#VUID-VkExecutionGraphPipelineCreateInfoAMDX-pStages-09129"
  id="VUID-VkExecutionGraphPipelineCreateInfoAMDX-pStages-09129"></a>
  VUID-VkExecutionGraphPipelineCreateInfoAMDX-pStages-09129  
  The shader code for the entry point identified by each element of
  `pStages` and the rest of the state identified by this structure
  **must** adhere to the pipeline linking rules described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces"
  target="_blank" rel="noopener">Shader Interfaces</a> chapter

- <a href="#VUID-VkExecutionGraphPipelineCreateInfoAMDX-layout-09130"
  id="VUID-VkExecutionGraphPipelineCreateInfoAMDX-layout-09130"></a>
  VUID-VkExecutionGraphPipelineCreateInfoAMDX-layout-09130  
  `layout` **must** be <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-pipelinelayout-consistency"
  target="_blank" rel="noopener">consistent</a> with the layout of the
  shaders specified in `pStages`

- <a
  href="#VUID-VkExecutionGraphPipelineCreateInfoAMDX-pLibraryInfo-09131"
  id="VUID-VkExecutionGraphPipelineCreateInfoAMDX-pLibraryInfo-09131"></a>
  VUID-VkExecutionGraphPipelineCreateInfoAMDX-pLibraryInfo-09131  
  If `pLibraryInfo` is not `NULL`, each element of its `pLibraries`
  member **must** have been created with a `layout` that is compatible
  with the `layout` in this pipeline

- <a href="#VUID-VkExecutionGraphPipelineCreateInfoAMDX-layout-09132"
  id="VUID-VkExecutionGraphPipelineCreateInfoAMDX-layout-09132"></a>
  VUID-VkExecutionGraphPipelineCreateInfoAMDX-layout-09132  
  The number of resources in `layout` accessible to each shader stage
  that is used by the pipeline **must** be less than or equal to
  `VkPhysicalDeviceLimits`::`maxPerStageResources`

- <a
  href="#VUID-VkExecutionGraphPipelineCreateInfoAMDX-pLibraryInfo-09133"
  id="VUID-VkExecutionGraphPipelineCreateInfoAMDX-pLibraryInfo-09133"></a>
  VUID-VkExecutionGraphPipelineCreateInfoAMDX-pLibraryInfo-09133  
  If `pLibraryInfo` is not `NULL`, each element of
  `pLibraryInfo->libraries` **must** be either a compute pipeline or an
  execution graph pipeline

- <a href="#VUID-VkExecutionGraphPipelineCreateInfoAMDX-None-09134"
  id="VUID-VkExecutionGraphPipelineCreateInfoAMDX-None-09134"></a>
  VUID-VkExecutionGraphPipelineCreateInfoAMDX-None-09134  
  There **must** be no two nodes in the pipeline that share both the
  same shader name and index, as specified by
  [VkPipelineShaderStageNodeCreateInfoAMDX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageNodeCreateInfoAMDX.html)

- <a href="#VUID-VkExecutionGraphPipelineCreateInfoAMDX-None-09135"
  id="VUID-VkExecutionGraphPipelineCreateInfoAMDX-None-09135"></a>
  VUID-VkExecutionGraphPipelineCreateInfoAMDX-None-09135  
  There **must** be no two nodes in the pipeline that share the same
  shader name and have input payload declarations with different sizes

- <a href="#VUID-VkExecutionGraphPipelineCreateInfoAMDX-None-09136"
  id="VUID-VkExecutionGraphPipelineCreateInfoAMDX-None-09136"></a>
  VUID-VkExecutionGraphPipelineCreateInfoAMDX-None-09136  
  There **must** be no two nodes in the pipeline that share the same
  name but have different execution models

- <a
  href="#VUID-VkExecutionGraphPipelineCreateInfoAMDX-CoalescedInputCountAMDX-09137"
  id="VUID-VkExecutionGraphPipelineCreateInfoAMDX-CoalescedInputCountAMDX-09137"></a>
  VUID-VkExecutionGraphPipelineCreateInfoAMDX-CoalescedInputCountAMDX-09137  
  There **must** be no two nodes in the pipeline that share the same
  name where one includes `CoalescedInputCountAMDX` and the other does
  not

- <a
  href="#VUID-VkExecutionGraphPipelineCreateInfoAMDX-StaticNumWorkgroupsAMDX-09138"
  id="VUID-VkExecutionGraphPipelineCreateInfoAMDX-StaticNumWorkgroupsAMDX-09138"></a>
  VUID-VkExecutionGraphPipelineCreateInfoAMDX-StaticNumWorkgroupsAMDX-09138  
  There **must** be no two nodes in the pipeline that share the same
  name where one includes `StaticNumWorkgroupsAMDX` and the other does
  not

- <a
  href="#VUID-VkExecutionGraphPipelineCreateInfoAMDX-PayloadNodeNameAMDX-09139"
  id="VUID-VkExecutionGraphPipelineCreateInfoAMDX-PayloadNodeNameAMDX-09139"></a>
  VUID-VkExecutionGraphPipelineCreateInfoAMDX-PayloadNodeNameAMDX-09139  
  If an output payload declared in any shader in the pipeline has a
  `PayloadNodeNameAMDX` decoration with a `Node` `Name` that matches the
  shader name of any other node in the graph, the size of the output
  payload **must** match the size of the input payload in the matching
  node

Valid Usage (Implicit)

- <a href="#VUID-VkExecutionGraphPipelineCreateInfoAMDX-sType-sType"
  id="VUID-VkExecutionGraphPipelineCreateInfoAMDX-sType-sType"></a>
  VUID-VkExecutionGraphPipelineCreateInfoAMDX-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_EXECUTION_GRAPH_PIPELINE_CREATE_INFO_AMDX`

- <a href="#VUID-VkExecutionGraphPipelineCreateInfoAMDX-pNext-pNext"
  id="VUID-VkExecutionGraphPipelineCreateInfoAMDX-pNext-pNext"></a>
  VUID-VkExecutionGraphPipelineCreateInfoAMDX-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the
  `pNext` chain **must** be either `NULL` or a pointer to a valid
  instance of
  [VkPipelineCompilerControlCreateInfoAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCompilerControlCreateInfoAMD.html)
  or
  [VkPipelineCreationFeedbackCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreationFeedbackCreateInfo.html)

- <a href="#VUID-VkExecutionGraphPipelineCreateInfoAMDX-sType-unique"
  id="VUID-VkExecutionGraphPipelineCreateInfoAMDX-sType-unique"></a>
  VUID-VkExecutionGraphPipelineCreateInfoAMDX-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

- <a href="#VUID-VkExecutionGraphPipelineCreateInfoAMDX-pStages-parameter"
  id="VUID-VkExecutionGraphPipelineCreateInfoAMDX-pStages-parameter"></a>
  VUID-VkExecutionGraphPipelineCreateInfoAMDX-pStages-parameter  
  If `stageCount` is not `0`, and `pStages` is not `NULL`, `pStages`
  **must** be a valid pointer to an array of `stageCount` valid
  [VkPipelineShaderStageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageCreateInfo.html)
  structures

- <a
  href="#VUID-VkExecutionGraphPipelineCreateInfoAMDX-pLibraryInfo-parameter"
  id="VUID-VkExecutionGraphPipelineCreateInfoAMDX-pLibraryInfo-parameter"></a>
  VUID-VkExecutionGraphPipelineCreateInfoAMDX-pLibraryInfo-parameter  
  If `pLibraryInfo` is not `NULL`, `pLibraryInfo` **must** be a valid
  pointer to a valid
  [VkPipelineLibraryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLibraryCreateInfoKHR.html)
  structure

- <a href="#VUID-VkExecutionGraphPipelineCreateInfoAMDX-layout-parameter"
  id="VUID-VkExecutionGraphPipelineCreateInfoAMDX-layout-parameter"></a>
  VUID-VkExecutionGraphPipelineCreateInfoAMDX-layout-parameter  
  `layout` **must** be a valid [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html)
  handle

- <a href="#VUID-VkExecutionGraphPipelineCreateInfoAMDX-commonparent"
  id="VUID-VkExecutionGraphPipelineCreateInfoAMDX-commonparent"></a>
  VUID-VkExecutionGraphPipelineCreateInfoAMDX-commonparent  
  Both of `basePipelineHandle`, and `layout` that are valid handles of
  non-ignored parameters **must** have been created, allocated, or
  retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_see_also" class="anchor"></a>See Also

[VK_AMDX_shader_enqueue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_AMDX_shader_enqueue.html),
[VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html),
[VkPipelineCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlags.html),
[VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html),
[VkPipelineLibraryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLibraryCreateInfoKHR.html),
[VkPipelineShaderStageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageCreateInfo.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCreateExecutionGraphPipelinesAMDX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateExecutionGraphPipelinesAMDX.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkExecutionGraphPipelineCreateInfoAMDX"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
