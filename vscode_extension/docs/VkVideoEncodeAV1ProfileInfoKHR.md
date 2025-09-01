# VkVideoEncodeAV1ProfileInfoKHR(3) Manual Page

## Name

VkVideoEncodeAV1ProfileInfoKHR - Structure specifying AV1 encode-specific video profile parameters



## [](#_c_specification)C Specification

A video profile supporting AV1 video encode operations is specified by setting [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoProfileInfoKHR.html)::`videoCodecOperation` to `VK_VIDEO_CODEC_OPERATION_ENCODE_AV1_BIT_KHR` and adding a `VkVideoEncodeAV1ProfileInfoKHR` structure to the [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoProfileInfoKHR.html)::`pNext` chain.

The `VkVideoEncodeAV1ProfileInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_video_encode_av1
typedef struct VkVideoEncodeAV1ProfileInfoKHR {
    VkStructureType       sType;
    const void*           pNext;
    StdVideoAV1Profile    stdProfile;
} VkVideoEncodeAV1ProfileInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `stdProfile` is a `StdVideoAV1Profile` value specifying the AV1 codec profile, as defined in section A.2 of the [AV1 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#aomedia-av1).

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkVideoEncodeAV1ProfileInfoKHR-sType-sType)VUID-VkVideoEncodeAV1ProfileInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_ENCODE_AV1_PROFILE_INFO_KHR`

## [](#_see_also)See Also

[VK\_KHR\_video\_encode\_av1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_av1.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoEncodeAV1ProfileInfoKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0