# PrimitivePointIndicesEXT(3) Manual Page

## Name

PrimitivePointIndicesEXT - Indices of point primitives in a mesh shader



## [](#_description)Description

`PrimitivePointIndicesEXT`

Decorating a variable with the `PrimitivePointIndicesEXT` decoration will make that variable contain the output array of vertex index values for point primitives.

Valid Usage

- [](#VUID-PrimitivePointIndicesEXT-PrimitivePointIndicesEXT-07041)VUID-PrimitivePointIndicesEXT-PrimitivePointIndicesEXT-07041  
  The `PrimitivePointIndicesEXT` decoration **must** be used only within the `MeshEXT` `Execution` `Model`
- [](#VUID-PrimitivePointIndicesEXT-PrimitivePointIndicesEXT-07042)VUID-PrimitivePointIndicesEXT-PrimitivePointIndicesEXT-07042  
  The `PrimitivePointIndicesEXT` decoration **must** be used with the `OutputPoints` `Execution` `Mode`
- [](#VUID-PrimitivePointIndicesEXT-PrimitivePointIndicesEXT-07043)VUID-PrimitivePointIndicesEXT-PrimitivePointIndicesEXT-07043  
  The variable decorated with `PrimitivePointIndicesEXT` **must** be declared using the `Output` `Storage` `Class`
- [](#VUID-PrimitivePointIndicesEXT-PrimitivePointIndicesEXT-07044)VUID-PrimitivePointIndicesEXT-PrimitivePointIndicesEXT-07044  
  The variable decorated with `PrimitivePointIndicesEXT` **must** be declared as an array of scalar 32-bit integer values
- [](#VUID-PrimitivePointIndicesEXT-PrimitivePointIndicesEXT-07045)VUID-PrimitivePointIndicesEXT-PrimitivePointIndicesEXT-07045  
  All index values of the array decorated with `PrimitivePointIndicesEXT` **must** be in the range \[0, N-1], where N is the value specified by the `OutputVertices` `Execution` `Mode`
- [](#VUID-PrimitivePointIndicesEXT-PrimitivePointIndicesEXT-07046)VUID-PrimitivePointIndicesEXT-PrimitivePointIndicesEXT-07046  
  The size of the array decorated with `PrimitivePointIndicesEXT` **must** match the value specified by `OutputPrimitivesEXT`

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#PrimitivePointIndicesEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0