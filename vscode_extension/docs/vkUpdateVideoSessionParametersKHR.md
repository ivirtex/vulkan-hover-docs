# vkUpdateVideoSessionParametersKHR(3) Manual Page

## Name

vkUpdateVideoSessionParametersKHR - Update video session parameters object



## [](#_c_specification)C Specification

To update video session parameters object with new parameters, call:

```c++
// Provided by VK_KHR_video_queue
VkResult vkUpdateVideoSessionParametersKHR(
    VkDevice                                    device,
    VkVideoSessionParametersKHR                 videoSessionParameters,
    const VkVideoSessionParametersUpdateInfoKHR* pUpdateInfo);
```

## [](#_parameters)Parameters

- `device` is the logical device that updates the video session parameters.
- `videoSessionParameters` is the video session parameters object to update.
- `pUpdateInfo` is a pointer to a [VkVideoSessionParametersUpdateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionParametersUpdateInfoKHR.html) structure specifying the parameter update information.

## [](#_description)Description

After a successful call to this command, the [update sequence counter](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#video-session-parameters) of `videoSessionParameters` is changed to the value specified in `pUpdateInfo->updateSequenceCount`.

Note

As each update issued to a video session parameters object needs to specify the next available update sequence count value, concurrent updates of the same video session parameters object are inherently disallowed. However, recording video coding operations to command buffers referring to parameters previously added to the video session parameters object is allowed, even if there is a concurrent update in progress adding some new entries to the object.

If `videoSessionParameters` was created with the video codec operation `VK_VIDEO_CODEC_OPERATION_DECODE_H264_BIT_KHR` and the `pUpdateInfo->pNext` chain includes a [VkVideoDecodeH264SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH264SessionParametersAddInfoKHR.html) structure, then this command adds the following parameter entries to `videoSessionParameters`:

- The [H.264 SPS](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-h264-sps) entries specified in [VkVideoDecodeH264SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH264SessionParametersAddInfoKHR.html)::`pStdSPSs`.
- The [H.264 PPS](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-h264-pps) entries specified in [VkVideoDecodeH264SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH264SessionParametersAddInfoKHR.html)::`pStdPPSs`.

If `videoSessionParameters` was created with the video codec operation `VK_VIDEO_CODEC_OPERATION_DECODE_H265_BIT_KHR` and the `pUpdateInfo->pNext` chain includes a [VkVideoDecodeH265SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH265SessionParametersAddInfoKHR.html) structure, then this command adds the following parameter entries to `videoSessionParameters`:

- The [H.265 VPS](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-h265-vps) entries specified in [VkVideoDecodeH265SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH265SessionParametersAddInfoKHR.html)::`pStdVPSs`.
- The [H.265 SPS](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-h265-sps) entries specified in [VkVideoDecodeH265SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH265SessionParametersAddInfoKHR.html)::`pStdSPSs`.
- The [H.265 PPS](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-h265-pps) entries specified in [VkVideoDecodeH265SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH265SessionParametersAddInfoKHR.html)::`pStdPPSs`.

If `videoSessionParameters` was created with the video codec operation `VK_VIDEO_CODEC_OPERATION_ENCODE_H264_BIT_KHR` and the `pUpdateInfo->pNext` chain includes a [VkVideoEncodeH264SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264SessionParametersAddInfoKHR.html) structure, then this command adds the following parameter entries to `videoSessionParameters`:

- The [H.264 SPS](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h264-sps) entries specified in [VkVideoEncodeH264SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264SessionParametersAddInfoKHR.html)::`pStdSPSs`.
- The [H.264 PPS](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h264-pps) entries specified in [VkVideoEncodeH264SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264SessionParametersAddInfoKHR.html)::`pStdPPSs`.

If `videoSessionParameters` was created with the video codec operation `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR` and the `pUpdateInfo->pNext` chain includes a [VkVideoEncodeH265SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265SessionParametersAddInfoKHR.html) structure, then this command adds the following parameter entries to `videoSessionParameters`:

- The [H.265 VPS](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h265-vps) entries specified in [VkVideoEncodeH265SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265SessionParametersAddInfoKHR.html)::`pStdVPSs`.
- The [H.265 SPS](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h265-sps) entries specified in [VkVideoEncodeH265SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265SessionParametersAddInfoKHR.html)::`pStdSPSs`.
- The [H.265 PPS](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h265-pps) entries specified in [VkVideoEncodeH265SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265SessionParametersAddInfoKHR.html)::`pStdPPSs`.

In case of video session parameters objects created with a video encode operation, implementations **may** return the `VK_ERROR_INVALID_VIDEO_STD_PARAMETERS_KHR` error if any of the specified Video Std parameters do not adhere to the syntactic or semantic requirements of the used video compression standard, or if values derived from parameters according to the rules defined by the used video compression standard do not adhere to the capabilities of the video compression standard or the implementation.

Note

Applications **should** not rely on the `VK_ERROR_INVALID_VIDEO_STD_PARAMETERS_KHR` error being returned by any command as a means to verify Video Std parameters, as implementations are not required to report the error in any specific set of cases.

Valid Usage

- [](#VUID-vkUpdateVideoSessionParametersKHR-pUpdateInfo-07215)VUID-vkUpdateVideoSessionParametersKHR-pUpdateInfo-07215  
  `pUpdateInfo->updateSequenceCount` **must** equal the current [update sequence counter](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#video-session-parameters) of `videoSessionParameters` plus one
- [](#VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07216)VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07216  
  If `videoSessionParameters` was created with the video codec operation `VK_VIDEO_CODEC_OPERATION_DECODE_H264_BIT_KHR` and the `pNext` chain of `pUpdateInfo` includes a [VkVideoDecodeH264SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH264SessionParametersAddInfoKHR.html) structure, then `videoSessionParameters` **must** not already contain a `StdVideoH264SequenceParameterSet` entry with `seq_parameter_set_id` matching any of the elements of [VkVideoDecodeH264SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH264SessionParametersAddInfoKHR.html)::`pStdSPSs`
- [](#VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07217)VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07217  
  If `videoSessionParameters` was created with the video codec operation `VK_VIDEO_CODEC_OPERATION_DECODE_H264_BIT_KHR`, then the number of `StdVideoH264SequenceParameterSet` entries already stored in it plus the value of the `stdSPSCount` member of the [VkVideoDecodeH264SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH264SessionParametersAddInfoKHR.html) structure included in the `pUpdateInfo->pNext` chain **must** be less than or equal to the [VkVideoDecodeH264SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH264SessionParametersCreateInfoKHR.html)::`maxStdSPSCount` `videoSessionParameters` was created with
- [](#VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07218)VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07218  
  If `videoSessionParameters` was created with the video codec operation `VK_VIDEO_CODEC_OPERATION_DECODE_H264_BIT_KHR` and the `pNext` chain of `pUpdateInfo` includes a [VkVideoDecodeH264SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH264SessionParametersAddInfoKHR.html) structure, then `videoSessionParameters` **must** not already contain a `StdVideoH264PictureParameterSet` entry with both `seq_parameter_set_id` and `pic_parameter_set_id` matching any of the elements of [VkVideoDecodeH264SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH264SessionParametersAddInfoKHR.html)::`pStdPPSs`
- [](#VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07219)VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07219  
  If `videoSessionParameters` was created with the video codec operation `VK_VIDEO_CODEC_OPERATION_DECODE_H264_BIT_KHR`, then the number of `StdVideoH264PictureParameterSet` entries already stored in it plus the value of the `stdPPSCount` member of the [VkVideoDecodeH264SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH264SessionParametersAddInfoKHR.html) structure included in the `pUpdateInfo->pNext` chain **must** be less than or equal to the [VkVideoDecodeH264SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH264SessionParametersCreateInfoKHR.html)::`maxStdPPSCount` `videoSessionParameters` was created with
- [](#VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07220)VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07220  
  If `videoSessionParameters` was created with the video codec operation `VK_VIDEO_CODEC_OPERATION_DECODE_H265_BIT_KHR` and the `pNext` chain of `pUpdateInfo` includes a [VkVideoDecodeH265SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH265SessionParametersAddInfoKHR.html) structure, then `videoSessionParameters` **must** not already contain a `StdVideoH265VideoParameterSet` entry with `vps_video_parameter_set_id` matching any of the elements of [VkVideoDecodeH265SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH265SessionParametersAddInfoKHR.html)::`pStdVPSs`
- [](#VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07221)VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07221  
  If `videoSessionParameters` was created with the video codec operation `VK_VIDEO_CODEC_OPERATION_DECODE_H265_BIT_KHR`, then the number of `StdVideoH265VideoParameterSet` entries already stored in it plus the value of the `stdVPSCount` member of the [VkVideoDecodeH265SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH265SessionParametersAddInfoKHR.html) structure included in the `pUpdateInfo->pNext` chain **must** be less than or equal to the [VkVideoDecodeH265SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH265SessionParametersCreateInfoKHR.html)::`maxStdVPSCount` `videoSessionParameters` was created with
- [](#VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07222)VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07222  
  If `videoSessionParameters` was created with the video codec operation `VK_VIDEO_CODEC_OPERATION_DECODE_H265_BIT_KHR` and the `pNext` chain of `pUpdateInfo` includes a [VkVideoDecodeH265SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH265SessionParametersAddInfoKHR.html) structure, then `videoSessionParameters` **must** not already contain a `StdVideoH265SequenceParameterSet` entry with both `sps_video_parameter_set_id` and `sps_seq_parameter_set_id` matching any of the elements of [VkVideoDecodeH265SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH265SessionParametersAddInfoKHR.html)::`pStdSPSs`
- [](#VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07223)VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07223  
  If `videoSessionParameters` was created with the video codec operation `VK_VIDEO_CODEC_OPERATION_DECODE_H265_BIT_KHR`, then the number of `StdVideoH265SequenceParameterSet` entries already stored in it plus the value of the `stdSPSCount` member of the [VkVideoDecodeH265SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH265SessionParametersAddInfoKHR.html) structure included in the `pUpdateInfo->pNext` chain **must** be less than or equal to the [VkVideoDecodeH265SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH265SessionParametersCreateInfoKHR.html)::`maxStdSPSCount` `videoSessionParameters` was created with
- [](#VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07224)VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07224  
  If `videoSessionParameters` was created with the video codec operation `VK_VIDEO_CODEC_OPERATION_DECODE_H265_BIT_KHR` and the `pNext` chain of `pUpdateInfo` includes a [VkVideoDecodeH265SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH265SessionParametersAddInfoKHR.html) structure, then `videoSessionParameters` **must** not already contain a `StdVideoH265PictureParameterSet` entry with `sps_video_parameter_set_id`, `pps_seq_parameter_set_id`, and `pps_pic_parameter_set_id` all matching any of the elements of [VkVideoDecodeH265SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH265SessionParametersAddInfoKHR.html)::`pStdPPSs`
- [](#VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07225)VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07225  
  If `videoSessionParameters` was created with the video codec operation `VK_VIDEO_CODEC_OPERATION_DECODE_H265_BIT_KHR`, then the number of `StdVideoH265PictureParameterSet` entries already stored in it plus the value of the `stdPPSCount` member of the [VkVideoDecodeH265SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH265SessionParametersAddInfoKHR.html) structure included in the `pUpdateInfo->pNext` chain **must** be less than or equal to the [VkVideoDecodeH265SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH265SessionParametersCreateInfoKHR.html)::`maxStdPPSCount` `videoSessionParameters` was created with
- [](#VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-09260)VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-09260  
  `videoSessionParameters` **must** not have been created with the video codec operation `VK_VIDEO_CODEC_OPERATION_DECODE_AV1_BIT_KHR`
- [](#VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07226)VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07226  
  If `videoSessionParameters` was created with the video codec operation `VK_VIDEO_CODEC_OPERATION_ENCODE_H264_BIT_KHR` and the `pNext` chain of `pUpdateInfo` includes a [VkVideoEncodeH264SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264SessionParametersAddInfoKHR.html) structure, then `videoSessionParameters` **must** not already contain a `StdVideoH264SequenceParameterSet` entry with `seq_parameter_set_id` matching any of the elements of [VkVideoEncodeH264SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264SessionParametersAddInfoKHR.html)::`pStdSPSs`
- [](#VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-06441)VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-06441  
  If `videoSessionParameters` was created with the video codec operation `VK_VIDEO_CODEC_OPERATION_ENCODE_H264_BIT_KHR`, then the number of `StdVideoH264SequenceParameterSet` entries already stored in it plus the value of the `stdSPSCount` member of the [VkVideoEncodeH264SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264SessionParametersAddInfoKHR.html) structure included in the `pUpdateInfo->pNext` chain **must** be less than or equal to the [VkVideoEncodeH264SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264SessionParametersCreateInfoKHR.html)::`maxStdSPSCount` `videoSessionParameters` was created with
- [](#VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07227)VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07227  
  If `videoSessionParameters` was created with the video codec operation `VK_VIDEO_CODEC_OPERATION_ENCODE_H264_BIT_KHR` and the `pNext` chain of `pUpdateInfo` includes a [VkVideoEncodeH264SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264SessionParametersAddInfoKHR.html) structure, then `videoSessionParameters` **must** not already contain a `StdVideoH264PictureParameterSet` entry with both `seq_parameter_set_id` and `pic_parameter_set_id` matching any of the elements of [VkVideoEncodeH264SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264SessionParametersAddInfoKHR.html)::`pStdPPSs`
- [](#VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-06442)VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-06442  
  If `videoSessionParameters` was created with the video codec operation `VK_VIDEO_CODEC_OPERATION_ENCODE_H264_BIT_KHR`, then the number of `StdVideoH264PictureParameterSet` entries already stored in it plus the value of the `stdPPSCount` member of the [VkVideoEncodeH264SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264SessionParametersAddInfoKHR.html) structure included in the `pUpdateInfo->pNext` chain **must** be less than or equal to the [VkVideoEncodeH264SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264SessionParametersCreateInfoKHR.html)::`maxStdPPSCount` `videoSessionParameters` was created with
- [](#VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07228)VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07228  
  If `videoSessionParameters` was created with the video codec operation `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR` and the `pNext` chain of `pUpdateInfo` includes a [VkVideoEncodeH265SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265SessionParametersAddInfoKHR.html) structure, then `videoSessionParameters` **must** not already contain a `StdVideoH265VideoParameterSet` entry with `vps_video_parameter_set_id` matching any of the elements of [VkVideoEncodeH265SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265SessionParametersAddInfoKHR.html)::`pStdVPSs`
- [](#VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-06443)VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-06443  
  If `videoSessionParameters` was created with the video codec operation `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR`, then the number of `StdVideoH265VideoParameterSet` entries already stored in it plus the value of the `stdVPSCount` member of the [VkVideoEncodeH265SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265SessionParametersAddInfoKHR.html) structure included in the `pUpdateInfo->pNext` chain **must** be less than or equal to the [VkVideoEncodeH265SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265SessionParametersCreateInfoKHR.html)::`maxStdVPSCount` `videoSessionParameters` was created with
- [](#VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07229)VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07229  
  If `videoSessionParameters` was created with the video codec operation `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR` and the `pNext` chain of `pUpdateInfo` includes a [VkVideoEncodeH265SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265SessionParametersAddInfoKHR.html) structure, then `videoSessionParameters` **must** not already contain a `StdVideoH265SequenceParameterSet` entry with both `sps_video_parameter_set_id` and `sps_seq_parameter_set_id` matching any of the elements of [VkVideoEncodeH265SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265SessionParametersAddInfoKHR.html)::`pStdSPSs`
- [](#VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-06444)VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-06444  
  If `videoSessionParameters` was created with the video codec operation `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR`, then the number of `StdVideoH265SequenceParameterSet` entries already stored in it plus the value of the `stdSPSCount` member of the [VkVideoEncodeH265SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265SessionParametersAddInfoKHR.html) structure included in the `pUpdateInfo->pNext` chain **must** be less than or equal to the [VkVideoEncodeH265SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265SessionParametersCreateInfoKHR.html)::`maxStdSPSCount` `videoSessionParameters` was created with
- [](#VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07230)VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07230  
  If `videoSessionParameters` was created with the video codec operation `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR` and the `pNext` chain of `pUpdateInfo` includes a [VkVideoEncodeH265SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265SessionParametersAddInfoKHR.html) structure, then `videoSessionParameters` **must** not already contain a `StdVideoH265PictureParameterSet` entry with `sps_video_parameter_set_id`, `pps_seq_parameter_set_id`, and `pps_pic_parameter_set_id` all matching any of the elements of [VkVideoEncodeH265SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265SessionParametersAddInfoKHR.html)::`pStdPPSs`
- [](#VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-06445)VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-06445  
  If `videoSessionParameters` was created with the video codec operation `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR`, then the number of `StdVideoH265PictureParameterSet` entries already stored in it plus the value of the `stdPPSCount` member of the [VkVideoEncodeH265SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265SessionParametersAddInfoKHR.html) structure included in the `pUpdateInfo->pNext` chain **must** be less than or equal to the [VkVideoEncodeH265SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265SessionParametersCreateInfoKHR.html)::`maxStdPPSCount` `videoSessionParameters` was created with
- [](#VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-08321)VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-08321  
  If `videoSessionParameters` was created with the video codec operation `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR` and the `pNext` chain of `pUpdateInfo` includes a [VkVideoEncodeH265SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265SessionParametersAddInfoKHR.html) structure, then `num_tile_columns_minus1` **must** be less than [VkVideoEncodeH265CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265CapabilitiesKHR.html)::`maxTiles.width`, as returned by [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html) for the video profile `videoSessionParameters` was created with, for each element of [VkVideoEncodeH265SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265SessionParametersAddInfoKHR.html)::`pStdPPSs`
- [](#VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-08322)VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-08322  
  If `videoSessionParameters` was created with the video codec operation `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR` and the `pNext` chain of `pUpdateInfo` includes a [VkVideoEncodeH265SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265SessionParametersAddInfoKHR.html) structure, then `num_tile_rows_minus1` **must** be less than [VkVideoEncodeH265CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265CapabilitiesKHR.html)::`maxTiles.height`, as returned by [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html) for the video profile `videoSessionParameters` was created with, for each element of [VkVideoEncodeH265SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265SessionParametersAddInfoKHR.html)::`pStdPPSs`
- [](#VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-10281)VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-10281  
  `videoSessionParameters` **must** not have been created with the video codec operation `VK_VIDEO_CODEC_OPERATION_ENCODE_AV1_BIT_KHR`

Valid Usage (Implicit)

- [](#VUID-vkUpdateVideoSessionParametersKHR-device-parameter)VUID-vkUpdateVideoSessionParametersKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-parameter)VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-parameter  
  `videoSessionParameters` **must** be a valid [VkVideoSessionParametersKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionParametersKHR.html) handle
- [](#VUID-vkUpdateVideoSessionParametersKHR-pUpdateInfo-parameter)VUID-vkUpdateVideoSessionParametersKHR-pUpdateInfo-parameter  
  `pUpdateInfo` **must** be a valid pointer to a valid [VkVideoSessionParametersUpdateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionParametersUpdateInfoKHR.html) structure
- [](#VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-parent)VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-parent  
  `videoSessionParameters` **must** have been created, allocated, or retrieved from `device`

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_INVALID_VIDEO_STD_PARAMETERS_KHR`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_KHR\_video\_queue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_queue.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkVideoSessionParametersKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionParametersKHR.html), [VkVideoSessionParametersUpdateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionParametersUpdateInfoKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkUpdateVideoSessionParametersKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0