# VkVideoEncodeH265NaluSliceSegmentInfoKHR(3) Manual Page

## Name

VkVideoEncodeH265NaluSliceSegmentInfoKHR - Structure specifies H.265
encode slice segment NALU parameters



## <a href="#_c_specification" class="anchor"></a>C Specification

The
[VkVideoEncodeH265NaluSliceSegmentInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265NaluSliceSegmentInfoKHR.html)
structure is defined as:

``` c
// Provided by VK_KHR_video_encode_h265
typedef struct VkVideoEncodeH265NaluSliceSegmentInfoKHR {
    VkStructureType                                sType;
    const void*                                    pNext;
    int32_t                                        constantQp;
    const StdVideoEncodeH265SliceSegmentHeader*    pStdSliceSegmentHeader;
} VkVideoEncodeH265NaluSliceSegmentInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `constantQp` is the QP to use for the slice segment if the current <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-rate-control-modes"
  target="_blank" rel="noopener">rate control mode</a> configured for
  the video session is
  `VK_VIDEO_ENCODE_RATE_CONTROL_MODE_DISABLED_BIT_KHR`.

- `pStdSliceSegmentHeader` is a pointer to a
  `StdVideoEncodeH265SliceSegmentHeader` structure specifying <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265-slice-segment-header-params"
  target="_blank" rel="noopener">H.265 slice segment header parameters</a>
  for the slice segment.

## <a href="#_description" class="anchor"></a>Description

Std Slice Segment Header Parameters  
The members of the `StdVideoEncodeH265SliceSegmentHeader` structure
pointed to by `pStdSliceSegmentHeader` are interpreted as follows:

- `flags.reserved` and `reserved1` are used only for padding purposes
  and are otherwise ignored;

- if `pWeightTable` is not `NULL`, then it is a pointer to a
  `StdVideoEncodeH265WeightTable` that is interpreted as follows:

  - `flags.luma_weight_l0_flag`, `flags.chroma_weight_l0_flag`,
    `flags.luma_weight_l1_flag`, and `flags.chroma_weight_l1_flag` are
    bitmasks where bit index i corresponds to `luma_weight_l0_flag[i]`,
    `chroma_weight_l0_flag[i]`, `luma_weight_l1_flag[i]`, and
    `chroma_weight_l1_flag[i]`, respectively, as defined in section
    7.4.7.3 of the <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h265"
    target="_blank" rel="noopener">ITU-T H.265 Specification</a>;

  - all other members of `StdVideoEncodeH265WeightTable` are interpreted
    as defined in section 7.4.7.3 of the <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h265"
    target="_blank" rel="noopener">ITU-T H.265 Specification</a>;

- all other members are interpreted as defined in section 7.4.7.1 of the
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h265"
  target="_blank" rel="noopener">ITU-T H.265 Specification</a>.

Valid Usage (Implicit)

- <a href="#VUID-VkVideoEncodeH265NaluSliceSegmentInfoKHR-sType-sType"
  id="VUID-VkVideoEncodeH265NaluSliceSegmentInfoKHR-sType-sType"></a>
  VUID-VkVideoEncodeH265NaluSliceSegmentInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_VIDEO_ENCODE_H265_NALU_SLICE_SEGMENT_INFO_KHR`

- <a href="#VUID-VkVideoEncodeH265NaluSliceSegmentInfoKHR-pNext-pNext"
  id="VUID-VkVideoEncodeH265NaluSliceSegmentInfoKHR-pNext-pNext"></a>
  VUID-VkVideoEncodeH265NaluSliceSegmentInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`

- <a
  href="#VUID-VkVideoEncodeH265NaluSliceSegmentInfoKHR-pStdSliceSegmentHeader-parameter"
  id="VUID-VkVideoEncodeH265NaluSliceSegmentInfoKHR-pStdSliceSegmentHeader-parameter"></a>
  VUID-VkVideoEncodeH265NaluSliceSegmentInfoKHR-pStdSliceSegmentHeader-parameter  
  `pStdSliceSegmentHeader` **must** be a valid pointer to a valid
  `StdVideoEncodeH265SliceSegmentHeader` value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_encode_h265](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_encode_h265.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkVideoEncodeH265PictureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265PictureInfoKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoEncodeH265NaluSliceSegmentInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
