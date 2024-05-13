# HitTNV(3) Manual Page

## Name

HitTNV - T value of a ray intersection



## <a href="#_description" class="anchor"></a>Description

`HitTNV`  
A variable decorated with the `HitTNV` decoration is equivalent to a
variable decorated with the `RayTmaxKHR` decoration.

Valid Usage

- <a href="#VUID-HitTNV-HitTNV-04245" id="VUID-HitTNV-HitTNV-04245"></a>
  VUID-HitTNV-HitTNV-04245  
  The `HitTNV` decoration **must** be used only within the `AnyHitNV` or
  `ClosestHitNV` `Execution` `Model`

- <a href="#VUID-HitTNV-HitTNV-04246" id="VUID-HitTNV-HitTNV-04246"></a>
  VUID-HitTNV-HitTNV-04246  
  The variable decorated with `HitTNV` **must** be declared using the
  `Input` `Storage` `Class`

- <a href="#VUID-HitTNV-HitTNV-04247" id="VUID-HitTNV-HitTNV-04247"></a>
  VUID-HitTNV-HitTNV-04247  
  The variable decorated with `HitTNV` **must** be declared as a scalar
  32-bit floating-point value

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#HitTNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
