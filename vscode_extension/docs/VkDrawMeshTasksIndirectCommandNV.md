# VkDrawMeshTasksIndirectCommandNV(3) Manual Page

## Name

VkDrawMeshTasksIndirectCommandNV - Structure specifying a mesh tasks draw indirect command



## [](#_c_specification)C Specification

The `VkDrawMeshTasksIndirectCommandNV` structure is defined as:

```c++
// Provided by VK_NV_mesh_shader
typedef struct VkDrawMeshTasksIndirectCommandNV {
    uint32_t    taskCount;
    uint32_t    firstTask;
} VkDrawMeshTasksIndirectCommandNV;
```

## [](#_members)Members

- `taskCount` is the number of local workgroups to dispatch in the X dimension. Y and Z dimension are implicitly set to one.
- `firstTask` is the X component of the first workgroup ID.

## [](#_description)Description

The members of `VkDrawMeshTasksIndirectCommandNV` have the same meaning as the similarly named parameters of [vkCmdDrawMeshTasksNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDrawMeshTasksNV.html).

Valid Usage

- [](#VUID-VkDrawMeshTasksIndirectCommandNV-taskCount-02175)VUID-VkDrawMeshTasksIndirectCommandNV-taskCount-02175  
  `taskCount` **must** be less than or equal to `VkPhysicalDeviceMeshShaderPropertiesNV`::`maxDrawMeshTasksCount`

## [](#_see_also)See Also

[VK\_NV\_mesh\_shader](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_mesh_shader.html), [vkCmdDrawMeshTasksIndirectNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDrawMeshTasksIndirectNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDrawMeshTasksIndirectCommandNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0