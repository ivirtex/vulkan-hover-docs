# VkVideoEncodeH265PictureInfoKHR(3) Manual Page

## Name

VkVideoEncodeH265PictureInfoKHR - Structure specifies H.265 encode frame parameters



## [](#_c_specification)C Specification

The [VkVideoEncodeH265PictureInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265PictureInfoKHR.html) structure is defined as:

```c++
// Provided by VK_KHR_video_encode_h265
typedef struct VkVideoEncodeH265PictureInfoKHR {
    VkStructureType                                    sType;
    const void*                                        pNext;
    uint32_t                                           naluSliceSegmentEntryCount;
    const VkVideoEncodeH265NaluSliceSegmentInfoKHR*    pNaluSliceSegmentEntries;
    const StdVideoEncodeH265PictureInfo*               pStdPictureInfo;
} VkVideoEncodeH265PictureInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `naluSliceSegmentEntryCount` is the number of elements in `pNaluSliceSegmentEntries`.
- `pNaluSliceSegmentEntries` is a pointer to an array of `naluSliceSegmentEntryCount` [VkVideoEncodeH265NaluSliceSegmentInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265NaluSliceSegmentInfoKHR.html) structures specifying the parameters of the individual H.265 slice segments to encode for the input picture.
- `pStdPictureInfo` is a pointer to a `StdVideoEncodeH265PictureInfo` structure specifying [H.265 picture information](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h265-picture-info).

## [](#_description)Description

This structure is specified in the `pNext` chain of the [VkVideoEncodeInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeInfoKHR.html) structure passed to [vkCmdEncodeVideoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdEncodeVideoKHR.html) to specify the codec-specific picture information for an [H.265 encode operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h265).

Encode Input Picture Information

When this structure is specified in the `pNext` chain of the [VkVideoEncodeInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeInfoKHR.html) structure passed to [vkCmdEncodeVideoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdEncodeVideoKHR.html), the information related to the [encode input picture](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-input-picture-info) is defined as follows:

- The image subregion used is determined according to the [H.265 Encode Picture Data Access](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h265-picture-data-access) section.
- The encode input picture is associated with the [H.265 picture information](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h265-picture-info) provided in `pStdPictureInfo`.

Std Picture Information

The members of the `StdVideoEncodeH265PictureInfo` structure pointed to by `pStdPictureInfo` are interpreted as follows:

- `flags.reserved` and `reserved1` are used only for padding purposes and are otherwise ignored;
- `flags.is_reference` as defined in section 3.132 of the [ITU-T H.265 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#itu-t-h265);
- `flags.IrapPicFlag` as defined in section 3.73 of the [ITU-T H.265 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#itu-t-h265);
- `flags.used_for_long_term_reference` is used to indicate whether the picture is marked as “used for long-term reference” as defined in section 8.3.2 of the [ITU-T H.265 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#itu-t-h265);
- `flags.discardable_flag` and `cross_layer_bla_flag` as defined in section F.7.4.7.1 of the [ITU-T H.265 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#itu-t-h265);
- `pic_type` as defined in section 7.4.3.5 of the [ITU-T H.265 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#itu-t-h265);
- `sps_video_parameter_set_id`, `pps_seq_parameter_set_id`, and `pps_pic_parameter_set_id` are used to identify the active parameter sets, as described below;
- `PicOrderCntVal` as defined in section 8.3.1 of the [ITU-T H.265 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#itu-t-h265);
- `TemporalId` as defined in section 7.4.2.2 of the [ITU-T H.265 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#itu-t-h265);
- if `pRefLists` is not `NULL`, then it is a pointer to a `StdVideoEncodeH265ReferenceListsInfo` structure that is interpreted as follows:
  
  - `flags.reserved` is used only for padding purposes and is otherwise ignored;
  - `ref_pic_list_modification_flag_l0` and `ref_pic_list_modification_flag_l1` as defined in section 7.4.7.2 of the [ITU-T H.265 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#itu-t-h265);
  - `num_ref_idx_l0_active_minus1` and `num_ref_idx_l1_active_minus1` as defined in section 7.4.7.1 of the [ITU-T H.265 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#itu-t-h265);
  - `RefPicList0` and `RefPicList1` as defined in section 8.3.4 of the [ITU-T H.265 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#itu-t-h265) where each element of these arrays either identifies an [active reference picture](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-active-reference-picture-info) using its [DPB slot](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#dpb-slot) index or contains the value `STD_VIDEO_H265_NO_REFERENCE_PICTURE` to indicate “no reference picture”;
  - `list_entry_l0` and `list_entry_l1` as defined in section 7.4.7.2 of the [ITU-T H.265 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#itu-t-h265);
- if `flags.short_term_ref_pic_set_sps_flag` is set, then the `StdVideoH265ShortTermRefPicSet` structure pointed to by `pShortTermRefPicSet` is interpreted as defined for the elements of the `pShortTermRefPicSet` array specified in [H.265 sequence parameter sets](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h265-sps).
- if `flags.long_term_ref_pics_present_flag` is set in the [active SPS](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h265-active-sps), then the `StdVideoEncodeH265LongTermRefPics` structure pointed to by `pLongTermRefPics` is interpreted as follows:
  
  - `used_by_curr_pic_lt_flag` is a bitmask where bit index i corresponds to `used_by_curr_pic_lt_flag[i]` as defined in section 7.4.7.1 of the [ITU-T H.265 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#itu-t-h265);
  - all other members of `StdVideoEncodeH265LongTermRefPics` are interpreted as defined in section 7.4.7.1 of the [ITU-T H.265 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#itu-t-h265);
- all other members are interpreted as defined in section 7.4.7.1 of the [ITU-T H.265 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#itu-t-h265).

Reference picture setup is controlled by the value of `StdVideoEncodeH265PictureInfo`::`flags.is_reference`. If it is set and a [reconstructed picture](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-reconstructed-picture-info) is specified, then the latter is used as the target of picture reconstruction to [activate](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#dpb-slot-states) the [DPB slot](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#dpb-slot) specified in `pEncodeInfo->pSetupReferenceSlot->slotIndex`. If `StdVideoEncodeH265PictureInfo`::`flags.is_reference` is not set, but a [reconstructed picture](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-reconstructed-picture-info) is specified, then the corresponding picture reference associated with the [DPB slot](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#dpb-slot) is invalidated, as described in the [DPB Slot States](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#dpb-slot-states) section.

Active Parameter Sets

The members of the `StdVideoEncodeH265PictureInfo` structure pointed to by `pStdPictureInfo` are used to select the active parameter sets to use from the bound video session parameters object, as follows:

- []()The *active VPS* is the [VPS](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h265-vps) identified by the key specified in `StdVideoEncodeH265PictureInfo`::`sps_video_parameter_set_id`.
- []()The *active SPS* is the [SPS](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h265-sps) identified by the key specified by the pair constructed from `StdVideoEncodeH265PictureInfo`::`sps_video_parameter_set_id` and `StdVideoEncodeH265PictureInfo`::`pps_seq_parameter_set_id`.
- []()The *active PPS* is the [PPS](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h265-pps) identified by the key specified by the triplet constructed from `StdVideoEncodeH265PictureInfo`::`sps_video_parameter_set_id`, `StdVideoEncodeH265PictureInfo`::`pps_seq_parameter_set_id`, and `StdVideoEncodeH265PictureInfo`::`pps_pic_parameter_set_id`.

H.265 encoding uses *explicit weighted sample prediction* for a slice segment, as defined in section 8.5.3.3.4 of the [ITU-T H.265 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#itu-t-h265), if any of the following conditions are true for the active [PPS](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h265-pps) and the `pStdSliceSegmentHeader` member of the corresponding element of `pNaluSliceSegmentEntries`:

- `pStdSliceSegmentHeader->slice_type` is `STD_VIDEO_H265_SLICE_TYPE_P` and `weighted_pred_flag` is enabled in the active PPS.
- `pStdSliceSegmentHeader->slice_type` is `STD_VIDEO_H265_SLICE_TYPE_B` and `weighted_bipred_flag` is enabled in the active PPS.

The number of H.265 tiles, as defined in section 3.174 of the [ITU-T H.265 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#itu-t-h265), is derived from the `num_tile_columns_minus1` and `num_tile_rows_minus1` members of the active [PPS](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h265-pps) as follows:

(`num_tile_columns_minus1` + 1) × (`num_tile_rows_minus1` + 1)

Valid Usage

- [](#VUID-VkVideoEncodeH265PictureInfoKHR-naluSliceSegmentEntryCount-08306)VUID-VkVideoEncodeH265PictureInfoKHR-naluSliceSegmentEntryCount-08306  
  `naluSliceSegmentEntryCount` **must** be between `1` and [VkVideoEncodeH265CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265CapabilitiesKHR.html)::`maxSliceSegmentCount`, inclusive, as returned by [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html) for the used video profile
- [](#VUID-VkVideoEncodeH265PictureInfoKHR-flags-08323)VUID-VkVideoEncodeH265PictureInfoKHR-flags-08323  
  If [VkVideoEncodeH265CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265CapabilitiesKHR.html)::`flags`, as returned by [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html) for the used video profile, does not include `VK_VIDEO_ENCODE_H265_CAPABILITY_MULTIPLE_TILES_PER_SLICE_SEGMENT_BIT_KHR`, then `naluSliceSegmentEntryCount` **must** be greater than or equal to the [number of H.265 tiles in the picture](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h265-tile-count)
- [](#VUID-VkVideoEncodeH265PictureInfoKHR-flags-08324)VUID-VkVideoEncodeH265PictureInfoKHR-flags-08324  
  If [VkVideoEncodeH265CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265CapabilitiesKHR.html)::`flags`, as returned by [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html) for the used video profile, does not include `VK_VIDEO_ENCODE_H265_CAPABILITY_MULTIPLE_SLICE_SEGMENTS_PER_TILE_BIT_KHR`, then `naluSliceSegmentEntryCount` **must** be less than or equal to the [number of H.265 tiles in the picture](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h265-tile-count)
- [](#VUID-VkVideoEncodeH265PictureInfoKHR-flags-08316)VUID-VkVideoEncodeH265PictureInfoKHR-flags-08316  
  If [VkVideoEncodeH265CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265CapabilitiesKHR.html)::`flags`, as returned by [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html) for the used video profile, does not include `VK_VIDEO_ENCODE_H265_CAPABILITY_PREDICTION_WEIGHT_TABLE_GENERATED_BIT_KHR` and the slice segment corresponding to any element of `pNaluSliceSegmentEntries` uses [explicit weighted sample prediction](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h265-weighted-pred), then [VkVideoEncodeH265NaluSliceSegmentInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265NaluSliceSegmentInfoKHR.html)::`pStdSliceSegmentHeader->pWeightTable` **must** not be `NULL` for that element of `pNaluSliceSegmentEntries`
- [](#VUID-VkVideoEncodeH265PictureInfoKHR-flags-08317)VUID-VkVideoEncodeH265PictureInfoKHR-flags-08317  
  If [VkVideoEncodeH265CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265CapabilitiesKHR.html)::`flags`, as returned by [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html) for the used video profile, does not include `VK_VIDEO_ENCODE_H265_CAPABILITY_DIFFERENT_SLICE_SEGMENT_TYPE_BIT_KHR`, then [VkVideoEncodeH265NaluSliceSegmentInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265NaluSliceSegmentInfoKHR.html)::`pStdSliceSegmentHeader->slice_type` **must** be identical for all elements of `pNaluSliceSegmentEntries`

Valid Usage (Implicit)

- [](#VUID-VkVideoEncodeH265PictureInfoKHR-sType-sType)VUID-VkVideoEncodeH265PictureInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_ENCODE_H265_PICTURE_INFO_KHR`
- [](#VUID-VkVideoEncodeH265PictureInfoKHR-pNaluSliceSegmentEntries-parameter)VUID-VkVideoEncodeH265PictureInfoKHR-pNaluSliceSegmentEntries-parameter  
  `pNaluSliceSegmentEntries` **must** be a valid pointer to an array of `naluSliceSegmentEntryCount` valid [VkVideoEncodeH265NaluSliceSegmentInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265NaluSliceSegmentInfoKHR.html) structures
- [](#VUID-VkVideoEncodeH265PictureInfoKHR-pStdPictureInfo-parameter)VUID-VkVideoEncodeH265PictureInfoKHR-pStdPictureInfo-parameter  
  `pStdPictureInfo` **must** be a valid pointer to a valid `StdVideoEncodeH265PictureInfo` value
- [](#VUID-VkVideoEncodeH265PictureInfoKHR-naluSliceSegmentEntryCount-arraylength)VUID-VkVideoEncodeH265PictureInfoKHR-naluSliceSegmentEntryCount-arraylength  
  `naluSliceSegmentEntryCount` **must** be greater than `0`

## [](#_see_also)See Also

[VK\_KHR\_video\_encode\_h265](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_h265.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkVideoEncodeH265NaluSliceSegmentInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265NaluSliceSegmentInfoKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoEncodeH265PictureInfoKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0