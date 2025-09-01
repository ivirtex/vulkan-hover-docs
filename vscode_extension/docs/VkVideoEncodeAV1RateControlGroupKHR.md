# VkVideoEncodeAV1RateControlGroupKHR(3) Manual Page

## Name

VkVideoEncodeAV1RateControlGroupKHR - AV1 encode rate control group



## [](#_c_specification)C Specification

Possible AV1 encode rate control groups are as follows:

```c++
// Provided by VK_KHR_video_encode_av1
typedef enum VkVideoEncodeAV1RateControlGroupKHR {
    VK_VIDEO_ENCODE_AV1_RATE_CONTROL_GROUP_INTRA_KHR = 0,
    VK_VIDEO_ENCODE_AV1_RATE_CONTROL_GROUP_PREDICTIVE_KHR = 1,
    VK_VIDEO_ENCODE_AV1_RATE_CONTROL_GROUP_BIPREDICTIVE_KHR = 2,
} VkVideoEncodeAV1RateControlGroupKHR;
```

## [](#_description)Description

- `VK_VIDEO_ENCODE_AV1_RATE_CONTROL_GROUP_INTRA_KHR` **should** be specified when encoding AV1 frames that use intra-only prediction (e.g. when encoding AV1 frames of type `STD_VIDEO_AV1_FRAME_TYPE_KEY` or `STD_VIDEO_AV1_FRAME_TYPE_INTRA_ONLY`).
- `VK_VIDEO_ENCODE_AV1_RATE_CONTROL_GROUP_PREDICTIVE_KHR` **should** be specified when encoding AV1 frames that only have forward references in display order.
- `VK_VIDEO_ENCODE_AV1_RATE_CONTROL_GROUP_BIPREDICTIVE_KHR` **should** be specified when encoding AV1 frames that have backward references in display order.

Note

While the application can specify any rate control group for any frame, indifferent of the frame type, prediction mode, or prediction direction, specifying a rate control group that does not reflect the prediction direction used by the encoded frame may result in unexpected behavior of the implementation’s rate control algorithm.

## [](#_see_also)See Also

[VK\_KHR\_video\_encode\_av1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_av1.html), [VkVideoEncodeAV1PictureInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1PictureInfoKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoEncodeAV1RateControlGroupKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0