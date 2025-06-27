# VkMappedMemoryRange(3) Manual Page

## Name

VkMappedMemoryRange - Structure specifying a mapped memory range



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkMappedMemoryRange` structure is defined as:

``` c
// Provided by VK_VERSION_1_0
typedef struct VkMappedMemoryRange {
    VkStructureType    sType;
    const void*        pNext;
    VkDeviceMemory     memory;
    VkDeviceSize       offset;
    VkDeviceSize       size;
} VkMappedMemoryRange;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `memory` is the memory object to which this range belongs.

- `offset` is the zero-based byte offset from the beginning of the
  memory object.

- `size` is either the size of range, or `VK_WHOLE_SIZE` to affect the
  range from `offset` to the end of the current mapping of the
  allocation.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkMappedMemoryRange-memory-00684"
  id="VUID-VkMappedMemoryRange-memory-00684"></a>
  VUID-VkMappedMemoryRange-memory-00684  
  `memory` **must** be currently host mapped

- <a href="#VUID-VkMappedMemoryRange-size-00685"
  id="VUID-VkMappedMemoryRange-size-00685"></a>
  VUID-VkMappedMemoryRange-size-00685  
  If `size` is not equal to `VK_WHOLE_SIZE`, `offset` and `size`
  **must** specify a range contained within the currently mapped range
  of `memory`

- <a href="#VUID-VkMappedMemoryRange-size-00686"
  id="VUID-VkMappedMemoryRange-size-00686"></a>
  VUID-VkMappedMemoryRange-size-00686  
  If `size` is equal to `VK_WHOLE_SIZE`, `offset` **must** be within the
  currently mapped range of `memory`

- <a href="#VUID-VkMappedMemoryRange-offset-00687"
  id="VUID-VkMappedMemoryRange-offset-00687"></a>
  VUID-VkMappedMemoryRange-offset-00687  
  `offset` **must** be a multiple of
  [VkPhysicalDeviceLimits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceLimits.html)::`nonCoherentAtomSize`

- <a href="#VUID-VkMappedMemoryRange-size-01389"
  id="VUID-VkMappedMemoryRange-size-01389"></a>
  VUID-VkMappedMemoryRange-size-01389  
  If `size` is equal to `VK_WHOLE_SIZE`, the end of the current mapping
  of `memory` **must** either be a multiple of
  [VkPhysicalDeviceLimits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceLimits.html)::`nonCoherentAtomSize`
  bytes from the beginning of the memory object, or be equal to the end
  of the memory object

- <a href="#VUID-VkMappedMemoryRange-size-01390"
  id="VUID-VkMappedMemoryRange-size-01390"></a>
  VUID-VkMappedMemoryRange-size-01390  
  If `size` is not equal to `VK_WHOLE_SIZE`, `size` **must** either be a
  multiple of
  [VkPhysicalDeviceLimits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceLimits.html)::`nonCoherentAtomSize`,
  or `offset` plus `size` **must** equal the size of `memory`

Valid Usage (Implicit)

- <a href="#VUID-VkMappedMemoryRange-sType-sType"
  id="VUID-VkMappedMemoryRange-sType-sType"></a>
  VUID-VkMappedMemoryRange-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_MAPPED_MEMORY_RANGE`

- <a href="#VUID-VkMappedMemoryRange-pNext-pNext"
  id="VUID-VkMappedMemoryRange-pNext-pNext"></a>
  VUID-VkMappedMemoryRange-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkMappedMemoryRange-memory-parameter"
  id="VUID-VkMappedMemoryRange-memory-parameter"></a>
  VUID-VkMappedMemoryRange-memory-parameter  
  `memory` **must** be a valid [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html)
  handle

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html),
[VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkFlushMappedMemoryRanges](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkFlushMappedMemoryRanges.html),
[vkInvalidateMappedMemoryRanges](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkInvalidateMappedMemoryRanges.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkMappedMemoryRange"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
