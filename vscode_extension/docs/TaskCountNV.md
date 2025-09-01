# TaskCountNV(3) Manual Page

## Name

TaskCountNV - Number of mesh shader workgroups that will be generated



## [](#_description)Description

`TaskCountNV`

Decorating a variable with the `TaskCountNV` decoration will make that variable contain the task count. The task count specifies the number of subsequent mesh shader workgroups that get generated upon completion of the task shader.

Valid Usage

- [](#VUID-TaskCountNV-TaskCountNV-04384)VUID-TaskCountNV-TaskCountNV-04384  
  The `TaskCountNV` decoration **must** be used only within the `TaskNV` `Execution` `Model`
- [](#VUID-TaskCountNV-TaskCountNV-04385)VUID-TaskCountNV-TaskCountNV-04385  
  The variable decorated with `TaskCountNV` **must** be declared using the `Output` `Storage` `Class`
- [](#VUID-TaskCountNV-TaskCountNV-04386)VUID-TaskCountNV-TaskCountNV-04386  
  The variable decorated with `TaskCountNV` **must** be declared as a scalar 32-bit integer value

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#TaskCountNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0