# VkVideoEncodeH264CapabilityFlagBitsKHR(3) Manual Page

## Name

VkVideoEncodeH264CapabilityFlagBitsKHR - H.264 encode capability flags



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **may** be set in
[VkVideoEncodeH264CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264CapabilitiesKHR.html)::`flags`,
indicating the H.264 encoding capabilities supported, are:

``` c
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
} VkVideoEncodeH264CapabilityFlagBitsKHR;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_VIDEO_ENCODE_H264_CAPABILITY_HRD_COMPLIANCE_BIT_KHR` indicates
  whether the implementation **may** be able to generate HRD compliant
  bitstreams if any of the `nal_hrd_parameters_present_flag` or
  `vcl_hrd_parameters_present_flag` members of `StdVideoH264SpsVuiFlags`
  are set to `1` in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h264-active-sps"
  target="_blank" rel="noopener">active SPS</a>.

- `VK_VIDEO_ENCODE_H264_CAPABILITY_PREDICTION_WEIGHT_TABLE_GENERATED_BIT_KHR`
  indicates that if `StdVideoH264PpsFlags`::`weighted_pred_flag` is set
  to `1` or `StdVideoH264PictureParameterSet`::`weighted_bipred_idc` is
  set to `STD_VIDEO_H264_WEIGHTED_BIPRED_IDC_EXPLICIT` in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h264-active-pps"
  target="_blank" rel="noopener">active PPS</a> when encoding a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h264-p-pic"
  target="_blank" rel="noopener">P picture</a> or <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h264-b-pic"
  target="_blank" rel="noopener">B picture</a>, respectively, then the
  implementation is able to internally decide syntax for
  `pred_weight_table`, as defined in section 7.4.3.2 of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h264"
  target="_blank" rel="noopener">ITU-T H.264 Specification</a>, and the
  application is not **required** to provide a weight table in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h264-slice-header-params"
  target="_blank" rel="noopener">H.264 slice header parameters</a>.

- `VK_VIDEO_ENCODE_H264_CAPABILITY_ROW_UNALIGNED_SLICE_BIT_KHR`
  indicates that each slice in a frame with multiple slices may begin or
  finish at any offset in a macroblock row. If not supported, all slices
  in the frame **must** begin at the start of a macroblock row (and
  hence each slice **must** finish at the end of a macroblock row).

- `VK_VIDEO_ENCODE_H264_CAPABILITY_DIFFERENT_SLICE_TYPE_BIT_KHR`
  indicates that when a frame is encoded with multiple slices, the
  implementation allows encoding each slice with a different
  `StdVideoEncodeH264SliceHeader`::`slice_type` specified in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h264-slice-header-params"
  target="_blank" rel="noopener">H.264 slice header parameters</a>. If
  not supported, all slices of the frame **must** be encoded with the
  same `slice_type` which corresponds to the picture type of the frame.

- `VK_VIDEO_ENCODE_H264_CAPABILITY_B_FRAME_IN_L0_LIST_BIT_KHR` indicates
  support for using a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h264-b-pic"
  target="_blank" rel="noopener">B frame</a> as L0 reference, as
  specified in `StdVideoEncodeH264ReferenceListsInfo`::`RefPicList0` in
  the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h264-picture-info"
  target="_blank" rel="noopener">H.264 picture information</a>.

- `VK_VIDEO_ENCODE_H264_CAPABILITY_B_FRAME_IN_L1_LIST_BIT_KHR` indicates
  support for using a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h264-b-pic"
  target="_blank" rel="noopener">B frame</a> as L1 reference, as
  specified in `StdVideoEncodeH264ReferenceListsInfo`::`RefPicList1` in
  the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h264-picture-info"
  target="_blank" rel="noopener">H.264 picture information</a>.

- `VK_VIDEO_ENCODE_H264_CAPABILITY_PER_PICTURE_TYPE_MIN_MAX_QP_BIT_KHR`
  indicates support for specifying different QP values in the members of
  [VkVideoEncodeH264QpKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264QpKHR.html).

- `VK_VIDEO_ENCODE_H264_CAPABILITY_PER_SLICE_CONSTANT_QP_BIT_KHR`
  indicates support for specifying different constant QP values for each
  slice.

- `VK_VIDEO_ENCODE_H264_CAPABILITY_GENERATE_PREFIX_NALU_BIT_KHR`
  indicates support for generating prefix NAL units by setting
  [VkVideoEncodeH264PictureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264PictureInfoKHR.html)::`generatePrefixNalu`
  to `VK_TRUE`.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_encode_h264](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_encode_h264.html),
[VkVideoEncodeH264CapabilityFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264CapabilityFlagsKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoEncodeH264CapabilityFlagBitsKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
