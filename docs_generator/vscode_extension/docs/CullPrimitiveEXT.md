# CullPrimitiveEXT(3) Manual Page

## Name

CullPrimitiveEXT - Application-specified culling state per primitive



## [](#_description)Description

`CullPrimitiveEXT`

Decorating a variable with the `CullPrimitiveEXT` built-in decoration will make that variable contain the culling state of output primitives. If the per-primitive boolean value is `true`, the primitive will be culled, if it is `false` it will not be culled.

Valid Usage

- [](#VUID-CullPrimitiveEXT-CullPrimitiveEXT-07034)VUID-CullPrimitiveEXT-CullPrimitiveEXT-07034  
  The `CullPrimitiveEXT` decoration **must** be used only within the `MeshEXT` `Execution` `Model`
- [](#VUID-CullPrimitiveEXT-CullPrimitiveEXT-07035)VUID-CullPrimitiveEXT-CullPrimitiveEXT-07035  
  The variable decorated with `CullPrimitiveEXT` **must** be declared using the `Output` `Storage` `Class`
- [](#VUID-CullPrimitiveEXT-CullPrimitiveEXT-07036)VUID-CullPrimitiveEXT-CullPrimitiveEXT-07036  
  `CullPrimitiveEXT` **must** decorate a scalar boolean member of a structure decorated as `Block`, or decorate a variable of type `OpTypeArray` of boolean values
- [](#VUID-CullPrimitiveEXT-CullPrimitiveEXT-10589)VUID-CullPrimitiveEXT-CullPrimitiveEXT-10589  
  If `CullPrimitiveEXT` is declared as an array of boolean values, the size of the array **must** match the value specified by `OutputPrimitivesEXT`
- [](#VUID-CullPrimitiveEXT-CullPrimitiveEXT-10590)VUID-CullPrimitiveEXT-CullPrimitiveEXT-10590  
  If `CullPrimitiveEXT` decorates a member of a structure, the variable declaration of the containing `Block` type **must** have an array size that matches the value specified by `OutputPrimitivesEXT`
- [](#VUID-CullPrimitiveEXT-CullPrimitiveEXT-10591)VUID-CullPrimitiveEXT-CullPrimitiveEXT-10591  
  There must be only one declaration of the `CullPrimitiveEXT` associated with a entry point’s interface
- [](#VUID-CullPrimitiveEXT-CullPrimitiveEXT-07038)VUID-CullPrimitiveEXT-CullPrimitiveEXT-07038  
  The variable decorated with `CullPrimitiveEXT` within the `MeshEXT` `Execution` `Model` **must** also be decorated with the `PerPrimitiveEXT` decoration

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#CullPrimitiveEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0