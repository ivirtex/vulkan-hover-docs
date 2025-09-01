# VkVideoEncodeH264CapabilityFlagBitsKHR(3) Manual Page

## Name

VkVideoEncodeH264CapabilityFlagBitsKHR - H.264 encode capability flags



## [](#_c_specification)C Specification

Bits which **may** be set in [VkVideoEncodeH264CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264CapabilitiesKHR.html)::`flags`, indicating the H.264 encoding capabilities supported, are:

```c++
// Provided by VK_KHR_video_encode_h264
typedef enum VkVideoEncodeH264CapabilityFlagBitsKHR {
    VK_VIDEO_ENCODE_H264_CAPABILITY_HRD_COMPLIANCE_BIT_KHR = 0x00000001,
    VK_VIDEO_ENCODE_H264_CAPABILITY_PREDICTION_WEIGHT_TABLE_GENERATED_BIT_KHR = 0x00000002,
    VK_VIDEO_ENCODE_H264_CAPABILITY_ROW_UNALIGNED_SLICE_BIT_KHR = 0x00000004,
    VK_VIDEO_ENCODE_H264_CAPABILITY_DIFFERENT_SLICE_TYPE_BIT_KHR = 0x00000008,
    VK_VIDEO_ENCODE_H264_CAPABILITY_B_FRAME_IN_L0_LIST_BIT_KHR = 0x00000010,
    VK_VIDEO_ENCODE_H264_CAPABILITY_B_FRAME_IN_L1_LIST_BIT_KHR = 0x00000020,
    VK_VIDEO_ENCODE_H264_CAPABILITY_PER_PICTURE_TYPE_MIN_MAX_QP_BIT_KHR = 0x00000040,
    VK_VIDEO_ENCODE_H264_CAPABILITY_PER_SLICE_CONSTANT_QP_BIT_KHR = 0x00000080,
    VK_VIDEO_ENCODE_H264_CAPABILITY_GENERATE_PREFIX_NALU_BIT_KHR = 0x00000100,
  // Provided by VK_KHR_video_encode_h264 with VK_KHR_video_encode_intra_refresh
    VK_VIDEO_ENCODE_H264_CAPABILITY_B_PICTURE_INTRA_REFRESH_BIT_KHR = 0x00000400,
  // Provided by VK_KHR_video_encode_h264 with VK_KHR_video_encode_quantization_map
    VK_VIDEO_ENCODE_H264_CAPABILITY_MB_QP_DIFF_WRAPAROUND_BIT_KHR = 0x00000200,
} VkVideoEncodeH264CapabilityFlagBitsKHR;
```

## [](#_description)Description

- `VK_VIDEO_ENCODE_H264_CAPABILITY_HRD_COMPLIANCE_BIT_KHR` specifies whether the implementation **may** be able to generate HRD compliant bitstreams if any of the `nal_hrd_parameters_present_flag` or `vcl_hrd_parameters_present_flag` members of `StdVideoH264SpsVuiFlags` are set to `1` in the [active SPS](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h264-active-sps).
- `VK_VIDEO_ENCODE_H264_CAPABILITY_PREDICTION_WEIGHT_TABLE_GENERATED_BIT_KHR` specifies that if `StdVideoH264PpsFlags`::`weighted_pred_flag` is set to `1` or `StdVideoH264PictureParameterSet`::`weighted_bipred_idc` is set to `STD_VIDEO_H264_WEIGHTED_BIPRED_IDC_EXPLICIT` in the [active PPS](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h264-active-pps) when encoding a [P picture](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h264-p-pic) or [B picture](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h264-b-pic), respectively, then the implementation is able to internally decide syntax for `pred_weight_table`, as defined in section 7.4.3.2 of the [ITU-T H.264 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#itu-t-h264), and the application is not **required** to provide a weight table in the [H.264 slice header parameters](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h264-slice-header-params).
- `VK_VIDEO_ENCODE_H264_CAPABILITY_ROW_UNALIGNED_SLICE_BIT_KHR` specifies that each slice in a frame with multiple slices may begin or finish at any offset in a macroblock row. If not supported, all slices in the frame **must** begin at the start of a macroblock row (and hence each slice **must** finish at the end of a macroblock row). When a picture is encoded with [intra refresh](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-intra-refresh), encoding non-rectangular slices also requires: support for the [VkVideoEncodeIntraRefreshCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeIntraRefreshCapabilitiesKHR.html)::`nonRectangularIntraRefreshRegions` capability.
- `VK_VIDEO_ENCODE_H264_CAPABILITY_DIFFERENT_SLICE_TYPE_BIT_KHR` specifies that when a frame is encoded with multiple slices, the implementation allows encoding each slice with a different `StdVideoEncodeH264SliceHeader`::`slice_type` specified in the [H.264 slice header parameters](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h264-slice-header-params). If not supported, all slices of the frame **must** be encoded with the same `slice_type` which corresponds to the picture type of the frame. There is one exception to this rule: if the picture is encoded with the [intra refresh mode](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-intra-refresh-modes) `VK_VIDEO_ENCODE_INTRA_REFRESH_MODE_PER_PICTURE_PARTITION_BIT_KHR`, then the currently refreshed slice **must** specify the `slice_type` `STD_VIDEO_H264_SLICE_TYPE_I` and **can** differ from the `slice_type` of the other slices regardless of whether `VK_VIDEO_ENCODE_H264_CAPABILITY_DIFFERENT_SLICE_TYPE_BIT_KHR` is supported.
- `VK_VIDEO_ENCODE_H264_CAPABILITY_B_FRAME_IN_L0_LIST_BIT_KHR` specifies support for using a [B frame](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h264-b-pic) as L0 reference, as specified in `StdVideoEncodeH264ReferenceListsInfo`::`RefPicList0` in the [H.264 picture information](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h264-picture-info).
- `VK_VIDEO_ENCODE_H264_CAPABILITY_B_FRAME_IN_L1_LIST_BIT_KHR` specifies support for using a [B frame](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h264-b-pic) as L1 reference, as specified in `StdVideoEncodeH264ReferenceListsInfo`::`RefPicList1` in the [H.264 picture information](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h264-picture-info).
- `VK_VIDEO_ENCODE_H264_CAPABILITY_PER_PICTURE_TYPE_MIN_MAX_QP_BIT_KHR` specifies support for specifying different QP values in the members of [VkVideoEncodeH264QpKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264QpKHR.html).
- `VK_VIDEO_ENCODE_H264_CAPABILITY_PER_SLICE_CONSTANT_QP_BIT_KHR` specifies support for specifying different constant QP values for each slice.
- `VK_VIDEO_ENCODE_H264_CAPABILITY_GENERATE_PREFIX_NALU_BIT_KHR` specifies support for generating prefix NAL units by setting [VkVideoEncodeH264PictureInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264PictureInfoKHR.html)::`generatePrefixNalu` to `VK_TRUE`.
- []()`VK_VIDEO_ENCODE_H264_CAPABILITY_MB_QP_DIFF_WRAPAROUND_BIT_KHR` indicates support for wraparound during the calculation of the QP values of subsequently encoded macroblocks, as defined in equation 7-37 of the [ITU-T H.264 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#itu-t-h264). If not supported, equation 7-37 of the [ITU-T H.264 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#itu-t-h264) is effectively reduced to the following:
  
  QPY = QPY,PREV + `mb_qp_delta`
  
  Note
  
  The effect of this is that the maximum QP difference across subsequent macroblocks is limited to the \[-(26 + QpBdOffsetY / 2), 25 + QpBdOffsetY / 2] range.
- `VK_VIDEO_ENCODE_H264_CAPABILITY_B_PICTURE_INTRA_REFRESH_BIT_KHR` indicates support for encoding [B pictures](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h264-b-pic) with [intra refresh](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-intra-refresh) enabled.

## [](#_see_also)See Also

[VK\_KHR\_video\_encode\_h264](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_h264.html), [VkVideoEncodeH264CapabilityFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264CapabilityFlagsKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoEncodeH264CapabilityFlagBitsKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0