# VertexOffsetHUAWEI(3) Manual Page

## Name

VertexOffsetHUAWEI - cluster culling shader output variable



## [](#_description)Description

`VertexOffsetHUAWEI`

The `VertexOffsetHUAWEI` decoration can be used to decorate a cluster culling shader output variable,this indexed mode specific variable will contain an integer value that specifies an offset value added to the vertex index of a cluster before indexing into the vertex buffer.

Valid Usage

- [](#VUID-VertexOffsetHUAWEI-VertexOffsetHUAWEI-07811)VUID-VertexOffsetHUAWEI-VertexOffsetHUAWEI-07811  
  The `VertexOffsetHUAWEI` decoration **must** be used only within the `ClusterCullingHUAWEI` `Execution` `Model`
- [](#VUID-VertexOffsetHUAWEI-VertexOffsetHUAWEI-07812)VUID-VertexOffsetHUAWEI-VertexOffsetHUAWEI-07812  
  The variable decorated with `VertexOffsetHUAWEI` **must** be declared as a scalar 32-bit integer value

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VertexOffsetHUAWEI).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0