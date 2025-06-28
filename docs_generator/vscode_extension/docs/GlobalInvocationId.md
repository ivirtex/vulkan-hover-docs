# GlobalInvocationId(3) Manual Page

## Name

GlobalInvocationId - Global invocation ID



## [](#_description)Description

`GlobalInvocationId`

Decorating a variable with the `GlobalInvocationId` built-in decoration will make that variable contain the location of the current invocation within the global workgroup. Each component is equal to the index of the local workgroup multiplied by the size of the local workgroup plus `LocalInvocationId`.

Valid Usage

- [](#VUID-GlobalInvocationId-GlobalInvocationId-04236)VUID-GlobalInvocationId-GlobalInvocationId-04236  
  The `GlobalInvocationId` decoration **must** be used only within the `GLCompute`, `MeshEXT`, `TaskEXT`, `MeshNV`, or `TaskNV` `Execution` `Model`
- [](#VUID-GlobalInvocationId-GlobalInvocationId-04237)VUID-GlobalInvocationId-GlobalInvocationId-04237  
  The variable decorated with `GlobalInvocationId` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-GlobalInvocationId-GlobalInvocationId-04238)VUID-GlobalInvocationId-GlobalInvocationId-04238  
  The variable decorated with `GlobalInvocationId` **must** be declared as a three-component vector of 32-bit integer values

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#GlobalInvocationId)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0