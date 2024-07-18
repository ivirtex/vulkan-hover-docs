# VkQueueFamilyVideoPropertiesKHR(3) Manual Page

## Name

VkQueueFamilyVideoPropertiesKHR - Structure describing video codec
operations supported by a queue family



## <a href="#_c_specification" class="anchor"></a>C Specification

The
[VkQueueFamilyVideoPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueueFamilyVideoPropertiesKHR.html)
structure is defined as:

``` c
// Provided by VK_KHR_video_queue
typedef struct VkQueueFamilyVideoPropertiesKHR {
    VkStructureType                  sType;
    void*                            pNext;
    VkVideoCodecOperationFlagsKHR    videoCodecOperations;
} VkQueueFamilyVideoPropertiesKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `videoCodecOperations` is a bitmask of
  [VkVideoCodecOperationFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCodecOperationFlagBitsKHR.html)
  that indicates the set of video codec operations supported by the
  queue family.

## <a href="#_description" class="anchor"></a>Description

If this structure is included in the `pNext` chain of the
[VkQueueFamilyProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueueFamilyProperties2.html) structure
passed to
[vkGetPhysicalDeviceQueueFamilyProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceQueueFamilyProperties2.html),
then it is filled with the set of video codec operations supported by
the specified queue family.

Valid Usage (Implicit)

- <a href="#VUID-VkQueueFamilyVideoPropertiesKHR-sType-sType"
  id="VUID-VkQueueFamilyVideoPropertiesKHR-sType-sType"></a>
  VUID-VkQueueFamilyVideoPropertiesKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_QUEUE_FAMILY_VIDEO_PROPERTIES_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_queue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_queue.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkVideoCodecOperationFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCodecOperationFlagsKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkQueueFamilyVideoPropertiesKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
