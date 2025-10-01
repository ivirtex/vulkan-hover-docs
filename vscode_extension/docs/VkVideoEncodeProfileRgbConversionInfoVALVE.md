# VkVideoEncodeProfileRgbConversionInfoVALVE(3) Manual Page

## Name

VkVideoEncodeProfileRgbConversionInfoVALVE - Structure describing video encode RGB conversion for a video profile



## [](#_c_specification)C Specification

A video profile supporting video encode R′G′B′ conversion is specified by an `pVideoProfile->videoCodecOperation` specifying an encode operation and including a `VkVideoEncodeProfileRgbConversionInfoVALVE` structure in the `pNext` chain of the [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoProfileInfoKHR.html) structure and enabling `performEncodeRgbConversion`.

The `VkVideoEncodeProfileRgbConversionInfoVALVE` structure is defined as:

```c++
// Provided by VK_VALVE_video_encode_rgb_conversion
typedef struct VkVideoEncodeProfileRgbConversionInfoVALVE {
    VkStructureType    sType;
    const void*        pNext;
    VkBool32           performEncodeRgbConversion;
} VkVideoEncodeProfileRgbConversionInfoVALVE;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `performEncodeRgbConversion` is a boolean value indicating whether video encode R′G′B′ conversion will be used for the encode operation.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkVideoEncodeProfileRgbConversionInfoVALVE-sType-sType)VUID-VkVideoEncodeProfileRgbConversionInfoVALVE-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_ENCODE_PROFILE_RGB_CONVERSION_INFO_VALVE`

## [](#_see_also)See Also

[VK\_VALVE\_video\_encode\_rgb\_conversion](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VALVE_video_encode_rgb_conversion.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoEncodeProfileRgbConversionInfoVALVE).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0