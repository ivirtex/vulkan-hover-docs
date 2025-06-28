# PrimitiveLineIndicesEXT(3) Manual Page

## Name

PrimitiveLineIndicesEXT - Indices of line primitives in a mesh shader



## [](#_description)Description

`PrimitiveLineIndicesEXT`

Decorating a variable with the `PrimitiveLineIndicesEXT` decoration will make that variable contain the output array of vertex index values for line primitives.

Valid Usage

- [](#VUID-PrimitiveLineIndicesEXT-PrimitiveLineIndicesEXT-07047)VUID-PrimitiveLineIndicesEXT-PrimitiveLineIndicesEXT-07047  
  The `PrimitiveLineIndicesEXT` decoration **must** be used only within the `MeshEXT` `Execution` `Model`
- [](#VUID-PrimitiveLineIndicesEXT-PrimitiveLineIndicesEXT-07048)VUID-PrimitiveLineIndicesEXT-PrimitiveLineIndicesEXT-07048  
  The `PrimitiveLineIndicesEXT` decoration **must** be used with the `OutputLinesEXT` `Execution` `Mode`
- [](#VUID-PrimitiveLineIndicesEXT-PrimitiveLineIndicesEXT-07049)VUID-PrimitiveLineIndicesEXT-PrimitiveLineIndicesEXT-07049  
  The variable decorated with `PrimitiveLineIndicesEXT` **must** be declared using the `Output` `Storage` `Class`
- [](#VUID-PrimitiveLineIndicesEXT-PrimitiveLineIndicesEXT-07050)VUID-PrimitiveLineIndicesEXT-PrimitiveLineIndicesEXT-07050  
  The variable decorated with `PrimitiveLineIndicesEXT` **must** be declared as an array of two component vector 32-bit integer values
- [](#VUID-PrimitiveLineIndicesEXT-PrimitiveLineIndicesEXT-07051)VUID-PrimitiveLineIndicesEXT-PrimitiveLineIndicesEXT-07051  
  All index values of the array decorated with `PrimitiveLineIndicesEXT` **must** be in the range \[0, N-1], where N is the value specified by the `OutputVertices` `Execution` `Mode`
- [](#VUID-PrimitiveLineIndicesEXT-PrimitiveLineIndicesEXT-07052)VUID-PrimitiveLineIndicesEXT-PrimitiveLineIndicesEXT-07052  
  The size of the array decorated with `PrimitiveLineIndicesEXT` **must** match the value specified by `OutputPrimitivesEXT`

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#PrimitiveLineIndicesEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0