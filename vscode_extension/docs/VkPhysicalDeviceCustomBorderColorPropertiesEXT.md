# VkPhysicalDeviceCustomBorderColorPropertiesEXT(3) Manual Page

## Name

VkPhysicalDeviceCustomBorderColorPropertiesEXT - Structure describing whether custom border colors can be supported by an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceCustomBorderColorPropertiesEXT` structure is defined as:

```c++
// Provided by VK_EXT_custom_border_color
typedef struct VkPhysicalDeviceCustomBorderColorPropertiesEXT {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           maxCustomBorderColorSamplers;
} VkPhysicalDeviceCustomBorderColorPropertiesEXT;
```

## [](#_members)Members

- []()`maxCustomBorderColorSamplers` indicates the maximum number of samplers with custom border colors which **can** simultaneously exist on a device.

## [](#_description)Description

If the `VkPhysicalDeviceCustomBorderColorPropertiesEXT` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceCustomBorderColorPropertiesEXT-sType-sType)VUID-VkPhysicalDeviceCustomBorderColorPropertiesEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_CUSTOM_BORDER_COLOR_PROPERTIES_EXT`

## [](#_see_also)See Also

[VK\_EXT\_custom\_border\_color](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_custom_border_color.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceCustomBorderColorPropertiesEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0