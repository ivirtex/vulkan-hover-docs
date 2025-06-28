# VkPhysicalDeviceFragmentDensityMapLayeredPropertiesVALVE(3) Manual Page

## Name

VkPhysicalDeviceFragmentDensityMapLayeredPropertiesVALVE - Structure describing fragment density map layered properties that can be supported by an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceFragmentDensityMapLayeredPropertiesVALVE` structure is defined as:

```c++
// Provided by VK_VALVE_fragment_density_map_layered
typedef struct VkPhysicalDeviceFragmentDensityMapLayeredPropertiesVALVE {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           maxFragmentDensityMapLayers;
} VkPhysicalDeviceFragmentDensityMapLayeredPropertiesVALVE;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`maxFragmentDensityMapLayers` is the maximum number of layers to use with a layered fragment density map.

## [](#_description)Description

If the `VkPhysicalDeviceFragmentDensityMapLayeredPropertiesVALVE` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceFragmentDensityMapLayeredPropertiesVALVE-sType-sType)VUID-VkPhysicalDeviceFragmentDensityMapLayeredPropertiesVALVE-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAGMENT_DENSITY_MAP_LAYERED_PROPERTIES_VALVE`

## [](#_see_also)See Also

[VK\_VALVE\_fragment\_density\_map\_layered](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VALVE_fragment_density_map_layered.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceFragmentDensityMapLayeredPropertiesVALVE)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0