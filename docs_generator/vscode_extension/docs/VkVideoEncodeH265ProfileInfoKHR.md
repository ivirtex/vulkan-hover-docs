# VkVideoEncodeH265ProfileInfoKHR(3) Manual Page

## Name

VkVideoEncodeH265ProfileInfoKHR - Structure specifying H.265 encode-specific video profile parameters



## [](#_c_specification)C Specification

A video profile supporting H.265 video encode operations is specified by setting [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoProfileInfoKHR.html)::`videoCodecOperation` to `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR` and adding a `VkVideoEncodeH265ProfileInfoKHR` structure to the [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoProfileInfoKHR.html)::`pNext` chain.

The `VkVideoEncodeH265ProfileInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_video_encode_h265
typedef struct VkVideoEncodeH265ProfileInfoKHR {
    VkStructureType           sType;
    const void*               pNext;
    StdVideoH265ProfileIdc    stdProfileIdc;
} VkVideoEncodeH265ProfileInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `stdProfileIdc` is a `StdVideoH265ProfileIdc` value specifying the H.265 codec profile IDC, as defined in section A.3 of the [ITU-T H.265 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#itu-t-h265).

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkVideoEncodeH265ProfileInfoKHR-sType-sType)VUID-VkVideoEncodeH265ProfileInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_ENCODE_H265_PROFILE_INFO_KHR`

## [](#_see_also)See Also

[VK\_KHR\_video\_encode\_h265](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_h265.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoEncodeH265ProfileInfoKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0