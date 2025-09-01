# ObjectRayOriginKHR(3) Manual Page

## Name

ObjectRayOriginKHR - Ray origin in object space



## [](#_description)Description

`ObjectRayOriginKHR`

A variable decorated with the `ObjectRayOriginKHR` decoration will specify the origin of the ray being processed, in object space.

Valid Usage

- [](#VUID-ObjectRayOriginKHR-ObjectRayOriginKHR-04302)VUID-ObjectRayOriginKHR-ObjectRayOriginKHR-04302  
  The `ObjectRayOriginKHR` decoration **must** be used only within the `IntersectionKHR`, `AnyHitKHR`, or `ClosestHitKHR` `Execution` `Model`
- [](#VUID-ObjectRayOriginKHR-ObjectRayOriginKHR-04303)VUID-ObjectRayOriginKHR-ObjectRayOriginKHR-04303  
  The variable decorated with `ObjectRayOriginKHR` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-ObjectRayOriginKHR-ObjectRayOriginKHR-04304)VUID-ObjectRayOriginKHR-ObjectRayOriginKHR-04304  
  The variable decorated with `ObjectRayOriginKHR` **must** be declared as a three-component vector of 32-bit floating-point values

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#ObjectRayOriginKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0