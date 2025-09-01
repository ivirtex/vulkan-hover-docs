# BaryCoordKHR(3) Manual Page

## Name

BaryCoordKHR - Barycentric coordinates of a fragment



## [](#_description)Description

`BaryCoordKHR`

The `BaryCoordKHR` decoration **can** be used to decorate a fragment shader input variable. This variable will contain a three-component floating-point vector with barycentric weights that indicate the location of the fragment relative to the screen-space locations of vertices of its primitive, obtained using perspective interpolation.

Valid Usage

- [](#VUID-BaryCoordKHR-BaryCoordKHR-04154)VUID-BaryCoordKHR-BaryCoordKHR-04154  
  The `BaryCoordKHR` decoration **must** be used only within the `Fragment` `Execution` `Model`
- [](#VUID-BaryCoordKHR-BaryCoordKHR-04155)VUID-BaryCoordKHR-BaryCoordKHR-04155  
  The variable decorated with `BaryCoordKHR` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-BaryCoordKHR-BaryCoordKHR-04156)VUID-BaryCoordKHR-BaryCoordKHR-04156  
  The variable decorated with `BaryCoordKHR` **must** be declared as a three-component vector of 32-bit floating-point values

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#BaryCoordKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0