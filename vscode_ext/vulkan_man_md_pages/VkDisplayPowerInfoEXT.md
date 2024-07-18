# VkDisplayPowerInfoEXT(3) Manual Page

## Name

VkDisplayPowerInfoEXT - Describe the power state of a display



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkDisplayPowerInfoEXT` structure is defined as:

``` c
// Provided by VK_EXT_display_control
typedef struct VkDisplayPowerInfoEXT {
    VkStructureType           sType;
    const void*               pNext;
    VkDisplayPowerStateEXT    powerState;
} VkDisplayPowerInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `powerState` is a
  [VkDisplayPowerStateEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayPowerStateEXT.html) value specifying
  the new power state of the display.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkDisplayPowerInfoEXT-sType-sType"
  id="VUID-VkDisplayPowerInfoEXT-sType-sType"></a>
  VUID-VkDisplayPowerInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DISPLAY_POWER_INFO_EXT`

- <a href="#VUID-VkDisplayPowerInfoEXT-pNext-pNext"
  id="VUID-VkDisplayPowerInfoEXT-pNext-pNext"></a>
  VUID-VkDisplayPowerInfoEXT-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkDisplayPowerInfoEXT-powerState-parameter"
  id="VUID-VkDisplayPowerInfoEXT-powerState-parameter"></a>
  VUID-VkDisplayPowerInfoEXT-powerState-parameter  
  `powerState` **must** be a valid
  [VkDisplayPowerStateEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayPowerStateEXT.html) value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_display_control](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_display_control.html),
[VkDisplayPowerStateEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayPowerStateEXT.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkDisplayPowerControlEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDisplayPowerControlEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDisplayPowerInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
