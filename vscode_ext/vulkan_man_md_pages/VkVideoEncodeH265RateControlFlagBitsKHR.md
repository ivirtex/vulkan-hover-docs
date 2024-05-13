# VkVideoEncodeH265RateControlFlagBitsKHR(3) Manual Page

## Name

VkVideoEncodeH265RateControlFlagBitsKHR - H.265 encode rate control bits



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **can** be set in
[VkVideoEncodeH265RateControlInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265RateControlInfoKHR.html)::`flags`,
specifying H.265 rate control flags, are:

``` c
// Provided by VK_KHR_video_encode_h265
typedef enum VkVideoEncodeH265RateControlFlagBitsKHR {
    VK_VIDEO_ENCODE_H265_RATE_CONTROL_ATTEMPT_HRD_COMPLIANCE_BIT_KHR = 0x00000001,
    VK_VIDEO_ENCODE_H265_RATE_CONTROL_REGULAR_GOP_BIT_KHR = 0x00000002,
    VK_VIDEO_ENCODE_H265_RATE_CONTROL_REFERENCE_PATTERN_FLAT_BIT_KHR = 0x00000004,
    VK_VIDEO_ENCODE_H265_RATE_CONTROL_REFERENCE_PATTERN_DYADIC_BIT_KHR = 0x00000008,
    VK_VIDEO_ENCODE_H265_RATE_CONTROL_TEMPORAL_SUB_LAYER_PATTERN_DYADIC_BIT_KHR = 0x00000010,
} VkVideoEncodeH265RateControlFlagBitsKHR;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_VIDEO_ENCODE_H265_RATE_CONTROL_ATTEMPT_HRD_COMPLIANCE_BIT_KHR`
  specifies that rate control **should** attempt to produce an HRD
  compliant bitstream, as defined in annex C of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h265"
  target="_blank" rel="noopener">ITU-T H.265 Specification</a>.

- `VK_VIDEO_ENCODE_H265_RATE_CONTROL_REGULAR_GOP_BIT_KHR` specifies that
  the application intends to use a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265-regular-gop"
  target="_blank" rel="noopener">regular GOP structure</a> according to
  the parameters specified in the `gopFrameCount`, `idrPeriod`, and
  `consecutiveBFrameCount` members of the
  [VkVideoEncodeH265RateControlInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265RateControlInfoKHR.html)
  structure.

- `VK_VIDEO_ENCODE_H265_RATE_CONTROL_REFERENCE_PATTERN_FLAT_BIT_KHR`
  specifies that the application intends to follow a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265-ref-pattern-flat"
  target="_blank" rel="noopener">flat reference pattern</a> in the GOP.

- `VK_VIDEO_ENCODE_H265_RATE_CONTROL_REFERENCE_PATTERN_DYADIC_BIT_KHR`
  specifies that the application intends to follow a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265-ref-pattern-dyadic"
  target="_blank" rel="noopener">dyadic reference pattern</a> in the
  GOP.

- `VK_VIDEO_ENCODE_H265_RATE_CONTROL_TEMPORAL_SUB_LAYER_PATTERN_DYADIC_BIT_KHR`
  specifies that the application intends to follow a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265-layer-pattern-dyadic"
  target="_blank" rel="noopener">dyadic temporal sub-layer pattern</a>.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_encode_h265](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_encode_h265.html),
[VkVideoEncodeH265RateControlFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265RateControlFlagsKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoEncodeH265RateControlFlagBitsKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
