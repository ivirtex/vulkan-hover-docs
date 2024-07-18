# VkColorBlendEquationEXT(3) Manual Page

## Name

VkColorBlendEquationEXT - Structure specifying the color blend factors
and operations for an attachment



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkColorBlendEquationEXT` structure is defined as:

``` c
// Provided by VK_EXT_extended_dynamic_state3, VK_EXT_shader_object
typedef struct VkColorBlendEquationEXT {
    VkBlendFactor    srcColorBlendFactor;
    VkBlendFactor    dstColorBlendFactor;
    VkBlendOp        colorBlendOp;
    VkBlendFactor    srcAlphaBlendFactor;
    VkBlendFactor    dstAlphaBlendFactor;
    VkBlendOp        alphaBlendOp;
} VkColorBlendEquationEXT;
```

## <a href="#_members" class="anchor"></a>Members

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

- `alphaBlendOp` selects which blend operation is use to calculate the
  alpha values to write to the color attachment.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkColorBlendEquationEXT-dualSrcBlend-07357"
  id="VUID-VkColorBlendEquationEXT-dualSrcBlend-07357"></a>
  VUID-VkColorBlendEquationEXT-dualSrcBlend-07357  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-dualSrcBlend"
  target="_blank" rel="noopener"><code>dualSrcBlend</code></a> feature
  is not enabled, `srcColorBlendFactor` **must** not be
  `VK_BLEND_FACTOR_SRC1_COLOR`, `VK_BLEND_FACTOR_ONE_MINUS_SRC1_COLOR`,
  `VK_BLEND_FACTOR_SRC1_ALPHA`, or
  `VK_BLEND_FACTOR_ONE_MINUS_SRC1_ALPHA`

- <a href="#VUID-VkColorBlendEquationEXT-dualSrcBlend-07358"
  id="VUID-VkColorBlendEquationEXT-dualSrcBlend-07358"></a>
  VUID-VkColorBlendEquationEXT-dualSrcBlend-07358  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-dualSrcBlend"
  target="_blank" rel="noopener"><code>dualSrcBlend</code></a> feature
  is not enabled, `dstColorBlendFactor` **must** not be
  `VK_BLEND_FACTOR_SRC1_COLOR`, `VK_BLEND_FACTOR_ONE_MINUS_SRC1_COLOR`,
  `VK_BLEND_FACTOR_SRC1_ALPHA`, or
  `VK_BLEND_FACTOR_ONE_MINUS_SRC1_ALPHA`

- <a href="#VUID-VkColorBlendEquationEXT-dualSrcBlend-07359"
  id="VUID-VkColorBlendEquationEXT-dualSrcBlend-07359"></a>
  VUID-VkColorBlendEquationEXT-dualSrcBlend-07359  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-dualSrcBlend"
  target="_blank" rel="noopener"><code>dualSrcBlend</code></a> feature
  is not enabled, `srcAlphaBlendFactor` **must** not be
  `VK_BLEND_FACTOR_SRC1_COLOR`, `VK_BLEND_FACTOR_ONE_MINUS_SRC1_COLOR`,
  `VK_BLEND_FACTOR_SRC1_ALPHA`, or
  `VK_BLEND_FACTOR_ONE_MINUS_SRC1_ALPHA`

- <a href="#VUID-VkColorBlendEquationEXT-dualSrcBlend-07360"
  id="VUID-VkColorBlendEquationEXT-dualSrcBlend-07360"></a>
  VUID-VkColorBlendEquationEXT-dualSrcBlend-07360  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-dualSrcBlend"
  target="_blank" rel="noopener"><code>dualSrcBlend</code></a> feature
  is not enabled, `dstAlphaBlendFactor` **must** not be
  `VK_BLEND_FACTOR_SRC1_COLOR`, `VK_BLEND_FACTOR_ONE_MINUS_SRC1_COLOR`,
  `VK_BLEND_FACTOR_SRC1_ALPHA`, or
  `VK_BLEND_FACTOR_ONE_MINUS_SRC1_ALPHA`

- <a href="#VUID-VkColorBlendEquationEXT-colorBlendOp-07361"
  id="VUID-VkColorBlendEquationEXT-colorBlendOp-07361"></a>
  VUID-VkColorBlendEquationEXT-colorBlendOp-07361  
  `colorBlendOp` and `alphaBlendOp` **must** not be
  `VK_BLEND_OP_ZERO_EXT`, `VK_BLEND_OP_SRC_EXT`, `VK_BLEND_OP_DST_EXT`,
  `VK_BLEND_OP_SRC_OVER_EXT`, `VK_BLEND_OP_DST_OVER_EXT`,
  `VK_BLEND_OP_SRC_IN_EXT`, `VK_BLEND_OP_DST_IN_EXT`,
  `VK_BLEND_OP_SRC_OUT_EXT`, `VK_BLEND_OP_DST_OUT_EXT`,
  `VK_BLEND_OP_SRC_ATOP_EXT`, `VK_BLEND_OP_DST_ATOP_EXT`,
  `VK_BLEND_OP_XOR_EXT`, `VK_BLEND_OP_MULTIPLY_EXT`,
  `VK_BLEND_OP_SCREEN_EXT`, `VK_BLEND_OP_OVERLAY_EXT`,
  `VK_BLEND_OP_DARKEN_EXT`, `VK_BLEND_OP_LIGHTEN_EXT`,
  `VK_BLEND_OP_COLORDODGE_EXT`, `VK_BLEND_OP_COLORBURN_EXT`,
  `VK_BLEND_OP_HARDLIGHT_EXT`, `VK_BLEND_OP_SOFTLIGHT_EXT`,
  `VK_BLEND_OP_DIFFERENCE_EXT`, `VK_BLEND_OP_EXCLUSION_EXT`,
  `VK_BLEND_OP_INVERT_EXT`, `VK_BLEND_OP_INVERT_RGB_EXT`,
  `VK_BLEND_OP_LINEARDODGE_EXT`, `VK_BLEND_OP_LINEARBURN_EXT`,
  `VK_BLEND_OP_VIVIDLIGHT_EXT`, `VK_BLEND_OP_LINEARLIGHT_EXT`,
  `VK_BLEND_OP_PINLIGHT_EXT`, `VK_BLEND_OP_HARDMIX_EXT`,
  `VK_BLEND_OP_HSL_HUE_EXT`, `VK_BLEND_OP_HSL_SATURATION_EXT`,
  `VK_BLEND_OP_HSL_COLOR_EXT`, `VK_BLEND_OP_HSL_LUMINOSITY_EXT`,
  `VK_BLEND_OP_PLUS_EXT`, `VK_BLEND_OP_PLUS_CLAMPED_EXT`,
  `VK_BLEND_OP_PLUS_CLAMPED_ALPHA_EXT`, `VK_BLEND_OP_PLUS_DARKER_EXT`,
  `VK_BLEND_OP_MINUS_EXT`, `VK_BLEND_OP_MINUS_CLAMPED_EXT`,
  `VK_BLEND_OP_CONTRAST_EXT`, `VK_BLEND_OP_INVERT_OVG_EXT`,
  `VK_BLEND_OP_RED_EXT`, `VK_BLEND_OP_GREEN_EXT`, or
  `VK_BLEND_OP_BLUE_EXT`

- <a
  href="#VUID-VkColorBlendEquationEXT-constantAlphaColorBlendFactors-07362"
  id="VUID-VkColorBlendEquationEXT-constantAlphaColorBlendFactors-07362"></a>
  VUID-VkColorBlendEquationEXT-constantAlphaColorBlendFactors-07362  
  If the [`VK_KHR_portability_subset`](VK_KHR_portability_subset.html)
  extension is enabled, and
  [VkPhysicalDevicePortabilitySubsetFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevicePortabilitySubsetFeaturesKHR.html)::`constantAlphaColorBlendFactors`
  is `VK_FALSE`, `srcColorBlendFactor` **must** not be
  `VK_BLEND_FACTOR_CONSTANT_ALPHA` or
  `VK_BLEND_FACTOR_ONE_MINUS_CONSTANT_ALPHA`

- <a
  href="#VUID-VkColorBlendEquationEXT-constantAlphaColorBlendFactors-07363"
  id="VUID-VkColorBlendEquationEXT-constantAlphaColorBlendFactors-07363"></a>
  VUID-VkColorBlendEquationEXT-constantAlphaColorBlendFactors-07363  
  If the [`VK_KHR_portability_subset`](VK_KHR_portability_subset.html)
  extension is enabled, and
  [VkPhysicalDevicePortabilitySubsetFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevicePortabilitySubsetFeaturesKHR.html)::`constantAlphaColorBlendFactors`
  is `VK_FALSE`, `dstColorBlendFactor` **must** not be
  `VK_BLEND_FACTOR_CONSTANT_ALPHA` or
  `VK_BLEND_FACTOR_ONE_MINUS_CONSTANT_ALPHA`

Valid Usage (Implicit)

- <a href="#VUID-VkColorBlendEquationEXT-srcColorBlendFactor-parameter"
  id="VUID-VkColorBlendEquationEXT-srcColorBlendFactor-parameter"></a>
  VUID-VkColorBlendEquationEXT-srcColorBlendFactor-parameter  
  `srcColorBlendFactor` **must** be a valid
  [VkBlendFactor](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBlendFactor.html) value

- <a href="#VUID-VkColorBlendEquationEXT-dstColorBlendFactor-parameter"
  id="VUID-VkColorBlendEquationEXT-dstColorBlendFactor-parameter"></a>
  VUID-VkColorBlendEquationEXT-dstColorBlendFactor-parameter  
  `dstColorBlendFactor` **must** be a valid
  [VkBlendFactor](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBlendFactor.html) value

- <a href="#VUID-VkColorBlendEquationEXT-colorBlendOp-parameter"
  id="VUID-VkColorBlendEquationEXT-colorBlendOp-parameter"></a>
  VUID-VkColorBlendEquationEXT-colorBlendOp-parameter  
  `colorBlendOp` **must** be a valid [VkBlendOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBlendOp.html) value

- <a href="#VUID-VkColorBlendEquationEXT-srcAlphaBlendFactor-parameter"
  id="VUID-VkColorBlendEquationEXT-srcAlphaBlendFactor-parameter"></a>
  VUID-VkColorBlendEquationEXT-srcAlphaBlendFactor-parameter  
  `srcAlphaBlendFactor` **must** be a valid
  [VkBlendFactor](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBlendFactor.html) value

- <a href="#VUID-VkColorBlendEquationEXT-dstAlphaBlendFactor-parameter"
  id="VUID-VkColorBlendEquationEXT-dstAlphaBlendFactor-parameter"></a>
  VUID-VkColorBlendEquationEXT-dstAlphaBlendFactor-parameter  
  `dstAlphaBlendFactor` **must** be a valid
  [VkBlendFactor](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBlendFactor.html) value

- <a href="#VUID-VkColorBlendEquationEXT-alphaBlendOp-parameter"
  id="VUID-VkColorBlendEquationEXT-alphaBlendOp-parameter"></a>
  VUID-VkColorBlendEquationEXT-alphaBlendOp-parameter  
  `alphaBlendOp` **must** be a valid [VkBlendOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBlendOp.html) value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_extended_dynamic_state3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_extended_dynamic_state3.html),
[VK_EXT_shader_object](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_shader_object.html),
[VkBlendFactor](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBlendFactor.html), [VkBlendOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBlendOp.html),
[vkCmdSetColorBlendEquationEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetColorBlendEquationEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkColorBlendEquationEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
