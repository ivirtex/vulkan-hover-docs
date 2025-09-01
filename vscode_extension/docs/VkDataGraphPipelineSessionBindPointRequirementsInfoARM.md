# VkDataGraphPipelineSessionBindPointRequirementsInfoARM(3) Manual Page

## Name

VkDataGraphPipelineSessionBindPointRequirementsInfoARM - Structure specifying info to query the bind point requirements of a data graph pipeline session



## [](#_c_specification)C Specification

```c++
// Provided by VK_ARM_data_graph
typedef struct VkDataGraphPipelineSessionBindPointRequirementsInfoARM {
    VkStructureType                  sType;
    const void*                      pNext;
    VkDataGraphPipelineSessionARM    session;
} VkDataGraphPipelineSessionBindPointRequirementsInfoARM;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `session` is a [VkDataGraphPipelineSessionARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineSessionARM.html) specifying the data graph pipeline session whose bind point requirements are being queried.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkDataGraphPipelineSessionBindPointRequirementsInfoARM-sType-sType)VUID-VkDataGraphPipelineSessionBindPointRequirementsInfoARM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DATA_GRAPH_PIPELINE_SESSION_BIND_POINT_REQUIREMENTS_INFO_ARM`
- [](#VUID-VkDataGraphPipelineSessionBindPointRequirementsInfoARM-pNext-pNext)VUID-VkDataGraphPipelineSessionBindPointRequirementsInfoARM-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkDataGraphPipelineSessionBindPointRequirementsInfoARM-session-parameter)VUID-VkDataGraphPipelineSessionBindPointRequirementsInfoARM-session-parameter  
  `session` **must** be a valid [VkDataGraphPipelineSessionARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineSessionARM.html) handle

## [](#_see_also)See Also

[VK\_ARM\_data\_graph](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_data_graph.html), [VkDataGraphPipelineSessionARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineSessionARM.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetDataGraphPipelineSessionBindPointRequirementsARM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDataGraphPipelineSessionBindPointRequirementsARM.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDataGraphPipelineSessionBindPointRequirementsInfoARM).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0