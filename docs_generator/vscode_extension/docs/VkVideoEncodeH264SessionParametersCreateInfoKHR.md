# VkVideoEncodeH264SessionParametersCreateInfoKHR(3) Manual Page

## Name

VkVideoEncodeH264SessionParametersCreateInfoKHR - Structure specifies H.264 encoder parameter set information



## [](#_c_specification)C Specification

When a [video session parameters](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#video-session-parameters) object is created with the codec operation `VK_VIDEO_CODEC_OPERATION_ENCODE_H264_BIT_KHR`, the [VkVideoSessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionParametersCreateInfoKHR.html)::`pNext` chain **must** include a `VkVideoEncodeH264SessionParametersCreateInfoKHR` structure specifying the capacity and initial contents of the object.

The `VkVideoEncodeH264SessionParametersCreateInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_video_encode_h264
typedef struct VkVideoEncodeH264SessionParametersCreateInfoKHR {
    VkStructureType                                        sType;
    const void*                                            pNext;
    uint32_t                                               maxStdSPSCount;
    uint32_t                                               maxStdPPSCount;
    const VkVideoEncodeH264SessionParametersAddInfoKHR*    pParametersAddInfo;
} VkVideoEncodeH264SessionParametersCreateInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `maxStdSPSCount` is the maximum number of [H.264 SPS](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h264-sps) entries the created `VkVideoSessionParametersKHR` **can** contain.
- `maxStdPPSCount` is the maximum number of [H.264 PPS](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h264-pps) entries the created `VkVideoSessionParametersKHR` **can** contain.
- `pParametersAddInfo` is `NULL` or a pointer to a [VkVideoEncodeH264SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264SessionParametersAddInfoKHR.html) structure specifying H.264 parameters to add upon object creation.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkVideoEncodeH264SessionParametersCreateInfoKHR-sType-sType)VUID-VkVideoEncodeH264SessionParametersCreateInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_ENCODE_H264_SESSION_PARAMETERS_CREATE_INFO_KHR`
- [](#VUID-VkVideoEncodeH264SessionParametersCreateInfoKHR-pParametersAddInfo-parameter)VUID-VkVideoEncodeH264SessionParametersCreateInfoKHR-pParametersAddInfo-parameter  
  If `pParametersAddInfo` is not `NULL`, `pParametersAddInfo` **must** be a valid pointer to a valid [VkVideoEncodeH264SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264SessionParametersAddInfoKHR.html) structure

## [](#_see_also)See Also

[VK\_KHR\_video\_encode\_h264](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_h264.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkVideoEncodeH264SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264SessionParametersAddInfoKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoEncodeH264SessionParametersCreateInfoKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0