# VkPhysicalDeviceVideoFormatInfoKHR(3) Manual Page

## Name

VkPhysicalDeviceVideoFormatInfoKHR - Structure specifying the codec
video format



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceVideoFormatInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_video_queue
typedef struct VkPhysicalDeviceVideoFormatInfoKHR {
    VkStructureType      sType;
    const void*          pNext;
    VkImageUsageFlags    imageUsage;
} VkPhysicalDeviceVideoFormatInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `imageUsage` is a bitmask of
  [VkImageUsageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageUsageFlagBits.html) specifying the
  intended usage of the video images.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceVideoFormatInfoKHR-sType-sType"
  id="VUID-VkPhysicalDeviceVideoFormatInfoKHR-sType-sType"></a>
  VUID-VkPhysicalDeviceVideoFormatInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VIDEO_FORMAT_INFO_KHR`

- <a href="#VUID-VkPhysicalDeviceVideoFormatInfoKHR-pNext-pNext"
  id="VUID-VkPhysicalDeviceVideoFormatInfoKHR-pNext-pNext"></a>
  VUID-VkPhysicalDeviceVideoFormatInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL` or a pointer to a valid instance of
  [VkVideoProfileListInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileListInfoKHR.html)

- <a href="#VUID-VkPhysicalDeviceVideoFormatInfoKHR-sType-unique"
  id="VUID-VkPhysicalDeviceVideoFormatInfoKHR-sType-unique"></a>
  VUID-VkPhysicalDeviceVideoFormatInfoKHR-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

- <a href="#VUID-VkPhysicalDeviceVideoFormatInfoKHR-imageUsage-parameter"
  id="VUID-VkPhysicalDeviceVideoFormatInfoKHR-imageUsage-parameter"></a>
  VUID-VkPhysicalDeviceVideoFormatInfoKHR-imageUsage-parameter  
  `imageUsage` **must** be a valid combination of
  [VkImageUsageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageUsageFlagBits.html) values

- <a
  href="#VUID-VkPhysicalDeviceVideoFormatInfoKHR-imageUsage-requiredbitmask"
  id="VUID-VkPhysicalDeviceVideoFormatInfoKHR-imageUsage-requiredbitmask"></a>
  VUID-VkPhysicalDeviceVideoFormatInfoKHR-imageUsage-requiredbitmask  
  `imageUsage` **must** not be `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_queue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_queue.html),
[VkImageUsageFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageUsageFlags.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetPhysicalDeviceVideoFormatPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoFormatPropertiesKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceVideoFormatInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
