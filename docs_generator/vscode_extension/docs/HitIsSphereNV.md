# HitIsSphereNV(3) Manual Page

## Name

HitIsSphereNV - Indicates if a sphere primitive was hit



## [](#_description)Description

`HitIsSphereNV`

A variable decorated with the `HitIsSphereNV` decoration will contain a non-zero value if the current ray hit a sphere primitive or zero otherwise.

Valid Usage

- [](#VUID-HitIsSphereNV-HitIsSphereNV-10513)VUID-HitIsSphereNV-HitIsSphereNV-10513  
  The `HitIsSphereNV` decoration **must** be used only within the `AnyHitKHR`, or `ClosestHitKHR` `Execution` `Model`
- [](#VUID-HitIsSphereNV-HitIsSphereNV-10514)VUID-HitIsSphereNV-HitIsSphereNV-10514  
  The variable decorated with `HitIsSphereNV` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-HitIsSphereNV-HitIsSphereNV-10515)VUID-HitIsSphereNV-HitIsSphereNV-10515  
  The variable decorated with `HitIsSphereNV` **must** be declared as a boolean value

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#HitIsSphereNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0