# VkPhysicalDeviceTileMemoryHeapPropertiesQCOM(3) Manual Page

## Name

VkPhysicalDeviceTileMemoryHeapPropertiesQCOM - Structure describing tile memory heap properties that can be supported by an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceTileMemoryHeapPropertiesQCOM` structure is defined as:

```c++
// Provided by VK_QCOM_tile_memory_heap
typedef struct VkPhysicalDeviceTileMemoryHeapPropertiesQCOM {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           queueSubmitBoundary;
    VkBool32           tileBufferTransfers;
} VkPhysicalDeviceTileMemoryHeapPropertiesQCOM;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`queueSubmitBoundary` is a boolean describing if tile memory becomes undefined at a queue submit boundary instead of the default command buffer submission batch boundary.
- []()`tileBufferTransfers` is a boolean describing if buffers bound to tile memory support transfer operations.

## [](#_description)Description

If the `VkPhysicalDeviceTileMemoryHeapPropertiesQCOM` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceTileMemoryHeapPropertiesQCOM-sType-sType)VUID-VkPhysicalDeviceTileMemoryHeapPropertiesQCOM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_TILE_MEMORY_HEAP_PROPERTIES_QCOM`

## [](#_see_also)See Also

[VK\_QCOM\_tile\_memory\_heap](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_QCOM_tile_memory_heap.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceTileMemoryHeapPropertiesQCOM)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0