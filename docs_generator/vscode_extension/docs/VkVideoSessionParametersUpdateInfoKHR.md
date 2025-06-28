# VkVideoSessionParametersUpdateInfoKHR(3) Manual Page

## Name

VkVideoSessionParametersUpdateInfoKHR - Structure specifying video session parameters update information



## [](#_c_specification)C Specification

The `VkVideoSessionParametersUpdateInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_video_queue
typedef struct VkVideoSessionParametersUpdateInfoKHR {
    VkStructureType    sType;
    const void*        pNext;
    uint32_t           updateSequenceCount;
} VkVideoSessionParametersUpdateInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `updateSequenceCount` is the new [update sequence count](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#video-session-parameters) to set for the video session parameters object.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkVideoSessionParametersUpdateInfoKHR-sType-sType)VUID-VkVideoSessionParametersUpdateInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_SESSION_PARAMETERS_UPDATE_INFO_KHR`
- [](#VUID-VkVideoSessionParametersUpdateInfoKHR-pNext-pNext)VUID-VkVideoSessionParametersUpdateInfoKHR-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the `pNext` chain **must** be either `NULL` or a pointer to a valid instance of [VkVideoDecodeH264SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH264SessionParametersAddInfoKHR.html), [VkVideoDecodeH265SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH265SessionParametersAddInfoKHR.html), [VkVideoEncodeH264SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264SessionParametersAddInfoKHR.html), or [VkVideoEncodeH265SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265SessionParametersAddInfoKHR.html)
- [](#VUID-VkVideoSessionParametersUpdateInfoKHR-sType-unique)VUID-VkVideoSessionParametersUpdateInfoKHR-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique

## [](#_see_also)See Also

[VK\_KHR\_video\_queue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_queue.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkUpdateVideoSessionParametersKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkUpdateVideoSessionParametersKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoSessionParametersUpdateInfoKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0