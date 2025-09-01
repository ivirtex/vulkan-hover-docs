# VkVideoEncodeAV1SessionParametersCreateInfoKHR(3) Manual Page

## Name

VkVideoEncodeAV1SessionParametersCreateInfoKHR - Structure specifies AV1 encoder parameter set information



## [](#_c_specification)C Specification

When a [video session parameters](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#video-session-parameters) object is created with the codec operation `VK_VIDEO_CODEC_OPERATION_ENCODE_AV1_BIT_KHR`, the [VkVideoSessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionParametersCreateInfoKHR.html)::`pNext` chain **must** include a `VkVideoEncodeAV1SessionParametersCreateInfoKHR` structure specifying the contents of the object.

The `VkVideoEncodeAV1SessionParametersCreateInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_video_encode_av1
typedef struct VkVideoEncodeAV1SessionParametersCreateInfoKHR {
    VkStructureType                               sType;
    const void*                                   pNext;
    const StdVideoAV1SequenceHeader*              pStdSequenceHeader;
    const StdVideoEncodeAV1DecoderModelInfo*      pStdDecoderModelInfo;
    uint32_t                                      stdOperatingPointCount;
    const StdVideoEncodeAV1OperatingPointInfo*    pStdOperatingPoints;
} VkVideoEncodeAV1SessionParametersCreateInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `pStdSequenceHeader` is a pointer to a `StdVideoAV1SequenceHeader` structure describing parameters of the [AV1 sequence header](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-sequence-header) entry to store in the created object.
- `pStdDecoderModelInfo` is `NULL` or a pointer to a `StdVideoEncodeAV1DecoderModelInfo` structure specifying the [AV1 decoder model information](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-decoder-model-info) to store in the created object.
- `stdOperatingPointCount` is the number of elements in the `pStdOperatingPoints` array.
- `pStdOperatingPoints` is `NULL` or a pointer to an array of `stdOperatingPointCount` number of `StdVideoEncodeAV1OperatingPointInfo` structures specifying the [AV1 operating point information](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-operating-points) to store in the created object. Each element i specifies the parameter values corresponding to element i of the syntax elements defined in section 6.4 of the [AV1 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#aomedia-av1).

## [](#_description)Description

Valid Usage

- [](#VUID-VkVideoEncodeAV1SessionParametersCreateInfoKHR-pStdSequenceHeader-10288)VUID-VkVideoEncodeAV1SessionParametersCreateInfoKHR-pStdSequenceHeader-10288  
  `pStdSequenceHeader->flags.film_grain_params_present` **must** be zero

Valid Usage (Implicit)

- [](#VUID-VkVideoEncodeAV1SessionParametersCreateInfoKHR-sType-sType)VUID-VkVideoEncodeAV1SessionParametersCreateInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_ENCODE_AV1_SESSION_PARAMETERS_CREATE_INFO_KHR`
- [](#VUID-VkVideoEncodeAV1SessionParametersCreateInfoKHR-pStdSequenceHeader-parameter)VUID-VkVideoEncodeAV1SessionParametersCreateInfoKHR-pStdSequenceHeader-parameter  
  `pStdSequenceHeader` **must** be a valid pointer to a valid `StdVideoAV1SequenceHeader` value
- [](#VUID-VkVideoEncodeAV1SessionParametersCreateInfoKHR-pStdDecoderModelInfo-parameter)VUID-VkVideoEncodeAV1SessionParametersCreateInfoKHR-pStdDecoderModelInfo-parameter  
  If `pStdDecoderModelInfo` is not `NULL`, `pStdDecoderModelInfo` **must** be a valid pointer to a valid `StdVideoEncodeAV1DecoderModelInfo` value
- [](#VUID-VkVideoEncodeAV1SessionParametersCreateInfoKHR-pStdOperatingPoints-parameter)VUID-VkVideoEncodeAV1SessionParametersCreateInfoKHR-pStdOperatingPoints-parameter  
  If `stdOperatingPointCount` is not `0`, and `pStdOperatingPoints` is not `NULL`, `pStdOperatingPoints` **must** be a valid pointer to an array of `stdOperatingPointCount` `StdVideoEncodeAV1OperatingPointInfo` values

## [](#_see_also)See Also

[VK\_KHR\_video\_encode\_av1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_av1.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoEncodeAV1SessionParametersCreateInfoKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0