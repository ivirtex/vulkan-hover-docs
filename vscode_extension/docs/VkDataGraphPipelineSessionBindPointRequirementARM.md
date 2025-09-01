# VkDataGraphPipelineSessionBindPointRequirementARM(3) Manual Page

## Name

VkDataGraphPipelineSessionBindPointRequirementARM - Structure specifying the requirements of a bind point of a data graph pipeline session



## [](#_c_specification)C Specification

```c++
// Provided by VK_ARM_data_graph
typedef struct VkDataGraphPipelineSessionBindPointRequirementARM {
    VkStructureType                               sType;
    const void*                                   pNext;
    VkDataGraphPipelineSessionBindPointARM        bindPoint;
    VkDataGraphPipelineSessionBindPointTypeARM    bindPointType;
    uint32_t                                      numObjects;
} VkDataGraphPipelineSessionBindPointRequirementARM;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `bindPoint` is a [VkDataGraphPipelineSessionBindPointARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineSessionBindPointARM.html) specifying the data graph pipeline session bind point being required.
- `bindPointType` is a [VkDataGraphPipelineSessionBindPointTypeARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineSessionBindPointTypeARM.html) specifying the type of object required for `bindPoint`.
- `numObjects` is the number of objects required for `bindPoint`.

## [](#_description)Description

Implementations **must** always return 1 for `numObjects` if `bindPoint` is one of the following bind points:

- `VK_DATA_GRAPH_PIPELINE_SESSION_BIND_POINT_TRANSIENT_ARM`

Valid Usage (Implicit)

- [](#VUID-VkDataGraphPipelineSessionBindPointRequirementARM-sType-sType)VUID-VkDataGraphPipelineSessionBindPointRequirementARM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DATA_GRAPH_PIPELINE_SESSION_BIND_POINT_REQUIREMENT_ARM`
- [](#VUID-VkDataGraphPipelineSessionBindPointRequirementARM-pNext-pNext)VUID-VkDataGraphPipelineSessionBindPointRequirementARM-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkDataGraphPipelineSessionBindPointRequirementARM-bindPoint-parameter)VUID-VkDataGraphPipelineSessionBindPointRequirementARM-bindPoint-parameter  
  `bindPoint` **must** be a valid [VkDataGraphPipelineSessionBindPointARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineSessionBindPointARM.html) value
- [](#VUID-VkDataGraphPipelineSessionBindPointRequirementARM-bindPointType-parameter)VUID-VkDataGraphPipelineSessionBindPointRequirementARM-bindPointType-parameter  
  `bindPointType` **must** be a valid [VkDataGraphPipelineSessionBindPointTypeARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineSessionBindPointTypeARM.html) value

## [](#_see_also)See Also

[VK\_ARM\_data\_graph](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_data_graph.html), [VkDataGraphPipelineSessionBindPointARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineSessionBindPointARM.html), [VkDataGraphPipelineSessionBindPointTypeARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineSessionBindPointTypeARM.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetDataGraphPipelineSessionBindPointRequirementsARM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDataGraphPipelineSessionBindPointRequirementsARM.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDataGraphPipelineSessionBindPointRequirementARM).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0