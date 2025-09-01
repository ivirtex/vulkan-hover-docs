# VkMemoryMapInfo(3) Manual Page

## Name

VkMemoryMapInfo - Structure containing parameters of a memory map operation



## [](#_c_specification)C Specification

The `VkMemoryMapInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_4
typedef struct VkMemoryMapInfo {
    VkStructureType     sType;
    const void*         pNext;
    VkMemoryMapFlags    flags;
    VkDeviceMemory      memory;
    VkDeviceSize        offset;
    VkDeviceSize        size;
} VkMemoryMapInfo;
```

or the equivalent

```c++
// Provided by VK_KHR_map_memory2
typedef VkMemoryMapInfo VkMemoryMapInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is a bitmask of [VkMemoryMapFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryMapFlagBits.html) specifying additional parameters of the memory map operation.
- `memory` is the [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) object to be mapped.
- `offset` is a zero-based byte offset from the beginning of the memory object.
- `size` is the size of the memory range to map, or `VK_WHOLE_SIZE` to map from `offset` to the end of the allocation.

## [](#_description)Description

Valid Usage

- [](#VUID-VkMemoryMapInfo-memory-07958)VUID-VkMemoryMapInfo-memory-07958  
  `memory` **must** not be currently host mapped
- [](#VUID-VkMemoryMapInfo-offset-07959)VUID-VkMemoryMapInfo-offset-07959  
  `offset` **must** be less than the size of `memory`
- [](#VUID-VkMemoryMapInfo-size-07960)VUID-VkMemoryMapInfo-size-07960  
  If `size` is not equal to `VK_WHOLE_SIZE`, `size` **must** be greater than `0`
- [](#VUID-VkMemoryMapInfo-size-07961)VUID-VkMemoryMapInfo-size-07961  
  If `size` is not equal to `VK_WHOLE_SIZE`, `size` **must** be less than or equal to the size of the `memory` minus `offset`
- [](#VUID-VkMemoryMapInfo-memory-07962)VUID-VkMemoryMapInfo-memory-07962  
  `memory` **must** have been created with a memory type that reports `VK_MEMORY_PROPERTY_HOST_VISIBLE_BIT`
- [](#VUID-VkMemoryMapInfo-memory-07963)VUID-VkMemoryMapInfo-memory-07963  
  `memory` **must** not have been allocated with multiple instances
- [](#VUID-VkMemoryMapInfo-flags-09569)VUID-VkMemoryMapInfo-flags-09569  
  If `VK_MEMORY_MAP_PLACED_BIT_EXT` is set in `flags`, the [`memoryMapPlaced`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-memoryMapPlaced) feature **must** be enabled
- [](#VUID-VkMemoryMapInfo-flags-09570)VUID-VkMemoryMapInfo-flags-09570  
  If `VK_MEMORY_MAP_PLACED_BIT_EXT` is set in `flags`, the `pNext` chain **must** include a [VkMemoryMapPlacedInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryMapPlacedInfoEXT.html) structure and `VkMemoryMapPlacedInfoEXT`::`pPlacedAddress` **must** not be `NULL`
- [](#VUID-VkMemoryMapInfo-flags-09571)VUID-VkMemoryMapInfo-flags-09571  
  If `VK_MEMORY_MAP_PLACED_BIT_EXT` is set in `flags` and the [`memoryMapRangePlaced`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-memoryMapRangePlaced) feature is not enabled, `offset` **must** be zero
- [](#VUID-VkMemoryMapInfo-flags-09572)VUID-VkMemoryMapInfo-flags-09572  
  If `VK_MEMORY_MAP_PLACED_BIT_EXT` is set in `flags` and the [`memoryMapRangePlaced`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-memoryMapRangePlaced) feature is not enabled, `size` **must** be `VK_WHOLE_SIZE` or `VkMemoryAllocateInfo`::`allocationSize`
- [](#VUID-VkMemoryMapInfo-flags-09573)VUID-VkMemoryMapInfo-flags-09573  
  If `VK_MEMORY_MAP_PLACED_BIT_EXT` is set in `flags` and the [`memoryMapRangePlaced`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-memoryMapRangePlaced) feature is enabled, `offset` **must** be aligned to an integer multiple of `VkPhysicalDeviceMapMemoryPlacedPropertiesEXT`::`minPlacedMemoryMapAlignment`
- [](#VUID-VkMemoryMapInfo-flags-09574)VUID-VkMemoryMapInfo-flags-09574  
  If `VK_MEMORY_MAP_PLACED_BIT_EXT` is set in `flags` and `size` is not `VK_WHOLE_SIZE`, `size` **must** be aligned to an integer multiple of `VkPhysicalDeviceMapMemoryPlacedPropertiesEXT`::`minPlacedMemoryMapAlignment`
- [](#VUID-VkMemoryMapInfo-flags-09651)VUID-VkMemoryMapInfo-flags-09651  
  If `VK_MEMORY_MAP_PLACED_BIT_EXT` is set in `flags` and `size` is `VK_WHOLE_SIZE`, `VkMemoryAllocateInfo`::`allocationSize` **must** be aligned to an integer multiple of `VkPhysicalDeviceMapMemoryPlacedPropertiesEXT`::`minPlacedMemoryMapAlignment`
- [](#VUID-VkMemoryMapInfo-flags-09575)VUID-VkMemoryMapInfo-flags-09575  
  If `VK_MEMORY_MAP_PLACED_BIT_EXT` is set in `flags`, the memory object **must** not have been imported from a handle type of `VK_EXTERNAL_MEMORY_HANDLE_TYPE_HOST_ALLOCATION_BIT_EXT` or `VK_EXTERNAL_MEMORY_HANDLE_TYPE_HOST_MAPPED_FOREIGN_MEMORY_BIT_EXT`

Valid Usage (Implicit)

- [](#VUID-VkMemoryMapInfo-sType-sType)VUID-VkMemoryMapInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_MEMORY_MAP_INFO`
- [](#VUID-VkMemoryMapInfo-pNext-pNext)VUID-VkMemoryMapInfo-pNext-pNext  
  `pNext` **must** be `NULL` or a pointer to a valid instance of [VkMemoryMapPlacedInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryMapPlacedInfoEXT.html)
- [](#VUID-VkMemoryMapInfo-sType-unique)VUID-VkMemoryMapInfo-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique
- [](#VUID-VkMemoryMapInfo-flags-parameter)VUID-VkMemoryMapInfo-flags-parameter  
  `flags` **must** be a valid combination of [VkMemoryMapFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryMapFlagBits.html) values
- [](#VUID-VkMemoryMapInfo-memory-parameter)VUID-VkMemoryMapInfo-memory-parameter  
  `memory` **must** be a valid [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) handle

Host Synchronization

- Host access to `memory` **must** be externally synchronized

## [](#_see_also)See Also

[VK\_KHR\_map\_memory2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_map_memory2.html), [VK\_VERSION\_1\_4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_4.html), [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html), [VkMemoryMapFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryMapFlags.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkMapMemory2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkMapMemory2.html), [vkMapMemory2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkMapMemory2.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkMemoryMapInfo).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0