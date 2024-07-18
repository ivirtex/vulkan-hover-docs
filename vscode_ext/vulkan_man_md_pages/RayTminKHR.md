# RayTminKHR(3) Manual Page

## Name

RayTminKHR - Minimum T value of a ray



## <a href="#_description" class="anchor"></a>Description

`RayTminKHR`  
A variable decorated with the `RayTminKHR` decoration will contain the
parametric t<sub>min</sub> value of the ray being processed. The value
is independent of the space in which the ray origin and direction exist.
The value is the parameter passed into the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#glossary-pipeline-trace-ray"
target="_blank" rel="noopener">pipeline trace ray</a> instruction.

The t<sub>min</sub> value remains constant for the duration of the ray
query.

Valid Usage

- <a href="#VUID-RayTminKHR-RayTminKHR-04351"
  id="VUID-RayTminKHR-RayTminKHR-04351"></a>
  VUID-RayTminKHR-RayTminKHR-04351  
  The `RayTminKHR` decoration **must** be used only within the
  `IntersectionKHR`, `AnyHitKHR`, `ClosestHitKHR`, or `MissKHR`
  `Execution` `Model`

- <a href="#VUID-RayTminKHR-RayTminKHR-04352"
  id="VUID-RayTminKHR-RayTminKHR-04352"></a>
  VUID-RayTminKHR-RayTminKHR-04352  
  The variable decorated with `RayTminKHR` **must** be declared using
  the `Input` `Storage` `Class`

- <a href="#VUID-RayTminKHR-RayTminKHR-04353"
  id="VUID-RayTminKHR-RayTminKHR-04353"></a>
  VUID-RayTminKHR-RayTminKHR-04353  
  The variable decorated with `RayTminKHR` **must** be declared as a
  scalar 32-bit floating-point value

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#RayTminKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
