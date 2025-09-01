# VkVideoCodingControlInfoKHR(3) Manual Page

## Name

VkVideoCodingControlInfoKHR - Structure specifying video coding control parameters



## [](#_c_specification)C Specification

The `VkVideoCodingControlInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_video_queue
typedef struct VkVideoCodingControlInfoKHR {
    VkStructureType                 sType;
    const void*                     pNext;
    VkVideoCodingControlFlagsKHR    flags;
} VkVideoCodingControlInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is a bitmask of [VkVideoCodingControlFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoCodingControlFlagsKHR.html) specifying control flags.

## [](#_description)Description

Valid Usage

- [](#VUID-VkVideoCodingControlInfoKHR-flags-07018)VUID-VkVideoCodingControlInfoKHR-flags-07018  
  If `flags` includes `VK_VIDEO_CODING_CONTROL_ENCODE_RATE_CONTROL_BIT_KHR`, then the `pNext` chain **must** include a [VkVideoEncodeRateControlInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeRateControlInfoKHR.html) structure
- [](#VUID-VkVideoCodingControlInfoKHR-flags-08349)VUID-VkVideoCodingControlInfoKHR-flags-08349  
  If `flags` includes `VK_VIDEO_CODING_CONTROL_ENCODE_QUALITY_LEVEL_BIT_KHR`, then the `pNext` chain **must** include a [VkVideoEncodeQualityLevelInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeQualityLevelInfoKHR.html) structure

Valid Usage (Implicit)

- [](#VUID-VkVideoCodingControlInfoKHR-sType-sType)VUID-VkVideoCodingControlInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_CODING_CONTROL_INFO_KHR`
- [](#VUID-VkVideoCodingControlInfoKHR-pNext-pNext)VUID-VkVideoCodingControlInfoKHR-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the `pNext` chain **must** be either `NULL` or a pointer to a valid instance of [VkVideoEncodeAV1RateControlInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1RateControlInfoKHR.html), [VkVideoEncodeH264RateControlInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264RateControlInfoKHR.html), [VkVideoEncodeH265RateControlInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265RateControlInfoKHR.html), [VkVideoEncodeQualityLevelInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeQualityLevelInfoKHR.html), or [VkVideoEncodeRateControlInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeRateControlInfoKHR.html)
- [](#VUID-VkVideoCodingControlInfoKHR-sType-unique)VUID-VkVideoCodingControlInfoKHR-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique
- [](#VUID-VkVideoCodingControlInfoKHR-flags-parameter)VUID-VkVideoCodingControlInfoKHR-flags-parameter  
  `flags` **must** be a valid combination of [VkVideoCodingControlFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoCodingControlFlagBitsKHR.html) values
- [](#VUID-VkVideoCodingControlInfoKHR-flags-requiredbitmask)VUID-VkVideoCodingControlInfoKHR-flags-requiredbitmask  
  `flags` **must** not be `0`

## [](#_see_also)See Also

[VK\_KHR\_video\_queue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_queue.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkVideoCodingControlFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoCodingControlFlagsKHR.html), [vkCmdControlVideoCodingKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdControlVideoCodingKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoCodingControlInfoKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0