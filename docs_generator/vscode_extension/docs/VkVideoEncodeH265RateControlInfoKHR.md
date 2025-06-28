# VkVideoEncodeH265RateControlInfoKHR(3) Manual Page

## Name

VkVideoEncodeH265RateControlInfoKHR - Structure describing H.265 stream rate control parameters



## [](#_c_specification)C Specification

The `VkVideoEncodeH265RateControlInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_video_encode_h265
typedef struct VkVideoEncodeH265RateControlInfoKHR {
    VkStructureType                         sType;
    const void*                             pNext;
    VkVideoEncodeH265RateControlFlagsKHR    flags;
    uint32_t                                gopFrameCount;
    uint32_t                                idrPeriod;
    uint32_t                                consecutiveBFrameCount;
    uint32_t                                subLayerCount;
} VkVideoEncodeH265RateControlInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is a bitmask of [VkVideoEncodeH265RateControlFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265RateControlFlagBitsKHR.html) specifying H.265 rate control flags.
- `gopFrameCount` is the number of frames within a [group of pictures (GOP)](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h265-gop) intended to be used by the application. If it is 0, the rate control algorithm **may** assume an implementation-dependent GOP length. If it is `UINT32_MAX`, the GOP length is treated as infinite.
- `idrPeriod` is the interval, in terms of number of frames, between two [IDR frames](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h265-idr-pic) (see [IDR period](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h265-idr-period)). If it is 0, the rate control algorithm **may** assume an implementation-dependent IDR period. If it is `UINT32_MAX`, the IDR period is treated as infinite.
- `consecutiveBFrameCount` is the number of consecutive B frames between I and/or P frames within the [GOP](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h265-gop).
- `subLayerCount` specifies the number of H.265 sub-layers that the application intends to use.

## [](#_description)Description

When an instance of this structure is included in the `pNext` chain of the [VkVideoCodingControlInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoCodingControlInfoKHR.html) structure passed to the [vkCmdControlVideoCodingKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdControlVideoCodingKHR.html) command, and [VkVideoCodingControlInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoCodingControlInfoKHR.html)::`flags` includes `VK_VIDEO_CODING_CONTROL_ENCODE_RATE_CONTROL_BIT_KHR`, the parameters in this structure are used as guidance for the implementation’s rate control algorithm (see [Video Coding Control](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#video-coding-control)).

If `flags` includes `VK_VIDEO_ENCODE_H265_RATE_CONTROL_ATTEMPT_HRD_COMPLIANCE_BIT_KHR`, then the rate control state is reset to an initial state to meet HRD compliance requirements. Otherwise the new rate control state **may** be applied without a reset depending on the implementation and the specified rate control parameters.

Note

It would be possible to infer the picture type to be used when encoding a frame, on the basis of the values provided for `consecutiveBFrameCount`, `idrPeriod`, and `gopFrameCount`, but this inferred picture type will not be used by implementations to override the picture type provided to the video encode operation.

Valid Usage

- [](#VUID-VkVideoEncodeH265RateControlInfoKHR-flags-08291)VUID-VkVideoEncodeH265RateControlInfoKHR-flags-08291  
  If [VkVideoEncodeH265CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265CapabilitiesKHR.html)::`flags`, as returned by [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html) for the used video profile, does not include `VK_VIDEO_ENCODE_H265_CAPABILITY_HRD_COMPLIANCE_BIT_KHR`, then `flags` **must** not contain `VK_VIDEO_ENCODE_H265_RATE_CONTROL_ATTEMPT_HRD_COMPLIANCE_BIT_KHR`
- [](#VUID-VkVideoEncodeH265RateControlInfoKHR-flags-08292)VUID-VkVideoEncodeH265RateControlInfoKHR-flags-08292  
  If `flags` contains `VK_VIDEO_ENCODE_H265_RATE_CONTROL_REFERENCE_PATTERN_FLAT_BIT_KHR` or `VK_VIDEO_ENCODE_H265_RATE_CONTROL_REFERENCE_PATTERN_DYADIC_BIT_KHR`, then it **must** also contain `VK_VIDEO_ENCODE_H265_RATE_CONTROL_REGULAR_GOP_BIT_KHR`
- [](#VUID-VkVideoEncodeH265RateControlInfoKHR-flags-08293)VUID-VkVideoEncodeH265RateControlInfoKHR-flags-08293  
  If `flags` contains `VK_VIDEO_ENCODE_H265_RATE_CONTROL_REFERENCE_PATTERN_FLAT_BIT_KHR`, then it **must** not also contain `VK_VIDEO_ENCODE_H265_RATE_CONTROL_REFERENCE_PATTERN_DYADIC_BIT_KHR`
- [](#VUID-VkVideoEncodeH265RateControlInfoKHR-flags-08294)VUID-VkVideoEncodeH265RateControlInfoKHR-flags-08294  
  If `flags` contains `VK_VIDEO_ENCODE_H265_RATE_CONTROL_REGULAR_GOP_BIT_KHR`, then `gopFrameCount` **must** be greater than `0`
- [](#VUID-VkVideoEncodeH265RateControlInfoKHR-idrPeriod-08295)VUID-VkVideoEncodeH265RateControlInfoKHR-idrPeriod-08295  
  If `idrPeriod` is not `0`, then it **must** be greater than or equal to `gopFrameCount`
- [](#VUID-VkVideoEncodeH265RateControlInfoKHR-consecutiveBFrameCount-08296)VUID-VkVideoEncodeH265RateControlInfoKHR-consecutiveBFrameCount-08296  
  If `consecutiveBFrameCount` is not `0`, then it **must** be less than `gopFrameCount`

Valid Usage (Implicit)

- [](#VUID-VkVideoEncodeH265RateControlInfoKHR-sType-sType)VUID-VkVideoEncodeH265RateControlInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_ENCODE_H265_RATE_CONTROL_INFO_KHR`
- [](#VUID-VkVideoEncodeH265RateControlInfoKHR-flags-parameter)VUID-VkVideoEncodeH265RateControlInfoKHR-flags-parameter  
  `flags` **must** be a valid combination of [VkVideoEncodeH265RateControlFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265RateControlFlagBitsKHR.html) values

## [](#_see_also)See Also

[VK\_KHR\_video\_encode\_h265](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_h265.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkVideoEncodeH265RateControlFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265RateControlFlagsKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoEncodeH265RateControlInfoKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0