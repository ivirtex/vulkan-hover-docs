# VkVideoEncodeH264SessionParametersGetInfoKHR(3) Manual Page

## Name

VkVideoEncodeH264SessionParametersGetInfoKHR - Structure specifying parameters for retrieving encoded H.264 parameter set data



## [](#_c_specification)C Specification

The `VkVideoEncodeH264SessionParametersGetInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_video_encode_h264
typedef struct VkVideoEncodeH264SessionParametersGetInfoKHR {
    VkStructureType    sType;
    const void*        pNext;
    VkBool32           writeStdSPS;
    VkBool32           writeStdPPS;
    uint32_t           stdSPSId;
    uint32_t           stdPPSId;
} VkVideoEncodeH264SessionParametersGetInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `writeStdSPS` indicates whether the encoded [H.264 sequence parameter set](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h264-sps) identified by `stdSPSId` is requested to be retrieved.
- `writeStdPPS` indicates whether the encoded [H.264 picture parameter set](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h264-pps) identified by the pair constructed from `stdSPSId` and `stdPPSId` is requested to be retrieved.
- `stdSPSId` specifies the H.264 sequence parameter set ID used to identify the retrieved H.264 sequence and/or picture parameter set(s).
- `stdPPSId` specifies the H.264 picture parameter set ID used to identify the retrieved H.264 picture parameter set when `writeStdPPS` is `VK_TRUE`.

## [](#_description)Description

When this structure is specified in the `pNext` chain of the [VkVideoEncodeSessionParametersGetInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeSessionParametersGetInfoKHR.html) structure passed to [vkGetEncodedVideoSessionParametersKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetEncodedVideoSessionParametersKHR.html), the command will write encoded parameter data to the output buffer in the following order:

1. The [H.264 sequence parameter set](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h264-sps) identified by `stdSPSId`, if `writeStdSPS` is `VK_TRUE`.
2. The [H.264 picture parameter set](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h264-pps) identified by the pair constructed from `stdSPSId` and `stdPPSId`, if `writeStdPPS` is `VK_TRUE`.

Valid Usage

- [](#VUID-VkVideoEncodeH264SessionParametersGetInfoKHR-writeStdSPS-08279)VUID-VkVideoEncodeH264SessionParametersGetInfoKHR-writeStdSPS-08279  
  At least one of `writeStdSPS` and `writeStdPPS` **must** be `VK_TRUE`

Valid Usage (Implicit)

- [](#VUID-VkVideoEncodeH264SessionParametersGetInfoKHR-sType-sType)VUID-VkVideoEncodeH264SessionParametersGetInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_ENCODE_H264_SESSION_PARAMETERS_GET_INFO_KHR`

## [](#_see_also)See Also

[VK\_KHR\_video\_encode\_h264](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_h264.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoEncodeH264SessionParametersGetInfoKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0