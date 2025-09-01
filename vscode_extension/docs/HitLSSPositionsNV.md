# HitLSSPositionsNV(3) Manual Page

## Name

HitLSSPositionsNV - Contains the position of the hit LSS primitive



## [](#_description)Description

`HitLSSPositionsNV`

A variable decorated with the `HitLSSPositionsNV` decoration will contain the position of the LSS primitive intersected by current ray.

Valid Usage

- [](#VUID-HitLSSPositionsNV-HitLSSPositionsNV-10525)VUID-HitLSSPositionsNV-HitLSSPositionsNV-10525  
  The `HitLSSPositionsNV` decoration **must** be used only within the `AnyHitKHR`, or `ClosestHitKHR` `Execution` `Model`
- [](#VUID-HitLSSPositionsNV-HitLSSPositionsNV-10526)VUID-HitLSSPositionsNV-HitLSSPositionsNV-10526  
  The variable decorated with `HitLSSPositionsNV` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-HitLSSPositionsNV-HitLSSPositionsNV-10527)VUID-HitLSSPositionsNV-HitLSSPositionsNV-10527  
  The variable decorated with `HitLSSPositionsNV` **must** be declared as an array of size two, containing three-component vector of 32-bit floating-point values

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#HitLSSPositionsNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0