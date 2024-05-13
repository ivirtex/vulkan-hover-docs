# BaryCoordKHR(3) Manual Page

## Name

BaryCoordKHR - Barycentric coordinates of a fragment



## <a href="#_description" class="anchor"></a>Description

`BaryCoordKHR`  
The `BaryCoordKHR` decoration **can** be used to decorate a fragment
shader input variable. This variable will contain a three-component
floating-point vector with barycentric weights that indicate the
location of the fragment relative to the screen-space locations of
vertices of its primitive, obtained using perspective interpolation.

Valid Usage

- <a href="#VUID-BaryCoordKHR-BaryCoordKHR-04154"
  id="VUID-BaryCoordKHR-BaryCoordKHR-04154"></a>
  VUID-BaryCoordKHR-BaryCoordKHR-04154  
  The `BaryCoordKHR` decoration **must** be used only within the
  `Fragment` `Execution` `Model`

- <a href="#VUID-BaryCoordKHR-BaryCoordKHR-04155"
  id="VUID-BaryCoordKHR-BaryCoordKHR-04155"></a>
  VUID-BaryCoordKHR-BaryCoordKHR-04155  
  The variable decorated with `BaryCoordKHR` **must** be declared using
  the `Input` `Storage` `Class`

- <a href="#VUID-BaryCoordKHR-BaryCoordKHR-04156"
  id="VUID-BaryCoordKHR-BaryCoordKHR-04156"></a>
  VUID-BaryCoordKHR-BaryCoordKHR-04156  
  The variable decorated with `BaryCoordKHR` **must** be declared as a
  three-component vector of 32-bit floating-point values

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#BaryCoordKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
