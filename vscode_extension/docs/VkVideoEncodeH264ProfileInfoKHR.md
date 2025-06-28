# VkVideoEncodeH264ProfileInfoKHR(3) Manual Page

## Name

VkVideoEncodeH264ProfileInfoKHR - Structure specifying H.264 encode-specific video profile parameters



## [](#_c_specification)C Specification

A video profile supporting H.264 video encode operations is specified by setting [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoProfileInfoKHR.html)::`videoCodecOperation` to `VK_VIDEO_CODEC_OPERATION_ENCODE_H264_BIT_KHR` and adding a `VkVideoEncodeH264ProfileInfoKHR` structure to the [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoProfileInfoKHR.html)::`pNext` chain.

The `VkVideoEncodeH264ProfileInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_video_encode_h264
typedef struct VkVideoEncodeH264ProfileInfoKHR {
    VkStructureType           sType;
    const void*               pNext;
    StdVideoH264ProfileIdc    stdProfileIdc;
} VkVideoEncodeH264ProfileInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `stdProfileIdc` is a `StdVideoH264ProfileIdc` value specifying the H.264 codec profile IDC, where enum constant `STD_VIDEO_H264_PROFILE_IDC_BASELINE` identifies the Constrained Baseline profile as defined in A.2.1.1 of the [ITU-T H.264 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#itu-t-h264), and all other values correspond to profiles defined in section A.2 of the [ITU-T H.264 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#itu-t-h264).

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkVideoEncodeH264ProfileInfoKHR-sType-sType)VUID-VkVideoEncodeH264ProfileInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_ENCODE_H264_PROFILE_INFO_KHR`

## [](#_see_also)See Also

[VK\_KHR\_video\_encode\_h264](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_h264.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoEncodeH264ProfileInfoKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0