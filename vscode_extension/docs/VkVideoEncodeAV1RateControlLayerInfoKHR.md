# VkVideoEncodeAV1RateControlLayerInfoKHR(3) Manual Page

## Name

VkVideoEncodeAV1RateControlLayerInfoKHR - Structure describing AV1 per-layer rate control parameters



## [](#_c_specification)C Specification

The `VkVideoEncodeAV1RateControlLayerInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_video_encode_av1
typedef struct VkVideoEncodeAV1RateControlLayerInfoKHR {
    VkStructureType                 sType;
    const void*                     pNext;
    VkBool32                        useMinQIndex;
    VkVideoEncodeAV1QIndexKHR       minQIndex;
    VkBool32                        useMaxQIndex;
    VkVideoEncodeAV1QIndexKHR       maxQIndex;
    VkBool32                        useMaxFrameSize;
    VkVideoEncodeAV1FrameSizeKHR    maxFrameSize;
} VkVideoEncodeAV1RateControlLayerInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `useMinQIndex` indicates whether the quantizer index values determined by rate control will be clamped to the lower bounds on the quantizer index values specified in `minQIndex`.
- `minQIndex` specifies the lower bounds on the quantizer index values, for each [rate control group](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-rate-control-group), that the implementation’s rate control algorithm will use when `useMinQIndex` is set to `VK_TRUE`.
- `useMaxQIndex` indicates whether the quantizer index values determined by rate control will be clamped to the upper bounds on the quantizer index values specified in `maxQIndex`.
- `maxQIndex` specifies the upper bounds on the quantizer index values, for each [rate control group](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-rate-control-group), that the implementation’s rate control algorithm will use when `useMaxQIndex` is set to `VK_TRUE`.
- `useMaxFrameSize` indicates whether the implementation’s rate control algorithm **should** use the values specified in `maxFrameSize` as the upper bounds on the encoded frame size for each [rate control group](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-rate-control-group).
- `maxFrameSize` specifies the upper bounds on the encoded frame size, for each [rate control group](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-rate-control-group), when `useMaxFrameSize` is set to `VK_TRUE`.

## [](#_description)Description

When used, the values in `minQIndex` and `maxQIndex` guarantee that the effective quantizer index values used by the implementation will respect those lower and upper bounds, respectively. However, limiting the range of quantizer index values that the implementation is able to use will also limit the capabilities of the implementation’s rate control algorithm to comply to other constraints. In particular, the implementation **may** not be able to comply to the following:

- The average and/or peak [bitrate](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-bitrate) values to be used for the encoded bitstream specified in the `averageBitrate` and `maxBitrate` members of the [VkVideoEncodeRateControlLayerInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeRateControlLayerInfoKHR.html) structure.
- The upper bounds on the encoded frame size, for each [rate control group](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-rate-control-group), specified in the `maxFrameSize` member of `VkVideoEncodeAV1RateControlLayerInfoKHR`.

Note

In general, applications need to configure rate control parameters appropriately in order to be able to get the desired rate control behavior, as described in the [Video Encode Rate Control](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-rate-control) section.

When an instance of this structure is included in the `pNext` chain of a [VkVideoEncodeRateControlLayerInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeRateControlLayerInfoKHR.html) structure specified in one of the elements of the `pLayers` array member of the [VkVideoEncodeRateControlInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeRateControlInfoKHR.html) structure passed to the [vkCmdControlVideoCodingKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdControlVideoCodingKHR.html) command, [VkVideoCodingControlInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoCodingControlInfoKHR.html)::`flags` includes `VK_VIDEO_CODING_CONTROL_ENCODE_RATE_CONTROL_BIT_KHR`, and the bound video session was created with the video codec operation `VK_VIDEO_CODEC_OPERATION_ENCODE_AV1_BIT_KHR`, it specifies the AV1-specific rate control parameters of the rate control layer corresponding to that element of `pLayers`.

Valid Usage

- [](#VUID-VkVideoEncodeAV1RateControlLayerInfoKHR-useMinQIndex-10300)VUID-VkVideoEncodeAV1RateControlLayerInfoKHR-useMinQIndex-10300  
  If `useMinQIndex` is `VK_TRUE`, then the `intraQIndex`, `predictiveQIndex`, and `bipredictiveQIndex` members of `minQIndex` **must** all be between [VkVideoEncodeAV1CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1CapabilitiesKHR.html)::`minQIndex` and [VkVideoEncodeAV1CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1CapabilitiesKHR.html)::`maxQIndex`, as returned by [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html) for the used video profile
- [](#VUID-VkVideoEncodeAV1RateControlLayerInfoKHR-useMinQIndex-10301)VUID-VkVideoEncodeAV1RateControlLayerInfoKHR-useMinQIndex-10301  
  If `useMinQIndex` is `VK_TRUE` and [VkVideoEncodeAV1CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1CapabilitiesKHR.html)::`flags`, as returned by [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html) for the used video profile, does not include `VK_VIDEO_ENCODE_AV1_CAPABILITY_PER_RATE_CONTROL_GROUP_MIN_MAX_Q_INDEX_BIT_KHR`, then the `intraQIndex`, `predictiveQIndex`, and `bipredictiveQIndex` members of `minQIndex` **must** all specify the same value
- [](#VUID-VkVideoEncodeAV1RateControlLayerInfoKHR-useMaxQIndex-10302)VUID-VkVideoEncodeAV1RateControlLayerInfoKHR-useMaxQIndex-10302  
  If `useMaxQIndex` is `VK_TRUE`, then the `intraQIndex`, `predictiveQIndex`, and `bipredictiveQIndex` members of `maxQIndex` **must** all be between [VkVideoEncodeAV1CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1CapabilitiesKHR.html)::`minQIndex` and [VkVideoEncodeAV1CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1CapabilitiesKHR.html)::`maxQIndex`, as returned by [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html) for the used video profile
- [](#VUID-VkVideoEncodeAV1RateControlLayerInfoKHR-useMaxQIndex-10303)VUID-VkVideoEncodeAV1RateControlLayerInfoKHR-useMaxQIndex-10303  
  If `useMaxQIndex` is `VK_TRUE` and [VkVideoEncodeAV1CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1CapabilitiesKHR.html)::`flags`, as returned by [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html) for the used video profile, does not include `VK_VIDEO_ENCODE_AV1_CAPABILITY_PER_RATE_CONTROL_GROUP_MIN_MAX_Q_INDEX_BIT_KHR`, then the `intraQIndex`, `predictiveQIndex`, and `bipredictiveQIndex` members of `maxQIndex` **must** all specify the same value
- [](#VUID-VkVideoEncodeAV1RateControlLayerInfoKHR-useMinQIndex-10304)VUID-VkVideoEncodeAV1RateControlLayerInfoKHR-useMinQIndex-10304  
  If `useMinQIndex` and `useMaxQIndex` are both `VK_TRUE`, then the `intraQIndex`, `predictiveQIndex`, and `bipredictiveQIndex` members of `minQIndex` **must** all be less than or equal to the respective members of `maxQIndex`

Valid Usage (Implicit)

- [](#VUID-VkVideoEncodeAV1RateControlLayerInfoKHR-sType-sType)VUID-VkVideoEncodeAV1RateControlLayerInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_ENCODE_AV1_RATE_CONTROL_LAYER_INFO_KHR`
- [](#VUID-VkVideoEncodeAV1RateControlLayerInfoKHR-minQIndex-parameter)VUID-VkVideoEncodeAV1RateControlLayerInfoKHR-minQIndex-parameter  
  `minQIndex` **must** be a valid [VkVideoEncodeAV1QIndexKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1QIndexKHR.html) structure
- [](#VUID-VkVideoEncodeAV1RateControlLayerInfoKHR-maxQIndex-parameter)VUID-VkVideoEncodeAV1RateControlLayerInfoKHR-maxQIndex-parameter  
  `maxQIndex` **must** be a valid [VkVideoEncodeAV1QIndexKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1QIndexKHR.html) structure
- [](#VUID-VkVideoEncodeAV1RateControlLayerInfoKHR-maxFrameSize-parameter)VUID-VkVideoEncodeAV1RateControlLayerInfoKHR-maxFrameSize-parameter  
  `maxFrameSize` **must** be a valid [VkVideoEncodeAV1FrameSizeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1FrameSizeKHR.html) structure

## [](#_see_also)See Also

[VK\_KHR\_video\_encode\_av1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_av1.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkVideoEncodeAV1FrameSizeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1FrameSizeKHR.html), [VkVideoEncodeAV1QIndexKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1QIndexKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoEncodeAV1RateControlLayerInfoKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0