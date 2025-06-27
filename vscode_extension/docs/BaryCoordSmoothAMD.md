# BaryCoordSmoothAMD(3) Manual Page

## Name

BaryCoordSmoothAMD - Barycentric coordinates of a fragment center



## <a href="#_description" class="anchor"></a>Description

`BaryCoordSmoothAMD`  
The `BaryCoordSmoothAMD` decoration **can** be used to decorate a
fragment shader input variable. This variable will contain the (I,J)
pair of the barycentric coordinates corresponding to the fragment
evaluated using perspective interpolation at the fragment’s center. The
K coordinate of the barycentric coordinates **can** be derived given the
identity I + J + K = 1.0.

Valid Usage

- <a href="#VUID-BaryCoordSmoothAMD-BaryCoordSmoothAMD-04172"
  id="VUID-BaryCoordSmoothAMD-BaryCoordSmoothAMD-04172"></a>
  VUID-BaryCoordSmoothAMD-BaryCoordSmoothAMD-04172  
  The `BaryCoordSmoothAMD` decoration **must** be used only within the
  `Fragment` `Execution` `Model`

- <a href="#VUID-BaryCoordSmoothAMD-BaryCoordSmoothAMD-04173"
  id="VUID-BaryCoordSmoothAMD-BaryCoordSmoothAMD-04173"></a>
  VUID-BaryCoordSmoothAMD-BaryCoordSmoothAMD-04173  
  The variable decorated with `BaryCoordSmoothAMD` **must** be declared
  using the `Input` `Storage` `Class`

- <a href="#VUID-BaryCoordSmoothAMD-BaryCoordSmoothAMD-04174"
  id="VUID-BaryCoordSmoothAMD-BaryCoordSmoothAMD-04174"></a>
  VUID-BaryCoordSmoothAMD-BaryCoordSmoothAMD-04174  
  The variable decorated with `BaryCoordSmoothAMD` **must** be declared
  as a two-component vector of 32-bit floating-point values

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#BaryCoordSmoothAMD"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
