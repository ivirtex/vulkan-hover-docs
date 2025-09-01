# HitLSSRadiiNV(3) Manual Page

## Name

HitLSSRadiiNV - Contains the radii of the hit LSS primitive



## [](#_description)Description

`HitLSSRadiiNV`

A variable decorated with the `HitLSSRadiiNV` decoration will contain the radii of LSS primitive intersected by current ray.

Valid Usage

- [](#VUID-HitLSSRadiiNV-HitLSSRadiiNV-10528)VUID-HitLSSRadiiNV-HitLSSRadiiNV-10528  
  The `HitLSSRadiiNV` decoration **must** be used only within the `AnyHitKHR`, or `ClosestHitKHR` `Execution` `Model`
- [](#VUID-HitLSSRadiiNV-HitLSSRadiiNV-10529)VUID-HitLSSRadiiNV-HitLSSRadiiNV-10529  
  The variable decorated with `HitLSSRadiiNV` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-HitLSSRadiiNV-HitLSSRadiiNV-10530)VUID-HitLSSRadiiNV-HitLSSRadiiNV-10530  
  The variable decorated with `HitLSSRadiiNV` **must** be declared as an array of size two, containing 32-bit floating-point values

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#HitLSSRadiiNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0