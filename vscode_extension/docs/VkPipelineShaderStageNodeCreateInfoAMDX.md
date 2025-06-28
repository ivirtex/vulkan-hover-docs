# VkPipelineShaderStageNodeCreateInfoAMDX(3) Manual Page

## Name

VkPipelineShaderStageNodeCreateInfoAMDX - Structure specifying the shader name and index with an execution graph



## [](#_c_specification)C Specification

The `VkPipelineShaderStageNodeCreateInfoAMDX` structure is defined as:

```c++
// Provided by VK_AMDX_shader_enqueue
typedef struct VkPipelineShaderStageNodeCreateInfoAMDX {
      VkStructureType    sType;
    const void*          pNext;
    const char*          pName;
    uint32_t             index;
} VkPipelineShaderStageNodeCreateInfoAMDX;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `pName` is the shader name to use when creating a node in an execution graph. If `pName` is `NULL`, the name of the entry point specified in SPIR-V is used as the shader name.
- `index` is the shader index to use when creating a node in an execution graph. If `index` is `VK_SHADER_INDEX_UNUSED_AMDX` then the original index is used, either as specified by the `ShaderIndexAMDX` execution mode, or `0` if that too is not specified.

## [](#_description)Description

When included in the `pNext` chain of a [VkPipelineShaderStageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineShaderStageCreateInfo.html) structure, this structure specifies the shader name and shader index of a node when creating an execution graph pipeline. If this structure is omitted, the shader name is set to the name of the entry point in SPIR-V and the shader index is set to `0`.

When dispatching a node from another shader, the name is fixed at pipeline creation, but the index **can** be set dynamically. By associating multiple shaders with the same name but different indexes, applications can dynamically select different nodes to execute. Applications **must** ensure each node has a unique name and index.

Note

Shaders with the same name **must** be of the same type - e.g. a compute and graphics shader, or even two compute shaders where one is coalescing and the other is not, cannot share the same name.

Valid Usage (Implicit)

- [](#VUID-VkPipelineShaderStageNodeCreateInfoAMDX-sType-sType)VUID-VkPipelineShaderStageNodeCreateInfoAMDX-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PIPELINE_SHADER_STAGE_NODE_CREATE_INFO_AMDX`
- [](#VUID-VkPipelineShaderStageNodeCreateInfoAMDX-pName-parameter)VUID-VkPipelineShaderStageNodeCreateInfoAMDX-pName-parameter  
  If `pName` is not `NULL`, `pName` **must** be a null-terminated UTF-8 string

## [](#_see_also)See Also

[VK\_AMDX\_shader\_enqueue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_AMDX_shader_enqueue.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetExecutionGraphPipelineNodeIndexAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetExecutionGraphPipelineNodeIndexAMDX.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPipelineShaderStageNodeCreateInfoAMDX)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0