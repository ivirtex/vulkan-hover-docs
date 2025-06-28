# VkVideoEncodeH265NaluSliceSegmentInfoKHR(3) Manual Page

## Name

VkVideoEncodeH265NaluSliceSegmentInfoKHR - Structure specifies H.265 encode slice segment NALU parameters



## [](#_c_specification)C Specification

The [VkVideoEncodeH265NaluSliceSegmentInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265NaluSliceSegmentInfoKHR.html) structure is defined as:

```c++
// Provided by VK_KHR_video_encode_h265
typedef struct VkVideoEncodeH265NaluSliceSegmentInfoKHR {
    VkStructureType                                sType;
    const void*                                    pNext;
    int32_t                                        constantQp;
    const StdVideoEncodeH265SliceSegmentHeader*    pStdSliceSegmentHeader;
} VkVideoEncodeH265NaluSliceSegmentInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `constantQp` is the QP to use for the slice segment if the current [rate control mode](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-rate-control-modes) configured for the video session is `VK_VIDEO_ENCODE_RATE_CONTROL_MODE_DISABLED_BIT_KHR`.
- `pStdSliceSegmentHeader` is a pointer to a `StdVideoEncodeH265SliceSegmentHeader` structure specifying [H.265 slice segment header parameters](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h265-slice-segment-header-params) for the slice segment.

## [](#_description)Description

Std Slice Segment Header Parameters

The members of the `StdVideoEncodeH265SliceSegmentHeader` structure pointed to by `pStdSliceSegmentHeader` are interpreted as follows:

- `flags.reserved` and `reserved1` are used only for padding purposes and are otherwise ignored;
- if `pWeightTable` is not `NULL`, then it is a pointer to a `StdVideoEncodeH265WeightTable` that is interpreted as follows:
  
  - `flags.luma_weight_l0_flag`, `flags.chroma_weight_l0_flag`, `flags.luma_weight_l1_flag`, and `flags.chroma_weight_l1_flag` are bitmasks where bit index i corresponds to `luma_weight_l0_flag[i]`, `chroma_weight_l0_flag[i]`, `luma_weight_l1_flag[i]`, and `chroma_weight_l1_flag[i]`, respectively, as defined in section 7.4.7.3 of the [ITU-T H.265 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#itu-t-h265);
  - all other members of `StdVideoEncodeH265WeightTable` are interpreted as defined in section 7.4.7.3 of the [ITU-T H.265 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#itu-t-h265);
- all other members are interpreted as defined in section 7.4.7.1 of the [ITU-T H.265 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#itu-t-h265).

Valid Usage (Implicit)

- [](#VUID-VkVideoEncodeH265NaluSliceSegmentInfoKHR-sType-sType)VUID-VkVideoEncodeH265NaluSliceSegmentInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_ENCODE_H265_NALU_SLICE_SEGMENT_INFO_KHR`
- [](#VUID-VkVideoEncodeH265NaluSliceSegmentInfoKHR-pNext-pNext)VUID-VkVideoEncodeH265NaluSliceSegmentInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkVideoEncodeH265NaluSliceSegmentInfoKHR-pStdSliceSegmentHeader-parameter)VUID-VkVideoEncodeH265NaluSliceSegmentInfoKHR-pStdSliceSegmentHeader-parameter  
  `pStdSliceSegmentHeader` **must** be a valid pointer to a valid `StdVideoEncodeH265SliceSegmentHeader` value

## [](#_see_also)See Also

[VK\_KHR\_video\_encode\_h265](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_h265.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkVideoEncodeH265PictureInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265PictureInfoKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoEncodeH265NaluSliceSegmentInfoKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0