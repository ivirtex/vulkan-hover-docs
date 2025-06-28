# VkVideoDecodeH264DpbSlotInfoKHR(3) Manual Page

## Name

VkVideoDecodeH264DpbSlotInfoKHR - Structure specifies H.264 decode DPB picture information



## [](#_c_specification)C Specification

The `VkVideoDecodeH264DpbSlotInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_video_decode_h264
typedef struct VkVideoDecodeH264DpbSlotInfoKHR {
    VkStructureType                           sType;
    const void*                               pNext;
    const StdVideoDecodeH264ReferenceInfo*    pStdReferenceInfo;
} VkVideoDecodeH264DpbSlotInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `pStdReferenceInfo` is a pointer to a `StdVideoDecodeH264ReferenceInfo` structure specifying [H.264 reference information](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-h264-reference-info).

## [](#_description)Description

This structure is specified in the `pNext` chain of [VkVideoDecodeInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeInfoKHR.html)::`pSetupReferenceSlot`, if not `NULL`, and the `pNext` chain of the elements of [VkVideoDecodeInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeInfoKHR.html)::`pReferenceSlots` to specify the codec-specific reference picture information for an [H.264 decode operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-h264).

Active Reference Picture Information

When this structure is specified in the `pNext` chain of the elements of [VkVideoDecodeInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeInfoKHR.html)::`pReferenceSlots`, one or two elements are added to the list of [active reference pictures](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-active-reference-picture-info) used by the video decode operation for each element of [VkVideoDecodeInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeInfoKHR.html)::`pReferenceSlots` as follows:

- If neither `pStdReferenceInfo->flags.top_field_flag` nor `pStdReferenceInfo->flags.bottom_field_flag` is set, then the picture is added as a frame reference to the list of active reference pictures.
- If `pStdReferenceInfo->flags.top_field_flag` is set, then the picture is added as a top field reference to the list of active reference pictures.
- If `pStdReferenceInfo->flags.bottom_field_flag` is set, then the picture is added as a bottom field reference to the list of active reference pictures.
- For each added reference picture, the corresponding image subregion used is determined according to the [H.264 Decode Picture Data Access](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-h264-picture-data-access) section.
- Each added reference picture is associated with the [DPB slot](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#dpb-slot) index specified in the `slotIndex` member of the corresponding element of [VkVideoDecodeInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeInfoKHR.html)::`pReferenceSlots`.
- Each added reference picture is associated with the [H.264 reference information](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-h264-reference-info) provided in `pStdReferenceInfo`.

Note

When both the top and bottom field of an interlaced frame currently associated with a DPB slot is intended to be used as an active reference picture and both fields are stored in the same image subregion (which is the case when using `VK_VIDEO_DECODE_H264_PICTURE_LAYOUT_INTERLACED_INTERLEAVED_LINES_BIT_KHR` which stores the two fields at even and odd scanlines of the same image subregion), both references have to be provided through a single [VkVideoReferenceSlotInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoReferenceSlotInfoKHR.html) structure that has both `flags.top_field_flag` and `flags.bottom_field_flag` set in the `StdVideoDecodeH264ReferenceInfo` structure pointed to by the `pStdReferenceInfo` member of the [VkVideoDecodeH264DpbSlotInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH264DpbSlotInfoKHR.html) structure included in the corresponding [VkVideoReferenceSlotInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoReferenceSlotInfoKHR.html) structure’s `pNext` chain. However, this approach can only be used when both fields are stored in the same image subregion. If that is not the case (e.g. when using `VK_VIDEO_DECODE_H264_PICTURE_LAYOUT_INTERLACED_SEPARATE_PLANES_BIT_KHR` which requires separate `codedOffset` values for the two fields and also allows storing the two fields of a frame in separate image layers or entirely separate images), then a separate [VkVideoReferenceSlotInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoReferenceSlotInfoKHR.html) structure needs to be provided for referencing the two fields, each only setting one of `flags.top_field_flag` or `flags.bottom_field_flag`, and providing the appropriate video picture resource information in [VkVideoReferenceSlotInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoReferenceSlotInfoKHR.html)::`pPictureResource`.

Reconstructed Picture Information

When this structure is specified in the `pNext` chain of [VkVideoDecodeInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeInfoKHR.html)::`pSetupReferenceSlot`, the information related to the [reconstructed picture](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-reconstructed-picture-info) is defined as follows:

- If neither `pStdReferenceInfo->flags.top_field_flag` nor `pStdReferenceInfo->flags.bottom_field_flag` is set, then the picture represents a frame.
- If `pStdReferenceInfo->flags.top_field_flag` is set, then the picture represents a field, specifically, the top field of the frame.
- If `pStdReferenceInfo->flags.bottom_field_flag` is set, then the picture represents a field, specifically, the bottom field of the frame.
- The image subregion used is determined according to the [H.264 Decode Picture Data Access](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-h264-picture-data-access) section.
- If [reference picture setup](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-ref-pic-setup) is requested, then the reconstructed picture is used to [activate](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#dpb-slot-states) the [DPB slot](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#dpb-slot) with the index specified in [VkVideoDecodeInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeInfoKHR.html)::`pSetupReferenceSlot->slotIndex`.
- The reconstructed picture is associated with the [H.264 reference information](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-h264-reference-info) provided in `pStdReferenceInfo`.

Std Reference Information

The members of the `StdVideoDecodeH264ReferenceInfo` structure pointed to by `pStdReferenceInfo` are interpreted as follows:

- `flags.top_field_flag` is used to indicate whether the reference is used as top field reference;
- `flags.bottom_field_flag` is used to indicate whether the reference is used as bottom field reference;
- `flags.used_for_long_term_reference` is used to indicate whether the picture is marked as “used for long-term reference” as defined in section 8.2.5.1 of the [ITU-T H.264 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#itu-t-h264);
- `flags.is_non_existing` is used to indicate whether the picture is marked as “non-existing” as defined in section 8.2.5.2 of the [ITU-T H.264 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#itu-t-h264);
- all other members are interpreted as defined in section 8.2 of the [ITU-T H.264 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#itu-t-h264).

Valid Usage (Implicit)

- [](#VUID-VkVideoDecodeH264DpbSlotInfoKHR-sType-sType)VUID-VkVideoDecodeH264DpbSlotInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_DECODE_H264_DPB_SLOT_INFO_KHR`
- [](#VUID-VkVideoDecodeH264DpbSlotInfoKHR-pStdReferenceInfo-parameter)VUID-VkVideoDecodeH264DpbSlotInfoKHR-pStdReferenceInfo-parameter  
  `pStdReferenceInfo` **must** be a valid pointer to a valid `StdVideoDecodeH264ReferenceInfo` value

## [](#_see_also)See Also

[VK\_KHR\_video\_decode\_h264](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_decode_h264.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoDecodeH264DpbSlotInfoKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0