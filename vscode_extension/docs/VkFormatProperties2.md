# VkFormatProperties2(3) Manual Page

## Name

VkFormatProperties2 - Structure specifying image format properties



## [](#_c_specification)C Specification

The `VkFormatProperties2` structure is defined as:

```c++
// Provided by VK_VERSION_1_1
typedef struct VkFormatProperties2 {
    VkStructureType       sType;
    void*                 pNext;
    VkFormatProperties    formatProperties;
} VkFormatProperties2;
```

or the equivalent

```c++
// Provided by VK_KHR_get_physical_device_properties2
typedef VkFormatProperties2 VkFormatProperties2KHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `formatProperties` is a [VkFormatProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatProperties.html) structure describing features supported by the requested format.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkFormatProperties2-sType-sType)VUID-VkFormatProperties2-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_FORMAT_PROPERTIES_2`
- [](#VUID-VkFormatProperties2-pNext-pNext)VUID-VkFormatProperties2-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the `pNext` chain **must** be either `NULL` or a pointer to a valid instance of [VkDrmFormatModifierPropertiesList2EXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDrmFormatModifierPropertiesList2EXT.html), [VkDrmFormatModifierPropertiesListEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDrmFormatModifierPropertiesListEXT.html), [VkFormatProperties3](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatProperties3.html), [VkSubpassResolvePerformanceQueryEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassResolvePerformanceQueryEXT.html), or [VkTensorFormatPropertiesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorFormatPropertiesARM.html)
- [](#VUID-VkFormatProperties2-sType-unique)VUID-VkFormatProperties2-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique

## [](#_see_also)See Also

[VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkFormatProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatProperties.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetPhysicalDeviceFormatProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFormatProperties2.html), [vkGetPhysicalDeviceFormatProperties2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFormatProperties2KHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkFormatProperties2)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0