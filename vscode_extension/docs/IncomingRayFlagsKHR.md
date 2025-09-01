# IncomingRayFlagsKHR(3) Manual Page

## Name

IncomingRayFlagsKHR - Flags used to trace a ray



## [](#_description)Description

`IncomingRayFlagsKHR`

A variable with the `IncomingRayFlagsKHR` decoration will contain the ray flags passed in to the trace call that invoked this particular shader. Setting pipeline flags on the ray tracing pipeline **must** not cause any corresponding flags to be set in variables with this decoration.

Valid Usage

- [](#VUID-IncomingRayFlagsKHR-IncomingRayFlagsKHR-04248)VUID-IncomingRayFlagsKHR-IncomingRayFlagsKHR-04248  
  The `IncomingRayFlagsKHR` decoration **must** be used only within the `IntersectionKHR`, `AnyHitKHR`, `ClosestHitKHR`, or `MissKHR` `Execution` `Model`
- [](#VUID-IncomingRayFlagsKHR-IncomingRayFlagsKHR-04249)VUID-IncomingRayFlagsKHR-IncomingRayFlagsKHR-04249  
  The variable decorated with `IncomingRayFlagsKHR` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-IncomingRayFlagsKHR-IncomingRayFlagsKHR-04250)VUID-IncomingRayFlagsKHR-IncomingRayFlagsKHR-04250  
  The variable decorated with `IncomingRayFlagsKHR` **must** be declared as a scalar 32-bit integer value

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#IncomingRayFlagsKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0