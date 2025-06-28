# ObjectToWorldKHR(3) Manual Page

## Name

ObjectToWorldKHR - Transformation matrix from object to world space



## [](#_description)Description

`ObjectToWorldKHR`

A variable decorated with the `ObjectToWorldKHR` decoration will contain the current object-to-world transformation matrix, which is determined by the instance of the current intersection.

Valid Usage

- [](#VUID-ObjectToWorldKHR-ObjectToWorldKHR-04305)VUID-ObjectToWorldKHR-ObjectToWorldKHR-04305  
  The `ObjectToWorldKHR` decoration **must** be used only within the `IntersectionKHR`, `AnyHitKHR`, or `ClosestHitKHR` `Execution` `Model`
- [](#VUID-ObjectToWorldKHR-ObjectToWorldKHR-04306)VUID-ObjectToWorldKHR-ObjectToWorldKHR-04306  
  The variable decorated with `ObjectToWorldKHR` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-ObjectToWorldKHR-ObjectToWorldKHR-04307)VUID-ObjectToWorldKHR-ObjectToWorldKHR-04307  
  The variable decorated with `ObjectToWorldKHR` **must** be declared as a matrix with four columns of three-component vectors of 32-bit floating-point values

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#ObjectToWorldKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0