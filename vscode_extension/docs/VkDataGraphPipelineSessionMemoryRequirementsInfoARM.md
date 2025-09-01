# VkDataGraphPipelineSessionMemoryRequirementsInfoARM(3) Manual Page

## Name

VkDataGraphPipelineSessionMemoryRequirementsInfoARM - Structure specifying parameters to query the memory requirements of a data graph pipeline session



## [](#_c_specification)C Specification

The `VkDataGraphPipelineSessionMemoryRequirementsInfoARM` structure is defined as:

```c++
// Provided by VK_ARM_data_graph
typedef struct VkDataGraphPipelineSessionMemoryRequirementsInfoARM {
    VkStructureType                           sType;
    const void*                               pNext;
    VkDataGraphPipelineSessionARM             session;
    VkDataGraphPipelineSessionBindPointARM    bindPoint;
    uint32_t                                  objectIndex;
} VkDataGraphPipelineSessionMemoryRequirementsInfoARM;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `session` is the data graph pipeline session to query.
- `bindPoint` is the bind point of a data graph pipeline session for which memory requirements are being queried.
- `objectIndex` is the index of the object whose memory requirements are being queried.

## [](#_description)Description

Valid Usage

- [](#VUID-VkDataGraphPipelineSessionMemoryRequirementsInfoARM-objectIndex-09855)VUID-VkDataGraphPipelineSessionMemoryRequirementsInfoARM-objectIndex-09855  
  `objectIndex` **must** be less than the number of objects returned by [vkGetDataGraphPipelineSessionBindPointRequirementsARM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDataGraphPipelineSessionBindPointRequirementsARM.html) via [VkDataGraphPipelineSessionBindPointRequirementARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineSessionBindPointRequirementARM.html)::`numObjects` with [VkDataGraphPipelineSessionMemoryRequirementsInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineSessionMemoryRequirementsInfoARM.html)::`bindPoint` equal to `bindPoint`

Valid Usage (Implicit)

- [](#VUID-VkDataGraphPipelineSessionMemoryRequirementsInfoARM-sType-sType)VUID-VkDataGraphPipelineSessionMemoryRequirementsInfoARM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DATA_GRAPH_PIPELINE_SESSION_MEMORY_REQUIREMENTS_INFO_ARM`
- [](#VUID-VkDataGraphPipelineSessionMemoryRequirementsInfoARM-pNext-pNext)VUID-VkDataGraphPipelineSessionMemoryRequirementsInfoARM-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkDataGraphPipelineSessionMemoryRequirementsInfoARM-session-parameter)VUID-VkDataGraphPipelineSessionMemoryRequirementsInfoARM-session-parameter  
  `session` **must** be a valid [VkDataGraphPipelineSessionARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineSessionARM.html) handle
- [](#VUID-VkDataGraphPipelineSessionMemoryRequirementsInfoARM-bindPoint-parameter)VUID-VkDataGraphPipelineSessionMemoryRequirementsInfoARM-bindPoint-parameter  
  `bindPoint` **must** be a valid [VkDataGraphPipelineSessionBindPointARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineSessionBindPointARM.html) value

## [](#_see_also)See Also

[VK\_ARM\_data\_graph](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_data_graph.html), [VkDataGraphPipelineSessionARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineSessionARM.html), [VkDataGraphPipelineSessionBindPointARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineSessionBindPointARM.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetDataGraphPipelineSessionMemoryRequirementsARM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDataGraphPipelineSessionMemoryRequirementsARM.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDataGraphPipelineSessionMemoryRequirementsInfoARM).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0