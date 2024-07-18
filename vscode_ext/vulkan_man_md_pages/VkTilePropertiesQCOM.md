# VkTilePropertiesQCOM(3) Manual Page

## Name

VkTilePropertiesQCOM - Structure holding available tile properties



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkTilePropertiesQCOM` structure is defined as:

``` c
// Provided by VK_QCOM_tile_properties
typedef struct VkTilePropertiesQCOM {
    VkStructureType    sType;
    void*              pNext;
    VkExtent3D         tileSize;
    VkExtent2D         apronSize;
    VkOffset2D         origin;
} VkTilePropertiesQCOM;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `tileSize` is the dimensions of a tile, with width and height
  describing the width and height of a tile in pixels, and depth
  corresponding to the number of slices the tile spans.

- `apronSize` is the dimension of the apron.

- `origin` is the top-left corner of the first tile in attachment space.

## <a href="#_description" class="anchor"></a>Description

All tiles will be tightly packed around the first tile, with edges being
multiples of tile width and/or height from the origin.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>Reported value for <code>apronSize</code> will be zero and its
functionality will be described in a future extension.</p></td>
</tr>
</tbody>
</table>

Valid Usage (Implicit)

- <a href="#VUID-VkTilePropertiesQCOM-sType-sType"
  id="VUID-VkTilePropertiesQCOM-sType-sType"></a>
  VUID-VkTilePropertiesQCOM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_TILE_PROPERTIES_QCOM`

- <a href="#VUID-VkTilePropertiesQCOM-pNext-pNext"
  id="VUID-VkTilePropertiesQCOM-pNext-pNext"></a>
  VUID-VkTilePropertiesQCOM-pNext-pNext  
  `pNext` **must** be `NULL`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_QCOM_tile_properties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_QCOM_tile_properties.html),
[VkExtent2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtent2D.html), [VkExtent3D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtent3D.html),
[VkOffset2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOffset2D.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetDynamicRenderingTilePropertiesQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDynamicRenderingTilePropertiesQCOM.html),
[vkGetFramebufferTilePropertiesQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetFramebufferTilePropertiesQCOM.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkTilePropertiesQCOM"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
