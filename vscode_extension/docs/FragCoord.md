# FragCoord(3) Manual Page

## Name

FragCoord - Coordinates of the fragment



## [](#_description)Description

`FragCoord`

Decorating a variable with the `FragCoord` built-in decoration will make that variable contain the coordinates (x,y,z,1/w) of the fragment being processed.

The (x,y) values are the framebuffer coordinates (xf,yf) of the fragment.

When [Sample Shading](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#primsrast-sampleshading) is enabled, the x and y components of `FragCoord` reflect the location of one of the samples corresponding to the shader invocation.

Otherwise, the x and y components of `FragCoord` reflect the location of the center of the fragment.

The z component of `FragCoord` is the interpolated depth value of the primitive.

The w component is the interpolated w1​.

The `Centroid` interpolation decoration is ignored, but allowed, on `FragCoord`.

Valid Usage

- [](#VUID-FragCoord-FragCoord-04210)VUID-FragCoord-FragCoord-04210  
  The `FragCoord` decoration **must** be used only within the `Fragment` `Execution` `Model`
- [](#VUID-FragCoord-FragCoord-04211)VUID-FragCoord-FragCoord-04211  
  The variable decorated with `FragCoord` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-FragCoord-FragCoord-04212)VUID-FragCoord-FragCoord-04212  
  The variable decorated with `FragCoord` **must** be declared as a four-component vector of 32-bit floating-point values

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#FragCoord).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0