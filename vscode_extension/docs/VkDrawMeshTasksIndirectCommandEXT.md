# VkDrawMeshTasksIndirectCommandEXT(3) Manual Page

## Name

VkDrawMeshTasksIndirectCommandEXT - Structure specifying a mesh tasks draw indirect command



## [](#_c_specification)C Specification

The `VkDrawMeshTasksIndirectCommandEXT` structure is defined as:

```c++
// Provided by VK_EXT_mesh_shader
typedef struct VkDrawMeshTasksIndirectCommandEXT {
    uint32_t    groupCountX;
    uint32_t    groupCountY;
    uint32_t    groupCountZ;
} VkDrawMeshTasksIndirectCommandEXT;
```

## [](#_members)Members

- `groupCountX` is the number of local workgroups to dispatch in the X dimension.
- `groupCountY` is the number of local workgroups to dispatch in the Y dimension.
- `groupCountZ` is the number of local workgroups to dispatch in the Z dimension.

## [](#_description)Description

The members of `VkDrawMeshTasksIndirectCommandEXT` have the same meaning as the similarly named parameters of [vkCmdDrawMeshTasksEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDrawMeshTasksEXT.html).

Valid Usage

- [](#VUID-VkDrawMeshTasksIndirectCommandEXT-TaskEXT-07322)VUID-VkDrawMeshTasksIndirectCommandEXT-TaskEXT-07322  
  If the current pipeline bound to `VK_PIPELINE_BIND_POINT_GRAPHICS` contains a shader using the `TaskEXT` `Execution` `Model`, `groupCountX` **must** be less than or equal to `VkPhysicalDeviceMeshShaderPropertiesEXT`::`maxTaskWorkGroupCount`\[0]
- [](#VUID-VkDrawMeshTasksIndirectCommandEXT-TaskEXT-07323)VUID-VkDrawMeshTasksIndirectCommandEXT-TaskEXT-07323  
  If the current pipeline bound to `VK_PIPELINE_BIND_POINT_GRAPHICS` contains a shader using the `TaskEXT` `Execution` `Model`, `groupCountY` **must** be less than or equal to `VkPhysicalDeviceMeshShaderPropertiesEXT`::`maxTaskWorkGroupCount`\[1]
- [](#VUID-VkDrawMeshTasksIndirectCommandEXT-TaskEXT-07324)VUID-VkDrawMeshTasksIndirectCommandEXT-TaskEXT-07324  
  If the current pipeline bound to `VK_PIPELINE_BIND_POINT_GRAPHICS` contains a shader using the `TaskEXT` `Execution` `Model`, `groupCountZ` **must** be less than or equal to `VkPhysicalDeviceMeshShaderPropertiesEXT`::`maxTaskWorkGroupCount`\[2]
- [](#VUID-VkDrawMeshTasksIndirectCommandEXT-TaskEXT-07325)VUID-VkDrawMeshTasksIndirectCommandEXT-TaskEXT-07325  
  If the current pipeline bound to `VK_PIPELINE_BIND_POINT_GRAPHICS` contains a shader using the `TaskEXT` `Execution` `Model`, The product of `groupCountX`, `groupCountY` and `groupCountZ` **must** be less than or equal to `VkPhysicalDeviceMeshShaderPropertiesEXT`::`maxTaskWorkGroupTotalCount`
- [](#VUID-VkDrawMeshTasksIndirectCommandEXT-TaskEXT-07326)VUID-VkDrawMeshTasksIndirectCommandEXT-TaskEXT-07326  
  If the current pipeline bound to `VK_PIPELINE_BIND_POINT_GRAPHICS` does not contain a shader using the `TaskEXT` `Execution` `Model`, `groupCountX` **must** be less than or equal to `VkPhysicalDeviceMeshShaderPropertiesEXT`::`maxMeshWorkGroupCount`\[0]
- [](#VUID-VkDrawMeshTasksIndirectCommandEXT-TaskEXT-07327)VUID-VkDrawMeshTasksIndirectCommandEXT-TaskEXT-07327  
  If the current pipeline bound to `VK_PIPELINE_BIND_POINT_GRAPHICS` does not contain a shader using the `TaskEXT` `Execution` `Model`, `groupCountY` **must** be less than or equal to `VkPhysicalDeviceMeshShaderPropertiesEXT`::`maxMeshWorkGroupCount`\[1]
- [](#VUID-VkDrawMeshTasksIndirectCommandEXT-TaskEXT-07328)VUID-VkDrawMeshTasksIndirectCommandEXT-TaskEXT-07328  
  If the current pipeline bound to `VK_PIPELINE_BIND_POINT_GRAPHICS` does not contain a shader using the `TaskEXT` `Execution` `Model`, `groupCountZ` **must** be less than or equal to `VkPhysicalDeviceMeshShaderPropertiesEXT`::`maxMeshWorkGroupCount`\[2]
- [](#VUID-VkDrawMeshTasksIndirectCommandEXT-TaskEXT-07329)VUID-VkDrawMeshTasksIndirectCommandEXT-TaskEXT-07329  
  If the current pipeline bound to `VK_PIPELINE_BIND_POINT_GRAPHICS` does not contain a shader using the `TaskEXT` `Execution` `Model`, The product of `groupCountX`, `groupCountY` and `groupCountZ` **must** be less than or equal to `VkPhysicalDeviceMeshShaderPropertiesEXT`::`maxMeshWorkGroupTotalCount`

## [](#_see_also)See Also

[VK\_EXT\_mesh\_shader](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_mesh_shader.html), [vkCmdDrawMeshTasksIndirectEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDrawMeshTasksIndirectEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDrawMeshTasksIndirectCommandEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0