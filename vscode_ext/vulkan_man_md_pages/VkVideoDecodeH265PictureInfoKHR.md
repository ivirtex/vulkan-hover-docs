# VkVideoDecodeH265PictureInfoKHR(3) Manual Page

## Name

VkVideoDecodeH265PictureInfoKHR - Structure specifies H.265 picture
information when decoding a frame



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkVideoDecodeH265PictureInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_video_decode_h265
typedef struct VkVideoDecodeH265PictureInfoKHR {
    VkStructureType                         sType;
    const void*                             pNext;
    const StdVideoDecodeH265PictureInfo*    pStdPictureInfo;
    uint32_t                                sliceSegmentCount;
    const uint32_t*                         pSliceSegmentOffsets;
} VkVideoDecodeH265PictureInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `pStdPictureInfo` is a pointer to a `StdVideoDecodeH265PictureInfo`
  structure specifying <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h265-picture-info"
  target="_blank" rel="noopener">H.265 picture information</a>.

- `sliceSegmentCount` is the number of elements in
  `pSliceSegmentOffsets`.

- `pSliceSegmentOffsets` is a pointer to an array of `sliceSegmentCount`
  offsets specifying the start offset of the slice segments of the
  picture within the video bitstream buffer range specified in
  [VkVideoDecodeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeInfoKHR.html).

## <a href="#_description" class="anchor"></a>Description

This structure is specified in the `pNext` chain of the
[VkVideoDecodeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeInfoKHR.html) structure passed to
[vkCmdDecodeVideoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDecodeVideoKHR.html) to specify the
codec-specific picture information for an <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h265"
target="_blank" rel="noopener">H.265 decode operation</a>.

Decode Output Picture Information  
When this structure is specified in the `pNext` chain of the
[VkVideoDecodeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeInfoKHR.html) structure passed to
[vkCmdDecodeVideoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDecodeVideoKHR.html), the information related
to the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-output-picture-info"
target="_blank" rel="noopener">decode output picture</a> is defined as
follows:

- The image subregion used is determined according to the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h265-picture-data-access"
  target="_blank" rel="noopener">H.265 Decode Picture Data Access</a>
  section.

- The decode output picture is associated with the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h265-picture-info"
  target="_blank" rel="noopener">H.265 picture information</a> provided
  in `pStdPictureInfo`.

<!-- -->

Std Picture Information  
The members of the `StdVideoDecodeH265PictureInfo` structure pointed to
by `pStdPictureInfo` are interpreted as follows:

- `reserved` is used only for padding purposes and is otherwise ignored;

- `flags.IrapPicFlag` as defined in section 3.73 of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h265"
  target="_blank" rel="noopener">ITU-T H.265 Specification</a>;

- `flags.IdrPicFlag` as defined in section 3.67 of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h265"
  target="_blank" rel="noopener">ITU-T H.265 Specification</a>;

- `flags.IsReference` as defined in section 3.132 of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h265"
  target="_blank" rel="noopener">ITU-T H.265 Specification</a>;

- `sps_video_parameter_set_id`, `pps_seq_parameter_set_id`, and
  `pps_pic_parameter_set_id` are used to identify the active parameter
  sets, as described below;

- `PicOrderCntVal` as defined in section 8.3.1 of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h265"
  target="_blank" rel="noopener">ITU-T H.265 Specification</a>;

- `NumBitsForSTRefPicSetInSlice` is the number of bits used in
  `st_ref_pic_set` when `short_term_ref_pic_set_sps_flag` is `0`, or `0`
  otherwise, as defined in sections 7.4.7 and 7.4.8 of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h265"
  target="_blank" rel="noopener">ITU-T H.265 Specification</a>;

- `NumDeltaPocsOfRefRpsIdx` is the value of `NumDeltaPocs[RefRpsIdx]`
  when `short_term_ref_pic_set_sps_flag` is `1`, or `0` otherwise, as
  defined in sections 7.4.7 and 7.4.8 of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h265"
  target="_blank" rel="noopener">ITU-T H.265 Specification</a>;

- `RefPicSetStCurrBefore`, `RefPicSetStCurrAfter`, and `RefPicSetLtCurr`
  are interpreted as defined in section 8.3.2 of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h265"
  target="_blank" rel="noopener">ITU-T H.265 Specification</a> where
  each element of these arrays either identifies an <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-active-reference-picture-info"
  target="_blank" rel="noopener">active reference picture</a> using its
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#dpb-slot"
  target="_blank" rel="noopener">DPB slot</a> index or contains the
  value `STD_VIDEO_H265_NO_REFERENCE_PICTURE` to indicate “no reference
  picture”;

- all other members are interpreted as defined in section 8.3.2 of the
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h265"
  target="_blank" rel="noopener">ITU-T H.265 Specification</a>.

Reference picture setup is controlled by the value of
`StdVideoDecodeH265PictureInfo`::`flags.IsReference`. If it is set and a
<a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-reconstructed-picture-info"
target="_blank" rel="noopener">reconstructed picture</a> is specified,
then the latter is used as the target of picture reconstruction to <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#dpb-slot-states"
target="_blank" rel="noopener">activate</a> the corresponding <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#dpb-slot"
target="_blank" rel="noopener">DPB slot</a>. If
`StdVideoDecodeH265PictureInfo`::`flags.IsReference` is not set, but a
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
The members of the `StdVideoDecodeH265PictureInfo` structure pointed to
by `pStdPictureInfo` are used to select the active parameter sets to use
from the bound video session parameters object, as follows:

- <span id="decode-h265-active-vps"></span> The *active VPS* is the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h265-vps"
  target="_blank" rel="noopener">VPS</a> identified by the key specified
  in `StdVideoDecodeH265PictureInfo`::`sps_video_parameter_set_id`.

- <span id="decode-h265-active-sps"></span> The *active SPS* is the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h265-sps"
  target="_blank" rel="noopener">SPS</a> identified by the key specified
  by the pair constructed from
  `StdVideoDecodeH265PictureInfo`::`sps_video_parameter_set_id` and
  `StdVideoDecodeH265PictureInfo`::`pps_seq_parameter_set_id`.

- <span id="decode-h265-active-pps"></span> The *active PPS* is the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h265-pps"
  target="_blank" rel="noopener">PPS</a> identified by the key specified
  by the triplet constructed from
  `StdVideoDecodeH265PictureInfo`::`sps_video_parameter_set_id`,
  `StdVideoDecodeH265PictureInfo`::`pps_seq_parameter_set_id`, and
  `StdVideoDecodeH265PictureInfo`::`pps_pic_parameter_set_id`.

Valid Usage (Implicit)

- <a href="#VUID-VkVideoDecodeH265PictureInfoKHR-sType-sType"
  id="VUID-VkVideoDecodeH265PictureInfoKHR-sType-sType"></a>
  VUID-VkVideoDecodeH265PictureInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_VIDEO_DECODE_H265_PICTURE_INFO_KHR`

- <a
  href="#VUID-VkVideoDecodeH265PictureInfoKHR-pStdPictureInfo-parameter"
  id="VUID-VkVideoDecodeH265PictureInfoKHR-pStdPictureInfo-parameter"></a>
  VUID-VkVideoDecodeH265PictureInfoKHR-pStdPictureInfo-parameter  
  `pStdPictureInfo` **must** be a valid pointer to a valid
  `StdVideoDecodeH265PictureInfo` value

- <a
  href="#VUID-VkVideoDecodeH265PictureInfoKHR-pSliceSegmentOffsets-parameter"
  id="VUID-VkVideoDecodeH265PictureInfoKHR-pSliceSegmentOffsets-parameter"></a>
  VUID-VkVideoDecodeH265PictureInfoKHR-pSliceSegmentOffsets-parameter  
  `pSliceSegmentOffsets` **must** be a valid pointer to an array of
  `sliceSegmentCount` `uint32_t` values

- <a
  href="#VUID-VkVideoDecodeH265PictureInfoKHR-sliceSegmentCount-arraylength"
  id="VUID-VkVideoDecodeH265PictureInfoKHR-sliceSegmentCount-arraylength"></a>
  VUID-VkVideoDecodeH265PictureInfoKHR-sliceSegmentCount-arraylength  
  `sliceSegmentCount` **must** be greater than `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_decode_h265](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_decode_h265.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoDecodeH265PictureInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
