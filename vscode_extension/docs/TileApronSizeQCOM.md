# TileApronSizeQCOM(3) Manual Page

## Name

TileApronSizeQCOM - Tile apron size of a shader invocation



## [](#_description)Description

`TileApronSizeQCOM`

The `TileApronSizeQCOM` decoration **can** be applied to a shader input which will be filled with the width and height of the active tile’s apron.

If [per-tile execution model](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-tile-shading) is enabled for the current shader invocation, and is executing in a dynamic render pass or is executing in a subpass where `VK_SUBPASS_DESCRIPTION_TILE_SHADING_APRON_BIT_QCOM` is included in its `flags`, x and y components of `TileApronSizeQCOM` reflect the with and height of the tile apron corresponding to the shader invocation.

Otherwise, the x and y components of `TileApronSizeQCOM` are filled with (0,0).

Valid Usage

- [](#VUID-TileApronSizeQCOM-TileApronSizeQCOM-10632)VUID-TileApronSizeQCOM-TileApronSizeQCOM-10632  
  The `TileApronSizeQCOM` decoration **must** be used only within the `Fragment` `Execution` `Model` or `GLCompute` `Execution` `Model`
- [](#VUID-TileApronSizeQCOM-TileApronSizeQCOM-10633)VUID-TileApronSizeQCOM-TileApronSizeQCOM-10633  
  The variable decorated with `TileApronSizeQCOM` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-TileApronSizeQCOM-TileApronSizeQCOM-10634)VUID-TileApronSizeQCOM-TileApronSizeQCOM-10634  
  The variable decorated with `TileApronSizeQCOM` **must** be declared as a two-component vector of 32-bit integer values

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#TileApronSizeQCOM).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0