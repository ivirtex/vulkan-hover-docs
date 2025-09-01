# VkMappedMemoryRange(3) Manual Page

## Name

VkMappedMemoryRange - Structure specifying a mapped memory range



## [](#_c_specification)C Specification

The `VkMappedMemoryRange` structure is defined as:

```c++
// Provided by VK_VERSION_1_0
typedef struct VkMappedMemoryRange {
    VkStructureType    sType;
    const void*        pNext;
    VkDeviceMemory     memory;
    VkDeviceSize       offset;
    VkDeviceSize       size;
} VkMappedMemoryRange;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `memory` is the memory object to which this range belongs.
- `offset` is the zero-based byte offset from the beginning of the memory object.
- `size` is either the size of range, or `VK_WHOLE_SIZE` to affect the range from `offset` to the end of the current mapping of the allocation.

## [](#_description)Description

Valid Usage

- [](#VUID-VkMappedMemoryRange-memory-00684)VUID-VkMappedMemoryRange-memory-00684  
  `memory` **must** be currently host mapped
- [](#VUID-VkMappedMemoryRange-size-00685)VUID-VkMappedMemoryRange-size-00685  
  If `size` is not equal to `VK_WHOLE_SIZE`, `offset` and `size` **must** specify a range contained within the currently mapped range of `memory`
- [](#VUID-VkMappedMemoryRange-size-00686)VUID-VkMappedMemoryRange-size-00686  
  If `size` is equal to `VK_WHOLE_SIZE`, `offset` **must** be within the currently mapped range of `memory`
- [](#VUID-VkMappedMemoryRange-offset-00687)VUID-VkMappedMemoryRange-offset-00687  
  `offset` **must** be a multiple of [VkPhysicalDeviceLimits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceLimits.html)::`nonCoherentAtomSize`
- [](#VUID-VkMappedMemoryRange-size-01389)VUID-VkMappedMemoryRange-size-01389  
  If `size` is equal to `VK_WHOLE_SIZE`, the end of the current mapping of `memory` **must** either be a multiple of [VkPhysicalDeviceLimits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceLimits.html)::`nonCoherentAtomSize` bytes from the beginning of the memory object, or be equal to the end of the memory object
- [](#VUID-VkMappedMemoryRange-size-01390)VUID-VkMappedMemoryRange-size-01390  
  If `size` is not equal to `VK_WHOLE_SIZE`, `size` **must** either be a multiple of [VkPhysicalDeviceLimits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceLimits.html)::`nonCoherentAtomSize`, or `offset` plus `size` **must** equal the size of `memory`

Valid Usage (Implicit)

- [](#VUID-VkMappedMemoryRange-sType-sType)VUID-VkMappedMemoryRange-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_MAPPED_MEMORY_RANGE`
- [](#VUID-VkMappedMemoryRange-pNext-pNext)VUID-VkMappedMemoryRange-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkMappedMemoryRange-memory-parameter)VUID-VkMappedMemoryRange-memory-parameter  
  `memory` **must** be a valid [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) handle

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkFlushMappedMemoryRanges](https://registry.khronos.org/vulkan/specs/latest/man/html/vkFlushMappedMemoryRanges.html), [vkInvalidateMappedMemoryRanges](https://registry.khronos.org/vulkan/specs/latest/man/html/vkInvalidateMappedMemoryRanges.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkMappedMemoryRange).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0