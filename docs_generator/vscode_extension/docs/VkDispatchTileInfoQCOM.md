# VkDispatchTileInfoQCOM(3) Manual Page

## Name

VkDispatchTileInfoQCOM - Structure specifying dispatch tile info



## [](#_c_specification)C Specification

The `VkDispatchTileInfoQCOM` structure is defined as:

```c++
// Provided by VK_QCOM_tile_shading
typedef struct VkDispatchTileInfoQCOM {
    VkStructureType    sType;
    const void*        pNext;
} VkDispatchTileInfoQCOM;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkDispatchTileInfoQCOM-sType-sType)VUID-VkDispatchTileInfoQCOM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DISPATCH_TILE_INFO_QCOM`
- [](#VUID-VkDispatchTileInfoQCOM-pNext-pNext)VUID-VkDispatchTileInfoQCOM-pNext-pNext  
  `pNext` **must** be `NULL`

## [](#_see_also)See Also

[VK\_QCOM\_tile\_shading](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_QCOM_tile_shading.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCmdDispatchTileQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDispatchTileQCOM.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDispatchTileInfoQCOM)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0