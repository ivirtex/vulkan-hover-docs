# Layer(3) Manual Page

## Name

Layer - Layer index for layered rendering



## <a href="#_description" class="anchor"></a>Description

`Layer`  
Decorating a variable with the `Layer` built-in decoration will make
that variable contain the select layer of a multi-layer framebuffer
attachment.

In a mesh, vertex, tessellation evaluation, or geometry shader, any
variable decorated with `Layer` can be written with the framebuffer
layer index to which the primitive produced by that shader will be
directed.

The last active <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
target="_blank" rel="noopener">pre-rasterization shader stage</a> (in
pipeline order) controls the `Layer` that is used. Outputs in previous
shader stages are not used, even if the last stage fails to write the
`Layer`.

If the last active <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
target="_blank" rel="noopener">pre-rasterization shader stage</a> shader
entry point’s interface does not include a variable decorated with
`Layer`, then the first layer is used. If a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
target="_blank" rel="noopener">pre-rasterization shader stage</a> shader
entry point’s interface includes a variable decorated with `Layer`, it
**must** write the same value to `Layer` for all output vertices of a
given primitive. If the `Layer` value is less than 0 or greater than or
equal to the number of layers in the framebuffer, then primitives
**may** still be rasterized, fragment shaders **may** be executed, and
the framebuffer values for all layers are undefined. In a mesh shader
this also applies when the `Layer` value is greater than or equal to the
`maxMeshOutputLayers` limit.

If a variable with the `Layer` decoration is also decorated with
`ViewportRelativeNV`, then the `ViewportIndex` is added to the layer
that is used for rendering and that is made available in the fragment
shader.

If the shader writes to a variable decorated `ViewportMaskNV`, then the
layer selected has a different value for each viewport a primitive is
rendered to.

In a fragment shader, a variable decorated with `Layer` contains the
layer index of the primitive that the fragment invocation belongs to.

Valid Usage

- <a href="#VUID-Layer-Layer-04272" id="VUID-Layer-Layer-04272"></a>
  VUID-Layer-Layer-04272  
  The `Layer` decoration **must** be used only within the `MeshEXT`,
  `MeshNV`, `Vertex`, `TessellationEvaluation`, `Geometry`, or
  `Fragment` `Execution` `Model`

- <a href="#VUID-Layer-Layer-04273" id="VUID-Layer-Layer-04273"></a>
  VUID-Layer-Layer-04273  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderOutputLayer"
  target="_blank" rel="noopener"><code>shaderOutputLayer</code></a>
  feature is not enabled then the `Layer` decoration **must** be used
  only within the `Geometry` or `Fragment` `Execution` `Model`

- <a href="#VUID-Layer-Layer-04274" id="VUID-Layer-Layer-04274"></a>
  VUID-Layer-Layer-04274  
  The variable decorated with `Layer` within the `MeshEXT`, `MeshNV`,
  `Vertex`, `TessellationEvaluation`, or `Geometry` `Execution` `Model`
  **must** be declared using the `Output` `Storage` `Class`

- <a href="#VUID-Layer-Layer-04275" id="VUID-Layer-Layer-04275"></a>
  VUID-Layer-Layer-04275  
  The variable decorated with `Layer` within the `Fragment` `Execution`
  `Model` **must** be declared using the `Input` `Storage` `Class`

- <a href="#VUID-Layer-Layer-04276" id="VUID-Layer-Layer-04276"></a>
  VUID-Layer-Layer-04276  
  The variable decorated with `Layer` **must** be declared as a scalar
  32-bit integer value

- <a href="#VUID-Layer-Layer-07039" id="VUID-Layer-Layer-07039"></a>
  VUID-Layer-Layer-07039  
  The variable decorated with `Layer` within the `MeshEXT` `Execution`
  `Model` **must** also be decorated with the `PerPrimitiveEXT`
  decoration

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#Layer"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
