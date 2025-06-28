# VkVideoEncodeH264QualityLevelPropertiesKHR(3) Manual Page

## Name

VkVideoEncodeH264QualityLevelPropertiesKHR - Structure describing the H.264 encode quality level properties



## [](#_c_specification)C Specification

When calling [vkGetPhysicalDeviceVideoEncodeQualityLevelPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoEncodeQualityLevelPropertiesKHR.html) with `pVideoProfile->videoCodecOperation` specified as `VK_VIDEO_CODEC_OPERATION_ENCODE_H264_BIT_KHR`, the [VkVideoEncodeH264QualityLevelPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264QualityLevelPropertiesKHR.html) structure **must** be included in the `pNext` chain of the [VkVideoEncodeQualityLevelPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeQualityLevelPropertiesKHR.html) structure to retrieve additional video encode quality level properties specific to H.264 encoding.

The [VkVideoEncodeH264QualityLevelPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264QualityLevelPropertiesKHR.html) structure is defined as:

```c++
// Provided by VK_KHR_video_encode_h264
typedef struct VkVideoEncodeH264QualityLevelPropertiesKHR {
    VkStructureType                         sType;
    void*                                   pNext;
    VkVideoEncodeH264RateControlFlagsKHR    preferredRateControlFlags;
    uint32_t                                preferredGopFrameCount;
    uint32_t                                preferredIdrPeriod;
    uint32_t                                preferredConsecutiveBFrameCount;
    uint32_t                                preferredTemporalLayerCount;
    VkVideoEncodeH264QpKHR                  preferredConstantQp;
    uint32_t                                preferredMaxL0ReferenceCount;
    uint32_t                                preferredMaxL1ReferenceCount;
    VkBool32                                preferredStdEntropyCodingModeFlag;
} VkVideoEncodeH264QualityLevelPropertiesKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `preferredRateControlFlags` is a bitmask of [VkVideoEncodeH264RateControlFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264RateControlFlagBitsKHR.html) values indicating the preferred flags to use for [VkVideoEncodeH264RateControlInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264RateControlInfoKHR.html)::`flags`.
- `preferredGopFrameCount` indicates the preferred value to use for [VkVideoEncodeH264RateControlInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264RateControlInfoKHR.html)::`gopFrameCount`.
- `preferredIdrPeriod` indicates the preferred value to use for [VkVideoEncodeH264RateControlInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264RateControlInfoKHR.html)::`idrPeriod`.
- `preferredConsecutiveBFrameCount` indicates the preferred value to use for [VkVideoEncodeH264RateControlInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264RateControlInfoKHR.html)::`consecutiveBFrameCount`.
- `preferredTemporalLayerCount` indicates the preferred value to use for [VkVideoEncodeH264RateControlInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264RateControlInfoKHR.html)::`temporalLayerCount`.
- `preferredConstantQp` indicates the preferred values to use for [VkVideoEncodeH264NaluSliceInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264NaluSliceInfoKHR.html)::`constantQp` for each picture type when using [rate control mode](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-rate-control-modes) `VK_VIDEO_ENCODE_RATE_CONTROL_MODE_DISABLED_BIT_KHR`.
- `preferredMaxL0ReferenceCount` indicates the preferred maximum number of reference pictures to use in the reference list L0.
- `preferredMaxL1ReferenceCount` indicates the preferred maximum number of reference pictures to use in the reference list L1.
- `preferredStdEntropyCodingModeFlag` indicates the preferred value to use for `entropy_coding_mode_flag` in `StdVideoH264PpsFlags`.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkVideoEncodeH264QualityLevelPropertiesKHR-sType-sType)VUID-VkVideoEncodeH264QualityLevelPropertiesKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_ENCODE_H264_QUALITY_LEVEL_PROPERTIES_KHR`

## [](#_see_also)See Also

[VK\_KHR\_video\_encode\_h264](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_h264.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkVideoEncodeH264QpKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264QpKHR.html), [VkVideoEncodeH264RateControlFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264RateControlFlagsKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoEncodeH264QualityLevelPropertiesKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0