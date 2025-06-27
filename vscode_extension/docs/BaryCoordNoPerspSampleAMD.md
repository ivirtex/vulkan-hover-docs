# BaryCoordNoPerspSampleAMD(3) Manual Page

## Name

BaryCoordNoPerspSampleAMD - Barycentric coordinates of a sample center
in screen-space



## <a href="#_description" class="anchor"></a>Description

`BaryCoordNoPerspSampleAMD`  
The `BaryCoordNoPerspSampleAMD` decoration **can** be used to decorate a
fragment shader input variable. This variable will contain the (I,J)
pair of the barycentric coordinates corresponding to the fragment
evaluated using linear interpolation at each covered sample. The K
coordinate of the barycentric coordinates **can** be derived given the
identity I + J + K = 1.0.

Valid Usage

- <a
  href="#VUID-BaryCoordNoPerspSampleAMD-BaryCoordNoPerspSampleAMD-04166"
  id="VUID-BaryCoordNoPerspSampleAMD-BaryCoordNoPerspSampleAMD-04166"></a>
  VUID-BaryCoordNoPerspSampleAMD-BaryCoordNoPerspSampleAMD-04166  
  The `BaryCoordNoPerspSampleAMD` decoration **must** be used only
  within the `Fragment` `Execution` `Model`

- <a
  href="#VUID-BaryCoordNoPerspSampleAMD-BaryCoordNoPerspSampleAMD-04167"
  id="VUID-BaryCoordNoPerspSampleAMD-BaryCoordNoPerspSampleAMD-04167"></a>
  VUID-BaryCoordNoPerspSampleAMD-BaryCoordNoPerspSampleAMD-04167  
  The variable decorated with `BaryCoordNoPerspSampleAMD` **must** be
  declared using the `Input` `Storage` `Class`

- <a
  href="#VUID-BaryCoordNoPerspSampleAMD-BaryCoordNoPerspSampleAMD-04168"
  id="VUID-BaryCoordNoPerspSampleAMD-BaryCoordNoPerspSampleAMD-04168"></a>
  VUID-BaryCoordNoPerspSampleAMD-BaryCoordNoPerspSampleAMD-04168  
  The variable decorated with `BaryCoordNoPerspSampleAMD` **must** be
  declared as a two-component vector of 32-bit floating-point values

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#BaryCoordNoPerspSampleAMD"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
