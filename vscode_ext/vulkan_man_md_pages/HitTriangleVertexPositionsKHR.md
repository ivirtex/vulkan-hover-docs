# HitTriangleVertexPositionsKHR(3) Manual Page

## Name

HitTriangleVertexPositionsKHR - Vertices of an intersected triangle



## <a href="#_description" class="anchor"></a>Description

`HitTriangleVertexPositionsKHR`  
A variable decorated with the `HitTriangleVertexPositionsKHR` decoration
will specify the object space vertices of the triangle at the current
intersection in application-provided order. The positions returned are
transformed by the geometry transform, which is performed at standard <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fundamentals-floating-point"
target="_blank" rel="noopener">floating-point</a> precision, but without
a specifically defined order of floating-point operations to perform the
matrix multiplication.

Valid Usage

- <a
  href="#VUID-HitTriangleVertexPositionsKHR-HitTriangleVertexPositionsKHR-08747"
  id="VUID-HitTriangleVertexPositionsKHR-HitTriangleVertexPositionsKHR-08747"></a>
  VUID-HitTriangleVertexPositionsKHR-HitTriangleVertexPositionsKHR-08747  
  The `HitTriangleVertexPositionsKHR` decoration **must** be used only
  within the `AnyHitKHR` or `ClosestHitKHR` `Execution` `Model`

- <a
  href="#VUID-HitTriangleVertexPositionsKHR-HitTriangleVertexPositionsKHR-08748"
  id="VUID-HitTriangleVertexPositionsKHR-HitTriangleVertexPositionsKHR-08748"></a>
  VUID-HitTriangleVertexPositionsKHR-HitTriangleVertexPositionsKHR-08748  
  The variable decorated with `HitTriangleVertexPositionsKHR` **must**
  be declared using the `Input` `Storage` `Class`

- <a
  href="#VUID-HitTriangleVertexPositionsKHR-HitTriangleVertexPositionsKHR-08749"
  id="VUID-HitTriangleVertexPositionsKHR-HitTriangleVertexPositionsKHR-08749"></a>
  VUID-HitTriangleVertexPositionsKHR-HitTriangleVertexPositionsKHR-08749  
  The variable decorated with `HitTriangleVertexPositionsKHR` **must**
  be declared as an array of three vectors of three 32-bit float values

- <a
  href="#VUID-HitTriangleVertexPositionsKHR-HitTriangleVertexPositionsKHR-08750"
  id="VUID-HitTriangleVertexPositionsKHR-HitTriangleVertexPositionsKHR-08750"></a>
  VUID-HitTriangleVertexPositionsKHR-HitTriangleVertexPositionsKHR-08750  
  The variable decorated with `HitTriangleVertexPositionsKHR` **must**
  be used only if the value of `HitKindKHR` is
  `HitKindFrontFacingTriangleKHR` or `HitKindBackFacingTriangleKHR`

- <a href="#VUID-HitTriangleVertexPositionsKHR-None-08751"
  id="VUID-HitTriangleVertexPositionsKHR-None-08751"></a>
  VUID-HitTriangleVertexPositionsKHR-None-08751  
  The acceleration structure corresponding to the current intersection
  **must** have been built with
  `VK_BUILD_ACCELERATION_STRUCTURE_ALLOW_DATA_ACCESS_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#HitTriangleVertexPositionsKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
