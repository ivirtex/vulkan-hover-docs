# VkPhysicalDeviceMapMemoryPlacedPropertiesEXT(3) Manual Page

## Name

VkPhysicalDeviceMapMemoryPlacedPropertiesEXT - Structure describing the alignment requirements of placed memory maps for a physical device



## [](#_c_specification)C Specification

The `VkPhysicalDeviceMapMemoryPlacedPropertiesEXT` structure is defined as:

```c++
// Provided by VK_EXT_map_memory_placed
typedef struct VkPhysicalDeviceMapMemoryPlacedPropertiesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkDeviceSize       minPlacedMemoryMapAlignment;
} VkPhysicalDeviceMapMemoryPlacedPropertiesEXT;
```

## [](#_members)Members

The members of the `VkPhysicalDeviceMapMemoryPlacedPropertiesEXT` structure describe the following:

## [](#_description)Description

- []()`minPlacedMemoryMapAlignment` is the minimum alignment required for memory object offsets and virtual address ranges when using placed memory mapping.

If the `VkPhysicalDeviceMapMemoryPlacedPropertiesEXT` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceMapMemoryPlacedPropertiesEXT-sType-sType)VUID-VkPhysicalDeviceMapMemoryPlacedPropertiesEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAP_MEMORY_PLACED_PROPERTIES_EXT`

## [](#_see_also)See Also

[VK\_EXT\_map\_memory\_placed](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_map_memory_placed.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceMapMemoryPlacedPropertiesEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0