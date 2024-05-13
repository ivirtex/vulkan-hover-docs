# HitKindKHR(3) Manual Page

## Name

HitKindKHR - Kind of hit that triggered an any-hit or closest hit ray
shader



## <a href="#_description" class="anchor"></a>Description

`HitKindKHR`  
A variable decorated with the `HitKindKHR` decoration will describe the
intersection that triggered the execution of the current shader. The
values are determined by the intersection shader. For user-defined
intersection shaders this is the value that was passed to the “Hit Kind”
operand of `OpReportIntersectionKHR`. For triangle intersection
candidates, this will be one of `HitKindFrontFacingTriangleKHR` or
`HitKindBackFacingTriangleKHR`.

Valid Usage

- <a href="#VUID-HitKindKHR-HitKindKHR-04242"
  id="VUID-HitKindKHR-HitKindKHR-04242"></a>
  VUID-HitKindKHR-HitKindKHR-04242  
  The `HitKindKHR` decoration **must** be used only within the
  `AnyHitKHR` or `ClosestHitKHR` `Execution` `Model`

- <a href="#VUID-HitKindKHR-HitKindKHR-04243"
  id="VUID-HitKindKHR-HitKindKHR-04243"></a>
  VUID-HitKindKHR-HitKindKHR-04243  
  The variable decorated with `HitKindKHR` **must** be declared using
  the `Input` `Storage` `Class`

- <a href="#VUID-HitKindKHR-HitKindKHR-04244"
  id="VUID-HitKindKHR-HitKindKHR-04244"></a>
  VUID-HitKindKHR-HitKindKHR-04244  
  The variable decorated with `HitKindKHR` **must** be declared as a
  scalar 32-bit integer value

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#HitKindKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
