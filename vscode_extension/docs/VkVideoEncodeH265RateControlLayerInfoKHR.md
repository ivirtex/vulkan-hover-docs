# VkVideoEncodeH265RateControlLayerInfoKHR(3) Manual Page

## Name

VkVideoEncodeH265RateControlLayerInfoKHR - Structure describing H.265 per-layer rate control parameters



## [](#_c_specification)C Specification

The `VkVideoEncodeH265RateControlLayerInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_video_encode_h265
typedef struct VkVideoEncodeH265RateControlLayerInfoKHR {
    VkStructureType                  sType;
    const void*                      pNext;
    VkBool32                         useMinQp;
    VkVideoEncodeH265QpKHR           minQp;
    VkBool32                         useMaxQp;
    VkVideoEncodeH265QpKHR           maxQp;
    VkBool32                         useMaxFrameSize;
    VkVideoEncodeH265FrameSizeKHR    maxFrameSize;
} VkVideoEncodeH265RateControlLayerInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `useMinQp` indicates whether the QP values determined by rate control will be clamped to the lower bounds on the QP values specified in `minQp`.
- `minQp` specifies the lower bounds on the QP values, for each picture type, that the implementation’s rate control algorithm will use when `useMinQp` is `VK_TRUE`.
- `useMaxQp` indicates whether the QP values determined by rate control will be clamped to the upper bounds on the QP values specified in `maxQp`.
- `maxQp` specifies the upper bounds on the QP values, for each picture type, that the implementation’s rate control algorithm will use when `useMaxQp` is `VK_TRUE`.
- `useMaxFrameSize` indicates whether the implementation’s rate control algorithm **should** use the values specified in `maxFrameSize` as the upper bounds on the encoded frame size for each picture type.
- `maxFrameSize` specifies the upper bounds on the encoded frame size, for each picture type, when `useMaxFrameSize` is `VK_TRUE`.

## [](#_description)Description

When used, the values in `minQp` and `maxQp` guarantee that the effective QP values used by the implementation will respect those lower and upper bounds, respectively. However, limiting the range of QP values that the implementation is able to use will also limit the capabilities of the implementation’s rate control algorithm to comply to other constraints. In particular, the implementation **may** not be able to comply to the following:

- The average and/or peak [bitrate](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-bitrate) values to be used for the encoded bitstream specified in the `averageBitrate` and `maxBitrate` members of the [VkVideoEncodeRateControlLayerInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeRateControlLayerInfoKHR.html) structure.
- The upper bounds on the encoded frame size, for each picture type, specified in the `maxFrameSize` member of `VkVideoEncodeH265RateControlLayerInfoKHR`.

Note

In general, applications need to configure rate control parameters appropriately in order to be able to get the desired rate control behavior, as described in the [Video Encode Rate Control](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-rate-control) section.

When an instance of this structure is included in the `pNext` chain of a [VkVideoEncodeRateControlLayerInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeRateControlLayerInfoKHR.html) structure specified in one of the elements of the `pLayers` array member of the [VkVideoEncodeRateControlInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeRateControlInfoKHR.html) structure passed to the [vkCmdControlVideoCodingKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdControlVideoCodingKHR.html) command, [VkVideoCodingControlInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoCodingControlInfoKHR.html)::`flags` includes `VK_VIDEO_CODING_CONTROL_ENCODE_RATE_CONTROL_BIT_KHR`, and the bound video session was created with the video codec operation `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR`, it specifies the H.265-specific rate control parameters of the rate control layer corresponding to that element of `pLayers`.

Valid Usage

- [](#VUID-VkVideoEncodeH265RateControlLayerInfoKHR-useMinQp-08297)VUID-VkVideoEncodeH265RateControlLayerInfoKHR-useMinQp-08297  
  If `useMinQp` is `VK_TRUE`, then the `qpI`, `qpP`, and `qpB` members of `minQp` **must** all be between [VkVideoEncodeH265CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265CapabilitiesKHR.html)::`minQp` and [VkVideoEncodeH265CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265CapabilitiesKHR.html)::`maxQp`, as returned by [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html) for the used video profile
- [](#VUID-VkVideoEncodeH265RateControlLayerInfoKHR-useMaxQp-08298)VUID-VkVideoEncodeH265RateControlLayerInfoKHR-useMaxQp-08298  
  If `useMaxQp` is `VK_TRUE`, then the `qpI`, `qpP`, and `qpB` members of `maxQp` **must** all be between [VkVideoEncodeH265CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265CapabilitiesKHR.html)::`minQp` and [VkVideoEncodeH265CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265CapabilitiesKHR.html)::`maxQp`, as returned by [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html) for the used video profile
- [](#VUID-VkVideoEncodeH265RateControlLayerInfoKHR-useMinQp-08299)VUID-VkVideoEncodeH265RateControlLayerInfoKHR-useMinQp-08299  
  If `useMinQp` is `VK_TRUE` and [VkVideoEncodeH265CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265CapabilitiesKHR.html)::`flags`, as returned by [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html) for the used video profile, does not include `VK_VIDEO_ENCODE_H265_CAPABILITY_PER_PICTURE_TYPE_MIN_MAX_QP_BIT_KHR`, then the `qpI`, `qpP`, and `qpB` members of `minQp` **must** all specify the same value
- [](#VUID-VkVideoEncodeH265RateControlLayerInfoKHR-useMaxQp-08300)VUID-VkVideoEncodeH265RateControlLayerInfoKHR-useMaxQp-08300  
  If `useMaxQp` is `VK_TRUE` and [VkVideoEncodeH265CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265CapabilitiesKHR.html)::`flags`, as returned by [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html) for the used video profile, does not include `VK_VIDEO_ENCODE_H265_CAPABILITY_PER_PICTURE_TYPE_MIN_MAX_QP_BIT_KHR`, then the `qpI`, `qpP`, and `qpB` members of `maxQp` **must** all specify the same value
- [](#VUID-VkVideoEncodeH265RateControlLayerInfoKHR-useMinQp-08375)VUID-VkVideoEncodeH265RateControlLayerInfoKHR-useMinQp-08375  
  If `useMinQp` and `useMaxQp` are both `VK_TRUE`, then the `qpI`, `qpP`, and `qpB` members of `minQp` **must** all be less than or equal to the respective members of `maxQp`

Valid Usage (Implicit)

- [](#VUID-VkVideoEncodeH265RateControlLayerInfoKHR-sType-sType)VUID-VkVideoEncodeH265RateControlLayerInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_ENCODE_H265_RATE_CONTROL_LAYER_INFO_KHR`
- [](#VUID-VkVideoEncodeH265RateControlLayerInfoKHR-minQp-parameter)VUID-VkVideoEncodeH265RateControlLayerInfoKHR-minQp-parameter  
  `minQp` **must** be a valid [VkVideoEncodeH265QpKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265QpKHR.html) structure
- [](#VUID-VkVideoEncodeH265RateControlLayerInfoKHR-maxQp-parameter)VUID-VkVideoEncodeH265RateControlLayerInfoKHR-maxQp-parameter  
  `maxQp` **must** be a valid [VkVideoEncodeH265QpKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265QpKHR.html) structure
- [](#VUID-VkVideoEncodeH265RateControlLayerInfoKHR-maxFrameSize-parameter)VUID-VkVideoEncodeH265RateControlLayerInfoKHR-maxFrameSize-parameter  
  `maxFrameSize` **must** be a valid [VkVideoEncodeH265FrameSizeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265FrameSizeKHR.html) structure

## [](#_see_also)See Also

[VK\_KHR\_video\_encode\_h265](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_h265.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkVideoEncodeH265FrameSizeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265FrameSizeKHR.html), [VkVideoEncodeH265QpKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265QpKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoEncodeH265RateControlLayerInfoKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0