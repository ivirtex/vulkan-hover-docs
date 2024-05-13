# VkVideoEncodeH265CtbSizeFlagBitsKHR(3) Manual Page

## Name

VkVideoEncodeH265CtbSizeFlagBitsKHR - Supported CTB sizes for H.265
video encode



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **may** be set in
[VkVideoEncodeH265CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265CapabilitiesKHR.html)::`ctbSizes`,
indicating the CTB sizes supported by the implementation, are:

``` c
// Provided by VK_KHR_video_encode_h265
typedef enum VkVideoEncodeH265CtbSizeFlagBitsKHR {
    VK_VIDEO_ENCODE_H265_CTB_SIZE_16_BIT_KHR = 0x00000001,
    VK_VIDEO_ENCODE_H265_CTB_SIZE_32_BIT_KHR = 0x00000002,
    VK_VIDEO_ENCODE_H265_CTB_SIZE_64_BIT_KHR = 0x00000004,
} VkVideoEncodeH265CtbSizeFlagBitsKHR;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_VIDEO_ENCODE_H265_CTB_SIZE_16_BIT_KHR` specifies that a CTB size
  of 16x16 is supported.

- `VK_VIDEO_ENCODE_H265_CTB_SIZE_32_BIT_KHR` specifies that a CTB size
  of 32x32 is supported.

- `VK_VIDEO_ENCODE_H265_CTB_SIZE_64_BIT_KHR` specifies that a CTB size
  of 64x64 is supported.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_encode_h265](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_encode_h265.html),
[VkVideoEncodeH265CtbSizeFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265CtbSizeFlagsKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoEncodeH265CtbSizeFlagBitsKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
