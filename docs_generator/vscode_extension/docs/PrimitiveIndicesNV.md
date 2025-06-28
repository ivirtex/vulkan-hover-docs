# PrimitiveIndicesNV(3) Manual Page

## Name

PrimitiveIndicesNV - Indices of primitives in a mesh shader



## [](#_description)Description

`PrimitiveIndicesNV`

Decorating a variable with the `PrimitiveIndicesNV` decoration will make that variable contain the output array of vertex index values. Depending on the output primitive type declared using the execution mode, the indices are split into groups of one (`OutputPoints`), two (`OutputLinesNV`), or three (`OutputTrianglesNV`) indices and each group generates a primitive.

Valid Usage

- [](#VUID-PrimitiveIndicesNV-PrimitiveIndicesNV-04338)VUID-PrimitiveIndicesNV-PrimitiveIndicesNV-04338  
  The `PrimitiveIndicesNV` decoration **must** be used only within the `MeshNV` `Execution` `Model`
- [](#VUID-PrimitiveIndicesNV-PrimitiveIndicesNV-04339)VUID-PrimitiveIndicesNV-PrimitiveIndicesNV-04339  
  The variable decorated with `PrimitiveIndicesNV` **must** be declared using the `Output` `Storage` `Class`
- [](#VUID-PrimitiveIndicesNV-PrimitiveIndicesNV-04340)VUID-PrimitiveIndicesNV-PrimitiveIndicesNV-04340  
  The variable decorated with `PrimitiveIndicesNV` **must** be declared as an array of scalar 32-bit integer values
- [](#VUID-PrimitiveIndicesNV-PrimitiveIndicesNV-04341)VUID-PrimitiveIndicesNV-PrimitiveIndicesNV-04341  
  All index values of the array decorated with `PrimitiveIndicesNV` **must** be in the range \[0, N-1], where N is the value specified by the `OutputVertices` `Execution` `Mode`
- [](#VUID-PrimitiveIndicesNV-OutputPoints-04342)VUID-PrimitiveIndicesNV-OutputPoints-04342  
  If the `Execution` `Mode` is `OutputPoints`, then the array decorated with `PrimitiveIndicesNV` **must** be the size of the value specified by `OutputPrimitivesNV`
- [](#VUID-PrimitiveIndicesNV-OutputLinesNV-04343)VUID-PrimitiveIndicesNV-OutputLinesNV-04343  
  If the `Execution` `Mode` is `OutputLinesNV`, then the array decorated with `PrimitiveIndicesNV` **must** be the size of two times the value specified by `OutputPrimitivesNV`
- [](#VUID-PrimitiveIndicesNV-OutputTrianglesNV-04344)VUID-PrimitiveIndicesNV-OutputTrianglesNV-04344  
  If the `Execution` `Mode` is `OutputTrianglesNV`, then the array decorated with `PrimitiveIndicesNV` **must** be the size of three times the value specified by `OutputPrimitivesNV`

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#PrimitiveIndicesNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0