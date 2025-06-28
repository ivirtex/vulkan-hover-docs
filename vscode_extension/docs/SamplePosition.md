# SamplePosition(3) Manual Page

## Name

SamplePosition - Position of a shaded sample



## [](#_description)Description

`SamplePosition`

Decorating a variable with the `SamplePosition` built-in decoration will make that variable contain the sub-pixel position of the sample being shaded. The top left of the pixel is considered to be at coordinate (0,0) and the bottom right of the pixel is considered to be at coordinate (1,1).

If the render pass has a fragment density map attachment, the variable will instead contain the sub-fragment position of the sample being shaded. The top left of the fragment is considered to be at coordinate (0,0) and the bottom right of the fragment is considered to be at coordinate (1,1) for any fragment area.

If a fragment shader entry point’s interface includes an input variable decorated with `SamplePosition`, [Sample Shading](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#primsrast-sampleshading) is considered enabled with a `minSampleShading` value of 1.0.

If the current pipeline uses [custom sample locations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#primsrast-samplelocations) the value of any variable decorated with the `SamplePosition` built-in decoration is undefined.

Valid Usage

- [](#VUID-SamplePosition-SamplePosition-04360)VUID-SamplePosition-SamplePosition-04360  
  The `SamplePosition` decoration **must** be used only within the `Fragment` `Execution` `Model`
- [](#VUID-SamplePosition-SamplePosition-04361)VUID-SamplePosition-SamplePosition-04361  
  The variable decorated with `SamplePosition` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-SamplePosition-SamplePosition-04362)VUID-SamplePosition-SamplePosition-04362  
  The variable decorated with `SamplePosition` **must** be declared as a two-component vector of 32-bit floating-point values

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#SamplePosition)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0