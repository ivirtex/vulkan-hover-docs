# VkVideoDecodeH264PictureLayoutFlagBitsKHR(3) Manual Page

## Name

VkVideoDecodeH264PictureLayoutFlagBitsKHR - H.264 video decode picture layout flags



## [](#_c_specification)C Specification

The H.264 video decode picture layout flags are defined as follows:

```c++
// Provided by VK_KHR_video_decode_h264
typedef enum VkVideoDecodeH264PictureLayoutFlagBitsKHR {
    VK_VIDEO_DECODE_H264_PICTURE_LAYOUT_PROGRESSIVE_KHR = 0,
    VK_VIDEO_DECODE_H264_PICTURE_LAYOUT_INTERLACED_INTERLEAVED_LINES_BIT_KHR = 0x00000001,
    VK_VIDEO_DECODE_H264_PICTURE_LAYOUT_INTERLACED_SEPARATE_PLANES_BIT_KHR = 0x00000002,
} VkVideoDecodeH264PictureLayoutFlagBitsKHR;
```

## [](#_description)Description

- `VK_VIDEO_DECODE_H264_PICTURE_LAYOUT_PROGRESSIVE_KHR` specifies support for progressive content. This flag has the value `0`.
- `VK_VIDEO_DECODE_H264_PICTURE_LAYOUT_INTERLACED_INTERLEAVED_LINES_BIT_KHR` specifies support for or use of a picture layout for interlaced content where all lines belonging to the top field are decoded to the even-numbered lines within the picture resource, and all lines belonging to the bottom field are decoded to the odd-numbered lines within the picture resource.
- `VK_VIDEO_DECODE_H264_PICTURE_LAYOUT_INTERLACED_SEPARATE_PLANES_BIT_KHR` specifies support for or use of a picture layout for interlaced content where all lines belonging to a field are grouped together in a single image subregion, and the two fields comprising the frame **can** be stored in separate image subregions of the same image subresource or in separate image subresources.

## [](#_see_also)See Also

[VK\_KHR\_video\_decode\_h264](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_decode_h264.html), [VkVideoDecodeH264PictureLayoutFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH264PictureLayoutFlagsKHR.html), [VkVideoDecodeH264ProfileInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH264ProfileInfoKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoDecodeH264PictureLayoutFlagBitsKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0