# PatchVertices(3) Manual Page

## Name

PatchVertices - Number of vertices in an input patch



## <a href="#_description" class="anchor"></a>Description

`PatchVertices`  
Decorating a variable with the `PatchVertices` built-in decoration will
make that variable contain the number of vertices in the input patch
being processed by the shader. In a Tessellation Control Shader, this is
the same as the name:patchControlPoints member of
[VkPipelineTessellationStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineTessellationStateCreateInfo.html).
In a Tessellation Evaluation Shader, `PatchVertices` is equal to the
tessellation control output patch size. When the same shader is used in
different pipelines where the patch sizes are configured differently,
the value of the `PatchVertices` variable will also differ.

Valid Usage

- <a href="#VUID-PatchVertices-PatchVertices-04308"
  id="VUID-PatchVertices-PatchVertices-04308"></a>
  VUID-PatchVertices-PatchVertices-04308  
  The `PatchVertices` decoration **must** be used only within the
  `TessellationControl` or `TessellationEvaluation` `Execution` `Model`

- <a href="#VUID-PatchVertices-PatchVertices-04309"
  id="VUID-PatchVertices-PatchVertices-04309"></a>
  VUID-PatchVertices-PatchVertices-04309  
  The variable decorated with `PatchVertices` **must** be declared using
  the `Input` `Storage` `Class`

- <a href="#VUID-PatchVertices-PatchVertices-04310"
  id="VUID-PatchVertices-PatchVertices-04310"></a>
  VUID-PatchVertices-PatchVertices-04310  
  The variable decorated with `PatchVertices` **must** be declared as a
  scalar 32-bit integer value

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#PatchVertices"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
