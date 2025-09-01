# ObjectRayDirectionKHR(3) Manual Page

## Name

ObjectRayDirectionKHR - Ray direction in object space



## [](#_description)Description

`ObjectRayDirectionKHR`

A variable decorated with the `ObjectRayDirectionKHR` decoration will specify the direction of the ray being processed, in object space.

Valid Usage

- [](#VUID-ObjectRayDirectionKHR-ObjectRayDirectionKHR-04299)VUID-ObjectRayDirectionKHR-ObjectRayDirectionKHR-04299  
  The `ObjectRayDirectionKHR` decoration **must** be used only within the `IntersectionKHR`, `AnyHitKHR`, or `ClosestHitKHR` `Execution` `Model`
- [](#VUID-ObjectRayDirectionKHR-ObjectRayDirectionKHR-04300)VUID-ObjectRayDirectionKHR-ObjectRayDirectionKHR-04300  
  The variable decorated with `ObjectRayDirectionKHR` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-ObjectRayDirectionKHR-ObjectRayDirectionKHR-04301)VUID-ObjectRayDirectionKHR-ObjectRayDirectionKHR-04301  
  The variable decorated with `ObjectRayDirectionKHR` **must** be declared as a three-component vector of 32-bit floating-point values

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#ObjectRayDirectionKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0