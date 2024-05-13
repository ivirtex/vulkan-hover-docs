# VertexIndex(3) Manual Page

## Name

VertexIndex - Vertex index of a shader invocation



## <a href="#_description" class="anchor"></a>Description

`VertexIndex`  
Decorating a variable with the `VertexIndex` built-in decoration will
make that variable contain the index of the vertex that is being
processed by the current vertex shader invocation. For non-indexed
draws, this variable begins at the `firstVertex` parameter to
[vkCmdDraw](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDraw.html) or the `firstVertex` member of a structure
consumed by [vkCmdDrawIndirect](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDrawIndirect.html) and increments
by one for each vertex in the draw. For indexed draws, its value is the
content of the index buffer for the vertex plus the `vertexOffset`
parameter to [vkCmdDrawIndexed](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDrawIndexed.html) or the
`vertexOffset` member of the structure consumed by
[vkCmdDrawIndexedIndirect](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDrawIndexedIndirect.html).

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p><code>VertexIndex</code> starts at the same starting value for each
instance.</p></td>
</tr>
</tbody>
</table>

Valid Usage

- <a href="#VUID-VertexIndex-VertexIndex-04398"
  id="VUID-VertexIndex-VertexIndex-04398"></a>
  VUID-VertexIndex-VertexIndex-04398  
  The `VertexIndex` decoration **must** be used only within the `Vertex`
  `Execution` `Model`

- <a href="#VUID-VertexIndex-VertexIndex-04399"
  id="VUID-VertexIndex-VertexIndex-04399"></a>
  VUID-VertexIndex-VertexIndex-04399  
  The variable decorated with `VertexIndex` **must** be declared using
  the `Input` `Storage` `Class`

- <a href="#VUID-VertexIndex-VertexIndex-04400"
  id="VUID-VertexIndex-VertexIndex-04400"></a>
  VUID-VertexIndex-VertexIndex-04400  
  The variable decorated with `VertexIndex` **must** be declared as a
  scalar 32-bit integer value

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VertexIndex"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
