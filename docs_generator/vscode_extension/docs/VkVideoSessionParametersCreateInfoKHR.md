# VkVideoSessionParametersCreateInfoKHR(3) Manual Page

## Name

VkVideoSessionParametersCreateInfoKHR - Structure specifying parameters of a newly created video session parameters object



## [](#_c_specification)C Specification

The `VkVideoSessionParametersCreateInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_video_queue
typedef struct VkVideoSessionParametersCreateInfoKHR {
    VkStructureType                           sType;
    const void*                               pNext;
    VkVideoSessionParametersCreateFlagsKHR    flags;
    VkVideoSessionParametersKHR               videoSessionParametersTemplate;
    VkVideoSessionKHR                         videoSession;
} VkVideoSessionParametersCreateInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is a bitmask of [VkVideoSessionParametersCreateFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionParametersCreateFlagBitsKHR.html) specifying create flags.
- `videoSessionParametersTemplate` is `VK_NULL_HANDLE` or a valid handle to a [VkVideoSessionParametersKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionParametersKHR.html) object used as a template for constructing the new video session parameters object.
- `videoSession` is the video session object against which the video session parameters object is going to be created.

## [](#_description)Description

Limiting values are defined below that are referenced by the relevant valid usage statements of this structure.

- If `videoSession` was created with the codec operation `VK_VIDEO_CODEC_OPERATION_DECODE_H264_BIT_KHR`, then let `StdVideoH264SequenceParameterSet spsAddList[]` be the list of [H.264 SPS](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-h264-sps) entries to add to the created video session parameters object, defined as follows:
  
  - If the `pParametersAddInfo` member of the [VkVideoDecodeH264SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH264SessionParametersCreateInfoKHR.html) structure provided in the `pNext` chain is not `NULL`, then the set of `StdVideoH264SequenceParameterSet` entries specified in `pParametersAddInfo->pStdSPSs` are added to `spsAddList`;
  - If `videoSessionParametersTemplate` is not `VK_NULL_HANDLE`, then each `StdVideoH264SequenceParameterSet` entry stored in it with `seq_parameter_set_id` not matching any of the entries already in `spsAddList` is added to `spsAddList`.
- If `videoSession` was created with the codec operation `VK_VIDEO_CODEC_OPERATION_DECODE_H264_BIT_KHR`, then let `StdVideoH264PictureParameterSet ppsAddList[]` be the list of [H.264 PPS](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-h264-pps) entries to add to the created video session parameters object, defined as follows:
  
  - If the `pParametersAddInfo` member of the [VkVideoDecodeH264SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH264SessionParametersCreateInfoKHR.html) structure provided in the `pNext` chain is not `NULL`, then the set of `StdVideoH264PictureParameterSet` entries specified in `pParametersAddInfo->pStdPPSs` are added to `ppsAddList`;
  - If `videoSessionParametersTemplate` is not `VK_NULL_HANDLE`, then each `StdVideoH264PictureParameterSet` entry stored in it with `seq_parameter_set_id` or `pic_parameter_set_id` not matching any of the entries already in `ppsAddList` is added to `ppsAddList`.
- If `videoSession` was created with the codec operation `VK_VIDEO_CODEC_OPERATION_DECODE_H265_BIT_KHR`, then let `StdVideoH265VideoParameterSet vpsAddList[]` be the list of [H.265 VPS](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-h265-vps) entries to add to the created video session parameters object, defined as follows:
  
  - If the `pParametersAddInfo` member of the [VkVideoDecodeH265SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH265SessionParametersCreateInfoKHR.html) structure provided in the `pNext` chain is not `NULL`, then the set of `StdVideoH265VideoParameterSet` entries specified in `pParametersAddInfo->pStdVPSs` are added to `vpsAddList`;
  - If `videoSessionParametersTemplate` is not `VK_NULL_HANDLE`, then each `StdVideoH265VideoParameterSet` entry stored in it with `vps_video_parameter_set_id` not matching any of the entries already in `vpsAddList` is added to `vpsAddList`.
- If `videoSession` was created with the codec operation `VK_VIDEO_CODEC_OPERATION_DECODE_H265_BIT_KHR`, then let `StdVideoH265SequenceParameterSet spsAddList[]` be the list of [H.265 SPS](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-h265-sps) entries to add to the created video session parameters object, defined as follows:
  
  - If the `pParametersAddInfo` member of the [VkVideoDecodeH265SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH265SessionParametersCreateInfoKHR.html) structure provided in the `pNext` chain is not `NULL`, then the set of `StdVideoH265SequenceParameterSet` entries specified in `pParametersAddInfo->pStdSPSs` are added to `spsAddList`;
  - If `videoSessionParametersTemplate` is not `VK_NULL_HANDLE`, then each `StdVideoH265SequenceParameterSet` entry stored in it with `sps_video_parameter_set_id` or `sps_seq_parameter_set_id` not matching any of the entries already in `spsAddList` is added to `spsAddList`.
- If `videoSession` was created with the codec operation `VK_VIDEO_CODEC_OPERATION_DECODE_H265_BIT_KHR`, then let `StdVideoH265PictureParameterSet ppsAddList[]` be the list of [H.265 PPS](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-h265-pps) entries to add to the created video session parameters object, defined as follows:
  
  - If the `pParametersAddInfo` member of the [VkVideoDecodeH265SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH265SessionParametersCreateInfoKHR.html) structure provided in the `pNext` chain is not `NULL`, then the set of `StdVideoH265PictureParameterSet` entries specified in `pParametersAddInfo->pStdPPSs` are added to `ppsAddList`;
  - If `videoSessionParametersTemplate` is not `VK_NULL_HANDLE`, then each `StdVideoH265PictureParameterSet` entry stored in it with `sps_video_parameter_set_id`, `pps_seq_parameter_set_id`, or `pps_pic_parameter_set_id` not matching any of the entries already in `ppsAddList` is added to `ppsAddList`.
- If `videoSession` was created with an encode operation, then let `uint32_t qualityLevel` be the [video encode quality level](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-quality-level) of the created video session parameters object, defined as follows:
  
  - If the `pNext` chain of this structure includes a [VkVideoEncodeQualityLevelInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeQualityLevelInfoKHR.html) structure, then `qualityLevel` is equal to [VkVideoEncodeQualityLevelInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeQualityLevelInfoKHR.html)::`qualityLevel`.
  - Otherwise `qualityLevel` is `0`
- If `videoSession` was created with the codec operation `VK_VIDEO_CODEC_OPERATION_ENCODE_H264_BIT_KHR`, then let `StdVideoH264SequenceParameterSet spsAddList[]` be the list of [H.264 SPS](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h264-sps) entries to add to the created video session parameters object, defined as follows:
  
  - If the `pParametersAddInfo` member of the [VkVideoEncodeH264SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264SessionParametersCreateInfoKHR.html) structure provided in the `pNext` chain is not `NULL`, then the set of `StdVideoH264SequenceParameterSet` entries specified in `pParametersAddInfo->pStdSPSs` are added to `spsAddList`;
  - If `videoSessionParametersTemplate` is not `VK_NULL_HANDLE`, then each `StdVideoH264SequenceParameterSet` entry stored in it with `seq_parameter_set_id` not matching any of the entries already in `spsAddList` is added to `spsAddList`.
- If `videoSession` was created with the codec operation `VK_VIDEO_CODEC_OPERATION_ENCODE_H264_BIT_KHR`, then let `StdVideoH264PictureParameterSet ppsAddList[]` be the list of [H.264 PPS](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h264-pps) entries to add to the created video session parameters object, defined as follows:
  
  - If the `pParametersAddInfo` member of the [VkVideoEncodeH264SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264SessionParametersCreateInfoKHR.html) structure provided in the `pNext` chain is not `NULL`, then the set of `StdVideoH264PictureParameterSet` entries specified in `pParametersAddInfo->pStdPPSs` are added to `ppsAddList`;
  - If `videoSessionParametersTemplate` is not `VK_NULL_HANDLE`, then each `StdVideoH264PictureParameterSet` entry stored in it with `seq_parameter_set_id` or `pic_parameter_set_id` not matching any of the entries already in `ppsAddList` is added to `ppsAddList`.
- If `videoSession` was created with the codec operation `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR`, then let `StdVideoH265VideoParameterSet vpsAddList[]` be the list of [H.265 VPS](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h265-vps) entries to add to the created video session parameters object, defined as follows:
  
  - If the `pParametersAddInfo` member of the [VkVideoEncodeH265SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265SessionParametersCreateInfoKHR.html) structure provided in the `pNext` chain is not `NULL`, then the set of `StdVideoH265VideoParameterSet` entries specified in `pParametersAddInfo->pStdVPSs` are added to `vpsAddList`;
  - If `videoSessionParametersTemplate` is not `VK_NULL_HANDLE`, then each `StdVideoH265VideoParameterSet` entry stored in it with `vps_video_parameter_set_id` not matching any of the entries already in `vpsAddList` is added to `vpsAddList`.
- If `videoSession` was created with the codec operation `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR`, then let `StdVideoH265SequenceParameterSet spsAddList[]` be the list of [H.265 SPS](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h265-sps) entries to add to the created video session parameters object, defined as follows:
  
  - If the `pParametersAddInfo` member of the [VkVideoEncodeH265SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265SessionParametersCreateInfoKHR.html) structure provided in the `pNext` chain is not `NULL`, then the set of `StdVideoH265SequenceParameterSet` entries specified in `pParametersAddInfo->pStdSPSs` are added to `spsAddList`;
  - If `videoSessionParametersTemplate` is not `VK_NULL_HANDLE`, then each `StdVideoH265SequenceParameterSet` entry stored in it with `sps_video_parameter_set_id` or `sps_seq_parameter_set_id` not matching any of the entries already in `spsAddList` is added to `spsAddList`.
- If `videoSession` was created with the codec operation `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR`, then let `StdVideoH265PictureParameterSet ppsAddList[]` be the list of [H.265 PPS](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h265-pps) entries to add to the created video session parameters object, defined as follows:
  
  - If the `pParametersAddInfo` member of the [VkVideoEncodeH265SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265SessionParametersCreateInfoKHR.html) structure provided in the `pNext` chain is not `NULL`, then the set of `StdVideoH265PictureParameterSet` entries specified in `pParametersAddInfo->pStdPPSs` are added to `ppsAddList`;
  - If `videoSessionParametersTemplate` is not `VK_NULL_HANDLE`, then each `StdVideoH265PictureParameterSet` entry stored in it with `sps_video_parameter_set_id`, `pps_seq_parameter_set_id`, or `pps_pic_parameter_set_id` not matching any of the entries already in `ppsAddList` is added to `ppsAddList`.

Valid Usage

- [](#VUID-VkVideoSessionParametersCreateInfoKHR-videoSessionParametersTemplate-04855)VUID-VkVideoSessionParametersCreateInfoKHR-videoSessionParametersTemplate-04855  
  If `videoSessionParametersTemplate` is not `VK_NULL_HANDLE`, it **must** have been created against `videoSession`
- [](#VUID-VkVideoSessionParametersCreateInfoKHR-videoSessionParametersTemplate-08310)VUID-VkVideoSessionParametersCreateInfoKHR-videoSessionParametersTemplate-08310  
  If `videoSessionParametersTemplate` is not `VK_NULL_HANDLE` and `videoSession` was created with an encode operation, then `qualityLevel` **must** equal the [video encode quality](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-quality-level) level `videoSessionParametersTemplate` was created with
- [](#VUID-VkVideoSessionParametersCreateInfoKHR-flags-10271)VUID-VkVideoSessionParametersCreateInfoKHR-flags-10271  
  If `flags` includes `VK_VIDEO_SESSION_PARAMETERS_CREATE_QUANTIZATION_MAP_COMPATIBLE_BIT_KHR`, then `videoSession` **must** have been created with `VK_VIDEO_SESSION_CREATE_ALLOW_ENCODE_QUANTIZATION_DELTA_MAP_BIT_KHR` or `VK_VIDEO_SESSION_CREATE_ALLOW_ENCODE_EMPHASIS_MAP_BIT_KHR`
- [](#VUID-VkVideoSessionParametersCreateInfoKHR-flags-10272)VUID-VkVideoSessionParametersCreateInfoKHR-flags-10272  
  If `flags` includes `VK_VIDEO_SESSION_PARAMETERS_CREATE_QUANTIZATION_MAP_COMPATIBLE_BIT_KHR`, then the `pNext` chain **must** include a [VkVideoEncodeQuantizationMapSessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeQuantizationMapSessionParametersCreateInfoKHR.html) structure
- [](#VUID-VkVideoSessionParametersCreateInfoKHR-flags-10273)VUID-VkVideoSessionParametersCreateInfoKHR-flags-10273  
  If `flags` includes `VK_VIDEO_SESSION_PARAMETERS_CREATE_QUANTIZATION_MAP_COMPATIBLE_BIT_KHR` and `videoSession` was created with `VK_VIDEO_SESSION_CREATE_ALLOW_ENCODE_QUANTIZATION_DELTA_MAP_BIT_KHR`, then the [list of video format properties](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#supported-video-format-properties) supported for the image usage flag `VK_IMAGE_USAGE_VIDEO_ENCODE_QUANTIZATION_DELTA_MAP_BIT_KHR` **must** have an element for which [VkVideoFormatQuantizationMapPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoFormatQuantizationMapPropertiesKHR.html)::`quantizationMapTexelSize` equals the `quantizationMapTexelSize` member of the [VkVideoEncodeQuantizationMapSessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeQuantizationMapSessionParametersCreateInfoKHR.html) structure included in the `pNext` chain
- [](#VUID-VkVideoSessionParametersCreateInfoKHR-flags-10274)VUID-VkVideoSessionParametersCreateInfoKHR-flags-10274  
  If `flags` includes `VK_VIDEO_SESSION_PARAMETERS_CREATE_QUANTIZATION_MAP_COMPATIBLE_BIT_KHR` and `videoSession` was created with `VK_VIDEO_SESSION_CREATE_ALLOW_ENCODE_EMPHASIS_MAP_BIT_KHR`, then the [list of video format properties](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#supported-video-format-properties) supported for the image usage flag `VK_IMAGE_USAGE_VIDEO_ENCODE_EMPHASIS_MAP_BIT_KHR` **must** have an element for which [VkVideoFormatQuantizationMapPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoFormatQuantizationMapPropertiesKHR.html)::`quantizationMapTexelSize` equals the `quantizationMapTexelSize` member of the [VkVideoEncodeQuantizationMapSessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeQuantizationMapSessionParametersCreateInfoKHR.html) structure included in the `pNext` chain
- [](#VUID-VkVideoSessionParametersCreateInfoKHR-videoSessionParametersTemplate-10275)VUID-VkVideoSessionParametersCreateInfoKHR-videoSessionParametersTemplate-10275  
  If `videoSessionParametersTemplate` is not `VK_NULL_HANDLE` and `flags` includes `VK_VIDEO_SESSION_PARAMETERS_CREATE_QUANTIZATION_MAP_COMPATIBLE_BIT_KHR`, then `videoSessionParametersTemplate` **must** have been created with `VK_VIDEO_SESSION_PARAMETERS_CREATE_QUANTIZATION_MAP_COMPATIBLE_BIT_KHR`
- [](#VUID-VkVideoSessionParametersCreateInfoKHR-videoSessionParametersTemplate-10276)VUID-VkVideoSessionParametersCreateInfoKHR-videoSessionParametersTemplate-10276  
  If `videoSessionParametersTemplate` is not `VK_NULL_HANDLE` and `flags` includes `VK_VIDEO_SESSION_PARAMETERS_CREATE_QUANTIZATION_MAP_COMPATIBLE_BIT_KHR`, then `videoSessionParametersTemplate` **must** have been created with the same [quantization map texel size](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-quantization-map-texel-size) as the one specified in the `quantizationMapTexelSize` member of the [VkVideoEncodeQuantizationMapSessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeQuantizationMapSessionParametersCreateInfoKHR.html) structure included in the `pNext` chain
- [](#VUID-VkVideoSessionParametersCreateInfoKHR-videoSessionParametersTemplate-10277)VUID-VkVideoSessionParametersCreateInfoKHR-videoSessionParametersTemplate-10277  
  If `videoSessionParametersTemplate` is not `VK_NULL_HANDLE` and `flags` does not include `VK_VIDEO_SESSION_PARAMETERS_CREATE_QUANTIZATION_MAP_COMPATIBLE_BIT_KHR`, then `videoSessionParametersTemplate` **must** have been created without `VK_VIDEO_SESSION_PARAMETERS_CREATE_QUANTIZATION_MAP_COMPATIBLE_BIT_KHR`
- [](#VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-07203)VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-07203  
  If `videoSession` was created with the video codec operation `VK_VIDEO_CODEC_OPERATION_DECODE_H264_BIT_KHR`, then the `pNext` chain **must** include a [VkVideoDecodeH264SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH264SessionParametersCreateInfoKHR.html) structure
- [](#VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-07204)VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-07204  
  If `videoSession` was created with the video codec operation `VK_VIDEO_CODEC_OPERATION_DECODE_H264_BIT_KHR`, then the number of elements of `spsAddList` **must** be less than or equal to the `maxStdSPSCount` specified in the [VkVideoDecodeH264SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH264SessionParametersCreateInfoKHR.html) structure included in the `pNext` chain
- [](#VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-07205)VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-07205  
  If `videoSession` was created with the video codec operation `VK_VIDEO_CODEC_OPERATION_DECODE_H264_BIT_KHR`, then the number of elements of `ppsAddList` **must** be less than or equal to the `maxStdPPSCount` specified in the [VkVideoDecodeH264SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH264SessionParametersCreateInfoKHR.html) structure included in the `pNext` chain
- [](#VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-07206)VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-07206  
  If `videoSession` was created with the video codec operation `VK_VIDEO_CODEC_OPERATION_DECODE_H265_BIT_KHR`, then the `pNext` chain **must** include a [VkVideoDecodeH265SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH265SessionParametersCreateInfoKHR.html) structure
- [](#VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-07207)VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-07207  
  If `videoSession` was created with the video codec operation `VK_VIDEO_CODEC_OPERATION_DECODE_H265_BIT_KHR`, then the number of elements of `vpsAddList` **must** be less than or equal to the `maxStdVPSCount` specified in the [VkVideoDecodeH265SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH265SessionParametersCreateInfoKHR.html) structure included in the `pNext` chain
- [](#VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-07208)VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-07208  
  If `videoSession` was created with the video codec operation `VK_VIDEO_CODEC_OPERATION_DECODE_H265_BIT_KHR`, then the number of elements of `spsAddList` **must** be less than or equal to the `maxStdSPSCount` specified in the [VkVideoDecodeH265SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH265SessionParametersCreateInfoKHR.html) structure included in the `pNext` chain
- [](#VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-07209)VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-07209  
  If `videoSession` was created with the video codec operation `VK_VIDEO_CODEC_OPERATION_DECODE_H265_BIT_KHR`, then the number of elements of `ppsAddList` **must** be less than or equal to the `maxStdPPSCount` specified in the [VkVideoDecodeH265SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH265SessionParametersCreateInfoKHR.html) structure included in the `pNext` chain
- [](#VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-10794)VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-10794  
  `videoSession` **must** not have been created with the codec operation `VK_VIDEO_CODEC_OPERATION_DECODE_VP9_BIT_KHR`
- [](#VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-09258)VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-09258  
  If `videoSession` was created with the video codec operation `VK_VIDEO_CODEC_OPERATION_DECODE_AV1_BIT_KHR`, then `videoSessionParametersTemplate` **must** be `VK_NULL_HANDLE`
- [](#VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-09259)VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-09259  
  If `videoSession` was created with the video codec operation `VK_VIDEO_CODEC_OPERATION_DECODE_AV1_BIT_KHR`, then the `pNext` chain **must** include a [VkVideoDecodeAV1SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeAV1SessionParametersCreateInfoKHR.html) structure
- [](#VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-07210)VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-07210  
  If `videoSession` was created with the video codec operation `VK_VIDEO_CODEC_OPERATION_ENCODE_H264_BIT_KHR`, then the `pNext` chain **must** include a [VkVideoEncodeH264SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264SessionParametersCreateInfoKHR.html) structure
- [](#VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-04839)VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-04839  
  If `videoSession` was created with the video codec operation `VK_VIDEO_CODEC_OPERATION_ENCODE_H264_BIT_KHR`, then the number of elements of `spsAddList` **must** be less than or equal to the `maxStdSPSCount` specified in the [VkVideoEncodeH264SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264SessionParametersCreateInfoKHR.html) structure included in the `pNext` chain
- [](#VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-04840)VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-04840  
  If `videoSession` was created with the video codec operation `VK_VIDEO_CODEC_OPERATION_ENCODE_H264_BIT_KHR`, then the number of elements of `ppsAddList` **must** be less than or equal to the `maxStdPPSCount` specified in the [VkVideoEncodeH264SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264SessionParametersCreateInfoKHR.html) structure included in the `pNext` chain
- [](#VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-07211)VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-07211  
  If `videoSession` was created with the video codec operation `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR`, then the `pNext` chain **must** include a [VkVideoEncodeH265SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265SessionParametersCreateInfoKHR.html) structure
- [](#VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-04841)VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-04841  
  If `videoSession` was created with the video codec operation `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR`, then the number of elements of `vpsAddList` **must** be less than or equal to the `maxStdVPSCount` specified in the [VkVideoEncodeH265SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265SessionParametersCreateInfoKHR.html) structure included in the `pNext` chain
- [](#VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-04842)VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-04842  
  If `videoSession` was created with the video codec operation `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR`, then the number of elements of `spsAddList` **must** be less than or equal to the `maxStdSPSCount` specified in the [VkVideoEncodeH265SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265SessionParametersCreateInfoKHR.html) structure included in the `pNext` chain
- [](#VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-04843)VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-04843  
  If `videoSession` was created with the video codec operation `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR`, then the number of elements of `ppsAddList` **must** be less than or equal to the `maxStdPPSCount` specified in the [VkVideoEncodeH265SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265SessionParametersCreateInfoKHR.html) structure included in the `pNext` chain
- [](#VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-08319)VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-08319  
  If `videoSession` was created with the video codec operation `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR`, then `num_tile_columns_minus1` **must** be less than [VkVideoEncodeH265CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265CapabilitiesKHR.html)::`maxTiles.width`, as returned by [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html) for the video profile `videoSession` was created with, for each element of `ppsAddList`
- [](#VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-08320)VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-08320  
  If `videoSession` was created with the video codec operation `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR`, then `num_tile_rows_minus1` **must** be less than [VkVideoEncodeH265CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265CapabilitiesKHR.html)::`maxTiles.height`, as returned by [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html) for the video profile `videoSession` was created with, for each element of `ppsAddList`
- [](#VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-10278)VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-10278  
  If `videoSession` was created with the video codec operation `VK_VIDEO_CODEC_OPERATION_ENCODE_AV1_BIT_KHR`, then `videoSessionParametersTemplate` **must** be `VK_NULL_HANDLE`
- [](#VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-10279)VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-10279  
  If `videoSession` was created with the video codec operation `VK_VIDEO_CODEC_OPERATION_ENCODE_AV1_BIT_KHR`, then the `pNext` chain **must** include a [VkVideoEncodeAV1SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1SessionParametersCreateInfoKHR.html) structure
- [](#VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-10280)VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-10280  
  If `videoSession` was created with the video codec operation `VK_VIDEO_CODEC_OPERATION_ENCODE_AV1_BIT_KHR`, then the `stdOperatingPointCount` member of the [VkVideoEncodeAV1SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1SessionParametersCreateInfoKHR.html) structure included in the `pNext` chain **must** be less than or equal to [VkVideoEncodeAV1CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1CapabilitiesKHR.html)::`maxOperatingPoints`, as returned by [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html) for the video profile `videoSession` was created with

Valid Usage (Implicit)

- [](#VUID-VkVideoSessionParametersCreateInfoKHR-sType-sType)VUID-VkVideoSessionParametersCreateInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_SESSION_PARAMETERS_CREATE_INFO_KHR`
- [](#VUID-VkVideoSessionParametersCreateInfoKHR-pNext-pNext)VUID-VkVideoSessionParametersCreateInfoKHR-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the `pNext` chain **must** be either `NULL` or a pointer to a valid instance of [VkVideoDecodeAV1SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeAV1SessionParametersCreateInfoKHR.html), [VkVideoDecodeH264SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH264SessionParametersCreateInfoKHR.html), [VkVideoDecodeH265SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH265SessionParametersCreateInfoKHR.html), [VkVideoEncodeAV1SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1SessionParametersCreateInfoKHR.html), [VkVideoEncodeH264SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264SessionParametersCreateInfoKHR.html), [VkVideoEncodeH265SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265SessionParametersCreateInfoKHR.html), [VkVideoEncodeQualityLevelInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeQualityLevelInfoKHR.html), or [VkVideoEncodeQuantizationMapSessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeQuantizationMapSessionParametersCreateInfoKHR.html)
- [](#VUID-VkVideoSessionParametersCreateInfoKHR-sType-unique)VUID-VkVideoSessionParametersCreateInfoKHR-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique
- [](#VUID-VkVideoSessionParametersCreateInfoKHR-flags-parameter)VUID-VkVideoSessionParametersCreateInfoKHR-flags-parameter  
  `flags` **must** be a valid combination of [VkVideoSessionParametersCreateFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionParametersCreateFlagBitsKHR.html) values
- [](#VUID-VkVideoSessionParametersCreateInfoKHR-videoSessionParametersTemplate-parameter)VUID-VkVideoSessionParametersCreateInfoKHR-videoSessionParametersTemplate-parameter  
  If `videoSessionParametersTemplate` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `videoSessionParametersTemplate` **must** be a valid [VkVideoSessionParametersKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionParametersKHR.html) handle
- [](#VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-parameter)VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-parameter  
  `videoSession` **must** be a valid [VkVideoSessionKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionKHR.html) handle
- [](#VUID-VkVideoSessionParametersCreateInfoKHR-videoSessionParametersTemplate-parent)VUID-VkVideoSessionParametersCreateInfoKHR-videoSessionParametersTemplate-parent  
  If `videoSessionParametersTemplate` is a valid handle, it **must** have been created, allocated, or retrieved from `videoSession`
- [](#VUID-VkVideoSessionParametersCreateInfoKHR-commonparent)VUID-VkVideoSessionParametersCreateInfoKHR-commonparent  
  Both of `videoSession`, and `videoSessionParametersTemplate` that are valid handles of non-ignored parameters **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_see_also)See Also

[VK\_KHR\_video\_queue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_queue.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkVideoSessionKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionKHR.html), [VkVideoSessionParametersCreateFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionParametersCreateFlagsKHR.html), [VkVideoSessionParametersKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionParametersKHR.html), [vkCreateVideoSessionParametersKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateVideoSessionParametersKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoSessionParametersCreateInfoKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0