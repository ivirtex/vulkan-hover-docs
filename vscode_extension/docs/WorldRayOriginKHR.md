# WorldRayOriginKHR(3) Manual Page

## Name

WorldRayOriginKHR - Ray origin in world space



## [](#_description)Description

`WorldRayOriginKHR`

A variable decorated with the `WorldRayOriginKHR` decoration will specify the origin of the ray being processed, in world space. The value is the parameter passed into the [pipeline trace ray](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#glossary-pipeline-trace-ray) instruction.

Valid Usage

- [](#VUID-WorldRayOriginKHR-WorldRayOriginKHR-04431)VUID-WorldRayOriginKHR-WorldRayOriginKHR-04431  
  The `WorldRayOriginKHR` decoration **must** be used only within the `IntersectionKHR`, `AnyHitKHR`, `ClosestHitKHR`, or `MissKHR` `Execution` `Model`
- [](#VUID-WorldRayOriginKHR-WorldRayOriginKHR-04432)VUID-WorldRayOriginKHR-WorldRayOriginKHR-04432  
  The variable decorated with `WorldRayOriginKHR` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-WorldRayOriginKHR-WorldRayOriginKHR-04433)VUID-WorldRayOriginKHR-WorldRayOriginKHR-04433  
  The variable decorated with `WorldRayOriginKHR` **must** be declared as a three-component vector of 32-bit floating-point values

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#WorldRayOriginKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0