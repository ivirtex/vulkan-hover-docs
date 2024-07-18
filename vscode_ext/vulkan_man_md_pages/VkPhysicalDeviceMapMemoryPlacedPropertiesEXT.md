# VkPhysicalDeviceMapMemoryPlacedPropertiesEXT(3) Manual Page

## Name

VkPhysicalDeviceMapMemoryPlacedPropertiesEXT - Structure describing the
alignment requirements of placed memory maps for a physical device



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceMapMemoryPlacedPropertiesEXT` structure is defined
as:

``` c
// Provided by VK_EXT_map_memory_placed
typedef struct VkPhysicalDeviceMapMemoryPlacedPropertiesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkDeviceSize       minPlacedMemoryMapAlignment;
} VkPhysicalDeviceMapMemoryPlacedPropertiesEXT;
```

## <a href="#_members" class="anchor"></a>Members

The members of the `VkPhysicalDeviceMapMemoryPlacedPropertiesEXT`
structure describe the following:

## <a href="#_description" class="anchor"></a>Description

- <span id="limits-minPlacedMemoryMapAlignment"></span>
  `minPlacedMemoryMapAlignment` is the minimum alignment required for
  memory object offsets and virtual address ranges when using placed
  memory mapping.

If the `VkPhysicalDeviceMapMemoryPlacedPropertiesEXT` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)
structure passed to
[vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html),
it is filled in with each corresponding implementation-dependent
property.

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceMapMemoryPlacedPropertiesEXT-sType-sType"
  id="VUID-VkPhysicalDeviceMapMemoryPlacedPropertiesEXT-sType-sType"></a>
  VUID-VkPhysicalDeviceMapMemoryPlacedPropertiesEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAP_MEMORY_PLACED_PROPERTIES_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_map_memory_placed](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_map_memory_placed.html),
[VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceMapMemoryPlacedPropertiesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
