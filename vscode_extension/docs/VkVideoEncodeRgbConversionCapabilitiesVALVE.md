# VkVideoEncodeRgbConversionCapabilitiesVALVE(3) Manual Page

## Name

VkVideoEncodeRgbConversionCapabilitiesVALVE - Structure describing video encode rgb conversion capabilities for a video profile



## [](#_c_specification)C Specification

When calling [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html) with `pVideoProfile->videoCodecOperation` specifying an encode operation, the [VkVideoEncodeRgbConversionCapabilitiesVALVE](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeRgbConversionCapabilitiesVALVE.html) structure **can** be included in the `pNext` chain of the [VkVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoCapabilitiesKHR.html) structure to retrieve capabilities specific to video encode R′G′B′ conversion.

The `VkVideoEncodeRgbConversionCapabilitiesVALVE` structure is defined as:

```c++
// Provided by VK_VALVE_video_encode_rgb_conversion
typedef struct VkVideoEncodeRgbConversionCapabilitiesVALVE {
    VkStructureType                               sType;
    void*                                         pNext;
    VkVideoEncodeRgbModelConversionFlagsVALVE     rgbModels;
    VkVideoEncodeRgbRangeCompressionFlagsVALVE    rgbRanges;
    VkVideoEncodeRgbChromaOffsetFlagsVALVE        xChromaOffsets;
    VkVideoEncodeRgbChromaOffsetFlagsVALVE        yChromaOffsets;
} VkVideoEncodeRgbConversionCapabilitiesVALVE;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `rgbModels` is a bitmask of [VkVideoEncodeRgbModelConversionFlagBitsVALVE](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeRgbModelConversionFlagBitsVALVE.html) describing supported [model conversions](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-rgb-conversion-model-conversion) for video encode R′G′B′ conversion.
- `rgbRanges` is a bitmask of [VkVideoEncodeRgbRangeCompressionFlagBitsVALVE](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeRgbRangeCompressionFlagBitsVALVE.html) describing supported [range compressions](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-rgb-conversion-range-compression) for video encode R′G′B′ conversion.
- `xChromaOffsets` is a bitmask of [VkVideoEncodeRgbChromaOffsetFlagBitsVALVE](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeRgbChromaOffsetFlagBitsVALVE.html) describing supported offsets of the output location of the downsampled chroma component on the X axis for video encode R′G′B′ conversion.
- `yChromaOffsets` is a bitmask of [VkVideoEncodeRgbChromaOffsetFlagBitsVALVE](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeRgbChromaOffsetFlagBitsVALVE.html) describing supported offsets of the output location of the downsampled chroma component on the Y axis for video encode R′G′B′ conversion.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkVideoEncodeRgbConversionCapabilitiesVALVE-sType-sType)VUID-VkVideoEncodeRgbConversionCapabilitiesVALVE-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_ENCODE_RGB_CONVERSION_CAPABILITIES_VALVE`

## [](#_see_also)See Also

[VK\_VALVE\_video\_encode\_rgb\_conversion](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VALVE_video_encode_rgb_conversion.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkVideoEncodeRgbChromaOffsetFlagsVALVE](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeRgbChromaOffsetFlagsVALVE.html), [VkVideoEncodeRgbModelConversionFlagsVALVE](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeRgbModelConversionFlagsVALVE.html), [VkVideoEncodeRgbRangeCompressionFlagsVALVE](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeRgbRangeCompressionFlagsVALVE.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoEncodeRgbConversionCapabilitiesVALVE).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0