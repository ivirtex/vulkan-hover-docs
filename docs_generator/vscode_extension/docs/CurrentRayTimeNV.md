# CurrentRayTimeNV(3) Manual Page

## Name

CurrentRayTimeNV - Time value of a ray intersection



## [](#_description)Description

`CurrentRayTimeNV`

A variable decorated with the `CurrentRayTimeNV` decoration contains the time value passed in to `OpTraceRayMotionNV` which called this shader.

Valid Usage

- [](#VUID-CurrentRayTimeNV-CurrentRayTimeNV-04942)VUID-CurrentRayTimeNV-CurrentRayTimeNV-04942  
  The `CurrentRayTimeNV` decoration **must** be used only within the `IntersectionKHR`, `AnyHitKHR`, `ClosestHitKHR`, or `MissKHR` `Execution` `Model`
- [](#VUID-CurrentRayTimeNV-CurrentRayTimeNV-04943)VUID-CurrentRayTimeNV-CurrentRayTimeNV-04943  
  The variable decorated with `CurrentRayTimeNV` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-CurrentRayTimeNV-CurrentRayTimeNV-04944)VUID-CurrentRayTimeNV-CurrentRayTimeNV-04944  
  The variable decorated with `CurrentRayTimeNV` **must** be declared as a scalar 32-bit floating-point value

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#CurrentRayTimeNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0