# VkDisplayModeProperties2KHR(3) Manual Page

## Name

VkDisplayModeProperties2KHR - Structure describing an available display mode



## [](#_c_specification)C Specification

The `VkDisplayModeProperties2KHR` structure is defined as:

```c++
// Provided by VK_KHR_get_display_properties2
typedef struct VkDisplayModeProperties2KHR {
    VkStructureType               sType;
    void*                         pNext;
    VkDisplayModePropertiesKHR    displayModeProperties;
} VkDisplayModeProperties2KHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `displayModeProperties` is a [VkDisplayModePropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayModePropertiesKHR.html) structure.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkDisplayModeProperties2KHR-sType-sType)VUID-VkDisplayModeProperties2KHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DISPLAY_MODE_PROPERTIES_2_KHR`
- [](#VUID-VkDisplayModeProperties2KHR-pNext-pNext)VUID-VkDisplayModeProperties2KHR-pNext-pNext  
  `pNext` **must** be `NULL` or a pointer to a valid instance of [VkDisplayModeStereoPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayModeStereoPropertiesNV.html)
- [](#VUID-VkDisplayModeProperties2KHR-sType-unique)VUID-VkDisplayModeProperties2KHR-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique

## [](#_see_also)See Also

[VK\_KHR\_get\_display\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_display_properties2.html), [VkDisplayModePropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayModePropertiesKHR.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetDisplayModeProperties2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDisplayModeProperties2KHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDisplayModeProperties2KHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0