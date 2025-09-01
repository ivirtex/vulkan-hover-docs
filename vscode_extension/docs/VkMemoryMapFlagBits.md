# VkMemoryMapFlagBits(3) Manual Page

## Name

VkMemoryMapFlagBits - Bitmask specifying additional parameters of a memory map



## [](#_c_specification)C Specification

Bits which **can** be set in [vkMapMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/vkMapMemory.html)::`flags` and [VkMemoryMapInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryMapInfo.html)::`flags`, specifying additional properties of a memory map, are:

```c++
// Provided by VK_VERSION_1_0
typedef enum VkMemoryMapFlagBits {
  // Provided by VK_EXT_map_memory_placed
    VK_MEMORY_MAP_PLACED_BIT_EXT = 0x00000001,
} VkMemoryMapFlagBits;
```

## [](#_description)Description

- `VK_MEMORY_MAP_PLACED_BIT_EXT` requests that the implementation place the memory map at the virtual address specified by the application via [VkMemoryMapPlacedInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryMapPlacedInfoEXT.html)::`pPlacedAddress`, replacing any existing mapping at that address. This flag **must** not be used with [vkMapMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/vkMapMemory.html) as there is no way to specify the placement address.

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkMemoryMapFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryMapFlags.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkMemoryMapFlagBits).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0