# VkVideoEncodeRateControlLayerInfoKHR(3) Manual Page

## Name

VkVideoEncodeRateControlLayerInfoKHR - Structure to set encode per-layer rate control parameters



## [](#_c_specification)C Specification

The `VkVideoEncodeRateControlLayerInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_video_encode_queue
typedef struct VkVideoEncodeRateControlLayerInfoKHR {
    VkStructureType    sType;
    const void*        pNext;
    uint64_t           averageBitrate;
    uint64_t           maxBitrate;
    uint32_t           frameRateNumerator;
    uint32_t           frameRateDenominator;
} VkVideoEncodeRateControlLayerInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is a pointer to a structure extending this structure.
- `averageBitrate` is the average [bitrate](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-bitrate) to be targeted by the implementation’s rate control algorithm.
- `maxBitrate` is the peak [bitrate](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-bitrate) to be targeted by the implementation’s rate control algorithm.
- `frameRateNumerator` is the numerator of the frame rate assumed by the implementation’s rate control algorithm.
- `frameRateDenominator` is the denominator of the frame rate assumed by the implementation’s rate control algorithm.

## [](#_description)Description

Note

The ability of the implementation’s rate control algorithm to be able to match the requested average and/or peak bitrates **may** be limited by the set of other codec-independent and codec-specific rate control parameters specified by the application, the input content, as well as the application conforming to the rate control guidance provided to the implementation, as described [earlier](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-rate-control).

Additional structures providing codec-specific rate control parameters **can** be included in the `pNext` chain of `VkVideoEncodeRateControlLayerInfoKHR` depending on the [video profile](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#video-profiles) the bound video session was created with. For further details see:

- [Video Coding Control](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#video-coding-control)
- [H.264 Encode Rate Control](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h264-rate-control)
- [H.265 Encode Rate Control](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h265-rate-control)
- [AV1 Encode Rate Control](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-rate-control)

Valid Usage

- [](#VUID-VkVideoEncodeRateControlLayerInfoKHR-frameRateNumerator-08350)VUID-VkVideoEncodeRateControlLayerInfoKHR-frameRateNumerator-08350  
  `frameRateNumerator` **must** be greater than zero
- [](#VUID-VkVideoEncodeRateControlLayerInfoKHR-frameRateDenominator-08351)VUID-VkVideoEncodeRateControlLayerInfoKHR-frameRateDenominator-08351  
  `frameRateDenominator` **must** be greater than zero

Valid Usage (Implicit)

- [](#VUID-VkVideoEncodeRateControlLayerInfoKHR-sType-sType)VUID-VkVideoEncodeRateControlLayerInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_ENCODE_RATE_CONTROL_LAYER_INFO_KHR`
- [](#VUID-VkVideoEncodeRateControlLayerInfoKHR-pNext-pNext)VUID-VkVideoEncodeRateControlLayerInfoKHR-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the `pNext` chain **must** be either `NULL` or a pointer to a valid instance of [VkVideoEncodeAV1RateControlLayerInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1RateControlLayerInfoKHR.html), [VkVideoEncodeH264RateControlLayerInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264RateControlLayerInfoKHR.html), or [VkVideoEncodeH265RateControlLayerInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265RateControlLayerInfoKHR.html)
- [](#VUID-VkVideoEncodeRateControlLayerInfoKHR-sType-unique)VUID-VkVideoEncodeRateControlLayerInfoKHR-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique

## [](#_see_also)See Also

[VK\_KHR\_video\_encode\_queue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_queue.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkVideoEncodeRateControlInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeRateControlInfoKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoEncodeRateControlLayerInfoKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0