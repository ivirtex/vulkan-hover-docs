# BaryCoordNoPerspCentroidAMD(3) Manual Page

## Name

BaryCoordNoPerspCentroidAMD - Barycentric coordinates of a fragment
centroid in screen-space



## <a href="#_description" class="anchor"></a>Description

`BaryCoordNoPerspCentroidAMD`  
The `BaryCoordNoPerspCentroidAMD` decoration **can** be used to decorate
a fragment shader input variable. This variable will contain the (I,J)
pair of the barycentric coordinates corresponding to the fragment
evaluated using linear interpolation at the centroid. The K coordinate
of the barycentric coordinates **can** be derived given the identity I +
J + K = 1.0.

Valid Usage

- <a
  href="#VUID-BaryCoordNoPerspCentroidAMD-BaryCoordNoPerspCentroidAMD-04163"
  id="VUID-BaryCoordNoPerspCentroidAMD-BaryCoordNoPerspCentroidAMD-04163"></a>
  VUID-BaryCoordNoPerspCentroidAMD-BaryCoordNoPerspCentroidAMD-04163  
  The `BaryCoordNoPerspCentroidAMD` decoration **must** be used only
  within the `Fragment` `Execution` `Model`

- <a
  href="#VUID-BaryCoordNoPerspCentroidAMD-BaryCoordNoPerspCentroidAMD-04164"
  id="VUID-BaryCoordNoPerspCentroidAMD-BaryCoordNoPerspCentroidAMD-04164"></a>
  VUID-BaryCoordNoPerspCentroidAMD-BaryCoordNoPerspCentroidAMD-04164  
  The variable decorated with `BaryCoordNoPerspCentroidAMD` **must** be
  declared using the `Input` `Storage` `Class`

- <a
  href="#VUID-BaryCoordNoPerspCentroidAMD-BaryCoordNoPerspCentroidAMD-04165"
  id="VUID-BaryCoordNoPerspCentroidAMD-BaryCoordNoPerspCentroidAMD-04165"></a>
  VUID-BaryCoordNoPerspCentroidAMD-BaryCoordNoPerspCentroidAMD-04165  
  The variable decorated with `BaryCoordNoPerspCentroidAMD` **must** be
  declared as a three-component vector of 32-bit floating-point values

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#BaryCoordNoPerspCentroidAMD"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
