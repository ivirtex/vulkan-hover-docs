# VkComputePipelineCreateInfo(3) Manual Page

## Name

VkComputePipelineCreateInfo - Structure specifying parameters of a newly created compute pipeline



## [](#_c_specification)C Specification

The `VkComputePipelineCreateInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_0
typedef struct VkComputePipelineCreateInfo {
    VkStructureType                    sType;
    const void*                        pNext;
    VkPipelineCreateFlags              flags;
    VkPipelineShaderStageCreateInfo    stage;
    VkPipelineLayout                   layout;
    VkPipeline                         basePipelineHandle;
    int32_t                            basePipelineIndex;
} VkComputePipelineCreateInfo;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is a bitmask of [VkPipelineCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlagBits.html) specifying how the pipeline will be generated.
- `stage` is a [VkPipelineShaderStageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineShaderStageCreateInfo.html) structure describing the compute shader.
- `layout` is the description of binding locations used by both the pipeline and descriptor sets used with the pipeline. If [VkPhysicalDeviceProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties.html)::`apiVersion` is greater than or equal to Vulkan 1.3 or [VK\_KHR\_maintenance4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance4.html) is enabled `layout` **must** not be accessed outside of the duration of the command this structure is passed to.
- `basePipelineHandle` is a pipeline to derive from.
- `basePipelineIndex` is an index into the `pCreateInfos` parameter to use as a pipeline to derive from.

## [](#_description)Description

The parameters `basePipelineHandle` and `basePipelineIndex` are described in more detail in [Pipeline Derivatives](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-pipeline-derivatives).

If the `pNext` chain includes a [VkPipelineCreateFlags2CreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlags2CreateInfo.html) structure, [VkPipelineCreateFlags2CreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlags2CreateInfo.html)::`flags` from that structure is used instead of `flags` from this structure.

Valid Usage

- [](#VUID-VkComputePipelineCreateInfo-None-09497)VUID-VkComputePipelineCreateInfo-None-09497  
  If the `pNext` chain does not include a [VkPipelineCreateFlags2CreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlags2CreateInfo.html) structure, `flags` **must** be a valid combination of [VkPipelineCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlagBits.html) values
- [](#VUID-VkComputePipelineCreateInfo-flags-07984)VUID-VkComputePipelineCreateInfo-flags-07984  
  If `flags` contains the `VK_PIPELINE_CREATE_DERIVATIVE_BIT` flag, and `basePipelineIndex` is -1, `basePipelineHandle` **must** be a valid compute `VkPipeline` handle
- [](#VUID-VkComputePipelineCreateInfo-flags-07985)VUID-VkComputePipelineCreateInfo-flags-07985  
  If `flags` contains the `VK_PIPELINE_CREATE_DERIVATIVE_BIT` flag, and `basePipelineHandle` is [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `basePipelineIndex` **must** be a valid index into the calling command’s `pCreateInfos` parameter
- [](#VUID-VkComputePipelineCreateInfo-flags-07986)VUID-VkComputePipelineCreateInfo-flags-07986  
  If `flags` contains the `VK_PIPELINE_CREATE_DERIVATIVE_BIT` flag, `basePipelineIndex` **must** be -1 or `basePipelineHandle` **must** be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html)
- [](#VUID-VkComputePipelineCreateInfo-layout-07987)VUID-VkComputePipelineCreateInfo-layout-07987  
  If a push constant block is declared in a shader, a push constant range in `layout` **must** match the shader stage
- [](#VUID-VkComputePipelineCreateInfo-layout-10069)VUID-VkComputePipelineCreateInfo-layout-10069  
  If a push constant block is declared in a shader, the block must be contained inside the push constant range in `layout` that matches the stage
- [](#VUID-VkComputePipelineCreateInfo-layout-07988)VUID-VkComputePipelineCreateInfo-layout-07988  
  If a [resource variable](#interfaces-resources) is declared in a shader, the corresponding descriptor set in `layout` **must** match the shader stage
- [](#VUID-VkComputePipelineCreateInfo-layout-07990)VUID-VkComputePipelineCreateInfo-layout-07990  
  If a [resource variable](#interfaces-resources) is declared in a shader, and the descriptor type is not `VK_DESCRIPTOR_TYPE_MUTABLE_EXT`, the corresponding descriptor set in `layout` **must** match the descriptor type
- [](#VUID-VkComputePipelineCreateInfo-layout-07991)VUID-VkComputePipelineCreateInfo-layout-07991  
  If a [resource variable](#interfaces-resources) is declared in a shader as an array, the corresponding descriptor set in `layout` **must** match the descriptor count
- [](#VUID-VkComputePipelineCreateInfo-None-10391)VUID-VkComputePipelineCreateInfo-None-10391  
  If a [resource variables](#interfaces-resources) is declared in a shader as an array of descriptors, then the descriptor type of that variable **must** not be `VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK`

<!--THE END-->

- [](#VUID-VkComputePipelineCreateInfo-flags-03365)VUID-VkComputePipelineCreateInfo-flags-03365  
  `flags` **must** not include `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_ANY_HIT_SHADERS_BIT_KHR`
- [](#VUID-VkComputePipelineCreateInfo-flags-03366)VUID-VkComputePipelineCreateInfo-flags-03366  
  `flags` **must** not include `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_CLOSEST_HIT_SHADERS_BIT_KHR`
- [](#VUID-VkComputePipelineCreateInfo-flags-03367)VUID-VkComputePipelineCreateInfo-flags-03367  
  `flags` **must** not include `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_MISS_SHADERS_BIT_KHR`
- [](#VUID-VkComputePipelineCreateInfo-flags-03368)VUID-VkComputePipelineCreateInfo-flags-03368  
  `flags` **must** not include `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_INTERSECTION_SHADERS_BIT_KHR`
- [](#VUID-VkComputePipelineCreateInfo-flags-03369)VUID-VkComputePipelineCreateInfo-flags-03369  
  `flags` **must** not include `VK_PIPELINE_CREATE_RAY_TRACING_SKIP_TRIANGLES_BIT_KHR`
- [](#VUID-VkComputePipelineCreateInfo-flags-03370)VUID-VkComputePipelineCreateInfo-flags-03370  
  `flags` **must** not include `VK_PIPELINE_CREATE_RAY_TRACING_SKIP_AABBS_BIT_KHR`
- [](#VUID-VkComputePipelineCreateInfo-flags-03576)VUID-VkComputePipelineCreateInfo-flags-03576  
  `flags` **must** not include `VK_PIPELINE_CREATE_RAY_TRACING_SHADER_GROUP_HANDLE_CAPTURE_REPLAY_BIT_KHR`
- [](#VUID-VkComputePipelineCreateInfo-flags-04945)VUID-VkComputePipelineCreateInfo-flags-04945  
  `flags` **must** not include `VK_PIPELINE_CREATE_RAY_TRACING_ALLOW_MOTION_BIT_NV`
- [](#VUID-VkComputePipelineCreateInfo-flags-09007)VUID-VkComputePipelineCreateInfo-flags-09007  
  If the [`VkPhysicalDeviceDeviceGeneratedCommandsComputeFeaturesNV`::`deviceGeneratedComputePipelines`](#features-deviceGeneratedComputePipelines) feature is not enabled, `flags` **must** not include `VK_PIPELINE_CREATE_INDIRECT_BINDABLE_BIT_NV`
- [](#VUID-VkComputePipelineCreateInfo-flags-09008)VUID-VkComputePipelineCreateInfo-flags-09008  
  If `flags` includes `VK_PIPELINE_CREATE_INDIRECT_BINDABLE_BIT_NV`, then the `pNext` chain **must** include a pointer to a valid instance of [VkComputePipelineIndirectBufferInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComputePipelineIndirectBufferInfoNV.html) specifying the address where the pipeline’s metadata will be saved
- [](#VUID-VkComputePipelineCreateInfo-flags-11007)VUID-VkComputePipelineCreateInfo-flags-11007  
  If `flags` includes `VK_PIPELINE_CREATE_2_INDIRECT_BINDABLE_BIT_EXT`, then the [`VkPhysicalDeviceDeviceGeneratedCommandsFeaturesEXT`::`deviceGeneratedCommands`](#features-deviceGeneratedCommands) feature **must** be enabled
- [](#VUID-VkComputePipelineCreateInfo-pipelineCreationCacheControl-02875)VUID-VkComputePipelineCreateInfo-pipelineCreationCacheControl-02875  
  If the [`pipelineCreationCacheControl`](#features-pipelineCreationCacheControl) feature is not enabled, `flags` **must** not include `VK_PIPELINE_CREATE_FAIL_ON_PIPELINE_COMPILE_REQUIRED_BIT` or `VK_PIPELINE_CREATE_EARLY_RETURN_ON_FAILURE_BIT`
- [](#VUID-VkComputePipelineCreateInfo-stage-00701)VUID-VkComputePipelineCreateInfo-stage-00701  
  The `stage` member of `stage` **must** be `VK_SHADER_STAGE_COMPUTE_BIT`
- [](#VUID-VkComputePipelineCreateInfo-stage-00702)VUID-VkComputePipelineCreateInfo-stage-00702  
  The shader code for the entry point identified by `stage` and the rest of the state identified by this structure **must** adhere to the pipeline linking rules described in the [Shader Interfaces](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces) chapter
- [](#VUID-VkComputePipelineCreateInfo-layout-01687)VUID-VkComputePipelineCreateInfo-layout-01687  
  The number of resources in `layout` accessible to the compute shader stage **must** be less than or equal to `VkPhysicalDeviceLimits`::`maxPerStageResources`
- [](#VUID-VkComputePipelineCreateInfo-shaderEnqueue-09177)VUID-VkComputePipelineCreateInfo-shaderEnqueue-09177  
  If the [`shaderEnqueue`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-shaderEnqueue) feature is not enabled, `flags` **must** not include `VK_PIPELINE_CREATE_LIBRARY_BIT_KHR`
- [](#VUID-VkComputePipelineCreateInfo-flags-09178)VUID-VkComputePipelineCreateInfo-flags-09178  
  If `flags` does not include `VK_PIPELINE_CREATE_LIBRARY_BIT_KHR`, the shader specified by `stage` **must** not declare the `ShaderEnqueueAMDX` capability
- [](#VUID-VkComputePipelineCreateInfo-pipelineStageCreationFeedbackCount-06566)VUID-VkComputePipelineCreateInfo-pipelineStageCreationFeedbackCount-06566  
  If [VkPipelineCreationFeedbackCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreationFeedbackCreateInfo.html)::`pipelineStageCreationFeedbackCount` is not `0`, it **must** be `1`
- [](#VUID-VkComputePipelineCreateInfo-flags-07367)VUID-VkComputePipelineCreateInfo-flags-07367  
  `flags` **must** not include `VK_PIPELINE_CREATE_RAY_TRACING_OPACITY_MICROMAP_BIT_EXT`
- [](#VUID-VkComputePipelineCreateInfo-flags-07996)VUID-VkComputePipelineCreateInfo-flags-07996  
  `flags` **must** not include `VK_PIPELINE_CREATE_RAY_TRACING_DISPLACEMENT_MICROMAP_BIT_NV`

Valid Usage (Implicit)

- [](#VUID-VkComputePipelineCreateInfo-sType-sType)VUID-VkComputePipelineCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_COMPUTE_PIPELINE_CREATE_INFO`
- [](#VUID-VkComputePipelineCreateInfo-pNext-pNext)VUID-VkComputePipelineCreateInfo-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the `pNext` chain **must** be either `NULL` or a pointer to a valid instance of [VkComputePipelineIndirectBufferInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComputePipelineIndirectBufferInfoNV.html), [VkPipelineBinaryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryInfoKHR.html), [VkPipelineCompilerControlCreateInfoAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCompilerControlCreateInfoAMD.html), [VkPipelineCreateFlags2CreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlags2CreateInfo.html), [VkPipelineCreationFeedbackCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreationFeedbackCreateInfo.html), [VkPipelineRobustnessCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRobustnessCreateInfo.html), or [VkSubpassShadingPipelineCreateInfoHUAWEI](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassShadingPipelineCreateInfoHUAWEI.html)
- [](#VUID-VkComputePipelineCreateInfo-sType-unique)VUID-VkComputePipelineCreateInfo-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique
- [](#VUID-VkComputePipelineCreateInfo-stage-parameter)VUID-VkComputePipelineCreateInfo-stage-parameter  
  `stage` **must** be a valid [VkPipelineShaderStageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineShaderStageCreateInfo.html) structure
- [](#VUID-VkComputePipelineCreateInfo-layout-parameter)VUID-VkComputePipelineCreateInfo-layout-parameter  
  `layout` **must** be a valid [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayout.html) handle
- [](#VUID-VkComputePipelineCreateInfo-commonparent)VUID-VkComputePipelineCreateInfo-commonparent  
  Both of `basePipelineHandle`, and `layout` that are valid handles of non-ignored parameters **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html), [VkPipelineCreateFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlags.html), [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayout.html), [VkPipelineShaderStageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineShaderStageCreateInfo.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCreateComputePipelines](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateComputePipelines.html), [vkGetPipelineIndirectMemoryRequirementsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPipelineIndirectMemoryRequirementsNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkComputePipelineCreateInfo).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0