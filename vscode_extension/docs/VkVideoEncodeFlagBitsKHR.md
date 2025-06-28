# VkVideoEncodeFlagBitsKHR(3) Manual Page

## Name

VkVideoEncodeFlagBitsKHR - Video encode flags



## [](#_c_specification)C Specification

Bits which **can** be set in [VkVideoEncodeInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeInfoKHR.html)::`flags`, specifying video encode flags, are:

```c++
// Provided by VK_KHR_video_encode_quantization_map
typedef enum VkVideoEncodeFlagBitsKHR {
  // Provided by VK_KHR_video_encode_quantization_map
    VK_VIDEO_ENCODE_WITH_QUANTIZATION_DELTA_MAP_BIT_KHR = 0x00000001,
  // Provided by VK_KHR_video_encode_quantization_map
    VK_VIDEO_ENCODE_WITH_EMPHASIS_MAP_BIT_KHR = 0x00000002,
} VkVideoEncodeFlagBitsKHR;
```

## [](#_description)Description

- `VK_VIDEO_ENCODE_WITH_QUANTIZATION_DELTA_MAP_BIT_KHR` specifies the use of a [quantization delta map](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-quantization-delta-map) in the issued [video encode operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#video-encode-operations).
- `VK_VIDEO_ENCODE_WITH_EMPHASIS_MAP_BIT_KHR` specifies the use of an [emphasis map](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-emphasis-map) in the issued [video encode operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#video-encode-operations).

## [](#_see_also)See Also

[VK\_KHR\_video\_encode\_quantization\_map](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_quantization_map.html), [VkVideoEncodeFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeFlagsKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoEncodeFlagBitsKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0