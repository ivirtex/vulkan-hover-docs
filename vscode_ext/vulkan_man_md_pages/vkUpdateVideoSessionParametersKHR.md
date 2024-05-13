# vkUpdateVideoSessionParametersKHR(3) Manual Page

## Name

vkUpdateVideoSessionParametersKHR - Update video session parameters
object



## <a href="#_c_specification" class="anchor"></a>C Specification

To update video session parameters object with new parameters, call:

``` c
// Provided by VK_KHR_video_queue
VkResult vkUpdateVideoSessionParametersKHR(
    VkDevice                                    device,
    VkVideoSessionParametersKHR                 videoSessionParameters,
    const VkVideoSessionParametersUpdateInfoKHR* pUpdateInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that updates the video session
  parameters.

- `videoSessionParameters` is the video session parameters object to
  update.

- `pUpdateInfo` is a pointer to a
  [VkVideoSessionParametersUpdateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionParametersUpdateInfoKHR.html)
  structure specifying the parameter update information.

## <a href="#_description" class="anchor"></a>Description

After a successful call to this command, the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-session-parameters"
target="_blank" rel="noopener">update sequence counter</a> of
`videoSessionParameters` is changed to the value specified in
`pUpdateInfo->updateSequenceCount`.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>As each update issued to a video session parameters object needs to
specify the next available update sequence count value, concurrent
updates of the same video session parameters object are inherently
disallowed. However, recording video coding operations to command
buffers referring to parameters previously added to the video session
parameters object is allowed, even if there is a concurrent update in
progress adding some new entries to the object.</p></td>
</tr>
</tbody>
</table>

If `videoSessionParameters` was created with the video codec operation
`VK_VIDEO_CODEC_OPERATION_DECODE_H264_BIT_KHR` and the
`pUpdateInfo->pNext` chain includes a
[VkVideoDecodeH264SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH264SessionParametersAddInfoKHR.html)
structure, then this command adds the following parameter entries to
`videoSessionParameters`:

- The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h264-sps"
  target="_blank" rel="noopener">H.264 SPS</a> entries specified in
  [VkVideoDecodeH264SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH264SessionParametersAddInfoKHR.html)::`pStdSPSs`.

- The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h264-pps"
  target="_blank" rel="noopener">H.264 PPS</a> entries specified in
  [VkVideoDecodeH264SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH264SessionParametersAddInfoKHR.html)::`pStdPPSs`.

If `videoSessionParameters` was created with the video codec operation
`VK_VIDEO_CODEC_OPERATION_DECODE_H265_BIT_KHR` and the
`pUpdateInfo->pNext` chain includes a
[VkVideoDecodeH265SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH265SessionParametersAddInfoKHR.html)
structure, then this command adds the following parameter entries to
`videoSessionParameters`:

- The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h265-vps"
  target="_blank" rel="noopener">H.265 VPS</a> entries specified in
  [VkVideoDecodeH265SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH265SessionParametersAddInfoKHR.html)::`pStdVPSs`.

- The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h265-sps"
  target="_blank" rel="noopener">H.265 SPS</a> entries specified in
  [VkVideoDecodeH265SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH265SessionParametersAddInfoKHR.html)::`pStdSPSs`.

- The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h265-pps"
  target="_blank" rel="noopener">H.265 PPS</a> entries specified in
  [VkVideoDecodeH265SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH265SessionParametersAddInfoKHR.html)::`pStdPPSs`.

If `videoSessionParameters` was created with the video codec operation
`VK_VIDEO_CODEC_OPERATION_ENCODE_H264_BIT_KHR` and the
`pUpdateInfo->pNext` chain includes a
[VkVideoEncodeH264SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264SessionParametersAddInfoKHR.html)
structure, then this command adds the following parameter entries to
`videoSessionParameters`:

- The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h264-sps"
  target="_blank" rel="noopener">H.264 SPS</a> entries specified in
  [VkVideoEncodeH264SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264SessionParametersAddInfoKHR.html)::`pStdSPSs`.

- The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h264-pps"
  target="_blank" rel="noopener">H.264 PPS</a> entries specified in
  [VkVideoEncodeH264SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264SessionParametersAddInfoKHR.html)::`pStdPPSs`.

If `videoSessionParameters` was created with the video codec operation
`VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR` and the
`pUpdateInfo->pNext` chain includes a
[VkVideoEncodeH265SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265SessionParametersAddInfoKHR.html)
structure, then this command adds the following parameter entries to
`videoSessionParameters`:

- The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265-vps"
  target="_blank" rel="noopener">H.265 VPS</a> entries specified in
  [VkVideoEncodeH265SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265SessionParametersAddInfoKHR.html)::`pStdVPSs`.

- The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265-sps"
  target="_blank" rel="noopener">H.265 SPS</a> entries specified in
  [VkVideoEncodeH265SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265SessionParametersAddInfoKHR.html)::`pStdSPSs`.

- The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265-pps"
  target="_blank" rel="noopener">H.265 PPS</a> entries specified in
  [VkVideoEncodeH265SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265SessionParametersAddInfoKHR.html)::`pStdPPSs`.

In case of video session parameters objects created with a video encode
operation, implementations **may** return the
`VK_ERROR_INVALID_VIDEO_STD_PARAMETERS_KHR` error if any of the
specified Video Std parameters do not adhere to the syntactic or
semantic requirements of the used video compression standard, or if
values derived from parameters according to the rules defined by the
used video compression standard do not adhere to the capabilities of the
video compression standard or the implementation.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>Applications <strong>should</strong> not rely on the
<code>VK_ERROR_INVALID_VIDEO_STD_PARAMETERS_KHR</code> error being
returned by any command as a means to verify Video Std parameters, as
implementations are not required to report the error in any specific set
of cases.</p></td>
</tr>
</tbody>
</table>

Valid Usage

- <a href="#VUID-vkUpdateVideoSessionParametersKHR-pUpdateInfo-07215"
  id="VUID-vkUpdateVideoSessionParametersKHR-pUpdateInfo-07215"></a>
  VUID-vkUpdateVideoSessionParametersKHR-pUpdateInfo-07215  
  `pUpdateInfo->updateSequenceCount` **must** equal the current <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-session-parameters"
  target="_blank" rel="noopener">update sequence counter</a> of
  `videoSessionParameters` plus one

- <a
  href="#VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07216"
  id="VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07216"></a>
  VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07216  
  If `videoSessionParameters` was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_DECODE_H264_BIT_KHR` and the `pNext` chain
  of `pUpdateInfo` includes a
  [VkVideoDecodeH264SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH264SessionParametersAddInfoKHR.html)
  structure, then `videoSessionParameters` **must** not already contain
  a `StdVideoH264SequenceParameterSet` entry with `seq_parameter_set_id`
  matching any of the elements of
  [VkVideoDecodeH264SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH264SessionParametersAddInfoKHR.html)::`pStdSPSs`

- <a
  href="#VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07217"
  id="VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07217"></a>
  VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07217  
  If `videoSessionParameters` was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_DECODE_H264_BIT_KHR`, then the number of
  `StdVideoH264SequenceParameterSet` entries already stored in it plus
  the value of the `stdSPSCount` member of the
  [VkVideoDecodeH264SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH264SessionParametersAddInfoKHR.html)
  structure included in the `pUpdateInfo->pNext` chain **must** be less
  than or equal to the
  [VkVideoDecodeH264SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH264SessionParametersCreateInfoKHR.html)::`maxStdSPSCount`
  `videoSessionParameters` was created with

- <a
  href="#VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07218"
  id="VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07218"></a>
  VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07218  
  If `videoSessionParameters` was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_DECODE_H264_BIT_KHR` and the `pNext` chain
  of `pUpdateInfo` includes a
  [VkVideoDecodeH264SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH264SessionParametersAddInfoKHR.html)
  structure, then `videoSessionParameters` **must** not already contain
  a `StdVideoH264PictureParameterSet` entry with both
  `seq_parameter_set_id` and `pic_parameter_set_id` matching any of the
  elements of
  [VkVideoDecodeH264SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH264SessionParametersAddInfoKHR.html)::`pStdPPSs`

- <a
  href="#VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07219"
  id="VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07219"></a>
  VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07219  
  If `videoSessionParameters` was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_DECODE_H264_BIT_KHR`, then the number of
  `StdVideoH264PictureParameterSet` entries already stored in it plus
  the value of the `stdPPSCount` member of the
  [VkVideoDecodeH264SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH264SessionParametersAddInfoKHR.html)
  structure included in the `pUpdateInfo->pNext` chain **must** be less
  than or equal to the
  [VkVideoDecodeH264SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH264SessionParametersCreateInfoKHR.html)::`maxStdPPSCount`
  `videoSessionParameters` was created with

- <a
  href="#VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07220"
  id="VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07220"></a>
  VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07220  
  If `videoSessionParameters` was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_DECODE_H265_BIT_KHR` and the `pNext` chain
  of `pUpdateInfo` includes a
  [VkVideoDecodeH265SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH265SessionParametersAddInfoKHR.html)
  structure, then `videoSessionParameters` **must** not already contain
  a `StdVideoH265VideoParameterSet` entry with
  `vps_video_parameter_set_id` matching any of the elements of
  [VkVideoDecodeH265SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH265SessionParametersAddInfoKHR.html)::`pStdVPSs`

- <a
  href="#VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07221"
  id="VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07221"></a>
  VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07221  
  If `videoSessionParameters` was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_DECODE_H265_BIT_KHR`, then the number of
  `StdVideoH265VideoParameterSet` entries already stored in it plus the
  value of the `stdVPSCount` member of the
  [VkVideoDecodeH265SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH265SessionParametersAddInfoKHR.html)
  structure included in the `pUpdateInfo->pNext` chain **must** be less
  than or equal to the
  [VkVideoDecodeH265SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH265SessionParametersCreateInfoKHR.html)::`maxStdVPSCount`
  `videoSessionParameters` was created with

- <a
  href="#VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07222"
  id="VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07222"></a>
  VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07222  
  If `videoSessionParameters` was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_DECODE_H265_BIT_KHR` and the `pNext` chain
  of `pUpdateInfo` includes a
  [VkVideoDecodeH265SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH265SessionParametersAddInfoKHR.html)
  structure, then `videoSessionParameters` **must** not already contain
  a `StdVideoH265SequenceParameterSet` entry with both
  `sps_video_parameter_set_id` and `sps_seq_parameter_set_id` matching
  any of the elements of
  [VkVideoDecodeH265SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH265SessionParametersAddInfoKHR.html)::`pStdSPSs`

- <a
  href="#VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07223"
  id="VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07223"></a>
  VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07223  
  If `videoSessionParameters` was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_DECODE_H265_BIT_KHR`, then the number of
  `StdVideoH265SequenceParameterSet` entries already stored in it plus
  the value of the `stdSPSCount` member of the
  [VkVideoDecodeH265SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH265SessionParametersAddInfoKHR.html)
  structure included in the `pUpdateInfo->pNext` chain **must** be less
  than or equal to the
  [VkVideoDecodeH265SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH265SessionParametersCreateInfoKHR.html)::`maxStdSPSCount`
  `videoSessionParameters` was created with

- <a
  href="#VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07224"
  id="VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07224"></a>
  VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07224  
  If `videoSessionParameters` was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_DECODE_H265_BIT_KHR` and the `pNext` chain
  of `pUpdateInfo` includes a
  [VkVideoDecodeH265SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH265SessionParametersAddInfoKHR.html)
  structure, then `videoSessionParameters` **must** not already contain
  a `StdVideoH265PictureParameterSet` entry with
  `sps_video_parameter_set_id`, `pps_seq_parameter_set_id`, and
  `pps_pic_parameter_set_id` all matching any of the elements of
  [VkVideoDecodeH265SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH265SessionParametersAddInfoKHR.html)::`pStdPPSs`

- <a
  href="#VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07225"
  id="VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07225"></a>
  VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07225  
  If `videoSessionParameters` was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_DECODE_H265_BIT_KHR`, then the number of
  `StdVideoH265PictureParameterSet` entries already stored in it plus
  the value of the `stdPPSCount` member of the
  [VkVideoDecodeH265SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH265SessionParametersAddInfoKHR.html)
  structure included in the `pUpdateInfo->pNext` chain **must** be less
  than or equal to the
  [VkVideoDecodeH265SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH265SessionParametersCreateInfoKHR.html)::`maxStdPPSCount`
  `videoSessionParameters` was created with

- <a
  href="#VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-09260"
  id="VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-09260"></a>
  VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-09260  
  `videoSessionParameters` **must** not have been created with the video
  codec operation `VK_VIDEO_CODEC_OPERATION_DECODE_AV1_BIT_KHR`

- <a
  href="#VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07226"
  id="VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07226"></a>
  VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07226  
  If `videoSessionParameters` was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H264_BIT_KHR` and the `pNext` chain
  of `pUpdateInfo` includes a
  [VkVideoEncodeH264SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264SessionParametersAddInfoKHR.html)
  structure, then `videoSessionParameters` **must** not already contain
  a `StdVideoH264SequenceParameterSet` entry with `seq_parameter_set_id`
  matching any of the elements of
  [VkVideoEncodeH264SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264SessionParametersAddInfoKHR.html)::`pStdSPSs`

- <a
  href="#VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-06441"
  id="VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-06441"></a>
  VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-06441  
  If `videoSessionParameters` was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H264_BIT_KHR`, then the number of
  `StdVideoH264SequenceParameterSet` entries already stored in it plus
  the value of the `stdSPSCount` member of the
  [VkVideoEncodeH264SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264SessionParametersAddInfoKHR.html)
  structure included in the `pUpdateInfo->pNext` chain **must** be less
  than or equal to the
  [VkVideoEncodeH264SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264SessionParametersCreateInfoKHR.html)::`maxStdSPSCount`
  `videoSessionParameters` was created with

- <a
  href="#VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07227"
  id="VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07227"></a>
  VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07227  
  If `videoSessionParameters` was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H264_BIT_KHR` and the `pNext` chain
  of `pUpdateInfo` includes a
  [VkVideoEncodeH264SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264SessionParametersAddInfoKHR.html)
  structure, then `videoSessionParameters` **must** not already contain
  a `StdVideoH264PictureParameterSet` entry with both
  `seq_parameter_set_id` and `pic_parameter_set_id` matching any of the
  elements of
  [VkVideoEncodeH264SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264SessionParametersAddInfoKHR.html)::`pStdPPSs`

- <a
  href="#VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-06442"
  id="VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-06442"></a>
  VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-06442  
  If `videoSessionParameters` was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H264_BIT_KHR`, then the number of
  `StdVideoH264PictureParameterSet` entries already stored in it plus
  the value of the `stdPPSCount` member of the
  [VkVideoEncodeH264SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264SessionParametersAddInfoKHR.html)
  structure included in the `pUpdateInfo->pNext` chain **must** be less
  than or equal to the
  [VkVideoEncodeH264SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264SessionParametersCreateInfoKHR.html)::`maxStdPPSCount`
  `videoSessionParameters` was created with

- <a
  href="#VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07228"
  id="VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07228"></a>
  VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07228  
  If `videoSessionParameters` was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR` and the `pNext` chain
  of `pUpdateInfo` includes a
  [VkVideoEncodeH265SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265SessionParametersAddInfoKHR.html)
  structure, then `videoSessionParameters` **must** not already contain
  a `StdVideoH265VideoParameterSet` entry with
  `vps_video_parameter_set_id` matching any of the elements of
  [VkVideoEncodeH265SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265SessionParametersAddInfoKHR.html)::`pStdVPSs`

- <a
  href="#VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-06443"
  id="VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-06443"></a>
  VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-06443  
  If `videoSessionParameters` was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR`, then the number of
  `StdVideoH265VideoParameterSet` entries already stored in it plus the
  value of the `stdVPSCount` member of the
  [VkVideoEncodeH265SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265SessionParametersAddInfoKHR.html)
  structure included in the `pUpdateInfo->pNext` chain **must** be less
  than or equal to the
  [VkVideoEncodeH265SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265SessionParametersCreateInfoKHR.html)::`maxStdVPSCount`
  `videoSessionParameters` was created with

- <a
  href="#VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07229"
  id="VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07229"></a>
  VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07229  
  If `videoSessionParameters` was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR` and the `pNext` chain
  of `pUpdateInfo` includes a
  [VkVideoEncodeH265SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265SessionParametersAddInfoKHR.html)
  structure, then `videoSessionParameters` **must** not already contain
  a `StdVideoH265SequenceParameterSet` entry with both
  `sps_video_parameter_set_id` and `sps_seq_parameter_set_id` matching
  any of the elements of
  [VkVideoEncodeH265SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265SessionParametersAddInfoKHR.html)::`pStdSPSs`

- <a
  href="#VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-06444"
  id="VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-06444"></a>
  VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-06444  
  If `videoSessionParameters` was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR`, then the number of
  `StdVideoH265SequenceParameterSet` entries already stored in it plus
  the value of the `stdSPSCount` member of the
  [VkVideoEncodeH265SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265SessionParametersAddInfoKHR.html)
  structure included in the `pUpdateInfo->pNext` chain **must** be less
  than or equal to the
  [VkVideoEncodeH265SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265SessionParametersCreateInfoKHR.html)::`maxStdSPSCount`
  `videoSessionParameters` was created with

- <a
  href="#VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07230"
  id="VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07230"></a>
  VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-07230  
  If `videoSessionParameters` was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR` and the `pNext` chain
  of `pUpdateInfo` includes a
  [VkVideoEncodeH265SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265SessionParametersAddInfoKHR.html)
  structure, then `videoSessionParameters` **must** not already contain
  a `StdVideoH265PictureParameterSet` entry with
  `sps_video_parameter_set_id`, `pps_seq_parameter_set_id`, and
  `pps_pic_parameter_set_id` all matching any of the elements of
  [VkVideoEncodeH265SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265SessionParametersAddInfoKHR.html)::`pStdPPSs`

- <a
  href="#VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-06445"
  id="VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-06445"></a>
  VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-06445  
  If `videoSessionParameters` was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR`, then the number of
  `StdVideoH265PictureParameterSet` entries already stored in it plus
  the value of the `stdPPSCount` member of the
  [VkVideoEncodeH265SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265SessionParametersAddInfoKHR.html)
  structure included in the `pUpdateInfo->pNext` chain **must** be less
  than or equal to the
  [VkVideoEncodeH265SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265SessionParametersCreateInfoKHR.html)::`maxStdPPSCount`
  `videoSessionParameters` was created with

- <a
  href="#VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-08321"
  id="VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-08321"></a>
  VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-08321  
  If `videoSessionParameters` was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR` and the `pNext` chain
  of `pUpdateInfo` includes a
  [VkVideoEncodeH265SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265SessionParametersAddInfoKHR.html)
  structure, then `num_tile_columns_minus1` **must** be less than
  [VkVideoEncodeH265CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265CapabilitiesKHR.html)::`maxTiles.width`,
  as returned by
  [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)
  for the video profile `videoSessionParameters` was created with, for
  each element of
  [VkVideoEncodeH265SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265SessionParametersAddInfoKHR.html)::`pStdPPSs`

- <a
  href="#VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-08322"
  id="VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-08322"></a>
  VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-08322  
  If `videoSessionParameters` was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR` and the `pNext` chain
  of `pUpdateInfo` includes a
  [VkVideoEncodeH265SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265SessionParametersAddInfoKHR.html)
  structure, then `num_tile_rows_minus1` **must** be less than
  [VkVideoEncodeH265CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265CapabilitiesKHR.html)::`maxTiles.height`,
  as returned by
  [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)
  for the video profile `videoSessionParameters` was created with, for
  each element of
  [VkVideoEncodeH265SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265SessionParametersAddInfoKHR.html)::`pStdPPSs`

Valid Usage (Implicit)

- <a href="#VUID-vkUpdateVideoSessionParametersKHR-device-parameter"
  id="VUID-vkUpdateVideoSessionParametersKHR-device-parameter"></a>
  VUID-vkUpdateVideoSessionParametersKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a
  href="#VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-parameter"
  id="VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-parameter"></a>
  VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-parameter  
  `videoSessionParameters` **must** be a valid
  [VkVideoSessionParametersKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionParametersKHR.html) handle

- <a href="#VUID-vkUpdateVideoSessionParametersKHR-pUpdateInfo-parameter"
  id="VUID-vkUpdateVideoSessionParametersKHR-pUpdateInfo-parameter"></a>
  VUID-vkUpdateVideoSessionParametersKHR-pUpdateInfo-parameter  
  `pUpdateInfo` **must** be a valid pointer to a valid
  [VkVideoSessionParametersUpdateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionParametersUpdateInfoKHR.html)
  structure

- <a
  href="#VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-parent"
  id="VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-parent"></a>
  VUID-vkUpdateVideoSessionParametersKHR-videoSessionParameters-parent  
  `videoSessionParameters` **must** have been created, allocated, or
  retrieved from `device`

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

- `VK_ERROR_INVALID_VIDEO_STD_PARAMETERS_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_queue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_queue.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkVideoSessionParametersKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionParametersKHR.html),
[VkVideoSessionParametersUpdateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionParametersUpdateInfoKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkUpdateVideoSessionParametersKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
