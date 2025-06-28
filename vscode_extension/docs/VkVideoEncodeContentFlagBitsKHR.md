# VkVideoEncodeContentFlagBitsKHR(3) Manual Page

## Name

VkVideoEncodeContentFlagBitsKHR - Video encode content flags



## [](#_c_specification)C Specification

The following bits **can** be specified in [VkVideoEncodeUsageInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeUsageInfoKHR.html)::`videoContentHints` as a hint about the encoded video content:

```c++
// Provided by VK_KHR_video_encode_queue
typedef enum VkVideoEncodeContentFlagBitsKHR {
    VK_VIDEO_ENCODE_CONTENT_DEFAULT_KHR = 0,
    VK_VIDEO_ENCODE_CONTENT_CAMERA_BIT_KHR = 0x00000001,
    VK_VIDEO_ENCODE_CONTENT_DESKTOP_BIT_KHR = 0x00000002,
    VK_VIDEO_ENCODE_CONTENT_RENDERED_BIT_KHR = 0x00000004,
} VkVideoEncodeContentFlagBitsKHR;
```

## [](#_description)Description

- `VK_VIDEO_ENCODE_CONTENT_CAMERA_BIT_KHR` specifies that video encoding is intended to be used to encode camera content.
- `VK_VIDEO_ENCODE_CONTENT_DESKTOP_BIT_KHR` specifies that video encoding is intended to be used to encode desktop content.
- `VK_VIDEO_ENCODE_CONTENT_RENDERED_BIT_KHR` specified that video encoding is intended to be used to encode rendered (e.g. game) content.

Note

There are no restrictions on the combination of bits that **can** be specified by the application. However, applications **should** use reasonable combinations in order for the implementation to be able to select the most appropriate mode of operation for the particular content type.

## [](#_see_also)See Also

[VK\_KHR\_video\_encode\_queue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_queue.html), [VkVideoEncodeContentFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeContentFlagsKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoEncodeContentFlagBitsKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0