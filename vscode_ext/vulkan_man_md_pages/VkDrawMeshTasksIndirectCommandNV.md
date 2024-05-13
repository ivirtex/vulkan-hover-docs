# VkDrawMeshTasksIndirectCommandNV(3) Manual Page

## Name

VkDrawMeshTasksIndirectCommandNV - Structure specifying a mesh tasks
draw indirect command



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkDrawMeshTasksIndirectCommandNV` structure is defined as:

``` c
// Provided by VK_NV_mesh_shader
typedef struct VkDrawMeshTasksIndirectCommandNV {
    uint32_t    taskCount;
    uint32_t    firstTask;
} VkDrawMeshTasksIndirectCommandNV;
```

## <a href="#_members" class="anchor"></a>Members

- `taskCount` is the number of local workgroups to dispatch in the X
  dimension. Y and Z dimension are implicitly set to one.

- `firstTask` is the X component of the first workgroup ID.

## <a href="#_description" class="anchor"></a>Description

The members of `VkDrawMeshTasksIndirectCommandNV` have the same meaning
as the similarly named parameters of
[vkCmdDrawMeshTasksNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDrawMeshTasksNV.html).

Valid Usage

- <a href="#VUID-VkDrawMeshTasksIndirectCommandNV-taskCount-02175"
  id="VUID-VkDrawMeshTasksIndirectCommandNV-taskCount-02175"></a>
  VUID-VkDrawMeshTasksIndirectCommandNV-taskCount-02175  
  `taskCount` **must** be less than or equal to
  `VkPhysicalDeviceMeshShaderPropertiesNV`::`maxDrawMeshTasksCount`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_mesh_shader](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_mesh_shader.html),
[vkCmdDrawMeshTasksIndirectNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDrawMeshTasksIndirectNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDrawMeshTasksIndirectCommandNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
