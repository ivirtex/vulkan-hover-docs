# VkDrawMeshTasksIndirectCommandEXT(3) Manual Page

## Name

VkDrawMeshTasksIndirectCommandEXT - Structure specifying a mesh tasks
draw indirect command



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkDrawMeshTasksIndirectCommandEXT` structure is defined as:

``` c
// Provided by VK_EXT_mesh_shader
typedef struct VkDrawMeshTasksIndirectCommandEXT {
    uint32_t    groupCountX;
    uint32_t    groupCountY;
    uint32_t    groupCountZ;
} VkDrawMeshTasksIndirectCommandEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `groupCountX` is the number of local workgroups to dispatch in the X
  dimension.

- `groupCountY` is the number of local workgroups to dispatch in the Y
  dimension.

- `groupCountZ` is the number of local workgroups to dispatch in the Z
  dimension.

## <a href="#_description" class="anchor"></a>Description

The members of `VkDrawMeshTasksIndirectCommandEXT` have the same meaning
as the similarly named parameters of
[vkCmdDrawMeshTasksEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDrawMeshTasksEXT.html).

Valid Usage

- <a href="#VUID-VkDrawMeshTasksIndirectCommandEXT-TaskEXT-07322"
  id="VUID-VkDrawMeshTasksIndirectCommandEXT-TaskEXT-07322"></a>
  VUID-VkDrawMeshTasksIndirectCommandEXT-TaskEXT-07322  
  If the current pipeline bound to `VK_PIPELINE_BIND_POINT_GRAPHICS`
  contains a shader using the `TaskEXT` `Execution` `Model`,
  `groupCountX` **must** be less than or equal to
  `VkPhysicalDeviceMeshShaderPropertiesEXT`::`maxTaskWorkGroupCount`\[0\]

- <a href="#VUID-VkDrawMeshTasksIndirectCommandEXT-TaskEXT-07323"
  id="VUID-VkDrawMeshTasksIndirectCommandEXT-TaskEXT-07323"></a>
  VUID-VkDrawMeshTasksIndirectCommandEXT-TaskEXT-07323  
  If the current pipeline bound to `VK_PIPELINE_BIND_POINT_GRAPHICS`
  contains a shader using the `TaskEXT` `Execution` `Model`,
  `groupCountY` **must** be less than or equal to
  `VkPhysicalDeviceMeshShaderPropertiesEXT`::`maxTaskWorkGroupCount`\[1\]

- <a href="#VUID-VkDrawMeshTasksIndirectCommandEXT-TaskEXT-07324"
  id="VUID-VkDrawMeshTasksIndirectCommandEXT-TaskEXT-07324"></a>
  VUID-VkDrawMeshTasksIndirectCommandEXT-TaskEXT-07324  
  If the current pipeline bound to `VK_PIPELINE_BIND_POINT_GRAPHICS`
  contains a shader using the `TaskEXT` `Execution` `Model`,
  `groupCountZ` **must** be less than or equal to
  `VkPhysicalDeviceMeshShaderPropertiesEXT`::`maxTaskWorkGroupCount`\[2\]

- <a href="#VUID-VkDrawMeshTasksIndirectCommandEXT-TaskEXT-07325"
  id="VUID-VkDrawMeshTasksIndirectCommandEXT-TaskEXT-07325"></a>
  VUID-VkDrawMeshTasksIndirectCommandEXT-TaskEXT-07325  
  If the current pipeline bound to `VK_PIPELINE_BIND_POINT_GRAPHICS`
  contains a shader using the `TaskEXT` `Execution` `Model`, The product
  of `groupCountX`, `groupCountY` and `groupCountZ` **must** be less
  than or equal to
  `VkPhysicalDeviceMeshShaderPropertiesEXT`::`maxTaskWorkGroupTotalCount`

- <a href="#VUID-VkDrawMeshTasksIndirectCommandEXT-TaskEXT-07326"
  id="VUID-VkDrawMeshTasksIndirectCommandEXT-TaskEXT-07326"></a>
  VUID-VkDrawMeshTasksIndirectCommandEXT-TaskEXT-07326  
  If the current pipeline bound to `VK_PIPELINE_BIND_POINT_GRAPHICS`
  does not contain a shader using the `TaskEXT` `Execution` `Model`,
  `groupCountX` **must** be less than or equal to
  `VkPhysicalDeviceMeshShaderPropertiesEXT`::`maxMeshWorkGroupCount`\[0\]

- <a href="#VUID-VkDrawMeshTasksIndirectCommandEXT-TaskEXT-07327"
  id="VUID-VkDrawMeshTasksIndirectCommandEXT-TaskEXT-07327"></a>
  VUID-VkDrawMeshTasksIndirectCommandEXT-TaskEXT-07327  
  If the current pipeline bound to `VK_PIPELINE_BIND_POINT_GRAPHICS`
  does not contain a shader using the `TaskEXT` `Execution` `Model`,
  `groupCountY` **must** be less than or equal to
  `VkPhysicalDeviceMeshShaderPropertiesEXT`::`maxMeshWorkGroupCount`\[1\]

- <a href="#VUID-VkDrawMeshTasksIndirectCommandEXT-TaskEXT-07328"
  id="VUID-VkDrawMeshTasksIndirectCommandEXT-TaskEXT-07328"></a>
  VUID-VkDrawMeshTasksIndirectCommandEXT-TaskEXT-07328  
  If the current pipeline bound to `VK_PIPELINE_BIND_POINT_GRAPHICS`
  does not contain a shader using the `TaskEXT` `Execution` `Model`,
  `groupCountZ` **must** be less than or equal to
  `VkPhysicalDeviceMeshShaderPropertiesEXT`::`maxMeshWorkGroupCount`\[2\]

- <a href="#VUID-VkDrawMeshTasksIndirectCommandEXT-TaskEXT-07329"
  id="VUID-VkDrawMeshTasksIndirectCommandEXT-TaskEXT-07329"></a>
  VUID-VkDrawMeshTasksIndirectCommandEXT-TaskEXT-07329  
  If the current pipeline bound to `VK_PIPELINE_BIND_POINT_GRAPHICS`
  does not contain a shader using the `TaskEXT` `Execution` `Model`, The
  product of `groupCountX`, `groupCountY` and `groupCountZ` **must** be
  less than or equal to
  `VkPhysicalDeviceMeshShaderPropertiesEXT`::`maxMeshWorkGroupTotalCount`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_mesh_shader](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_mesh_shader.html),
[vkCmdDrawMeshTasksIndirectEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDrawMeshTasksIndirectEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDrawMeshTasksIndirectCommandEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
