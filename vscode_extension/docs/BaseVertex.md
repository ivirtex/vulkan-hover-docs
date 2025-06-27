# BaseVertex(3) Manual Page

## Name

BaseVertex - First vertex being rendered



## <a href="#_description" class="anchor"></a>Description

`BaseVertex`  
Decorating a variable with the `BaseVertex` built-in will make that
variable contain the integer value corresponding to the first vertex or
vertex offset that was passed to the command that invoked the current
vertex shader invocation. For *non-indexed drawing commands*, this
variable is the `firstVertex` parameter to a *direct drawing command* or
the `firstVertex` member of the structure consumed by an *indirect
drawing command*. For *indexed drawing commands*, this variable is the
`vertexOffset` parameter to a *direct drawing command* or the
`vertexOffset` member of the structure consumed by an *indirect drawing
command*.

Valid Usage

- <a href="#VUID-BaseVertex-BaseVertex-04184"
  id="VUID-BaseVertex-BaseVertex-04184"></a>
  VUID-BaseVertex-BaseVertex-04184  
  The `BaseVertex` decoration **must** be used only within the `Vertex`
  `Execution` `Model`

- <a href="#VUID-BaseVertex-BaseVertex-04185"
  id="VUID-BaseVertex-BaseVertex-04185"></a>
  VUID-BaseVertex-BaseVertex-04185  
  The variable decorated with `BaseVertex` **must** be declared using
  the `Input` `Storage` `Class`

- <a href="#VUID-BaseVertex-BaseVertex-04186"
  id="VUID-BaseVertex-BaseVertex-04186"></a>
  VUID-BaseVertex-BaseVertex-04186  
  The variable decorated with `BaseVertex` **must** be declared as a
  scalar 32-bit integer value

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#BaseVertex"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
