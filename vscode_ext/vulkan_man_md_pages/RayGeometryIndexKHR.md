# RayGeometryIndexKHR(3) Manual Page

## Name

RayGeometryIndexKHR - Geometry index in a ray shader



## <a href="#_description" class="anchor"></a>Description

`RayGeometryIndexKHR`  
A variable decorated with the `RayGeometryIndexKHR` decoration will
contain the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#acceleration-structure-geometry-index"
target="_blank" rel="noopener">geometry index</a> for the acceleration
structure geometry currently being shaded.

Valid Usage

- <a href="#VUID-RayGeometryIndexKHR-RayGeometryIndexKHR-04345"
  id="VUID-RayGeometryIndexKHR-RayGeometryIndexKHR-04345"></a>
  VUID-RayGeometryIndexKHR-RayGeometryIndexKHR-04345  
  The `RayGeometryIndexKHR` decoration **must** be used only within the
  `IntersectionKHR`, `AnyHitKHR`, or `ClosestHitKHR` `Execution` `Model`

- <a href="#VUID-RayGeometryIndexKHR-RayGeometryIndexKHR-04346"
  id="VUID-RayGeometryIndexKHR-RayGeometryIndexKHR-04346"></a>
  VUID-RayGeometryIndexKHR-RayGeometryIndexKHR-04346  
  The variable decorated with `RayGeometryIndexKHR` **must** be declared
  using the `Input` `Storage` `Class`

- <a href="#VUID-RayGeometryIndexKHR-RayGeometryIndexKHR-04347"
  id="VUID-RayGeometryIndexKHR-RayGeometryIndexKHR-04347"></a>
  VUID-RayGeometryIndexKHR-RayGeometryIndexKHR-04347  
  The variable decorated with `RayGeometryIndexKHR` **must** be declared
  as a scalar 32-bit integer value

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#RayGeometryIndexKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
