# InvocationId(3) Manual Page

## Name

InvocationId - Invocation ID in a geometry or tessellation control
shader



## <a href="#_description" class="anchor"></a>Description

`InvocationId`  
Decorating a variable with the `InvocationId` built-in decoration will
make that variable contain the index of the current shader invocation in
a geometry shader, or the index of the output patch vertex in a
tessellation control shader.

In a geometry shader, the index of the current shader invocation ranges
from zero to the number of <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#geometry-invocations"
target="_blank" rel="noopener">instances</a> declared in the shader
minus one. If the instance count of the geometry shader is one or is not
specified, then `InvocationId` will be zero.

Valid Usage

- <a href="#VUID-InvocationId-InvocationId-04257"
  id="VUID-InvocationId-InvocationId-04257"></a>
  VUID-InvocationId-InvocationId-04257  
  The `InvocationId` decoration **must** be used only within the
  `TessellationControl` or `Geometry` `Execution` `Model`

- <a href="#VUID-InvocationId-InvocationId-04258"
  id="VUID-InvocationId-InvocationId-04258"></a>
  VUID-InvocationId-InvocationId-04258  
  The variable decorated with `InvocationId` **must** be declared using
  the `Input` `Storage` `Class`

- <a href="#VUID-InvocationId-InvocationId-04259"
  id="VUID-InvocationId-InvocationId-04259"></a>
  VUID-InvocationId-InvocationId-04259  
  The variable decorated with `InvocationId` **must** be declared as a
  scalar 32-bit integer value

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#InvocationId"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
