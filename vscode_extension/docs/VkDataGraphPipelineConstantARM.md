# VkDataGraphPipelineConstantARM(3) Manual Page

## Name

VkDataGraphPipelineConstantARM - Structure specifying parameters of a data graph pipeline constant



## [](#_c_specification)C Specification

The `VkDataGraphPipelineConstantARM` structure is defined as:

```c++
// Provided by VK_ARM_data_graph
typedef struct VkDataGraphPipelineConstantARM {
    VkStructureType    sType;
    const void*        pNext;
    uint32_t           id;
    const void*        pConstantData;
} VkDataGraphPipelineConstantARM;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is a pointer to a structure extending this structure.
- `id` is the unique identifier of the graph constant this structure describes.
- `pConstantData` is a pointer to the data for this graph constant.

## [](#_description)Description

The size and layout of the data pointed to by `pConstantData` is specified by a specific structure in the `pNext` chain for each type of graph constant.

For graph constants of tensor type, the layout of the data is specified by a [VkTensorDescriptionARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorDescriptionARM.html) structure. The data **must** be laid out according to the following members of this structure:

- [VkTensorDescriptionARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorDescriptionARM.html)::`tiling`
- [VkTensorDescriptionARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorDescriptionARM.html)::`format`
- [VkTensorDescriptionARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorDescriptionARM.html)::`dimensionCount`
- [VkTensorDescriptionARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorDescriptionARM.html)::`pDimensions`
- [VkTensorDescriptionARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorDescriptionARM.html)::`pStrides`

The presence of a [VkDataGraphPipelineConstantTensorSemiStructuredSparsityInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineConstantTensorSemiStructuredSparsityInfoARM.html) structure in the `pNext` chain has no impact on the expected layout of the data pointed to by `pConstantData`.

Valid Usage

- [](#VUID-VkDataGraphPipelineConstantARM-pNext-09775)VUID-VkDataGraphPipelineConstantARM-pNext-09775  
  If the `pNext` chain of this structure includes one or more [VkDataGraphPipelineConstantTensorSemiStructuredSparsityInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineConstantTensorSemiStructuredSparsityInfoARM.html) structures then it **must** also include a [VkTensorDescriptionARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorDescriptionARM.html) structure
- [](#VUID-VkDataGraphPipelineConstantARM-pNext-09776)VUID-VkDataGraphPipelineConstantARM-pNext-09776  
  If the `pNext` chain of this structure includes one or more [VkDataGraphPipelineConstantTensorSemiStructuredSparsityInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineConstantTensorSemiStructuredSparsityInfoARM.html) structures then, for each structure, [VkDataGraphPipelineConstantTensorSemiStructuredSparsityInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineConstantTensorSemiStructuredSparsityInfoARM.html)::`dimension` **must** be less than [VkTensorDescriptionARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorDescriptionARM.html)::`dimensionCount`
- [](#VUID-VkDataGraphPipelineConstantARM-pNext-09777)VUID-VkDataGraphPipelineConstantARM-pNext-09777  
  If the `pNext` chain of this structure includes a [VkDataGraphPipelineConstantTensorSemiStructuredSparsityInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineConstantTensorSemiStructuredSparsityInfoARM.html) structure then, for each structure, [VkTensorDescriptionARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorDescriptionARM.html)::`pDimensions`\[[VkDataGraphPipelineConstantTensorSemiStructuredSparsityInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineConstantTensorSemiStructuredSparsityInfoARM.html)::`dimension`] **must** be a multiple of [VkDataGraphPipelineConstantTensorSemiStructuredSparsityInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineConstantTensorSemiStructuredSparsityInfoARM.html)::`groupSize`
- [](#VUID-VkDataGraphPipelineConstantARM-pNext-09870)VUID-VkDataGraphPipelineConstantARM-pNext-09870  
  If the `pNext` chain of this structure includes multiple [VkDataGraphPipelineConstantTensorSemiStructuredSparsityInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineConstantTensorSemiStructuredSparsityInfoARM.html) structures then no two structures **may** have their [VkDataGraphPipelineConstantTensorSemiStructuredSparsityInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineConstantTensorSemiStructuredSparsityInfoARM.html)::`dimension` member set to the same value
- [](#VUID-VkDataGraphPipelineConstantARM-id-09850)VUID-VkDataGraphPipelineConstantARM-id-09850  
  If `id` corresponds to a graph constant of tensor type, then a [VkTensorDescriptionARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorDescriptionARM.html) structure whose `usage` member contains `VK_TENSOR_USAGE_DATA_GRAPH_BIT_ARM` **must** be included in the `pNext` chain
- [](#VUID-VkDataGraphPipelineConstantARM-pNext-09917)VUID-VkDataGraphPipelineConstantARM-pNext-09917  
  If the `pNext` chain of this structure includes a [VkTensorDescriptionARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorDescriptionARM.html) structure, then its `tiling` member **must** be `VK_TENSOR_TILING_LINEAR_ARM`

Valid Usage (Implicit)

- [](#VUID-VkDataGraphPipelineConstantARM-sType-sType)VUID-VkDataGraphPipelineConstantARM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DATA_GRAPH_PIPELINE_CONSTANT_ARM`
- [](#VUID-VkDataGraphPipelineConstantARM-pNext-pNext)VUID-VkDataGraphPipelineConstantARM-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the `pNext` chain **must** be either `NULL` or a pointer to a valid instance of [VkDataGraphPipelineConstantTensorSemiStructuredSparsityInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineConstantTensorSemiStructuredSparsityInfoARM.html) or [VkTensorDescriptionARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorDescriptionARM.html)
- [](#VUID-VkDataGraphPipelineConstantARM-sType-unique)VUID-VkDataGraphPipelineConstantARM-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique, with the exception of structures of type [VkDataGraphPipelineConstantTensorSemiStructuredSparsityInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineConstantTensorSemiStructuredSparsityInfoARM.html)
- [](#VUID-VkDataGraphPipelineConstantARM-pConstantData-parameter)VUID-VkDataGraphPipelineConstantARM-pConstantData-parameter  
  `pConstantData` **must** be a pointer value

## [](#_see_also)See Also

[VK\_ARM\_data\_graph](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_data_graph.html), [VkDataGraphPipelineShaderModuleCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineShaderModuleCreateInfoARM.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDataGraphPipelineConstantARM).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0