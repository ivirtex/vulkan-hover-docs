# VkTileMemorySizeInfoQCOM(3) Manual Page

## Name

VkTileMemorySizeInfoQCOM - Structure describing tile memory size in use in a render pass instance



## [](#_c_specification)C Specification

The tile properties queried using [VK\_QCOM\_tile\_properties](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_QCOM_tile_properties.html) depend on the size of the reserved tile memory by the application. This size **can** be specified by the following structure to [VkRenderPassCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassCreateInfo.html), [VkRenderPassCreateInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassCreateInfo2.html) , or [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingInfo.html) to specify the reserved tile memory size for the render pass object.

For dynamic render passes, this structure **can** be attached to the `pNext` member of [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingInfo.html) passed to [vkGetDynamicRenderingTilePropertiesQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDynamicRenderingTilePropertiesQCOM.html).

The `VkTileMemorySizeInfoQCOM` structure is defined as:

```c++
// Provided by VK_QCOM_tile_memory_heap with VK_QCOM_tile_properties
typedef struct VkTileMemorySizeInfoQCOM {
    VkStructureType    sType;
    const void*        pNext;
    VkDeviceSize       size;
} VkTileMemorySizeInfoQCOM;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `size` is the size in bytes of tile memory used by the render pass or preserved for later use.

## [](#_description)Description

The returned tile properties are invalid if the `size` is not equal to the [bound tile memory’s](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-bind-tile-memory) allocation size when the render pass is executed.

If this structure is not provided, the `size` of the reserved region defaults to `0`.

Note

Tile memory is reserved for application use by binding tile memory objects to the command buffer.

The size provided by this command is informational only for use when evaluating tile properties. If the application does not need to query the tile properties, then this size **can** be safely omitted.

Valid Usage

- [](#VUID-VkTileMemorySizeInfoQCOM-size-10729)VUID-VkTileMemorySizeInfoQCOM-size-10729  
  `size` must be less than or equal to the largest size memory heap with the `VK_MEMORY_HEAP_TILE_MEMORY_BIT_QCOM` property

Valid Usage (Implicit)

- [](#VUID-VkTileMemorySizeInfoQCOM-sType-sType)VUID-VkTileMemorySizeInfoQCOM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_TILE_MEMORY_SIZE_INFO_QCOM`

## [](#_see_also)See Also

[VK\_QCOM\_tile\_memory\_heap](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_QCOM_tile_memory_heap.html), [VK\_QCOM\_tile\_properties](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_QCOM_tile_properties.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkTileMemorySizeInfoQCOM).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0