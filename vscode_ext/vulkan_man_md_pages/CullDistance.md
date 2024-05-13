# CullDistance(3) Manual Page

## Name

CullDistance - Application-specified cull distances



## <a href="#_description" class="anchor"></a>Description

`CullDistance`  
Decorating a variable with the `CullDistance` built-in decoration will
make that variable contain the mechanism for controlling user culling.
If any member of this array is assigned a negative value for all
vertices belonging to a primitive, then the primitive is discarded
before rasterization.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>In fragment shaders, the values of the <code>CullDistance</code>
array are linearly interpolated across each primitive.</p></td>
</tr>
</tbody>
</table>

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>If <code>CullDistance</code> decorates an input variable, that
variable will contain the corresponding value from the
<code>CullDistance</code> decorated output variable from the previous
shader stage.</p></td>
</tr>
</tbody>
</table>

Valid Usage

- <a href="#VUID-CullDistance-CullDistance-04196"
  id="VUID-CullDistance-CullDistance-04196"></a>
  VUID-CullDistance-CullDistance-04196  
  The `CullDistance` decoration **must** be used only within the
  `MeshEXT`, `MeshNV`, `Vertex`, `Fragment`, `TessellationControl`,
  `TessellationEvaluation`, or `Geometry` `Execution` `Model`

- <a href="#VUID-CullDistance-CullDistance-04197"
  id="VUID-CullDistance-CullDistance-04197"></a>
  VUID-CullDistance-CullDistance-04197  
  The variable decorated with `CullDistance` within the `MeshEXT`,
  `MeshNV` or `Vertex` `Execution` `Model` **must** be declared using
  the `Output` `Storage` `Class`

- <a href="#VUID-CullDistance-CullDistance-04198"
  id="VUID-CullDistance-CullDistance-04198"></a>
  VUID-CullDistance-CullDistance-04198  
  The variable decorated with `CullDistance` within the `Fragment`
  `Execution` `Model` **must** be declared using the `Input` `Storage`
  `Class`

- <a href="#VUID-CullDistance-CullDistance-04199"
  id="VUID-CullDistance-CullDistance-04199"></a>
  VUID-CullDistance-CullDistance-04199  
  The variable decorated with `CullDistance` within the
  `TessellationControl`, `TessellationEvaluation`, or `Geometry`
  `Execution` `Model` **must** not be declared using a `Storage` `Class`
  other than `Input` or `Output`

- <a href="#VUID-CullDistance-CullDistance-04200"
  id="VUID-CullDistance-CullDistance-04200"></a>
  VUID-CullDistance-CullDistance-04200  
  The variable decorated with `CullDistance` **must** be declared as an
  array of 32-bit floating-point values

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#CullDistance"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
