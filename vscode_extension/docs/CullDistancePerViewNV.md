# CullDistancePerViewNV(3) Manual Page

## Name

CullDistancePerViewNV - Application-specified cull distances per view



## [](#_description)Description

`CullDistancePerViewNV`

Decorating a variable with the `CullDistancePerViewNV` built-in decoration will make that variable contain the per-view cull distances. The per-view cull distances have the same semantics as `CullDistance`.

Valid Usage

- [](#VUID-CullDistancePerViewNV-CullDistancePerViewNV-04201)VUID-CullDistancePerViewNV-CullDistancePerViewNV-04201  
  The `CullDistancePerViewNV` decoration **must** be used only within the `MeshNV` `Execution` `Model`
- [](#VUID-CullDistancePerViewNV-CullDistancePerViewNV-04202)VUID-CullDistancePerViewNV-CullDistancePerViewNV-04202  
  The variable decorated with `CullDistancePerViewNV` **must** be declared using the `Output` `Storage` `Class`
- [](#VUID-CullDistancePerViewNV-CullDistancePerViewNV-04203)VUID-CullDistancePerViewNV-CullDistancePerViewNV-04203  
  The variable decorated with `CullDistancePerViewNV` **must** also be decorated with the `PerViewNV` decoration
- [](#VUID-CullDistancePerViewNV-CullDistancePerViewNV-04204)VUID-CullDistancePerViewNV-CullDistancePerViewNV-04204  
  The variable decorated with `CullDistancePerViewNV` **must** be declared as a two-dimensional array of 32-bit floating-point values

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#CullDistancePerViewNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0