# ViewportMaskNV(3) Manual Page

## Name

ViewportMaskNV - Mask of the viewports used



## <a href="#_description" class="anchor"></a>Description

`ViewportMaskNV`  
Decorating a variable with the `ViewportMaskNV` built-in decoration will
make that variable contain the viewport mask.

In a mesh, vertex, tessellation evaluation, or geometry shader, the
variable decorated with `ViewportMaskNV` can be written to with the mask
of which viewports the primitive produced by that shader will directed.

The `ViewportMaskNV` variable **must** be an array that has
⌈(`VkPhysicalDeviceLimits`::`maxViewports` / 32)⌉ elements. When a
shader writes to this variable, bit B of element M controls whether a
primitive is emitted to viewport 32 × M + B. The viewports indicated by
the mask are used to select the viewport transform, scissor rectangle,
and exclusive scissor rectangle that a primitive will be transformed by.

The last active *<a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
target="_blank" rel="noopener">pre-rasterization shader stage</a>* (in
pipeline order) controls the `ViewportMaskNV` that is used. Outputs in
previous shader stages are not used, even if the last stage fails to
write the `ViewportMaskNV`. When `ViewportMaskNV` is written by the
final <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
target="_blank" rel="noopener">pre-rasterization shader stage</a>, any
variable decorated with `ViewportIndex` in the fragment shader will have
the index of the viewport that was used in generating that fragment.

If a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
target="_blank" rel="noopener">pre-rasterization shader stage</a> shader
entry point’s interface includes a variable decorated with
`ViewportMaskNV`, it **must** write the same value to `ViewportMaskNV`
for all output vertices of a given primitive.

Valid Usage

- <a href="#VUID-ViewportMaskNV-ViewportMaskNV-04409"
  id="VUID-ViewportMaskNV-ViewportMaskNV-04409"></a>
  VUID-ViewportMaskNV-ViewportMaskNV-04409  
  The `ViewportMaskNV` decoration **must** be used only within the
  `Vertex`, `MeshNV`, `TessellationEvaluation`, or `Geometry`
  `Execution` `Model`

- <a href="#VUID-ViewportMaskNV-ViewportMaskNV-04410"
  id="VUID-ViewportMaskNV-ViewportMaskNV-04410"></a>
  VUID-ViewportMaskNV-ViewportMaskNV-04410  
  The variable decorated with `ViewportMaskNV` **must** be declared
  using the `Output` `Storage` `Class`

- <a href="#VUID-ViewportMaskNV-ViewportMaskNV-04411"
  id="VUID-ViewportMaskNV-ViewportMaskNV-04411"></a>
  VUID-ViewportMaskNV-ViewportMaskNV-04411  
  The variable decorated with `ViewportMaskNV` **must** be declared as
  an array of 32-bit integer values

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#ViewportMaskNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
