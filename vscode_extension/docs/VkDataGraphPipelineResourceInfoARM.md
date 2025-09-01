# VkDataGraphPipelineResourceInfoARM(3) Manual Page

## Name

VkDataGraphPipelineResourceInfoARM - Structure specifying parameters of a data graph pipeline resource



## [](#_c_specification)C Specification

The `VkDataGraphPipelineResourceInfoARM` structure is defined as:

```c++
// Provided by VK_ARM_data_graph
typedef struct VkDataGraphPipelineResourceInfoARM {
    VkStructureType    sType;
    const void*        pNext;
    uint32_t           descriptorSet;
    uint32_t           binding;
    uint32_t           arrayElement;
} VkDataGraphPipelineResourceInfoARM;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `descriptorSet` is the descriptor set number of the resource being described.
- `binding` is the binding number of the resource being described.
- `arrayElement` is the element in the resource array if `descriptorSet` and `binding` identifies an array of resources or `0` otherwise.

## [](#_description)Description

Valid Usage

- [](#VUID-VkDataGraphPipelineResourceInfoARM-arrayElement-09779)VUID-VkDataGraphPipelineResourceInfoARM-arrayElement-09779  
  `arrayElement` **must** be `0`
- [](#VUID-VkDataGraphPipelineResourceInfoARM-descriptorSet-09851)VUID-VkDataGraphPipelineResourceInfoARM-descriptorSet-09851  
  If `descriptorSet` and `binding` identify a tensor resource or an array of tensor resources, then a [VkTensorDescriptionARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorDescriptionARM.html) structure whose `usage` contains `VK_TENSOR_USAGE_DATA_GRAPH_BIT_ARM` **must** be included in the `pNext` chain

Valid Usage (Implicit)

- [](#VUID-VkDataGraphPipelineResourceInfoARM-sType-sType)VUID-VkDataGraphPipelineResourceInfoARM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DATA_GRAPH_PIPELINE_RESOURCE_INFO_ARM`
- [](#VUID-VkDataGraphPipelineResourceInfoARM-pNext-pNext)VUID-VkDataGraphPipelineResourceInfoARM-pNext-pNext  
  `pNext` **must** be `NULL` or a pointer to a valid instance of [VkTensorDescriptionARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorDescriptionARM.html)
- [](#VUID-VkDataGraphPipelineResourceInfoARM-sType-unique)VUID-VkDataGraphPipelineResourceInfoARM-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique

## [](#_see_also)See Also

[VK\_ARM\_data\_graph](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_data_graph.html), [VkDataGraphPipelineCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineCreateInfoARM.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDataGraphPipelineResourceInfoARM).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0