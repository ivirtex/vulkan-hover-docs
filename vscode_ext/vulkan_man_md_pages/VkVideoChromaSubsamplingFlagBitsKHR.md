# VkVideoChromaSubsamplingFlagBitsKHR(3) Manual Page

## Name

VkVideoChromaSubsamplingFlagBitsKHR - Video format chroma subsampling
bits



## <a href="#_c_specification" class="anchor"></a>C Specification

The video format chroma subsampling is defined with the following enums:

``` c
// Provided by VK_KHR_video_queue
typedef enum VkVideoChromaSubsamplingFlagBitsKHR {
    VK_VIDEO_CHROMA_SUBSAMPLING_INVALID_KHR = 0,
    VK_VIDEO_CHROMA_SUBSAMPLING_MONOCHROME_BIT_KHR = 0x00000001,
    VK_VIDEO_CHROMA_SUBSAMPLING_420_BIT_KHR = 0x00000002,
    VK_VIDEO_CHROMA_SUBSAMPLING_422_BIT_KHR = 0x00000004,
    VK_VIDEO_CHROMA_SUBSAMPLING_444_BIT_KHR = 0x00000008,
} VkVideoChromaSubsamplingFlagBitsKHR;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_VIDEO_CHROMA_SUBSAMPLING_MONOCHROME_BIT_KHR` specifies that the
  format is monochrome.

- `VK_VIDEO_CHROMA_SUBSAMPLING_420_BIT_KHR` specified that the format is
  4:2:0 chroma subsampled, i.e. the two chroma components are sampled
  horizontally and vertically at half the sample rate of the luma
  component.

- `VK_VIDEO_CHROMA_SUBSAMPLING_422_BIT_KHR` - the format is 4:2:2 chroma
  subsampled, i.e. the two chroma components are sampled horizontally at
  half the sample rate of luma component.

- `VK_VIDEO_CHROMA_SUBSAMPLING_444_BIT_KHR` - the format is 4:4:4 chroma
  sampled, i.e. all three components of the Y′C<sub>B</sub>C<sub>R</sub>
  format are sampled at the same rate, thus there is no chroma
  subsampling.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_queue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_queue.html),
[VkVideoChromaSubsamplingFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoChromaSubsamplingFlagsKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoChromaSubsamplingFlagBitsKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
