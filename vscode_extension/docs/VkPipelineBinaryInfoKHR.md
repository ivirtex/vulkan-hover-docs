# VkPipelineBinaryInfoKHR(3) Manual Page

## Name

VkPipelineBinaryInfoKHR - Structure specifying pipeline binaries to use during pipeline creation



## [](#_c_specification)C Specification

The `VkPipelineBinaryInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_pipeline_binary
typedef struct VkPipelineBinaryInfoKHR {
    VkStructureType               sType;
    const void*                   pNext;
    uint32_t                      binaryCount;
    const VkPipelineBinaryKHR*    pPipelineBinaries;
} VkPipelineBinaryInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `binaryCount` is the number of elements in the `pPipelineBinaries` array.
- `pPipelineBinaries` is a pointer to an array of [VkPipelineBinaryKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryKHR.html) handles.

## [](#_description)Description

If a `VkPipelineBinaryInfoKHR` structure with a `binaryCount` greater than 0 is included in the `pNext` chain of any `Vk*PipelineCreateInfo` structure when creating a pipeline, implementations **must** use the data in `pPipelineBinaries` instead of recalculating it. Any shader module identifiers or shader modules declared in [VkPipelineShaderStageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineShaderStageCreateInfo.html) instances are ignored.

If this structure is not included in the `pNext` chain, it is equivalent to specifying this structure with a `binaryCount` of `0`.

Valid Usage

- [](#VUID-VkPipelineBinaryInfoKHR-binaryCount-09603)VUID-VkPipelineBinaryInfoKHR-binaryCount-09603  
  `binaryCount` and the order of the elements in `pPipelineBinaries` **must** exactly match that returned by [vkCreatePipelineBinariesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreatePipelineBinariesKHR.html) for the matching `Vk*PipelineCreateInfo` structure and its `pNext` chain, ignoring the presence of the `VkPipelineBinaryInfoKHR` structure, the presence of the `VK_PIPELINE_CREATE_2_CAPTURE_DATA_BIT_KHR` flag, and absence of any shader module identifiers or shader modules, for the same [global pipeline key](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#global-pipeline-key), from either:
  
  - [VkPipelineBinaryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryCreateInfoKHR.html)::`pPipelineCreateInfo`, or
  - [VkPipelineBinaryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryCreateInfoKHR.html)::`pipeline`

Valid Usage (Implicit)

- [](#VUID-VkPipelineBinaryInfoKHR-sType-sType)VUID-VkPipelineBinaryInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PIPELINE_BINARY_INFO_KHR`
- [](#VUID-VkPipelineBinaryInfoKHR-pPipelineBinaries-parameter)VUID-VkPipelineBinaryInfoKHR-pPipelineBinaries-parameter  
  If `binaryCount` is not `0`, `pPipelineBinaries` **must** be a valid pointer to an array of `binaryCount` valid [VkPipelineBinaryKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryKHR.html) handles

## [](#_see_also)See Also

[VK\_KHR\_pipeline\_binary](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_pipeline_binary.html), [VkPipelineBinaryKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryKHR.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPipelineBinaryInfoKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0