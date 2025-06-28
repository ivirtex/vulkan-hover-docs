# VkQueueFamilyVideoPropertiesKHR(3) Manual Page

## Name

VkQueueFamilyVideoPropertiesKHR - Structure describing video codec operations supported by a queue family



## [](#_c_specification)C Specification

The [VkQueueFamilyVideoPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueFamilyVideoPropertiesKHR.html) structure is defined as:

```c++
// Provided by VK_KHR_video_queue
typedef struct VkQueueFamilyVideoPropertiesKHR {
    VkStructureType                  sType;
    void*                            pNext;
    VkVideoCodecOperationFlagsKHR    videoCodecOperations;
} VkQueueFamilyVideoPropertiesKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `videoCodecOperations` is a bitmask of [VkVideoCodecOperationFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoCodecOperationFlagBitsKHR.html) that indicates the set of video codec operations supported by the queue family.

## [](#_description)Description

If this structure is included in the `pNext` chain of the [VkQueueFamilyProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueFamilyProperties2.html) structure passed to [vkGetPhysicalDeviceQueueFamilyProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceQueueFamilyProperties2.html), then it is filled with the set of video codec operations supported by the specified queue family.

Valid Usage (Implicit)

- [](#VUID-VkQueueFamilyVideoPropertiesKHR-sType-sType)VUID-VkQueueFamilyVideoPropertiesKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_QUEUE_FAMILY_VIDEO_PROPERTIES_KHR`

## [](#_see_also)See Also

[VK\_KHR\_video\_queue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_queue.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkVideoCodecOperationFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoCodecOperationFlagsKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkQueueFamilyVideoPropertiesKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0