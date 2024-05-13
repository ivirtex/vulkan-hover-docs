# VkVideoEncodeH264QualityLevelPropertiesKHR(3) Manual Page

## Name

VkVideoEncodeH264QualityLevelPropertiesKHR - Structure describing the
H.264 encode quality level properties



## <a href="#_c_specification" class="anchor"></a>C Specification

When calling
[vkGetPhysicalDeviceVideoEncodeQualityLevelPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoEncodeQualityLevelPropertiesKHR.html)
with `pVideoProfile->videoCodecOperation` specified as
`VK_VIDEO_CODEC_OPERATION_ENCODE_H264_BIT_KHR`, the
[VkVideoEncodeH264QualityLevelPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264QualityLevelPropertiesKHR.html)
structure **must** be included in the `pNext` chain of the
[VkVideoEncodeQualityLevelPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeQualityLevelPropertiesKHR.html)
structure to retrieve additional video encode quality level properties
specific to H.264 encoding.

The
[VkVideoEncodeH264QualityLevelPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264QualityLevelPropertiesKHR.html)
structure is defined as:

``` c
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

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `preferredRateControlFlags` is a bitmask of
  [VkVideoEncodeH264RateControlFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264RateControlFlagBitsKHR.html)
  values indicating the preferred flags to use for
  [VkVideoEncodeH264RateControlInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264RateControlInfoKHR.html)::`flags`.

- `preferredGopFrameCount` indicates the preferred value to use for
  [VkVideoEncodeH264RateControlInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264RateControlInfoKHR.html)::`gopFrameCount`.

- `preferredIdrPeriod` indicates the preferred value to use for
  [VkVideoEncodeH264RateControlInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264RateControlInfoKHR.html)::`idrPeriod`.

- `preferredConsecutiveBFrameCount` indicates the preferred value to use
  for
  [VkVideoEncodeH264RateControlInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264RateControlInfoKHR.html)::`consecutiveBFrameCount`.

- `preferredTemporalLayerCount` indicates the preferred value to use for
  [VkVideoEncodeH264RateControlInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264RateControlInfoKHR.html)::`temporalLayerCount`.

- `preferredConstantQp` indicates the preferred values to use for
  [VkVideoEncodeH264NaluSliceInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264NaluSliceInfoKHR.html)::`constantQp`
  for each picture type when using <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-rate-control-modes"
  target="_blank" rel="noopener">rate control mode</a>
  `VK_VIDEO_ENCODE_RATE_CONTROL_MODE_DISABLED_BIT_KHR`.

- `preferredMaxL0ReferenceCount` indicates the preferred maximum number
  of reference pictures to use in the reference list L0.

- `preferredMaxL1ReferenceCount` indicates the preferred maximum number
  of reference pictures to use in the reference list L1.

- `preferredStdEntropyCodingModeFlag` indicates the preferred value to
  use for `entropy_coding_mode_flag` in `StdVideoH264PpsFlags`.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkVideoEncodeH264QualityLevelPropertiesKHR-sType-sType"
  id="VUID-VkVideoEncodeH264QualityLevelPropertiesKHR-sType-sType"></a>
  VUID-VkVideoEncodeH264QualityLevelPropertiesKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_VIDEO_ENCODE_H264_QUALITY_LEVEL_PROPERTIES_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_encode_h264](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_encode_h264.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkVideoEncodeH264QpKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264QpKHR.html),
[VkVideoEncodeH264RateControlFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264RateControlFlagsKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoEncodeH264QualityLevelPropertiesKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
