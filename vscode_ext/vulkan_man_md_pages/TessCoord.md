# TessCoord(3) Manual Page

## Name

TessCoord - Barycentric coordinate of a tessellated vertex within a
patch



## <a href="#_description" class="anchor"></a>Description

`TessCoord`  
Decorating a variable with the `TessCoord` built-in decoration will make
that variable contain the three-dimensional (u,v,w) barycentric
coordinate of the tessellated vertex within the patch. u, v, and w are
in the range \[0,1\] and vary linearly across the primitive being
subdivided. For the tessellation modes of `Quads` or `IsoLines`, the
third component is always zero.

Valid Usage

- <a href="#VUID-TessCoord-TessCoord-04387"
  id="VUID-TessCoord-TessCoord-04387"></a>
  VUID-TessCoord-TessCoord-04387  
  The `TessCoord` decoration **must** be used only within the
  `TessellationEvaluation` `Execution` `Model`

- <a href="#VUID-TessCoord-TessCoord-04388"
  id="VUID-TessCoord-TessCoord-04388"></a>
  VUID-TessCoord-TessCoord-04388  
  The variable decorated with `TessCoord` **must** be declared using the
  `Input` `Storage` `Class`

- <a href="#VUID-TessCoord-TessCoord-04389"
  id="VUID-TessCoord-TessCoord-04389"></a>
  VUID-TessCoord-TessCoord-04389  
  The variable decorated with `TessCoord` **must** be declared as a
  three-component vector of 32-bit floating-point values

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#TessCoord"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
