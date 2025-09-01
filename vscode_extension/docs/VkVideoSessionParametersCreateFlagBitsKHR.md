# VkVideoSessionParametersCreateFlagBitsKHR(3) Manual Page

## Name

VkVideoSessionParametersCreateFlagBitsKHR - Video session parameters creation flags



## [](#_c_specification)C Specification

Bits which **can** be set in [VkVideoSessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionParametersCreateInfoKHR.html)::`flags` are:

```c++
// Provided by VK_KHR_video_encode_quantization_map
typedef enum VkVideoSessionParametersCreateFlagBitsKHR {
  // Provided by VK_KHR_video_encode_quantization_map
    VK_VIDEO_SESSION_PARAMETERS_CREATE_QUANTIZATION_MAP_COMPATIBLE_BIT_KHR = 0x00000001,
} VkVideoSessionParametersCreateFlagBitsKHR;
```

## [](#_description)Description

- `VK_VIDEO_SESSION_PARAMETERS_CREATE_QUANTIZATION_MAP_COMPATIBLE_BIT_KHR` specifies that the created video session parameters object **can** be used with [quantization maps](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-quantization-map).

## [](#_see_also)See Also

[VK\_KHR\_video\_encode\_quantization\_map](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_quantization_map.html), [VkVideoSessionParametersCreateFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionParametersCreateFlagsKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoSessionParametersCreateFlagBitsKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0