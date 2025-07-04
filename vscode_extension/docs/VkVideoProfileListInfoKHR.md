# VkVideoProfileListInfoKHR(3) Manual Page

## Name

VkVideoProfileListInfoKHR - Structure specifying one or more video profiles used in conjunction



## [](#_c_specification)C Specification

The `VkVideoProfileListInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_video_queue
typedef struct VkVideoProfileListInfoKHR {
    VkStructureType                 sType;
    const void*                     pNext;
    uint32_t                        profileCount;
    const VkVideoProfileInfoKHR*    pProfiles;
} VkVideoProfileListInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `profileCount` is the number of elements in the `pProfiles` array.
- `pProfiles` is a pointer to an array of [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoProfileInfoKHR.html) structures.

## [](#_description)Description

Note

Video transcoding is an example of a use case that necessitates the specification of multiple profiles in various contexts.

When the application provides a video decode profile and one or more video encode profiles in the profile list, the implementation ensures that any capabilitities returned or resources created are suitable for the video transcoding use cases without the need for manual data transformations.

Valid Usage

- [](#VUID-VkVideoProfileListInfoKHR-pProfiles-06813)VUID-VkVideoProfileListInfoKHR-pProfiles-06813  
  `pProfiles` **must** not contain more than one element whose `videoCodecOperation` member specifies a decode operation

Valid Usage (Implicit)

- [](#VUID-VkVideoProfileListInfoKHR-sType-sType)VUID-VkVideoProfileListInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_PROFILE_LIST_INFO_KHR`
- [](#VUID-VkVideoProfileListInfoKHR-pProfiles-parameter)VUID-VkVideoProfileListInfoKHR-pProfiles-parameter  
  If `profileCount` is not `0`, `pProfiles` **must** be a valid pointer to an array of `profileCount` valid [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoProfileInfoKHR.html) structures

## [](#_see_also)See Also

[VK\_KHR\_video\_queue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_queue.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoProfileInfoKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoProfileListInfoKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0