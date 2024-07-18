# VkVideoComponentBitDepthFlagBitsKHR(3) Manual Page

## Name

VkVideoComponentBitDepthFlagBitsKHR - Video format component bit depth



## <a href="#_c_specification" class="anchor"></a>C Specification

Possible values for the video format component bit depth are:

``` c
// Provided by VK_KHR_video_queue
typedef enum VkVideoComponentBitDepthFlagBitsKHR {
    VK_VIDEO_COMPONENT_BIT_DEPTH_INVALID_KHR = 0,
    VK_VIDEO_COMPONENT_BIT_DEPTH_8_BIT_KHR = 0x00000001,
    VK_VIDEO_COMPONENT_BIT_DEPTH_10_BIT_KHR = 0x00000004,
    VK_VIDEO_COMPONENT_BIT_DEPTH_12_BIT_KHR = 0x00000010,
} VkVideoComponentBitDepthFlagBitsKHR;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_VIDEO_COMPONENT_BIT_DEPTH_8_BIT_KHR` specifies a component bit
  depth of 8 bits.

- `VK_VIDEO_COMPONENT_BIT_DEPTH_10_BIT_KHR` specifies a component bit
  depth of 10 bits.

- `VK_VIDEO_COMPONENT_BIT_DEPTH_12_BIT_KHR` specifies a component bit
  depth of 12 bits.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_queue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_queue.html),
[VkVideoComponentBitDepthFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoComponentBitDepthFlagsKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoComponentBitDepthFlagBitsKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
