# VkPhysicalDeviceCustomBorderColorPropertiesEXT(3) Manual Page

## Name

VkPhysicalDeviceCustomBorderColorPropertiesEXT - Structure describing
whether custom border colors can be supported by an implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceCustomBorderColorPropertiesEXT` structure is
defined as:

``` c
// Provided by VK_EXT_custom_border_color
typedef struct VkPhysicalDeviceCustomBorderColorPropertiesEXT {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           maxCustomBorderColorSamplers;
} VkPhysicalDeviceCustomBorderColorPropertiesEXT;
```

## <a href="#_members" class="anchor"></a>Members

- <span id="limits-maxCustomBorderColorSamplers"></span>
  `maxCustomBorderColorSamplers` indicates the maximum number of
  samplers with custom border colors which **can** simultaneously exist
  on a device.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceCustomBorderColorPropertiesEXT` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)
structure passed to
[vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html),
it is filled in with each corresponding implementation-dependent
property.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceCustomBorderColorPropertiesEXT-sType-sType"
  id="VUID-VkPhysicalDeviceCustomBorderColorPropertiesEXT-sType-sType"></a>
  VUID-VkPhysicalDeviceCustomBorderColorPropertiesEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_CUSTOM_BORDER_COLOR_PROPERTIES_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_custom_border_color](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_custom_border_color.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceCustomBorderColorPropertiesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
