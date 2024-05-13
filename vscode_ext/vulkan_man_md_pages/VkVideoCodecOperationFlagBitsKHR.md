# VkVideoCodecOperationFlagBitsKHR(3) Manual Page

## Name

VkVideoCodecOperationFlagBitsKHR - Video codec operation bits



## <a href="#_c_specification" class="anchor"></a>C Specification

Possible values of
[VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileInfoKHR.html)::`videoCodecOperation`,
specifying the type of video coding operation and video compression
standard used by a video profile, are:

``` c
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
} VkVideoCodecOperationFlagBitsKHR;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_VIDEO_CODEC_OPERATION_NONE_KHR` indicates no support for any video
  codec operations.

- `VK_VIDEO_CODEC_OPERATION_DECODE_H264_BIT_KHR` specifies support for
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h264"
  target="_blank" rel="noopener">H.264 decode operations</a>.

- `VK_VIDEO_CODEC_OPERATION_DECODE_H265_BIT_KHR` specifies support for
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h265"
  target="_blank" rel="noopener">H.265 decode operations</a>.

- `VK_VIDEO_CODEC_OPERATION_DECODE_AV1_BIT_KHR` specifies support for <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-av1"
  target="_blank" rel="noopener">AV1 decode operations</a>.

- `VK_VIDEO_CODEC_OPERATION_ENCODE_H264_BIT_KHR` specifies support for
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h264"
  target="_blank" rel="noopener">H.264 encode operations</a>.

- `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR` specifies support for
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265"
  target="_blank" rel="noopener">H.265 encode operations</a>.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_queue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_queue.html),
[VkVideoCodecOperationFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCodecOperationFlagsKHR.html),
[VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileInfoKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoCodecOperationFlagBitsKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
