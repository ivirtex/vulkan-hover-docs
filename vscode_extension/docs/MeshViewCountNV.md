# MeshViewCountNV(3) Manual Page

## Name

MeshViewCountNV - Number of views processed by a mesh or task shader



## [](#_description)Description

`MeshViewCountNV`

Decorating a variable with the `MeshViewCountNV` built-in decoration will make that variable contain the number of views processed by the current mesh or task shader invocations.

Valid Usage

- [](#VUID-MeshViewCountNV-MeshViewCountNV-04287)VUID-MeshViewCountNV-MeshViewCountNV-04287  
  The `MeshViewCountNV` decoration **must** be used only within the `MeshNV` or `TaskNV` `Execution` `Model`
- [](#VUID-MeshViewCountNV-MeshViewCountNV-04288)VUID-MeshViewCountNV-MeshViewCountNV-04288  
  The variable decorated with `MeshViewCountNV` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-MeshViewCountNV-MeshViewCountNV-04289)VUID-MeshViewCountNV-MeshViewCountNV-04289  
  The variable decorated with `MeshViewCountNV` **must** be declared as a scalar 32-bit integer value

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#MeshViewCountNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0