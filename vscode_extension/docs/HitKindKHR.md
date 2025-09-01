# HitKindKHR(3) Manual Page

## Name

HitKindKHR - Kind of hit that triggered an any-hit or closest hit ray shader



## [](#_description)Description

`HitKindKHR`

A variable decorated with the `HitKindKHR` decoration will describe the intersection that triggered the execution of the current shader. The values are determined by the intersection shader. For user-defined intersection shaders this is the value that was passed to the “Hit Kind” operand of `OpReportIntersectionKHR`. For triangle intersection candidates, this will be one of `HitKindFrontFacingTriangleKHR` or `HitKindBackFacingTriangleKHR`.

Valid Usage

- [](#VUID-HitKindKHR-HitKindKHR-04242)VUID-HitKindKHR-HitKindKHR-04242  
  The `HitKindKHR` decoration **must** be used only within the `AnyHitKHR` or `ClosestHitKHR` `Execution` `Model`
- [](#VUID-HitKindKHR-HitKindKHR-04243)VUID-HitKindKHR-HitKindKHR-04243  
  The variable decorated with `HitKindKHR` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-HitKindKHR-HitKindKHR-04244)VUID-HitKindKHR-HitKindKHR-04244  
  The variable decorated with `HitKindKHR` **must** be declared as a scalar 32-bit integer value

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#HitKindKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0