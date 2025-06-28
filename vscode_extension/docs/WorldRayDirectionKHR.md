# WorldRayDirectionKHR(3) Manual Page

## Name

WorldRayDirectionKHR - Ray direction in world space



## [](#_description)Description

`WorldRayDirectionKHR`

A variable decorated with the `WorldRayDirectionKHR` decoration will specify the direction of the ray being processed, in world space. The value is the parameter passed into the [pipeline trace ray](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#glossary-pipeline-trace-ray) instruction.

Valid Usage

- [](#VUID-WorldRayDirectionKHR-WorldRayDirectionKHR-04428)VUID-WorldRayDirectionKHR-WorldRayDirectionKHR-04428  
  The `WorldRayDirectionKHR` decoration **must** be used only within the `IntersectionKHR`, `AnyHitKHR`, `ClosestHitKHR`, or `MissKHR` `Execution` `Model`
- [](#VUID-WorldRayDirectionKHR-WorldRayDirectionKHR-04429)VUID-WorldRayDirectionKHR-WorldRayDirectionKHR-04429  
  The variable decorated with `WorldRayDirectionKHR` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-WorldRayDirectionKHR-WorldRayDirectionKHR-04430)VUID-WorldRayDirectionKHR-WorldRayDirectionKHR-04430  
  The variable decorated with `WorldRayDirectionKHR` **must** be declared as a three-component vector of 32-bit floating-point values

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#WorldRayDirectionKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0