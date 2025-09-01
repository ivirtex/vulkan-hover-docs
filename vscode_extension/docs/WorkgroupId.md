# WorkgroupId(3) Manual Page

## Name

WorkgroupId - Workgroup ID of a shader



## [](#_description)Description

`WorkgroupId`

Decorating a variable with the `WorkgroupId` built-in decoration will make that variable contain the global coordinate of the local workgroup that the current invocation is a member of. Each component is in the range [base,base + count), where base and count are based on the parameters passed into the dispatching or drawing commands in each dimension.

Valid Usage

- [](#VUID-WorkgroupId-WorkgroupId-04422)VUID-WorkgroupId-WorkgroupId-04422  
  The `WorkgroupId` decoration **must** be used only within the `GLCompute`, `MeshEXT`, `TaskEXT`, `MeshNV`, or `TaskNV` `Execution` `Model`
- [](#VUID-WorkgroupId-WorkgroupId-04423)VUID-WorkgroupId-WorkgroupId-04423  
  The variable decorated with `WorkgroupId` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-WorkgroupId-WorkgroupId-04424)VUID-WorkgroupId-WorkgroupId-04424  
  The variable decorated with `WorkgroupId` **must** be declared as a three-component vector of 32-bit integer values

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#WorkgroupId).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0