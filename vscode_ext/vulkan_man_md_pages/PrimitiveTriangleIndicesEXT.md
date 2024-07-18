# PrimitiveTriangleIndicesEXT(3) Manual Page

## Name

PrimitiveTriangleIndicesEXT - Indices of triangle primitives in a mesh
shader



## <a href="#_description" class="anchor"></a>Description

`PrimitiveTriangleIndicesEXT`  
Decorating a variable with the `PrimitiveTriangleIndicesEXT` decoration
will make that variable contain the output array of vertex index values
for triangle primitives.

Valid Usage

- <a
  href="#VUID-PrimitiveTriangleIndicesEXT-PrimitiveTriangleIndicesEXT-07053"
  id="VUID-PrimitiveTriangleIndicesEXT-PrimitiveTriangleIndicesEXT-07053"></a>
  VUID-PrimitiveTriangleIndicesEXT-PrimitiveTriangleIndicesEXT-07053  
  The `PrimitiveTriangleIndicesEXT` decoration **must** be used only
  within the `MeshEXT` `Execution` `Model`

- <a
  href="#VUID-PrimitiveTriangleIndicesEXT-PrimitiveTriangleIndicesEXT-07054"
  id="VUID-PrimitiveTriangleIndicesEXT-PrimitiveTriangleIndicesEXT-07054"></a>
  VUID-PrimitiveTriangleIndicesEXT-PrimitiveTriangleIndicesEXT-07054  
  The `PrimitiveTriangleIndicesEXT` decoration **must** be used with the
  `OutputTrianglesEXT` `Execution` `Mode`

- <a
  href="#VUID-PrimitiveTriangleIndicesEXT-PrimitiveTriangleIndicesEXT-07055"
  id="VUID-PrimitiveTriangleIndicesEXT-PrimitiveTriangleIndicesEXT-07055"></a>
  VUID-PrimitiveTriangleIndicesEXT-PrimitiveTriangleIndicesEXT-07055  
  The variable decorated with `PrimitiveTriangleIndicesEXT` **must** be
  declared using the `Output` `Storage` `Class`

- <a
  href="#VUID-PrimitiveTriangleIndicesEXT-PrimitiveTriangleIndicesEXT-07056"
  id="VUID-PrimitiveTriangleIndicesEXT-PrimitiveTriangleIndicesEXT-07056"></a>
  VUID-PrimitiveTriangleIndicesEXT-PrimitiveTriangleIndicesEXT-07056  
  The variable decorated with `PrimitiveTriangleIndicesEXT` **must** be
  declared as an array of three component vector 32-bit integer values

- <a
  href="#VUID-PrimitiveTriangleIndicesEXT-PrimitiveTriangleIndicesEXT-07057"
  id="VUID-PrimitiveTriangleIndicesEXT-PrimitiveTriangleIndicesEXT-07057"></a>
  VUID-PrimitiveTriangleIndicesEXT-PrimitiveTriangleIndicesEXT-07057  
  All index values of the array decorated with
  `PrimitiveTriangleIndicesEXT` **must** be in the range \[0, N-1\],
  where N is the value specified by the `OutputVertices` `Execution`
  `Mode`

- <a
  href="#VUID-PrimitiveTriangleIndicesEXT-PrimitiveTriangleIndicesEXT-07058"
  id="VUID-PrimitiveTriangleIndicesEXT-PrimitiveTriangleIndicesEXT-07058"></a>
  VUID-PrimitiveTriangleIndicesEXT-PrimitiveTriangleIndicesEXT-07058  
  The size of the array decorated with `PrimitiveTriangleIndicesEXT`
  **must** match the value specified by `OutputPrimitivesEXT`

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#PrimitiveTriangleIndicesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
