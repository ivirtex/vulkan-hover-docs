# VkDataGraphPipelineConstantTensorSemiStructuredSparsityInfoARM(3) Manual Page

## Name

VkDataGraphPipelineConstantTensorSemiStructuredSparsityInfoARM - Structure specifying semi-structured sparsity parameters of a tensor data graph pipeline constant



## [](#_c_specification)C Specification

The `VkDataGraphPipelineConstantTensorSemiStructuredSparsityInfoARM` structure is defined as:

```c++
// Provided by VK_ARM_data_graph with VK_ARM_tensors
typedef struct VkDataGraphPipelineConstantTensorSemiStructuredSparsityInfoARM {
    VkStructureType    sType;
    const void*        pNext;
    uint32_t           dimension;
    uint32_t           zeroCount;
    uint32_t           groupSize;
} VkDataGraphPipelineConstantTensorSemiStructuredSparsityInfoARM;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `dimension` is the dimension of the tensor along which its data is sparse.
- `zeroCount` is the number of tensor elements that **must** be zero in every group of `groupSize` elements.
- `groupSize` is the number of tensor elements in a group.

## [](#_description)Description

Note

This extension does not provide applications with a way of knowing which combinations of `dimension`, `zeroCount`, and `groupSize` an implementation **can** take advantage of. Providing sparsity information for a graph constant is always valid and recommended, regardless of the specific combinations an implementation **can** take advantage of. When they **can** not take advantage of the sparsity information, implementations will ignore it and treat the data as dense.

Valid Usage (Implicit)

- [](#VUID-VkDataGraphPipelineConstantTensorSemiStructuredSparsityInfoARM-sType-sType)VUID-VkDataGraphPipelineConstantTensorSemiStructuredSparsityInfoARM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DATA_GRAPH_PIPELINE_CONSTANT_TENSOR_SEMI_STRUCTURED_SPARSITY_INFO_ARM`

## [](#_see_also)See Also

[VK\_ARM\_data\_graph](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_data_graph.html), [VK\_ARM\_tensors](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_tensors.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDataGraphPipelineConstantTensorSemiStructuredSparsityInfoARM).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0