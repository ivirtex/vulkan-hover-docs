# VertexCountHUAWEI(3) Manual Page

## Name

VertexCountHUAWEI - cluster culling shader output variable



## [](#_description)Description

`VertexCountHUAWEI`

The `VertexCountHUAWEI` decoration can be used to decorate a cluster culling shader output variable,this non-indexed mode specific variable will contain an integer value that specifies the number of vertices in a cluster to draw.

Valid Usage

- [](#VUID-VertexCountHUAWEI-VertexCountHUAWEI-07809)VUID-VertexCountHUAWEI-VertexCountHUAWEI-07809  
  The `VertexCountHUAWEI` decoration **must** be used only within the `ClusterCullingHUAWEI` `Execution` `Model`
- [](#VUID-VertexCountHUAWEI-VertexCountHUAWEI-07810)VUID-VertexCountHUAWEI-VertexCountHUAWEI-07810  
  The variable decorated with `VertexCountHUAWEI` **must** be declared as a scalar 32-bit integer value

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VertexCountHUAWEI).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0