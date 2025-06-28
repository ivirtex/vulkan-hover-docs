# PrimitiveShadingRateKHR(3) Manual Page

## Name

PrimitiveShadingRateKHR - Primitive contribution to fragment shading rate



## [](#_description)Description

`PrimitiveShadingRateKHR`

Decorating a variable with the `PrimitiveShadingRateKHR` built-in decoration will make that variable contain the [primitive fragment shading rate](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#primsrast-fragment-shading-rate-primitive).

The value written to the variable decorated with `PrimitiveShadingRateKHR` by the last [pre-rasterization shader stage](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization) in the pipeline is used as the [primitive fragment shading rate](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#primsrast-fragment-shading-rate-primitive). Outputs in previous shader stages are ignored.

If the last active [pre-rasterization shader stage](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization) shader entry point’s interface does not include a variable decorated with `PrimitiveShadingRateKHR`, then it is as if the shader specified a fragment shading rate value of 0, indicating a horizontal and vertical rate of 1 pixel.

If a shader has `PrimitiveShadingRateKHR` in the output interface and there is an execution path through the shader that does not write to it, its value is undefined for executions of the shader that take that path.

Valid Usage

- [](#VUID-PrimitiveShadingRateKHR-PrimitiveShadingRateKHR-04484)VUID-PrimitiveShadingRateKHR-PrimitiveShadingRateKHR-04484  
  The `PrimitiveShadingRateKHR` decoration **must** be used only within the `MeshEXT`, `MeshNV`, `Vertex`, or `Geometry` `Execution` `Model`
- [](#VUID-PrimitiveShadingRateKHR-PrimitiveShadingRateKHR-04485)VUID-PrimitiveShadingRateKHR-PrimitiveShadingRateKHR-04485  
  The variable decorated with `PrimitiveShadingRateKHR` **must** be declared using the `Output` `Storage` `Class`
- [](#VUID-PrimitiveShadingRateKHR-PrimitiveShadingRateKHR-04486)VUID-PrimitiveShadingRateKHR-PrimitiveShadingRateKHR-04486  
  The variable decorated with `PrimitiveShadingRateKHR` **must** be declared as a scalar 32-bit integer value for all supported execution models except `MeshEXT`
- [](#VUID-PrimitiveShadingRateKHR-PrimitiveShadingRateKHR-04487)VUID-PrimitiveShadingRateKHR-PrimitiveShadingRateKHR-04487  
  The value written to `PrimitiveShadingRateKHR` **must** include no more than one of `Vertical2Pixels` and `Vertical4Pixels`
- [](#VUID-PrimitiveShadingRateKHR-PrimitiveShadingRateKHR-04488)VUID-PrimitiveShadingRateKHR-PrimitiveShadingRateKHR-04488  
  The value written to `PrimitiveShadingRateKHR` **must** include no more than one of `Horizontal2Pixels` and `Horizontal4Pixels`
- [](#VUID-PrimitiveShadingRateKHR-PrimitiveShadingRateKHR-04489)VUID-PrimitiveShadingRateKHR-PrimitiveShadingRateKHR-04489  
  The value written to `PrimitiveShadingRateKHR` **must** not have any bits set other than those defined by **Fragment Shading Rate Flags** enumerants in the SPIR-V specification
- [](#VUID-PrimitiveShadingRateKHR-PrimitiveShadingRateKHR-07059)VUID-PrimitiveShadingRateKHR-PrimitiveShadingRateKHR-07059  
  The variable decorated with `PrimitiveShadingRateKHR` within the `MeshEXT` `Execution` `Model` **must** also be decorated with the `PerPrimitiveEXT` decoration
- [](#VUID-PrimitiveShadingRateKHR-PrimitiveShadingRateKHR-10598)VUID-PrimitiveShadingRateKHR-PrimitiveShadingRateKHR-10598  
  `PrimitiveShadingRateKHR` within the `MeshEXT` `Execution` `Model` **must** decorate a scalar 32-bit integer member of a structure decorated as `Block`, or decorate a variable of type `OpTypeArray` of 32-bit integer values
- [](#VUID-PrimitiveShadingRateKHR-PrimitiveShadingRateKHR-10599)VUID-PrimitiveShadingRateKHR-PrimitiveShadingRateKHR-10599  
  If `PrimitiveShadingRateKHR` is declared as an array of boolean values, the size of the array **must** match the value specified by `OutputPrimitivesEXT`
- [](#VUID-PrimitiveShadingRateKHR-PrimitiveShadingRateKHR-10600)VUID-PrimitiveShadingRateKHR-PrimitiveShadingRateKHR-10600  
  If `PrimitiveShadingRateKHR` decorates a member of a structure, the variable declaration of the containing `Block` type **must** have an array size that matches the value specified by `OutputPrimitivesEXT`

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#PrimitiveShadingRateKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0