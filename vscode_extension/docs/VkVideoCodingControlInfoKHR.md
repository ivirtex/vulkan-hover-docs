# VkVideoCodingControlInfoKHR(3) Manual Page

## Name

VkVideoCodingControlInfoKHR - Structure specifying video coding control
parameters



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkVideoCodingControlInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_video_queue
typedef struct VkVideoCodingControlInfoKHR {
    VkStructureType                 sType;
    const void*                     pNext;
    VkVideoCodingControlFlagsKHR    flags;
} VkVideoCodingControlInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is a bitmask of
  [VkVideoCodingControlFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCodingControlFlagsKHR.html)
  specifying control flags.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkVideoCodingControlInfoKHR-flags-07018"
  id="VUID-VkVideoCodingControlInfoKHR-flags-07018"></a>
  VUID-VkVideoCodingControlInfoKHR-flags-07018  
  If `flags` includes
  `VK_VIDEO_CODING_CONTROL_ENCODE_RATE_CONTROL_BIT_KHR`, then the
  `pNext` chain **must** include a
  [VkVideoEncodeRateControlInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeRateControlInfoKHR.html)
  structure

- <a href="#VUID-VkVideoCodingControlInfoKHR-flags-08349"
  id="VUID-VkVideoCodingControlInfoKHR-flags-08349"></a>
  VUID-VkVideoCodingControlInfoKHR-flags-08349  
  If `flags` includes
  `VK_VIDEO_CODING_CONTROL_ENCODE_QUALITY_LEVEL_BIT_KHR`, then the
  `pNext` chain **must** include a
  [VkVideoEncodeQualityLevelInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeQualityLevelInfoKHR.html)
  structure

Valid Usage (Implicit)

- <a href="#VUID-VkVideoCodingControlInfoKHR-sType-sType"
  id="VUID-VkVideoCodingControlInfoKHR-sType-sType"></a>
  VUID-VkVideoCodingControlInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_CODING_CONTROL_INFO_KHR`

- <a href="#VUID-VkVideoCodingControlInfoKHR-pNext-pNext"
  id="VUID-VkVideoCodingControlInfoKHR-pNext-pNext"></a>
  VUID-VkVideoCodingControlInfoKHR-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the
  `pNext` chain **must** be either `NULL` or a pointer to a valid
  instance of
  [VkVideoEncodeH264RateControlInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264RateControlInfoKHR.html),
  [VkVideoEncodeH265RateControlInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265RateControlInfoKHR.html),
  [VkVideoEncodeQualityLevelInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeQualityLevelInfoKHR.html),
  or
  [VkVideoEncodeRateControlInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeRateControlInfoKHR.html)

- <a href="#VUID-VkVideoCodingControlInfoKHR-sType-unique"
  id="VUID-VkVideoCodingControlInfoKHR-sType-unique"></a>
  VUID-VkVideoCodingControlInfoKHR-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

- <a href="#VUID-VkVideoCodingControlInfoKHR-flags-parameter"
  id="VUID-VkVideoCodingControlInfoKHR-flags-parameter"></a>
  VUID-VkVideoCodingControlInfoKHR-flags-parameter  
  `flags` **must** be a valid combination of
  [VkVideoCodingControlFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCodingControlFlagBitsKHR.html)
  values

- <a href="#VUID-VkVideoCodingControlInfoKHR-flags-requiredbitmask"
  id="VUID-VkVideoCodingControlInfoKHR-flags-requiredbitmask"></a>
  VUID-VkVideoCodingControlInfoKHR-flags-requiredbitmask  
  `flags` **must** not be `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_queue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_queue.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkVideoCodingControlFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCodingControlFlagsKHR.html),
[vkCmdControlVideoCodingKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdControlVideoCodingKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoCodingControlInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
