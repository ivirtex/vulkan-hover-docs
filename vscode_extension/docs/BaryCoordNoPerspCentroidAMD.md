# BaryCoordNoPerspCentroidAMD(3) Manual Page

## Name

BaryCoordNoPerspCentroidAMD - Barycentric coordinates of a fragment centroid in screen-space



## [](#_description)Description

`BaryCoordNoPerspCentroidAMD`

The `BaryCoordNoPerspCentroidAMD` decoration **can** be used to decorate a fragment shader input variable. This variable will contain the (I,J) pair of the barycentric coordinates corresponding to the fragment evaluated using linear interpolation at the centroid. The K coordinate of the barycentric coordinates **can** be derived given the identity I + J + K = 1.0.

Valid Usage

- [](#VUID-BaryCoordNoPerspCentroidAMD-BaryCoordNoPerspCentroidAMD-04163)VUID-BaryCoordNoPerspCentroidAMD-BaryCoordNoPerspCentroidAMD-04163  
  The `BaryCoordNoPerspCentroidAMD` decoration **must** be used only within the `Fragment` `Execution` `Model`
- [](#VUID-BaryCoordNoPerspCentroidAMD-BaryCoordNoPerspCentroidAMD-04164)VUID-BaryCoordNoPerspCentroidAMD-BaryCoordNoPerspCentroidAMD-04164  
  The variable decorated with `BaryCoordNoPerspCentroidAMD` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-BaryCoordNoPerspCentroidAMD-BaryCoordNoPerspCentroidAMD-04165)VUID-BaryCoordNoPerspCentroidAMD-BaryCoordNoPerspCentroidAMD-04165  
  The variable decorated with `BaryCoordNoPerspCentroidAMD` **must** be declared as a three-component vector of 32-bit floating-point values

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#BaryCoordNoPerspCentroidAMD)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0