# ClipDistance(3) Manual Page

## Name

ClipDistance - Application-specified clip distances



## <a href="#_description" class="anchor"></a>Description

`ClipDistance`  
Decorating a variable with the `ClipDistance` built-in decoration will
make that variable contain the mechanism for controlling user clipping.
`ClipDistance` is an array such that the i<sup>th</sup> element of the
array specifies the clip distance for plane i. A clip distance of 0
means the vertex is on the plane, a positive distance means the vertex
is inside the clip half-space, and a negative distance means the vertex
is outside the clip half-space.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>The array variable decorated with <code>ClipDistance</code> is
explicitly sized by the shader.</p></td>
</tr>
</tbody>
</table>

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>In the last <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
target="_blank" rel="noopener">pre-rasterization shader stage</a>, these
values will be linearly interpolated across the primitive and the
portion of the primitive with interpolated distances less than 0 will be
considered outside the clip volume. If <code>ClipDistance</code> is then
used by a fragment shader, <code>ClipDistance</code> contains these
linearly interpolated values.</p></td>
</tr>
</tbody>
</table>

Valid Usage

- <a href="#VUID-ClipDistance-ClipDistance-04187"
  id="VUID-ClipDistance-ClipDistance-04187"></a>
  VUID-ClipDistance-ClipDistance-04187  
  The `ClipDistance` decoration **must** be used only within the
  `MeshEXT`, `MeshNV`, `Vertex`, `Fragment`, `TessellationControl`,
  `TessellationEvaluation`, or `Geometry` `Execution` `Model`

- <a href="#VUID-ClipDistance-ClipDistance-04188"
  id="VUID-ClipDistance-ClipDistance-04188"></a>
  VUID-ClipDistance-ClipDistance-04188  
  The variable decorated with `ClipDistance` within the `MeshEXT`,
  `MeshNV`, or `Vertex` `Execution` `Model` **must** be declared using
  the `Output` `Storage` `Class`

- <a href="#VUID-ClipDistance-ClipDistance-04189"
  id="VUID-ClipDistance-ClipDistance-04189"></a>
  VUID-ClipDistance-ClipDistance-04189  
  The variable decorated with `ClipDistance` within the `Fragment`
  `Execution` `Model` **must** be declared using the `Input` `Storage`
  `Class`

- <a href="#VUID-ClipDistance-ClipDistance-04190"
  id="VUID-ClipDistance-ClipDistance-04190"></a>
  VUID-ClipDistance-ClipDistance-04190  
  The variable decorated with `ClipDistance` within the
  `TessellationControl`, `TessellationEvaluation`, or `Geometry`
  `Execution` `Model` **must** not be declared in a `Storage` `Class`
  other than `Input` or `Output`

- <a href="#VUID-ClipDistance-ClipDistance-04191"
  id="VUID-ClipDistance-ClipDistance-04191"></a>
  VUID-ClipDistance-ClipDistance-04191  
  The variable decorated with `ClipDistance` **must** be declared as an
  array of 32-bit floating-point values

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#ClipDistance"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
