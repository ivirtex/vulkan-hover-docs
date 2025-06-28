# PointCoord(3) Manual Page

## Name

PointCoord - Fragment coordinates in screen-space within a point primitive



## [](#_description)Description

`PointCoord`

Decorating a variable with the `PointCoord` built-in decoration will make that variable contain the coordinate of the current fragment within the point being rasterized, normalized to the size of the point with origin in the upper left corner of the point, as described in [Basic Point Rasterization](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#primsrast-points-basic). If the primitive the fragment shader invocation belongs to is not a point, then the variable decorated with `PointCoord` contains an undefined value.

Note

Depending on how the point is rasterized, `PointCoord` **may** never reach (0,0) or (1,1).

Valid Usage

- [](#VUID-PointCoord-PointCoord-04311)VUID-PointCoord-PointCoord-04311  
  The `PointCoord` decoration **must** be used only within the `Fragment` `Execution` `Model`
- [](#VUID-PointCoord-PointCoord-04312)VUID-PointCoord-PointCoord-04312  
  The variable decorated with `PointCoord` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-PointCoord-PointCoord-04313)VUID-PointCoord-PointCoord-04313  
  The variable decorated with `PointCoord` **must** be declared as a two-component vector of 32-bit floating-point values

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#PointCoord)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0