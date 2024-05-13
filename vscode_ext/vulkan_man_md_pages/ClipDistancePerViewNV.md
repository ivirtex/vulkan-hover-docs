# ClipDistancePerViewNV(3) Manual Page

## Name

ClipDistancePerViewNV - Application-specified clip distances per view



## <a href="#_description" class="anchor"></a>Description

`ClipDistancePerViewNV`  
Decorating a variable with the `ClipDistancePerViewNV` built-in
decoration will make that variable contain the per-view clip distances.
The per-view clip distances have the same semantics as `ClipDistance`.

Valid Usage

- <a href="#VUID-ClipDistancePerViewNV-ClipDistancePerViewNV-04192"
  id="VUID-ClipDistancePerViewNV-ClipDistancePerViewNV-04192"></a>
  VUID-ClipDistancePerViewNV-ClipDistancePerViewNV-04192  
  The `ClipDistancePerViewNV` decoration **must** be used only within
  the `MeshNV` `Execution` `Model`

- <a href="#VUID-ClipDistancePerViewNV-ClipDistancePerViewNV-04193"
  id="VUID-ClipDistancePerViewNV-ClipDistancePerViewNV-04193"></a>
  VUID-ClipDistancePerViewNV-ClipDistancePerViewNV-04193  
  The variable decorated with `ClipDistancePerViewNV` **must** be
  declared using the `Output` `Storage` `Class`

- <a href="#VUID-ClipDistancePerViewNV-ClipDistancePerViewNV-04194"
  id="VUID-ClipDistancePerViewNV-ClipDistancePerViewNV-04194"></a>
  VUID-ClipDistancePerViewNV-ClipDistancePerViewNV-04194  
  The variable decorated with `ClipDistancePerViewNV` **must** also be
  decorated with the `PerViewNV` decoration

- <a href="#VUID-ClipDistancePerViewNV-ClipDistancePerViewNV-04195"
  id="VUID-ClipDistancePerViewNV-ClipDistancePerViewNV-04195"></a>
  VUID-ClipDistancePerViewNV-ClipDistancePerViewNV-04195  
  The variable decorated with `ClipDistancePerViewNV` **must** be
  declared as a two-dimensional array of 32-bit floating-point values

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#ClipDistancePerViewNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
