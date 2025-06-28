# BaseVertex(3) Manual Page

## Name

BaseVertex - First vertex being rendered



## [](#_description)Description

`BaseVertex`

Decorating a variable with the `BaseVertex` built-in will make that variable contain the integer value corresponding to the first vertex or vertex offset that was passed to the command that invoked the current vertex shader invocation. For *non-indexed drawing commands*, this variable is the `firstVertex` parameter to a *direct drawing command* or the `firstVertex` member of the structure consumed by an *indirect drawing command*. For *indexed drawing commands*, this variable is the `vertexOffset` parameter to a *direct drawing command* or the `vertexOffset` member of the structure consumed by an *indirect drawing command*.

Valid Usage

- [](#VUID-BaseVertex-BaseVertex-04184)VUID-BaseVertex-BaseVertex-04184  
  The `BaseVertex` decoration **must** be used only within the `Vertex` `Execution` `Model`
- [](#VUID-BaseVertex-BaseVertex-04185)VUID-BaseVertex-BaseVertex-04185  
  The variable decorated with `BaseVertex` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-BaseVertex-BaseVertex-04186)VUID-BaseVertex-BaseVertex-04186  
  The variable decorated with `BaseVertex` **must** be declared as a scalar 32-bit integer value

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#BaseVertex)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0