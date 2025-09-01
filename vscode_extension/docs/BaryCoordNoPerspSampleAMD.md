# BaryCoordNoPerspSampleAMD(3) Manual Page

## Name

BaryCoordNoPerspSampleAMD - Barycentric coordinates of a sample center in screen-space



## [](#_description)Description

`BaryCoordNoPerspSampleAMD`

The `BaryCoordNoPerspSampleAMD` decoration **can** be used to decorate a fragment shader input variable. This variable will contain the (I,J) pair of the barycentric coordinates corresponding to the fragment evaluated using linear interpolation at each covered sample. The K coordinate of the barycentric coordinates **can** be derived given the identity I + J + K = 1.0.

Valid Usage

- [](#VUID-BaryCoordNoPerspSampleAMD-BaryCoordNoPerspSampleAMD-04166)VUID-BaryCoordNoPerspSampleAMD-BaryCoordNoPerspSampleAMD-04166  
  The `BaryCoordNoPerspSampleAMD` decoration **must** be used only within the `Fragment` `Execution` `Model`
- [](#VUID-BaryCoordNoPerspSampleAMD-BaryCoordNoPerspSampleAMD-04167)VUID-BaryCoordNoPerspSampleAMD-BaryCoordNoPerspSampleAMD-04167  
  The variable decorated with `BaryCoordNoPerspSampleAMD` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-BaryCoordNoPerspSampleAMD-BaryCoordNoPerspSampleAMD-04168)VUID-BaryCoordNoPerspSampleAMD-BaryCoordNoPerspSampleAMD-04168  
  The variable decorated with `BaryCoordNoPerspSampleAMD` **must** be declared as a two-component vector of 32-bit floating-point values

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#BaryCoordNoPerspSampleAMD).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0