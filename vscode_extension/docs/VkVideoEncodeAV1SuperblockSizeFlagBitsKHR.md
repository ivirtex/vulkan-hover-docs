# VkVideoEncodeAV1SuperblockSizeFlagBitsKHR(3) Manual Page

## Name

VkVideoEncodeAV1SuperblockSizeFlagBitsKHR - Supported superblock sizes for AV1 video encode



## [](#_c_specification)C Specification

Bits which **may** be set in [VkVideoEncodeAV1CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1CapabilitiesKHR.html)::`superblockSizes`, indicating the superblock sizes supported by the implementation, are:

```c++
// Provided by VK_KHR_video_encode_av1
typedef enum VkVideoEncodeAV1SuperblockSizeFlagBitsKHR {
    VK_VIDEO_ENCODE_AV1_SUPERBLOCK_SIZE_64_BIT_KHR = 0x00000001,
    VK_VIDEO_ENCODE_AV1_SUPERBLOCK_SIZE_128_BIT_KHR = 0x00000002,
} VkVideoEncodeAV1SuperblockSizeFlagBitsKHR;
```

## [](#_description)Description

- `VK_VIDEO_ENCODE_AV1_SUPERBLOCK_SIZE_64_BIT_KHR` specifies that a superblock size of 64x64 is supported.
- `VK_VIDEO_ENCODE_AV1_SUPERBLOCK_SIZE_128_BIT_KHR` specifies that a superblock size of 128x128 is supported.

## [](#_see_also)See Also

[VK\_KHR\_video\_encode\_av1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_av1.html), [VkVideoEncodeAV1SuperblockSizeFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1SuperblockSizeFlagsKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoEncodeAV1SuperblockSizeFlagBitsKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0