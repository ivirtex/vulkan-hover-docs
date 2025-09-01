# VkVideoEncodeUsageInfoKHR(3) Manual Page

## Name

VkVideoEncodeUsageInfoKHR - Structure specifying video encode usage information



## [](#_c_specification)C Specification

Additional information about the video encode use case **can** be provided by adding a `VkVideoEncodeUsageInfoKHR` structure to the `pNext` chain of [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoProfileInfoKHR.html).

The `VkVideoEncodeUsageInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_video_encode_queue
typedef struct VkVideoEncodeUsageInfoKHR {
    VkStructureType                 sType;
    const void*                     pNext;
    VkVideoEncodeUsageFlagsKHR      videoUsageHints;
    VkVideoEncodeContentFlagsKHR    videoContentHints;
    VkVideoEncodeTuningModeKHR      tuningMode;
} VkVideoEncodeUsageInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `videoUsageHints` is a bitmask of [VkVideoEncodeUsageFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeUsageFlagBitsKHR.html) specifying hints about the intended use of the video encode profile.
- `videoContentHints` is a bitmask of [VkVideoEncodeContentFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeContentFlagBitsKHR.html) specifying hints about the content to be encoded using the video encode profile.
- `tuningMode` is a [VkVideoEncodeTuningModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeTuningModeKHR.html) value specifying the tuning mode to use when encoding with the video profile.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkVideoEncodeUsageInfoKHR-sType-sType)VUID-VkVideoEncodeUsageInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_ENCODE_USAGE_INFO_KHR`
- [](#VUID-VkVideoEncodeUsageInfoKHR-videoUsageHints-parameter)VUID-VkVideoEncodeUsageInfoKHR-videoUsageHints-parameter  
  `videoUsageHints` **must** be a valid combination of [VkVideoEncodeUsageFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeUsageFlagBitsKHR.html) values
- [](#VUID-VkVideoEncodeUsageInfoKHR-videoContentHints-parameter)VUID-VkVideoEncodeUsageInfoKHR-videoContentHints-parameter  
  `videoContentHints` **must** be a valid combination of [VkVideoEncodeContentFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeContentFlagBitsKHR.html) values
- [](#VUID-VkVideoEncodeUsageInfoKHR-tuningMode-parameter)VUID-VkVideoEncodeUsageInfoKHR-tuningMode-parameter  
  If `tuningMode` is not `0`, `tuningMode` **must** be a valid [VkVideoEncodeTuningModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeTuningModeKHR.html) value

## [](#_see_also)See Also

[VK\_KHR\_video\_encode\_queue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_queue.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkVideoEncodeContentFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeContentFlagsKHR.html), [VkVideoEncodeTuningModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeTuningModeKHR.html), [VkVideoEncodeUsageFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeUsageFlagsKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoEncodeUsageInfoKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0