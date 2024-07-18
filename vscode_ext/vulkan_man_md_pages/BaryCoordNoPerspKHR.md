# BaryCoordNoPerspKHR(3) Manual Page

## Name

BaryCoordNoPerspKHR - Barycentric coordinates of a fragment in
screen-space



## <a href="#_description" class="anchor"></a>Description

`BaryCoordNoPerspKHR`  
The `BaryCoordNoPerspKHR` decoration **can** be used to decorate a
fragment shader input variable. This variable will contain a
three-component floating-point vector with barycentric weights that
indicate the location of the fragment relative to the screen-space
locations of vertices of its primitive, obtained using linear
interpolation.

Valid Usage

- <a href="#VUID-BaryCoordNoPerspKHR-BaryCoordNoPerspKHR-04160"
  id="VUID-BaryCoordNoPerspKHR-BaryCoordNoPerspKHR-04160"></a>
  VUID-BaryCoordNoPerspKHR-BaryCoordNoPerspKHR-04160  
  The `BaryCoordNoPerspKHR` decoration **must** be used only within the
  `Fragment` `Execution` `Model`

- <a href="#VUID-BaryCoordNoPerspKHR-BaryCoordNoPerspKHR-04161"
  id="VUID-BaryCoordNoPerspKHR-BaryCoordNoPerspKHR-04161"></a>
  VUID-BaryCoordNoPerspKHR-BaryCoordNoPerspKHR-04161  
  The variable decorated with `BaryCoordNoPerspKHR` **must** be declared
  using the `Input` `Storage` `Class`

- <a href="#VUID-BaryCoordNoPerspKHR-BaryCoordNoPerspKHR-04162"
  id="VUID-BaryCoordNoPerspKHR-BaryCoordNoPerspKHR-04162"></a>
  VUID-BaryCoordNoPerspKHR-BaryCoordNoPerspKHR-04162  
  The variable decorated with `BaryCoordNoPerspKHR` **must** be declared
  as a three-component vector of 32-bit floating-point values

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#BaryCoordNoPerspKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
