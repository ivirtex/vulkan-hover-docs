# TessLevelOuter(3) Manual Page

## Name

TessLevelOuter - Outer tessellation levels



## <a href="#_description" class="anchor"></a>Description

`TessLevelOuter`  
Decorating a variable with the `TessLevelOuter` built-in decoration will
make that variable contain the outer tessellation levels for the current
patch.

In tessellation control shaders, the variable decorated with
`TessLevelOuter` **can** be written to, controlling the tessellation
factors for the resulting patch. These values are used by the
tessellator to control primitive tessellation and **can** be read by
tessellation evaluation shaders.

In tessellation evaluation shaders, the variable decorated with
`TessLevelOuter` **can** read the values written by the tessellation
control shader.

Valid Usage

- <a href="#VUID-TessLevelOuter-TessLevelOuter-04390"
  id="VUID-TessLevelOuter-TessLevelOuter-04390"></a>
  VUID-TessLevelOuter-TessLevelOuter-04390  
  The `TessLevelOuter` decoration **must** be used only within the
  `TessellationControl` or `TessellationEvaluation` `Execution` `Model`

- <a href="#VUID-TessLevelOuter-TessLevelOuter-04391"
  id="VUID-TessLevelOuter-TessLevelOuter-04391"></a>
  VUID-TessLevelOuter-TessLevelOuter-04391  
  The variable decorated with `TessLevelOuter` within the
  `TessellationControl` `Execution` `Model` **must** be declared using
  the `Output` `Storage` `Class`

- <a href="#VUID-TessLevelOuter-TessLevelOuter-04392"
  id="VUID-TessLevelOuter-TessLevelOuter-04392"></a>
  VUID-TessLevelOuter-TessLevelOuter-04392  
  The variable decorated with `TessLevelOuter` within the
  `TessellationEvaluation` `Execution` `Model` **must** be declared
  using the `Input` `Storage` `Class`

- <a href="#VUID-TessLevelOuter-TessLevelOuter-04393"
  id="VUID-TessLevelOuter-TessLevelOuter-04393"></a>
  VUID-TessLevelOuter-TessLevelOuter-04393  
  The variable decorated with `TessLevelOuter` **must** be declared as
  an array of size four, containing 32-bit floating-point values

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#TessLevelOuter"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
