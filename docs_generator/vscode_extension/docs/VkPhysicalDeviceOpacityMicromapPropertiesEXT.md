# VkPhysicalDeviceOpacityMicromapPropertiesEXT(3) Manual Page

## Name

VkPhysicalDeviceOpacityMicromapPropertiesEXT - Structure describing the opacity micromap properties of a physical device



## [](#_c_specification)C Specification

The `VkPhysicalDeviceOpacityMicromapPropertiesEXT` structure is defined as:

```c++
// Provided by VK_EXT_opacity_micromap
typedef struct VkPhysicalDeviceOpacityMicromapPropertiesEXT {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           maxOpacity2StateSubdivisionLevel;
    uint32_t           maxOpacity4StateSubdivisionLevel;
} VkPhysicalDeviceOpacityMicromapPropertiesEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `maxOpacity2StateSubdivisionLevel` is the maximum allowed `subdivisionLevel` when `format` is `VK_OPACITY_MICROMAP_FORMAT_2_STATE_EXT`
- `maxOpacity4StateSubdivisionLevel` is the maximum allowed `subdivisionLevel` when `format` is `VK_OPACITY_MICROMAP_FORMAT_4_STATE_EXT`

## [](#_description)Description

If the `VkPhysicalDeviceOpacityMicromapPropertiesEXT` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceOpacityMicromapPropertiesEXT-sType-sType)VUID-VkPhysicalDeviceOpacityMicromapPropertiesEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_OPACITY_MICROMAP_PROPERTIES_EXT`

## [](#_see_also)See Also

[VK\_EXT\_opacity\_micromap](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_opacity_micromap.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceOpacityMicromapPropertiesEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0