# HitIsLSSNV(3) Manual Page

## Name

HitIsLSSNV - Indicates if a LSS primitive was hit



## [](#_description)Description

`HitIsLSSNV`

A variable decorated with the `HitIsLSSNV` decoration will contain a non-zero value if the current ray hit a LSS primitive or zero otherwise.

Valid Usage

- [](#VUID-HitIsLSSNV-HitIsLSSNV-10516)VUID-HitIsLSSNV-HitIsLSSNV-10516  
  The `HitIsLSSNV` decoration **must** be used only within the `AnyHitKHR`, or `ClosestHitKHR` `Execution` `Model`
- [](#VUID-HitIsLSSNV-HitIsLSSNV-10517)VUID-HitIsLSSNV-HitIsLSSNV-10517  
  The variable decorated with `HitIsLSSNV` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-HitIsLSSNV-HitIsLSSNV-10518)VUID-HitIsLSSNV-HitIsLSSNV-10518  
  The variable decorated with `HitIsLSSNV` **must** be declared as a boolean value

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#HitIsLSSNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0