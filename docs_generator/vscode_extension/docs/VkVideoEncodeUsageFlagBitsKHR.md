# VkVideoEncodeUsageFlagBitsKHR(3) Manual Page

## Name

VkVideoEncodeUsageFlagBitsKHR - Video encode usage flags



## [](#_c_specification)C Specification

The following bits **can** be specified in [VkVideoEncodeUsageInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeUsageInfoKHR.html)::`videoUsageHints` as a hint about the video encode use case:

```c++
// Provided by VK_KHR_video_encode_queue
typedef enum VkVideoEncodeUsageFlagBitsKHR {
    VK_VIDEO_ENCODE_USAGE_DEFAULT_KHR = 0,
    VK_VIDEO_ENCODE_USAGE_TRANSCODING_BIT_KHR = 0x00000001,
    VK_VIDEO_ENCODE_USAGE_STREAMING_BIT_KHR = 0x00000002,
    VK_VIDEO_ENCODE_USAGE_RECORDING_BIT_KHR = 0x00000004,
    VK_VIDEO_ENCODE_USAGE_CONFERENCING_BIT_KHR = 0x00000008,
} VkVideoEncodeUsageFlagBitsKHR;
```

## [](#_description)Description

- `VK_VIDEO_ENCODE_USAGE_TRANSCODING_BIT_KHR` specifies that video encoding is intended to be used in conjunction with video decoding to transcode a video bitstream with the same and/or different codecs.
- `VK_VIDEO_ENCODE_USAGE_STREAMING_BIT_KHR` specifies that video encoding is intended to be used to produce a video bitstream that is expected to be sent as a continuous flow over network.
- `VK_VIDEO_ENCODE_USAGE_RECORDING_BIT_KHR` specifies that video encoding is intended to be used for real-time recording for offline consumption.
- `VK_VIDEO_ENCODE_USAGE_CONFERENCING_BIT_KHR` specifies that video encoding is intended to be used in a video conferencing scenario.

Note

There are no restrictions on the combination of bits that **can** be specified by the application. However, applications **should** use reasonable combinations in order for the implementation to be able to select the most appropriate mode of operation for the particular use case.

## [](#_see_also)See Also

[VK\_KHR\_video\_encode\_queue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_queue.html), [VkVideoEncodeUsageFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeUsageFlagsKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoEncodeUsageFlagBitsKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0