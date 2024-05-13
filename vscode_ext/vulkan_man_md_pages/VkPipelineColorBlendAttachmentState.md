# VkPipelineColorBlendAttachmentState(3) Manual Page

## Name

VkPipelineColorBlendAttachmentState - Structure specifying a pipeline
color blend attachment state



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPipelineColorBlendAttachmentState` structure is defined as:

``` c
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

## <a href="#_members" class="anchor"></a>Members

- `blendEnable` controls whether blending is enabled for the
  corresponding color attachment. If blending is not enabled, the source
  fragment’s color for that attachment is passed through unmodified.

- `srcColorBlendFactor` selects which blend factor is used to determine
  the source factors (S<sub>r</sub>,S<sub>g</sub>,S<sub>b</sub>).

- `dstColorBlendFactor` selects which blend factor is used to determine
  the destination factors (D<sub>r</sub>,D<sub>g</sub>,D<sub>b</sub>).

- `colorBlendOp` selects which blend operation is used to calculate the
  RGB values to write to the color attachment.

- `srcAlphaBlendFactor` selects which blend factor is used to determine
  the source factor S<sub>a</sub>.

- `dstAlphaBlendFactor` selects which blend factor is used to determine
  the destination factor D<sub>a</sub>.

- `alphaBlendOp` selects which blend operation is used to calculate the
  alpha values to write to the color attachment.

- `colorWriteMask` is a bitmask of
  [VkColorComponentFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkColorComponentFlagBits.html) specifying
  which of the R, G, B, and/or A components are enabled for writing, as
  described for the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#framebuffer-color-write-mask"
  target="_blank" rel="noopener">Color Write Mask</a>.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a
  href="#VUID-VkPipelineColorBlendAttachmentState-srcColorBlendFactor-00608"
  id="VUID-VkPipelineColorBlendAttachmentState-srcColorBlendFactor-00608"></a>
  VUID-VkPipelineColorBlendAttachmentState-srcColorBlendFactor-00608  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-dualSrcBlend"
  target="_blank" rel="noopener"><code>dualSrcBlend</code></a> feature
  is not enabled, `srcColorBlendFactor` **must** not be
  `VK_BLEND_FACTOR_SRC1_COLOR`, `VK_BLEND_FACTOR_ONE_MINUS_SRC1_COLOR`,
  `VK_BLEND_FACTOR_SRC1_ALPHA`, or
  `VK_BLEND_FACTOR_ONE_MINUS_SRC1_ALPHA`

- <a
  href="#VUID-VkPipelineColorBlendAttachmentState-dstColorBlendFactor-00609"
  id="VUID-VkPipelineColorBlendAttachmentState-dstColorBlendFactor-00609"></a>
  VUID-VkPipelineColorBlendAttachmentState-dstColorBlendFactor-00609  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-dualSrcBlend"
  target="_blank" rel="noopener"><code>dualSrcBlend</code></a> feature
  is not enabled, `dstColorBlendFactor` **must** not be
  `VK_BLEND_FACTOR_SRC1_COLOR`, `VK_BLEND_FACTOR_ONE_MINUS_SRC1_COLOR`,
  `VK_BLEND_FACTOR_SRC1_ALPHA`, or
  `VK_BLEND_FACTOR_ONE_MINUS_SRC1_ALPHA`

- <a
  href="#VUID-VkPipelineColorBlendAttachmentState-srcAlphaBlendFactor-00610"
  id="VUID-VkPipelineColorBlendAttachmentState-srcAlphaBlendFactor-00610"></a>
  VUID-VkPipelineColorBlendAttachmentState-srcAlphaBlendFactor-00610  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-dualSrcBlend"
  target="_blank" rel="noopener"><code>dualSrcBlend</code></a> feature
  is not enabled, `srcAlphaBlendFactor` **must** not be
  `VK_BLEND_FACTOR_SRC1_COLOR`, `VK_BLEND_FACTOR_ONE_MINUS_SRC1_COLOR`,
  `VK_BLEND_FACTOR_SRC1_ALPHA`, or
  `VK_BLEND_FACTOR_ONE_MINUS_SRC1_ALPHA`

- <a
  href="#VUID-VkPipelineColorBlendAttachmentState-dstAlphaBlendFactor-00611"
  id="VUID-VkPipelineColorBlendAttachmentState-dstAlphaBlendFactor-00611"></a>
  VUID-VkPipelineColorBlendAttachmentState-dstAlphaBlendFactor-00611  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-dualSrcBlend"
  target="_blank" rel="noopener"><code>dualSrcBlend</code></a> feature
  is not enabled, `dstAlphaBlendFactor` **must** not be
  `VK_BLEND_FACTOR_SRC1_COLOR`, `VK_BLEND_FACTOR_ONE_MINUS_SRC1_COLOR`,
  `VK_BLEND_FACTOR_SRC1_ALPHA`, or
  `VK_BLEND_FACTOR_ONE_MINUS_SRC1_ALPHA`

- <a href="#VUID-VkPipelineColorBlendAttachmentState-colorBlendOp-01406"
  id="VUID-VkPipelineColorBlendAttachmentState-colorBlendOp-01406"></a>
  VUID-VkPipelineColorBlendAttachmentState-colorBlendOp-01406  
  If either of `colorBlendOp` or `alphaBlendOp` is an <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#framebuffer-blend-advanced"
  target="_blank" rel="noopener">advanced blend operation</a>, then
  `colorBlendOp` **must** equal `alphaBlendOp`

- <a
  href="#VUID-VkPipelineColorBlendAttachmentState-advancedBlendIndependentBlend-01407"
  id="VUID-VkPipelineColorBlendAttachmentState-advancedBlendIndependentBlend-01407"></a>
  VUID-VkPipelineColorBlendAttachmentState-advancedBlendIndependentBlend-01407  
  If
  [VkPhysicalDeviceBlendOperationAdvancedPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceBlendOperationAdvancedPropertiesEXT.html)::`advancedBlendIndependentBlend`
  is `VK_FALSE` and `colorBlendOp` is an <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#framebuffer-blend-advanced"
  target="_blank" rel="noopener">advanced blend operation</a>, then
  `colorBlendOp` **must** be the same for all attachments

- <a
  href="#VUID-VkPipelineColorBlendAttachmentState-advancedBlendIndependentBlend-01408"
  id="VUID-VkPipelineColorBlendAttachmentState-advancedBlendIndependentBlend-01408"></a>
  VUID-VkPipelineColorBlendAttachmentState-advancedBlendIndependentBlend-01408  
  If
  [VkPhysicalDeviceBlendOperationAdvancedPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceBlendOperationAdvancedPropertiesEXT.html)::`advancedBlendIndependentBlend`
  is `VK_FALSE` and `alphaBlendOp` is an <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#framebuffer-blend-advanced"
  target="_blank" rel="noopener">advanced blend operation</a>, then
  `alphaBlendOp` **must** be the same for all attachments

- <a
  href="#VUID-VkPipelineColorBlendAttachmentState-advancedBlendAllOperations-01409"
  id="VUID-VkPipelineColorBlendAttachmentState-advancedBlendAllOperations-01409"></a>
  VUID-VkPipelineColorBlendAttachmentState-advancedBlendAllOperations-01409  
  If
  [VkPhysicalDeviceBlendOperationAdvancedPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceBlendOperationAdvancedPropertiesEXT.html)::`advancedBlendAllOperations`
  is `VK_FALSE`, then `colorBlendOp` **must** not be
  `VK_BLEND_OP_ZERO_EXT`, `VK_BLEND_OP_SRC_EXT`, `VK_BLEND_OP_DST_EXT`,
  `VK_BLEND_OP_SRC_OVER_EXT`, `VK_BLEND_OP_DST_OVER_EXT`,
  `VK_BLEND_OP_SRC_IN_EXT`, `VK_BLEND_OP_DST_IN_EXT`,
  `VK_BLEND_OP_SRC_OUT_EXT`, `VK_BLEND_OP_DST_OUT_EXT`,
  `VK_BLEND_OP_SRC_ATOP_EXT`, `VK_BLEND_OP_DST_ATOP_EXT`,
  `VK_BLEND_OP_XOR_EXT`, `VK_BLEND_OP_INVERT_EXT`,
  `VK_BLEND_OP_INVERT_RGB_EXT`, `VK_BLEND_OP_LINEARDODGE_EXT`,
  `VK_BLEND_OP_LINEARBURN_EXT`, `VK_BLEND_OP_VIVIDLIGHT_EXT`,
  `VK_BLEND_OP_LINEARLIGHT_EXT`, `VK_BLEND_OP_PINLIGHT_EXT`,
  `VK_BLEND_OP_HARDMIX_EXT`, `VK_BLEND_OP_PLUS_EXT`,
  `VK_BLEND_OP_PLUS_CLAMPED_EXT`, `VK_BLEND_OP_PLUS_CLAMPED_ALPHA_EXT`,
  `VK_BLEND_OP_PLUS_DARKER_EXT`, `VK_BLEND_OP_MINUS_EXT`,
  `VK_BLEND_OP_MINUS_CLAMPED_EXT`, `VK_BLEND_OP_CONTRAST_EXT`,
  `VK_BLEND_OP_INVERT_OVG_EXT`, `VK_BLEND_OP_RED_EXT`,
  `VK_BLEND_OP_GREEN_EXT`, or `VK_BLEND_OP_BLUE_EXT`

- <a href="#VUID-VkPipelineColorBlendAttachmentState-colorBlendOp-01410"
  id="VUID-VkPipelineColorBlendAttachmentState-colorBlendOp-01410"></a>
  VUID-VkPipelineColorBlendAttachmentState-colorBlendOp-01410  
  If `colorBlendOp` or `alphaBlendOp` is an <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#framebuffer-blend-advanced"
  target="_blank" rel="noopener">advanced blend operation</a>, then
  `colorAttachmentCount` of the subpass this pipeline is compiled
  against **must** be less than or equal to
  [VkPhysicalDeviceBlendOperationAdvancedPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceBlendOperationAdvancedPropertiesEXT.html)::`advancedBlendMaxColorAttachments`

- <a
  href="#VUID-VkPipelineColorBlendAttachmentState-constantAlphaColorBlendFactors-04454"
  id="VUID-VkPipelineColorBlendAttachmentState-constantAlphaColorBlendFactors-04454"></a>
  VUID-VkPipelineColorBlendAttachmentState-constantAlphaColorBlendFactors-04454  
  If the [`VK_KHR_portability_subset`](VK_KHR_portability_subset.html)
  extension is enabled, and
  [VkPhysicalDevicePortabilitySubsetFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevicePortabilitySubsetFeaturesKHR.html)::`constantAlphaColorBlendFactors`
  is `VK_FALSE`, `srcColorBlendFactor` **must** not be
  `VK_BLEND_FACTOR_CONSTANT_ALPHA` or
  `VK_BLEND_FACTOR_ONE_MINUS_CONSTANT_ALPHA`

- <a
  href="#VUID-VkPipelineColorBlendAttachmentState-constantAlphaColorBlendFactors-04455"
  id="VUID-VkPipelineColorBlendAttachmentState-constantAlphaColorBlendFactors-04455"></a>
  VUID-VkPipelineColorBlendAttachmentState-constantAlphaColorBlendFactors-04455  
  If the [`VK_KHR_portability_subset`](VK_KHR_portability_subset.html)
  extension is enabled, and
  [VkPhysicalDevicePortabilitySubsetFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevicePortabilitySubsetFeaturesKHR.html)::`constantAlphaColorBlendFactors`
  is `VK_FALSE`, `dstColorBlendFactor` **must** not be
  `VK_BLEND_FACTOR_CONSTANT_ALPHA` or
  `VK_BLEND_FACTOR_ONE_MINUS_CONSTANT_ALPHA`

Valid Usage (Implicit)

- <a
  href="#VUID-VkPipelineColorBlendAttachmentState-srcColorBlendFactor-parameter"
  id="VUID-VkPipelineColorBlendAttachmentState-srcColorBlendFactor-parameter"></a>
  VUID-VkPipelineColorBlendAttachmentState-srcColorBlendFactor-parameter  
  `srcColorBlendFactor` **must** be a valid
  [VkBlendFactor](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBlendFactor.html) value

- <a
  href="#VUID-VkPipelineColorBlendAttachmentState-dstColorBlendFactor-parameter"
  id="VUID-VkPipelineColorBlendAttachmentState-dstColorBlendFactor-parameter"></a>
  VUID-VkPipelineColorBlendAttachmentState-dstColorBlendFactor-parameter  
  `dstColorBlendFactor` **must** be a valid
  [VkBlendFactor](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBlendFactor.html) value

- <a
  href="#VUID-VkPipelineColorBlendAttachmentState-colorBlendOp-parameter"
  id="VUID-VkPipelineColorBlendAttachmentState-colorBlendOp-parameter"></a>
  VUID-VkPipelineColorBlendAttachmentState-colorBlendOp-parameter  
  `colorBlendOp` **must** be a valid [VkBlendOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBlendOp.html) value

- <a
  href="#VUID-VkPipelineColorBlendAttachmentState-srcAlphaBlendFactor-parameter"
  id="VUID-VkPipelineColorBlendAttachmentState-srcAlphaBlendFactor-parameter"></a>
  VUID-VkPipelineColorBlendAttachmentState-srcAlphaBlendFactor-parameter  
  `srcAlphaBlendFactor` **must** be a valid
  [VkBlendFactor](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBlendFactor.html) value

- <a
  href="#VUID-VkPipelineColorBlendAttachmentState-dstAlphaBlendFactor-parameter"
  id="VUID-VkPipelineColorBlendAttachmentState-dstAlphaBlendFactor-parameter"></a>
  VUID-VkPipelineColorBlendAttachmentState-dstAlphaBlendFactor-parameter  
  `dstAlphaBlendFactor` **must** be a valid
  [VkBlendFactor](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBlendFactor.html) value

- <a
  href="#VUID-VkPipelineColorBlendAttachmentState-alphaBlendOp-parameter"
  id="VUID-VkPipelineColorBlendAttachmentState-alphaBlendOp-parameter"></a>
  VUID-VkPipelineColorBlendAttachmentState-alphaBlendOp-parameter  
  `alphaBlendOp` **must** be a valid [VkBlendOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBlendOp.html) value

- <a
  href="#VUID-VkPipelineColorBlendAttachmentState-colorWriteMask-parameter"
  id="VUID-VkPipelineColorBlendAttachmentState-colorWriteMask-parameter"></a>
  VUID-VkPipelineColorBlendAttachmentState-colorWriteMask-parameter  
  `colorWriteMask` **must** be a valid combination of
  [VkColorComponentFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkColorComponentFlagBits.html) values

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkBlendFactor](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBlendFactor.html), [VkBlendOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBlendOp.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html),
[VkColorComponentFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkColorComponentFlags.html),
[VkPipelineColorBlendStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineColorBlendStateCreateInfo.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPipelineColorBlendAttachmentState"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
