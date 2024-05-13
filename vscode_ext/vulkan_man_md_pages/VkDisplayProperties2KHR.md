# VkDisplayProperties2KHR(3) Manual Page

## Name

VkDisplayProperties2KHR - Structure describing an available display
device



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkDisplayProperties2KHR` structure is defined as:

``` c
// Provided by VK_KHR_get_display_properties2
typedef struct VkDisplayProperties2KHR {
    VkStructureType           sType;
    void*                     pNext;
    VkDisplayPropertiesKHR    displayProperties;
} VkDisplayProperties2KHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `displayProperties` is a
  [VkDisplayPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayPropertiesKHR.html) structure.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkDisplayProperties2KHR-sType-sType"
  id="VUID-VkDisplayProperties2KHR-sType-sType"></a>
  VUID-VkDisplayProperties2KHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DISPLAY_PROPERTIES_2_KHR`

- <a href="#VUID-VkDisplayProperties2KHR-pNext-pNext"
  id="VUID-VkDisplayProperties2KHR-pNext-pNext"></a>
  VUID-VkDisplayProperties2KHR-pNext-pNext  
  `pNext` **must** be `NULL`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_get_display_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_display_properties2.html),
[VkDisplayPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayPropertiesKHR.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetPhysicalDeviceDisplayProperties2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceDisplayProperties2KHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDisplayProperties2KHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
