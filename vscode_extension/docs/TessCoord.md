# TessCoord(3) Manual Page

## Name

TessCoord - Barycentric coordinate of a tessellated vertex within a patch



## [](#_description)Description

`TessCoord`

Decorating a variable with the `TessCoord` built-in decoration will make that variable contain the three-dimensional (u,v,w) barycentric coordinate of the tessellated vertex within the patch. u, v, and w are in the range \[0,1] and vary linearly across the primitive being subdivided. For the tessellation modes of `Quads` or `IsoLines`, the third component is always zero.

Valid Usage

- [](#VUID-TessCoord-TessCoord-04387)VUID-TessCoord-TessCoord-04387  
  The `TessCoord` decoration **must** be used only within the `TessellationEvaluation` `Execution` `Model`
- [](#VUID-TessCoord-TessCoord-04388)VUID-TessCoord-TessCoord-04388  
  The variable decorated with `TessCoord` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-TessCoord-TessCoord-04389)VUID-TessCoord-TessCoord-04389  
  The variable decorated with `TessCoord` **must** be declared as a three-component vector of 32-bit floating-point values

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#TessCoord)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0