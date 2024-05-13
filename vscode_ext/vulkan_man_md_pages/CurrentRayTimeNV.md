# CurrentRayTimeNV(3) Manual Page

## Name

CurrentRayTimeNV - Time value of a ray intersection



## <a href="#_description" class="anchor"></a>Description

`CurrentRayTimeNV`  
A variable decorated with the `CurrentRayTimeNV` decoration contains the
time value passed in to `OpTraceRayMotionNV` which called this shader.

Valid Usage

- <a href="#VUID-CurrentRayTimeNV-CurrentRayTimeNV-04942"
  id="VUID-CurrentRayTimeNV-CurrentRayTimeNV-04942"></a>
  VUID-CurrentRayTimeNV-CurrentRayTimeNV-04942  
  The `CurrentRayTimeNV` decoration **must** be used only within the
  `IntersectionKHR`, `AnyHitKHR`, `ClosestHitKHR`, or `MissKHR`
  `Execution` `Model`

- <a href="#VUID-CurrentRayTimeNV-CurrentRayTimeNV-04943"
  id="VUID-CurrentRayTimeNV-CurrentRayTimeNV-04943"></a>
  VUID-CurrentRayTimeNV-CurrentRayTimeNV-04943  
  The variable decorated with `CurrentRayTimeNV` **must** be declared
  using the `Input` `Storage` `Class`

- <a href="#VUID-CurrentRayTimeNV-CurrentRayTimeNV-04944"
  id="VUID-CurrentRayTimeNV-CurrentRayTimeNV-04944"></a>
  VUID-CurrentRayTimeNV-CurrentRayTimeNV-04944  
  The variable decorated with `CurrentRayTimeNV` **must** be declared as
  a scalar 32-bit floating-point value

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#CurrentRayTimeNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
