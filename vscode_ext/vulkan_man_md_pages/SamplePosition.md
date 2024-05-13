# SamplePosition(3) Manual Page

## Name

SamplePosition - Position of a shaded sample



## <a href="#_description" class="anchor"></a>Description

`SamplePosition`  
Decorating a variable with the `SamplePosition` built-in decoration will
make that variable contain the sub-pixel position of the sample being
shaded. The top left of the pixel is considered to be at coordinate
(0,0) and the bottom right of the pixel is considered to be at
coordinate (1,1).

If the render pass has a fragment density map attachment, the variable
will instead contain the sub-fragment position of the sample being
shaded. The top left of the fragment is considered to be at coordinate
(0,0) and the bottom right of the fragment is considered to be at
coordinate (1,1) for any fragment area.

If a fragment shader entry point’s interface includes an input variable
decorated with `SamplePosition`, <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-sampleshading"
target="_blank" rel="noopener">Sample Shading</a> is considered enabled
with a `minSampleShading` value of 1.0.

If the current pipeline uses <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-samplelocations"
target="_blank" rel="noopener">custom sample locations</a> the value of
any variable decorated with the `SamplePosition` built-in decoration is
undefined.

Valid Usage

- <a href="#VUID-SamplePosition-SamplePosition-04360"
  id="VUID-SamplePosition-SamplePosition-04360"></a>
  VUID-SamplePosition-SamplePosition-04360  
  The `SamplePosition` decoration **must** be used only within the
  `Fragment` `Execution` `Model`

- <a href="#VUID-SamplePosition-SamplePosition-04361"
  id="VUID-SamplePosition-SamplePosition-04361"></a>
  VUID-SamplePosition-SamplePosition-04361  
  The variable decorated with `SamplePosition` **must** be declared
  using the `Input` `Storage` `Class`

- <a href="#VUID-SamplePosition-SamplePosition-04362"
  id="VUID-SamplePosition-SamplePosition-04362"></a>
  VUID-SamplePosition-SamplePosition-04362  
  The variable decorated with `SamplePosition` **must** be declared as a
  two-component vector of 32-bit floating-point values

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#SamplePosition"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
