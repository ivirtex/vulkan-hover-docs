# ClipDistancePerViewNV(3) Manual Page

## Name

ClipDistancePerViewNV - Application-specified clip distances per view



## [](#_description)Description

`ClipDistancePerViewNV`

Decorating a variable with the `ClipDistancePerViewNV` built-in decoration will make that variable contain the per-view clip distances. The per-view clip distances have the same semantics as `ClipDistance`.

Valid Usage

- [](#VUID-ClipDistancePerViewNV-ClipDistancePerViewNV-04192)VUID-ClipDistancePerViewNV-ClipDistancePerViewNV-04192  
  The `ClipDistancePerViewNV` decoration **must** be used only within the `MeshNV` `Execution` `Model`
- [](#VUID-ClipDistancePerViewNV-ClipDistancePerViewNV-04193)VUID-ClipDistancePerViewNV-ClipDistancePerViewNV-04193  
  The variable decorated with `ClipDistancePerViewNV` **must** be declared using the `Output` `Storage` `Class`
- [](#VUID-ClipDistancePerViewNV-ClipDistancePerViewNV-04194)VUID-ClipDistancePerViewNV-ClipDistancePerViewNV-04194  
  The variable decorated with `ClipDistancePerViewNV` **must** also be decorated with the `PerViewNV` decoration
- [](#VUID-ClipDistancePerViewNV-ClipDistancePerViewNV-04195)VUID-ClipDistancePerViewNV-ClipDistancePerViewNV-04195  
  The variable decorated with `ClipDistancePerViewNV` **must** be declared as a two-dimensional array of 32-bit floating-point values

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#ClipDistancePerViewNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0