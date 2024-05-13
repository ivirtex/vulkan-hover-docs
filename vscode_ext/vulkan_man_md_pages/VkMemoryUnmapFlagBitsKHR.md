# VkMemoryUnmapFlagBitsKHR(3) Manual Page

## Name

VkMemoryUnmapFlagBitsKHR - Bitmask specifying additional parameters of a
memory unmap



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **can** be set in
[VkMemoryUnmapInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryUnmapInfoKHR.html)::`flags`, specifying
additional properties of a memory unmap, are:

``` c
// Provided by VK_KHR_map_memory2
typedef enum VkMemoryUnmapFlagBitsKHR {
  // Provided by VK_EXT_map_memory_placed
    VK_MEMORY_UNMAP_RESERVE_BIT_EXT = 0x00000001,
} VkMemoryUnmapFlagBitsKHR;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_MEMORY_UNMAP_RESERVE_BIT_EXT` requests that virtual address range
  currently occupied by the memory map remain reserved after the
  [vkUnmapMemory2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkUnmapMemory2KHR.html) call completes. Future
  system memory map operations or calls to
  [vkMapMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkMapMemory.html) or
  [vkMapMemory2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkMapMemory2KHR.html) will not return addresses in
  that range unless the range has since been unreserved by the client or
  the mapping is explicitly placed in that range by calling
  [vkMapMemory2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkMapMemory2KHR.html) with
  `VK_MEMORY_MAP_PLACED_BIT_EXT`, or doing the system memory map
  equivalent. When `VK_MEMORY_UNMAP_RESERVE_BIT_EXT` is set, the memory
  unmap operation **may** fail, in which case the memory object will
  remain host mapped and [vkUnmapMemory2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkUnmapMemory2KHR.html)
  will return `VK_ERROR_MEMORY_MAP_FAILED`.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_map_memory2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_map_memory2.html),
[VkMemoryUnmapFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryUnmapFlagsKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkMemoryUnmapFlagBitsKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
