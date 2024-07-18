# VkVideoEndCodingInfoKHR(3) Manual Page

## Name

VkVideoEndCodingInfoKHR - Structure specifying video coding scope end
information



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkVideoEndCodingInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_video_queue
typedef struct VkVideoEndCodingInfoKHR {
    VkStructureType             sType;
    const void*                 pNext;
    VkVideoEndCodingFlagsKHR    flags;
} VkVideoEndCodingInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is reserved for future use.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkVideoEndCodingInfoKHR-sType-sType"
  id="VUID-VkVideoEndCodingInfoKHR-sType-sType"></a>
  VUID-VkVideoEndCodingInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_END_CODING_INFO_KHR`

- <a href="#VUID-VkVideoEndCodingInfoKHR-pNext-pNext"
  id="VUID-VkVideoEndCodingInfoKHR-pNext-pNext"></a>
  VUID-VkVideoEndCodingInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkVideoEndCodingInfoKHR-flags-zerobitmask"
  id="VUID-VkVideoEndCodingInfoKHR-flags-zerobitmask"></a>
  VUID-VkVideoEndCodingInfoKHR-flags-zerobitmask  
  `flags` **must** be `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_queue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_queue.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkVideoEndCodingFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEndCodingFlagsKHR.html),
[vkCmdEndVideoCodingKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdEndVideoCodingKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoEndCodingInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
