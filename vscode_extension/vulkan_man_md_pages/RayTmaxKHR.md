# RayTmaxKHR(3) Manual Page

## Name

RayTmaxKHR - Maximum T value of a ray



## <a href="#_description" class="anchor"></a>Description

`RayTmaxKHR`  
A variable decorated with the `RayTmaxKHR` decoration will contain the
parametric t<sub>max</sub> value of the ray being processed. The value
is independent of the space in which the ray origin and direction exist.
The value is initialized to the parameter passed into the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#glossary-pipeline-trace-ray"
target="_blank" rel="noopener">pipeline trace ray</a> instruction.

The t<sub>max</sub> value changes throughout the lifetime of the ray
that produced the intersection. In the closest hit shader, the value
reflects the closest distance to the intersected primitive. In the
any-hit shader, it reflects the distance to the primitive currently
being intersected. In the intersection shader, it reflects the distance
to the closest primitive intersected so far or the initial value. The
value can change in the intersection shader after calling
`OpReportIntersectionKHR` if the corresponding any-hit shader does not
ignore the intersection. In a miss shader, the value is identical to the
parameter passed into the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#glossary-pipeline-trace-ray"
target="_blank" rel="noopener">pipeline trace ray</a> instruction.

Valid Usage

- <a href="#VUID-RayTmaxKHR-RayTmaxKHR-04348"
  id="VUID-RayTmaxKHR-RayTmaxKHR-04348"></a>
  VUID-RayTmaxKHR-RayTmaxKHR-04348  
  The `RayTmaxKHR` decoration **must** be used only within the
  `IntersectionKHR`, `AnyHitKHR`, `ClosestHitKHR`, or `MissKHR`
  `Execution` `Model`

- <a href="#VUID-RayTmaxKHR-RayTmaxKHR-04349"
  id="VUID-RayTmaxKHR-RayTmaxKHR-04349"></a>
  VUID-RayTmaxKHR-RayTmaxKHR-04349  
  The variable decorated with `RayTmaxKHR` **must** be declared using
  the `Input` `Storage` `Class`

- <a href="#VUID-RayTmaxKHR-RayTmaxKHR-04350"
  id="VUID-RayTmaxKHR-RayTmaxKHR-04350"></a>
  VUID-RayTmaxKHR-RayTmaxKHR-04350  
  The variable decorated with `RayTmaxKHR` **must** be declared as a
  scalar 32-bit floating-point value

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#RayTmaxKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
