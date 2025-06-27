# FragCoord(3) Manual Page

## Name

FragCoord - Screen-space coordinate of the fragment center



## <a href="#_description" class="anchor"></a>Description

`FragCoord`  
Decorating a variable with the `FragCoord` built-in decoration will make
that variable contain the framebuffer coordinate (x,y,z,w1​) of the
fragment being processed. The (x,y) coordinate (0,0) is the upper left
corner of the upper left pixel in the framebuffer.

When <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-sampleshading"
target="_blank" rel="noopener">Sample Shading</a> is enabled, the x and
y components of `FragCoord` reflect the location of one of the samples
corresponding to the shader invocation.

Otherwise, the x and y components of `FragCoord` reflect the location of
the center of the fragment.

The z component of `FragCoord` is the interpolated depth value of the
primitive.

The w component is the interpolated w1​.

The `Centroid` interpolation decoration is ignored, but allowed, on
`FragCoord`.

Valid Usage

- <a href="#VUID-FragCoord-FragCoord-04210"
  id="VUID-FragCoord-FragCoord-04210"></a>
  VUID-FragCoord-FragCoord-04210  
  The `FragCoord` decoration **must** be used only within the `Fragment`
  `Execution` `Model`

- <a href="#VUID-FragCoord-FragCoord-04211"
  id="VUID-FragCoord-FragCoord-04211"></a>
  VUID-FragCoord-FragCoord-04211  
  The variable decorated with `FragCoord` **must** be declared using the
  `Input` `Storage` `Class`

- <a href="#VUID-FragCoord-FragCoord-04212"
  id="VUID-FragCoord-FragCoord-04212"></a>
  VUID-FragCoord-FragCoord-04212  
  The variable decorated with `FragCoord` **must** be declared as a
  four-component vector of 32-bit floating-point values

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#FragCoord"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
