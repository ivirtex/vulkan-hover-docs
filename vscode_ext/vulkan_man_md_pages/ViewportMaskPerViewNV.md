# ViewportMaskPerViewNV(3) Manual Page

## Name

ViewportMaskPerViewNV - Mask of viewports broadcast to per view



## <a href="#_description" class="anchor"></a>Description

`ViewportMaskPerViewNV`  
Decorating a variable with the `ViewportMaskPerViewNV` built-in
decoration will make that variable contain the mask of viewports
primitives are broadcast to, for each view.

The value written to an element of `ViewportMaskPerViewNV` in the last
<a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
target="_blank" rel="noopener">pre-rasterization shader stage</a> is a
bitmask indicating which viewports the primitive will be directed to.
The primitive will be broadcast to the viewport corresponding to each
non-zero bit of the bitmask, and that viewport index is used to select
the viewport transform, scissor rectangle, and exclusive scissor
rectangle, for each view. The same values **must** be written to all
vertices in a given primitive, or else the set of viewports used for
that primitive is undefined.

Elements of the array correspond to views in a multiview subpass, and
those elements corresponding to views in the view mask of the subpass
the shader is compiled against will be used as the viewport mask value
for those views. `ViewportMaskPerViewNV` output in an earlier <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
target="_blank" rel="noopener">pre-rasterization shader stage</a> is not
available as an input in the subsequent <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
target="_blank" rel="noopener">pre-rasterization shader stage</a>.

Although `ViewportMaskNV` is an array, `ViewportMaskPerViewNV` is not a
two-dimensional array. Instead, `ViewportMaskPerViewNV` is limited to 32
viewports.

Valid Usage

- <a href="#VUID-ViewportMaskPerViewNV-ViewportMaskPerViewNV-04412"
  id="VUID-ViewportMaskPerViewNV-ViewportMaskPerViewNV-04412"></a>
  VUID-ViewportMaskPerViewNV-ViewportMaskPerViewNV-04412  
  The `ViewportMaskPerViewNV` decoration **must** be used only within
  the `Vertex`, `MeshNV`, `TessellationControl`,
  `TessellationEvaluation`, or `Geometry` `Execution` `Model`

- <a href="#VUID-ViewportMaskPerViewNV-ViewportMaskPerViewNV-04413"
  id="VUID-ViewportMaskPerViewNV-ViewportMaskPerViewNV-04413"></a>
  VUID-ViewportMaskPerViewNV-ViewportMaskPerViewNV-04413  
  The variable decorated with `ViewportMaskPerViewNV` **must** be
  declared using the `Output` `Storage` `Class`

- <a href="#VUID-ViewportMaskPerViewNV-ViewportMaskPerViewNV-04414"
  id="VUID-ViewportMaskPerViewNV-ViewportMaskPerViewNV-04414"></a>
  VUID-ViewportMaskPerViewNV-ViewportMaskPerViewNV-04414  
  The variable decorated with `ViewportMaskPerViewNV` **must** be
  declared as an array of 32-bit integer values

- <a href="#VUID-ViewportMaskPerViewNV-ViewportMaskPerViewNV-04415"
  id="VUID-ViewportMaskPerViewNV-ViewportMaskPerViewNV-04415"></a>
  VUID-ViewportMaskPerViewNV-ViewportMaskPerViewNV-04415  
  The array decorated with `ViewportMaskPerViewNV` **must** be a size
  less than or equal to 32

- <a href="#VUID-ViewportMaskPerViewNV-ViewportMaskPerViewNV-04416"
  id="VUID-ViewportMaskPerViewNV-ViewportMaskPerViewNV-04416"></a>
  VUID-ViewportMaskPerViewNV-ViewportMaskPerViewNV-04416  
  The array decorated with `ViewportMaskPerViewNV` **must** be a size
  greater than the maximum view in the subpass’s view mask

- <a href="#VUID-ViewportMaskPerViewNV-ViewportMaskPerViewNV-04417"
  id="VUID-ViewportMaskPerViewNV-ViewportMaskPerViewNV-04417"></a>
  VUID-ViewportMaskPerViewNV-ViewportMaskPerViewNV-04417  
  The array variable decorated with `ViewportMaskPerViewNV` **must**
  only be indexed by a constant or specialization constant

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#ViewportMaskPerViewNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
