# VkVideoDecodeH264PictureLayoutFlagBitsKHR(3) Manual Page

## Name

VkVideoDecodeH264PictureLayoutFlagBitsKHR - H.264 video decode picture
layout flags



## <a href="#_c_specification" class="anchor"></a>C Specification

The H.264 video decode picture layout flags are defined as follows:

``` c
// Provided by VK_KHR_video_decode_h264
typedef enum VkVideoDecodeH264PictureLayoutFlagBitsKHR {
    VK_VIDEO_DECODE_H264_PICTURE_LAYOUT_PROGRESSIVE_KHR = 0,
    VK_VIDEO_DECODE_H264_PICTURE_LAYOUT_INTERLACED_INTERLEAVED_LINES_BIT_KHR = 0x00000001,
    VK_VIDEO_DECODE_H264_PICTURE_LAYOUT_INTERLACED_SEPARATE_PLANES_BIT_KHR = 0x00000002,
} VkVideoDecodeH264PictureLayoutFlagBitsKHR;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_VIDEO_DECODE_H264_PICTURE_LAYOUT_PROGRESSIVE_KHR` specifies
  support for progressive content. This flag has the value `0`.

- `VK_VIDEO_DECODE_H264_PICTURE_LAYOUT_INTERLACED_INTERLEAVED_LINES_BIT_KHR`
  specifies support for or use of a picture layout for interlaced
  content where all lines belonging to the top field are decoded to the
  even-numbered lines within the picture resource, and all lines
  belonging to the bottom field are decoded to the odd-numbered lines
  within the picture resource.

- `VK_VIDEO_DECODE_H264_PICTURE_LAYOUT_INTERLACED_SEPARATE_PLANES_BIT_KHR`
  specifies support for or use of a picture layout for interlaced
  content where all lines belonging to a field are grouped together in a
  single image subregion, and the two fields comprising the frame
  **can** be stored in separate image subregions of the same image
  subresource or in separate image subresources.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_decode_h264](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_decode_h264.html),
[VkVideoDecodeH264PictureLayoutFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH264PictureLayoutFlagsKHR.html),
[VkVideoDecodeH264ProfileInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH264ProfileInfoKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoDecodeH264PictureLayoutFlagBitsKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
