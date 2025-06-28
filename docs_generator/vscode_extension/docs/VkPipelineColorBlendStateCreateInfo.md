# VkPipelineColorBlendStateCreateInfo(3) Manual Page

## Name

VkPipelineColorBlendStateCreateInfo - Structure specifying parameters of a newly created pipeline color blend state



## [](#_c_specification)C Specification

The `VkPipelineColorBlendStateCreateInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_0
typedef struct VkPipelineColorBlendStateCreateInfo {
    VkStructureType                               sType;
    const void*                                   pNext;
    VkPipelineColorBlendStateCreateFlags          flags;
    VkBool32                                      logicOpEnable;
    VkLogicOp                                     logicOp;
    uint32_t                                      attachmentCount;
    const VkPipelineColorBlendAttachmentState*    pAttachments;
    float                                         blendConstants[4];
} VkPipelineColorBlendStateCreateInfo;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is a bitmask of [VkPipelineColorBlendStateCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineColorBlendStateCreateFlagBits.html) specifying additional color blending information.
- `logicOpEnable` controls whether to apply [Logical Operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#framebuffer-logicop).
- `logicOp` selects which logical operation to apply.
- `attachmentCount` is the number of [VkPipelineColorBlendAttachmentState](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineColorBlendAttachmentState.html) elements in `pAttachments`. It is ignored if the pipeline is created with `VK_DYNAMIC_STATE_COLOR_BLEND_ENABLE_EXT`, `VK_DYNAMIC_STATE_COLOR_BLEND_EQUATION_EXT`, and `VK_DYNAMIC_STATE_COLOR_WRITE_MASK_EXT` dynamic states set, and either `VK_DYNAMIC_STATE_COLOR_BLEND_ADVANCED_EXT` set or the [advancedBlendCoherentOperations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-advancedBlendCoherentOperations) feature is not enabled.
- `pAttachments` is a pointer to an array of [VkPipelineColorBlendAttachmentState](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineColorBlendAttachmentState.html) structures defining blend state for each color attachment. It is ignored if the pipeline is created with `VK_DYNAMIC_STATE_COLOR_BLEND_ENABLE_EXT`, `VK_DYNAMIC_STATE_COLOR_BLEND_EQUATION_EXT`, and `VK_DYNAMIC_STATE_COLOR_WRITE_MASK_EXT` dynamic states set, and either `VK_DYNAMIC_STATE_COLOR_BLEND_ADVANCED_EXT` set or the [advancedBlendCoherentOperations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-advancedBlendCoherentOperations) feature is not enabled.
- `blendConstants` is a pointer to an array of four values used as the R, G, B, and A components of the blend constant that are used in blending, depending on the [blend factor](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#framebuffer-blendfactors).

## [](#_description)Description

Valid Usage

- [](#VUID-VkPipelineColorBlendStateCreateInfo-pAttachments-00605)VUID-VkPipelineColorBlendStateCreateInfo-pAttachments-00605  
  If the [`independentBlend`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-independentBlend) feature is not enabled, all elements of `pAttachments` **must** be identical
- [](#VUID-VkPipelineColorBlendStateCreateInfo-logicOpEnable-00606)VUID-VkPipelineColorBlendStateCreateInfo-logicOpEnable-00606  
  If the [`logicOp`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-logicOp) feature is not enabled, `logicOpEnable` **must** be `VK_FALSE`
- [](#VUID-VkPipelineColorBlendStateCreateInfo-logicOpEnable-00607)VUID-VkPipelineColorBlendStateCreateInfo-logicOpEnable-00607  
  If `logicOpEnable` is `VK_TRUE`, `logicOp` **must** be a valid [VkLogicOp](https://registry.khronos.org/vulkan/specs/latest/man/html/VkLogicOp.html) value
- [](#VUID-VkPipelineColorBlendStateCreateInfo-rasterizationOrderColorAttachmentAccess-06465)VUID-VkPipelineColorBlendStateCreateInfo-rasterizationOrderColorAttachmentAccess-06465  
  If the [`rasterizationOrderColorAttachmentAccess`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-rasterizationOrderColorAttachmentAccess) feature is not enabled, `flags` **must** not include `VK_PIPELINE_COLOR_BLEND_STATE_CREATE_RASTERIZATION_ORDER_ATTACHMENT_ACCESS_BIT_EXT`
- [](#VUID-VkPipelineColorBlendStateCreateInfo-pAttachments-07353)VUID-VkPipelineColorBlendStateCreateInfo-pAttachments-07353  
  If `attachmentCount` is not `0` , and any of `VK_DYNAMIC_STATE_COLOR_BLEND_ADVANCED_EXT`, `VK_DYNAMIC_STATE_COLOR_BLEND_ENABLE_EXT`, `VK_DYNAMIC_STATE_COLOR_BLEND_EQUATION_EXT`, or `VK_DYNAMIC_STATE_COLOR_WRITE_MASK_EXT` are not set, `pAttachments` **must** be a valid pointer to an array of `attachmentCount` valid [VkPipelineColorBlendAttachmentState](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineColorBlendAttachmentState.html) structures

Valid Usage (Implicit)

- [](#VUID-VkPipelineColorBlendStateCreateInfo-sType-sType)VUID-VkPipelineColorBlendStateCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PIPELINE_COLOR_BLEND_STATE_CREATE_INFO`
- [](#VUID-VkPipelineColorBlendStateCreateInfo-pNext-pNext)VUID-VkPipelineColorBlendStateCreateInfo-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the `pNext` chain **must** be either `NULL` or a pointer to a valid instance of [VkPipelineColorBlendAdvancedStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineColorBlendAdvancedStateCreateInfoEXT.html) or [VkPipelineColorWriteCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineColorWriteCreateInfoEXT.html)
- [](#VUID-VkPipelineColorBlendStateCreateInfo-sType-unique)VUID-VkPipelineColorBlendStateCreateInfo-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique
- [](#VUID-VkPipelineColorBlendStateCreateInfo-flags-parameter)VUID-VkPipelineColorBlendStateCreateInfo-flags-parameter  
  `flags` **must** be a valid combination of [VkPipelineColorBlendStateCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineColorBlendStateCreateFlagBits.html) values
- [](#VUID-VkPipelineColorBlendStateCreateInfo-pAttachments-parameter)VUID-VkPipelineColorBlendStateCreateInfo-pAttachments-parameter  
  If `attachmentCount` is not `0`, and `pAttachments` is not `NULL`, `pAttachments` **must** be a valid pointer to an array of `attachmentCount` valid [VkPipelineColorBlendAttachmentState](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineColorBlendAttachmentState.html) structures

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineCreateInfo.html), [VkLogicOp](https://registry.khronos.org/vulkan/specs/latest/man/html/VkLogicOp.html), [VkPipelineColorBlendAttachmentState](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineColorBlendAttachmentState.html), [VkPipelineColorBlendStateCreateFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineColorBlendStateCreateFlags.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPipelineColorBlendStateCreateInfo)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0