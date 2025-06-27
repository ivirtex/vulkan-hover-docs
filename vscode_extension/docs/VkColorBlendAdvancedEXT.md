# VkColorBlendAdvancedEXT(3) Manual Page

## Name

VkColorBlendAdvancedEXT - Structure specifying the advanced blend
operation parameters for an attachment



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkColorBlendAdvancedEXT` structure is defined as:

``` c
// Provided by VK_EXT_extended_dynamic_state3, VK_EXT_shader_object
typedef struct VkColorBlendAdvancedEXT {
    VkBlendOp            advancedBlendOp;
    VkBool32             srcPremultiplied;
    VkBool32             dstPremultiplied;
    VkBlendOverlapEXT    blendOverlap;
    VkBool32             clampResults;
} VkColorBlendAdvancedEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `advancedBlendOp` selects which blend operation is used to calculate
  the RGB values to write to the color attachment.

- `srcPremultiplied` specifies whether the source color of the blend
  operation is treated as premultiplied.

- `dstPremultiplied` specifies whether the destination color of the
  blend operation is treated as premultiplied.

- `blendOverlap` is a [VkBlendOverlapEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBlendOverlapEXT.html) value
  specifying how the source and destination sample’s coverage is
  correlated.

- `clampResults` specifies the results must be clamped to the \[0,1\]
  range before writing to the attachment, which is useful when the
  attachment format is not normalized fixed-point.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkColorBlendAdvancedEXT-srcPremultiplied-07505"
  id="VUID-VkColorBlendAdvancedEXT-srcPremultiplied-07505"></a>
  VUID-VkColorBlendAdvancedEXT-srcPremultiplied-07505  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-advancedBlendNonPremultipliedSrcColor"
  target="_blank" rel="noopener">non-premultiplied source color</a>
  property is not supported, `srcPremultiplied` **must** be `VK_TRUE`

- <a href="#VUID-VkColorBlendAdvancedEXT-dstPremultiplied-07506"
  id="VUID-VkColorBlendAdvancedEXT-dstPremultiplied-07506"></a>
  VUID-VkColorBlendAdvancedEXT-dstPremultiplied-07506  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-advancedBlendNonPremultipliedDstColor"
  target="_blank" rel="noopener">non-premultiplied destination color</a>
  property is not supported, `dstPremultiplied` **must** be `VK_TRUE`

- <a href="#VUID-VkColorBlendAdvancedEXT-blendOverlap-07507"
  id="VUID-VkColorBlendAdvancedEXT-blendOverlap-07507"></a>
  VUID-VkColorBlendAdvancedEXT-blendOverlap-07507  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-advancedBlendCorrelatedOverlap"
  target="_blank" rel="noopener">correlated overlap</a> property is not
  supported, `blendOverlap` **must** be
  `VK_BLEND_OVERLAP_UNCORRELATED_EXT`

Valid Usage (Implicit)

- <a href="#VUID-VkColorBlendAdvancedEXT-advancedBlendOp-parameter"
  id="VUID-VkColorBlendAdvancedEXT-advancedBlendOp-parameter"></a>
  VUID-VkColorBlendAdvancedEXT-advancedBlendOp-parameter  
  `advancedBlendOp` **must** be a valid [VkBlendOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBlendOp.html)
  value

- <a href="#VUID-VkColorBlendAdvancedEXT-blendOverlap-parameter"
  id="VUID-VkColorBlendAdvancedEXT-blendOverlap-parameter"></a>
  VUID-VkColorBlendAdvancedEXT-blendOverlap-parameter  
  `blendOverlap` **must** be a valid
  [VkBlendOverlapEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBlendOverlapEXT.html) value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_extended_dynamic_state3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_extended_dynamic_state3.html),
[VK_EXT_shader_object](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_shader_object.html),
[VkBlendOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBlendOp.html),
[VkBlendOverlapEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBlendOverlapEXT.html), [VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html),
[vkCmdSetColorBlendAdvancedEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetColorBlendAdvancedEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkColorBlendAdvancedEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
