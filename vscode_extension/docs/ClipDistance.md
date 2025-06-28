# ClipDistance(3) Manual Page

## Name

ClipDistance - Application-specified clip distances



## [](#_description)Description

`ClipDistance`

Decorating a variable with the `ClipDistance` built-in decoration will make that variable contain the mechanism for controlling user clipping. `ClipDistance` is an array such that the ith element of the array specifies the clip distance for plane i. A clip distance of 0 means the vertex is on the plane, a positive distance means the vertex is inside the clip half-space, and a negative distance means the vertex is outside the clip half-space.

Note

The array variable decorated with `ClipDistance` is explicitly sized by the shader.

Note

In the last [pre-rasterization shader stage](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization), these values will be linearly interpolated across the primitive and the portion of the primitive with interpolated distances less than 0 will be considered outside the clip volume. If `ClipDistance` is then used by a fragment shader, `ClipDistance` contains these linearly interpolated values.

Valid Usage

- [](#VUID-ClipDistance-ClipDistance-04187)VUID-ClipDistance-ClipDistance-04187  
  The `ClipDistance` decoration **must** be used only within the `MeshEXT`, `MeshNV`, `Vertex`, `Fragment`, `TessellationControl`, `TessellationEvaluation`, or `Geometry` `Execution` `Model`
- [](#VUID-ClipDistance-ClipDistance-04188)VUID-ClipDistance-ClipDistance-04188  
  The variable decorated with `ClipDistance` within the `MeshEXT`, `MeshNV`, or `Vertex` `Execution` `Model` **must** be declared using the `Output` `Storage` `Class`
- [](#VUID-ClipDistance-ClipDistance-04189)VUID-ClipDistance-ClipDistance-04189  
  The variable decorated with `ClipDistance` within the `Fragment` `Execution` `Model` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-ClipDistance-ClipDistance-04190)VUID-ClipDistance-ClipDistance-04190  
  The variable decorated with `ClipDistance` within the `TessellationControl`, `TessellationEvaluation`, or `Geometry` `Execution` `Model` **must** not be declared in a `Storage` `Class` other than `Input` or `Output`
- [](#VUID-ClipDistance-ClipDistance-04191)VUID-ClipDistance-ClipDistance-04191  
  The variable decorated with `ClipDistance` **must** be declared as an array of 32-bit floating-point values

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#ClipDistance)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0