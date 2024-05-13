# VkDisplayEventInfoEXT(3) Manual Page

## Name

VkDisplayEventInfoEXT - Describe a display event to create



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkDisplayEventInfoEXT` structure is defined as:

``` c
// Provided by VK_EXT_display_control
typedef struct VkDisplayEventInfoEXT {
    VkStructureType          sType;
    const void*              pNext;
    VkDisplayEventTypeEXT    displayEvent;
} VkDisplayEventInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `displayEvent` is a
  [VkDisplayEventTypeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayEventTypeEXT.html) specifying when
  the fence will be signaled.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkDisplayEventInfoEXT-sType-sType"
  id="VUID-VkDisplayEventInfoEXT-sType-sType"></a>
  VUID-VkDisplayEventInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DISPLAY_EVENT_INFO_EXT`

- <a href="#VUID-VkDisplayEventInfoEXT-pNext-pNext"
  id="VUID-VkDisplayEventInfoEXT-pNext-pNext"></a>
  VUID-VkDisplayEventInfoEXT-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkDisplayEventInfoEXT-displayEvent-parameter"
  id="VUID-VkDisplayEventInfoEXT-displayEvent-parameter"></a>
  VUID-VkDisplayEventInfoEXT-displayEvent-parameter  
  `displayEvent` **must** be a valid
  [VkDisplayEventTypeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayEventTypeEXT.html) value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_display_control](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_display_control.html),
[VkDisplayEventTypeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayEventTypeEXT.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkRegisterDisplayEventEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkRegisterDisplayEventEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDisplayEventInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
