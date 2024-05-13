# VkVideoEncodeH265TransformBlockSizeFlagBitsKHR(3) Manual Page

## Name

VkVideoEncodeH265TransformBlockSizeFlagBitsKHR - Supported transform
block sizes for H.265 video encode



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **may** be set in
[VkVideoEncodeH265CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265CapabilitiesKHR.html)::`transformBlockSizes`,
indicating the transform block sizes supported by the implementation,
are:

``` c
// Provided by VK_KHR_video_encode_h265
typedef enum VkVideoEncodeH265TransformBlockSizeFlagBitsKHR {
    VK_VIDEO_ENCODE_H265_TRANSFORM_BLOCK_SIZE_4_BIT_KHR = 0x00000001,
    VK_VIDEO_ENCODE_H265_TRANSFORM_BLOCK_SIZE_8_BIT_KHR = 0x00000002,
    VK_VIDEO_ENCODE_H265_TRANSFORM_BLOCK_SIZE_16_BIT_KHR = 0x00000004,
    VK_VIDEO_ENCODE_H265_TRANSFORM_BLOCK_SIZE_32_BIT_KHR = 0x00000008,
} VkVideoEncodeH265TransformBlockSizeFlagBitsKHR;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_VIDEO_ENCODE_H265_TRANSFORM_BLOCK_SIZE_4_BIT_KHR` specifies that a
  transform block size of 4x4 is supported.

- `VK_VIDEO_ENCODE_H265_TRANSFORM_BLOCK_SIZE_8_BIT_KHR` specifies that a
  transform block size of 8x8 is supported.

- `VK_VIDEO_ENCODE_H265_TRANSFORM_BLOCK_SIZE_16_BIT_KHR` specifies that
  a transform block size of 16x16 is supported.

- `VK_VIDEO_ENCODE_H265_TRANSFORM_BLOCK_SIZE_32_BIT_KHR` specifies that
  a transform block size of 32x32 is supported.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_encode_h265](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_encode_h265.html),
[VkVideoEncodeH265TransformBlockSizeFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265TransformBlockSizeFlagsKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoEncodeH265TransformBlockSizeFlagBitsKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
