# VkFragmentShadingRateNV(3) Manual Page

## Name

VkFragmentShadingRateNV - Enumeration with fragment shading rates



## <a href="#_c_specification" class="anchor"></a>C Specification

If the `fragmentShadingRateEnums` feature is enabled, fragment shading
rates may be specified using the
[VkFragmentShadingRateNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFragmentShadingRateNV.html) enumerated type
defined as:

``` c
// Provided by VK_NV_fragment_shading_rate_enums
typedef enum VkFragmentShadingRateNV {
    VK_FRAGMENT_SHADING_RATE_1_INVOCATION_PER_PIXEL_NV = 0,
    VK_FRAGMENT_SHADING_RATE_1_INVOCATION_PER_1X2_PIXELS_NV = 1,
    VK_FRAGMENT_SHADING_RATE_1_INVOCATION_PER_2X1_PIXELS_NV = 4,
    VK_FRAGMENT_SHADING_RATE_1_INVOCATION_PER_2X2_PIXELS_NV = 5,
    VK_FRAGMENT_SHADING_RATE_1_INVOCATION_PER_2X4_PIXELS_NV = 6,
    VK_FRAGMENT_SHADING_RATE_1_INVOCATION_PER_4X2_PIXELS_NV = 9,
    VK_FRAGMENT_SHADING_RATE_1_INVOCATION_PER_4X4_PIXELS_NV = 10,
    VK_FRAGMENT_SHADING_RATE_2_INVOCATIONS_PER_PIXEL_NV = 11,
    VK_FRAGMENT_SHADING_RATE_4_INVOCATIONS_PER_PIXEL_NV = 12,
    VK_FRAGMENT_SHADING_RATE_8_INVOCATIONS_PER_PIXEL_NV = 13,
    VK_FRAGMENT_SHADING_RATE_16_INVOCATIONS_PER_PIXEL_NV = 14,
    VK_FRAGMENT_SHADING_RATE_NO_INVOCATIONS_NV = 15,
} VkFragmentShadingRateNV;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_FRAGMENT_SHADING_RATE_1_INVOCATION_PER_PIXEL_NV` specifies a
  fragment size of 1x1 pixels.

- `VK_FRAGMENT_SHADING_RATE_1_INVOCATION_PER_1X2_PIXELS_NV` specifies a
  fragment size of 1x2 pixels.

- `VK_FRAGMENT_SHADING_RATE_1_INVOCATION_PER_2X1_PIXELS_NV` specifies a
  fragment size of 2x1 pixels.

- `VK_FRAGMENT_SHADING_RATE_1_INVOCATION_PER_2X2_PIXELS_NV` specifies a
  fragment size of 2x2 pixels.

- `VK_FRAGMENT_SHADING_RATE_1_INVOCATION_PER_2X4_PIXELS_NV` specifies a
  fragment size of 2x4 pixels.

- `VK_FRAGMENT_SHADING_RATE_1_INVOCATION_PER_4X2_PIXELS_NV` specifies a
  fragment size of 4x2 pixels.

- `VK_FRAGMENT_SHADING_RATE_1_INVOCATION_PER_4X4_PIXELS_NV` specifies a
  fragment size of 4x4 pixels.

- `VK_FRAGMENT_SHADING_RATE_2_INVOCATIONS_PER_PIXEL_NV` specifies a
  fragment size of 1x1 pixels, with two fragment shader invocations per
  fragment.

- `VK_FRAGMENT_SHADING_RATE_4_INVOCATIONS_PER_PIXEL_NV` specifies a
  fragment size of 1x1 pixels, with four fragment shader invocations per
  fragment.

- `VK_FRAGMENT_SHADING_RATE_8_INVOCATIONS_PER_PIXEL_NV` specifies a
  fragment size of 1x1 pixels, with eight fragment shader invocations
  per fragment.

- `VK_FRAGMENT_SHADING_RATE_16_INVOCATIONS_PER_PIXEL_NV` specifies a
  fragment size of 1x1 pixels, with sixteen fragment shader invocations
  per fragment.

- `VK_FRAGMENT_SHADING_RATE_NO_INVOCATIONS_NV` specifies that any
  portions of a primitive that use that shading rate should be discarded
  without invoking any fragment shader.

To use the shading rates
`VK_FRAGMENT_SHADING_RATE_2_INVOCATIONS_PER_PIXEL_NV`,
`VK_FRAGMENT_SHADING_RATE_4_INVOCATIONS_PER_PIXEL_NV`,
`VK_FRAGMENT_SHADING_RATE_8_INVOCATIONS_PER_PIXEL_NV`, and
`VK_FRAGMENT_SHADING_RATE_16_INVOCATIONS_PER_PIXEL_NV` as a pipeline,
primitive, or attachment shading rate, the
`supersampleFragmentShadingRates` feature **must** be enabled. To use
the shading rate `VK_FRAGMENT_SHADING_RATE_NO_INVOCATIONS_NV` as a
pipeline, primitive, or attachment shading rate, the
`noInvocationFragmentShadingRates` feature **must** be enabled.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_fragment_shading_rate_enums](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_fragment_shading_rate_enums.html),
[VkPipelineFragmentShadingRateEnumStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineFragmentShadingRateEnumStateCreateInfoNV.html),
[vkCmdSetFragmentShadingRateEnumNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetFragmentShadingRateEnumNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkFragmentShadingRateNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
