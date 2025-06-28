# VkExecutionGraphPipelineCreateInfoAMDX(3) Manual Page

## Name

VkExecutionGraphPipelineCreateInfoAMDX - Structure specifying parameters of a newly created execution graph pipeline



## [](#_c_specification)C Specification

The `VkExecutionGraphPipelineCreateInfoAMDX` structure is defined as:

```c++
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

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is a bitmask of [VkPipelineCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlagBits.html) specifying how the pipeline will be generated.
- `stageCount` is the number of entries in the `pStages` array.
- `pStages` is a pointer to an array of `stageCount` [VkPipelineShaderStageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineShaderStageCreateInfo.html) structures describing the set of the shader stages to be included in the execution graph pipeline.
- `pLibraryInfo` is a pointer to a [VkPipelineLibraryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLibraryCreateInfoKHR.html) structure defining pipeline libraries to include.
- `layout` is the description of binding locations used by both the pipeline and descriptor sets used with the pipeline. The implementation **must** not access this object outside of the duration of the command this structure is passed to.
- `basePipelineHandle` is a pipeline to derive from
- `basePipelineIndex` is an index into the `pCreateInfos` parameter to use as a pipeline to derive from

## [](#_description)Description

The parameters `basePipelineHandle` and `basePipelineIndex` are described in more detail in [Pipeline Derivatives](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-pipeline-derivatives).

Each shader stage provided when creating an execution graph pipeline (including those in libraries) is associated with a name and an index, determined by the inclusion or omission of a [VkPipelineShaderStageNodeCreateInfoAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineShaderStageNodeCreateInfoAMDX.html) structure in its `pNext` chain. For any graphics pipeline libraries, only the name and index of the vertex or mesh shader stage is linked directly to the graph as a node - other shader stages in the pipeline will be executed after those shader stages as normal. Task shaders cannot be included in a graphics pipeline used for a draw node.

In addition to the shader name and index, an internal "node index" is also generated for each node, which can be queried with [vkGetExecutionGraphPipelineNodeIndexAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetExecutionGraphPipelineNodeIndexAMDX.html), and is used exclusively for initial dispatch of an execution graph.

Valid Usage

- [](#VUID-VkExecutionGraphPipelineCreateInfoAMDX-None-09497)VUID-VkExecutionGraphPipelineCreateInfoAMDX-None-09497  
  If the `pNext` chain does not include a [VkPipelineCreateFlags2CreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlags2CreateInfo.html) structure, `flags` **must** be a valid combination of [VkPipelineCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlagBits.html) values
- [](#VUID-VkExecutionGraphPipelineCreateInfoAMDX-flags-07984)VUID-VkExecutionGraphPipelineCreateInfoAMDX-flags-07984  
  If `flags` contains the `VK_PIPELINE_CREATE_DERIVATIVE_BIT` flag, and `basePipelineIndex` is -1, `basePipelineHandle` **must** be a valid execution graph `VkPipeline` handle
- [](#VUID-VkExecutionGraphPipelineCreateInfoAMDX-flags-07985)VUID-VkExecutionGraphPipelineCreateInfoAMDX-flags-07985  
  If `flags` contains the `VK_PIPELINE_CREATE_DERIVATIVE_BIT` flag, and `basePipelineHandle` is [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `basePipelineIndex` **must** be a valid index into the calling command’s `pCreateInfos` parameter
- [](#VUID-VkExecutionGraphPipelineCreateInfoAMDX-flags-07986)VUID-VkExecutionGraphPipelineCreateInfoAMDX-flags-07986  
  If `flags` contains the `VK_PIPELINE_CREATE_DERIVATIVE_BIT` flag, `basePipelineIndex` **must** be -1 or `basePipelineHandle` **must** be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html)
- [](#VUID-VkExecutionGraphPipelineCreateInfoAMDX-layout-07987)VUID-VkExecutionGraphPipelineCreateInfoAMDX-layout-07987  
  If a push constant block is declared in a shader, a push constant range in `layout` **must** match the shader stage
- [](#VUID-VkExecutionGraphPipelineCreateInfoAMDX-layout-10069)VUID-VkExecutionGraphPipelineCreateInfoAMDX-layout-10069  
  If a push constant block is declared in a shader, the block must be contained inside the push constant range in `layout` that matches the stage
- [](#VUID-VkExecutionGraphPipelineCreateInfoAMDX-layout-07988)VUID-VkExecutionGraphPipelineCreateInfoAMDX-layout-07988  
  If a [resource variable](#interfaces-resources) is declared in a shader, the corresponding descriptor set in `layout` **must** match the shader stage
- [](#VUID-VkExecutionGraphPipelineCreateInfoAMDX-layout-07990)VUID-VkExecutionGraphPipelineCreateInfoAMDX-layout-07990  
  If a [resource variable](#interfaces-resources) is declared in a shader, and the descriptor type is not `VK_DESCRIPTOR_TYPE_MUTABLE_EXT`, the corresponding descriptor set in `layout` **must** match the descriptor type
- [](#VUID-VkExecutionGraphPipelineCreateInfoAMDX-layout-07991)VUID-VkExecutionGraphPipelineCreateInfoAMDX-layout-07991  
  If a [resource variable](#interfaces-resources) is declared in a shader as an array, the corresponding descriptor set in `layout` **must** match the descriptor count
- [](#VUID-VkExecutionGraphPipelineCreateInfoAMDX-None-10391)VUID-VkExecutionGraphPipelineCreateInfoAMDX-None-10391  
  If a [resource variables](#interfaces-resources) is declared in a shader as an array of descriptors, then the descriptor type of that variable **must** not be `VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK`

<!--THE END-->

- [](#VUID-VkExecutionGraphPipelineCreateInfoAMDX-flags-03365)VUID-VkExecutionGraphPipelineCreateInfoAMDX-flags-03365  
  `flags` **must** not include `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_ANY_HIT_SHADERS_BIT_KHR`
- [](#VUID-VkExecutionGraphPipelineCreateInfoAMDX-flags-03366)VUID-VkExecutionGraphPipelineCreateInfoAMDX-flags-03366  
  `flags` **must** not include `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_CLOSEST_HIT_SHADERS_BIT_KHR`
- [](#VUID-VkExecutionGraphPipelineCreateInfoAMDX-flags-03367)VUID-VkExecutionGraphPipelineCreateInfoAMDX-flags-03367  
  `flags` **must** not include `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_MISS_SHADERS_BIT_KHR`
- [](#VUID-VkExecutionGraphPipelineCreateInfoAMDX-flags-03368)VUID-VkExecutionGraphPipelineCreateInfoAMDX-flags-03368  
  `flags` **must** not include `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_INTERSECTION_SHADERS_BIT_KHR`
- [](#VUID-VkExecutionGraphPipelineCreateInfoAMDX-flags-03369)VUID-VkExecutionGraphPipelineCreateInfoAMDX-flags-03369  
  `flags` **must** not include `VK_PIPELINE_CREATE_RAY_TRACING_SKIP_TRIANGLES_BIT_KHR`
- [](#VUID-VkExecutionGraphPipelineCreateInfoAMDX-flags-03370)VUID-VkExecutionGraphPipelineCreateInfoAMDX-flags-03370  
  `flags` **must** not include `VK_PIPELINE_CREATE_RAY_TRACING_SKIP_AABBS_BIT_KHR`
- [](#VUID-VkExecutionGraphPipelineCreateInfoAMDX-flags-03576)VUID-VkExecutionGraphPipelineCreateInfoAMDX-flags-03576  
  `flags` **must** not include `VK_PIPELINE_CREATE_RAY_TRACING_SHADER_GROUP_HANDLE_CAPTURE_REPLAY_BIT_KHR`
- [](#VUID-VkExecutionGraphPipelineCreateInfoAMDX-flags-04945)VUID-VkExecutionGraphPipelineCreateInfoAMDX-flags-04945  
  `flags` **must** not include `VK_PIPELINE_CREATE_RAY_TRACING_ALLOW_MOTION_BIT_NV`
- [](#VUID-VkExecutionGraphPipelineCreateInfoAMDX-flags-09007)VUID-VkExecutionGraphPipelineCreateInfoAMDX-flags-09007  
  If the [`VkPhysicalDeviceDeviceGeneratedCommandsComputeFeaturesNV`::`deviceGeneratedComputePipelines`](#features-deviceGeneratedComputePipelines) feature is not enabled, `flags` **must** not include `VK_PIPELINE_CREATE_INDIRECT_BINDABLE_BIT_NV`
- [](#VUID-VkExecutionGraphPipelineCreateInfoAMDX-flags-09008)VUID-VkExecutionGraphPipelineCreateInfoAMDX-flags-09008  
  If `flags` includes `VK_PIPELINE_CREATE_INDIRECT_BINDABLE_BIT_NV`, then the `pNext` chain **must** include a pointer to a valid instance of [VkComputePipelineIndirectBufferInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComputePipelineIndirectBufferInfoNV.html) specifying the address where the pipeline’s metadata will be saved
- [](#VUID-VkExecutionGraphPipelineCreateInfoAMDX-flags-11007)VUID-VkExecutionGraphPipelineCreateInfoAMDX-flags-11007  
  If `flags` includes `VK_PIPELINE_CREATE_2_INDIRECT_BINDABLE_BIT_EXT`, then the [`VkPhysicalDeviceDeviceGeneratedCommandsFeaturesEXT`::`deviceGeneratedCommands`](#features-deviceGeneratedCommands) feature **must** be enabled
- [](#VUID-VkExecutionGraphPipelineCreateInfoAMDX-pipelineCreationCacheControl-02875)VUID-VkExecutionGraphPipelineCreateInfoAMDX-pipelineCreationCacheControl-02875  
  If the [`pipelineCreationCacheControl`](#features-pipelineCreationCacheControl) feature is not enabled, `flags` **must** not include `VK_PIPELINE_CREATE_FAIL_ON_PIPELINE_COMPILE_REQUIRED_BIT` or `VK_PIPELINE_CREATE_EARLY_RETURN_ON_FAILURE_BIT`
- [](#VUID-VkExecutionGraphPipelineCreateInfoAMDX-stage-09128)VUID-VkExecutionGraphPipelineCreateInfoAMDX-stage-09128  
  The `stage` member of any element of `pStages` **must** be `VK_SHADER_STAGE_COMPUTE_BIT`
- [](#VUID-VkExecutionGraphPipelineCreateInfoAMDX-pStages-09129)VUID-VkExecutionGraphPipelineCreateInfoAMDX-pStages-09129  
  The shader code for the entry point identified by each element of `pStages` and the rest of the state identified by this structure **must** adhere to the pipeline linking rules described in the [Shader Interfaces](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces) chapter
- [](#VUID-VkExecutionGraphPipelineCreateInfoAMDX-layout-09130)VUID-VkExecutionGraphPipelineCreateInfoAMDX-layout-09130  
  `layout` **must** be [consistent](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-pipelinelayout-consistency) with the layout of the shaders specified in `pStages`
- [](#VUID-VkExecutionGraphPipelineCreateInfoAMDX-pLibraryInfo-09131)VUID-VkExecutionGraphPipelineCreateInfoAMDX-pLibraryInfo-09131  
  If `pLibraryInfo` is not `NULL`, each element of its `pLibraries` member **must** have been created with a `layout` that is compatible with the `layout` in this pipeline
- [](#VUID-VkExecutionGraphPipelineCreateInfoAMDX-layout-09132)VUID-VkExecutionGraphPipelineCreateInfoAMDX-layout-09132  
  The number of resources in `layout` accessible to each shader stage that is used by the pipeline **must** be less than or equal to `VkPhysicalDeviceLimits`::`maxPerStageResources`
- [](#VUID-VkExecutionGraphPipelineCreateInfoAMDX-pLibraryInfo-09133)VUID-VkExecutionGraphPipelineCreateInfoAMDX-pLibraryInfo-09133  
  If `pLibraryInfo` is not `NULL`, each element of `pLibraryInfo->pLibraries` **must** be either a compute pipeline, an execution graph pipeline, or a graphics pipeline
- [](#VUID-VkExecutionGraphPipelineCreateInfoAMDX-pLibraryInfo-10181)VUID-VkExecutionGraphPipelineCreateInfoAMDX-pLibraryInfo-10181  
  If `pLibraryInfo` is not `NULL`, each element of `pLibraryInfo->pLibraries` that is a compute pipeline or a graphics pipeline **must** have been created with `VK_PIPELINE_CREATE_2_EXECUTION_GRAPH_BIT_AMDX` set
- [](#VUID-VkExecutionGraphPipelineCreateInfoAMDX-shaderMeshEnqueue-10182)VUID-VkExecutionGraphPipelineCreateInfoAMDX-shaderMeshEnqueue-10182  
  If the [`shaderMeshEnqueue`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-shaderMeshEnqueue) feature is not enabled, and `pLibraryInfo->pLibraries` is not `NULL`, `pLibraryInfo->pLibraries` **must** not contain any graphics pipelines
- [](#VUID-VkExecutionGraphPipelineCreateInfoAMDX-pLibraryInfo-10183)VUID-VkExecutionGraphPipelineCreateInfoAMDX-pLibraryInfo-10183  
  Any element of `pLibraryInfo->pLibraries` identifying a graphics pipeline **must** have been created with [all possible state subsets](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-graphics-subsets-complete)
- [](#VUID-VkExecutionGraphPipelineCreateInfoAMDX-None-09134)VUID-VkExecutionGraphPipelineCreateInfoAMDX-None-09134  
  There **must** be no two nodes in the pipeline that share both the same shader name and index, as specified by [VkPipelineShaderStageNodeCreateInfoAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineShaderStageNodeCreateInfoAMDX.html)
- [](#VUID-VkExecutionGraphPipelineCreateInfoAMDX-None-09135)VUID-VkExecutionGraphPipelineCreateInfoAMDX-None-09135  
  There **must** be no two nodes in the pipeline that share the same shader name and have input payload declarations with different sizes
- [](#VUID-VkExecutionGraphPipelineCreateInfoAMDX-None-09136)VUID-VkExecutionGraphPipelineCreateInfoAMDX-None-09136  
  There **must** be no two nodes in the pipeline that share the same name but have different execution models
- [](#VUID-VkExecutionGraphPipelineCreateInfoAMDX-CoalescedInputCountAMDX-09137)VUID-VkExecutionGraphPipelineCreateInfoAMDX-CoalescedInputCountAMDX-09137  
  There **must** be no two nodes in the pipeline that share the same name where one includes `CoalescedInputCountAMDX` and the other does not
- [](#VUID-VkExecutionGraphPipelineCreateInfoAMDX-StaticNumWorkgroupsAMDX-09138)VUID-VkExecutionGraphPipelineCreateInfoAMDX-StaticNumWorkgroupsAMDX-09138  
  There **must** be no two nodes in the pipeline that share the same name where one includes `StaticNumWorkgroupsAMDX` and the other does not
- [](#VUID-VkExecutionGraphPipelineCreateInfoAMDX-PayloadNodeNameAMDX-09139)VUID-VkExecutionGraphPipelineCreateInfoAMDX-PayloadNodeNameAMDX-09139  
  If an output payload declared in any shader in the pipeline has a `PayloadNodeNameAMDX` decoration with a `Node` `Name` that matches the shader name of any other node in the graph, the size of the output payload **must** match the size of the input payload in the matching node
- [](#VUID-VkExecutionGraphPipelineCreateInfoAMDX-flags-10184)VUID-VkExecutionGraphPipelineCreateInfoAMDX-flags-10184  
  If `flags` does not include `VK_PIPELINE_CREATE_LIBRARY_BIT_KHR`, and an output payload declared in any shader in the pipeline does not have a `PayloadNodeSparseArrayAMDX` decoration, there **must** be a node in the graph corresponding to every index from 0 to its `PayloadNodeArraySizeAMDX` decoration

Valid Usage (Implicit)

- [](#VUID-VkExecutionGraphPipelineCreateInfoAMDX-sType-sType)VUID-VkExecutionGraphPipelineCreateInfoAMDX-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_EXECUTION_GRAPH_PIPELINE_CREATE_INFO_AMDX`
- [](#VUID-VkExecutionGraphPipelineCreateInfoAMDX-pNext-pNext)VUID-VkExecutionGraphPipelineCreateInfoAMDX-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the `pNext` chain **must** be either `NULL` or a pointer to a valid instance of [VkPipelineCompilerControlCreateInfoAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCompilerControlCreateInfoAMD.html) or [VkPipelineCreationFeedbackCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreationFeedbackCreateInfo.html)
- [](#VUID-VkExecutionGraphPipelineCreateInfoAMDX-sType-unique)VUID-VkExecutionGraphPipelineCreateInfoAMDX-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique
- [](#VUID-VkExecutionGraphPipelineCreateInfoAMDX-pStages-parameter)VUID-VkExecutionGraphPipelineCreateInfoAMDX-pStages-parameter  
  If `stageCount` is not `0`, and `pStages` is not `NULL`, `pStages` **must** be a valid pointer to an array of `stageCount` valid [VkPipelineShaderStageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineShaderStageCreateInfo.html) structures
- [](#VUID-VkExecutionGraphPipelineCreateInfoAMDX-pLibraryInfo-parameter)VUID-VkExecutionGraphPipelineCreateInfoAMDX-pLibraryInfo-parameter  
  If `pLibraryInfo` is not `NULL`, `pLibraryInfo` **must** be a valid pointer to a valid [VkPipelineLibraryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLibraryCreateInfoKHR.html) structure
- [](#VUID-VkExecutionGraphPipelineCreateInfoAMDX-layout-parameter)VUID-VkExecutionGraphPipelineCreateInfoAMDX-layout-parameter  
  `layout` **must** be a valid [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayout.html) handle
- [](#VUID-VkExecutionGraphPipelineCreateInfoAMDX-commonparent)VUID-VkExecutionGraphPipelineCreateInfoAMDX-commonparent  
  Both of `basePipelineHandle`, and `layout` that are valid handles of non-ignored parameters **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_see_also)See Also

[VK\_AMDX\_shader\_enqueue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_AMDX_shader_enqueue.html), [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html), [VkPipelineCreateFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlags.html), [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayout.html), [VkPipelineLibraryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLibraryCreateInfoKHR.html), [VkPipelineShaderStageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineShaderStageCreateInfo.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCreateExecutionGraphPipelinesAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateExecutionGraphPipelinesAMDX.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkExecutionGraphPipelineCreateInfoAMDX)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0