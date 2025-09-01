# BaryCoordNoPerspKHR(3) Manual Page

## Name

BaryCoordNoPerspKHR - Barycentric coordinates of a fragment in screen-space



## [](#_description)Description

`BaryCoordNoPerspKHR`

The `BaryCoordNoPerspKHR` decoration **can** be used to decorate a fragment shader input variable. This variable will contain a three-component floating-point vector with barycentric weights that indicate the location of the fragment relative to the screen-space locations of vertices of its primitive, obtained using linear interpolation.

Valid Usage

- [](#VUID-BaryCoordNoPerspKHR-BaryCoordNoPerspKHR-04160)VUID-BaryCoordNoPerspKHR-BaryCoordNoPerspKHR-04160  
  The `BaryCoordNoPerspKHR` decoration **must** be used only within the `Fragment` `Execution` `Model`
- [](#VUID-BaryCoordNoPerspKHR-BaryCoordNoPerspKHR-04161)VUID-BaryCoordNoPerspKHR-BaryCoordNoPerspKHR-04161  
  The variable decorated with `BaryCoordNoPerspKHR` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-BaryCoordNoPerspKHR-BaryCoordNoPerspKHR-04162)VUID-BaryCoordNoPerspKHR-BaryCoordNoPerspKHR-04162  
  The variable decorated with `BaryCoordNoPerspKHR` **must** be declared as a three-component vector of 32-bit floating-point values

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#BaryCoordNoPerspKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0