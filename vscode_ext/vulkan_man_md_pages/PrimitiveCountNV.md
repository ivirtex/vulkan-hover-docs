# PrimitiveCountNV(3) Manual Page

## Name

PrimitiveCountNV - Number of primitives output by a mesh shader



## <a href="#_description" class="anchor"></a>Description

`PrimitiveCountNV`  
Decorating a variable with the `PrimitiveCountNV` decoration will make
that variable contain the primitive count. The primitive count specifies
the number of primitives in the output mesh produced by the mesh shader
that will be processed by subsequent pipeline stages.

Valid Usage

- <a href="#VUID-PrimitiveCountNV-PrimitiveCountNV-04327"
  id="VUID-PrimitiveCountNV-PrimitiveCountNV-04327"></a>
  VUID-PrimitiveCountNV-PrimitiveCountNV-04327  
  The `PrimitiveCountNV` decoration **must** be used only within the
  `MeshNV` `Execution` `Model`

- <a href="#VUID-PrimitiveCountNV-PrimitiveCountNV-04328"
  id="VUID-PrimitiveCountNV-PrimitiveCountNV-04328"></a>
  VUID-PrimitiveCountNV-PrimitiveCountNV-04328  
  The variable decorated with `PrimitiveCountNV` **must** be declared
  using the `Output` `Storage` `Class`

- <a href="#VUID-PrimitiveCountNV-PrimitiveCountNV-04329"
  id="VUID-PrimitiveCountNV-PrimitiveCountNV-04329"></a>
  VUID-PrimitiveCountNV-PrimitiveCountNV-04329  
  The variable decorated with `PrimitiveCountNV` **must** be declared as
  a scalar 32-bit integer value

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#PrimitiveCountNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
