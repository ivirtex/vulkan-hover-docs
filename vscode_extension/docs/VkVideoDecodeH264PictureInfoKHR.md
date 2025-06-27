# VkVideoDecodeH264PictureInfoKHR(3) Manual Page

## Name

VkVideoDecodeH264PictureInfoKHR - Structure specifies H.264 decode
picture parameters when decoding a picture



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkVideoDecodeH264PictureInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_video_decode_h264
typedef struct VkVideoDecodeH264PictureInfoKHR {
    VkStructureType                         sType;
    const void*                             pNext;
    const StdVideoDecodeH264PictureInfo*    pStdPictureInfo;
    uint32_t                                sliceCount;
    const uint32_t*                         pSliceOffsets;
} VkVideoDecodeH264PictureInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `pStdPictureInfo` is a pointer to a `StdVideoDecodeH264PictureInfo`
  structure specifying <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h264-picture-info"
  target="_blank" rel="noopener">H.264 picture information</a>.

- `sliceCount` is the number of elements in `pSliceOffsets`.

- `pSliceOffsets` is a pointer to an array of `sliceCount` offsets
  specifying the start offset of the slices of the picture within the
  video bitstream buffer range specified in
  [VkVideoDecodeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeInfoKHR.html).

## <a href="#_description" class="anchor"></a>Description

This structure is specified in the `pNext` chain of the
[VkVideoDecodeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeInfoKHR.html) structure passed to
[vkCmdDecodeVideoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDecodeVideoKHR.html) to specify the
codec-specific picture information for an <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h264"
target="_blank" rel="noopener">H.264 decode operation</a>.

Decode Output Picture Information  
When this structure is specified in the `pNext` chain of the
[VkVideoDecodeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeInfoKHR.html) structure passed to
[vkCmdDecodeVideoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDecodeVideoKHR.html), the information related
to the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-output-picture-info"
target="_blank" rel="noopener">decode output picture</a> is defined as
follows:

- If `pStdPictureInfo->flags.field_pic_flag` is not set, then the
  picture represents a frame.

- If `pStdPictureInfo->flags.field_pic_flag` is set, then the picture
  represents a field. Specifically:

  - If `pStdPictureInfo->flags.bottom_field_flag` is not set, then the
    picture represents the top field of the frame.

  - If `pStdPictureInfo->flags.bottom_field_flag` is set, then the
    picture represents the bottom field of the frame.

- The image subregion used is determined according to the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h264-picture-data-access"
  target="_blank" rel="noopener">H.264 Decode Picture Data Access</a>
  section.

- The decode output picture is associated with the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h264-picture-info"
  target="_blank" rel="noopener">H.264 picture information</a> provided
  in `pStdPictureInfo`.

<!-- -->

Std Picture Information  
The members of the `StdVideoDecodeH264PictureInfo` structure pointed to
by `pStdPictureInfo` are interpreted as follows:

- `reserved1` and `reserved2` are used only for padding purposes and are
  otherwise ignored;

- `flags.is_intra` as defined in section 3.73 of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h264"
  target="_blank" rel="noopener">ITU-T H.264 Specification</a>;

- `flags.is_reference` as defined in section 3.136 of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h264"
  target="_blank" rel="noopener">ITU-T H.264 Specification</a>;

- `flags.complementary_field_pair` as defined in section 3.35 of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h264"
  target="_blank" rel="noopener">ITU-T H.264 Specification</a>;

- `seq_parameter_set_id` and `pic_parameter_set_id` are used to identify
  the active parameter sets, as described below;

- all other members are interpreted as defined in section 7.4.3 of the
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h264"
  target="_blank" rel="noopener">ITU-T H.264 Specification</a>.

Reference picture setup is controlled by the value of
`StdVideoDecodeH264PictureInfo`::`flags.is_reference`. If it is set and
a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-reconstructed-picture-info"
target="_blank" rel="noopener">reconstructed picture</a> is specified,
then the latter is used as the target of picture reconstruction to <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#dpb-slot-states"
target="_blank" rel="noopener">activate</a> the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#dpb-slot"
target="_blank" rel="noopener">DPB slot</a> specified in
`pDecodeInfo->pSetupReferenceSlot->slotIndex`. If
`StdVideoDecodeH264PictureInfo`::`flags.is_reference` is not set, but a
<a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-reconstructed-picture-info"
target="_blank" rel="noopener">reconstructed picture</a> is specified,
then the corresponding picture reference associated with the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#dpb-slot"
target="_blank" rel="noopener">DPB slot</a> is invalidated, as described
in the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#dpb-slot-states"
target="_blank" rel="noopener">DPB Slot States</a> section.

Active Parameter Sets  
The members of the `StdVideoDecodeH264PictureInfo` structure pointed to
by `pStdPictureInfo` are used to select the active parameter sets to use
from the bound video session parameters object, as follows:

- <span id="decode-h264-active-sps"></span> The *active SPS* is the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h264-sps"
  target="_blank" rel="noopener">SPS</a> identified by the key specified
  in `StdVideoDecodeH264PictureInfo`::`seq_parameter_set_id`.

- <span id="decode-h264-active-pps"></span> The *active PPS* is the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h264-pps"
  target="_blank" rel="noopener">PPS</a> identified by the key specified
  by the pair constructed from
  `StdVideoDecodeH264PictureInfo`::`seq_parameter_set_id` and
  `StdVideoDecodeH264PictureInfo`::`pic_parameter_set_id`.

Valid Usage (Implicit)

- <a href="#VUID-VkVideoDecodeH264PictureInfoKHR-sType-sType"
  id="VUID-VkVideoDecodeH264PictureInfoKHR-sType-sType"></a>
  VUID-VkVideoDecodeH264PictureInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_VIDEO_DECODE_H264_PICTURE_INFO_KHR`

- <a
  href="#VUID-VkVideoDecodeH264PictureInfoKHR-pStdPictureInfo-parameter"
  id="VUID-VkVideoDecodeH264PictureInfoKHR-pStdPictureInfo-parameter"></a>
  VUID-VkVideoDecodeH264PictureInfoKHR-pStdPictureInfo-parameter  
  `pStdPictureInfo` **must** be a valid pointer to a valid
  `StdVideoDecodeH264PictureInfo` value

- <a href="#VUID-VkVideoDecodeH264PictureInfoKHR-pSliceOffsets-parameter"
  id="VUID-VkVideoDecodeH264PictureInfoKHR-pSliceOffsets-parameter"></a>
  VUID-VkVideoDecodeH264PictureInfoKHR-pSliceOffsets-parameter  
  `pSliceOffsets` **must** be a valid pointer to an array of
  `sliceCount` `uint32_t` values

- <a href="#VUID-VkVideoDecodeH264PictureInfoKHR-sliceCount-arraylength"
  id="VUID-VkVideoDecodeH264PictureInfoKHR-sliceCount-arraylength"></a>
  VUID-VkVideoDecodeH264PictureInfoKHR-sliceCount-arraylength  
  `sliceCount` **must** be greater than `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_decode_h264](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_decode_h264.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoDecodeH264PictureInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
