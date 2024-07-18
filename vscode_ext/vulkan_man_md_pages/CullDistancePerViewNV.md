# CullDistancePerViewNV(3) Manual Page

## Name

CullDistancePerViewNV - Application-specified cull distances per view



## <a href="#_description" class="anchor"></a>Description

`CullDistancePerViewNV`  
Decorating a variable with the `CullDistancePerViewNV` built-in
decoration will make that variable contain the per-view cull distances.
The per-view cull distances have the same semantics as `CullDistance`.

Valid Usage

- <a href="#VUID-CullDistancePerViewNV-CullDistancePerViewNV-04201"
  id="VUID-CullDistancePerViewNV-CullDistancePerViewNV-04201"></a>
  VUID-CullDistancePerViewNV-CullDistancePerViewNV-04201  
  The `CullDistancePerViewNV` decoration **must** be used only within
  the `MeshNV` `Execution` `Model`

- <a href="#VUID-CullDistancePerViewNV-CullDistancePerViewNV-04202"
  id="VUID-CullDistancePerViewNV-CullDistancePerViewNV-04202"></a>
  VUID-CullDistancePerViewNV-CullDistancePerViewNV-04202  
  The variable decorated with `CullDistancePerViewNV` **must** be
  declared using the `Output` `Storage` `Class`

- <a href="#VUID-CullDistancePerViewNV-CullDistancePerViewNV-04203"
  id="VUID-CullDistancePerViewNV-CullDistancePerViewNV-04203"></a>
  VUID-CullDistancePerViewNV-CullDistancePerViewNV-04203  
  The variable decorated with `CullDistancePerViewNV` **must** also be
  decorated with the `PerViewNV` decoration

- <a href="#VUID-CullDistancePerViewNV-CullDistancePerViewNV-04204"
  id="VUID-CullDistancePerViewNV-CullDistancePerViewNV-04204"></a>
  VUID-CullDistancePerViewNV-CullDistancePerViewNV-04204  
  The variable decorated with `CullDistancePerViewNV` **must** be
  declared as a two-dimensional array of 32-bit floating-point values

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#CullDistancePerViewNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
