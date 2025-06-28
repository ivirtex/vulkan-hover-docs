# VkVideoEncodeAV1GopRemainingFrameInfoKHR(3) Manual Page

## Name

VkVideoEncodeAV1GopRemainingFrameInfoKHR - Structure specifying AV1 encode rate control GOP remaining frame counts



## [](#_c_specification)C Specification

The `VkVideoEncodeAV1GopRemainingFrameInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_video_encode_av1
typedef struct VkVideoEncodeAV1GopRemainingFrameInfoKHR {
    VkStructureType    sType;
    const void*        pNext;
    VkBool32           useGopRemainingFrames;
    uint32_t           gopRemainingIntra;
    uint32_t           gopRemainingPredictive;
    uint32_t           gopRemainingBipredictive;
} VkVideoEncodeAV1GopRemainingFrameInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `useGopRemainingFrames` indicates whether the implementation’s rate control algorithm **should** use the values specified in `gopRemainingIntra`, `gopRemainingPredictive`, and `gopRemainingBipredictive`. If `useGopRemainingFrames` is `VK_FALSE`, then the values of `gopRemainingIntra`, `gopRemainingPredictive`, and `gopRemainingBipredictive` are ignored.
- `gopRemainingIntra` specifies the number of frames encoded with `VK_VIDEO_ENCODE_AV1_RATE_CONTROL_GROUP_INTRA_KHR` the implementation’s rate control algorithm **should** assume to be remaining in the [GOP](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-gop) prior to executing the next video encode operation.
- `gopRemainingPredictive` specifies the number of frames encoded with `VK_VIDEO_ENCODE_AV1_RATE_CONTROL_GROUP_PREDICTIVE_KHR` the implementation’s rate control algorithm **should** assume to be remaining in the [GOP](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-gop) prior to executing the next video encode operation.
- `gopRemainingBipredictive` specifies the number of frames encoded with `VK_VIDEO_ENCODE_AV1_RATE_CONTROL_GROUP_BIPREDICTIVE_KHR` the implementation’s rate control algorithm **should** assume to be remaining in the [GOP](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-gop) prior to executing the next video encode operation.

## [](#_description)Description

Setting `useGopRemainingFrames` to `VK_TRUE` and including this structure in the `pNext` chain of [VkVideoBeginCodingInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoBeginCodingInfoKHR.html) is only mandatory if the [VkVideoEncodeAV1CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1CapabilitiesKHR.html)::`requiresGopRemainingFrames` reported for the used [video profile](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#video-profiles) is `VK_TRUE`. However, implementations **may** use these remaining frame counts, when specified, even when it is not required. In particular, when the application does not use a [regular GOP structure](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-regular-gop), these values **may** provide additional guidance for the implementation’s rate control algorithm.

The [VkVideoEncodeAV1CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1CapabilitiesKHR.html)::`prefersGopRemainingFrames` capability is also used to indicate that the implementation’s rate control algorithm **may** operate more accurately if the application specifies the remaining frame counts using this structure.

As with other rate control guidance values, if the effective order and number of frames encoded by the application are not in line with the remaining frame counts specified in this structure at any given point, then the behavior of the implementation’s rate control algorithm **may** deviate from the one expected by the application.

Valid Usage (Implicit)

- [](#VUID-VkVideoEncodeAV1GopRemainingFrameInfoKHR-sType-sType)VUID-VkVideoEncodeAV1GopRemainingFrameInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_ENCODE_AV1_GOP_REMAINING_FRAME_INFO_KHR`

## [](#_see_also)See Also

[VK\_KHR\_video\_encode\_av1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_av1.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoEncodeAV1GopRemainingFrameInfoKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0