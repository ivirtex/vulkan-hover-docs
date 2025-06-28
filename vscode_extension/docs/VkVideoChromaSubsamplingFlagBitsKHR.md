# VkVideoChromaSubsamplingFlagBitsKHR(3) Manual Page

## Name

VkVideoChromaSubsamplingFlagBitsKHR - Video format chroma subsampling bits



## [](#_c_specification)C Specification

The video format chroma subsampling is defined with the following enums:

```c++
// Provided by VK_KHR_video_queue
typedef enum VkVideoChromaSubsamplingFlagBitsKHR {
    VK_VIDEO_CHROMA_SUBSAMPLING_INVALID_KHR = 0,
    VK_VIDEO_CHROMA_SUBSAMPLING_MONOCHROME_BIT_KHR = 0x00000001,
    VK_VIDEO_CHROMA_SUBSAMPLING_420_BIT_KHR = 0x00000002,
    VK_VIDEO_CHROMA_SUBSAMPLING_422_BIT_KHR = 0x00000004,
    VK_VIDEO_CHROMA_SUBSAMPLING_444_BIT_KHR = 0x00000008,
} VkVideoChromaSubsamplingFlagBitsKHR;
```

## [](#_description)Description

- `VK_VIDEO_CHROMA_SUBSAMPLING_MONOCHROME_BIT_KHR` specifies that the format is monochrome.
- `VK_VIDEO_CHROMA_SUBSAMPLING_420_BIT_KHR` specified that the format is 4:2:0 chroma subsampled, i.e. the two chroma components are sampled horizontally and vertically at half the sample rate of the luma component.
- `VK_VIDEO_CHROMA_SUBSAMPLING_422_BIT_KHR` - the format is 4:2:2 chroma subsampled, i.e. the two chroma components are sampled horizontally at half the sample rate of luma component.
- `VK_VIDEO_CHROMA_SUBSAMPLING_444_BIT_KHR` - the format is 4:4:4 chroma sampled, i.e. all three components of the Y′CBCR format are sampled at the same rate, thus there is no chroma subsampling.

## [](#_see_also)See Also

[VK\_KHR\_video\_queue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_queue.html), [VkVideoChromaSubsamplingFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoChromaSubsamplingFlagsKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoChromaSubsamplingFlagBitsKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0