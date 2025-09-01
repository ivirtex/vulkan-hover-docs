# VkVideoDecodeVP9ProfileInfoKHR(3) Manual Page

## Name

VkVideoDecodeVP9ProfileInfoKHR - Structure specifying VP9 decode profile



## [](#_c_specification)C Specification

A video profile supporting VP9 video decode operations is specified by setting [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoProfileInfoKHR.html)::`videoCodecOperation` to `VK_VIDEO_CODEC_OPERATION_DECODE_VP9_BIT_KHR` and adding a `VkVideoDecodeVP9ProfileInfoKHR` structure to the [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoProfileInfoKHR.html)::`pNext` chain.

The `VkVideoDecodeVP9ProfileInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_video_decode_vp9
typedef struct VkVideoDecodeVP9ProfileInfoKHR {
    VkStructureType       sType;
    const void*           pNext;
    StdVideoVP9Profile    stdProfile;
} VkVideoDecodeVP9ProfileInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `stdProfile` is a `StdVideoVP9Profile` value specifying the VP9 codec profile, as defined in section 7.2 of the [VP9 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#google-vp9).

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkVideoDecodeVP9ProfileInfoKHR-sType-sType)VUID-VkVideoDecodeVP9ProfileInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_DECODE_VP9_PROFILE_INFO_KHR`

## [](#_see_also)See Also

[VK\_KHR\_video\_decode\_vp9](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_decode_vp9.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoDecodeVP9ProfileInfoKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0