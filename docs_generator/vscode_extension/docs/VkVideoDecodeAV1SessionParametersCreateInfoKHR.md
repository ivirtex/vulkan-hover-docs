# VkVideoDecodeAV1SessionParametersCreateInfoKHR(3) Manual Page

## Name

VkVideoDecodeAV1SessionParametersCreateInfoKHR - Structure specifies AV1 decoder parameter set information



## [](#_c_specification)C Specification

When a [video session parameters](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#video-session-parameters) object is created with the codec operation `VK_VIDEO_CODEC_OPERATION_DECODE_AV1_BIT_KHR`, the [VkVideoSessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionParametersCreateInfoKHR.html)::`pNext` chain **must** include a `VkVideoDecodeAV1SessionParametersCreateInfoKHR` structure specifying the contents of the object.

The `VkVideoDecodeAV1SessionParametersCreateInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_video_decode_av1
typedef struct VkVideoDecodeAV1SessionParametersCreateInfoKHR {
    VkStructureType                     sType;
    const void*                         pNext;
    const StdVideoAV1SequenceHeader*    pStdSequenceHeader;
} VkVideoDecodeAV1SessionParametersCreateInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `pStdSequenceHeader` is a pointer to a `StdVideoAV1SequenceHeader` structure describing the [AV1 sequence header](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-av1-sequence-header) entry to store in the created object.

## [](#_description)Description

Note

As AV1 video session parameters objects will only ever contain a single AV1 sequence header, this has to be specified at object creation time and such video session parameters objects cannot be updated using the [vkUpdateVideoSessionParametersKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkUpdateVideoSessionParametersKHR.html) command. When a new AV1 sequence header is decoded from the input video bitstream the application needs to create a new video session parameters object to store it.

Valid Usage (Implicit)

- [](#VUID-VkVideoDecodeAV1SessionParametersCreateInfoKHR-sType-sType)VUID-VkVideoDecodeAV1SessionParametersCreateInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_DECODE_AV1_SESSION_PARAMETERS_CREATE_INFO_KHR`
- [](#VUID-VkVideoDecodeAV1SessionParametersCreateInfoKHR-pStdSequenceHeader-parameter)VUID-VkVideoDecodeAV1SessionParametersCreateInfoKHR-pStdSequenceHeader-parameter  
  `pStdSequenceHeader` **must** be a valid pointer to a valid `StdVideoAV1SequenceHeader` value

## [](#_see_also)See Also

[VK\_KHR\_video\_decode\_av1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_decode_av1.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoDecodeAV1SessionParametersCreateInfoKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0