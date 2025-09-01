# VkVideoEncodeSessionParametersGetInfoKHR(3) Manual Page

## Name

VkVideoEncodeSessionParametersGetInfoKHR - Structure specifying parameters for retrieving encoded video session parameter data



## [](#_c_specification)C Specification

The `VkVideoEncodeSessionParametersGetInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_video_encode_queue
typedef struct VkVideoEncodeSessionParametersGetInfoKHR {
    VkStructureType                sType;
    const void*                    pNext;
    VkVideoSessionParametersKHR    videoSessionParameters;
} VkVideoEncodeSessionParametersGetInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `videoSessionParameters` is the [VkVideoSessionParametersKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionParametersKHR.html) object to retrieve encoded parameter data from.

## [](#_description)Description

Depending on the used video encode operation, additional codec-specific structures **may** need to be included in the `pNext` chain of this structure to identify the specific video session parameters to retrieve encoded parameter data for, as described in the corresponding sections.

Valid Usage (Implicit)

- [](#VUID-VkVideoEncodeSessionParametersGetInfoKHR-sType-sType)VUID-VkVideoEncodeSessionParametersGetInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_ENCODE_SESSION_PARAMETERS_GET_INFO_KHR`
- [](#VUID-VkVideoEncodeSessionParametersGetInfoKHR-pNext-pNext)VUID-VkVideoEncodeSessionParametersGetInfoKHR-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the `pNext` chain **must** be either `NULL` or a pointer to a valid instance of [VkVideoEncodeH264SessionParametersGetInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264SessionParametersGetInfoKHR.html) or [VkVideoEncodeH265SessionParametersGetInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265SessionParametersGetInfoKHR.html)
- [](#VUID-VkVideoEncodeSessionParametersGetInfoKHR-sType-unique)VUID-VkVideoEncodeSessionParametersGetInfoKHR-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique
- [](#VUID-VkVideoEncodeSessionParametersGetInfoKHR-videoSessionParameters-parameter)VUID-VkVideoEncodeSessionParametersGetInfoKHR-videoSessionParameters-parameter  
  `videoSessionParameters` **must** be a valid [VkVideoSessionParametersKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionParametersKHR.html) handle

## [](#_see_also)See Also

[VK\_KHR\_video\_encode\_queue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_queue.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkVideoSessionParametersKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionParametersKHR.html), [vkGetEncodedVideoSessionParametersKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetEncodedVideoSessionParametersKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoEncodeSessionParametersGetInfoKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0