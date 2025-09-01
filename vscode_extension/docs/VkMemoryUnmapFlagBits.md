# VkMemoryUnmapFlagBits(3) Manual Page

## Name

VkMemoryUnmapFlagBits - Bitmask specifying additional parameters of a memory unmap



## [](#_c_specification)C Specification

Bits which **can** be set in [VkMemoryUnmapInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryUnmapInfo.html)::`flags`, specifying additional properties of a memory unmap, are:

```c++
// Provided by VK_VERSION_1_4
typedef enum VkMemoryUnmapFlagBits {
  // Provided by VK_EXT_map_memory_placed
    VK_MEMORY_UNMAP_RESERVE_BIT_EXT = 0x00000001,
} VkMemoryUnmapFlagBits;
```

or the equivalent

```c++
// Provided by VK_KHR_map_memory2
typedef VkMemoryUnmapFlagBits VkMemoryUnmapFlagBitsKHR;
```

## [](#_description)Description

- `VK_MEMORY_UNMAP_RESERVE_BIT_EXT` requests that virtual address range currently occupied by the memory map remain reserved after the [vkUnmapMemory2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkUnmapMemory2.html) call completes. Future system memory map operations or calls to [vkMapMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/vkMapMemory.html) or [vkMapMemory2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkMapMemory2.html) will not return addresses in that range unless the range has since been unreserved by the client or the mapping is explicitly placed in that range by calling [vkMapMemory2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkMapMemory2.html) with `VK_MEMORY_MAP_PLACED_BIT_EXT`, or doing the system memory map equivalent. When `VK_MEMORY_UNMAP_RESERVE_BIT_EXT` is set, the memory unmap operation **may** fail, in which case the memory object will remain host mapped and [vkUnmapMemory2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkUnmapMemory2.html) will return `VK_ERROR_MEMORY_MAP_FAILED`.

## [](#_see_also)See Also

[VK\_KHR\_map\_memory2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_map_memory2.html), [VK\_VERSION\_1\_4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_4.html), [VkMemoryUnmapFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryUnmapFlags.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkMemoryUnmapFlagBits).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0