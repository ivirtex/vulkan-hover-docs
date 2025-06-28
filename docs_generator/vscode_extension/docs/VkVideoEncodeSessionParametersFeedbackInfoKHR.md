# VkVideoEncodeSessionParametersFeedbackInfoKHR(3) Manual Page

## Name

VkVideoEncodeSessionParametersFeedbackInfoKHR - Structure providing feedback about the requested video session parameters



## [](#_c_specification)C Specification

The `VkVideoEncodeSessionParametersFeedbackInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_video_encode_queue
typedef struct VkVideoEncodeSessionParametersFeedbackInfoKHR {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           hasOverrides;
} VkVideoEncodeSessionParametersFeedbackInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `hasOverrides` indicates whether any of the requested parameter data were [overridden](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-overrides) by the implementation.

## [](#_description)Description

Depending on the used video encode operation, additional codec-specific structures **can** be included in the `pNext` chain of this structure to capture codec-specific feedback information about the requested parameter data, as described in the corresponding sections.

Valid Usage (Implicit)

- [](#VUID-VkVideoEncodeSessionParametersFeedbackInfoKHR-sType-sType)VUID-VkVideoEncodeSessionParametersFeedbackInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_ENCODE_SESSION_PARAMETERS_FEEDBACK_INFO_KHR`
- [](#VUID-VkVideoEncodeSessionParametersFeedbackInfoKHR-pNext-pNext)VUID-VkVideoEncodeSessionParametersFeedbackInfoKHR-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the `pNext` chain **must** be either `NULL` or a pointer to a valid instance of [VkVideoEncodeH264SessionParametersFeedbackInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264SessionParametersFeedbackInfoKHR.html) or [VkVideoEncodeH265SessionParametersFeedbackInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265SessionParametersFeedbackInfoKHR.html)
- [](#VUID-VkVideoEncodeSessionParametersFeedbackInfoKHR-sType-unique)VUID-VkVideoEncodeSessionParametersFeedbackInfoKHR-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique

## [](#_see_also)See Also

[VK\_KHR\_video\_encode\_queue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_queue.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetEncodedVideoSessionParametersKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetEncodedVideoSessionParametersKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoEncodeSessionParametersFeedbackInfoKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0