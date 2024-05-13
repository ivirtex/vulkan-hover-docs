# VkDisplayPlaneProperties2KHR(3) Manual Page

## Name

VkDisplayPlaneProperties2KHR - Structure describing an available display
plane



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkDisplayPlaneProperties2KHR` structure is defined as:

``` c
// Provided by VK_KHR_get_display_properties2
typedef struct VkDisplayPlaneProperties2KHR {
    VkStructureType                sType;
    void*                          pNext;
    VkDisplayPlanePropertiesKHR    displayPlaneProperties;
} VkDisplayPlaneProperties2KHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `displayPlaneProperties` is a
  [VkDisplayPlanePropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayPlanePropertiesKHR.html)
  structure.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkDisplayPlaneProperties2KHR-sType-sType"
  id="VUID-VkDisplayPlaneProperties2KHR-sType-sType"></a>
  VUID-VkDisplayPlaneProperties2KHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DISPLAY_PLANE_PROPERTIES_2_KHR`

- <a href="#VUID-VkDisplayPlaneProperties2KHR-pNext-pNext"
  id="VUID-VkDisplayPlaneProperties2KHR-pNext-pNext"></a>
  VUID-VkDisplayPlaneProperties2KHR-pNext-pNext  
  `pNext` **must** be `NULL`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_get_display_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_display_properties2.html),
[VkDisplayPlanePropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayPlanePropertiesKHR.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetPhysicalDeviceDisplayPlaneProperties2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceDisplayPlaneProperties2KHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDisplayPlaneProperties2KHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
