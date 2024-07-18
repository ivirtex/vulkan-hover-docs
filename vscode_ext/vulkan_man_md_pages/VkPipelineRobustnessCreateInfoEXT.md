# VkPipelineRobustnessCreateInfoEXT(3) Manual Page

## Name

VkPipelineRobustnessCreateInfoEXT - Structure controlling the robustness
of a newly created pipeline shader stage



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPipelineRobustnessCreateInfoEXT` structure is defined as:

``` c
// Provided by VK_EXT_pipeline_robustness
typedef struct VkPipelineRobustnessCreateInfoEXT {
    VkStructureType                          sType;
    const void*                              pNext;
    VkPipelineRobustnessBufferBehaviorEXT    storageBuffers;
    VkPipelineRobustnessBufferBehaviorEXT    uniformBuffers;
    VkPipelineRobustnessBufferBehaviorEXT    vertexInputs;
    VkPipelineRobustnessImageBehaviorEXT     images;
} VkPipelineRobustnessCreateInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `storageBuffers` sets the behavior of out of bounds accesses made to
  resources bound as:

  - `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER`

  - `VK_DESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFER`

  - `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC`

- `uniformBuffers` describes the behavior of out of bounds accesses made
  to resources bound as:

  - `VK_DESCRIPTOR_TYPE_UNIFORM_TEXEL_BUFFER`

  - `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER`

  - `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC`

  - `VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK`

- `vertexInputs` describes the behavior of out of bounds accesses made
  to vertex input attributes

- `images` describes the behavior of out of bounds accesses made to
  resources bound as:

  - `VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE`

  - `VK_DESCRIPTOR_TYPE_STORAGE_IMAGE`

## <a href="#_description" class="anchor"></a>Description

Resources bound as `VK_DESCRIPTOR_TYPE_MUTABLE_EXT` will have the
robustness behavior that covers its active descriptor type.

The scope of the effect of `VkPipelineRobustnessCreateInfoEXT` depends
on which structure’s `pNext` chain it is included in.

- `VkGraphicsPipelineCreateInfo`, `VkRayTracingPipelineCreateInfoKHR`,
  `VkComputePipelineCreateInfo`:  
  The robustness behavior described by
  `VkPipelineRobustnessCreateInfoEXT` applies to all accesses through
  this pipeline

- `VkPipelineShaderStageCreateInfo`:  
  The robustness behavior described by
  `VkPipelineRobustnessCreateInfoEXT` applies to all accesses emanating
  from the shader code of this shader stage

If `VkPipelineRobustnessCreateInfoEXT` is specified for both a pipeline
and a pipeline stage, the `VkPipelineRobustnessCreateInfoEXT` specified
for the pipeline stage will take precedence.

When `VkPipelineRobustnessCreateInfoEXT` is specified for a pipeline, it
only affects the subset of the pipeline that is specified by the create
info, as opposed to subsets linked from pipeline libraries. For
[VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineCreateInfo.html), that
subset is specified by
[VkGraphicsPipelineLibraryCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineLibraryCreateInfoEXT.html)::`flags`.
For
[VkRayTracingPipelineCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingPipelineCreateInfoKHR.html),
that subset is specified by the specific stages in
[VkRayTracingPipelineCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingPipelineCreateInfoKHR.html)::`pStages`.

Valid Usage

- <a
  href="#VUID-VkPipelineRobustnessCreateInfoEXT-pipelineRobustness-06926"
  id="VUID-VkPipelineRobustnessCreateInfoEXT-pipelineRobustness-06926"></a>
  VUID-VkPipelineRobustnessCreateInfoEXT-pipelineRobustness-06926  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-pipelineRobustness"
  target="_blank" rel="noopener"><code>pipelineRobustness</code></a>
  feature is not enabled, `storageBuffers` **must** be
  `VK_PIPELINE_ROBUSTNESS_BUFFER_BEHAVIOR_DEVICE_DEFAULT_EXT`

- <a
  href="#VUID-VkPipelineRobustnessCreateInfoEXT-pipelineRobustness-06927"
  id="VUID-VkPipelineRobustnessCreateInfoEXT-pipelineRobustness-06927"></a>
  VUID-VkPipelineRobustnessCreateInfoEXT-pipelineRobustness-06927  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-pipelineRobustness"
  target="_blank" rel="noopener"><code>pipelineRobustness</code></a>
  feature is not enabled, `uniformBuffers` **must** be
  `VK_PIPELINE_ROBUSTNESS_BUFFER_BEHAVIOR_DEVICE_DEFAULT_EXT`

- <a
  href="#VUID-VkPipelineRobustnessCreateInfoEXT-pipelineRobustness-06928"
  id="VUID-VkPipelineRobustnessCreateInfoEXT-pipelineRobustness-06928"></a>
  VUID-VkPipelineRobustnessCreateInfoEXT-pipelineRobustness-06928  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-pipelineRobustness"
  target="_blank" rel="noopener"><code>pipelineRobustness</code></a>
  feature is not enabled, `vertexInputs` **must** be
  `VK_PIPELINE_ROBUSTNESS_BUFFER_BEHAVIOR_DEVICE_DEFAULT_EXT`

- <a
  href="#VUID-VkPipelineRobustnessCreateInfoEXT-pipelineRobustness-06929"
  id="VUID-VkPipelineRobustnessCreateInfoEXT-pipelineRobustness-06929"></a>
  VUID-VkPipelineRobustnessCreateInfoEXT-pipelineRobustness-06929  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-pipelineRobustness"
  target="_blank" rel="noopener"><code>pipelineRobustness</code></a>
  feature is not enabled, `images` **must** be
  `VK_PIPELINE_ROBUSTNESS_IMAGE_BEHAVIOR_DEVICE_DEFAULT_EXT`

- <a
  href="#VUID-VkPipelineRobustnessCreateInfoEXT-robustImageAccess-06930"
  id="VUID-VkPipelineRobustnessCreateInfoEXT-robustImageAccess-06930"></a>
  VUID-VkPipelineRobustnessCreateInfoEXT-robustImageAccess-06930  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-robustImageAccess"
  target="_blank" rel="noopener"><code>robustImageAccess</code></a>
  feature is not supported, `images` **must** not be
  `VK_PIPELINE_ROBUSTNESS_IMAGE_BEHAVIOR_ROBUST_IMAGE_ACCESS_EXT`

- <a
  href="#VUID-VkPipelineRobustnessCreateInfoEXT-robustBufferAccess2-06931"
  id="VUID-VkPipelineRobustnessCreateInfoEXT-robustBufferAccess2-06931"></a>
  VUID-VkPipelineRobustnessCreateInfoEXT-robustBufferAccess2-06931  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-robustBufferAccess2"
  target="_blank" rel="noopener"><code>robustBufferAccess2</code></a>
  feature is not supported, `storageBuffers` **must** not be
  `VK_PIPELINE_ROBUSTNESS_BUFFER_BEHAVIOR_ROBUST_BUFFER_ACCESS_2_EXT`

- <a
  href="#VUID-VkPipelineRobustnessCreateInfoEXT-robustBufferAccess2-06932"
  id="VUID-VkPipelineRobustnessCreateInfoEXT-robustBufferAccess2-06932"></a>
  VUID-VkPipelineRobustnessCreateInfoEXT-robustBufferAccess2-06932  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-robustBufferAccess2"
  target="_blank" rel="noopener"><code>robustBufferAccess2</code></a>
  feature is not supported, `uniformBuffers` **must** not be
  `VK_PIPELINE_ROBUSTNESS_BUFFER_BEHAVIOR_ROBUST_BUFFER_ACCESS_2_EXT`

- <a
  href="#VUID-VkPipelineRobustnessCreateInfoEXT-robustBufferAccess2-06933"
  id="VUID-VkPipelineRobustnessCreateInfoEXT-robustBufferAccess2-06933"></a>
  VUID-VkPipelineRobustnessCreateInfoEXT-robustBufferAccess2-06933  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-robustBufferAccess2"
  target="_blank" rel="noopener"><code>robustBufferAccess2</code></a>
  feature is not supported, `vertexInputs` **must** not be
  `VK_PIPELINE_ROBUSTNESS_BUFFER_BEHAVIOR_ROBUST_BUFFER_ACCESS_2_EXT`

- <a
  href="#VUID-VkPipelineRobustnessCreateInfoEXT-robustImageAccess2-06934"
  id="VUID-VkPipelineRobustnessCreateInfoEXT-robustImageAccess2-06934"></a>
  VUID-VkPipelineRobustnessCreateInfoEXT-robustImageAccess2-06934  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-robustImageAccess2"
  target="_blank" rel="noopener"><code>robustImageAccess2</code></a>
  feature is not supported, `images` **must** not be
  `VK_PIPELINE_ROBUSTNESS_IMAGE_BEHAVIOR_ROBUST_IMAGE_ACCESS_2_EXT`

Valid Usage (Implicit)

- <a href="#VUID-VkPipelineRobustnessCreateInfoEXT-sType-sType"
  id="VUID-VkPipelineRobustnessCreateInfoEXT-sType-sType"></a>
  VUID-VkPipelineRobustnessCreateInfoEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PIPELINE_ROBUSTNESS_CREATE_INFO_EXT`

- <a
  href="#VUID-VkPipelineRobustnessCreateInfoEXT-storageBuffers-parameter"
  id="VUID-VkPipelineRobustnessCreateInfoEXT-storageBuffers-parameter"></a>
  VUID-VkPipelineRobustnessCreateInfoEXT-storageBuffers-parameter  
  `storageBuffers` **must** be a valid
  [VkPipelineRobustnessBufferBehaviorEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRobustnessBufferBehaviorEXT.html)
  value

- <a
  href="#VUID-VkPipelineRobustnessCreateInfoEXT-uniformBuffers-parameter"
  id="VUID-VkPipelineRobustnessCreateInfoEXT-uniformBuffers-parameter"></a>
  VUID-VkPipelineRobustnessCreateInfoEXT-uniformBuffers-parameter  
  `uniformBuffers` **must** be a valid
  [VkPipelineRobustnessBufferBehaviorEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRobustnessBufferBehaviorEXT.html)
  value

- <a href="#VUID-VkPipelineRobustnessCreateInfoEXT-vertexInputs-parameter"
  id="VUID-VkPipelineRobustnessCreateInfoEXT-vertexInputs-parameter"></a>
  VUID-VkPipelineRobustnessCreateInfoEXT-vertexInputs-parameter  
  `vertexInputs` **must** be a valid
  [VkPipelineRobustnessBufferBehaviorEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRobustnessBufferBehaviorEXT.html)
  value

- <a href="#VUID-VkPipelineRobustnessCreateInfoEXT-images-parameter"
  id="VUID-VkPipelineRobustnessCreateInfoEXT-images-parameter"></a>
  VUID-VkPipelineRobustnessCreateInfoEXT-images-parameter  
  `images` **must** be a valid
  [VkPipelineRobustnessImageBehaviorEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRobustnessImageBehaviorEXT.html)
  value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_pipeline_robustness](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_pipeline_robustness.html),
[VkPipelineRobustnessBufferBehaviorEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRobustnessBufferBehaviorEXT.html),
[VkPipelineRobustnessImageBehaviorEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRobustnessImageBehaviorEXT.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPipelineRobustnessCreateInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
