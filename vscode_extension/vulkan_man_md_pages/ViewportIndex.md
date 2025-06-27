# ViewportIndex(3) Manual Page

## Name

ViewportIndex - Viewport index used



## <a href="#_description" class="anchor"></a>Description

`ViewportIndex`  
Decorating a variable with the `ViewportIndex` built-in decoration will
make that variable contain the index of the viewport.

In a mesh, vertex, tessellation evaluation, or geometry shader, the
variable decorated with `ViewportIndex` can be written to with the
viewport index to which the primitive produced by that shader will be
directed.

The selected viewport index is used to select the viewport transform,
scissor rectangle, and exclusive scissor rectangle.

The last active *<a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
target="_blank" rel="noopener">pre-rasterization shader stage</a>* (in
pipeline order) controls the `ViewportIndex` that is used. Outputs in
previous shader stages are not used, even if the last stage fails to
write the `ViewportIndex`.

If the last active <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
target="_blank" rel="noopener">pre-rasterization shader stage</a> shader
entry point’s interface does not include a variable decorated with
`ViewportIndex` , and if <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-multiview-per-view-viewports"
target="_blank" rel="noopener">multiviewPerViewViewports</a> is not
enabled, then the first viewport is used. If a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
target="_blank" rel="noopener">pre-rasterization shader stage</a> shader
entry point’s interface includes a variable decorated with
`ViewportIndex`, it **must** write the same value to `ViewportIndex` for
all output vertices of a given primitive.

In a fragment shader, the variable decorated with `ViewportIndex`
contains the viewport index of the primitive that the fragment
invocation belongs to.

If <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-multiview-per-view-viewports"
target="_blank"
rel="noopener"><code>multiviewPerViewViewports</code></a> is enabled,
and if the last active <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
target="_blank" rel="noopener">pre-rasterization shader stage</a> shader
entry point’s interface does not include a variable decorated with
`ViewportIndex`, then the value of `ViewIndex` is used as an index to
select the viewport transform and scissor rectangle, and the value of
`ViewportIndex` in the fragment shader is undefined:.

Valid Usage

- <a href="#VUID-ViewportIndex-ViewportIndex-04404"
  id="VUID-ViewportIndex-ViewportIndex-04404"></a>
  VUID-ViewportIndex-ViewportIndex-04404  
  The `ViewportIndex` decoration **must** be used only within the
  `MeshEXT`, `MeshNV`, `Vertex`, `TessellationEvaluation`, `Geometry`,
  or `Fragment` `Execution` `Model`

- <a href="#VUID-ViewportIndex-ViewportIndex-04405"
  id="VUID-ViewportIndex-ViewportIndex-04405"></a>
  VUID-ViewportIndex-ViewportIndex-04405  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderOutputViewportIndex"
  target="_blank"
  rel="noopener"><code>shaderOutputViewportIndex</code></a> feature is
  not enabled then the `ViewportIndex` decoration **must** be used only
  within the `Geometry` or `Fragment` `Execution` `Model`

- <a href="#VUID-ViewportIndex-ViewportIndex-04406"
  id="VUID-ViewportIndex-ViewportIndex-04406"></a>
  VUID-ViewportIndex-ViewportIndex-04406  
  The variable decorated with `ViewportIndex` within the `MeshEXT`,
  `MeshNV`, `Vertex`, `TessellationEvaluation`, or `Geometry`
  `Execution` `Model` **must** be declared using the `Output` `Storage`
  `Class`

- <a href="#VUID-ViewportIndex-ViewportIndex-04407"
  id="VUID-ViewportIndex-ViewportIndex-04407"></a>
  VUID-ViewportIndex-ViewportIndex-04407  
  The variable decorated with `ViewportIndex` within the `Fragment`
  `Execution` `Model` **must** be declared using the `Input` `Storage`
  `Class`

- <a href="#VUID-ViewportIndex-ViewportIndex-04408"
  id="VUID-ViewportIndex-ViewportIndex-04408"></a>
  VUID-ViewportIndex-ViewportIndex-04408  
  The variable decorated with `ViewportIndex` **must** be declared as a
  scalar 32-bit integer value

- <a href="#VUID-ViewportIndex-ViewportIndex-07060"
  id="VUID-ViewportIndex-ViewportIndex-07060"></a>
  VUID-ViewportIndex-ViewportIndex-07060  
  The variable decorated with `ViewportIndex` within the `MeshEXT`
  `Execution` `Model` **must** also be decorated with the
  `PerPrimitiveEXT` decoration

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#ViewportIndex"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
