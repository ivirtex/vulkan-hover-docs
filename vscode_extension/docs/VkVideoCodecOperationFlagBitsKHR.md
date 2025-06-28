# VkVideoCodecOperationFlagBitsKHR(3) Manual Page

## Name

VkVideoCodecOperationFlagBitsKHR - Video codec operation bits



## [](#_c_specification)C Specification

Possible values of [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoProfileInfoKHR.html)::`videoCodecOperation`, specifying the type of video coding operation and video compression standard used by a video profile, are:

```c++
// Provided by VK_KHR_video_queue
typedef enum VkVideoCodecOperationFlagBitsKHR {
    VK_VIDEO_CODEC_OPERATION_NONE_KHR = 0,
  // Provided by VK_KHR_video_encode_h264
    VK_VIDEO_CODEC_OPERATION_ENCODE_H264_BIT_KHR = 0x00010000,
  // Provided by VK_KHR_video_encode_h265
    VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR = 0x00020000,
  // Provided by VK_KHR_video_decode_h264
    VK_VIDEO_CODEC_OPERATION_DECODE_H264_BIT_KHR = 0x00000001,
  // Provided by VK_KHR_video_decode_h265
    VK_VIDEO_CODEC_OPERATION_DECODE_H265_BIT_KHR = 0x00000002,
  // Provided by VK_KHR_video_decode_av1
    VK_VIDEO_CODEC_OPERATION_DECODE_AV1_BIT_KHR = 0x00000004,
  // Provided by VK_KHR_video_encode_av1
    VK_VIDEO_CODEC_OPERATION_ENCODE_AV1_BIT_KHR = 0x00040000,
  // Provided by VK_KHR_video_decode_vp9
    VK_VIDEO_CODEC_OPERATION_DECODE_VP9_BIT_KHR = 0x00000008,
} VkVideoCodecOperationFlagBitsKHR;
```

## [](#_description)Description

- `VK_VIDEO_CODEC_OPERATION_NONE_KHR` specifies that no video codec operations are supported.
- `VK_VIDEO_CODEC_OPERATION_DECODE_H264_BIT_KHR` specifies support for [H.264 decode operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-h264).
- `VK_VIDEO_CODEC_OPERATION_DECODE_H265_BIT_KHR` specifies support for [H.265 decode operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-h265).
- `VK_VIDEO_CODEC_OPERATION_DECODE_VP9_BIT_KHR` specifies support for [VP9 decode operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-vp9).
- `VK_VIDEO_CODEC_OPERATION_DECODE_AV1_BIT_KHR` specifies support for [AV1 decode operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-av1).
- `VK_VIDEO_CODEC_OPERATION_ENCODE_H264_BIT_KHR` specifies support for [H.264 encode operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h264).
- `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR` specifies support for [H.265 encode operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h265).
- `VK_VIDEO_CODEC_OPERATION_ENCODE_AV1_BIT_KHR` specifies support for [AV1 encode operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1).

## [](#_see_also)See Also

[VK\_KHR\_video\_queue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_queue.html), [VkVideoCodecOperationFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoCodecOperationFlagsKHR.html), [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoProfileInfoKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoCodecOperationFlagBitsKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0