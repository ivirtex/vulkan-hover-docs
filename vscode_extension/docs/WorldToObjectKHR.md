# WorldToObjectKHR(3) Manual Page

## Name

WorldToObjectKHR - Transformation matrix from world to object space



## [](#_description)Description

`WorldToObjectKHR`

A variable decorated with the `WorldToObjectKHR` decoration will contain the current world-to-object transformation matrix, which is determined by the instance of the current intersection.

Valid Usage

- [](#VUID-WorldToObjectKHR-WorldToObjectKHR-04434)VUID-WorldToObjectKHR-WorldToObjectKHR-04434  
  The `WorldToObjectKHR` decoration **must** be used only within the `IntersectionKHR`, `AnyHitKHR`, or `ClosestHitKHR` `Execution` `Model`
- [](#VUID-WorldToObjectKHR-WorldToObjectKHR-04435)VUID-WorldToObjectKHR-WorldToObjectKHR-04435  
  The variable decorated with `WorldToObjectKHR` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-WorldToObjectKHR-WorldToObjectKHR-04436)VUID-WorldToObjectKHR-WorldToObjectKHR-04436  
  The variable decorated with `WorldToObjectKHR` **must** be declared as a matrix with four columns of three-component vectors of 32-bit floating-point values

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#WorldToObjectKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0