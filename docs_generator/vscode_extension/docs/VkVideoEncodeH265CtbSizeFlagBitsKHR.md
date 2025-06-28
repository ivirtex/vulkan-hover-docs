# VkVideoEncodeH265CtbSizeFlagBitsKHR(3) Manual Page

## Name

VkVideoEncodeH265CtbSizeFlagBitsKHR - Supported CTB sizes for H.265 video encode



## [](#_c_specification)C Specification

Bits which **may** be set in [VkVideoEncodeH265CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265CapabilitiesKHR.html)::`ctbSizes`, indicating the CTB sizes supported by the implementation, are:

```c++
// Provided by VK_KHR_video_encode_h265
typedef enum VkVideoEncodeH265CtbSizeFlagBitsKHR {
    VK_VIDEO_ENCODE_H265_CTB_SIZE_16_BIT_KHR = 0x00000001,
    VK_VIDEO_ENCODE_H265_CTB_SIZE_32_BIT_KHR = 0x00000002,
    VK_VIDEO_ENCODE_H265_CTB_SIZE_64_BIT_KHR = 0x00000004,
} VkVideoEncodeH265CtbSizeFlagBitsKHR;
```

## [](#_description)Description

- `VK_VIDEO_ENCODE_H265_CTB_SIZE_16_BIT_KHR` specifies that a CTB size of 16x16 is supported.
- `VK_VIDEO_ENCODE_H265_CTB_SIZE_32_BIT_KHR` specifies that a CTB size of 32x32 is supported.
- `VK_VIDEO_ENCODE_H265_CTB_SIZE_64_BIT_KHR` specifies that a CTB size of 64x64 is supported.

## [](#_see_also)See Also

[VK\_KHR\_video\_encode\_h265](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_h265.html), [VkVideoEncodeH265CtbSizeFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265CtbSizeFlagsKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoEncodeH265CtbSizeFlagBitsKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0