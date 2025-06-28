# VkMemoryUnmapInfo(3) Manual Page

## Name

VkMemoryUnmapInfo - Structure containing parameters of a memory unmap operation



## [](#_c_specification)C Specification

The `VkMemoryUnmapInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_4
typedef struct VkMemoryUnmapInfo {
    VkStructureType       sType;
    const void*           pNext;
    VkMemoryUnmapFlags    flags;
    VkDeviceMemory        memory;
} VkMemoryUnmapInfo;
```

or the equivalent

```c++
// Provided by VK_KHR_map_memory2
typedef VkMemoryUnmapInfo VkMemoryUnmapInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is a bitmask of [VkMemoryUnmapFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryUnmapFlagBits.html) specifying additional parameters of the memory map operation.
- `memory` is the [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) object to be unmapped.

## [](#_description)Description

Valid Usage

- [](#VUID-VkMemoryUnmapInfo-memory-07964)VUID-VkMemoryUnmapInfo-memory-07964  
  `memory` **must** be currently host mapped
- [](#VUID-VkMemoryUnmapInfo-flags-09579)VUID-VkMemoryUnmapInfo-flags-09579  
  If `VK_MEMORY_UNMAP_RESERVE_BIT_EXT` is set in `flags`, the [`memoryUnmapReserve`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-memoryUnmapReserve) **must** be enabled
- [](#VUID-VkMemoryUnmapInfo-flags-09580)VUID-VkMemoryUnmapInfo-flags-09580  
  If `VK_MEMORY_UNMAP_RESERVE_BIT_EXT` is set in `flags`, the memory object **must** not have been imported from a handle type of `VK_EXTERNAL_MEMORY_HANDLE_TYPE_HOST_ALLOCATION_BIT_EXT` or `VK_EXTERNAL_MEMORY_HANDLE_TYPE_HOST_MAPPED_FOREIGN_MEMORY_BIT_EXT`

Valid Usage (Implicit)

- [](#VUID-VkMemoryUnmapInfo-sType-sType)VUID-VkMemoryUnmapInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_MEMORY_UNMAP_INFO`
- [](#VUID-VkMemoryUnmapInfo-pNext-pNext)VUID-VkMemoryUnmapInfo-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkMemoryUnmapInfo-flags-parameter)VUID-VkMemoryUnmapInfo-flags-parameter  
  `flags` **must** be a valid combination of [VkMemoryUnmapFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryUnmapFlagBits.html) values
- [](#VUID-VkMemoryUnmapInfo-memory-parameter)VUID-VkMemoryUnmapInfo-memory-parameter  
  `memory` **must** be a valid [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) handle

Host Synchronization

- Host access to `memory` **must** be externally synchronized

## [](#_see_also)See Also

[VK\_KHR\_map\_memory2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_map_memory2.html), [VK\_VERSION\_1\_4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_4.html), [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html), [VkMemoryUnmapFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryUnmapFlags.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkUnmapMemory2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkUnmapMemory2.html), [vkUnmapMemory2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkUnmapMemory2KHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkMemoryUnmapInfo)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0