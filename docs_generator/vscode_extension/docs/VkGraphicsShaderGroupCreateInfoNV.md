# VkGraphicsShaderGroupCreateInfoNV(3) Manual Page

## Name

VkGraphicsShaderGroupCreateInfoNV - Structure specifying override parameters for each shader group



## [](#_c_specification)C Specification

The `VkGraphicsShaderGroupCreateInfoNV` structure provides the state overrides for each shader group. Each shader group behaves like a pipeline that was created from its state as well as the remaining parent’s state. It is defined as:

```c++
// Provided by VK_NV_device_generated_commands
typedef struct VkGraphicsShaderGroupCreateInfoNV {
    VkStructureType                                 sType;
    const void*                                     pNext;
    uint32_t                                        stageCount;
    const VkPipelineShaderStageCreateInfo*          pStages;
    const VkPipelineVertexInputStateCreateInfo*     pVertexInputState;
    const VkPipelineTessellationStateCreateInfo*    pTessellationState;
} VkGraphicsShaderGroupCreateInfoNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `stageCount` is the number of entries in the `pStages` array.
- `pStages` is a pointer to an array [VkPipelineShaderStageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineShaderStageCreateInfo.html) structures specifying the set of the shader stages to be included in this shader group.
- `pVertexInputState` is a pointer to a [VkPipelineVertexInputStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineVertexInputStateCreateInfo.html) structure.
- `pTessellationState` is a pointer to a [VkPipelineTessellationStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineTessellationStateCreateInfo.html) structure, and is ignored if the shader group does not include a tessellation control shader stage and tessellation evaluation shader stage.

## [](#_description)Description

Valid Usage

- [](#VUID-VkGraphicsShaderGroupCreateInfoNV-stageCount-02888)VUID-VkGraphicsShaderGroupCreateInfoNV-stageCount-02888  
  For `stageCount`, the same restrictions as in [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineCreateInfo.html)::`stageCount` apply
- [](#VUID-VkGraphicsShaderGroupCreateInfoNV-pStages-02889)VUID-VkGraphicsShaderGroupCreateInfoNV-pStages-02889  
  For `pStages`, the same restrictions as in [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineCreateInfo.html)::`pStages` apply
- [](#VUID-VkGraphicsShaderGroupCreateInfoNV-pVertexInputState-02890)VUID-VkGraphicsShaderGroupCreateInfoNV-pVertexInputState-02890  
  For `pVertexInputState`, the same restrictions as in [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineCreateInfo.html)::`pVertexInputState` apply
- [](#VUID-VkGraphicsShaderGroupCreateInfoNV-pTessellationState-02891)VUID-VkGraphicsShaderGroupCreateInfoNV-pTessellationState-02891  
  For `pTessellationState`, the same restrictions as in [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineCreateInfo.html)::`pTessellationState` apply

Valid Usage (Implicit)

- [](#VUID-VkGraphicsShaderGroupCreateInfoNV-sType-sType)VUID-VkGraphicsShaderGroupCreateInfoNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_GRAPHICS_SHADER_GROUP_CREATE_INFO_NV`
- [](#VUID-VkGraphicsShaderGroupCreateInfoNV-pNext-pNext)VUID-VkGraphicsShaderGroupCreateInfoNV-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkGraphicsShaderGroupCreateInfoNV-pStages-parameter)VUID-VkGraphicsShaderGroupCreateInfoNV-pStages-parameter  
  `pStages` **must** be a valid pointer to an array of `stageCount` valid [VkPipelineShaderStageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineShaderStageCreateInfo.html) structures
- [](#VUID-VkGraphicsShaderGroupCreateInfoNV-stageCount-arraylength)VUID-VkGraphicsShaderGroupCreateInfoNV-stageCount-arraylength  
  `stageCount` **must** be greater than `0`

## [](#_see_also)See Also

[VK\_NV\_device\_generated\_commands](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_device_generated_commands.html), [VkGraphicsPipelineShaderGroupsCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineShaderGroupsCreateInfoNV.html), [VkPipelineShaderStageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineShaderStageCreateInfo.html), [VkPipelineTessellationStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineTessellationStateCreateInfo.html), [VkPipelineVertexInputStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineVertexInputStateCreateInfo.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkGraphicsShaderGroupCreateInfoNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0