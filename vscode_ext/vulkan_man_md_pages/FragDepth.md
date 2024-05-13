# FragDepth(3) Manual Page

## Name

FragDepth - Application-specified depth for depth testing



## <a href="#_description" class="anchor"></a>Description

`FragDepth`  
To have a shader supply a fragment-depth value, the shader **must**
declare the `DepthReplacing` execution mode. Such a shader’s
fragment-depth value will come from the variable decorated with the
`FragDepth` built-in decoration.

This value will be used for any subsequent depth testing performed by
the implementation or writes to the depth attachment. See <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fragops-shader-depthreplacement"
target="_blank" rel="noopener">fragment shader depth replacement</a> for
details.

Valid Usage

- <a href="#VUID-FragDepth-FragDepth-04213"
  id="VUID-FragDepth-FragDepth-04213"></a>
  VUID-FragDepth-FragDepth-04213  
  The `FragDepth` decoration **must** be used only within the `Fragment`
  `Execution` `Model`

- <a href="#VUID-FragDepth-FragDepth-04214"
  id="VUID-FragDepth-FragDepth-04214"></a>
  VUID-FragDepth-FragDepth-04214  
  The variable decorated with `FragDepth` **must** be declared using the
  `Output` `Storage` `Class`

- <a href="#VUID-FragDepth-FragDepth-04215"
  id="VUID-FragDepth-FragDepth-04215"></a>
  VUID-FragDepth-FragDepth-04215  
  The variable decorated with `FragDepth` **must** be declared as a
  scalar 32-bit floating-point value

- <a href="#VUID-FragDepth-FragDepth-04216"
  id="VUID-FragDepth-FragDepth-04216"></a>
  VUID-FragDepth-FragDepth-04216  
  If the shader dynamically writes to the variable decorated with
  `FragDepth`, the `DepthReplacing` `Execution` `Mode` **must** be
  declared

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#FragDepth"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
