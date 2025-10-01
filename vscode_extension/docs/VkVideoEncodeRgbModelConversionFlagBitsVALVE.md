# VkVideoEncodeRgbModelConversionFlagBitsVALVE(3) Manual Page

## Name

VkVideoEncodeRgbModelConversionFlagBitsVALVE - Color model conversions for encode RGB conversion



## [](#_c_specification)C Specification

The range-compressed values are converted between color models, according to the color model conversion specified in the `rgbModel` member of the [VkVideoEncodeSessionRgbConversionCreateInfoVALVE](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeSessionRgbConversionCreateInfoVALVE.html) structure.

[VkVideoEncodeRgbModelConversionFlagBitsVALVE](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeRgbModelConversionFlagBitsVALVE.html) defines the conversion from the [encode input picture](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-input-picture-info)'s color model to the encode color model.

```c++
// Provided by VK_VALVE_video_encode_rgb_conversion
typedef enum VkVideoEncodeRgbModelConversionFlagBitsVALVE {
    VK_VIDEO_ENCODE_RGB_MODEL_CONVERSION_RGB_IDENTITY_BIT_VALVE = 0x00000001,
    VK_VIDEO_ENCODE_RGB_MODEL_CONVERSION_YCBCR_IDENTITY_BIT_VALVE = 0x00000002,
    VK_VIDEO_ENCODE_RGB_MODEL_CONVERSION_YCBCR_709_BIT_VALVE = 0x00000004,
    VK_VIDEO_ENCODE_RGB_MODEL_CONVERSION_YCBCR_601_BIT_VALVE = 0x00000008,
    VK_VIDEO_ENCODE_RGB_MODEL_CONVERSION_YCBCR_2020_BIT_VALVE = 0x00000010,
} VkVideoEncodeRgbModelConversionFlagBitsVALVE;
```

## [](#_description)Description

- `VK_VIDEO_ENCODE_RGB_MODEL_CONVERSION_RGB_IDENTITY_BIT_VALVE` specifies the color components are not modified by the color model conversion since they are assumed to represent the desired color model for video coding; R′G′B′ range compression is applied to the components.
- `VK_VIDEO_ENCODE_RGB_MODEL_CONVERSION_YCBCR_IDENTITY_BIT_VALVE` specifies the color components are not modified by the color model conversion are assumed to be treated as though in Y′CBCR form; video encode R′G′B′ range compression and video encode R′G′B′ chroma subsampling is also ignored.
- `VK_VIDEO_ENCODE_RGB_MODEL_CONVERSION_YCBCR_709_BIT_VALVE` specifies the color components are transformed from an R′G′B′ representation to a Y′CBCR representation as described in the “BT.709 Y′CBCR conversion” section of the [Khronos Data Format Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#data-format).
- `VK_VIDEO_ENCODE_RGB_MODEL_CONVERSION_YCBCR_601_BIT_VALVE` specifies the color components are transformed from an R′G′B′ representation to a Y′CBCR representation as described in the “BT.601 Y′CBCR conversion” section of the [Khronos Data Format Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#data-format).
- `VK_VIDEO_ENCODE_RGB_MODEL_CONVERSION_YCBCR_2020_BIT_VALVE` specifies the color components are transformed from an R′G′B′ representation to a Y′CBCR representation as described in the “BT.2020 Y′CBCR conversion” section of the [Khronos Data Format Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#data-format).

Note

Video encode R′G′B′ model conversion transformations have the inverse definition of [sampler Y′CBCR model conversion](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#textures-sampler-YCbCr-conversion-modelconversion) transformations.

Note

The video encode R′G′B′ model conversion step does not apply any transfer function, only converting from R′G′B′ to Y′CBCR using the primaries of the specified `rgbModel` color model.

## [](#_see_also)See Also

[VK\_VALVE\_video\_encode\_rgb\_conversion](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VALVE_video_encode_rgb_conversion.html), [VkVideoEncodeRgbModelConversionFlagsVALVE](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeRgbModelConversionFlagsVALVE.html), [VkVideoEncodeSessionRgbConversionCreateInfoVALVE](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeSessionRgbConversionCreateInfoVALVE.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoEncodeRgbModelConversionFlagBitsVALVE).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0