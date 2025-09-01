# TileDimensionQCOM(3) Manual Page

## Name

TileDimensionQCOM - Tile offset of a shader invocation



## [](#_description)Description

`TileDimensionQCOM`

The `TileDimensionQCOM` decoration **can** be applied to a shader input which will be filled with the width and height of the active tile.

When [per-tile execution model](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-tile-shading) is enabled, for the current shader invocation x and y components of `TileDimensionQCOM` reflect the with and height of the tile corresponding to the shader invocation.

Otherwise, the x and y components of `TileDimensionQCOM` are filled with (0,0).

Valid Usage

- [](#VUID-TileDimensionQCOM-TileDimensionQCOM-10629)VUID-TileDimensionQCOM-TileDimensionQCOM-10629  
  The `TileDimensionQCOM` decoration **must** be used only within the `Fragment` `Execution` `Model` or `GLCompute` `Execution` `Model`
- [](#VUID-TileDimensionQCOM-TileDimensionQCOM-10630)VUID-TileDimensionQCOM-TileDimensionQCOM-10630  
  The variable decorated with `TileDimensionQCOM` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-TileDimensionQCOM-TileDimensionQCOM-10631)VUID-TileDimensionQCOM-TileDimensionQCOM-10631  
  The variable decorated with `TileDimensionQCOM` **must** be declared as a two-component vector of 32-bit integer values

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#TileDimensionQCOM).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0