# VkMemoryMapPlacedInfoEXT(3) Manual Page

## Name

VkMemoryMapPlacedInfoEXT - Structure containing memory map placement
parameters



## <a href="#_c_specification" class="anchor"></a>C Specification

If `VK_MEMORY_MAP_PLACED_BIT_EXT` is set in
`VkMemoryMapInfoKHR`::`flags` and the `pNext` chain of
[VkMemoryMapInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryMapInfoKHR.html) includes a
`VkMemoryMapPlacedInfoEXT` structure, then that structure specifies the
placement address of the memory map. The implementation will place the
memory map at the specified address, replacing any existing maps in the
specified memory range. Replacing memory maps in this way does not
implicitly unmap Vulkan memory objects. Instead, the client **must**
ensure no other Vulkan memory objects are mapped anywhere in the
specified virtual address range. If successful, `ppData` will be set to
the same value as `VkMemoryMapPlacedInfoEXT`::`pPlacedAddress` and
`vkMapMemory2KHR` will return `VK_SUCCESS`. If it cannot place the map
at the requested address for any reason, the memory object is left
unmapped and `vkMapMemory2KHR` will return `VK_ERROR_MEMORY_MAP_FAILED`.

The `VkMemoryMapPlacedInfoEXT` structure is defined as:

``` c
// Provided by VK_EXT_map_memory_placed
typedef struct VkMemoryMapPlacedInfoEXT {
    VkStructureType    sType;
    const void*        pNext;
    void*              pPlacedAddress;
} VkMemoryMapPlacedInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `pPlacedAddress` is the virtual address at which to place the address.
  If `VkMemoryMapInfoKHR`::`flags` does not contain
  `VK_MEMORY_MAP_PLACED_BIT_EXT`, this value is ignored.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkMemoryMapPlacedInfoEXT-flags-09576"
  id="VUID-VkMemoryMapPlacedInfoEXT-flags-09576"></a>
  VUID-VkMemoryMapPlacedInfoEXT-flags-09576  
  If `VkMemoryMapInfoKHR`::`flags` contains
  `VK_MEMORY_MAP_PLACED_BIT_EXT`, `pPlacedAddress` **must** not be
  `NULL`

- <a href="#VUID-VkMemoryMapPlacedInfoEXT-pPlacedAddress-09577"
  id="VUID-VkMemoryMapPlacedInfoEXT-pPlacedAddress-09577"></a>
  VUID-VkMemoryMapPlacedInfoEXT-pPlacedAddress-09577  
  `pPlacedAddress` **must** be aligned to an integer multiple of
  `VkPhysicalDeviceMapMemoryPlacedPropertiesEXT`::`minPlacedMemoryMapAlignment`

- <a href="#VUID-VkMemoryMapPlacedInfoEXT-pPlacedAddress-09578"
  id="VUID-VkMemoryMapPlacedInfoEXT-pPlacedAddress-09578"></a>
  VUID-VkMemoryMapPlacedInfoEXT-pPlacedAddress-09578  
  The address range specified by `pPlacedAddress` and
  `VkMemoryMapInfoKHR`::`size` **must** not overlap any existing Vulkan
  memory object mapping

Valid Usage (Implicit)

- <a href="#VUID-VkMemoryMapPlacedInfoEXT-sType-sType"
  id="VUID-VkMemoryMapPlacedInfoEXT-sType-sType"></a>
  VUID-VkMemoryMapPlacedInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_MEMORY_MAP_PLACED_INFO_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_map_memory_placed](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_map_memory_placed.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkMemoryMapPlacedInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
