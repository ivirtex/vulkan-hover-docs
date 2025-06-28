# VkVideoEncodeAV1QualityLevelPropertiesKHR(3) Manual Page

## Name

VkVideoEncodeAV1QualityLevelPropertiesKHR - Structure describing the AV1 encode quality level properties



## [](#_c_specification)C Specification

When calling [vkGetPhysicalDeviceVideoEncodeQualityLevelPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoEncodeQualityLevelPropertiesKHR.html) with `pVideoProfile->videoCodecOperation` specified as `VK_VIDEO_CODEC_OPERATION_ENCODE_AV1_BIT_KHR`, the [VkVideoEncodeAV1QualityLevelPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1QualityLevelPropertiesKHR.html) structure **must** be included in the `pNext` chain of the [VkVideoEncodeQualityLevelPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeQualityLevelPropertiesKHR.html) structure to retrieve additional video encode quality level properties specific to AV1 encoding.

The [VkVideoEncodeAV1QualityLevelPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1QualityLevelPropertiesKHR.html) structure is defined as:

```c++
// Provided by VK_KHR_video_encode_av1
typedef struct VkVideoEncodeAV1QualityLevelPropertiesKHR {
    VkStructureType                        sType;
    void*                                  pNext;
    VkVideoEncodeAV1RateControlFlagsKHR    preferredRateControlFlags;
    uint32_t                               preferredGopFrameCount;
    uint32_t                               preferredKeyFramePeriod;
    uint32_t                               preferredConsecutiveBipredictiveFrameCount;
    uint32_t                               preferredTemporalLayerCount;
    VkVideoEncodeAV1QIndexKHR              preferredConstantQIndex;
    uint32_t                               preferredMaxSingleReferenceCount;
    uint32_t                               preferredSingleReferenceNameMask;
    uint32_t                               preferredMaxUnidirectionalCompoundReferenceCount;
    uint32_t                               preferredMaxUnidirectionalCompoundGroup1ReferenceCount;
    uint32_t                               preferredUnidirectionalCompoundReferenceNameMask;
    uint32_t                               preferredMaxBidirectionalCompoundReferenceCount;
    uint32_t                               preferredMaxBidirectionalCompoundGroup1ReferenceCount;
    uint32_t                               preferredMaxBidirectionalCompoundGroup2ReferenceCount;
    uint32_t                               preferredBidirectionalCompoundReferenceNameMask;
} VkVideoEncodeAV1QualityLevelPropertiesKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `preferredRateControlFlags` is a bitmask of [VkVideoEncodeAV1RateControlFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1RateControlFlagBitsKHR.html) values indicating the preferred flags to use for [VkVideoEncodeAV1RateControlInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1RateControlInfoKHR.html)::`flags`.
- `preferredGopFrameCount` indicates the preferred value to use for [VkVideoEncodeAV1RateControlInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1RateControlInfoKHR.html)::`gopFrameCount`.
- `preferredKeyFramePeriod` indicates the preferred value to use for [VkVideoEncodeAV1RateControlInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1RateControlInfoKHR.html)::`keyFramePeriod`.
- `preferredConsecutiveBipredictiveFrameCount` indicates the preferred value to use for [VkVideoEncodeAV1RateControlInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1RateControlInfoKHR.html)::`consecutiveBipredictiveFrameCount`.
- `preferredTemporalLayerCount` indicates the preferred value to use for [VkVideoEncodeAV1RateControlInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1RateControlInfoKHR.html)::`temporalLayerCount`.
- `preferredConstantQIndex` indicates the preferred value to use for [VkVideoEncodeAV1PictureInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1PictureInfoKHR.html)::`constantQIndex` when using [rate control mode](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-rate-control-modes) `VK_VIDEO_ENCODE_RATE_CONTROL_MODE_DISABLED_BIT_KHR`.
- `preferredMaxSingleReferenceCount` indicates the preferred maximum number of reference pictures to use with [single reference prediction mode](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-prediction-modes).
- `preferredSingleReferenceNameMask` is a bitmask of preferred [AV1 reference names](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-reference-names) when using [single reference prediction mode](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-prediction-modes).
- `preferredMaxUnidirectionalCompoundReferenceCount` indicates the preferred maximum number of reference pictures to use with [unidirectional compound prediction mode](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-prediction-modes).
- `preferredMaxUnidirectionalCompoundGroup1ReferenceCount` indicates the preferred maximum number of reference pictures to use with [unidirectional compound prediction mode](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-prediction-modes) from reference frame group 1, as defined in section 6.10.24 of the [AV1 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#aomedia-av1).
- `preferredUnidirectionalCompoundReferenceNameMask` is a bitmask of preferred [AV1 reference names](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-reference-names) when using [unidirectional compound prediction mode](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-prediction-modes).
- `preferredMaxBidirectionalCompoundReferenceCount` indicates the preferred maximum number of reference pictures to use with [bidirectional compound prediction mode](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-prediction-modes).
- `preferredMaxBidirectionalCompoundGroup1ReferenceCount` indicates the preferred maximum number of reference pictures to use with [bidirectional compound prediction mode](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-prediction-modes) from reference frame group 1, as defined in section 6.10.24 of the [AV1 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#aomedia-av1).
- `preferredMaxBidirectionalCompoundGroup2ReferenceCount` indicates the preferred maximum number of reference pictures to use with [bidirectional compound prediction mode](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-prediction-modes) from reference frame group 2, as defined in section 6.10.24 of the [AV1 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#aomedia-av1).
- `preferredBidirectionalCompoundReferenceNameMask` is a bitmask of preferred [AV1 reference names](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-reference-names) when using [bidirectional compound prediction mode](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-prediction-modes).

## [](#_description)Description

`preferredSingleReferenceNameMask`, `preferredUnidirectionalCompoundReferenceNameMask`, and `preferredBidirectionalCompoundReferenceNameMask` are encoded such that when bit index i is set, it indicates preference for using the [AV1 reference name](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-reference-names) `STD_VIDEO_AV1_REFERENCE_NAME_LAST_FRAME` + i.

Valid Usage (Implicit)

- [](#VUID-VkVideoEncodeAV1QualityLevelPropertiesKHR-sType-sType)VUID-VkVideoEncodeAV1QualityLevelPropertiesKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_ENCODE_AV1_QUALITY_LEVEL_PROPERTIES_KHR`

## [](#_see_also)See Also

[VK\_KHR\_video\_encode\_av1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_av1.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkVideoEncodeAV1QIndexKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1QIndexKHR.html), [VkVideoEncodeAV1RateControlFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1RateControlFlagsKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoEncodeAV1QualityLevelPropertiesKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0