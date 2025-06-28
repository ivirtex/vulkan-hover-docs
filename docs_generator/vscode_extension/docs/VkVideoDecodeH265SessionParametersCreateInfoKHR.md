# VkVideoDecodeH265SessionParametersCreateInfoKHR(3) Manual Page

## Name

VkVideoDecodeH265SessionParametersCreateInfoKHR - Structure specifies H.265 decoder parameter set information



## [](#_c_specification)C Specification

When a [video session parameters](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#video-session-parameters) object is created with the codec operation `VK_VIDEO_CODEC_OPERATION_DECODE_H265_BIT_KHR`, the [VkVideoSessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionParametersCreateInfoKHR.html)::`pNext` chain **must** include a `VkVideoDecodeH265SessionParametersCreateInfoKHR` structure specifying the capacity and initial contents of the object.

The `VkVideoDecodeH265SessionParametersCreateInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_video_decode_h265
typedef struct VkVideoDecodeH265SessionParametersCreateInfoKHR {
    VkStructureType                                        sType;
    const void*                                            pNext;
    uint32_t                                               maxStdVPSCount;
    uint32_t                                               maxStdSPSCount;
    uint32_t                                               maxStdPPSCount;
    const VkVideoDecodeH265SessionParametersAddInfoKHR*    pParametersAddInfo;
} VkVideoDecodeH265SessionParametersCreateInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `maxStdVPSCount` is the maximum number of [H.265 VPS](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-h265-vps) entries the created `VkVideoSessionParametersKHR` **can** contain.
- `maxStdSPSCount` is the maximum number of [H.265 SPS](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-h265-sps) entries the created `VkVideoSessionParametersKHR` **can** contain.
- `maxStdPPSCount` is the maximum number of [H.265 PPS](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-h265-pps) entries the created `VkVideoSessionParametersKHR` **can** contain.
- `pParametersAddInfo` is `NULL` or a pointer to a [VkVideoDecodeH265SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH265SessionParametersAddInfoKHR.html) structure specifying H.265 parameters to add upon object creation.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkVideoDecodeH265SessionParametersCreateInfoKHR-sType-sType)VUID-VkVideoDecodeH265SessionParametersCreateInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_DECODE_H265_SESSION_PARAMETERS_CREATE_INFO_KHR`
- [](#VUID-VkVideoDecodeH265SessionParametersCreateInfoKHR-pParametersAddInfo-parameter)VUID-VkVideoDecodeH265SessionParametersCreateInfoKHR-pParametersAddInfo-parameter  
  If `pParametersAddInfo` is not `NULL`, `pParametersAddInfo` **must** be a valid pointer to a valid [VkVideoDecodeH265SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH265SessionParametersAddInfoKHR.html) structure

## [](#_see_also)See Also

[VK\_KHR\_video\_decode\_h265](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_decode_h265.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkVideoDecodeH265SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH265SessionParametersAddInfoKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoDecodeH265SessionParametersCreateInfoKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0