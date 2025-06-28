# VkPipelineVertexInputStateCreateInfo(3) Manual Page

## Name

VkPipelineVertexInputStateCreateInfo - Structure specifying parameters of a newly created pipeline vertex input state



## [](#_c_specification)C Specification

The `VkPipelineVertexInputStateCreateInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_0
typedef struct VkPipelineVertexInputStateCreateInfo {
    VkStructureType                             sType;
    const void*                                 pNext;
    VkPipelineVertexInputStateCreateFlags       flags;
    uint32_t                                    vertexBindingDescriptionCount;
    const VkVertexInputBindingDescription*      pVertexBindingDescriptions;
    uint32_t                                    vertexAttributeDescriptionCount;
    const VkVertexInputAttributeDescription*    pVertexAttributeDescriptions;
} VkPipelineVertexInputStateCreateInfo;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is reserved for future use.
- `vertexBindingDescriptionCount` is the number of vertex binding descriptions provided in `pVertexBindingDescriptions`.
- `pVertexBindingDescriptions` is a pointer to an array of [VkVertexInputBindingDescription](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVertexInputBindingDescription.html) structures.
- `vertexAttributeDescriptionCount` is the number of vertex attribute descriptions provided in `pVertexAttributeDescriptions`.
- `pVertexAttributeDescriptions` is a pointer to an array of [VkVertexInputAttributeDescription](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVertexInputAttributeDescription.html) structures.

## [](#_description)Description

Valid Usage

- [](#VUID-VkPipelineVertexInputStateCreateInfo-vertexBindingDescriptionCount-00613)VUID-VkPipelineVertexInputStateCreateInfo-vertexBindingDescriptionCount-00613  
  `vertexBindingDescriptionCount` **must** be less than or equal to `VkPhysicalDeviceLimits`::`maxVertexInputBindings`
- [](#VUID-VkPipelineVertexInputStateCreateInfo-vertexAttributeDescriptionCount-00614)VUID-VkPipelineVertexInputStateCreateInfo-vertexAttributeDescriptionCount-00614  
  `vertexAttributeDescriptionCount` **must** be less than or equal to `VkPhysicalDeviceLimits`::`maxVertexInputAttributes`
- [](#VUID-VkPipelineVertexInputStateCreateInfo-binding-00615)VUID-VkPipelineVertexInputStateCreateInfo-binding-00615  
  For every `binding` specified by each element of `pVertexAttributeDescriptions`, a [VkVertexInputBindingDescription](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVertexInputBindingDescription.html) **must** exist in `pVertexBindingDescriptions` with the same value of `binding`
- [](#VUID-VkPipelineVertexInputStateCreateInfo-pVertexBindingDescriptions-00616)VUID-VkPipelineVertexInputStateCreateInfo-pVertexBindingDescriptions-00616  
  All elements of `pVertexBindingDescriptions` **must** describe distinct binding numbers
- [](#VUID-VkPipelineVertexInputStateCreateInfo-pVertexAttributeDescriptions-00617)VUID-VkPipelineVertexInputStateCreateInfo-pVertexAttributeDescriptions-00617  
  All elements of `pVertexAttributeDescriptions` **must** describe distinct attribute locations

Valid Usage (Implicit)

- [](#VUID-VkPipelineVertexInputStateCreateInfo-sType-sType)VUID-VkPipelineVertexInputStateCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PIPELINE_VERTEX_INPUT_STATE_CREATE_INFO`
- [](#VUID-VkPipelineVertexInputStateCreateInfo-pNext-pNext)VUID-VkPipelineVertexInputStateCreateInfo-pNext-pNext  
  `pNext` **must** be `NULL` or a pointer to a valid instance of [VkPipelineVertexInputDivisorStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineVertexInputDivisorStateCreateInfo.html)
- [](#VUID-VkPipelineVertexInputStateCreateInfo-sType-unique)VUID-VkPipelineVertexInputStateCreateInfo-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique
- [](#VUID-VkPipelineVertexInputStateCreateInfo-flags-zerobitmask)VUID-VkPipelineVertexInputStateCreateInfo-flags-zerobitmask  
  `flags` **must** be `0`
- [](#VUID-VkPipelineVertexInputStateCreateInfo-pVertexBindingDescriptions-parameter)VUID-VkPipelineVertexInputStateCreateInfo-pVertexBindingDescriptions-parameter  
  If `vertexBindingDescriptionCount` is not `0`, `pVertexBindingDescriptions` **must** be a valid pointer to an array of `vertexBindingDescriptionCount` valid [VkVertexInputBindingDescription](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVertexInputBindingDescription.html) structures
- [](#VUID-VkPipelineVertexInputStateCreateInfo-pVertexAttributeDescriptions-parameter)VUID-VkPipelineVertexInputStateCreateInfo-pVertexAttributeDescriptions-parameter  
  If `vertexAttributeDescriptionCount` is not `0`, `pVertexAttributeDescriptions` **must** be a valid pointer to an array of `vertexAttributeDescriptionCount` valid [VkVertexInputAttributeDescription](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVertexInputAttributeDescription.html) structures

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineCreateInfo.html), [VkGraphicsShaderGroupCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsShaderGroupCreateInfoNV.html), [VkPipelineVertexInputStateCreateFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineVertexInputStateCreateFlags.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkVertexInputAttributeDescription](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVertexInputAttributeDescription.html), [VkVertexInputBindingDescription](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVertexInputBindingDescription.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPipelineVertexInputStateCreateInfo)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0