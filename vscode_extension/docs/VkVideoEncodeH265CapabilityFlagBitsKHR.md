# VkVideoEncodeH265CapabilityFlagBitsKHR(3) Manual Page

## Name

VkVideoEncodeH265CapabilityFlagBitsKHR - Video encode H.265 capability
flags



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **may** be set in
[VkVideoEncodeH265CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265CapabilitiesKHR.html)::`flags`,
indicating the H.265 encoding capabilities supported, are:

``` c
// Provided by VK_KHR_video_encode_h265
typedef enum VkVideoEncodeH265CapabilityFlagBitsKHR {
    VK_VIDEO_ENCODE_H265_CAPABILITY_HRD_COMPLIANCE_BIT_KHR = 0x00000001,
    VK_VIDEO_ENCODE_H265_CAPABILITY_PREDICTION_WEIGHT_TABLE_GENERATED_BIT_KHR = 0x00000002,
    VK_VIDEO_ENCODE_H265_CAPABILITY_ROW_UNALIGNED_SLICE_SEGMENT_BIT_KHR = 0x00000004,
    VK_VIDEO_ENCODE_H265_CAPABILITY_DIFFERENT_SLICE_SEGMENT_TYPE_BIT_KHR = 0x00000008,
    VK_VIDEO_ENCODE_H265_CAPABILITY_B_FRAME_IN_L0_LIST_BIT_KHR = 0x00000010,
    VK_VIDEO_ENCODE_H265_CAPABILITY_B_FRAME_IN_L1_LIST_BIT_KHR = 0x00000020,
    VK_VIDEO_ENCODE_H265_CAPABILITY_PER_PICTURE_TYPE_MIN_MAX_QP_BIT_KHR = 0x00000040,
    VK_VIDEO_ENCODE_H265_CAPABILITY_PER_SLICE_SEGMENT_CONSTANT_QP_BIT_KHR = 0x00000080,
    VK_VIDEO_ENCODE_H265_CAPABILITY_MULTIPLE_TILES_PER_SLICE_SEGMENT_BIT_KHR = 0x00000100,
    VK_VIDEO_ENCODE_H265_CAPABILITY_MULTIPLE_SLICE_SEGMENTS_PER_TILE_BIT_KHR = 0x00000200,
} VkVideoEncodeH265CapabilityFlagBitsKHR;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_VIDEO_ENCODE_H265_CAPABILITY_HRD_COMPLIANCE_BIT_KHR` indicates if
  the implementation **may** be able to generate HRD compliant
  bitstreams if any of the `nal_hrd_parameters_present_flag`,
  `vcl_hrd_parameters_present_flag`, or
  `sub_pic_hrd_params_present_flag` members of `StdVideoH265HrdFlags`
  are set to `1` in the HRD parameters of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265-active-vps"
  target="_blank" rel="noopener">active VPS</a> or <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265-active-sps"
  target="_blank" rel="noopener">active SPS</a>, or if
  `StdVideoH265SpsVuiFlags`::`vui_hrd_parameters_present_flag` is set to
  `1` in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265-active-sps"
  target="_blank" rel="noopener">active SPS</a>.

- `VK_VIDEO_ENCODE_H265_CAPABILITY_PREDICTION_WEIGHT_TABLE_GENERATED_BIT_KHR`
  indicates that if the `weighted_pred_flag` or the
  `weighted_bipred_flag` member of `StdVideoH265PpsFlags` is set to `1`
  in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265-active-pps"
  target="_blank" rel="noopener">active PPS</a> when encoding a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265-p-pic"
  target="_blank" rel="noopener">P picture</a> or <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265-b-pic"
  target="_blank" rel="noopener">B picture</a>, respectively, then the
  implementation is able to internally decide syntax for
  `pred_weight_table`, as defined in section 7.4.7.3 of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h265"
  target="_blank" rel="noopener">ITU-T H.265 Specification</a>, and the
  application is not **required** to provide a weight table in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265-slice-segment-header-params"
  target="_blank" rel="noopener">H.265 slice segment header parameters</a>.

- `VK_VIDEO_ENCODE_H265_CAPABILITY_ROW_UNALIGNED_SLICE_SEGMENT_BIT_KHR`
  indicates that each slice segment in a frame with a single or multiple
  tiles per slice may begin or finish at any offset in a CTB row. If not
  supported, all slice segments in such a frame **must** begin at the
  start of a CTB row (and hence each slice segment **must** finish at
  the end of a CTB row). Also indicates that each slice segment in a
  frame with multiple slices per tile may begin or finish at any offset
  within the enclosing tile’s CTB row. If not supported, slice segments
  in such a frame **must** begin at the start of the enclosing tile’s
  CTB row (and hence each slice segment **must** finish at the end of
  the enclosing tile’s CTB row).

- `VK_VIDEO_ENCODE_H265_CAPABILITY_DIFFERENT_SLICE_SEGMENT_TYPE_BIT_KHR`
  indicates that when a frame is encoded with multiple slice segments,
  the implementation allows encoding each slice segment with a different
  `StdVideoEncodeH265SliceSegmentHeader`::`slice_type` specified in the
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265-slice-segment-header-params"
  target="_blank" rel="noopener">H.265 slice segment header parameters</a>.
  If not supported, all slice segments of the frame **must** be encoded
  with the same `slice_type` which corresponds to the picture type of
  the frame.

- `VK_VIDEO_ENCODE_H265_CAPABILITY_B_FRAME_IN_L0_LIST_BIT_KHR` indicates
  support for using a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265-b-pic"
  target="_blank" rel="noopener">B frame</a> as L0 reference, as
  specified in `StdVideoEncodeH265ReferenceListsInfo`::`RefPicList0` in
  the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265-picture-info"
  target="_blank" rel="noopener">H.265 picture information</a>.

- `VK_VIDEO_ENCODE_H265_CAPABILITY_B_FRAME_IN_L1_LIST_BIT_KHR` indicates
  support for using a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265-b-pic"
  target="_blank" rel="noopener">B frame</a> as L1 reference, as
  specified in `StdVideoEncodeH265ReferenceListsInfo`::`RefPicList1` in
  the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265-picture-info"
  target="_blank" rel="noopener">H.265 picture information</a>.

- `VK_VIDEO_ENCODE_H265_CAPABILITY_PER_PICTURE_TYPE_MIN_MAX_QP_BIT_KHR`
  indicates support for specifying different QP values in the members of
  [VkVideoEncodeH265QpKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265QpKHR.html).

- `VK_VIDEO_ENCODE_H265_CAPABILITY_PER_SLICE_SEGMENT_CONSTANT_QP_BIT_KHR`
  indicates support for specifying different constant QP values for each
  slice segment.

- `VK_VIDEO_ENCODE_H265_CAPABILITY_MULTIPLE_TILES_PER_SLICE_SEGMENT_BIT_KHR`
  indicates if encoding multiple tiles per slice segment, as defined in
  section 6.3.1 of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h265"
  target="_blank" rel="noopener">ITU-T H.265 Specification</a>, is
  supported. If this capability flag is not present, then the
  implementation is only able to encode a single tile for each slice
  segment.

- `VK_VIDEO_ENCODE_H265_CAPABILITY_MULTIPLE_SLICE_SEGMENTS_PER_TILE_BIT_KHR`
  indicates if encoding multiple slice segments per tile, as defined in
  section 6.3.1 of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h265"
  target="_blank" rel="noopener">ITU-T H.265 Specification</a>, is
  supported. If this capability flag is not present, then the
  implementation is only able to encode a single slice segment for each
  tile.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_encode_h265](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_encode_h265.html),
[VkVideoEncodeH265CapabilityFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265CapabilityFlagsKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoEncodeH265CapabilityFlagBitsKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
