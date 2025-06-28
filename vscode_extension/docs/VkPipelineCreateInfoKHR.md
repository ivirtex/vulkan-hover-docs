# VkPipelineCreateInfoKHR(3) Manual Page

## Name

VkPipelineCreateInfoKHR - Structure specifying a pipeline createinfo chain



## [](#_c_specification)C Specification

The `VkPipelineCreateInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_pipeline_binary
typedef struct VkPipelineCreateInfoKHR {
    VkStructureType    sType;
    void*              pNext;
} VkPipelineCreateInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is a pointer to a structure extending this structure.

## [](#_description)Description

Valid Usage

- [](#VUID-VkPipelineCreateInfoKHR-pNext-09604)VUID-VkPipelineCreateInfoKHR-pNext-09604  
  `pNext` **must** be pointer to a valid instance of [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineCreateInfo.html), [VkExecutionGraphPipelineCreateInfoAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExecutionGraphPipelineCreateInfoAMDX.html), [VkRayTracingPipelineCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineCreateInfoKHR.html), or [VkComputePipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComputePipelineCreateInfo.html)

Valid Usage (Implicit)

- [](#VUID-VkPipelineCreateInfoKHR-sType-sType)VUID-VkPipelineCreateInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PIPELINE_CREATE_INFO_KHR`

## [](#_see_also)See Also

[VK\_KHR\_pipeline\_binary](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_pipeline_binary.html), [VkPipelineBinaryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryCreateInfoKHR.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetPipelineKeyKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPipelineKeyKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPipelineCreateInfoKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0