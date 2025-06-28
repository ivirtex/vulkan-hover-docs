# VkVideoEncodeAV1DpbSlotInfoKHR(3) Manual Page

## Name

VkVideoEncodeAV1DpbSlotInfoKHR - Structure specifies AV1 encode DPB picture information



## [](#_c_specification)C Specification

The [VkVideoEncodeAV1DpbSlotInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1DpbSlotInfoKHR.html) structure is defined as:

```c++
// Provided by VK_KHR_video_encode_av1
typedef struct VkVideoEncodeAV1DpbSlotInfoKHR {
    VkStructureType                          sType;
    const void*                              pNext;
    const StdVideoEncodeAV1ReferenceInfo*    pStdReferenceInfo;
} VkVideoEncodeAV1DpbSlotInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `pStdReferenceInfo` is a pointer to a `StdVideoEncodeAV1ReferenceInfo` structure specifying [AV1 reference information](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-reference-info).

## [](#_description)Description

This structure is specified in the `pNext` chain of [VkVideoEncodeInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeInfoKHR.html)::`pSetupReferenceSlot`, if not `NULL`, and the `pNext` chain of the elements of [VkVideoEncodeInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeInfoKHR.html)::`pReferenceSlots` to specify the codec-specific reference picture information for an [AV1 encode operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1).

Active Reference Picture Information

When this structure is specified in the `pNext` chain of the elements of [VkVideoEncodeInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeInfoKHR.html)::`pReferenceSlots`, one element is added to the list of [active reference pictures](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-active-reference-picture-info) used by the video encode operation for each element of [VkVideoEncodeInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeInfoKHR.html)::`pReferenceSlots` as follows:

- The image subregion used is determined according to the [AV1 Encode Picture Data Access](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-picture-data-access) section.
- The reference picture is associated with the [DPB slot](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#dpb-slot) index specified in the `slotIndex` member of the corresponding element of [VkVideoEncodeInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeInfoKHR.html)::`pReferenceSlots`.
- The reference picture is associated with the [AV1 reference information](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-reference-info) provided in `pStdReferenceInfo`.

Reconstructed Picture Information

When this structure is specified in the `pNext` chain of [VkVideoEncodeInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeInfoKHR.html)::`pSetupReferenceSlot`, the information related to the [reconstructed picture](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-reconstructed-picture-info) is defined as follows:

- The image subregion used is determined according to the [AV1 Encode Picture Data Access](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-picture-data-access) section.
- If [reference picture setup](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-ref-pic-setup) is requested, then the reconstructed picture is used to [activate](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#dpb-slot-states) the [DPB slot](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#dpb-slot) with the index specified in [VkVideoEncodeInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeInfoKHR.html)::`pSetupReferenceSlot->slotIndex`.
- The reconstructed picture is associated with the [AV1 reference information](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-reference-info) provided in `pStdReferenceInfo`.

Std Reference Information

The members of the `StdVideoEncodeAV1ReferenceInfo` structure pointed to by `pStdReferenceInfo` are interpreted as follows:

- `flags.reserved` and `reserved1` are used only for padding purposes and are otherwise ignored;
- `flags.disable_frame_end_update_cdf` is interpreted as defined in section 6.8.2 of the [AV1 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#aomedia-av1);
- `flags.segmentation_enabled` is interpreted as defined in section 6.8.13 of the [AV1 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#aomedia-av1);
- `RefFrameId` is interpreted as the element of the `RefFrameId` array defined in section 6.8.2 of the [AV1 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#aomedia-av1) corresponding to the reference frame;
- `frame_type` is interpreted as defined in section 6.8.2 of the [AV1 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#aomedia-av1);
- `OrderHint` is interpreted as defined in section 6.8.2 of the [AV1 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#aomedia-av1);
- `pExtensionHeader` is `NULL` or a pointer to a `StdVideoEncodeAV1ExtensionHeader` structure whose `temporal_id` and `spatial_id` members specify the temporal and spatial layer ID of the reference frame, respectively.

Valid Usage (Implicit)

- [](#VUID-VkVideoEncodeAV1DpbSlotInfoKHR-sType-sType)VUID-VkVideoEncodeAV1DpbSlotInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_ENCODE_AV1_DPB_SLOT_INFO_KHR`
- [](#VUID-VkVideoEncodeAV1DpbSlotInfoKHR-pStdReferenceInfo-parameter)VUID-VkVideoEncodeAV1DpbSlotInfoKHR-pStdReferenceInfo-parameter  
  `pStdReferenceInfo` **must** be a valid pointer to a valid `StdVideoEncodeAV1ReferenceInfo` value

## [](#_see_also)See Also

[VK\_KHR\_video\_encode\_av1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_av1.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoEncodeAV1DpbSlotInfoKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0