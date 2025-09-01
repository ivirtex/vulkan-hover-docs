# VkVideoDecodeUsageInfoKHR(3) Manual Page

## Name

VkVideoDecodeUsageInfoKHR - Structure specifying video decode usage information



## [](#_c_specification)C Specification

Additional information about the video decode use case **can** be provided by adding a `VkVideoDecodeUsageInfoKHR` structure to the `pNext` chain of [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoProfileInfoKHR.html).

The `VkVideoDecodeUsageInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_video_decode_queue
typedef struct VkVideoDecodeUsageInfoKHR {
    VkStructureType               sType;
    const void*                   pNext;
    VkVideoDecodeUsageFlagsKHR    videoUsageHints;
} VkVideoDecodeUsageInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `videoUsageHints` is a bitmask of [VkVideoDecodeUsageFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeUsageFlagBitsKHR.html) specifying hints about the intended use of the video decode profile.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkVideoDecodeUsageInfoKHR-sType-sType)VUID-VkVideoDecodeUsageInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_DECODE_USAGE_INFO_KHR`
- [](#VUID-VkVideoDecodeUsageInfoKHR-videoUsageHints-parameter)VUID-VkVideoDecodeUsageInfoKHR-videoUsageHints-parameter  
  `videoUsageHints` **must** be a valid combination of [VkVideoDecodeUsageFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeUsageFlagBitsKHR.html) values

## [](#_see_also)See Also

[VK\_KHR\_video\_decode\_queue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_decode_queue.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkVideoDecodeUsageFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeUsageFlagsKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoDecodeUsageInfoKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0