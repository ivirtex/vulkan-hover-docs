# RayTminKHR(3) Manual Page

## Name

RayTminKHR - Minimum T value of a ray



## [](#_description)Description

`RayTminKHR`

A variable decorated with the `RayTminKHR` decoration will contain the parametric tmin value of the ray being processed. The value is independent of the space in which the ray origin and direction exist. The value is the parameter passed into the [pipeline trace ray](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#glossary-pipeline-trace-ray) instruction.

The tmin value remains constant for the duration of the ray query.

Valid Usage

- [](#VUID-RayTminKHR-RayTminKHR-04351)VUID-RayTminKHR-RayTminKHR-04351  
  The `RayTminKHR` decoration **must** be used only within the `IntersectionKHR`, `AnyHitKHR`, `ClosestHitKHR`, or `MissKHR` `Execution` `Model`
- [](#VUID-RayTminKHR-RayTminKHR-04352)VUID-RayTminKHR-RayTminKHR-04352  
  The variable decorated with `RayTminKHR` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-RayTminKHR-RayTminKHR-04353)VUID-RayTminKHR-RayTminKHR-04353  
  The variable decorated with `RayTminKHR` **must** be declared as a scalar 32-bit floating-point value

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#RayTminKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0