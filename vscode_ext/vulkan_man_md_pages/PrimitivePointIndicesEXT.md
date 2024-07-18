# PrimitivePointIndicesEXT(3) Manual Page

## Name

PrimitivePointIndicesEXT - Indices of point primitives in a mesh shader



## <a href="#_description" class="anchor"></a>Description

`PrimitivePointIndicesEXT`  
Decorating a variable with the `PrimitivePointIndicesEXT` decoration
will make that variable contain the output array of vertex index values
for point primitives.

Valid Usage

- <a href="#VUID-PrimitivePointIndicesEXT-PrimitivePointIndicesEXT-07041"
  id="VUID-PrimitivePointIndicesEXT-PrimitivePointIndicesEXT-07041"></a>
  VUID-PrimitivePointIndicesEXT-PrimitivePointIndicesEXT-07041  
  The `PrimitivePointIndicesEXT` decoration **must** be used only within
  the `MeshEXT` `Execution` `Model`

- <a href="#VUID-PrimitivePointIndicesEXT-PrimitivePointIndicesEXT-07042"
  id="VUID-PrimitivePointIndicesEXT-PrimitivePointIndicesEXT-07042"></a>
  VUID-PrimitivePointIndicesEXT-PrimitivePointIndicesEXT-07042  
  The `PrimitivePointIndicesEXT` decoration **must** be used with the
  `OutputPoints` `Execution` `Mode`

- <a href="#VUID-PrimitivePointIndicesEXT-PrimitivePointIndicesEXT-07043"
  id="VUID-PrimitivePointIndicesEXT-PrimitivePointIndicesEXT-07043"></a>
  VUID-PrimitivePointIndicesEXT-PrimitivePointIndicesEXT-07043  
  The variable decorated with `PrimitivePointIndicesEXT` **must** be
  declared using the `Output` `Storage` `Class`

- <a href="#VUID-PrimitivePointIndicesEXT-PrimitivePointIndicesEXT-07044"
  id="VUID-PrimitivePointIndicesEXT-PrimitivePointIndicesEXT-07044"></a>
  VUID-PrimitivePointIndicesEXT-PrimitivePointIndicesEXT-07044  
  The variable decorated with `PrimitivePointIndicesEXT` **must** be
  declared as an array of scalar 32-bit integer values

- <a href="#VUID-PrimitivePointIndicesEXT-PrimitivePointIndicesEXT-07045"
  id="VUID-PrimitivePointIndicesEXT-PrimitivePointIndicesEXT-07045"></a>
  VUID-PrimitivePointIndicesEXT-PrimitivePointIndicesEXT-07045  
  All index values of the array decorated with
  `PrimitivePointIndicesEXT` **must** be in the range \[0, N-1\], where
  N is the value specified by the `OutputVertices` `Execution` `Mode`

- <a href="#VUID-PrimitivePointIndicesEXT-PrimitivePointIndicesEXT-07046"
  id="VUID-PrimitivePointIndicesEXT-PrimitivePointIndicesEXT-07046"></a>
  VUID-PrimitivePointIndicesEXT-PrimitivePointIndicesEXT-07046  
  The size of the array decorated with `PrimitivePointIndicesEXT`
  **must** match the value specified by `OutputPrimitivesEXT`

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#PrimitivePointIndicesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
