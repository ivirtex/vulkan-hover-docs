# HitSphereRadiusNV(3) Manual Page

## Name

HitSphereRadiusNV - Contains the radius of the hit sphere



## [](#_description)Description

`HitSphereRadiusNV`

A variable decorated with the `HitSphereRadiusNV` decoration will contain the radius of sphere primitive intersected by current ray.

Valid Usage

- [](#VUID-HitSphereRadiusNV-HitSphereRadiusNV-10522)VUID-HitSphereRadiusNV-HitSphereRadiusNV-10522  
  The `HitSphereRadiusNV` decoration **must** be used only within the `AnyHitKHR`, or `ClosestHitKHR` `Execution` `Model`
- [](#VUID-HitSphereRadiusNV-HitSphereRadiusNV-10523)VUID-HitSphereRadiusNV-HitSphereRadiusNV-10523  
  The variable decorated with `HitSphereRadiusNV` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-HitSphereRadiusNV-HitSphereRadiusNV-10524)VUID-HitSphereRadiusNV-HitSphereRadiusNV-10524  
  The variable decorated with `HitSphereRadiusNV` **must** be declared as a scalar 32-bit floating-point value

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#HitSphereRadiusNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0