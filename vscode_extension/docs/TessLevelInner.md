# TessLevelInner(3) Manual Page

## Name

TessLevelInner - Inner tessellation levels



## [](#_description)Description

`TessLevelInner`

Decorating a variable with the `TessLevelInner` built-in decoration will make that variable contain the inner tessellation levels for the current patch.

In tessellation control shaders, the variable decorated with `TessLevelInner` **can** be written to, controlling the tessellation factors for the resulting patch. These values are used by the tessellator to control primitive tessellation and **can** be read by tessellation evaluation shaders.

In tessellation evaluation shaders, the variable decorated with `TessLevelInner` **can** read the values written by the tessellation control shader.

Valid Usage

- [](#VUID-TessLevelInner-TessLevelInner-04394)VUID-TessLevelInner-TessLevelInner-04394  
  The `TessLevelInner` decoration **must** be used only within the `TessellationControl` or `TessellationEvaluation` `Execution` `Model`
- [](#VUID-TessLevelInner-TessLevelInner-04395)VUID-TessLevelInner-TessLevelInner-04395  
  The variable decorated with `TessLevelInner` within the `TessellationControl` `Execution` `Model` **must** be declared using the `Output` `Storage` `Class`
- [](#VUID-TessLevelInner-TessLevelInner-04396)VUID-TessLevelInner-TessLevelInner-04396  
  The variable decorated with `TessLevelInner` within the `TessellationEvaluation` `Execution` `Model` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-TessLevelInner-TessLevelInner-04397)VUID-TessLevelInner-TessLevelInner-04397  
  The variable decorated with `TessLevelInner` **must** be declared as an array of size two, containing 32-bit floating-point values

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#TessLevelInner)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0