# VkPhysicalDeviceFragmentDensityMapPropertiesEXT(3) Manual Page

## Name

VkPhysicalDeviceFragmentDensityMapPropertiesEXT - Structure describing fragment density map properties that can be supported by an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceFragmentDensityMapPropertiesEXT` structure is defined as:

```c++
// Provided by VK_EXT_fragment_density_map
typedef struct VkPhysicalDeviceFragmentDensityMapPropertiesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkExtent2D         minFragmentDensityTexelSize;
    VkExtent2D         maxFragmentDensityTexelSize;
    VkBool32           fragmentDensityInvocations;
} VkPhysicalDeviceFragmentDensityMapPropertiesEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`minFragmentDensityTexelSize` is the minimum [fragment density texel size](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#glossary-fragment-density-texel-size).
- []()`maxFragmentDensityTexelSize` is the maximum fragment density texel size.
- []()`fragmentDensityInvocations` specifies whether the implementation **may** invoke additional fragment shader invocations for each covered sample.

## [](#_description)Description

If the `VkPhysicalDeviceFragmentDensityMapPropertiesEXT` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceFragmentDensityMapPropertiesEXT-sType-sType)VUID-VkPhysicalDeviceFragmentDensityMapPropertiesEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAGMENT_DENSITY_MAP_PROPERTIES_EXT`

## [](#_see_also)See Also

[VK\_EXT\_fragment\_density\_map](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_fragment_density_map.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkExtent2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtent2D.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceFragmentDensityMapPropertiesEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0