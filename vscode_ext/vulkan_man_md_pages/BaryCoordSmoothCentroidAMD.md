# BaryCoordSmoothCentroidAMD(3) Manual Page

## Name

BaryCoordSmoothCentroidAMD - Barycentric coordinates of a fragment
centroid



## <a href="#_description" class="anchor"></a>Description

`BaryCoordSmoothCentroidAMD`  
The `BaryCoordSmoothCentroidAMD` decoration **can** be used to decorate
a fragment shader input variable. This variable will contain the (I,J)
pair of the barycentric coordinates corresponding to the fragment
evaluated using perspective interpolation at the centroid. The K
coordinate of the barycentric coordinates **can** be derived given the
identity I + J + K = 1.0.

Valid Usage

- <a
  href="#VUID-BaryCoordSmoothCentroidAMD-BaryCoordSmoothCentroidAMD-04175"
  id="VUID-BaryCoordSmoothCentroidAMD-BaryCoordSmoothCentroidAMD-04175"></a>
  VUID-BaryCoordSmoothCentroidAMD-BaryCoordSmoothCentroidAMD-04175  
  The `BaryCoordSmoothCentroidAMD` decoration **must** be used only
  within the `Fragment` `Execution` `Model`

- <a
  href="#VUID-BaryCoordSmoothCentroidAMD-BaryCoordSmoothCentroidAMD-04176"
  id="VUID-BaryCoordSmoothCentroidAMD-BaryCoordSmoothCentroidAMD-04176"></a>
  VUID-BaryCoordSmoothCentroidAMD-BaryCoordSmoothCentroidAMD-04176  
  The variable decorated with `BaryCoordSmoothCentroidAMD` **must** be
  declared using the `Input` `Storage` `Class`

- <a
  href="#VUID-BaryCoordSmoothCentroidAMD-BaryCoordSmoothCentroidAMD-04177"
  id="VUID-BaryCoordSmoothCentroidAMD-BaryCoordSmoothCentroidAMD-04177"></a>
  VUID-BaryCoordSmoothCentroidAMD-BaryCoordSmoothCentroidAMD-04177  
  The variable decorated with `BaryCoordSmoothCentroidAMD` **must** be
  declared as a two-component vector of 32-bit floating-point values

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#BaryCoordSmoothCentroidAMD"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
