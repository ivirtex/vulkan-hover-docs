# BaryCoordSmoothSampleAMD(3) Manual Page

## Name

BaryCoordSmoothSampleAMD - Barycentric coordinates of a sample center



## <a href="#_description" class="anchor"></a>Description

`BaryCoordSmoothSampleAMD`  
The `BaryCoordSmoothSampleAMD` decoration **can** be used to decorate a
fragment shader input variable. This variable will contain the (I,J)
pair of the barycentric coordinates corresponding to the fragment
evaluated using perspective interpolation at each covered sample. The K
coordinate of the barycentric coordinates **can** be derived given the
identity I + J + K = 1.0.

Valid Usage

- <a href="#VUID-BaryCoordSmoothSampleAMD-BaryCoordSmoothSampleAMD-04178"
  id="VUID-BaryCoordSmoothSampleAMD-BaryCoordSmoothSampleAMD-04178"></a>
  VUID-BaryCoordSmoothSampleAMD-BaryCoordSmoothSampleAMD-04178  
  The `BaryCoordSmoothSampleAMD` decoration **must** be used only within
  the `Fragment` `Execution` `Model`

- <a href="#VUID-BaryCoordSmoothSampleAMD-BaryCoordSmoothSampleAMD-04179"
  id="VUID-BaryCoordSmoothSampleAMD-BaryCoordSmoothSampleAMD-04179"></a>
  VUID-BaryCoordSmoothSampleAMD-BaryCoordSmoothSampleAMD-04179  
  The variable decorated with `BaryCoordSmoothSampleAMD` **must** be
  declared using the `Input` `Storage` `Class`

- <a href="#VUID-BaryCoordSmoothSampleAMD-BaryCoordSmoothSampleAMD-04180"
  id="VUID-BaryCoordSmoothSampleAMD-BaryCoordSmoothSampleAMD-04180"></a>
  VUID-BaryCoordSmoothSampleAMD-BaryCoordSmoothSampleAMD-04180  
  The variable decorated with `BaryCoordSmoothSampleAMD` **must** be
  declared as a two-component vector of 32-bit floating-point values

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#BaryCoordSmoothSampleAMD"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
