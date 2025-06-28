# VkPhysicalDeviceDisplacementMicromapPropertiesNV(3) Manual Page

## Name

VkPhysicalDeviceDisplacementMicromapPropertiesNV - Structure describing the displacement micromap properties of a physical device



## [](#_c_specification)C Specification

The `VkPhysicalDeviceDisplacementMicromapPropertiesNV` structure is defined as:

```c++
// Provided by VK_NV_displacement_micromap
typedef struct VkPhysicalDeviceDisplacementMicromapPropertiesNV {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           maxDisplacementMicromapSubdivisionLevel;
} VkPhysicalDeviceDisplacementMicromapPropertiesNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `maxDisplacementMicromapSubdivisionLevel` is the maximum allowed `subdivisionLevel` for displacement micromaps.

## [](#_description)Description

If the `VkPhysicalDeviceDisplacementMicromapPropertiesNV` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceDisplacementMicromapPropertiesNV-sType-sType)VUID-VkPhysicalDeviceDisplacementMicromapPropertiesNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DISPLACEMENT_MICROMAP_PROPERTIES_NV`

## [](#_see_also)See Also

[VK\_NV\_displacement\_micromap](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_displacement_micromap.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceDisplacementMicromapPropertiesNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0