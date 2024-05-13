# VkVideoDecodeH265ProfileInfoKHR(3) Manual Page

## Name

VkVideoDecodeH265ProfileInfoKHR - Structure specifying H.265 decode
profile



## <a href="#_c_specification" class="anchor"></a>C Specification

A video profile supporting H.265 video decode operations is specified by
setting
[VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileInfoKHR.html)::`videoCodecOperation`
to `VK_VIDEO_CODEC_OPERATION_DECODE_H265_BIT_KHR` and adding a
`VkVideoDecodeH265ProfileInfoKHR` structure to the
[VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileInfoKHR.html)::`pNext` chain.

The `VkVideoDecodeH265ProfileInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_video_decode_h265
typedef struct VkVideoDecodeH265ProfileInfoKHR {
    VkStructureType           sType;
    const void*               pNext;
    StdVideoH265ProfileIdc    stdProfileIdc;
} VkVideoDecodeH265ProfileInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `stdProfileIdc` is a `StdVideoH265ProfileIdc` value specifying the
  H.265 codec profile IDC, as defined in section A.3 of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h265"
  target="_blank" rel="noopener">ITU-T H.265 Specification</a>.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkVideoDecodeH265ProfileInfoKHR-sType-sType"
  id="VUID-VkVideoDecodeH265ProfileInfoKHR-sType-sType"></a>
  VUID-VkVideoDecodeH265ProfileInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_VIDEO_DECODE_H265_PROFILE_INFO_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_decode_h265](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_decode_h265.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoDecodeH265ProfileInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
