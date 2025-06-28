# BaryCoordSmoothAMD(3) Manual Page

## Name

BaryCoordSmoothAMD - Barycentric coordinates of a fragment center



## [](#_description)Description

`BaryCoordSmoothAMD`

The `BaryCoordSmoothAMD` decoration **can** be used to decorate a fragment shader input variable. This variable will contain the (I,J) pair of the barycentric coordinates corresponding to the fragment evaluated using perspective interpolation at the fragment’s center. The K coordinate of the barycentric coordinates **can** be derived given the identity I + J + K = 1.0.

Valid Usage

- [](#VUID-BaryCoordSmoothAMD-BaryCoordSmoothAMD-04172)VUID-BaryCoordSmoothAMD-BaryCoordSmoothAMD-04172  
  The `BaryCoordSmoothAMD` decoration **must** be used only within the `Fragment` `Execution` `Model`
- [](#VUID-BaryCoordSmoothAMD-BaryCoordSmoothAMD-04173)VUID-BaryCoordSmoothAMD-BaryCoordSmoothAMD-04173  
  The variable decorated with `BaryCoordSmoothAMD` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-BaryCoordSmoothAMD-BaryCoordSmoothAMD-04174)VUID-BaryCoordSmoothAMD-BaryCoordSmoothAMD-04174  
  The variable decorated with `BaryCoordSmoothAMD` **must** be declared as a two-component vector of 32-bit floating-point values

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#BaryCoordSmoothAMD)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0