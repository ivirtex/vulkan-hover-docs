# InvocationId(3) Manual Page

## Name

InvocationId - Invocation ID in a geometry or tessellation control shader



## [](#_description)Description

`InvocationId`

Decorating a variable with the `InvocationId` built-in decoration will make that variable contain the index of the current shader invocation in a geometry shader, or the index of the output patch vertex in a tessellation control shader.

In a geometry shader, the index of the current shader invocation ranges from zero to the number of [instances](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#geometry-invocations) declared in the shader minus one. If the instance count of the geometry shader is one or is not specified, then `InvocationId` will be zero.

Valid Usage

- [](#VUID-InvocationId-InvocationId-04257)VUID-InvocationId-InvocationId-04257  
  The `InvocationId` decoration **must** be used only within the `TessellationControl` or `Geometry` `Execution` `Model`
- [](#VUID-InvocationId-InvocationId-04258)VUID-InvocationId-InvocationId-04258  
  The variable decorated with `InvocationId` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-InvocationId-InvocationId-04259)VUID-InvocationId-InvocationId-04259  
  The variable decorated with `InvocationId` **must** be declared as a scalar 32-bit integer value

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#InvocationId)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0