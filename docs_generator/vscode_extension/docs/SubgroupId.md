# SubgroupId(3) Manual Page

## Name

SubgroupId - Subgroup ID



## [](#_description)Description

`SubgroupId`

Decorating a variable with the `SubgroupId` built-in decoration will make that variable contain the index of the subgroup within the local workgroup. This variable is in range \[0, `NumSubgroups`-1].

Valid Usage

- [](#VUID-SubgroupId-SubgroupId-04367)VUID-SubgroupId-SubgroupId-04367  
  The `SubgroupId` decoration **must** be used only within the `GLCompute`, `MeshEXT`, `TaskEXT`, `MeshNV`, or `TaskNV` `Execution` `Model`
- [](#VUID-SubgroupId-SubgroupId-04368)VUID-SubgroupId-SubgroupId-04368  
  The variable decorated with `SubgroupId` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-SubgroupId-SubgroupId-04369)VUID-SubgroupId-SubgroupId-04369  
  The variable decorated with `SubgroupId` **must** be declared as a scalar 32-bit integer value

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#SubgroupId)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0