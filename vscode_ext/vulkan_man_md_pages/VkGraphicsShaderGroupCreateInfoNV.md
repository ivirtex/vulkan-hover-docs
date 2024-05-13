# VkGraphicsShaderGroupCreateInfoNV(3) Manual Page

## Name

VkGraphicsShaderGroupCreateInfoNV - Structure specifying override
parameters for each shader group



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkGraphicsShaderGroupCreateInfoNV` structure provides the state
overrides for each shader group. Each shader group behaves like a
pipeline that was created from its state as well as the remaining
parent’s state. It is defined as:

``` c
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

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `stageCount` is the number of entries in the `pStages` array.

- `pStages` is a pointer to an array
  [VkPipelineShaderStageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageCreateInfo.html)
  structures specifying the set of the shader stages to be included in
  this shader group.

- `pVertexInputState` is a pointer to a
  [VkPipelineVertexInputStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineVertexInputStateCreateInfo.html)
  structure.

- `pTessellationState` is a pointer to a
  [VkPipelineTessellationStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineTessellationStateCreateInfo.html)
  structure, and is ignored if the shader group does not include a
  tessellation control shader stage and tessellation evaluation shader
  stage.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkGraphicsShaderGroupCreateInfoNV-stageCount-02888"
  id="VUID-VkGraphicsShaderGroupCreateInfoNV-stageCount-02888"></a>
  VUID-VkGraphicsShaderGroupCreateInfoNV-stageCount-02888  
  For `stageCount`, the same restrictions as in
  [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineCreateInfo.html)::`stageCount`
  apply

- <a href="#VUID-VkGraphicsShaderGroupCreateInfoNV-pStages-02889"
  id="VUID-VkGraphicsShaderGroupCreateInfoNV-pStages-02889"></a>
  VUID-VkGraphicsShaderGroupCreateInfoNV-pStages-02889  
  For `pStages`, the same restrictions as in
  [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineCreateInfo.html)::`pStages`
  apply

- <a
  href="#VUID-VkGraphicsShaderGroupCreateInfoNV-pVertexInputState-02890"
  id="VUID-VkGraphicsShaderGroupCreateInfoNV-pVertexInputState-02890"></a>
  VUID-VkGraphicsShaderGroupCreateInfoNV-pVertexInputState-02890  
  For `pVertexInputState`, the same restrictions as in
  [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineCreateInfo.html)::`pVertexInputState`
  apply

- <a
  href="#VUID-VkGraphicsShaderGroupCreateInfoNV-pTessellationState-02891"
  id="VUID-VkGraphicsShaderGroupCreateInfoNV-pTessellationState-02891"></a>
  VUID-VkGraphicsShaderGroupCreateInfoNV-pTessellationState-02891  
  For `pTessellationState`, the same restrictions as in
  [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineCreateInfo.html)::`pTessellationState`
  apply

Valid Usage (Implicit)

- <a href="#VUID-VkGraphicsShaderGroupCreateInfoNV-sType-sType"
  id="VUID-VkGraphicsShaderGroupCreateInfoNV-sType-sType"></a>
  VUID-VkGraphicsShaderGroupCreateInfoNV-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_GRAPHICS_SHADER_GROUP_CREATE_INFO_NV`

- <a href="#VUID-VkGraphicsShaderGroupCreateInfoNV-pNext-pNext"
  id="VUID-VkGraphicsShaderGroupCreateInfoNV-pNext-pNext"></a>
  VUID-VkGraphicsShaderGroupCreateInfoNV-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkGraphicsShaderGroupCreateInfoNV-pStages-parameter"
  id="VUID-VkGraphicsShaderGroupCreateInfoNV-pStages-parameter"></a>
  VUID-VkGraphicsShaderGroupCreateInfoNV-pStages-parameter  
  `pStages` **must** be a valid pointer to an array of `stageCount`
  valid
  [VkPipelineShaderStageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageCreateInfo.html)
  structures

- <a href="#VUID-VkGraphicsShaderGroupCreateInfoNV-stageCount-arraylength"
  id="VUID-VkGraphicsShaderGroupCreateInfoNV-stageCount-arraylength"></a>
  VUID-VkGraphicsShaderGroupCreateInfoNV-stageCount-arraylength  
  `stageCount` **must** be greater than `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_device_generated_commands](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_device_generated_commands.html),
[VkGraphicsPipelineShaderGroupsCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineShaderGroupsCreateInfoNV.html),
[VkPipelineShaderStageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageCreateInfo.html),
[VkPipelineTessellationStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineTessellationStateCreateInfo.html),
[VkPipelineVertexInputStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineVertexInputStateCreateInfo.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkGraphicsShaderGroupCreateInfoNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
