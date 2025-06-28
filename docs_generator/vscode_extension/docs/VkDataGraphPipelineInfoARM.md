# VkDataGraphPipelineInfoARM(3) Manual Page

## Name

VkDataGraphPipelineInfoARM - Structure describing a data graph pipeline



## [](#_c_specification)C Specification

The `VkDataGraphPipelineInfoARM` structure is defined as:

```c++
// Provided by VK_ARM_data_graph
typedef struct VkDataGraphPipelineInfoARM {
    VkStructureType    sType;
    const void*        pNext;
    VkPipeline         dataGraphPipeline;
} VkDataGraphPipelineInfoARM;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `dataGraphPipeline` is a [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) handle.

## [](#_description)Description

Valid Usage

- [](#VUID-VkDataGraphPipelineInfoARM-dataGraphPipeline-09803)VUID-VkDataGraphPipelineInfoARM-dataGraphPipeline-09803  
  `dataGraphPipeline` **must** have been created with [vkCreateDataGraphPipelinesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateDataGraphPipelinesARM.html)

Valid Usage (Implicit)

- [](#VUID-VkDataGraphPipelineInfoARM-sType-sType)VUID-VkDataGraphPipelineInfoARM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DATA_GRAPH_PIPELINE_INFO_ARM`
- [](#VUID-VkDataGraphPipelineInfoARM-pNext-pNext)VUID-VkDataGraphPipelineInfoARM-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkDataGraphPipelineInfoARM-dataGraphPipeline-parameter)VUID-VkDataGraphPipelineInfoARM-dataGraphPipeline-parameter  
  `dataGraphPipeline` **must** be a valid [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) handle

## [](#_see_also)See Also

[VK\_ARM\_data\_graph](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_data_graph.html), [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetDataGraphPipelineAvailablePropertiesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDataGraphPipelineAvailablePropertiesARM.html), [vkGetDataGraphPipelinePropertiesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDataGraphPipelinePropertiesARM.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDataGraphPipelineInfoARM)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0