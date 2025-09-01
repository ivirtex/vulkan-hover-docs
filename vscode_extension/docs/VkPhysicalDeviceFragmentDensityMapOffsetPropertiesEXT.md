# VkPhysicalDeviceFragmentDensityMapOffsetPropertiesEXT(3) Manual Page

## Name

VkPhysicalDeviceFragmentDensityMapOffsetPropertiesEXT - Structure describing fragment density map offset properties that can be supported by an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceFragmentDensityMapOffsetPropertiesEXT` structure is defined as:

```c++
// Provided by VK_EXT_fragment_density_map_offset
typedef struct VkPhysicalDeviceFragmentDensityMapOffsetPropertiesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkExtent2D         fragmentDensityOffsetGranularity;
} VkPhysicalDeviceFragmentDensityMapOffsetPropertiesEXT;
```

or the equivalent

```c++
// Provided by VK_QCOM_fragment_density_map_offset
typedef VkPhysicalDeviceFragmentDensityMapOffsetPropertiesEXT VkPhysicalDeviceFragmentDensityMapOffsetPropertiesQCOM;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`fragmentDensityOffsetGranularity` is the granularity for [fragment density offsets](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-fragmentdensitymapoffsets).

## [](#_description)Description

If the `VkPhysicalDeviceFragmentDensityMapOffsetPropertiesEXT` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceFragmentDensityMapOffsetPropertiesEXT-sType-sType)VUID-VkPhysicalDeviceFragmentDensityMapOffsetPropertiesEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAGMENT_DENSITY_MAP_OFFSET_PROPERTIES_EXT`

## [](#_see_also)See Also

[VK\_EXT\_fragment\_density\_map\_offset](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_fragment_density_map_offset.html), [VkExtent2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtent2D.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceFragmentDensityMapOffsetPropertiesEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0