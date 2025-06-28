# LayerPerViewNV(3) Manual Page

## Name

LayerPerViewNV - Layer index per view for layered rendering



## [](#_description)Description

`LayerPerViewNV`

Decorating a variable with the `LayerPerViewNV` built-in decoration will make that variable contain the per-view layer information. The per-view layer has the same semantics as `Layer`, for each view.

Valid Usage

- [](#VUID-LayerPerViewNV-LayerPerViewNV-04277)VUID-LayerPerViewNV-LayerPerViewNV-04277  
  The `LayerPerViewNV` decoration **must** be used only within the `MeshNV` `Execution` `Model`
- [](#VUID-LayerPerViewNV-LayerPerViewNV-04278)VUID-LayerPerViewNV-LayerPerViewNV-04278  
  The variable decorated with `LayerPerViewNV` **must** be declared using the `Output` `Storage` `Class`
- [](#VUID-LayerPerViewNV-LayerPerViewNV-04279)VUID-LayerPerViewNV-LayerPerViewNV-04279  
  The variable decorated with `LayerPerViewNV` **must** also be decorated with the `PerViewNV` decoration
- [](#VUID-LayerPerViewNV-LayerPerViewNV-04280)VUID-LayerPerViewNV-LayerPerViewNV-04280  
  The variable decorated with `LayerPerViewNV` **must** be declared as an array of scalar 32-bit integer values

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#LayerPerViewNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0