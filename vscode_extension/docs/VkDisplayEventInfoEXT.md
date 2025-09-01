# VkDisplayEventInfoEXT(3) Manual Page

## Name

VkDisplayEventInfoEXT - Describe a display event to create



## [](#_c_specification)C Specification

The `VkDisplayEventInfoEXT` structure is defined as:

```c++
// Provided by VK_EXT_display_control
typedef struct VkDisplayEventInfoEXT {
    VkStructureType          sType;
    const void*              pNext;
    VkDisplayEventTypeEXT    displayEvent;
} VkDisplayEventInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `displayEvent` is a [VkDisplayEventTypeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayEventTypeEXT.html) specifying when the fence will be signaled.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkDisplayEventInfoEXT-sType-sType)VUID-VkDisplayEventInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DISPLAY_EVENT_INFO_EXT`
- [](#VUID-VkDisplayEventInfoEXT-pNext-pNext)VUID-VkDisplayEventInfoEXT-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkDisplayEventInfoEXT-displayEvent-parameter)VUID-VkDisplayEventInfoEXT-displayEvent-parameter  
  `displayEvent` **must** be a valid [VkDisplayEventTypeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayEventTypeEXT.html) value

## [](#_see_also)See Also

[VK\_EXT\_display\_control](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_display_control.html), [VkDisplayEventTypeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayEventTypeEXT.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkRegisterDisplayEventEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkRegisterDisplayEventEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDisplayEventInfoEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0