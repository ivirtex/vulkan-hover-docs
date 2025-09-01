# BaryCoordSmoothCentroidAMD(3) Manual Page

## Name

BaryCoordSmoothCentroidAMD - Barycentric coordinates of a fragment centroid



## [](#_description)Description

`BaryCoordSmoothCentroidAMD`

The `BaryCoordSmoothCentroidAMD` decoration **can** be used to decorate a fragment shader input variable. This variable will contain the (I,J) pair of the barycentric coordinates corresponding to the fragment evaluated using perspective interpolation at the centroid. The K coordinate of the barycentric coordinates **can** be derived given the identity I + J + K = 1.0.

Valid Usage

- [](#VUID-BaryCoordSmoothCentroidAMD-BaryCoordSmoothCentroidAMD-04175)VUID-BaryCoordSmoothCentroidAMD-BaryCoordSmoothCentroidAMD-04175  
  The `BaryCoordSmoothCentroidAMD` decoration **must** be used only within the `Fragment` `Execution` `Model`
- [](#VUID-BaryCoordSmoothCentroidAMD-BaryCoordSmoothCentroidAMD-04176)VUID-BaryCoordSmoothCentroidAMD-BaryCoordSmoothCentroidAMD-04176  
  The variable decorated with `BaryCoordSmoothCentroidAMD` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-BaryCoordSmoothCentroidAMD-BaryCoordSmoothCentroidAMD-04177)VUID-BaryCoordSmoothCentroidAMD-BaryCoordSmoothCentroidAMD-04177  
  The variable decorated with `BaryCoordSmoothCentroidAMD` **must** be declared as a two-component vector of 32-bit floating-point values

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#BaryCoordSmoothCentroidAMD).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0