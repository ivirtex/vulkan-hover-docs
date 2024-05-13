# VkShaderStageFlagBits(3) Manual Page

## Name

VkShaderStageFlagBits - Bitmask specifying a pipeline stage



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **can** be set by commands and structures, specifying one or
more shader stages, are:

``` c
// Provided by VK_VERSION_1_0
typedef enum VkShaderStageFlagBits {
    VK_SHADER_STAGE_VERTEX_BIT = 0x00000001,
    VK_SHADER_STAGE_TESSELLATION_CONTROL_BIT = 0x00000002,
    VK_SHADER_STAGE_TESSELLATION_EVALUATION_BIT = 0x00000004,
    VK_SHADER_STAGE_GEOMETRY_BIT = 0x00000008,
    VK_SHADER_STAGE_FRAGMENT_BIT = 0x00000010,
    VK_SHADER_STAGE_COMPUTE_BIT = 0x00000020,
    VK_SHADER_STAGE_ALL_GRAPHICS = 0x0000001F,
    VK_SHADER_STAGE_ALL = 0x7FFFFFFF,
  // Provided by VK_KHR_ray_tracing_pipeline
    VK_SHADER_STAGE_RAYGEN_BIT_KHR = 0x00000100,
  // Provided by VK_KHR_ray_tracing_pipeline
    VK_SHADER_STAGE_ANY_HIT_BIT_KHR = 0x00000200,
  // Provided by VK_KHR_ray_tracing_pipeline
    VK_SHADER_STAGE_CLOSEST_HIT_BIT_KHR = 0x00000400,
  // Provided by VK_KHR_ray_tracing_pipeline
    VK_SHADER_STAGE_MISS_BIT_KHR = 0x00000800,
  // Provided by VK_KHR_ray_tracing_pipeline
    VK_SHADER_STAGE_INTERSECTION_BIT_KHR = 0x00001000,
  // Provided by VK_KHR_ray_tracing_pipeline
    VK_SHADER_STAGE_CALLABLE_BIT_KHR = 0x00002000,
  // Provided by VK_EXT_mesh_shader
    VK_SHADER_STAGE_TASK_BIT_EXT = 0x00000040,
  // Provided by VK_EXT_mesh_shader
    VK_SHADER_STAGE_MESH_BIT_EXT = 0x00000080,
  // Provided by VK_HUAWEI_subpass_shading
    VK_SHADER_STAGE_SUBPASS_SHADING_BIT_HUAWEI = 0x00004000,
  // Provided by VK_HUAWEI_cluster_culling_shader
    VK_SHADER_STAGE_CLUSTER_CULLING_BIT_HUAWEI = 0x00080000,
  // Provided by VK_NV_ray_tracing
    VK_SHADER_STAGE_RAYGEN_BIT_NV = VK_SHADER_STAGE_RAYGEN_BIT_KHR,
  // Provided by VK_NV_ray_tracing
    VK_SHADER_STAGE_ANY_HIT_BIT_NV = VK_SHADER_STAGE_ANY_HIT_BIT_KHR,
  // Provided by VK_NV_ray_tracing
    VK_SHADER_STAGE_CLOSEST_HIT_BIT_NV = VK_SHADER_STAGE_CLOSEST_HIT_BIT_KHR,
  // Provided by VK_NV_ray_tracing
    VK_SHADER_STAGE_MISS_BIT_NV = VK_SHADER_STAGE_MISS_BIT_KHR,
  // Provided by VK_NV_ray_tracing
    VK_SHADER_STAGE_INTERSECTION_BIT_NV = VK_SHADER_STAGE_INTERSECTION_BIT_KHR,
  // Provided by VK_NV_ray_tracing
    VK_SHADER_STAGE_CALLABLE_BIT_NV = VK_SHADER_STAGE_CALLABLE_BIT_KHR,
  // Provided by VK_NV_mesh_shader
    VK_SHADER_STAGE_TASK_BIT_NV = VK_SHADER_STAGE_TASK_BIT_EXT,
  // Provided by VK_NV_mesh_shader
    VK_SHADER_STAGE_MESH_BIT_NV = VK_SHADER_STAGE_MESH_BIT_EXT,
} VkShaderStageFlagBits;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_SHADER_STAGE_VERTEX_BIT` specifies the vertex stage.

- `VK_SHADER_STAGE_TESSELLATION_CONTROL_BIT` specifies the tessellation
  control stage.

- `VK_SHADER_STAGE_TESSELLATION_EVALUATION_BIT` specifies the
  tessellation evaluation stage.

- `VK_SHADER_STAGE_GEOMETRY_BIT` specifies the geometry stage.

- `VK_SHADER_STAGE_FRAGMENT_BIT` specifies the fragment stage.

- `VK_SHADER_STAGE_COMPUTE_BIT` specifies the compute stage.

- `VK_SHADER_STAGE_ALL_GRAPHICS` is a combination of bits used as
  shorthand to specify all graphics stages defined above (excluding the
  compute stage).

- `VK_SHADER_STAGE_ALL` is a combination of bits used as shorthand to
  specify all shader stages supported by the device, including all
  additional stages which are introduced by extensions.

- `VK_SHADER_STAGE_TASK_BIT_EXT` specifies the task stage.

- `VK_SHADER_STAGE_MESH_BIT_EXT` specifies the mesh stage.

- `VK_SHADER_STAGE_CLUSTER_CULLING_BIT_HUAWEI` specifies the cluster
  culling stage.

- `VK_SHADER_STAGE_RAYGEN_BIT_KHR` specifies the ray generation stage.

- `VK_SHADER_STAGE_ANY_HIT_BIT_KHR` specifies the any-hit stage.

- `VK_SHADER_STAGE_CLOSEST_HIT_BIT_KHR` specifies the closest hit stage.

- `VK_SHADER_STAGE_MISS_BIT_KHR` specifies the miss stage.

- `VK_SHADER_STAGE_INTERSECTION_BIT_KHR` specifies the intersection
  stage.

- `VK_SHADER_STAGE_CALLABLE_BIT_KHR` specifies the callable stage.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p><code>VK_SHADER_STAGE_ALL_GRAPHICS</code> only includes the original
five graphics stages included in Vulkan 1.0, and not any stages added by
extensions. Thus, it may not have the desired effect in all
cases.</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkPipelineShaderStageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageCreateInfo.html),
[VkShaderCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderCreateInfoEXT.html),
[VkShaderStageFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderStageFlags.html),
[vkCmdBindShadersEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBindShadersEXT.html),
[vkGetShaderInfoAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetShaderInfoAMD.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkShaderStageFlagBits"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
