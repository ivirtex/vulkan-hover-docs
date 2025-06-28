# VkPerTileBeginInfoQCOM(3) Manual Page

## Name

VkPerTileBeginInfoQCOM - Structure specifying per-tile begin information



## [](#_c_specification)C Specification

The `VkPerTileBeginInfoQCOM` structure is defined as:

```c++
// Provided by VK_QCOM_tile_shading
typedef struct VkPerTileBeginInfoQCOM {
    VkStructureType    sType;
    const void*        pNext;
} VkPerTileBeginInfoQCOM;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkPerTileBeginInfoQCOM-sType-sType)VUID-VkPerTileBeginInfoQCOM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PER_TILE_BEGIN_INFO_QCOM`
- [](#VUID-VkPerTileBeginInfoQCOM-pNext-pNext)VUID-VkPerTileBeginInfoQCOM-pNext-pNext  
  `pNext` **must** be `NULL`

## [](#_see_also)See Also

[VK\_QCOM\_tile\_shading](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_QCOM_tile_shading.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCmdBeginPerTileExecutionQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBeginPerTileExecutionQCOM.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPerTileBeginInfoQCOM)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0