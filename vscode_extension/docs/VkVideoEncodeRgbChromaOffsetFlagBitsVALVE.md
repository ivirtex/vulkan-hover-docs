# VkVideoEncodeRgbChromaOffsetFlagBitsVALVE(3) Manual Page

## Name

VkVideoEncodeRgbChromaOffsetFlagBitsVALVE - Position of downsampled chroma samples for encode RGB conversion



## [](#_c_specification)C Specification

The model-converted values are chroma subsampled and quantized, according to the `chromaSubsampling`, `lumaBitDepth` and `chromaBitDepth` values specified by the bound video session.

The [VkVideoEncodeRgbChromaOffsetFlagBitsVALVE](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeRgbChromaOffsetFlagBitsVALVE.html) enum defines the location of downsampled chroma component samples relative to the luma samples for video encode R′G′B′ conversion, and is defined as:

```c++
// Provided by VK_VALVE_video_encode_rgb_conversion
typedef enum VkVideoEncodeRgbChromaOffsetFlagBitsVALVE {
    VK_VIDEO_ENCODE_RGB_CHROMA_OFFSET_COSITED_EVEN_BIT_VALVE = 0x00000001,
    VK_VIDEO_ENCODE_RGB_CHROMA_OFFSET_MIDPOINT_BIT_VALVE = 0x00000002,
} VkVideoEncodeRgbChromaOffsetFlagBitsVALVE;
```

## [](#_description)Description

- `VK_VIDEO_ENCODE_RGB_CHROMA_OFFSET_COSITED_EVEN_BIT_VALVE` specifies that downsampled chroma samples are aligned with luma samples with even coordinates.
- `VK_VIDEO_ENCODE_RGB_CHROMA_OFFSET_MIDPOINT_BIT_VALVE` specifies that downsampled chroma samples are located half way between each even luma sample and the nearest higher odd luma sample.

The output location of downsampled chroma components are specified by the `xChromaOffset` and `yChromaOffset` values of the [VkVideoEncodeSessionRgbConversionCreateInfoVALVE](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeSessionRgbConversionCreateInfoVALVE.html) structure:

Chroma subsampling is described in more detail in the [Chroma Reconstruction](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#textures-chroma-reconstruction) section.

## [](#_see_also)See Also

[VK\_VALVE\_video\_encode\_rgb\_conversion](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VALVE_video_encode_rgb_conversion.html), [VkVideoEncodeRgbChromaOffsetFlagsVALVE](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeRgbChromaOffsetFlagsVALVE.html), [VkVideoEncodeSessionRgbConversionCreateInfoVALVE](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeSessionRgbConversionCreateInfoVALVE.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoEncodeRgbChromaOffsetFlagBitsVALVE).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0