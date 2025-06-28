# VkDisplayPowerInfoEXT(3) Manual Page

## Name

VkDisplayPowerInfoEXT - Describe the power state of a display



## [](#_c_specification)C Specification

The `VkDisplayPowerInfoEXT` structure is defined as:

```c++
// Provided by VK_EXT_display_control
typedef struct VkDisplayPowerInfoEXT {
    VkStructureType           sType;
    const void*               pNext;
    VkDisplayPowerStateEXT    powerState;
} VkDisplayPowerInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `powerState` is a [VkDisplayPowerStateEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayPowerStateEXT.html) value specifying the new power state of the display.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkDisplayPowerInfoEXT-sType-sType)VUID-VkDisplayPowerInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DISPLAY_POWER_INFO_EXT`
- [](#VUID-VkDisplayPowerInfoEXT-pNext-pNext)VUID-VkDisplayPowerInfoEXT-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkDisplayPowerInfoEXT-powerState-parameter)VUID-VkDisplayPowerInfoEXT-powerState-parameter  
  `powerState` **must** be a valid [VkDisplayPowerStateEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayPowerStateEXT.html) value

## [](#_see_also)See Also

[VK\_EXT\_display\_control](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_display_control.html), [VkDisplayPowerStateEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayPowerStateEXT.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkDisplayPowerControlEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkDisplayPowerControlEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDisplayPowerInfoEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0