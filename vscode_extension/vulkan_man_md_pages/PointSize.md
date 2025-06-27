# PointSize(3) Manual Page

## Name

PointSize - Size of a point primitive



## <a href="#_description" class="anchor"></a>Description

`PointSize`  
Decorating a variable with the `PointSize` built-in decoration will make
that variable contain the size of point primitives or the final
rasterization of polygons if <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-polygonmode"
target="_blank" rel="noopener">polygon mode</a> is
`VK_POLYGON_MODE_POINT` when
`VkPhysicalDeviceMaintenance5PropertiesKHR`::`polygonModePointSize` is
set to `VK_TRUE` . The value written to the variable decorated with
`PointSize` by the last <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
target="_blank" rel="noopener">pre-rasterization shader stage</a> in the
pipeline is used as the framebuffer-space size of points produced by
rasterization. If <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-maintenance5"
target="_blank" rel="noopener"><code>maintenance5</code></a> is enabled
and a value is not written to a variable decorated with `PointSize`, a
value of 1.0 is used as the size of points.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>When <code>PointSize</code> decorates a variable in the
<code>Input</code> <code>Storage</code> <code>Class</code>, it contains
the data written to the output variable decorated with
<code>PointSize</code> from the previous shader stage.</p></td>
</tr>
</tbody>
</table>

Valid Usage

- <a href="#VUID-PointSize-PointSize-04314"
  id="VUID-PointSize-PointSize-04314"></a>
  VUID-PointSize-PointSize-04314  
  The `PointSize` decoration **must** be used only within the `MeshEXT`,
  `MeshNV`, `Vertex`, `TessellationControl`, `TessellationEvaluation`,
  or `Geometry` `Execution` `Model`

- <a href="#VUID-PointSize-PointSize-04315"
  id="VUID-PointSize-PointSize-04315"></a>
  VUID-PointSize-PointSize-04315  
  The variable decorated with `PointSize` within the `MeshEXT`,
  `MeshNV`, or `Vertex` `Execution` `Model` **must** be declared using
  the `Output` `Storage` `Class`

- <a href="#VUID-PointSize-PointSize-04316"
  id="VUID-PointSize-PointSize-04316"></a>
  VUID-PointSize-PointSize-04316  
  The variable decorated with `PointSize` within the
  `TessellationControl`, `TessellationEvaluation`, or `Geometry`
  `Execution` `Model` **must** not be declared using a `Storage` `Class`
  other than `Input` or `Output`

- <a href="#VUID-PointSize-PointSize-04317"
  id="VUID-PointSize-PointSize-04317"></a>
  VUID-PointSize-PointSize-04317  
  The variable decorated with `PointSize` **must** be declared as a
  scalar 32-bit floating-point value

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#PointSize"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
