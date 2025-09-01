# VkDataGraphPipelineDispatchInfoARM(3) Manual Page

## Name

VkDataGraphPipelineDispatchInfoARM - Structure specifying parameters of a data graph pipeline dispatch



## [](#_c_specification)C Specification

The [VkDataGraphPipelineDispatchInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineDispatchInfoARM.html) structure is defined as:

```c++
// Provided by VK_ARM_data_graph
typedef struct VkDataGraphPipelineDispatchInfoARM {
    VkStructureType                        sType;
    void*                                  pNext;
    VkDataGraphPipelineDispatchFlagsARM    flags;
} VkDataGraphPipelineDispatchInfoARM;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is a bitmask of [VkDataGraphPipelineDispatchFlagBitsARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineDispatchFlagBitsARM.html) describing additional parameters of the dispatch.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkDataGraphPipelineDispatchInfoARM-sType-sType)VUID-VkDataGraphPipelineDispatchInfoARM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DATA_GRAPH_PIPELINE_DISPATCH_INFO_ARM`
- [](#VUID-VkDataGraphPipelineDispatchInfoARM-pNext-pNext)VUID-VkDataGraphPipelineDispatchInfoARM-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkDataGraphPipelineDispatchInfoARM-flags-zerobitmask)VUID-VkDataGraphPipelineDispatchInfoARM-flags-zerobitmask  
  `flags` **must** be `0`

## [](#_see_also)See Also

[VK\_ARM\_data\_graph](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_data_graph.html), [VkDataGraphPipelineDispatchFlagsARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineDispatchFlagsARM.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCmdDispatchDataGraphARM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDispatchDataGraphARM.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDataGraphPipelineDispatchInfoARM).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0