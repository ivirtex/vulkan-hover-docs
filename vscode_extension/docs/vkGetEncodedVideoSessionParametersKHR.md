# vkGetEncodedVideoSessionParametersKHR(3) Manual Page

## Name

vkGetEncodedVideoSessionParametersKHR - Get encoded parameter sets from a video session parameters object



## [](#_c_specification)C Specification

Encoded parameter data **can** be retrieved from a video session parameters object created with a video encode operation using the command:

```c++
// Provided by VK_KHR_video_encode_queue
VkResult vkGetEncodedVideoSessionParametersKHR(
    VkDevice                                    device,
    const VkVideoEncodeSessionParametersGetInfoKHR* pVideoSessionParametersInfo,
    VkVideoEncodeSessionParametersFeedbackInfoKHR* pFeedbackInfo,
    size_t*                                     pDataSize,
    void*                                       pData);
```

## [](#_parameters)Parameters

- `device` is the logical device that owns the video session parameters object.
- `pVideoSessionParametersInfo` is a pointer to a [VkVideoEncodeSessionParametersGetInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeSessionParametersGetInfoKHR.html) structure specifying the parameters of the encoded parameter data to retrieve.
- `pFeedbackInfo` is either `NULL` or a pointer to a [VkVideoEncodeSessionParametersFeedbackInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeSessionParametersFeedbackInfoKHR.html) structure in which feedback about the requested parameter data is returned.
- `pDataSize` is a pointer to a `size_t` value related to the amount of encode parameter data returned, as described below.
- `pData` is either `NULL` or a pointer to a buffer to write the encoded parameter data to.

## [](#_description)Description

If `pData` is `NULL`, then the size of the encoded parameter data, in bytes, that **can** be retrieved is returned in `pDataSize`. Otherwise, `pDataSize` **must** point to a variable set by the application to the size of the buffer, in bytes, pointed to by `pData`, and on return the variable is overwritten with the number of bytes actually written to `pData`. If `pDataSize` is less than the size of the encoded parameter data that **can** be retrieved, then no data will be written to `pData`, zero will be written to `pDataSize`, and `VK_INCOMPLETE` will be returned instead of `VK_SUCCESS`, to indicate that no encoded parameter data was returned.

If `pFeedbackInfo` is not `NULL` then the members of the [VkVideoEncodeSessionParametersFeedbackInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeSessionParametersFeedbackInfoKHR.html) structure and any additional structures included in its `pNext` chain that are applicable to the video session parameters object specified in `pVideoSessionParametersInfo->videoSessionParameters` will be filled with feedback about the requested parameter data on all successful calls to this command.

Note

This includes the cases when `pData` is `NULL` or when `VK_INCOMPLETE` is returned by the command, and enables the application to determine whether the implementation [overrode](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-overrides) any of the requested video session parameters without actually needing to retrieve the encoded parameter data itself.

Note

This query does not behave consistently with the behavior described in [Opaque Binary Data Results](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#fundamentals-binaryresults), for historical reasons.

If the amount of data available is larger than the passed `pDataSize`, the query returns a `VK_INCOMPLETE` success status instead of a `VK_ERROR_NOT_ENOUGH_SPACE_KHR` error status, and writes zero to `pDataSize`.

Valid Usage

- [](#VUID-vkGetEncodedVideoSessionParametersKHR-pVideoSessionParametersInfo-08359)VUID-vkGetEncodedVideoSessionParametersKHR-pVideoSessionParametersInfo-08359  
  `pVideoSessionParametersInfo->videoSessionParameters` **must** have been created with an encode operation
- [](#VUID-vkGetEncodedVideoSessionParametersKHR-pVideoSessionParametersInfo-08262)VUID-vkGetEncodedVideoSessionParametersKHR-pVideoSessionParametersInfo-08262  
  If `pVideoSessionParametersInfo->videoSessionParameters` was created with the video codec operation `VK_VIDEO_CODEC_OPERATION_ENCODE_H264_BIT_KHR`, then the `pNext` chain of `pVideoSessionParametersInfo` **must** include a [VkVideoEncodeH264SessionParametersGetInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264SessionParametersGetInfoKHR.html) structure
- [](#VUID-vkGetEncodedVideoSessionParametersKHR-pVideoSessionParametersInfo-08263)VUID-vkGetEncodedVideoSessionParametersKHR-pVideoSessionParametersInfo-08263  
  If `pVideoSessionParametersInfo->videoSessionParameters` was created with the video codec operation `VK_VIDEO_CODEC_OPERATION_ENCODE_H264_BIT_KHR`, then for the [VkVideoEncodeH264SessionParametersGetInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264SessionParametersGetInfoKHR.html) structure included in the `pNext` chain of `pVideoSessionParametersInfo`, if its `writeStdSPS` member is `VK_TRUE`, then `pVideoSessionParametersInfo->videoSessionParameters` **must** contain a `StdVideoH264SequenceParameterSet` entry with `seq_parameter_set_id` matching [VkVideoEncodeH264SessionParametersGetInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264SessionParametersGetInfoKHR.html)::`stdSPSId`
- [](#VUID-vkGetEncodedVideoSessionParametersKHR-pVideoSessionParametersInfo-08264)VUID-vkGetEncodedVideoSessionParametersKHR-pVideoSessionParametersInfo-08264  
  If `pVideoSessionParametersInfo->videoSessionParameters` was created with the video codec operation `VK_VIDEO_CODEC_OPERATION_ENCODE_H264_BIT_KHR`, then for the [VkVideoEncodeH264SessionParametersGetInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264SessionParametersGetInfoKHR.html) structure included in the `pNext` chain of `pVideoSessionParametersInfo`, if its `writeStdPPS` member is `VK_TRUE`, then `pVideoSessionParametersInfo->videoSessionParameters` **must** contain a `StdVideoH264PictureParameterSet` entry with `seq_parameter_set_id` and `pic_parameter_set_id` matching [VkVideoEncodeH264SessionParametersGetInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264SessionParametersGetInfoKHR.html)::`stdSPSId` and [VkVideoEncodeH264SessionParametersGetInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264SessionParametersGetInfoKHR.html)::`stdPPSId`, respectively
- [](#VUID-vkGetEncodedVideoSessionParametersKHR-pVideoSessionParametersInfo-08265)VUID-vkGetEncodedVideoSessionParametersKHR-pVideoSessionParametersInfo-08265  
  If `pVideoSessionParametersInfo->videoSessionParameters` was created with the video codec operation `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR`, then the `pNext` chain of `pVideoSessionParametersInfo` **must** include a [VkVideoEncodeH265SessionParametersGetInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265SessionParametersGetInfoKHR.html) structure
- [](#VUID-vkGetEncodedVideoSessionParametersKHR-pVideoSessionParametersInfo-08266)VUID-vkGetEncodedVideoSessionParametersKHR-pVideoSessionParametersInfo-08266  
  If `pVideoSessionParametersInfo->videoSessionParameters` was created with the video codec operation `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR`, then for the [VkVideoEncodeH265SessionParametersGetInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265SessionParametersGetInfoKHR.html) structure included in the `pNext` chain of `pVideoSessionParametersInfo`, if its `writeStdVPS` member is `VK_TRUE`, then `pVideoSessionParametersInfo->videoSessionParameters` **must** contain a `StdVideoH265VideoParameterSet` entry with `vps_video_parameter_set_id` matching [VkVideoEncodeH265SessionParametersGetInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265SessionParametersGetInfoKHR.html)::`stdVPSId`
- [](#VUID-vkGetEncodedVideoSessionParametersKHR-pVideoSessionParametersInfo-08267)VUID-vkGetEncodedVideoSessionParametersKHR-pVideoSessionParametersInfo-08267  
  If `pVideoSessionParametersInfo->videoSessionParameters` was created with the video codec operation `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR`, then for the [VkVideoEncodeH265SessionParametersGetInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265SessionParametersGetInfoKHR.html) structure included in the `pNext` chain of `pVideoSessionParametersInfo`, if its `writeStdSPS` member is `VK_TRUE`, then `pVideoSessionParametersInfo->videoSessionParameters` **must** contain a `StdVideoH265SequenceParameterSet` entry with `sps_video_parameter_set_id` and `sps_seq_parameter_set_id` matching [VkVideoEncodeH265SessionParametersGetInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265SessionParametersGetInfoKHR.html)::`stdVPSId` and [VkVideoEncodeH265SessionParametersGetInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265SessionParametersGetInfoKHR.html)::`stdSPSId`, respectively
- [](#VUID-vkGetEncodedVideoSessionParametersKHR-pVideoSessionParametersInfo-08268)VUID-vkGetEncodedVideoSessionParametersKHR-pVideoSessionParametersInfo-08268  
  If `pVideoSessionParametersInfo->videoSessionParameters` was created with the video codec operation `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR`, then for the [VkVideoEncodeH265SessionParametersGetInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265SessionParametersGetInfoKHR.html) structure included in the `pNext` chain of `pVideoSessionParametersInfo`, if its `writeStdPPS` member is `VK_TRUE`, then `pVideoSessionParametersInfo->videoSessionParameters` **must** contain a `StdVideoH265PictureParameterSet` entry with `sps_video_parameter_set_id`, `pps_seq_parameter_set_id`, and `pps_pic_parameter_set_id` matching [VkVideoEncodeH265SessionParametersGetInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265SessionParametersGetInfoKHR.html)::`stdVPSId`, [VkVideoEncodeH265SessionParametersGetInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265SessionParametersGetInfoKHR.html)::`stdSPSId`, and [VkVideoEncodeH265SessionParametersGetInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265SessionParametersGetInfoKHR.html)::`stdPPSId`, respectively

Valid Usage (Implicit)

- [](#VUID-vkGetEncodedVideoSessionParametersKHR-device-parameter)VUID-vkGetEncodedVideoSessionParametersKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetEncodedVideoSessionParametersKHR-pVideoSessionParametersInfo-parameter)VUID-vkGetEncodedVideoSessionParametersKHR-pVideoSessionParametersInfo-parameter  
  `pVideoSessionParametersInfo` **must** be a valid pointer to a valid [VkVideoEncodeSessionParametersGetInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeSessionParametersGetInfoKHR.html) structure
- [](#VUID-vkGetEncodedVideoSessionParametersKHR-pFeedbackInfo-parameter)VUID-vkGetEncodedVideoSessionParametersKHR-pFeedbackInfo-parameter  
  If `pFeedbackInfo` is not `NULL`, `pFeedbackInfo` **must** be a valid pointer to a [VkVideoEncodeSessionParametersFeedbackInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeSessionParametersFeedbackInfoKHR.html) structure
- [](#VUID-vkGetEncodedVideoSessionParametersKHR-pDataSize-parameter)VUID-vkGetEncodedVideoSessionParametersKHR-pDataSize-parameter  
  `pDataSize` **must** be a valid pointer to a `size_t` value
- [](#VUID-vkGetEncodedVideoSessionParametersKHR-pData-parameter)VUID-vkGetEncodedVideoSessionParametersKHR-pData-parameter  
  If the value referenced by `pDataSize` is not `0`, and `pData` is not `NULL`, `pData` **must** be a valid pointer to an array of `pDataSize` bytes

Return Codes

On success, this command returns

- `VK_SUCCESS`
- `VK_INCOMPLETE`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## [](#_see_also)See Also

[VK\_KHR\_video\_encode\_queue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_queue.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkVideoEncodeSessionParametersFeedbackInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeSessionParametersFeedbackInfoKHR.html), [VkVideoEncodeSessionParametersGetInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeSessionParametersGetInfoKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetEncodedVideoSessionParametersKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0