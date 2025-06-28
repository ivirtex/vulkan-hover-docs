# VkVideoEncodeAV1CapabilityFlagBitsKHR(3) Manual Page

## Name

VkVideoEncodeAV1CapabilityFlagBitsKHR - AV1 encode capability flags



## [](#_c_specification)C Specification

Bits which **may** be set in [VkVideoEncodeAV1CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1CapabilitiesKHR.html)::`flags`, indicating the AV1 encoding capabilities supported, are:

```c++
// Provided by VK_KHR_video_encode_av1
typedef enum VkVideoEncodeAV1CapabilityFlagBitsKHR {
    VK_VIDEO_ENCODE_AV1_CAPABILITY_PER_RATE_CONTROL_GROUP_MIN_MAX_Q_INDEX_BIT_KHR = 0x00000001,
    VK_VIDEO_ENCODE_AV1_CAPABILITY_GENERATE_OBU_EXTENSION_HEADER_BIT_KHR = 0x00000002,
    VK_VIDEO_ENCODE_AV1_CAPABILITY_PRIMARY_REFERENCE_CDF_ONLY_BIT_KHR = 0x00000004,
    VK_VIDEO_ENCODE_AV1_CAPABILITY_FRAME_SIZE_OVERRIDE_BIT_KHR = 0x00000008,
    VK_VIDEO_ENCODE_AV1_CAPABILITY_MOTION_VECTOR_SCALING_BIT_KHR = 0x00000010,
} VkVideoEncodeAV1CapabilityFlagBitsKHR;
```

## [](#_description)Description

- `VK_VIDEO_ENCODE_AV1_CAPABILITY_PER_RATE_CONTROL_GROUP_MIN_MAX_Q_INDEX_BIT_KHR` specifies support for specifying different quantizer index values in the members of [VkVideoEncodeAV1QIndexKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1QIndexKHR.html).
- `VK_VIDEO_ENCODE_AV1_CAPABILITY_GENERATE_OBU_EXTENSION_HEADER_BIT_KHR` specifies support for generating OBU extension headers, as defined in section 5.3.3 of the [AV1 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#aomedia-av1).
- `VK_VIDEO_ENCODE_AV1_CAPABILITY_PRIMARY_REFERENCE_CDF_ONLY_BIT_KHR` specifies support for using the primary reference frame indicated by the value of `StdVideoEncodeAV1PictureInfo`::`primary_ref_frame` in the [AV1 picture information](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-picture-info) only for CDF data reference, as defined in section 6.8.2 of the [AV1 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#aomedia-av1).
- `VK_VIDEO_ENCODE_AV1_CAPABILITY_FRAME_SIZE_OVERRIDE_BIT_KHR` specifies support for encoding a picture with a frame size different from the maximum frame size defined in the [active AV1 sequence header](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-active-sequence-header). If this capability is not supported, then `frame_size_override_flag` **must** not be set in the [AV1 picture information](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-picture-info) of the encoded frame and the coded extent of the [encode input picture](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-input-picture) **must** match the maximum coded extent allowed by the [active AV1 sequence header](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-active-sequence-header), i.e. (`max_frame_width_minus_1` + 1, `max_frame_height_minus_1` + 1).
- `VK_VIDEO_ENCODE_AV1_CAPABILITY_MOTION_VECTOR_SCALING_BIT_KHR` specifies support for motion vector scaling, as defined in section 7.11.3.3 of the [AV1 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#aomedia-av1). If this capability is not supported, then the coded extent of all [active reference pictures](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#active-reference-pictures) **must** match the coded extent of the [encode input picture](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-input-picture). This capability **may** only be supported by a video profile when `VK_VIDEO_ENCODE_AV1_CAPABILITY_FRAME_SIZE_OVERRIDE_BIT_KHR` is also supported.

## [](#_see_also)See Also

[VK\_KHR\_video\_encode\_av1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_av1.html), [VkVideoEncodeAV1CapabilityFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1CapabilityFlagsKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoEncodeAV1CapabilityFlagBitsKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0