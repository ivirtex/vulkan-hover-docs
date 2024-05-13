# VkMemoryUnmapInfoKHR(3) Manual Page

## Name

VkMemoryUnmapInfoKHR - Structure containing parameters of a memory unmap
operation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkMemoryUnmapInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_map_memory2
typedef struct VkMemoryUnmapInfoKHR {
    VkStructureType          sType;
    const void*              pNext;
    VkMemoryUnmapFlagsKHR    flags;
    VkDeviceMemory           memory;
} VkMemoryUnmapInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is a bitmask of
  [VkMemoryUnmapFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryUnmapFlagBitsKHR.html) specifying
  additional parameters of the memory map operation.

- `memory` is the [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html) object to be
  unmapped.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkMemoryUnmapInfoKHR-memory-07964"
  id="VUID-VkMemoryUnmapInfoKHR-memory-07964"></a>
  VUID-VkMemoryUnmapInfoKHR-memory-07964  
  `memory` **must** be currently host mapped

- <a href="#VUID-VkMemoryUnmapInfoKHR-flags-09579"
  id="VUID-VkMemoryUnmapInfoKHR-flags-09579"></a>
  VUID-VkMemoryUnmapInfoKHR-flags-09579  
  If `VK_MEMORY_UNMAP_RESERVE_BIT_EXT` is set in `flags`, the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-memoryUnmapReserve"
  target="_blank" rel="noopener"><code>memoryUnmapReserve</code></a>
  **must** be enabled

- <a href="#VUID-VkMemoryUnmapInfoKHR-flags-09580"
  id="VUID-VkMemoryUnmapInfoKHR-flags-09580"></a>
  VUID-VkMemoryUnmapInfoKHR-flags-09580  
  If `VK_MEMORY_UNMAP_RESERVE_BIT_EXT` is set in `flags`, the memory
  object **must** not have been imported from a handle type of
  `VK_EXTERNAL_MEMORY_HANDLE_TYPE_HOST_ALLOCATION_BIT_EXT` or
  `VK_EXTERNAL_MEMORY_HANDLE_TYPE_HOST_MAPPED_FOREIGN_MEMORY_BIT_EXT`

Valid Usage (Implicit)

- <a href="#VUID-VkMemoryUnmapInfoKHR-sType-sType"
  id="VUID-VkMemoryUnmapInfoKHR-sType-sType"></a>
  VUID-VkMemoryUnmapInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_MEMORY_UNMAP_INFO_KHR`

- <a href="#VUID-VkMemoryUnmapInfoKHR-pNext-pNext"
  id="VUID-VkMemoryUnmapInfoKHR-pNext-pNext"></a>
  VUID-VkMemoryUnmapInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkMemoryUnmapInfoKHR-flags-parameter"
  id="VUID-VkMemoryUnmapInfoKHR-flags-parameter"></a>
  VUID-VkMemoryUnmapInfoKHR-flags-parameter  
  `flags` **must** be a valid combination of
  [VkMemoryUnmapFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryUnmapFlagBitsKHR.html) values

- <a href="#VUID-VkMemoryUnmapInfoKHR-memory-parameter"
  id="VUID-VkMemoryUnmapInfoKHR-memory-parameter"></a>
  VUID-VkMemoryUnmapInfoKHR-memory-parameter  
  `memory` **must** be a valid [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html)
  handle

Host Synchronization

- Host access to `memory` **must** be externally synchronized

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_map_memory2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_map_memory2.html),
[VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html),
[VkMemoryUnmapFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryUnmapFlagsKHR.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkUnmapMemory2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkUnmapMemory2KHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkMemoryUnmapInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
