# VkVideoDecodeUsageInfoKHR(3) Manual Page

## Name

VkVideoDecodeUsageInfoKHR - Structure specifying video decode usage
information



## <a href="#_c_specification" class="anchor"></a>C Specification

Additional information about the video decode use case **can** be
provided by adding a `VkVideoDecodeUsageInfoKHR` structure to the
`pNext` chain of [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileInfoKHR.html).

The `VkVideoDecodeUsageInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_video_decode_queue
typedef struct VkVideoDecodeUsageInfoKHR {
    VkStructureType               sType;
    const void*                   pNext;
    VkVideoDecodeUsageFlagsKHR    videoUsageHints;
} VkVideoDecodeUsageInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `videoUsageHints` is a bitmask of
  [VkVideoDecodeUsageFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeUsageFlagBitsKHR.html)
  specifying hints about the intended use of the video decode profile.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkVideoDecodeUsageInfoKHR-sType-sType"
  id="VUID-VkVideoDecodeUsageInfoKHR-sType-sType"></a>
  VUID-VkVideoDecodeUsageInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_DECODE_USAGE_INFO_KHR`

- <a href="#VUID-VkVideoDecodeUsageInfoKHR-videoUsageHints-parameter"
  id="VUID-VkVideoDecodeUsageInfoKHR-videoUsageHints-parameter"></a>
  VUID-VkVideoDecodeUsageInfoKHR-videoUsageHints-parameter  
  `videoUsageHints` **must** be a valid combination of
  [VkVideoDecodeUsageFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeUsageFlagBitsKHR.html)
  values

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_decode_queue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_decode_queue.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkVideoDecodeUsageFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeUsageFlagsKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoDecodeUsageInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
