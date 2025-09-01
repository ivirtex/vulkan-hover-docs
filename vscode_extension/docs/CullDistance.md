# CullDistance(3) Manual Page

## Name

CullDistance - Application-specified cull distances



## [](#_description)Description

`CullDistance`

Decorating a variable with the `CullDistance` built-in decoration will make that variable contain the mechanism for controlling user culling. If any member of this array is assigned a negative value for all vertices belonging to a primitive, then the primitive is discarded before rasterization.

Note

In fragment shaders, the values of the `CullDistance` array are linearly interpolated across each primitive.

Note

If `CullDistance` decorates an input variable, that variable will contain the corresponding value from the `CullDistance` decorated output variable from the previous shader stage.

Valid Usage

- [](#VUID-CullDistance-CullDistance-04196)VUID-CullDistance-CullDistance-04196  
  The `CullDistance` decoration **must** be used only within the `MeshEXT`, `MeshNV`, `Vertex`, `Fragment`, `TessellationControl`, `TessellationEvaluation`, or `Geometry` `Execution` `Model`
- [](#VUID-CullDistance-CullDistance-04197)VUID-CullDistance-CullDistance-04197  
  The variable decorated with `CullDistance` within the `MeshEXT`, `MeshNV` or `Vertex` `Execution` `Model` **must** be declared using the `Output` `Storage` `Class`
- [](#VUID-CullDistance-CullDistance-04198)VUID-CullDistance-CullDistance-04198  
  The variable decorated with `CullDistance` within the `Fragment` `Execution` `Model` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-CullDistance-CullDistance-04199)VUID-CullDistance-CullDistance-04199  
  The variable decorated with `CullDistance` within the `TessellationControl`, `TessellationEvaluation`, or `Geometry` `Execution` `Model` **must** not be declared using a `Storage` `Class` other than `Input` or `Output`
- [](#VUID-CullDistance-CullDistance-04200)VUID-CullDistance-CullDistance-04200  
  The variable decorated with `CullDistance` **must** be declared as an array of 32-bit floating-point values

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#CullDistance).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0