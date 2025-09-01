# VkPipelineDynamicStateCreateInfo(3) Manual Page

## Name

VkPipelineDynamicStateCreateInfo - Structure specifying parameters of a newly created pipeline dynamic state



## [](#_c_specification)C Specification

The `VkPipelineDynamicStateCreateInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_0
typedef struct VkPipelineDynamicStateCreateInfo {
    VkStructureType                      sType;
    const void*                          pNext;
    VkPipelineDynamicStateCreateFlags    flags;
    uint32_t                             dynamicStateCount;
    const VkDynamicState*                pDynamicStates;
} VkPipelineDynamicStateCreateInfo;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is reserved for future use.
- `dynamicStateCount` is the number of elements in the `pDynamicStates` array.
- `pDynamicStates` is a pointer to an array of [VkDynamicState](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDynamicState.html) values specifying which pieces of pipeline state will use the values from dynamic state commands rather than from pipeline state creation information.

## [](#_description)Description

Valid Usage

- [](#VUID-VkPipelineDynamicStateCreateInfo-pDynamicStates-01442)VUID-VkPipelineDynamicStateCreateInfo-pDynamicStates-01442  
  Each element of `pDynamicStates` **must** be unique

Valid Usage (Implicit)

- [](#VUID-VkPipelineDynamicStateCreateInfo-sType-sType)VUID-VkPipelineDynamicStateCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PIPELINE_DYNAMIC_STATE_CREATE_INFO`
- [](#VUID-VkPipelineDynamicStateCreateInfo-pNext-pNext)VUID-VkPipelineDynamicStateCreateInfo-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkPipelineDynamicStateCreateInfo-flags-zerobitmask)VUID-VkPipelineDynamicStateCreateInfo-flags-zerobitmask  
  `flags` **must** be `0`
- [](#VUID-VkPipelineDynamicStateCreateInfo-pDynamicStates-parameter)VUID-VkPipelineDynamicStateCreateInfo-pDynamicStates-parameter  
  If `dynamicStateCount` is not `0`, `pDynamicStates` **must** be a valid pointer to an array of `dynamicStateCount` valid [VkDynamicState](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDynamicState.html) values

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkDynamicState](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDynamicState.html), [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineCreateInfo.html), [VkPipelineDynamicStateCreateFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineDynamicStateCreateFlags.html), [VkRayTracingPipelineCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineCreateInfoKHR.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPipelineDynamicStateCreateInfo).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0