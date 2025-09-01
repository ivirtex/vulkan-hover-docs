# TileOffsetQCOM(3) Manual Page

## Name

TileOffsetQCOM - Tile offset of a shader invocation



## [](#_description)Description

`TileOffsetQCOM`

The `TileOffsetQCOM` decoration **can** be applied to a shader input which will be filled with the framebuffer coordinates of the active tile.

When [per-tile execution model](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-tile-shading) is enabled, for the current shader invocation x and y components of `TileOffsetQCOM` reflect the framebuffer integer coordinates of the top-left texel of the tile corresponding to the shader invocation.

Otherwise, the x and y components of `TileOffsetQCOM` are filled with (0,0).

Valid Usage

- [](#VUID-TileOffsetQCOM-TileOffsetQCOM-10626)VUID-TileOffsetQCOM-TileOffsetQCOM-10626  
  The `TileOffsetQCOM` decoration **must** be used only within the `Fragment` `Execution` `Model` or `GLCompute` `Execution` `Model`
- [](#VUID-TileOffsetQCOM-TileOffsetQCOM-10627)VUID-TileOffsetQCOM-TileOffsetQCOM-10627  
  The variable decorated with `TileOffsetQCOM` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-TileOffsetQCOM-TileOffsetQCOM-10628)VUID-TileOffsetQCOM-TileOffsetQCOM-10628  
  The variable decorated with `TileOffsetQCOM` **must** be declared as a two-component vector of 32-bit integer values

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#TileOffsetQCOM).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0