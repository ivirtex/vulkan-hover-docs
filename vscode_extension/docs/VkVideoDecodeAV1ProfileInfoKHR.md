# VkVideoDecodeAV1ProfileInfoKHR(3) Manual Page

## Name

VkVideoDecodeAV1ProfileInfoKHR - Structure specifying AV1 decode profile



## [](#_c_specification)C Specification

A video profile supporting AV1 video decode operations is specified by setting [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoProfileInfoKHR.html)::`videoCodecOperation` to `VK_VIDEO_CODEC_OPERATION_DECODE_AV1_BIT_KHR` and adding a `VkVideoDecodeAV1ProfileInfoKHR` structure to the [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoProfileInfoKHR.html)::`pNext` chain.

The `VkVideoDecodeAV1ProfileInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_video_decode_av1
typedef struct VkVideoDecodeAV1ProfileInfoKHR {
    VkStructureType       sType;
    const void*           pNext;
    StdVideoAV1Profile    stdProfile;
    VkBool32              filmGrainSupport;
} VkVideoDecodeAV1ProfileInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `stdProfile` is a `StdVideoAV1Profile` value specifying the AV1 codec profile, as defined in section A.2 of the [AV1 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#aomedia-av1).
- []()`filmGrainSupport` specifies whether AV1 film grain, as defined in section 7.8.3 of the [AV1 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#aomedia-av1), **can** be used with the video profile. When this member is `VK_TRUE`, video session objects created against the video profile will be able to decode pictures that have [film grain](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-av1-film-grain) enabled.

## [](#_description)Description

Note

Enabling `filmGrainSupport` **may** increase the memory requirements of video sessions and/or video picture resources on some implementations.

Valid Usage (Implicit)

- [](#VUID-VkVideoDecodeAV1ProfileInfoKHR-sType-sType)VUID-VkVideoDecodeAV1ProfileInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_DECODE_AV1_PROFILE_INFO_KHR`

## [](#_see_also)See Also

[VK\_KHR\_video\_decode\_av1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_decode_av1.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoDecodeAV1ProfileInfoKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0