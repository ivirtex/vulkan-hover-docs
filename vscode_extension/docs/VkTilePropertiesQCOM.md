# VkTilePropertiesQCOM(3) Manual Page

## Name

VkTilePropertiesQCOM - Structure holding available tile properties



## [](#_c_specification)C Specification

The `VkTilePropertiesQCOM` structure is defined as:

```c++
// Provided by VK_QCOM_tile_properties
typedef struct VkTilePropertiesQCOM {
    VkStructureType    sType;
    void*              pNext;
    VkExtent3D         tileSize;
    VkExtent2D         apronSize;
    VkOffset2D         origin;
} VkTilePropertiesQCOM;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `tileSize` is the dimensions of a tile, with width and height describing the width and height of a tile in pixels, and depth corresponding to the number of slices the tile spans.
- `apronSize` is the dimension of the [apron](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-tile-shading-aprons).
- `origin` is the top-left corner of the first tile in attachment space.

## [](#_description)Description

All tiles will be tightly packed around the first tile, with edges being multiples of tile width and/or height from the origin.

The `tileSize` is guaranteed to be a multiple of [`tileGranularity`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-tileGranularity).

Valid Usage (Implicit)

- [](#VUID-VkTilePropertiesQCOM-sType-sType)VUID-VkTilePropertiesQCOM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_TILE_PROPERTIES_QCOM`
- [](#VUID-VkTilePropertiesQCOM-pNext-pNext)VUID-VkTilePropertiesQCOM-pNext-pNext  
  `pNext` **must** be `NULL`

## [](#_see_also)See Also

[VK\_QCOM\_tile\_properties](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_QCOM_tile_properties.html), [VkExtent2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtent2D.html), [VkExtent3D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtent3D.html), [VkOffset2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOffset2D.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetDynamicRenderingTilePropertiesQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDynamicRenderingTilePropertiesQCOM.html), [vkGetFramebufferTilePropertiesQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetFramebufferTilePropertiesQCOM.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkTilePropertiesQCOM)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0