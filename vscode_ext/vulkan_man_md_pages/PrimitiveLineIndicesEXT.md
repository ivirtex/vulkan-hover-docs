# PrimitiveLineIndicesEXT(3) Manual Page

## Name

PrimitiveLineIndicesEXT - Indices of line primitives in a mesh shader



## <a href="#_description" class="anchor"></a>Description

`PrimitiveLineIndicesEXT`  
Decorating a variable with the `PrimitiveLineIndicesEXT` decoration will
make that variable contain the output array of vertex index values for
line primitives.

Valid Usage

- <a href="#VUID-PrimitiveLineIndicesEXT-PrimitiveLineIndicesEXT-07047"
  id="VUID-PrimitiveLineIndicesEXT-PrimitiveLineIndicesEXT-07047"></a>
  VUID-PrimitiveLineIndicesEXT-PrimitiveLineIndicesEXT-07047  
  The `PrimitiveLineIndicesEXT` decoration **must** be used only within
  the `MeshEXT` `Execution` `Model`

- <a href="#VUID-PrimitiveLineIndicesEXT-PrimitiveLineIndicesEXT-07048"
  id="VUID-PrimitiveLineIndicesEXT-PrimitiveLineIndicesEXT-07048"></a>
  VUID-PrimitiveLineIndicesEXT-PrimitiveLineIndicesEXT-07048  
  The `PrimitiveLineIndicesEXT` decoration **must** be used with the
  `OutputLinesEXT` `Execution` `Mode`

- <a href="#VUID-PrimitiveLineIndicesEXT-PrimitiveLineIndicesEXT-07049"
  id="VUID-PrimitiveLineIndicesEXT-PrimitiveLineIndicesEXT-07049"></a>
  VUID-PrimitiveLineIndicesEXT-PrimitiveLineIndicesEXT-07049  
  The variable decorated with `PrimitiveLineIndicesEXT` **must** be
  declared using the `Output` `Storage` `Class`

- <a href="#VUID-PrimitiveLineIndicesEXT-PrimitiveLineIndicesEXT-07050"
  id="VUID-PrimitiveLineIndicesEXT-PrimitiveLineIndicesEXT-07050"></a>
  VUID-PrimitiveLineIndicesEXT-PrimitiveLineIndicesEXT-07050  
  The variable decorated with `PrimitiveLineIndicesEXT` **must** be
  declared as an array of two component vector 32-bit integer values

- <a href="#VUID-PrimitiveLineIndicesEXT-PrimitiveLineIndicesEXT-07051"
  id="VUID-PrimitiveLineIndicesEXT-PrimitiveLineIndicesEXT-07051"></a>
  VUID-PrimitiveLineIndicesEXT-PrimitiveLineIndicesEXT-07051  
  All index values of the array decorated with `PrimitiveLineIndicesEXT`
  **must** be in the range \[0, N-1\], where N is the value specified by
  the `OutputVertices` `Execution` `Mode`

- <a href="#VUID-PrimitiveLineIndicesEXT-PrimitiveLineIndicesEXT-07052"
  id="VUID-PrimitiveLineIndicesEXT-PrimitiveLineIndicesEXT-07052"></a>
  VUID-PrimitiveLineIndicesEXT-PrimitiveLineIndicesEXT-07052  
  The size of the array decorated with `PrimitiveLineIndicesEXT`
  **must** match the value specified by `OutputPrimitivesEXT`

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#PrimitiveLineIndicesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
