# PositionPerViewNV(3) Manual Page

## Name

PositionPerViewNV - Vertex position per view



## <a href="#_description" class="anchor"></a>Description

`PositionPerViewNV`  
Decorating a variable with the `PositionPerViewNV` built-in decoration
will make that variable contain the position of the current vertex, for
each view.

Elements of the array correspond to views in a multiview subpass, and
those elements corresponding to views in the view mask of the subpass
the shader is compiled against will be used as the position value for
those views. For the final <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
target="_blank" rel="noopener">pre-rasterization shader stage</a> in the
pipeline, values written to an output variable decorated with
`PositionPerViewNV` are used in subsequent primitive assembly, clipping,
and rasterization operations, as with `Position`. `PositionPerViewNV`
output in an earlier <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
target="_blank" rel="noopener">pre-rasterization shader stage</a> is
available as an input in the subsequent <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
target="_blank" rel="noopener">pre-rasterization shader stage</a>.

If a shader is compiled against a subpass that has the
`VK_SUBPASS_DESCRIPTION_PER_VIEW_POSITION_X_ONLY_BIT_NVX` bit set, then
the position values for each view **must** not differ in any component
other than the X component. If the values do differ, one will be chosen
in an implementation-dependent manner.

Valid Usage

- <a href="#VUID-PositionPerViewNV-PositionPerViewNV-04322"
  id="VUID-PositionPerViewNV-PositionPerViewNV-04322"></a>
  VUID-PositionPerViewNV-PositionPerViewNV-04322  
  The `PositionPerViewNV` decoration **must** be used only within the
  `MeshNV`, `Vertex`, `TessellationControl`, `TessellationEvaluation`,
  or `Geometry` `Execution` `Model`

- <a href="#VUID-PositionPerViewNV-PositionPerViewNV-04323"
  id="VUID-PositionPerViewNV-PositionPerViewNV-04323"></a>
  VUID-PositionPerViewNV-PositionPerViewNV-04323  
  The variable decorated with `PositionPerViewNV` within the `Vertex`,
  or `MeshNV` `Execution` `Model` **must** be declared using the
  `Output` `Storage` `Class`

- <a href="#VUID-PositionPerViewNV-PositionPerViewNV-04324"
  id="VUID-PositionPerViewNV-PositionPerViewNV-04324"></a>
  VUID-PositionPerViewNV-PositionPerViewNV-04324  
  The variable decorated with `PositionPerViewNV` within the
  `TessellationControl`, `TessellationEvaluation`, or `Geometry`
  `Execution` `Model` **must** not be declared using a `Storage` `Class`
  other than `Input` or `Output`

- <a href="#VUID-PositionPerViewNV-PositionPerViewNV-04325"
  id="VUID-PositionPerViewNV-PositionPerViewNV-04325"></a>
  VUID-PositionPerViewNV-PositionPerViewNV-04325  
  The variable decorated with `PositionPerViewNV` **must** be declared
  as an array of four-component vector of 32-bit floating-point values
  with at least as many elements as the maximum view in the subpass’s
  view mask plus one

- <a href="#VUID-PositionPerViewNV-PositionPerViewNV-04326"
  id="VUID-PositionPerViewNV-PositionPerViewNV-04326"></a>
  VUID-PositionPerViewNV-PositionPerViewNV-04326  
  The array variable decorated with `PositionPerViewNV` **must** only be
  indexed by a constant or specialization constant

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#PositionPerViewNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
