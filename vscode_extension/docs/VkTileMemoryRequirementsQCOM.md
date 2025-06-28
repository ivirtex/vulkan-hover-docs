# VkTileMemoryRequirementsQCOM(3) Manual Page

## Name

VkTileMemoryRequirementsQCOM - Structure specifying tile memory requirements



## [](#_c_specification)C Specification

To determine the tile memory allocation requirements of a buffer or image resource, add a `VkTileMemoryRequirementsQCOM` structure to the `pNext` chain of the `VkMemoryRequirements2` structure passed as the `pMemoryRequirements` parameter of [vkGetBufferMemoryRequirements2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetBufferMemoryRequirements2.html) or [vkGetImageMemoryRequirements2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetImageMemoryRequirements2.html), respectively. The `VkTileMemoryRequirementsQCOM` structure is defined as:

```c++
// Provided by VK_QCOM_tile_memory_heap
typedef struct VkTileMemoryRequirementsQCOM {
    VkStructureType    sType;
    void*              pNext;
    VkDeviceSize       size;
    VkDeviceSize       alignment;
} VkTileMemoryRequirementsQCOM;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `size` size is the size, in bytes, of the tile memory allocation required for the resource.
- `alignment` is the alignment, in bytes, of the offset within the tile memory allocation required for the resource.

## [](#_description)Description

The `size` and `alignment` **must** be used when the resource is bound to a [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) object that was allocated from a [VkMemoryType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryType.html) that has a `heapIndex` that corresponds to a [VkMemoryHeap](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryHeap.html) with the `VK_MEMORY_HEAP_TILE_MEMORY_BIT_QCOM` property.

If the resource cannot be bound to tile memory, then `size` and `alignment` is filled with zero by the implementation.

Valid Usage (Implicit)

- [](#VUID-VkTileMemoryRequirementsQCOM-sType-sType)VUID-VkTileMemoryRequirementsQCOM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_TILE_MEMORY_REQUIREMENTS_QCOM`

## [](#_see_also)See Also

[VK\_QCOM\_tile\_memory\_heap](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_QCOM_tile_memory_heap.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkTileMemoryRequirementsQCOM)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0