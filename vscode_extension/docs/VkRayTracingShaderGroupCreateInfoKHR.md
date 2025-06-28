# VkRayTracingShaderGroupCreateInfoKHR(3) Manual Page

## Name

VkRayTracingShaderGroupCreateInfoKHR - Structure specifying shaders in a shader group



## [](#_c_specification)C Specification

The `VkRayTracingShaderGroupCreateInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_ray_tracing_pipeline
typedef struct VkRayTracingShaderGroupCreateInfoKHR {
    VkStructureType                   sType;
    const void*                       pNext;
    VkRayTracingShaderGroupTypeKHR    type;
    uint32_t                          generalShader;
    uint32_t                          closestHitShader;
    uint32_t                          anyHitShader;
    uint32_t                          intersectionShader;
    const void*                       pShaderGroupCaptureReplayHandle;
} VkRayTracingShaderGroupCreateInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `type` is the type of hit group specified in this structure.
- `generalShader` is the index of the ray generation, miss, or callable shader from [VkRayTracingPipelineCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineCreateInfoKHR.html)::`pStages` in the group if the shader group has `type` of `VK_RAY_TRACING_SHADER_GROUP_TYPE_GENERAL_KHR`, and `VK_SHADER_UNUSED_KHR` otherwise.
- `closestHitShader` is the optional index of the closest hit shader from [VkRayTracingPipelineCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineCreateInfoKHR.html)::`pStages` in the group if the shader group has `type` of `VK_RAY_TRACING_SHADER_GROUP_TYPE_TRIANGLES_HIT_GROUP_KHR` or `VK_RAY_TRACING_SHADER_GROUP_TYPE_PROCEDURAL_HIT_GROUP_KHR`, and `VK_SHADER_UNUSED_KHR` otherwise.
- `anyHitShader` is the optional index of the any-hit shader from [VkRayTracingPipelineCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineCreateInfoKHR.html)::`pStages` in the group if the shader group has `type` of `VK_RAY_TRACING_SHADER_GROUP_TYPE_TRIANGLES_HIT_GROUP_KHR` or `VK_RAY_TRACING_SHADER_GROUP_TYPE_PROCEDURAL_HIT_GROUP_KHR`, and `VK_SHADER_UNUSED_KHR` otherwise.
- `intersectionShader` is the index of the intersection shader from [VkRayTracingPipelineCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineCreateInfoKHR.html)::`pStages` in the group if the shader group has `type` of `VK_RAY_TRACING_SHADER_GROUP_TYPE_PROCEDURAL_HIT_GROUP_KHR`, and `VK_SHADER_UNUSED_KHR` otherwise.
- `pShaderGroupCaptureReplayHandle` is `NULL` or a pointer to replay information for this shader group queried from [vkGetRayTracingCaptureReplayShaderGroupHandlesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetRayTracingCaptureReplayShaderGroupHandlesKHR.html), as described in [Ray Tracing Capture Replay](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#ray-tracing-capture-replay). Ignored if [VkPhysicalDeviceRayTracingPipelineFeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceRayTracingPipelineFeaturesKHR.html)::`rayTracingPipelineShaderGroupHandleCaptureReplay` is `VK_FALSE`.

## [](#_description)Description

If the pipeline is created with `VK_PIPELINE_CREATE_LIBRARY_BIT_KHR` and the [pipelineLibraryGroupHandles](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-pipelineLibraryGroupHandles) feature is enabled, `pShaderGroupCaptureReplayHandle` is inherited by all pipelines which link against this pipeline and remains bitwise identical for any pipeline which references this pipeline library.

Valid Usage

- [](#VUID-VkRayTracingShaderGroupCreateInfoKHR-type-03474)VUID-VkRayTracingShaderGroupCreateInfoKHR-type-03474  
  If `type` is `VK_RAY_TRACING_SHADER_GROUP_TYPE_GENERAL_KHR` then `generalShader` **must** be a valid index into [VkRayTracingPipelineCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineCreateInfoKHR.html)::`pStages` referring to a shader of `VK_SHADER_STAGE_RAYGEN_BIT_KHR`, `VK_SHADER_STAGE_MISS_BIT_KHR`, or `VK_SHADER_STAGE_CALLABLE_BIT_KHR`
- [](#VUID-VkRayTracingShaderGroupCreateInfoKHR-type-03475)VUID-VkRayTracingShaderGroupCreateInfoKHR-type-03475  
  If `type` is `VK_RAY_TRACING_SHADER_GROUP_TYPE_GENERAL_KHR` then `closestHitShader`, `anyHitShader`, and `intersectionShader` **must** be `VK_SHADER_UNUSED_KHR`
- [](#VUID-VkRayTracingShaderGroupCreateInfoKHR-type-03476)VUID-VkRayTracingShaderGroupCreateInfoKHR-type-03476  
  If `type` is `VK_RAY_TRACING_SHADER_GROUP_TYPE_PROCEDURAL_HIT_GROUP_KHR` then `intersectionShader` **must** be a valid index into [VkRayTracingPipelineCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineCreateInfoKHR.html)::`pStages` referring to a shader of `VK_SHADER_STAGE_INTERSECTION_BIT_KHR`
- [](#VUID-VkRayTracingShaderGroupCreateInfoKHR-type-03477)VUID-VkRayTracingShaderGroupCreateInfoKHR-type-03477  
  If `type` is `VK_RAY_TRACING_SHADER_GROUP_TYPE_TRIANGLES_HIT_GROUP_KHR` then `intersectionShader` **must** be `VK_SHADER_UNUSED_KHR`
- [](#VUID-VkRayTracingShaderGroupCreateInfoKHR-closestHitShader-03478)VUID-VkRayTracingShaderGroupCreateInfoKHR-closestHitShader-03478  
  `closestHitShader` **must** be either `VK_SHADER_UNUSED_KHR` or a valid index into [VkRayTracingPipelineCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineCreateInfoKHR.html)::`pStages` referring to a shader of `VK_SHADER_STAGE_CLOSEST_HIT_BIT_KHR`
- [](#VUID-VkRayTracingShaderGroupCreateInfoKHR-anyHitShader-03479)VUID-VkRayTracingShaderGroupCreateInfoKHR-anyHitShader-03479  
  `anyHitShader` **must** be either `VK_SHADER_UNUSED_KHR` or a valid index into [VkRayTracingPipelineCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineCreateInfoKHR.html)::`pStages` referring to a shader of `VK_SHADER_STAGE_ANY_HIT_BIT_KHR`
- [](#VUID-VkRayTracingShaderGroupCreateInfoKHR-rayTracingPipelineShaderGroupHandleCaptureReplayMixed-03603)VUID-VkRayTracingShaderGroupCreateInfoKHR-rayTracingPipelineShaderGroupHandleCaptureReplayMixed-03603  
  If [VkPhysicalDeviceRayTracingPipelineFeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceRayTracingPipelineFeaturesKHR.html)::`rayTracingPipelineShaderGroupHandleCaptureReplayMixed` is `VK_FALSE` then `pShaderGroupCaptureReplayHandle` **must** not be provided if it has not been provided on a previous call to ray tracing pipeline creation
- [](#VUID-VkRayTracingShaderGroupCreateInfoKHR-rayTracingPipelineShaderGroupHandleCaptureReplayMixed-03604)VUID-VkRayTracingShaderGroupCreateInfoKHR-rayTracingPipelineShaderGroupHandleCaptureReplayMixed-03604  
  If [VkPhysicalDeviceRayTracingPipelineFeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceRayTracingPipelineFeaturesKHR.html)::`rayTracingPipelineShaderGroupHandleCaptureReplayMixed` is `VK_FALSE` then the caller **must** guarantee that no ray tracing pipeline creation commands with `pShaderGroupCaptureReplayHandle` provided execute simultaneously with ray tracing pipeline creation commands without `pShaderGroupCaptureReplayHandle` provided

Valid Usage (Implicit)

- [](#VUID-VkRayTracingShaderGroupCreateInfoKHR-sType-sType)VUID-VkRayTracingShaderGroupCreateInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_RAY_TRACING_SHADER_GROUP_CREATE_INFO_KHR`
- [](#VUID-VkRayTracingShaderGroupCreateInfoKHR-pNext-pNext)VUID-VkRayTracingShaderGroupCreateInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkRayTracingShaderGroupCreateInfoKHR-type-parameter)VUID-VkRayTracingShaderGroupCreateInfoKHR-type-parameter  
  `type` **must** be a valid [VkRayTracingShaderGroupTypeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingShaderGroupTypeKHR.html) value

## [](#_see_also)See Also

[VK\_KHR\_ray\_tracing\_pipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_ray_tracing_pipeline.html), [VkRayTracingPipelineCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineCreateInfoKHR.html), [VkRayTracingShaderGroupTypeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingShaderGroupTypeKHR.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkRayTracingShaderGroupCreateInfoKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0