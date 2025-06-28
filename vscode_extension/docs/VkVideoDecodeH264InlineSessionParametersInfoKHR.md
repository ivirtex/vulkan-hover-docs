# VkVideoDecodeH264InlineSessionParametersInfoKHR(3) Manual Page

## Name

VkVideoDecodeH264InlineSessionParametersInfoKHR - Structure specifies inline H.264 decoder parameter set information



## [](#_c_specification)C Specification

The `VkVideoDecodeH264InlineSessionParametersInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_video_decode_h264 with VK_KHR_video_maintenance2
typedef struct VkVideoDecodeH264InlineSessionParametersInfoKHR {
    VkStructureType                            sType;
    const void*                                pNext;
    const StdVideoH264SequenceParameterSet*    pStdSPS;
    const StdVideoH264PictureParameterSet*     pStdPPS;
} VkVideoDecodeH264InlineSessionParametersInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `pStdSPS` is `NULL` or a pointer to an instance of the `StdVideoH264SequenceParameterSet` structure describing the [active H.264 SPS](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-h264-active-sps).
- `pStdPPS` is `NULL` or a pointer to an instance of the `StdVideoH264PictureParameterSet` structure describing the [active H.264 PPS](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-h264-active-pps).

## [](#_description)Description

If `pStdSPS` or `pStdPPS` is not `NULL`, the issued video decode operations will use the parameter sets specified by them, respectively, instead of the corresponding parameter sets being sourced from the bound video session parameters object.

Valid Usage (Implicit)

- [](#VUID-VkVideoDecodeH264InlineSessionParametersInfoKHR-sType-sType)VUID-VkVideoDecodeH264InlineSessionParametersInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_DECODE_H264_INLINE_SESSION_PARAMETERS_INFO_KHR`
- [](#VUID-VkVideoDecodeH264InlineSessionParametersInfoKHR-pStdSPS-parameter)VUID-VkVideoDecodeH264InlineSessionParametersInfoKHR-pStdSPS-parameter  
  If `pStdSPS` is not `NULL`, `pStdSPS` **must** be a valid pointer to a valid `StdVideoH264SequenceParameterSet` value
- [](#VUID-VkVideoDecodeH264InlineSessionParametersInfoKHR-pStdPPS-parameter)VUID-VkVideoDecodeH264InlineSessionParametersInfoKHR-pStdPPS-parameter  
  If `pStdPPS` is not `NULL`, `pStdPPS` **must** be a valid pointer to a valid `StdVideoH264PictureParameterSet` value

## [](#_see_also)See Also

[VK\_KHR\_video\_decode\_h264](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_decode_h264.html), [VK\_KHR\_video\_maintenance2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_maintenance2.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoDecodeH264InlineSessionParametersInfoKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0