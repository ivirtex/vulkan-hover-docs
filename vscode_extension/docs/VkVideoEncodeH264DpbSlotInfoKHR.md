# VkVideoEncodeH264DpbSlotInfoKHR(3) Manual Page

## Name

VkVideoEncodeH264DpbSlotInfoKHR - Structure specifies H.264 encode DPB picture information



## [](#_c_specification)C Specification

The [VkVideoEncodeH264DpbSlotInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264DpbSlotInfoKHR.html) structure is defined as:

```c++
// Provided by VK_KHR_video_encode_h264
typedef struct VkVideoEncodeH264DpbSlotInfoKHR {
    VkStructureType                           sType;
    const void*                               pNext;
    const StdVideoEncodeH264ReferenceInfo*    pStdReferenceInfo;
} VkVideoEncodeH264DpbSlotInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `pStdReferenceInfo` is a pointer to a `StdVideoEncodeH264ReferenceInfo` structure specifying [H.264 reference information](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h264-reference-info).

## [](#_description)Description

This structure is specified in the `pNext` chain of [VkVideoEncodeInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeInfoKHR.html)::`pSetupReferenceSlot`, if not `NULL`, and the `pNext` chain of the elements of [VkVideoEncodeInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeInfoKHR.html)::`pReferenceSlots` to specify the codec-specific reference picture information for an [H.264 encode operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h264).

Active Reference Picture Information

When this structure is specified in the `pNext` chain of the elements of [VkVideoEncodeInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeInfoKHR.html)::`pReferenceSlots`, one element is added to the list of [active reference pictures](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-active-reference-picture-info) used by the video encode operation for each element of [VkVideoEncodeInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeInfoKHR.html)::`pReferenceSlots` as follows:

- The image subregion used is determined according to the [H.264 Encode Picture Data Access](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h264-picture-data-access) section.
- The reference picture is associated with the [DPB slot](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#dpb-slot) index specified in the `slotIndex` member of the corresponding element of [VkVideoEncodeInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeInfoKHR.html)::`pReferenceSlots`.
- The reference picture is associated with the [H.264 reference information](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h264-reference-info) provided in `pStdReferenceInfo`.

Reconstructed Picture Information

When this structure is specified in the `pNext` chain of [VkVideoEncodeInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeInfoKHR.html)::`pSetupReferenceSlot`, the information related to the [reconstructed picture](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-reconstructed-picture-info) is defined as follows:

- The image subregion used is determined according to the [H.264 Encode Picture Data Access](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h264-picture-data-access) section.
- If [reference picture setup](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h264-ref-pic-setup) is requested, then the reconstructed picture is used to [activate](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#dpb-slot-states) the [DPB slot](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#dpb-slot) with the index specified in [VkVideoEncodeInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeInfoKHR.html)::`pSetupReferenceSlot->slotIndex`.
- The reconstructed picture is associated with the [H.264 reference information](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h264-reference-info) provided in `pStdReferenceInfo`.

Std Reference Information

The members of the `StdVideoEncodeH264ReferenceInfo` structure pointed to by `pStdReferenceInfo` are interpreted as follows:

- `flags.reserved` is used only for padding purposes and is otherwise ignored;
- `flags.used_for_long_term_reference` is used to indicate whether the picture is marked as “used for long-term reference” as defined in section 8.2.5.1 of the [ITU-T H.264 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#itu-t-h264);
- `primary_pic_type` as defined in section 7.4.2 of the [ITU-T H.264 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#itu-t-h264);
- `long_term_pic_num` and `long_term_frame_idx` as defined in section 7.4.3 of the [ITU-T H.264 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#itu-t-h264);
- `temporal_id` as defined in section G.7.4.1.1 of the [ITU-T H.264 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#itu-t-h264);
- all other members are interpreted as defined in section 8.2 of the [ITU-T H.264 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#itu-t-h264).

Valid Usage (Implicit)

- [](#VUID-VkVideoEncodeH264DpbSlotInfoKHR-sType-sType)VUID-VkVideoEncodeH264DpbSlotInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_ENCODE_H264_DPB_SLOT_INFO_KHR`
- [](#VUID-VkVideoEncodeH264DpbSlotInfoKHR-pStdReferenceInfo-parameter)VUID-VkVideoEncodeH264DpbSlotInfoKHR-pStdReferenceInfo-parameter  
  `pStdReferenceInfo` **must** be a valid pointer to a valid `StdVideoEncodeH264ReferenceInfo` value

## [](#_see_also)See Also

[VK\_KHR\_video\_encode\_h264](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_h264.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoEncodeH264DpbSlotInfoKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0