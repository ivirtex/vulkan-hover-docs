# VkVideoEncodeH265SessionParametersFeedbackInfoKHR(3) Manual Page

## Name

VkVideoEncodeH265SessionParametersFeedbackInfoKHR - Structure providing feedback about the requested H.265 video session parameters



## [](#_c_specification)C Specification

The `VkVideoEncodeH265SessionParametersFeedbackInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_video_encode_h265
typedef struct VkVideoEncodeH265SessionParametersFeedbackInfoKHR {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           hasStdVPSOverrides;
    VkBool32           hasStdSPSOverrides;
    VkBool32           hasStdPPSOverrides;
} VkVideoEncodeH265SessionParametersFeedbackInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `hasStdVPSOverrides` indicates whether any of the parameters of the requested [H.265 video parameter set](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h265-vps), if one was requested via [VkVideoEncodeH265SessionParametersGetInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265SessionParametersGetInfoKHR.html)::`writeStdVPS`, were [overridden](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-overrides) by the implementation.
- `hasStdSPSOverrides` indicates whether any of the parameters of the requested [H.265 sequence parameter set](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h265-sps), if one was requested via [VkVideoEncodeH265SessionParametersGetInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265SessionParametersGetInfoKHR.html)::`writeStdSPS`, were [overridden](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-overrides) by the implementation.
- `hasStdPPSOverrides` indicates whether any of the parameters of the requested [H.265 picture parameter set](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h265-pps), if one was requested via [VkVideoEncodeH265SessionParametersGetInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265SessionParametersGetInfoKHR.html)::`writeStdPPS`, were [overridden](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-overrides) by the implementation.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkVideoEncodeH265SessionParametersFeedbackInfoKHR-sType-sType)VUID-VkVideoEncodeH265SessionParametersFeedbackInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_ENCODE_H265_SESSION_PARAMETERS_FEEDBACK_INFO_KHR`

## [](#_see_also)See Also

[VK\_KHR\_video\_encode\_h265](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_h265.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoEncodeH265SessionParametersFeedbackInfoKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0