# VkVideoEndCodingInfoKHR(3) Manual Page

## Name

VkVideoEndCodingInfoKHR - Structure specifying video coding scope end information



## [](#_c_specification)C Specification

The `VkVideoEndCodingInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_video_queue
typedef struct VkVideoEndCodingInfoKHR {
    VkStructureType             sType;
    const void*                 pNext;
    VkVideoEndCodingFlagsKHR    flags;
} VkVideoEndCodingInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is reserved for future use.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkVideoEndCodingInfoKHR-sType-sType)VUID-VkVideoEndCodingInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_END_CODING_INFO_KHR`
- [](#VUID-VkVideoEndCodingInfoKHR-pNext-pNext)VUID-VkVideoEndCodingInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkVideoEndCodingInfoKHR-flags-zerobitmask)VUID-VkVideoEndCodingInfoKHR-flags-zerobitmask  
  `flags` **must** be `0`

## [](#_see_also)See Also

[VK\_KHR\_video\_queue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_queue.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkVideoEndCodingFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEndCodingFlagsKHR.html), [vkCmdEndVideoCodingKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdEndVideoCodingKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoEndCodingInfoKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0