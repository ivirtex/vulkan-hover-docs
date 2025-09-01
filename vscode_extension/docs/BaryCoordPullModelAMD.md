# BaryCoordPullModelAMD(3) Manual Page

## Name

BaryCoordPullModelAMD - Inverse barycentric coordinates of a fragment center



## [](#_description)Description

`BaryCoordPullModelAMD`

The `BaryCoordPullModelAMD` decoration **can** be used to decorate a fragment shader input variable. This variable will contain (1/W, 1/I, 1/J) evaluated at the fragment center and **can** be used to calculate gradients and then interpolate I, J, and W at any desired sample location.

Valid Usage

- [](#VUID-BaryCoordPullModelAMD-BaryCoordPullModelAMD-04169)VUID-BaryCoordPullModelAMD-BaryCoordPullModelAMD-04169  
  The `BaryCoordPullModelAMD` decoration **must** be used only within the `Fragment` `Execution` `Model`
- [](#VUID-BaryCoordPullModelAMD-BaryCoordPullModelAMD-04170)VUID-BaryCoordPullModelAMD-BaryCoordPullModelAMD-04170  
  The variable decorated with `BaryCoordPullModelAMD` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-BaryCoordPullModelAMD-BaryCoordPullModelAMD-04171)VUID-BaryCoordPullModelAMD-BaryCoordPullModelAMD-04171  
  The variable decorated with `BaryCoordPullModelAMD` **must** be declared as a three-component vector of 32-bit floating-point values

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#BaryCoordPullModelAMD).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0