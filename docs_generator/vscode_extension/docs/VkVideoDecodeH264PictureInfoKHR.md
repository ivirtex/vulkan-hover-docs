# VkVideoDecodeH264PictureInfoKHR(3) Manual Page

## Name

VkVideoDecodeH264PictureInfoKHR - Structure specifies H.264 decode picture parameters when decoding a picture



## [](#_c_specification)C Specification

The `VkVideoDecodeH264PictureInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_video_decode_h264
typedef struct VkVideoDecodeH264PictureInfoKHR {
    VkStructureType                         sType;
    const void*                             pNext;
    const StdVideoDecodeH264PictureInfo*    pStdPictureInfo;
    uint32_t                                sliceCount;
    const uint32_t*                         pSliceOffsets;
} VkVideoDecodeH264PictureInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `pStdPictureInfo` is a pointer to a `StdVideoDecodeH264PictureInfo` structure specifying [H.264 picture information](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-h264-picture-info).
- `sliceCount` is the number of elements in `pSliceOffsets`.
- `pSliceOffsets` is a pointer to an array of `sliceCount` offsets specifying the start offset of the slices of the picture within the video bitstream buffer range specified in [VkVideoDecodeInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeInfoKHR.html).

## [](#_description)Description

This structure is specified in the `pNext` chain of the [VkVideoDecodeInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeInfoKHR.html) structure passed to [vkCmdDecodeVideoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDecodeVideoKHR.html) to specify the codec-specific picture information for an [H.264 decode operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-h264).

Decode Output Picture Information

When this structure is specified in the `pNext` chain of the [VkVideoDecodeInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeInfoKHR.html) structure passed to [vkCmdDecodeVideoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDecodeVideoKHR.html), the information related to the [decode output picture](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-output-picture-info) is defined as follows:

- If `pStdPictureInfo->flags.field_pic_flag` is not set, then the picture represents a frame.
- If `pStdPictureInfo->flags.field_pic_flag` is set, then the picture represents a field. Specifically:
  
  - If `pStdPictureInfo->flags.bottom_field_flag` is not set, then the picture represents the top field of the frame.
  - If `pStdPictureInfo->flags.bottom_field_flag` is set, then the picture represents the bottom field of the frame.
- The image subregion used is determined according to the [H.264 Decode Picture Data Access](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-h264-picture-data-access) section.
- The decode output picture is associated with the [H.264 picture information](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-h264-picture-info) provided in `pStdPictureInfo`.

Std Picture Information

The members of the `StdVideoDecodeH264PictureInfo` structure pointed to by `pStdPictureInfo` are interpreted as follows:

- `reserved1` and `reserved2` are used only for padding purposes and are otherwise ignored;
- `flags.is_intra` as defined in section 3.73 of the [ITU-T H.264 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#itu-t-h264);
- `flags.is_reference` as defined in section 3.136 of the [ITU-T H.264 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#itu-t-h264);
- `flags.complementary_field_pair` as defined in section 3.35 of the [ITU-T H.264 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#itu-t-h264);
- `seq_parameter_set_id` and `pic_parameter_set_id` are used to identify the active parameter sets, as described below;
- all other members are interpreted as defined in section 7.4.3 of the [ITU-T H.264 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#itu-t-h264).

Reference picture setup is controlled by the value of `StdVideoDecodeH264PictureInfo`::`flags.is_reference`. If it is set and a [reconstructed picture](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-reconstructed-picture-info) is specified, then the latter is used as the target of picture reconstruction to [activate](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#dpb-slot-states) the [DPB slot](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#dpb-slot) specified in `pDecodeInfo->pSetupReferenceSlot->slotIndex`. If `StdVideoDecodeH264PictureInfo`::`flags.is_reference` is not set, but a [reconstructed picture](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-reconstructed-picture-info) is specified, then the corresponding picture reference associated with the [DPB slot](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#dpb-slot) is invalidated, as described in the [DPB Slot States](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#dpb-slot-states) section.

Active Parameter Sets

The members of the `StdVideoDecodeH264PictureInfo` structure pointed to by `pStdPictureInfo` are used to select the active parameter sets to use from the specified [inline parameter sets](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-h264-inline-parameter-sets) or from the bound video session parameters object, as follows:

- []()The *active SPS* is the [SPS](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-h264-sps) identified by the key specified in `StdVideoDecodeH264PictureInfo`::`seq_parameter_set_id`.
- []()The *active PPS* is the [PPS](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-h264-pps) identified by the key specified by the pair constructed from `StdVideoDecodeH264PictureInfo`::`seq_parameter_set_id` and `StdVideoDecodeH264PictureInfo`::`pic_parameter_set_id`.

If the bound video session was created with `VK_VIDEO_SESSION_CREATE_INLINE_SESSION_PARAMETERS_BIT_KHR` and any of the active parameter sets are specified [inline](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-h264-inline-parameter-sets), then the corresponding inline parameter set is used. Otherwise, the parameter set in question is sourced from the bound video session parameters object.

Valid Usage (Implicit)

- [](#VUID-VkVideoDecodeH264PictureInfoKHR-sType-sType)VUID-VkVideoDecodeH264PictureInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_DECODE_H264_PICTURE_INFO_KHR`
- [](#VUID-VkVideoDecodeH264PictureInfoKHR-pStdPictureInfo-parameter)VUID-VkVideoDecodeH264PictureInfoKHR-pStdPictureInfo-parameter  
  `pStdPictureInfo` **must** be a valid pointer to a valid `StdVideoDecodeH264PictureInfo` value
- [](#VUID-VkVideoDecodeH264PictureInfoKHR-pSliceOffsets-parameter)VUID-VkVideoDecodeH264PictureInfoKHR-pSliceOffsets-parameter  
  `pSliceOffsets` **must** be a valid pointer to an array of `sliceCount` `uint32_t` values
- [](#VUID-VkVideoDecodeH264PictureInfoKHR-sliceCount-arraylength)VUID-VkVideoDecodeH264PictureInfoKHR-sliceCount-arraylength  
  `sliceCount` **must** be greater than `0`

## [](#_see_also)See Also

[VK\_KHR\_video\_decode\_h264](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_decode_h264.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoDecodeH264PictureInfoKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0