# VkPipelineColorBlendStateCreateInfo(3) Manual Page

## Name

VkPipelineColorBlendStateCreateInfo - Structure specifying parameters of
a newly created pipeline color blend state



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPipelineColorBlendStateCreateInfo` structure is defined as:

``` c
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

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is a bitmask of
  [VkPipelineColorBlendStateCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineColorBlendStateCreateFlagBits.html)
  specifying additional color blending information.

- `logicOpEnable` controls whether to apply <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#framebuffer-logicop"
  target="_blank" rel="noopener">Logical Operations</a>.

- `logicOp` selects which logical operation to apply.

- `attachmentCount` is the number of
  [VkPipelineColorBlendAttachmentState](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineColorBlendAttachmentState.html)
  elements in `pAttachments`. It is ignored if the pipeline is created
  with `VK_DYNAMIC_STATE_COLOR_BLEND_ENABLE_EXT`,
  `VK_DYNAMIC_STATE_COLOR_BLEND_EQUATION_EXT`, and
  `VK_DYNAMIC_STATE_COLOR_WRITE_MASK_EXT` dynamic states set, and either
  `VK_DYNAMIC_STATE_COLOR_BLEND_ADVANCED_EXT` set or <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-advancedBlendCoherentOperations"
  target="_blank" rel="noopener">advancedBlendCoherentOperations</a> is
  not enabled on the device.

- `pAttachments` is a pointer to an array of
  [VkPipelineColorBlendAttachmentState](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineColorBlendAttachmentState.html)
  structures defining blend state for each color attachment. It is
  ignored if the pipeline is created with
  `VK_DYNAMIC_STATE_COLOR_BLEND_ENABLE_EXT`,
  `VK_DYNAMIC_STATE_COLOR_BLEND_EQUATION_EXT`, and
  `VK_DYNAMIC_STATE_COLOR_WRITE_MASK_EXT` dynamic states set, and either
  `VK_DYNAMIC_STATE_COLOR_BLEND_ADVANCED_EXT` set or <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-advancedBlendCoherentOperations"
  target="_blank" rel="noopener">advancedBlendCoherentOperations</a> is
  not enabled on the device.

- `blendConstants` is a pointer to an array of four values used as the
  R, G, B, and A components of the blend constant that are used in
  blending, depending on the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#framebuffer-blendfactors"
  target="_blank" rel="noopener">blend factor</a>.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkPipelineColorBlendStateCreateInfo-pAttachments-00605"
  id="VUID-VkPipelineColorBlendStateCreateInfo-pAttachments-00605"></a>
  VUID-VkPipelineColorBlendStateCreateInfo-pAttachments-00605  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-independentBlend"
  target="_blank" rel="noopener"><code>independentBlend</code></a>
  feature is not enabled, all elements of `pAttachments` **must** be
  identical

- <a href="#VUID-VkPipelineColorBlendStateCreateInfo-logicOpEnable-00606"
  id="VUID-VkPipelineColorBlendStateCreateInfo-logicOpEnable-00606"></a>
  VUID-VkPipelineColorBlendStateCreateInfo-logicOpEnable-00606  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-logicOp"
  target="_blank" rel="noopener"><code>logicOp</code></a> feature is not
  enabled, `logicOpEnable` **must** be `VK_FALSE`

- <a href="#VUID-VkPipelineColorBlendStateCreateInfo-logicOpEnable-00607"
  id="VUID-VkPipelineColorBlendStateCreateInfo-logicOpEnable-00607"></a>
  VUID-VkPipelineColorBlendStateCreateInfo-logicOpEnable-00607  
  If `logicOpEnable` is `VK_TRUE`, `logicOp` **must** be a valid
  [VkLogicOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLogicOp.html) value

- <a
  href="#VUID-VkPipelineColorBlendStateCreateInfo-rasterizationOrderColorAttachmentAccess-06465"
  id="VUID-VkPipelineColorBlendStateCreateInfo-rasterizationOrderColorAttachmentAccess-06465"></a>
  VUID-VkPipelineColorBlendStateCreateInfo-rasterizationOrderColorAttachmentAccess-06465  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-rasterizationOrderColorAttachmentAccess"
  target="_blank"
  rel="noopener"><code>rasterizationOrderColorAttachmentAccess</code></a>
  feature is not enabled, `flags` **must** not include
  `VK_PIPELINE_COLOR_BLEND_STATE_CREATE_RASTERIZATION_ORDER_ATTACHMENT_ACCESS_BIT_EXT`

- <a href="#VUID-VkPipelineColorBlendStateCreateInfo-pAttachments-07353"
  id="VUID-VkPipelineColorBlendStateCreateInfo-pAttachments-07353"></a>
  VUID-VkPipelineColorBlendStateCreateInfo-pAttachments-07353  
  If `attachmentCount` is not `0` , and any of
  `VK_DYNAMIC_STATE_COLOR_BLEND_ADVANCED_EXT`,
  `VK_DYNAMIC_STATE_COLOR_BLEND_ENABLE_EXT`,
  `VK_DYNAMIC_STATE_COLOR_BLEND_EQUATION_EXT`, or
  `VK_DYNAMIC_STATE_COLOR_WRITE_MASK_EXT` are not set, `pAttachments`
  **must** be a valid pointer to an array of `attachmentCount` valid
  [VkPipelineColorBlendAttachmentState](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineColorBlendAttachmentState.html)
  structures

Valid Usage (Implicit)

- <a href="#VUID-VkPipelineColorBlendStateCreateInfo-sType-sType"
  id="VUID-VkPipelineColorBlendStateCreateInfo-sType-sType"></a>
  VUID-VkPipelineColorBlendStateCreateInfo-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PIPELINE_COLOR_BLEND_STATE_CREATE_INFO`

- <a href="#VUID-VkPipelineColorBlendStateCreateInfo-pNext-pNext"
  id="VUID-VkPipelineColorBlendStateCreateInfo-pNext-pNext"></a>
  VUID-VkPipelineColorBlendStateCreateInfo-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the
  `pNext` chain **must** be either `NULL` or a pointer to a valid
  instance of
  [VkPipelineColorBlendAdvancedStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineColorBlendAdvancedStateCreateInfoEXT.html)
  or
  [VkPipelineColorWriteCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineColorWriteCreateInfoEXT.html)

- <a href="#VUID-VkPipelineColorBlendStateCreateInfo-sType-unique"
  id="VUID-VkPipelineColorBlendStateCreateInfo-sType-unique"></a>
  VUID-VkPipelineColorBlendStateCreateInfo-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

- <a href="#VUID-VkPipelineColorBlendStateCreateInfo-flags-parameter"
  id="VUID-VkPipelineColorBlendStateCreateInfo-flags-parameter"></a>
  VUID-VkPipelineColorBlendStateCreateInfo-flags-parameter  
  `flags` **must** be a valid combination of
  [VkPipelineColorBlendStateCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineColorBlendStateCreateFlagBits.html)
  values

- <a
  href="#VUID-VkPipelineColorBlendStateCreateInfo-pAttachments-parameter"
  id="VUID-VkPipelineColorBlendStateCreateInfo-pAttachments-parameter"></a>
  VUID-VkPipelineColorBlendStateCreateInfo-pAttachments-parameter  
  If `attachmentCount` is not `0`, and `pAttachments` is not `NULL`,
  `pAttachments` **must** be a valid pointer to an array of
  `attachmentCount` valid
  [VkPipelineColorBlendAttachmentState](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineColorBlendAttachmentState.html)
  structures

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html), [VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html),
[VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineCreateInfo.html),
[VkLogicOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLogicOp.html),
[VkPipelineColorBlendAttachmentState](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineColorBlendAttachmentState.html),
[VkPipelineColorBlendStateCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineColorBlendStateCreateFlags.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPipelineColorBlendStateCreateInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
