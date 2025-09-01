# BaryCoordNoPerspAMD(3) Manual Page

## Name

BaryCoordNoPerspAMD - Barycentric coordinates of a fragment center in screen-space



## [](#_description)Description

`BaryCoordNoPerspAMD`

The `BaryCoordNoPerspAMD` decoration **can** be used to decorate a fragment shader input variable. This variable will contain the (I,J) pair of the barycentric coordinates corresponding to the fragment evaluated using linear interpolation at the fragment’s center. The K coordinate of the barycentric coordinates **can** be derived given the identity I + J + K = 1.0.

Valid Usage

- [](#VUID-BaryCoordNoPerspAMD-BaryCoordNoPerspAMD-04157)VUID-BaryCoordNoPerspAMD-BaryCoordNoPerspAMD-04157  
  The `BaryCoordNoPerspAMD` decoration **must** be used only within the `Fragment` `Execution` `Model`
- [](#VUID-BaryCoordNoPerspAMD-BaryCoordNoPerspAMD-04158)VUID-BaryCoordNoPerspAMD-BaryCoordNoPerspAMD-04158  
  The variable decorated with `BaryCoordNoPerspAMD` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-BaryCoordNoPerspAMD-BaryCoordNoPerspAMD-04159)VUID-BaryCoordNoPerspAMD-BaryCoordNoPerspAMD-04159  
  The variable decorated with `BaryCoordNoPerspAMD` **must** be declared as a two-component vector of 32-bit floating-point values

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#BaryCoordNoPerspAMD).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0