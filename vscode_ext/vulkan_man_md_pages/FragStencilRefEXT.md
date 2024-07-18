# FragStencilRefEXT(3) Manual Page

## Name

FragStencilRefEXT - Application-specified stencil reference value used
in stencil tests



## <a href="#_description" class="anchor"></a>Description

`FragStencilRefEXT`  
Decorating a variable with the `FragStencilRefEXT` built-in decoration
will make that variable contain the new stencil reference value for all
samples covered by the fragment. This value will be used as the stencil
reference value used in stencil testing.

To write to `FragStencilRefEXT`, a shader **must** declare the
`StencilRefReplacingEXT` execution mode. If a shader declares the
`StencilRefReplacingEXT` execution mode and there is an execution path
through the shader that does not set `FragStencilRefEXT`, then the
fragment’s stencil reference value is undefined for executions of the
shader that take that path.

Only the least significant **s** bits of the integer value of the
variable decorated with `FragStencilRefEXT` are considered for stencil
testing, where **s** is the number of bits in the stencil framebuffer
attachment, and higher order bits are discarded.

See <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fragops-shader-stencilrefreplacement"
target="_blank" rel="noopener">fragment shader stencil reference
replacement</a> for more details.

Valid Usage

- <a href="#VUID-FragStencilRefEXT-FragStencilRefEXT-04223"
  id="VUID-FragStencilRefEXT-FragStencilRefEXT-04223"></a>
  VUID-FragStencilRefEXT-FragStencilRefEXT-04223  
  The `FragStencilRefEXT` decoration **must** be used only within the
  `Fragment` `Execution` `Model`

- <a href="#VUID-FragStencilRefEXT-FragStencilRefEXT-04224"
  id="VUID-FragStencilRefEXT-FragStencilRefEXT-04224"></a>
  VUID-FragStencilRefEXT-FragStencilRefEXT-04224  
  The variable decorated with `FragStencilRefEXT` **must** be declared
  using the `Output` `Storage` `Class`

- <a href="#VUID-FragStencilRefEXT-FragStencilRefEXT-04225"
  id="VUID-FragStencilRefEXT-FragStencilRefEXT-04225"></a>
  VUID-FragStencilRefEXT-FragStencilRefEXT-04225  
  The variable decorated with `FragStencilRefEXT` **must** be declared
  as a scalar integer value

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#FragStencilRefEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
