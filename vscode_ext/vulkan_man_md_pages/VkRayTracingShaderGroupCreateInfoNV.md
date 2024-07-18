# VkRayTracingShaderGroupCreateInfoNV(3) Manual Page

## Name

VkRayTracingShaderGroupCreateInfoNV - Structure specifying shaders in a
shader group



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkRayTracingShaderGroupCreateInfoNV` structure is defined as:

``` c
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

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `type` is the type of hit group specified in this structure.

- `generalShader` is the index of the ray generation, miss, or callable
  shader from
  [VkRayTracingPipelineCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingPipelineCreateInfoNV.html)::`pStages`
  in the group if the shader group has `type` of
  `VK_RAY_TRACING_SHADER_GROUP_TYPE_GENERAL_NV`, and
  `VK_SHADER_UNUSED_NV` otherwise.

- `closestHitShader` is the optional index of the closest hit shader
  from
  [VkRayTracingPipelineCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingPipelineCreateInfoNV.html)::`pStages`
  in the group if the shader group has `type` of
  `VK_RAY_TRACING_SHADER_GROUP_TYPE_TRIANGLES_HIT_GROUP_NV` or
  `VK_RAY_TRACING_SHADER_GROUP_TYPE_PROCEDURAL_HIT_GROUP_NV`, and
  `VK_SHADER_UNUSED_NV` otherwise.

- `anyHitShader` is the optional index of the any-hit shader from
  [VkRayTracingPipelineCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingPipelineCreateInfoNV.html)::`pStages`
  in the group if the shader group has `type` of
  `VK_RAY_TRACING_SHADER_GROUP_TYPE_TRIANGLES_HIT_GROUP_NV` or
  `VK_RAY_TRACING_SHADER_GROUP_TYPE_PROCEDURAL_HIT_GROUP_NV`, and
  `VK_SHADER_UNUSED_NV` otherwise.

- `intersectionShader` is the index of the intersection shader from
  [VkRayTracingPipelineCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingPipelineCreateInfoNV.html)::`pStages`
  in the group if the shader group has `type` of
  `VK_RAY_TRACING_SHADER_GROUP_TYPE_PROCEDURAL_HIT_GROUP_NV`, and
  `VK_SHADER_UNUSED_NV` otherwise.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkRayTracingShaderGroupCreateInfoNV-type-02413"
  id="VUID-VkRayTracingShaderGroupCreateInfoNV-type-02413"></a>
  VUID-VkRayTracingShaderGroupCreateInfoNV-type-02413  
  If `type` is `VK_RAY_TRACING_SHADER_GROUP_TYPE_GENERAL_NV` then
  `generalShader` **must** be a valid index into
  [VkRayTracingPipelineCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingPipelineCreateInfoNV.html)::`pStages`
  referring to a shader of `VK_SHADER_STAGE_RAYGEN_BIT_NV`,
  `VK_SHADER_STAGE_MISS_BIT_NV`, or `VK_SHADER_STAGE_CALLABLE_BIT_NV`

- <a href="#VUID-VkRayTracingShaderGroupCreateInfoNV-type-02414"
  id="VUID-VkRayTracingShaderGroupCreateInfoNV-type-02414"></a>
  VUID-VkRayTracingShaderGroupCreateInfoNV-type-02414  
  If `type` is `VK_RAY_TRACING_SHADER_GROUP_TYPE_GENERAL_NV` then
  `closestHitShader`, `anyHitShader`, and `intersectionShader` **must**
  be `VK_SHADER_UNUSED_NV`

- <a href="#VUID-VkRayTracingShaderGroupCreateInfoNV-type-02415"
  id="VUID-VkRayTracingShaderGroupCreateInfoNV-type-02415"></a>
  VUID-VkRayTracingShaderGroupCreateInfoNV-type-02415  
  If `type` is
  `VK_RAY_TRACING_SHADER_GROUP_TYPE_PROCEDURAL_HIT_GROUP_NV` then
  `intersectionShader` **must** be a valid index into
  [VkRayTracingPipelineCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingPipelineCreateInfoNV.html)::`pStages`
  referring to a shader of `VK_SHADER_STAGE_INTERSECTION_BIT_NV`

- <a href="#VUID-VkRayTracingShaderGroupCreateInfoNV-type-02416"
  id="VUID-VkRayTracingShaderGroupCreateInfoNV-type-02416"></a>
  VUID-VkRayTracingShaderGroupCreateInfoNV-type-02416  
  If `type` is `VK_RAY_TRACING_SHADER_GROUP_TYPE_TRIANGLES_HIT_GROUP_NV`
  then `intersectionShader` **must** be `VK_SHADER_UNUSED_NV`

- <a
  href="#VUID-VkRayTracingShaderGroupCreateInfoNV-closestHitShader-02417"
  id="VUID-VkRayTracingShaderGroupCreateInfoNV-closestHitShader-02417"></a>
  VUID-VkRayTracingShaderGroupCreateInfoNV-closestHitShader-02417  
  `closestHitShader` **must** be either `VK_SHADER_UNUSED_NV` or a valid
  index into
  [VkRayTracingPipelineCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingPipelineCreateInfoNV.html)::`pStages`
  referring to a shader of `VK_SHADER_STAGE_CLOSEST_HIT_BIT_NV`

- <a href="#VUID-VkRayTracingShaderGroupCreateInfoNV-anyHitShader-02418"
  id="VUID-VkRayTracingShaderGroupCreateInfoNV-anyHitShader-02418"></a>
  VUID-VkRayTracingShaderGroupCreateInfoNV-anyHitShader-02418  
  `anyHitShader` **must** be either `VK_SHADER_UNUSED_NV` or a valid
  index into
  [VkRayTracingPipelineCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingPipelineCreateInfoNV.html)::`pStages`
  referring to a shader of `VK_SHADER_STAGE_ANY_HIT_BIT_NV`

Valid Usage (Implicit)

- <a href="#VUID-VkRayTracingShaderGroupCreateInfoNV-sType-sType"
  id="VUID-VkRayTracingShaderGroupCreateInfoNV-sType-sType"></a>
  VUID-VkRayTracingShaderGroupCreateInfoNV-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_RAY_TRACING_SHADER_GROUP_CREATE_INFO_NV`

- <a href="#VUID-VkRayTracingShaderGroupCreateInfoNV-pNext-pNext"
  id="VUID-VkRayTracingShaderGroupCreateInfoNV-pNext-pNext"></a>
  VUID-VkRayTracingShaderGroupCreateInfoNV-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkRayTracingShaderGroupCreateInfoNV-type-parameter"
  id="VUID-VkRayTracingShaderGroupCreateInfoNV-type-parameter"></a>
  VUID-VkRayTracingShaderGroupCreateInfoNV-type-parameter  
  `type` **must** be a valid
  [VkRayTracingShaderGroupTypeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingShaderGroupTypeKHR.html)
  value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_ray_tracing](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_ray_tracing.html),
[VkRayTracingPipelineCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingPipelineCreateInfoNV.html),
[VkRayTracingShaderGroupTypeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingShaderGroupTypeKHR.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkRayTracingShaderGroupCreateInfoNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
