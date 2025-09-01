# VkVideoEncodeH265SessionParametersAddInfoKHR(3) Manual Page

## Name

VkVideoEncodeH265SessionParametersAddInfoKHR - Structure specifies H.265 encoder parameter set information



## [](#_c_specification)C Specification

The `VkVideoEncodeH265SessionParametersAddInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_video_encode_h265
typedef struct VkVideoEncodeH265SessionParametersAddInfoKHR {
    VkStructureType                            sType;
    const void*                                pNext;
    uint32_t                                   stdVPSCount;
    const StdVideoH265VideoParameterSet*       pStdVPSs;
    uint32_t                                   stdSPSCount;
    const StdVideoH265SequenceParameterSet*    pStdSPSs;
    uint32_t                                   stdPPSCount;
    const StdVideoH265PictureParameterSet*     pStdPPSs;
} VkVideoEncodeH265SessionParametersAddInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `stdVPSCount` is the number of elements in the `pStdVPSs` array.
- `pStdVPSs` is a pointer to an array of `StdVideoH265VideoParameterSet` structures describing the [H.265 VPS](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h265-vps) entries to add.
- `stdSPSCount` is the number of elements in the `pStdSPSs` array.
- `pStdSPSs` is a pointer to an array of `StdVideoH265SequenceParameterSet` structures describing the [H.265 SPS](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h265-sps) entries to add.
- `stdPPSCount` is the number of elements in the `pStdPPSs` array.
- `pStdPPSs` is a pointer to an array of `StdVideoH265PictureParameterSet` structures describing the [H.265 PPS](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h265-pps) entries to add.

## [](#_description)Description

This structure **can** be specified in the following places:

- In the `pParametersAddInfo` member of the [VkVideoEncodeH265SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265SessionParametersCreateInfoKHR.html) structure specified in the `pNext` chain of [VkVideoSessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionParametersCreateInfoKHR.html) used to create a [video session parameters](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#video-session-parameters) object. In this case, if the video codec operation the video session parameters object is created with is `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR`, then it defines the set of initial parameters to add to the created object (see [Creating Video Session Parameters](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#creating-video-session-parameters)).
- In the `pNext` chain of [VkVideoSessionParametersUpdateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionParametersUpdateInfoKHR.html). In this case, if the video codec operation the [video session parameters](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#video-session-parameters) object to be updated was created with is `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR`, then it defines the set of parameters to add to it (see [Updating Video Session Parameters](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#video-session-parameters-update)).

Valid Usage (Implicit)

- [](#VUID-VkVideoEncodeH265SessionParametersAddInfoKHR-sType-sType)VUID-VkVideoEncodeH265SessionParametersAddInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_ENCODE_H265_SESSION_PARAMETERS_ADD_INFO_KHR`
- [](#VUID-VkVideoEncodeH265SessionParametersAddInfoKHR-pStdVPSs-parameter)VUID-VkVideoEncodeH265SessionParametersAddInfoKHR-pStdVPSs-parameter  
  If `stdVPSCount` is not `0`, and `pStdVPSs` is not `NULL`, `pStdVPSs` **must** be a valid pointer to an array of `stdVPSCount` `StdVideoH265VideoParameterSet` values
- [](#VUID-VkVideoEncodeH265SessionParametersAddInfoKHR-pStdSPSs-parameter)VUID-VkVideoEncodeH265SessionParametersAddInfoKHR-pStdSPSs-parameter  
  If `stdSPSCount` is not `0`, and `pStdSPSs` is not `NULL`, `pStdSPSs` **must** be a valid pointer to an array of `stdSPSCount` `StdVideoH265SequenceParameterSet` values
- [](#VUID-VkVideoEncodeH265SessionParametersAddInfoKHR-pStdPPSs-parameter)VUID-VkVideoEncodeH265SessionParametersAddInfoKHR-pStdPPSs-parameter  
  If `stdPPSCount` is not `0`, and `pStdPPSs` is not `NULL`, `pStdPPSs` **must** be a valid pointer to an array of `stdPPSCount` `StdVideoH265PictureParameterSet` values

Valid Usage

- [](#VUID-VkVideoEncodeH265SessionParametersAddInfoKHR-None-06438)VUID-VkVideoEncodeH265SessionParametersAddInfoKHR-None-06438  
  The `vps_video_parameter_set_id` member of each `StdVideoH265VideoParameterSet` structure specified in the elements of `pStdVPSs` **must** be unique within `pStdVPSs`
- [](#VUID-VkVideoEncodeH265SessionParametersAddInfoKHR-None-06439)VUID-VkVideoEncodeH265SessionParametersAddInfoKHR-None-06439  
  The pair constructed from the `sps_video_parameter_set_id` and `sps_seq_parameter_set_id` members of each `StdVideoH265SequenceParameterSet` structure specified in the elements of `pStdSPSs` **must** be unique within `pStdSPSs`
- [](#VUID-VkVideoEncodeH265SessionParametersAddInfoKHR-None-06440)VUID-VkVideoEncodeH265SessionParametersAddInfoKHR-None-06440  
  The triplet constructed from the `sps_video_parameter_set_id`, `pps_seq_parameter_set_id`, and `pps_pic_parameter_set_id` members of each `StdVideoH265PictureParameterSet` structure specified in the elements of `pStdPPSs` **must** be unique within `pStdPPSs`

## [](#_see_also)See Also

[VK\_KHR\_video\_encode\_h265](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_h265.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkVideoEncodeH265SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265SessionParametersCreateInfoKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoEncodeH265SessionParametersAddInfoKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0