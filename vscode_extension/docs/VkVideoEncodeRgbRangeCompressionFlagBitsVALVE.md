# VkVideoEncodeRgbRangeCompressionFlagBitsVALVE(3) Manual Page

## Name

VkVideoEncodeRgbRangeCompressionFlagBitsVALVE - Range compression operation to perform for encode rgb conversion



## [](#_c_specification)C Specification

The video encode R′G′B′ range compression to be applied to color component values of the [encode input picture](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-input-picture) before video coding is defined by the `rgbRange` member of the [VkVideoEncodeSessionRgbConversionCreateInfoVALVE](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeSessionRgbConversionCreateInfoVALVE.html) structure.

The [VkVideoEncodeRgbRangeCompressionFlagBitsVALVE](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeRgbRangeCompressionFlagBitsVALVE.html) enum describes whether color components are encoded using the full range of numerical values or whether values are reserved for headroom and foot room:

```c++
// Provided by VK_VALVE_video_encode_rgb_conversion
typedef enum VkVideoEncodeRgbRangeCompressionFlagBitsVALVE {
    VK_VIDEO_ENCODE_RGB_RANGE_COMPRESSION_FULL_RANGE_BIT_VALVE = 0x00000001,
    VK_VIDEO_ENCODE_RGB_RANGE_COMPRESSION_NARROW_RANGE_BIT_VALVE = 0x00000002,
} VkVideoEncodeRgbRangeCompressionFlagBitsVALVE;
```

## [](#_description)Description

- `VK_VIDEO_ENCODE_RGB_RANGE_COMPRESSION_FULL_RANGE_BIT_VALVE` specifies the following transformations are applied:
  
  Y′CB​CR​​=Crgba′​\[G]=Crgba′​\[B]−(2n)−12(n−1)​=Crgba′​\[R]−(2n)−12(n−1)​​
  
  Note
  
  These formulae correspond to the “full range” encoding in the “Quantization schemes” chapter of the [Khronos Data Format Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#data-format).
  
  Should any future amendments be made to the ITU specifications from which these equations are derived, the formulae used by Vulkan **may** also be updated to maintain parity.
- `VK_VIDEO_ENCODE_RGB_RANGE_COMPRESSION_NARROW_RANGE_BIT_VALVE` specifies the following transformations are applied:
  
  Y′CB​CR​​=219×2n−8Crgba′​\[G]×(2n−1)−16×2n−8​=224×2n−8Crgba′​\[B]×(2n−1)−128×2n−8​=224×2n−8Crgba′​\[R]×(2n−1)−128×2n−8​​
  
  Note
  
  These formulae correspond to the “narrow range” encoding in the “Quantization schemes” chapter of the [Khronos Data Format Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#data-format).
  
  Unlike [sampler Y′CBCR range expansion](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#textures-sampler-YCbCr-conversion-rangeexpand), no precision guarantees are made for video encode R′G′B′ range compression.
- *n* is the bit-depth of the components in the bound video session’s `pictureFormat`.

Note

Video encode R′G′B′ range compression transformations have the inverse definition of the [sampler Y′CBCR range expansion](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#textures-sampler-YCbCr-conversion-rangeexpand) transformations.

## [](#_see_also)See Also

[VK\_VALVE\_video\_encode\_rgb\_conversion](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VALVE_video_encode_rgb_conversion.html), [VkVideoEncodeRgbRangeCompressionFlagsVALVE](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeRgbRangeCompressionFlagsVALVE.html), [VkVideoEncodeSessionRgbConversionCreateInfoVALVE](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeSessionRgbConversionCreateInfoVALVE.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoEncodeRgbRangeCompressionFlagBitsVALVE).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0