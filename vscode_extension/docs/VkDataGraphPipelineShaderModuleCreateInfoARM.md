# VkDataGraphPipelineShaderModuleCreateInfoARM(3) Manual Page

## Name

VkDataGraphPipelineShaderModuleCreateInfoARM - Structure specifying shader module parameters of a newly created data graph pipeline



## [](#_c_specification)C Specification

The `VkDataGraphPipelineShaderModuleCreateInfoARM` structure is defined as:

```c++
// Provided by VK_ARM_data_graph
typedef struct VkDataGraphPipelineShaderModuleCreateInfoARM {
    VkStructureType                          sType;
    const void*                              pNext;
    VkShaderModule                           module;
    const char*                              pName;
    const VkSpecializationInfo*              pSpecializationInfo;
    uint32_t                                 constantCount;
    const VkDataGraphPipelineConstantARM*    pConstants;
} VkDataGraphPipelineShaderModuleCreateInfoARM;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `module` is optionally a [VkShaderModule](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderModule.html) object containing the description of the graph.
- `pName` is a pointer to a null-terminated UTF-8 string specifying the graph entry point name for this pipeline.
- `pSpecializationInfo` is a pointer to a [VkSpecializationInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSpecializationInfo.html) structure as described in [Specialization Constants](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-specialization-constants), or `NULL`.
- `constantCount` is the length of the `pConstants` array.
- `pConstants` is a pointer to an array of [VkDataGraphPipelineConstantARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineConstantARM.html) structures.

## [](#_description)Description

If `module` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), the pipeline’s graph is defined by `module`. If `module` is [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), the pipeline’s graph is defined by the chained [VkShaderModuleCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderModuleCreateInfo.html).

Valid Usage

- [](#VUID-VkDataGraphPipelineShaderModuleCreateInfoARM-id-09774)VUID-VkDataGraphPipelineShaderModuleCreateInfoARM-id-09774  
  The `id` member of all structures in `pConstants` **must** be a valid `GraphConstantID` used by a `OpGraphConstantARM` instruction in `module`
- [](#VUID-VkDataGraphPipelineShaderModuleCreateInfoARM-dataGraphSpecializationConstants-09849)VUID-VkDataGraphPipelineShaderModuleCreateInfoARM-dataGraphSpecializationConstants-09849  
  If the [`dataGraphSpecializationConstants`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-dataGraphSpecializationConstants) feature is not enabled then `pSpecializationInfo` **must** be `NULL` and `module` **must** not contain any `OpSpec*` instructions
- [](#VUID-VkDataGraphPipelineShaderModuleCreateInfoARM-pName-09872)VUID-VkDataGraphPipelineShaderModuleCreateInfoARM-pName-09872  
  `pName` **must** be the name of an `OpGraphEntryPointARM` in `module`
- [](#VUID-VkDataGraphPipelineShaderModuleCreateInfoARM-pNext-09873)VUID-VkDataGraphPipelineShaderModuleCreateInfoARM-pNext-09873  
  If the `pNext` chain includes a [VkShaderModuleCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderModuleCreateInfo.html) structure, then `module` **must** be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html)
- [](#VUID-VkDataGraphPipelineShaderModuleCreateInfoARM-pNext-09874)VUID-VkDataGraphPipelineShaderModuleCreateInfoARM-pNext-09874  
  If the `pNext` chain does not include a [VkShaderModuleCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderModuleCreateInfo.html) structure, then `module` **must** be a valid [VkShaderModule](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderModule.html)

Valid Usage (Implicit)

- [](#VUID-VkDataGraphPipelineShaderModuleCreateInfoARM-sType-sType)VUID-VkDataGraphPipelineShaderModuleCreateInfoARM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DATA_GRAPH_PIPELINE_SHADER_MODULE_CREATE_INFO_ARM`
- [](#VUID-VkDataGraphPipelineShaderModuleCreateInfoARM-module-parameter)VUID-VkDataGraphPipelineShaderModuleCreateInfoARM-module-parameter  
  If `module` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `module` **must** be a valid [VkShaderModule](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderModule.html) handle
- [](#VUID-VkDataGraphPipelineShaderModuleCreateInfoARM-pName-parameter)VUID-VkDataGraphPipelineShaderModuleCreateInfoARM-pName-parameter  
  `pName` **must** be a null-terminated UTF-8 string
- [](#VUID-VkDataGraphPipelineShaderModuleCreateInfoARM-pSpecializationInfo-parameter)VUID-VkDataGraphPipelineShaderModuleCreateInfoARM-pSpecializationInfo-parameter  
  If `pSpecializationInfo` is not `NULL`, `pSpecializationInfo` **must** be a valid pointer to a valid [VkSpecializationInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSpecializationInfo.html) structure
- [](#VUID-VkDataGraphPipelineShaderModuleCreateInfoARM-pConstants-parameter)VUID-VkDataGraphPipelineShaderModuleCreateInfoARM-pConstants-parameter  
  If `constantCount` is not `0`, and `pConstants` is not `NULL`, `pConstants` **must** be a valid pointer to an array of `constantCount` valid [VkDataGraphPipelineConstantARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineConstantARM.html) structures

## [](#_see_also)See Also

[VK\_ARM\_data\_graph](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_data_graph.html), [VkDataGraphPipelineConstantARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineConstantARM.html), [VkShaderModule](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderModule.html), [VkSpecializationInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSpecializationInfo.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDataGraphPipelineShaderModuleCreateInfoARM)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0