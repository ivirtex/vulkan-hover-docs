# VkMemoryMapFlagBits(3) Manual Page

## Name

VkMemoryMapFlagBits - Bitmask specifying additional parameters of a
memory map



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **can** be set in [vkMapMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkMapMemory.html)::`flags`
and [VkMemoryMapInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryMapInfoKHR.html)::`flags`, specifying
additional properties of a memory map, are:

``` c
// Provided by VK_VERSION_1_0
typedef enum VkMemoryMapFlagBits {
  // Provided by VK_EXT_map_memory_placed
    VK_MEMORY_MAP_PLACED_BIT_EXT = 0x00000001,
} VkMemoryMapFlagBits;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_MEMORY_MAP_PLACED_BIT_EXT` requests that the implementation place
  the memory map at the virtual address specified by the client via
  [VkMemoryMapPlacedInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryMapPlacedInfoEXT.html)::`pPlacedAddress`,
  replacing any existing mapping at that address. This flag **must** not
  be used with [vkMapMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkMapMemory.html) as there is no way to
  specify the placement address.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkMemoryMapFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryMapFlags.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkMemoryMapFlagBits"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
