# HitTNV(3) Manual Page

## Name

HitTNV - T value of a ray intersection



## [](#_description)Description

`HitTNV`

A variable decorated with the `HitTNV` decoration is equivalent to a variable decorated with the `RayTmaxKHR` decoration.

Valid Usage

- [](#VUID-HitTNV-HitTNV-04245)VUID-HitTNV-HitTNV-04245  
  The `HitTNV` decoration **must** be used only within the `AnyHitNV` or `ClosestHitNV` `Execution` `Model`
- [](#VUID-HitTNV-HitTNV-04246)VUID-HitTNV-HitTNV-04246  
  The variable decorated with `HitTNV` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-HitTNV-HitTNV-04247)VUID-HitTNV-HitTNV-04247  
  The variable decorated with `HitTNV` **must** be declared as a scalar 32-bit floating-point value

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#HitTNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0