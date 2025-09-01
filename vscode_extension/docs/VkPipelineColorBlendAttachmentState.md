# VkPipelineColorBlendAttachmentState(3) Manual Page

## Name

VkPipelineColorBlendAttachmentState - Structure specifying a pipeline color blend attachment state



## [](#_c_specification)C Specification

The `VkPipelineColorBlendAttachmentState` structure is defined as:

```c++
// Provided by VK_VERSION_1_0
typedef struct VkPipelineColorBlendAttachmentState {
    VkBool32                 blendEnable;
    VkBlendFactor            srcColorBlendFactor;
    VkBlendFactor            dstColorBlendFactor;
    VkBlendOp                colorBlendOp;
    VkBlendFactor            srcAlphaBlendFactor;
    VkBlendFactor            dstAlphaBlendFactor;
    VkBlendOp                alphaBlendOp;
    VkColorComponentFlags    colorWriteMask;
} VkPipelineColorBlendAttachmentState;
```

## [](#_members)Members

- `blendEnable` controls whether blending is enabled for the corresponding color attachment. If blending is not enabled, the source fragment’s color for that attachment is passed through unmodified.
- `srcColorBlendFactor` selects which blend factor is used to determine the source factors (Sr,Sg,Sb).
- `dstColorBlendFactor` selects which blend factor is used to determine the destination factors (Dr,Dg,Db).
- `colorBlendOp` selects which blend operation is used to calculate the RGB values to write to the color attachment.
- `srcAlphaBlendFactor` selects which blend factor is used to determine the source factor Sa.
- `dstAlphaBlendFactor` selects which blend factor is used to determine the destination factor Da.
- `alphaBlendOp` selects which blend operation is used to calculate the alpha values to write to the color attachment.
- `colorWriteMask` is a bitmask of [VkColorComponentFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkColorComponentFlagBits.html) specifying which of the R, G, B, and/or A components are enabled for writing, as described for the [Color Write Mask](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#framebuffer-color-write-mask).

## [](#_description)Description

Valid Usage

- [](#VUID-VkPipelineColorBlendAttachmentState-srcColorBlendFactor-00608)VUID-VkPipelineColorBlendAttachmentState-srcColorBlendFactor-00608  
  If the [`dualSrcBlend`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-dualSrcBlend) feature is not enabled, `srcColorBlendFactor` **must** not be `VK_BLEND_FACTOR_SRC1_COLOR`, `VK_BLEND_FACTOR_ONE_MINUS_SRC1_COLOR`, `VK_BLEND_FACTOR_SRC1_ALPHA`, or `VK_BLEND_FACTOR_ONE_MINUS_SRC1_ALPHA`
- [](#VUID-VkPipelineColorBlendAttachmentState-dstColorBlendFactor-00609)VUID-VkPipelineColorBlendAttachmentState-dstColorBlendFactor-00609  
  If the [`dualSrcBlend`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-dualSrcBlend) feature is not enabled, `dstColorBlendFactor` **must** not be `VK_BLEND_FACTOR_SRC1_COLOR`, `VK_BLEND_FACTOR_ONE_MINUS_SRC1_COLOR`, `VK_BLEND_FACTOR_SRC1_ALPHA`, or `VK_BLEND_FACTOR_ONE_MINUS_SRC1_ALPHA`
- [](#VUID-VkPipelineColorBlendAttachmentState-srcAlphaBlendFactor-00610)VUID-VkPipelineColorBlendAttachmentState-srcAlphaBlendFactor-00610  
  If the [`dualSrcBlend`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-dualSrcBlend) feature is not enabled, `srcAlphaBlendFactor` **must** not be `VK_BLEND_FACTOR_SRC1_COLOR`, `VK_BLEND_FACTOR_ONE_MINUS_SRC1_COLOR`, `VK_BLEND_FACTOR_SRC1_ALPHA`, or `VK_BLEND_FACTOR_ONE_MINUS_SRC1_ALPHA`
- [](#VUID-VkPipelineColorBlendAttachmentState-dstAlphaBlendFactor-00611)VUID-VkPipelineColorBlendAttachmentState-dstAlphaBlendFactor-00611  
  If the [`dualSrcBlend`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-dualSrcBlend) feature is not enabled, `dstAlphaBlendFactor` **must** not be `VK_BLEND_FACTOR_SRC1_COLOR`, `VK_BLEND_FACTOR_ONE_MINUS_SRC1_COLOR`, `VK_BLEND_FACTOR_SRC1_ALPHA`, or `VK_BLEND_FACTOR_ONE_MINUS_SRC1_ALPHA`
- [](#VUID-VkPipelineColorBlendAttachmentState-colorBlendOp-01406)VUID-VkPipelineColorBlendAttachmentState-colorBlendOp-01406  
  If either of `colorBlendOp` or `alphaBlendOp` is an [advanced blend operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#framebuffer-blend-advanced), then `colorBlendOp` **must** equal `alphaBlendOp`
- [](#VUID-VkPipelineColorBlendAttachmentState-advancedBlendIndependentBlend-01407)VUID-VkPipelineColorBlendAttachmentState-advancedBlendIndependentBlend-01407  
  If [VkPhysicalDeviceBlendOperationAdvancedPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceBlendOperationAdvancedPropertiesEXT.html)::`advancedBlendIndependentBlend` is `VK_FALSE` and `colorBlendOp` is an [advanced blend operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#framebuffer-blend-advanced), then `colorBlendOp` **must** be the same for all attachments
- [](#VUID-VkPipelineColorBlendAttachmentState-advancedBlendIndependentBlend-01408)VUID-VkPipelineColorBlendAttachmentState-advancedBlendIndependentBlend-01408  
  If [VkPhysicalDeviceBlendOperationAdvancedPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceBlendOperationAdvancedPropertiesEXT.html)::`advancedBlendIndependentBlend` is `VK_FALSE` and `alphaBlendOp` is an [advanced blend operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#framebuffer-blend-advanced), then `alphaBlendOp` **must** be the same for all attachments
- [](#VUID-VkPipelineColorBlendAttachmentState-advancedBlendAllOperations-01409)VUID-VkPipelineColorBlendAttachmentState-advancedBlendAllOperations-01409  
  If [VkPhysicalDeviceBlendOperationAdvancedPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceBlendOperationAdvancedPropertiesEXT.html)::`advancedBlendAllOperations` is `VK_FALSE`, then `colorBlendOp` **must** not be `VK_BLEND_OP_ZERO_EXT`, `VK_BLEND_OP_SRC_EXT`, `VK_BLEND_OP_DST_EXT`, `VK_BLEND_OP_SRC_OVER_EXT`, `VK_BLEND_OP_DST_OVER_EXT`, `VK_BLEND_OP_SRC_IN_EXT`, `VK_BLEND_OP_DST_IN_EXT`, `VK_BLEND_OP_SRC_OUT_EXT`, `VK_BLEND_OP_DST_OUT_EXT`, `VK_BLEND_OP_SRC_ATOP_EXT`, `VK_BLEND_OP_DST_ATOP_EXT`, `VK_BLEND_OP_XOR_EXT`, `VK_BLEND_OP_INVERT_EXT`, `VK_BLEND_OP_INVERT_RGB_EXT`, `VK_BLEND_OP_LINEARDODGE_EXT`, `VK_BLEND_OP_LINEARBURN_EXT`, `VK_BLEND_OP_VIVIDLIGHT_EXT`, `VK_BLEND_OP_LINEARLIGHT_EXT`, `VK_BLEND_OP_PINLIGHT_EXT`, `VK_BLEND_OP_HARDMIX_EXT`, `VK_BLEND_OP_PLUS_EXT`, `VK_BLEND_OP_PLUS_CLAMPED_EXT`, `VK_BLEND_OP_PLUS_CLAMPED_ALPHA_EXT`, `VK_BLEND_OP_PLUS_DARKER_EXT`, `VK_BLEND_OP_MINUS_EXT`, `VK_BLEND_OP_MINUS_CLAMPED_EXT`, `VK_BLEND_OP_CONTRAST_EXT`, `VK_BLEND_OP_INVERT_OVG_EXT`, `VK_BLEND_OP_RED_EXT`, `VK_BLEND_OP_GREEN_EXT`, or `VK_BLEND_OP_BLUE_EXT`
- [](#VUID-VkPipelineColorBlendAttachmentState-colorBlendOp-01410)VUID-VkPipelineColorBlendAttachmentState-colorBlendOp-01410  
  If `colorBlendOp` or `alphaBlendOp` is an [advanced blend operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#framebuffer-blend-advanced), then `colorAttachmentCount` of the subpass this pipeline is compiled against **must** be less than or equal to [VkPhysicalDeviceBlendOperationAdvancedPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceBlendOperationAdvancedPropertiesEXT.html)::`advancedBlendMaxColorAttachments`
- [](#VUID-VkPipelineColorBlendAttachmentState-constantAlphaColorBlendFactors-04454)VUID-VkPipelineColorBlendAttachmentState-constantAlphaColorBlendFactors-04454  
  If the `VK_KHR_portability_subset` extension is enabled, and [VkPhysicalDevicePortabilitySubsetFeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevicePortabilitySubsetFeaturesKHR.html)::`constantAlphaColorBlendFactors` is `VK_FALSE`, `srcColorBlendFactor` **must** not be `VK_BLEND_FACTOR_CONSTANT_ALPHA` or `VK_BLEND_FACTOR_ONE_MINUS_CONSTANT_ALPHA`
- [](#VUID-VkPipelineColorBlendAttachmentState-constantAlphaColorBlendFactors-04455)VUID-VkPipelineColorBlendAttachmentState-constantAlphaColorBlendFactors-04455  
  If the `VK_KHR_portability_subset` extension is enabled, and [VkPhysicalDevicePortabilitySubsetFeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevicePortabilitySubsetFeaturesKHR.html)::`constantAlphaColorBlendFactors` is `VK_FALSE`, `dstColorBlendFactor` **must** not be `VK_BLEND_FACTOR_CONSTANT_ALPHA` or `VK_BLEND_FACTOR_ONE_MINUS_CONSTANT_ALPHA`

Valid Usage (Implicit)

- [](#VUID-VkPipelineColorBlendAttachmentState-srcColorBlendFactor-parameter)VUID-VkPipelineColorBlendAttachmentState-srcColorBlendFactor-parameter  
  `srcColorBlendFactor` **must** be a valid [VkBlendFactor](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBlendFactor.html) value
- [](#VUID-VkPipelineColorBlendAttachmentState-dstColorBlendFactor-parameter)VUID-VkPipelineColorBlendAttachmentState-dstColorBlendFactor-parameter  
  `dstColorBlendFactor` **must** be a valid [VkBlendFactor](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBlendFactor.html) value
- [](#VUID-VkPipelineColorBlendAttachmentState-colorBlendOp-parameter)VUID-VkPipelineColorBlendAttachmentState-colorBlendOp-parameter  
  `colorBlendOp` **must** be a valid [VkBlendOp](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBlendOp.html) value
- [](#VUID-VkPipelineColorBlendAttachmentState-srcAlphaBlendFactor-parameter)VUID-VkPipelineColorBlendAttachmentState-srcAlphaBlendFactor-parameter  
  `srcAlphaBlendFactor` **must** be a valid [VkBlendFactor](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBlendFactor.html) value
- [](#VUID-VkPipelineColorBlendAttachmentState-dstAlphaBlendFactor-parameter)VUID-VkPipelineColorBlendAttachmentState-dstAlphaBlendFactor-parameter  
  `dstAlphaBlendFactor` **must** be a valid [VkBlendFactor](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBlendFactor.html) value
- [](#VUID-VkPipelineColorBlendAttachmentState-alphaBlendOp-parameter)VUID-VkPipelineColorBlendAttachmentState-alphaBlendOp-parameter  
  `alphaBlendOp` **must** be a valid [VkBlendOp](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBlendOp.html) value
- [](#VUID-VkPipelineColorBlendAttachmentState-colorWriteMask-parameter)VUID-VkPipelineColorBlendAttachmentState-colorWriteMask-parameter  
  `colorWriteMask` **must** be a valid combination of [VkColorComponentFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkColorComponentFlagBits.html) values

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkBlendFactor](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBlendFactor.html), [VkBlendOp](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBlendOp.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkColorComponentFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkColorComponentFlags.html), [VkPipelineColorBlendStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineColorBlendStateCreateInfo.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPipelineColorBlendAttachmentState).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0