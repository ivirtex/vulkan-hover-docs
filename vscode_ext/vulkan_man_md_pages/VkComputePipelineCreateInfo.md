# VkComputePipelineCreateInfo(3) Manual Page

## Name

VkComputePipelineCreateInfo - Structure specifying parameters of a newly
created compute pipeline



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkComputePipelineCreateInfo` structure is defined as:

``` c
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

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is a bitmask of
  [VkPipelineCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlagBits.html) specifying
  how the pipeline will be generated.

- `stage` is a
  [VkPipelineShaderStageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageCreateInfo.html)
  structure describing the compute shader.

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

- <a href="#VUID-VkComputePipelineCreateInfo-None-09497"
  id="VUID-VkComputePipelineCreateInfo-None-09497"></a>
  VUID-VkComputePipelineCreateInfo-None-09497  
  If the `pNext` chain does not include a
  [VkPipelineCreateFlags2CreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlags2CreateInfoKHR.html)
  structure, `flags` must be a valid combination of
  [VkPipelineCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlagBits.html) values

- <a href="#VUID-VkComputePipelineCreateInfo-flags-07984"
  id="VUID-VkComputePipelineCreateInfo-flags-07984"></a>
  VUID-VkComputePipelineCreateInfo-flags-07984  
  If `flags` contains the `VK_PIPELINE_CREATE_DERIVATIVE_BIT` flag, and
  `basePipelineIndex` is -1, `basePipelineHandle` **must** be a valid
  compute `VkPipeline` handle

- <a href="#VUID-VkComputePipelineCreateInfo-flags-07985"
  id="VUID-VkComputePipelineCreateInfo-flags-07985"></a>
  VUID-VkComputePipelineCreateInfo-flags-07985  
  If `flags` contains the `VK_PIPELINE_CREATE_DERIVATIVE_BIT` flag, and
  `basePipelineHandle` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `basePipelineIndex` **must** be a valid index into the calling
  command’s `pCreateInfos` parameter

- <a href="#VUID-VkComputePipelineCreateInfo-flags-07986"
  id="VUID-VkComputePipelineCreateInfo-flags-07986"></a>
  VUID-VkComputePipelineCreateInfo-flags-07986  
  If `flags` contains the `VK_PIPELINE_CREATE_DERIVATIVE_BIT` flag,
  `basePipelineIndex` **must** be -1 or `basePipelineHandle` **must** be
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

- <a href="#VUID-VkComputePipelineCreateInfo-layout-07987"
  id="VUID-VkComputePipelineCreateInfo-layout-07987"></a>
  VUID-VkComputePipelineCreateInfo-layout-07987  
  If a push constant block is declared in a shader, a push constant
  range in `layout` **must** match both the shader stage and range

- <a href="#VUID-VkComputePipelineCreateInfo-layout-07988"
  id="VUID-VkComputePipelineCreateInfo-layout-07988"></a>
  VUID-VkComputePipelineCreateInfo-layout-07988  
  If a [resource variables](#interfaces-resources) is declared in a
  shader, a descriptor slot in `layout` **must** match the shader stage

- <a href="#VUID-VkComputePipelineCreateInfo-layout-07990"
  id="VUID-VkComputePipelineCreateInfo-layout-07990"></a>
  VUID-VkComputePipelineCreateInfo-layout-07990  
  If a [resource variables](#interfaces-resources) is declared in a
  shader, and the descriptor type is not
  `VK_DESCRIPTOR_TYPE_MUTABLE_EXT`, a descriptor slot in `layout`
  **must** match the descriptor type

- <a href="#VUID-VkComputePipelineCreateInfo-layout-07991"
  id="VUID-VkComputePipelineCreateInfo-layout-07991"></a>
  VUID-VkComputePipelineCreateInfo-layout-07991  
  If a [resource variables](#interfaces-resources) is declared in a
  shader as an array, a descriptor slot in `layout` **must** match the
  descriptor count

<!-- -->

- <a href="#VUID-VkComputePipelineCreateInfo-flags-03365"
  id="VUID-VkComputePipelineCreateInfo-flags-03365"></a>
  VUID-VkComputePipelineCreateInfo-flags-03365  
  `flags` **must** not include
  `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_ANY_HIT_SHADERS_BIT_KHR`

- <a href="#VUID-VkComputePipelineCreateInfo-flags-03366"
  id="VUID-VkComputePipelineCreateInfo-flags-03366"></a>
  VUID-VkComputePipelineCreateInfo-flags-03366  
  `flags` **must** not include
  `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_CLOSEST_HIT_SHADERS_BIT_KHR`

- <a href="#VUID-VkComputePipelineCreateInfo-flags-03367"
  id="VUID-VkComputePipelineCreateInfo-flags-03367"></a>
  VUID-VkComputePipelineCreateInfo-flags-03367  
  `flags` **must** not include
  `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_MISS_SHADERS_BIT_KHR`

- <a href="#VUID-VkComputePipelineCreateInfo-flags-03368"
  id="VUID-VkComputePipelineCreateInfo-flags-03368"></a>
  VUID-VkComputePipelineCreateInfo-flags-03368  
  `flags` **must** not include
  `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_INTERSECTION_SHADERS_BIT_KHR`

- <a href="#VUID-VkComputePipelineCreateInfo-flags-03369"
  id="VUID-VkComputePipelineCreateInfo-flags-03369"></a>
  VUID-VkComputePipelineCreateInfo-flags-03369  
  `flags` **must** not include
  `VK_PIPELINE_CREATE_RAY_TRACING_SKIP_TRIANGLES_BIT_KHR`

- <a href="#VUID-VkComputePipelineCreateInfo-flags-03370"
  id="VUID-VkComputePipelineCreateInfo-flags-03370"></a>
  VUID-VkComputePipelineCreateInfo-flags-03370  
  `flags` **must** not include
  `VK_PIPELINE_CREATE_RAY_TRACING_SKIP_AABBS_BIT_KHR`

- <a href="#VUID-VkComputePipelineCreateInfo-flags-03576"
  id="VUID-VkComputePipelineCreateInfo-flags-03576"></a>
  VUID-VkComputePipelineCreateInfo-flags-03576  
  `flags` **must** not include
  `VK_PIPELINE_CREATE_RAY_TRACING_SHADER_GROUP_HANDLE_CAPTURE_REPLAY_BIT_KHR`

- <a href="#VUID-VkComputePipelineCreateInfo-flags-04945"
  id="VUID-VkComputePipelineCreateInfo-flags-04945"></a>
  VUID-VkComputePipelineCreateInfo-flags-04945  
  `flags` **must** not include
  `VK_PIPELINE_CREATE_RAY_TRACING_ALLOW_MOTION_BIT_NV`

- <a href="#VUID-VkComputePipelineCreateInfo-flags-09007"
  id="VUID-VkComputePipelineCreateInfo-flags-09007"></a>
  VUID-VkComputePipelineCreateInfo-flags-09007  
  If the
  [`VkPhysicalDeviceDeviceGeneratedCommandsComputeFeaturesNV`::`deviceGeneratedComputePipelines`](#features-deviceGeneratedComputePipelines)
  is not enabled, `flags` **must** not include
  `VK_PIPELINE_CREATE_INDIRECT_BINDABLE_BIT_NV`

- <a href="#VUID-VkComputePipelineCreateInfo-flags-09008"
  id="VUID-VkComputePipelineCreateInfo-flags-09008"></a>
  VUID-VkComputePipelineCreateInfo-flags-09008  
  If `flags` includes `VK_PIPELINE_CREATE_INDIRECT_BINDABLE_BIT_NV`,
  then the `pNext` chain **must** include a pointer to a valid instance
  of
  [VkComputePipelineIndirectBufferInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkComputePipelineIndirectBufferInfoNV.html)
  specifying the address where the pipeline’s metadata will be saved

- <a
  href="#VUID-VkComputePipelineCreateInfo-pipelineCreationCacheControl-02875"
  id="VUID-VkComputePipelineCreateInfo-pipelineCreationCacheControl-02875"></a>
  VUID-VkComputePipelineCreateInfo-pipelineCreationCacheControl-02875  
  If the
  [`pipelineCreationCacheControl`](#features-pipelineCreationCacheControl)
  feature is not enabled, `flags` **must** not include
  `VK_PIPELINE_CREATE_FAIL_ON_PIPELINE_COMPILE_REQUIRED_BIT` or
  `VK_PIPELINE_CREATE_EARLY_RETURN_ON_FAILURE_BIT`

- <a href="#VUID-VkComputePipelineCreateInfo-stage-00701"
  id="VUID-VkComputePipelineCreateInfo-stage-00701"></a>
  VUID-VkComputePipelineCreateInfo-stage-00701  
  The `stage` member of `stage` **must** be
  `VK_SHADER_STAGE_COMPUTE_BIT`

- <a href="#VUID-VkComputePipelineCreateInfo-stage-00702"
  id="VUID-VkComputePipelineCreateInfo-stage-00702"></a>
  VUID-VkComputePipelineCreateInfo-stage-00702  
  The shader code for the entry point identified by `stage` and the rest
  of the state identified by this structure **must** adhere to the
  pipeline linking rules described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces"
  target="_blank" rel="noopener">Shader Interfaces</a> chapter

- <a href="#VUID-VkComputePipelineCreateInfo-layout-01687"
  id="VUID-VkComputePipelineCreateInfo-layout-01687"></a>
  VUID-VkComputePipelineCreateInfo-layout-01687  
  The number of resources in `layout` accessible to the compute shader
  stage **must** be less than or equal to
  `VkPhysicalDeviceLimits`::`maxPerStageResources`

- <a href="#VUID-VkComputePipelineCreateInfo-shaderEnqueue-09177"
  id="VUID-VkComputePipelineCreateInfo-shaderEnqueue-09177"></a>
  VUID-VkComputePipelineCreateInfo-shaderEnqueue-09177  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderEnqueue"
  target="_blank" rel="noopener"><code>shaderEnqueue</code></a> is not
  enabled, `flags` **must** not include
  `VK_PIPELINE_CREATE_LIBRARY_BIT_KHR`

- <a href="#VUID-VkComputePipelineCreateInfo-flags-09178"
  id="VUID-VkComputePipelineCreateInfo-flags-09178"></a>
  VUID-VkComputePipelineCreateInfo-flags-09178  
  If `flags` does not include `VK_PIPELINE_CREATE_LIBRARY_BIT_KHR`, the
  shader specified by `stage` **must** not declare the
  `ShaderEnqueueAMDX` capability

- <a
  href="#VUID-VkComputePipelineCreateInfo-pipelineStageCreationFeedbackCount-06566"
  id="VUID-VkComputePipelineCreateInfo-pipelineStageCreationFeedbackCount-06566"></a>
  VUID-VkComputePipelineCreateInfo-pipelineStageCreationFeedbackCount-06566  
  If
  [VkPipelineCreationFeedbackCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreationFeedbackCreateInfo.html)::`pipelineStageCreationFeedbackCount`
  is not `0`, it **must** be `1`

- <a href="#VUID-VkComputePipelineCreateInfo-flags-07367"
  id="VUID-VkComputePipelineCreateInfo-flags-07367"></a>
  VUID-VkComputePipelineCreateInfo-flags-07367  
  `flags` **must** not include
  `VK_PIPELINE_CREATE_RAY_TRACING_OPACITY_MICROMAP_BIT_EXT`

- <a href="#VUID-VkComputePipelineCreateInfo-flags-07996"
  id="VUID-VkComputePipelineCreateInfo-flags-07996"></a>
  VUID-VkComputePipelineCreateInfo-flags-07996  
  `flags` **must** not include
  `VK_PIPELINE_CREATE_RAY_TRACING_DISPLACEMENT_MICROMAP_BIT_NV`

Valid Usage (Implicit)

- <a href="#VUID-VkComputePipelineCreateInfo-sType-sType"
  id="VUID-VkComputePipelineCreateInfo-sType-sType"></a>
  VUID-VkComputePipelineCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_COMPUTE_PIPELINE_CREATE_INFO`

- <a href="#VUID-VkComputePipelineCreateInfo-pNext-pNext"
  id="VUID-VkComputePipelineCreateInfo-pNext-pNext"></a>
  VUID-VkComputePipelineCreateInfo-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the
  `pNext` chain **must** be either `NULL` or a pointer to a valid
  instance of
  [VkComputePipelineIndirectBufferInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkComputePipelineIndirectBufferInfoNV.html),
  [VkPipelineCompilerControlCreateInfoAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCompilerControlCreateInfoAMD.html),
  [VkPipelineCreateFlags2CreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlags2CreateInfoKHR.html),
  [VkPipelineCreationFeedbackCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreationFeedbackCreateInfo.html),
  [VkPipelineRobustnessCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRobustnessCreateInfoEXT.html),
  or
  [VkSubpassShadingPipelineCreateInfoHUAWEI](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassShadingPipelineCreateInfoHUAWEI.html)

- <a href="#VUID-VkComputePipelineCreateInfo-sType-unique"
  id="VUID-VkComputePipelineCreateInfo-sType-unique"></a>
  VUID-VkComputePipelineCreateInfo-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

- <a href="#VUID-VkComputePipelineCreateInfo-stage-parameter"
  id="VUID-VkComputePipelineCreateInfo-stage-parameter"></a>
  VUID-VkComputePipelineCreateInfo-stage-parameter  
  `stage` **must** be a valid
  [VkPipelineShaderStageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageCreateInfo.html)
  structure

- <a href="#VUID-VkComputePipelineCreateInfo-layout-parameter"
  id="VUID-VkComputePipelineCreateInfo-layout-parameter"></a>
  VUID-VkComputePipelineCreateInfo-layout-parameter  
  `layout` **must** be a valid [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html)
  handle

- <a href="#VUID-VkComputePipelineCreateInfo-commonparent"
  id="VUID-VkComputePipelineCreateInfo-commonparent"></a>
  VUID-VkComputePipelineCreateInfo-commonparent  
  Both of `basePipelineHandle`, and `layout` that are valid handles of
  non-ignored parameters **must** have been created, allocated, or
  retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html), [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html),
[VkPipelineCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlags.html),
[VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html),
[VkPipelineShaderStageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageCreateInfo.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCreateComputePipelines](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateComputePipelines.html),
[vkGetPipelineIndirectMemoryRequirementsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPipelineIndirectMemoryRequirementsNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkComputePipelineCreateInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
