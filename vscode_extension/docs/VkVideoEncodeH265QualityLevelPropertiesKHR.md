# VkVideoEncodeH265QualityLevelPropertiesKHR(3) Manual Page

## Name

VkVideoEncodeH265QualityLevelPropertiesKHR - Structure describing the H.265 encode quality level properties



## [](#_c_specification)C Specification

When calling [vkGetPhysicalDeviceVideoEncodeQualityLevelPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoEncodeQualityLevelPropertiesKHR.html) with `pVideoProfile->videoCodecOperation` specified as `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR`, the [VkVideoEncodeH265QualityLevelPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265QualityLevelPropertiesKHR.html) structure **must** be included in the `pNext` chain of the [VkVideoEncodeQualityLevelPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeQualityLevelPropertiesKHR.html) structure to retrieve additional video encode quality level properties specific to H.265 encoding.

The [VkVideoEncodeH265QualityLevelPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265QualityLevelPropertiesKHR.html) structure is defined as:

```c++
// Provided by VK_KHR_video_encode_h265
typedef struct VkVideoEncodeH265QualityLevelPropertiesKHR {
    VkStructureType                         sType;
    void*                                   pNext;
    VkVideoEncodeH265RateControlFlagsKHR    preferredRateControlFlags;
    uint32_t                                preferredGopFrameCount;
    uint32_t                                preferredIdrPeriod;
    uint32_t                                preferredConsecutiveBFrameCount;
    uint32_t                                preferredSubLayerCount;
    VkVideoEncodeH265QpKHR                  preferredConstantQp;
    uint32_t                                preferredMaxL0ReferenceCount;
    uint32_t                                preferredMaxL1ReferenceCount;
} VkVideoEncodeH265QualityLevelPropertiesKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `preferredRateControlFlags` is a bitmask of [VkVideoEncodeH265RateControlFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265RateControlFlagBitsKHR.html) values indicating the preferred flags to use for [VkVideoEncodeH265RateControlInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265RateControlInfoKHR.html)::`flags`.
- `preferredGopFrameCount` indicates the preferred value to use for [VkVideoEncodeH265RateControlInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265RateControlInfoKHR.html)::`gopFrameCount`.
- `preferredIdrPeriod` indicates the preferred value to use for [VkVideoEncodeH265RateControlInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265RateControlInfoKHR.html)::`idrPeriod`.
- `preferredConsecutiveBFrameCount` indicates the preferred value to use for [VkVideoEncodeH265RateControlInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265RateControlInfoKHR.html)::`consecutiveBFrameCount`.
- `preferredSubLayerCount` indicates the preferred value to use for [VkVideoEncodeH265RateControlInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265RateControlInfoKHR.html)::`subLayerCount`.
- `preferredConstantQp` indicates the preferred values to use for [VkVideoEncodeH265NaluSliceSegmentInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265NaluSliceSegmentInfoKHR.html)::`constantQp` for each picture type when using [rate control mode](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-rate-control-modes) `VK_VIDEO_ENCODE_RATE_CONTROL_MODE_DISABLED_BIT_KHR`.
- `preferredMaxL0ReferenceCount` indicates the preferred maximum number of reference pictures to use in the reference list L0.
- `preferredMaxL1ReferenceCount` indicates the preferred maximum number of reference pictures to use in the reference list L1.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkVideoEncodeH265QualityLevelPropertiesKHR-sType-sType)VUID-VkVideoEncodeH265QualityLevelPropertiesKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_ENCODE_H265_QUALITY_LEVEL_PROPERTIES_KHR`

## [](#_see_also)See Also

[VK\_KHR\_video\_encode\_h265](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_h265.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkVideoEncodeH265QpKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265QpKHR.html), [VkVideoEncodeH265RateControlFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265RateControlFlagsKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoEncodeH265QualityLevelPropertiesKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0