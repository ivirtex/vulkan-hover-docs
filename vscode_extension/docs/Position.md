# Position(3) Manual Page

## Name

Position - Vertex position



## [](#_description)Description

`Position`

Decorating a variable with the `Position` built-in decoration will make that variable contain the position of the current vertex. In the last [pre-rasterization shader stage](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization), the value of the variable decorated with `Position` is used in subsequent primitive assembly, clipping, and rasterization operations.

Note

When `Position` decorates a variable in the `Input` `Storage` `Class`, it contains the data written to the output variable decorated with `Position` from the previous shader stage.

Valid Usage

- [](#VUID-Position-Position-04318)VUID-Position-Position-04318  
  The `Position` decoration **must** be used only within the `MeshEXT`, `MeshNV`, `Vertex`, `TessellationControl`, `TessellationEvaluation`, or `Geometry` `Execution` `Model`
- [](#VUID-Position-Position-04319)VUID-Position-Position-04319  
  The variable decorated with `Position` within the `MeshEXT`, `MeshNV`, or `Vertex` `Execution` `Model` **must** be declared using the `Output` `Storage` `Class`
- [](#VUID-Position-Position-04320)VUID-Position-Position-04320  
  The variable decorated with `Position` within the `TessellationControl`, `TessellationEvaluation`, or `Geometry` `Execution` `Model` **must** not be declared using a `Storage` `Class` other than `Input` or `Output`
- [](#VUID-Position-Position-04321)VUID-Position-Position-04321  
  The variable decorated with `Position` **must** be declared as a four-component vector of 32-bit floating-point values

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#Position)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0