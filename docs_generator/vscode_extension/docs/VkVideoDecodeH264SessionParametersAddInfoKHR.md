# VkVideoDecodeH264SessionParametersAddInfoKHR(3) Manual Page

## Name

VkVideoDecodeH264SessionParametersAddInfoKHR - Structure specifies H.264 decoder parameter set information



## [](#_c_specification)C Specification

The `VkVideoDecodeH264SessionParametersAddInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_video_decode_h264
typedef struct VkVideoDecodeH264SessionParametersAddInfoKHR {
    VkStructureType                            sType;
    const void*                                pNext;
    uint32_t                                   stdSPSCount;
    const StdVideoH264SequenceParameterSet*    pStdSPSs;
    uint32_t                                   stdPPSCount;
    const StdVideoH264PictureParameterSet*     pStdPPSs;
} VkVideoDecodeH264SessionParametersAddInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `stdSPSCount` is the number of elements in the `pStdSPSs` array.
- `pStdSPSs` is a pointer to an array of `StdVideoH264SequenceParameterSet` structures describing the [H.264 SPS](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-h264-sps) entries to add.
- `stdPPSCount` is the number of elements in the `pStdPPSs` array.
- `pStdPPSs` is a pointer to an array of `StdVideoH264PictureParameterSet` structures describing the [H.264 PPS](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-h264-pps) entries to add.

## [](#_description)Description

This structure **can** be specified in the following places:

- In the `pParametersAddInfo` member of the [VkVideoDecodeH264SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH264SessionParametersCreateInfoKHR.html) structure specified in the `pNext` chain of [VkVideoSessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionParametersCreateInfoKHR.html) used to create a [video session parameters](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#video-session-parameters) object. In this case, if the video codec operation the video session parameters object is created with is `VK_VIDEO_CODEC_OPERATION_DECODE_H264_BIT_KHR`, then it defines the set of initial parameters to add to the created object (see [Creating Video Session Parameters](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#creating-video-session-parameters)).
- In the `pNext` chain of [VkVideoSessionParametersUpdateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionParametersUpdateInfoKHR.html). In this case, if the video codec operation the [video session parameters](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#video-session-parameters) object to be updated was created with is `VK_VIDEO_CODEC_OPERATION_DECODE_H264_BIT_KHR`, then it defines the set of parameters to add to it (see [Updating Video Session Parameters](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#video-session-parameters-update)).

Valid Usage

- [](#VUID-VkVideoDecodeH264SessionParametersAddInfoKHR-None-04825)VUID-VkVideoDecodeH264SessionParametersAddInfoKHR-None-04825  
  The `seq_parameter_set_id` member of each `StdVideoH264SequenceParameterSet` structure specified in the elements of `pStdSPSs` **must** be unique within `pStdSPSs`
- [](#VUID-VkVideoDecodeH264SessionParametersAddInfoKHR-None-04826)VUID-VkVideoDecodeH264SessionParametersAddInfoKHR-None-04826  
  The pair constructed from the `seq_parameter_set_id` and `pic_parameter_set_id` members of each `StdVideoH264PictureParameterSet` structure specified in the elements of `pStdPPSs` **must** be unique within `pStdPPSs`

Valid Usage (Implicit)

- [](#VUID-VkVideoDecodeH264SessionParametersAddInfoKHR-sType-sType)VUID-VkVideoDecodeH264SessionParametersAddInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_DECODE_H264_SESSION_PARAMETERS_ADD_INFO_KHR`
- [](#VUID-VkVideoDecodeH264SessionParametersAddInfoKHR-pStdSPSs-parameter)VUID-VkVideoDecodeH264SessionParametersAddInfoKHR-pStdSPSs-parameter  
  If `stdSPSCount` is not `0`, `pStdSPSs` **must** be a valid pointer to an array of `stdSPSCount` `StdVideoH264SequenceParameterSet` values
- [](#VUID-VkVideoDecodeH264SessionParametersAddInfoKHR-pStdPPSs-parameter)VUID-VkVideoDecodeH264SessionParametersAddInfoKHR-pStdPPSs-parameter  
  If `stdPPSCount` is not `0`, `pStdPPSs` **must** be a valid pointer to an array of `stdPPSCount` `StdVideoH264PictureParameterSet` values

## [](#_see_also)See Also

[VK\_KHR\_video\_decode\_h264](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_decode_h264.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkVideoDecodeH264SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH264SessionParametersCreateInfoKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoDecodeH264SessionParametersAddInfoKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0