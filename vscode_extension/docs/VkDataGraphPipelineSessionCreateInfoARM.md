# VkDataGraphPipelineSessionCreateInfoARM(3) Manual Page

## Name

VkDataGraphPipelineSessionCreateInfoARM - Structure specifying parameters of a newly created data graph pipeline session



## [](#_c_specification)C Specification

The `VkDataGraphPipelineSessionCreateInfoARM` structure is defined as:

```c++
// Provided by VK_ARM_data_graph
typedef struct VkDataGraphPipelineSessionCreateInfoARM {
    VkStructureType                             sType;
    const void*                                 pNext;
    VkDataGraphPipelineSessionCreateFlagsARM    flags;
    VkPipeline                                  dataGraphPipeline;
} VkDataGraphPipelineSessionCreateInfoARM;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is a bitmask of [VkDataGraphPipelineSessionCreateFlagBitsARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineSessionCreateFlagBitsARM.html) describing additional parameters of the session.
- `dataGraphPipeline` is the [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) handle of the data graph pipeline for which a session is being created.

## [](#_description)Description

Valid Usage

- [](#VUID-VkDataGraphPipelineSessionCreateInfoARM-dataGraphPipeline-09781)VUID-VkDataGraphPipelineSessionCreateInfoARM-dataGraphPipeline-09781  
  `dataGraphPipeline` **must** have been obtained via a call to [vkCreateDataGraphPipelinesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateDataGraphPipelinesARM.html)
- [](#VUID-VkDataGraphPipelineSessionCreateInfoARM-protectedMemory-09782)VUID-VkDataGraphPipelineSessionCreateInfoARM-protectedMemory-09782  
  If the [`protectedMemory`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-protectedMemory) feature is not enabled, `flags` **must** not contain `VK_DATA_GRAPH_PIPELINE_SESSION_CREATE_PROTECTED_BIT_ARM`

Valid Usage (Implicit)

- [](#VUID-VkDataGraphPipelineSessionCreateInfoARM-sType-sType)VUID-VkDataGraphPipelineSessionCreateInfoARM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DATA_GRAPH_PIPELINE_SESSION_CREATE_INFO_ARM`
- [](#VUID-VkDataGraphPipelineSessionCreateInfoARM-pNext-pNext)VUID-VkDataGraphPipelineSessionCreateInfoARM-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkDataGraphPipelineSessionCreateInfoARM-flags-parameter)VUID-VkDataGraphPipelineSessionCreateInfoARM-flags-parameter  
  `flags` **must** be a valid combination of [VkDataGraphPipelineSessionCreateFlagBitsARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineSessionCreateFlagBitsARM.html) values
- [](#VUID-VkDataGraphPipelineSessionCreateInfoARM-dataGraphPipeline-parameter)VUID-VkDataGraphPipelineSessionCreateInfoARM-dataGraphPipeline-parameter  
  `dataGraphPipeline` **must** be a valid [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) handle

## [](#_see_also)See Also

[VK\_ARM\_data\_graph](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_data_graph.html), [VkDataGraphPipelineSessionCreateFlagsARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineSessionCreateFlagsARM.html), [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCreateDataGraphPipelineSessionARM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateDataGraphPipelineSessionARM.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDataGraphPipelineSessionCreateInfoARM).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0