# VkVideoEncodeH264StdFlagBitsKHR(3) Manual Page

## Name

VkVideoEncodeH264StdFlagBitsKHR - Video encode H.264 syntax capability flags



## [](#_c_specification)C Specification

Bits which **may** be set in [VkVideoEncodeH264CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264CapabilitiesKHR.html)::`stdSyntaxFlags`, indicating the capabilities related to the H.264 syntax elements, are:

```c++
// Provided by VK_KHR_video_encode_h264
typedef enum VkVideoEncodeH264StdFlagBitsKHR {
    VK_VIDEO_ENCODE_H264_STD_SEPARATE_COLOR_PLANE_FLAG_SET_BIT_KHR = 0x00000001,
    VK_VIDEO_ENCODE_H264_STD_QPPRIME_Y_ZERO_TRANSFORM_BYPASS_FLAG_SET_BIT_KHR = 0x00000002,
    VK_VIDEO_ENCODE_H264_STD_SCALING_MATRIX_PRESENT_FLAG_SET_BIT_KHR = 0x00000004,
    VK_VIDEO_ENCODE_H264_STD_CHROMA_QP_INDEX_OFFSET_BIT_KHR = 0x00000008,
    VK_VIDEO_ENCODE_H264_STD_SECOND_CHROMA_QP_INDEX_OFFSET_BIT_KHR = 0x00000010,
    VK_VIDEO_ENCODE_H264_STD_PIC_INIT_QP_MINUS26_BIT_KHR = 0x00000020,
    VK_VIDEO_ENCODE_H264_STD_WEIGHTED_PRED_FLAG_SET_BIT_KHR = 0x00000040,
    VK_VIDEO_ENCODE_H264_STD_WEIGHTED_BIPRED_IDC_EXPLICIT_BIT_KHR = 0x00000080,
    VK_VIDEO_ENCODE_H264_STD_WEIGHTED_BIPRED_IDC_IMPLICIT_BIT_KHR = 0x00000100,
    VK_VIDEO_ENCODE_H264_STD_TRANSFORM_8X8_MODE_FLAG_SET_BIT_KHR = 0x00000200,
    VK_VIDEO_ENCODE_H264_STD_DIRECT_SPATIAL_MV_PRED_FLAG_UNSET_BIT_KHR = 0x00000400,
    VK_VIDEO_ENCODE_H264_STD_ENTROPY_CODING_MODE_FLAG_UNSET_BIT_KHR = 0x00000800,
    VK_VIDEO_ENCODE_H264_STD_ENTROPY_CODING_MODE_FLAG_SET_BIT_KHR = 0x00001000,
    VK_VIDEO_ENCODE_H264_STD_DIRECT_8X8_INFERENCE_FLAG_UNSET_BIT_KHR = 0x00002000,
    VK_VIDEO_ENCODE_H264_STD_CONSTRAINED_INTRA_PRED_FLAG_SET_BIT_KHR = 0x00004000,
    VK_VIDEO_ENCODE_H264_STD_DEBLOCKING_FILTER_DISABLED_BIT_KHR = 0x00008000,
    VK_VIDEO_ENCODE_H264_STD_DEBLOCKING_FILTER_ENABLED_BIT_KHR = 0x00010000,
    VK_VIDEO_ENCODE_H264_STD_DEBLOCKING_FILTER_PARTIAL_BIT_KHR = 0x00020000,
    VK_VIDEO_ENCODE_H264_STD_SLICE_QP_DELTA_BIT_KHR = 0x00080000,
    VK_VIDEO_ENCODE_H264_STD_DIFFERENT_SLICE_QP_DELTA_BIT_KHR = 0x00100000,
} VkVideoEncodeH264StdFlagBitsKHR;
```

## [](#_description)Description

- `VK_VIDEO_ENCODE_H264_STD_SEPARATE_COLOR_PLANE_FLAG_SET_BIT_KHR` specifies whether the implementation supports using the application-provided value for `StdVideoH264SpsFlags`::`separate_colour_plane_flag` in the [SPS](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h264-sps) when that value is `1`.
- `VK_VIDEO_ENCODE_H264_STD_QPPRIME_Y_ZERO_TRANSFORM_BYPASS_FLAG_SET_BIT_KHR` specifies whether the implementation supports using the application-provided value for `StdVideoH264SpsFlags`::`qpprime_y_zero_transform_bypass_flag` in the [SPS](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h264-sps) when that value is `1`.
- `VK_VIDEO_ENCODE_H264_STD_SCALING_MATRIX_PRESENT_FLAG_SET_BIT_KHR` specifies whether the implementation supports using the application-provided values for `StdVideoH264SpsFlags`::`seq_scaling_matrix_present_flag` in the [SPS](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h264-sps) and `StdVideoH264PpsFlags`::`pic_scaling_matrix_present_flag` in the [PPS](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h264-pps) when any of those values are `1`.
- `VK_VIDEO_ENCODE_H264_STD_CHROMA_QP_INDEX_OFFSET_BIT_KHR` specifies whether the implementation supports using the application-provided value for `StdVideoH264PictureParameterSet`::`chroma_qp_index_offset` in the [PPS](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h264-pps) when that value is non-zero.
- `VK_VIDEO_ENCODE_H264_STD_SECOND_CHROMA_QP_INDEX_OFFSET_BIT_KHR` specifies whether the implementation supports using the application-provided value for `StdVideoH264PictureParameterSet`::`second_chroma_qp_index_offset` in the [PPS](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h264-pps) when that value is non-zero.
- `VK_VIDEO_ENCODE_H264_STD_PIC_INIT_QP_MINUS26_BIT_KHR` specifies whether the implementation supports using the application-provided value for `StdVideoH264PictureParameterSet`::`pic_init_qp_minus26` in the [PPS](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h264-pps) when that value is non-zero.
- `VK_VIDEO_ENCODE_H264_STD_WEIGHTED_PRED_FLAG_SET_BIT_KHR` specifies whether the implementation supports using the application-provided value for `StdVideoH264PpsFlags`::`weighted_pred_flag` in the [PPS](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h264-pps) when that value is `1`.
- `VK_VIDEO_ENCODE_H264_STD_WEIGHTED_BIPRED_IDC_EXPLICIT_BIT_KHR` specifies whether the implementation supports using the application-provided value for `StdVideoH264PictureParameterSet`::`weighted_bipred_idc` in the [PPS](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h264-pps) when that value is `STD_VIDEO_H264_WEIGHTED_BIPRED_IDC_EXPLICIT`.
- `VK_VIDEO_ENCODE_H264_STD_WEIGHTED_BIPRED_IDC_IMPLICIT_BIT_KHR` specifies whether the implementation supports using the application-provided value for `StdVideoH264PictureParameterSet`::`weighted_bipred_idc` in the [PPS](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h264-pps) when that value is `STD_VIDEO_H264_WEIGHTED_BIPRED_IDC_IMPLICIT`.
- `VK_VIDEO_ENCODE_H264_STD_TRANSFORM_8X8_MODE_FLAG_SET_BIT_KHR` specifies whether the implementation supports using the application-provided value for `StdVideoH264PpsFlags`::`transform_8x8_mode_flag` in the [PPS](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h264-pps) when that value is `1`.
- `VK_VIDEO_ENCODE_H264_STD_DIRECT_SPATIAL_MV_PRED_FLAG_UNSET_BIT_KHR` specifies whether the implementation supports using the application-provided value for `StdVideoEncodeH264SliceHeaderFlags`::`direct_spatial_mv_pred_flag` in the [H.264 slice header parameters](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h264-slice-header-params) when that value is `0`.
- `VK_VIDEO_ENCODE_H264_STD_ENTROPY_CODING_MODE_FLAG_UNSET_BIT_KHR` specifies whether the implementation supports CAVLC entropy coding, as defined in section 9.2 of the [ITU-T H.264 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#itu-t-h264), and thus supports using the application-provided value for `StdVideoH264PpsFlags`::`entropy_coding_mode_flag` in the [PPS](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h264-pps) when that value is `0`.
- `VK_VIDEO_ENCODE_H264_STD_ENTROPY_CODING_MODE_FLAG_SET_BIT_KHR` specifies whether the implementation supports CABAC entropy coding, as defined in section 9.3 of the [ITU-T H.264 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#itu-t-h264), and thus supports using the application-provided value for `StdVideoH264PpsFlags`::`entropy_coding_mode_flag` in the [PPS](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h264-pps) when that value is `1`.
- `VK_VIDEO_ENCODE_H264_STD_DIRECT_8X8_INFERENCE_FLAG_UNSET_BIT_KHR` specifies whether the implementation supports using the application-provided value for `StdVideoH264SpsFlags`::`direct_8x8_inference_flag` in the [SPS](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h264-sps) when that value is `0`.
- `VK_VIDEO_ENCODE_H264_STD_CONSTRAINED_INTRA_PRED_FLAG_SET_BIT_KHR` specifies whether the implementation supports using the application-provided value for `StdVideoH264PpsFlags`::`constrained_intra_pred_flag` in the [PPS](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h264-pps) when that value is `1`.
- `VK_VIDEO_ENCODE_H264_STD_DEBLOCKING_FILTER_DISABLED_BIT_KHR` specifies whether the implementation supports using the application-provided value for `StdVideoEncodeH264SliceHeader`::`disable_deblocking_filter_idc` in the [H.264 slice header parameters](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h264-slice-header-params) when that value is `STD_VIDEO_H264_DISABLE_DEBLOCKING_FILTER_IDC_DISABLED`.
- `VK_VIDEO_ENCODE_H264_STD_DEBLOCKING_FILTER_ENABLED_BIT_KHR` specifies whether the implementation supports using the application-provided value for `StdVideoEncodeH264SliceHeader`::`disable_deblocking_filter_idc` in the [H.264 slice header parameters](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h264-slice-header-params) when that value is `STD_VIDEO_H264_DISABLE_DEBLOCKING_FILTER_IDC_ENABLED`.
- `VK_VIDEO_ENCODE_H264_STD_DEBLOCKING_FILTER_PARTIAL_BIT_KHR` specifies whether the implementation supports using the application-provided value for `StdVideoEncodeH264SliceHeader`::`disable_deblocking_filter_idc` in the [H.264 slice header parameters](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h264-slice-header-params) when that value is `STD_VIDEO_H264_DISABLE_DEBLOCKING_FILTER_IDC_PARTIAL`.
- `VK_VIDEO_ENCODE_H264_STD_SLICE_QP_DELTA_BIT_KHR` specifies whether the implementation supports using the application-provided value for `StdVideoEncodeH264SliceHeader`::`slice_qp_delta` in the [H.264 slice header parameters](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h264-slice-header-params) when that value is identical across the slices of the encoded frame.
- `VK_VIDEO_ENCODE_H264_STD_DIFFERENT_SLICE_QP_DELTA_BIT_KHR` specifies whether the implementation supports using the application-provided value for `StdVideoEncodeH264SliceHeader`::`slice_qp_delta` in the [H.264 slice header parameters](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h264-slice-header-params) when that value is different across the slices of the encoded frame.

These capability flags provide information to the application about specific H.264 syntax element values that the implementation supports without having to [override](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h264-overrides) them and do not otherwise restrict the values that the application **can** specify for any of the mentioned H.264 syntax elements.

## [](#_see_also)See Also

[VK\_KHR\_video\_encode\_h264](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_h264.html), [VkVideoEncodeH264StdFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264StdFlagsKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoEncodeH264StdFlagBitsKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0