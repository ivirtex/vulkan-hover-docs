# VkVideoEncodeAV1QIndexKHR(3) Manual Page

## Name

VkVideoEncodeAV1QIndexKHR - Structure describing AV1 quantizer index values per prediction mode



## [](#_c_specification)C Specification

The `VkVideoEncodeAV1QIndexKHR` structure is defined as:

```c++
// Provided by VK_KHR_video_encode_av1
typedef struct VkVideoEncodeAV1QIndexKHR {
    uint32_t    intraQIndex;
    uint32_t    predictiveQIndex;
    uint32_t    bipredictiveQIndex;
} VkVideoEncodeAV1QIndexKHR;
```

## [](#_members)Members

- `intraQIndex` is the quantizer index to be used for frames encoded with `VK_VIDEO_ENCODE_AV1_RATE_CONTROL_GROUP_INTRA_KHR`.
- `predictiveQIndex` is the quantizer index to be used for frames encoded with `VK_VIDEO_ENCODE_AV1_RATE_CONTROL_GROUP_PREDICTIVE_KHR`.
- `bipredictiveQIndex` is the quantizer index to be used for frames encoded with `VK_VIDEO_ENCODE_AV1_RATE_CONTROL_GROUP_BIPREDICTIVE_KHR`.

## [](#_description)Description

## [](#_see_also)See Also

[VK\_KHR\_video\_encode\_av1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_av1.html), [VkVideoEncodeAV1QualityLevelPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1QualityLevelPropertiesKHR.html), [VkVideoEncodeAV1RateControlLayerInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1RateControlLayerInfoKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoEncodeAV1QIndexKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0