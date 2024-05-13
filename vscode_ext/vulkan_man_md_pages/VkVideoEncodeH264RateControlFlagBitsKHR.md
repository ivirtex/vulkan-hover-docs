# VkVideoEncodeH264RateControlFlagBitsKHR(3) Manual Page

## Name

VkVideoEncodeH264RateControlFlagBitsKHR - H.264 encode rate control bits



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **can** be set in
[VkVideoEncodeH264RateControlInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264RateControlInfoKHR.html)::`flags`,
specifying H.264 rate control flags, are:

``` c
// Provided by VK_KHR_video_encode_h264
typedef enum VkVideoEncodeH264RateControlFlagBitsKHR {
    VK_VIDEO_ENCODE_H264_RATE_CONTROL_ATTEMPT_HRD_COMPLIANCE_BIT_KHR = 0x00000001,
    VK_VIDEO_ENCODE_H264_RATE_CONTROL_REGULAR_GOP_BIT_KHR = 0x00000002,
    VK_VIDEO_ENCODE_H264_RATE_CONTROL_REFERENCE_PATTERN_FLAT_BIT_KHR = 0x00000004,
    VK_VIDEO_ENCODE_H264_RATE_CONTROL_REFERENCE_PATTERN_DYADIC_BIT_KHR = 0x00000008,
    VK_VIDEO_ENCODE_H264_RATE_CONTROL_TEMPORAL_LAYER_PATTERN_DYADIC_BIT_KHR = 0x00000010,
} VkVideoEncodeH264RateControlFlagBitsKHR;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_VIDEO_ENCODE_H264_RATE_CONTROL_ATTEMPT_HRD_COMPLIANCE_BIT_KHR`
  specifies that rate control **should** attempt to produce an HRD
  compliant bitstream, as defined in annex C of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h264"
  target="_blank" rel="noopener">ITU-T H.264 Specification</a>.

- `VK_VIDEO_ENCODE_H264_RATE_CONTROL_REGULAR_GOP_BIT_KHR` specifies that
  the application intends to use a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h264-regular-gop"
  target="_blank" rel="noopener">regular GOP structure</a> according to
  the parameters specified in the `gopFrameCount`, `idrPeriod`, and
  `consecutiveBFrameCount` members of the
  [VkVideoEncodeH264RateControlInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264RateControlInfoKHR.html)
  structure.

- `VK_VIDEO_ENCODE_H264_RATE_CONTROL_REFERENCE_PATTERN_FLAT_BIT_KHR`
  specifies that the application intends to follow a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h264-ref-pattern-flat"
  target="_blank" rel="noopener">flat reference pattern</a> in the GOP.

- `VK_VIDEO_ENCODE_H264_RATE_CONTROL_REFERENCE_PATTERN_DYADIC_BIT_KHR`
  specifies that the application intends to follow a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h264-ref-pattern-dyadic"
  target="_blank" rel="noopener">dyadic reference pattern</a> in the
  GOP.

- `VK_VIDEO_ENCODE_H264_RATE_CONTROL_TEMPORAL_LAYER_PATTERN_DYADIC_BIT_KHR`
  specifies that the application intends to follow a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h264-layer-pattern-dyadic"
  target="_blank" rel="noopener">dyadic temporal layer pattern</a>.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_encode_h264](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_encode_h264.html),
[VkVideoEncodeH264RateControlFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264RateControlFlagsKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoEncodeH264RateControlFlagBitsKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
