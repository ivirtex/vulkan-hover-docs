# VkPipelineRobustnessCreateInfo(3) Manual Page

## Name

VkPipelineRobustnessCreateInfo - Structure controlling the robustness of a newly created pipeline shader stage



## [](#_c_specification)C Specification

The `VkPipelineRobustnessCreateInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_4
typedef struct VkPipelineRobustnessCreateInfo {
    VkStructureType                       sType;
    const void*                           pNext;
    VkPipelineRobustnessBufferBehavior    storageBuffers;
    VkPipelineRobustnessBufferBehavior    uniformBuffers;
    VkPipelineRobustnessBufferBehavior    vertexInputs;
    VkPipelineRobustnessImageBehavior     images;
} VkPipelineRobustnessCreateInfo;
```

or the equivalent

```c++
// Provided by VK_EXT_pipeline_robustness
typedef VkPipelineRobustnessCreateInfo VkPipelineRobustnessCreateInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `storageBuffers` sets the behavior of out of bounds accesses made to resources bound as:
  
  - `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER`
  - `VK_DESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFER`
  - `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC`
- `uniformBuffers` describes the behavior of out of bounds accesses made to resources bound as:
  
  - `VK_DESCRIPTOR_TYPE_UNIFORM_TEXEL_BUFFER`
  - `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER`
  - `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC`
  - `VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK`
- `vertexInputs` describes the behavior of out of bounds accesses made to vertex input attributes
- `images` describes the behavior of out of bounds accesses made to resources bound as:
  
  - `VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE`
  - `VK_DESCRIPTOR_TYPE_STORAGE_IMAGE`

## [](#_description)Description

Resources bound as `VK_DESCRIPTOR_TYPE_MUTABLE_EXT` will have the robustness behavior that covers its active descriptor type.

The scope of the effect of `VkPipelineRobustnessCreateInfo` depends on which structure’s `pNext` chain it is included in.

- `VkGraphicsPipelineCreateInfo`, `VkRayTracingPipelineCreateInfoKHR`, `VkComputePipelineCreateInfo`:  
  The robustness behavior described by `VkPipelineRobustnessCreateInfo` applies to all accesses through this pipeline
- `VkPipelineShaderStageCreateInfo`:  
  The robustness behavior described by `VkPipelineRobustnessCreateInfo` applies to all accesses emanating from the shader code of this shader stage

If `VkPipelineRobustnessCreateInfo` is specified for both a pipeline and a pipeline stage, the `VkPipelineRobustnessCreateInfo` specified for the pipeline stage will take precedence.

When `VkPipelineRobustnessCreateInfo` is specified for a pipeline, it only affects the subset of the pipeline that is specified by the create info, as opposed to subsets linked from pipeline libraries. For [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineCreateInfo.html), that subset is specified by [VkGraphicsPipelineLibraryCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineLibraryCreateInfoEXT.html)::`flags`. For [VkRayTracingPipelineCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineCreateInfoKHR.html), that subset is specified by the specific stages in [VkRayTracingPipelineCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineCreateInfoKHR.html)::`pStages`.

Valid Usage

- [](#VUID-VkPipelineRobustnessCreateInfo-pipelineRobustness-06926)VUID-VkPipelineRobustnessCreateInfo-pipelineRobustness-06926  
  If the [`pipelineRobustness`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-pipelineRobustness) feature is not enabled, `storageBuffers` **must** be `VK_PIPELINE_ROBUSTNESS_BUFFER_BEHAVIOR_DEVICE_DEFAULT`
- [](#VUID-VkPipelineRobustnessCreateInfo-pipelineRobustness-06927)VUID-VkPipelineRobustnessCreateInfo-pipelineRobustness-06927  
  If the [`pipelineRobustness`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-pipelineRobustness) feature is not enabled, `uniformBuffers` **must** be `VK_PIPELINE_ROBUSTNESS_BUFFER_BEHAVIOR_DEVICE_DEFAULT`
- [](#VUID-VkPipelineRobustnessCreateInfo-pipelineRobustness-06928)VUID-VkPipelineRobustnessCreateInfo-pipelineRobustness-06928  
  If the [`pipelineRobustness`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-pipelineRobustness) feature is not enabled, `vertexInputs` **must** be `VK_PIPELINE_ROBUSTNESS_BUFFER_BEHAVIOR_DEVICE_DEFAULT`
- [](#VUID-VkPipelineRobustnessCreateInfo-pipelineRobustness-06929)VUID-VkPipelineRobustnessCreateInfo-pipelineRobustness-06929  
  If the [`pipelineRobustness`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-pipelineRobustness) feature is not enabled, `images` **must** be `VK_PIPELINE_ROBUSTNESS_IMAGE_BEHAVIOR_DEVICE_DEFAULT`
- [](#VUID-VkPipelineRobustnessCreateInfo-robustImageAccess-06930)VUID-VkPipelineRobustnessCreateInfo-robustImageAccess-06930  
  If the [`robustImageAccess`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-robustImageAccess) feature is not supported, `images` **must** not be `VK_PIPELINE_ROBUSTNESS_IMAGE_BEHAVIOR_ROBUST_IMAGE_ACCESS`
- [](#VUID-VkPipelineRobustnessCreateInfo-robustBufferAccess2-06931)VUID-VkPipelineRobustnessCreateInfo-robustBufferAccess2-06931  
  If the [`robustBufferAccess2`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-robustBufferAccess2) feature is not supported, `storageBuffers` **must** not be `VK_PIPELINE_ROBUSTNESS_BUFFER_BEHAVIOR_ROBUST_BUFFER_ACCESS_2`
- [](#VUID-VkPipelineRobustnessCreateInfo-robustBufferAccess2-06932)VUID-VkPipelineRobustnessCreateInfo-robustBufferAccess2-06932  
  If the [`robustBufferAccess2`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-robustBufferAccess2) feature is not supported, `uniformBuffers` **must** not be `VK_PIPELINE_ROBUSTNESS_BUFFER_BEHAVIOR_ROBUST_BUFFER_ACCESS_2`
- [](#VUID-VkPipelineRobustnessCreateInfo-robustBufferAccess2-06933)VUID-VkPipelineRobustnessCreateInfo-robustBufferAccess2-06933  
  If the [`robustBufferAccess2`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-robustBufferAccess2) feature is not supported, `vertexInputs` **must** not be `VK_PIPELINE_ROBUSTNESS_BUFFER_BEHAVIOR_ROBUST_BUFFER_ACCESS_2`
- [](#VUID-VkPipelineRobustnessCreateInfo-robustImageAccess2-06934)VUID-VkPipelineRobustnessCreateInfo-robustImageAccess2-06934  
  If the [`robustImageAccess2`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-robustImageAccess2) feature is not supported, `images` **must** not be `VK_PIPELINE_ROBUSTNESS_IMAGE_BEHAVIOR_ROBUST_IMAGE_ACCESS_2`
- [](#VUID-VkPipelineRobustnessCreateInfo-storageBuffers-10636)VUID-VkPipelineRobustnessCreateInfo-storageBuffers-10636  
  If `storageBuffers` is `VK_PIPELINE_ROBUSTNESS_BUFFER_BEHAVIOR_ROBUST_BUFFER_ACCESS` or `VK_PIPELINE_ROBUSTNESS_BUFFER_BEHAVIOR_ROBUST_BUFFER_ACCESS_2`, and either the [`descriptorBindingStorageBufferUpdateAfterBind`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-descriptorBindingStorageBufferUpdateAfterBind) feature or the [`descriptorBindingStorageTexelBufferUpdateAfterBind`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-descriptorBindingStorageTexelBufferUpdateAfterBind) feature is enabled on the device, [`robustBufferAccessUpdateAfterBind`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-robustBufferAccessUpdateAfterBind) **must** be `VK_TRUE`
- [](#VUID-VkPipelineRobustnessCreateInfo-uniformBuffers-10637)VUID-VkPipelineRobustnessCreateInfo-uniformBuffers-10637  
  If `uniformBuffers` is `VK_PIPELINE_ROBUSTNESS_BUFFER_BEHAVIOR_ROBUST_BUFFER_ACCESS` or `VK_PIPELINE_ROBUSTNESS_BUFFER_BEHAVIOR_ROBUST_BUFFER_ACCESS_2`, and either the [`descriptorBindingInlineUniformBlockUpdateAfterBind`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-descriptorBindingInlineUniformBlockUpdateAfterBind) feature, the [`descriptorBindingUniformBufferUpdateAfterBind`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-descriptorBindingUniformBufferUpdateAfterBind) feature, or the [`descriptorBindingUniformTexelBufferUpdateAfterBind`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-descriptorBindingUniformTexelBufferUpdateAfterBind) feature is enabled on the device, [`robustBufferAccessUpdateAfterBind`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-robustBufferAccessUpdateAfterBind) **must** be `VK_TRUE`
- [](#VUID-VkPipelineRobustnessCreateInfo-images-10638)VUID-VkPipelineRobustnessCreateInfo-images-10638  
  If `images` is `VK_PIPELINE_ROBUSTNESS_IMAGE_BEHAVIOR_ROBUST_IMAGE_ACCESS` or `VK_PIPELINE_ROBUSTNESS_IMAGE_BEHAVIOR_ROBUST_IMAGE_ACCESS_2`, and either the [`descriptorBindingStorageImageUpdateAfterBind`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-descriptorBindingStorageImageUpdateAfterBind) feature or the [`descriptorBindingSampledImageUpdateAfterBind`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-descriptorBindingSampledImageUpdateAfterBind) feature is enabled on the device, [`robustBufferAccessUpdateAfterBind`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-robustBufferAccessUpdateAfterBind) **must** be `VK_TRUE`

Valid Usage (Implicit)

- [](#VUID-VkPipelineRobustnessCreateInfo-sType-sType)VUID-VkPipelineRobustnessCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PIPELINE_ROBUSTNESS_CREATE_INFO`
- [](#VUID-VkPipelineRobustnessCreateInfo-storageBuffers-parameter)VUID-VkPipelineRobustnessCreateInfo-storageBuffers-parameter  
  `storageBuffers` **must** be a valid [VkPipelineRobustnessBufferBehavior](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRobustnessBufferBehavior.html) value
- [](#VUID-VkPipelineRobustnessCreateInfo-uniformBuffers-parameter)VUID-VkPipelineRobustnessCreateInfo-uniformBuffers-parameter  
  `uniformBuffers` **must** be a valid [VkPipelineRobustnessBufferBehavior](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRobustnessBufferBehavior.html) value
- [](#VUID-VkPipelineRobustnessCreateInfo-vertexInputs-parameter)VUID-VkPipelineRobustnessCreateInfo-vertexInputs-parameter  
  `vertexInputs` **must** be a valid [VkPipelineRobustnessBufferBehavior](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRobustnessBufferBehavior.html) value
- [](#VUID-VkPipelineRobustnessCreateInfo-images-parameter)VUID-VkPipelineRobustnessCreateInfo-images-parameter  
  `images` **must** be a valid [VkPipelineRobustnessImageBehavior](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRobustnessImageBehavior.html) value

## [](#_see_also)See Also

[VK\_EXT\_pipeline\_robustness](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_pipeline_robustness.html), [VK\_VERSION\_1\_4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_4.html), [VkPipelineRobustnessBufferBehavior](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRobustnessBufferBehavior.html), [VkPipelineRobustnessImageBehavior](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRobustnessImageBehavior.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPipelineRobustnessCreateInfo)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0