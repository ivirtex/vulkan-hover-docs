# VertexIndex(3) Manual Page

## Name

VertexIndex - Vertex index of a shader invocation



## [](#_description)Description

`VertexIndex`

Decorating a variable with the `VertexIndex` built-in decoration will make that variable contain the index of the vertex that is being processed by the current vertex shader invocation. For non-indexed draws, this variable begins at the `firstVertex` parameter to [vkCmdDraw](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDraw.html) or the `firstVertex` member of a structure consumed by [vkCmdDrawIndirect](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDrawIndirect.html) and increments by one for each vertex in the draw. For indexed draws, its value is the content of the index buffer for the vertex plus the `vertexOffset` parameter to [vkCmdDrawIndexed](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDrawIndexed.html) or the `vertexOffset` member of the structure consumed by [vkCmdDrawIndexedIndirect](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDrawIndexedIndirect.html).

Note

`VertexIndex` starts at the same starting value for each instance.

Valid Usage

- [](#VUID-VertexIndex-VertexIndex-04398)VUID-VertexIndex-VertexIndex-04398  
  The `VertexIndex` decoration **must** be used only within the `Vertex` `Execution` `Model`
- [](#VUID-VertexIndex-VertexIndex-04399)VUID-VertexIndex-VertexIndex-04399  
  The variable decorated with `VertexIndex` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-VertexIndex-VertexIndex-04400)VUID-VertexIndex-VertexIndex-04400  
  The variable decorated with `VertexIndex` **must** be declared as a scalar 32-bit integer value

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VertexIndex)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0