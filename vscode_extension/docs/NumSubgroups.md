# NumSubgroups(3) Manual Page

## Name

NumSubgroups - Number of subgroups in a workgroup



## [](#_description)Description

`NumSubgroups`

Decorating a variable with the `NumSubgroups` built-in decoration will make that variable contain the number of subgroups in the local workgroup.

Valid Usage

- [](#VUID-NumSubgroups-NumSubgroups-04293)VUID-NumSubgroups-NumSubgroups-04293  
  The `NumSubgroups` decoration **must** be used only within the `GLCompute`, `MeshEXT`, `TaskEXT`, `MeshNV`, or `TaskNV` `Execution` `Model`
- [](#VUID-NumSubgroups-NumSubgroups-04294)VUID-NumSubgroups-NumSubgroups-04294  
  The variable decorated with `NumSubgroups` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-NumSubgroups-NumSubgroups-04295)VUID-NumSubgroups-NumSubgroups-04295  
  The variable decorated with `NumSubgroups` **must** be declared as a scalar 32-bit integer value

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#NumSubgroups).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0