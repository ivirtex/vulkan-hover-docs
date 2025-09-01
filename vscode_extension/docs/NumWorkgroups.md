# NumWorkgroups(3) Manual Page

## Name

NumWorkgroups - Number of workgroups in a dispatch



## [](#_description)Description

`NumWorkgroups`

Decorating a variable with the `NumWorkgroups` built-in decoration will make that variable contain the number of local workgroups that are part of the dispatch that the invocation belongs to. Each component is equal to the values of the workgroup count parameters passed into the dispatching commands.

Valid Usage

- [](#VUID-NumWorkgroups-NumWorkgroups-04296)VUID-NumWorkgroups-NumWorkgroups-04296  
  The `NumWorkgroups` decoration **must** be used only within the `GLCompute`, `MeshEXT`, or `TaskEXT` `Execution` `Model`
- [](#VUID-NumWorkgroups-NumWorkgroups-04297)VUID-NumWorkgroups-NumWorkgroups-04297  
  The variable decorated with `NumWorkgroups` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-NumWorkgroups-NumWorkgroups-04298)VUID-NumWorkgroups-NumWorkgroups-04298  
  The variable decorated with `NumWorkgroups` **must** be declared as a three-component vector of 32-bit integer values

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#NumWorkgroups).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0