# LocalInvocationId(3) Manual Page

## Name

LocalInvocationId - Local invocation ID



## [](#_description)Description

`LocalInvocationId`

Decorating a variable with the `LocalInvocationId` built-in decoration will make that variable contain the location of the current cluster culling, task, mesh, or compute shader invocation within the local workgroup. Each component ranges from zero through to the size of the workgroup in that dimension minus one.

Note

If the size of the workgroup in a particular dimension is one, then the `LocalInvocationId` in that dimension will be zero. If the workgroup is effectively two-dimensional, then `LocalInvocationId.z` will be zero. If the workgroup is effectively one-dimensional, then both `LocalInvocationId.y` and `LocalInvocationId.z` will be zero.

Valid Usage

- [](#VUID-LocalInvocationId-LocalInvocationId-04281)VUID-LocalInvocationId-LocalInvocationId-04281  
  The `LocalInvocationId` decoration **must** be used only within the `GLCompute`, `MeshEXT`, `TaskEXT`, `MeshNV`, or `TaskNV` `Execution` `Model`
- [](#VUID-LocalInvocationId-LocalInvocationId-04282)VUID-LocalInvocationId-LocalInvocationId-04282  
  The variable decorated with `LocalInvocationId` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-LocalInvocationId-LocalInvocationId-04283)VUID-LocalInvocationId-LocalInvocationId-04283  
  The variable decorated with `LocalInvocationId` **must** be declared as a three-component vector of 32-bit integer values

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#LocalInvocationId).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0