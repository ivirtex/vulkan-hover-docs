# VkVideoEncodeAV1RateControlFlagBitsKHR(3) Manual Page

## Name

VkVideoEncodeAV1RateControlFlagBitsKHR - AV1 encode rate control bits



## [](#_c_specification)C Specification

Bits which **can** be set in [VkVideoEncodeAV1RateControlInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1RateControlInfoKHR.html)::`flags`, specifying AV1 rate control flags, are:

```c++
// Provided by VK_KHR_video_encode_av1
typedef enum VkVideoEncodeAV1RateControlFlagBitsKHR {
    VK_VIDEO_ENCODE_AV1_RATE_CONTROL_REGULAR_GOP_BIT_KHR = 0x00000001,
    VK_VIDEO_ENCODE_AV1_RATE_CONTROL_TEMPORAL_LAYER_PATTERN_DYADIC_BIT_KHR = 0x00000002,
    VK_VIDEO_ENCODE_AV1_RATE_CONTROL_REFERENCE_PATTERN_FLAT_BIT_KHR = 0x00000004,
    VK_VIDEO_ENCODE_AV1_RATE_CONTROL_REFERENCE_PATTERN_DYADIC_BIT_KHR = 0x00000008,
} VkVideoEncodeAV1RateControlFlagBitsKHR;
```

## [](#_description)Description

- `VK_VIDEO_ENCODE_AV1_RATE_CONTROL_REGULAR_GOP_BIT_KHR` specifies that the application intends to use a [regular GOP structure](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-regular-gop) according to the parameters specified in the `gopFrameCount` and `keyFramePeriod` members of the [VkVideoEncodeAV1RateControlInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1RateControlInfoKHR.html) structure.
- `VK_VIDEO_ENCODE_AV1_RATE_CONTROL_TEMPORAL_LAYER_PATTERN_DYADIC_BIT_KHR` specifies that the application intends to follow a [dyadic temporal layer pattern](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-layer-pattern-dyadic).
- `VK_VIDEO_ENCODE_AV1_RATE_CONTROL_REFERENCE_PATTERN_FLAT_BIT_KHR` specifies that the application intends to follow a [flat reference pattern](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-ref-pattern-flat) in the GOP.
- `VK_VIDEO_ENCODE_AV1_RATE_CONTROL_REFERENCE_PATTERN_DYADIC_BIT_KHR` specifies that the application intends to follow a [dyadic reference pattern](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-ref-pattern-dyadic) in the GOP.

## [](#_see_also)See Also

[VK\_KHR\_video\_encode\_av1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_av1.html), [VkVideoEncodeAV1RateControlFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1RateControlFlagsKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoEncodeAV1RateControlFlagBitsKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0