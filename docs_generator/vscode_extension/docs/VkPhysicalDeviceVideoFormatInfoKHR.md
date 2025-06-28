# VkPhysicalDeviceVideoFormatInfoKHR(3) Manual Page

## Name

VkPhysicalDeviceVideoFormatInfoKHR - Structure specifying the codec video format



## [](#_c_specification)C Specification

The `VkPhysicalDeviceVideoFormatInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_video_queue
typedef struct VkPhysicalDeviceVideoFormatInfoKHR {
    VkStructureType      sType;
    const void*          pNext;
    VkImageUsageFlags    imageUsage;
} VkPhysicalDeviceVideoFormatInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `imageUsage` is a bitmask of [VkImageUsageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageUsageFlagBits.html) specifying the intended usage of the video images.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceVideoFormatInfoKHR-sType-sType)VUID-VkPhysicalDeviceVideoFormatInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VIDEO_FORMAT_INFO_KHR`
- [](#VUID-VkPhysicalDeviceVideoFormatInfoKHR-pNext-pNext)VUID-VkPhysicalDeviceVideoFormatInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL` or a pointer to a valid instance of [VkVideoProfileListInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoProfileListInfoKHR.html)
- [](#VUID-VkPhysicalDeviceVideoFormatInfoKHR-sType-unique)VUID-VkPhysicalDeviceVideoFormatInfoKHR-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique
- [](#VUID-VkPhysicalDeviceVideoFormatInfoKHR-imageUsage-parameter)VUID-VkPhysicalDeviceVideoFormatInfoKHR-imageUsage-parameter  
  `imageUsage` **must** be a valid combination of [VkImageUsageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageUsageFlagBits.html) values
- [](#VUID-VkPhysicalDeviceVideoFormatInfoKHR-imageUsage-requiredbitmask)VUID-VkPhysicalDeviceVideoFormatInfoKHR-imageUsage-requiredbitmask  
  `imageUsage` **must** not be `0`

## [](#_see_also)See Also

[VK\_KHR\_video\_queue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_queue.html), [VkImageUsageFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageUsageFlags.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetPhysicalDeviceVideoFormatPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoFormatPropertiesKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceVideoFormatInfoKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0