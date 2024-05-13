# Position(3) Manual Page

## Name

Position - Vertex position



## <a href="#_description" class="anchor"></a>Description

`Position`  
Decorating a variable with the `Position` built-in decoration will make
that variable contain the position of the current vertex. In the last <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
target="_blank" rel="noopener">pre-rasterization shader stage</a>, the
value of the variable decorated with `Position` is used in subsequent
primitive assembly, clipping, and rasterization operations.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>When <code>Position</code> decorates a variable in the
<code>Input</code> <code>Storage</code> <code>Class</code>, it contains
the data written to the output variable decorated with
<code>Position</code> from the previous shader stage.</p></td>
</tr>
</tbody>
</table>

Valid Usage

- <a href="#VUID-Position-Position-04318"
  id="VUID-Position-Position-04318"></a> VUID-Position-Position-04318  
  The `Position` decoration **must** be used only within the `MeshEXT`,
  `MeshNV`, `Vertex`, `TessellationControl`, `TessellationEvaluation`,
  or `Geometry` `Execution` `Model`

- <a href="#VUID-Position-Position-04319"
  id="VUID-Position-Position-04319"></a> VUID-Position-Position-04319  
  The variable decorated with `Position` within the `MeshEXT`, `MeshNV`,
  or `Vertex` `Execution` `Model` **must** be declared using the
  `Output` `Storage` `Class`

- <a href="#VUID-Position-Position-04320"
  id="VUID-Position-Position-04320"></a> VUID-Position-Position-04320  
  The variable decorated with `Position` within the
  `TessellationControl`, `TessellationEvaluation`, or `Geometry`
  `Execution` `Model` **must** not be declared using a `Storage` `Class`
  other than `Input` or `Output`

- <a href="#VUID-Position-Position-04321"
  id="VUID-Position-Position-04321"></a> VUID-Position-Position-04321  
  The variable decorated with `Position` **must** be declared as a
  four-component vector of 32-bit floating-point values

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#Position"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
