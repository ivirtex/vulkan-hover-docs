# WorldRayOriginKHR(3) Manual Page

## Name

WorldRayOriginKHR - Ray origin in world space



## <a href="#_description" class="anchor"></a>Description

`WorldRayOriginKHR`  
A variable decorated with the `WorldRayOriginKHR` decoration will
specify the origin of the ray being processed, in world space. The value
is the parameter passed into the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#glossary-pipeline-trace-ray"
target="_blank" rel="noopener">pipeline trace ray</a> instruction.

Valid Usage

- <a href="#VUID-WorldRayOriginKHR-WorldRayOriginKHR-04431"
  id="VUID-WorldRayOriginKHR-WorldRayOriginKHR-04431"></a>
  VUID-WorldRayOriginKHR-WorldRayOriginKHR-04431  
  The `WorldRayOriginKHR` decoration **must** be used only within the
  `IntersectionKHR`, `AnyHitKHR`, `ClosestHitKHR`, or `MissKHR`
  `Execution` `Model`

- <a href="#VUID-WorldRayOriginKHR-WorldRayOriginKHR-04432"
  id="VUID-WorldRayOriginKHR-WorldRayOriginKHR-04432"></a>
  VUID-WorldRayOriginKHR-WorldRayOriginKHR-04432  
  The variable decorated with `WorldRayOriginKHR` **must** be declared
  using the `Input` `Storage` `Class`

- <a href="#VUID-WorldRayOriginKHR-WorldRayOriginKHR-04433"
  id="VUID-WorldRayOriginKHR-WorldRayOriginKHR-04433"></a>
  VUID-WorldRayOriginKHR-WorldRayOriginKHR-04433  
  The variable decorated with `WorldRayOriginKHR` **must** be declared
  as a three-component vector of 32-bit floating-point values

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#WorldRayOriginKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
