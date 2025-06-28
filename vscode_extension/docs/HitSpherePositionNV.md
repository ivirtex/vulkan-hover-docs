# HitSpherePositionNV(3) Manual Page

## Name

HitSpherePositionNV - Contains the position of the hit sphere



## [](#_description)Description

`HitSpherePositionNV`

A variable decorated with the `HitSpherePositionNV` decoration will contain the position of sphere primitive intersected by current ray.

Valid Usage

- [](#VUID-HitSpherePositionNV-HitSpherePositionNV-10519)VUID-HitSpherePositionNV-HitSpherePositionNV-10519  
  The `HitSpherePositionNV` decoration **must** be used only within the `AnyHitKHR`, or `ClosestHitKHR` `Execution` `Model`
- [](#VUID-HitSpherePositionNV-HitSpherePositionNV-10520)VUID-HitSpherePositionNV-HitSpherePositionNV-10520  
  The variable decorated with `HitSpherePositionNV` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-HitSpherePositionNV-HitSpherePositionNV-10521)VUID-HitSpherePositionNV-HitSpherePositionNV-10521  
  The variable decorated with `HitSpherePositionNV` **must** be declared as a three-component vector of 32-bit floating-point values

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#HitSpherePositionNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0