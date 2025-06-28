# vkCreateVideoSessionParametersKHR(3) Manual Page

## Name

vkCreateVideoSessionParametersKHR - Creates video session parameters object



## [](#_c_specification)C Specification

To create a video session parameters object, call:

```c++
// Provided by VK_KHR_video_queue
VkResult vkCreateVideoSessionParametersKHR(
    VkDevice                                    device,
    const VkVideoSessionParametersCreateInfoKHR* pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkVideoSessionParametersKHR*                pVideoSessionParameters);
```

## [](#_parameters)Parameters

- `device` is the logical device that creates the video session parameters object.
- `pCreateInfo` is a pointer to [VkVideoSessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionParametersCreateInfoKHR.html) structure containing parameters to be used to create the video session parameters object.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.
- `pVideoSessionParameters` is a pointer to a [VkVideoSessionParametersKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionParametersKHR.html) handle in which the resulting video session parameters object is returned.

## [](#_description)Description

The resulting video session parameters object is said to be created with the video codec operation `pCreateInfo->videoSession` was created with.

Video session parameters objects created with an encode operation are always created with respect to a [video encode quality level](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-quality-level). By default, the created video session parameters objects are created with quality level zero, unless otherwise specified by including a [VkVideoEncodeQualityLevelInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeQualityLevelInfoKHR.html) structure in the `pCreateInfo->pNext` chain, in which case the video session parameters object is created with the quality level specified in [VkVideoEncodeQualityLevelInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeQualityLevelInfoKHR.html)::`qualityLevel`.

If `pCreateInfo->videoSessionParametersTemplate` is not `VK_NULL_HANDLE`, then it will be used as a template for constructing the new video session parameters object. This happens by first adding any parameters according to the additional creation parameters provided in the `pCreateInfo->pNext` chain, followed by adding any parameters from the template object that have a key that does not match the key of any of the already added parameters.

For video session parameters objects created with an encode operation, the template object specified in `pCreateInfo->videoSessionParametersTemplate` **must** have been created with the same [video encode quality level](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-quality-level) as the newly created object.

Note

This means that codec-specific parameters stored in video session parameters objects **can** only be reused across different video encode quality levels by re-specifying them, as previously created video session parameters against other quality levels **cannot** be used as template because the original codec-specific parameters (before the implementation **may** have applied [parameter overrides](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-overrides)) **may** no longer be available in them for the purposes of constructing the derived object.

Video session parameters objects are only compatible with [quantization maps](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-quantization-map) if they are created with `pCreateInfo->flags` including `VK_VIDEO_SESSION_PARAMETERS_CREATE_QUANTIZATION_MAP_COMPATIBLE_BIT_KHR`.

Video session parameters objects created with `VK_VIDEO_SESSION_PARAMETERS_CREATE_QUANTIZATION_MAP_COMPATIBLE_BIT_KHR` against a video session object that was created with `VK_VIDEO_SESSION_CREATE_ALLOW_ENCODE_QUANTIZATION_DELTA_MAP_BIT_KHR` or `VK_VIDEO_SESSION_CREATE_ALLOW_ENCODE_EMPHASIS_MAP_BIT_KHR` are created with a specific compatible [quantization map texel size](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-quantization-map-texel-size) specified in the `quantizationMapTexelSize` member of the [VkVideoEncodeQuantizationMapSessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeQuantizationMapSessionParametersCreateInfoKHR.html) structure included in the `pNext` chain of `pCreateInfo`.

Note

This means that the quantization map texel size that such a video session parameters object is compatible with is fixed for the lifetime of the object. Applications have to create separate video session parameters objects to use different quantization map texel sizes with a single video session object. This is necessary because the used quantization map texel size may affect the [parameter overrides](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-overrides) the implementation has to perform and thus the final values of the used codec-specific parameters.

For video session parameters objects created with `VK_VIDEO_SESSION_PARAMETERS_CREATE_QUANTIZATION_MAP_COMPATIBLE_BIT_KHR`, the template object specified in `pCreateInfo->videoSessionParametersTemplate` **must** also have been created with `VK_VIDEO_SESSION_PARAMETERS_CREATE_QUANTIZATION_MAP_COMPATIBLE_BIT_KHR` and the same compatible [quantization map texel size](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-quantization-map-texel-size) specified in [VkVideoEncodeQuantizationMapSessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeQuantizationMapSessionParametersCreateInfoKHR.html)::`quantizationMapTexelSize`.

Note

This means that codec-specific parameters stored in video session parameters objects **can** only be reused with different quantization map texel sizes by re-specifying them, as previously created video session parameters against other quantization map texel sizes **cannot** be used as template because the original codec-specific parameters (before the implementation **may** have applied [parameter overrides](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-overrides)) **may** no longer be available in them for the purposes of constructing the derived object.

For video session parameters objects created without `VK_VIDEO_SESSION_PARAMETERS_CREATE_QUANTIZATION_MAP_COMPATIBLE_BIT_KHR`, the template object specified in `pCreateInfo->videoSessionParametersTemplate` **must** also have been created without `VK_VIDEO_SESSION_PARAMETERS_CREATE_QUANTIZATION_MAP_COMPATIBLE_BIT_KHR`.

If `pCreateInfo->videoSession` was created with the video codec operation `VK_VIDEO_CODEC_OPERATION_DECODE_H264_BIT_KHR`, then the created video session parameters object will initially contain the following sets of parameter entries:

- `StdVideoH264SequenceParameterSet` structures representing [H.264 SPS](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-h264-sps) entries, as follows:
  
  - If the `pParametersAddInfo` member of the [VkVideoDecodeH264SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH264SessionParametersCreateInfoKHR.html) structure provided in the `pCreateInfo->pNext` chain is not `NULL`, then the set of `StdVideoH264SequenceParameterSet` entries specified in `pParametersAddInfo->pStdSPSs` are added first;
  - If `pCreateInfo->videoSessionParametersTemplate` is not `VK_NULL_HANDLE`, then each `StdVideoH264SequenceParameterSet` entry stored in it is copied to the created video session parameters object if the created object does not already contain such an entry with the same `seq_parameter_set_id`.
- `StdVideoH264PictureParameterSet` structures representing [H.264 PPS](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-h264-pps) entries, as follows:
  
  - If the `pParametersAddInfo` member of the [VkVideoDecodeH264SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH264SessionParametersCreateInfoKHR.html) structure provided in the `pCreateInfo->pNext` chain is not `NULL`, then the set of `StdVideoH264PictureParameterSet` entries specified in `pParametersAddInfo->pStdPPSs` are added first;
  - If `pCreateInfo->videoSessionParametersTemplate` is not `VK_NULL_HANDLE`, then each `StdVideoH264PictureParameterSet` entry stored in it is copied to the created video session parameters object if the created object does not already contain such an entry with the same `seq_parameter_set_id` and `pic_parameter_set_id`.

If `pCreateInfo->videoSession` was created with the video codec operation `VK_VIDEO_CODEC_OPERATION_DECODE_H265_BIT_KHR`, then the created video session parameters object will initially contain the following sets of parameter entries:

- `StdVideoH265VideoParameterSet` structures representing [H.265 VPS](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-h265-vps) entries, as follows:
  
  - If the `pParametersAddInfo` member of the [VkVideoDecodeH265SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH265SessionParametersCreateInfoKHR.html) structure provided in the `pCreateInfo->pNext` chain is not `NULL`, then the set of `StdVideoH265VideoParameterSet` entries specified in `pParametersAddInfo->pStdVPSs` are added first;
  - If `pCreateInfo->videoSessionParametersTemplate` is not `VK_NULL_HANDLE`, then each `StdVideoH265VideoParameterSet` entry stored in it is copied to the created video session parameters object if the created object does not already contain such an entry with the same `vps_video_parameter_set_id`.
- `StdVideoH265SequenceParameterSet` structures representing [H.265 SPS](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-h265-sps) entries, as follows:
  
  - If the `pParametersAddInfo` member of the [VkVideoDecodeH265SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH265SessionParametersCreateInfoKHR.html) structure provided in the `pCreateInfo->pNext` chain is not `NULL`, then the set of `StdVideoH265SequenceParameterSet` entries specified in `pParametersAddInfo->pStdSPSs` are added first;
  - If `pCreateInfo->videoSessionParametersTemplate` is not `VK_NULL_HANDLE`, then each `StdVideoH265SequenceParameterSet` entry stored in it is copied to the created video session parameters object if the created object does not already contain such an entry with the same `sps_video_parameter_set_id` and `sps_seq_parameter_set_id`.
- `StdVideoH265PictureParameterSet` structures representing [H.265 PPS](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-h265-pps) entries, as follows:
  
  - If the `pParametersAddInfo` member of the [VkVideoDecodeH265SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH265SessionParametersCreateInfoKHR.html) structure provided in the `pCreateInfo->pNext` chain is not `NULL`, then the set of `StdVideoH265PictureParameterSet` entries specified in `pParametersAddInfo->pStdPPSs` are added first;
  - If `pCreateInfo->videoSessionParametersTemplate` is not `VK_NULL_HANDLE`, then each `StdVideoH265PictureParameterSet` entry stored in it is copied to the created video session parameters object if the created object does not already contain such an entry with the same `sps_video_parameter_set_id`, `pps_seq_parameter_set_id`, and `pps_pic_parameter_set_id`.

If `pCreateInfo->videoSession` was created with the video codec operation `VK_VIDEO_CODEC_OPERATION_DECODE_AV1_BIT_KHR`, then the created video session parameters object will contain a single [AV1 sequence header](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-av1-sequence-header) represented by a `StdVideoAV1SequenceHeader` structure specified through the `pStdSequenceHeader` member of the [VkVideoDecodeAV1SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeAV1SessionParametersCreateInfoKHR.html) structure provided in the `pCreateInfo->pNext` chain. As such video session parameters objects **can** only contain a single [AV1 sequence header](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-av1-sequence-header), it is not possible to use a previously created object as a template or subsequently update the created video session parameters object.

If `pCreateInfo->videoSession` was created with the video codec operation `VK_VIDEO_CODEC_OPERATION_ENCODE_H264_BIT_KHR`, then the created video session parameters object will initially contain the following sets of parameter entries:

- `StdVideoH264SequenceParameterSet` structures representing [H.264 SPS](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h264-sps) entries, as follows:
  
  - If the `pParametersAddInfo` member of the [VkVideoEncodeH264SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264SessionParametersCreateInfoKHR.html) structure provided in the `pCreateInfo->pNext` chain is not `NULL`, then the set of `StdVideoH264SequenceParameterSet` entries specified in `pParametersAddInfo->pStdSPSs` are added first;
  - If `pCreateInfo->videoSessionParametersTemplate` is not `VK_NULL_HANDLE`, then each `StdVideoH264SequenceParameterSet` entry stored in it is copied to the created video session parameters object if the created object does not already contain such an entry with the same `seq_parameter_set_id`.
- `StdVideoH264PictureParameterSet` structures representing [H.264 PPS](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h264-pps) entries, as follows:
  
  - If the `pParametersAddInfo` member of the [VkVideoEncodeH264SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264SessionParametersCreateInfoKHR.html) structure provided in the `pCreateInfo->pNext` chain is not `NULL`, then the set of `StdVideoH264PictureParameterSet` entries specified in `pParametersAddInfo->pStdPPSs` are added first;
  - If `pCreateInfo->videoSessionParametersTemplate` is not `VK_NULL_HANDLE`, then each `StdVideoH264PictureParameterSet` entry stored in it is copied to the created video session parameters object if the created object does not already contain such an entry with the same `seq_parameter_set_id` and `pic_parameter_set_id`.

If `pCreateInfo->videoSession` was created with the video codec operation `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR`, then the created video session parameters object will initially contain the following sets of parameter entries:

- `StdVideoH265VideoParameterSet` structures representing [H.265 VPS](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h265-vps) entries, as follows:
  
  - If the `pParametersAddInfo` member of the [VkVideoEncodeH265SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265SessionParametersCreateInfoKHR.html) structure provided in the `pCreateInfo->pNext` chain is not `NULL`, then the set of `StdVideoH265VideoParameterSet` entries specified in `pParametersAddInfo->pStdVPSs` are added first;
  - If `pCreateInfo->videoSessionParametersTemplate` is not `VK_NULL_HANDLE`, then each `StdVideoH265VideoParameterSet` entry stored in it is copied to the created video session parameters object if the created object does not already contain such an entry with the same `vps_video_parameter_set_id`.
- `StdVideoH265SequenceParameterSet` structures representing [H.265 SPS](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h265-sps) entries, as follows:
  
  - If the `pParametersAddInfo` member of the [VkVideoEncodeH265SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265SessionParametersCreateInfoKHR.html) structure provided in the `pCreateInfo->pNext` chain is not `NULL`, then the set of `StdVideoH265SequenceParameterSet` entries specified in `pParametersAddInfo->pStdSPSs` are added first;
  - If `pCreateInfo->videoSessionParametersTemplate` is not `VK_NULL_HANDLE`, then each `StdVideoH265SequenceParameterSet` entry stored in it is copied to the created video session parameters object if the created object does not already contain such an entry with the same `sps_video_parameter_set_id` and `sps_seq_parameter_set_id`.
- `StdVideoH265PictureParameterSet` structures representing [H.265 PPS](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h265-pps) entries, as follows:
  
  - If the `pParametersAddInfo` member of the [VkVideoEncodeH265SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265SessionParametersCreateInfoKHR.html) structure provided in the `pCreateInfo->pNext` chain is not `NULL`, then the set of `StdVideoH265PictureParameterSet` entries specified in `pParametersAddInfo->pStdPPSs` are added first;
  - If `pCreateInfo->videoSessionParametersTemplate` is not `VK_NULL_HANDLE`, then each `StdVideoH265PictureParameterSet` entry stored in it is copied to the created video session parameters object if the created object does not already contain such an entry with the same `sps_video_parameter_set_id`, `pps_seq_parameter_set_id`, and `pps_pic_parameter_set_id`.

If `pCreateInfo->videoSession` was created with the video codec operation `VK_VIDEO_CODEC_OPERATION_ENCODE_AV1_BIT_KHR`, then the created video session parameters object will contain a single [AV1 sequence header](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-sequence-header) specified through the members of the [VkVideoEncodeAV1SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1SessionParametersCreateInfoKHR.html) structure provided in the `pCreateInfo->pNext` chain. As such video session parameters objects **can** only contain a single [AV1 sequence header](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-sequence-header), it is not possible to use a previously created object as a template or subsequently update the created video session parameters object.

In case of video session parameters objects created with a video encode operation, implementations **may** return the `VK_ERROR_INVALID_VIDEO_STD_PARAMETERS_KHR` error if any of the specified Video Std parameters do not adhere to the syntactic or semantic requirements of the used video compression standard, or if values derived from parameters according to the rules defined by the used video compression standard do not adhere to the capabilities of the video compression standard or the implementation.

Note

Applications **should** not rely on the `VK_ERROR_INVALID_VIDEO_STD_PARAMETERS_KHR` error being returned by any command as a means to verify Video Std parameters, as implementations are not required to report the error in any specific set of cases.

Valid Usage (Implicit)

- [](#VUID-vkCreateVideoSessionParametersKHR-device-parameter)VUID-vkCreateVideoSessionParametersKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkCreateVideoSessionParametersKHR-pCreateInfo-parameter)VUID-vkCreateVideoSessionParametersKHR-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid [VkVideoSessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionParametersCreateInfoKHR.html) structure
- [](#VUID-vkCreateVideoSessionParametersKHR-pAllocator-parameter)VUID-vkCreateVideoSessionParametersKHR-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkCreateVideoSessionParametersKHR-pVideoSessionParameters-parameter)VUID-vkCreateVideoSessionParametersKHR-pVideoSessionParameters-parameter  
  `pVideoSessionParameters` **must** be a valid pointer to a [VkVideoSessionParametersKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionParametersKHR.html) handle
- [](#VUID-vkCreateVideoSessionParametersKHR-device-queuecount)VUID-vkCreateVideoSessionParametersKHR-device-queuecount  
  The device **must** have been created with at least `1` queue

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_INITIALIZATION_FAILED`
- `VK_ERROR_INVALID_VIDEO_STD_PARAMETERS_KHR`

## [](#_see_also)See Also

[VK\_KHR\_video\_queue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_queue.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkVideoSessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionParametersCreateInfoKHR.html), [VkVideoSessionParametersKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionParametersKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCreateVideoSessionParametersKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0