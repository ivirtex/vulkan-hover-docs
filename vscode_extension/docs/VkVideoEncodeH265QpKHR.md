# VkVideoEncodeH265QpKHR(3) Manual Page

## Name

VkVideoEncodeH265QpKHR - Structure describing H.265 QP values per picture type



## [](#_c_specification)C Specification

The `VkVideoEncodeH265QpKHR` structure is defined as:

```c++
// Provided by VK_KHR_video_encode_h265
typedef struct VkVideoEncodeH265QpKHR {
    int32_t    qpI;
    int32_t    qpP;
    int32_t    qpB;
} VkVideoEncodeH265QpKHR;
```

## [](#_members)Members

- `qpI` is the QP to be used for [I pictures](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h265-i-pic).
- `qpP` is the QP to be used for [P pictures](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h265-p-pic).
- `qpB` is the QP to be used for [B pictures](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h265-b-pic).

## [](#_description)Description

## [](#_see_also)See Also

[VK\_KHR\_video\_encode\_h265](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_h265.html), [VkVideoEncodeH265QualityLevelPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265QualityLevelPropertiesKHR.html), [VkVideoEncodeH265RateControlLayerInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265RateControlLayerInfoKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoEncodeH265QpKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0