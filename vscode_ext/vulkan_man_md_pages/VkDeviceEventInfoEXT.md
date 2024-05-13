# VkDeviceEventInfoEXT(3) Manual Page

## Name

VkDeviceEventInfoEXT - Describe a device event to create



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkDeviceEventInfoEXT` structure is defined as:

``` c
// Provided by VK_EXT_display_control
typedef struct VkDeviceEventInfoEXT {
    VkStructureType         sType;
    const void*             pNext;
    VkDeviceEventTypeEXT    deviceEvent;
} VkDeviceEventInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `device` is a [VkDeviceEventTypeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceEventTypeEXT.html) value
  specifying when the fence will be signaled.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkDeviceEventInfoEXT-sType-sType"
  id="VUID-VkDeviceEventInfoEXT-sType-sType"></a>
  VUID-VkDeviceEventInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DEVICE_EVENT_INFO_EXT`

- <a href="#VUID-VkDeviceEventInfoEXT-pNext-pNext"
  id="VUID-VkDeviceEventInfoEXT-pNext-pNext"></a>
  VUID-VkDeviceEventInfoEXT-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkDeviceEventInfoEXT-deviceEvent-parameter"
  id="VUID-VkDeviceEventInfoEXT-deviceEvent-parameter"></a>
  VUID-VkDeviceEventInfoEXT-deviceEvent-parameter  
  `deviceEvent` **must** be a valid
  [VkDeviceEventTypeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceEventTypeEXT.html) value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_display_control](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_display_control.html),
[VkDeviceEventTypeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceEventTypeEXT.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkRegisterDeviceEventEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkRegisterDeviceEventEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDeviceEventInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
