# PrimitiveTriangleIndicesEXT(3) Manual Page

## Name

PrimitiveTriangleIndicesEXT - Indices of triangle primitives in a mesh shader



## [](#_description)Description

`PrimitiveTriangleIndicesEXT`

Decorating a variable with the `PrimitiveTriangleIndicesEXT` decoration will make that variable contain the output array of vertex index values for triangle primitives.

Valid Usage

- [](#VUID-PrimitiveTriangleIndicesEXT-PrimitiveTriangleIndicesEXT-07053)VUID-PrimitiveTriangleIndicesEXT-PrimitiveTriangleIndicesEXT-07053  
  The `PrimitiveTriangleIndicesEXT` decoration **must** be used only within the `MeshEXT` `Execution` `Model`
- [](#VUID-PrimitiveTriangleIndicesEXT-PrimitiveTriangleIndicesEXT-07054)VUID-PrimitiveTriangleIndicesEXT-PrimitiveTriangleIndicesEXT-07054  
  The `PrimitiveTriangleIndicesEXT` decoration **must** be used with the `OutputTrianglesEXT` `Execution` `Mode`
- [](#VUID-PrimitiveTriangleIndicesEXT-PrimitiveTriangleIndicesEXT-07055)VUID-PrimitiveTriangleIndicesEXT-PrimitiveTriangleIndicesEXT-07055  
  The variable decorated with `PrimitiveTriangleIndicesEXT` **must** be declared using the `Output` `Storage` `Class`
- [](#VUID-PrimitiveTriangleIndicesEXT-PrimitiveTriangleIndicesEXT-07056)VUID-PrimitiveTriangleIndicesEXT-PrimitiveTriangleIndicesEXT-07056  
  The variable decorated with `PrimitiveTriangleIndicesEXT` **must** be declared as an array of three component vector 32-bit integer values
- [](#VUID-PrimitiveTriangleIndicesEXT-PrimitiveTriangleIndicesEXT-07057)VUID-PrimitiveTriangleIndicesEXT-PrimitiveTriangleIndicesEXT-07057  
  All index values of the array decorated with `PrimitiveTriangleIndicesEXT` **must** be in the range \[0, N-1], where N is the value specified by the `OutputVertices` `Execution` `Mode`
- [](#VUID-PrimitiveTriangleIndicesEXT-PrimitiveTriangleIndicesEXT-07058)VUID-PrimitiveTriangleIndicesEXT-PrimitiveTriangleIndicesEXT-07058  
  The size of the array decorated with `PrimitiveTriangleIndicesEXT` **must** match the value specified by `OutputPrimitivesEXT`

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#PrimitiveTriangleIndicesEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0