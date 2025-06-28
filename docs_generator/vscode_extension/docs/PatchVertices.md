# PatchVertices(3) Manual Page

## Name

PatchVertices - Number of vertices in an input patch



## [](#_description)Description

`PatchVertices`

Decorating a variable with the `PatchVertices` built-in decoration will make that variable contain the number of vertices in the input patch being processed by the shader. In a Tessellation Control Shader, this is the same as the name:patchControlPoints member of [VkPipelineTessellationStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineTessellationStateCreateInfo.html). In a Tessellation Evaluation Shader, `PatchVertices` is equal to the tessellation control output patch size. When the same shader is used in different pipelines where the patch sizes are configured differently, the value of the `PatchVertices` variable will also differ.

Valid Usage

- [](#VUID-PatchVertices-PatchVertices-04308)VUID-PatchVertices-PatchVertices-04308  
  The `PatchVertices` decoration **must** be used only within the `TessellationControl` or `TessellationEvaluation` `Execution` `Model`
- [](#VUID-PatchVertices-PatchVertices-04309)VUID-PatchVertices-PatchVertices-04309  
  The variable decorated with `PatchVertices` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-PatchVertices-PatchVertices-04310)VUID-PatchVertices-PatchVertices-04310  
  The variable decorated with `PatchVertices` **must** be declared as a scalar 32-bit integer value

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#PatchVertices)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0