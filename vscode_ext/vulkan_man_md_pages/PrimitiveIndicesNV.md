# PrimitiveIndicesNV(3) Manual Page

## Name

PrimitiveIndicesNV - Indices of primitives in a mesh shader



## <a href="#_description" class="anchor"></a>Description

`PrimitiveIndicesNV`  
Decorating a variable with the `PrimitiveIndicesNV` decoration will make
that variable contain the output array of vertex index values. Depending
on the output primitive type declared using the execution mode, the
indices are split into groups of one (`OutputPoints`), two
(`OutputLinesNV`), or three (`OutputTrianglesNV`) indices and each group
generates a primitive.

Valid Usage

- <a href="#VUID-PrimitiveIndicesNV-PrimitiveIndicesNV-04338"
  id="VUID-PrimitiveIndicesNV-PrimitiveIndicesNV-04338"></a>
  VUID-PrimitiveIndicesNV-PrimitiveIndicesNV-04338  
  The `PrimitiveIndicesNV` decoration **must** be used only within the
  `MeshNV` `Execution` `Model`

- <a href="#VUID-PrimitiveIndicesNV-PrimitiveIndicesNV-04339"
  id="VUID-PrimitiveIndicesNV-PrimitiveIndicesNV-04339"></a>
  VUID-PrimitiveIndicesNV-PrimitiveIndicesNV-04339  
  The variable decorated with `PrimitiveIndicesNV` **must** be declared
  using the `Output` `Storage` `Class`

- <a href="#VUID-PrimitiveIndicesNV-PrimitiveIndicesNV-04340"
  id="VUID-PrimitiveIndicesNV-PrimitiveIndicesNV-04340"></a>
  VUID-PrimitiveIndicesNV-PrimitiveIndicesNV-04340  
  The variable decorated with `PrimitiveIndicesNV` **must** be declared
  as an array of scalar 32-bit integer values

- <a href="#VUID-PrimitiveIndicesNV-PrimitiveIndicesNV-04341"
  id="VUID-PrimitiveIndicesNV-PrimitiveIndicesNV-04341"></a>
  VUID-PrimitiveIndicesNV-PrimitiveIndicesNV-04341  
  All index values of the array decorated with `PrimitiveIndicesNV`
  **must** be in the range \[0, N-1\], where N is the value specified by
  the `OutputVertices` `Execution` `Mode`

- <a href="#VUID-PrimitiveIndicesNV-OutputPoints-04342"
  id="VUID-PrimitiveIndicesNV-OutputPoints-04342"></a>
  VUID-PrimitiveIndicesNV-OutputPoints-04342  
  If the `Execution` `Mode` is `OutputPoints`, then the array decorated
  with `PrimitiveIndicesNV` **must** be the size of the value specified
  by `OutputPrimitivesNV`

- <a href="#VUID-PrimitiveIndicesNV-OutputLinesNV-04343"
  id="VUID-PrimitiveIndicesNV-OutputLinesNV-04343"></a>
  VUID-PrimitiveIndicesNV-OutputLinesNV-04343  
  If the `Execution` `Mode` is `OutputLinesNV`, then the array decorated
  with `PrimitiveIndicesNV` **must** be the size of two times the value
  specified by `OutputPrimitivesNV`

- <a href="#VUID-PrimitiveIndicesNV-OutputTrianglesNV-04344"
  id="VUID-PrimitiveIndicesNV-OutputTrianglesNV-04344"></a>
  VUID-PrimitiveIndicesNV-OutputTrianglesNV-04344  
  If the `Execution` `Mode` is `OutputTrianglesNV`, then the array
  decorated with `PrimitiveIndicesNV` **must** be the size of three
  times the value specified by `OutputPrimitivesNV`

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#PrimitiveIndicesNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
