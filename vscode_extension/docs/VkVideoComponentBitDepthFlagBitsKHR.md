# VkVideoComponentBitDepthFlagBitsKHR(3) Manual Page

## Name

VkVideoComponentBitDepthFlagBitsKHR - Video format component bit depth



## [](#_c_specification)C Specification

Possible values for the video format component bit depth are:

```c++
// Provided by VK_KHR_video_queue
typedef enum VkVideoComponentBitDepthFlagBitsKHR {
    VK_VIDEO_COMPONENT_BIT_DEPTH_INVALID_KHR = 0,
    VK_VIDEO_COMPONENT_BIT_DEPTH_8_BIT_KHR = 0x00000001,
    VK_VIDEO_COMPONENT_BIT_DEPTH_10_BIT_KHR = 0x00000004,
    VK_VIDEO_COMPONENT_BIT_DEPTH_12_BIT_KHR = 0x00000010,
} VkVideoComponentBitDepthFlagBitsKHR;
```

## [](#_description)Description

- `VK_VIDEO_COMPONENT_BIT_DEPTH_8_BIT_KHR` specifies a component bit depth of 8 bits.
- `VK_VIDEO_COMPONENT_BIT_DEPTH_10_BIT_KHR` specifies a component bit depth of 10 bits.
- `VK_VIDEO_COMPONENT_BIT_DEPTH_12_BIT_KHR` specifies a component bit depth of 12 bits.

## [](#_see_also)See Also

[VK\_KHR\_video\_queue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_queue.html), [VkVideoComponentBitDepthFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoComponentBitDepthFlagsKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoComponentBitDepthFlagBitsKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0