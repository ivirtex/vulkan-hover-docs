# VkTileMemoryBindInfoQCOM(3) Manual Page

## Name

VkTileMemoryBindInfoQCOM - Structure specifying tile memory to bind



## [](#_c_specification)C Specification

The `VkTileMemoryBindInfoQCOM` structure is defined as:

```c++
// Provided by VK_QCOM_tile_memory_heap
typedef struct VkTileMemoryBindInfoQCOM {
    VkStructureType    sType;
    const void*        pNext;
    VkDeviceMemory     memory;
} VkTileMemoryBindInfoQCOM;
```

## [](#_members)Members

- `memory` is the tile memory object to be bound.

## [](#_description)Description

`memory` is used to bind this memory object to tile memory for all subsequent commands in the `commandBuffer`. Tile memory contents for ranges in the heap outside the bound `memory` are discarded and become undefined for the [active *tile memory scope*](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-tile-heaps) if an action command is executed.

For secondary command buffers executing within a render pass instance, the active bound tile memory object is provided with this structure included in the `pNext` chain of `VkCommandBufferInheritanceInfo`.

If this structure was not specified since recording started for `commandBuffer`, no tile memory is bound to the command buffer and all contents become undefined for the *tile memory scope* if an action command is executed.

Valid Usage

- [](#VUID-VkTileMemoryBindInfoQCOM-memory-10726)VUID-VkTileMemoryBindInfoQCOM-memory-10726  
  `memory` **must** have been allocated from a [VkMemoryHeap](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryHeap.html) with the `VK_MEMORY_HEAP_TILE_MEMORY_BIT_QCOM` property

Valid Usage (Implicit)

- [](#VUID-VkTileMemoryBindInfoQCOM-sType-sType)VUID-VkTileMemoryBindInfoQCOM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_TILE_MEMORY_BIND_INFO_QCOM`
- [](#VUID-VkTileMemoryBindInfoQCOM-memory-parameter)VUID-VkTileMemoryBindInfoQCOM-memory-parameter  
  `memory` **must** be a valid [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) handle

## [](#_see_also)See Also

[VK\_QCOM\_tile\_memory\_heap](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_QCOM_tile_memory_heap.html), [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCmdBindTileMemoryQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBindTileMemoryQCOM.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkTileMemoryBindInfoQCOM).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0