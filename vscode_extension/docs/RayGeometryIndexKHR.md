# RayGeometryIndexKHR(3) Manual Page

## Name

RayGeometryIndexKHR - Geometry index in a ray shader



## [](#_description)Description

`RayGeometryIndexKHR`

A variable decorated with the `RayGeometryIndexKHR` decoration will contain the [geometry index](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#acceleration-structure-geometry-index) for the acceleration structure geometry currently being shaded.

Valid Usage

- [](#VUID-RayGeometryIndexKHR-RayGeometryIndexKHR-04345)VUID-RayGeometryIndexKHR-RayGeometryIndexKHR-04345  
  The `RayGeometryIndexKHR` decoration **must** be used only within the `IntersectionKHR`, `AnyHitKHR`, or `ClosestHitKHR` `Execution` `Model`
- [](#VUID-RayGeometryIndexKHR-RayGeometryIndexKHR-04346)VUID-RayGeometryIndexKHR-RayGeometryIndexKHR-04346  
  The variable decorated with `RayGeometryIndexKHR` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-RayGeometryIndexKHR-RayGeometryIndexKHR-04347)VUID-RayGeometryIndexKHR-RayGeometryIndexKHR-04347  
  The variable decorated with `RayGeometryIndexKHR` **must** be declared as a scalar 32-bit integer value

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#RayGeometryIndexKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0