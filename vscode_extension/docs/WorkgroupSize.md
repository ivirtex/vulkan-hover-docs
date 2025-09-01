# WorkgroupSize(3) Manual Page

## Name

WorkgroupSize - Size of a workgroup



## [](#_description)Description

`WorkgroupSize`

Note

SPIR-V 1.6 deprecated `WorkgroupSize` in favor of using the `LocalSizeId` Execution Mode instead. Support for `LocalSizeId` was added with `VK_KHR_maintenance4` and promoted to core in Version 1.3.

Decorating an object with the `WorkgroupSize` built-in decoration will make that object contain the dimensions of a local workgroup. If an object is decorated with the `WorkgroupSize` decoration, this takes precedence over any `LocalSize` or `LocalSizeId` execution mode.

Valid Usage

- [](#VUID-WorkgroupSize-WorkgroupSize-04425)VUID-WorkgroupSize-WorkgroupSize-04425  
  The `WorkgroupSize` decoration **must** be used only within the `GLCompute`, `MeshEXT`, `TaskEXT`, `MeshNV`, or `TaskNV` `Execution` `Model`
- [](#VUID-WorkgroupSize-WorkgroupSize-04426)VUID-WorkgroupSize-WorkgroupSize-04426  
  The variable decorated with `WorkgroupSize` **must** be a specialization constant or a constant
- [](#VUID-WorkgroupSize-WorkgroupSize-04427)VUID-WorkgroupSize-WorkgroupSize-04427  
  The variable decorated with `WorkgroupSize` **must** be declared as a three-component vector of 32-bit integer values
- [](#VUID-WorkgroupSize-TileShadingRateQCOM-10635)VUID-WorkgroupSize-TileShadingRateQCOM-10635  
  If the `TileShadingRateQCOM` `Execution` `Mode` is used, variables decorated with`WorkgroupSize` **must** be declared using the `Input` `Storage` `Class`

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#WorkgroupSize).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0