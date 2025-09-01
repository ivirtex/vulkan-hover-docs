# VkVideoEncodeH265FrameSizeKHR(3) Manual Page

## Name

VkVideoEncodeH265FrameSizeKHR - Structure describing frame size values per H.265 picture type



## [](#_c_specification)C Specification

The `VkVideoEncodeH265FrameSizeKHR` structure is defined as:

```c++
// Provided by VK_KHR_video_encode_h265
typedef struct VkVideoEncodeH265FrameSizeKHR {
    uint32_t    frameISize;
    uint32_t    framePSize;
    uint32_t    frameBSize;
} VkVideoEncodeH265FrameSizeKHR;
```

## [](#_members)Members

- `frameISize` is the size in bytes to be used for [I frames](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h265-i-pic).
- `framePSize` is the size in bytes to be used for [P frames](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h265-p-pic).
- `frameBSize` is the size in bytes to be used for [B frames](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h265-b-pic).

## [](#_description)Description

## [](#_see_also)See Also

[VK\_KHR\_video\_encode\_h265](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_h265.html), [VkVideoEncodeH265RateControlLayerInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265RateControlLayerInfoKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoEncodeH265FrameSizeKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0