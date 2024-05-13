# VkMemoryMapInfoKHR(3) Manual Page

## Name

VkMemoryMapInfoKHR - Structure containing parameters of a memory map
operation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkMemoryMapInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_map_memory2
typedef struct VkMemoryMapInfoKHR {
    VkStructureType     sType;
    const void*         pNext;
    VkMemoryMapFlags    flags;
    VkDeviceMemory      memory;
    VkDeviceSize        offset;
    VkDeviceSize        size;
} VkMemoryMapInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is a bitmask of
  [VkMemoryMapFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryMapFlagBits.html) specifying additional
  parameters of the memory map operation.

- `memory` is the [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html) object to be
  mapped.

- `offset` is a zero-based byte offset from the beginning of the memory
  object.

- `size` is the size of the memory range to map, or `VK_WHOLE_SIZE` to
  map from `offset` to the end of the allocation.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkMemoryMapInfoKHR-memory-07958"
  id="VUID-VkMemoryMapInfoKHR-memory-07958"></a>
  VUID-VkMemoryMapInfoKHR-memory-07958  
  `memory` **must** not be currently host mapped

- <a href="#VUID-VkMemoryMapInfoKHR-offset-07959"
  id="VUID-VkMemoryMapInfoKHR-offset-07959"></a>
  VUID-VkMemoryMapInfoKHR-offset-07959  
  `offset` **must** be less than the size of `memory`

- <a href="#VUID-VkMemoryMapInfoKHR-size-07960"
  id="VUID-VkMemoryMapInfoKHR-size-07960"></a>
  VUID-VkMemoryMapInfoKHR-size-07960  
  If `size` is not equal to `VK_WHOLE_SIZE`, `size` **must** be greater
  than `0`

- <a href="#VUID-VkMemoryMapInfoKHR-size-07961"
  id="VUID-VkMemoryMapInfoKHR-size-07961"></a>
  VUID-VkMemoryMapInfoKHR-size-07961  
  If `size` is not equal to `VK_WHOLE_SIZE`, `size` **must** be less
  than or equal to the size of the `memory` minus `offset`

- <a href="#VUID-VkMemoryMapInfoKHR-memory-07962"
  id="VUID-VkMemoryMapInfoKHR-memory-07962"></a>
  VUID-VkMemoryMapInfoKHR-memory-07962  
  `memory` **must** have been created with a memory type that reports
  `VK_MEMORY_PROPERTY_HOST_VISIBLE_BIT`

- <a href="#VUID-VkMemoryMapInfoKHR-memory-07963"
  id="VUID-VkMemoryMapInfoKHR-memory-07963"></a>
  VUID-VkMemoryMapInfoKHR-memory-07963  
  `memory` **must** not have been allocated with multiple instances

- <a href="#VUID-VkMemoryMapInfoKHR-flags-09569"
  id="VUID-VkMemoryMapInfoKHR-flags-09569"></a>
  VUID-VkMemoryMapInfoKHR-flags-09569  
  If `VK_MEMORY_MAP_PLACED_BIT_EXT` is set in `flags`, the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-memoryMapPlaced"
  target="_blank" rel="noopener"><code>memoryMapPlaced</code></a>
  feature **must** be enabled

- <a href="#VUID-VkMemoryMapInfoKHR-flags-09570"
  id="VUID-VkMemoryMapInfoKHR-flags-09570"></a>
  VUID-VkMemoryMapInfoKHR-flags-09570  
  If `VK_MEMORY_MAP_PLACED_BIT_EXT` is set in `flags`, the `pNext` chain
  **must** include a
  [VkMemoryMapPlacedInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryMapPlacedInfoEXT.html) structure
  and `VkMemoryMapPlacedInfoEXT`::`pPlacedAddress` **must** not be
  `NULL`

- <a href="#VUID-VkMemoryMapInfoKHR-flags-09571"
  id="VUID-VkMemoryMapInfoKHR-flags-09571"></a>
  VUID-VkMemoryMapInfoKHR-flags-09571  
  If `VK_MEMORY_MAP_PLACED_BIT_EXT` is set in `flags` and the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-memoryMapRangePlaced"
  target="_blank" rel="noopener"><code>memoryMapRangePlaced</code></a>
  feature is not enabled, `offset` **must** be zero

- <a href="#VUID-VkMemoryMapInfoKHR-flags-09572"
  id="VUID-VkMemoryMapInfoKHR-flags-09572"></a>
  VUID-VkMemoryMapInfoKHR-flags-09572  
  If `VK_MEMORY_MAP_PLACED_BIT_EXT` is set in `flags` and the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-memoryMapRangePlaced"
  target="_blank" rel="noopener"><code>memoryMapRangePlaced</code></a>
  feature is not enabled, `size` **must** be `VK_WHOLE_SIZE` or
  `VkMemoryAllocateInfo`::`allocationSize`

- <a href="#VUID-VkMemoryMapInfoKHR-flags-09573"
  id="VUID-VkMemoryMapInfoKHR-flags-09573"></a>
  VUID-VkMemoryMapInfoKHR-flags-09573  
  If `VK_MEMORY_MAP_PLACED_BIT_EXT` is set in `flags` and the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-memoryMapRangePlaced"
  target="_blank" rel="noopener"><code>memoryMapRangePlaced</code></a>
  feature is enabled, `offset` **must** be aligned to an integer
  multiple of
  `VkPhysicalDeviceMapMemoryPlacedPropertiesEXT`::`minPlacedMemoryMapAlignment`

- <a href="#VUID-VkMemoryMapInfoKHR-flags-09574"
  id="VUID-VkMemoryMapInfoKHR-flags-09574"></a>
  VUID-VkMemoryMapInfoKHR-flags-09574  
  If `VK_MEMORY_MAP_PLACED_BIT_EXT` is set in `flags` and `size` is not
  `VK_WHOLE_SIZE`, `size` **must** be aligned to an integer multiple of
  `VkPhysicalDeviceMapMemoryPlacedPropertiesEXT`::`minPlacedMemoryMapAlignment`

- <a href="#VUID-VkMemoryMapInfoKHR-flags-09651"
  id="VUID-VkMemoryMapInfoKHR-flags-09651"></a>
  VUID-VkMemoryMapInfoKHR-flags-09651  
  If `VK_MEMORY_MAP_PLACED_BIT_EXT` is set in `flags` and `size` is
  `VK_WHOLE_SIZE`, `VkMemoryAllocateInfo`::`allocationSize` **must** be
  aligned to an integer multiple of
  `VkPhysicalDeviceMapMemoryPlacedPropertiesEXT`::`minPlacedMemoryMapAlignment`

- <a href="#VUID-VkMemoryMapInfoKHR-flags-09575"
  id="VUID-VkMemoryMapInfoKHR-flags-09575"></a>
  VUID-VkMemoryMapInfoKHR-flags-09575  
  If `VK_MEMORY_MAP_PLACED_BIT_EXT` is set in `flags`, the memory object
  **must** not have been imported from a handle type of
  `VK_EXTERNAL_MEMORY_HANDLE_TYPE_HOST_ALLOCATION_BIT_EXT` or
  `VK_EXTERNAL_MEMORY_HANDLE_TYPE_HOST_MAPPED_FOREIGN_MEMORY_BIT_EXT`

Valid Usage (Implicit)

- <a href="#VUID-VkMemoryMapInfoKHR-sType-sType"
  id="VUID-VkMemoryMapInfoKHR-sType-sType"></a>
  VUID-VkMemoryMapInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_MEMORY_MAP_INFO_KHR`

- <a href="#VUID-VkMemoryMapInfoKHR-pNext-pNext"
  id="VUID-VkMemoryMapInfoKHR-pNext-pNext"></a>
  VUID-VkMemoryMapInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL` or a pointer to a valid instance of
  [VkMemoryMapPlacedInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryMapPlacedInfoEXT.html)

- <a href="#VUID-VkMemoryMapInfoKHR-sType-unique"
  id="VUID-VkMemoryMapInfoKHR-sType-unique"></a>
  VUID-VkMemoryMapInfoKHR-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

- <a href="#VUID-VkMemoryMapInfoKHR-flags-parameter"
  id="VUID-VkMemoryMapInfoKHR-flags-parameter"></a>
  VUID-VkMemoryMapInfoKHR-flags-parameter  
  `flags` **must** be a valid combination of
  [VkMemoryMapFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryMapFlagBits.html) values

- <a href="#VUID-VkMemoryMapInfoKHR-memory-parameter"
  id="VUID-VkMemoryMapInfoKHR-memory-parameter"></a>
  VUID-VkMemoryMapInfoKHR-memory-parameter  
  `memory` **must** be a valid [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html)
  handle

Host Synchronization

- Host access to `memory` **must** be externally synchronized

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_map_memory2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_map_memory2.html),
[VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html),
[VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html),
[VkMemoryMapFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryMapFlags.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkMapMemory2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkMapMemory2KHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkMemoryMapInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
