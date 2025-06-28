# VkMemoryMapPlacedInfoEXT(3) Manual Page

## Name

VkMemoryMapPlacedInfoEXT - Structure containing memory map placement parameters



## [](#_c_specification)C Specification

If `VK_MEMORY_MAP_PLACED_BIT_EXT` is set in `VkMemoryMapInfo`::`flags` and the `pNext` chain of [VkMemoryMapInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryMapInfo.html) includes a `VkMemoryMapPlacedInfoEXT` structure, then that structure specifies the placement address of the memory map. The implementation will place the memory map at the specified address, replacing any existing maps in the specified memory range. Replacing memory maps in this way does not implicitly unmap Vulkan memory objects. Instead, the application **must** ensure no other Vulkan memory objects are mapped anywhere in the specified virtual address range. If successful, `ppData` will be set to the same value as `VkMemoryMapPlacedInfoEXT`::`pPlacedAddress` and `vkMapMemory2` will return `VK_SUCCESS`. If it cannot place the map at the requested address for any reason, the memory object is left unmapped and `vkMapMemory2` will return `VK_ERROR_MEMORY_MAP_FAILED`.

The `VkMemoryMapPlacedInfoEXT` structure is defined as:

```c++
// Provided by VK_EXT_map_memory_placed
typedef struct VkMemoryMapPlacedInfoEXT {
    VkStructureType    sType;
    const void*        pNext;
    void*              pPlacedAddress;
} VkMemoryMapPlacedInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `pPlacedAddress` is the virtual address at which to place the address. If `VkMemoryMapInfo`::`flags` does not contain `VK_MEMORY_MAP_PLACED_BIT_EXT`, this value is ignored.

## [](#_description)Description

Valid Usage

- [](#VUID-VkMemoryMapPlacedInfoEXT-flags-09576)VUID-VkMemoryMapPlacedInfoEXT-flags-09576  
  If `VkMemoryMapInfo`::`flags` contains `VK_MEMORY_MAP_PLACED_BIT_EXT`, `pPlacedAddress` **must** not be `NULL`
- [](#VUID-VkMemoryMapPlacedInfoEXT-pPlacedAddress-09577)VUID-VkMemoryMapPlacedInfoEXT-pPlacedAddress-09577  
  `pPlacedAddress` **must** be aligned to an integer multiple of `VkPhysicalDeviceMapMemoryPlacedPropertiesEXT`::`minPlacedMemoryMapAlignment`
- [](#VUID-VkMemoryMapPlacedInfoEXT-pPlacedAddress-09578)VUID-VkMemoryMapPlacedInfoEXT-pPlacedAddress-09578  
  The address range specified by `pPlacedAddress` and `VkMemoryMapInfo`::`size` **must** not overlap any existing Vulkan memory object mapping

Valid Usage (Implicit)

- [](#VUID-VkMemoryMapPlacedInfoEXT-sType-sType)VUID-VkMemoryMapPlacedInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_MEMORY_MAP_PLACED_INFO_EXT`

## [](#_see_also)See Also

[VK\_EXT\_map\_memory\_placed](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_map_memory_placed.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkMemoryMapPlacedInfoEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0