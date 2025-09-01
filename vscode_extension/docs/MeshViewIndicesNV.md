# MeshViewIndicesNV(3) Manual Page

## Name

MeshViewIndicesNV - Indices of views processed by a mesh or task shader



## [](#_description)Description

`MeshViewIndicesNV`

Decorating a variable with the `MeshViewIndicesNV` built-in decoration will make that variable contain the mesh view indices. The mesh view indices is an array of values where each element holds the view number of one of the views being processed by the current mesh or task shader invocations. The values of array elements with indices greater than or equal to `MeshViewCountNV` are undefined. If the value of `MeshViewIndicesNV`\[i] is j, then any outputs decorated with `PerViewNV` will take on the value of array element i when processing primitives for view index j.

Valid Usage

- [](#VUID-MeshViewIndicesNV-MeshViewIndicesNV-04290)VUID-MeshViewIndicesNV-MeshViewIndicesNV-04290  
  The `MeshViewIndicesNV` decoration **must** be used only within the `MeshNV` or `TaskNV` `Execution` `Model`
- [](#VUID-MeshViewIndicesNV-MeshViewIndicesNV-04291)VUID-MeshViewIndicesNV-MeshViewIndicesNV-04291  
  The variable decorated with `MeshViewIndicesNV` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-MeshViewIndicesNV-MeshViewIndicesNV-04292)VUID-MeshViewIndicesNV-MeshViewIndicesNV-04292  
  The variable decorated with `MeshViewIndicesNV` **must** be declared as an array of scalar 32-bit integer values

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#MeshViewIndicesNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0