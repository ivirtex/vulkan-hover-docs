# PrimitiveId(3) Manual Page

## Name

PrimitiveId - Primitive ID



## <a href="#_description" class="anchor"></a>Description

`PrimitiveId`  
Decorating a variable with the `PrimitiveId` built-in decoration will
make that variable contain the index of the current primitive.

The index of the first primitive generated by a drawing command is zero,
and the index is incremented after every individual point, line, or
triangle primitive is processed.

For triangles drawn as points or line segments (see <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-polygonmode"
target="_blank" rel="noopener">Polygon Mode</a>), the primitive index is
incremented only once, even if multiple points or lines are eventually
drawn.

Variables decorated with `PrimitiveId` are reset to zero between each
instance drawn.

Restarting a primitive topology using primitive restart has no effect on
the value of variables decorated with `PrimitiveId`.

In tessellation control and tessellation evaluation shaders, it will
contain the index of the patch within the current set of rendering
primitives that corresponds to the shader invocation.

In a geometry shader, it will contain the number of primitives presented
as input to the shader since the current set of rendering primitives was
started.

In a fragment shader, it will contain the primitive index written by the
mesh shader if a mesh shader is present, or the primitive index written
by the geometry shader if a geometry shader is present, or with the
value that would have been presented as input to the geometry shader had
it been present.

In an intersection, any-hit, or closest hit shader, it will contain the
index within the geometry of the triangle or bounding box being
processed.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>When the <code>PrimitiveId</code> decoration is applied to an output
variable in the mesh shader or geometry shader, the resulting value is
seen through the <code>PrimitiveId</code> decorated input variable in
the fragment shader.</p>
<p>The fragment shader using <code>PrimitiveId</code> will need to
declare either the <code>MeshShadingNV</code>,
<code>MeshShadingEXT</code>, <code>Geometry</code> or
<code>Tessellation</code> capability to satisfy the requirement SPIR-V
has to use <code>PrimitiveId</code>.</p></td>
</tr>
</tbody>
</table>

Valid Usage

- <a href="#VUID-PrimitiveId-PrimitiveId-04330"
  id="VUID-PrimitiveId-PrimitiveId-04330"></a>
  VUID-PrimitiveId-PrimitiveId-04330  
  The `PrimitiveId` decoration **must** be used only within the
  `MeshEXT`, `MeshNV`, `IntersectionKHR`, `AnyHitKHR`, `ClosestHitKHR`,
  `TessellationControl`, `TessellationEvaluation`, `Geometry`, or
  `Fragment` `Execution` `Model`

- <a href="#VUID-PrimitiveId-Fragment-04331"
  id="VUID-PrimitiveId-Fragment-04331"></a>
  VUID-PrimitiveId-Fragment-04331  
  If pipeline contains both the `Fragment` and `Geometry` `Execution`
  `Model` and a variable decorated with `PrimitiveId` is read from
  `Fragment` shader, then the `Geometry` shader **must** write to the
  output variables decorated with `PrimitiveId` in all execution paths

- <a href="#VUID-PrimitiveId-Fragment-04332"
  id="VUID-PrimitiveId-Fragment-04332"></a>
  VUID-PrimitiveId-Fragment-04332  
  If pipeline contains both the `Fragment` and `MeshEXT` or `MeshNV`
  `Execution` `Model` and a variable decorated with `PrimitiveId` is
  read from `Fragment` shader, then the `MeshEXT` or `MeshNV` shader
  **must** write to the output variables decorated with `PrimitiveId` in
  all execution paths

- <a href="#VUID-PrimitiveId-Fragment-04333"
  id="VUID-PrimitiveId-Fragment-04333"></a>
  VUID-PrimitiveId-Fragment-04333  
  If `Fragment` `Execution` `Model` contains a variable decorated with
  `PrimitiveId`, then either the `MeshShadingEXT`, `MeshShadingNV`,
  `Geometry` or `Tessellation` capability **must** also be declared

- <a href="#VUID-PrimitiveId-PrimitiveId-04334"
  id="VUID-PrimitiveId-PrimitiveId-04334"></a>
  VUID-PrimitiveId-PrimitiveId-04334  
  The variable decorated with `PrimitiveId` within the
  `TessellationControl`, `TessellationEvaluation`, `Fragment`,
  `IntersectionKHR`, `AnyHitKHR`, or `ClosestHitKHR` `Execution` `Model`
  **must** be declared using the `Input` `Storage` `Class`

- <a href="#VUID-PrimitiveId-PrimitiveId-04335"
  id="VUID-PrimitiveId-PrimitiveId-04335"></a>
  VUID-PrimitiveId-PrimitiveId-04335  
  The variable decorated with `PrimitiveId` within the `Geometry`
  `Execution` `Model` **must** be declared using the `Input` or `Output`
  `Storage` `Class`

- <a href="#VUID-PrimitiveId-PrimitiveId-04336"
  id="VUID-PrimitiveId-PrimitiveId-04336"></a>
  VUID-PrimitiveId-PrimitiveId-04336  
  The variable decorated with `PrimitiveId` within the `MeshEXT` or
  `MeshNV` `Execution` `Model` **must** be declared using the `Output`
  `Storage` `Class`

- <a href="#VUID-PrimitiveId-PrimitiveId-04337"
  id="VUID-PrimitiveId-PrimitiveId-04337"></a>
  VUID-PrimitiveId-PrimitiveId-04337  
  The variable decorated with `PrimitiveId` **must** be declared as a
  scalar 32-bit integer value

- <a href="#VUID-PrimitiveId-PrimitiveId-07040"
  id="VUID-PrimitiveId-PrimitiveId-07040"></a>
  VUID-PrimitiveId-PrimitiveId-07040  
  The variable decorated with `PrimitiveId` within the `MeshEXT`
  `Execution` `Model` **must** also be decorated with the
  `PerPrimitiveEXT` decoration

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#PrimitiveId"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700