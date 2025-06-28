# VkPipelineTessellationStateCreateInfo(3) Manual Page

## Name

VkPipelineTessellationStateCreateInfo - Structure specifying parameters of a newly created pipeline tessellation state



## [](#_c_specification)C Specification

The `VkPipelineTessellationStateCreateInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_0
typedef struct VkPipelineTessellationStateCreateInfo {
    VkStructureType                           sType;
    const void*                               pNext;
    VkPipelineTessellationStateCreateFlags    flags;
    uint32_t                                  patchControlPoints;
} VkPipelineTessellationStateCreateInfo;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is reserved for future use.
- `patchControlPoints` is the number of control points per patch.

## [](#_description)Description

Valid Usage

- [](#VUID-VkPipelineTessellationStateCreateInfo-patchControlPoints-01214)VUID-VkPipelineTessellationStateCreateInfo-patchControlPoints-01214  
  `patchControlPoints` **must** be greater than zero and less than or equal to `VkPhysicalDeviceLimits`::`maxTessellationPatchSize`

Valid Usage (Implicit)

- [](#VUID-VkPipelineTessellationStateCreateInfo-sType-sType)VUID-VkPipelineTessellationStateCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PIPELINE_TESSELLATION_STATE_CREATE_INFO`
- [](#VUID-VkPipelineTessellationStateCreateInfo-pNext-pNext)VUID-VkPipelineTessellationStateCreateInfo-pNext-pNext  
  `pNext` **must** be `NULL` or a pointer to a valid instance of [VkPipelineTessellationDomainOriginStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineTessellationDomainOriginStateCreateInfo.html)
- [](#VUID-VkPipelineTessellationStateCreateInfo-sType-unique)VUID-VkPipelineTessellationStateCreateInfo-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique
- [](#VUID-VkPipelineTessellationStateCreateInfo-flags-zerobitmask)VUID-VkPipelineTessellationStateCreateInfo-flags-zerobitmask  
  `flags` **must** be `0`

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineCreateInfo.html), [VkGraphicsShaderGroupCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsShaderGroupCreateInfoNV.html), [VkPipelineTessellationStateCreateFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineTessellationStateCreateFlags.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPipelineTessellationStateCreateInfo)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0