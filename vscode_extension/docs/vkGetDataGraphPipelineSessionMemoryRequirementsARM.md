# vkGetDataGraphPipelineSessionMemoryRequirementsARM(3) Manual Page

## Name

vkGetDataGraphPipelineSessionMemoryRequirementsARM - Get the memory requirements of a data graph pipeline session



## [](#_c_specification)C Specification

To determine the memory requirements for a data graph pipeline session, call:

```c++
// Provided by VK_ARM_data_graph
void vkGetDataGraphPipelineSessionMemoryRequirementsARM(
    VkDevice                                    device,
    const VkDataGraphPipelineSessionMemoryRequirementsInfoARM* pInfo,
    VkMemoryRequirements2*                      pMemoryRequirements);
```

## [](#_parameters)Parameters

- `device` is the logical device that owns the data graph pipeline session.
- `pInfo` is a pointer to a [VkDataGraphPipelineSessionMemoryRequirementsInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineSessionMemoryRequirementsInfoARM.html) structure containing parameters for the memory requirements query.
- `pMemoryRequirements` is a pointer to a [VkMemoryRequirements2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryRequirements2.html) structure in which the memory requirements of the data graph pipeline session object are returned.

## [](#_description)Description

Valid Usage

- [](#VUID-vkGetDataGraphPipelineSessionMemoryRequirementsARM-bindPoint-09784)VUID-vkGetDataGraphPipelineSessionMemoryRequirementsARM-bindPoint-09784  
  The `bindPoint` member of `pInfo` **must** have been returned as part of a [VkDataGraphPipelineSessionBindPointRequirementARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineSessionBindPointRequirementARM.html) whose `bindPointType` member is `VK_DATA_GRAPH_PIPELINE_SESSION_BIND_POINT_TYPE_MEMORY_ARM` by a prior call to [vkGetDataGraphPipelineSessionMemoryRequirementsARM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDataGraphPipelineSessionMemoryRequirementsARM.html) for the `session` member of `pInfo`

Valid Usage (Implicit)

- [](#VUID-vkGetDataGraphPipelineSessionMemoryRequirementsARM-device-parameter)VUID-vkGetDataGraphPipelineSessionMemoryRequirementsARM-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetDataGraphPipelineSessionMemoryRequirementsARM-pInfo-parameter)VUID-vkGetDataGraphPipelineSessionMemoryRequirementsARM-pInfo-parameter  
  `pInfo` **must** be a valid pointer to a valid [VkDataGraphPipelineSessionMemoryRequirementsInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineSessionMemoryRequirementsInfoARM.html) structure
- [](#VUID-vkGetDataGraphPipelineSessionMemoryRequirementsARM-pMemoryRequirements-parameter)VUID-vkGetDataGraphPipelineSessionMemoryRequirementsARM-pMemoryRequirements-parameter  
  `pMemoryRequirements` **must** be a valid pointer to a [VkMemoryRequirements2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryRequirements2.html) structure

## [](#_see_also)See Also

[VK\_ARM\_data\_graph](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_data_graph.html), [VkDataGraphPipelineSessionMemoryRequirementsInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineSessionMemoryRequirementsInfoARM.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkMemoryRequirements2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryRequirements2.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetDataGraphPipelineSessionMemoryRequirementsARM)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0