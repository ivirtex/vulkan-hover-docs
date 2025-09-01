# FragStencilRefEXT(3) Manual Page

## Name

FragStencilRefEXT - Application-specified stencil reference value used in stencil tests



## [](#_description)Description

`FragStencilRefEXT`

Decorating a variable with the `FragStencilRefEXT` built-in decoration will make that variable contain the new stencil reference value for all samples covered by the fragment. This value will be used as the stencil reference value used in stencil testing.

To write to `FragStencilRefEXT`, a shader **must** declare the `StencilRefReplacingEXT` execution mode. If a shader declares the `StencilRefReplacingEXT` execution mode and there is an execution path through the shader that does not set `FragStencilRefEXT`, then the fragment’s stencil reference value is undefined for executions of the shader that take that path.

Only the least significant **s** bits of the integer value of the variable decorated with `FragStencilRefEXT` are considered for stencil testing, where **s** is the number of bits in the stencil framebuffer attachment, and higher order bits are discarded.

See [fragment shader stencil reference replacement](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#fragops-shader-stencilrefreplacement) for more details.

Valid Usage

- [](#VUID-FragStencilRefEXT-FragStencilRefEXT-04223)VUID-FragStencilRefEXT-FragStencilRefEXT-04223  
  The `FragStencilRefEXT` decoration **must** be used only within the `Fragment` `Execution` `Model`
- [](#VUID-FragStencilRefEXT-FragStencilRefEXT-04224)VUID-FragStencilRefEXT-FragStencilRefEXT-04224  
  The variable decorated with `FragStencilRefEXT` **must** be declared using the `Output` `Storage` `Class`
- [](#VUID-FragStencilRefEXT-FragStencilRefEXT-04225)VUID-FragStencilRefEXT-FragStencilRefEXT-04225  
  The variable decorated with `FragStencilRefEXT` **must** be declared as a scalar integer value

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#FragStencilRefEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0