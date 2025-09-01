# VkVideoDecodeUsageFlagBitsKHR(3) Manual Page

## Name

VkVideoDecodeUsageFlagBitsKHR - Video decode usage flags



## [](#_c_specification)C Specification

The following bits **can** be specified in [VkVideoDecodeUsageInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeUsageInfoKHR.html)::`videoUsageHints` as a hint about the video decode use case:

```c++
// Provided by VK_KHR_video_decode_queue
typedef enum VkVideoDecodeUsageFlagBitsKHR {
    VK_VIDEO_DECODE_USAGE_DEFAULT_KHR = 0,
    VK_VIDEO_DECODE_USAGE_TRANSCODING_BIT_KHR = 0x00000001,
    VK_VIDEO_DECODE_USAGE_OFFLINE_BIT_KHR = 0x00000002,
    VK_VIDEO_DECODE_USAGE_STREAMING_BIT_KHR = 0x00000004,
} VkVideoDecodeUsageFlagBitsKHR;
```

## [](#_description)Description

- `VK_VIDEO_DECODE_USAGE_TRANSCODING_BIT_KHR` specifies that video decoding is intended to be used in conjunction with video encoding to transcode a video bitstream with the same and/or different codecs.
- `VK_VIDEO_DECODE_USAGE_OFFLINE_BIT_KHR` specifies that video decoding is intended to be used to consume a local video bitstream.
- `VK_VIDEO_DECODE_USAGE_STREAMING_BIT_KHR` specifies that video decoding is intended to be used to consume a video bitstream received as a continuous flow over network.

Note

There are no restrictions on the combination of bits that **can** be specified by the application. However, applications **should** use reasonable combinations in order for the implementation to be able to select the most appropriate mode of operation for the particular use case.

## [](#_see_also)See Also

[VK\_KHR\_video\_decode\_queue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_decode_queue.html), [VkVideoDecodeUsageFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeUsageFlagsKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoDecodeUsageFlagBitsKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0