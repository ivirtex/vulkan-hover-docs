# DrawIndex(3) Manual Page

## Name

DrawIndex - Index of the current draw



## [](#_description)Description

`DrawIndex`

Decorating a variable with the `DrawIndex` built-in will make that variable contain the integer value corresponding to the zero-based index of the draw that invoked the current task, mesh, or vertex shader invocation. For *indirect drawing commands*, `DrawIndex` begins at zero and increments by one for each draw executed. The number of draws is given by the `drawCount` parameter. For *direct drawing commands*, if [vkCmdDrawMultiEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDrawMultiEXT.html) or [vkCmdDrawMultiIndexedEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDrawMultiIndexedEXT.html) is used, this variable contains the integer value corresponding to the zero-based index of the draw. Otherwise `DrawIndex` is always zero. `DrawIndex` is dynamically uniform.

When task or mesh shaders are used, only the first active stage will have proper access to the variable. The value read by other stages is undefined.

Valid Usage

- [](#VUID-DrawIndex-DrawIndex-04207)VUID-DrawIndex-DrawIndex-04207  
  The `DrawIndex` decoration **must** be used only within the `Vertex`, `MeshEXT`, `TaskEXT`, `MeshNV`, or `TaskNV` `Execution` `Model`
- [](#VUID-DrawIndex-DrawIndex-04208)VUID-DrawIndex-DrawIndex-04208  
  The variable decorated with `DrawIndex` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-DrawIndex-DrawIndex-04209)VUID-DrawIndex-DrawIndex-04209  
  The variable decorated with `DrawIndex` **must** be declared as a scalar 32-bit integer value

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#DrawIndex)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0