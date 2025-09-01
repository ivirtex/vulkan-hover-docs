# VkVideoDecodeH265PictureInfoKHR(3) Manual Page

## Name

VkVideoDecodeH265PictureInfoKHR - Structure specifies H.265 picture information when decoding a frame



## [](#_c_specification)C Specification

The `VkVideoDecodeH265PictureInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_video_decode_h265
typedef struct VkVideoDecodeH265PictureInfoKHR {
    VkStructureType                         sType;
    const void*                             pNext;
    const StdVideoDecodeH265PictureInfo*    pStdPictureInfo;
    uint32_t                                sliceSegmentCount;
    const uint32_t*                         pSliceSegmentOffsets;
} VkVideoDecodeH265PictureInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `pStdPictureInfo` is a pointer to a `StdVideoDecodeH265PictureInfo` structure specifying [H.265 picture information](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-h265-picture-info).
- `sliceSegmentCount` is the number of elements in `pSliceSegmentOffsets`.
- `pSliceSegmentOffsets` is a pointer to an array of `sliceSegmentCount` offsets specifying the start offset of the slice segments of the picture within the video bitstream buffer range specified in [VkVideoDecodeInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeInfoKHR.html).

## [](#_description)Description

This structure is specified in the `pNext` chain of the [VkVideoDecodeInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeInfoKHR.html) structure passed to [vkCmdDecodeVideoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDecodeVideoKHR.html) to specify the codec-specific picture information for an [H.265 decode operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-h265).

Decode Output Picture Information

When this structure is specified in the `pNext` chain of the [VkVideoDecodeInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeInfoKHR.html) structure passed to [vkCmdDecodeVideoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDecodeVideoKHR.html), the information related to the [decode output picture](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-output-picture-info) is defined as follows:

- The image subregion used is determined according to the [H.265 Decode Picture Data Access](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-h265-picture-data-access) section.
- The decode output picture is associated with the [H.265 picture information](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-h265-picture-info) provided in `pStdPictureInfo`.

Std Picture Information

The members of the `StdVideoDecodeH265PictureInfo` structure pointed to by `pStdPictureInfo` are interpreted as follows:

- `reserved` is used only for padding purposes and is otherwise ignored;
- `flags.IrapPicFlag` as defined in section 3.73 of the [ITU-T H.265 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#itu-t-h265);
- `flags.IdrPicFlag` as defined in section 3.67 of the [ITU-T H.265 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#itu-t-h265);
- `flags.IsReference` as defined in section 3.132 of the [ITU-T H.265 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#itu-t-h265);
- `sps_video_parameter_set_id`, `pps_seq_parameter_set_id`, and `pps_pic_parameter_set_id` are used to identify the active parameter sets, as described below;
- `PicOrderCntVal` as defined in section 8.3.1 of the [ITU-T H.265 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#itu-t-h265);
- `NumBitsForSTRefPicSetInSlice` is the number of bits used in `st_ref_pic_set` when `short_term_ref_pic_set_sps_flag` is `0`, or `0` otherwise, as defined in sections 7.4.7 and 7.4.8 of the [ITU-T H.265 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#itu-t-h265);
- `NumDeltaPocsOfRefRpsIdx` is the value of `NumDeltaPocs[RefRpsIdx]` when `short_term_ref_pic_set_sps_flag` is `1`, or `0` otherwise, as defined in sections 7.4.7 and 7.4.8 of the [ITU-T H.265 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#itu-t-h265);
- `RefPicSetStCurrBefore`, `RefPicSetStCurrAfter`, and `RefPicSetLtCurr` are interpreted as defined in section 8.3.2 of the [ITU-T H.265 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#itu-t-h265) where each element of these arrays either identifies an [active reference picture](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-active-reference-picture-info) using its [DPB slot](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#dpb-slot) index or contains the value `STD_VIDEO_H265_NO_REFERENCE_PICTURE` to indicate “no reference picture”;
- all other members are interpreted as defined in section 8.3.2 of the [ITU-T H.265 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#itu-t-h265).

Reference picture setup is controlled by the value of `StdVideoDecodeH265PictureInfo`::`flags.IsReference`. If it is set and a [reconstructed picture](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-reconstructed-picture-info) is specified, then the latter is used as the target of picture reconstruction to [activate](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#dpb-slot-states) the corresponding [DPB slot](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#dpb-slot). If `StdVideoDecodeH265PictureInfo`::`flags.IsReference` is not set, but a [reconstructed picture](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-reconstructed-picture-info) is specified, then the corresponding picture reference associated with the [DPB slot](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#dpb-slot) is invalidated, as described in the [DPB Slot States](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#dpb-slot-states) section.

Active Parameter Sets

The members of the `StdVideoDecodeH265PictureInfo` structure pointed to by `pStdPictureInfo` are used to select the active parameter sets to use from the specified [inline parameter sets](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-h265-inline-parameter-sets) or from the bound video session parameters object, as follows:

- []()The *active VPS* is the [VPS](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-h265-vps) identified by the key specified in `StdVideoDecodeH265PictureInfo`::`sps_video_parameter_set_id`.
- []()The *active SPS* is the [SPS](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-h265-sps) identified by the key specified by the pair constructed from `StdVideoDecodeH265PictureInfo`::`sps_video_parameter_set_id` and `StdVideoDecodeH265PictureInfo`::`pps_seq_parameter_set_id`.
- []()The *active PPS* is the [PPS](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-h265-pps) identified by the key specified by the triplet constructed from `StdVideoDecodeH265PictureInfo`::`sps_video_parameter_set_id`, `StdVideoDecodeH265PictureInfo`::`pps_seq_parameter_set_id`, and `StdVideoDecodeH265PictureInfo`::`pps_pic_parameter_set_id`.

If the bound video session was created with `VK_VIDEO_SESSION_CREATE_INLINE_SESSION_PARAMETERS_BIT_KHR` and any of the active parameter sets are specified [inline](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-h265-inline-parameter-sets), then the corresponding inline parameter set is used. Otherwise, the parameter set in question is sourced from the bound video session parameters object.

Valid Usage (Implicit)

- [](#VUID-VkVideoDecodeH265PictureInfoKHR-sType-sType)VUID-VkVideoDecodeH265PictureInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_DECODE_H265_PICTURE_INFO_KHR`
- [](#VUID-VkVideoDecodeH265PictureInfoKHR-pStdPictureInfo-parameter)VUID-VkVideoDecodeH265PictureInfoKHR-pStdPictureInfo-parameter  
  `pStdPictureInfo` **must** be a valid pointer to a valid `StdVideoDecodeH265PictureInfo` value
- [](#VUID-VkVideoDecodeH265PictureInfoKHR-pSliceSegmentOffsets-parameter)VUID-VkVideoDecodeH265PictureInfoKHR-pSliceSegmentOffsets-parameter  
  `pSliceSegmentOffsets` **must** be a valid pointer to an array of `sliceSegmentCount` `uint32_t` values
- [](#VUID-VkVideoDecodeH265PictureInfoKHR-sliceSegmentCount-arraylength)VUID-VkVideoDecodeH265PictureInfoKHR-sliceSegmentCount-arraylength  
  `sliceSegmentCount` **must** be greater than `0`

## [](#_see_also)See Also

[VK\_KHR\_video\_decode\_h265](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_decode_h265.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoDecodeH265PictureInfoKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0