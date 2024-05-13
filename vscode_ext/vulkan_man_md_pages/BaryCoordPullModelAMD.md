# BaryCoordPullModelAMD(3) Manual Page

## Name

BaryCoordPullModelAMD - Inverse barycentric coordinates of a fragment
center



## <a href="#_description" class="anchor"></a>Description

`BaryCoordPullModelAMD`  
The `BaryCoordPullModelAMD` decoration **can** be used to decorate a
fragment shader input variable. This variable will contain (1/W, 1/I,
1/J) evaluated at the fragment center and **can** be used to calculate
gradients and then interpolate I, J, and W at any desired sample
location.

Valid Usage

- <a href="#VUID-BaryCoordPullModelAMD-BaryCoordPullModelAMD-04169"
  id="VUID-BaryCoordPullModelAMD-BaryCoordPullModelAMD-04169"></a>
  VUID-BaryCoordPullModelAMD-BaryCoordPullModelAMD-04169  
  The `BaryCoordPullModelAMD` decoration **must** be used only within
  the `Fragment` `Execution` `Model`

- <a href="#VUID-BaryCoordPullModelAMD-BaryCoordPullModelAMD-04170"
  id="VUID-BaryCoordPullModelAMD-BaryCoordPullModelAMD-04170"></a>
  VUID-BaryCoordPullModelAMD-BaryCoordPullModelAMD-04170  
  The variable decorated with `BaryCoordPullModelAMD` **must** be
  declared using the `Input` `Storage` `Class`

- <a href="#VUID-BaryCoordPullModelAMD-BaryCoordPullModelAMD-04171"
  id="VUID-BaryCoordPullModelAMD-BaryCoordPullModelAMD-04171"></a>
  VUID-BaryCoordPullModelAMD-BaryCoordPullModelAMD-04171  
  The variable decorated with `BaryCoordPullModelAMD` **must** be
  declared as a three-component vector of 32-bit floating-point values

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#BaryCoordPullModelAMD"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
