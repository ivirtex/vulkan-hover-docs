# VkVideoSessionParametersCreateInfoKHR(3) Manual Page

## Name

VkVideoSessionParametersCreateInfoKHR - Structure specifying parameters
of a newly created video session parameters object



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkVideoSessionParametersCreateInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_video_queue
typedef struct VkVideoSessionParametersCreateInfoKHR {
    VkStructureType                           sType;
    const void*                               pNext;
    VkVideoSessionParametersCreateFlagsKHR    flags;
    VkVideoSessionParametersKHR               videoSessionParametersTemplate;
    VkVideoSessionKHR                         videoSession;
} VkVideoSessionParametersCreateInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is reserved for future use.

- `videoSessionParametersTemplate` is `VK_NULL_HANDLE` or a valid handle
  to a [VkVideoSessionParametersKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionParametersKHR.html)
  object used as a template for constructing the new video session
  parameters object.

- `videoSession` is the video session object against which the video
  session parameters object is going to be created.

## <a href="#_description" class="anchor"></a>Description

Limiting values are defined below that are referenced by the relevant
valid usage statements of this structure.

- If `videoSession` was created with the codec operation
  `VK_VIDEO_CODEC_OPERATION_DECODE_H264_BIT_KHR`, then let
  `StdVideoH264SequenceParameterSet spsAddList[]` be the list of <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h264-sps"
  target="_blank" rel="noopener">H.264 SPS</a> entries to add to the
  created video session parameters object, defined as follows:

  - If the `pParametersAddInfo` member of the
    [VkVideoDecodeH264SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH264SessionParametersCreateInfoKHR.html)
    structure provided in the `pNext` chain is not `NULL`, then the set
    of `StdVideoH264SequenceParameterSet` entries specified in
    `pParametersAddInfo->pStdSPSs` are added to `spsAddList`;

  - If `videoSessionParametersTemplate` is not `VK_NULL_HANDLE`, then
    each `StdVideoH264SequenceParameterSet` entry stored in it with
    `seq_parameter_set_id` not matching any of the entries already in
    `spsAddList` is added to `spsAddList`.

- If `videoSession` was created with the codec operation
  `VK_VIDEO_CODEC_OPERATION_DECODE_H264_BIT_KHR`, then let
  `StdVideoH264PictureParameterSet ppsAddList[]` be the list of <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h264-pps"
  target="_blank" rel="noopener">H.264 PPS</a> entries to add to the
  created video session parameters object, defined as follows:

  - If the `pParametersAddInfo` member of the
    [VkVideoDecodeH264SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH264SessionParametersCreateInfoKHR.html)
    structure provided in the `pNext` chain is not `NULL`, then the set
    of `StdVideoH264PictureParameterSet` entries specified in
    `pParametersAddInfo->pStdPPSs` are added to `ppsAddList`;

  - If `videoSessionParametersTemplate` is not `VK_NULL_HANDLE`, then
    each `StdVideoH264PictureParameterSet` entry stored in it with
    `seq_parameter_set_id` or `pic_parameter_set_id` not matching any of
    the entries already in `ppsAddList` is added to `ppsAddList`.

- If `videoSession` was created with the codec operation
  `VK_VIDEO_CODEC_OPERATION_DECODE_H265_BIT_KHR`, then let
  `StdVideoH265VideoParameterSet vpsAddList[]` be the list of <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h265-vps"
  target="_blank" rel="noopener">H.265 VPS</a> entries to add to the
  created video session parameters object, defined as follows:

  - If the `pParametersAddInfo` member of the
    [VkVideoDecodeH265SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH265SessionParametersCreateInfoKHR.html)
    structure provided in the `pNext` chain is not `NULL`, then the set
    of `StdVideoH265VideoParameterSet` entries specified in
    `pParametersAddInfo->pStdVPSs` are added to `vpsAddList`;

  - If `videoSessionParametersTemplate` is not `VK_NULL_HANDLE`, then
    each `StdVideoH265VideoParameterSet` entry stored in it with
    `vps_video_parameter_set_id` not matching any of the entries already
    in `vpsAddList` is added to `vpsAddList`.

- If `videoSession` was created with the codec operation
  `VK_VIDEO_CODEC_OPERATION_DECODE_H265_BIT_KHR`, then let
  `StdVideoH265SequenceParameterSet spsAddList[]` be the list of <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h265-sps"
  target="_blank" rel="noopener">H.265 SPS</a> entries to add to the
  created video session parameters object, defined as follows:

  - If the `pParametersAddInfo` member of the
    [VkVideoDecodeH265SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH265SessionParametersCreateInfoKHR.html)
    structure provided in the `pNext` chain is not `NULL`, then the set
    of `StdVideoH265SequenceParameterSet` entries specified in
    `pParametersAddInfo->pStdSPSs` are added to `spsAddList`;

  - If `videoSessionParametersTemplate` is not `VK_NULL_HANDLE`, then
    each `StdVideoH265SequenceParameterSet` entry stored in it with
    `sps_video_parameter_set_id` or `sps_seq_parameter_set_id` not
    matching any of the entries already in `spsAddList` is added to
    `spsAddList`.

- If `videoSession` was created with the codec operation
  `VK_VIDEO_CODEC_OPERATION_DECODE_H265_BIT_KHR`, then let
  `StdVideoH265PictureParameterSet ppsAddList[]` be the list of <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h265-pps"
  target="_blank" rel="noopener">H.265 PPS</a> entries to add to the
  created video session parameters object, defined as follows:

  - If the `pParametersAddInfo` member of the
    [VkVideoDecodeH265SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH265SessionParametersCreateInfoKHR.html)
    structure provided in the `pNext` chain is not `NULL`, then the set
    of `StdVideoH265PictureParameterSet` entries specified in
    `pParametersAddInfo->pStdPPSs` are added to `ppsAddList`;

  - If `videoSessionParametersTemplate` is not `VK_NULL_HANDLE`, then
    each `StdVideoH265PictureParameterSet` entry stored in it with
    `sps_video_parameter_set_id`, `pps_seq_parameter_set_id`, or
    `pps_pic_parameter_set_id` not matching any of the entries already
    in `ppsAddList` is added to `ppsAddList`.

- If `videoSession` was created with an encode operation, then let
  `uint32_t qualityLevel` be the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-quality-level"
  target="_blank" rel="noopener">video encode quality level</a> of the
  created video session parameters object, defined as follows:

  - If the `pNext` chain of this structure includes a
    [VkVideoEncodeQualityLevelInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeQualityLevelInfoKHR.html)
    structure, then `qualityLevel` is equal to
    [VkVideoEncodeQualityLevelInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeQualityLevelInfoKHR.html)::`qualityLevel`.

  - Otherwise `qualityLevel` is `0`

- If `videoSession` was created with the codec operation
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H264_BIT_KHR`, then let
  `StdVideoH264SequenceParameterSet spsAddList[]` be the list of <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h264-sps"
  target="_blank" rel="noopener">H.264 SPS</a> entries to add to the
  created video session parameters object, defined as follows:

  - If the `pParametersAddInfo` member of the
    [VkVideoEncodeH264SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264SessionParametersCreateInfoKHR.html)
    structure provided in the `pNext` chain is not `NULL`, then the set
    of `StdVideoH264SequenceParameterSet` entries specified in
    `pParametersAddInfo->pStdSPSs` are added to `spsAddList`;

  - If `videoSessionParametersTemplate` is not `VK_NULL_HANDLE`, then
    each `StdVideoH264SequenceParameterSet` entry stored in it with
    `seq_parameter_set_id` not matching any of the entries already in
    `spsAddList` is added to `spsAddList`.

- If `videoSession` was created with the codec operation
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H264_BIT_KHR`, then let
  `StdVideoH264PictureParameterSet ppsAddList[]` be the list of <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h264-pps"
  target="_blank" rel="noopener">H.264 PPS</a> entries to add to the
  created video session parameters object, defined as follows:

  - If the `pParametersAddInfo` member of the
    [VkVideoEncodeH264SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264SessionParametersCreateInfoKHR.html)
    structure provided in the `pNext` chain is not `NULL`, then the set
    of `StdVideoH264PictureParameterSet` entries specified in
    `pParametersAddInfo->pStdPPSs` are added to `ppsAddList`;

  - If `videoSessionParametersTemplate` is not `VK_NULL_HANDLE`, then
    each `StdVideoH264PictureParameterSet` entry stored in it with
    `seq_parameter_set_id` or `pic_parameter_set_id` not matching any of
    the entries already in `ppsAddList` is added to `ppsAddList`.

- If `videoSession` was created with the codec operation
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR`, then let
  `StdVideoH265VideoParameterSet vpsAddList[]` be the list of <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265-vps"
  target="_blank" rel="noopener">H.265 VPS</a> entries to add to the
  created video session parameters object, defined as follows:

  - If the `pParametersAddInfo` member of the
    [VkVideoEncodeH265SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265SessionParametersCreateInfoKHR.html)
    structure provided in the `pNext` chain is not `NULL`, then the set
    of `StdVideoH265VideoParameterSet` entries specified in
    `pParametersAddInfo->pStdVPSs` are added to `vpsAddList`;

  - If `videoSessionParametersTemplate` is not `VK_NULL_HANDLE`, then
    each `StdVideoH265VideoParameterSet` entry stored in it with
    `vps_video_parameter_set_id` not matching any of the entries already
    in `vpsAddList` is added to `vpsAddList`.

- If `videoSession` was created with the codec operation
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR`, then let
  `StdVideoH265SequenceParameterSet spsAddList[]` be the list of <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265-sps"
  target="_blank" rel="noopener">H.265 SPS</a> entries to add to the
  created video session parameters object, defined as follows:

  - If the `pParametersAddInfo` member of the
    [VkVideoEncodeH265SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265SessionParametersCreateInfoKHR.html)
    structure provided in the `pNext` chain is not `NULL`, then the set
    of `StdVideoH265SequenceParameterSet` entries specified in
    `pParametersAddInfo->pStdSPSs` are added to `spsAddList`;

  - If `videoSessionParametersTemplate` is not `VK_NULL_HANDLE`, then
    each `StdVideoH265SequenceParameterSet` entry stored in it with
    `sps_video_parameter_set_id` or `sps_seq_parameter_set_id` not
    matching any of the entries already in `spsAddList` is added to
    `spsAddList`.

- If `videoSession` was created with the codec operation
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR`, then let
  `StdVideoH265PictureParameterSet ppsAddList[]` be the list of <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265-pps"
  target="_blank" rel="noopener">H.265 PPS</a> entries to add to the
  created video session parameters object, defined as follows:

  - If the `pParametersAddInfo` member of the
    [VkVideoEncodeH265SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265SessionParametersCreateInfoKHR.html)
    structure provided in the `pNext` chain is not `NULL`, then the set
    of `StdVideoH265PictureParameterSet` entries specified in
    `pParametersAddInfo->pStdPPSs` are added to `ppsAddList`;

  - If `videoSessionParametersTemplate` is not `VK_NULL_HANDLE`, then
    each `StdVideoH265PictureParameterSet` entry stored in it with
    `sps_video_parameter_set_id`, `pps_seq_parameter_set_id`, or
    `pps_pic_parameter_set_id` not matching any of the entries already
    in `ppsAddList` is added to `ppsAddList`.

Valid Usage

- <a
  href="#VUID-VkVideoSessionParametersCreateInfoKHR-videoSessionParametersTemplate-04855"
  id="VUID-VkVideoSessionParametersCreateInfoKHR-videoSessionParametersTemplate-04855"></a>
  VUID-VkVideoSessionParametersCreateInfoKHR-videoSessionParametersTemplate-04855  
  If `videoSessionParametersTemplate` is not `VK_NULL_HANDLE`, it
  **must** have been created against `videoSession`

- <a
  href="#VUID-VkVideoSessionParametersCreateInfoKHR-videoSessionParametersTemplate-08310"
  id="VUID-VkVideoSessionParametersCreateInfoKHR-videoSessionParametersTemplate-08310"></a>
  VUID-VkVideoSessionParametersCreateInfoKHR-videoSessionParametersTemplate-08310  
  If `videoSessionParametersTemplate` is not `VK_NULL_HANDLE` and
  `videoSession` was created with an encode operation, then
  `qualityLevel` **must** equal the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-quality-level"
  target="_blank" rel="noopener">video encode quality</a> level
  `videoSessionParametersTemplate` was created with

- <a href="#VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-07203"
  id="VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-07203"></a>
  VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-07203  
  If `videoSession` was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_DECODE_H264_BIT_KHR`, then the `pNext` chain
  **must** include a
  [VkVideoDecodeH264SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH264SessionParametersCreateInfoKHR.html)
  structure

- <a href="#VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-07204"
  id="VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-07204"></a>
  VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-07204  
  If `videoSession` was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_DECODE_H264_BIT_KHR`, then the number of
  elements of `spsAddList` **must** be less than or equal to the
  `maxStdSPSCount` specified in the
  [VkVideoDecodeH264SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH264SessionParametersCreateInfoKHR.html)
  structure included in the `pNext` chain

- <a href="#VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-07205"
  id="VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-07205"></a>
  VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-07205  
  If `videoSession` was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_DECODE_H264_BIT_KHR`, then the number of
  elements of `ppsAddList` **must** be less than or equal to the
  `maxStdPPSCount` specified in the
  [VkVideoDecodeH264SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH264SessionParametersCreateInfoKHR.html)
  structure included in the `pNext` chain

- <a href="#VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-07206"
  id="VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-07206"></a>
  VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-07206  
  If `videoSession` was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_DECODE_H265_BIT_KHR`, then the `pNext` chain
  **must** include a
  [VkVideoDecodeH265SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH265SessionParametersCreateInfoKHR.html)
  structure

- <a href="#VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-07207"
  id="VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-07207"></a>
  VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-07207  
  If `videoSession` was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_DECODE_H265_BIT_KHR`, then the number of
  elements of `vpsAddList` **must** be less than or equal to the
  `maxStdVPSCount` specified in the
  [VkVideoDecodeH265SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH265SessionParametersCreateInfoKHR.html)
  structure included in the `pNext` chain

- <a href="#VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-07208"
  id="VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-07208"></a>
  VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-07208  
  If `videoSession` was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_DECODE_H265_BIT_KHR`, then the number of
  elements of `spsAddList` **must** be less than or equal to the
  `maxStdSPSCount` specified in the
  [VkVideoDecodeH265SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH265SessionParametersCreateInfoKHR.html)
  structure included in the `pNext` chain

- <a href="#VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-07209"
  id="VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-07209"></a>
  VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-07209  
  If `videoSession` was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_DECODE_H265_BIT_KHR`, then the number of
  elements of `ppsAddList` **must** be less than or equal to the
  `maxStdPPSCount` specified in the
  [VkVideoDecodeH265SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH265SessionParametersCreateInfoKHR.html)
  structure included in the `pNext` chain

- <a href="#VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-09258"
  id="VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-09258"></a>
  VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-09258  
  If `videoSession` was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_DECODE_AV1_BIT_KHR`, then
  `videoSessionParametersTemplate` **must** be `VK_NULL_HANDLE`

- <a href="#VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-09259"
  id="VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-09259"></a>
  VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-09259  
  If `videoSession` was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_DECODE_AV1_BIT_KHR`, then the `pNext` chain
  **must** include a
  [VkVideoDecodeAV1SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeAV1SessionParametersCreateInfoKHR.html)
  structure

- <a href="#VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-07210"
  id="VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-07210"></a>
  VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-07210  
  If `videoSession` was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H264_BIT_KHR`, then the `pNext` chain
  **must** include a
  [VkVideoEncodeH264SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264SessionParametersCreateInfoKHR.html)
  structure

- <a href="#VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-04839"
  id="VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-04839"></a>
  VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-04839  
  If `videoSession` was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H264_BIT_KHR`, then the number of
  elements of `spsAddList` **must** be less than or equal to the
  `maxStdSPSCount` specified in the
  [VkVideoEncodeH264SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264SessionParametersCreateInfoKHR.html)
  structure included in the `pNext` chain

- <a href="#VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-04840"
  id="VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-04840"></a>
  VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-04840  
  If `videoSession` was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H264_BIT_KHR`, then the number of
  elements of `ppsAddList` **must** be less than or equal to the
  `maxStdPPSCount` specified in the
  [VkVideoEncodeH264SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264SessionParametersCreateInfoKHR.html)
  structure included in the `pNext` chain

- <a href="#VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-07211"
  id="VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-07211"></a>
  VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-07211  
  If `videoSession` was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR`, then the `pNext` chain
  **must** include a
  [VkVideoEncodeH265SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265SessionParametersCreateInfoKHR.html)
  structure

- <a href="#VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-04841"
  id="VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-04841"></a>
  VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-04841  
  If `videoSession` was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR`, then the number of
  elements of `vpsAddList` **must** be less than or equal to the
  `maxStdVPSCount` specified in the
  [VkVideoEncodeH265SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265SessionParametersCreateInfoKHR.html)
  structure included in the `pNext` chain

- <a href="#VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-04842"
  id="VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-04842"></a>
  VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-04842  
  If `videoSession` was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR`, then the number of
  elements of `spsAddList` **must** be less than or equal to the
  `maxStdSPSCount` specified in the
  [VkVideoEncodeH265SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265SessionParametersCreateInfoKHR.html)
  structure included in the `pNext` chain

- <a href="#VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-04843"
  id="VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-04843"></a>
  VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-04843  
  If `videoSession` was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR`, then the number of
  elements of `ppsAddList` **must** be less than or equal to the
  `maxStdPPSCount` specified in the
  [VkVideoEncodeH265SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265SessionParametersCreateInfoKHR.html)
  structure included in the `pNext` chain

- <a href="#VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-08319"
  id="VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-08319"></a>
  VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-08319  
  If `videoSession` was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR`, then
  `num_tile_columns_minus1` **must** be less than
  [VkVideoEncodeH265CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265CapabilitiesKHR.html)::`maxTiles.width`,
  as returned by
  [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)
  for the video profile `videoSession` was created with, for each
  element of `ppsAddList`

- <a href="#VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-08320"
  id="VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-08320"></a>
  VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-08320  
  If `videoSession` was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR`, then
  `num_tile_rows_minus1` **must** be less than
  [VkVideoEncodeH265CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265CapabilitiesKHR.html)::`maxTiles.height`,
  as returned by
  [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)
  for the video profile `videoSession` was created with, for each
  element of `ppsAddList`

Valid Usage (Implicit)

- <a href="#VUID-VkVideoSessionParametersCreateInfoKHR-sType-sType"
  id="VUID-VkVideoSessionParametersCreateInfoKHR-sType-sType"></a>
  VUID-VkVideoSessionParametersCreateInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_VIDEO_SESSION_PARAMETERS_CREATE_INFO_KHR`

- <a href="#VUID-VkVideoSessionParametersCreateInfoKHR-pNext-pNext"
  id="VUID-VkVideoSessionParametersCreateInfoKHR-pNext-pNext"></a>
  VUID-VkVideoSessionParametersCreateInfoKHR-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the
  `pNext` chain **must** be either `NULL` or a pointer to a valid
  instance of
  [VkVideoDecodeAV1SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeAV1SessionParametersCreateInfoKHR.html),
  [VkVideoDecodeH264SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH264SessionParametersCreateInfoKHR.html),
  [VkVideoDecodeH265SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH265SessionParametersCreateInfoKHR.html),
  [VkVideoEncodeH264SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264SessionParametersCreateInfoKHR.html),
  [VkVideoEncodeH265SessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265SessionParametersCreateInfoKHR.html),
  or
  [VkVideoEncodeQualityLevelInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeQualityLevelInfoKHR.html)

- <a href="#VUID-VkVideoSessionParametersCreateInfoKHR-sType-unique"
  id="VUID-VkVideoSessionParametersCreateInfoKHR-sType-unique"></a>
  VUID-VkVideoSessionParametersCreateInfoKHR-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

- <a href="#VUID-VkVideoSessionParametersCreateInfoKHR-flags-zerobitmask"
  id="VUID-VkVideoSessionParametersCreateInfoKHR-flags-zerobitmask"></a>
  VUID-VkVideoSessionParametersCreateInfoKHR-flags-zerobitmask  
  `flags` **must** be `0`

- <a
  href="#VUID-VkVideoSessionParametersCreateInfoKHR-videoSessionParametersTemplate-parameter"
  id="VUID-VkVideoSessionParametersCreateInfoKHR-videoSessionParametersTemplate-parameter"></a>
  VUID-VkVideoSessionParametersCreateInfoKHR-videoSessionParametersTemplate-parameter  
  If `videoSessionParametersTemplate` is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `videoSessionParametersTemplate` **must** be a valid
  [VkVideoSessionParametersKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionParametersKHR.html) handle

- <a
  href="#VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-parameter"
  id="VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-parameter"></a>
  VUID-VkVideoSessionParametersCreateInfoKHR-videoSession-parameter  
  `videoSession` **must** be a valid
  [VkVideoSessionKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionKHR.html) handle

- <a
  href="#VUID-VkVideoSessionParametersCreateInfoKHR-videoSessionParametersTemplate-parent"
  id="VUID-VkVideoSessionParametersCreateInfoKHR-videoSessionParametersTemplate-parent"></a>
  VUID-VkVideoSessionParametersCreateInfoKHR-videoSessionParametersTemplate-parent  
  If `videoSessionParametersTemplate` is a valid handle, it **must**
  have been created, allocated, or retrieved from `videoSession`

- <a href="#VUID-VkVideoSessionParametersCreateInfoKHR-commonparent"
  id="VUID-VkVideoSessionParametersCreateInfoKHR-commonparent"></a>
  VUID-VkVideoSessionParametersCreateInfoKHR-commonparent  
  Both of `videoSession`, and `videoSessionParametersTemplate` that are
  valid handles of non-ignored parameters **must** have been created,
  allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_queue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_queue.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkVideoSessionKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionKHR.html),
[VkVideoSessionParametersCreateFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionParametersCreateFlagsKHR.html),
[VkVideoSessionParametersKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionParametersKHR.html),
[vkCreateVideoSessionParametersKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateVideoSessionParametersKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoSessionParametersCreateInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
