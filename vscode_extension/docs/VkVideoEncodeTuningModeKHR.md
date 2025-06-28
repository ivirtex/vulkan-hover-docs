# VkVideoEncodeTuningModeKHR(3) Manual Page

## Name

VkVideoEncodeTuningModeKHR - Video encode tuning mode



## [](#_c_specification)C Specification

Possible video encode tuning mode values are as follows:

```c++
// Provided by VK_KHR_video_encode_queue
typedef enum VkVideoEncodeTuningModeKHR {
    VK_VIDEO_ENCODE_TUNING_MODE_DEFAULT_KHR = 0,
    VK_VIDEO_ENCODE_TUNING_MODE_HIGH_QUALITY_KHR = 1,
    VK_VIDEO_ENCODE_TUNING_MODE_LOW_LATENCY_KHR = 2,
    VK_VIDEO_ENCODE_TUNING_MODE_ULTRA_LOW_LATENCY_KHR = 3,
    VK_VIDEO_ENCODE_TUNING_MODE_LOSSLESS_KHR = 4,
} VkVideoEncodeTuningModeKHR;
```

## [](#_description)Description

- `VK_VIDEO_ENCODE_TUNING_MODE_DEFAULT_KHR` specifies the default tuning mode.
- `VK_VIDEO_ENCODE_TUNING_MODE_HIGH_QUALITY_KHR` specifies that video encoding is tuned for high quality. When using this tuning mode, the implementation **may** compromise the latency of video encoding operations to improve quality.
- `VK_VIDEO_ENCODE_TUNING_MODE_LOW_LATENCY_KHR` specifies that video encoding is tuned for low latency. When using this tuning mode, the implementation **may** compromise quality to increase the performance and lower the latency of video encode operations.
- `VK_VIDEO_ENCODE_TUNING_MODE_ULTRA_LOW_LATENCY_KHR` specifies that video encoding is tuned for ultra-low latency. When using this tuning mode, the implementation **may** compromise quality to maximize the performance and minimize the latency of video encoding operations.
- `VK_VIDEO_ENCODE_TUNING_MODE_LOSSLESS_KHR` specifies that video encoding is tuned for lossless encoding. When using this tuning mode, video encode operations produce lossless output.

## [](#_see_also)See Also

[VK\_KHR\_video\_encode\_queue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_queue.html), [VkVideoEncodeUsageInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeUsageInfoKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoEncodeTuningModeKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0