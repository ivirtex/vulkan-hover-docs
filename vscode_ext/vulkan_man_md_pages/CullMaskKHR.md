# CullMaskKHR(3) Manual Page

## Name

CullMaskKHR - OpTrace specified ray cull mask



## <a href="#_description" class="anchor"></a>Description

`CullMaskKHR`  
A variable decorated with the `CullMaskKHR` decoration will specify the
cull mask of the ray being processed. The value is given by the
`Cull Mask` parameter passed into one of the `OpTrace*` instructions.

Valid Usage

- <a href="#VUID-CullMaskKHR-CullMaskKHR-06735"
  id="VUID-CullMaskKHR-CullMaskKHR-06735"></a>
  VUID-CullMaskKHR-CullMaskKHR-06735  
  The `CullMaskKHR` decoration **must** be used only within the
  `IntersectionKHR`, `AnyHitKHR`, `ClosestHitKHR`, or `MissKHR`
  `Execution` `Model`

- <a href="#VUID-CullMaskKHR-CullMaskKHR-06736"
  id="VUID-CullMaskKHR-CullMaskKHR-06736"></a>
  VUID-CullMaskKHR-CullMaskKHR-06736  
  The variable decorated with `CullMaskKHR` **must** be declared using
  the `Input` `Storage` `Class`

- <a href="#VUID-CullMaskKHR-CullMaskKHR-06737"
  id="VUID-CullMaskKHR-CullMaskKHR-06737"></a>
  VUID-CullMaskKHR-CullMaskKHR-06737  
  The variable decorated with `CullMaskKHR` **must** be declared as a
  scalar 32-bit integer value

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#CullMaskKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
