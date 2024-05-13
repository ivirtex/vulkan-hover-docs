# VkRayTracingShaderGroupTypeKHR(3) Manual Page

## Name

VkRayTracingShaderGroupTypeKHR - Shader group types



## <a href="#_c_specification" class="anchor"></a>C Specification

Possible values of `type` in `VkRayTracingShaderGroupCreateInfoKHR` are:

``` c
// Provided by VK_KHR_ray_tracing_pipeline
typedef enum VkRayTracingShaderGroupTypeKHR {
    VK_RAY_TRACING_SHADER_GROUP_TYPE_GENERAL_KHR = 0,
    VK_RAY_TRACING_SHADER_GROUP_TYPE_TRIANGLES_HIT_GROUP_KHR = 1,
    VK_RAY_TRACING_SHADER_GROUP_TYPE_PROCEDURAL_HIT_GROUP_KHR = 2,
  // Provided by VK_NV_ray_tracing
    VK_RAY_TRACING_SHADER_GROUP_TYPE_GENERAL_NV = VK_RAY_TRACING_SHADER_GROUP_TYPE_GENERAL_KHR,
  // Provided by VK_NV_ray_tracing
    VK_RAY_TRACING_SHADER_GROUP_TYPE_TRIANGLES_HIT_GROUP_NV = VK_RAY_TRACING_SHADER_GROUP_TYPE_TRIANGLES_HIT_GROUP_KHR,
  // Provided by VK_NV_ray_tracing
    VK_RAY_TRACING_SHADER_GROUP_TYPE_PROCEDURAL_HIT_GROUP_NV = VK_RAY_TRACING_SHADER_GROUP_TYPE_PROCEDURAL_HIT_GROUP_KHR,
} VkRayTracingShaderGroupTypeKHR;
```

or the equivalent

``` c
// Provided by VK_NV_ray_tracing
typedef VkRayTracingShaderGroupTypeKHR VkRayTracingShaderGroupTypeNV;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_RAY_TRACING_SHADER_GROUP_TYPE_GENERAL_KHR` indicates a shader
  group with a single `VK_SHADER_STAGE_RAYGEN_BIT_KHR`,
  `VK_SHADER_STAGE_MISS_BIT_KHR`, or `VK_SHADER_STAGE_CALLABLE_BIT_KHR`
  shader in it.

- `VK_RAY_TRACING_SHADER_GROUP_TYPE_TRIANGLES_HIT_GROUP_KHR` specifies a
  shader group that only hits triangles and **must** not contain an
  intersection shader, only closest hit and any-hit shaders.

- `VK_RAY_TRACING_SHADER_GROUP_TYPE_PROCEDURAL_HIT_GROUP_KHR` specifies
  a shader group that only intersects with custom geometry and **must**
  contain an intersection shader and **may** contain closest hit and
  any-hit shaders.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>For current group types, the hit group type could be inferred from
the presence or absence of the intersection shader, but we provide the
type explicitly for future hit groups that do not have that
property.</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_ray_tracing_pipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_ray_tracing_pipeline.html),
[VK_NV_ray_tracing](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_ray_tracing.html),
[VkRayTracingShaderGroupCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingShaderGroupCreateInfoKHR.html),
[VkRayTracingShaderGroupCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingShaderGroupCreateInfoNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkRayTracingShaderGroupTypeKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
