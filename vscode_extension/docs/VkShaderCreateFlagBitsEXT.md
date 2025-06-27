# VkShaderCreateFlagBitsEXT(3) Manual Page

## Name

VkShaderCreateFlagBitsEXT - Bitmask controlling how a shader object is
created



## <a href="#_c_specification" class="anchor"></a>C Specification

Possible values of the `flags` member of
[VkShaderCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderCreateInfoEXT.html) specifying how a
shader object is created, are:

``` c
// Provided by VK_EXT_shader_object
typedef enum VkShaderCreateFlagBitsEXT {
    VK_SHADER_CREATE_LINK_STAGE_BIT_EXT = 0x00000001,
  // Provided by VK_EXT_shader_object with VK_EXT_subgroup_size_control or VK_VERSION_1_3
    VK_SHADER_CREATE_ALLOW_VARYING_SUBGROUP_SIZE_BIT_EXT = 0x00000002,
  // Provided by VK_EXT_shader_object with VK_EXT_subgroup_size_control or VK_VERSION_1_3
    VK_SHADER_CREATE_REQUIRE_FULL_SUBGROUPS_BIT_EXT = 0x00000004,
  // Provided by VK_EXT_shader_object with VK_EXT_mesh_shader or VK_NV_mesh_shader
    VK_SHADER_CREATE_NO_TASK_SHADER_BIT_EXT = 0x00000008,
  // Provided by VK_EXT_shader_object with VK_KHR_device_group or VK_VERSION_1_1
    VK_SHADER_CREATE_DISPATCH_BASE_BIT_EXT = 0x00000010,
  // Provided by VK_KHR_fragment_shading_rate with VK_EXT_shader_object
    VK_SHADER_CREATE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_EXT = 0x00000020,
  // Provided by VK_EXT_fragment_density_map with VK_EXT_shader_object
    VK_SHADER_CREATE_FRAGMENT_DENSITY_MAP_ATTACHMENT_BIT_EXT = 0x00000040,
} VkShaderCreateFlagBitsEXT;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_SHADER_CREATE_LINK_STAGE_BIT_EXT` specifies that a shader is
  linked to all other shaders created in the same
  [vkCreateShadersEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateShadersEXT.html) call whose
  [VkShaderCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderCreateInfoEXT.html) structures'
  `flags` include `VK_SHADER_CREATE_LINK_STAGE_BIT_EXT`.

- `VK_SHADER_CREATE_ALLOW_VARYING_SUBGROUP_SIZE_BIT_EXT` specifies that
  the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-sgs"
  target="_blank" rel="noopener"><code>SubgroupSize</code></a> **may**
  vary in a task, mesh, or compute shader.

- `VK_SHADER_CREATE_REQUIRE_FULL_SUBGROUPS_BIT_EXT` specifies that the
  subgroup sizes **must** be launched with all invocations active in a
  task, mesh, or compute shader.

- `VK_SHADER_CREATE_NO_TASK_SHADER_BIT_EXT` specifies that a mesh shader
  **must** only be used without a task shader. Otherwise, the mesh
  shader **must** only be used with a task shader.

- `VK_SHADER_CREATE_DISPATCH_BASE_BIT_EXT` specifies that a compute
  shader **can** be used with
  [vkCmdDispatchBase](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDispatchBase.html) with a non-zero base
  workgroup.

- `VK_SHADER_CREATE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_EXT` specifies
  that a fragment shader **can** be used with a fragment shading rate
  attachment.

- `VK_SHADER_CREATE_FRAGMENT_DENSITY_MAP_ATTACHMENT_BIT_EXT` specifies
  that a fragment shader **can** be used with a fragment density map
  attachment.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_shader_object](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_shader_object.html),
[VkShaderCreateFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderCreateFlagsEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkShaderCreateFlagBitsEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
