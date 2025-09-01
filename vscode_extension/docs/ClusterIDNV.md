# ClusterIDNV(3) Manual Page

## Name

ClusterIDNV - Contains the triangle cluster ID of a hit triangle in cluster acceleration structure



## [](#_description)Description

`ClusterIDNV`

A variable decorated with the `ClusterIDNV` decoration will contain the triangle cluster ID of a hit triangle in a cluster acceleration structure if the current ray hit a triangle primitive or `-1` otherwise.

Valid Usage

- [](#VUID-ClusterIDNV-ClusterIDNV-10531)VUID-ClusterIDNV-ClusterIDNV-10531  
  The `ClusterIDNV` decoration **must** be used only within the `AnyHitKHR`, or `ClosestHitKHR` `Execution` `Model`
- [](#VUID-ClusterIDNV-ClusterIDNV-10532)VUID-ClusterIDNV-ClusterIDNV-10532  
  The variable decorated with `ClusterIDNV` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-ClusterIDNV-ClusterIDNV-10533)VUID-ClusterIDNV-ClusterIDNV-10533  
  The variable decorated with `ClusterIDNV` **must** be declared as a scalar 32-bit integer value

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#ClusterIDNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0