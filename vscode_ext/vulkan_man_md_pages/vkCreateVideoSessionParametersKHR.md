# vkCreateVideoSessionParametersKHR(3) Manual Page

## Name

vkCreateVideoSessionParametersKHR - Creates video session parameters
object



## <a href="#_c_specification" class="anchor"></a>C Specification

To create a video session parameters object, call:

``` c
// Provided by VK_KHR_video_queue
VkResult vkCreateVideoSessionParametersKHR(
    VkDevice                                    device,
    const VkVideoSessionParametersCreateInfoKHR* pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkVideoSessionParametersKHR*                pVideoSessionParameters);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that creates the video session
  parameters object.

- `pCreateInfo` is a pointer to
  [VkVideoSessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionParametersCreateInfoKHR.html)
  structure containing parameters to be used to create the video session
  parameters object.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

- `pVideoSessionParameters` is a pointer to a
  [VkVideoSessionParametersKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionParametersKHR.html) handle
  in which the resulting video session parameters object is returned.

## <a href="#_description" class="anchor"></a>Description

The resulting video session parameters object is said to be created with
the video codec operation `pCreateInfo->videoSession` was created with.

Video session parameters objects created with an encode operation are
always created with respect to a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-quality-level"
target="_blank" rel="noopener">video encode quality level</a>. By
default, the created video session parameters objects are created with
quality level zero, unless otherwise specified by including a
[VkVideoEncodeQualityLevelInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeQualityLevelInfoKHR.html)
structure in the `pCreateInfo->pNext` chain, in which case the video
session parameters object is created with the quality level specified in
[VkVideoEncodeQualityLevelInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeQualityLevelInfoKHR.html)::`qualityLevel`.

If `pCreateInfo->videoSessionParametersTemplate` is not
`VK_NULL_HANDLE`, then it will be used as a template for constructing
the new video session parameters object. This happens by first adding
any parameters according to the additional creation parameters provided
in the `pCreateInfo->pNext` chain, followed by adding any parameters
from the template object that have a key that does not match the key of
any of the already added parameters.

For video session parameters objects created with an encode operation,
the template object specified in
`pCreateInfo->videoSessionParametersTemplate` **must** have been created
with the same <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-quality-level"
target="_blank" rel="noopener">video encode quality level</a> as the
newly created object.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>This means that codec-specific parameters stored in video session
parameters objects <strong>can</strong> only be reused across different
video encode quality levels by re-specifying them, as previously created
video session parameters against other quality levels
<strong>cannot</strong> be used as template because the original
codec-specific parameters (before the implementation
<strong>may</strong> have applied <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-overrides"
target="_blank" rel="noopener">parameter overrides</a>)
<strong>may</strong> no longer be available in them for the purposes of
constructing the derived object.</p></td>
</tr>
</tbody>
</table>

If `pCreateInfo->videoSession` was created with the video codec
operation `VK_VIDEO_CODEC_OPERATION_DECODE_H264_BIT_KHR`, then the
created video session parameters object will initially contain the
following sets of parameter entries:

- `StdVideoH264SequenceParameterSet` structures representing <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h264-sps"
  target="_blank" rel="noopener">H.264 SPS</a> entries, as follows:

  - If the `pParametersAddInfo` member of the
    [VkVideoDecodeH264SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH264SessionParametersCreateInfoKHR.html)
    structure provided in the `pCreateInfo->pNext` chain is not `NULL`,
    then the set of `StdVideoH264SequenceParameterSet` entries specified
    in `pParametersAddInfo->pStdSPSs` are added first;

  - If `pCreateInfo->videoSessionParametersTemplate` is not
    `VK_NULL_HANDLE`, then each `StdVideoH264SequenceParameterSet` entry
    stored in it is copied to the created video session parameters
    object if the created object does not already contain such an entry
    with the same `seq_parameter_set_id`.

- `StdVideoH264PictureParameterSet` structures representing <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h264-pps"
  target="_blank" rel="noopener">H.264 PPS</a> entries, as follows:

  - If the `pParametersAddInfo` member of the
    [VkVideoDecodeH264SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH264SessionParametersCreateInfoKHR.html)
    structure provided in the `pCreateInfo->pNext` chain is not `NULL`,
    then the set of `StdVideoH264PictureParameterSet` entries specified
    in `pParametersAddInfo->pStdPPSs` are added first;

  - If `pCreateInfo->videoSessionParametersTemplate` is not
    `VK_NULL_HANDLE`, then each `StdVideoH264PictureParameterSet` entry
    stored in it is copied to the created video session parameters
    object if the created object does not already contain such an entry
    with the same `seq_parameter_set_id` and `pic_parameter_set_id`.

If `pCreateInfo->videoSession` was created with the video codec
operation `VK_VIDEO_CODEC_OPERATION_DECODE_H265_BIT_KHR`, then the
created video session parameters object will initially contain the
following sets of parameter entries:

- `StdVideoH265VideoParameterSet` structures representing <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h265-vps"
  target="_blank" rel="noopener">H.265 VPS</a> entries, as follows:

  - If the `pParametersAddInfo` member of the
    [VkVideoDecodeH265SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH265SessionParametersCreateInfoKHR.html)
    structure provided in the `pCreateInfo->pNext` chain is not `NULL`,
    then the set of `StdVideoH265VideoParameterSet` entries specified in
    `pParametersAddInfo->pStdVPSs` are added first;

  - If `pCreateInfo->videoSessionParametersTemplate` is not
    `VK_NULL_HANDLE`, then each `StdVideoH265VideoParameterSet` entry
    stored in it is copied to the created video session parameters
    object if the created object does not already contain such an entry
    with the same `vps_video_parameter_set_id`.

- `StdVideoH265SequenceParameterSet` structures representing <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h265-sps"
  target="_blank" rel="noopener">H.265 SPS</a> entries, as follows:

  - If the `pParametersAddInfo` member of the
    [VkVideoDecodeH265SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH265SessionParametersCreateInfoKHR.html)
    structure provided in the `pCreateInfo->pNext` chain is not `NULL`,
    then the set of `StdVideoH265SequenceParameterSet` entries specified
    in `pParametersAddInfo->pStdSPSs` are added first;

  - If `pCreateInfo->videoSessionParametersTemplate` is not
    `VK_NULL_HANDLE`, then each `StdVideoH265SequenceParameterSet` entry
    stored in it is copied to the created video session parameters
    object if the created object does not already contain such an entry
    with the same `sps_video_parameter_set_id` and
    `sps_seq_parameter_set_id`.

- `StdVideoH265PictureParameterSet` structures representing <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h265-pps"
  target="_blank" rel="noopener">H.265 PPS</a> entries, as follows:

  - If the `pParametersAddInfo` member of the
    [VkVideoDecodeH265SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH265SessionParametersCreateInfoKHR.html)
    structure provided in the `pCreateInfo->pNext` chain is not `NULL`,
    then the set of `StdVideoH265PictureParameterSet` entries specified
    in `pParametersAddInfo->pStdPPSs` are added first;

  - If `pCreateInfo->videoSessionParametersTemplate` is not
    `VK_NULL_HANDLE`, then each `StdVideoH265PictureParameterSet` entry
    stored in it is copied to the created video session parameters
    object if the created object does not already contain such an entry
    with the same `sps_video_parameter_set_id`,
    `pps_seq_parameter_set_id`, and `pps_pic_parameter_set_id`.

If `pCreateInfo->videoSession` was created with the video codec
operation `VK_VIDEO_CODEC_OPERATION_DECODE_AV1_BIT_KHR`, then the
created video session parameters object will contain a single <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-av1-sequence-header"
target="_blank" rel="noopener">AV1 sequence header</a> represented by a
`StdVideoAV1SequenceHeader` structure specified through the
`pStdSequenceHeader` member of the
[VkVideoDecodeAV1SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeAV1SessionParametersCreateInfoKHR.html)
structure provided in the `pCreateInfo->pNext` chain. As such video
session parameters objects **can** only contain a single <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-av1-sequence-header"
target="_blank" rel="noopener">AV1 sequence header</a>, it is not
possible to use a previously created object as a template or
subsequently update the created video session parameters object.

If `pCreateInfo->videoSession` was created with the video codec
operation `VK_VIDEO_CODEC_OPERATION_ENCODE_H264_BIT_KHR`, then the
created video session parameters object will initially contain the
following sets of parameter entries:

- `StdVideoH264SequenceParameterSet` structures representing <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h264-sps"
  target="_blank" rel="noopener">H.264 SPS</a> entries, as follows:

  - If the `pParametersAddInfo` member of the
    [VkVideoEncodeH264SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264SessionParametersCreateInfoKHR.html)
    structure provided in the `pCreateInfo->pNext` chain is not `NULL`,
    then the set of `StdVideoH264SequenceParameterSet` entries specified
    in `pParametersAddInfo->pStdSPSs` are added first;

  - If `pCreateInfo->videoSessionParametersTemplate` is not
    `VK_NULL_HANDLE`, then each `StdVideoH264SequenceParameterSet` entry
    stored in it is copied to the created video session parameters
    object if the created object does not already contain such an entry
    with the same `seq_parameter_set_id`.

- `StdVideoH264PictureParameterSet` structures representing <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h264-pps"
  target="_blank" rel="noopener">H.264 PPS</a> entries, as follows:

  - If the `pParametersAddInfo` member of the
    [VkVideoEncodeH264SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264SessionParametersCreateInfoKHR.html)
    structure provided in the `pCreateInfo->pNext` chain is not `NULL`,
    then the set of `StdVideoH264PictureParameterSet` entries specified
    in `pParametersAddInfo->pStdPPSs` are added first;

  - If `pCreateInfo->videoSessionParametersTemplate` is not
    `VK_NULL_HANDLE`, then each `StdVideoH264PictureParameterSet` entry
    stored in it is copied to the created video session parameters
    object if the created object does not already contain such an entry
    with the same `seq_parameter_set_id` and `pic_parameter_set_id`.

If `pCreateInfo->videoSession` was created with the video codec
operation `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR`, then the
created video session parameters object will initially contain the
following sets of parameter entries:

- `StdVideoH265VideoParameterSet` structures representing <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265-vps"
  target="_blank" rel="noopener">H.265 VPS</a> entries, as follows:

  - If the `pParametersAddInfo` member of the
    [VkVideoEncodeH265SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265SessionParametersCreateInfoKHR.html)
    structure provided in the `pCreateInfo->pNext` chain is not `NULL`,
    then the set of `StdVideoH265VideoParameterSet` entries specified in
    `pParametersAddInfo->pStdVPSs` are added first;

  - If `pCreateInfo->videoSessionParametersTemplate` is not
    `VK_NULL_HANDLE`, then each `StdVideoH265VideoParameterSet` entry
    stored in it is copied to the created video session parameters
    object if the created object does not already contain such an entry
    with the same `vps_video_parameter_set_id`.

- `StdVideoH265SequenceParameterSet` structures representing <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265-sps"
  target="_blank" rel="noopener">H.265 SPS</a> entries, as follows:

  - If the `pParametersAddInfo` member of the
    [VkVideoEncodeH265SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265SessionParametersCreateInfoKHR.html)
    structure provided in the `pCreateInfo->pNext` chain is not `NULL`,
    then the set of `StdVideoH265SequenceParameterSet` entries specified
    in `pParametersAddInfo->pStdSPSs` are added first;

  - If `pCreateInfo->videoSessionParametersTemplate` is not
    `VK_NULL_HANDLE`, then each `StdVideoH265SequenceParameterSet` entry
    stored in it is copied to the created video session parameters
    object if the created object does not already contain such an entry
    with the same `sps_video_parameter_set_id` and
    `sps_seq_parameter_set_id`.

- `StdVideoH265PictureParameterSet` structures representing <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265-pps"
  target="_blank" rel="noopener">H.265 PPS</a> entries, as follows:

  - If the `pParametersAddInfo` member of the
    [VkVideoEncodeH265SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265SessionParametersCreateInfoKHR.html)
    structure provided in the `pCreateInfo->pNext` chain is not `NULL`,
    then the set of `StdVideoH265PictureParameterSet` entries specified
    in `pParametersAddInfo->pStdPPSs` are added first;

  - If `pCreateInfo->videoSessionParametersTemplate` is not
    `VK_NULL_HANDLE`, then each `StdVideoH265PictureParameterSet` entry
    stored in it is copied to the created video session parameters
    object if the created object does not already contain such an entry
    with the same `sps_video_parameter_set_id`,
    `pps_seq_parameter_set_id`, and `pps_pic_parameter_set_id`.

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

Valid Usage (Implicit)

- <a href="#VUID-vkCreateVideoSessionParametersKHR-device-parameter"
  id="VUID-vkCreateVideoSessionParametersKHR-device-parameter"></a>
  VUID-vkCreateVideoSessionParametersKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkCreateVideoSessionParametersKHR-pCreateInfo-parameter"
  id="VUID-vkCreateVideoSessionParametersKHR-pCreateInfo-parameter"></a>
  VUID-vkCreateVideoSessionParametersKHR-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid
  [VkVideoSessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionParametersCreateInfoKHR.html)
  structure

- <a href="#VUID-vkCreateVideoSessionParametersKHR-pAllocator-parameter"
  id="VUID-vkCreateVideoSessionParametersKHR-pAllocator-parameter"></a>
  VUID-vkCreateVideoSessionParametersKHR-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a
  href="#VUID-vkCreateVideoSessionParametersKHR-pVideoSessionParameters-parameter"
  id="VUID-vkCreateVideoSessionParametersKHR-pVideoSessionParameters-parameter"></a>
  VUID-vkCreateVideoSessionParametersKHR-pVideoSessionParameters-parameter  
  `pVideoSessionParameters` **must** be a valid pointer to a
  [VkVideoSessionParametersKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionParametersKHR.html) handle

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

- `VK_ERROR_INITIALIZATION_FAILED`

- `VK_ERROR_INVALID_VIDEO_STD_PARAMETERS_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_queue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_queue.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkVideoSessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionParametersCreateInfoKHR.html),
[VkVideoSessionParametersKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionParametersKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCreateVideoSessionParametersKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
