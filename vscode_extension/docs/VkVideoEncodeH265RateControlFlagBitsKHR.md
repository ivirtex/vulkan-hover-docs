# VkVideoEncodeH265RateControlFlagBitsKHR(3) Manual Page

## Name

VkVideoEncodeH265RateControlFlagBitsKHR - H.265 encode rate control bits



## [](#_c_specification)C Specification

Bits which **can** be set in [VkVideoEncodeH265RateControlInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265RateControlInfoKHR.html)::`flags`, specifying H.265 rate control flags, are:

```c++
// Provided by VK_KHR_video_encode_h265
typedef enum VkVideoEncodeH265RateControlFlagBitsKHR {
    VK_VIDEO_ENCODE_H265_RATE_CONTROL_ATTEMPT_HRD_COMPLIANCE_BIT_KHR = 0x00000001,
    VK_VIDEO_ENCODE_H265_RATE_CONTROL_REGULAR_GOP_BIT_KHR = 0x00000002,
    VK_VIDEO_ENCODE_H265_RATE_CONTROL_REFERENCE_PATTERN_FLAT_BIT_KHR = 0x00000004,
    VK_VIDEO_ENCODE_H265_RATE_CONTROL_REFERENCE_PATTERN_DYADIC_BIT_KHR = 0x00000008,
    VK_VIDEO_ENCODE_H265_RATE_CONTROL_TEMPORAL_SUB_LAYER_PATTERN_DYADIC_BIT_KHR = 0x00000010,
} VkVideoEncodeH265RateControlFlagBitsKHR;
```

## [](#_description)Description

- `VK_VIDEO_ENCODE_H265_RATE_CONTROL_ATTEMPT_HRD_COMPLIANCE_BIT_KHR` specifies that rate control **should** attempt to produce an HRD compliant bitstream, as defined in annex C of the [ITU-T H.265 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#itu-t-h265).
- `VK_VIDEO_ENCODE_H265_RATE_CONTROL_REGULAR_GOP_BIT_KHR` specifies that the application intends to use a [regular GOP structure](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h265-regular-gop) according to the parameters specified in the `gopFrameCount`, `idrPeriod`, and `consecutiveBFrameCount` members of the [VkVideoEncodeH265RateControlInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265RateControlInfoKHR.html) structure.
- `VK_VIDEO_ENCODE_H265_RATE_CONTROL_REFERENCE_PATTERN_FLAT_BIT_KHR` specifies that the application intends to follow a [flat reference pattern](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h265-ref-pattern-flat) in the GOP.
- `VK_VIDEO_ENCODE_H265_RATE_CONTROL_REFERENCE_PATTERN_DYADIC_BIT_KHR` specifies that the application intends to follow a [dyadic reference pattern](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h265-ref-pattern-dyadic) in the GOP.
- `VK_VIDEO_ENCODE_H265_RATE_CONTROL_TEMPORAL_SUB_LAYER_PATTERN_DYADIC_BIT_KHR` specifies that the application intends to follow a [dyadic temporal sub-layer pattern](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h265-layer-pattern-dyadic).

## [](#_see_also)See Also

[VK\_KHR\_video\_encode\_h265](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_h265.html), [VkVideoEncodeH265RateControlFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265RateControlFlagsKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoEncodeH265RateControlFlagBitsKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0