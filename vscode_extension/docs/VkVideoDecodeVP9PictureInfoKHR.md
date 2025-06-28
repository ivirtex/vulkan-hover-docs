# VkVideoDecodeVP9PictureInfoKHR(3) Manual Page

## Name

VkVideoDecodeVP9PictureInfoKHR - Structure specifies VP9 picture information when decoding a frame



## [](#_c_specification)C Specification

The `VkVideoDecodeVP9PictureInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_video_decode_vp9
typedef struct VkVideoDecodeVP9PictureInfoKHR {
    VkStructureType                        sType;
    const void*                            pNext;
    const StdVideoDecodeVP9PictureInfo*    pStdPictureInfo;
    int32_t                                referenceNameSlotIndices[VK_MAX_VIDEO_VP9_REFERENCES_PER_FRAME_KHR];
    uint32_t                               uncompressedHeaderOffset;
    uint32_t                               compressedHeaderOffset;
    uint32_t                               tilesOffset;
} VkVideoDecodeVP9PictureInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `pStdPictureInfo` is a pointer to a `StdVideoDecodeVP9PictureInfo` structure specifying [VP9 picture information](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-vp9-picture-info).
- `referenceNameSlotIndices` is an array of three (`VK_MAX_VIDEO_VP9_REFERENCES_PER_FRAME_KHR`, which is equal to the Video Std definition `STD_VIDEO_VP9_REFS_PER_FRAME`) signed integer values specifying the index of the [DPB slot](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#dpb-slot) or a negative integer value for each [VP9 reference name](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-vp9-reference-names) used for inter coding. In particular, the DPB slot index for the VP9 reference name `frame` is specified in `referenceNameSlotIndices`\[`frame` - `STD_VIDEO_VP9_REFERENCE_NAME_LAST_FRAME`].
- `uncompressedHeaderOffset` is the byte offset of the uncompressed frame header, as defined in section 6.2 of the [VP9 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#google-vp9).
- `compressedHeaderOffset` is the byte offset of the compressed frame header, as defined in section 6.3 of the [VP9 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#google-vp9).
- `tilesOffset` is the byte offset of the frame tile data, as defined in section 6.4 of the [VP9 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#google-vp9).

## [](#_description)Description

This structure is specified in the `pNext` chain of the [VkVideoDecodeInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeInfoKHR.html) structure passed to [vkCmdDecodeVideoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDecodeVideoKHR.html) to specify the codec-specific picture information for an [VP9 decode operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-vp9).

Decode Output Picture Information

When this structure is specified in the `pNext` chain of the [VkVideoDecodeInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeInfoKHR.html) structure passed to [vkCmdDecodeVideoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDecodeVideoKHR.html), the information related to the [decode output picture](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-output-picture-info) is defined as follows:

- The image subregion used is determined according to the [VP9 Decode Picture Data Access](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-vp9-picture-data-access) section.
- The decode output picture is associated with the [VP9 picture information](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-vp9-picture-info) provided in `pStdPictureInfo`.

Std Picture Information

The members of the `StdVideoDecodeVP9PictureInfo` structure pointed to by `pStdPictureInfo` are interpreted as follows:

- `flags.reserved` and `reserved1` are used only for padding purposes and are otherwise ignored;
- `ref_frame_sign_bias_mask` is a bitmask where bit index i corresponds to `ref_frame_sign_bias[i]` as defined in section 7.2 of the [VP9 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#google-vp9);
- the `StdVideoVP9ColorConfig` structure pointed to by `pColorConfig` is interpreted as follows:
  
  - `flags.reserved` and `reserved1` are used only for padding purposes and are otherwise ignored;
  - all other members of `StdVideoVP9ColorConfig` are interpreted as defined in sections 6.2, 6.2.2, and 7.2.2 of the [VP9 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#google-vp9);
- the `StdVideoVP9LoopFilter` structure pointed to by `pLoopFilter` is interpreted as follows:
  
  - `flags.reserved` is used only for padding purposes and is otherwise ignored;
  - `update_ref_delta` is a bitmask where bit index i is interpreted as the value of `update_ref_delta` corresponding to element i of `loop_filter_ref_deltas` as defined in section 7.2.8 of the [VP9 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#google-vp9);
  - `update_mode_delta` is a bitmask where bit index i is interpreted as the value of `update_mode_delta` corresponding to element i of `loop_filter_mode_deltas` as defined in section 7.2.8 of the [VP9 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#google-vp9);
  - all other members of `StdVideoVP9LoopFilter` are interpreted as defined in section 7.2.8 of the [VP9 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#google-vp9);
    
    Note
    
    If the syntax elements corresponding to `loop_filter_ref_deltas` and `loop_filter_mode_deltas` are not present, the application should specify the previous values, as defined in section 7.2.8 of the [VP9 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#google-vp9).
- if `flags.segmentation_enabled` is set, then the `StdVideoVP9Segmentation` structure pointed to by `pSegmentation` is interpreted as follows:
  
  - `flags.reserved` is used only for padding purposes and is otherwise ignored;
  - the elements of `FeatureEnabled` are bitmasks where bit index j of element i corresponds to `FeatureEnabled[i][j]` as defined in section 6.2.11 of the [VP9 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#google-vp9);
  - `FeatureData` is interpreted as defined in section 6.2.11 of the [VP9 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#google-vp9);
  - all other members of `StdVideoVP9Segmentation` are interpreted as defined in section 7.2.10 of the [VP9 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#google-vp9);
- all other members are interpreted as defined in section 7.2 of the [VP9 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#google-vp9).

Reference picture setup is controlled by the value of `StdVideoDecodeVP9PictureInfo`::`refresh_frame_flags`. If it is not zero and a [reconstructed picture](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-reconstructed-picture-info) is specified, then the latter is used as the target of picture reconstruction to [activate](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#dpb-slot-states) the [DPB slot](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#dpb-slot) specified in `pDecodeInfo->pSetupReferenceSlot->slotIndex`. If `StdVideoDecodeVP9PictureInfo`::`refresh_frame_flags` is zero, but a [reconstructed picture](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-reconstructed-picture-info) is specified, then the corresponding picture reference associated with the [DPB slot](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#dpb-slot) is invalidated, as described in the [DPB Slot States](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#dpb-slot-states) section.

Valid Usage (Implicit)

- [](#VUID-VkVideoDecodeVP9PictureInfoKHR-sType-sType)VUID-VkVideoDecodeVP9PictureInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_DECODE_VP9_PICTURE_INFO_KHR`
- [](#VUID-VkVideoDecodeVP9PictureInfoKHR-pStdPictureInfo-parameter)VUID-VkVideoDecodeVP9PictureInfoKHR-pStdPictureInfo-parameter  
  `pStdPictureInfo` **must** be a valid pointer to a valid `StdVideoDecodeVP9PictureInfo` value

## [](#_see_also)See Also

[VK\_KHR\_video\_decode\_vp9](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_decode_vp9.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoDecodeVP9PictureInfoKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0