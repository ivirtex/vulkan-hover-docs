# VkVideoDecodeH264ProfileInfoKHR(3) Manual Page

## Name

VkVideoDecodeH264ProfileInfoKHR - Structure specifying H.264
decode-specific video profile parameters



## <a href="#_c_specification" class="anchor"></a>C Specification

A video profile supporting H.264 video decode operations is specified by
setting
[VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileInfoKHR.html)::`videoCodecOperation`
to `VK_VIDEO_CODEC_OPERATION_DECODE_H264_BIT_KHR` and adding a
`VkVideoDecodeH264ProfileInfoKHR` structure to the
[VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileInfoKHR.html)::`pNext` chain.

The `VkVideoDecodeH264ProfileInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_video_decode_h264
typedef struct VkVideoDecodeH264ProfileInfoKHR {
    VkStructureType                              sType;
    const void*                                  pNext;
    StdVideoH264ProfileIdc                       stdProfileIdc;
    VkVideoDecodeH264PictureLayoutFlagBitsKHR    pictureLayout;
} VkVideoDecodeH264ProfileInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `stdProfileIdc` is a `StdVideoH264ProfileIdc` value specifying the
  H.264 codec profile IDC, as defined in section A.2 of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h264"
  target="_blank" rel="noopener">ITU-T H.264 Specification</a>.

- `pictureLayout` is a
  [VkVideoDecodeH264PictureLayoutFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH264PictureLayoutFlagBitsKHR.html)
  value specifying the picture layout used by the H.264 video sequence
  to be decoded.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkVideoDecodeH264ProfileInfoKHR-sType-sType"
  id="VUID-VkVideoDecodeH264ProfileInfoKHR-sType-sType"></a>
  VUID-VkVideoDecodeH264ProfileInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_VIDEO_DECODE_H264_PROFILE_INFO_KHR`

- <a href="#VUID-VkVideoDecodeH264ProfileInfoKHR-pictureLayout-parameter"
  id="VUID-VkVideoDecodeH264ProfileInfoKHR-pictureLayout-parameter"></a>
  VUID-VkVideoDecodeH264ProfileInfoKHR-pictureLayout-parameter  
  If `pictureLayout` is not `0`, `pictureLayout` **must** be a valid
  [VkVideoDecodeH264PictureLayoutFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH264PictureLayoutFlagBitsKHR.html)
  value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_decode_h264](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_decode_h264.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkVideoDecodeH264PictureLayoutFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH264PictureLayoutFlagBitsKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoDecodeH264ProfileInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
