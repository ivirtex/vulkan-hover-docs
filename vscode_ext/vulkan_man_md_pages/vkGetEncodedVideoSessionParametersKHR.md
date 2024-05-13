# vkGetEncodedVideoSessionParametersKHR(3) Manual Page

## Name

vkGetEncodedVideoSessionParametersKHR - Get encoded parameter sets from
a video session parameters object



## <a href="#_c_specification" class="anchor"></a>C Specification

Encoded parameter data **can** be retrieved from a video session
parameters object created with a video encode operation using the
command:

``` c
// Provided by VK_KHR_video_encode_queue
VkResult vkGetEncodedVideoSessionParametersKHR(
    VkDevice                                    device,
    const VkVideoEncodeSessionParametersGetInfoKHR* pVideoSessionParametersInfo,
    VkVideoEncodeSessionParametersFeedbackInfoKHR* pFeedbackInfo,
    size_t*                                     pDataSize,
    void*                                       pData);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that owns the video session parameters
  object.

- `pVideoSessionParametersInfo` is a pointer to a
  [VkVideoEncodeSessionParametersGetInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeSessionParametersGetInfoKHR.html)
  structure specifying the parameters of the encoded parameter data to
  retrieve.

- `pFeedbackInfo` is either `NULL` or a pointer to a
  [VkVideoEncodeSessionParametersFeedbackInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeSessionParametersFeedbackInfoKHR.html)
  structure in which feedback about the requested parameter data is
  returned.

- `pDataSize` is a pointer to a `size_t` value related to the amount of
  encode parameter data returned, as described below.

- `pData` is either `NULL` or a pointer to a buffer to write the encoded
  parameter data to.

## <a href="#_description" class="anchor"></a>Description

If `pData` is `NULL`, then the size of the encoded parameter data, in
bytes, that **can** be retrieved is returned in `pDataSize`. Otherwise,
`pDataSize` **must** point to a variable set by the application to the
size of the buffer, in bytes, pointed to by `pData`, and on return the
variable is overwritten with the number of bytes actually written to
`pData`. If `pDataSize` is less than the size of the encoded parameter
data that **can** be retrieved, then no data will be written to `pData`,
zero will be written to `pDataSize`, and `VK_INCOMPLETE` will be
returned instead of `VK_SUCCESS`, to indicate that no encoded parameter
data was returned.

If `pFeedbackInfo` is not `NULL` then the members of the
[VkVideoEncodeSessionParametersFeedbackInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeSessionParametersFeedbackInfoKHR.html)
structure and any additional structures included in its `pNext` chain
that are applicable to the video session parameters object specified in
`pVideoSessionParametersInfo->videoSessionParameters` will be filled
with feedback about the requested parameter data on all successful calls
to this command.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>This includes the cases when <code>pData</code> is <code>NULL</code>
or when <code>VK_INCOMPLETE</code> is returned by the command, and
enables the application to determine whether the implementation <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-overrides"
target="_blank" rel="noopener">overrode</a> any of the requested video
session parameters without actually needing to retrieve the encoded
parameter data itself.</p></td>
</tr>
</tbody>
</table>

Valid Usage

- <a
  href="#VUID-vkGetEncodedVideoSessionParametersKHR-pVideoSessionParametersInfo-08359"
  id="VUID-vkGetEncodedVideoSessionParametersKHR-pVideoSessionParametersInfo-08359"></a>
  VUID-vkGetEncodedVideoSessionParametersKHR-pVideoSessionParametersInfo-08359  
  `pVideoSessionParametersInfo->videoSessionParameters` **must** have
  been created with an encode operation

- <a
  href="#VUID-vkGetEncodedVideoSessionParametersKHR-pVideoSessionParametersInfo-08262"
  id="VUID-vkGetEncodedVideoSessionParametersKHR-pVideoSessionParametersInfo-08262"></a>
  VUID-vkGetEncodedVideoSessionParametersKHR-pVideoSessionParametersInfo-08262  
  If `pVideoSessionParametersInfo->videoSessionParameters` was created
  with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H264_BIT_KHR`, then the `pNext` chain
  of `pVideoSessionParametersInfo` **must** include a
  [VkVideoEncodeH264SessionParametersGetInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264SessionParametersGetInfoKHR.html)
  structure

- <a
  href="#VUID-vkGetEncodedVideoSessionParametersKHR-pVideoSessionParametersInfo-08263"
  id="VUID-vkGetEncodedVideoSessionParametersKHR-pVideoSessionParametersInfo-08263"></a>
  VUID-vkGetEncodedVideoSessionParametersKHR-pVideoSessionParametersInfo-08263  
  If `pVideoSessionParametersInfo->videoSessionParameters` was created
  with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H264_BIT_KHR`, then for the
  [VkVideoEncodeH264SessionParametersGetInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264SessionParametersGetInfoKHR.html)
  structure included in the `pNext` chain of
  `pVideoSessionParametersInfo`, if its `writeStdSPS` member is
  `VK_TRUE`, then `pVideoSessionParametersInfo->videoSessionParameters`
  **must** contain a `StdVideoH264SequenceParameterSet` entry with
  `seq_parameter_set_id` matching
  [VkVideoEncodeH264SessionParametersGetInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264SessionParametersGetInfoKHR.html)::`stdSPSId`

- <a
  href="#VUID-vkGetEncodedVideoSessionParametersKHR-pVideoSessionParametersInfo-08264"
  id="VUID-vkGetEncodedVideoSessionParametersKHR-pVideoSessionParametersInfo-08264"></a>
  VUID-vkGetEncodedVideoSessionParametersKHR-pVideoSessionParametersInfo-08264  
  If `pVideoSessionParametersInfo->videoSessionParameters` was created
  with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H264_BIT_KHR`, then for the
  [VkVideoEncodeH264SessionParametersGetInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264SessionParametersGetInfoKHR.html)
  structure included in the `pNext` chain of
  `pVideoSessionParametersInfo`, if its `writeStdPPS` member is
  `VK_TRUE`, then `pVideoSessionParametersInfo->videoSessionParameters`
  **must** contain a `StdVideoH264PictureParameterSet` entry with
  `seq_parameter_set_id` and `pic_parameter_set_id` matching
  [VkVideoEncodeH264SessionParametersGetInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264SessionParametersGetInfoKHR.html)::`stdSPSId`
  and
  [VkVideoEncodeH264SessionParametersGetInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264SessionParametersGetInfoKHR.html)::`stdPPSId`,
  respectively

- <a
  href="#VUID-vkGetEncodedVideoSessionParametersKHR-pVideoSessionParametersInfo-08265"
  id="VUID-vkGetEncodedVideoSessionParametersKHR-pVideoSessionParametersInfo-08265"></a>
  VUID-vkGetEncodedVideoSessionParametersKHR-pVideoSessionParametersInfo-08265  
  If `pVideoSessionParametersInfo->videoSessionParameters` was created
  with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR`, then the `pNext` chain
  of `pVideoSessionParametersInfo` **must** include a
  [VkVideoEncodeH265SessionParametersGetInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265SessionParametersGetInfoKHR.html)
  structure

- <a
  href="#VUID-vkGetEncodedVideoSessionParametersKHR-pVideoSessionParametersInfo-08266"
  id="VUID-vkGetEncodedVideoSessionParametersKHR-pVideoSessionParametersInfo-08266"></a>
  VUID-vkGetEncodedVideoSessionParametersKHR-pVideoSessionParametersInfo-08266  
  If `pVideoSessionParametersInfo->videoSessionParameters` was created
  with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR`, then for the
  [VkVideoEncodeH265SessionParametersGetInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265SessionParametersGetInfoKHR.html)
  structure included in the `pNext` chain of
  `pVideoSessionParametersInfo`, if its `writeStdVPS` member is
  `VK_TRUE`, then `pVideoSessionParametersInfo->videoSessionParameters`
  **must** contain a `StdVideoH265VideoParameterSet` entry with
  `vps_video_parameter_set_id` matching
  [VkVideoEncodeH265SessionParametersGetInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265SessionParametersGetInfoKHR.html)::`stdVPSId`

- <a
  href="#VUID-vkGetEncodedVideoSessionParametersKHR-pVideoSessionParametersInfo-08267"
  id="VUID-vkGetEncodedVideoSessionParametersKHR-pVideoSessionParametersInfo-08267"></a>
  VUID-vkGetEncodedVideoSessionParametersKHR-pVideoSessionParametersInfo-08267  
  If `pVideoSessionParametersInfo->videoSessionParameters` was created
  with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR`, then for the
  [VkVideoEncodeH265SessionParametersGetInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265SessionParametersGetInfoKHR.html)
  structure included in the `pNext` chain of
  `pVideoSessionParametersInfo`, if its `writeStdSPS` member is
  `VK_TRUE`, then `pVideoSessionParametersInfo->videoSessionParameters`
  **must** contain a `StdVideoH265SequenceParameterSet` entry with
  `sps_video_parameter_set_id` and `sps_seq_parameter_set_id` matching
  [VkVideoEncodeH265SessionParametersGetInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265SessionParametersGetInfoKHR.html)::`stdVPSId`
  and
  [VkVideoEncodeH265SessionParametersGetInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265SessionParametersGetInfoKHR.html)::`stdSPSId`,
  respectively

- <a
  href="#VUID-vkGetEncodedVideoSessionParametersKHR-pVideoSessionParametersInfo-08268"
  id="VUID-vkGetEncodedVideoSessionParametersKHR-pVideoSessionParametersInfo-08268"></a>
  VUID-vkGetEncodedVideoSessionParametersKHR-pVideoSessionParametersInfo-08268  
  If `pVideoSessionParametersInfo->videoSessionParameters` was created
  with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR`, then for the
  [VkVideoEncodeH265SessionParametersGetInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265SessionParametersGetInfoKHR.html)
  structure included in the `pNext` chain of
  `pVideoSessionParametersInfo`, if its `writeStdPPS` member is
  `VK_TRUE`, then `pVideoSessionParametersInfo->videoSessionParameters`
  **must** contain a `StdVideoH265PictureParameterSet` entry with
  `sps_video_parameter_set_id`, `pps_seq_parameter_set_id`, and
  `pps_pic_parameter_set_id` matching
  [VkVideoEncodeH265SessionParametersGetInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265SessionParametersGetInfoKHR.html)::`stdVPSId`,
  [VkVideoEncodeH265SessionParametersGetInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265SessionParametersGetInfoKHR.html)::`stdSPSId`,
  and
  [VkVideoEncodeH265SessionParametersGetInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265SessionParametersGetInfoKHR.html)::`stdPPSId`,
  respectively

Valid Usage (Implicit)

- <a href="#VUID-vkGetEncodedVideoSessionParametersKHR-device-parameter"
  id="VUID-vkGetEncodedVideoSessionParametersKHR-device-parameter"></a>
  VUID-vkGetEncodedVideoSessionParametersKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a
  href="#VUID-vkGetEncodedVideoSessionParametersKHR-pVideoSessionParametersInfo-parameter"
  id="VUID-vkGetEncodedVideoSessionParametersKHR-pVideoSessionParametersInfo-parameter"></a>
  VUID-vkGetEncodedVideoSessionParametersKHR-pVideoSessionParametersInfo-parameter  
  `pVideoSessionParametersInfo` **must** be a valid pointer to a valid
  [VkVideoEncodeSessionParametersGetInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeSessionParametersGetInfoKHR.html)
  structure

- <a
  href="#VUID-vkGetEncodedVideoSessionParametersKHR-pFeedbackInfo-parameter"
  id="VUID-vkGetEncodedVideoSessionParametersKHR-pFeedbackInfo-parameter"></a>
  VUID-vkGetEncodedVideoSessionParametersKHR-pFeedbackInfo-parameter  
  If `pFeedbackInfo` is not `NULL`, `pFeedbackInfo` **must** be a valid
  pointer to a
  [VkVideoEncodeSessionParametersFeedbackInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeSessionParametersFeedbackInfoKHR.html)
  structure

- <a
  href="#VUID-vkGetEncodedVideoSessionParametersKHR-pDataSize-parameter"
  id="VUID-vkGetEncodedVideoSessionParametersKHR-pDataSize-parameter"></a>
  VUID-vkGetEncodedVideoSessionParametersKHR-pDataSize-parameter  
  `pDataSize` **must** be a valid pointer to a `size_t` value

- <a href="#VUID-vkGetEncodedVideoSessionParametersKHR-pData-parameter"
  id="VUID-vkGetEncodedVideoSessionParametersKHR-pData-parameter"></a>
  VUID-vkGetEncodedVideoSessionParametersKHR-pData-parameter  
  If the value referenced by `pDataSize` is not `0`, and `pData` is not
  `NULL`, `pData` **must** be a valid pointer to an array of `pDataSize`
  bytes

Return Codes

On success, this command returns  
- `VK_SUCCESS`

- `VK_INCOMPLETE`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_encode_queue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_encode_queue.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkVideoEncodeSessionParametersFeedbackInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeSessionParametersFeedbackInfoKHR.html),
[VkVideoEncodeSessionParametersGetInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeSessionParametersGetInfoKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetEncodedVideoSessionParametersKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
