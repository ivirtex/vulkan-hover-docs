# VkVideoEncodeH265PictureInfoKHR(3) Manual Page

## Name

VkVideoEncodeH265PictureInfoKHR - Structure specifies H.265 encode frame
parameters



## <a href="#_c_specification" class="anchor"></a>C Specification

The
[VkVideoEncodeH265PictureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265PictureInfoKHR.html)
structure is defined as:

``` c
// Provided by VK_KHR_video_encode_h265
typedef struct VkVideoEncodeH265PictureInfoKHR {
    VkStructureType                                    sType;
    const void*                                        pNext;
    uint32_t                                           naluSliceSegmentEntryCount;
    const VkVideoEncodeH265NaluSliceSegmentInfoKHR*    pNaluSliceSegmentEntries;
    const StdVideoEncodeH265PictureInfo*               pStdPictureInfo;
} VkVideoEncodeH265PictureInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `naluSliceSegmentEntryCount` is the number of elements in
  `pNaluSliceSegmentEntries`.

- `pNaluSliceSegmentEntries` is a pointer to an array of
  `naluSliceSegmentEntryCount`
  [VkVideoEncodeH265NaluSliceSegmentInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265NaluSliceSegmentInfoKHR.html)
  structures specifying the parameters of the individual H.265 slice
  segments to encode for the input picture.

- `pStdPictureInfo` is a pointer to a `StdVideoEncodeH265PictureInfo`
  structure specifying <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265-picture-info"
  target="_blank" rel="noopener">H.265 picture information</a>.

## <a href="#_description" class="anchor"></a>Description

This structure is specified in the `pNext` chain of the
[VkVideoEncodeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeInfoKHR.html) structure passed to
[vkCmdEncodeVideoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdEncodeVideoKHR.html) to specify the
codec-specific picture information for an <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265"
target="_blank" rel="noopener">H.265 encode operation</a>.

Encode Input Picture Information  
When this structure is specified in the `pNext` chain of the
[VkVideoEncodeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeInfoKHR.html) structure passed to
[vkCmdEncodeVideoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdEncodeVideoKHR.html), the information related
to the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-input-picture-info"
target="_blank" rel="noopener">encode input picture</a> is defined as
follows:

- The image subregion used is determined according to the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265-picture-data-access"
  target="_blank" rel="noopener">H.265 Encode Picture Data Access</a>
  section.

- The encode input picture is associated with the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265-picture-info"
  target="_blank" rel="noopener">H.265 picture information</a> provided
  in `pStdPictureInfo`.

<!-- -->

Std Picture Information  
The members of the `StdVideoEncodeH265PictureInfo` structure pointed to
by `pStdPictureInfo` are interpreted as follows:

- `flags.reserved` and `reserved1` are used only for padding purposes
  and are otherwise ignored;

- `flags.is_reference` as defined in section 3.132 of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h265"
  target="_blank" rel="noopener">ITU-T H.265 Specification</a>;

- `flags.IrapPicFlag` as defined in section 3.73 of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h265"
  target="_blank" rel="noopener">ITU-T H.265 Specification</a>;

- `flags.used_for_long_term_reference` is used to indicate whether the
  picture is marked as “used for long-term reference” as defined in
  section 8.3.2 of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h265"
  target="_blank" rel="noopener">ITU-T H.265 Specification</a>;

- `flags.discardable_flag` and `cross_layer_bla_flag` as defined in
  section F.7.4.7.1 of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h265"
  target="_blank" rel="noopener">ITU-T H.265 Specification</a>;

- `pic_type` as defined in section 7.4.3.5 of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h265"
  target="_blank" rel="noopener">ITU-T H.265 Specification</a>;

- `sps_video_parameter_set_id`, `pps_seq_parameter_set_id`, and
  `pps_pic_parameter_set_id` are used to identify the active parameter
  sets, as described below;

- `PicOrderCntVal` as defined in section 8.3.1 of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h265"
  target="_blank" rel="noopener">ITU-T H.265 Specification</a>;

- `TemporalId` as defined in section 7.4.2.2 of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h265"
  target="_blank" rel="noopener">ITU-T H.265 Specification</a>;

- if `pRefLists` is not `NULL`, then it is a pointer to a
  `StdVideoEncodeH265ReferenceListsInfo` structure that is interpreted
  as follows:

  - `flags.reserved` is used only for padding purposes and is otherwise
    ignored;

  - `ref_pic_list_modification_flag_l0` and
    `ref_pic_list_modification_flag_l1` as defined in section 7.4.7.2 of
    the <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h265"
    target="_blank" rel="noopener">ITU-T H.265 Specification</a>;

  - `num_ref_idx_l0_active_minus1` and `num_ref_idx_l1_active_minus1` as
    defined in section 7.4.7.1 of the <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h265"
    target="_blank" rel="noopener">ITU-T H.265 Specification</a>;

  - `RefPicList0` and `RefPicList1` as defined in section 8.3.4 of the
    <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h265"
    target="_blank" rel="noopener">ITU-T H.265 Specification</a> where
    each element of these arrays either identifies an <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-active-reference-picture-info"
    target="_blank" rel="noopener">active reference picture</a> using
    its <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#dpb-slot"
    target="_blank" rel="noopener">DPB slot</a> index or contains the
    value `STD_VIDEO_H265_NO_REFERENCE_PICTURE` to indicate “no
    reference picture”;

  - `list_entry_l0` and `list_entry_l1` as defined in section 7.4.7.2 of
    the <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h265"
    target="_blank" rel="noopener">ITU-T H.265 Specification</a>;

- if `flags.short_term_ref_pic_set_sps_flag` is set, then the
  `StdVideoH265ShortTermRefPicSet` structure pointed to by
  `pShortTermRefPicSet` is interpreted as defined for the elements of
  the `pShortTermRefPicSet` array specified in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265-sps"
  target="_blank" rel="noopener">H.265 sequence parameter sets</a>.

- if `flags.long_term_ref_pics_present_flag` is set in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265-active-sps"
  target="_blank" rel="noopener">active SPS</a>, then the
  `StdVideoEncodeH265LongTermRefPics` structure pointed to by
  `pLongTermRefPics` is interpreted as follows:

  - `used_by_curr_pic_lt_flag` is a bitmask where bit index i
    corresponds to `used_by_curr_pic_lt_flag[i]` as defined in section
    7.4.7.1 of the <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h265"
    target="_blank" rel="noopener">ITU-T H.265 Specification</a>;

  - all other members of `StdVideoEncodeH265LongTermRefPics` are
    interpreted as defined in section 7.4.7.1 of the <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h265"
    target="_blank" rel="noopener">ITU-T H.265 Specification</a>;

- all other members are interpreted as defined in section 7.4.7.1 of the
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h265"
  target="_blank" rel="noopener">ITU-T H.265 Specification</a>.

Reference picture setup is controlled by the value of
`StdVideoEncodeH265PictureInfo`::`flags.is_reference`. If it is set and
a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-reconstructed-picture-info"
target="_blank" rel="noopener">reconstructed picture</a> is specified,
then the latter is used as the target of picture reconstruction to <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#dpb-slot-states"
target="_blank" rel="noopener">activate</a> the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#dpb-slot"
target="_blank" rel="noopener">DPB slot</a> specified in
`pEncodeInfo->pSetupReferenceSlot->slotIndex`. If
`StdVideoEncodeH265PictureInfo`::`flags.is_reference` is not set, but a
<a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-reconstructed-picture-info"
target="_blank" rel="noopener">reconstructed picture</a> is specified,
then the corresponding picture reference associated with the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#dpb-slot"
target="_blank" rel="noopener">DPB slot</a> is invalidated, as described
in the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#dpb-slot-states"
target="_blank" rel="noopener">DPB Slot States</a> section.

Active Parameter Sets  
The members of the `StdVideoEncodeH265PictureInfo` structure pointed to
by `pStdPictureInfo` are used to select the active parameter sets to use
from the bound video session parameters object, as follows:

- <span id="encode-h265-active-vps"></span> The *active VPS* is the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265-vps"
  target="_blank" rel="noopener">VPS</a> identified by the key specified
  in `StdVideoEncodeH265PictureInfo`::`sps_video_parameter_set_id`.

- <span id="encode-h265-active-sps"></span> The *active SPS* is the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265-sps"
  target="_blank" rel="noopener">SPS</a> identified by the key specified
  by the pair constructed from
  `StdVideoEncodeH265PictureInfo`::`sps_video_parameter_set_id` and
  `StdVideoEncodeH265PictureInfo`::`pps_seq_parameter_set_id`.

- <span id="encode-h265-active-pps"></span> The *active PPS* is the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265-pps"
  target="_blank" rel="noopener">PPS</a> identified by the key specified
  by the triplet constructed from
  `StdVideoEncodeH265PictureInfo`::`sps_video_parameter_set_id`,
  `StdVideoEncodeH265PictureInfo`::`pps_seq_parameter_set_id`, and
  `StdVideoEncodeH265PictureInfo`::`pps_pic_parameter_set_id`.

H.265 encoding uses *explicit weighted sample prediction* for a slice
segment, as defined in section 8.5.3.3.4 of the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h265"
target="_blank" rel="noopener">ITU-T H.265 Specification</a>, if any of
the following conditions are true for the active <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265-pps"
target="_blank" rel="noopener">PPS</a> and the `pStdSliceSegmentHeader`
member of the corresponding element of `pNaluSliceSegmentEntries`:

- `pStdSliceSegmentHeader->slice_type` is `STD_VIDEO_H265_SLICE_TYPE_P`
  and `weighted_pred_flag` is enabled in the active PPS.

- `pStdSliceSegmentHeader->slice_type` is `STD_VIDEO_H265_SLICE_TYPE_B`
  and `weighted_bipred_flag` is enabled in the active PPS.

The number of H.265 tiles, as defined in section 3.174 of the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h265"
target="_blank" rel="noopener">ITU-T H.265 Specification</a>, is derived
from the `num_tile_columns_minus1` and `num_tile_rows_minus1` members of
the active <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265-pps"
target="_blank" rel="noopener">PPS</a> as follows:

  
(`num_tile_columns_minus1` + 1) × (`num_tile_rows_minus1` + 1)

Valid Usage

- <a
  href="#VUID-VkVideoEncodeH265PictureInfoKHR-naluSliceSegmentEntryCount-08306"
  id="VUID-VkVideoEncodeH265PictureInfoKHR-naluSliceSegmentEntryCount-08306"></a>
  VUID-VkVideoEncodeH265PictureInfoKHR-naluSliceSegmentEntryCount-08306  
  `naluSliceSegmentEntryCount` **must** be between `1` and
  [VkVideoEncodeH265CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265CapabilitiesKHR.html)::`maxSliceSegmentCount`,
  inclusive, as returned by
  [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)
  for the used video profile

- <a href="#VUID-VkVideoEncodeH265PictureInfoKHR-flags-08323"
  id="VUID-VkVideoEncodeH265PictureInfoKHR-flags-08323"></a>
  VUID-VkVideoEncodeH265PictureInfoKHR-flags-08323  
  If
  [VkVideoEncodeH265CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265CapabilitiesKHR.html)::`flags`,
  as returned by
  [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)
  for the used video profile, does not include
  `VK_VIDEO_ENCODE_H265_CAPABILITY_MULTIPLE_TILES_PER_SLICE_SEGMENT_BIT_KHR`,
  then `naluSliceSegmentEntryCount` **must** be greater than or equal to
  the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265-tile-count"
  target="_blank" rel="noopener">number of H.265 tiles in the picture</a>

- <a href="#VUID-VkVideoEncodeH265PictureInfoKHR-flags-08324"
  id="VUID-VkVideoEncodeH265PictureInfoKHR-flags-08324"></a>
  VUID-VkVideoEncodeH265PictureInfoKHR-flags-08324  
  If
  [VkVideoEncodeH265CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265CapabilitiesKHR.html)::`flags`,
  as returned by
  [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)
  for the used video profile, does not include
  `VK_VIDEO_ENCODE_H265_CAPABILITY_MULTIPLE_SLICE_SEGMENTS_PER_TILE_BIT_KHR`,
  then `naluSliceSegmentEntryCount` **must** be less than or equal to
  the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265-tile-count"
  target="_blank" rel="noopener">number of H.265 tiles in the picture</a>

- <a href="#VUID-VkVideoEncodeH265PictureInfoKHR-flags-08316"
  id="VUID-VkVideoEncodeH265PictureInfoKHR-flags-08316"></a>
  VUID-VkVideoEncodeH265PictureInfoKHR-flags-08316  
  If
  [VkVideoEncodeH265CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265CapabilitiesKHR.html)::`flags`,
  as returned by
  [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)
  for the used video profile, does not include
  `VK_VIDEO_ENCODE_H265_CAPABILITY_PREDICTION_WEIGHT_TABLE_GENERATED_BIT_KHR`
  and the slice segment corresponding to any element of
  `pNaluSliceSegmentEntries` uses <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265-weighted-pred"
  target="_blank" rel="noopener">explicit weighted sample prediction</a>,
  then
  [VkVideoEncodeH265NaluSliceSegmentInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265NaluSliceSegmentInfoKHR.html)::`pStdSliceSegmentHeader->pWeightTable`
  **must** not be `NULL` for that element of `pNaluSliceSegmentEntries`

- <a href="#VUID-VkVideoEncodeH265PictureInfoKHR-flags-08317"
  id="VUID-VkVideoEncodeH265PictureInfoKHR-flags-08317"></a>
  VUID-VkVideoEncodeH265PictureInfoKHR-flags-08317  
  If
  [VkVideoEncodeH265CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265CapabilitiesKHR.html)::`flags`,
  as returned by
  [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)
  for the used video profile, does not include
  `VK_VIDEO_ENCODE_H265_CAPABILITY_DIFFERENT_SLICE_SEGMENT_TYPE_BIT_KHR`,
  then
  [VkVideoEncodeH265NaluSliceSegmentInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265NaluSliceSegmentInfoKHR.html)::`pStdSliceSegmentHeader->slice_type`
  **must** be identical for all elements of `pNaluSliceSegmentEntries`

Valid Usage (Implicit)

- <a href="#VUID-VkVideoEncodeH265PictureInfoKHR-sType-sType"
  id="VUID-VkVideoEncodeH265PictureInfoKHR-sType-sType"></a>
  VUID-VkVideoEncodeH265PictureInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_VIDEO_ENCODE_H265_PICTURE_INFO_KHR`

- <a
  href="#VUID-VkVideoEncodeH265PictureInfoKHR-pNaluSliceSegmentEntries-parameter"
  id="VUID-VkVideoEncodeH265PictureInfoKHR-pNaluSliceSegmentEntries-parameter"></a>
  VUID-VkVideoEncodeH265PictureInfoKHR-pNaluSliceSegmentEntries-parameter  
  `pNaluSliceSegmentEntries` **must** be a valid pointer to an array of
  `naluSliceSegmentEntryCount` valid
  [VkVideoEncodeH265NaluSliceSegmentInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265NaluSliceSegmentInfoKHR.html)
  structures

- <a
  href="#VUID-VkVideoEncodeH265PictureInfoKHR-pStdPictureInfo-parameter"
  id="VUID-VkVideoEncodeH265PictureInfoKHR-pStdPictureInfo-parameter"></a>
  VUID-VkVideoEncodeH265PictureInfoKHR-pStdPictureInfo-parameter  
  `pStdPictureInfo` **must** be a valid pointer to a valid
  `StdVideoEncodeH265PictureInfo` value

- <a
  href="#VUID-VkVideoEncodeH265PictureInfoKHR-naluSliceSegmentEntryCount-arraylength"
  id="VUID-VkVideoEncodeH265PictureInfoKHR-naluSliceSegmentEntryCount-arraylength"></a>
  VUID-VkVideoEncodeH265PictureInfoKHR-naluSliceSegmentEntryCount-arraylength  
  `naluSliceSegmentEntryCount` **must** be greater than `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_encode_h265](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_encode_h265.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkVideoEncodeH265NaluSliceSegmentInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265NaluSliceSegmentInfoKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoEncodeH265PictureInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
