# VkColorBlendAdvancedEXT(3) Manual Page

## Name

VkColorBlendAdvancedEXT - Structure specifying the advanced blend operation parameters for an attachment



## [](#_c_specification)C Specification

The `VkColorBlendAdvancedEXT` structure is defined as:

```c++
// Provided by VK_EXT_extended_dynamic_state3, VK_EXT_shader_object
typedef struct VkColorBlendAdvancedEXT {
    VkBlendOp            advancedBlendOp;
    VkBool32             srcPremultiplied;
    VkBool32             dstPremultiplied;
    VkBlendOverlapEXT    blendOverlap;
    VkBool32             clampResults;
} VkColorBlendAdvancedEXT;
```

## [](#_members)Members

- `advancedBlendOp` selects which blend operation is used to calculate the RGB values to write to the color attachment.
- `srcPremultiplied` specifies whether the source color of the blend operation is treated as premultiplied.
- `dstPremultiplied` specifies whether the destination color of the blend operation is treated as premultiplied.
- `blendOverlap` is a [VkBlendOverlapEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBlendOverlapEXT.html) value specifying how the source and destination sample’s coverage is correlated.
- `clampResults` specifies that results **must** be clamped to the \[0,1] range before writing to the attachment, which is useful when the attachment format is not normalized fixed-point.

## [](#_description)Description

Valid Usage

- [](#VUID-VkColorBlendAdvancedEXT-srcPremultiplied-07505)VUID-VkColorBlendAdvancedEXT-srcPremultiplied-07505  
  If the [non-premultiplied source color](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-advancedBlendNonPremultipliedSrcColor) property is not supported, `srcPremultiplied` **must** be `VK_TRUE`
- [](#VUID-VkColorBlendAdvancedEXT-dstPremultiplied-07506)VUID-VkColorBlendAdvancedEXT-dstPremultiplied-07506  
  If the [non-premultiplied destination color](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-advancedBlendNonPremultipliedDstColor) property is not supported, `dstPremultiplied` **must** be `VK_TRUE`
- [](#VUID-VkColorBlendAdvancedEXT-blendOverlap-07507)VUID-VkColorBlendAdvancedEXT-blendOverlap-07507  
  If the [correlated overlap](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-advancedBlendCorrelatedOverlap) property is not supported, `blendOverlap` **must** be `VK_BLEND_OVERLAP_UNCORRELATED_EXT`

Valid Usage (Implicit)

- [](#VUID-VkColorBlendAdvancedEXT-advancedBlendOp-parameter)VUID-VkColorBlendAdvancedEXT-advancedBlendOp-parameter  
  `advancedBlendOp` **must** be a valid [VkBlendOp](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBlendOp.html) value
- [](#VUID-VkColorBlendAdvancedEXT-blendOverlap-parameter)VUID-VkColorBlendAdvancedEXT-blendOverlap-parameter  
  `blendOverlap` **must** be a valid [VkBlendOverlapEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBlendOverlapEXT.html) value

## [](#_see_also)See Also

[VK\_EXT\_extended\_dynamic\_state3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_extended_dynamic_state3.html), [VK\_EXT\_shader\_object](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_shader_object.html), [VkBlendOp](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBlendOp.html), [VkBlendOverlapEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBlendOverlapEXT.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [vkCmdSetColorBlendAdvancedEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetColorBlendAdvancedEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkColorBlendAdvancedEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0