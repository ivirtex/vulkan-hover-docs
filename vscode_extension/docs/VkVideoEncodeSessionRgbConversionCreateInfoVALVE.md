# VkVideoEncodeSessionRgbConversionCreateInfoVALVE(3) Manual Page

## Name

VkVideoEncodeSessionRgbConversionCreateInfoVALVE - Structure specifying video encode RGB conversion parameters for a video session



## [](#_c_specification)C Specification

The `VkVideoEncodeSessionRgbConversionCreateInfoVALVE` structure is defined as:

```c++
// Provided by VK_VALVE_video_encode_rgb_conversion
typedef struct VkVideoEncodeSessionRgbConversionCreateInfoVALVE {
    VkStructureType                                  sType;
    const void*                                      pNext;
    VkVideoEncodeRgbModelConversionFlagBitsVALVE     rgbModel;
    VkVideoEncodeRgbRangeCompressionFlagBitsVALVE    rgbRange;
    VkVideoEncodeRgbChromaOffsetFlagBitsVALVE        xChromaOffset;
    VkVideoEncodeRgbChromaOffsetFlagBitsVALVE        yChromaOffset;
} VkVideoEncodeSessionRgbConversionCreateInfoVALVE;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `rgbModel` is the used [R′G′B′ model conversion](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-rgb-conversion-model-conversion) for the R′G′B′ conversion.
- `rgbRange` is the used [R′G′B′ range compression](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-rgb-conversion-range-compression) for the R′G′B′ conversion.
- `xChromaOffset` describes the output location of downsampled chroma components in the x dimension for the R′G′B′ conversion.
- `yChromaOffset` describes the output location of downsampled chroma components in the y dimension for the R′G′B′ conversion.

## [](#_description)Description

Valid Usage

- [](#VUID-VkVideoEncodeSessionRgbConversionCreateInfoVALVE-rgbModel-10930)VUID-VkVideoEncodeSessionRgbConversionCreateInfoVALVE-rgbModel-10930  
  `rgbModel` **must** only be a bit set in [VkVideoEncodeRgbConversionCapabilitiesVALVE](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeRgbConversionCapabilitiesVALVE.html)::`rgbModels` as returned by [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html) with `VkVideoEncodeRgbConversionCapabilitiesVALVE` in the `pNext` chain of `VkVideoCapabilitiesKHR` with the given `pVideoProfile`
- [](#VUID-VkVideoEncodeSessionRgbConversionCreateInfoVALVE-rgbRange-10931)VUID-VkVideoEncodeSessionRgbConversionCreateInfoVALVE-rgbRange-10931  
  `rgbRange` **must** only be a bit set in [VkVideoEncodeRgbConversionCapabilitiesVALVE](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeRgbConversionCapabilitiesVALVE.html)::`rgbRanges` as returned by [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html) with `VkVideoEncodeRgbConversionCapabilitiesVALVE` in the `pNext` chain of `VkVideoCapabilitiesKHR` with the given `pVideoProfile`
- [](#VUID-VkVideoEncodeSessionRgbConversionCreateInfoVALVE-xChromaOffset-10932)VUID-VkVideoEncodeSessionRgbConversionCreateInfoVALVE-xChromaOffset-10932  
  `xChromaOffset` **must** only be a bit set in [VkVideoEncodeRgbConversionCapabilitiesVALVE](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeRgbConversionCapabilitiesVALVE.html)::`xChromaOffsets` as returned by [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html) with `VkVideoEncodeRgbConversionCapabilitiesVALVE` in the `pNext` chain of `VkVideoCapabilitiesKHR` with the given `pVideoProfile`
- [](#VUID-VkVideoEncodeSessionRgbConversionCreateInfoVALVE-yChromaOffset-10933)VUID-VkVideoEncodeSessionRgbConversionCreateInfoVALVE-yChromaOffset-10933  
  `yChromaOffset` **must** only be a bit set in [VkVideoEncodeRgbConversionCapabilitiesVALVE](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeRgbConversionCapabilitiesVALVE.html)::`yChromaOffsets` as returned by [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html) with `VkVideoEncodeRgbConversionCapabilitiesVALVE` in the `pNext` chain of `VkVideoCapabilitiesKHR` with the given `pVideoProfile`

Valid Usage (Implicit)

- [](#VUID-VkVideoEncodeSessionRgbConversionCreateInfoVALVE-sType-sType)VUID-VkVideoEncodeSessionRgbConversionCreateInfoVALVE-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_ENCODE_SESSION_RGB_CONVERSION_CREATE_INFO_VALVE`
- [](#VUID-VkVideoEncodeSessionRgbConversionCreateInfoVALVE-rgbModel-parameter)VUID-VkVideoEncodeSessionRgbConversionCreateInfoVALVE-rgbModel-parameter  
  `rgbModel` **must** be a valid [VkVideoEncodeRgbModelConversionFlagBitsVALVE](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeRgbModelConversionFlagBitsVALVE.html) value
- [](#VUID-VkVideoEncodeSessionRgbConversionCreateInfoVALVE-rgbRange-parameter)VUID-VkVideoEncodeSessionRgbConversionCreateInfoVALVE-rgbRange-parameter  
  `rgbRange` **must** be a valid [VkVideoEncodeRgbRangeCompressionFlagBitsVALVE](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeRgbRangeCompressionFlagBitsVALVE.html) value
- [](#VUID-VkVideoEncodeSessionRgbConversionCreateInfoVALVE-xChromaOffset-parameter)VUID-VkVideoEncodeSessionRgbConversionCreateInfoVALVE-xChromaOffset-parameter  
  `xChromaOffset` **must** be a valid [VkVideoEncodeRgbChromaOffsetFlagBitsVALVE](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeRgbChromaOffsetFlagBitsVALVE.html) value
- [](#VUID-VkVideoEncodeSessionRgbConversionCreateInfoVALVE-yChromaOffset-parameter)VUID-VkVideoEncodeSessionRgbConversionCreateInfoVALVE-yChromaOffset-parameter  
  `yChromaOffset` **must** be a valid [VkVideoEncodeRgbChromaOffsetFlagBitsVALVE](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeRgbChromaOffsetFlagBitsVALVE.html) value

## [](#_see_also)See Also

[VK\_VALVE\_video\_encode\_rgb\_conversion](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VALVE_video_encode_rgb_conversion.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkVideoEncodeRgbChromaOffsetFlagBitsVALVE](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeRgbChromaOffsetFlagBitsVALVE.html), [VkVideoEncodeRgbModelConversionFlagBitsVALVE](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeRgbModelConversionFlagBitsVALVE.html), [VkVideoEncodeRgbRangeCompressionFlagBitsVALVE](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeRgbRangeCompressionFlagBitsVALVE.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoEncodeSessionRgbConversionCreateInfoVALVE).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0