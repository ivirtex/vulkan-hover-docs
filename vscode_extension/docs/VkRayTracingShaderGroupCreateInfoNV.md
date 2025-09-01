# VkRayTracingShaderGroupCreateInfoNV(3) Manual Page

## Name

VkRayTracingShaderGroupCreateInfoNV - Structure specifying shaders in a shader group



## [](#_c_specification)C Specification

The `VkRayTracingShaderGroupCreateInfoNV` structure is defined as:

```c++
// Provided by VK_NV_ray_tracing
typedef struct VkRayTracingShaderGroupCreateInfoNV {
    VkStructureType                   sType;
    const void*                       pNext;
    VkRayTracingShaderGroupTypeKHR    type;
    uint32_t                          generalShader;
    uint32_t                          closestHitShader;
    uint32_t                          anyHitShader;
    uint32_t                          intersectionShader;
} VkRayTracingShaderGroupCreateInfoNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `type` is the type of hit group specified in this structure.
- `generalShader` is the index of the ray generation, miss, or callable shader from [VkRayTracingPipelineCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineCreateInfoNV.html)::`pStages` in the group if the shader group has `type` of `VK_RAY_TRACING_SHADER_GROUP_TYPE_GENERAL_NV`, and `VK_SHADER_UNUSED_NV` otherwise.
- `closestHitShader` is the optional index of the closest hit shader from [VkRayTracingPipelineCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineCreateInfoNV.html)::`pStages` in the group if the shader group has `type` of `VK_RAY_TRACING_SHADER_GROUP_TYPE_TRIANGLES_HIT_GROUP_NV` or `VK_RAY_TRACING_SHADER_GROUP_TYPE_PROCEDURAL_HIT_GROUP_NV`, and `VK_SHADER_UNUSED_NV` otherwise.
- `anyHitShader` is the optional index of the any-hit shader from [VkRayTracingPipelineCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineCreateInfoNV.html)::`pStages` in the group if the shader group has `type` of `VK_RAY_TRACING_SHADER_GROUP_TYPE_TRIANGLES_HIT_GROUP_NV` or `VK_RAY_TRACING_SHADER_GROUP_TYPE_PROCEDURAL_HIT_GROUP_NV`, and `VK_SHADER_UNUSED_NV` otherwise.
- `intersectionShader` is the index of the intersection shader from [VkRayTracingPipelineCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineCreateInfoNV.html)::`pStages` in the group if the shader group has `type` of `VK_RAY_TRACING_SHADER_GROUP_TYPE_PROCEDURAL_HIT_GROUP_NV`, and `VK_SHADER_UNUSED_NV` otherwise.

## [](#_description)Description

Valid Usage

- [](#VUID-VkRayTracingShaderGroupCreateInfoNV-type-02413)VUID-VkRayTracingShaderGroupCreateInfoNV-type-02413  
  If `type` is `VK_RAY_TRACING_SHADER_GROUP_TYPE_GENERAL_NV` then `generalShader` **must** be a valid index into [VkRayTracingPipelineCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineCreateInfoNV.html)::`pStages` referring to a shader of `VK_SHADER_STAGE_RAYGEN_BIT_NV`, `VK_SHADER_STAGE_MISS_BIT_NV`, or `VK_SHADER_STAGE_CALLABLE_BIT_NV`
- [](#VUID-VkRayTracingShaderGroupCreateInfoNV-type-02414)VUID-VkRayTracingShaderGroupCreateInfoNV-type-02414  
  If `type` is `VK_RAY_TRACING_SHADER_GROUP_TYPE_GENERAL_NV` then `closestHitShader`, `anyHitShader`, and `intersectionShader` **must** be `VK_SHADER_UNUSED_NV`
- [](#VUID-VkRayTracingShaderGroupCreateInfoNV-type-02415)VUID-VkRayTracingShaderGroupCreateInfoNV-type-02415  
  If `type` is `VK_RAY_TRACING_SHADER_GROUP_TYPE_PROCEDURAL_HIT_GROUP_NV` then `intersectionShader` **must** be a valid index into [VkRayTracingPipelineCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineCreateInfoNV.html)::`pStages` referring to a shader of `VK_SHADER_STAGE_INTERSECTION_BIT_NV`
- [](#VUID-VkRayTracingShaderGroupCreateInfoNV-type-02416)VUID-VkRayTracingShaderGroupCreateInfoNV-type-02416  
  If `type` is `VK_RAY_TRACING_SHADER_GROUP_TYPE_TRIANGLES_HIT_GROUP_NV` then `intersectionShader` **must** be `VK_SHADER_UNUSED_NV`
- [](#VUID-VkRayTracingShaderGroupCreateInfoNV-closestHitShader-02417)VUID-VkRayTracingShaderGroupCreateInfoNV-closestHitShader-02417  
  `closestHitShader` **must** be either `VK_SHADER_UNUSED_NV` or a valid index into [VkRayTracingPipelineCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineCreateInfoNV.html)::`pStages` referring to a shader of `VK_SHADER_STAGE_CLOSEST_HIT_BIT_NV`
- [](#VUID-VkRayTracingShaderGroupCreateInfoNV-anyHitShader-02418)VUID-VkRayTracingShaderGroupCreateInfoNV-anyHitShader-02418  
  `anyHitShader` **must** be either `VK_SHADER_UNUSED_NV` or a valid index into [VkRayTracingPipelineCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineCreateInfoNV.html)::`pStages` referring to a shader of `VK_SHADER_STAGE_ANY_HIT_BIT_NV`

Valid Usage (Implicit)

- [](#VUID-VkRayTracingShaderGroupCreateInfoNV-sType-sType)VUID-VkRayTracingShaderGroupCreateInfoNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_RAY_TRACING_SHADER_GROUP_CREATE_INFO_NV`
- [](#VUID-VkRayTracingShaderGroupCreateInfoNV-pNext-pNext)VUID-VkRayTracingShaderGroupCreateInfoNV-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkRayTracingShaderGroupCreateInfoNV-type-parameter)VUID-VkRayTracingShaderGroupCreateInfoNV-type-parameter  
  `type` **must** be a valid [VkRayTracingShaderGroupTypeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingShaderGroupTypeKHR.html) value

## [](#_see_also)See Also

[VK\_NV\_ray\_tracing](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_ray_tracing.html), [VkRayTracingPipelineCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineCreateInfoNV.html), [VkRayTracingShaderGroupTypeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingShaderGroupTypeKHR.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkRayTracingShaderGroupCreateInfoNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0